//version 1.16++
package main

import (
	"cs2/router"
	"cs2/utils"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
)

//go:embed frontend/dist/*
var Fs embed.FS

func main() {
	//生成端口
	port := utils.GeneratorPort()
	fmt.Println("端口：", port)
	// 结构化目录
	staticFiles, _ := fs.Sub(Fs, "frontend/dist")
	router := router.InitRouter(staticFiles)
	go router.Run(":" + port)
	//chan
	chChromeDie := make(chan struct{})
	chBackendDie := make(chan struct{})
	//启动chrome
	go utils.Chrome_start(chChromeDie, chBackendDie, port)

	chSignal := signalInterrupt()
	//case只执行一次所以要死循环
	for {
		select {
		//Ctrl+C时关闭
		case <-chSignal:
			fmt.Println("监听到Ctrl+c")
			chBackendDie <- struct{}{}
		//浏览器关闭后关闭主程序
		case <-chChromeDie:
			os.Exit(0)
		}
	}

}

//检测主程序关闭
func signalInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
