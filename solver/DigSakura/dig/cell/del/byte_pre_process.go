package del

//
//import (
//	"DigSakura/dig/point"
//	"time"
//)
//
//func Readdd(poll *BuffPool, content *[]byte, bufflen int, cb func(data *[]*Cell)) {
//	buff := *content
//	startTime := time.Now()
//	// 去新的携程中处理二进制流
//	datas := make([]*Cell, 2000010)
//	datasIndex := 0
//
//	data := &Cell{Pool: poll}
//
//	startLocal := 0
//	startMagic := 0
//	isLocalRead := false
//	origin := make([]byte, 64)
//	originLen := 0
//
//	for i := 0; i < bufflen; {
//		elem := buff[i]
//		// 开始符号
//		//if elem == byte('\n') {
//		//}
//
//		// 头部
//		if i+3 < bufflen && buff[i] == '{' && buff[i+1] == '"' && buff[i+2] == 'l' && buff[i+3] == 'o' {
//			//跳过 {"locationid":"
//			i += 15
//
//			// 记录 local 起点
//			startLocal = i
//			isLocalRead = true
//			// 分配内存
//			origin = make([]byte, 64)
//			// 索引归零
//			originLen = 0
//
//			//开始记录 Locationid
//			continue
//		}
//		// 中部
//		if i+3 < bufflen && buff[i] == '"' && buff[i+1] == ',' && buff[i+2] == '"' && buff[i+3] == 'm' {
//			//println(string(buff[startLocal:endLocal]))
//			// local 读取完成，写入切片
//			data.Locationid = buff[startLocal:i]
//			// 写入处理的数据
//			data.Origin = origin[0:originLen]
//			isLocalRead = false
//
//			//跳过 ","magic":"
//			i += 11
//
//			startMagic = i
//			continue
//		}
//
//		// 尾部
//		if i+2 < bufflen && buff[i] == '"' && buff[i+1] == '}' {
//			// Magic 读取完成，写入切片
//			data.Magic = buff[startMagic:i]
//
//			//跳过 "}
//			i += 2
//
//			datas[datasIndex] = data
//			datasIndex++
//			//结束记录 创建信息的Cell
//			data = &Cell{Pool: poll}
//			continue
//		}
//
//		// 正常字符部分，移动指针向后
//		i++
//		if isLocalRead && '0' <= elem && elem <= '9' {
//			origin[originLen] = elem
//			originLen++
//		}
//
//	}
//
//	point.DtTime(startTime, "二进制过滤时间")
//
//	// 归还切边
//	//buff = nil
//
//	org := datas[0:datasIndex]
//	//ioutil.WriteFile("./bin/output2.txt", org[0].Locationid, 0666) //写入文件(字节数组)
//	cb(&org)
//}
