package btime

import (
	"github.com/Qoiuy/gotool/basic/bstring"
	"strconv"
	"strings"
	"time"
)

// IsTimedOut 判断输入的时间是否超时
// (time.Now - t) > timeout 则超时
func IsTimedOut(t time.Time, timeout time.Duration) bool {
	return !t.Add(timeout).After(time.Now())
}

// Timestamp 返回时间戳
func Timestamp(t time.Time) int64 {
	result := t.Unix()
	if result > 0 {
		return result
	}
	return 0
}

// MinuteTimestamp 返回当前分钟数
func MinuteTimestamp(t time.Time) int64 {
	result := t.Unix()
	if result > 0 {
		return result / 60
	}
	return 0
}

// TimestampMS 返回时间戳（毫秒）
func TimestampMS(t time.Time) int64 {
	result := t.UnixNano() / 1000000
	if result > 0 {
		return result
	}
	return 0
}

// TimestampSStr 返回时间（秒）
func TimestampSStr() string {
	result := time.Now().UnixNano() / 1000000000
	if result > 0 {
		return bstring.Int2Str(result)
	}
	return "0"
}

// TimestampUS 返回时间戳（微秒）
func TimestampUS(t time.Time) int64 {
	result := t.UnixNano() / 1000
	if result > 0 {
		return result
	}
	return 0
}

// TimestampMS 返回时间戳（毫秒）
func TimestampNS(t time.Time) int64 {
	result := t.UnixNano()
	if result > 0 {
		return result
	}
	return 0
}

// NowNsStr 获取当前时间的时间戳（纳秒）字符串
func NowNsStr() string {
	return bstring.Int2Str(TimestampNS(time.Now()))
}

// 获取时间戳   return:1464059756
func GetTimeUnix(t time.Time) int64 {
	return t.Local().Unix()
}

// 获取当前时间的时间戳 return:1464059756
func GetNowTimeUnix() int64 {
	return GetTimeUnix(time.Now())
}

// 获取当前时间的时间戳 return:1464059756
func GetNowTimestampUS() int64 {
	return TimestampUS(time.Now())
}

func GetNowTimeUnixStr() string {
	return strconv.FormatInt(GetTimeUnix(time.Now()), 10)
}

// GetNonceStr 纳秒级别的不重复
func GetNonceStr() string {
	return strconv.FormatInt(time.Now().UnixNano(), 16)
}

// 获取当日晚上24点（次日0点）的时间
func Get24time(t time.Time) time.Time {
	dateStr := TimeToDate(t.Add(time.Hour * 24))
	return DateStrToTime(dateStr)
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetFirstDateOfWeek 获取本周周一的时间
func GetFirstDateOfWeek() time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}

// 获取当日晚上24点（次日0点）的时间戳
func Get24timeUnix(t time.Time) int64 {
	t24 := Get24time(t)
	return GetTimeUnix(t24)
}

// 获取今天晚上24点（次日0点）的时间
func GetToday24time() time.Time {
	return Get24time(time.Now())
}

// 获取今天晚上24点（次日0点）的时间戳
func GetToday24timeUnix() int64 {
	return Get24timeUnix(time.Now())
}

// 获取今天0点 的时间
func GetToday0time() time.Time {
	return today0()
}

// 获取今天0点 的时间戳，毫秒
func GetToday0timeMs() int64 {
	return TimestampMS(GetToday0time())
}

//GetTodayCertainTime 获取当天的指定时间
func GetTodayCertainTime(d time.Time, hour, min, sec, nsec int) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), hour, min, sec, nsec, d.Location())
}

// 时间转换成日期字符串 (time.Time to "2006-01-02")
func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// 时间转换成日期字符串 (time.Time to "200601")
func TimeToMonth(t time.Time) string {
	return t.Format("200601")
}

// 时间转换成日期字符串, 精确到分钟 (time.Time to "2006-01-02 15:04")
func TimeToMinuteStr(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

// 获取当前时间的日期字符串（"2006-01-02"）
func GetNowDateStr() string {
	return TimeToDate(time.Now())
}

// 获取当前时间的日期字符串（"200601"）
func GetNowMonthStr() string {
	return TimeToMonth(time.Now())
}

// 获取当前时间的日期字符串（"2006-01-02 15:04"）
func GetNowMinuteStr() string {
	return TimeToMinuteStr(time.Now())
}

// 时间转换成日期+时间字符串 (time.Time to "2006-01-02 15:04:05")
func TimeToDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 获取当前的时期+时间字符串
func GetNowStr() string {
	return TimeToDateTime(time.Now())
}

// 日期字符串转换成时间 ("2006-01-02" to time.Time)
func DateStrToTime(d string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02", d, time.Local)
	return t
}

// 日期字符串转换成时间 ("2006-01-02" to time.Time)
func DateStrToTimeErr(d string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", d, time.Local)
}

// 日期时间字符串转换成时间 ("2006-01-02 00:00:00" to time.Time)
func DateStrToDateTimeErr(d string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", d, time.Local)
}

// 日期字符串转换成时间 ("2020-09-21T16:00:00.000Z" to time.Time)
func DateStr1ToTime(d string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02T15:04:05.000Z", d, time.Local)
	return t
}

// 日期字符串转换成时间 ("2006-01-02T15:04:05Z/2006-01-02T15:04:05+07:00" to time.Time)
func DateStrRfc3339ToTime(d string) time.Time {
	t, _ := time.ParseInLocation(time.RFC3339, d, time.Local)
	return t
}

// 日期字符串转换成时间 自定义format
func DateStrToTimeFormatErr(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}

// 日期+时间字符串转换成时间 ("2006-01-02 15:04:05" to time.Time)
func DateTimeStrToTime(dt string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", dt, time.Local)
	return t
}

// 日期+时间字符串转换成时间 ("2006-01-02 15:04" to time.Time)
func DateTimeStr1ToTime(dt string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04", dt, time.Local)
	return t
}

func DateTimeStr2Time(dt string) time.Time {
	var arr []string
	if len(dt) > 0 && dt != ":" {
		arr = strings.Split(dt, ":")
	}

	switch len(arr) {
	case 1:
		return DateStrToTime(dt)
	case 2:
		return DateTimeStr1ToTime(dt)
	case 3:
		return DateTimeStrToTime(dt)
	default:
		return DateStrToTime(dt)
	}
}

// 时间字符串转换成时间 ("15:04:05" to time.Time)
func TimeTodayStrToTime(dt string) time.Time {
	now := time.Now()
	strNowDate := TimeToDate(now)
	return DateTimeStrToTime(strNowDate + " " + dt)
}

// 到今日结束剩余时间(秒)
func TodayEndSeconds() int {
	return int(GetToday24timeUnix() - GetNowTimeUnix())
}

// 是否是周末
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	if wd == time.Sunday || wd == time.Saturday {
		return true
	}
	return false
}

// Timestamp2Time 时间戳转时间
func Timestamp2Time(t int64) time.Time {
	if t > 1e10 {
		return time.Unix(t/1000, t%1000*1e6)
	}
	return time.Unix(t, 0)
}

// ConvertTimestampMs 秒级转毫秒级
func ConvertTimestampMs(t int64) int64 {
	if t < 1e10 {
		t = t * 1000
	}
	return t
}

//获取本月1号
func ThisMonthFirstDay() string {
	return TimeToDate(GetFirstDateOfMonth(time.Now()))
}

//获取本周一
func FirstDateOfWeek() string {
	return TimeToDate(GetFirstDateOfWeek())
}

//获取昨天
func DateOfYesterday() string {
	return TimeToDate(time.Now().AddDate(0, 0, -1))
}

//获取昨天
func DateOfLastYear() string {
	return TimeToDate(time.Now().AddDate(-1, 0, 0))
}

//获取昨天
func TimeOfLastYear() time.Time {
	return time.Now().AddDate(-1, 0, 0)
}

//获取昨天
func TimeOfToday() time.Time {
	return time.Now()
}

//获取今天0点
func DateOfToday() string {
	return TimeToDate(today0())
}

//获取昨天 2006-01-02 15:04
func MinuteOfYesterday() string {
	return TimeToMinuteStr(yesterday0())
}

//获取今天 2006-01-02 15:04
func MinuteOfToday() string {
	return TimeToMinuteStr(today0())
}

//获取昨天的 时间戳
func TimestampOfYesterday() int64 {
	return time.Now().AddDate(0, 0, -1).Unix()
}

//获取昨天的 时间戳
func TimestampOfYesterday0Time() int64 {
	return yesterday0().Unix()
}

// 今天0点Time
func today0() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 昨天0点Time
func yesterday0() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location())
}

// 获取今天0点 时间戳
func TimestampOfToday0Time() int64 {
	return today0().Unix()
}

//获取上周一
func FirstDateOfLastWeek() string {
	return TimeToDate(GetFirstDateOfWeek().AddDate(0, 0, -7))
}

//获取上周周日
func EndDateOfLastWeek() string {
	return TimeToDate(GetFirstDateOfWeek().AddDate(0, 0, -1))
}

//获取上月1号
func LastMonthFirstDay() string {
	return TimeToDate(GetLastMonth(time.Now()))
}

//获取本月1号
func LastMonthEndDay() string {
	return TimeToDate(GetLastMonthEndDay(time.Now()))
}

// GetLastMonth 获取传入的时间所在月份的 上一个月 第一天
func GetLastMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1).AddDate(0, -1, 0)
	return GetZeroTime(d)
}

// GetLastMonthEndDay 获取传入的时间所在月份的 上一个月 最后一天
func GetLastMonthEndDay(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day())
	return GetZeroTime(d)
}

// Timestamp2TimeStr 时间戳转格式化时间字符串
func Timestamp2TimeStr(t int64, format string) string {
	if t > 1e10 {
		t = t / 1000
	}
	return time.Unix(t, 0).Format(format)
}

// TimeSubDay 计算两个时间相差的天数
func TimeSubDay(timestamp1, timestamp2 int64) int64 {

	t1 := Timestamp2Time(timestamp1)
	t2 := Timestamp2Time(timestamp2)
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	dayDif := t1.Sub(t2).Hours() / 24
	if dayDif < 0 {
		dayDif = -dayDif
	}
	return int64(dayDif)
}

type DateScope struct {
	StartDate time.Time
	EndDate   time.Time
}

// 获取今年到现在 每个月的时间
func EveryMonthOfThisYear() []DateScope {
	now := time.Now()
	var allDateScope []DateScope
	//去年的
	for i := 1; i <= 12; i++ {
		allDateScope = append(allDateScope,
			DateScope{
				StartDate: time.Date(now.Year()-1, time.Month(i), 1, 0, 0, 0, 0, now.Location()),
				EndDate:   time.Date(now.Year()-1, time.Month(i+1), 0, 0, 0, 0, 0, now.Location()),
			})
	}

	// 今年的
	for i := 1; i <= int(now.Month()); i++ {
		// 能跨年 已测试
		allDateScope = append(allDateScope,
			DateScope{
				StartDate: time.Date(now.Year(), time.Month(i), 1, 0, 0, 0, 0, now.Location()),
				EndDate:   time.Date(now.Year(), time.Month(i+1), 0, 0, 0, 0, 0, now.Location()),
			})
	}
	return allDateScope
}

// 获取今年到现在 每个月的时间
func GetThisYearNMonthScope(n int) []DateScope {
	now := time.Now()
	var allDateScope []DateScope
	for i := 1; i <= n; i++ {
		// 能跨年 已测试
		allDateScope = append(allDateScope,
			DateScope{
				StartDate: time.Date(now.Year(), time.Month(i), 1, 0, 0, 0, 0, now.Location()),
				EndDate:   time.Date(now.Year(), time.Month(i+1), 0, 0, 0, 0, 0, now.Location()),
			})
	}
	return allDateScope
}

//获取从1970年到现在的第n周
func DateStrToWeek(times string) int {
	var t time.Time
	t = DateTimeStr2Time(times)
	unixtime := t.Unix()
	weektime := 7 * 24 * 3600
	week := int(unixtime) / weektime
	isweek := int(unixtime) % weektime

	if isweek != 0 {
		week++
	}
	return week
}

/**
*计算两个日前间的间隔天数
 */
func TimeSubDays(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {

			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24)
		} else { // more than 24 hours
			return int(hours/24) + 1
		}
	}

}
