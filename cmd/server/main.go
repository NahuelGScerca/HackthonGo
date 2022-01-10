package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/NahuelGScerca/HackthonGo/internal/sales"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/DBHackathon")
	if err != nil {
		fmt.Println("aqa")
		panic(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("aqa ping")

		panic(err)
	}
	log.Println("Database Configured")

	//SALES
	repoSales := sales.NewRepository(db)
	serviceSales := sales.NewService(repoSales)
	serviceSales.ExportData(context.Background())

	//PRODUCTS
	// repoProducts := products.NewRepository(db)
	// serviceProducts := products.NewService(repoProducts)
	// serviceProducts.ExportData(context.Background())

	//CUSTOMER
	// repoCustomer := customers.NewRepository(db)
	// serviceCustomer := customers.NewService(repoCustomer)
	// serviceCustomer.ExportData(context.Background())

	//INVOICES
	// repoInvoices := invoices.NewRepository(db)
	// serviceInvoices := invoices.NewService(repoInvoices)
	// serviceInvoices.ExportData(context.Background())
}
