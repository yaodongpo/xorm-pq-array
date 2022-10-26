package main

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	pqarray "xorm-pq-array"
)

var Db *xorm.Engine

func main() {
	db, err := xorm.NewEngine("postgres", "postgresql://postgres:postgres@127.0.0.1:5432/test?sslmode=disable")
	if err != nil {
		fmt.Println("NewEngine", err)
	}
	//2.显示sql语句
	db.ShowSQL(true)
	Db = db
	err = initTables()
	if err != nil {
		fmt.Println("initTables:", err)
		return
	}
	// TODO 手动添加点数据进去根据数据做查询
	err = GetInfo("ydp", "123456", 17673600227)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
func initTables() error {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS account_chain_data (localpart TEXT NOT NULL PRIMARY KEY, tel_numbers BIGINT [], blacklist TEXT [], whitelist TEXT [], address_book TEXT []);")
	if err != nil {
		return err
	}
	return nil
}

type UserNumbers struct {
	Localpart   string              `json:"localpart"`
	TelNumbers  pqarray.Int64Array  `json:"tel_numbers"`
	AddressBook pqarray.StringArray `json:"address_book"`
}

func GetInfo(from, to string, number int64) error {
	selectSql := "select localpart, tel_numbers, address_book from account_chain_data where localpart=?"
	usersNums := new(UserNumbers)
	get, err := Db.SQL(selectSql, from).Get(usersNums)
	if err != nil {
		return err
	}
	if !get {
		return errors.New("user not found")
	}
	return nil
}
