/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWNAVIGATOR_CLASS.html */
package notesviewnavigator_test

import (
	"fmt"
	"testing"

	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdocument"
	"github.com/monstermichl/domigo/domino/notessession"
	"github.com/monstermichl/domigo/domino/notesviewentry"
	"github.com/monstermichl/domigo/domino/notesviewnavigator"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var entry notesviewentry.NotesViewEntry
var document notesdocument.NotesDocument
var viewnavigator notesviewnavigator.NotesViewNavigator

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	var info string

	session, err := notessession.Initialize()
	defer session.Release()

	defer func() {
		fmt.Println(err)
		fmt.Println(info)
	}()

	if err != nil {
		info = "Session could not be initialized"
		return
	}

	db, err := testhelpers.CreateTestDatabase(session)
	defer db.Release()
	defer db.Remove()

	if err != nil {
		info = "Database could not be created"
		return
	}

	document, err = testhelpers.CreateTestDocument(db)

	if err != nil {
		info = "Document could not be saved"
		return
	}

	view, err := testhelpers.CreateTestView(db, []domino.String{"Column 1", "Column 2", "Column 2"})

	if err != nil {
		info = "View could not be created"
		return
	}

	entries, err := view.AllEntries()
	defer entries.Release()

	if err != nil {
		info = "Entries could not be retrieved"
		return
	}

	count, err := entries.Count()

	if err != nil {
		info = "View entries count could not be retrieved"
		return
	} else if count <= 0 {
		info = "No view entries found"
		return
	}

	entry, err = entries.GetFirstEntry()
	defer entry.Release()

	if err != nil {
		info = "First entry could not be retrieved"
		return
	}

	viewnavigator, err = notesviewnavigator.NotesViewCreateViewNav(view)
	defer viewnavigator.Release()

	if err != nil {
		info = "NotesViewNavigator could not be created"
		return
	}

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func TestCacheSize(t *testing.T) {
	_, err := viewnavigator.CacheSize()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func TestSetCacheSize(t *testing.T) {
	s, err := viewnavigator.CacheSize()
	require.Nil(t, err)

	err = viewnavigator.SetCacheSize(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_VIEWNAV.html */
func TestCount(t *testing.T) {
	_, err := viewnavigator.Count()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func TestMaxLevel(t *testing.T) {
	_, err := viewnavigator.MaxLevel()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func TestSetMaxLevel(t *testing.T) {
	s, err := viewnavigator.MaxLevel()
	require.Nil(t, err)

	err = viewnavigator.SetMaxLevel(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY_VIEWNAV.html */
func TestParentView(t *testing.T) {
	_, err := viewnavigator.ParentView()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_VIEWNAV.html */
func TestGetChild(t *testing.T) {
	_, err := viewnavigator.GetChild(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCURRENT_METHOD_VIEWNAV.html */
func TestGetCurrent(t *testing.T) {
	_, err := viewnavigator.GetCurrent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_VIEWNAV.html */
func TestGetEntry(t *testing.T) {
	_, err := viewnavigator.GetEntry(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_VIEWNAV.html */
func TestGetFirst(t *testing.T) {
	_, err := viewnavigator.GetFirst()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func TestGetFirstDocument(t *testing.T) {
	_, err := viewnavigator.GetFirstDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_VIEWNAV.html */
func TestGetLast(t *testing.T) {
	_, err := viewnavigator.GetLast()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEWNAV.html */
func TestGetLastDocument(t *testing.T) {
	_, err := viewnavigator.GetLastDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_VIEWNAV.html */
func TestGetNext(t *testing.T) {
	_, err := viewnavigator.GetNext(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTCATEGORY_METHOD_VIEWNAV.html */
func TestGetNextCategory(t *testing.T) {
	_, err := viewnavigator.GetNextCategory(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEWNAV.html */
func TestGetNextDocument(t *testing.T) {
	_, err := viewnavigator.GetNextDocument(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_VIEWNAV.html */
func TestGetNextSibling(t *testing.T) {
	_, err := viewnavigator.GetNextSibling(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTH_METHOD_VIEWNAV.html */
func TestGetNth(t *testing.T) {
	_, err := viewnavigator.GetNth(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_VIEWNAV.html */
func TestGetParent(t *testing.T) {
	_, err := viewnavigator.GetParent(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOS_METHOD_VIEWNAV.html */
func TestGetPos(t *testing.T) {
	_, err := viewnavigator.GetPos("1.2.3", ".")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_VIEWNAV.html */
func TestGetPrev(t *testing.T) {
	_, err := viewnavigator.GetPrev(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVCATEGORY_METHOD_VIEWNAV.html */
func TestGetPrevCategory(t *testing.T) {
	_, err := viewnavigator.GetPrevCategory(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEWNAV.html */
func TestGetPrevDocument(t *testing.T) {
	_, err := viewnavigator.GetPrevDocument(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_VIEWNAV.html */
func TestGetPrevSibling(t *testing.T) {
	_, err := viewnavigator.GetPrevSibling(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOCHILD_METHOD_NOTESVIEWNAV.html */
func TestGotoChild(t *testing.T) {
	err := viewnavigator.GotoChild(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOENTRY_METHOD_VIEWNAV.html */
func TestGotoEntry(t *testing.T) {
	err := viewnavigator.GotoEntry(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRST_METHOD_VIEWNAV.html */
func TestGotoFirst(t *testing.T) {
	err := viewnavigator.GotoFirst()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func TestGotoFirstDocument(t *testing.T) {
	err := viewnavigator.GotoFirstDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLAST_METHOD_VIEWNAV.html */
func TestGotoLast(t *testing.T) {
	err := viewnavigator.GotoLast()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLASTDOCUMENT_METHOD_VIEWNAV.html */
func TestGotoLastDocument(t *testing.T) {
	err := viewnavigator.GotoLastDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXT_METHOD_VIEWNAV.html */
func TestGotoNext(t *testing.T) {
	err := viewnavigator.GotoNext(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTCATEGORY_METHOD_VIEWNAV.html */
func TestGotoNextCategory(t *testing.T) {
	err := viewnavigator.GotoNextCategory(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTDOCUMENT_METHOD_VIEWNAV.html */
func TestGotoNextDocument(t *testing.T) {
	err := viewnavigator.GotoNextDocument(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTSIBLING_METHOD_VIEWNAV.html */
func TestGotoNextSibling(t *testing.T) {
	err := viewnavigator.GotoNextSibling(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPARENT_METHOD_VIEWNAV.html */
func TestGotoParent(t *testing.T) {
	err := viewnavigator.GotoParent(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPOS_METHOD_VIEWNAV.html */
func TestGotoPos(t *testing.T) {
	err := viewnavigator.GotoPos("1.2.3", ".")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREV_METHOD_VIEWNAV.html */
func TestGotoPrev(t *testing.T) {
	err := viewnavigator.GotoPrev(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVCATEGORY_METHOD_VIEWNAV.html */
func TestGotoPrevCategory(t *testing.T) {
	err := viewnavigator.GotoPrevCategory(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVDOCUMENT_METHOD_VIEWNAV.html */
func TestGotoPrevDocument(t *testing.T) {
	err := viewnavigator.GotoPrevDocument(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVSIBLING_METHOD_VIEWNAV.html */
func TestGotoPrevSibling(t *testing.T) {
	err := viewnavigator.GotoPrevSibling(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEWNAV.html */
func TestMarkAllRead(t *testing.T) {
	err := viewnavigator.MarkAllRead()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEWNAV.html */
func TestMarkAllUnread(t *testing.T) {
	err := viewnavigator.MarkAllUnread()
	require.Nil(t, err)
}
