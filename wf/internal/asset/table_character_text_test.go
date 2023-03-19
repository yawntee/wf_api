package asset

import (
	"fmt"
	"testing"
)

func TestAsset_GetCharacterTextTable(t *testing.T) {
	for id, text := range GlobalAsset.GetCharacterTextTable() {
		fmt.Println(id, text)
	}
}
