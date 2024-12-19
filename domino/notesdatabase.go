/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATABASE_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDatabase struct {
	NotesStruct
}

func NewNotesDatabase(dispatchPtr *ole.IDispatch) NotesDatabase {
	return NotesDatabase{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACL_PROPERTY.html */
func (d NotesDatabase) ACL() (NotesACL, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("ACL")
	return NewNotesACL(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLACTIVITYLOG_PROPERTY_DB.html */
func (d NotesDatabase) ACLActivityLog() ([]String, error) {
	vals, err := d.Com().GetArrayProperty("ACLActivityLog")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AGENTS_PROPERTY.html */
func (d NotesDatabase) Agents() ([]NotesAgent, error) {
	return com.GetObjectArrayProperty(d.Com(), NewNotesAgent, "Agents")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALLDOCUMENTS_PROPERTY.html */
func (d NotesDatabase) AllDocuments() (NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("AllDocuments")
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func (d NotesDatabase) Categories() (String, error) {
	val, err := d.Com().GetProperty("Categories")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func (d NotesDatabase) SetCategories(v String) error {
	return d.Com().PutProperty("Categories", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DB.html */
func (d NotesDatabase) Created() (Time, error) {
	val, err := d.Com().GetProperty("Created")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTACCESSLEVEL_PROPERTY.html */
func (d NotesDatabase) CurrentAccessLevel() (Long, error) {
	val, err := d.Com().GetProperty("CurrentAccessLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func (d NotesDatabase) DelayUpdates() (Boolean, error) {
	val, err := d.Com().GetProperty("DelayUpdates")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func (d NotesDatabase) SetDelayUpdates(v Boolean) error {
	return d.Com().PutProperty("DelayUpdates", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNTEMPLATENAME_PROPERTY.html */
func (d NotesDatabase) DesignTemplateName() (String, error) {
	val, err := d.Com().GetProperty("DesignTemplateName")
	return helpers.CastValue[String](val), err
}

/* TODO: Access type for EncryptionStrength could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func (d NotesDatabase) EncryptionStrength() (Integer, error) {
	val, err := d.Com().GetProperty("EncryptionStrength")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func (d NotesDatabase) SetEncryptionStrength(v Integer) error {
	return d.Com().PutProperty("EncryptionStrength", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEFORMAT_PROPERTY_DB.html */
func (d NotesDatabase) FileFormat() (Long, error) {
	val, err := d.Com().GetProperty("FileFormat")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILENAME_PROPERTY.html */
func (d NotesDatabase) FileName() (String, error) {
	val, err := d.Com().GetProperty("FileName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEPATH_PROPERTY.html */
func (d NotesDatabase) FilePath() (String, error) {
	val, err := d.Com().GetProperty("FilePath")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func (d NotesDatabase) FolderReferencesEnabled() (Boolean, error) {
	val, err := d.Com().GetProperty("FolderReferencesEnabled")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func (d NotesDatabase) SetFolderReferencesEnabled(v Boolean) error {
	return d.Com().PutProperty("FolderReferencesEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMS_PROPERTY.html */
func (d NotesDatabase) Forms() ([]NotesForm, error) {
	return com.GetObjectArrayProperty(d.Com(), NewNotesForm, "Forms")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func (d NotesDatabase) FTIndexFrequency() (Long, error) {
	val, err := d.Com().GetProperty("FTIndexFrequency")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func (d NotesDatabase) SetFTIndexFrequency(v Long) error {
	return d.Com().PutProperty("FTIndexFrequency", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DATABASE_COM.html */
func (d NotesDatabase) HttpURL() (String, error) {
	val, err := d.Com().GetProperty("HttpURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func (d NotesDatabase) IsClusterReplication() (Boolean, error) {
	val, err := d.Com().GetProperty("IsClusterReplication")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func (d NotesDatabase) SetIsClusterReplication(v Boolean) error {
	return d.Com().PutProperty("IsClusterReplication", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFIGURATIONDIRECTORY_PROPERTY_DB.html */
func (d NotesDatabase) IsConfigurationDirectory() (Boolean, error) {
	val, err := d.Com().GetProperty("IsConfigurationDirectory")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICREADER_PROPERTY_DB.html */
func (d NotesDatabase) IsCurrentAccessPublicReader() (Boolean, error) {
	val, err := d.Com().GetProperty("IsCurrentAccessPublicReader")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICWRITER_PROPERTY_DB.html */
func (d NotesDatabase) IsCurrentAccessPublicWriter() (Boolean, error) {
	val, err := d.Com().GetProperty("IsCurrentAccessPublicWriter")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) IsDesignLockingEnabled() (Boolean, error) {
	val, err := d.Com().GetProperty("IsDesignLockingEnabled")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsDesignLockingEnabled(v Boolean) error {
	return d.Com().PutProperty("IsDesignLockingEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDIRECTORYCATALOG_DB.html */
func (d NotesDatabase) IsDirectoryCatalog() (Boolean, error) {
	val, err := d.Com().GetProperty("IsDirectoryCatalog")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) IsDocumentLockingEnabled() (Boolean, error) {
	val, err := d.Com().GetProperty("IsDocumentLockingEnabled")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsDocumentLockingEnabled(v Boolean) error {
	return d.Com().PutProperty("IsDocumentLockingEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFTINDEXED_PROPERTY.html */
func (d NotesDatabase) IsFTIndexed() (Boolean, error) {
	val, err := d.Com().GetProperty("IsFTIndexed")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func (d NotesDatabase) IsInMultiDbIndexing() (Boolean, error) {
	val, err := d.Com().GetProperty("IsInMultiDbIndexing")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func (d NotesDatabase) SetIsInMultiDbIndexing(v Boolean) error {
	return d.Com().PutProperty("IsInMultiDbIndexing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func (d NotesDatabase) IsInService() (Boolean, error) {
	val, err := d.Com().GetProperty("IsInService")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func (d NotesDatabase) SetIsInService(v Boolean) error {
	return d.Com().PutProperty("IsInService", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLINK_PROPERTY_DB.html */
func (d NotesDatabase) IsLink() (Boolean, error) {
	val, err := d.Com().GetProperty("IsLink")
	return helpers.CastValue[Boolean](val), err
}

/* TODO: Access type for IsLocallyEncrypted could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func (d NotesDatabase) IsLocallyEncrypted() (Boolean, error) {
	val, err := d.Com().GetProperty("IsLocallyEncrypted")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsLocallyEncrypted(v Boolean) error {
	return d.Com().PutProperty("IsLocallyEncrypted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMULTIDBSEARCH_PROPERTY.html */
func (d NotesDatabase) IsMultiDbSearch() (Boolean, error) {
	val, err := d.Com().GetProperty("IsMultiDbSearch")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISOPEN_PROPERTY.html */
func (d NotesDatabase) IsOpen() (Boolean, error) {
	val, err := d.Com().GetProperty("IsOpen")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPENDINGDELETE_PROPERTY_DB.html */
func (d NotesDatabase) IsPendingDelete() (Boolean, error) {
	val, err := d.Com().GetProperty("IsPendingDelete")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATEADDRESSBOOK_PROPERTY.html */
func (d NotesDatabase) IsPrivateAddressBook() (Boolean, error) {
	val, err := d.Com().GetProperty("IsPrivateAddressBook")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICADDRESSBOOK_PROPERTY.html */
func (d NotesDatabase) IsPublicAddressBook() (Boolean, error) {
	val, err := d.Com().GetProperty("IsPublicAddressBook")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFIXUP_PROPERTY_DB.html */
func (d NotesDatabase) LastFixup() (NotesDateTime, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("LastFixup")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFTINDEXED_PROPERTY.html */
func (d NotesDatabase) LastFTIndexed() (Time, error) {
	val, err := d.Com().GetProperty("LastFTIndexed")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DB.html */
func (d NotesDatabase) LastModified() (Time, error) {
	val, err := d.Com().GetProperty("LastModified")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func (d NotesDatabase) LimitRevisions() (Double, error) {
	val, err := d.Com().GetProperty("LimitRevisions")
	return helpers.CastValue[Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func (d NotesDatabase) SetLimitRevisions(v Double) error {
	return d.Com().PutProperty("LimitRevisions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func (d NotesDatabase) LimitUpdatedBy() (Double, error) {
	val, err := d.Com().GetProperty("LimitUpdatedBy")
	return helpers.CastValue[Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func (d NotesDatabase) SetLimitUpdatedBy(v Double) error {
	return d.Com().PutProperty("LimitUpdatedBy", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func (d NotesDatabase) ListInDbCatalog() (Boolean, error) {
	val, err := d.Com().GetProperty("ListInDbCatalog")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func (d NotesDatabase) SetListInDbCatalog(v Boolean) error {
	return d.Com().PutProperty("ListInDbCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MANAGERS_PROPERTY.html */
func (d NotesDatabase) Managers() ([]String, error) {
	vals, err := d.Com().GetArrayProperty("Managers")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSIZE_PROPERTY_6579_ABOUT.html */
func (d NotesDatabase) MaxSize() (Double, error) {
	val, err := d.Com().GetProperty("MaxSize")
	return helpers.CastValue[Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DATABASE_COM.html */
func (d NotesDatabase) NotesURL() (String, error) {
	val, err := d.Com().GetProperty("NotesURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DB.html */
func (d NotesDatabase) Parent() (NotesSession, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PERCENTUSED_PROPERTY.html */
func (d NotesDatabase) PercentUsed() (Double, error) {
	val, err := d.Com().GetProperty("PercentUsed")
	return helpers.CastValue[Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAID_PROPERTY.html */
func (d NotesDatabase) ReplicaID() (String, error) {
	val, err := d.Com().GetProperty("ReplicaID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATIONINFO_PROPERTY_7679_ABOUT.html */
func (d NotesDatabase) ReplicationInfo() (NotesReplication, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("ReplicationInfo")
	return NewNotesReplication(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY.html */
func (d NotesDatabase) Server() (String, error) {
	val, err := d.Com().GetProperty("Server")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DB.html */
func (d NotesDatabase) Size() (Double, error) {
	val, err := d.Com().GetProperty("Size")
	return helpers.CastValue[Double](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func (d NotesDatabase) SizeQuota() (Long, error) {
	val, err := d.Com().GetProperty("SizeQuota")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func (d NotesDatabase) SetSizeQuota(v Long) error {
	return d.Com().PutProperty("SizeQuota", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func (d NotesDatabase) SizeWarning() (Long, error) {
	val, err := d.Com().GetProperty("SizeWarning")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func (d NotesDatabase) SetSizeWarning(v Long) error {
	return d.Com().PutProperty("SizeWarning", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEMPLATENAME_PROPERTY.html */
func (d NotesDatabase) TemplateName() (String, error) {
	val, err := d.Com().GetProperty("TemplateName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func (d NotesDatabase) Title() (String, error) {
	val, err := d.Com().GetProperty("Title")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func (d NotesDatabase) SetTitle(v String) error {
	return d.Com().PutProperty("Title", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_DB.html */
func (d NotesDatabase) Type() (Long, error) {
	val, err := d.Com().GetProperty("Type")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func (d NotesDatabase) UndeleteExpireTime() (Long, error) {
	val, err := d.Com().GetProperty("UndeleteExpireTime")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func (d NotesDatabase) SetUndeleteExpireTime(v Long) error {
	return d.Com().PutProperty("UndeleteExpireTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDDOCUMENTS_PROPERTY.html */
func (d NotesDatabase) UnprocessedDocuments() (NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("UnprocessedDocuments")
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY.html */
func (d NotesDatabase) Views() ([]NotesView, error) {
	return com.GetObjectArrayProperty(d.Com(), NewNotesView, "Views")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACT_METHOD.html */
func (d NotesDatabase) Compact() (Long, error) {
	val, err := d.Com().CallMethod("Compact")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACTWITHOPTIONS_METHOD_DB.html */
type notesDatabaseCompactWithOptionsParams struct {
	spaceThreshold *String
}

type notesDatabaseCompactWithOptionsParam func(*notesDatabaseCompactWithOptionsParams)

func WithNotesDatabaseCompactWithOptionsSpaceThreshold(spaceThreshold String) notesDatabaseCompactWithOptionsParam {
	return func(c *notesDatabaseCompactWithOptionsParams) {
		c.spaceThreshold = &spaceThreshold
	}
}

func (d NotesDatabase) CompactWithOptions(options any, params ...notesDatabaseCompactWithOptionsParam) (Long, error) {
	paramsStruct := &notesDatabaseCompactWithOptionsParams{}
	paramsOrdered := []interface{}{options}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.spaceThreshold != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.spaceThreshold)
	}
	val, err := d.Com().CallMethod("CompactWithOptions", paramsOrdered...)
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATE_METHOD.html */
type notesDatabaseCreateParams struct {
	maxsize *Byte
}

type notesDatabaseCreateParam func(*notesDatabaseCreateParams)

func WithNotesDatabaseCreateMaxsize(maxsize Byte) notesDatabaseCreateParam {
	return func(c *notesDatabaseCreateParams) {
		c.maxsize = &maxsize
	}
}

func (d NotesDatabase) Create(server String, dbfile String, openFlag Boolean, params ...notesDatabaseCreateParam) error {
	paramsStruct := &notesDatabaseCreateParams{}
	paramsOrdered := []interface{}{server, dbfile, openFlag}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.maxsize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.maxsize)
	}
	_, err := d.Com().CallMethod("Create", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOPY_METHOD.html */
type notesDatabaseCreateCopyParams struct {
	maxsize *Byte
}

type notesDatabaseCreateCopyParam func(*notesDatabaseCreateCopyParams)

func WithNotesDatabaseCreateCopyMaxsize(maxsize Byte) notesDatabaseCreateCopyParam {
	return func(c *notesDatabaseCreateCopyParams) {
		c.maxsize = &maxsize
	}
}

func (d NotesDatabase) CreateCopy(newServer String, newDbFile String, params ...notesDatabaseCreateCopyParam) (NotesDatabase, error) {
	paramsStruct := &notesDatabaseCreateCopyParams{}
	paramsOrdered := []interface{}{newServer, newDbFile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.maxsize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.maxsize)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateCopy", paramsOrdered...)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENT_METHOD.html */
func (d NotesDatabase) CreateDocument() (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDocument")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENTCOLLECTION_METHOD.html */
func (d NotesDatabase) CreateDocumentCollection() (NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDocumentCollection")
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOMINOQUERY_METHOD.html */
func (d NotesDatabase) CreateDominoQuery() (NotesDominoQuery, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDominoQuery")
	return NewNotesDominoQuery(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFROMTEMPLATE_METHOD.html */
type notesDatabaseCreateFromTemplateParams struct {
	maxsize *Byte
}

type notesDatabaseCreateFromTemplateParam func(*notesDatabaseCreateFromTemplateParams)

func WithNotesDatabaseCreateFromTemplateMaxsize(maxsize Byte) notesDatabaseCreateFromTemplateParam {
	return func(c *notesDatabaseCreateFromTemplateParams) {
		c.maxsize = &maxsize
	}
}

func (d NotesDatabase) CreateFromTemplate(newServer String, newDbFile String, inheritFlag Boolean, params ...notesDatabaseCreateFromTemplateParam) (NotesDatabase, error) {
	paramsStruct := &notesDatabaseCreateFromTemplateParams{}
	paramsOrdered := []interface{}{newServer, newDbFile, inheritFlag}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.maxsize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.maxsize)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateFromTemplate", paramsOrdered...)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_METHOD_DB.html */
func (d NotesDatabase) CreateFTIndex(options Long, recreate Boolean) error {
	_, err := d.Com().CallMethod("CreateFTIndex", options, recreate)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENOTECOLLECTION_METHOD_DATABASE.html */
func (d NotesDatabase) CreateNoteCollection(selectAllFlag Boolean) (NotesNoteCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateNoteCollection", selectAllFlag)
	return NewNotesNoteCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEOUTLINE_METHOD_DATABASE.html */
type notesDatabaseCreateOutlineParams struct {
	defaultOutline *Boolean
}

type notesDatabaseCreateOutlineParam func(*notesDatabaseCreateOutlineParams)

func WithNotesDatabaseCreateOutlineDefaultOutline(defaultOutline Boolean) notesDatabaseCreateOutlineParam {
	return func(c *notesDatabaseCreateOutlineParams) {
		c.defaultOutline = &defaultOutline
	}
}

func (d NotesDatabase) CreateOutline(outlinename String, params ...notesDatabaseCreateOutlineParam) (NotesOutline, error) {
	paramsStruct := &notesDatabaseCreateOutlineParams{}
	paramsOrdered := []interface{}{outlinename}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.defaultOutline != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.defaultOutline)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateOutline", paramsOrdered...)
	return NewNotesOutline(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD.html */
func (d NotesDatabase) CreateReplica(newServer String, newDbFile String) (NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateReplica", newServer, newDbFile)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEW_METHOD_DB.html */
type notesDatabaseCreateViewParams struct {
	viewName                           *String
	viewSelectionFormula               *String
	templateView                       *NotesView
	prohibitDesignRefreshModifications *Boolean
}

type notesDatabaseCreateViewParam func(*notesDatabaseCreateViewParams)

func WithNotesDatabaseCreateViewViewName(viewName String) notesDatabaseCreateViewParam {
	return func(c *notesDatabaseCreateViewParams) {
		c.viewName = &viewName
	}
}

func WithNotesDatabaseCreateViewViewSelectionFormula(viewSelectionFormula String) notesDatabaseCreateViewParam {
	return func(c *notesDatabaseCreateViewParams) {
		c.viewSelectionFormula = &viewSelectionFormula
	}
}

func WithNotesDatabaseCreateViewTemplateView(templateView NotesView) notesDatabaseCreateViewParam {
	return func(c *notesDatabaseCreateViewParams) {
		c.templateView = &templateView
	}
}

func WithNotesDatabaseCreateViewProhibitDesignRefreshModifications(prohibitDesignRefreshModifications Boolean) notesDatabaseCreateViewParam {
	return func(c *notesDatabaseCreateViewParams) {
		c.prohibitDesignRefreshModifications = &prohibitDesignRefreshModifications
	}
}

func (d NotesDatabase) CreateView(params ...notesDatabaseCreateViewParam) (NotesView, error) {
	paramsStruct := &notesDatabaseCreateViewParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.viewName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.viewName)
		if paramsStruct.viewSelectionFormula != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.viewSelectionFormula)
			if paramsStruct.templateView != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.templateView)
				if paramsStruct.prohibitDesignRefreshModifications != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.prohibitDesignRefreshModifications)
				}
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateView", paramsOrdered...)
	return NewNotesView(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECRYPT_METHOD_DATABASE.html */
func (d NotesDatabase) Decrypt(def Boolean) error {
	_, err := d.Com().CallMethod("Decrypt", def)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEFOLDER_METHOD_DATABASE.html */
func (d NotesDatabase) EnableFolder(foldername String) error {
	_, err := d.Com().CallMethod("EnableFolder", foldername)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD_DATABASE.html */
func (d NotesDatabase) Encrypt(encryptionStrength Integer, def Boolean) error {
	_, err := d.Com().CallMethod("Encrypt", encryptionStrength, def)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIXUP_METHOD_DB.html */
type notesDatabaseFixupParams struct {
	options *Long
}

type notesDatabaseFixupParam func(*notesDatabaseFixupParams)

func WithNotesDatabaseFixupOptions(options Long) notesDatabaseFixupParam {
	return func(c *notesDatabaseFixupParams) {
		c.options = &options
	}
}

func (d NotesDatabase) Fixup(params ...notesDatabaseFixupParam) error {
	paramsStruct := &notesDatabaseFixupParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	_, err := d.Com().CallMethod("Fixup", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTDOMAINSEARCH_METHOD_9515.html */
type notesDatabaseFTDomainSearchParams struct {
	sortoptions  *Integer
	otheroptions *Integer
	start        *Long
	count        *Long
}

type notesDatabaseFTDomainSearchParam func(*notesDatabaseFTDomainSearchParams)

func WithNotesDatabaseFTDomainSearchSortoptions(sortoptions Integer) notesDatabaseFTDomainSearchParam {
	return func(c *notesDatabaseFTDomainSearchParams) {
		c.sortoptions = &sortoptions
	}
}

func WithNotesDatabaseFTDomainSearchOtheroptions(otheroptions Integer) notesDatabaseFTDomainSearchParam {
	return func(c *notesDatabaseFTDomainSearchParams) {
		c.otheroptions = &otheroptions
	}
}

func WithNotesDatabaseFTDomainSearchStart(start Long) notesDatabaseFTDomainSearchParam {
	return func(c *notesDatabaseFTDomainSearchParams) {
		c.start = &start
	}
}

func WithNotesDatabaseFTDomainSearchCount(count Long) notesDatabaseFTDomainSearchParam {
	return func(c *notesDatabaseFTDomainSearchParams) {
		c.count = &count
	}
}

func (d NotesDatabase) FTDomainSearch(query String, maxDocs Integer, entryform String, params ...notesDatabaseFTDomainSearchParam) (NotesDocument, error) {
	paramsStruct := &notesDatabaseFTDomainSearchParams{}
	paramsOrdered := []interface{}{query, maxDocs, entryform}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.sortoptions)
		if paramsStruct.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.otheroptions)
			if paramsStruct.start != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.start)
				if paramsStruct.count != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.count)
				}
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("FTDomainSearch", paramsOrdered...)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_DB.html */
type notesDatabaseFTSearchParams struct {
	sortoptions  *Integer
	otheroptions *Integer
}

type notesDatabaseFTSearchParam func(*notesDatabaseFTSearchParams)

func WithNotesDatabaseFTSearchSortoptions(sortoptions Integer) notesDatabaseFTSearchParam {
	return func(c *notesDatabaseFTSearchParams) {
		c.sortoptions = &sortoptions
	}
}

func WithNotesDatabaseFTSearchOtheroptions(otheroptions Integer) notesDatabaseFTSearchParam {
	return func(c *notesDatabaseFTSearchParams) {
		c.otheroptions = &otheroptions
	}
}

func (d NotesDatabase) FTSearch(query String, maxdocs Integer, params ...notesDatabaseFTSearchParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDatabaseFTSearchParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.sortoptions)
		if paramsStruct.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.otheroptions)
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("FTSearch", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHRANGE_METHOD_DB.html */
type notesDatabaseFTSearchRangeParams struct {
	sortoptions  *Integer
	otheroptions *Integer
	start        *Long
}

type notesDatabaseFTSearchRangeParam func(*notesDatabaseFTSearchRangeParams)

func WithNotesDatabaseFTSearchRangeSortoptions(sortoptions Integer) notesDatabaseFTSearchRangeParam {
	return func(c *notesDatabaseFTSearchRangeParams) {
		c.sortoptions = &sortoptions
	}
}

func WithNotesDatabaseFTSearchRangeOtheroptions(otheroptions Integer) notesDatabaseFTSearchRangeParam {
	return func(c *notesDatabaseFTSearchRangeParams) {
		c.otheroptions = &otheroptions
	}
}

func WithNotesDatabaseFTSearchRangeStart(start Long) notesDatabaseFTSearchRangeParam {
	return func(c *notesDatabaseFTSearchRangeParams) {
		c.start = &start
	}
}

func (d NotesDatabase) FTSearchRange(query String, maxdocs Integer, params ...notesDatabaseFTSearchRangeParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDatabaseFTSearchRangeParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.sortoptions)
		if paramsStruct.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.otheroptions)
			if paramsStruct.start != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.start)
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("FTSearchRange", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETAGENT_METHOD.html */
func (d NotesDatabase) GetAgent(agentName String) (NotesAgent, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetAgent", agentName)
	return NewNotesAgent(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLREADDOCUMENTS_DATABASE.html */
type notesDatabaseGetAllReadDocumentsParams struct {
	username *String
}

type notesDatabaseGetAllReadDocumentsParam func(*notesDatabaseGetAllReadDocumentsParams)

func WithNotesDatabaseGetAllReadDocumentsUsername(username String) notesDatabaseGetAllReadDocumentsParam {
	return func(c *notesDatabaseGetAllReadDocumentsParams) {
		c.username = &username
	}
}

func (d NotesDatabase) GetAllReadDocuments(params ...notesDatabaseGetAllReadDocumentsParam) (NotesNoteCollection, error) {
	paramsStruct := &notesDatabaseGetAllReadDocumentsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetAllReadDocuments", paramsOrdered...)
	return NewNotesNoteCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLUNREADDOCUMENTS_DATABASE.html */
type notesDatabaseGetAllUnreadDocumentsParams struct {
	username *String
}

type notesDatabaseGetAllUnreadDocumentsParam func(*notesDatabaseGetAllUnreadDocumentsParams)

func WithNotesDatabaseGetAllUnreadDocumentsUsername(username String) notesDatabaseGetAllUnreadDocumentsParam {
	return func(c *notesDatabaseGetAllUnreadDocumentsParams) {
		c.username = &username
	}
}

func (d NotesDatabase) GetAllUnreadDocuments(params ...notesDatabaseGetAllUnreadDocumentsParam) (NotesNoteCollection, error) {
	paramsStruct := &notesDatabaseGetAllUnreadDocumentsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetAllUnreadDocuments", paramsOrdered...)
	return NewNotesNoteCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYID_METHOD.html */
func (d NotesDatabase) GetDocumentByID(noteID String) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocumentByID", noteID)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYUNID_METHOD.html */
func (d NotesDatabase) GetDocumentByUNID(unid String) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocumentByUNID", unid)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYURL_METHOD.html */
type notesDatabaseGetDocumentByURLParams struct {
	reload            *Integer
	urllist           *Integer
	charset           *String
	webusername       *String
	webpassword       *String
	proxywebusername  *String
	proxywebpassword  *String
	returnimmediately *Boolean
}

type notesDatabaseGetDocumentByURLParam func(*notesDatabaseGetDocumentByURLParams)

func WithNotesDatabaseGetDocumentByURLReload(reload Integer) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.reload = &reload
	}
}

func WithNotesDatabaseGetDocumentByURLUrllist(urllist Integer) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.urllist = &urllist
	}
}

func WithNotesDatabaseGetDocumentByURLCharset(charset String) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.charset = &charset
	}
}

func WithNotesDatabaseGetDocumentByURLWebusername(webusername String) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.webusername = &webusername
	}
}

func WithNotesDatabaseGetDocumentByURLWebpassword(webpassword String) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.webpassword = &webpassword
	}
}

func WithNotesDatabaseGetDocumentByURLProxywebusername(proxywebusername String) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.proxywebusername = &proxywebusername
	}
}

func WithNotesDatabaseGetDocumentByURLProxywebpassword(proxywebpassword String) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.proxywebpassword = &proxywebpassword
	}
}

func WithNotesDatabaseGetDocumentByURLReturnimmediately(returnimmediately Boolean) notesDatabaseGetDocumentByURLParam {
	return func(c *notesDatabaseGetDocumentByURLParams) {
		c.returnimmediately = &returnimmediately
	}
}

func (d NotesDatabase) GetDocumentByURL(URL String, params ...notesDatabaseGetDocumentByURLParam) (NotesDocument, error) {
	paramsStruct := &notesDatabaseGetDocumentByURLParams{}
	paramsOrdered := []interface{}{URL}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.reload != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.reload)
		if paramsStruct.urllist != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.urllist)
			if paramsStruct.charset != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.charset)
				if paramsStruct.webusername != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.webusername)
					if paramsStruct.webpassword != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.webpassword)
						if paramsStruct.proxywebusername != nil {
							paramsOrdered = append(paramsOrdered, *paramsStruct.proxywebusername)
							if paramsStruct.proxywebpassword != nil {
								paramsOrdered = append(paramsOrdered, *paramsStruct.proxywebpassword)
								if paramsStruct.returnimmediately != nil {
									paramsOrdered = append(paramsOrdered, *paramsStruct.returnimmediately)
								}
							}
						}
					}
				}
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocumentByURL", paramsOrdered...)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFORM_METHOD.html */
func (d NotesDatabase) GetForm(name String) (NotesForm, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetForm", name)
	return NewNotesForm(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTS_METHOD_DB.html */
type notesDatabaseGetModifiedDocumentsParams struct {
	since     *NotesDateTime
	noteClass *Integer
}

type notesDatabaseGetModifiedDocumentsParam func(*notesDatabaseGetModifiedDocumentsParams)

func WithNotesDatabaseGetModifiedDocumentsSince(since NotesDateTime) notesDatabaseGetModifiedDocumentsParam {
	return func(c *notesDatabaseGetModifiedDocumentsParams) {
		c.since = &since
	}
}

func WithNotesDatabaseGetModifiedDocumentsNoteClass(noteClass Integer) notesDatabaseGetModifiedDocumentsParam {
	return func(c *notesDatabaseGetModifiedDocumentsParams) {
		c.noteClass = &noteClass
	}
}

func (d NotesDatabase) GetModifiedDocuments(params ...notesDatabaseGetModifiedDocumentsParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDatabaseGetModifiedDocumentsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.since != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.since.Com().Dispatch())
		if paramsStruct.noteClass != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.noteClass)
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetModifiedDocuments", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTWITHOPTIONS_METHOD_DATABASE.html */
func (d NotesDatabase) GetModifiedDocumentsWithOptions(modifiedSince NotesDateTime, modifiedUntil NotesDateTime, options Integer) (NotesNoteCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetModifiedDocumentsWithOptions", modifiedSince.Com().Dispatch(), modifiedUntil.Com().Dispatch(), options)
	return NewNotesNoteCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENT_METHOD.html */
func (d NotesDatabase) GetNamedDocument() error {
	_, err := d.Com().CallMethod("GetNamedDocument")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENTCOLLECTION_METHOD.html */
type notesDatabaseGetNamedDocCollectionParams struct {
	name *String
}

type notesDatabaseGetNamedDocCollectionParam func(*notesDatabaseGetNamedDocCollectionParams)

func WithNotesDatabaseGetNamedDocCollectionName(name String) notesDatabaseGetNamedDocCollectionParam {
	return func(c *notesDatabaseGetNamedDocCollectionParams) {
		c.name = &name
	}
}

func (d NotesDatabase) GetNamedDocCollection(params ...notesDatabaseGetNamedDocCollectionParam) error {
	paramsStruct := &notesDatabaseGetNamedDocCollectionParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	_, err := d.Com().CallMethod("GetNamedDocCollection", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOPTION_METHOD_DB.html */
func (d NotesDatabase) GetOption(optionName Integer) (Boolean, error) {
	val, err := d.Com().CallMethod("GetOption", optionName)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOUTLINE_METHOD_DATABASE.html */
func (d NotesDatabase) GetOutline(outlinename String) (NotesOutline, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetOutline", outlinename)
	return NewNotesOutline(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROFILEDOCCOLLECTION_METHOD_DATABASE.html */
type notesDatabaseGetProfileDocCollectionParams struct {
	profilename *String
}

type notesDatabaseGetProfileDocCollectionParam func(*notesDatabaseGetProfileDocCollectionParams)

func WithNotesDatabaseGetProfileDocCollectionProfilename(profilename String) notesDatabaseGetProfileDocCollectionParam {
	return func(c *notesDatabaseGetProfileDocCollectionParams) {
		c.profilename = &profilename
	}
}

func (d NotesDatabase) GetProfileDocCollection(params ...notesDatabaseGetProfileDocCollectionParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDatabaseGetProfileDocCollectionParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.profilename != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.profilename)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetProfileDocCollection", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROFILEDOCUMENT_METHOD.html */
type notesDatabaseGetProfileDocumentParams struct {
	uniqueKey *String
}

type notesDatabaseGetProfileDocumentParam func(*notesDatabaseGetProfileDocumentParams)

func WithNotesDatabaseGetProfileDocumentUniqueKey(uniqueKey String) notesDatabaseGetProfileDocumentParam {
	return func(c *notesDatabaseGetProfileDocumentParams) {
		c.uniqueKey = &uniqueKey
	}
}

func (d NotesDatabase) GetProfileDocument(profilename String, params ...notesDatabaseGetProfileDocumentParam) (NotesDocument, error) {
	paramsStruct := &notesDatabaseGetProfileDocumentParams{}
	paramsOrdered := []interface{}{profilename}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.uniqueKey != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.uniqueKey)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("GetProfileDocument", paramsOrdered...)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETQUERYRESULTSPROCESSOR_METHOD.html */
func (d NotesDatabase) GetQueryResultsProcessor() (NotesQueryResultsProcessor, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetQueryResultsProcessor")
	return NewNotesQueryResultsProcessor(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETURLHEADERINFO_METHOD.html */
type notesDatabaseGetURLHeaderInfoParams struct {
	URL              *String
	headername       *Integer
	Webusername      *String
	Webpassword      *String
	Proxywebusername *String
	Proxywebpassword *String
}

type notesDatabaseGetURLHeaderInfoParam func(*notesDatabaseGetURLHeaderInfoParams)

func WithNotesDatabaseGetURLHeaderInfoURL(URL String) notesDatabaseGetURLHeaderInfoParam {
	return func(c *notesDatabaseGetURLHeaderInfoParams) {
		c.URL = &URL
	}
}

func WithNotesDatabaseGetURLHeaderInfoHeadername(headername Integer) notesDatabaseGetURLHeaderInfoParam {
	return func(c *notesDatabaseGetURLHeaderInfoParams) {
		c.headername = &headername
	}
}

func WithNotesDatabaseGetURLHeaderInfoWebusername(Webusername String) notesDatabaseGetURLHeaderInfoParam {
	return func(c *notesDatabaseGetURLHeaderInfoParams) {
		c.Webusername = &Webusername
	}
}

func WithNotesDatabaseGetURLHeaderInfoWebpassword(Webpassword String) notesDatabaseGetURLHeaderInfoParam {
	return func(c *notesDatabaseGetURLHeaderInfoParams) {
		c.Webpassword = &Webpassword
	}
}

func WithNotesDatabaseGetURLHeaderInfoProxywebusername(Proxywebusername String) notesDatabaseGetURLHeaderInfoParam {
	return func(c *notesDatabaseGetURLHeaderInfoParams) {
		c.Proxywebusername = &Proxywebusername
	}
}

func WithNotesDatabaseGetURLHeaderInfoProxywebpassword(Proxywebpassword String) notesDatabaseGetURLHeaderInfoParam {
	return func(c *notesDatabaseGetURLHeaderInfoParams) {
		c.Proxywebpassword = &Proxywebpassword
	}
}

func (d NotesDatabase) GetURLHeaderInfo(params ...notesDatabaseGetURLHeaderInfoParam) (String, error) {
	paramsStruct := &notesDatabaseGetURLHeaderInfoParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.URL != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.URL)
		if paramsStruct.headername != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.headername)
			if paramsStruct.Webusername != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.Webusername)
				if paramsStruct.Webpassword != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.Webpassword)
					if paramsStruct.Proxywebusername != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.Proxywebusername)
						if paramsStruct.Proxywebpassword != nil {
							paramsOrdered = append(paramsOrdered, *paramsStruct.Proxywebpassword)
						}
					}
				}
			}
		}
	}
	val, err := d.Com().CallMethod("GetURLHeaderInfo", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVIEW_METHOD.html */
func (d NotesDatabase) GetView(viewName String) (NotesView, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetView", viewName)
	return NewNotesView(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GRANTACCESS_METHOD.html */
func (d NotesDatabase) GrantAccess(name String, level Integer) error {
	_, err := d.Com().CallMethod("GrantAccess", name, level)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKFORDELETE_METHOD_DB.html */
func (d NotesDatabase) MarkForDelete() error {
	_, err := d.Com().CallMethod("MarkForDelete")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD.html */
func (d NotesDatabase) Open(server String, dbfile String) (Boolean, error) {
	val, err := d.Com().CallMethod("Open", server, dbfile)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENBYREPLICAID_METHOD.html */
func (d NotesDatabase) OpenByReplicaID(server String, replicaID String) (Boolean, error) {
	val, err := d.Com().CallMethod("OpenByReplicaID", server, replicaID)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENIFMODIFIED_METHOD.html */
func (d NotesDatabase) OpenIfModified(server String, dbfile String, notesDateTime NotesDateTime) (Boolean, error) {
	val, err := d.Com().CallMethod("OpenIfModified", server, dbfile, notesDateTime.Com().Dispatch())
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAIL_METHOD.html */
func (d NotesDatabase) OpenMail() error {
	_, err := d.Com().CallMethod("OpenMail")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENURLDB_METHOD.html */
func (d NotesDatabase) OpenURLDb() (Boolean, error) {
	val, err := d.Com().CallMethod("OpenURLDb")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENWITHFAILOVER_METHOD.html */
func (d NotesDatabase) OpenWithFailover(server String, dbfile String) (Boolean, error) {
	val, err := d.Com().CallMethod("OpenWithFailover", server, dbfile)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESS_METHOD.html */
func (d NotesDatabase) QueryAccess(name String) (Integer, error) {
	val, err := d.Com().CallMethod("QueryAccess", name)
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSPRIVILEGES_METHOD_DB.html */
func (d NotesDatabase) QueryAccessPrivileges(name String) (Long, error) {
	val, err := d.Com().CallMethod("QueryAccessPrivileges", name)
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSROLES_METHOD_DB.html */
func (d NotesDatabase) QueryAccessRoles(name String) ([]String, error) {
	vals, err := d.Com().CallArrayMethod("QueryAccessRoles", name)
	return helpers.CastSlice[String](vals), err
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
func (d NotesDatabase) Replicate(serverName String) (Boolean, error) {
	val, err := d.Com().CallMethod("Replicate", serverName)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REVOKEACCESS_METHOD.html */
func (d NotesDatabase) RevokeAccess(name String) error {
	_, err := d.Com().CallMethod("RevokeAccess", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCH_METHOD.html */
func (d NotesDatabase) Search(formula String, notesDateTime NotesDateTime, maxDocs Integer) (NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("Search", formula, notesDateTime.Com().Dispatch(), maxDocs)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETOPTION_METHOD_DB.html */
func (d NotesDatabase) SetOption(optionName Integer, flag any) error {
	_, err := d.Com().CallMethod("SetOption", optionName, flag)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERIDFORDECRYPT_DATABASE.html */
func (d NotesDatabase) SetUserIDForDecrypt(uid NotesUserID, idFile string, password string) error {
	_, err := d.Com().CallMethod("SetUserIDForDecrypt", uid.Com().Dispatch(), idFile, password)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD_DB.html */
type notesDatabaseSignParams struct {
	documentType     *Integer
	existingSigsOnly *Boolean
	nameStr          *String
	nameStrIsNoteID  *Boolean
}

type notesDatabaseSignParam func(*notesDatabaseSignParams)

func WithNotesDatabaseSignDocumentType(documentType Integer) notesDatabaseSignParam {
	return func(c *notesDatabaseSignParams) {
		c.documentType = &documentType
	}
}

func WithNotesDatabaseSignExistingSigsOnly(existingSigsOnly Boolean) notesDatabaseSignParam {
	return func(c *notesDatabaseSignParams) {
		c.existingSigsOnly = &existingSigsOnly
	}
}

func WithNotesDatabaseSignNameStr(nameStr String) notesDatabaseSignParam {
	return func(c *notesDatabaseSignParams) {
		c.nameStr = &nameStr
	}
}

func WithNotesDatabaseSignNameStrIsNoteID(nameStrIsNoteID Boolean) notesDatabaseSignParam {
	return func(c *notesDatabaseSignParams) {
		c.nameStrIsNoteID = &nameStrIsNoteID
	}
}

func (d NotesDatabase) Sign(params ...notesDatabaseSignParam) error {
	paramsStruct := &notesDatabaseSignParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.documentType != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.documentType)
		if paramsStruct.existingSigsOnly != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.existingSigsOnly)
			if paramsStruct.nameStr != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.nameStr)
				if paramsStruct.nameStrIsNoteID != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.nameStrIsNoteID)
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
type notesDatabaseUnprocessedFTSearchParams struct {
	sortoptions  *Integer
	otheroptions *Integer
}

type notesDatabaseUnprocessedFTSearchParam func(*notesDatabaseUnprocessedFTSearchParams)

func WithNotesDatabaseUnprocessedFTSearchSortoptions(sortoptions Integer) notesDatabaseUnprocessedFTSearchParam {
	return func(c *notesDatabaseUnprocessedFTSearchParams) {
		c.sortoptions = &sortoptions
	}
}

func WithNotesDatabaseUnprocessedFTSearchOtheroptions(otheroptions Integer) notesDatabaseUnprocessedFTSearchParam {
	return func(c *notesDatabaseUnprocessedFTSearchParams) {
		c.otheroptions = &otheroptions
	}
}

func (d NotesDatabase) UnprocessedFTSearch(query String, maxdocs Integer, params ...notesDatabaseUnprocessedFTSearchParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDatabaseUnprocessedFTSearchParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.sortoptions)
		if paramsStruct.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.otheroptions)
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("UnprocessedFTSearch", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDFTSEARCHRANGE_METHOD_DB.html */
type notesDatabaseUnprocessedFTSearchRangeParams struct {
	sortoptions  *Integer
	otheroptions *Integer
	start        *Long
}

type notesDatabaseUnprocessedFTSearchRangeParam func(*notesDatabaseUnprocessedFTSearchRangeParams)

func WithNotesDatabaseUnprocessedFTSearchRangeSortoptions(sortoptions Integer) notesDatabaseUnprocessedFTSearchRangeParam {
	return func(c *notesDatabaseUnprocessedFTSearchRangeParams) {
		c.sortoptions = &sortoptions
	}
}

func WithNotesDatabaseUnprocessedFTSearchRangeOtheroptions(otheroptions Integer) notesDatabaseUnprocessedFTSearchRangeParam {
	return func(c *notesDatabaseUnprocessedFTSearchRangeParams) {
		c.otheroptions = &otheroptions
	}
}

func WithNotesDatabaseUnprocessedFTSearchRangeStart(start Long) notesDatabaseUnprocessedFTSearchRangeParam {
	return func(c *notesDatabaseUnprocessedFTSearchRangeParams) {
		c.start = &start
	}
}

func (d NotesDatabase) UnprocessedFTSearchRange(query String, maxdocs Integer, params ...notesDatabaseUnprocessedFTSearchRangeParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDatabaseUnprocessedFTSearchRangeParams{}
	paramsOrdered := []interface{}{query, maxdocs}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.sortoptions != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.sortoptions)
		if paramsStruct.otheroptions != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.otheroptions)
			if paramsStruct.start != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.start)
			}
		}
	}
	dispatchPtr, err := d.Com().CallObjectMethod("UnprocessedFTSearchRange", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDSEARCH_METHOD.html */
func (d NotesDatabase) UnprocessedSearch(formula String, notesDateTime NotesDateTime, maxDocs Integer) (NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("UnprocessedSearch", formula, notesDateTime.Com().Dispatch(), maxDocs)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEFTINDEX_METHOD.html */
func (d NotesDatabase) UpdateFTIndex(createFlag Boolean) error {
	_, err := d.Com().CallMethod("UpdateFTIndex", createFlag)
	return err
}
