package del

//
//import (
//	"DigSakura/dig/point"
//	"math/big"
//	"time"
//)
//
//// 默认工作队列数量
//var workerNum = 3
//
//// 正在工作中的队列数量
//var workeringNum = 0
//
//var workers = make([]chan *Cell, 0)
//
//// 开启 cell 处理线程
//func StartCellWorker(workNum int, endCb func()) {
//	if workNum > 0 {
//		workerNum = workNum
//	}
//
//	workeringNum = workerNum
//	println("开启处理管线", workerNum)
//	for i := 0; i < workerNum; i++ {
//		// cell 管道大小
//		ch := make(chan *Cell, 2000000)
//		workers = append(workers, ch)
//		go func() {
//			startTime := time.Now()
//			procesPack(ch)
//			point.DtTime(startTime, "cell 通道关闭,当前管道处理时长")
//			workeringNum--
//			if workeringNum == 0 {
//				endCb()
//			}
//		}()
//	}
//
//}
//
//// 处理数据包线程
//func procesPack(Ch chan *Cell) {
//
//	// 制作好大数据对象反复在当前计算线程中使用
//	var add = big.NewInt(0)
//	var b = big.NewInt(0)
//	var a = big.NewInt(0)
//	var const1024 = big.NewInt(1024)
//	for {
//		cell, ok := <-Ch
//		if ok {
//			cell.Process(a, b, add, const1024)
//			// 处理完成后减少引用计数
//			cell.Pool.RefReduce()
//		} else {
//			// 通道关闭
//			return
//		}
//	}
//
//}
//
//var sele = 0
//
//func SubmitCellsToChan(data *[]*Cell) {
//	for _, cell := range *data {
//		cell.Pool.RefPlus()
//		// 数据平均分配到每条管道中
//		workers[sele%workerNum] <- cell
//		sele++
//	}
//}
//
//func SubmitCellToChan(cell *Cell) {
//	workers[sele%workerNum] <- cell
//	sele++
//}
//
//// 关闭所有管道
//func CloseProcess() {
//	for _, worker := range workers {
//		close(worker)
//	}
//}
