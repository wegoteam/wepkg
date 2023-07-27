package snowflake

import "strconv"

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
