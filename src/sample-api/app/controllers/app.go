package controllers

import "github.com/revel/revel"

type App struct {
  *revel.Controller
}

func (c App) Index() revel.Result {
  title := "index page"

  c.Flash.Success("Welcome, Yoshi")
  return c.Render(title)
}
