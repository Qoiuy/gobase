package uuid

import (
	"github.com/Qoiuy/gotool/basic/btime"
	"sort"
)

// GenIdByTime 根据时间戳生成id
func GenIdByTime(t int64) (uint64, error) {
	// 转成毫秒级
	t = btime.ConvertTimestampMs(t)
	// 生成节点实例
	node, err := NewNode(t, RandNum(0, nodeMax+1), RandNum(0, stepMask+1))
	if err != nil {
		return 0, err
	}
	return uint64(node.Generate()), nil
}

func GenUintIdByTime(t int64) (uint, error) {
	id, err := GenIdByTime(t)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// GenUintSequentialIds 经过测试 当num = 4929时，出现重复Id 具体测试见 Test_Duplicate_Id
func GenUintSequentialIds(t int64, num int) ([]uint, error) {

	Ids, err := genUintSequentialIds(t, num)
	if err != nil {
		return nil, err
	}

	// 找到重复的数据
	tmp, errIdx := make(map[uint]bool, 0), make([]int, 0)
	for i := range Ids {
		if _, ok := tmp[Ids[i]]; !ok {
			tmp[Ids[i]] = true
		} else {
			errIdx = append(errIdx, i)
		}
	}

	// 去除重复的数据，对数组进行补充
	for i := 0; i < len(errIdx); {
		newId, err1 := genUintSequentialIds(t, 1) //不差这点效率
		if err1 != nil {
			return nil, err1
		}
		if _, ok := tmp[newId[0]]; !ok {
			tmp[newId[0]] = true
			// 替换掉原始数组数据
			Ids[errIdx[i]] = newId[0]
			i++
		}
	}

	return Ids, err
}

// genUintSequentialIds 根据时间戳生成num个id 并排序
func genUintSequentialIds(t int64, num int) ([]uint, error) {
	res := make([]uint, num)
	for i := 0; i < num; i++ {
		id, err := GenIdByTime(t)
		if err != nil {
			return nil, err
		}
		res[i] = uint(id)
	}

	// 对结果进行排序
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res, nil
}
