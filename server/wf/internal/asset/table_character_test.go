package asset

import (
	"fmt"
	"testing"
)

func TestAsset_GetCharacterTable(t *testing.T) {
	for id, character := range GlobalAsset.GetCharacterTable() {
		fmt.Println(id, character)
	}
}
