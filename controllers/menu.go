package controllers

import (
	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/services"
	"github.com/fajarpnugroho/omelettebot/models"
)

// MenuController menu controllers
type MenuController struct {
	beego.Controller
}

// Get index page
func (c *MenuController) Get() {
	c.Data["menu"] = "active"
	c.Data["titlePage"] = "Menu"
	c.Layout = "main.tpl"
	c.TplName = "menu.tpl"
}

func (m *MenuController) GetMenuByCategoryId() {
	categoryID := m.Ctx.Input.Param(":categoryid")

	menuModel := models.Menus{}
	services.GetMenuByCategory(&menuModel, categoryID)
	m.Data["menus"] = menuModel
}

func (m *MenuController) AddMenuByCategoryId() {
}
