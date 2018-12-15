package utils

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// json 实现序列化方法
func (jt JsonTime) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%s\"", time.Time(jt).Format("2006-01-02 15:04:05"))
	return []byte(str), nil
}
