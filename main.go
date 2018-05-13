package main

import (
	"culfi/apps"
	"culfi/spaces"
	"fmt"
)

func main() {
	// TODO start server
	// for now test code
	spaces, _ := spaces.GetSpaces()
	fmt.Println(spaces)
	apps, _ := apps.GetApps(spaces[0].GUID)
	fmt.Println(apps)
}
