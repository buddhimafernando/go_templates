package main

import (
	"fmt"
	"os"
	"text/template"
)

type Header struct {
	Company       string
	FormatDate    string
	FormatCurrency string
}

type MonthlySales struct {
	Month          string
	PreviousRevenue int
	MonthlyRevenue int
	Profit         int
}

type Employee struct {
	Position    string
	Performance int
}

type Product struct {
	Product  string
	Quantity int
	Price    int
}

func HeaderInfo() Header {
	header := Header{"Evolza", "01/01/2021", "LKR"}
	return header
}

func ProductInventory() {
	products := []Product{
		{"Laptop", 10, 100000},
		{"Mobile", 20, 50000},
		{"Tablet", 30, 30000},
		{"Desktop", 40, 80000},
		{"Printer", 50, 20000},
	}

	header := HeaderInfo()
	tmpl, err := template.ParseFiles("template3.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("templates/Template 3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, struct {
		Header   Header
		Products []Product
	}{
		Header:   header,
		Products: products,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func EmployeePerformance() {
	employees := map[string]Employee{
		"John Doe":    {"Manager", 60},
		"Jane Smith":  {"Assistant Manager", 69},
		"Tom White":   {"Sales Executive", 80},
		"Jerry Black": {"Sales Executive", 90},
		"Harry Green": {"Sales Executive", 70},
	}

	header := HeaderInfo()
	tmpl, err := template.ParseFiles("template2.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	file, err := os.Create("templates/Template 2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, struct {
		Header      Header
		Employees   map[string]Employee
	}{
		Header:    header,
		Employees: employees,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func SalesReport() {
	salesInfo := MonthlySales{"January", 900000, 1000000, 0}
	salesInfo.Profit = salesInfo.MonthlyRevenue - salesInfo.PreviousRevenue

	header := HeaderInfo()
	tmpl, err := template.ParseFiles("template.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	file, err := os.Create("templates/Template 1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, struct {
		Header     Header
		MonthlySales MonthlySales
	}{
		Header:       header,
		MonthlySales: salesInfo,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Print("Select the report type you need to generate: \n 1. Sales Report \n 2. Employee Performance \n 3. Product Inventory \n")
	fmt.Print("Enter the number of the report type: ")
	var reportType int
	fmt.Scan(&reportType)

	switch reportType {
	case 1:
		SalesReport()
	case 2:
		EmployeePerformance()
	case 3:
		ProductInventory()
	default:
		fmt.Println("Invalid selection")
	}
}
