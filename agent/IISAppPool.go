package agent

import (
	"fmt"
	"github.com/Deansquirrel/goAgentWin/global"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris/core/errors"
	"os/exec"
	"strings"
	"time"
)

type iisAppPool struct {
	Name string
}

func NewIISAppPool(name string) *iisAppPool {
	return &iisAppPool{
		Name: strings.Trim(name, " "),
	}
}

//重启IIS应用程序池
func (ap *iisAppPool) Restart() error {
	exist, err := ap.isExist()
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("AppPool is not exist")
	}
	b, err := ap.IsRunning()
	if err != nil {
		return err
	}
	if b {
		err = ap.Stop()
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 5)
	}
	outTime := time.Now().Add(time.Duration(global.SysConfig.AppPool.StartTimeout * 1000 * 1000 * 1000))
	for {
		err = ap.Start()
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(global.SysConfig.AppPool.StartDelay * 1000 * 1000 * 1000))
		check, err := ap.IsRunning()
		if err != nil {
			return err
		}
		if check {
			return nil
		}
		if time.Now().After(outTime) {
			return errors.New("timeout")
		}
	}
}

//检测IIS应用程序池是否存在
func (ap *iisAppPool) isExist() (bool, error) {
	cmd := exec.Command(global.SysConfig.AppPool.Path, "list", "appPools")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("检测IIS应用程序池是否存在时遇到错误：" + err.Error())
		return false, err
	}
	if strings.Index(strings.ToLower(string(out)), strings.ToLower(ap.Name)) > 0 {
		return true, nil
	}
	return false, nil
}

//检测IIS应用程序池是否在运行
func (ap *iisAppPool) IsRunning() (bool, error) {
	exist, err := ap.isExist()
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errors.New("AppPool is not exist")
	}
	cmd := exec.Command(global.SysConfig.AppPool.Path, "list", "appPools", "/state:started")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("检测IIS应用程序池是否在运行时遇到错误：" + err.Error())
		return false, err
	}
	outStr := string(out)
	outStr = strings.ToLower(outStr)
	if strings.Index(outStr, strings.ToLower(ap.Name)) > 0 {
		return true, nil
	}
	return false, nil
}

//启动IIS应用程序池
func (ap *iisAppPool) Start() error {
	exist, err := ap.isExist()
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("AppPool is not exist")
	}
	running, err := ap.IsRunning()
	if err != nil {
		return err
	}
	if running {
		return nil
	}
	cmd := exec.Command(global.SysConfig.AppPool.Path, "start", "appPool", fmt.Sprintf("/appPool.name:%s", ap.Name))
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Error("启动应用程序池时遇到错误：" + err.Error())
		return err
	}
	return nil
}

//停止IIS应用程序池
func (ap *iisAppPool) Stop() error {
	exist, err := ap.isExist()
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("AppPool is not exist")
	}
	running, err := ap.IsRunning()
	if err != nil {
		return err
	}
	if !running {
		return nil
	}
	cmd := exec.Command(global.SysConfig.AppPool.Path, "stop", "appPool", fmt.Sprintf("/appPool.name:%s", ap.Name))
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Error("停止应用程序池时遇到错误：" + err.Error())
		return err
	}
	return nil
}
