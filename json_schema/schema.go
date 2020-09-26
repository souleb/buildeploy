package json_schema

import (
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

var schemaData = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "config_schema_v1.json",
  "type": "object",
  "patternProperties": {
    "jobs": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/job"
      },
      "default": []
    },
    "name": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "definitions": {
    "job": {
      "id": "#/definitions/job",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "executor": {
          "$ref": "#/definitions/executor"
        },
        "env": {
          "type": "string"
        },
        "steps": {
          "type": "array",
          "items": {
            "oneOf": [
              {
                "type": "string"
              },
              {
                "type": "object"
              }
            ]
          }
        },
        "branches": {
          "type": "string"
        }
      },
        
      "additionalProperties": false
    },
    "executor": {
      "id": "#/definitions/executor",
      "type": "object",
      "minProperties": 1,
      "maxProperties": 1,
      "patternProperties": {
        "^docker": {
          "$ref": "#/definitions/docker"
        },
        "^machine": {
          "$ref": "#/definitions/machine"
        }
      },
      "additionalProperties": false
    },
    "docker": {
      "id": "#/definitions/docker",
      "type": "object",
      "properties": {
        "image": {
          "type": "string"
        },
        "tags": {
          "type": "string"
        }
      },
      "additionalProperties": false
    },
    "machine": {
      "id": "#/definitions/machine",
      "type": "object",
      "properties": {
        "os": {
          "type": "string"
        },
        "cpu_cores": {
          "type": [
            "number",
            "string"
          ]
        },
        "mem": {
          "type": [
            "number",
            "string"
          ]
        }
      },
      "additionalProperties": false
    }
  }
}
`

type SchemaService struct {
	schema *gojsonschema.Schema
}

func NewSchemaService() *SchemaService {
	schemaService := &SchemaService{}
	schemaService.setUpSchema()

	return schemaService
}

func (v *SchemaService) setUpSchema() {
	schemaLoader := gojsonschema.NewStringLoader(schemaData)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		panic(err.Error())
	}
	v.schema = schema
}

// Validate take a given interface and validate it.
// It uses it internal schema do to so.
func (v *SchemaService) Validate(data interface{}) error {
	documentLoader := gojsonschema.NewGoLoader(data)

	//result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	result, err := v.schema.Validate(documentLoader)
	if err != nil {
		return err
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
		return nil
	}

	fmt.Printf("The document is not valid. see errors :\n")
	var validationErrors []string
	for _, desc := range result.Errors() {
		validationErrors = append(validationErrors, fmt.Sprintf("- %s\n", desc))
	}

	return fmt.Errorf(strings.Join(validationErrors, "\n"))

}
