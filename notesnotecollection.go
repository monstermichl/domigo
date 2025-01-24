/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNOTECOLLECTION_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
	"github.com/monstermichl/domigo/internal/helpers"
)

type NotesNoteCollection struct {
	NotesStruct
}

func newNotesNoteCollection(dispatchPtr *ole.IDispatch) NotesNoteCollection {
	return NotesNoteCollection{newNotesStruct(dispatchPtr)}
}

func (n NotesNoteCollection) checkCombinableTypes(val any) error {
	return helpers.CheckTypeNames(val, []string{"number", "string", "NotesDocument", "NotesDocumentCollection", "NotesViewEntry", "NotesViewEntryCollection"})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_NOTECOLLECTION.html */
func (n NotesNoteCollection) Count() (Long, error) {
	return getComProperty[Long](n, "Count")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTBUILDTIME_PROPERTY_NC.html */
func (n NotesNoteCollection) LastBuildTime() (NotesDateTime, error) {
	return getComObjectProperty(n, newNotesDateTime, "LastBuildTime")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTECOLLECTION.html */
func (n NotesNoteCollection) Parent() (NotesNoteCollection, error) {
	return getComObjectProperty(n, newNotesNoteCollection, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectACL() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectACL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectACL(v Boolean) error {
	return putComProperty(n, "SelectACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectActions() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectActions")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectActions(v Boolean) error {
	return putComProperty(n, "SelectActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectAgents() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectAgents")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectAgents(v Boolean) error {
	return putComProperty(n, "SelectAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDatabaseScript() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectDatabaseScript")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDatabaseScript(v Boolean) error {
	return putComProperty(n, "SelectDatabaseScript", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDataConnections() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectDataConnections")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDataConnections(v Boolean) error {
	return putComProperty(n, "SelectDataConnections", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDocuments() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectDocuments")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDocuments(v Boolean) error {
	return putComProperty(n, "SelectDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFolders() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectFolders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFolders(v Boolean) error {
	return putComProperty(n, "SelectFolders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectForms() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectForms")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectForms(v Boolean) error {
	return putComProperty(n, "SelectForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFrameSets() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectFrameSets")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFrameSets(v Boolean) error {
	return putComProperty(n, "SelectFrameSets", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpAbout() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectHelpAbout")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpAbout(v Boolean) error {
	return putComProperty(n, "SelectHelpAbout", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpIndex() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectHelpIndex")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpIndex(v Boolean) error {
	return putComProperty(n, "SelectHelpIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpUsing() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectHelpUsing")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpUsing(v Boolean) error {
	return putComProperty(n, "SelectHelpUsing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectIcon() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectIcon")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectIcon(v Boolean) error {
	return putComProperty(n, "SelectIcon", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectImageResources() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectImageResources")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectImageResources(v Boolean) error {
	return putComProperty(n, "SelectImageResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectionFormula() (String, error) {
	return getComProperty[String](n, "SelectionFormula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectionFormula(v String) error {
	return putComProperty(n, "SelectionFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectJavaResources() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectJavaResources")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectJavaResources(v Boolean) error {
	return putComProperty(n, "SelectJavaResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscCodeElements() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectMiscCodeElements")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscCodeElements(v Boolean) error {
	return putComProperty(n, "SelectMiscCodeElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscFormatElements() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectMiscFormatElements")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscFormatElements(v Boolean) error {
	return putComProperty(n, "SelectMiscFormatElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscIndexElements() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectMiscIndexElements")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscIndexElements(v Boolean) error {
	return putComProperty(n, "SelectMiscIndexElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectNavigators() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectNavigators")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectNavigators(v Boolean) error {
	return putComProperty(n, "SelectNavigators", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectOutlines() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectOutlines")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectOutlines(v Boolean) error {
	return putComProperty(n, "SelectOutlines", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectPages() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectPages")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectPages(v Boolean) error {
	return putComProperty(n, "SelectPages", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectProfiles() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectProfiles")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectProfiles(v Boolean) error {
	return putComProperty(n, "SelectProfiles", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectReplicationFormulas() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectReplicationFormulas")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectReplicationFormulas(v Boolean) error {
	return putComProperty(n, "SelectReplicationFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectScriptLibraries() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectScriptLibraries")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectScriptLibraries(v Boolean) error {
	return putComProperty(n, "SelectScriptLibraries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSharedFields() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectSharedFields")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSharedFields(v Boolean) error {
	return putComProperty(n, "SelectSharedFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectStyleSheetResources() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectStyleSheetResources")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectStyleSheetResources(v Boolean) error {
	return putComProperty(n, "SelectStyleSheetResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSubforms() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectSubforms")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSubforms(v Boolean) error {
	return putComProperty(n, "SelectSubforms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectViews() (Boolean, error) {
	return getComProperty[Boolean](n, "SelectViews")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectViews(v Boolean) error {
	return putComProperty(n, "SelectViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SinceTime() (NotesDateTime, error) {
	return getComObjectProperty(n, newNotesDateTime, "SinceTime")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSinceTime(v NotesDateTime) error {
	return putComProperty(n, "SinceTime", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADD_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Add(additionSpecifier any) error {
	err := n.checkCombinableTypes(additionSpecifier)

	if err != nil {
		return err
	}
	return callComVoidMethod(n, "Add", additionSpecifier)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BUILDCOLLECTION_METHOD_NC.html */
func (n NotesNoteCollection) BuildCollection() error {
	return callComVoidMethod(n, "BuildCollection")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARCOLLECTION_METHOD_NC.html */
func (n NotesNoteCollection) ClearCollection() error {
	return callComVoidMethod(n, "ClearCollection")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetFirstNoteID() (String, error) {
	return callComMethod[String](n, "GetFirstNoteID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetNextNoteID(noteID String) (String, error) {
	return callComMethod[String](n, "GetNextNoteID", noteID)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Intersect(intersectionSpecifier any) error {
	err := n.checkCombinableTypes(intersectionSpecifier)

	if err != nil {
		return err
	}
	return callComVoidMethod(n, "Intersect", intersectionSpecifier)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Remove(removalSpecifier any) error {
	err := n.checkCombinableTypes(removalSpecifier)

	if err != nil {
		return err
	}
	return callComVoidMethod(n, "Remove", removalSpecifier)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLADMINNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllAdminNotes(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllAdminNotes", selectorValue)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLCODEELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllCodeElements(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllCodeElements", selectorValue)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDATANOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDataNotes(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllDataNotes", selectorValue)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDESIGNELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDesignElements(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllDesignElements", selectorValue)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLFORMATELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllFormatElements(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllFormatElements", selectorValue)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLINDEXELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllIndexElements(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllIndexElements", selectorValue)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllNotes(selectorValue Boolean) error {
	return callComVoidMethod(n, "SelectAllNotes", selectorValue)
}
