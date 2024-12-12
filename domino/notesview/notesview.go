package notesview

import (
	"domigo/domino"
	"domigo/domino/notesdocument"
	"domigo/domino/notesdocumentcollection"
	"domigo/domino/notesrichtextitem"
	"domigo/domino/notesviewcolumn"
	"domigo/domino/notesviewentry"
	"domigo/domino/notesviewentrycollection"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesView struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesView {
	return NotesView{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY.html */
/* Moved from NotesDocument. */
func NotesDocumentParentView(d notesdocument.NotesDocument) (NotesView, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("ParentView")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDDOCLINK_METHOD.html */
/* Moved from NotesRichTextItem. */
type appendDocLinkParams struct {
	HotSpotText *domino.String
}

type appendDocLinkParam func(*appendDocLinkParams)

func WithAppendDocLinkHotSpotText(HotSpotText domino.String) appendDocLinkParam {
	return func(c *appendDocLinkParams) {
		c.HotSpotText = &HotSpotText
	}
}

func NotesRichTextItemAppendDocLink(r notesrichtextitem.NotesRichTextItem, linkTo NotesView, comment domino.String, params ...appendDocLinkParam) error {
	paramsStruct := &appendDocLinkParams{}
	paramsOrdered := []interface{}{linkTo.Com().Dispatch(), comment}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.HotSpotText != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.HotSpotText)
	}
	_, err := r.Com().CallMethod("AppendDocLink", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTESVIEWCOLUMN_CLASS.html */
/* Moved from NotesViewColumn. */
func NotesViewColumnParent(c notesviewcolumn.NotesViewColumn) (NotesView, error) {
	dispatchPtr, err := c.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_3044.html */
/* Moved from NotesViewEntry. */
func NotesViewEntryParent(v notesviewentry.NotesViewEntry) (NotesView, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_9184.html */
/* Moved from NotesViewEntryCollection. */
func  NotesViewEntryCollectionParent(v notesviewentrycollection.NotesViewEntryCollection) (NotesView, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func (v NotesView) Aliases() (domino.String, error) {
	val, err := v.Com().GetProperty("Aliases")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIES_PROPERTY_6084.html */
func (v NotesView) AllEntries() (notesviewentrycollection.NotesViewEntryCollection, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("AllEntries")
	return notesviewentrycollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func (v NotesView) AutoUpdate() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("AutoUpdate")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func (v NotesView) SetAutoUpdate(val domino.Boolean) error {
	return v.Com().PutProperty("AutoUpdate", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func (v NotesView) BackgroundColor() (domino.Integer, error) {
	val, err := v.Com().GetProperty("BackgroundColor")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func (v NotesView) SetBackgroundColor(val domino.Integer) error {
	return v.Com().PutProperty("BackgroundColor", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNCOUNT_PROPERTY_7753.html */
func (v NotesView) ColumnCount() (domino.Integer, error) {
	val, err := v.Com().GetProperty("ColumnCount")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNNAMES_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) ColumnNames() ([]domino.String, error) {
	vals, err := v.Com().GetArrayProperty("ColumnNames")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNS_PROPERTY.html */
func (v NotesView) Columns() (notesviewcolumn.NotesViewColumn, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Columns")
	return notesviewcolumn.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_VIEW.html */
func (v NotesView) Created() (domino.Time, error) {
	val, err := v.Com().GetProperty("Created")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCOUNT_PROPERTY_VIEW.html */
func (v NotesView) EntryCount() (domino.Long, error) {
	val, err := v.Com().GetProperty("EntryCount")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERLINES_PROPERTY_1209.html */
func (v NotesView) HeaderLines() (domino.Integer, error) {
	val, err := v.Com().GetProperty("HeaderLines")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) HttpURL() (domino.String, error) {
	val, err := v.Com().GetProperty("HttpURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCALENDAR_PROPERTY.html */
func (v NotesView) IsCalendar() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsCalendar")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORIZED_PROPERTY_2814.html */
func (v NotesView) IsCategorized() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsCategorized")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1242.html */
func (v NotesView) IsConflict() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsConflict")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func (v NotesView) IsDefaultView() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsDefaultView")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func (v NotesView) SetIsDefaultView(val domino.Boolean) error {
	return v.Com().PutProperty("IsDefaultView", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFOLDER_PROPERTY.html */
func (v NotesView) IsFolder() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsFolder")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIERARCHICAL_PROPERTY_7038.html */
func (v NotesView) IsHierarchical() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsHierarchical")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMODIFIED_PROPERTY_6416.html */
func (v NotesView) IsModified() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsModified")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_VIEW.html */
func (v NotesView) IsPrivate() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsPrivate")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func (v NotesView) IsProhibitDesignRefresh() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsProhibitDesignRefresh")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func (v NotesView) SetIsProhibitDesignRefresh(val domino.Boolean) error {
	return v.Com().PutProperty("IsProhibitDesignRefresh", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_VIEW.html */
func (v NotesView) LastModified() (domino.Time, error) {
	val, err := v.Com().GetProperty("LastModified")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_VIEW.html */
func (v NotesView) LockHolders() (domino.String, error) {
	val, err := v.Com().GetProperty("LockHolders")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func (v NotesView) Name() (domino.String, error) {
	val, err := v.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func (v NotesView) SetName(val domino.String) error {
	return v.Com().PutProperty("Name", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) NotesURL() (domino.String, error) {
	val, err := v.Com().GetProperty("NotesURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func (v NotesView) ProtectReaders() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("ProtectReaders")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func (v NotesView) SetProtectReaders(val domino.Boolean) error {
	return v.Com().PutProperty("ProtectReaders", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func (v NotesView) Readers() (domino.String, error) {
	val, err := v.Com().GetProperty("Readers")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func (v NotesView) SetReaders(val domino.String) error {
	return v.Com().PutProperty("Readers", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROWLINES_PROPERTY_4578.html */
func (v NotesView) RowLines() (domino.Integer, error) {
	val, err := v.Com().GetProperty("RowLines")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func (v NotesView) SelectionFormula() (domino.String, error) {
	val, err := v.Com().GetProperty("SelectionFormula")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func (v NotesView) SetSelectionFormula(val domino.String) error {
	return v.Com().PutProperty("SelectionFormula", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func (v NotesView) Spacing() (domino.Integer, error) {
	val, err := v.Com().GetProperty("Spacing")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func (v NotesView) SetSpacing(val domino.Integer) error {
	return v.Com().PutProperty("Spacing", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TOPLEVELENTRYCOUNT_PROPERTY_6487.html */
func (v NotesView) TopLevelEntryCount() (domino.Integer, error) {
	val, err := v.Com().GetProperty("TopLevelEntryCount")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_VIEW.html */
func (v NotesView) UniversalID() (domino.String, error) {
	val, err := v.Com().GetProperty("UniversalID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWINHERITEDNAME_PROPERTY_VIEW.html */
func (v NotesView) ViewInheritedName() (domino.String, error) {
	val, err := v.Com().GetProperty("ViewInheritedName")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_VIEW.html */
func (v NotesView) Clear() error {
	_, err := v.Com().CallMethod("Clear")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYCOLUMN_METHOD_VIEW.html */
func (v NotesView) CopyColumn(sourceColumn domino.Integer, destinationIndex domino.Long) (notesviewcolumn.NotesViewColumn, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CopyColumn", sourceColumn, destinationIndex)
	return notesviewcolumn.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLUMN_METHOD_VIEW.html */
func (v NotesView) CreateColumn(position domino.Long, columnName domino.String, formula domino.String) (notesviewcolumn.NotesViewColumn, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateColumn", position, columnName, formula)
	return notesviewcolumn.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_VIEW.html */
func (v NotesView) FTSearch(query domino.String, maxDocs domino.Integer) (domino.Long, error) {
	val, err := v.Com().CallMethod("FTSearch", query, maxDocs)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLDOCUMENTSBYKEY_METHOD.html */
func (v NotesView) GetAllDocumentsByKey(keyArray []domino.Integer, exactMatch domino.Boolean) (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetAllDocumentsByKey", keyArray, exactMatch)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIESBYKEY_METHOD_9837.html */
func (v NotesView) GetAllEntriesByKey(keyArray []domino.Integer, exactMatch domino.Boolean) (notesviewentrycollection.NotesViewEntryCollection, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetAllEntriesByKey", keyArray, exactMatch)
	return notesviewentrycollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLREADENTRIES.html */
func (v NotesView) GetAllReadEntries(username domino.String) (notesviewentrycollection.NotesViewEntryCollection, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetAllReadEntries", username)
	return notesviewentrycollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLUNREADENTRIES.html */
func (v NotesView) GetAllUnreadEntries(username domino.String) (notesviewentrycollection.NotesViewEntryCollection, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetAllUnreadEntries", username)
	return notesviewentrycollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD.html */
func (v NotesView) GetChild(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetChild", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCOLUMN_METHOD_NOTESVIEW_CLASS.html */
func (v NotesView) GetColumn(columnNumber domino.Long) (notesviewcolumn.NotesViewColumn, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetColumn", columnNumber)
	return notesviewcolumn.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYKEY_METHOD.html */
func (v NotesView) GetDocumentByKey(keyArray []domino.Integer, exactMatch domino.Boolean) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetDocumentByKey", keyArray, exactMatch)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYKEY_METHOD_3846.html */
func (v NotesView) GetEntryByKey(keyArray []domino.Integer, exactMatch domino.Boolean) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetEntryByKey", keyArray, exactMatch)
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetFirstDocument() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetFirstDocument")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetLastDocument() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetLastDocument")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetNextDocument(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextDocument", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD.html */
func (v NotesView) GetNextSibling(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextSibling", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetNthDocument(index domino.Long) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNthDocument", index)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTDOCUMENT_METHOD.html */
func (v NotesView) GetParentDocument(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetParentDocument", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetPrevDocument(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevDocument", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD.html */
func (v NotesView) GetPrevSibling(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevSibling", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_VIEW.html */
func (v NotesView) Lock(name []domino.String, provisionalOK domino.Boolean) (domino.Boolean, error) {
	val, err := v.Com().CallMethod("Lock", name, provisionalOK)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_VIEW.html */
func (v NotesView) LockProvisional(name []domino.String) (domino.Boolean, error) {
	val, err := v.Com().CallMethod("LockProvisional", name)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEW.html */
func (v NotesView) MarkAllRead(username domino.String) error {
	_, err := v.Com().CallMethod("MarkAllRead", username)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEW.html */
func (v NotesView) MarkAllUnread(username domino.String) error {
	_, err := v.Com().CallMethod("MarkAllUnread", username)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESH_METHOD_VIEW.html */
func (v NotesView) Refresh() error {
	_, err := v.Com().CallMethod("Refresh")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_VIEW.html */
func (v NotesView) Remove() error {
	_, err := v.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTVIEW_METHOD_VIEW.html */
func (v NotesView) ResortView(columnName domino.String, ascendingFlag domino.Boolean) error {
	_, err := v.Com().CallMethod("ResortView", columnName, ascendingFlag)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVECOLUMN_METHOD_VIEW.html */
func (v NotesView) RemoveColumn(columnNameOrIndex domino.Long) error {
	_, err := v.Com().CallMethod("RemoveColumn", columnNameOrIndex)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETALIASES_METHOD_VIEW.html */
func (v NotesView) SetAliases(alias domino.String) (domino.Long, error) {
	val, err := v.Com().CallMethod("SetAliases", alias)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_VIEW.html */
func (v NotesView) UnLock() error {
	_, err := v.Com().CallMethod("UnLock")
	return err
}
