package generator

import (
	"log"
	"os"

	"github.com/Arman92/go-tl-parser/tlparser"
)

const (
	commonFileName = "common.go"
)

func GenerateCode(schema *tlparser.TlSchema, basePackageUri, packageName, typesOutputDir, methodsOutputDir string) {

	os.RemoveAll(typesOutputDir)
	os.MkdirAll(typesOutputDir, os.ModePerm)

	err := generateCommonFiles(packageName, typesOutputDir)
	if err != nil {
		log.Fatal("Failed to generate/write common file:", err)
	}

	err = GenerateInterfaceAndEnums(schema, packageName, typesOutputDir)
	if err != nil {
		log.Fatal("Failed to generate/write interface/enum files:", err)
	}

	err = GenerateClasses(schema, packageName, typesOutputDir)
	if err != nil {
		log.Fatal("Failed to generate/write classes files:", err)
	}

	err = GenerateMethods(schema, basePackageUri, packageName, "client", methodsOutputDir)
	if err != nil {
		log.Fatal("Failed to generate/write method files:", err)
	}

}
