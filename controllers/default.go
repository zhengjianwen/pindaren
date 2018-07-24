package controllers

type MainController struct {
	Base
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}
