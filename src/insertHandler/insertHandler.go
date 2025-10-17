package insertHandler

import (
	"fmt"

	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Insert(kvPair models.KeyValuePair) {
	fmt.Printf("%+v", kvPair)
}
