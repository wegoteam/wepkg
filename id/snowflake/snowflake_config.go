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

func ExtractTime(id int64) time.Time {
	return defaultSnowflake.ExtractTime(id)
}
