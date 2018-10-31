package main

func main() {
}

type Document interface {
	ToXML() bool
	ToJson() bool
}

type Exporter interface {
	ExportXML() bool
	ExportJson() bool
}

 //---------