package yaml

import (
	"testing"
)

func TestMapToInterface(t *testing.T) {
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
		d := Handler{}
		err := d.mapToInterface(test)
		if err != nil {
			t.Fatalf("mapping does not contain any key. got %s",
				err)
		}
	}
}
