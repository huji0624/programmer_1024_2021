package find_1024

import "bytes"

//废除替换已经被使用的cells
func ban_cells_math(t_loctionids *[][]byte) {
	t_loctionids_v := *t_loctionids
	for ti := 0; ti < len(t_loctionids_v); ti++ {
		find_change_cells_in_formule(&t_loctionids_v[ti])
	}
}

//找到被占用的cells，然后替换
func find_change_cells_in_formule(t_loctionid *[]byte) {
	t_loctionid_v := *t_loctionid
	//主框架
	for i := 0; i < v_mathend_1024_cells_len; i++ {
		for j := 0; j < 2; j++ {
			//find
			if bytes.Equal(v_mathend_1024_cells[i][j].Locationid, t_loctionid_v) {
				//drop
				j_frined := 1
				if j == 1 {
					j_frined = 0
				}
				b := magic_bucket(v_mathend_1024_cells[i][j_frined], false)
				v_math_databucket_cells[b][0][v_math_databucket_cells_lens[b][0]] = v_mathend_1024_cells[i][j_frined]
				v_math_databucket_cells_lens[b][0]++
				v_math_databucket_cells_lens_all++
				//change
				if v_math_mathbucket_cells_lens > 0 {
					v_mathend_1024_cells[i] = v_math_mathbucket_cells[v_math_mathbucket_cells_lens-1]
					v_math_mathbucket_cells_lens--
				} else {
					v_mathend_1024_cells_len--
					for i < v_mathend_1024_cells_len {
						v_mathend_1024_cells[i] = v_mathend_1024_cells[i+1]
						i++
					}
				}
				//return
				return
			}
		}
	}

	//0框架
	if bytes.Equal(v_mathend_0_cells.Locationid, t_loctionid_v) {
		//nochange
		v_mathend_0_cells_len = 0
		//return
		return
	}

	//00框架
	if bytes.Equal(v_mathend_00_cells.Locationid, t_loctionid_v) {
		//nochange
		v_mathend_00_cells_len = 0
		//return
		return
	}

	///000_databucket框架
	if v_math_databucket_cells_lens_all > 0 {
		for i := 0; i < len(v_math_databucket_cells_lens); i++ {
			for j := 0; j < len(v_math_databucket_cells_lens[i]); j++ {
				for k := 0; k < v_math_databucket_cells_lens[i][j]; k++ {
					if bytes.Equal(v_math_databucket_cells[i][j][k].Locationid, t_loctionid_v) {
						//change
						v_math_databucket_cells_lens[i][j]--
						v_math_databucket_cells_lens_all--
						for k < v_math_databucket_cells_lens[i][j] {
							v_math_databucket_cells[i][j][k] = v_math_databucket_cells[i][j][k+1]
							k++
						}
						//return
						return
					}
				}
			}
		}
	}

	///000_mathbucket框架
	for i := 0; i < v_math_mathbucket_cells_lens; i++ {
		for j := 0; j < 2; j++ {
			if bytes.Equal(v_math_mathbucket_cells[i][j].Locationid, t_loctionid_v) {
				//drop
				j_frined := 1
				if j == 1 {
					j_frined = 0
				}
				b := magic_bucket(v_math_mathbucket_cells[i][j_frined], false)
				v_math_databucket_cells[b][0][v_math_databucket_cells_lens[b][0]] = v_math_mathbucket_cells[i][j_frined]
				v_math_databucket_cells_lens[b][0]++
				v_math_databucket_cells_lens_all++
				//change
				v_math_mathbucket_cells_lens--
				for i < v_math_mathbucket_cells_lens {
					v_math_mathbucket_cells[i] = v_math_mathbucket_cells[i+1]
					i++
				}
				//return
				return
			}
		}
	}

}

//————————————————————————————策略工具————————————————————————————
//是否将剩下的缓存数据全部放入计算池
func Tactics_dataAllIn(b bool) {
	v_tactics_dataAllIn = b
	if v_tactics_dataAllIn && !v_pre_math_bool {
		input_math(v_math_data_bool)
	}
}

//————————————————————————————辅助工具————————————————————————————

//处理完成数据，存储数据
func pre_save(t_cells_data *Cells, t_bucket int) {
	if t_bucket > 9 && t_cells_data.Magic[len(t_cells_data.Magic)+6-t_bucket] <= '1' {
		v_pre_databucket_cells[t_bucket][1][v_pre_databucket_cells_lens[t_bucket][1]] = t_cells_data
		v_pre_databucket_cells_lens[t_bucket][1]++
		v_pre_databucket_cells_lens_all++
		return
	} else if t_bucket > 0 {
		if v_pre_databucket_cells_lens[t_bucket][0] > 0 {
			for i := 0; i < v_pre_databucket_cells_lens[t_bucket][0]; i++ {
				t_compare := magic_compare(t_cells_data, v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]-1-i], t_bucket)
				//save math bucket
				if t_compare > 0 {
					v_pre_mathbucket_cells[v_pre_mathbucket_cells_lens][0] = t_cells_data
					v_pre_mathbucket_cells[v_pre_mathbucket_cells_lens][1] = v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]-1-i]
				} else if t_compare < 0 {
					v_pre_mathbucket_cells[v_pre_mathbucket_cells_lens][0] = v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]-1-i]
					v_pre_mathbucket_cells[v_pre_mathbucket_cells_lens][1] = t_cells_data
				} else {
					continue
				}
				v_pre_mathbucket_cells_lens++

				v_pre_databucket_cells_lens[t_bucket][0]--
				v_pre_databucket_cells_lens_all--
				for i < v_pre_databucket_cells_lens[t_bucket][0] {
					v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]-1-i] = v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]-i]
					i++
				}
				return
			}
		}

		v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]] = t_cells_data
		v_pre_databucket_cells_lens[t_bucket][0]++
		v_pre_databucket_cells_lens_all++
	} else if t_bucket == 0 && t_cells_data.Magic[len(t_cells_data.Magic)-1] == '1' {
		v_pre_databucket_cells[t_bucket][0][v_pre_databucket_cells_lens[t_bucket][0]] = t_cells_data
		v_pre_databucket_cells_lens[t_bucket][0]++
		v_pre_databucket_cells_lens_all++
	} else {
		v_pre_databucket_cells[t_bucket][1][v_pre_databucket_cells_lens[t_bucket][1]] = t_cells_data
		v_pre_databucket_cells_lens[t_bucket][1]++
		v_pre_databucket_cells_lens_all++
	}
}

//预计算2倍内的大小比较,
func magic_compare(c1 *Cells, c2 *Cells, t_bucket int) int {
	m1 := c1.Magic
	m2 := c2.Magic

	l1 := len(m1)
	l2 := len(m2)

	if t_bucket < 10 {
		if l1 > l2 {
			return 1
		} else if l1 < l2 {
			return -1
		} else {
			for i := 0; i < l1; i++ {
				if m1[i] > m2[i] {
					return 1
				} else if m1[i] < m2[i] {
					return -1
				}
			}
		}
		return 0
	} else {
		lmin := l1
		if l1 > l2 {
			lmin = l2
		}

		s1 := m1[l1-lmin]
		s2 := m2[l2-lmin]
		if s1 < 50 || s2 < 50 {
			return 0
		}
		if s1 == s2 {
			return 0
		}

		min_s1 := (s1-48)/2 + 48
		max_s1 := (s1-48)*2 + 48
		if min_s1 < s2 && s2 < s1 {
			return 1
		} else if s1 < s2 && s2 < max_s1 {
			return -1
		}
		return 0
	}
}

func magic_true_compare(c1 *Cells, c2 *Cells) int {
	m1 := c1.Magic
	m2 := c2.Magic

	l1 := len(m1)
	l2 := len(m2)

	for i := 0; i < len(m1); i++ {
		if m1[i] == '0' {
			l1--
		} else {
			break
		}
	}
	for i := 0; i < len(m2); i++ {
		if m2[i] == '0' {
			l2--
		} else {
			break
		}
	}

	if l1 > l2 {
		return 1
	} else if l1 < l2 {
		return -1
	} else {
		for i := 0; i < l1; i++ {
			if m1[len(m1)-l1+i] > m2[len(m2)-l2+i] {
				return 1
			} else if m1[len(m1)-l1+i] < m2[len(m2)-l2+i] {
				return -1
			}
		}
	}

	return 0
}

//计算cells的值在哪个桶
func magic_bucket(t_cells_data *Cells, isremainder bool) int {

	//处理求余数据
	t_magic := t_cells_data.Magic
	t_magiclen := len(t_magic)
	var i int
	for i = 0; i < len(t_magic); i++ {
		if t_magic[i] == '0' {
			t_magiclen--
		} else {
			break
		}
	}
	//处理数据长度
	if t_magiclen > 4 {
		return t_magiclen + 6
	} else if t_magiclen == 4 {
		if isremainder {
			return 9
		} else {
			if t_magic[i] >= '2' {
				return 10
			} else {
				if t_magic[i+1] >= '1' {
					return 10
				} else {
					if t_magic[i+2] >= '3' {
						return 10
					} else {
						if t_magic[i+3] >= '5' {
							return 10
						} else {
							return 9
						}
					}
				}
			}
		}
	} else if t_magiclen == 3 {
		if t_magic[i] >= '6' {
			return 9
		} else if t_magic[i] == '5' {
			if t_magic[i+1] >= '2' {
				return 9
			} else if t_magic[i+1] == '1' {
				if t_magic[i+2] >= '2' {
					return 9
				} else {
					return 8
				}
			} else {
				return 8
			}
		} else if t_magic[i] >= '3' {
			return 8
		} else if t_magic[i] == '2' {
			if t_magic[i+1] >= '6' {
				return 8
			} else if t_magic[i+1] == '5' {
				if t_magic[i+2] >= '6' {
					return 8
				} else {
					return 7
				}
			} else {
				return 7
			}
		} else {
			if t_magic[i+1] >= '3' {
				return 7
			} else if t_magic[i+1] == '2' {
				if t_magic[i+2] >= '8' {
					return 7
				} else {
					return 6
				}
			} else {
				return 6
			}
		}
	} else if t_magiclen == 2 {
		if t_magic[i] >= '7' {
			return 6
		} else if t_magic[i] == '6' {
			if t_magic[i+1] >= '4' {
				return 6
			} else {
				return 5
			}
		} else if t_magic[i] > '4' {
			return 5
		} else if t_magic[i] == '3' {
			if t_magic[i+1] >= '2' {
				return 5
			} else {
				return 4
			}
		} else if t_magic[i] == '2' {
			return 4
		} else {
			if t_magic[i+1] >= '6' {
				return 4
			} else {
				return 3
			}
		}
	} else if t_magiclen == 1 {
		if t_magic[i] >= '8' {
			return 3
		} else if t_magic[i] >= '4' {
			return 2
		} else if t_magic[i] >= '2' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}

}

//宝藏结构
type Cells struct {
	// 宝藏原始值
	Locationid []byte
	// 宝藏值
	Magic []byte
}
