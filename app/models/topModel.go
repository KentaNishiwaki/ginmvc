package models

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
	"unsafe"
)

type TopPageData struct {
	User              *LoginUser
	Title             string
	Description       string
	Error             bool
	DevMode           bool
	MoreStyles        []string
	MoreScripts       []string
	MoreModule        []string
	FooterScripts     []string
	MicrowaveJson     string
	HighFrequencyJson string
	JWT_TOKEN         string
}

func GetAll(user *LoginUser, config *Config) (formdata TopPageData) {
	formdata.User = user
	formdata.Title = "Gin App"
	formdata.Error = false
	formdata.DevMode = false
	formdata.MoreStyles = []string{}
	formdata.MoreScripts = []string{"chart.js", "chartjs-adapter-date-fns.bundle.min.js", "jquery.cookie.js"}
	formdata.MoreModule = []string{}
	formdata.FooterScripts = []string{"topPage.js?time=" + time.Now().Format("2006/1/2 15:04:05")}
	return
}

func GetMicroWaveJSON(config *Config, selDate string) *string {
	microWave := KvMicrowave{}
	rc, err := microWave.GetKvMicrowave(config, selDate)
	if err != nil {
		empty := ""
		return &empty
	}
	strData, _ := json.Marshal(rc)
	return (*string)(unsafe.Pointer(&strData))
}
func GetHighFrequencyJSON(config *Config, selDate string) *string {
	highFrequency := KvHighFrequency{}
	rc, err := highFrequency.GetKvHighFrequency(config, selDate)
	if err != nil {
		empty := ""
		return &empty
	}
	strData, _ := json.Marshal(rc)
	return (*string)(unsafe.Pointer(&strData))
}
func GetNoUser(config *Config, user *LoginUser) (formdata TopPageData) {
	formdata.User = user
	formdata.Title = "App Login"
	formdata.Error = false
	formdata.DevMode = false
	return
}

func ExecShutdown(config *Config) (rc string, errMes string) {
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

func ExecReboot(config *Config) (rc string, errMes string) {
	time.Sleep(1000 * time.Millisecond)

	cmd := exec.Command("sh/reboot.sh")
	result, err := cmd.Output()

	rc = string(result)
	errMes = err.Error()
	fmt.Printf("ExecReboot mes -> " + rc + "\n")
	fmt.Printf("ExecReboot err -> " + errMes + "\n")
	return
}
