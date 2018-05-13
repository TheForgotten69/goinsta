// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media := inst.AcquireFeed()
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Println("Deleting", os.Args[2])
	err = media.Items[0].Delete()
	e.CheckErr(err)

	err = media.Sync()
	if err != nil {
		fmt.Println("Deleted!")
	} else {
		fmt.Println("error deleting...")
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
