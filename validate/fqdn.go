package validate

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func VariableFQDN(s string) error {
	if govalidator.IsDNSName(s) == false {
		return fmt.Errorf(NotAValidFQDN, s)
	}

	return nil
}
