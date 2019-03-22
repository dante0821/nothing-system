package daos

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/redis.v5"
	"log"
	"time"
)

var (
	mysql *xorm.Engine
	rc    *redis.Client
)

// SetDriver 设置驱动
func SetMySql(e *xorm.Engine) {
	mysql = e
	go func() {
		for {
			mysql.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()
}

func SetRedis(_rc *redis.Client) {
	rc = _rc
}

func Transaction(fs ...func(s *xorm.Session) error) error {
	session := mysql.NewSession()
	session.Begin()
	for _, f := range fs {
		err := f(session)
		if err != nil {
			log.Println(err)
			session.Rollback()
			session.Close()
			return err
		}
	}
	session.Commit()
	session.Close()
	return nil
}

func TransactionInner(session *xorm.Session, fs ...func(s *xorm.Session) error) error {
	for _, f := range fs {
		err := f(session)
		if err != nil {
			return err
		}
	}
	return nil
}

func CloesMySQL() {
	mysql.Close()
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
