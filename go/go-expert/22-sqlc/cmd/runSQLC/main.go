package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Lucioschenkel/go-expert/22-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Test",
	// 	Description: sql.NullString{String: "Test", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.Name)
	// }
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "c9632acf-7f3d-499b-9318-c91aa5d16fb2",
		Name:        "Category 1",
		Description: sql.NullString{String: "Category 1 Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	if err := queries.DeleteCategory(ctx, "c9632acf-7f3d-499b-9318-c91aa5d16fb2"); err != nil {
		panic(err)
	}

	fmt.Println("Category deleted")
}
