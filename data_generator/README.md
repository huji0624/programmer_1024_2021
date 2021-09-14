# programmer_1024_2021

## 用于生成游戏数据的程序

游戏设定:

​	1.每一个数据文件包含N条数据，每条数据用\n换行符分割

​	2.每条数据的格式形如：{"locationid":"2ixewekdyvgbfmfzli9iifm1w9hnd2ij5kr1avy1zw3c7rl","magic":"9248642713483188"}

​	3.如果locationid代表的字符，去掉字母部分，剩下的数据部分，通过加/减/乘/取余 1024，恰好等于 magic 字符串所代表的数字，那么这个地点就表示存在宝藏。

​	例如 : {"locationid":"2mnab0kquw4uuu8nnm","magic":"2"} 抽取后的数字为2048，2048/1024 = 2 恰好等于magic代表的数字，所以这个地点就代表有宝藏.

​	4.找到宝藏地点后，需要把对应的locationid通过post请求，发送到我们的服务器，如果该宝藏还未被其他队伍找到，那么获得1分.

​	5.最后，得分最多的队伍获胜，如果有队伍得分一样，那么首先获得第1分的队伍获胜.