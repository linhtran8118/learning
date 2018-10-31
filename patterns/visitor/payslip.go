package main

import (
	"fmt"
)

type Payslip struct {
	Exporter
}

func (p Payslip) ToXML() bool {
	return p.ExportXML()
}

func (p Payslip) ToJson() bool {
	return p.ExportJson()
}

type PayslipExporter struct {}

func (e PayslipExporter) ExportXML() bool {
	fmt.Println("Payslip1.XML")
	return true
}

func (e PayslipExporter) ExportJson() bool {
	fmt.Println("Payslip1.Json")
	return true
}
