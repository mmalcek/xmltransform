package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/clbanning/mxj/v2"
)

func main() {
	inputFile := flag.String("i", "", "input file")
	outputFile := flag.String("o", "", "output file")
	textTemplate := flag.String("t", "", "template")
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("input file must be defined: -i input.xml", *inputFile)
	}
	if *outputFile == "" {
		log.Fatal("output file must be defined: -o output.csv")
	}
	if *textTemplate == "" {
		log.Fatal("template file must be defined: -t template.tmpl")
	}

	data, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal("readFile: ", err.Error())
	}
	mapXML, err := mxj.NewMapXml(data)
	if err != nil {
		log.Fatal("mapXML: ", err.Error())
	}
	template, err := template.ParseFiles(*textTemplate)
	if err != nil {
		log.Fatal("parseTemplate: ", err.Error())
	}
	output, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal("createOutputFile: ", err.Error())
	}
	defer output.Close()
	err = template.Execute(output, mapXML)
	if err != nil {
		log.Fatal("writeOutputFile: ", err.Error())
	}
}
