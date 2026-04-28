package subscribers

import (
	"fmt"

	"github.com/poisnoir/purifier/models"
)

func ControllerHandler(data models.XboxController) {
	fmt.Println(data)
}
