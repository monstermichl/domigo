/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENT_CLASS.html */
package domigo

import (
	"fmt"
	"time"

	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesDocument struct {
	NotesStruct
}

func NewNotesDocument(dispatchPtr *ole.IDispatch) NotesDocument {
	return NotesDocument{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTHORS_PROPERTY.html */
func (d NotesDocument) Authors() ([]String, error) {
	return getComArrayProperty[String](d, "Authors")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY.html */
/* TODO: Wrap values into struct. */
func (d NotesDocument) ColumnValues() ([]any, error) {
	return getComArrayProperty(d, "ColumnValues")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DOC.html */
func (d NotesDocument) Created() (time.Time, error) {
	return getComProperty(d, helpers, "Created")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EMBEDDEDOBJECTS_PROPERTY_DOC.html */
func (d NotesDocument) EmbeddedObjects() ([]NotesEmbeddedObject, error) {
	return com.GetObjectArrayProperty(d.com(), NewNotesEmbeddedObject, "EmbeddedObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func (d NotesDocument) EncryptionKeys() ([]String, error) {
	return getComArrayProperty[String](d, "EncryptionKeys")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func (d NotesDocument) SetEncryptionKeys(v []String) error {
	return putComProperty(d, "EncryptionKeys", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func (d NotesDocument) EncryptOnSend() (Boolean, error) {
	return getComProperty[Boolean](d, "EncryptOnSend")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func (d NotesDocument) SetEncryptOnSend(v Boolean) error {
	return putComProperty(d, "EncryptOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCES_PROPERTY_1349_ABOUT.html */
func (d NotesDocument) FolderReferences() ([]String, error) {
	return getComArrayProperty[String](d, "FolderReferences")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY.html */
func (d NotesDocument) FTSearchScore() (Long, error) {
	return getComProperty[Long](d, "FTSearchScore")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASEMBEDDED_PROPERTY.html */
func (d NotesDocument) HasEmbedded() (Boolean, error) {
	return getComProperty[Boolean](d, "HasEmbedded")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DOCUMENT_COM.html */
func (d NotesDocument) HttpURL() (String, error) {
	return getComProperty[String](d, "HttpURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDELETED_PROPERTY_8893_ABOUT.html */
func (d NotesDocument) IsDeleted() (Boolean, error) {
	return getComProperty[Boolean](d, "IsDeleted")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY_DOC.html */
func (d NotesDocument) IsEncrypted() (Boolean, error) {
	return getComProperty[Boolean](d, "IsEncrypted")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMEDDOC_PROPERTY.html */
func (d NotesDocument) IsNamedDoc() (Boolean, error) {
	return getComProperty[Boolean](d, "IsNamedDoc")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNEWNOTE_PROPERTY.html */
func (d NotesDocument) IsNewNote() (Boolean, error) {
	return getComProperty[Boolean](d, "IsNewNote")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROFILE_PROPERTY.html */
func (d NotesDocument) IsProfile() (Boolean, error) {
	return getComProperty[Boolean](d, "IsProfile")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESPONSE_PROPERTY_DOC.html */
func (d NotesDocument) IsResponse() (Boolean, error) {
	return getComProperty[Boolean](d, "IsResponse")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_DOC.html */
func (d NotesDocument) IsSigned() (Boolean, error) {
	return getComProperty[Boolean](d, "IsSigned")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISUIDOCOPEN_PROPERTY.html */
func (d NotesDocument) IsUIDocOpen() (Boolean, error) {
	return getComProperty[Boolean](d, "IsUIDocOpen")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_DOC.html */
func (d NotesDocument) IsValid() (Boolean, error) {
	return getComProperty[Boolean](d, "IsValid")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITEMS_PROPERTY.html */
func (d NotesDocument) Items() ([]NotesItem, error) {
	return com.GetObjectArrayProperty(d.com(), NewNotesItem, "Items")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEY_PROPERTY.html */
func (d NotesDocument) Key() (String, error) {
	return getComProperty[String](d, "Key")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTACCESSED_PROPERTY.html */
func (d NotesDocument) LastAccessed() (Time, error) {
	return getComProperty[Time](d, "LastAccessed")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DOC.html */
func (d NotesDocument) LastModified() (Time, error) {
	return getComProperty[Time](d, "LastModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_DOC.html */
func (d NotesDocument) LockHolders() ([]String, error) {
	return getComArrayProperty[String](d, "LockHolders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFDOC_PROPERTY.html */
func (d NotesDocument) NameOfDoc() (String, error) {
	return getComProperty[String](d, "NameOfDoc")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFPROFILE_PROPERTY.html */
func (d NotesDocument) NameOfProfile() (String, error) {
	return getComProperty[String](d, "NameOfProfile")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY.html */
func (d NotesDocument) NoteID() (String, error) {
	return getComProperty[String](d, "NoteID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DOCUMENT_COM.html */
func (d NotesDocument) NotesURL() (String, error) {
	return getComProperty[String](d, "NotesURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDATABASE_PROPERTY.html */
func (d NotesDocument) ParentDatabase() (NotesDatabase, error) {
	return getComObjectProperty(d, NewNotesDatabase, "ParentDatabase")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDOCUMENTUNID_PROPERTY.html */
func (d NotesDocument) ParentDocumentUNID() (String, error) {
	return getComProperty[String](d, "ParentDocumentUNID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY.html */
func (d NotesDocument) ParentView() (NotesView, error) {
	return getComObjectProperty(d, NewNotesView, "ParentView")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESPONSES_PROPERTY.html */
func (d NotesDocument) Responses() (NotesDocumentCollection, error) {
	return getComObjectProperty(d, NewNotesDocumentCollection, "Responses")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func (d NotesDocument) SaveMessageOnSend() (Boolean, error) {
	return getComProperty[Boolean](d, "SaveMessageOnSend")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func (d NotesDocument) SetSaveMessageOnSend(v Boolean) error {
	return putComProperty(d, "SaveMessageOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENTBYAGENT_PROPERTY.html */
func (d NotesDocument) SentByAgent() (Boolean, error) {
	return getComProperty[Boolean](d, "SentByAgent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNER_PROPERTY.html */
func (d NotesDocument) Signer() (String, error) {
	return getComProperty[String](d, "Signer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func (d NotesDocument) SignOnSend() (Boolean, error) {
	return getComProperty[Boolean](d, "SignOnSend")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func (d NotesDocument) SetSignOnSend(v Boolean) error {
	return putComProperty(d, "SignOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOFDOC_PROPERTY.html */
func (d NotesDocument) UserNameOfDoc() (String, error) {
	return getComProperty[String](d, "UserNameOfDoc")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DOC.html */
func (d NotesDocument) Size() (Long, error) {
	return getComProperty[Long](d, "Size")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func (d NotesDocument) UniversalID() (String, error) {
	return getComProperty[String](d, "UniversalID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func (d NotesDocument) SetUniversalID(v String) error {
	return putComProperty(d, "UniversalID", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFIER_PROPERTY.html */
func (d NotesDocument) Verifier() (String, error) {
	return getComProperty[String](d, "Verifier")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDITEMVALUE_METHOD.html */
func (d NotesDocument) AppendItemValue(itemName String, value any) (NotesItem, error) {
	return callComObjectMethod(d, NewNotesItem, "AppendItemValue", itemName, value)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHVCARD_METHOD.html */
func (d NotesDocument) AttachVCard(clientADTObject NotesDocument) error {
	return callComVoidMethod(d, "AttachVCard", clientADTObject)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSEMIMEENTITIES_METHOD_DOC.html */
type notesDocumentCloseMIMEEntitiesParams struct {
	savechanges    *Boolean
	entityitemname *String
}

type notesDocumentCloseMIMEEntitiesParam func(*notesDocumentCloseMIMEEntitiesParams)

func WithNotesDocumentCloseMIMEEntitiesSavechanges(savechanges Boolean) notesDocumentCloseMIMEEntitiesParam {
	return func(c *notesDocumentCloseMIMEEntitiesParams) {
		c.savechanges = &savechanges
	}
}

func WithNotesDocumentCloseMIMEEntitiesEntityitemname(entityitemname String) notesDocumentCloseMIMEEntitiesParam {
	return func(c *notesDocumentCloseMIMEEntitiesParams) {
		c.entityitemname = &entityitemname
	}
}

func (d NotesDocument) CloseMIMEEntities(params ...notesDocumentCloseMIMEEntitiesParam) (Boolean, error) {
	paramsStruct := &notesDocumentCloseMIMEEntitiesParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.savechanges != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.savechanges)
		if paramsStruct.entityitemname != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.entityitemname)
		}
	}
	return callComMethod[Boolean](d, "CloseMIMEEntities", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPUTEWITHFORM_METHOD.html */
func (d NotesDocument) ComputeWithForm(doDataTypes Boolean, raiseError Boolean) (Boolean, error) {
	return callComMethod[Boolean](d, "ComputeWithForm", doDataTypes, raiseError)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTTOMIME_METHOD.html */
type notesDocumentConvertToMIMEParams struct {
	conversionType *Integer
	options        *Long
}

type notesDocumentConvertToMIMEParam func(*notesDocumentConvertToMIMEParams)

func WithNotesDocumentConvertToMIMEConversionType(conversionType Integer) notesDocumentConvertToMIMEParam {
	return func(c *notesDocumentConvertToMIMEParams) {
		c.conversionType = &conversionType
	}
}

func WithNotesDocumentConvertToMIMEOptions(options Long) notesDocumentConvertToMIMEParam {
	return func(c *notesDocumentConvertToMIMEParams) {
		c.options = &options
	}
}

func (d NotesDocument) ConvertToMIME(params ...notesDocumentConvertToMIMEParam) error {
	paramsStruct := &notesDocumentConvertToMIMEParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.conversionType != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.conversionType)
		if paramsStruct.options != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.options)
		}
	}
	return callComVoidMethod(d, "ConvertToMIME", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYALLITEMS_METHOD.html */
type notesDocumentCopyAllItemsParams struct {
	replace *Boolean
}

type notesDocumentCopyAllItemsParam func(*notesDocumentCopyAllItemsParams)

func WithNotesDocumentCopyAllItemsReplace(replace Boolean) notesDocumentCopyAllItemsParam {
	return func(c *notesDocumentCopyAllItemsParams) {
		c.replace = &replace
	}
}

func (d NotesDocument) CopyAllItems(notesDocument NotesDocument, params ...notesDocumentCopyAllItemsParam) error {
	paramsStruct := &notesDocumentCopyAllItemsParams{}
	paramsOrdered := []interface{}{notesDocument.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.replace != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.replace)
	}
	return callComVoidMethod(d, "CopyAllItems", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEM_METHOD.html */
func (d NotesDocument) CopyItem(item NotesItem, newName String) (NotesItem, error) {
	return callComObjectMethod(d, NewNotesItem, "CopyItem", item.com().Dispatch(), newName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYTODATABASE_METHOD.html */
func (d NotesDocument) CopyToDatabase(notesDatabase NotesDatabase) (NotesDocument, error) {
	return callComObjectMethod(d, NewNotesDocument, "CopyToDatabase", notesDatabase.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMIMEENTITY_METHOD_DOC.html */
type notesDocumentCreateMIMEEntityParams struct {
	itemName *String
}

type notesDocumentCreateMIMEEntityParam func(*notesDocumentCreateMIMEEntityParams)

func WithNotesDocumentCreateMIMEEntityItemName(itemName String) notesDocumentCreateMIMEEntityParam {
	return func(c *notesDocumentCreateMIMEEntityParams) {
		c.itemName = &itemName
	}
}

func (d NotesDocument) CreateMIMEEntity(params ...notesDocumentCreateMIMEEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &notesDocumentCreateMIMEEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.itemName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.itemName)
	}
	return callComObjectMethod(d, NewNotesMIMEEntity, "CreateMIMEEntity", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLYMESSAGE_METHOD.html */
func (d NotesDocument) CreateReplyMessage(all Boolean) (NotesDocument, error) {
	return callComObjectMethod(d, NewNotesDocument, "CreateReplyMessage", all)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTITEM_METHOD.html */
func (d NotesDocument) CreateRichTextItem(name String) (NotesRichTextItem, error) {
	return callComObjectMethod(d, NewNotesRichTextItem, "CreateRichTextItem", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD.html */
func (d NotesDocument) Encrypt() error {
	return callComVoidMethod(d, "Encrypt")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETATTACHMENT_METHOD.html */
func (d NotesDocument) GetAttachment(fileName String) (NotesEmbeddedObject, error) {
	return callComObjectMethod(d, NewNotesEmbeddedObject, "GetAttachment", fileName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEM_METHOD.html */
func (d NotesDocument) GetFirstItem(name String) (NotesItem, error) {
	return callComObjectMethod(d, NewNotesItem, "GetFirstItem", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUE_METHOD.html */
func (d NotesDocument) GetItemValue(itemName String) (any, error) {
	return callComMethod(d, "GetItemValue", itemName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
func (d NotesDocument) GetItemValueCustomDataBytes(itemName String, dataTypeName String) ([]Byte, error) {
	return callComArrayMethod[Byte](d, "GetItemValueCustomDataBytes", itemName, dataTypeName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUEDATETIMEARRAY_METHOD.html */
/* TODO: Find out how to handle different return value types. */
func (d NotesDocument) GetItemValueDateTimeArray(itemName String) ([]NotesDateTime, error) {
	dispatchPtrs, err := callComObjectArrayMethod(d, "GetItemValueDateTimeArray", itemName)
	// vals := []any{}

	if err != nil {
		return []NotesDateTime{}, err
	}

	for _, ptr := range dispatchPtrs {
		fmt.Println(ptr)
	}
	return com.CallObjectArrayMethod(d.com(), NewNotesDateTime, "GetItemValueDateTimeArray")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_DOC.html */
type notesDocumentGetMIMEEntityParams struct {
	itemName *String
}

type notesDocumentGetMIMEEntityParam func(*notesDocumentGetMIMEEntityParams)

func WithNotesDocumentGetMIMEEntityItemName(itemName String) notesDocumentGetMIMEEntityParam {
	return func(c *notesDocumentGetMIMEEntityParams) {
		c.itemName = &itemName
	}
}

func (d NotesDocument) GetMIMEEntity(params ...notesDocumentGetMIMEEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &notesDocumentGetMIMEEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.itemName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.itemName)
	}
	return callComObjectMethod(d, NewNotesMIMEEntity, "GetMIMEEntity", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETREAD_METHOD_DOC.html */
type notesDocumentGetReadParams struct {
	username *String
}

type notesDocumentGetReadParam func(*notesDocumentGetReadParams)

func WithNotesDocumentGetReadUsername(username String) notesDocumentGetReadParam {
	return func(c *notesDocumentGetReadParams) {
		c.username = &username
	}
}

func (d NotesDocument) GetRead(params ...notesDocumentGetReadParam) (Boolean, error) {
	paramsStruct := &notesDocumentGetReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	return callComMethod[Boolean](d, "GetRead", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETRECEIVEDITEMTEXT_METHOD_DOC.html */
func (d NotesDocument) GetReceivedItemText() (String, error) {
	return callComMethod[String](d, "GetReceivedItemText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASITEM_METHOD.html */
func (d NotesDocument) HasItem(itemName String) (Boolean, error) {
	return callComMethod[Boolean](d, "HasItem", itemName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_DOC.html */
type notesDocumentLockParams struct {
	name          *[]String
	provisionalOK *Boolean
}

type notesDocumentLockParam func(*notesDocumentLockParams)

func WithNotesDocumentLockName(name []String) notesDocumentLockParam {
	return func(c *notesDocumentLockParams) {
		c.name = &name
	}
}

func WithNotesDocumentLockProvisionalOK(provisionalOK Boolean) notesDocumentLockParam {
	return func(c *notesDocumentLockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (d NotesDocument) Lock(params ...notesDocumentLockParam) (Boolean, error) {
	paramsStruct := &notesDocumentLockParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
		if paramsStruct.provisionalOK != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.provisionalOK)
		}
	}
	return callComMethod[Boolean](d, "Lock", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_DOC.html */
type notesDocumentLockProvisionalParams struct {
	name *[]String
}

type notesDocumentLockProvisionalParam func(*notesDocumentLockProvisionalParams)

func WithNotesDocumentLockProvisionalName(name []String) notesDocumentLockProvisionalParam {
	return func(c *notesDocumentLockProvisionalParams) {
		c.name = &name
	}
}

func (d NotesDocument) LockProvisional(params ...notesDocumentLockProvisionalParam) (Boolean, error) {
	paramsStruct := &notesDocumentLockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	return callComMethod[Boolean](d, "LockProvisional", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAKERESPONSE_METHOD.html */
func (d NotesDocument) MakeResponse(document NotesDocument) error {
	return callComVoidMethod(d, "MakeResponse", document.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKREAD_DOCUMENT.html */
type notesDocumentMarkReadParams struct {
	username *String
}

type notesDocumentMarkReadParam func(*notesDocumentMarkReadParams)

func WithNotesDocumentMarkReadUsername(username String) notesDocumentMarkReadParam {
	return func(c *notesDocumentMarkReadParams) {
		c.username = &username
	}
}

func (d NotesDocument) MarkRead(params ...notesDocumentMarkReadParam) error {
	paramsStruct := &notesDocumentMarkReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	return callComVoidMethod(d, "MarkRead", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKUNREAD_DOCUMENT.html */
type notesDocumentMarkUnreadParams struct {
	username *String
}

type notesDocumentMarkUnreadParam func(*notesDocumentMarkUnreadParams)

func WithNotesDocumentMarkUnreadUsername(username String) notesDocumentMarkUnreadParam {
	return func(c *notesDocumentMarkUnreadParams) {
		c.username = &username
	}
}

func (d NotesDocument) MarkUnread(params ...notesDocumentMarkUnreadParam) error {
	paramsStruct := &notesDocumentMarkUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	return callComVoidMethod(d, "MarkUnread", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTINFOLDER_METHOD.html */
type notesDocumentPutInFolderParams struct {
	createonfail *Boolean
}

type notesDocumentPutInFolderParam func(*notesDocumentPutInFolderParams)

func WithNotesDocumentPutInFolderCreateonfail(createonfail Boolean) notesDocumentPutInFolderParam {
	return func(c *notesDocumentPutInFolderParams) {
		c.createonfail = &createonfail
	}
}

func (d NotesDocument) PutInFolder(folderName String, params ...notesDocumentPutInFolderParam) error {
	paramsStruct := &notesDocumentPutInFolderParams{}
	paramsOrdered := []interface{}{folderName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	return callComVoidMethod(d, "PutInFolder", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DOC.html */
func (d NotesDocument) Remove(force Boolean) (Boolean, error) {
	return callComMethod[Boolean](d, "Remove", force)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFROMFOLDER_METHOD.html */
func (d NotesDocument) RemoveFromFolder(folderName String) error {
	return callComVoidMethod(d, "RemoveFromFolder", folderName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEITEM_METHOD.html */
func (d NotesDocument) RemoveItem(itemName String) error {
	return callComVoidMethod(d, "RemoveItem", itemName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEPERMANENTLY_METHOD_DOC.html */
func (d NotesDocument) RemovePermanently(force Boolean) (Boolean, error) {
	return callComMethod[Boolean](d, "RemovePermanently", force)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENDERTORTITEM_METHOD.html */
func (d NotesDocument) RenderToRTItem(notesRichTextItem NotesRichTextItem) (Boolean, error) {
	return callComMethod[Boolean](d, "RenderToRTItem", notesRichTextItem.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUE_METHOD.html */
func (d NotesDocument) ReplaceItemValue(itemName String, value any) (NotesItem, error) {
	s, err := GetNotesStruct(value)

	if err == nil {
		fmt.Println("casted", s)
		value = s.com().Dispatch()
	}
	return callComObjectMethod(d, NewNotesItem, "ReplaceItemValue", itemName, value)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
func (d NotesDocument) ReplaceItemValueCustomDataBytes(itemName String, dataTypeName String, byteArray []Byte) error {
	return callComVoidMethod(d, "ReplaceItemValueCustomDataBytes", itemName, dataTypeName, byteArray)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_DOC.html */
type notesDocumentSaveParams struct {
	markRead *Boolean
}

type notesDocumentSaveParam func(*notesDocumentSaveParams)

func WithNotesDocumentSaveMarkRead(markRead Boolean) notesDocumentSaveParam {
	return func(c *notesDocumentSaveParams) {
		c.markRead = &markRead
	}
}

func (d NotesDocument) Save(force Boolean, createResponse Boolean, params ...notesDocumentSaveParam) (Boolean, error) {
	paramsStruct := &notesDocumentSaveParams{}
	paramsOrdered := []interface{}{force, createResponse}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.markRead != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.markRead)
	}
	return callComMethod[Boolean](d, "Save", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEND_METHOD_DOC.html */
type notesDocumentSendParams struct {
	recipients *[]String
}

type notesDocumentSendParam func(*notesDocumentSendParams)

func WithNotesDocumentSendRecipients(recipients []String) notesDocumentSendParam {
	return func(c *notesDocumentSendParams) {
		c.recipients = &recipients
	}
}

func (d NotesDocument) Send(attachForm Boolean, params ...notesDocumentSendParam) error {
	paramsStruct := &notesDocumentSendParams{}
	paramsOrdered := []interface{}{attachForm}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.recipients != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.recipients)
	}
	return callComVoidMethod(d, "Send", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD.html */
func (d NotesDocument) Sign() error {
	return callComVoidMethod(d, "Sign")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_DOC.html */
func (d NotesDocument) UnLock() error {
	return callComVoidMethod(d, "UnLock")
}
