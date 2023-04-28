package snowflake

import (
	"fmt"
	"testing"
	"time"
)

/**
    Method：雪花计算方法,（1-漂移算法|2-传统算法），默认1
	BaseTime：基础时间（ms单位），不能超过当前系统时间
	WorkerId：机器码，必须由外部设定，最大值 2^WorkerIdBitLength-1
	WorkerIdBitLength：机器码位长，默认值6，取值范围 [1, 15]（要求：序列数位长+机器码位长不超过22）
	SeqBitLength：序列数位长，默认值6，取值范围 [3, 21]（要求：序列数位长+机器码位长不超过22）
	MaxSeqNumber：最大序列数（含），设置范围 [MinSeqNumber, 2^SeqBitLength-1]，默认值0，表示最大序列数取最大值（2^SeqBitLength-1]）
	MinSeqNumber：最小序列数（含），默认值5，取值范围 [5, MaxSeqNumber]，每毫秒的前5个序列数对应编号0-4是保留位，其中1-4是时间回拨相应预留位，0是手工新值预留位
	TopOverCostCount：最大漂移次数（含），默认2000，推荐范围500-10000（与计算能力有关）
*/
func TestSnowflake(t *testing.T) {
	// 创建对象，可在构造函数中输入 WorkerId：
	var options = NewSnowflakeOptions(1)
	options.Method = 1
	options.WorkerIdBitLength = 6
	options.SeqBitLength = 6
	// 保存参数（务必调用，否则参数设置不生效）：
	SetSnowflakeOptions(options)

	// 以上过程只需全局一次，且应在生成ID之前完成。

	// 初始化后，在任何需要生成ID的地方，调用以下方法：

	for {
		var newId = GenSnowflakeId()
		fmt.Println(newId)
		time.Sleep(time.Second)
	}
}

/**
使用默认配置生成
*/
func TestSnowflakeId(t *testing.T) {
	//返回字符串雪花算法ID
	var newStrId = GetSnowflakeId()

	fmt.Println(newStrId)

	//返回int64雪花算法ID
	newId := GenSnowflakeId()
	fmt.Println(newId)
}
