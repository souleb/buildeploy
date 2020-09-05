package yaml

import (
	"fmt"
	"testing"
)

func TestYamlMapper(t *testing.T) {
	tests := []string{
		`
name: "dev workflow"
jobs:
- name: "dev job"
  runner:
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
  runner:
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
		fmt.Println(yamlMapper(test))
		fmt.Println("done")
	}
}
