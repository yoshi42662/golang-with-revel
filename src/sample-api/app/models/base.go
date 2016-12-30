package models

//import  "time" // if you need/want
import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "time"
)



type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}

type User struct {          // example user fields
    Name                  string
    EcncryptedPassword    []byte
    Password              string      `sql:"-"`
}

// import (
//     "github.com/jinzhu/gorm"
//     _ "github.com/jinzhu/gorm/dialects/postgres"
//     "sample-api/app/models"
// )

// func main() {
//   db, err := gorm.Open("postgres", "port=5432 host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword")
//   defer db.Close()
// }

// it can be used for jobs
var Gdb gorm.DB

// init db
func InitDB() {
    var err error
    // open db
    Gdb, err := gorm.Open("postgres", "user=uname dbname=udbname sslmode=disable password=supersecret")
    if err != nil {
        //r.ERROR.Println("FATAL", err)
        panic( err )
    }
    Gdb.AutoMigrate(&User{})
    // uniquie index if need
    //Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}


// transactions

// This method fills the c.Txn before each transaction
// func (c *GormController) Begin() r.Result {
//     txn := Gdb.Begin()
//     if txn.Error != nil {
//         panic(txn.Error)
//     }
//     c.Txn = txn
//     return nil
// }
//
// // This method clears the c.Txn after each transaction
// func (c *GormController) Commit() r.Result {
//     if c.Txn == nil {
//         return nil
//     }
//     c.Txn.Commit()
//     if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
//         panic(err)
//     }
//     c.Txn = nil
//     return nil
// }
//
// // This method clears the c.Txn after each transaction, too
// func (c *GormController) Rollback() r.Result {
//     if c.Txn == nil {
//         return nil
//     }
//     c.Txn.Rollback()
//     if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
//         panic(err)
//     }
//     c.Txn = nil
//     return nil
// }
