package main

import (
	"domigo/domino/notessession"
	"fmt"
)

func main() {
	/* "http://test19web1.test-edv.net/at.gv.bgld.kulturgutschein-import-test.nsf/", "adsfas" */
	session, err := notessession.Initialize()

	if err != nil {
		fmt.Println(err)
	} else {
		db, err := session.GetDatabase("", "GoInterface.nsf")

		if err != nil {
			fmt.Println(err)
		} else {
			forms, _ := db.Forms()
			for _, form := range forms {
				fmt.Println(form.Name())
				break
			}
		}
		session.Release()
	}
}
