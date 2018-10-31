package main

import (
	"fmt"
)

type Timesheet struct {
	Exporter
}

func (p Timesheet) ToXML() bool {
	return p.ExportXML()
}

func (p Timesheet) ToJson() bool {
	return p.ExportJson()
}

type TimeSheetExporter struct {}

func (e TimeSheetExporter) ExportXML() bool {
	fmt.Println("Timesheet.XML")
	return true
}

func (e TimeSheetExporter) ExportJson() bool {
	fmt.Println("Timesheet.Json")
	return true
}