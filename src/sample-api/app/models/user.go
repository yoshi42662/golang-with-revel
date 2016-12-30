package models

import  "time" // if you need/want

type User struct {
    Id                    int64
    Name                  string
    Email                 string
    EcncryptedPassword    []byte
    Password              string      `sql:"-"`
    CreatedAt             time.Time
    UpdatedAt             time.Time
    DeletedAt             time.Time
}
