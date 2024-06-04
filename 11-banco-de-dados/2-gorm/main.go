package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
  ID int `gorm:"primaryKey"`
  Name string
  Price float64
  gorm.Model
}

func main()  {
  dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&Product{})

  products := []Product{
    {Name: "ASUS Vivobook", Price: 10000.0},
    {Name: "Logitech MX Master 3S", Price: 500.0},
  }
  db.Create(&products)
  // var product Product
  // db.First(&product, 1)
  // fmt.Println(product)
  // db.First(&product, "name = ?", "ASUS Vivobook")
  // fmt.Println(product)

  // select all
  // var products []Product
  // db.Find(&products)
  // for _, product := range products {
  //   fmt.Println(product)
  // }
  // var products []Product
  // db.Limit(2).Offset(2).Find(&products)
  // for _, product := range products {
  //   fmt.Println(product)
  // }

  // where
  // var products []Product
  // db.Where("price > ?", 500.0).Find(&products)
  // for _, product := range products {
  //   fmt.Println(product)
  // }

  // like
  // var products []Product
  // db.Where("name LIKE ?", "%book%").Find(&products)
  // for _, product := range products {
  //   fmt.Println(product)
  // }
  // var p Product
  // db.First(&p, 2)
  // p.Name = "New Mouse"
  // db.Save(&p)

  var p2 Product
  db.First(&p2, 2)
  fmt.Println(p2.Name)

  db.Delete(&p2)

}

