/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATABASE_CLASS.html */
package notesdatabase

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesacl"
	"domigo/domino/notesagent"
	"domigo/domino/notesdatetime"
	"domigo/domino/notesdocument"
	"domigo/domino/notesdocumentcollection"
	"domigo/domino/notesdominoquery"
	"domigo/domino/notesform"
	"domigo/domino/notesnotecollection"
	"domigo/domino/notesoutline"
	"domigo/domino/notesoutlineentry"
	"domigo/domino/notesqueryresultsprocessor"
	"domigo/domino/notesreplication"
	"domigo/domino/notesrichtextitem"
	"domigo/domino/notesuserid"
	"domigo/domino/notesview"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDatabase struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDatabase {
	return NotesDatabase{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDATABASE_PROPERTY.html */
/* Moved from NotesDocument */
func NotesDocumentParentDatabase(d notesdocument.NotesDocument) (NotesDatabase, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("ParentDatabase")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYTODATABASE_METHOD.html */
/* Moved from NotesDocument */
func NotesDocumentCopyToDatabase(d notesdocument.NotesDocument, notesDatabase NotesDatabase) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CopyToDatabase", notesDatabase.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ACL.html */
/* Moved from NotesACL. */
func NotesACLParent(a notesacl.NotesACL) (NotesDatabase, error) {
	dispatchPtr, err := a.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_COLLECTION.html */
/* Moved from NotesDocumentCollection. */
func NotesDocumentCollectionParent(d notesdocumentcollection.NotesDocumentCollection) (NotesDatabase, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_FORM_COM.html */
/* Moved from NotesForm. */
func NotesFormParent(f notesform.NotesForm) (NotesDatabase, error) {
	dispatchPtr, err := f.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTECOLLECTION.html */
/* Moved from NotesNoteCollection. */
func NotesNoteCollectionParent(n notesnotecollection.NotesNoteCollection) (NotesDatabase, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDATABASE_PROPERTY_OUTLINE_COM.html */
/* Moved from NotesOutline. */
func NotesOutlineParentDatabase(o notesoutline.NotesOutline) (NotesDatabase, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("ParentDatabase")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATABASE_PROPERTY_OUTLINEENTRY.html */
/* Moved from NotesOutlineEntry. */
func NotesOutlineEntryDatabase(o notesoutlineentry.NotesOutlineEntry) (NotesDatabase, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("Database")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDELEMENT_METHOD_OUTLINEENTRY.html */
/* Moved from NotesOutlineEntry. */
func NotesOutlineEntrySetNamedElement(o notesoutlineentry.NotesOutlineEntry, notesDatabase NotesDatabase, elementName domino.String, entryclass domino.Long) (domino.Boolean, error) {
	val, err := o.Com().CallMethod("SetNamedElement", notesDatabase.Com().Dispatch(), elementName, entryclass)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNOTELINK_METHOD_OUTLINEENTRY.html */
/* Moved from NotesOutlineEntry. */
func NotesOutlineEntrySetNoteLink(o notesoutlineentry.NotesOutlineEntry, notesDatabase NotesDatabase, notesView notesview.NotesView, notesDocument notesdocument.NotesDocument, obj any) (domino.Boolean, error) {
	val, err := o.Com().CallMethod("SetNoteLink", notesDatabase.Com().Dispatch(), notesView.Com().Dispatch(), notesDocument.Com().Dispatch(), obj)
	return helpers.CastValue[domino.Boolean](val), err
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

func NotesRichTextItemAppendDocLink(r notesrichtextitem.NotesRichTextItem, linkTo NotesDatabase, comment domino.String, params ...appendDocLinkParam) error {
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

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_VIEW.html */
/* Moved from NotesView. */
func NotesViewParent(v notesview.NotesView) (NotesDatabase, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACL_PROPERTY.html */
func (d NotesDatabase) ACL() (notesacl.NotesACL, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("ACL")
	return notesacl.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLACTIVITYLOG_PROPERTY_DB.html */
func (d NotesDatabase) ACLActivityLog() ([]domino.String, error) {
	vals, err := d.Com().GetArrayProperty("ACLActivityLog")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AGENTS_PROPERTY.html */
func (d NotesDatabase) Agents() ([]notesagent.NotesAgent, error) {
	return com.GetObjectArrayProperty(d.Com(), notesagent.New, "Agents")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALLDOCUMENTS_PROPERTY.html */
func (d NotesDatabase) AllDocuments() (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("AllDocuments")
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func (d NotesDatabase) Categories() (domino.String, error) {
	val, err := d.Com().GetProperty("Categories")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func (d NotesDatabase) SetCategories(v domino.String) error {
	return d.Com().PutProperty("Categories", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DB.html */
func (d NotesDatabase) Created() (domino.Time, error) {
	val, err := d.Com().GetProperty("Created")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTACCESSLEVEL_PROPERTY.html */
func (d NotesDatabase) CurrentAccessLevel() (domino.Long, error) {
	val, err := d.Com().GetProperty("CurrentAccessLevel")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func (d NotesDatabase) DelayUpdates() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("DelayUpdates")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func (d NotesDatabase) SetDelayUpdates(v domino.Boolean) error {
	return d.Com().PutProperty("DelayUpdates", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNTEMPLATENAME_PROPERTY.html */
func (d NotesDatabase) DesignTemplateName() (domino.String, error) {
	val, err := d.Com().GetProperty("DesignTemplateName")
	return helpers.CastValue[domino.String](val), err
}

/* TODO: Access type for EncryptionStrength could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func (d NotesDatabase) EncryptionStrength() (domino.Integer, error) {
	val, err := d.Com().GetProperty("EncryptionStrength")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func (d NotesDatabase) SetEncryptionStrength(v domino.Integer) error {
	return d.Com().PutProperty("EncryptionStrength", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEFORMAT_PROPERTY_DB.html */
func (d NotesDatabase) FileFormat() (domino.Long, error) {
	val, err := d.Com().GetProperty("FileFormat")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILENAME_PROPERTY.html */
func (d NotesDatabase) FileName() (domino.String, error) {
	val, err := d.Com().GetProperty("FileName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEPATH_PROPERTY.html */
func (d NotesDatabase) FilePath() (domino.String, error) {
	val, err := d.Com().GetProperty("FilePath")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func (d NotesDatabase) FolderReferencesEnabled() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("FolderReferencesEnabled")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func (d NotesDatabase) SetFolderReferencesEnabled(v domino.Boolean) error {
	return d.Com().PutProperty("FolderReferencesEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMS_PROPERTY.html */
func (d NotesDatabase) Forms() ([]notesform.NotesForm, error) {
	return com.GetObjectArrayProperty(d.Com(), notesform.New, "Forms")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func (d NotesDatabase) FTIndexFrequency() (domino.Long, error) {
	val, err := d.Com().GetProperty("FTIndexFrequency")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func (d NotesDatabase) SetFTIndexFrequency(v domino.Long) error {
	return d.Com().PutProperty("FTIndexFrequency", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DATABASE_COM.html */
func (d NotesDatabase) HttpURL() (domino.String, error) {
	val, err := d.Com().GetProperty("HttpURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func (d NotesDatabase) IsClusterReplication() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsClusterReplication")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func (d NotesDatabase) SetIsClusterReplication(v domino.Boolean) error {
	return d.Com().PutProperty("IsClusterReplication", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFIGURATIONDIRECTORY_PROPERTY_DB.html */
func (d NotesDatabase) IsConfigurationDirectory() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsConfigurationDirectory")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICREADER_PROPERTY_DB.html */
func (d NotesDatabase) IsCurrentAccessPublicReader() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsCurrentAccessPublicReader")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICWRITER_PROPERTY_DB.html */
func (d NotesDatabase) IsCurrentAccessPublicWriter() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsCurrentAccessPublicWriter")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) IsDesignLockingEnabled() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsDesignLockingEnabled")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsDesignLockingEnabled(v domino.Boolean) error {
	return d.Com().PutProperty("IsDesignLockingEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDIRECTORYCATALOG_DB.html */
func (d NotesDatabase) IsDirectoryCatalog() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsDirectoryCatalog")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) IsDocumentLockingEnabled() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsDocumentLockingEnabled")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsDocumentLockingEnabled(v domino.Boolean) error {
	return d.Com().PutProperty("IsDocumentLockingEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFTINDEXED_PROPERTY.html */
func (d NotesDatabase) IsFTIndexed() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsFTIndexed")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func (d NotesDatabase) IsInMultiDbIndexing() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsInMultiDbIndexing")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func (d NotesDatabase) SetIsInMultiDbIndexing(v domino.Boolean) error {
	return d.Com().PutProperty("IsInMultiDbIndexing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func (d NotesDatabase) IsInService() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsInService")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func (d NotesDatabase) SetIsInService(v domino.Boolean) error {
	return d.Com().PutProperty("IsInService", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLINK_PROPERTY_DB.html */
func (d NotesDatabase) IsLink() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsLink")
	return helpers.CastValue[domino.Boolean](val), err
}

/* TODO: Access type for IsLocallyEncrypted could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func (d NotesDatabase) IsLocallyEncrypted() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsLocallyEncrypted")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsLocallyEncrypted(v domino.Boolean) error {
	return d.Com().PutProperty("IsLocallyEncrypted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMULTIDBSEARCH_PROPERTY.html */
func (d NotesDatabase) IsMultiDbSearch() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsMultiDbSearch")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISOPEN_PROPERTY.html */
func (d NotesDatabase) IsOpen() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsOpen")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPENDINGDELETE_PROPERTY_DB.html */
func (d NotesDatabase) IsPendingDelete() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsPendingDelete")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATEADDRESSBOOK_PROPERTY.html */
func (d NotesDatabase) IsPrivateAddressBook() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsPrivateAddressBook")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICADDRESSBOOK_PROPERTY.html */
func (d NotesDatabase) IsPublicAddressBook() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsPublicAddressBook")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFIXUP_PROPERTY_DB.html */
func (d NotesDatabase) LastFixup() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("LastFixup")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFTINDEXED_PROPERTY.html */
func (d NotesDatabase) LastFTIndexed() (domino.Time, error) {
	val, err := d.Com().GetProperty("LastFTIndexed")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DB.html */
func (d NotesDatabase) LastModified() (domino.Time, error) {
	val, err := d.Com().GetProperty("LastModified")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func (d NotesDatabase) LimitRevisions() (domino.Double, error) {
	val, err := d.Com().GetProperty("LimitRevisions")
	return helpers.CastValue[domino.Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func (d NotesDatabase) SetLimitRevisions(v domino.Double) error {
	return d.Com().PutProperty("LimitRevisions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func (d NotesDatabase) LimitUpdatedBy() (domino.Double, error) {
	val, err := d.Com().GetProperty("LimitUpdatedBy")
	return helpers.CastValue[domino.Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func (d NotesDatabase) SetLimitUpdatedBy(v domino.Double) error {
	return d.Com().PutProperty("LimitUpdatedBy", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func (d NotesDatabase) ListInDbCatalog() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("ListInDbCatalog")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func (d NotesDatabase) SetListInDbCatalog(v domino.Boolean) error {
	return d.Com().PutProperty("ListInDbCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MANAGERS_PROPERTY.html */
func (d NotesDatabase) Managers() ([]domino.String, error) {
	vals, err := d.Com().GetArrayProperty("Managers")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSIZE_PROPERTY_6579_ABOUT.html */
func (d NotesDatabase) MaxSize() (domino.Double, error) {
	val, err := d.Com().GetProperty("MaxSize")
	return helpers.CastValue[domino.Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DATABASE_COM.html */
func (d NotesDatabase) NotesURL() (domino.String, error) {
	val, err := d.Com().GetProperty("NotesURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PERCENTUSED_PROPERTY.html */
func (d NotesDatabase) PercentUsed() (domino.Double, error) {
	val, err := d.Com().GetProperty("PercentUsed")
	return helpers.CastValue[domino.Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAID_PROPERTY.html */
func (d NotesDatabase) ReplicaID() (domino.String, error) {
	val, err := d.Com().GetProperty("ReplicaID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATIONINFO_PROPERTY_7679_ABOUT.html */
func (d NotesDatabase) ReplicationInfo() (notesreplication.NotesReplication, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("ReplicationInfo")
	return notesreplication.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY.html */
func (d NotesDatabase) Server() (domino.String, error) {
	val, err := d.Com().GetProperty("Server")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DB.html */
func (d NotesDatabase) Size() (domino.Double, error) {
	val, err := d.Com().GetProperty("Size")
	return helpers.CastValue[domino.Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func (d NotesDatabase) SizeQuota() (domino.Long, error) {
	val, err := d.Com().GetProperty("SizeQuota")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func (d NotesDatabase) SetSizeQuota(v domino.Long) error {
	return d.Com().PutProperty("SizeQuota", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func (d NotesDatabase) SizeWarning() (domino.Long, error) {
	val, err := d.Com().GetProperty("SizeWarning")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func (d NotesDatabase) SetSizeWarning(v domino.Long) error {
	return d.Com().PutProperty("SizeWarning", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEMPLATENAME_PROPERTY.html */
func (d NotesDatabase) TemplateName() (domino.String, error) {
	val, err := d.Com().GetProperty("TemplateName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func (d NotesDatabase) Title() (domino.String, error) {
	val, err := d.Com().GetProperty("Title")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func (d NotesDatabase) SetTitle(v domino.String) error {
	return d.Com().PutProperty("Title", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_DB.html */
func (d NotesDatabase) Type() (domino.Long, error) {
	val, err := d.Com().GetProperty("Type")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func (d NotesDatabase) UndeleteExpireTime() (domino.Long, error) {
	val, err := d.Com().GetProperty("UndeleteExpireTime")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func (d NotesDatabase) SetUndeleteExpireTime(v domino.Long) error {
	return d.Com().PutProperty("UndeleteExpireTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDDOCUMENTS_PROPERTY.html */
func (d NotesDatabase) UnprocessedDocuments() (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("UnprocessedDocuments")
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY.html */
func (d NotesDatabase) Views() ([]notesview.NotesView, error) {
	return com.GetObjectArrayProperty(d.Com(), notesview.New, "Views")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACT_METHOD.html */
func (d NotesDatabase) Compact() (domino.Long, error) {
	val, err := d.Com().CallMethod("Compact")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACTWITHOPTIONS_METHOD_DB.html */
type compactWithOptionsParams struct {
	spaceThreshold *domino.String
}

type compactWithOptionsParam func(*compactWithOptionsParams)

func WithCompactWithOptionsSpaceThreshold(spaceThreshold domino.String) compactWithOptionsParam {
	return func(c *compactWithOptionsParams) {
		c.spaceThreshold = &spaceThreshold
	}
}

func (d NotesDatabase) CompactWithOptions(options any, params ...compactWithOptionsParam) (domino.Long, error) {
	s := &compactWithOptionsParams{}
	paramsOrdered := []interface{}{options, options}

	for _, p := range params {
		p(s)
	}

	if s.spaceThreshold != nil {
		paramsOrdered = append(paramsOrdered, *s.spaceThreshold)
	}
	val, err := d.Com().CallMethod("CompactWithOptions", paramsOrdered...)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATE_METHOD.html */
type createParams struct {
	maxsize *domino.Integer
}

type createParam func(*createParams)

func WithCreateMaxsize(maxsize domino.Integer) createParam {
	return func(c *createParams) {
		c.maxsize = &maxsize
	}
}

func (d NotesDatabase) Create(server domino.String, dbfile domino.String, openFlag domino.Boolean, params ...createParam) error {
	s := &createParams{}
	paramsOrdered := []interface{}{server, dbfile, openFlag}

	for _, p := range params {
		p(s)
	}

	if s.maxsize != nil {
		paramsOrdered = append(paramsOrdered, *s.maxsize)
	}
	_, err := d.Com().CallMethod("Create", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOPY_METHOD.html */
type createCopyParams struct {
	maxsize *domino.Integer
}

type createCopyParam func(*createCopyParams)

func WithCreateCopyMaxsize(maxsize domino.Integer) createCopyParam {
	return func(c *createCopyParams) {
		c.maxsize = &maxsize
	}
}

func (d NotesDatabase) CreateCopy(newServer domino.String, newDbFile domino.String, params ...createCopyParam) (NotesDatabase, error) {
	s := &createCopyParams{}
	paramsOrdered := []interface{}{newServer, newDbFile}

	for _, p := range params {
		p(s)
	}

	if s.maxsize != nil {
		paramsOrdered = append(paramsOrdered, *s.maxsize)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateCopy", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENT_METHOD.html */
func (d NotesDatabase) CreateDocument() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDocument")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENTCOLLECTION_METHOD.html */
func (d NotesDatabase) CreateDocumentCollection() (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDocumentCollection")
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOMINOQUERY_METHOD.html */
func (d NotesDatabase) CreateDominoQuery() (notesdominoquery.NotesDominoQuery, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDominoQuery")
	return notesdominoquery.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFROMTEMPLATE_METHOD.html */
type createFromTemplateParams struct {
	maxsize *domino.Integer
}

type createFromTemplateParam func(*createFromTemplateParams)

func WithCreateFromTemplateMaxsize(maxsize domino.Integer) createFromTemplateParam {
	return func(c *createFromTemplateParams) {
		c.maxsize = &maxsize
	}
}

func (d NotesDatabase) CreateFromTemplate(newServer domino.String, newDbFile domino.String, inheritFlag domino.Boolean, params ...createFromTemplateParam) (NotesDatabase, error) {
	s := &createFromTemplateParams{}
	paramsOrdered := []interface{}{newServer, newDbFile, inheritFlag}

	for _, p := range params {
		p(s)
	}

	if s.maxsize != nil {
		paramsOrdered = append(paramsOrdered, *s.maxsize)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateFromTemplate", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_METHOD_DB.html */
func (d NotesDatabase) CreateFTIndex(options domino.Long, recreate domino.Boolean) error {
	_, err := d.Com().CallMethod("CreateFTIndex", options, recreate)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENOTECOLLECTION_METHOD_DATABASE.html */
func (d NotesDatabase) CreateNoteCollection(selectAllFlag domino.Boolean) (notesnotecollection.NotesNoteCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateNoteCollection", selectAllFlag)
	return notesnotecollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEOUTLINE_METHOD_DATABASE.html */
type createOutlineParams struct {
	defaultOutline *domino.Boolean
}

type createOutlineParam func(*createOutlineParams)

func WithCreateOutlineDefaultOutline(defaultOutline domino.Boolean) createOutlineParam {
	return func(c *createOutlineParams) {
		c.defaultOutline = &defaultOutline
	}
}

func (d NotesDatabase) CreateOutline(outlinename domino.String, params ...createOutlineParam) (notesoutline.NotesOutline, error) {
	s := &createOutlineParams{}
	paramsOrdered := []interface{}{outlinename}

	for _, p := range params {
		p(s)
	}

	if s.defaultOutline != nil {
		paramsOrdered = append(paramsOrdered, *s.defaultOutline)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateOutline", paramsOrdered...)
	return notesoutline.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD.html */
func (d NotesDatabase) CreateReplica(newServer domino.String, newDbFile domino.String) (NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateReplica", newServer, newDbFile)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEW_METHOD_DB.html */
type createViewParams struct {
	viewName                           *domino.String
	viewSelectionFormula               *domino.String
	templateView                       *notesview.NotesView
	prohibitDesignRefreshModifications *domino.Boolean
}

type createViewParam func(*createViewParams)

func WithCreateViewViewName(viewName domino.String) createViewParam {
	return func(c *createViewParams) {
		c.viewName = &viewName
	}
}

func WithCreateViewViewSelectionFormula(viewSelectionFormula domino.String) createViewParam {
	return func(c *createViewParams) {
		c.viewSelectionFormula = &viewSelectionFormula
	}
}

func WithCreateViewTemplateView(templateView notesview.NotesView) createViewParam {
	return func(c *createViewParams) {
		c.templateView = &templateView
	}
}

func WithCreateViewProhibitDesignRefreshModifications(prohibitDesignRefreshModifications domino.Boolean) createViewParam {
	return func(c *createViewParams) {
		c.prohibitDesignRefreshModifications = &prohibitDesignRefreshModifications
	}
}

func (d NotesDatabase) CreateView(params ...createViewParam) (notesview.NotesView, error) {
	s := &createViewParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.viewName != nil {
		paramsOrdered = append(paramsOrdered, *s.viewName)
		if s.viewSelectionFormula != nil {
			paramsOrdered = append(paramsOrdered, *s.viewSelectionFormula)
			if s.templateView != nil {
				paramsOrdered = append(paramsOrdered, *s.templateView)
				if s.prohibitDesignRefreshModifications != nil {
					paramsOrdered = append(paramsOrdered, *s.prohibitDesignRefreshModifications)
				}
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateView", paramsOrdered...)
	return notesview.New(dispatchPtr), err
}

func (d NotesDatabase) Decrypt(def domino.Boolean) error {
	_, err := d.Com().CallMethod("Decrypt", def)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEFOLDER_METHOD_DATABASE.html */
func (d NotesDatabase) EnableFolder(foldername domino.String) error {
	_, err := d.Com().CallMethod("EnableFolder", foldername)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD_DATABASE.html */
func (d NotesDatabase) Encrypt(encryptionStrength domino.Integer, def domino.Boolean) error {
	_, err := d.Com().CallMethod("Encrypt", encryptionStrength, def)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIXUP_METHOD_DB.html */
type fixupParams struct {
	options *domino.Long
}

type fixupParam func(*fixupParams)

func WithFixupOptions(options domino.Long) fixupParam {
	return func(c *fixupParams) {
		c.options = &options
	}
}

func (d NotesDatabase) Fixup(params ...fixupParam) error {
	s := &fixupParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.options != nil {
		paramsOrdered = append(paramsOrdered, *s.options)
	}
	_, err := d.Com().CallMethod("Fixup", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTDOMAINSEARCH_METHOD_9515.html */
type fTDomainSearchParams struct {
	sortoptions  *domino.Integer
	otheroptions *domino.Integer
	start        *domino.Long
	count        *domino.Integer
}

type fTDomainSearchParam func(*fTDomainSearchParams)

func WithFTDomainSearchSortoptions(sortoptions domino.Integer) fTDomainSearchParam {
	return func(c *fTDomainSearchParams) {
		c.sortoptions = &sortoptions
	}
}

func WithFTDomainSearchOtheroptions(otheroptions domino.Integer) fTDomainSearchParam {
	return func(c *fTDomainSearchParams) {
		c.otheroptions = &otheroptions
	}
}

func WithFTDomainSearchStart(start domino.Long) fTDomainSearchParam {
	return func(c *fTDomainSearchParams) {
		c.start = &start
	}
}

func WithFTDomainSearchCount(count domino.Integer) fTDomainSearchParam {
	return func(c *fTDomainSearchParams) {
		c.count = &count
	}
}

func (d NotesDatabase) FTDomainSearch(query domino.String, maxDocs domino.Integer, entryform domino.String, params ...fTDomainSearchParam) (notesdocument.NotesDocument, error) {
	s := &fTDomainSearchParams{}
	paramsOrdered := []interface{}{query, maxDocs, entryform}

	for _, p := range params {
		p(s)
	}

	if s.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *s.sortoptions)
		if s.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *s.otheroptions)
			if s.start != nil {
				paramsOrdered = append(paramsOrdered, *s.start)
				if s.count != nil {
					paramsOrdered = append(paramsOrdered, *s.count)
				}
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("FTDomainSearch", paramsOrdered...)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_DB.html */
type fTSearchParams struct {
	sortoptions  *domino.Integer
	otheroptions *domino.Integer
}

type fTSearchParam func(*fTSearchParams)

func WithFTSearchSortoptions(sortoptions domino.Integer) fTSearchParam {
	return func(c *fTSearchParams) {
		c.sortoptions = &sortoptions
	}
}

func WithFTSearchOtheroptions(otheroptions domino.Integer) fTSearchParam {
	return func(c *fTSearchParams) {
		c.otheroptions = &otheroptions
	}
}

func (d NotesDatabase) FTSearch(query domino.String, maxdocs domino.Integer, params ...fTSearchParam) (notesdocumentcollection.NotesDocumentCollection, error) {
	s := &fTSearchParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(s)
	}

	if s.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *s.sortoptions)
		if s.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *s.otheroptions)
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("FTSearch", paramsOrdered...)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHRANGE_METHOD_DB.html */
type fTSearchRangeParams struct {
	sortoptions  *domino.Integer
	otheroptions *domino.Integer
	start        *domino.Long
}

type fTSearchRangeParam func(*fTSearchRangeParams)

func WithFTSearchRangeSortoptions(sortoptions domino.Integer) fTSearchRangeParam {
	return func(c *fTSearchRangeParams) {
		c.sortoptions = &sortoptions
	}
}

func WithFTSearchRangeOtheroptions(otheroptions domino.Integer) fTSearchRangeParam {
	return func(c *fTSearchRangeParams) {
		c.otheroptions = &otheroptions
	}
}

func WithFTSearchRangeStart(start domino.Long) fTSearchRangeParam {
	return func(c *fTSearchRangeParams) {
		c.start = &start
	}
}

func (d NotesDatabase) FTSearchRange(query domino.String, maxdocs domino.Integer, params ...fTSearchRangeParam) (notesdocumentcollection.NotesDocumentCollection, error) {
	s := &fTSearchRangeParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(s)
	}

	if s.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *s.sortoptions)
		if s.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *s.otheroptions)
			if s.start != nil {
				paramsOrdered = append(paramsOrdered, *s.start)
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("FTSearchRange", paramsOrdered...)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETAGENT_METHOD.html */
func (d NotesDatabase) GetAgent(agentName domino.String) (notesagent.NotesAgent, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetAgent", agentName)
	return notesagent.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLREADDOCUMENTS_DATABASE.html */
type getAllReadDocumentsParams struct {
	username *domino.String
}

type getAllReadDocumentsParam func(*getAllReadDocumentsParams)

func WithGetAllReadDocumentsUsername(username domino.String) getAllReadDocumentsParam {
	return func(c *getAllReadDocumentsParams) {
		c.username = &username
	}
}

func (d NotesDatabase) GetAllReadDocuments(params ...getAllReadDocumentsParam) (notesnotecollection.NotesNoteCollection, error) {
	s := &getAllReadDocumentsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.username != nil {
		paramsOrdered = append(paramsOrdered, *s.username)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetAllReadDocuments", paramsOrdered...)
	return notesnotecollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLUNREADDOCUMENTS_DATABASE.html */
type getAllUnreadDocumentsParams struct {
	username *domino.String
}

type getAllUnreadDocumentsParam func(*getAllUnreadDocumentsParams)

func WithGetAllUnreadDocumentsUsername(username domino.String) getAllUnreadDocumentsParam {
	return func(c *getAllUnreadDocumentsParams) {
		c.username = &username
	}
}

func (d NotesDatabase) GetAllUnreadDocuments(params ...getAllUnreadDocumentsParam) (notesnotecollection.NotesNoteCollection, error) {
	s := &getAllUnreadDocumentsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.username != nil {
		paramsOrdered = append(paramsOrdered, *s.username)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetAllUnreadDocuments", paramsOrdered...)
	return notesnotecollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYID_METHOD.html */
func (d NotesDatabase) GetDocumentByID(noteID domino.String) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocumentByID", noteID)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYUNID_METHOD.html */
func (d NotesDatabase) GetDocumentByUNID(unid domino.String) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocumentByUNID", unid)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYURL_METHOD.html */
type getDocumentByURLParams struct {
	reload            *domino.Integer
	urllist           *domino.Integer
	charset           *domino.String
	webusername       *domino.String
	webpassword       *domino.String
	proxywebusername  *domino.String
	proxywebpassword  *domino.String
	returnimmediately *domino.Boolean
}

type getDocumentByURLParam func(*getDocumentByURLParams)

func WithGetDocumentByURLReload(reload domino.Integer) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.reload = &reload
	}
}

func WithGetDocumentByURLUrllist(urllist domino.Integer) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.urllist = &urllist
	}
}

func WithGetDocumentByURLCharset(charset domino.String) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.charset = &charset
	}
}

func WithGetDocumentByURLWebusername(webusername domino.String) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.webusername = &webusername
	}
}

func WithGetDocumentByURLWebpassword(webpassword domino.String) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.webpassword = &webpassword
	}
}

func WithGetDocumentByURLProxywebusername(proxywebusername domino.String) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.proxywebusername = &proxywebusername
	}
}

func WithGetDocumentByURLProxywebpassword(proxywebpassword domino.String) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.proxywebpassword = &proxywebpassword
	}
}

func WithGetDocumentByURLReturnimmediately(returnimmediately domino.Boolean) getDocumentByURLParam {
	return func(c *getDocumentByURLParams) {
		c.returnimmediately = &returnimmediately
	}
}

func (d NotesDatabase) GetDocumentByURL(URL domino.String, params ...getDocumentByURLParam) (notesdocument.NotesDocument, error) {
	s := &getDocumentByURLParams{}
	paramsOrdered := []interface{}{URL}

	for _, p := range params {
		p(s)
	}

	if s.reload != nil {
		paramsOrdered = append(paramsOrdered, *s.reload)
		if s.urllist != nil {
			paramsOrdered = append(paramsOrdered, *s.urllist)
			if s.charset != nil {
				paramsOrdered = append(paramsOrdered, *s.charset)
				if s.webusername != nil {
					paramsOrdered = append(paramsOrdered, *s.webusername)
					if s.webpassword != nil {
						paramsOrdered = append(paramsOrdered, *s.webpassword)
						if s.proxywebusername != nil {
							paramsOrdered = append(paramsOrdered, *s.proxywebusername)
							if s.proxywebpassword != nil {
								paramsOrdered = append(paramsOrdered, *s.proxywebpassword)
								if s.returnimmediately != nil {
									paramsOrdered = append(paramsOrdered, *s.returnimmediately)
								}
							}
						}
					}
				}
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocumentByURL", paramsOrdered...)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFORM_METHOD.html */
func (d NotesDatabase) GetForm(name domino.String) (notesform.NotesForm, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetForm", name)
	return notesform.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTS_METHOD_DB.html */
func (d NotesDatabase) GetModifiedDocuments(since notesdatetime.NotesDateTime, noteClass domino.Integer) (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetModifiedDocuments", since.Com().Dispatch(), noteClass)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTWITHOPTIONS_METHOD_DATABASE.html */
func (d NotesDatabase) GetModifiedDocumentsWithOptions(modifiedSince notesdatetime.NotesDateTime, modifiedUntil notesdatetime.NotesDateTime, options domino.Integer) (notesnotecollection.NotesNoteCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetModifiedDocumentsWithOptions", modifiedSince.Com().Dispatch(), modifiedUntil.Com().Dispatch(), options)
	return notesnotecollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENT_METHOD.html */
func (d NotesDatabase) GetNamedDocument() error {
	_, err := d.Com().CallMethod("GetNamedDocument")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENTCOLLECTION_METHOD.html */
type getNamedDocCollectionParams struct {
	name *domino.String
}

type getNamedDocCollectionParam func(*getNamedDocCollectionParams)

func WithGetNamedDocCollectionName(name domino.String) getNamedDocCollectionParam {
	return func(c *getNamedDocCollectionParams) {
		c.name = &name
	}
}

func (d NotesDatabase) GetNamedDocCollection(params ...getNamedDocCollectionParam) error {
	s := &getNamedDocCollectionParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.name != nil {
		paramsOrdered = append(paramsOrdered, *s.name)
	}
	_, err := d.Com().CallMethod("GetNamedDocCollection", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOPTION_METHOD_DB.html */
func (d NotesDatabase) GetOption(optionName domino.Integer) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("GetOption", optionName)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOUTLINE_METHOD_DATABASE.html */
func (d NotesDatabase) GetOutline(outlinename domino.String) (notesoutline.NotesOutline, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetOutline", outlinename)
	return notesoutline.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROFILEDOCCOLLECTION_METHOD_DATABASE.html */
type getProfileDocCollectionParams struct {
	profilename *domino.String
}

type getProfileDocCollectionParam func(*getProfileDocCollectionParams)

func WithGetProfileDocCollectionProfilename(profilename domino.String) getProfileDocCollectionParam {
	return func(c *getProfileDocCollectionParams) {
		c.profilename = &profilename
	}
}

func (d NotesDatabase) GetProfileDocCollection(params ...getProfileDocCollectionParam) (notesdocumentcollection.NotesDocumentCollection, error) {
	s := &getProfileDocCollectionParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.profilename != nil {
		paramsOrdered = append(paramsOrdered, *s.profilename)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetProfileDocCollection", paramsOrdered...)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROFILEDOCUMENT_METHOD.html */
type getProfileDocumentParams struct {
	uniqueKey *domino.String
}

type getProfileDocumentParam func(*getProfileDocumentParams)

func WithGetProfileDocumentUniqueKey(uniqueKey domino.String) getProfileDocumentParam {
	return func(c *getProfileDocumentParams) {
		c.uniqueKey = &uniqueKey
	}
}

func (d NotesDatabase) GetProfileDocument(profilename domino.String, params ...getProfileDocumentParam) (notesdocument.NotesDocument, error) {
	s := &getProfileDocumentParams{}
	paramsOrdered := []interface{}{profilename}

	for _, p := range params {
		p(s)
	}

	if s.uniqueKey != nil {
		paramsOrdered = append(paramsOrdered, *s.uniqueKey)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetProfileDocument", paramsOrdered...)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETQUERYRESULTSPROCESSOR_METHOD.html */
func (d NotesDatabase) GetQueryResultsProcessor() (notesqueryresultsprocessor.NotesQueryResultsProcessor, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetQueryResultsProcessor")
	return notesqueryresultsprocessor.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETURLHEADERINFO_METHOD.html */
type getURLHeaderInfoParams struct {
	Webusername      *domino.String
	Webpassword      *domino.String
	Proxywebusername *domino.String
	Proxywebpassword *domino.String
}

type getURLHeaderInfoParam func(*getURLHeaderInfoParams)

func WithGetURLHeaderInfoWebusername(Webusername domino.String) getURLHeaderInfoParam {
	return func(c *getURLHeaderInfoParams) {
		c.Webusername = &Webusername
	}
}

func WithGetURLHeaderInfoWebpassword(Webpassword domino.String) getURLHeaderInfoParam {
	return func(c *getURLHeaderInfoParams) {
		c.Webpassword = &Webpassword
	}
}

func WithGetURLHeaderInfoProxywebusername(Proxywebusername domino.String) getURLHeaderInfoParam {
	return func(c *getURLHeaderInfoParams) {
		c.Proxywebusername = &Proxywebusername
	}
}

func WithGetURLHeaderInfoProxywebpassword(Proxywebpassword domino.String) getURLHeaderInfoParam {
	return func(c *getURLHeaderInfoParams) {
		c.Proxywebpassword = &Proxywebpassword
	}
}

func (d NotesDatabase) GetURLHeaderInfo(URL domino.String, headername domino.String, params ...getURLHeaderInfoParam) (domino.String, error) {
	s := &getURLHeaderInfoParams{}
	paramsOrdered := []interface{}{URL, headername}

	for _, p := range params {
		p(s)
	}

	if s.Webusername != nil {
		paramsOrdered = append(paramsOrdered, *s.Webusername)
		if s.Webpassword != nil {
			paramsOrdered = append(paramsOrdered, *s.Webpassword)
			if s.Proxywebusername != nil {
				paramsOrdered = append(paramsOrdered, *s.Proxywebusername)
				if s.Proxywebpassword != nil {
					paramsOrdered = append(paramsOrdered, *s.Proxywebpassword)
				}
			}
		}
	}
	val, err := d.Com().CallMethod("GetURLHeaderInfo", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVIEW_METHOD.html */
func (d NotesDatabase) GetView(viewName domino.String) (notesview.NotesView, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetView", viewName)
	return notesview.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GRANTACCESS_METHOD.html */
func (d NotesDatabase) GrantAccess(name domino.String, level domino.Integer) error {
	_, err := d.Com().CallMethod("GrantAccess", name, level)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKFORDELETE_METHOD_DB.html */
func (d NotesDatabase) MarkForDelete() error {
	_, err := d.Com().CallMethod("MarkForDelete")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD.html */
func (d NotesDatabase) Open(server domino.String, dbfile domino.String) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("Open", server, dbfile)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENBYREPLICAID_METHOD.html */
func (d NotesDatabase) OpenByReplicaID(server domino.String, replicaID domino.String) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("OpenByReplicaID", server, replicaID)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENIFMODIFIED_METHOD.html */
func (d NotesDatabase) OpenIfModified(server domino.String, dbfile domino.String, notesDateTime notesdatetime.NotesDateTime) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("OpenIfModified", server, dbfile, notesDateTime.Com().Dispatch())
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAIL_METHOD.html */
func (d NotesDatabase) OpenMail() error {
	_, err := d.Com().CallMethod("OpenMail")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENURLDB_METHOD.html */
func (d NotesDatabase) OpenURLDb() (domino.Boolean, error) {
	val, err := d.Com().CallMethod("OpenURLDb")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENWITHFAILOVER_METHOD.html */
func (d NotesDatabase) OpenWithFailover(server domino.String, dbfile domino.String) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("OpenWithFailover", server, dbfile)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESS_METHOD.html */
func (d NotesDatabase) QueryAccess(name domino.String) (domino.Integer, error) {
	val, err := d.Com().CallMethod("QueryAccess", name)
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSPRIVILEGES_METHOD_DB.html */
func (d NotesDatabase) QueryAccessPrivileges(name domino.String) (domino.Long, error) {
	val, err := d.Com().CallMethod("QueryAccessPrivileges", name)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSROLES_METHOD_DB.html */
func (d NotesDatabase) QueryAccessRoles(name domino.String) ([]domino.String, error) {
	vals, err := d.Com().CallArrayMethod("QueryAccessRoles", name)
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DB.html */
func (d NotesDatabase) Remove() error {
	_, err := d.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFTINDEX_METHOD_DB.html */
func (d NotesDatabase) RemoveFTIndex() error {
	_, err := d.Com().CallMethod("RemoveFTIndex")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATE_METHOD.html */
func (d NotesDatabase) Replicate(serverName domino.String) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("Replicate", serverName)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REVOKEACCESS_METHOD.html */
func (d NotesDatabase) RevokeAccess(name domino.String) error {
	_, err := d.Com().CallMethod("RevokeAccess", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCH_METHOD.html */
func (d NotesDatabase) Search(formula domino.String, notesDateTime notesdatetime.NotesDateTime, maxDocs domino.Integer) (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("Search", formula, notesDateTime.Com().Dispatch(), maxDocs)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETOPTION_METHOD_DB.html */
func (d NotesDatabase) SetOption(optionName domino.Integer, flag any) error {
	_, err := d.Com().CallMethod("SetOption", optionName, flag)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERIDFORDECRYPT_DATABASE.html */
func (d NotesDatabase) SetUserIDForDecrypt(uid notesuserid.NotesUserID, idFile string, password string) error {
	_, err := d.Com().CallMethod("SetUserIDForDecrypt", uid.Com().Dispatch(), idFile, password)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD_DB.html */
type signParams struct {
	documentType     *domino.Integer
	existingSigsOnly *domino.Boolean
	nameStr          *domino.String
	nameStrIsNoteID  *domino.Boolean
}

type signParam func(*signParams)

func WithSignDocumentType(documentType domino.Integer) signParam {
	return func(c *signParams) {
		c.documentType = &documentType
	}
}

func WithSignExistingSigsOnly(existingSigsOnly domino.Boolean) signParam {
	return func(c *signParams) {
		c.existingSigsOnly = &existingSigsOnly
	}
}

func WithSignNameStr(nameStr domino.String) signParam {
	return func(c *signParams) {
		c.nameStr = &nameStr
	}
}

func WithSignNameStrIsNoteID(nameStrIsNoteID domino.Boolean) signParam {
	return func(c *signParams) {
		c.nameStrIsNoteID = &nameStrIsNoteID
	}
}

func (d NotesDatabase) Sign(params ...signParam) error {
	s := &signParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.documentType != nil {
		paramsOrdered = append(paramsOrdered, *s.documentType)
		if s.existingSigsOnly != nil {
			paramsOrdered = append(paramsOrdered, *s.existingSigsOnly)
			if s.nameStr != nil {
				paramsOrdered = append(paramsOrdered, *s.nameStr)
				if s.nameStrIsNoteID != nil {
					paramsOrdered = append(paramsOrdered, *s.nameStrIsNoteID)
				}
			}
		}
	}
	_, err := d.Com().CallMethod("Sign", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONBEGIN_METHOD.html */
func (d NotesDatabase) TransactionBegin() error {
	_, err := d.Com().CallMethod("TransactionBegin")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONCOMMIT_METHOD.html */
func (d NotesDatabase) TransactionCommit() error {
	_, err := d.Com().CallMethod("TransactionCommit")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONROLLBACK_METHOD.html */
func (d NotesDatabase) TransactionRollback() error {
	_, err := d.Com().CallMethod("TransactionRollback")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDFTSEARCH_METHOD.html */
type unprocessedFTSearchParams struct {
	sortoptions  *domino.Integer
	otheroptions *domino.Integer
}

type unprocessedFTSearchParam func(*unprocessedFTSearchParams)

func WithUnprocessedFTSearchSortoptions(sortoptions domino.Integer) unprocessedFTSearchParam {
	return func(c *unprocessedFTSearchParams) {
		c.sortoptions = &sortoptions
	}
}

func WithUnprocessedFTSearchOtheroptions(otheroptions domino.Integer) unprocessedFTSearchParam {
	return func(c *unprocessedFTSearchParams) {
		c.otheroptions = &otheroptions
	}
}

func (d NotesDatabase) UnprocessedFTSearch(query domino.String, maxdocs domino.Integer, params ...unprocessedFTSearchParam) (notesdocumentcollection.NotesDocumentCollection, error) {
	s := &unprocessedFTSearchParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(s)
	}

	if s.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *s.sortoptions)
		if s.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *s.otheroptions)
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("UnprocessedFTSearch", paramsOrdered...)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDFTSEARCHRANGE_METHOD_DB.html */
type unprocessedFTSearchRangeParams struct {
	sortoptions  *domino.Integer
	otheroptions *domino.Integer
	start        *domino.Long
}

type unprocessedFTSearchRangeParam func(*unprocessedFTSearchRangeParams)

func WithUnprocessedFTSearchRangeSortoptions(sortoptions domino.Integer) unprocessedFTSearchRangeParam {
	return func(c *unprocessedFTSearchRangeParams) {
		c.sortoptions = &sortoptions
	}
}

func WithUnprocessedFTSearchRangeOtheroptions(otheroptions domino.Integer) unprocessedFTSearchRangeParam {
	return func(c *unprocessedFTSearchRangeParams) {
		c.otheroptions = &otheroptions
	}
}

func WithUnprocessedFTSearchRangeStart(start domino.Long) unprocessedFTSearchRangeParam {
	return func(c *unprocessedFTSearchRangeParams) {
		c.start = &start
	}
}

func (d NotesDatabase) UnprocessedFTSearchRange(query domino.String, maxdocs domino.Integer, params ...unprocessedFTSearchRangeParam) (notesdocumentcollection.NotesDocumentCollection, error) {
	s := &unprocessedFTSearchRangeParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(s)
	}

	if s.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *s.sortoptions)
		if s.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *s.otheroptions)
			if s.start != nil {
				paramsOrdered = append(paramsOrdered, *s.start)
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("UnprocessedFTSearchRange", paramsOrdered...)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDSEARCH_METHOD.html */
func (d NotesDatabase) UnprocessedSearch(formula domino.String, notesDateTime notesdatetime.NotesDateTime, maxDocs domino.Integer) (notesdocumentcollection.NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("UnprocessedSearch", formula, notesDateTime.Com().Dispatch(), maxDocs)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEFTINDEX_METHOD.html */
func (d NotesDatabase) UpdateFTIndex(createFlag domino.Boolean) error {
	_, err := d.Com().CallMethod("UpdateFTIndex", createFlag)
	return err
}
