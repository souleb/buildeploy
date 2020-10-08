package json_schema

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestValidate(t *testing.T) {
	tests := []string{
		`
name: "dev workflow"
jobs:
- name: "dev job"
  executor:
    docker:
      image: "dev.docker.io/golang:stable"
      tags: "golang:true, npm:false, prod: false"
  env: "env: dev"
  steps:
    - checkout
    - command: "echo \"Hello World in dev!\" "
  branches: "develop"
`,
		`
name: "prod workflow"
jobs:
- name: "prod job"
  executor:
    docker:
      image: "prod.docker.io/golang:stable"
      tags: "golang:true, npm:false, prod: true"
  env: "env: prod"
  steps:
    - checkout
    - command: "echo \"Hello World in prod!\" "
  branches: "production"
`}

	for _, test := range tests {
		var res map[string]interface{}
		err := yaml.Unmarshal([]byte(test), &res)
		if err != nil {
			t.Fatalf("input is not a valid yaml got %s",
				err)
		}

		s := NewSchemaService()
		err = s.Validate(res)
		if err != nil {
			t.Fatalf("validation failed, got %s",
				err)
		}
	}
}
