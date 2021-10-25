package del

//
//import (
//	"DigSakura/dig/point"
//	"math/big"
//	"strconv"
//)
//
//type Cell struct {
//	// 宝藏原始值
//	Locationid []byte
//	// 从宝藏中找到的数字
//	Origin []byte
//	// 宝藏值
//	Magic []byte
//	// 切片内容所在的内存块
//	Pool *BuffPool
//}
//
//// 处理宝藏
//func (c *Cell) Process(a, b, add, tarNum *big.Int) {
//	magic := string(c.Magic)
//	origin := string(c.Origin)
//	// 先解析目标值
//	numB, err := strconv.ParseInt(magic, 10, 64)
//	if err != nil {
//		c.processBig(a, b, add, tarNum, magic, origin)
//		return
//	}
//	// 目标值不越界，原值肯定不越界(加减法时)，
//	numA, err := strconv.ParseInt(origin, 10, 64)
//	if err != nil {
//		c.processBig(a, b, add, tarNum, magic, origin)
//		return
//	}
//	var const1024 int64 = 1024
//
//	// 小数据处理
//	point.Rec(point.ProcesLowNum)
//
//	if numA+const1024 == numB {
//		c.submit("+", false)
//		return
//	}
//	if numA-const1024 == numB {
//		c.submit("-", false)
//		return
//	}
//
//	if numA*const1024 == numB {
//		c.submit("*", false)
//		return
//	}
//
//	//if len(c.Magic) < 5 {
//	if numA%const1024 == numB {
//		c.submit("%", false)
//		return
//	}
//	//} else {
//	//	// 统计跳过量
//	//	point.Rec(point.SkipModNum)
//	//}
//
//}
//
//// 大数据处理
//func (c *Cell) processBig(a, b, add, tarNum *big.Int, Magic, Origin string) {
//	point.Rec(point.ProcesBigNum)
//
//	a.SetString(Magic, 10)
//	b.SetString(Origin, 10)
//
//	add = add.Add(b, tarNum)
//	if a.Cmp(add) == 0 {
//		c.submit("+", true)
//		return
//	}
//
//	add = add.Sub(b, tarNum)
//	if a.Cmp(add) == 0 {
//		c.submit("-", true)
//		return
//	}
//
//	add = add.Mul(b, tarNum)
//	if a.Cmp(add) == 0 {
//		c.submit("*", true)
//		return
//	}
//	//宝藏长度小于5 ，有可能满足求余运算
//	//if len(c.Magic) < 5 {
//	add = add.Mod(b, tarNum)
//	if a.Cmp(add) == 0 {
//		c.submit("%", true)
//		return
//	}
//	//} else {
//	//	// 统计跳过量
//	//	point.Rec(point.SkipModNum)
//
//	//}
//
//}
//
//// 提交宝藏信息
//func (c *Cell) submit(method string, isBig bool) {
//	// 复制当前Cell数据！！！发送到宝藏中心。如果不复制内容，所在大块内存区将无法释放
//
//	// 提交结果到服务器
//	//submit.SubmitLocationId(c.Locationid)
//
//	// 统计宝藏数量
//	point.Rec(point.TreasuresNum)
//
//	switch method {
//	case "+":
//		point.Rec(point.AddNum)
//	case "-":
//		point.Rec(point.SubNum)
//	case "*":
//		point.Rec(point.MulNum)
//	case "%":
//		//println(method, string(c.Locationid), "  ", string(c.Magic))
//		point.Rec(point.ModNum)
//	}
//
//	if isBig {
//		point.Rec(point.TreasuresBigNum)
//	} else {
//		point.Rec(point.TreasuresLowNum)
//	}
//
//	//println("发现宝藏", methad, Mg)
//	//g.Dump(string(c.Locationid) + "  " + string(c.Origin) + "  " + string(c.Magic))
//
//}
