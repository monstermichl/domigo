/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATABASE_CLASS.html */
package domigo

import (
	"errors"

	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesDatabase struct {
	NotesStruct
}

func newNotesDatabase(dispatchPtr *ole.IDispatch) NotesDatabase {
	return NotesDatabase{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACL_PROPERTY.html */
func (d NotesDatabase) ACL() (NotesACL, error) {
	return getComObjectProperty(d, newNotesACL, "ACL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLACTIVITYLOG_PROPERTY_DB.html */
func (d NotesDatabase) ACLActivityLog() ([]String, error) {
	return getComArrayProperty[String](d, "ACLActivityLog")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AGENTS_PROPERTY.html */
func (d NotesDatabase) Agents() ([]NotesAgent, error) {
	return com.GetObjectArrayProperty(d.com(), newNotesAgent, "Agents")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALLDOCUMENTS_PROPERTY.html */
func (d NotesDatabase) AllDocuments() (NotesDocumentCollection, error) {
	return getComObjectProperty(d, newNotesDocumentCollection, "AllDocuments")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func (d NotesDatabase) Categories() (String, error) {
	return getComProperty[String](d, "Categories")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func (d NotesDatabase) SetCategories(v String) error {
	return putComProperty(d, "Categories", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DB.html */
func (d NotesDatabase) Created() (Time, error) {
	return getComProperty[Time](d, "Created")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTACCESSLEVEL_PROPERTY.html */
func (d NotesDatabase) CurrentAccessLevel() (Long, error) {
	return getComProperty[Long](d, "CurrentAccessLevel")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func (d NotesDatabase) DelayUpdates() (Boolean, error) {
	return getComProperty[Boolean](d, "DelayUpdates")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func (d NotesDatabase) SetDelayUpdates(v Boolean) error {
	return putComProperty(d, "DelayUpdates", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNTEMPLATENAME_PROPERTY.html */
func (d NotesDatabase) DesignTemplateName() (String, error) {
	return getComProperty[String](d, "DesignTemplateName")
}

/* TODO: Access type for EncryptionStrength could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func (d NotesDatabase) EncryptionStrength() (Integer, error) {
	return getComProperty[Integer](d, "EncryptionStrength")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func (d NotesDatabase) SetEncryptionStrength(v Integer) error {
	return putComProperty(d, "EncryptionStrength", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEFORMAT_PROPERTY_DB.html */
func (d NotesDatabase) FileFormat() (Long, error) {
	return getComProperty[Long](d, "FileFormat")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILENAME_PROPERTY.html */
func (d NotesDatabase) FileName() (String, error) {
	return getComProperty[String](d, "FileName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEPATH_PROPERTY.html */
func (d NotesDatabase) FilePath() (String, error) {
	return getComProperty[String](d, "FilePath")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func (d NotesDatabase) FolderReferencesEnabled() (Boolean, error) {
	return getComProperty[Boolean](d, "FolderReferencesEnabled")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func (d NotesDatabase) SetFolderReferencesEnabled(v Boolean) error {
	return putComProperty(d, "FolderReferencesEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMS_PROPERTY.html */
func (d NotesDatabase) Forms() ([]NotesForm, error) {
	return com.GetObjectArrayProperty(d.com(), newNotesForm, "Forms")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func (d NotesDatabase) FTIndexFrequency() (Long, error) {
	return getComProperty[Long](d, "FTIndexFrequency")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func (d NotesDatabase) SetFTIndexFrequency(v Long) error {
	return putComProperty(d, "FTIndexFrequency", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DATABASE_COM.html */
func (d NotesDatabase) HttpURL() (String, error) {
	return getComProperty[String](d, "HttpURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func (d NotesDatabase) IsClusterReplication() (Boolean, error) {
	return getComProperty[Boolean](d, "IsClusterReplication")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func (d NotesDatabase) SetIsClusterReplication(v Boolean) error {
	return putComProperty(d, "IsClusterReplication", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFIGURATIONDIRECTORY_PROPERTY_DB.html */
func (d NotesDatabase) IsConfigurationDirectory() (Boolean, error) {
	return getComProperty[Boolean](d, "IsConfigurationDirectory")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICREADER_PROPERTY_DB.html */
func (d NotesDatabase) IsCurrentAccessPublicReader() (Boolean, error) {
	return getComProperty[Boolean](d, "IsCurrentAccessPublicReader")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICWRITER_PROPERTY_DB.html */
func (d NotesDatabase) IsCurrentAccessPublicWriter() (Boolean, error) {
	return getComProperty[Boolean](d, "IsCurrentAccessPublicWriter")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) IsDesignLockingEnabled() (Boolean, error) {
	return getComProperty[Boolean](d, "IsDesignLockingEnabled")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsDesignLockingEnabled(v Boolean) error {
	return putComProperty(d, "IsDesignLockingEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDIRECTORYCATALOG_DB.html */
func (d NotesDatabase) IsDirectoryCatalog() (Boolean, error) {
	return getComProperty[Boolean](d, "IsDirectoryCatalog")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) IsDocumentLockingEnabled() (Boolean, error) {
	return getComProperty[Boolean](d, "IsDocumentLockingEnabled")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func (d NotesDatabase) SetIsDocumentLockingEnabled(v Boolean) error {
	return putComProperty(d, "IsDocumentLockingEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFTINDEXED_PROPERTY.html */
func (d NotesDatabase) IsFTIndexed() (Boolean, error) {
	return getComProperty[Boolean](d, "IsFTIndexed")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func (d NotesDatabase) IsInMultiDbIndexing() (Boolean, error) {
	return getComProperty[Boolean](d, "IsInMultiDbIndexing")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func (d NotesDatabase) SetIsInMultiDbIndexing(v Boolean) error {
	return putComProperty(d, "IsInMultiDbIndexing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func (d NotesDatabase) IsInService() (Boolean, error) {
	return getComProperty[Boolean](d, "IsInService")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func (d NotesDatabase) SetIsInService(v Boolean) error {
	return putComProperty(d, "IsInService", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLINK_PROPERTY_DB.html */
func (d NotesDatabase) IsLink() (Boolean, error) {
	return getComProperty[Boolean](d, "IsLink")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func (d NotesDatabase) IsLocallyEncrypted() (Boolean, error) {
	return getComProperty[Boolean](d, "IsLocallyEncrypted")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMULTIDBSEARCH_PROPERTY.html */
func (d NotesDatabase) IsMultiDbSearch() (Boolean, error) {
	return getComProperty[Boolean](d, "IsMultiDbSearch")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISOPEN_PROPERTY.html */
func (d NotesDatabase) IsOpen() (Boolean, error) {
	return getComProperty[Boolean](d, "IsOpen")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPENDINGDELETE_PROPERTY_DB.html */
func (d NotesDatabase) IsPendingDelete() (Boolean, error) {
	return getComProperty[Boolean](d, "IsPendingDelete")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATEADDRESSBOOK_PROPERTY.html */
func (d NotesDatabase) IsPrivateAddressBook() (Boolean, error) {
	return getComProperty[Boolean](d, "IsPrivateAddressBook")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICADDRESSBOOK_PROPERTY.html */
func (d NotesDatabase) IsPublicAddressBook() (Boolean, error) {
	return getComProperty[Boolean](d, "IsPublicAddressBook")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFIXUP_PROPERTY_DB.html */
func (d NotesDatabase) LastFixup() (NotesDateTime, error) {
	return getComObjectProperty(d, newNotesDateTime, "LastFixup")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFTINDEXED_PROPERTY.html */
func (d NotesDatabase) LastFTIndexed() (Time, error) {
	return getComProperty[Time](d, "LastFTIndexed")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DB.html */
func (d NotesDatabase) LastModified() (Time, error) {
	return getComProperty[Time](d, "LastModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func (d NotesDatabase) LimitRevisions() (Double, error) {
	return getComProperty[Double](d, "LimitRevisions")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func (d NotesDatabase) SetLimitRevisions(v Double) error {
	return putComProperty(d, "LimitRevisions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func (d NotesDatabase) LimitUpdatedBy() (Double, error) {
	return getComProperty[Double](d, "LimitUpdatedBy")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func (d NotesDatabase) SetLimitUpdatedBy(v Double) error {
	return putComProperty(d, "LimitUpdatedBy", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func (d NotesDatabase) ListInDbCatalog() (Boolean, error) {
	return getComProperty[Boolean](d, "ListInDbCatalog")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func (d NotesDatabase) SetListInDbCatalog(v Boolean) error {
	return putComProperty(d, "ListInDbCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MANAGERS_PROPERTY.html */
func (d NotesDatabase) Managers() ([]String, error) {
	return getComArrayProperty[String](d, "Managers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSIZE_PROPERTY_6579_ABOUT.html */
func (d NotesDatabase) MaxSize() (Double, error) {
	return getComProperty[Double](d, "MaxSize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DATABASE_COM.html */
func (d NotesDatabase) NotesURL() (String, error) {
	return getComProperty[String](d, "NotesURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DB.html */
func (d NotesDatabase) Parent() (NotesSession, error) {
	return getComObjectProperty(d, newNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PERCENTUSED_PROPERTY.html */
func (d NotesDatabase) PercentUsed() (Double, error) {
	return getComProperty[Double](d, "PercentUsed")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAID_PROPERTY.html */
func (d NotesDatabase) ReplicaID() (String, error) {
	return getComProperty[String](d, "ReplicaID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATIONINFO_PROPERTY_7679_ABOUT.html */
func (d NotesDatabase) ReplicationInfo() (NotesReplication, error) {
	return getComObjectProperty(d, newNotesReplication, "ReplicationInfo")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY.html */
func (d NotesDatabase) Server() (String, error) {
	return getComProperty[String](d, "Server")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DB.html */
func (d NotesDatabase) Size() (Double, error) {
	return getComProperty[Double](d, "Size")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func (d NotesDatabase) SizeQuota() (Long, error) {
	return getComProperty[Long](d, "SizeQuota")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func (d NotesDatabase) SetSizeQuota(v Long) error {
	return putComProperty(d, "SizeQuota", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func (d NotesDatabase) SizeWarning() (Long, error) {
	return getComProperty[Long](d, "SizeWarning")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func (d NotesDatabase) SetSizeWarning(v Long) error {
	return putComProperty(d, "SizeWarning", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEMPLATENAME_PROPERTY.html */
func (d NotesDatabase) TemplateName() (String, error) {
	return getComProperty[String](d, "TemplateName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func (d NotesDatabase) Title() (String, error) {
	return getComProperty[String](d, "Title")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func (d NotesDatabase) SetTitle(v String) error {
	return putComProperty(d, "Title", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_DB.html */
func (d NotesDatabase) Type() (Long, error) {
	return getComProperty[Long](d, "Type")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func (d NotesDatabase) UndeleteExpireTime() (Long, error) {
	return getComProperty[Long](d, "UndeleteExpireTime")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func (d NotesDatabase) SetUndeleteExpireTime(v Long) error {
	return putComProperty(d, "UndeleteExpireTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDDOCUMENTS_PROPERTY.html */
func (d NotesDatabase) UnprocessedDocuments() (NotesDocumentCollection, error) {
	return getComObjectProperty(d, newNotesDocumentCollection, "UnprocessedDocuments")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY.html */
func (d NotesDatabase) Views() ([]NotesView, error) {
	return com.GetObjectArrayProperty(d.com(), newNotesView, "Views")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACT_METHOD.html */
func (d NotesDatabase) Compact() (Long, error) {
	return callComMethod[Long](d, "Compact")
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
	return callComMethod[Long](d, "CompactWithOptions", paramsOrdered...)
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
	return errors.New("this method is supported in LotusScript® only. For COM, use CreateDatabase in NotesDbDirectory (see https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATE_METHOD.html)")
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
	return callComObjectMethod(d, newNotesDatabase, "CreateCopy", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENT_METHOD.html */
func (d NotesDatabase) CreateDocument() (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "CreateDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENTCOLLECTION_METHOD.html */
func (d NotesDatabase) CreateDocumentCollection() (NotesDocumentCollection, error) {
	return callComObjectMethod(d, newNotesDocumentCollection, "CreateDocumentCollection")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOMINOQUERY_METHOD.html */
func (d NotesDatabase) CreateDominoQuery() (NotesDominoQuery, error) {
	return callComObjectMethod(d, newNotesDominoQuery, "CreateDominoQuery")
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
	return callComObjectMethod(d, newNotesDatabase, "CreateFromTemplate", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_METHOD_DB.html */
func (d NotesDatabase) CreateFTIndex(options Long, recreate Boolean) error {
	return callComVoidMethod(d, "CreateFTIndex", options, recreate)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENOTECOLLECTION_METHOD_DATABASE.html */
func (d NotesDatabase) CreateNoteCollection(selectAllFlag Boolean) (NotesNoteCollection, error) {
	return callComObjectMethod(d, newNotesNoteCollection, "CreateNoteCollection", selectAllFlag)
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
	return callComObjectMethod(d, newNotesOutline, "CreateOutline", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD.html */
func (d NotesDatabase) CreateReplica(newServer String, newDbFile String) (NotesDatabase, error) {
	return callComObjectMethod(d, newNotesDatabase, "CreateReplica", newServer, newDbFile)
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
	return callComObjectMethod(d, newNotesView, "CreateView", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECRYPT_METHOD_DATABASE.html */
func (d NotesDatabase) Decrypt(def Boolean) error {
	return callComVoidMethod(d, "Decrypt", def)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEFOLDER_METHOD_DATABASE.html */
func (d NotesDatabase) EnableFolder(foldername String) error {
	return callComVoidMethod(d, "EnableFolder", foldername)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD_DATABASE.html */
func (d NotesDatabase) Encrypt(encryptionStrength Integer, def Boolean) error {
	return callComVoidMethod(d, "Encrypt", encryptionStrength, def)
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
	return callComVoidMethod(d, "Fixup", paramsOrdered...)
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
	return callComObjectMethod(d, newNotesDocument, "FTDomainSearch", paramsOrdered...)
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
	return callComObjectMethod(d, newNotesDocumentCollection, "FTSearch", paramsOrdered...)
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
	return callComObjectMethod(d, newNotesDocumentCollection, "FTSearchRange", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETAGENT_METHOD.html */
func (d NotesDatabase) GetAgent(agentName String) (NotesAgent, error) {
	return callComObjectMethod(d, newNotesAgent, "GetAgent", agentName)
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
	return callComObjectMethod(d, newNotesNoteCollection, "GetAllReadDocuments", paramsOrdered...)
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
	return callComObjectMethod(d, newNotesNoteCollection, "GetAllUnreadDocuments", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYID_METHOD.html */
func (d NotesDatabase) GetDocumentByID(noteID String) (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetDocumentByID", noteID)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYUNID_METHOD.html */
func (d NotesDatabase) GetDocumentByUNID(unid String) (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetDocumentByUNID", unid)
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
	return callComObjectMethod(d, newNotesDocument, "GetDocumentByURL", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFORM_METHOD.html */
func (d NotesDatabase) GetForm(name String) (NotesForm, error) {
	return callComObjectMethod(d, newNotesForm, "GetForm", name)
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
		paramsOrdered = append(paramsOrdered, *paramsStruct.since)
		if paramsStruct.noteClass != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.noteClass)
		}
	}
	return callComObjectMethod(d, newNotesDocumentCollection, "GetModifiedDocuments", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTWITHOPTIONS_METHOD_DATABASE.html */
func (d NotesDatabase) GetModifiedDocumentsWithOptions(modifiedSince NotesDateTime, modifiedUntil NotesDateTime, options Integer) (NotesNoteCollection, error) {
	return callComObjectMethod(d, newNotesNoteCollection, "GetModifiedDocumentsWithOptions", modifiedSince, modifiedUntil, options)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENT_METHOD.html */
func (d NotesDatabase) GetNamedDocument() error {
	return callComVoidMethod(d, "GetNamedDocument")
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
	return callComVoidMethod(d, "GetNamedDocCollection", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOPTION_METHOD_DB.html */
func (d NotesDatabase) GetOption(optionName Integer) (Boolean, error) {
	return callComMethod[Boolean](d, "GetOption", optionName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOUTLINE_METHOD_DATABASE.html */
func (d NotesDatabase) GetOutline(outlinename String) (NotesOutline, error) {
	return callComObjectMethod(d, newNotesOutline, "GetOutline", outlinename)
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
	return callComObjectMethod(d, newNotesDocumentCollection, "GetProfileDocCollection", paramsOrdered...)
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
	return callComObjectMethod(d, newNotesDocument, "GetProfileDocument", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETQUERYRESULTSPROCESSOR_METHOD.html */
func (d NotesDatabase) GetQueryResultsProcessor() (NotesQueryResultsProcessor, error) {
	return callComObjectMethod(d, newNotesQueryResultsProcessor, "GetQueryResultsProcessor")
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
	return callComMethod[String](d, "GetURLHeaderInfo", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVIEW_METHOD.html */
func (d NotesDatabase) GetView(viewName String) (NotesView, error) {
	return callComObjectMethod(d, newNotesView, "GetView", viewName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GRANTACCESS_METHOD.html */
func (d NotesDatabase) GrantAccess(name String, level Integer) error {
	return callComVoidMethod(d, "GrantAccess", name, level)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKFORDELETE_METHOD_DB.html */
func (d NotesDatabase) MarkForDelete() error {
	return callComVoidMethod(d, "MarkForDelete")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD.html */
func (d NotesDatabase) Open(server String, dbfile String) (Boolean, error) {
	return callComMethod[Boolean](d, "Open", server, dbfile)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENBYREPLICAID_METHOD.html */
func (d NotesDatabase) OpenByReplicaID(server String, replicaID String) (Boolean, error) {
	return callComMethod[Boolean](d, "OpenByReplicaID", server, replicaID)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENIFMODIFIED_METHOD.html */
func (d NotesDatabase) OpenIfModified(server String, dbfile String, notesDateTime NotesDateTime) (Boolean, error) {
	return callComMethod[Boolean](d, "OpenIfModified", server, dbfile, notesDateTime)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAIL_METHOD.html */
func (d NotesDatabase) OpenMail() error {
	return callComVoidMethod(d, "OpenMail")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENURLDB_METHOD.html */
func (d NotesDatabase) OpenURLDb() (Boolean, error) {
	return callComMethod[Boolean](d, "OpenURLDb")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENWITHFAILOVER_METHOD.html */
func (d NotesDatabase) OpenWithFailover(server String, dbfile String) (Boolean, error) {
	return callComMethod[Boolean](d, "OpenWithFailover", server, dbfile)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESS_METHOD.html */
func (d NotesDatabase) QueryAccess(name String) (Integer, error) {
	return callComMethod[Integer](d, "QueryAccess", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSPRIVILEGES_METHOD_DB.html */
func (d NotesDatabase) QueryAccessPrivileges(name String) (Long, error) {
	return callComMethod[Long](d, "QueryAccessPrivileges", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSROLES_METHOD_DB.html */
func (d NotesDatabase) QueryAccessRoles(name String) ([]String, error) {
	return callComArrayMethod[String](d, "QueryAccessRoles", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DB.html */
func (d NotesDatabase) Remove() error {
	return callComVoidMethod(d, "Remove")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFTINDEX_METHOD_DB.html */
func (d NotesDatabase) RemoveFTIndex() error {
	return callComVoidMethod(d, "RemoveFTIndex")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATE_METHOD.html */
func (d NotesDatabase) Replicate(serverName String) (Boolean, error) {
	return callComMethod[Boolean](d, "Replicate", serverName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REVOKEACCESS_METHOD.html */
func (d NotesDatabase) RevokeAccess(name String) error {
	return callComVoidMethod(d, "RevokeAccess", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCH_METHOD.html */
func (d NotesDatabase) Search(formula String, notesDateTime NotesDateTime, maxDocs Integer) (NotesDocumentCollection, error) {
	return callComObjectMethod(d, newNotesDocumentCollection, "Search", formula, notesDateTime, maxDocs)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETOPTION_METHOD_DB.html */
func (d NotesDatabase) SetOption(optionName Integer, flag any) error {
	return callComVoidMethod(d, "SetOption", optionName, flag)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERIDFORDECRYPT_DATABASE.html */
func (d NotesDatabase) SetUserIDForDecrypt(uid NotesUserID, idFile string, password string) error {
	return callComVoidMethod(d, "SetUserIDForDecrypt", uid, idFile, password)
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
	return callComVoidMethod(d, "Sign", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONBEGIN_METHOD.html */
func (d NotesDatabase) TransactionBegin() error {
	return callComVoidMethod(d, "TransactionBegin")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONCOMMIT_METHOD.html */
func (d NotesDatabase) TransactionCommit() error {
	return callComVoidMethod(d, "TransactionCommit")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONROLLBACK_METHOD.html */
func (d NotesDatabase) TransactionRollback() error {
	return callComVoidMethod(d, "TransactionRollback")
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
	return callComObjectMethod(d, newNotesDocumentCollection, "UnprocessedFTSearch", paramsOrdered...)
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
	return callComObjectMethod(d, newNotesDocumentCollection, "UnprocessedFTSearchRange", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDSEARCH_METHOD.html */
func (d NotesDatabase) UnprocessedSearch(formula String, notesDateTime NotesDateTime, maxDocs Integer) (NotesDocumentCollection, error) {
	return callComObjectMethod(d, newNotesDocumentCollection, "UnprocessedSearch", formula, notesDateTime, maxDocs)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEFTINDEX_METHOD.html */
func (d NotesDatabase) UpdateFTIndex(createFlag Boolean) error {
	return callComVoidMethod(d, "UpdateFTIndex", createFlag)
}
