/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEW_CLASS.html */
package notesview_test

import (
	"fmt"
	"testing"

	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdocument"
	"github.com/monstermichl/domigo/domino/notessession"
	"github.com/monstermichl/domigo/domino/notesview"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var document notesdocument.NotesDocument
var view notesview.NotesView

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

	err = db.SetIsDesignLockingEnabled(true)

	if err != nil {
		info = "Design locking could not be enabled"
		return
	}

	document, err = testhelpers.CreateTestDocument(db)

	if err != nil {
		info = "Document could not be saved"
		return
	}

	view, err = testhelpers.CreateTestView(db, []domino.String{"Column 1", "Column 2", "Column 2"})

	if err != nil {
		info = "View could not be created"
		return
	}

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func TestAliases(t *testing.T) {
	_, err := view.Aliases()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func TestSetAliases(t *testing.T) {
	s, err := view.Aliases()
	require.Nil(t, err)

	err = view.SetAliases(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIES_PROPERTY_6084.html */
func TestAllEntries(t *testing.T) {
	_, err := view.AllEntries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func TestAutoUpdate(t *testing.T) {
	_, err := view.AutoUpdate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func TestSetAutoUpdate(t *testing.T) {
	s, err := view.AutoUpdate()
	require.Nil(t, err)

	err = view.SetAutoUpdate(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func TestBackgroundColor(t *testing.T) {
	_, err := view.BackgroundColor()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func TestSetBackgroundColor(t *testing.T) {
	s, err := view.BackgroundColor()
	require.Nil(t, err)

	err = view.SetBackgroundColor(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNCOUNT_PROPERTY_7753.html */
func TestColumnCount(t *testing.T) {
	_, err := view.ColumnCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNNAMES_PROPERTY_NOTESVIEW_CLASS.html */
func TestColumnNames(t *testing.T) {
	_, err := view.ColumnNames()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNS_PROPERTY.html */
func TestColumns(t *testing.T) {
	_, err := view.Columns()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_VIEW.html */
func TestCreated(t *testing.T) {
	_, err := view.Created()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCOUNT_PROPERTY_VIEW.html */
func TestEntryCount(t *testing.T) {
	_, err := view.EntryCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERLINES_PROPERTY_1209.html */
func TestHeaderLines(t *testing.T) {
	_, err := view.HeaderLines()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_NOTESVIEW_CLASS.html */
func TestHttpURL(t *testing.T) {
	_, err := view.HttpURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCALENDAR_PROPERTY.html */
func TestIsCalendar(t *testing.T) {
	_, err := view.IsCalendar()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORIZED_PROPERTY_2814.html */
func TestIsCategorized(t *testing.T) {
	_, err := view.IsCategorized()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1242.html */
func TestIsConflict(t *testing.T) {
	_, err := view.IsConflict()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func TestIsDefaultView(t *testing.T) {
	_, err := view.IsDefaultView()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func TestSetIsDefaultView(t *testing.T) {
	s, err := view.IsDefaultView()
	require.Nil(t, err)

	err = view.SetIsDefaultView(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFOLDER_PROPERTY.html */
func TestIsFolder(t *testing.T) {
	_, err := view.IsFolder()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIERARCHICAL_PROPERTY_7038.html */
func TestIsHierarchical(t *testing.T) {
	_, err := view.IsHierarchical()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMODIFIED_PROPERTY_6416.html */
func TestIsModified(t *testing.T) {
	_, err := view.IsModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_VIEW.html */
func TestIsPrivate(t *testing.T) {
	_, err := view.IsPrivate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func TestIsProhibitDesignRefresh(t *testing.T) {
	_, err := view.IsProhibitDesignRefresh()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func TestSetIsProhibitDesignRefresh(t *testing.T) {
	s, err := view.IsProhibitDesignRefresh()
	require.Nil(t, err)

	err = view.SetIsProhibitDesignRefresh(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_VIEW.html */
func TestLastModified(t *testing.T) {
	_, err := view.LastModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_VIEW.html */
func TestLockHolders(t *testing.T) {
	_, err := view.LockHolders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func TestName(t *testing.T) {
	_, err := view.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func TestSetName(t *testing.T) {
	s, err := view.Name()
	require.Nil(t, err)

	err = view.SetName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_NOTESVIEW_CLASS.html */
func TestNotesURL(t *testing.T) {
	_, err := view.NotesURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func TestProtectReaders(t *testing.T) {
	_, err := view.ProtectReaders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func TestSetProtectReaders(t *testing.T) {
	s, err := view.ProtectReaders()
	require.Nil(t, err)

	err = view.SetProtectReaders(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func TestReaders(t *testing.T) {
	_, err := view.Readers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func TestSetReaders(t *testing.T) {
	s, err := view.Readers()
	require.Nil(t, err)

	err = view.SetReaders(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROWLINES_PROPERTY_4578.html */
func TestRowLines(t *testing.T) {
	_, err := view.RowLines()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func TestSelectionFormula(t *testing.T) {
	_, err := view.SelectionFormula()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func TestSetSelectionFormula(t *testing.T) {
	s, err := view.SelectionFormula()
	require.Nil(t, err)

	err = view.SetSelectionFormula(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func TestSpacing(t *testing.T) {
	_, err := view.Spacing()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func TestSetSpacing(t *testing.T) {
	s, err := view.Spacing()
	require.Nil(t, err)

	err = view.SetSpacing(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TOPLEVELENTRYCOUNT_PROPERTY_6487.html */
func TestTopLevelEntryCount(t *testing.T) {
	_, err := view.TopLevelEntryCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_VIEW.html */
func TestUniversalID(t *testing.T) {
	_, err := view.UniversalID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWINHERITEDNAME_PROPERTY_VIEW.html */
func TestViewInheritedName(t *testing.T) {
	_, err := view.ViewInheritedName()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_VIEW.html */
func TestClear(t *testing.T) {
	err := view.Clear()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYCOLUMN_METHOD_VIEW.html */
func TestCopyColumn(t *testing.T) {
	_, err := view.CopyColumn(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLUMN_METHOD_VIEW.html */
func TestCreateColumn(t *testing.T) {
	_, err := view.CreateColumn(notesview.WithCreateColumnColumnName("secondColumn"))
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_VIEW.html */
func TestFTSearch(t *testing.T) {
	_, err := view.FTSearch("*", 2)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLDOCUMENTSBYKEY_METHOD.html */
func TestGetAllDocumentsByKey(t *testing.T) {
	_, err := view.GetAllDocumentsByKey([]domino.String{"key"})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIESBYKEY_METHOD_9837.html */
func TestGetAllEntriesByKey(t *testing.T) {
	_, err := view.GetAllEntriesByKey([]domino.String{"key"})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLREADENTRIES.html */
func TestGetAllReadEntries(t *testing.T) {
	_, err := view.GetAllReadEntries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLUNREADENTRIES.html */
func TestGetAllUnreadEntries(t *testing.T) {
	_, err := view.GetAllUnreadEntries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD.html */
func TestGetChild(t *testing.T) {
	_, err := view.GetChild(document)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCOLUMN_METHOD_NOTESVIEW_CLASS.html */
func TestGetColumn(t *testing.T) {
	_, err := view.GetColumn(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYKEY_METHOD.html */
func TestGetDocumentByKey(t *testing.T) {
	_, err := view.GetDocumentByKey([]domino.String{"key"})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYKEY_METHOD_3846.html */
func TestGetEntryByKey(t *testing.T) {
	_, err := view.GetEntryByKey([]domino.String{"key"})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEW.html */
func TestGetFirstDocument(t *testing.T) {
	_, err := view.GetFirstDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEW.html */
func TestGetLastDocument(t *testing.T) {
	_, err := view.GetLastDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEW.html */
func TestGetNextDocument(t *testing.T) {
	_, err := view.GetNextDocument(document)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD.html */
func TestGetNextSibling(t *testing.T) {
	_, err := view.GetNextSibling(document)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_VIEW.html */
func TestGetNthDocument(t *testing.T) {
	_, err := view.GetNthDocument(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTDOCUMENT_METHOD.html */
func TestGetParentDocument(t *testing.T) {
	_, err := view.GetParentDocument(document)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEW.html */
func TestGetPrevDocument(t *testing.T) {
	_, err := view.GetPrevDocument(document)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD.html */
func TestGetPrevSibling(t *testing.T) {
	_, err := view.GetPrevSibling(document)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_VIEW.html */
func TestLock(t *testing.T) {
	_, err := view.Lock()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_VIEW.html */
func TestLockProvisional(t *testing.T) {
	_, err := view.LockProvisional()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEW.html */
func TestMarkAllRead(t *testing.T) {
	err := view.MarkAllRead()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEW.html */
func TestMarkAllUnread(t *testing.T) {
	err := view.MarkAllUnread()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESH_METHOD_VIEW.html */
func TestRefresh(t *testing.T) {
	err := view.Refresh()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTVIEW_METHOD_VIEW.html */
func TestResortView(t *testing.T) {
	err := view.ResortView()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVECOLUMN_METHOD_VIEW.html */
func TestRemoveColumn(t *testing.T) {
	err := view.RemoveColumn()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_VIEW.html */
func TestUnLock(t *testing.T) {
	err := view.UnLock()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_VIEW.html */
func TestRemove(t *testing.T) {
	err := view.Remove()
	require.Nil(t, err)
}
