package ginit

import (
	"database/sql"
)

var db *sql.DB

func InitMysql() error{
	var err error
	//BaseDB
	db , err = NewMysql("db")
	if err != nil{
		Scrlog.Error("DB error :%v", err)
		return err
	}
	//New Mysql
	//....
	return  nil
}

func NewMysql(DBName string) (*sql.DB, error ){
	var BaseDB string = Conf.DefaultString(DBName + "::user", "") + ":" +
		Conf.DefaultString(DBName + "::password", "") + "@tcp(" +
		Conf.DefaultString(DBName + "::host", "") + ":" +
		Conf.DefaultString(DBName + "::port", "3306") + ")/" +
		Conf.DefaultString(DBName + "::database", "") + "?charset=" +
		Conf.DefaultString(DBName + "+::charset", "utf8mb4")

	db, err := sql.Open("mysql", BaseDB)
	if err != nil {
		Scrlog.Error("dbsourcename=%v err :%v", BaseDB, err)
	}

	err = db.Ping()
	if err != nil {
		Scrlog.Error("dbsourcename=%v err :%v", BaseDB, err)
		return nil ,err
	}

	db.SetMaxIdleConns(20)
	db.SetMaxIdleConns(10)
	return db, nil
}
