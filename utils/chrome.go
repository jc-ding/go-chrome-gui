package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Chrome_start(chChromeDie, chBackendDie chan struct{}, port string) {
	args := []string{
		"--window-size=600,600",
		//"--no-first-run",//没有首页时显示空白页
		//"--remote-debugging-port=0",
	}
	url := fmt.Sprintf("http://localhost:%s/index.html", port)
	args = append(args, "--app="+url)
	args = append(args, "--user-data-dir=./"+port)
	c := exec.Command(LocateChrome(), args...)
	c.Start()
	//后端关闭时关闭浏览器
	go func() {
		<-chBackendDie
		fmt.Println("检测到后端已关闭")
		c.Process.Kill()
	}()
	//浏览器关闭时传出chan
	go func() {
		c.Wait()
		fmt.Println("检测到浏览器已关闭")
		chChromeDie <- struct{}{}
	}()
}

func LocateChrome() string {

	// If env variable "LORCACHROME" specified and it exists
	if path, ok := os.LookupEnv("LORCACHROME"); ok {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
		}
	case "windows":
		paths = []string{
			os.Getenv("LocalAppData") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("LocalAppData") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe",
		}
	default:
		paths = []string{
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
			"/snap/bin/chromium",
		}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}
