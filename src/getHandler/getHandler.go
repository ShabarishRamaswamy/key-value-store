package getHandler

import (
	"fmt"

	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Get(kvPair models.KeyValuePair) {
	fmt.Printf("%+v", kvPair)
}
