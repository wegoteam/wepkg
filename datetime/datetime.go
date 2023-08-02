package datetime

// 引用：https://github.com/golang-module/carbon
// 标准时间规则 2006-01-02 15:04:05.999999999
// Go的格式与yyyy-MM-dd HH:mm:ss格式的对应关系：
// 年：2006=yyyy 06=yy
// 月：01=MM 1=M Jan=MMM January=MMMM
// 日：02=dd 2=d
// 星期：Mon=EEE Monday=EEEE
// 小时：15=HH 03=KK 3=K
// 分钟：04=mm 4=m
// 秒：05=ss 5=s
// 上午/下午：PM=a pm无对应

import (
	timeUtil "github.com/golang-module/carbon/v2"
	"time"
)

const (
	DefaultDateTimePatternMilli  = "Y-m-d H:i:s"
	DefaultDateTimePatternMicro  = "Y-m-d H:i:s.u"
	DefaultDateTimePatternNano   = "Y-m-d H:i:s.U"
	DateTimeLayout               = "2006-01-02 15:04:05"
	DateTimeMilliLayout          = "2006-01-02 15:04:05.000"
	DateTimeMicroLayout          = "2006-01-02 15:04:05.000000"
	DateTimeNanoLayout           = "2006-01-02 15:04:05.000000000"
	PinyinDateTimeLayout         = "2006年01月02日 15时04分05秒"
	Pinyin2DateTimeLayout        = "2006年01月02日15时04分05秒"
	UnderlineDateTimeLayout      = "2006/01/02 15:04:05"
	UnderlineDateTimeMilliLayout = "2006/01/02 15:04:05.000"
	UnderlineDateTimeMicroLayout = "2006/01/02 15:04:05.000000"
	UnderlineDateTimeNanoLayout  = "2006/01/02 15:04:05.000000000"
)

// From
// @Description: 从time转换为Carbon
// @param: time
// @return timeUtil.Carbon
func From(data time.Time) timeUtil.Carbon {
	return timeUtil.FromStdTime(data)
}

// To
// @Description: 从Carbon转换为time
// @param: time
// @return time.Time
func To(data timeUtil.Carbon) time.Time {
	return data.ToStdTime()
}

// New
// @Description: 获取当前Carbon时间
// @return timeUtil.Carbon
func New() timeUtil.Carbon {
	return timeUtil.Now()
}

// As
// @Description: 字符串转Carbon
// @param: data
// @return timeUtil.Carbon
func As(data string) timeUtil.Carbon {
	return timeUtil.Parse(data)
}

// Now
// @Description: 获取当前时间
// @return time.Time
func Now() time.Time {
	return time.Now()
}

// Timestamp
// @Description: 获取当前秒级时间戳
func Timestamp() int64 {
	//return time.Now().Unix()
	return timeUtil.Now().Timestamp()
}

// TimestampMilli
// @Description: 获取当前毫秒级时间戳
// @return int64
func TimestampMilli() int64 {
	//return time.Now().UnixMilli()
	return timeUtil.Now().TimestampMilli()
}

// TimestampMicro
// @Description: 获取当前微秒级时间戳
// @return int64
func TimestampMicro() int64 {
	//return time.Now().UnixMicro()
	return timeUtil.Now().TimestampMicro()
}

// TimestampNano
// @Description: 获取当前纳秒级时间戳
// @return int64
func TimestampNano() int64 {
	//return time.Now().UnixNano()
	return timeUtil.Now().TimestampNano()
}

// Yesterday
// @Description: 获取昨天时间
// @return time.Time
func Yesterday() time.Time {
	return timeUtil.Yesterday().ToStdTime()
}

// Tomorrow
// @Description: 获取明天时间
// @return time.Time
func Tomorrow() time.Time {
	return timeUtil.Tomorrow().ToStdTime()
}

// Parse
// @Description: 字符串转time
// @param: data
// @return time.Time
func Parse(data string) time.Time {
	return timeUtil.Parse(data).ToStdTime()
}

// ToString
// @Description: time转字符串
// @param: data
// @return string
func ToString(data time.Time) string {
	from := From(data)
	return from.ToDateTimeString()
}

// Format
// @Description: time转字符串
// @param: data 时间
// @param: pattern 正则
// @return string
func Format(data time.Time, pattern string) string {
	from := From(data)
	return from.Format(pattern)
}

// Layout
// @Description: time转字符串
// @param: data 时间
// @param: pattern 正则
// @return string
func Layout(data time.Time, pattern string) string {
	from := From(data)
	return from.Layout(pattern)
}

// ChangeDays
// @Description: 修改天数
// @param: data 时间
// @param: days 天数
// @return time.Time
func ChangeDays(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddDays(num).ToStdTime()
	}
	return stdTime.SubDays(-num).ToStdTime()
}

// ChangeMonths
// @Description: 修改月数
// @param: data 时间
// @param: num 月数
// @return time.Time
func ChangeMonths(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddMonths(num).ToStdTime()
	}
	return stdTime.SubMonths(-num).ToStdTime()
}

// ChangeYears
// @Description: 修改年数
// @param: data 时间
// @param: num 年数
// @return time.Time
func ChangeYears(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddYears(num).ToStdTime()
	}
	return stdTime.SubYears(-num).ToStdTime()
}

// ChangeHours
// @Description: 修改小时数
// @param: data 时间
// @param: num 小时数
// @return time.Time
func ChangeHours(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddHours(num).ToStdTime()
	}
	return stdTime.SubHours(-num).ToStdTime()
}

// ChangeMinutes
// @Description: 修改分钟数
// @param: data 时间
// @param: num 分钟数
// @return time.Time
func ChangeMinutes(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddMinutes(num).ToStdTime()
	}
	return stdTime.SubMinutes(-num).ToStdTime()
}

// ChangeSeconds
// @Description: 修改秒数
// @param: data 时间
// @param: num 秒数
// @return time.Time
func ChangeSeconds(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddSeconds(num).ToStdTime()
	}
	return stdTime.SubSeconds(-num).ToStdTime()
}

// ChangeMilliseconds
// @Description: 修改毫秒数
// @param: data 时间
// @param: num 毫秒数
// @return time.Time
func ChangeMilliseconds(data time.Time, num int) time.Time {
	if num == 0 {
		return data
	}
	stdTime := timeUtil.FromStdTime(data)
	if num > 0 {
		return stdTime.AddMilliseconds(num).ToStdTime()
	}
	return stdTime.SubMilliseconds(-num).ToStdTime()
}

// DiffYear
// @Description: 计算两个时间相差年数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffYear(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInYears(timeUtil.FromStdTime(end))
}

// DiffAbsYear
// @Description: 计算两个时间相差年数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsYear(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInYears(timeUtil.FromStdTime(end))
}

// DiffMonth
// @Description: 计算两个时间相差月数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffMonth(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInMonths(timeUtil.FromStdTime(end))
}

// DiffAbsMonth
// @Description: 计算两个时间相差月数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsMonth(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInMonths(timeUtil.FromStdTime(end))
}

// DiffDay
// @Description: 计算两个时间相差天数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffDay(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInDays(timeUtil.FromStdTime(end))
}

// DiffAbsDay
// @Description: 计算两个时间相差天数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsDay(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInDays(timeUtil.FromStdTime(end))
}

// DiffHour
// @Description: 计算两个时间相差小时数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffHour(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInHours(timeUtil.FromStdTime(end))
}

// DiffAbsHour
// @Description: 计算两个时间相差小时数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsHour(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInHours(timeUtil.FromStdTime(end))
}

// DiffMinute
// @Description: 计算两个时间相差分钟数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffMinute(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInMinutes(timeUtil.FromStdTime(end))
}

// DiffAbsMinute
// @Description: 计算两个时间相差分钟数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsMinute(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInMinutes(timeUtil.FromStdTime(end))
}

// DiffSecond
// @Description: 计算两个时间相差秒数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffSecond(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInSeconds(timeUtil.FromStdTime(end))
}

// DiffAbsSecond
// @Description: 计算两个时间相差秒数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsSecond(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInSeconds(timeUtil.FromStdTime(end))
}

// DiffWeek
// @Description: 计算两个时间相差周数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffWeek(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffInWeeks(timeUtil.FromStdTime(end))
}

// DiffAbsWeek
// @Description: 计算两个时间相差周数
// @param: start 开始时间
// @param: end 结束时间
// @return int64
func DiffAbsWeek(start, end time.Time) int64 {
	return timeUtil.FromStdTime(start).DiffAbsInWeeks(timeUtil.FromStdTime(end))
}

// ToTimestamp
// @Description: 转换为时间戳
// @param: data 时间
// @return int64
func ToTimestamp(data time.Time) int64 {
	//return data.Unix()
	return timeUtil.FromStdTime(data).Timestamp()
}

// ToTimestampMilli
// @Description: 转换为时间戳(毫秒)
// @param: data
// @return int64
func ToTimestampMilli(data time.Time) int64 {
	//return data.UnixMilli()
	return timeUtil.FromStdTime(data).TimestampMilli()
}

// ToTimestampMicro
// @Description: 转换为时间戳(微秒)
// @param: data
// @return int64
func ToTimestampMicro(data time.Time) int64 {
	//return data.UnixMicro()
	return timeUtil.FromStdTime(data).TimestampMicro()
}

// ToTimestampNano
// @Description: 转换为时间戳(纳秒)
// @param: data
// @return int64
func ToTimestampNano(data time.Time) int64 {
	//return data.UnixNano()
	return timeUtil.FromStdTime(data).TimestampNano()
}

// IsEffective
// @Description: 判断时间是否有效
// @param: data 时间字符串
// @return bool
func IsEffective(data string) bool {
	return !timeUtil.Parse(data).IsInvalid()
}
