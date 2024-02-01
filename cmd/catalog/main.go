package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/DemetriusLeonardoBantim/goapi/internal/database"
	"github.com/DemetriusLeonardoBantim/goapi/internal/service"
	"github.com/DemetriusLeonardoBantim/goapi/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goapi17")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewCProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/products", webProductHandler.GetProducts)
	c.Post("/product", webProductHandler.CreateProduct)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)

	fmt.Printf("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
