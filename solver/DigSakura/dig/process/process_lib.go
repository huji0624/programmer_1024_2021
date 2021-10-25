package process

//————————————————————————————————————————————————————
func add99_table(input1, input2, inputconst byte) (output1, output2 byte) {
	output1 = input1 - 48 + input2 + inputconst
	output2 = 48
	if output1 > 57 {
		output1 = output1 - 10
		output2++
	}
	return output1, output2
}

func sub99_table(input1, input2, inputconst byte) (output1, output2 byte) {
	output1 = 58 + input1 - input2 - inputconst
	output2 = 48
	if output1 > 57 {
		output1 = output1 - 10
	} else {
		output2++
	}
	return output1, output2
}

func multi1024_table(input1, input2 byte) (output1, output2, output3, output4 byte) {
	if input1 == '0' {
		output1 = 48
		output2 = 48
		output3 = 48
		output4 = 48
	} else if input1 == '1' {
		output1 = 52
		output2 = 50
		output3 = 48
		output4 = 49
	} else if input1 == '2' {
		output1 = 56
		output2 = 52
		output3 = 48
		output4 = 50
	} else if input1 == '3' {
		output1 = 50
		output2 = 55
		output3 = 48
		output4 = 51
	} else if input1 == '4' {
		output1 = 54
		output2 = 57
		output3 = 48
		output4 = 52
	} else if input1 == '5' {
		output1 = 48
		output2 = 50
		output3 = 49
		output4 = 53
	} else if input1 == '6' {
		output1 = 52
		output2 = 52
		output3 = 49
		output4 = 54
	} else if input1 == '7' {
		output1 = 56
		output2 = 54
		output3 = 49
		output4 = 55
	} else if input1 == '8' {
		output1 = 50
		output2 = 57
		output3 = 49
		output4 = 56
	} else if input1 == '9' {
		output1 = 54
		output2 = 49
		output3 = 50
		output4 = 57
	}

	output4 = output4 + input2 - 48
	if output4 > 57 {
		output4 = output4 - 10
		output3++
		if output3 > 57 {
			output3 = output3 - 10
			output2++
		}
	}

	return output1, output2, output3, output4
}

func remainder1024_table_1(input1 byte) (output2, output3, output4 byte) {
	//output1 = 48
	output2 = 48
	output3 = 48
	output4 = 48
	if input1 == '4' {
		//output1 = 49
		output2 = 50
		output3 = 48
		output4 = 49
	} else if input1 == '8' {
		//output1 = 50
		output2 = 52
		output3 = 48
		output4 = 50
	} else if input1 == '2' {
		//output1 = 51
		output2 = 55
		output3 = 48
		output4 = 51
	} else if input1 == '6' {
		//output1 = 52
		output2 = 57
		output3 = 48
		output4 = 52
	}
	return output2, output3, output4
}

func remainder1024_table_2(input1 byte) (output2, output3, output4 byte) {
	//output1 = 53
	output2 = 50
	output3 = 49
	output4 = 53
	if input1 == '4' {
		//output1 = 54
		output2 = 52
		output3 = 49
		output4 = 54
	} else if input1 == '8' {
		//output1 = 55
		output2 = 54
		output3 = 49
		output4 = 55
	} else if input1 == '2' {
		//output1 = 56
		output2 = 57
		output3 = 49
		output4 = 56
	} else if input1 == '6' {
		//output1 = 57
		output2 = 49
		output3 = 50
		output4 = 57
	}
	return output2, output3, output4
}

//————————————————————————————————————————————————————
