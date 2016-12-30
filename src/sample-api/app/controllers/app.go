package controllers

import (
  "github.com/revel/revel"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "sample-api/app/models"
)

type App struct {
  *revel.Controller
  Txn *gorm.DB
}

func (c App) Index() revel.Result {
  ///////
  //

  db, err := gorm.Open("postgres", "port=5432 dbname=gorm sslmode=disable")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  db.AutoMigrate(&models.User{})
  db.Create(&models.User{ Name: "George", Email: "user@example.com" })
  var user models.User
  db.First(&user, "name = ?", "George").Update("name", "Member user")
  db.Delete(&user)
  ///
  ////////

  title := "index page"

  c.Flash.Success("Welcome, Yoshi")
  return c.Render(title)
}
