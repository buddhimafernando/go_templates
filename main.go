package main

import (
	"fmt"
	"os"
	"text/template"
)

type MonthlySales struct {
	Month  string
	PreviousRevenue  int
	MonthlyRevenue int
	Profit int
}

type Employee struct {
	Position string
	Performance int
}



func ProductInventory() {

}

func EmployeePerformance(){
	employees := map[string] Employee{
		"John Doe": { "Manager", 60},
		"Jane Smith" : {"Assistant Manager", 69},
		"Tom White" : {"Sales Executive", 80},
		"Jerry Black" : {"Sales Executive", 90},
		"Harry Green" : {"Sales Executive", 70},
	}

	tmpl1, err:= template.ParseFiles("template2.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl1.Execute(os.Stdout, employees)
	
	if err != nil {
		fmt.Println(err)
	}

}

func SalesReport() {
	salesInfo := MonthlySales{"January", 900000, 1000000, 0}
	profit := salesInfo.PreviousRevenue - salesInfo.MonthlyRevenue
	salesInfo.Profit = profit

	tmpl, err := template.ParseFiles("template.txt")
	if err!= nil{
		fmt.Println(err)
	}

	err = tmpl.Execute(os.Stdout, salesInfo)
	if err!= nil{
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
	}
}