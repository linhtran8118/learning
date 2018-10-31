package main

type Payslip struct {
}

func (p Payslip) export(v Visitor) bool {
	return v.exportPayslip()
}