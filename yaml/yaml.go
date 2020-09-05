package yaml

import (
	"github.com/SouleBA/buildeploy/app"
	"github.com/SouleBA/buildeploy/jsonschema"

	"gopkg.in/yaml.v3"
)

type handler struct {
	JobService  app.JobService
	StepService app.StepService
}

func yamlMapper(content string) (interface{}, error) {
	var workflow interface{}
	err := yaml.Unmarshal([]byte(content), &workflow)
	if err != nil {
		return app.Workflow{}, err
	}

	jsonschema.Validate(workflow)

	return workflow, nil

}
