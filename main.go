package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/clbanning/mxj/v2"
	lua "github.com/yuin/gopher-lua"
)

const version = "1.0.2"

var (
	luaData  *lua.LState
	luaReady = false
)

func init() {
	if _, err := os.Stat("./lua/functions.lua"); !os.IsNotExist(err) {
		luaData = lua.NewState()
		err := luaData.DoFile("./lua/functions.lua")
		if err != nil {
			log.Fatal("loadLuaFunctions", err.Error())
		}
		luaReady = true
	}
}

func main() {
	inputFile := flag.String("i", "", "input file")
	outputFile := flag.String("o", "", "output file, if not defined stdout is used")
	textTemplate := flag.String("t", "", "template")
	getVersion := flag.Bool("v", false, "template")
	flag.Parse()

	if *getVersion {
		fmt.Printf("Version: %s", version)
		os.Exit(0)
	}

	if *inputFile == "" {
		log.Fatal("input file must be defined: -i input.xml", *inputFile)
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
	templateFile, err := ioutil.ReadFile(*textTemplate)
	if err != nil {
		log.Fatal("readFile: ", err.Error())
	}
	template, err := template.New("new").Funcs(templateFunctions()).Parse(string(templateFile))
	if err != nil {
		log.Fatal("parseTemplate: ", err.Error())
	}
	if *outputFile == "" {
		output := new(bytes.Buffer)
		err = template.Execute(output, mapXML)
		if err != nil {
			log.Fatal("writeStdout: ", err.Error())
		}
		fmt.Print(output)
	} else {
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
	if luaReady {
		luaData.Close()
	}
}
