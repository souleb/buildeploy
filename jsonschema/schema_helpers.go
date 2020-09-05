package jsonschema

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func Validate(data interface{}) {
	schemaLoader := gojsonschema.NewStringLoader(schemaDataV1)
	documentLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
