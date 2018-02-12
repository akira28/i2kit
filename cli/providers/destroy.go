package providers

import (
	"fmt"

	"github.com/pchico83/i2kit/cli/providers/aws"
	"github.com/pchico83/i2kit/cli/schemas/environment"
	"github.com/pchico83/i2kit/cli/schemas/service"
)

//Destroy destroys a given service in a given environment
func Destroy(s *service.Service, e *environment.Environment) error {
	if e.Provider == nil {
		fmt.Printf("Service '%s' dry-run destroy was successful\n", s.Name)
		return nil
	}
	return aws.Destroy(s, e)
}
