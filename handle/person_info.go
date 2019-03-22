package handle

import (
	"MySystem/daos"
	"MySystem/globals"
	"MySystem/models"
	"MySystem/vm"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Rsp struct {
	Msg string `json:"msg"`
}

func CreatePerson(c echo.Context) error {
	req := &vm.CreateOrUpdatePersonInfo{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("请求体错误", err, http.StatusNotAcceptable))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, globals.NewErrorRsp("请求校验失败", err, http.StatusUnprocessableEntity))
	}

	m := &models.PersonInfo{
		PersonName: req.PersonName,
		PlatformId: c.Get("id").(int),
		Sex:        req.Sex,
		Age:        req.Age,
		IdCard:     req.IdCard,
		Address:    req.Address,
		Party:      req.Party,
		Phone:      req.Phone,
		Birthday:   req.Birthday,
	}
	err := daos.CreatePerson(m)
	if err != nil {
		return c.JSON(http.StatusBadGateway, globals.NewErrorRsp("创建失败", err, http.StatusBadGateway))
	}
	rsp := &globals.ResponseInfo{
		ReturnCode: http.StatusOK,
	}
	return c.JSON(http.StatusOK, rsp)
}

func DeletePerson(c echo.Context) error {
	idParam := c.Param("id")
	id, iderr := strconv.Atoi(idParam)
	if iderr != nil {
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("请求参数错误", nil, http.StatusNotAcceptable))
	}
	zeroUpdate, _ := strconv.ParseBool(c.QueryParam("cols"))
	m := daos.GetPersonById(id)
	if m == nil {
		return c.JSON(http.StatusNotFound, globals.NewErrorRsp("数据不存在", nil, http.StatusNotFound))
	}
	err := daos.DeletePerson(m, zeroUpdate)
	if err != nil {
		return c.JSON(http.StatusBadGateway, globals.NewErrorRsp("更新失败", err, http.StatusBadGateway))
	}
	rsp := &globals.ResponseInfo{
		ReturnCode: http.StatusOK,
	}
	return c.JSON(http.StatusOK, rsp)
}

func UpdatePerson(c echo.Context) error {
	idParam := c.Param("id")
	id, iderr := strconv.Atoi(idParam)
	if iderr != nil {
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("请求参数错误", nil, http.StatusNotAcceptable))
	}
	zeroUpdate, _ := strconv.ParseBool(c.QueryParam("cols"))
	req := &vm.CreateOrUpdatePersonInfo{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("请求体错误", err, http.StatusNotAcceptable))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, globals.NewErrorRsp("请求校验失败", err, http.StatusUnprocessableEntity))
	}
	m := daos.GetPersonById(id)
	if m == nil {
		return c.JSON(http.StatusNotFound, globals.NewErrorRsp("数据不存在", nil, http.StatusNotFound))
	}
	m.PersonName = req.PersonName
	m.Sex = req.Sex
	m.Age = req.Age
	m.IdCard = req.IdCard
	m.Address = req.Address
	m.Party = req.Party
	m.Phone = req.Phone
	m.Birthday = req.Birthday

	err := daos.UpdatePerson(m, zeroUpdate)
	if err != nil {
		return c.JSON(http.StatusBadGateway, globals.NewErrorRsp("更新失败", err, http.StatusBadGateway))
	}
	rsp := &globals.ResponseInfo{
		ReturnCode: http.StatusOK,
	}
	return c.JSON(http.StatusOK, rsp)
}

func GetPerson(c echo.Context) error {
	idParam := c.Param("id")
	id, iderr := strconv.Atoi(idParam)
	if iderr != nil {
		return c.JSON(http.StatusNotAcceptable, globals.NewErrorRsp("请求参数错误", nil, http.StatusNotAcceptable))
	}
	m := daos.GetPersonById(id)
	if m == nil {
		return c.JSON(http.StatusNotFound, globals.NewErrorRsp("数据不存在", nil, http.StatusNotFound))
	}
	info := &vm.PersonInfo{
		PersonId:    m.PersonId,
		PersonName:  m.PersonName,
		Sex:         m.Sex,
		Age:         m.Age,
		IdCard:      m.IdCard,
		Address:     m.Address,
		Party:       m.Party,
		Phone:       m.Phone,
		Birthday:    m.Birthday,
		CreatedTime: globals.FormatTime(m.CreatedTime),
		UpdatedTime: globals.FormatTime(m.UpdatedTime),
	}
	rsp := &globals.ResponseInfo{
		Data:       info,
		ReturnCode: http.StatusOK,
	}
	return c.JSON(http.StatusOK, rsp)
}

func GetPersonList(c echo.Context) error {
	filters := globals.CreateFilterMap()
	orderby := globals.CreateOrderSlice()
	limit, start := globals.GetLimitAndStart(c)
	filters["platform_id"] = c.Get("id").(int)
	list := daos.GetPersonList(filters, orderby, limit, start)
	info := &vm.GetPersonInfoListRsp{}
	info.List = make([]vm.PersonInfo, 0, len(list))
	for _, m := range list {
		item := vm.PersonInfo{
			PersonId:    m.PersonId,
			PersonName:  m.PersonName,
			Sex:         m.Sex,
			Age:         m.Age,
			IdCard:      m.IdCard,
			Address:     m.Address,
			Party:       m.Party,
			Phone:       m.Phone,
			Birthday:    m.Birthday,
			CreatedTime: globals.FormatTime(m.CreatedTime),
			UpdatedTime: globals.FormatTime(m.UpdatedTime),
		}
		info.List = append(info.List, item)
	}
	info.Count = daos.GetPersonCount(filters)

	rsp := &globals.ResponseInfo{
		Data:       info,
		ReturnCode: http.StatusOK,
	}
	return c.JSON(http.StatusOK, rsp)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func ImportPersonInfoByExcel(c echo.Context) error {
	//文件地址
	path := c.FormValue("path")
	if path == "" {
		return c.JSON(http.StatusBadRequest, globals.NewErrorRsp("请上传excel文件", nil, http.StatusBadGateway))
	}
	path = strings.TrimLeft(path, "/")
	if flag, _ := exists(path); !flag {
		return c.JSON(http.StatusBadRequest, globals.NewErrorRsp("未找到excel文件", nil, http.StatusBadGateway))
	}
	//path := "files/excel/personInfo_template_test.xlsx"
	//excel表
	sheetName := "Sheet1"
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, globals.NewErrorRsp("打开excel失败", err, http.StatusBadGateway))
	}

	//sheet名称
	rows := xlsx.GetRows(sheetName)

	var personInfo *models.PersonInfo
	var errorMap = map[int]interface{}{}
	var errorMsg []string

	//模板错误或无有效数据
	if len(rows) <= 2 || len(rows[1]) != 7 {
		return c.JSON(http.StatusBadRequest, globals.NewErrorRsp("excel模板错误或无有效数据", nil, http.StatusBadGateway))
	}

	for rIndex, row := range rows {
		//标题行
		if rIndex != 0 {
			i := 0

			errorMsg = []string{}
			personInfo = new(models.PersonInfo)
			//姓名
			personInfo.PersonName = Trim(row[i])
			if personInfo.PersonName == "" {
				errorMsg = append(errorMsg, "姓名不能为空")
			} else if len(personInfo.PersonName) >= 20 {
				errorMsg = append(errorMsg, "姓名字符过长")
			}
			i++
			//性别
			personInfo.Sex = Trim(row[i])
			i++
			//年龄
			age, err := strconv.Atoi(Trim(row[i]))
			if err == nil {
				personInfo.Age = age
			}
			i++
			//出生日期
			personInfo.Birthday = Trim(row[i])
			i++
			//手机号
			personInfo.Phone = Trim(row[i])
			i++
			//身份证
			personInfo.IdCard = Trim(row[i])
			if personInfo.IdCard != "" {
				if len(personInfo.IdCard) != 15 && len(personInfo.IdCard) != 18 {
					errorMsg = append(errorMsg, "身份证格式错误仅，支持15、18位")
				}
			}
			i++
			//家庭地址
			personInfo.Address = Trim(row[i])
			i++
			if len(personInfo.Address) >= 255 {
				errorMsg = append(errorMsg, "家庭地址字符过长")
			}

		}

		if len(errorMsg) > 0 {
			//错误记录
			errorMap[rIndex] = errorMsg
		} else {
			//添加的默认数据
			personInfo.PlatformId = c.Get("id").(int)

			//数据保存
			if err := daos.CreatePerson(personInfo); err != nil {
				//保存失败错误处理
				errorMsg = append(errorMsg, err.Error())
				errorMap[rIndex] = errorMsg
			}
		}
	}
	rsp := &globals.ResponseInfo{
		Data:       errorMap,
		ReturnCode: http.StatusOK,
	}
	return c.JSON(http.StatusOK, rsp)
}

//去除前后所有空格、空字符串、制表符
func Trim(str string) string {
	if str == "" {
		return ""
	}
	return Trim(strings.TrimPrefix(str, string('\uFEFF')))
}

//已处理数字型日期
func ParseExcelDate(date string) (d *time.Time, err error) {
	if date != "" {
		var date2 time.Time
		date2, err = ParseDate(date)
		if err != nil {
			return
		}
		d = &date2
	}
	return
}

//字符串日期转换
func ParseDate(date string) (time.Time, error) {
	date = strings.Replace(date, "/", "-", -1)
	date = strings.Replace(date, ".", "-", -1)
	date = strings.Replace(date, "-0", "-", -1)
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation("2006-1-2", date, local)
}
