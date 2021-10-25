package magic

import (
	"DigSakura/dig/find_1024"
	"DigSakura/dig/point"
	"DigSakura/dig/submit"
)

type Magic struct {
	// 宝藏原始值
	Locationid string
	// 宝藏值
	Magic string
	// 计算类型 1 +，2 -，3 *，4 /,5 %
	Type int8
}

// 提交宝藏信息
func (c *Magic) Submit() {
	// 提交结果到服务器
	submit.SubmitLocationId(c.Locationid)
	go func() {
		if r := recover(); r != nil {
			println("处理函数运行中发生错误", c.Locationid, "    ", c.Magic)
		}
		// 提交到公式计算器
		find_1024.Input_cells(c.Locationid, c.Magic, true)
	}()

	// 统计宝藏数量
	point.TreasuresNum.Count++

	switch c.Type {
	case 1:
		point.AddNum.Count++
	case 2:
		point.SubNum.Count++
	case 3:
		point.MulNum.Count++
	case 5:
		point.ModNum.Count++
	}

}
