package globals

import (
	"github.com/labstack/echo"
	"strconv"
)

// GetLimitAndStart 获取开始和结束
func GetLimitAndStart(c echo.Context) (int, int) {
	start, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		start = 1
	}
	if start < 1 {
		start = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	if limit < 0 || limit > 1000 {
		limit = 20
	}

	return limit, (start - 1) * limit
}

func GetLimitAndStartNew(page int, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 0 || limit > 1000 {
		limit = 20
	}
	return limit, (page - 1) * limit
}

// CreateFilterMap 创建过滤字段切片
func CreateFilterMap() map[string]interface{} {
	return make(map[string]interface{}, 0)
}

// CreateStringFilterMap 创建过滤字段切片
func CreateStringFilterMap() map[string]string {
	return make(map[string]string, 0)
}

// CreateOrderSlice 创建排序字段切片
func CreateOrderSlice() []string {
	return make([]string, 0)
}
