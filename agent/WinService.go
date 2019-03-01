package agent

import (
	"errors"
	log "github.com/Deansquirrel/goToolLog"
	"os/exec"
	"strings"
)

type winService struct {
	Name string
}

func NewWinService(name string) *winService {
	return &winService{
		Name: name,
	}
}

//重启Windows服务
func (ws *winService) Restart() error {
	b, err := ws.IsRunning()
	if err != nil {
		return err
	}
	if b {
		err = ws.Stop()
		if err != nil {
			return err
		}
	}
	err = ws.Start()
	if err != nil {
		return err
	}
	return nil
}

//监测Windows服务是否在运行
func (ws *winService) IsRunning() (bool, error) {
	cmd := exec.Command("sc", "query", ws.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if err.Error() == "exit status 1060" {
			return false, errors.New("指定的服务[" + ws.Name + "]未安装")
		}
		return false, err
	}
	if strings.Index(string(out), "STOPPED") > 0 {
		return false, nil
	}
	if strings.Index(string(out), "RUNNING") > 0 {
		return true, nil
	}
	return false, errors.New("状态检查失败")
}

//停止Windows服务
func (ws *winService) Stop() error {
	log.Info("stop service " + ws.Name)
	cmd := exec.Command("net", "stop", ws.Name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

//启动Windows服务
func (ws *winService) Start() error {
	log.Info("start service " + ws.Name)
	cmd := exec.Command("net", "start", ws.Name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
