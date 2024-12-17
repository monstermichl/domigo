/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENT_CLASS.html */
package notesdocument

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesdatetime"
	"domigo/domino/notesembeddedobject"
	"domigo/domino/notesitem"
	"domigo/domino/notesmimeentity"
	"domigo/domino/notesrichtextitem"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDocument struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDocument {
	return NotesDocument{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ITEM.html */
/* Moved from NotesItem. */
func NotesItemParent(i notesitem.NotesItem) (NotesDocument, error) {
	dispatchPtr, err := i.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEMTODOCUMENT_METHOD.html */
/* Moved from NotesItem. */
func NotesItemCopyItemToDocument(i notesitem.NotesItem, document NotesDocument, newName domino.String) (notesitem.NotesItem, error) {
	dispatchPtr, err := i.Com().CallObjectMethod("CopyItemToDocument", document.Com().Dispatch(), newName)
	return notesitem.New(dispatchPtr), err
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

func NotesRichTextItemAppendDocLink(r notesrichtextitem.NotesRichTextItem, linkTo NotesDocument, comment domino.String, params ...appendDocLinkParam) error {
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

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTHORS_PROPERTY.html */
func (d NotesDocument) Authors() ([]domino.String, error) {
	vals, err := d.Com().GetArrayProperty("Authors")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY.html */
func (d NotesDocument) ColumnValues() ([]domino.Variant, error) {
	vals, err := d.Com().GetArrayProperty("ColumnValues")
	return helpers.CastSlice[domino.Variant](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_DOC.html */
func (d NotesDocument) Created() (domino.Time, error) {
	val, err := d.Com().GetProperty("Created")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EMBEDDEDOBJECTS_PROPERTY_DOC.html */
func (d NotesDocument) EmbeddedObjects() ([]notesembeddedobject.NotesEmbeddedObject, error) {
	return com.GetObjectArrayProperty(d.Com(), notesembeddedobject.New, "EmbeddedObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func (d NotesDocument) EncryptionKeys() ([]domino.String, error) {
	vals, err := d.Com().GetArrayProperty("EncryptionKeys")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTIONKEYS_PROPERTY.html */
func (d NotesDocument) SetEncryptionKeys(v []domino.String) error {
	return d.Com().PutProperty("EncryptionKeys", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func (d NotesDocument) EncryptOnSend() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("EncryptOnSend")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPTONSEND_PROPERTY.html */
func (d NotesDocument) SetEncryptOnSend(v domino.Boolean) error {
	return d.Com().PutProperty("EncryptOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FOLDERREFERENCES_PROPERTY_1349_ABOUT.html */
func (d NotesDocument) FolderReferences() ([]domino.String, error) {
	vals, err := d.Com().GetArrayProperty("FolderReferences")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY.html */
func (d NotesDocument) FTSearchScore() (domino.Long, error) {
	val, err := d.Com().GetProperty("FTSearchScore")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASEMBEDDED_PROPERTY.html */
func (d NotesDocument) HasEmbedded() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("HasEmbedded")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_DOCUMENT_COM.html */
func (d NotesDocument) HttpURL() (domino.String, error) {
	val, err := d.Com().GetProperty("HttpURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDELETED_PROPERTY_8893_ABOUT.html */
func (d NotesDocument) IsDeleted() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsDeleted")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY_DOC.html */
func (d NotesDocument) IsEncrypted() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsEncrypted")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMEDDOC_PROPERTY.html */
func (d NotesDocument) IsNamedDoc() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsNamedDoc")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNEWNOTE_PROPERTY.html */
func (d NotesDocument) IsNewNote() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsNewNote")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROFILE_PROPERTY.html */
func (d NotesDocument) IsProfile() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsProfile")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESPONSE_PROPERTY_DOC.html */
func (d NotesDocument) IsResponse() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsResponse")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_DOC.html */
func (d NotesDocument) IsSigned() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsSigned")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISUIDOCOPEN_PROPERTY.html */
func (d NotesDocument) IsUIDocOpen() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsUIDocOpen")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_DOC.html */
func (d NotesDocument) IsValid() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsValid")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITEMS_PROPERTY.html */
func (d NotesDocument) Items() ([]notesitem.NotesItem, error) {
	return com.GetObjectArrayProperty(d.Com(), notesitem.New, "Items")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEY_PROPERTY.html */
func (d NotesDocument) Key() (domino.String, error) {
	val, err := d.Com().GetProperty("Key")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTACCESSED_PROPERTY.html */
func (d NotesDocument) LastAccessed() (domino.Time, error) {
	val, err := d.Com().GetProperty("LastAccessed")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_DOC.html */
func (d NotesDocument) LastModified() (domino.Time, error) {
	val, err := d.Com().GetProperty("LastModified")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_DOC.html */
func (d NotesDocument) LockHolders() ([]domino.String, error) {
	vals, err := d.Com().GetArrayProperty("LockHolders")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFDOC_PROPERTY.html */
func (d NotesDocument) NameOfDoc() (domino.String, error) {
	val, err := d.Com().GetProperty("NameOfDoc")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOFPROFILE_PROPERTY.html */
func (d NotesDocument) NameOfProfile() (domino.String, error) {
	val, err := d.Com().GetProperty("NameOfProfile")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY.html */
func (d NotesDocument) NoteID() (domino.String, error) {
	val, err := d.Com().GetProperty("NoteID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_DOCUMENT_COM.html */
func (d NotesDocument) NotesURL() (domino.String, error) {
	val, err := d.Com().GetProperty("NotesURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDOCUMENTUNID_PROPERTY.html */
func (d NotesDocument) ParentDocumentUNID() (domino.String, error) {
	val, err := d.Com().GetProperty("ParentDocumentUNID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func (d NotesDocument) SaveMessageOnSend() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("SaveMessageOnSend")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEMESSAGEONSEND_PROPERTY.html */
func (d NotesDocument) SetSaveMessageOnSend(v domino.Boolean) error {
	return d.Com().PutProperty("SaveMessageOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENTBYAGENT_PROPERTY.html */
func (d NotesDocument) SentByAgent() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("SentByAgent")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNER_PROPERTY.html */
func (d NotesDocument) Signer() (domino.String, error) {
	val, err := d.Com().GetProperty("Signer")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func (d NotesDocument) SignOnSend() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("SignOnSend")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNONSEND_PROPERTY.html */
func (d NotesDocument) SetSignOnSend(v domino.Boolean) error {
	return d.Com().PutProperty("SignOnSend", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOFDOC_PROPERTY.html */
func (d NotesDocument) UserNameOfDoc() (domino.String, error) {
	val, err := d.Com().GetProperty("UserNameOfDoc")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIZE_PROPERTY_DOC.html */
func (d NotesDocument) Size() (domino.Long, error) {
	val, err := d.Com().GetProperty("Size")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func (d NotesDocument) UniversalID() (domino.String, error) {
	val, err := d.Com().GetProperty("UniversalID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_DOC.html */
func (d NotesDocument) SetUniversalID(v domino.String) error {
	return d.Com().PutProperty("UniversalID", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFIER_PROPERTY.html */
func (d NotesDocument) Verifier() (domino.String, error) {
	val, err := d.Com().GetProperty("Verifier")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDITEMVALUE_METHOD.html */
func (d NotesDocument) AppendItemValue(itemName domino.String, value any) (notesitem.NotesItem, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("AppendItemValue", itemName, value)
	return notesitem.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHVCARD_METHOD.html */
func (d NotesDocument) AttachVCard(clientADTObject NotesDocument) error {
	_, err := d.Com().CallMethod("AttachVCard", clientADTObject)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSEMIMEENTITIES_METHOD_DOC.html */
type closeMIMEEntitiesParams struct {
	saveChanges    *domino.Boolean
	entityItemName *domino.String
}

type closeMIMEEntitiesParam func(*closeMIMEEntitiesParams)

func WithCloseMIMEEntitiesSavechanges(saveChanges domino.Boolean) closeMIMEEntitiesParam {
	return func(c *closeMIMEEntitiesParams) {
		c.saveChanges = &saveChanges
	}
}

func WithCloseMIMEEntitiesEntityitemname(entityItemName domino.String) closeMIMEEntitiesParam {
	return func(c *closeMIMEEntitiesParams) {
		c.entityItemName = &entityItemName
	}
}

func (d NotesDocument) CloseMIMEEntities(params ...closeMIMEEntitiesParam) (domino.Boolean, error) {
	paramsStruct := &closeMIMEEntitiesParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.saveChanges != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.saveChanges)
		if paramsStruct.entityItemName != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.entityItemName)
		}
	}
	val, err := d.Com().CallMethod("CloseMIMEEntities", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPUTEWITHFORM_METHOD.html */
func (d NotesDocument) ComputeWithForm(doDataTypes domino.Boolean, raiseError domino.Boolean) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("ComputeWithForm", doDataTypes, raiseError)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTTOMIME_METHOD.html */
type convertToMIMEParams struct {
	conversionType *domino.Integer
	options        *domino.Long
}

type convertToMIMEParam func(*convertToMIMEParams)

func WithConvertToMIMEConversionType(conversionType domino.Integer) convertToMIMEParam {
	return func(c *convertToMIMEParams) {
		c.conversionType = &conversionType
	}
}

func WithConvertToMIMEOptions(options domino.Long) convertToMIMEParam {
	return func(c *convertToMIMEParams) {
		c.options = &options
	}
}

func (d NotesDocument) ConvertToMIME(params ...convertToMIMEParam) error {
	paramsStruct := &convertToMIMEParams{}
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
	_, err := d.Com().CallMethod("ConvertToMIME", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYALLITEMS_METHOD.html */
func (d NotesDocument) CopyAllItems(notesDocument NotesDocument, replace domino.Boolean) error {
	_, err := d.Com().CallMethod("CopyAllItems", notesDocument.Com().Dispatch(), replace)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEM_METHOD.html */
func (d NotesDocument) CopyItem(item notesitem.NotesItem, newName domino.String) (notesitem.NotesItem, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CopyItem", item.Com().Dispatch(), newName)
	return notesitem.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMIMEENTITY_METHOD_DOC.html */
type createMIMEEntityParams struct {
	itemName *domino.String
}

type createMIMEEntityParam func(*createMIMEEntityParams)

func WithCreateMIMEEntityItemName(itemName domino.String) createMIMEEntityParam {
	return func(c *createMIMEEntityParams) {
		c.itemName = &itemName
	}
}

func (d NotesDocument) CreateMIMEEntity(params ...createMIMEEntityParam) (notesmimeentity.NotesMIMEEntity, error) {
	paramsStruct := &createMIMEEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.itemName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.itemName)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateMIMEEntity", paramsOrdered...)
	return notesmimeentity.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLYMESSAGE_METHOD.html */
func (d NotesDocument) CreateReplyMessage(all domino.Boolean) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateReplyMessage", all)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTITEM_METHOD.html */
func (d NotesDocument) CreateRichTextItem(name domino.String) (notesrichtextitem.NotesRichTextItem, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateRichTextItem", name)
	return notesrichtextitem.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCRYPT_METHOD.html */
func (d NotesDocument) Encrypt() error {
	_, err := d.Com().CallMethod("Encrypt")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETATTACHMENT_METHOD.html */
func (d NotesDocument) GetAttachment(fileName domino.String) (notesembeddedobject.NotesEmbeddedObject, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetAttachment", fileName)
	return notesembeddedobject.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEM_METHOD.html */
func (d NotesDocument) GetFirstItem(name domino.String) (notesitem.NotesItem, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetFirstItem", name)
	return notesitem.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUE_METHOD.html */
func (d NotesDocument) GetItemValue(itemName domino.String) (any, error) {
	return d.Com().CallMethod("GetItemValue", itemName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
func (d NotesDocument) GetItemValueCustomDataBytes(itemName domino.String, dataTypeName domino.String) ([]domino.Byte, error) {
	vals, err := d.Com().CallArrayMethod("GetItemValueCustomDataBytes", itemName, dataTypeName)
	return helpers.CastSlice[domino.Byte](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETITEMVALUEDATETIMEARRAY_METHOD.html */
/* TODO: Find out how to handle different return value types. */
func (d NotesDocument) GetItemValueDateTimeArray(itemName domino.String) ([]notesdatetime.NotesDateTime, error) {
	return com.GetObjectArrayProperty(d.Com(), notesdatetime.New, "GetItemValueDateTimeArray")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_DOC.html */
func (d NotesDocument) GetMIMEEntity(itemName domino.String) (notesmimeentity.NotesMIMEEntity, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetMIMEEntity", itemName)
	return notesmimeentity.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETREAD_METHOD_DOC.html */
type getReadParams struct {
	username *domino.String
}

type getReadParam func(*getReadParams)

func WithGetReadUsername(username domino.String) getReadParam {
	return func(c *getReadParams) {
		c.username = &username
	}
}

func (d NotesDocument) GetRead(params ...getReadParam) (domino.Boolean, error) {
	paramsStruct := &getReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	val, err := d.Com().CallMethod("GetRead", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETRECEIVEDITEMTEXT_METHOD_DOC.html */
func (d NotesDocument) GetReceivedItemText() (domino.String, error) {
	val, err := d.Com().CallMethod("GetReceivedItemText")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASITEM_METHOD.html */
func (d NotesDocument) HasItem(itemName domino.String) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("HasItem", itemName)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_DOC.html */
type lockParams struct {
	name          *[]domino.String
	provisionalOK *domino.Boolean
}

type lockParam func(*lockParams)

func WithLockName(name []domino.String) lockParam {
	return func(c *lockParams) {
		c.name = &name
	}
}

func WithLockProvisionalOK(provisionalOK domino.Boolean) lockParam {
	return func(c *lockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (d NotesDocument) Lock(params ...lockParam) (domino.Boolean, error) {
	paramsStruct := &lockParams{}
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
	val, err := d.Com().CallMethod("Lock", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_DOC.html */
type lockProvisionalParams struct {
	name *[]domino.String
}

type lockProvisionalParam func(*lockProvisionalParams)

func WithLockProvisionalName(name []domino.String) lockProvisionalParam {
	return func(c *lockProvisionalParams) {
		c.name = &name
	}
}

func (d NotesDocument) LockProvisional(params ...lockProvisionalParam) (domino.Boolean, error) {
	paramsStruct := &lockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	val, err := d.Com().CallMethod("LockProvisional", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAKERESPONSE_METHOD.html */
func (d NotesDocument) MakeResponse(document NotesDocument) error {
	_, err := d.Com().CallMethod("MakeResponse", document.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKREAD_DOCUMENT.html */
type markReadParams struct {
	username *domino.String
}

type markReadParam func(*markReadParams)

func WithMarkReadUsername(username domino.String) markReadParam {
	return func(c *markReadParams) {
		c.username = &username
	}
}

func (d NotesDocument) MarkRead(params ...markReadParam) error {
	paramsStruct := &markReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := d.Com().CallMethod("MarkRead", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKUNREAD_DOCUMENT.html */
type markUnreadParams struct {
	username *domino.String
}

type markUnreadParam func(*markUnreadParams)

func WithMarkUnreadUsername(username domino.String) markUnreadParam {
	return func(c *markUnreadParams) {
		c.username = &username
	}
}

func (d NotesDocument) MarkUnread(params ...markUnreadParam) error {
	paramsStruct := &markUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := d.Com().CallMethod("MarkUnread", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTINFOLDER_METHOD.html */
type putInFolderParams struct {
	createonfail *domino.Boolean
}

type putInFolderParam func(*putInFolderParams)

func WithPutInFolderCreateonfail(createonfail domino.Boolean) putInFolderParam {
	return func(c *putInFolderParams) {
		c.createonfail = &createonfail
	}
}

func (d NotesDocument) PutInFolder(folderName domino.String, params ...putInFolderParam) error {
	paramsStruct := &putInFolderParams{}
	paramsOrdered := []interface{}{folderName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	_, err := d.Com().CallMethod("PutInFolder", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_DOC.html */
func (d NotesDocument) Remove(force domino.Boolean) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("Remove", force)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEFROMFOLDER_METHOD.html */
func (d NotesDocument) RemoveFromFolder(folderName domino.String) error {
	_, err := d.Com().CallMethod("RemoveFromFolder", folderName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEITEM_METHOD.html */
func (d NotesDocument) RemoveItem(itemName domino.String) error {
	_, err := d.Com().CallMethod("RemoveItem", itemName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEPERMANENTLY_METHOD_DOC.html */
func (d NotesDocument) RemovePermanently(force domino.Boolean) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("RemovePermanently", force)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENDERTORTITEM_METHOD.html */
func (d NotesDocument) RenderToRTItem(notesRichTextItem notesrichtextitem.NotesRichTextItem) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("RenderToRTItem", notesRichTextItem.Com().Dispatch())
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUE_METHOD.html */
func (d NotesDocument) ReplaceItemValue(itemName domino.String, value any) (notesitem.NotesItem, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("ReplaceItemValue", itemName, value)
	return notesitem.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEITEMVALUECUSTOMDATABYTES_METHOD_DOC.html */
func (d NotesDocument) ReplaceItemValueCustomDataBytes(itemName domino.String, dataTypeName domino.String, byteArray []domino.Byte) error {
	_, err := d.Com().CallMethod("ReplaceItemValueCustomDataBytes", itemName, dataTypeName, byteArray)
	return err
}

type saveParams struct {
	markRead *domino.Boolean
}

type saveParam func(*saveParams)

func WithSaveMarkRead(markRead domino.Boolean) saveParam {
	return func(c *saveParams) {
		c.markRead = &markRead
	}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_DOC.html */
func (d NotesDocument) Save(force domino.Boolean, createResponse domino.Boolean, params ...saveParam) (domino.Boolean, error) {
	paramsStruct := &saveParams{}
	paramsOrdered := []interface{}{force, createResponse}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.markRead != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.markRead)
	}
	val, err := d.Com().CallMethod("Save", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEND_METHOD_DOC.html */
type sendParams struct {
	recipients *[]domino.String
}

type sendParam func(*sendParams)

func WithSendRecipients(recipients []domino.String) sendParam {
	return func(c *sendParams) {
		c.recipients = &recipients
	}
}

func (d NotesDocument) Send(attachForm domino.Boolean, params ...sendParam) error {
	paramsStruct := &sendParams{}
	paramsOrdered := []interface{}{attachForm}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.recipients != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.recipients)
	}
	_, err := d.Com().CallMethod("Send", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGN_METHOD.html */
func (d NotesDocument) Sign() error {
	_, err := d.Com().CallMethod("Sign")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_DOC.html */
func (d NotesDocument) UnLock() error {
	_, err := d.Com().CallMethod("UnLock")
	return err
}
