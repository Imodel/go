package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"webproject/models"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) GetUser() {
	beego.Info("test")

	id := u.Ctx.Input.Param(":id")
	beego.Info("id :", id)
	uid, err := strconv.Atoi(id)
	if err != nil {
		beego.Info("param error")
	}
	str := models.Get(uid)
	u.Data["json"] = str
	u.ServeJSON()
}

func (u *UserController) AddUser() {
	var user models.Userinfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	beego.Info(user.Username)
	id := models.Add(&user)
	u.Data["json"] = "{\"Id\":\"" + id + "\"}"
	u.ServeJSON()
}

func (u *UserController) ListUser() {
	user := models.List()
	u.Data["json"] = user
	u.ServeJSON()
}
