package models

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type Userinfo struct {
	Uid        int `orm:"pk"`
	Username   string
	Departname string
	Created    string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)
	orm.RegisterModel(new(Userinfo))
}

func Add(u *Userinfo) (str string) {
	o := orm.NewOrm()
	uid, err := o.Insert(u)
	checkErr(err)
	return string(uid)
}

func Get(id int) (name string) {
	o := orm.NewOrm()
	user := Userinfo{Uid: id}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Uid, user.Username)
		return user.Username
	}

	return "notFound"
}

func List() (list []*Userinfo) {

	var user []*Userinfo
	o := orm.NewOrm()
	qs := o.QueryTable("userinfo")
	qs.Filter("username", "liuxk")
	num, err := qs.All(&user)
	checkErr(err)
	fmt.Println("num is ", num)
	return user

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
