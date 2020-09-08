package yaml

import (
	"gopkg.in/yaml.v3"
)

type Handler struct {
	mapping interface{}
}

func (h *Handler) mapToInterface(data string) error {
	err := yaml.Unmarshal([]byte(data), &h.mapping)
	if err != nil {
		return err
	}

	return nil

}
