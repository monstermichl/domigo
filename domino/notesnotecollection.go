/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNOTECOLLECTION_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

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
	val, err := n.Com().GetProperty("Count")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTBUILDTIME_PROPERTY_NC.html */
func (n NotesNoteCollection) LastBuildTime() (NotesDateTime, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("LastBuildTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTECOLLECTION.html */
func (n NotesNoteCollection) Parent() (NotesNoteCollection, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("Parent")
	return NewNotesNoteCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectACL() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectACL")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectACL(v Boolean) error {
	return n.Com().PutProperty("SelectACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectActions() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectActions")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectActions(v Boolean) error {
	return n.Com().PutProperty("SelectActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectAgents() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectAgents")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectAgents(v Boolean) error {
	return n.Com().PutProperty("SelectAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDatabaseScript() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectDatabaseScript")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDatabaseScript(v Boolean) error {
	return n.Com().PutProperty("SelectDatabaseScript", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDataConnections() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectDataConnections")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDataConnections(v Boolean) error {
	return n.Com().PutProperty("SelectDataConnections", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDocuments() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectDocuments")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDocuments(v Boolean) error {
	return n.Com().PutProperty("SelectDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFolders() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectFolders")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFolders(v Boolean) error {
	return n.Com().PutProperty("SelectFolders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectForms() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectForms")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectForms(v Boolean) error {
	return n.Com().PutProperty("SelectForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFrameSets() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectFrameSets")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFrameSets(v Boolean) error {
	return n.Com().PutProperty("SelectFrameSets", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpAbout() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectHelpAbout")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpAbout(v Boolean) error {
	return n.Com().PutProperty("SelectHelpAbout", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpIndex() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectHelpIndex")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpIndex(v Boolean) error {
	return n.Com().PutProperty("SelectHelpIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpUsing() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectHelpUsing")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpUsing(v Boolean) error {
	return n.Com().PutProperty("SelectHelpUsing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectIcon() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectIcon")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectIcon(v Boolean) error {
	return n.Com().PutProperty("SelectIcon", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectImageResources() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectImageResources")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectImageResources(v Boolean) error {
	return n.Com().PutProperty("SelectImageResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectionFormula() (String, error) {
	val, err := n.Com().GetProperty("SelectionFormula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectionFormula(v String) error {
	return n.Com().PutProperty("SelectionFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectJavaResources() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectJavaResources")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectJavaResources(v Boolean) error {
	return n.Com().PutProperty("SelectJavaResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscCodeElements() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectMiscCodeElements")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscCodeElements(v Boolean) error {
	return n.Com().PutProperty("SelectMiscCodeElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscFormatElements() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectMiscFormatElements")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscFormatElements(v Boolean) error {
	return n.Com().PutProperty("SelectMiscFormatElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscIndexElements() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectMiscIndexElements")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscIndexElements(v Boolean) error {
	return n.Com().PutProperty("SelectMiscIndexElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectNavigators() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectNavigators")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectNavigators(v Boolean) error {
	return n.Com().PutProperty("SelectNavigators", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectOutlines() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectOutlines")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectOutlines(v Boolean) error {
	return n.Com().PutProperty("SelectOutlines", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectPages() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectPages")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectPages(v Boolean) error {
	return n.Com().PutProperty("SelectPages", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectProfiles() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectProfiles")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectProfiles(v Boolean) error {
	return n.Com().PutProperty("SelectProfiles", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectReplicationFormulas() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectReplicationFormulas")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectReplicationFormulas(v Boolean) error {
	return n.Com().PutProperty("SelectReplicationFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectScriptLibraries() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectScriptLibraries")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectScriptLibraries(v Boolean) error {
	return n.Com().PutProperty("SelectScriptLibraries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSharedFields() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectSharedFields")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSharedFields(v Boolean) error {
	return n.Com().PutProperty("SelectSharedFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectStyleSheetResources() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectStyleSheetResources")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectStyleSheetResources(v Boolean) error {
	return n.Com().PutProperty("SelectStyleSheetResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSubforms() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectSubforms")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSubforms(v Boolean) error {
	return n.Com().PutProperty("SelectSubforms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectViews() (Boolean, error) {
	val, err := n.Com().GetProperty("SelectViews")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectViews(v Boolean) error {
	return n.Com().PutProperty("SelectViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SinceTime() (NotesDateTime, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("SinceTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSinceTime(v NotesDateTime) error {
	return n.Com().PutProperty("SinceTime", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADD_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Add(additionSpecifier Variant) error {
	_, err := n.Com().CallMethod("Add", additionSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BUILDCOLLECTION_METHOD_NC.html */
func (n NotesNoteCollection) BuildCollection() error {
	_, err := n.Com().CallMethod("BuildCollection")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARCOLLECTION_METHOD_NC.html */
func (n NotesNoteCollection) ClearCollection() error {
	_, err := n.Com().CallMethod("ClearCollection")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetFirstNoteID() (String, error) {
	val, err := n.Com().CallMethod("GetFirstNoteID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetNextNoteID(noteID String) (String, error) {
	val, err := n.Com().CallMethod("GetNextNoteID", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Intersect(intersectionSpecifier Variant) error {
	_, err := n.Com().CallMethod("Intersect", intersectionSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Remove(removalSpecifier Variant) error {
	_, err := n.Com().CallMethod("Remove", removalSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLADMINNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllAdminNotes(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllAdminNotes", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLCODEELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllCodeElements(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllCodeElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDATANOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDataNotes(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllDataNotes", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDESIGNELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDesignElements(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllDesignElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLFORMATELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllFormatElements(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllFormatElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLINDEXELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllIndexElements(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllIndexElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllNotes(selectorValue Boolean) error {
	_, err := n.Com().CallMethod("SelectAllNotes", selectorValue)
	return err
}
