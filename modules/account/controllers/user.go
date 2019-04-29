package controllers

import (
	"app/core"
	"app/modules/account/models"

	"github.com/astaxie/beego"
)

// UserController ...
type UserController struct {
	beego.Controller
}

// Index 用户列表
func (c *UserController) Index() {

	var users []models.User

	core.DB.Find(&users)

	c.Data["json"] = users
	c.ServeJSON()
}

// Store 增加用户
func (c *UserController) Store() {

	name := c.GetString("name")
	password := c.GetString("password")

	if name == "" {
		c.Ctx.WriteString("name is empty")
		return
	}

	if password == "" {
		c.Ctx.WriteString("password is empty")
		return
	}

	user := models.User{
		Name:     name,
		Password: password,
	}

	if err := core.DB.Create(&user).Error; err != nil {
		// 创建失败 ...
		return
	}

	// 暂时返回匿名结构体
	c.Data["json"] = struct {
		ID string
	}{
		user.ID,
	}
	c.ServeJSON()
}

// Destroy 用户删除
func (c *UserController) Destroy() {
	id := c.Ctx.Input.Param(":id")
	if err := core.DB.Where("id = ?", id).Delete(models.User{}).Error; err != nil {
		// return err
	}
	c.ServeJSON()
}

// Update 用户更新
func (c *UserController) Update() {
	id := c.Ctx.Input.Param(":id")
	name := c.GetString("name")

	user := models.User{}
	if err := core.DB.Model(&user).Where("id = ?", id).Updates(models.User{Name: name}).Error; err != nil {
		// return err
	}

	c.ServeJSON()
}

// Show 用户查询
func (c *UserController) Show() {
	id := c.Ctx.Input.Param(":id")

	user := models.User{}
	if err := core.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// 查询失败 ...
		c.Data["json"] = err.Error()
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}
