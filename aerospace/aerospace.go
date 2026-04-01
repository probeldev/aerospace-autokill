package aerospace

import (
	"strings"

	"github.com/probeldev/aerospace-autokill/bash"
	"github.com/probeldev/aerospace-autokill/model"
)

func GetAllWindows() (
	[]model.AerospaceWindow,
	error,
) {
	command := `aerospace list-windows --all --format "%{app-pid}:%{app-name}"`

	output, err := bash.RunCommand(command)

	if err != nil {
		return nil, err
	}

	response, err := parseAerospaceOutput(output)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func parseAerospaceOutput(
	output string,
) (
	[]model.AerospaceWindow,
	error,
) {

	response := []model.AerospaceWindow{}

	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		data := strings.Split(line, ":")

		response = append(response, model.AerospaceWindow{
			Pid:  data[0],
			Name: data[1],
		})

	}

	return response, nil

}
