package json

/*type SchemaService struct {
	schema *gojsonschema.Schema
}

func (v *SchemaService) setUpSchema() {
	schemaLoader := gojsonschema.NewStringLoader(schemaDataV1)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		panic(err.Error())
	}
	v.schema = schema
}

func (v *SchemaService) Validate(data interface{}) {
	documentLoader := gojsonschema.NewGoLoader(data)

	//result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	result, err := v.schema.Validate(documentLoader)
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
*/
