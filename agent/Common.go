package agent

import (
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris/core/errors"
	"os/exec"
	"strings"
)

func RestartWinService(serviceName string) error {
	log.Debug("ServiceName:" + serviceName)
	b, err := IsServiceRunning(serviceName)
	if err != nil {
		return err
	}
	if b {
		err = StopService(serviceName)
		if err != nil {
			return err
		}
	}
	err = StartService(serviceName)
	if err != nil {
		return err
	}
	return nil
}

func IsServiceRunning(serviceName string) (bool, error) {
	cmd := exec.Command("sc", "query", serviceName)
	out, err := cmd.CombinedOutput()
	if err != nil {
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

func StopService(serviceName string) error {
	log.Info("stop service " + serviceName)
	cmd := exec.Command("net", "stop", serviceName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func StartService(serviceName string) error {
	log.Info("start service " + serviceName)
	cmd := exec.Command("net", "start", serviceName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
