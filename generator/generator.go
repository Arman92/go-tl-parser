package generator

import (
	"log"
	"os"

	"github.com/Arman92/go-tl-parser/tlparser"
)

const (
	commonFileName = "common.go"
)

func GenerateCode(schema *tlparser.TlSchema, packageName, outputDir string) {

	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir, os.ModePerm)

	err := generateCommonFiles(packageName, outputDir)
	if err != nil {
		log.Fatal("Failed to generate/write common file:", err)
	}

	err = GenerateInterfaceAndEnums(schema, packageName, outputDir)
	if err != nil {
		log.Fatal("Failed to generate/write interface/enum files:", err)
	}

	err = GenerateClasses(schema, packageName, outputDir)
	if err != nil {
		log.Fatal("Failed to generate/write classes files:", err)
	}

	err = GenerateFunctions(schema, packageName, outputDir)
	if err != nil {
		log.Fatal("Failed to generate/write function files:", err)
	}

}
