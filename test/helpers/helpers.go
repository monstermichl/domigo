package testhelpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdatabase"
	"github.com/monstermichl/domigo/domino/notesdocument"
	"github.com/monstermichl/domigo/domino/notessession"
	"github.com/monstermichl/domigo/domino/notesview"
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

func CreateTestView(db notesdatabase.NotesDatabase, columns []domino.String) (notesview.NotesView, error) {
	viewName := fmt.Sprintf("TestView-%d", time.Now().Unix())
	view, err := db.CreateView(notesdatabase.WithCreateViewViewName(viewName), notesdatabase.WithCreateViewProhibitDesignRefreshModifications(false))

	if err == nil {
		err = view.SetSelectionFormula("@All")
	}

	if err == nil {
		err = view.Refresh()
	}

	for _, cn := range columns {
		c, err := view.CreateColumn(notesview.WithCreateColumnColumnName(cn))
		defer c.Release()

		if err == nil {
			err = c.SetIsSorted(true)
		}

		if err != nil {
			break
		}
	}
	return view, err
}

func CreateTestDocument(db notesdatabase.NotesDatabase) (notesdocument.NotesDocument, error) {
	document, err := db.CreateDocument()
	defer document.Release()

	if err == nil {
		_, err = document.ReplaceItemValue("testField", fmt.Sprint(time.Now().Unix()))
	}

	if err == nil {
		_, err = document.Save(true, false)
	}
	return document, err
}
