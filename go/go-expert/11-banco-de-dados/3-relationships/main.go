package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
  ID int `gorm:"primaryKey"`
  Name string
}

type Product struct {
  ID int `gorm:"primaryKey"`
  Name string
  CategoryID int
  Category Category
  Price float64
  SerialNumber SerialNumber
  gorm.Model
}

type SerialNumber struct {
  ID int `gorm:"primaryKey"`
  Number string
  ProductID int
}

func main()  {
  dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

  // create category
  category := Category{Name: "Electronics"}
  db.Create(&category)

  // create product
  db.Create(&Product{
    Name: "Notebook",
    Price: 1000.0,
    CategoryID: category.ID,
  })

  db.Create(&SerialNumber{
    Number: "123456",
    ProductID: 1,
  })

  var products []Product
  db.Preload("Category").Preload("SerialNumber").Find(&products)
  for _, product := range products {
    fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
  }
}

