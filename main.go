package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Arman92/go-tl-parser/generator"
	"github.com/Arman92/go-tl-parser/tlparser"
)

type config struct {
	version          string
	packageName      string
	typesOutputDir   string
	methodsOutputDir string
	basePackageUri   string
}

func main() {
	var config config

	flag.StringVar(&config.version, "version", "v1.7.0", "TDLib version")
	flag.StringVar(&config.typesOutputDir, "typesOutputDir", "../go-tdlib/tdlib/", "output directory")
	flag.StringVar(&config.methodsOutputDir, "methodsOutputDir", "../go-tdlib/client/", "output directory")
	flag.StringVar(&config.basePackageUri, "basePackageUri", "github.com/Arman92/go-tdlib", "base package uri")
	flag.StringVar(&config.packageName, "package", "tdlib", "package name")

	flag.Parse()

	resp, err := http.Get("https://raw.githubusercontent.com/tdlib/td/" + config.version + "/td/generate/scheme/td_api.tl")
	if err != nil {
		log.Fatalf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	schema, err := tlparser.ParseInputSchema(resp.Body)
	if err != nil {
		log.Fatalf("schema parse error: %s", err)
		return
	}

	generator.GenerateCode(schema, config.basePackageUri, config.packageName, config.typesOutputDir, config.methodsOutputDir)

}
