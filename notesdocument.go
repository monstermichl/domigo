/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENT_CLASS.html */
package domigo

import (
	"fmt"
	"time"

	"github.com/monstermichl/domigo/internal/com"
	"github.com/monstermichl/domigo/internal/helpers"

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
	vals, err := getComArrayProperty(d, "Authors")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY.html */
/* TODO: Wrap values into struct. */
func (d NotesDocument) ColumnValues() ([]any, error) {
	return getComArrayProperty(d, "ColumnValues")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DOC.html */
func (d NotesDocument) Created() (time.Time, error) {
	val, err := getComProperty(d, "Created")
	return helpers.CastValue[time.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EMBEDDEDOBJECTS_PROPERTY_DOC.html */
func (d NotesDocument) EmbeddedObjects() ([]NotesEmbeddedObject, error) {
	return com.GetObjectArrayProperty(d.com(), NewNotesEmbeddedObject, "EmbeddedObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func (d NotesDocument) EncryptionKeys() ([]String, error) {
	vals, err := getComArrayProperty(d, "EncryptionKeys")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func (d NotesDocument) SetEncryptionKeys(v []String) error {
	return putComProperty(d, "EncryptionKeys", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func (d NotesDocument) EncryptOnSend() (Boolean, error) {
	val, err := getComProperty(d, "EncryptOnSend")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func (d NotesDocument) SetEncryptOnSend(v Boolean) error {
	return putComProperty(d, "EncryptOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCES_PROPERTY_1349_ABOUT.html */
func (d NotesDocument) FolderReferences() ([]String, error) {
	vals, err := getComArrayProperty(d, "FolderReferences")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY.html */
func (d NotesDocument) FTSearchScore() (Long, error) {
	val, err := getComProperty(d, "FTSearchScore")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASEMBEDDED_PROPERTY.html */
func (d NotesDocument) HasEmbedded() (Boolean, error) {
	val, err := getComProperty(d, "HasEmbedded")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DOCUMENT_COM.html */
func (d NotesDocument) HttpURL() (String, error) {
	val, err := getComProperty(d, "HttpURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDELETED_PROPERTY_8893_ABOUT.html */
func (d NotesDocument) IsDeleted() (Boolean, error) {
	val, err := getComProperty(d, "IsDeleted")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY_DOC.html */
func (d NotesDocument) IsEncrypted() (Boolean, error) {
	val, err := getComProperty(d, "IsEncrypted")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMEDDOC_PROPERTY.html */
func (d NotesDocument) IsNamedDoc() (Boolean, error) {
	val, err := getComProperty(d, "IsNamedDoc")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNEWNOTE_PROPERTY.html */
func (d NotesDocument) IsNewNote() (Boolean, error) {
	val, err := getComProperty(d, "IsNewNote")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROFILE_PROPERTY.html */
func (d NotesDocument) IsProfile() (Boolean, error) {
	val, err := getComProperty(d, "IsProfile")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESPONSE_PROPERTY_DOC.html */
func (d NotesDocument) IsResponse() (Boolean, error) {
	val, err := getComProperty(d, "IsResponse")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_DOC.html */
func (d NotesDocument) IsSigned() (Boolean, error) {
	val, err := getComProperty(d, "IsSigned")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISUIDOCOPEN_PROPERTY.html */
func (d NotesDocument) IsUIDocOpen() (Boolean, error) {
	val, err := getComProperty(d, "IsUIDocOpen")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_DOC.html */
func (d NotesDocument) IsValid() (Boolean, error) {
	val, err := getComProperty(d, "IsValid")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITEMS_PROPERTY.html */
func (d NotesDocument) Items() ([]NotesItem, error) {
	return com.GetObjectArrayProperty(d.com(), NewNotesItem, "Items")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEY_PROPERTY.html */
func (d NotesDocument) Key() (String, error) {
	val, err := getComProperty(d, "Key")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTACCESSED_PROPERTY.html */
func (d NotesDocument) LastAccessed() (Time, error) {
	val, err := getComProperty(d, "LastAccessed")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DOC.html */
func (d NotesDocument) LastModified() (Time, error) {
	val, err := getComProperty(d, "LastModified")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_DOC.html */
func (d NotesDocument) LockHolders() ([]String, error) {
	vals, err := getComArrayProperty(d, "LockHolders")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFDOC_PROPERTY.html */
func (d NotesDocument) NameOfDoc() (String, error) {
	val, err := getComProperty(d, "NameOfDoc")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFPROFILE_PROPERTY.html */
func (d NotesDocument) NameOfProfile() (String, error) {
	val, err := getComProperty(d, "NameOfProfile")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY.html */
func (d NotesDocument) NoteID() (String, error) {
	val, err := getComProperty(d, "NoteID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DOCUMENT_COM.html */
func (d NotesDocument) NotesURL() (String, error) {
	val, err := getComProperty(d, "NotesURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDATABASE_PROPERTY.html */
func (d NotesDocument) ParentDatabase() (NotesDatabase, error) {
	dispatchPtr, err := getComObjectProperty(d, "ParentDatabase")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDOCUMENTUNID_PROPERTY.html */
func (d NotesDocument) ParentDocumentUNID() (String, error) {
	val, err := getComProperty(d, "ParentDocumentUNID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY.html */
func (d NotesDocument) ParentView() (NotesView, error) {
	dispatchPtr, err := getComObjectProperty(d, "ParentView")
	return NewNotesView(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESPONSES_PROPERTY.html */
func (d NotesDocument) Responses() (NotesDocumentCollection, error) {
	dispatchPtr, err := getComObjectProperty(d, "Responses")
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func (d NotesDocument) SaveMessageOnSend() (Boolean, error) {
	val, err := getComProperty(d, "SaveMessageOnSend")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func (d NotesDocument) SetSaveMessageOnSend(v Boolean) error {
	return putComProperty(d, "SaveMessageOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENTBYAGENT_PROPERTY.html */
func (d NotesDocument) SentByAgent() (Boolean, error) {
	val, err := getComProperty(d, "SentByAgent")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNER_PROPERTY.html */
func (d NotesDocument) Signer() (String, error) {
	val, err := getComProperty(d, "Signer")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func (d NotesDocument) SignOnSend() (Boolean, error) {
	val, err := getComProperty(d, "SignOnSend")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func (d NotesDocument) SetSignOnSend(v Boolean) error {
	return putComProperty(d, "SignOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOFDOC_PROPERTY.html */
func (d NotesDocument) UserNameOfDoc() (String, error) {
	val, err := getComProperty(d, "UserNameOfDoc")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DOC.html */
func (d NotesDocument) Size() (Long, error) {
	val, err := getComProperty(d, "Size")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func (d NotesDocument) UniversalID() (String, error) {
	val, err := getComProperty(d, "UniversalID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func (d NotesDocument) SetUniversalID(v String) error {
	return putComProperty(d, "UniversalID", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFIER_PROPERTY.html */
func (d NotesDocument) Verifier() (String, error) {
	val, err := getComProperty(d, "Verifier")
	return helpers.CastValue[String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDITEMVALUE_METHOD.html */
func (d NotesDocument) AppendItemValue(itemName String, value any) (NotesItem, error) {
	dispatchPtr, err := callComObjectMethod(d, "AppendItemValue", itemName, value)
	return NewNotesItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHVCARD_METHOD.html */
func (d NotesDocument) AttachVCard(clientADTObject NotesDocument) error {
	_, err := callComMethod(d, "AttachVCard", clientADTObject)
	return err
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
	val, err := callComMethod(d, "CloseMIMEEntities", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPUTEWITHFORM_METHOD.html */
func (d NotesDocument) ComputeWithForm(doDataTypes Boolean, raiseError Boolean) (Boolean, error) {
	val, err := callComMethod(d, "ComputeWithForm", doDataTypes, raiseError)
	return helpers.CastValue[Boolean](val), err
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
	_, err := callComMethod(d, "ConvertToMIME", paramsOrdered...)
	return err
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
	_, err := callComMethod(d, "CopyAllItems", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEM_METHOD.html */
func (d NotesDocument) CopyItem(item NotesItem, newName String) (NotesItem, error) {
	dispatchPtr, err := callComObjectMethod(d, "CopyItem", item.com().Dispatch(), newName)
	return NewNotesItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYTODATABASE_METHOD.html */
func (d NotesDocument) CopyToDatabase(notesDatabase NotesDatabase) (NotesDocument, error) {
	dispatchPtr, err := callComObjectMethod(d, "CopyToDatabase", notesDatabase.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
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
	dispatchPtr, err := callComObjectMethod(d, "CreateMIMEEntity", paramsOrdered...)
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLYMESSAGE_METHOD.html */
func (d NotesDocument) CreateReplyMessage(all Boolean) (NotesDocument, error) {
	dispatchPtr, err := callComObjectMethod(d, "CreateReplyMessage", all)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTITEM_METHOD.html */
func (d NotesDocument) CreateRichTextItem(name String) (NotesRichTextItem, error) {
	dispatchPtr, err := callComObjectMethod(d, "CreateRichTextItem", name)
	return NewNotesRichTextItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD.html */
func (d NotesDocument) Encrypt() error {
	_, err := callComMethod(d, "Encrypt")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETATTACHMENT_METHOD.html */
func (d NotesDocument) GetAttachment(fileName String) (NotesEmbeddedObject, error) {
	dispatchPtr, err := callComObjectMethod(d, "GetAttachment", fileName)
	return NewNotesEmbeddedObject(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEM_METHOD.html */
func (d NotesDocument) GetFirstItem(name String) (NotesItem, error) {
	dispatchPtr, err := callComObjectMethod(d, "GetFirstItem", name)
	return NewNotesItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUE_METHOD.html */
func (d NotesDocument) GetItemValue(itemName String) (any, error) {
	return callComMethod(d, "GetItemValue", itemName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
func (d NotesDocument) GetItemValueCustomDataBytes(itemName String, dataTypeName String) ([]Byte, error) {
	vals, err := callComArrayMethod(d, "GetItemValueCustomDataBytes", itemName, dataTypeName)
	return helpers.CastSlice[Byte](vals), err
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
	dispatchPtr, err := callComObjectMethod(d, "GetMIMEEntity", paramsOrdered...)
	return NewNotesMIMEEntity(dispatchPtr), err
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
	val, err := callComMethod(d, "GetRead", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETRECEIVEDITEMTEXT_METHOD_DOC.html */
func (d NotesDocument) GetReceivedItemText() (String, error) {
	val, err := callComMethod(d, "GetReceivedItemText")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASITEM_METHOD.html */
func (d NotesDocument) HasItem(itemName String) (Boolean, error) {
	val, err := callComMethod(d, "HasItem", itemName)
	return helpers.CastValue[Boolean](val), err
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
	val, err := callComMethod(d, "Lock", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
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
	val, err := callComMethod(d, "LockProvisional", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAKERESPONSE_METHOD.html */
func (d NotesDocument) MakeResponse(document NotesDocument) error {
	_, err := callComMethod(d, "MakeResponse", document.com().Dispatch())
	return err
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
	_, err := callComMethod(d, "MarkRead", paramsOrdered...)
	return err
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
	_, err := callComMethod(d, "MarkUnread", paramsOrdered...)
	return err
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
	_, err := callComMethod(d, "PutInFolder", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DOC.html */
func (d NotesDocument) Remove(force Boolean) (Boolean, error) {
	val, err := callComMethod(d, "Remove", force)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFROMFOLDER_METHOD.html */
func (d NotesDocument) RemoveFromFolder(folderName String) error {
	_, err := callComMethod(d, "RemoveFromFolder", folderName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEITEM_METHOD.html */
func (d NotesDocument) RemoveItem(itemName String) error {
	_, err := callComMethod(d, "RemoveItem", itemName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEPERMANENTLY_METHOD_DOC.html */
func (d NotesDocument) RemovePermanently(force Boolean) (Boolean, error) {
	val, err := callComMethod(d, "RemovePermanently", force)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENDERTORTITEM_METHOD.html */
func (d NotesDocument) RenderToRTItem(notesRichTextItem NotesRichTextItem) (Boolean, error) {
	val, err := callComMethod(d, "RenderToRTItem", notesRichTextItem.com().Dispatch())
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUE_METHOD.html */
func (d NotesDocument) ReplaceItemValue(itemName String, value any) (NotesItem, error) {
	s, err := GetNotesStruct(value)

	if err == nil {
		fmt.Println("casted", s)
		value = s.com().Dispatch()
	}
	dispatchPtr, err := callComObjectMethod(d, "ReplaceItemValue", itemName, value)
	return NewNotesItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
func (d NotesDocument) ReplaceItemValueCustomDataBytes(itemName String, dataTypeName String, byteArray []Byte) error {
	_, err := callComMethod(d, "ReplaceItemValueCustomDataBytes", itemName, dataTypeName, byteArray)
	return err
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
	val, err := callComMethod(d, "Save", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
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
	_, err := callComMethod(d, "Send", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD.html */
func (d NotesDocument) Sign() error {
	_, err := callComMethod(d, "Sign")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_DOC.html */
func (d NotesDocument) UnLock() error {
	_, err := callComMethod(d, "UnLock")
	return err
}
