package main

import (
	"culfi/apps"
	"fmt"
)

func main() {
	// TODO start server
	// for now test code
	apps, _ := apps.GetApps()
	fmt.Println(apps)
}
