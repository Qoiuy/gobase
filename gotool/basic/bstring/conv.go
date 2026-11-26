package bstring

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Uint2Str 无符号整型到字符串
func Uint2Str(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// Str2Uint32 字符串到无符号整型
func Str2Uint32(s string) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(i)
}

// Str2Int32 字符串到符号整型
func Str2Int32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// Str2Uint64 字符串到无符号整型
func Str2Uint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Str2Uint64Err 字符串到无符号整型，并返回错误信息
func Str2Uint64Err(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Str2Float 字符串转浮点型
func Str2Float(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

// Float2Str 浮点型导致字符串
func Float2Str(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float2StrFmt 浮点型到字符串，digits 小数位数
// 示例：
//
//  Float2StrFmt(1.3564, 2)
//
// 输出 1.35
func Float2StrFmt(f float64, digits uint) string {
	format := fmt.Sprintf("%%.%df", digits)
	return fmt.Sprintf(format, f)
}

// I2Str 整型到字符串
func I2Str(n int) string {
	return strconv.Itoa(n)
}
func U2Str(n uint) string {
	return strconv.Itoa(int(n))
}

// Int2Str 整型到字符串
func Int2Str(n int64) string {
	return strconv.FormatInt(n, 10)
}

// Str2Int64 字符串到整型
func Str2Int64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Str2int64Err 字符串到整型，并返回错误信息
func Str2Int64Err(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Str2Int 字符串到整型
func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// Uint8s2Bytes uint8 -> byte
func Uint8s2Bytes(data []uint8) []byte {
	buf := make([]byte, 0, len(data))
	for _, v := range data {
		buf = append(buf, byte(v))
	}
	return buf
}

// FinanceFormat 金额格式化，float64加单位后保留2位小数
func FinanceFormat(value float64) string {
	if value < 0.01 && value > 0.000 {
		return "小于0.01"
	} else if value >= 10000 && value < 1e8 {
		value = value / 10000
		return Float2StrFmt(value, 2) + "万"
	} else if value >= 1e8 {
		value = value / 1e8
		return Float2StrFmt(value, 2) + "亿"
	}
	value = math.Trunc(value*1e2) * 1e-2
	s := fmt.Sprintf("%.2f", value)
	return s
}

// Float2StrPrefix float64转string，并添加判断数值过小
func Float2StrPrefix(value float64, digits uint) string {
	critical := 1 / (math.Pow(float64(10), float64(digits)))
	if value < critical && value > 0.00 {
		return "<" + Float2Str(critical)
	}
	format := fmt.Sprintf("%%.%df", digits)
	return fmt.Sprintf(format, value)
}

// ToUint64Arr string数组转uint64数组
func ToUint64Arr(strArr []string) []uint64 {
	uint64Arr := make([]uint64, 0, len(strArr))
	for _, str := range strArr {
		uint64Arr = append(uint64Arr, Str2Uint64(str))
	}
	return uint64Arr
}

// ToIntArr string数组转int数组
func ToIntArr(strArr []string) []int {
	uint64Arr := make([]int, 0, len(strArr))
	for _, str := range strArr {
		uint64Arr = append(uint64Arr, Str2Int(str))
	}
	return uint64Arr
}

// ToInt64Arr string数组转int64数组
func ToInt64Arr(strArr []string) []int64 {
	int64Arr := make([]int64, 0, len(strArr))
	for _, str := range strArr {
		int64Arr = append(int64Arr, Str2Int64(str))
	}
	return int64Arr
}

// ToStringArr int64数组转string数组
func ToStringArr(int64Arr []int64) []string {
	stringArr := make([]string, 0, len(int64Arr))
	for _, v := range int64Arr {
		stringArr = append(stringArr, Int2Str(v))
	}
	return stringArr
}

// UintArrToStringArr uint64数组转string数组
func UintArrToStringArr(uint64Arr []uint64) []string {
	stringArr := make([]string, 0, len(uint64Arr))
	for _, v := range uint64Arr {
		stringArr = append(stringArr, Uint2Str(v))
	}
	return stringArr
}

// IntArrToStringArr int数组转string数组
func IntArrToStringArr(intArr []int) []string {
	stringArr := make([]string, 0, len(intArr))
	for _, v := range intArr {
		stringArr = append(stringArr, I2Str(v))
	}
	return stringArr
}

// FloatArrToStringArr float64数组转string数组
func FloatArrToStringArr(floatArr []float64) []string {
	stringArr := make([]string, 0, len(floatArr))
	for _, v := range floatArr {
		stringArr = append(stringArr, Float2Str(v))
	}
	return stringArr
}

// StringArrToFloatArr string数组转float64数组
func StringArrToFloatArr(strArr []string) []float64 {
	floatArr := make([]float64, 0, len(strArr))
	for _, v := range strArr {
		floatArr = append(floatArr, Str2Float(v))
	}
	return floatArr
}

// StringArrToFloatArr string数组转float64数组
func StringArrToUint64Arr(strArr []string) []uint64 {
	intArr := make([]uint64, 0, len(strArr))
	for _, v := range strArr {
		intArr = append(intArr, Str2Uint64(v))
	}
	return intArr
}

// Str2PointString 字符串到 字符串指针  code generate 使用
func Str2PointString(s string) *string {
	return &s
}

func SplitStr2UintArr(s string) []uint64 {
	stringArr := Split(s, ",")
	res := []uint64{}
	for _, str := range stringArr {
		res = append(res, Str2Uint64(str))
	}
	return res
}

func RemoveSpacing(s string) string {
	return strings.Replace(s, " ", "", -1)
}

// GetStrRealLen 获取字符串长度，中英文均算一个长度
func GetStrRealLen(s string) int {
	return strings.Count(s, "") - 1
}

// Uint2StringArr 将一个 uint值转成一个 string数组
func Uint2StringArr(u uint) []string {
	return []string{Uint2Str(uint64(u))}
}

// intersectUint64 uint64 slice 求交集
func IntersectUint64(long, short []uint64) []uint64 {
	m := make(map[uint64]struct{})
	result := make([]uint64, 0)
	for _, v := range short {
		m[v] = struct{}{}
	}

	for _, v := range long {
		if _, found := m[v]; found {
			result = append(result, v)
		}
	}
	return result
}

func Split(s, sep string) []string {
	if len(s) > 0 && s != sep {
		return strings.Split(s, sep)
	}
	return []string{}
}

func JoinUint64Arr(arr []uint64, sep string) string {
	return strings.Join(UintArrToStringArr(arr), sep)
}

func JoinIntArr(arr []int, sep string) string {
	return strings.Join(IntArrToStringArr(arr), sep)
}
func JoinInt64Arr(arr []int64, sep string) string {
	return strings.Join(ToStringArr(arr), sep)
}

func ConvertByType(v string, t string) interface{} {
	var final interface{}
	switch t {
	case "int":
		final = Str2Int(v)
	case "int32":
		final = Str2Int32(v)
	case "int64":
		final = Str2Int64(v)
	case "uint32":
		final = Str2Uint32(v)
	case "uint64":
		final = Str2Uint64(v)
	case "float32":
		final = Str2Float(v)
	case "float64":
		final = Str2Float(v)
	case "string":
		final = v
	}
	return final
}
