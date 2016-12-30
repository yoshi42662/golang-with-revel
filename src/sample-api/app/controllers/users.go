package controllers

import "github.com/revel/revel"

type Users struct {
  *revel.Controller
}

func (c Users) Index() revel.Result {
  c.Flash.Success("Welcome, Yoshi")
  return c.Render()
}

func (c Users) Show() revel.Result {
  c.Flash.Success("Welcome, Yoshi")
  return c.Render()
}
