package main

import (
	"testing"
)


func TestNormal(t *testing.T) {
	documents := []Document{
		Payslip{},
		Timesheet{},
	}
	vistor := XMLVisitor{}
	for _, doc := range documents {
		doc.export(vistor)
	}
}