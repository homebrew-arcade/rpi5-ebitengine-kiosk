package kiosk

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

type CmdConfig struct {
	Cmd     *exec.Cmd
	AppPath string
}

type RunResult struct {
	Code    int
	Message string
}

func Main() error {
	appPath := os.Getenv("KIOSK_APP_PATH")
	if appPath == "" {
		return fmt.Errorf("KIOSK_APP_PATH not set in ENV")
	}
	appWatch := false
	if os.Getenv("KIOSK_APP_WATCH") == "1" {
		appWatch = true
	}

	cmdConf := &CmdConfig{
		AppPath: appPath,
	}
	appStat, statErr := os.Stat(appPath)
	if statErr != nil {
		return fmt.Errorf("KIOSK_APP_PATH inaccessible from stat")
	}
	sigsChan := make(chan os.Signal, 1)
	signal.Notify(sigsChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	resultChan := make(chan RunResult, 1)
	ticker := time.NewTicker(time.Second)
	isHotReloading := false
	if !appWatch {
		ticker.Stop()
	}

	RunProc(cmdConf, resultChan)
	for {
		select {
		case <-sigsChan:
			fmt.Println("Interupt signal receieved")
			CloseProc(cmdConf)
			os.Exit(0)
		case res := <-resultChan:
			fmt.Printf("Process completed. Code: %v Message: %v \n", res.Code, res.Message)
			if !isHotReloading {
				fmt.Println("Exiting without hotreload")
				os.Exit(0)
			}
			isHotReloading = false
		case <-ticker.C:
			stat, statErr := os.Stat(appPath)
			if statErr != nil {
				return fmt.Errorf("KIOSK_APP_PATH inaccessible from stat")
			}
			if stat.Size() != appStat.Size() || stat.ModTime() != appStat.ModTime() {
				appStat = stat
				fmt.Println("Hot reloading app")
				isHotReloading = true
				RunProc(cmdConf, resultChan)
			}
		}
	}
}

func CloseProc(cmdConf *CmdConfig) {
	if cmdConf.Cmd != nil && cmdConf.Cmd.Process != nil {
		fmt.Println("CloseProc via signal")
		err := cmdConf.Cmd.Process.Signal(syscall.SIGINT)
		if err != nil {
			fmt.Println("CloseProc via kill")
			cmdConf.Cmd.Process.Kill()
		}
		cmdConf.Cmd = nil
	}
}

func RunProc(cmdConf *CmdConfig, resultChan chan RunResult) {
	CloseProc(cmdConf)
	cmdConf.Cmd = exec.Command("cage", cmdConf.AppPath)
	err := cmdConf.Cmd.Start()
	if err != nil {
		resultChan <- RunResult{
			Code:    1,
			Message: err.Error(),
		}
		return
	}
	go func(resultChan chan RunResult) {
		result := RunResult{}
		if err := cmdConf.Cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				fmt.Println("Exit Status", exiterr.ExitCode())
				result.Code = exiterr.ExitCode()
			} else {
				result.Code = 1
				result.Message = err.Error()
			}
		}
		resultChan <- result
	}(resultChan)
}
