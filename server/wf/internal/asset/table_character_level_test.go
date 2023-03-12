package asset

import (
	"fmt"
	"testing"
)

func TestAsset_GetCharacterLevelTable(t *testing.T) {
	fmt.Println(GlobalAsset.GetCharacterLevelTable())
}
