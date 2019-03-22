package daos

import (
	"MySystem/globals"
	"MySystem/models"
	"errors"
	"time"
)

func CreatePerson(m *models.PersonInfo) error {
	m.CreatedTime = time.Now()
	m.UpdatedTime = time.Now()
	aff, err := mysql.Insert(m)
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New("insert agent false")
	}
	return nil
}

func DeletePerson(m *models.PersonInfo, zeroUpdate bool) error {
	m.UpdatedTime = time.Now()
	m.IsDeleted = 1
	query := mysql.Id(m.PersonId)
	if zeroUpdate {
		query.AllCols()
	}
	aff, err := query.Update(m)
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New("delete fails")
	}
	return nil
}

func UpdatePerson(m *models.PersonInfo, zeroUpdate bool) error {
	m.UpdatedTime = time.Now()
	query := mysql.Id(m.PersonId)
	if zeroUpdate {
		query.AllCols()
	}
	aff, err := query.Update(m)
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New("请勿频繁操作")
	}
	return nil
}

func GetPersonList(filter map[string]interface{}, orderBy []string, limit int, start int) []*models.PersonInfo {
	list := make([]*models.PersonInfo, 0)
	query := mysql.Limit(limit, start)

	query.Where("is_deleted = 0 ")

	for k, v := range filter {
		query.Where(k+"=?", v)
	}
	for _, v := range orderBy {
		query.OrderBy(v)
	}
	err := query.Find(&list)
	globals.CheckErr(err)
	return list
}

func GetPersonCount(filter map[string]interface{}) int {
	m := &models.PersonInfo{}
	query := mysql.NewSession()

	query.Where("is_deleted = 0")

	for k, v := range filter {
		query.Where(k+"=?", v)
	}
	count, err := query.Count(m)
	globals.CheckErr(err)
	return int(count)
}

func GetPersonById(id interface{}) *models.PersonInfo {
	m := &models.PersonInfo{}

	has, err := mysql.Id(id).Where("is_deleted = 0").Get(m)

	globals.CheckErr(err)
	if has {
		return m
	} else {
		return nil
	}
}
