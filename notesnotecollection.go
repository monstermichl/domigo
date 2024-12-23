/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNOTECOLLECTION_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesNoteCollection struct {
	NotesStruct
}

func NewNotesNoteCollection(dispatchPtr *ole.IDispatch) NotesNoteCollection {
	return NotesNoteCollection{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_NOTECOLLECTION.html */
func (n NotesNoteCollection) Count() (Long, error) {
	val, err := getComProperty(n, "Count")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTBUILDTIME_PROPERTY_NC.html */
func (n NotesNoteCollection) LastBuildTime() (NotesDateTime, error) {
	dispatchPtr, err := getComObjectProperty(n, "LastBuildTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTECOLLECTION.html */
func (n NotesNoteCollection) Parent() (NotesNoteCollection, error) {
	dispatchPtr, err := getComObjectProperty(n, "Parent")
	return NewNotesNoteCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectACL() (Boolean, error) {
	val, err := getComProperty(n, "SelectACL")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectACL(v Boolean) error {
	return putComProperty(n, "SelectACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectActions() (Boolean, error) {
	val, err := getComProperty(n, "SelectActions")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectActions(v Boolean) error {
	return putComProperty(n, "SelectActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectAgents() (Boolean, error) {
	val, err := getComProperty(n, "SelectAgents")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectAgents(v Boolean) error {
	return putComProperty(n, "SelectAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDatabaseScript() (Boolean, error) {
	val, err := getComProperty(n, "SelectDatabaseScript")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDatabaseScript(v Boolean) error {
	return putComProperty(n, "SelectDatabaseScript", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDataConnections() (Boolean, error) {
	val, err := getComProperty(n, "SelectDataConnections")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDataConnections(v Boolean) error {
	return putComProperty(n, "SelectDataConnections", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDocuments() (Boolean, error) {
	val, err := getComProperty(n, "SelectDocuments")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDocuments(v Boolean) error {
	return putComProperty(n, "SelectDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFolders() (Boolean, error) {
	val, err := getComProperty(n, "SelectFolders")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFolders(v Boolean) error {
	return putComProperty(n, "SelectFolders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectForms() (Boolean, error) {
	val, err := getComProperty(n, "SelectForms")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectForms(v Boolean) error {
	return putComProperty(n, "SelectForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFrameSets() (Boolean, error) {
	val, err := getComProperty(n, "SelectFrameSets")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFrameSets(v Boolean) error {
	return putComProperty(n, "SelectFrameSets", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpAbout() (Boolean, error) {
	val, err := getComProperty(n, "SelectHelpAbout")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpAbout(v Boolean) error {
	return putComProperty(n, "SelectHelpAbout", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpIndex() (Boolean, error) {
	val, err := getComProperty(n, "SelectHelpIndex")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpIndex(v Boolean) error {
	return putComProperty(n, "SelectHelpIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpUsing() (Boolean, error) {
	val, err := getComProperty(n, "SelectHelpUsing")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpUsing(v Boolean) error {
	return putComProperty(n, "SelectHelpUsing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectIcon() (Boolean, error) {
	val, err := getComProperty(n, "SelectIcon")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectIcon(v Boolean) error {
	return putComProperty(n, "SelectIcon", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectImageResources() (Boolean, error) {
	val, err := getComProperty(n, "SelectImageResources")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectImageResources(v Boolean) error {
	return putComProperty(n, "SelectImageResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectionFormula() (String, error) {
	val, err := getComProperty(n, "SelectionFormula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectionFormula(v String) error {
	return putComProperty(n, "SelectionFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectJavaResources() (Boolean, error) {
	val, err := getComProperty(n, "SelectJavaResources")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectJavaResources(v Boolean) error {
	return putComProperty(n, "SelectJavaResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscCodeElements() (Boolean, error) {
	val, err := getComProperty(n, "SelectMiscCodeElements")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscCodeElements(v Boolean) error {
	return putComProperty(n, "SelectMiscCodeElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscFormatElements() (Boolean, error) {
	val, err := getComProperty(n, "SelectMiscFormatElements")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscFormatElements(v Boolean) error {
	return putComProperty(n, "SelectMiscFormatElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscIndexElements() (Boolean, error) {
	val, err := getComProperty(n, "SelectMiscIndexElements")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscIndexElements(v Boolean) error {
	return putComProperty(n, "SelectMiscIndexElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectNavigators() (Boolean, error) {
	val, err := getComProperty(n, "SelectNavigators")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectNavigators(v Boolean) error {
	return putComProperty(n, "SelectNavigators", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectOutlines() (Boolean, error) {
	val, err := getComProperty(n, "SelectOutlines")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectOutlines(v Boolean) error {
	return putComProperty(n, "SelectOutlines", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectPages() (Boolean, error) {
	val, err := getComProperty(n, "SelectPages")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectPages(v Boolean) error {
	return putComProperty(n, "SelectPages", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectProfiles() (Boolean, error) {
	val, err := getComProperty(n, "SelectProfiles")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectProfiles(v Boolean) error {
	return putComProperty(n, "SelectProfiles", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectReplicationFormulas() (Boolean, error) {
	val, err := getComProperty(n, "SelectReplicationFormulas")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectReplicationFormulas(v Boolean) error {
	return putComProperty(n, "SelectReplicationFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectScriptLibraries() (Boolean, error) {
	val, err := getComProperty(n, "SelectScriptLibraries")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectScriptLibraries(v Boolean) error {
	return putComProperty(n, "SelectScriptLibraries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSharedFields() (Boolean, error) {
	val, err := getComProperty(n, "SelectSharedFields")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSharedFields(v Boolean) error {
	return putComProperty(n, "SelectSharedFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectStyleSheetResources() (Boolean, error) {
	val, err := getComProperty(n, "SelectStyleSheetResources")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectStyleSheetResources(v Boolean) error {
	return putComProperty(n, "SelectStyleSheetResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSubforms() (Boolean, error) {
	val, err := getComProperty(n, "SelectSubforms")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSubforms(v Boolean) error {
	return putComProperty(n, "SelectSubforms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectViews() (Boolean, error) {
	val, err := getComProperty(n, "SelectViews")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectViews(v Boolean) error {
	return putComProperty(n, "SelectViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SinceTime() (NotesDateTime, error) {
	dispatchPtr, err := getComObjectProperty(n, "SinceTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSinceTime(v NotesDateTime) error {
	return putComProperty(n, "SinceTime", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADD_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Add(additionSpecifier Variant) error {
	_, err := callComMethod(n, "Add", additionSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BUILDCOLLECTION_METHOD_NC.html */
func (n NotesNoteCollection) BuildCollection() error {
	_, err := callComMethod(n, "BuildCollection")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARCOLLECTION_METHOD_NC.html */
func (n NotesNoteCollection) ClearCollection() error {
	_, err := callComMethod(n, "ClearCollection")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetFirstNoteID() (String, error) {
	val, err := callComMethod(n, "GetFirstNoteID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetNextNoteID(noteID String) (String, error) {
	val, err := callComMethod(n, "GetNextNoteID", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Intersect(intersectionSpecifier Variant) error {
	_, err := callComMethod(n, "Intersect", intersectionSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Remove(removalSpecifier Variant) error {
	_, err := callComMethod(n, "Remove", removalSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLADMINNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllAdminNotes(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllAdminNotes", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLCODEELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllCodeElements(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllCodeElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDATANOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDataNotes(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllDataNotes", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDESIGNELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDesignElements(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllDesignElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLFORMATELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllFormatElements(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllFormatElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLINDEXELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllIndexElements(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllIndexElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllNotes(selectorValue Boolean) error {
	_, err := callComMethod(n, "SelectAllNotes", selectorValue)
	return err
}
