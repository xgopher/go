package controllers

import (
	"app/modules/helloworld/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// ObjectController ...
type ObjectController struct {
	beego.Controller
}

// Index 列表
// 不要用 this, self 这些关键字
// 统一用 i , 为什么？ iphone, ipad, i love you...
func (i *ObjectController) Index() {
	obs := models.GetAll()
	i.Data["json"] = obs
	i.ServeJSON()
}

// Store 创建
func (i *ObjectController) Store() {
	var ob models.Object
	json.Unmarshal(i.Ctx.Input.RequestBody, &ob)
	objectID := models.AddOne(ob)
	i.Data["json"] = map[string]string{"objectID": objectID}
	i.ServeJSON()
}

// Show 查看
func (i *ObjectController) Show() {
	objectID := i.Ctx.Input.Param(":id")
	if objectID != "" {
		ob, err := models.GetOne(objectID)
		if err != nil {
			i.Data["json"] = err.Error()
		} else {
			i.Data["json"] = ob
		}
	}
	i.ServeJSON()
}

// Update 更新
func (i *ObjectController) Update() {
	objectID := i.Ctx.Input.Param(":id")
	var ob models.Object
	json.Unmarshal(i.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectID, ob.Score)
	if err != nil {
		i.Data["json"] = err.Error()
	} else {
		i.Data["json"] = "update success!"
	}
	i.ServeJSON()
}

// Destroy 删除
func (i *ObjectController) Destroy() {
	objectID := i.Ctx.Input.Param(":id")
	models.Delete(objectID)
	i.Data["json"] = "delete success!"
	i.ServeJSON()
}
