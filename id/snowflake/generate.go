package snowflake

import (
	"strconv"
	"sync"
)

var once sync.Once

// init
// @Description: 初始化
func init() {
	once.Do(func() {
		snowflake, err := initSnowflake()
		if err != nil {
			return
		}
		var options = &SnowflakeOptions{
			Method:            uint16(snowflake.Method),
			WorkerId:          uint16(snowflake.WorkerId),
			BaseTime:          snowflake.BaseTime,
			WorkerIdBitLength: snowflake.BitLength,
			SeqBitLength:      snowflake.SeqBitLength,
			MaxSeqNumber:      0,
			MinSeqNumber:      5,
			TopOverCostCount:  2000,
		}
		SetSnowflakeOptions(options)
	})
}

// GenSnowflakeId 获取雪花算法ID
func GenSnowflakeId() int64 {
	if defaultSnowflake == nil {
		var options = NewSnowflakeOptions(1)
		// 保存参数（务必调用，否则参数设置不生效）：
		SetSnowflakeOptions(options)
	}
	return defaultSnowflake.NewNextId()
}

// GetSnowflakeId 获取雪花算法ID .
func GetSnowflakeId() string {
	if defaultSnowflake == nil {
		var options = NewSnowflakeOptions(1)
		// 保存参数（务必调用，否则参数设置不生效）：
		SetSnowflakeOptions(options)
	}
	newNextId := defaultSnowflake.NewNextId()
	return strconv.FormatInt(newNextId, 10)
}
