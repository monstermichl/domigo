package notesdocument_test

import (
	"domigo/domino/notesdatabase"
	"domigo/domino/notesdocument"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/require"
)

const TEST_ITEM_NAME = "test_item"

var db notesdatabase.NotesDatabase
var document notesdocument.NotesDocument

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	db, _ = session.GetDatabase("", "GoInterface.nsf")
	document, _ = db.CreateDocument()

	m.Run()
	session.Release()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTHORS_PROPERTY.html */
func TestAuthors(t *testing.T) {
	_, err := document.Authors()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY.html */
func TestColumnValues(t *testing.T) {
	_, err := document.ColumnValues()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DOC.html */
func TestCreated(t *testing.T) {
	_, err := document.Created()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EMBEDDEDOBJECTS_PROPERTY_DOC.html */
func TestEmbeddedObjects(t *testing.T) {
	_, err := document.EmbeddedObjects()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func TestEncryptionKeys(t *testing.T) {
	_, err := document.EncryptionKeys()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func TestSetEncryptionKeys(t *testing.T) {
	s, err := document.EncryptionKeys()
	require.Nil(t, err)

	err = document.SetEncryptionKeys(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func TestEncryptOnSend(t *testing.T) {
	_, err := document.EncryptOnSend()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func TestSetEncryptOnSend(t *testing.T) {
	s, err := document.EncryptOnSend()
	require.Nil(t, err)

	err = document.SetEncryptOnSend(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCES_PROPERTY_1349_ABOUT.html */
func TestFolderReferences(t *testing.T) {
	_, err := document.FolderReferences()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY.html */
func TestFTSearchScore(t *testing.T) {
	_, err := document.FTSearchScore()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASEMBEDDED_PROPERTY.html */
func TestHasEmbedded(t *testing.T) {
	_, err := document.HasEmbedded()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DOCUMENT_COM.html */
func TestHttpURL(t *testing.T) {
	_, err := document.HttpURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDELETED_PROPERTY_8893_ABOUT.html */
func TestIsDeleted(t *testing.T) {
	_, err := document.IsDeleted()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY_DOC.html */
func TestIsEncrypted(t *testing.T) {
	_, err := document.IsEncrypted()
	require.Nil(t, err)
}

/* TODO: Access type for IsNamedDoc could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMEDDOC_PROPERTY.html */
func TestIsNamedDoc(t *testing.T) {
	_, err := document.IsNamedDoc()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNEWNOTE_PROPERTY.html */
func TestIsNewNote(t *testing.T) {
	_, err := document.IsNewNote()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROFILE_PROPERTY.html */
func TestIsProfile(t *testing.T) {
	_, err := document.IsProfile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESPONSE_PROPERTY_DOC.html */
func TestIsResponse(t *testing.T) {
	_, err := document.IsResponse()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_DOC.html */
func TestIsSigned(t *testing.T) {
	_, err := document.IsSigned()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISUIDOCOPEN_PROPERTY.html */
func TestIsUIDocOpen(t *testing.T) {
	_, err := document.IsUIDocOpen()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_DOC.html */
func TestIsValid(t *testing.T) {
	_, err := document.IsValid()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITEMS_PROPERTY.html */
func TestItems(t *testing.T) {
	_, err := document.Items()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEY_PROPERTY.html */
func TestKey(t *testing.T) {
	_, err := document.Key()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTACCESSED_PROPERTY.html */
func TestLastAccessed(t *testing.T) {
	_, err := document.LastAccessed()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DOC.html */
func TestLastModified(t *testing.T) {
	_, err := document.LastModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_DOC.html */
func TestLockHolders(t *testing.T) {
	_, err := document.LockHolders()
	require.Nil(t, err)
}

/* TODO: Access type for NameOfDoc could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFDOC_PROPERTY.html */
func TestNameOfDoc(t *testing.T) {
	_, err := document.NameOfDoc()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFPROFILE_PROPERTY.html */
func TestNameOfProfile(t *testing.T) {
	_, err := document.NameOfProfile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY.html */
func TestNoteID(t *testing.T) {
	_, err := document.NoteID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DOCUMENT_COM.html */
func TestNotesURL(t *testing.T) {
	_, err := document.NotesURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDOCUMENTUNID_PROPERTY.html */
func TestParentDocumentUNID(t *testing.T) {
	_, err := document.ParentDocumentUNID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func TestSaveMessageOnSend(t *testing.T) {
	_, err := document.SaveMessageOnSend()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func TestSetSaveMessageOnSend(t *testing.T) {
	s, err := document.SaveMessageOnSend()
	require.Nil(t, err)

	err = document.SetSaveMessageOnSend(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENTBYAGENT_PROPERTY.html */
func TestSentByAgent(t *testing.T) {
	_, err := document.SentByAgent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNER_PROPERTY.html */
func TestSigner(t *testing.T) {
	_, err := document.Signer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func TestSignOnSend(t *testing.T) {
	_, err := document.SignOnSend()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func TestSetSignOnSend(t *testing.T) {
	s, err := document.SignOnSend()
	require.Nil(t, err)

	err = document.SetSignOnSend(s)
	require.Nil(t, err)
}

/* TODO: Access type for UserNameOfDoc could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOFDOC_PROPERTY.html */
func TestUserNameOfDoc(t *testing.T) {
	_, err := document.UserNameOfDoc()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DOC.html */
func TestSize(t *testing.T) {
	_, err := document.Size()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func TestUniversalID(t *testing.T) {
	_, err := document.UniversalID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func TestSetUniversalID(t *testing.T) {
	s, err := document.UniversalID()
	require.Nil(t, err)

	err = document.SetUniversalID(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFIER_PROPERTY.html */
func TestVerifier(t *testing.T) {
	_, err := document.Verifier()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDITEMVALUE_METHOD.html */
func TestAppendItemValue(t *testing.T) {
	_, err := document.AppendItemValue(TEST_ITEM_NAME, "appended value")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHVCARD_METHOD.html */
func TestAttachVCard(t *testing.T) {
	newDoc, _ := db.CreateDocument()
	err := document.AttachVCard(newDoc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSEMIMEENTITIES_METHOD_DOC.html */
func TestCloseMIMEEntities(t *testing.T) {
	_, err := document.CloseMIMEEntities()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPUTEWITHFORM_METHOD.html */
func TestComputeWithForm(t *testing.T) {
	_, err := document.ComputeWithForm(false, false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTTOMIME_METHOD.html */
func TestConvertToMIME(t *testing.T) {
	err := document.ConvertToMIME()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYALLITEMS_METHOD.html */
func TestCopyAllItems(t *testing.T) {
	targetDoc, _ := db.CreateDocument()
	err := document.CopyAllItems(targetDoc, true)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEM_METHOD.html */
// func TestCopyItem(t *testing.T) {
// 	_, err := document.CopyItem( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMIMEENTITY_METHOD_DOC.html */
// func TestCreateMIMEEntity(t *testing.T) {
// 	_, err := document.CreateMIMEEntity( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLYMESSAGE_METHOD.html */
// func TestCreateReplyMessage(t *testing.T) {
// 	_, err := document.CreateReplyMessage( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTITEM_METHOD.html */
// func TestCreateRichTextItem(t *testing.T) {
// 	_, err := document.CreateRichTextItem( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD.html */
// func TestEncrypt(t *testing.T) {
// 	err := document.Encrypt()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETATTACHMENT_METHOD.html */
// func TestGetAttachment(t *testing.T) {
// 	_, err := document.GetAttachment( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEM_METHOD.html */
// func TestGetFirstItem(t *testing.T) {
// 	_, err := document.GetFirstItem( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUE_METHOD.html */
// func TestGetItemValue(t *testing.T) {
// 	_, err := document.GetItemValue( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
// func TestGetItemValueCustomDataBytes(t *testing.T) {
// 	_, err := document.GetItemValueCustomDataBytes( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUEDATETIMEARRAY_METHOD.html */
// func TestGetItemValueDateTimeArray(t *testing.T) {
// 	_, err := document.GetItemValueDateTimeArray( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_DOC.html */
// func TestGetMIMEEntity(t *testing.T) {
// 	_, err := document.GetMIMEEntity( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETREAD_METHOD_DOC.html */
// func TestGetRead(t *testing.T) {
// 	_, err := document.GetRead( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETRECEIVEDITEMTEXT_METHOD_DOC.html */
// func TestGetReceivedItemText(t *testing.T) {
// 	_, err := document.GetReceivedItemText()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASITEM_METHOD.html */
// func TestHasItem(t *testing.T) {
// 	_, err := document.HasItem( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_DOC.html */
// func TestLock(t *testing.T) {
// 	_, err := document.Lock( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_DOC.html */
// func TestLockProvisional(t *testing.T) {
// 	_, err := document.LockProvisional( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAKERESPONSE_METHOD.html */
// func TestMakeResponse(t *testing.T) {
// 	err := document.MakeResponse( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKREAD_DOCUMENT.html */
// func TestMarkRead(t *testing.T) {
// 	err := document.MarkRead( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKUNREAD_DOCUMENT.html */
// func TestMarkUnread(t *testing.T) {
// 	err := document.MarkUnread( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTINFOLDER_METHOD.html */
// func TestPutInFolder(t *testing.T) {
// 	err := document.PutInFolder( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DOC.html */
// func TestRemove(t *testing.T) {
// 	_, err := document.Remove( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFROMFOLDER_METHOD.html */
// func TestRemoveFromFolder(t *testing.T) {
// 	err := document.RemoveFromFolder( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEITEM_METHOD.html */
// func TestRemoveItem(t *testing.T) {
// 	err := document.RemoveItem( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEPERMANENTLY_METHOD_DOC.html */
// func TestRemovePermanently(t *testing.T) {
// 	_, err := document.RemovePermanently( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENDERTORTITEM_METHOD.html */
// func TestRenderToRTItem(t *testing.T) {
// 	_, err := document.RenderToRTItem( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUE_METHOD.html */
// func TestReplaceItemValue(t *testing.T) {
// 	_, err := document.ReplaceItemValue( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
// func TestReplaceItemValueCustomDataBytes(t *testing.T) {
// 	err := document.ReplaceItemValueCustomDataBytes( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_DOC.html */
// func TestSave(t *testing.T) {
// 	_, err := document.Save( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEND_METHOD_DOC.html */
// func TestSend(t *testing.T) {
// 	err := document.Send( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD.html */
func TestSign(t *testing.T) {
	err := document.Sign()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_DOC.html */
func TestUnLock(t *testing.T) {
	err := document.UnLock()
	require.Nil(t, err)
}
