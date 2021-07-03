package seed

import (
  "../models"
  "../db"
)

func Seed(){
  db := db.SqlConnect()
  db.AutoMigrate(&models.User{})
  user := models.User{
    ID: 1,
    Name: "mizukami",
  }
  db.Create(&user)
  defer db.Close()
}
