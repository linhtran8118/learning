package main

import (
	"fmt"
)

func main() {
}

type Document interface {
	export(v Visitor) bool
}

type Exporter interface {
	ExportXML() bool
	ExportJson() bool
}

type Visitor interface {
	exportPayslip() bool
	exportTimesheet() bool
}

type XMLVisitor struct{}

func (v XMLVisitor) exportPayslip() bool {
	fmt.Println("Payslip-1.XML")
	return true
}

func (v XMLVisitor) exportTimesheet() bool {
	fmt.Println("Timesheet-1.XML")
	return true
}

type JsonVisitor struct{}

func (v JsonVisitor) exportPayslip() bool {
	fmt.Println("Payslip-1.Json")
	return true
}

func (v JsonVisitor) exportTimesheet() bool {
	fmt.Println("Timesheet-1.Json")
	return true
}