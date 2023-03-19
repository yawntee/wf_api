package asset

import (
	"fmt"
	"testing"
)

func TestAsset_GetBoxRewardTable(t *testing.T) {
	sum := 0
	for _, reward := range GlobalAsset.GetBoxRewardTable()[5][1] {
		sum += reward.Count
	}
	fmt.Println(sum)
}
