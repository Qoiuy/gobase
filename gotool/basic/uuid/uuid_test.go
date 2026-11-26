package uuid

import (
	"github.com/Qoiuy/gotool/basic/btime"
	"testing"
	"time"
)

func TestGenUintSequentialIds(t *testing.T) {
	num := 10
	uuidSequentialList, err := GenUintSequentialIds(btime.TimestampNS(time.Now()), num)

	if err != nil {
		t.Errorf("GenUintSequentialIds() error = %v", err)
		return
	}
	t.Logf("GenUintSequentialIds() = %v", uuidSequentialList)

}

func Test_Duplicate_Id(t *testing.T) {
	j := 4929
	jsonIds, err := genUintSequentialIds(btime.TimestampNS(time.Now()), j)
	if err != nil {
		t.Errorf("db import_pdf_to_text_info 结构体 主键uuid生成失败，p err:%+v", err)
	}
	t.Logf("判断一下Id是否有重复的:")
	tmp := make(map[uint]bool, 0)
	for i, _ := range jsonIds {
		// 如果tmp中已经存在 jsonIds[i] 打印错误

		if _, ok := tmp[jsonIds[i]]; !ok {
			tmp[jsonIds[i]] = true
		} else {
			t.Errorf("出现问题了， 生成的主键Id的确有重复的。id:%+v", jsonIds[i])
		}
	}
}

func Test_Duplicate_Id_New_func(t *testing.T) {
	j := 4929
	jsonIds, err := GenUintSequentialIds(btime.TimestampNS(time.Now()), j)
	if err != nil {
		t.Errorf("db import_pdf_to_text_info 结构体 主键uuid生成失败，p err:%+v", err)
	}
	t.Logf("判断一下Id是否有重复的:")
	tmp := make(map[uint]bool, 0)
	for i, _ := range jsonIds {
		// 如果tmp中已经存在 jsonIds[i] 打印错误

		if _, ok := tmp[jsonIds[i]]; !ok {
			tmp[jsonIds[i]] = true
		} else {
			t.Errorf("出现问题了， 生成的主键Id的确有重复的。id:%+v", jsonIds[i])
		}
	}
}
