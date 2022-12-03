package asset

import (
	"os"
	"testing"
)

func TestAsset_GetPicture(t *testing.T) {
	pic := GlobalAsset.GetCharacterFaceImage("fire_dragon", 1)
	err := os.WriteFile("out.png", pic, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
