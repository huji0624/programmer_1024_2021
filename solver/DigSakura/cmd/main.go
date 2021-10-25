package main

import (
	"DigSakura/dig/cell"
	"DigSakura/dig/find_1024"
	"DigSakura/dig/point"
	"DigSakura/dig/submit"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//var stub = make([]byte, 1024*1024*1024)
	//stub[0] = 1
	c := make(chan os.Signal, 1)

	startTime := time.Now()

	// 开启8条提交线程
	submit.StartSubmitWorker(64, func() {
		println("网络关闭")
		closeCallback()
		//// 全部提交后，退出应用程序
		//c <- syscall.SIGQUIT
	})

	// 开启16条 脱壳 处理管道
	cell.StartByteWorker(128, func() {
		println("文件读取完成")
	}, func() {
		// 所有文件处理完成
		// 查找值后，优化1024算式
		find_1024.Tactics_dataAllIn(true)
		point.DtTime(startTime, "总工作时长")
		submit.Close()
	})

	// 开启1024公式查找
	find_1024.StartFind1024()

	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		//g.Log().Info("get a signal :", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			//closeCallback()
			println("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}

// 关闭回调
func closeCallback() {
	// 打印处理报告
	point.Report()
}
