package snowflake

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wegoteam/wepkg/bean"
	"github.com/wegoteam/wepkg/config"
	"sync"
	"time"
)

// Snowflake
// @Description: 雪花算法
type Snowflake struct {
	Method       int   `json:"method"`       // 雪花计算方法,（1-漂移算法|2-传统算法），默认1
	BaseTime     int64 `json:"baseTime"`     // 基础时间（ms单位），不能超过当前系统时间
	WorkerId     int   `json:"workerId"`     // 机器码，必须由外部设定，最大值 2^WorkerIdBitLength-1
	BitLength    byte  `json:"bitLength"`    // 机器码位长，默认值6，取值范围 [1, 15]（要求：序列数位长+机器码位长不超过22）
	SeqBitLength byte  `json:"seqBitLength"` // 序列数位长，默认值6，取值范围 [3, 21]（要求：序列数位长+机器码位长不超过22）
}

var singletonMutex sync.Mutex
var defaultSnowflake *DefaultSnowflake

// SetSnowflakeOptions
// @Description: 设置雪花算法配置
// @param: options
func SetSnowflakeOptions(options *SnowflakeOptions) {
	singletonMutex.Lock()
	defaultSnowflake = NewDefaultSnowflake(options)
	singletonMutex.Unlock()
}

// ExtractTime
// @Description: 从ID中提取时间
// @param: id
// @return time.Time
func ExtractTime(id int64) time.Time {
	return defaultSnowflake.ExtractTime(id)
}

// initSnowflake
// @Description: 从配置文件初始化雪花算法
func initSnowflake() (*Snowflake, error) {
	c := config.GetConfig()
	var snowflake = &Snowflake{}
	err := c.Load("snowflake", snowflake)
	if err != nil {
		fmt.Errorf("Fatal error load snowflake config: %s \n", err)
		return snowflake, err
	}
	if bean.IsZero(snowflake) {
		fmt.Errorf("init snowflake error")
		return snowflake, errors.New("init snowflake error")
	}
	fmt.Printf("Init snowflake config: %v \n", snowflake)
	return snowflake, nil
}
