package snowflake

import (
	"sync"
	"time"
)

var singletonMutex sync.Mutex
var defaultSnowflake *DefaultSnowflake

// SetSnowflakeOptions .
func SetSnowflakeOptions(options *SnowflakeOptions) {
	singletonMutex.Lock()
	defaultSnowflake = NewDefaultSnowflake(options)
	singletonMutex.Unlock()
}

// GenSnowflakeId .
func GenSnowflakeId() int64 {
	if defaultSnowflake == nil {
		panic("Please Initialize SnowflakeOptions")
	}
	return defaultSnowflake.NewNextId()
}

func ExtractTime(id int64) time.Time {
	return defaultSnowflake.ExtractTime(id)
}
