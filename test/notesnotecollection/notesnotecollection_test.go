/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNOTECOLLECTION_CLASS.html */
package notesnotecollection_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var db domigo.NotesDatabase
var notecollection domigo.NotesNoteCollection

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, dbTemp domigo.NotesDatabase) (string, error) {
		var err error
		db = dbTemp

		notecollection, err = db.CreateNoteCollection(true)
		defer notecollection.Release()

		if err != nil {
			return "NotesNoteCollection could not be created", err
		}
		m.Run()
		return "", nil
	})
}

func createAndSaveDoc(t *testing.T) domigo.NotesDocument {
	doc, err := db.CreateDocument()
	require.Nil(t, err, "Document could not be created")

	saved, err := doc.Save(true, false)
	msg := "Document could not be saved"
	require.Nil(t, err, msg)
	require.True(t, saved, msg)

	return doc
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_NOTECOLLECTION.html */
func TestCount(t *testing.T) {
	_, err := notecollection.Count()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTBUILDTIME_PROPERTY_NC.html */
func TestLastBuildTime(t *testing.T) {
	_, err := notecollection.LastBuildTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTECOLLECTION.html */
func TestParent(t *testing.T) {
	_, err := notecollection.Parent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func TestSelectACL(t *testing.T) {
	_, err := notecollection.SelectACL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func TestSetSelectACL(t *testing.T) {
	s, err := notecollection.SelectACL()
	require.Nil(t, err)

	err = notecollection.SetSelectACL(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func TestSelectActions(t *testing.T) {
	_, err := notecollection.SelectActions()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func TestSetSelectActions(t *testing.T) {
	s, err := notecollection.SelectActions()
	require.Nil(t, err)

	err = notecollection.SetSelectActions(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func TestSelectAgents(t *testing.T) {
	_, err := notecollection.SelectAgents()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func TestSetSelectAgents(t *testing.T) {
	s, err := notecollection.SelectAgents()
	require.Nil(t, err)

	err = notecollection.SetSelectAgents(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func TestSelectDatabaseScript(t *testing.T) {
	_, err := notecollection.SelectDatabaseScript()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func TestSetSelectDatabaseScript(t *testing.T) {
	s, err := notecollection.SelectDatabaseScript()
	require.Nil(t, err)

	err = notecollection.SetSelectDatabaseScript(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func TestSelectDataConnections(t *testing.T) {
	_, err := notecollection.SelectDataConnections()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func TestSetSelectDataConnections(t *testing.T) {
	s, err := notecollection.SelectDataConnections()
	require.Nil(t, err)

	err = notecollection.SetSelectDataConnections(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func TestSelectDocuments(t *testing.T) {
	_, err := notecollection.SelectDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func TestSetSelectDocuments(t *testing.T) {
	s, err := notecollection.SelectDocuments()
	require.Nil(t, err)

	err = notecollection.SetSelectDocuments(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func TestSelectFolders(t *testing.T) {
	_, err := notecollection.SelectFolders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func TestSetSelectFolders(t *testing.T) {
	s, err := notecollection.SelectFolders()
	require.Nil(t, err)

	err = notecollection.SetSelectFolders(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func TestSelectForms(t *testing.T) {
	_, err := notecollection.SelectForms()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func TestSetSelectForms(t *testing.T) {
	s, err := notecollection.SelectForms()
	require.Nil(t, err)

	err = notecollection.SetSelectForms(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func TestSelectFrameSets(t *testing.T) {
	_, err := notecollection.SelectFrameSets()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func TestSetSelectFrameSets(t *testing.T) {
	s, err := notecollection.SelectFrameSets()
	require.Nil(t, err)

	err = notecollection.SetSelectFrameSets(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func TestSelectHelpAbout(t *testing.T) {
	_, err := notecollection.SelectHelpAbout()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func TestSetSelectHelpAbout(t *testing.T) {
	s, err := notecollection.SelectHelpAbout()
	require.Nil(t, err)

	err = notecollection.SetSelectHelpAbout(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func TestSelectHelpIndex(t *testing.T) {
	_, err := notecollection.SelectHelpIndex()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func TestSetSelectHelpIndex(t *testing.T) {
	s, err := notecollection.SelectHelpIndex()
	require.Nil(t, err)

	err = notecollection.SetSelectHelpIndex(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func TestSelectHelpUsing(t *testing.T) {
	_, err := notecollection.SelectHelpUsing()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func TestSetSelectHelpUsing(t *testing.T) {
	s, err := notecollection.SelectHelpUsing()
	require.Nil(t, err)

	err = notecollection.SetSelectHelpUsing(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func TestSelectIcon(t *testing.T) {
	_, err := notecollection.SelectIcon()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func TestSetSelectIcon(t *testing.T) {
	s, err := notecollection.SelectIcon()
	require.Nil(t, err)

	err = notecollection.SetSelectIcon(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func TestSelectImageResources(t *testing.T) {
	_, err := notecollection.SelectImageResources()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func TestSetSelectImageResources(t *testing.T) {
	s, err := notecollection.SelectImageResources()
	require.Nil(t, err)

	err = notecollection.SetSelectImageResources(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func TestSelectionFormula(t *testing.T) {
	_, err := notecollection.SelectionFormula()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func TestSetSelectionFormula(t *testing.T) {
	s, err := notecollection.SelectionFormula()
	require.Nil(t, err)

	err = notecollection.SetSelectionFormula(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func TestSelectJavaResources(t *testing.T) {
	_, err := notecollection.SelectJavaResources()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func TestSetSelectJavaResources(t *testing.T) {
	s, err := notecollection.SelectJavaResources()
	require.Nil(t, err)

	err = notecollection.SetSelectJavaResources(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func TestSelectMiscCodeElements(t *testing.T) {
	_, err := notecollection.SelectMiscCodeElements()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func TestSetSelectMiscCodeElements(t *testing.T) {
	s, err := notecollection.SelectMiscCodeElements()
	require.Nil(t, err)

	err = notecollection.SetSelectMiscCodeElements(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func TestSelectMiscFormatElements(t *testing.T) {
	_, err := notecollection.SelectMiscFormatElements()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func TestSetSelectMiscFormatElements(t *testing.T) {
	s, err := notecollection.SelectMiscFormatElements()
	require.Nil(t, err)

	err = notecollection.SetSelectMiscFormatElements(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func TestSelectMiscIndexElements(t *testing.T) {
	_, err := notecollection.SelectMiscIndexElements()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func TestSetSelectMiscIndexElements(t *testing.T) {
	s, err := notecollection.SelectMiscIndexElements()
	require.Nil(t, err)

	err = notecollection.SetSelectMiscIndexElements(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func TestSelectNavigators(t *testing.T) {
	_, err := notecollection.SelectNavigators()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func TestSetSelectNavigators(t *testing.T) {
	s, err := notecollection.SelectNavigators()
	require.Nil(t, err)

	err = notecollection.SetSelectNavigators(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func TestSelectOutlines(t *testing.T) {
	_, err := notecollection.SelectOutlines()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func TestSetSelectOutlines(t *testing.T) {
	s, err := notecollection.SelectOutlines()
	require.Nil(t, err)

	err = notecollection.SetSelectOutlines(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func TestSelectPages(t *testing.T) {
	_, err := notecollection.SelectPages()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func TestSetSelectPages(t *testing.T) {
	s, err := notecollection.SelectPages()
	require.Nil(t, err)

	err = notecollection.SetSelectPages(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func TestSelectProfiles(t *testing.T) {
	_, err := notecollection.SelectProfiles()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func TestSetSelectProfiles(t *testing.T) {
	s, err := notecollection.SelectProfiles()
	require.Nil(t, err)

	err = notecollection.SetSelectProfiles(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func TestSelectReplicationFormulas(t *testing.T) {
	_, err := notecollection.SelectReplicationFormulas()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func TestSetSelectReplicationFormulas(t *testing.T) {
	s, err := notecollection.SelectReplicationFormulas()
	require.Nil(t, err)

	err = notecollection.SetSelectReplicationFormulas(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func TestSelectScriptLibraries(t *testing.T) {
	_, err := notecollection.SelectScriptLibraries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func TestSetSelectScriptLibraries(t *testing.T) {
	s, err := notecollection.SelectScriptLibraries()
	require.Nil(t, err)

	err = notecollection.SetSelectScriptLibraries(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func TestSelectSharedFields(t *testing.T) {
	_, err := notecollection.SelectSharedFields()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func TestSetSelectSharedFields(t *testing.T) {
	s, err := notecollection.SelectSharedFields()
	require.Nil(t, err)

	err = notecollection.SetSelectSharedFields(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func TestSelectStyleSheetResources(t *testing.T) {
	_, err := notecollection.SelectStyleSheetResources()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func TestSetSelectStyleSheetResources(t *testing.T) {
	s, err := notecollection.SelectStyleSheetResources()
	require.Nil(t, err)

	err = notecollection.SetSelectStyleSheetResources(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func TestSelectSubforms(t *testing.T) {
	_, err := notecollection.SelectSubforms()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func TestSetSelectSubforms(t *testing.T) {
	s, err := notecollection.SelectSubforms()
	require.Nil(t, err)

	err = notecollection.SetSelectSubforms(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func TestSelectViews(t *testing.T) {
	_, err := notecollection.SelectViews()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func TestSetSelectViews(t *testing.T) {
	s, err := notecollection.SelectViews()
	require.Nil(t, err)

	err = notecollection.SetSelectViews(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func TestSinceTime(t *testing.T) {
	_, err := notecollection.SinceTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func TestSetSinceTime(t *testing.T) {
	s, err := notecollection.SinceTime()
	require.Nil(t, err)

	err = notecollection.SetSinceTime(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADD_METHOD_NOTECOLLECTION.html */
func TestAdd(t *testing.T) {
	doc := createAndSaveDoc(t)
	defer doc.Release()

	err := notecollection.Add(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BUILDCOLLECTION_METHOD_NC.html */
func TestBuildCollection(t *testing.T) {
	err := notecollection.BuildCollection()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARCOLLECTION_METHOD_NC.html */
func TestClearCollection(t *testing.T) {
	err := notecollection.ClearCollection()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTNOTEID_METHOD_NC.html */
func TestGetFirstNoteID(t *testing.T) {
	_, err := notecollection.GetFirstNoteID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_NC.html */
func TestGetNextNoteID(t *testing.T) {
	_, err := notecollection.GetNextNoteID("unknownId")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_NOTECOLLECTION.html */
func TestIntersect(t *testing.T) {
	doc := createAndSaveDoc(t)
	defer doc.Release()

	err := notecollection.Intersect(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_NOTECOLLECTION.html */
func TestRemove(t *testing.T) {
	doc := createAndSaveDoc(t)
	defer doc.Release()

	err := notecollection.Add(doc)
	require.Nil(t, err)

	err = notecollection.Remove(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLADMINNOTES_METHOD_NC.html */
func TestSelectAllAdminNotes(t *testing.T) {
	err := notecollection.SelectAllAdminNotes(true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLCODEELEMENTS_METHOD_NC.html */
func TestSelectAllCodeElements(t *testing.T) {
	err := notecollection.SelectAllCodeElements(true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDATANOTES_METHOD_NC.html */
func TestSelectAllDataNotes(t *testing.T) {
	err := notecollection.SelectAllDataNotes(true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDESIGNELEMENTS_METHOD_NC.html */
func TestSelectAllDesignElements(t *testing.T) {
	err := notecollection.SelectAllDesignElements(true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLFORMATELEMENTS_METHOD_NC.html */
func TestSelectAllFormatElements(t *testing.T) {
	err := notecollection.SelectAllFormatElements(true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLINDEXELEMENTS_METHOD_NC.html */
func TestSelectAllIndexElements(t *testing.T) {
	err := notecollection.SelectAllIndexElements(true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLNOTES_METHOD_NC.html */
func TestSelectAllNotes(t *testing.T) {
	err := notecollection.SelectAllNotes(true)
	require.Nil(t, err)
}
