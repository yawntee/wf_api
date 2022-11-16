package main

import (
	"fmt"
	"wf_api/wf"
)

func main() {
	client := wf.NewClient(wf.LEITING)
	err := client.Login("15295879613", "wzmwsadjkl123")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}
