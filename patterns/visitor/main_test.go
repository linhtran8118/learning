package main

import (
	"testing"
)


func TestNormal(t *testing.T) {
	documents := []Document{
		Payslip{PayslipExporter{}},
		Timesheet{TimeSheetExporter{}},
		Payslip{PayslipExporter{}},
		Timesheet{TimeSheetExporter{}},
	}
	for _, doc := range documents {
		doc.ToXML()
	}
}