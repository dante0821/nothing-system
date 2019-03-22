package globals

import "time"

const (
	// TimeFormat 时间格式
	TimeFormat = "2006-01-02 15:04:03"
)

// FormatTime 时间格式化
func FormatTime(t time.Time) int64 {
	return t.Unix()
}
