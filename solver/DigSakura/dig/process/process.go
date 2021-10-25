package process

import "DigSakura/dig/magic"

func Read(content *[]byte, bufflen int) {
	buff := *content
	//bufflen := n
	suancount := make([]int, 5)

	//倒序循环读
	for i := bufflen; i >= 0; {
		//elem := buff[i]
		//倒读
		if buff[i] == '}' {
			//是否进行如下筛选
			isadd := true
			issub := true
			ismulti := true
			isremainder := true
			//求余实际计算开关
			isremainder_math := false
			//暂定状态
			isadd_status := false
			issub_status := false
			ismulti_status := false
			isremainder_status := false
			//去除乘法筛选
			if buff[i-2] == '1' || buff[i-2] == '3' || buff[i-2] == '5' || buff[i-2] == '7' || buff[i-2] == '9' {
				ismulti = false
			}
			//去除求余筛选
			if '0' <= buff[i-6] && buff[i-6] <= '9' {
				isremainder = false
			} else {
				if '2' <= buff[i-5] && buff[i-5] <= '9' {
					isremainder = false
				} else if buff[i-5] == '1' {
					if buff[i-4] != '0' {
						isremainder = false
					} else {
						if '3' <= buff[i-3] && buff[i-3] <= '9' {
							isremainder = false
						} else if buff[i-3] == '2' {
							if '4' <= buff[i-2] && buff[i-2] <= '9' {
								isremainder = false
							}
						}
					}
				}
			}

			//向上查找,号
			var ii int
			for ii = i - 12; ii >= 0; ii-- {
				if buff[ii] == ',' {
					break
				}
			}
			//magic长度
			//magicLen := i - ii - 11

			//记录下一个{地址
			iii := ii - 82

			//定义位数
			forNum := 0
			//用于求余函数的最大10位的指针记录
			remainder_locNumFind := [10]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
			remainder_j := 0
			remainder_k := 0
			remainder_switch2trees := make([]bool, 11)
			//remainder_endValue := make([]byte, 10)
			remainder_maxyumathlength := 0

			//定义进位桶
			var addCarry byte = '0'
			var subCarry byte = '0'
			multiCarry := [4]byte{'0', '0', '0', '0'}
			remainderCarry := [10][30]byte{
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}}

			//定义magic位数
			magicNum := i - 1
			gold := func() bool {
				//loctionID右字符倒序运算
				for locNum := ii - 2; locNum >= ii-65; locNum-- {

					//当剩余loctionid长度没有剩余magic长度长时，无解
					if locNum < magicNum-80 {
						return false
					}

					//找到一个数字字符串
					if '0' <= buff[locNum] && buff[locNum] <= '9' {
						//超出预解值
						if buff[locNum] != '0' {
							if isadd_status {
								isadd_status = false
							}
							if issub_status {
								issub_status = false
							}
							if ismulti_status {
								ismulti_status = false
							}
						}

						//增加一位找到的数字
						forNum++
						magicNum--

						//计算固定值
						var constput byte = 0
						switch forNum {
						case 1:
							constput = 4
						case 2:
							constput = 2
						case 4:
							constput = 1
						}

						//加法运算
						if isadd {
							o1, o2 := add99_table(buff[locNum], addCarry, constput)
							if o1 == buff[magicNum] {
								addCarry = o2
								if magicNum == ii+10 && o2 == '0' {
									isadd_status = true
									isadd = false
								} else if magicNum <= ii+11 && o2 == buff[magicNum-1] {
									isadd_status = true
									isadd = false
								}
							} else {
								isadd = false
							}
						}
						//减法运算
						if issub {
							o1, o2 := sub99_table(buff[locNum], subCarry, constput)
							if o1 == buff[magicNum] {
								subCarry = o2
								if magicNum == ii+10 && o2 == '0' {
									issub_status = true
									issub = false
									/*  } else {
									issub_status = false */
								}
							} else if magicNum == ii+9 {
								if o2 == '0' {
									issub_status = true
									issub = false
								}
							} else {
								issub = false
							}
						}
						//乘法运算
						if ismulti {
							o1, o2, o3, o4 := multi1024_table(buff[locNum], '0')
							multiCarry[0], o1 = add99_table(multiCarry[0], o1, 0)
							multiCarry[1], o2 = add99_table(multiCarry[1], o2, o1-48)
							multiCarry[2], o3 = add99_table(multiCarry[2], o3, o2-48)
							multiCarry[3], o4 = add99_table(multiCarry[3], o4, o3-48)
							if multiCarry[0] == buff[magicNum] {
								multiCarry[0] = multiCarry[1]
								multiCarry[1] = multiCarry[2]
								multiCarry[2] = multiCarry[3]
								multiCarry[3] = o4
								if magicNum == ii+13 && multiCarry[0] == buff[magicNum-1] && multiCarry[1] == buff[magicNum-2] && multiCarry[2] == buff[magicNum-3] && o4 == '0' {
									ismulti_status = true
									ismulti = false
								} else if magicNum == ii+14 && multiCarry[0] == buff[magicNum-1] && multiCarry[1] == buff[magicNum-2] && multiCarry[2] == buff[magicNum-3] && o4 == buff[magicNum-4] {
									ismulti_status = true
									ismulti = false
								}
							} else {
								ismulti = false
							}
						}
						//求余入库
						if isremainder {
							remainder_locNumFind[remainder_maxyumathlength] = buff[locNum]
							remainder_maxyumathlength++
							if remainder_maxyumathlength == 10 {
								isremainder = false
								isremainder_math = true
							}
						}
					}

					//最后一位字符强行打开求余计算
					if locNum == ii-65 {
						//isremainder = false
						isremainder_math = true
					}
					//求余运算
					if isremainder_math {
						for imaxlength := remainder_maxyumathlength - 1; imaxlength > 0; imaxlength-- {
							if remainder_locNumFind[imaxlength] == '0' {
								remainder_maxyumathlength--
							} else {
								break
							}
						}
						func() {
							for remainder_j < remainder_maxyumathlength {
								//申明回滚变量
								rollback := false
								//———————————————magic读入—————————————————
								//读取magic对应的数字
								var magicGetNum byte = '0'
								if '0' <= buff[i-2-remainder_j] && buff[i-2-remainder_j] <= '9' {
									magicGetNum = buff[i-2-remainder_j]
								}
								remainderCarry[remainder_j][remainder_k] = magicGetNum
								remainder_j++ //增加找到的数字位数长度
								remainder_k++ //完成一个步骤
								//———————————————减法借位—————————————————
								jiesuanNum := remainder_locNumFind[remainder_j-1]
								//减少K步骤内的所有值

								ii_k := remainder_k - 12
								if ii_k < 0 {
									ii_k = 0
								}
								for ; ii_k < remainder_k; ii_k++ {
									var ii_output2 byte
									jiesuanNum, ii_output2 = sub99_table(jiesuanNum, remainderCarry[remainder_j-1][ii_k], 0)
									ii_j := remainder_j
									for ii_j < 10 {
										remainderCarry[ii_j][remainder_k], ii_output2 = add99_table(remainderCarry[ii_j][remainder_k], ii_output2, 0)
										ii_j++
										if ii_output2 == '0' {
											break
										}
									}
								}
								remainder_k++ //完成一个步骤
								//———————————————算解—————————————————
								if jiesuanNum == '1' || jiesuanNum == '3' || jiesuanNum == '5' || jiesuanNum == '7' || jiesuanNum == '9' {
									//错误，准备回滚
									rollback = true
								} else {
									//正确部分
									var chuyu_o2, chuyu_o3, chuyu_o4 byte
									if !remainder_switch2trees[remainder_j] {
										chuyu_o2, chuyu_o3, chuyu_o4 = remainder1024_table_1(jiesuanNum)
									} else {
										chuyu_o2, chuyu_o3, chuyu_o4 = remainder1024_table_2(jiesuanNum)
									}
									/*
										remainder_endValue[remainder_j-1] = chuyu_o1                                                    //test
										println("get->", "remainder_j:", remainder_j, "remainder_endValue", string(remainder_endValue)) //test
									*/
									if remainder_j+2 >= remainder_maxyumathlength {
										if chuyu_o4 != '0' {
											rollback = true //huigun
										}
									} else {
										remainderCarry[remainder_j+2][remainder_k] = chuyu_o4
									}
									if remainder_j+1 >= remainder_maxyumathlength {
										if chuyu_o3 != '0' {
											rollback = true //huigun
										}
									} else {
										remainderCarry[remainder_j+1][remainder_k] = chuyu_o3
									}
									if remainder_j >= remainder_maxyumathlength {
										if chuyu_o2 != '0' {
											rollback = true //huigun
										}
									} else {
										remainderCarry[remainder_j][remainder_k] = chuyu_o2
									}
								}
								remainder_k++ //完成一个步骤
								//—————————————————————回滚步骤———————————————————
								if rollback {
									//回退J进位，用二叉树判断回退多少
									remainder_j_trees_back := 0
									remainder_switch2trees[remainder_j] = false
									for remainder_j_trees_back = 0; remainder_j_trees_back < remainder_j; remainder_j_trees_back++ {
										if remainder_switch2trees[remainder_j-1-remainder_j_trees_back] {
											if remainder_j-1-remainder_j_trees_back == 0 {
												isremainder_math = false
												return
											}
											remainder_switch2trees[remainder_j-1-remainder_j_trees_back] = false
										} else {
											remainder_switch2trees[remainder_j-1-remainder_j_trees_back] = true
											break
										}
									}
									//正式回退
									for remainder_j_trees_back >= -1 {
										/*
											if remainder_j >= 0 {
												remainder_endValue[remainder_j] = 0
											} //获得最终值
											println("back->", "remainder_j:", remainder_j, "remainder_endValue", string(remainder_endValue)) //test
										*/
										//回退二叉树
										remainder_j--
										remainder_k = remainder_k - 3

										if remainder_j < 0 || remainder_k < 0 {
											remainder_j = 0
											remainder_k = 0
											isremainder_math = false
											return //找不到正确的解
										}
										//删除回退部分的数据
										for back_k := remainder_k; back_k < remainder_k+3; back_k++ {
											for back_j := remainder_j; back_j < 10; back_j++ {
												remainderCarry[back_j][back_k] = '0'
											}
										}

										//回退计数
										remainder_j_trees_back--
									}
								}

								//成功找到解的条件
								if remainder_k >= remainder_maxyumathlength*3-1 {
									isremainder_status = true
									return
								}
							}

						}()
						if isremainder_status {
							return true
						}
						isremainder_math = false
					}
				}
				//4个运算均不对，则不是宝藏，无解
				if isadd_status || issub_status || ismulti_status || isremainder_status {
					return true
				}

				return false
			}()

			//存储locationID和magic
			if gold {
				magicData := &magic.Magic{}
				suancount[0]++
				if isadd_status {
					magicData.Type = 1
					suancount[1]++
				} else if issub_status {
					magicData.Type = 2
					suancount[2]++
				} else if ismulti_status {
					magicData.Type = 3
					suancount[3]++
					//println(string(buff[ii-65:ii-1]) + "     " + string(buff[ii+10:i-1]))
				} else if isremainder_status {
					magicData.Type = 5
					suancount[4]++
				}

				magicData.Locationid = string(buff[ii-65 : ii-1])
				magicData.Magic = string(buff[ii+10 : i-1])
				magicData.Submit()
				//magic.SubmitMagicToChan(magicData)

				//println(string(buff[ii-65:ii-1]) + "     " + string(buff[ii+10:i-1]))
			}

			//println("gold:", gold, ",isadd:", isadd, ",issub:", issub, ",ismulti:", ismulti, ",isremainder:", isremainder, ",loctionID:"+string(buff[ii-65:ii-1]), ",magic:"+string(buff[ii+10:i-1]))

			//直接跳到下一个{
			i = iii
		} else {
			i--
		}

	}

}
