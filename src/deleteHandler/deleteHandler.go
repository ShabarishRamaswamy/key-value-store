package deleteHandler

import (
	"fmt"

	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Delete(kvPair models.KeyValuePair) {
	fmt.Printf("%+v", kvPair)
}
