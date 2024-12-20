package testhelpers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/monstermichl/domigo"
)

func TestDatabaseName() string {
	return fmt.Sprintf("GoInterfaceCopy-%d.nsf", time.Now().Unix())
}

func CreateTestDatabase(session domigo.NotesSession) (domigo.NotesDatabase, error) {
	now := time.Now().Unix()
	dbCopyFile := TestDatabaseName()
	dbTemp, err := session.GetDatabase("", "GoInterface.nsf")

	if err != nil {
		return domigo.NotesDatabase{}, errors.New("base database could not be retrieved")
	}
	db, err := dbTemp.CreateCopy("", dbCopyFile)

	if err != nil {
		return domigo.NotesDatabase{}, errors.New("base database could not be copied")
	}
	err = db.SetTitle(fmt.Sprintf("Go Interface Test %d", now))

	if err != nil {
		return domigo.NotesDatabase{}, errors.New("database title could not be set")
	}
	return db, err
}

func CreateTestView(db domigo.NotesDatabase, columns []domigo.String) (domigo.NotesView, error) {
	viewName := fmt.Sprintf("TestView-%d", time.Now().Unix())
	view, err := db.CreateView(domigo.WithNotesDatabaseCreateViewViewName(viewName), domigo.WithNotesDatabaseCreateViewProhibitDesignRefreshModifications(false))

	if err == nil {
		err = view.SetSelectionFormula("@All")
	}

	if err == nil {
		err = view.Refresh()
	}

	for _, cn := range columns {
		c, err := view.CreateColumn(domigo.WithNotesViewCreateColumnColumnName(cn))
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

func CreateTestDocument(db domigo.NotesDatabase) (domigo.NotesDocument, error) {
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

func Initialize(c func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error)) {
	var info string

	session, err := domigo.Initialize()
	defer session.Release()

	defer func() {
		if err != nil {
			fmt.Println(err)
		}
		info = strings.TrimSpace(info)

		if info != "" {
			fmt.Println(info)
		}
	}()

	if err != nil {
		info = "Session could not be initialized"
		return
	}

	db, err := CreateTestDatabase(session)
	defer db.Release()
	defer db.Remove()

	if err != nil {
		info = "Database could not be created"
		return
	}
	info, err = c(session, db)
}
