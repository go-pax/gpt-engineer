package main

import (
	"encoding/json"
	"flag"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var (
	model       string
	lang        string
	temperature float64
	steps       string
)

func init() {
	flag.StringVar(&model, "model", defaultModel, "The model to use or for Azure use deployment name")
	flag.Float64Var(&temperature, "temperature", defaultTemperature, "The temperature to use")
	flag.StringVar(&lang, "lang", defaultLang, "The language to use")
	flag.StringVar(&steps, "steps", "default", "The steps to run")
}

func main() {
	flag.Parse()
	projectPath := flag.Arg(0)
	if projectPath == "" {
		projectPath = "./projects/example"
	}
	ai := NewAI(model, temperature, lang)
	dbs, err := NewDBs(projectPath)
	if err != nil {
		print(err)
		os.Exit(1)
	}

	for _, step := range STEPS[steps] {
		messages := step(ai, dbs)

		pc := runtime.FuncForPC(reflect.ValueOf(step).Pointer())
		funcName := strings.ReplaceAll(pc.Name(), "main.", "") // 去除包名
		contents, _ := json.Marshal(messages)
		dbs.logs.Set(funcName, string(contents))
	}
}
