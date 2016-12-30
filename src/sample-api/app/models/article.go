package models

import  "time" // if you need/want

type Article struct {
    Id                    int64
    title                 string
    content               string
    thumnail              string
    PublishedAt           time.Time
    CreatedAt             time.Time
    UpdatedAt             time.Time
    DeletedAt             time.Time
}
