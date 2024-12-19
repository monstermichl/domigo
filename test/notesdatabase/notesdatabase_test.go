package notesform_test

import (
	"testing"

	domigo "github.com/monstermichl/domigo/domino"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var database domigo.NotesDatabase

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := domigo.Initialize()
	database, _ = testhelpers.CreateTestDatabase(session)

	defer database.Release()
	defer database.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACL_PROPERTY.html */
func TestACL(t *testing.T) {
	_, err := database.ACL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLACTIVITYLOG_PROPERTY_DB.html */
func TestACLActivityLog(t *testing.T) {
	_, err := database.ACLActivityLog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AGENTS_PROPERTY.html */
func TestAgents(t *testing.T) {
	_, err := database.Agents()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALLDOCUMENTS_PROPERTY.html */
func TestAllDocuments(t *testing.T) {
	_, err := database.AllDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func TestCategories(t *testing.T) {
	_, err := database.Categories()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CATEGORIES_PROPERTY.html */
func TestSetCategories(t *testing.T) {
	s, err := database.Categories()
	require.Nil(t, err)

	err = database.SetCategories(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DB.html */
func TestCreated(t *testing.T) {
	_, err := database.Created()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTACCESSLEVEL_PROPERTY.html */
func TestCurrentAccessLevel(t *testing.T) {
	_, err := database.CurrentAccessLevel()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func TestDelayUpdates(t *testing.T) {
	_, err := database.DelayUpdates()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELAYUPDATES_PROPERTY.html */
func TestSetDelayUpdates(t *testing.T) {
	s, err := database.DelayUpdates()
	require.Nil(t, err)

	err = database.SetDelayUpdates(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNTEMPLATENAME_PROPERTY.html */
func TestDesignTemplateName(t *testing.T) {
	_, err := database.DesignTemplateName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD_DATABASE.html */
func TestEncrypt(t *testing.T) {
	err := database.Encrypt(5, false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func TestEncryptionStrength(t *testing.T) {
	_, err := database.EncryptionStrength()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONSTRENGTH_PROPERTY_DB.html */
func TestSetEncryptionStrength(t *testing.T) {
	s, err := database.EncryptionStrength()
	require.Nil(t, err)

	err = database.SetEncryptionStrength(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEFORMAT_PROPERTY_DB.html */
func TestFileFormat(t *testing.T) {
	_, err := database.FileFormat()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILENAME_PROPERTY.html */
func TestFileName(t *testing.T) {
	_, err := database.FileName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEPATH_PROPERTY.html */
func TestFilePath(t *testing.T) {
	_, err := database.FilePath()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func TestFolderReferencesEnabled(t *testing.T) {
	_, err := database.FolderReferencesEnabled()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCESENABLED_PROPERTY_4434_ABOUT.html */
func TestSetFolderReferencesEnabled(t *testing.T) {
	s, err := database.FolderReferencesEnabled()
	require.Nil(t, err)

	err = database.SetFolderReferencesEnabled(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMS_PROPERTY.html */
func TestForms(t *testing.T) {
	_, err := database.Forms()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func TestFTIndexFrequency(t *testing.T) {
	_, err := database.FTIndexFrequency()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTINDEXFREQUENCY_PROPERTY_DB.html */
func TestSetFTIndexFrequency(t *testing.T) {
	s, err := database.FTIndexFrequency()
	require.Nil(t, err)

	err = database.SetFTIndexFrequency(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DATABASE_COM.html */
func TestHttpURL(t *testing.T) {
	_, err := database.HttpURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func TestIsClusterReplication(t *testing.T) {
	_, err := database.IsClusterReplication()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCLUSTERREPLICATION_PROPERTY_DB.html */
func TestSetIsClusterReplication(t *testing.T) {
	s, err := database.IsClusterReplication()
	require.Nil(t, err)

	err = database.SetIsClusterReplication(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFIGURATIONDIRECTORY_PROPERTY_DB.html */
func TestIsConfigurationDirectory(t *testing.T) {
	_, err := database.IsConfigurationDirectory()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICREADER_PROPERTY_DB.html */
func TestIsCurrentAccessPublicReader(t *testing.T) {
	_, err := database.IsCurrentAccessPublicReader()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENTACCESSPUBLICWRITER_PROPERTY_DB.html */
func TestIsCurrentAccessPublicWriter(t *testing.T) {
	_, err := database.IsCurrentAccessPublicWriter()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func TestIsDesignLockingEnabled(t *testing.T) {
	_, err := database.IsDesignLockingEnabled()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDESIGNLOCKINGENABLED_PROPERTY_DB.html */
func TestSetIsDesignLockingEnabled(t *testing.T) {
	s, err := database.IsDesignLockingEnabled()
	require.Nil(t, err)

	err = database.SetIsDesignLockingEnabled(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDIRECTORYCATALOG_DB.html */
func TestIsDirectoryCatalog(t *testing.T) {
	_, err := database.IsDirectoryCatalog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func TestIsDocumentLockingEnabled(t *testing.T) {
	_, err := database.IsDocumentLockingEnabled()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENTLOCKINGENABLED_PROPERTY_DB.html */
func TestSetIsDocumentLockingEnabled(t *testing.T) {
	s, err := database.IsDocumentLockingEnabled()
	require.Nil(t, err)

	err = database.SetIsDocumentLockingEnabled(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFTINDEXED_PROPERTY.html */
func TestIsFTIndexed(t *testing.T) {
	_, err := database.IsFTIndexed()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func TestIsInMultiDbIndexing(t *testing.T) {
	_, err := database.IsInMultiDbIndexing()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINMULTIDBINDEXING_PROPERTY_DB.html */
func TestSetIsInMultiDbIndexing(t *testing.T) {
	s, err := database.IsInMultiDbIndexing()
	require.Nil(t, err)

	err = database.SetIsInMultiDbIndexing(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func TestIsInService(t *testing.T) {
	_, err := database.IsInService()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINSERVICE_PROPERTY_DB.html */
func TestSetIsInService(t *testing.T) {
	s, err := database.IsInService()
	require.Nil(t, err)

	err = database.SetIsInService(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLINK_PROPERTY_DB.html */
func TestIsLink(t *testing.T) {
	_, err := database.IsLink()
	require.Nil(t, err)
}

/* TODO: Access type for IsLocallyEncrypted could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func TestIsLocallyEncrypted(t *testing.T) {
	_, err := database.IsLocallyEncrypted()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISLOCALLYENCRYPTED_PROPERTY_DB.html */
func TestSetIsLocallyEncrypted(t *testing.T) {
	s, err := database.IsLocallyEncrypted()
	require.Nil(t, err)

	err = database.SetIsLocallyEncrypted(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMULTIDBSEARCH_PROPERTY.html */
func TestIsMultiDbSearch(t *testing.T) {
	_, err := database.IsMultiDbSearch()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISOPEN_PROPERTY.html */
func TestIsOpen(t *testing.T) {
	_, err := database.IsOpen()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPENDINGDELETE_PROPERTY_DB.html */
func TestIsPendingDelete(t *testing.T) {
	_, err := database.IsPendingDelete()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATEADDRESSBOOK_PROPERTY.html */
func TestIsPrivateAddressBook(t *testing.T) {
	_, err := database.IsPrivateAddressBook()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICADDRESSBOOK_PROPERTY.html */
func TestIsPublicAddressBook(t *testing.T) {
	_, err := database.IsPublicAddressBook()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFIXUP_PROPERTY_DB.html */
func TestLastFixup(t *testing.T) {
	_, err := database.LastFixup()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTFTINDEXED_PROPERTY.html */
func TestLastFTIndexed(t *testing.T) {
	_, err := database.LastFTIndexed()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DB.html */
func TestLastModified(t *testing.T) {
	_, err := database.LastModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func TestLimitRevisions(t *testing.T) {
	_, err := database.LimitRevisions()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITREVISIONS_PROPERTY_DB.html */
func TestSetLimitRevisions(t *testing.T) {
	s, err := database.LimitRevisions()
	require.Nil(t, err)

	err = database.SetLimitRevisions(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func TestLimitUpdatedBy(t *testing.T) {
	_, err := database.LimitUpdatedBy()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITUPDATEDBY_PROPERTY_DB.html */
func TestSetLimitUpdatedBy(t *testing.T) {
	s, err := database.LimitUpdatedBy()
	require.Nil(t, err)

	err = database.SetLimitUpdatedBy(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func TestListInDbCatalog(t *testing.T) {
	_, err := database.ListInDbCatalog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDBCATALOG_PROPERTY_DB.html */
func TestSetListInDbCatalog(t *testing.T) {
	s, err := database.ListInDbCatalog()
	require.Nil(t, err)

	err = database.SetListInDbCatalog(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MANAGERS_PROPERTY.html */
func TestManagers(t *testing.T) {
	_, err := database.Managers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSIZE_PROPERTY_6579_ABOUT.html */
func TestMaxSize(t *testing.T) {
	_, err := database.MaxSize()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DATABASE_COM.html */
func TestNotesURL(t *testing.T) {
	_, err := database.NotesURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PERCENTUSED_PROPERTY.html */
func TestPercentUsed(t *testing.T) {
	_, err := database.PercentUsed()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAID_PROPERTY.html */
func TestReplicaID(t *testing.T) {
	_, err := database.ReplicaID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATIONINFO_PROPERTY_7679_ABOUT.html */
func TestReplicationInfo(t *testing.T) {
	_, err := database.ReplicationInfo()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY.html */
func TestServer(t *testing.T) {
	_, err := database.Server()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DB.html */
func TestSize(t *testing.T) {
	_, err := database.Size()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func TestSizeQuota(t *testing.T) {
	_, err := database.SizeQuota()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEQUOTA_PROPERTY.html */
func TestSetSizeQuota(t *testing.T) {
	s, err := database.SizeQuota()
	require.Nil(t, err)

	err = database.SetSizeQuota(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func TestSizeWarning(t *testing.T) {
	_, err := database.SizeWarning()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZEWARNING_PROPERTY_DB.html */
func TestSetSizeWarning(t *testing.T) {
	s, err := database.SizeWarning()
	require.Nil(t, err)

	err = database.SetSizeWarning(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEMPLATENAME_PROPERTY.html */
func TestTemplateName(t *testing.T) {
	_, err := database.TemplateName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func TestTitle(t *testing.T) {
	_, err := database.Title()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_DB.html */
func TestSetTitle(t *testing.T) {
	s, err := database.Title()
	require.Nil(t, err)

	err = database.SetTitle(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_DB.html */
func TestType(t *testing.T) {
	_, err := database.Type()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func TestUndeleteExpireTime(t *testing.T) {
	_, err := database.UndeleteExpireTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDELETEEXPIRETIME_PROPERTY_DB.html */
func TestSetUndeleteExpireTime(t *testing.T) {
	s, err := database.UndeleteExpireTime()
	require.Nil(t, err)

	err = database.SetUndeleteExpireTime(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDDOCUMENTS_PROPERTY.html */
func TestUnprocessedDocuments(t *testing.T) {
	_, err := database.UnprocessedDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY.html */
func TestViews(t *testing.T) {
	_, err := database.Views()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACT_METHOD.html */
func TestCompact(t *testing.T) {
	_, err := database.Compact()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPACTWITHOPTIONS_METHOD_DB.html */
// func TestCompactWithOptions(t *testing.T) {
// 	_, err := database.CompactWithOptions( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATE_METHOD.html */
func TestCreate(t *testing.T) {
	err := database.Create("", "TestDB.nsf", false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOPY_METHOD.html */
// func TestCreateCopy(t *testing.T) {
// 	_, err := database.CreateCopy( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENT_METHOD.html */
func TestCreateDocument(t *testing.T) {
	_, err := database.CreateDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOCUMENTCOLLECTION_METHOD.html */
func TestCreateDocumentCollection(t *testing.T) {
	_, err := database.CreateDocumentCollection()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOMINOQUERY_METHOD.html */
func TestCreateDominoQuery(t *testing.T) {
	_, err := database.CreateDominoQuery()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFROMTEMPLATE_METHOD.html */
// func TestCreateFromTemplate(t *testing.T) {
// 	_, err := database.CreateFromTemplate( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_METHOD_DB.html */
// func TestCreateFTIndex(t *testing.T) {
// 	err := database.CreateFTIndex( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENOTECOLLECTION_METHOD_DATABASE.html */
func TestCreateNoteCollection(t *testing.T) {
	_, err := database.CreateNoteCollection(false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEOUTLINE_METHOD_DATABASE.html */
func TestCreateOutline(t *testing.T) {
	_, err := database.CreateOutline("new-outline")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD.html */
// func TestCreateReplica(t *testing.T) {
// 	_, err := database.CreateReplica( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEW_METHOD_DB.html */
func TestCreateView(t *testing.T) {
	_, err := database.CreateView(domigo.WithNotesDatabaseCreateViewViewName("test-view"))
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECRYPT_METHOD_DATABASE.html */
func TestDecrypt(t *testing.T) {
	err := database.Decrypt(false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEFOLDER_METHOD_DATABASE.html */
// func TestEnableFolder(t *testing.T) {
// 	err := database.EnableFolder( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIXUP_METHOD_DB.html */
func TestFixup(t *testing.T) {
	err := database.Fixup()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTDOMAINSEARCH_METHOD_9515.html */
// func TestFTDomainSearch(t *testing.T) {
// 	_, err := database.FTDomainSearch( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_DB.html */
// func TestFTSearch(t *testing.T) {
// 	_, err := database.FTSearch( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHRANGE_METHOD_DB.html */
// func TestFTSearchRange(t *testing.T) {
// 	_, err := database.FTSearchRange( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETAGENT_METHOD.html */
func TestGetAgent(t *testing.T) {
	_, err := database.GetAgent("TestAgent")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLREADDOCUMENTS_DATABASE.html */
func TestGetAllReadDocuments(t *testing.T) {
	_, err := database.GetAllReadDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLUNREADDOCUMENTS_DATABASE.html */
func TestGetAllUnreadDocuments(t *testing.T) {
	_, err := database.GetAllUnreadDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYID_METHOD.html */
// func TestGetDocumentByID(t *testing.T) {
// 	_, err := database.GetDocumentByID( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYUNID_METHOD.html */
// func TestGetDocumentByUNID(t *testing.T) {
// 	_, err := database.GetDocumentByUNID( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYURL_METHOD.html */
// func TestGetDocumentByURL(t *testing.T) {
// 	_, err := database.GetDocumentByURL( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFORM_METHOD.html */
func TestGetForm(t *testing.T) {
	_, err := database.GetForm("TestForm")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTS_METHOD_DB.html */
// func TestGetModifiedDocuments(t *testing.T) {
// 	_, err := database.GetModifiedDocuments( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMODIFIEDDOCUMENTWITHOPTIONS_METHOD_DATABASE.html */
// func TestGetModifiedDocumentsWithOptions(t *testing.T) {
// 	_, err := database.GetModifiedDocumentsWithOptions( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENT_METHOD.html */
func TestGetNamedDocument(t *testing.T) {
	err := database.GetNamedDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNAMEDDOCUMENTCOLLECTION_METHOD.html */
// func TestGetNamedDocCollection(t *testing.T) {
// 	err := database.GetNamedDocCollection( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOPTION_METHOD_DB.html */
// func TestGetOption(t *testing.T) {
// 	_, err := database.GetOption( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOUTLINE_METHOD_DATABASE.html */
// func TestGetOutline(t *testing.T) {
// 	_, err := database.GetOutline( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROFILEDOCCOLLECTION_METHOD_DATABASE.html */
// func TestGetProfileDocCollection(t *testing.T) {
// 	_, err := database.GetProfileDocCollection( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROFILEDOCUMENT_METHOD.html */
// func TestGetProfileDocument(t *testing.T) {
// 	_, err := database.GetProfileDocument( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETQUERYRESULTSPROCESSOR_METHOD.html */
func TestGetQueryResultsProcessor(t *testing.T) {
	_, err := database.GetQueryResultsProcessor()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETURLHEADERINFO_METHOD.html */
// func TestGetURLHeaderInfo(t *testing.T) {
// 	_, err := database.GetURLHeaderInfo( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVIEW_METHOD.html */
func TestGetView(t *testing.T) {
	_, err := database.GetView("TestView")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GRANTACCESS_METHOD.html */
// func TestGrantAccess(t *testing.T) {
// 	err := database.GrantAccess( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKFORDELETE_METHOD_DB.html */
// func TestMarkForDelete(t *testing.T) {
// 	err := database.MarkForDelete() /* TODO: How to test this function? */
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD.html */
// func TestOpen(t *testing.T) {
// 	_, err := database.Open( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENBYREPLICAID_METHOD.html */
// func TestOpenByReplicaID(t *testing.T) {
// 	_, err := database.OpenByReplicaID( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENIFMODIFIED_METHOD.html */
// func TestOpenIfModified(t *testing.T) {
// 	_, err := database.OpenIfModified( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAIL_METHOD.html */
func TestOpenMail(t *testing.T) {
	err := database.OpenMail()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENURLDB_METHOD.html */
func TestOpenURLDb(t *testing.T) {
	_, err := database.OpenURLDb()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENWITHFAILOVER_METHOD.html */
// func TestOpenWithFailover(t *testing.T) {
// 	_, err := database.OpenWithFailover( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESS_METHOD.html */
// func TestQueryAccess(t *testing.T) {
// 	_, err := database.QueryAccess( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSPRIVILEGES_METHOD_DB.html */
// func TestQueryAccessPrivileges(t *testing.T) {
// 	_, err := database.QueryAccessPrivileges( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERYACCESSROLES_METHOD_DB.html */
// func TestQueryAccessRoles(t *testing.T) {
// 	_, err := database.QueryAccessRoles( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DB.html */
// func TestRemove(t *testing.T) {
// 	err := database.Remove() /* TODO: Test later. */
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFTINDEX_METHOD_DB.html */
func TestRemoveFTIndex(t *testing.T) {
	err := database.RemoveFTIndex()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICATE_METHOD.html */
// func TestReplicate(t *testing.T) {
// 	_, err := database.Replicate( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REVOKEACCESS_METHOD.html */
// func TestRevokeAccess(t *testing.T) {
// 	err := database.RevokeAccess( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCH_METHOD.html */
// func TestSearch(t *testing.T) {
// 	_, err := database.Search( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETOPTION_METHOD_DB.html */
// func TestSetOption(t *testing.T) {
// 	err := database.SetOption( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERIDFORDECRYPT_DATABASE.html */
// func TestSetUserIDForDecrypt(t *testing.T) {
// 	err := database.SetUserIDForDecrypt()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD_DB.html */
// func TestSign(t *testing.T) {
// 	err := database.Sign( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONBEGIN_METHOD.html */
func TestTransactionBegin(t *testing.T) {
	err := database.TransactionBegin()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONCOMMIT_METHOD.html */
func TestTransactionCommit(t *testing.T) {
	err := database.TransactionCommit()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRANSACTIONROLLBACK_METHOD.html */
func TestTransactionRollback(t *testing.T) {
	err := database.TransactionRollback()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDFTSEARCH_METHOD.html */
// func TestUnprocessedFTSearch(t *testing.T) {
// 	_, err := database.UnprocessedFTSearch( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDFTSEARCHRANGE_METHOD_DB.html */
// func TestUnprocessedFTSearchRange(t *testing.T) {
// 	_, err := database.UnprocessedFTSearchRange( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNPROCESSEDSEARCH_METHOD.html */
// func TestUnprocessedSearch(t *testing.T) {
// 	_, err := database.UnprocessedSearch( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEFTINDEX_METHOD.html */
func TestUpdateFTIndex(t *testing.T) {
	err := database.UpdateFTIndex(false)
	require.Nil(t, err)
}
