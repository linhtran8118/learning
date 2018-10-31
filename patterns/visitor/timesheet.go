package main

type Timesheet struct {
}

func (p Timesheet) export(v Visitor) bool {
	return v.exportTimesheet()
}
