package notesnotecollection

import (
	"domigo/domino"
	"domigo/domino/notesdatetime"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesNoteCollection struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesNoteCollection {
	return NotesNoteCollection{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_NOTECOLLECTION.html */
func (n NotesNoteCollection) Count() (domino.Long, error) {
	val, err := n.Com().GetProperty("Count")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTBUILDTIME_PROPERTY_NC.html */
func (n NotesNoteCollection) LastBuildTime() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("LastBuildTime")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectACL() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectACL")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACL_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectACL(v domino.Boolean) error {
	return n.Com().PutProperty("SelectACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectActions() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectActions")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTACTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectActions(v domino.Boolean) error {
	return n.Com().PutProperty("SelectActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectAgents() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectAgents")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTAGENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectAgents(v domino.Boolean) error {
	return n.Com().PutProperty("SelectAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDatabaseScript() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectDatabaseScript")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATABASESCRIPT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDatabaseScript(v domino.Boolean) error {
	return n.Com().PutProperty("SelectDatabaseScript", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDataConnections() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectDataConnections")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDATACONNECTIONS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDataConnections(v domino.Boolean) error {
	return n.Com().PutProperty("SelectDataConnections", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectDocuments() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectDocuments")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTDOCUMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectDocuments(v domino.Boolean) error {
	return n.Com().PutProperty("SelectDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFolders() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectFolders")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFOLDERS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFolders(v domino.Boolean) error {
	return n.Com().PutProperty("SelectFolders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectForms() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectForms")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectForms(v domino.Boolean) error {
	return n.Com().PutProperty("SelectForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectFrameSets() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectFrameSets")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTFRAMESETS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectFrameSets(v domino.Boolean) error {
	return n.Com().PutProperty("SelectFrameSets", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpAbout() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectHelpAbout")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPABOUT_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpAbout(v domino.Boolean) error {
	return n.Com().PutProperty("SelectHelpAbout", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpIndex() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectHelpIndex")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPINDEX_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpIndex(v domino.Boolean) error {
	return n.Com().PutProperty("SelectHelpIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectHelpUsing() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectHelpUsing")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTHELPUSING_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectHelpUsing(v domino.Boolean) error {
	return n.Com().PutProperty("SelectHelpUsing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectIcon() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectIcon")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTICON_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectIcon(v domino.Boolean) error {
	return n.Com().PutProperty("SelectIcon", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectImageResources() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectImageResources")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIMAGERESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectImageResources(v domino.Boolean) error {
	return n.Com().PutProperty("SelectImageResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectionFormula() (domino.String, error) {
	val, err := n.Com().GetProperty("SelectionFormula")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectionFormula(v domino.String) error {
	return n.Com().PutProperty("SelectionFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectJavaResources() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectJavaResources")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTJAVARESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectJavaResources(v domino.Boolean) error {
	return n.Com().PutProperty("SelectJavaResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscCodeElements() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectMiscCodeElements")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCCODEELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscCodeElements(v domino.Boolean) error {
	return n.Com().PutProperty("SelectMiscCodeElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscFormatElements() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectMiscFormatElements")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCFORMATELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscFormatElements(v domino.Boolean) error {
	return n.Com().PutProperty("SelectMiscFormatElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectMiscIndexElements() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectMiscIndexElements")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTMISCINDEXELEMENTS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectMiscIndexElements(v domino.Boolean) error {
	return n.Com().PutProperty("SelectMiscIndexElements", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectNavigators() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectNavigators")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTNAVIGATORS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectNavigators(v domino.Boolean) error {
	return n.Com().PutProperty("SelectNavigators", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectOutlines() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectOutlines")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTOUTLINES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectOutlines(v domino.Boolean) error {
	return n.Com().PutProperty("SelectOutlines", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectPages() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectPages")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPAGES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectPages(v domino.Boolean) error {
	return n.Com().PutProperty("SelectPages", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectProfiles() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectProfiles")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTPROFILES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectProfiles(v domino.Boolean) error {
	return n.Com().PutProperty("SelectProfiles", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectReplicationFormulas() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectReplicationFormulas")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTREPLICATIONFORMULAS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectReplicationFormulas(v domino.Boolean) error {
	return n.Com().PutProperty("SelectReplicationFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectScriptLibraries() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectScriptLibraries")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSCRIPTLIBRARIES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectScriptLibraries(v domino.Boolean) error {
	return n.Com().PutProperty("SelectScriptLibraries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSharedFields() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectSharedFields")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSHAREDFIELDS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSharedFields(v domino.Boolean) error {
	return n.Com().PutProperty("SelectSharedFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectStyleSheetResources() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectStyleSheetResources")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSTYLESHEETRESOURCES_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectStyleSheetResources(v domino.Boolean) error {
	return n.Com().PutProperty("SelectStyleSheetResources", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectSubforms() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectSubforms")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTSUBFORMS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectSubforms(v domino.Boolean) error {
	return n.Com().PutProperty("SelectSubforms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SelectViews() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("SelectViews")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTVIEWS_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSelectViews(v domino.Boolean) error {
	return n.Com().PutProperty("SelectViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SinceTime() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("SinceTime")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SINCETIME_PROPERTY_NC.html */
func (n NotesNoteCollection) SetSinceTime(v notesdatetime.NotesDateTime) error {
	return n.Com().PutProperty("SinceTime", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADD_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Add(additionSpecifier domino.Variant) error {
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
func (n NotesNoteCollection) GetFirstNoteID() (domino.String, error) {
	val, err := n.Com().CallMethod("GetFirstNoteID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_NC.html */
func (n NotesNoteCollection) GetNextNoteID(noteID domino.String) (domino.String, error) {
	val, err := n.Com().CallMethod("GetNextNoteID", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Intersect(intersectionSpecifier domino.Variant) error {
	_, err := n.Com().CallMethod("Intersect", intersectionSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_NOTECOLLECTION.html */
func (n NotesNoteCollection) Remove(removalSpecifier domino.Variant) error {
	_, err := n.Com().CallMethod("Remove", removalSpecifier)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLADMINNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllAdminNotes(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllAdminNotes", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLCODEELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllCodeElements(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllCodeElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDATANOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDataNotes(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllDataNotes", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLDESIGNELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllDesignElements(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllDesignElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLFORMATELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllFormatElements(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllFormatElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLINDEXELEMENTS_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllIndexElements(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllIndexElements", selectorValue)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTALLNOTES_METHOD_NC.html */
func (n NotesNoteCollection) SelectAllNotes(selectorValue domino.Boolean) error {
	_, err := n.Com().CallMethod("SelectAllNotes", selectorValue)
	return err
}
