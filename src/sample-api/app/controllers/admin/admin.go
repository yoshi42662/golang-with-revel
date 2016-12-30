package controllers

import "github.com/revel/revel"

type Admin struct {
  *revel.Controller
}

func (c Admin) Index() revel.Result {
  //Meta tags
  title := "Admin page"

  //Basic information.

  // Flash
  c.Flash.Success("Welcome, Yoshi")
  return c.Render(title)
}

func (c Admin) Show() revel.Result {
  c.Flash.Success("Welcome, Yoshi")
  return c.Render()
}
