package handle

import (
	"strings"
	"time"
	"github.com/labstack/echo"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
	"os"
)

type Rsp struct {
	Msg string `json:"msg"`
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func ImportAccountByExcel(c echo.Context) error {
	//文件地址
	path := c.FormValue("path")
	if path == "" {
		return c.JSON(http.StatusBadRequest, &Rsp{"请上传excel文件"})
	}
	path = strings.TrimLeft(path, "/")
	if flag, _ := exists(path); !flag {
		return c.JSON(http.StatusBadRequest, &Rsp{"未找到excel文件"})
	}
	//path := "files/excel/account_template_test.xlsx"
	//excel表
	sheetName := "Sheet1"
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Rsp{fmt.Sprintf("打开excel失败,error：%s", err.Error())})
	}

	//sheet名称
	rows := xlsx.GetRows(sheetName)

	var ip = c.RealIP()
	var now = time.Now()
	var partyId int64 = 1
	var account *model.Account
	var accountInfo *model.AccountInfo
	var dateStr, tempStr string
	var errorMap = map[int]interface{}{}
	var errorMsg []string
	var org model.PartyOrg
	var partyPost model.PartyPost
	//模板错误
	if len(rows) <= 2 {
		return utils.ErrorNull(c, "excel格式错误或无有效数据")
	}
	//无有效数据
	if len(rows[2]) < 25 {
		return utils.ErrorNull(c, "excel格式错误或无有效数据")
	}
	for rIndex, row := range rows {
		//跳过提示行、标题行
		if rIndex != 0 && rIndex != 1 {
			errorMsg = []string{}
			accountInfo = new(model.AccountInfo)
			account = new(model.Account)
			//姓名
			account.FullName = utils.Trim(row[0])
			if account.FullName == "" {
				errorMsg = append(errorMsg, "姓名不能为空")
			} else if len(account.FullName) >= 50 {
				errorMsg = append(errorMsg, "姓名字符过长")
			}

			//性别
			account.Gender = utils.Trim(row[1])
			if account.Gender == "" {
				errorMsg = append(errorMsg, "性别为必填项")
			} else if account.Gender != enum.GENDER_MALE && account.Gender != enum.GENDER_FEMALE {
				errorMsg = append(errorMsg, "性别错误，值范围：男、女")
			}

			//手机号
			account.Mobile = utils.Trim(row[2])
			if account.Mobile == "" {
				errorMsg = append(errorMsg, "手机号码为必填项")
			} else if !utils.IsMobile(account.Mobile) {
				errorMsg = append(errorMsg, "手机号码格式错误")
			}
			account.Name = account.Mobile

			//出生日期
			dateStr = utils.Trim(row[3])
			account.DateOfBirth, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "出生日期格式错误")
			}

			//在岗状态
			accountInfo.WorkStatus = utils.Trim(row[4])
			switch accountInfo.WorkStatus {
			case "在岗", "待聘人员", "农民工", "停薪留职", "排休人员", "离退休", "其他", "":
				break
			default:
				errorMsg = append(errorMsg, "在岗状态错误，值范围：在岗, 待聘人员, 农民工, 停薪留职, 排休人员, 离退休, 其他")
				break
			}

			//民族
			accountInfo.Nation = utils.Trim(row[5])
			if len(accountInfo.Nation) > 50 {
				errorMsg = append(errorMsg, "民族字符串过长")
			}

			//籍贯
			accountInfo.NativePlace = utils.Trim(row[6])
			if len(accountInfo.Nation) > 100 {
				errorMsg = append(errorMsg, "籍贯字符串过长")
			}

			//身份证
			accountInfo.Idcard = utils.Trim(row[7])

			if accountInfo.Idcard != "" {
				if len(accountInfo.Idcard) != 15 && len(accountInfo.Idcard) != 18 {
					errorMsg = append(errorMsg, "身份证格式错误仅，支持15、18位")
				}
			}
			//学历
			accountInfo.Education = utils.Trim(row[8])
			switch accountInfo.Education {
			case "博士", "硕士", "本科", "专科", "高中及以下", "":
				break
			default:
				errorMsg = append(errorMsg, "学历错误，值范围：博士,硕士,本科,专科,高中及以下")
				break
			}

			//人员类型
			accountInfo.PersonnelType = utils.Trim(row[9])
			if accountInfo.PersonnelType != "" {
				switch accountInfo.PersonnelType {
				case "正式党员", "预备党员", "":
					break
				default:
					errorMsg = append(errorMsg, "学历错误，值范围：正式党员,预备党员")
					break
				}
			}

			//党支部
			tempStr = utils.Trim(row[10])
			if tempStr != "" {
				org, err = GetPartyOrgByName(partyId, tempStr)
				if err != nil {
					errorMsg = append(errorMsg, "党支部不存在")
				} else {
					account.OrgId = org.ID
				}
			} else {
				errorMsg = append(errorMsg, "党支部为必填项")
			}

			//党内职务
			tempStr = utils.Trim(row[11])
			if tempStr != "" {
				partyPost, err = GetPartyPostByName(partyId, tempStr)
				if err != nil {
					errorMsg = append(errorMsg, "党内职务不存在")
				} else {
					accountInfo.PartyPostId = partyPost.ID
				}
			}

			//转为预备党员日期
			dateStr = utils.Trim(row[12])
			accountInfo.TurnPreparePartyDate, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "转为预备党员日期格式错误")
			}

			//转为正式党员日期
			dateStr = utils.Trim(row[13])
			accountInfo.TurnPartyDate, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "转为正式党员日期格式错误")
			}

			//工作岗位
			accountInfo.Post = utils.Trim(row[14])
			if len(accountInfo.Post) >= 50 {
				errorMsg = append(errorMsg, "工作岗位字符过长")
			}

			//税后工资
			tempStr = utils.Trim(row[15])
			if tempStr != "" {
				accountInfo.AfterTaxWages, err = convert.ToFloat64(utils.Trim(row[15]))
				if err != nil || accountInfo.AfterTaxWages < 0 {
					errorMsg = append(errorMsg, "税后工资格式错误")
				}
			}
			//固定电话
			accountInfo.Phone = utils.Trim(row[16])
			if len(accountInfo.Phone) >= 30 {
				errorMsg = append(errorMsg, "固定电话字符过长")
			}

			//家庭地址
			accountInfo.HomeAddress = utils.Trim(row[17])
			if len(accountInfo.HomeAddress) >= 255 {
				errorMsg = append(errorMsg, "家庭地址字符过长")
			}

			//党籍状态
			accountInfo.PartyStatus = utils.Trim(row[18])
			switch accountInfo.PartyStatus {
			case "正常", "异常", "":
				break
			default:
				errorMsg = append(errorMsg, "党籍状态错误，值范围：正常、异常")
				break
			}

			//是否为失联党员
			accountInfo.PartyLostStatus = utils.Trim(row[19])
			switch accountInfo.PartyLostStatus {
			case "是", "否", "":
				break
			default:
				errorMsg = append(errorMsg, "是否为失联党员错误，值范围：是、否")
				break
			}

			//失去联系的日期
			dateStr = utils.Trim(row[20])
			accountInfo.PartyLostDate, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "失去联系的日期格式错误")
			}

			//是否为流动党员
			accountInfo.PartyFlowStatus = utils.Trim(row[21])
			switch accountInfo.PartyFlowStatus {
			case "是", "否", "":
				break
			default:
				errorMsg = append(errorMsg, "是否为流动党员错误，值范围：是、否")
				break
			}

			//外出流向
			accountInfo.OutgoingFlow = utils.Trim(row[22])
			if len(accountInfo.OutgoingFlow) >= 500 {
				errorMsg = append(errorMsg, "外出流向字符过长")
			}

			//申请入党日期
			dateStr = utils.Trim(row[23])
			accountInfo.PartyApplyDate, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "申请入党日期格式错误")
			}

			//列为积极分子日期
			dateStr = utils.Trim(row[24])
			accountInfo.PartyActivistDate, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "列为积极分子日期格式错误")
			}

			//列为发展对象日期
			dateStr = utils.Trim(row[25])
			accountInfo.PartyDevelopDate, err = utils.ParseExcelDate(dateStr)
			if err != nil {
				errorMsg = append(errorMsg, "列为发展对象日期格式错误")
			}

			//判断手机号码是否存在
			acc, _ := GetAccountByMobile(account.Mobile)
			if acc.ID > 0 {
				errorMsg = append(errorMsg, "手机号码已存在")
			}

			if len(errorMsg) > 0 {
				//错误记录
				xlsx.SetCellDefault(sheetName, fmt.Sprintf("AA%v", rIndex+1), strings.Join(errorMsg, ";\r\n"))
				errorMap[rIndex] = errorMsg
			} else {
				//添加的默认数据
				account.ID = utils.ID()
				account.Status = enum.NORMAL
				account.CTime = &now
				account.UTime = account.CTime
				account.PartyId = partyId
				account.Ip = ip

				accountInfo.ID = utils.ID()
				accountInfo.AccountId = account.ID

				//数据保存
				if err := saveImportAccount(account, accountInfo); err != nil {
					//保存失败错误处理
					errorMsg = append(errorMsg, err.Error())
					xlsx.SetCellDefault(sheetName, fmt.Sprintf("AA%v", rIndex+1), strings.Join(errorMsg, ";\r\n"))
					errorMap[rIndex] = errorMsg
				}
			}
			//如果有错误，将背景设为警示颜色
			if len(errorMsg) > 0 {
				xlsx.SetCellStyle(sheetName, fmt.Sprintf("A%v", rIndex+1), fmt.Sprintf("AA%v", rIndex+1), style)
			}
			fmt.Println("-------------------------------------------------------------------------------------------")
		}
	}

	if len(errorMap) > 0 {
		//固定的标题栏位置
		xlsx.SetCellDefault(sheetName, "AA2", "错误说明")
		xlsx.Save()
		return utils.Confirm(c, "导入数据异常，请下载excel根据最后一列的错误说明进行修改调整", fmt.Sprintf("%s", path))
	}
	//需要自己处理返回
	return utils.SuccessNull(c, "导入成功")
}

func saveImportAccount(account *model.Account, accountInfo *model.AccountInfo) error {
	//保存事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		global.Log.Error("tx error：%v", tx.Error.Error())
		return errors.New(fmt.Sprintf("初始化事务失败：%s", tx.Error.Error()))
	}
	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback()
		global.Log.Error("tx create account error：%v", err)
		return errors.New(fmt.Sprintf("导入党员失败：%s", err.Error()))
	}
	if err := tx.Create(&accountInfo).Error; err != nil {
		tx.Rollback()
		global.Log.Error("tx create accountInfo error：%v", err)
		return errors.New(fmt.Sprintf("导入党员详细信息失败：%s", err.Error()))
	}

	if err := tx.Commit().Error; err != nil {
		global.Log.Error("commit error：%v", err)
		return errors.New(fmt.Sprintf("保存党员失败：%s", err.Error()))
	}
	return nil
}

//去除前后所有空格、空字符串、制表符
func Trim(str string) string {
	if str == "" {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(str, string('\uFEFF')))
}

//已处理数字型日期
func ParseExcelDate(date string) (d *time.Time, err error) {
	if date != "" {
		var date2 time.Time
		if !IsValidNumber(date) {
			date2, err = ParseDate(date)
			if err != nil {
				return
			}
			d = &date2
			return
		} else {
			date2, err = ParseDate("1900-1-1")
			if err != nil {
				return
			}
			days := convert.MustInt(date)
			date2 = date2.AddDate(0, 0, days-2)
			d = &date2
			return
		}
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
