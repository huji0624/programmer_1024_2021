package find_1024

import (
	"DigSakura/dig/point"
	"DigSakura/dig/submit"
	"bytes"
	"sync"
)

//策略工具
var v_tactics_dataAllIn bool = false
var v_tactics_du bool = true

//预处理数据
var v_pre_databucket_cells [39][2][10000]*Cells
var v_pre_databucket_cells_lens = [39][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
var v_pre_databucket_cells_lens_all = 0

var v_pre_mathbucket_cells [10000][2]*Cells
var v_pre_mathbucket_cells_lens = 0

var v_pre_math_bool = false

//计算中数据
var v_math_databucket_cells [39][2][10000]*Cells
var v_math_databucket_cells_lens = [39][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
var v_math_databucket_cells_lens_all = 0

var v_math_mathbucket_cells [10000][2]*Cells
var v_math_mathbucket_cells_lens = 0

var v_math_data_bool = false

//计算结果使用到的数据
var v_mathend_1024_cells [20][2]*Cells
var v_mathend_1024_cells_len = 0
var v_mathend_0_cells *Cells
var v_mathend_0_cells_len = 0
var v_mathend_00_cells *Cells
var v_mathend_00_cells_len = 0

//1024,512,256,128,64,32,16,8,4,2,0

var inputSync sync.Mutex

func StartFind1024() {

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				StartFind1024()
			}
		}()
		for true {
			// 单线程同步
			if !v_pre_math_bool && v_pre_mathbucket_cells_lens+v_math_mathbucket_cells_lens+v_mathend_1024_cells_len >= 20 {
				inputSync.Lock()
				//->计算队列
				if v_tactics_du {
					// 单车变摩托
					if v_pre_databucket_cells_lens_all+v_pre_mathbucket_cells_lens > 256 {
						v_tactics_du = false
						input_math(v_math_data_bool)
					}
				} else {
					input_math(v_math_data_bool)
				}
				inputSync.Unlock()
				//->计算
				math()
			}
		}
	}()
}

//->预处理队列
func Input_cells(Locationid string, Magic string, isremainder bool) {

	//// 宝藏原始值
	//Locationid []byte
	//// 宝藏值
	//Magic []byte
	t_cells_data := &Cells{}
	t_cells_data.Locationid = []byte(Locationid)
	t_cells_data.Magic = []byte(Magic)

	//println(Locationid,"   ",Magic )
	// 预处理
	b := magic_bucket(t_cells_data, isremainder)

	// 公共内存区加锁
	inputSync.Lock()
	defer inputSync.Unlock()
	pre_save(t_cells_data, b)

}

//导入数据
func input_math(data_appends bool) {

	v_pre_math_bool = true
	v_math_data_bool = true

	if !data_appends {
		//无math数据情况下
		v_math_databucket_cells = v_pre_databucket_cells

		v_math_databucket_cells_lens = v_pre_databucket_cells_lens
		v_pre_databucket_cells_lens = [39][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

		v_math_databucket_cells_lens_all = v_pre_databucket_cells_lens_all
		v_pre_databucket_cells_lens_all = 0

		v_math_mathbucket_cells = v_pre_mathbucket_cells

		v_math_mathbucket_cells_lens = v_pre_mathbucket_cells_lens
		v_pre_mathbucket_cells_lens = 0
	} else {
		//有math数据情况下
		for i := 0; i < len(v_pre_databucket_cells_lens); i++ {
			for j := 0; j < len(v_pre_databucket_cells_lens[i]); j++ {
				for k := 0; k < v_pre_databucket_cells_lens[i][j]; k++ {
					v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]+k] = v_pre_databucket_cells[i][j][k]
				}
				v_math_databucket_cells_lens[i][j] = v_math_databucket_cells_lens[i][j] + v_pre_databucket_cells_lens[i][j]
			}
		}
		v_pre_databucket_cells_lens = [39][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

		v_math_databucket_cells_lens_all = v_math_databucket_cells_lens_all + v_pre_databucket_cells_lens_all
		v_pre_databucket_cells_lens_all = 0

		for i := 0; i < v_pre_mathbucket_cells_lens; i++ {
			v_math_mathbucket_cells[v_math_mathbucket_cells_lens+i] = v_pre_mathbucket_cells[i]
		}
		v_math_mathbucket_cells_lens = v_math_mathbucket_cells_lens + v_pre_mathbucket_cells_lens
		v_pre_mathbucket_cells_lens = 0
	}
}

//计算算式
func math() {
	//尝试算术,算式主模型
	for v_mathend_1024_cells_len < 20 {
		if v_math_mathbucket_cells_lens < 1 {
			break
		}
		v_mathend_1024_cells[v_mathend_1024_cells_len] = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1]
		v_math_mathbucket_cells_lens--
		v_mathend_1024_cells_len++
	}

	if v_mathend_1024_cells_len < 20 {
		v_pre_math_bool = false
		return
	}

	//尝试附加0模型
	if v_mathend_0_cells_len == 0 {
		if v_mathend_00_cells_len == 0 {
			if v_math_mathbucket_cells_lens > 0 {
				v_mathend_0_cells = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1][1]
				v_mathend_00_cells = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1][0]
				v_mathend_0_cells_len = 1
				v_mathend_00_cells_len = 1
			} else if v_math_databucket_cells_lens_all > 1 {
				func() {
					for i := 0; i < len(v_math_databucket_cells_lens); i++ {
						for j := 0; j < len(v_math_databucket_cells_lens[i]); j++ {
							if v_math_databucket_cells_lens[i][j] > 0 {
								v_mathend_0_cells = v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1]
								v_math_databucket_cells_lens[i][j]--
								v_math_databucket_cells_lens_all--
								v_mathend_0_cells_len = 1
								return
							}
						}
					}
				}()
			}
		} else {
			if v_math_mathbucket_cells_lens > 0 {
				b := magic_bucket(v_mathend_00_cells, false)
				v_math_databucket_cells[b][1][v_math_databucket_cells_lens[b][1]] = v_mathend_00_cells
				v_math_databucket_cells_lens[b][1]++
				v_math_databucket_cells_lens_all++

				v_mathend_0_cells = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1][1]
				v_mathend_00_cells = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1][0]
				v_mathend_0_cells_len = 1
				v_mathend_00_cells_len = 1
			} else if v_math_databucket_cells_lens_all > 1 {
				func() {
					for i := 0; i < len(v_math_databucket_cells_lens); i++ {
						for j := 0; j < len(v_math_databucket_cells_lens[i]); j++ {
							if v_math_databucket_cells_lens[i][j] > 0 {
								t_true_compare := magic_true_compare(v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1], v_mathend_00_cells)

								if t_true_compare < 0 {
									v_mathend_0_cells = v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1]
								} else if t_true_compare > 0 {
									v_mathend_0_cells = v_mathend_00_cells
									v_mathend_00_cells = v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1]
								} else {
									continue
								}
								v_math_databucket_cells_lens[i][j]--
								v_math_databucket_cells_lens_all--
								v_mathend_0_cells_len = 1
								return
							}
						}
					}
				}()
			}
		}
	}

	//尝试附加00模型
	if v_mathend_00_cells_len == 0 && v_mathend_0_cells_len == 1 {
		if v_math_mathbucket_cells_lens > 0 {
			b := magic_bucket(v_mathend_0_cells, false)
			v_math_databucket_cells[b][1][v_math_databucket_cells_lens[b][1]] = v_mathend_00_cells
			v_math_databucket_cells_lens[b][1]++
			v_math_databucket_cells_lens_all++

			v_mathend_0_cells = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1][1]
			v_mathend_00_cells = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1][0]
			v_mathend_0_cells_len = 1
			v_mathend_00_cells_len = 1
		} else if v_math_databucket_cells_lens_all > 1 {
			func() {
				for i := 0; i < len(v_math_databucket_cells_lens); i++ {
					for j := 0; j < len(v_math_databucket_cells_lens[i]); j++ {
						if v_math_databucket_cells_lens[i][j] > 0 {
							t_true_compare := magic_true_compare(v_mathend_0_cells, v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1])

							if t_true_compare < 0 {
								v_mathend_00_cells = v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1]
							} else if t_true_compare > 0 {
								v_mathend_00_cells = v_mathend_0_cells
								v_mathend_0_cells = v_math_databucket_cells[i][j][v_math_databucket_cells_lens[i][j]-1]
							} else {
								continue
							}
							v_math_databucket_cells_lens[i][j]--
							v_math_databucket_cells_lens_all--
							v_mathend_00_cells_len = 1
							return
						}
					}
				}
			}()
		}
	}

	//合成公式
	if v_mathend_1024_cells_len == 20 {
		mixture()
	}
}

//————————————————————————————算式工具————————————————————————————
const c_math_left_kuo byte = '('
const c_math_right_kuo byte = ')'
const c_math_add byte = '+'
const c_math_sub byte = '-'
const c_math_multi byte = '*'
const c_math_div byte = '/'

//合成公式
func mixture() {

	var buffer bytes.Buffer
	for i := 0; i < v_mathend_1024_cells_len; i = i + 2 {
		if i != 0 {
			buffer.Write([]byte{c_math_multi})
		}
		buffer.Write([]byte{c_math_left_kuo})
		buffer.Write(v_mathend_1024_cells[i][0].Locationid)
		buffer.Write([]byte{c_math_div})
		buffer.Write(v_mathend_1024_cells[i][1].Locationid)
		buffer.Write([]byte{c_math_add})
		buffer.Write(v_mathend_1024_cells[i+1][0].Locationid)
		buffer.Write([]byte{c_math_div})
		buffer.Write(v_mathend_1024_cells[i+1][1].Locationid)
		buffer.Write([]byte{c_math_right_kuo})
	}

	if v_mathend_0_cells_len == 1 && v_mathend_00_cells_len == 1 {
		buffer.Write([]byte{c_math_sub})
		buffer.Write(v_mathend_0_cells.Locationid)
		buffer.Write([]byte{c_math_div})
		buffer.Write(v_mathend_00_cells.Locationid)
		if v_math_mathbucket_cells_lens > 0 {
			for i := 0; i < v_math_mathbucket_cells_lens; i++ {
				buffer.Write([]byte{c_math_div})
				buffer.Write(v_math_mathbucket_cells[i][1].Locationid)
				buffer.Write([]byte{c_math_div})
				buffer.Write(v_math_mathbucket_cells[i][0].Locationid)
			}
		}
		if v_math_databucket_cells_lens_all > 0 {
			for i := len(v_math_databucket_cells_lens) - 1; i >= 0; i-- {
				for j := 0; j < len(v_math_databucket_cells_lens[i]); j++ {
					for k := 0; k < v_math_databucket_cells_lens[i][j]; k++ {
						if i == 0 && j == 1 {
							buffer.Write([]byte{c_math_sub})
							buffer.Write(v_math_databucket_cells[i][j][k].Locationid)
						} else {
							buffer.Write([]byte{c_math_div})
							buffer.Write(v_math_databucket_cells[i][j][k].Locationid)
						}
					}
				}
			}
		}
	}

	//最终公式
	formule := buffer.Bytes()

	//push公式
	//println(string(formule)) //test

	resp := submit.SubmitFormula(string(formule), 2)
	if resp == nil {
		// 网络返回为空
		println("网络返回为空")
		Formule_success()
		return
	}
	//0表示成功，3表示算式有被使用了的locationid，其他返回值表示错误
	switch resp.Errorno {
	case 0:
		Formule_success()
	case 3:
		aa := [][]byte{}
		for _, loc := range resp.Data {
			aa = append(aa, []byte(loc))
			point.ServerLocalIdUse.Count++
		}
		Formule_failed(aa) //test
	default:
		Formule_success()
	}

}

//————————————————————————————网络反馈————————————————————————————
//算式成功
func Formule_success() {
	v_math_databucket_cells_lens = [39][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	v_math_databucket_cells_lens_all = 0
	v_math_mathbucket_cells_lens = 0
	v_mathend_1024_cells_len = 0
	v_mathend_0_cells_len = 0
	v_mathend_00_cells_len = 0

	v_math_data_bool = false
	v_pre_math_bool = false
	if v_tactics_dataAllIn {
		input_math(v_math_data_bool)
		math()
	}
}

//算式失败
func Formule_failed(data [][]byte) {
	t_loctionids := &data
	if len(*t_loctionids) >= v_mathend_1024_cells_len+v_mathend_0_cells_len+v_mathend_00_cells_len+v_math_databucket_cells_lens_all+v_math_mathbucket_cells_lens {
		Formule_success()
		return
	}
	v_pre_math_bool = false
	input_math(v_math_data_bool)
	ban_cells_math(t_loctionids)
	math()
}
