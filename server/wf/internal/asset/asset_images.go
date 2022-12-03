package asset

import "fmt"

func (a *Asset) GetCharacterFaceImage(thumbnailId string, evolution uint8) []byte {
	return a.GetPicture(fmt.Sprintf("character/%s/ui/square_%d", thumbnailId, evolution))
}
