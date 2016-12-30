package controllers

import (
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/jinzhu/gorm"
  "github.com/revel/revel"
  "sample-api/app/models"
)

type GormController struct {
    *revel.Controller
    Txn *gorm.DB
}

type App struct {
  GormController
}

var Gdb gorm.DB

// init db
func InitDB() {
    var err error
    // open db
    Gdb, err := gorm.Open("postgres", "dbname=gorm sslmode=disable password=supersecret")
    if err != nil {
        //r.ERROR.Println("FATAL", err)
        panic( err )
    }
    Gdb.AutoMigrate(&models.User{})
    // uniquie index if need
    //Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}

func (c App) Index() revel.Result {
  InitDB()
  title := "index page"

  //user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

  user := models.User{Name: "Jinzhup"}
  //c.Txn.NewRecord(user)
  //c.Txn.Create(&user)
  c.Flash.Success("Welcome, Yoshi")
  return c.Render(title, user)
}
