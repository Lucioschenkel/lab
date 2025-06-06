package main

import (

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "gorm.io/gorm/clause"
)

type Category struct {
  ID int `gorm:"primaryKey"`
  Name string
  Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
  ID int `gorm:"primaryKey"`
  Name string
  Categories []Category `gorm:"many2many:products_categories;"`
  Price float64
  gorm.Model
}

func main()  {
  dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&Product{}, &Category{})

  tx := db.Begin()
  var c Category
  err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
  if err != nil {
    panic(err)
  }

  c.Name = "Eletronicos"
  tx.Debug().Save(&c)
  tx.Debug().Commit()
}


// Lock otimista - versionamento
// Lock pessimista - bloqueia a linha da tabela (caro computacionalmente)
