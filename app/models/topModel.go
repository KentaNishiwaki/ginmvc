package models

import (
	"fmt"
	"os/exec"
	"time"
)

type TopPageData struct {
	User        *LoginUser
	Title       string
	Description string
	Error       bool
	DevMode     bool
}

func GetAll(user *LoginUser) (datas TopPageData) {
	datas.User = user
	datas.Title = "Raspberry Pi App"
	datas.Error = false
	datas.DevMode = false
	return
}

func GetNoUser() (datas TopPageData) {
	datas.User = &LoginUser{}
	datas.Title = "App Login"
	datas.Error = false
	datas.DevMode = false
	return
}

func ExecShutdown() (rc string, errMes string) {
	time.Sleep(1000 * time.Millisecond)

	cmd := exec.Command("sh/shutdown.sh")
	cmd.Output()
	result, err := cmd.Output()

	rc = string(result)
	errMes = err.Error()
	fmt.Printf("ExecShutdown mes -> " + rc + "\n")
	fmt.Printf("ExecShutdown err -> " + errMes + "\n")
	return
}

func ExecReboot() (rc string, errMes string) {
	time.Sleep(1000 * time.Millisecond)

	cmd := exec.Command("sh/reboot.sh")
	result, err := cmd.Output()

	rc = string(result)
	errMes = err.Error()
	fmt.Printf("ExecReboot mes -> " + rc + "\n")
	fmt.Printf("ExecReboot err -> " + errMes + "\n")
	return
}
