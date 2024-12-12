package testhelpers

import (
	"domigo/domino/notesdatabase"
	"domigo/domino/notessession"
	"errors"
	"fmt"
	"time"
)

func TestDatabaseName() string {
	return fmt.Sprintf("GoInterfaceCopy-%d.nsf", time.Now().Unix())
}

func CreateTestDatabase(session notessession.NotesSession) (notesdatabase.NotesDatabase, error) {
	now := time.Now().Unix()
	dbCopyFile := TestDatabaseName()
	dbTemp, err := session.GetDatabase("", "GoInterface.nsf")

	if err != nil {
		return notesdatabase.NotesDatabase{}, errors.New("base database could not be retrieved")
	}
	db, err := dbTemp.CreateCopy("", dbCopyFile)

	if err != nil {
		return notesdatabase.NotesDatabase{}, errors.New("base database could not be copied")
	}
	err = db.SetTitle(fmt.Sprintf("Go Interface Test %d", now))

	if err != nil {
		return notesdatabase.NotesDatabase{}, errors.New("database title could not be set")
	}
	return db, err
}
