/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESITEM_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesItemDataType Long

const (
	NOTESITEM_TYPE_ACTIONCD       NotesItemDataType = 16
	NOTESITEM_TYPE_ASSISTANTINFO  NotesItemDataType = 17
	NOTESITEM_TYPE_ATTACHMENT     NotesItemDataType = 1084
	NOTESITEM_TYPE_AUTHORS        NotesItemDataType = 1076
	NOTESITEM_TYPE_COLLATION      NotesItemDataType = 2
	NOTESITEM_TYPE_DATETIMES      NotesItemDataType = 1024
	NOTESITEM_TYPE_EMBEDDEDOBJECT NotesItemDataType = 1090
	NOTESITEM_TYPE_ERRORITEM      NotesItemDataType = 256
	NOTESITEM_TYPE_FORMULA        NotesItemDataType = 1536
	NOTESITEM_TYPE_HTML           NotesItemDataType = 21
	NOTESITEM_TYPE_ICON           NotesItemDataType = 6
	NOTESITEM_TYPE_LSOBJECT       NotesItemDataType = 20
	NOTESITEM_TYPE_MIME_PART      NotesItemDataType = 25
	NOTESITEM_TYPE_NAMES          NotesItemDataType = 1074
	NOTESITEM_TYPE_NOTELINKS      NotesItemDataType = 7
	NOTESITEM_TYPE_NOTEREFS       NotesItemDataType = 4
	NOTESITEM_TYPE_NUMBERS        NotesItemDataType = 768
	NOTESITEM_TYPE_OTHEROBJECT    NotesItemDataType = 1085
	NOTESITEM_TYPE_QUERYCD        NotesItemDataType = 15
	NOTESITEM_TYPE_READERS        NotesItemDataType = 1075
	NOTESITEM_TYPE_RFC822TEXT     NotesItemDataType = 1282
	NOTESITEM_TYPE_RICHTEXT       NotesItemDataType = 1
	NOTESITEM_TYPE_SIGNATURE      NotesItemDataType = 8
	NOTESITEM_TYPE_TEXT           NotesItemDataType = 1280
	NOTESITEM_TYPE_UNAVAILABLE    NotesItemDataType = 512
	NOTESITEM_TYPE_UNKNOWN        NotesItemDataType = 0
	NOTESITEM_TYPE_USERDATA       NotesItemDataType = 14
	NOTESITEM_TYPE_USERID         NotesItemDataType = 1792
	NOTESITEM_TYPE_VIEWMAPDATA    NotesItemDataType = 18
	NOTESITEM_TYPE_VIEWMAPLAYOUT  NotesItemDataType = 19
)

type NotesItem struct {
	NotesStruct
}

func NewNotesItem(dispatchPtr *ole.IDispatch) NotesItem {
	return NotesItem{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func (i NotesItem) DateTimeValue() (NotesDateTime, error) {
	dispatchPtr, err := i.Com().GetObjectProperty("DateTimeValue")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func (i NotesItem) SetDateTimeValue(v NotesDateTime) error {
	return i.Com().PutProperty("DateTimeValue", v.Com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func (i NotesItem) IsAuthors() (Boolean, error) {
	val, err := i.Com().GetProperty("IsAuthors")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func (i NotesItem) SetIsAuthors(v Boolean) error {
	return i.Com().PutProperty("IsAuthors", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func (i NotesItem) IsEncrypted() (Boolean, error) {
	val, err := i.Com().GetProperty("IsEncrypted")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func (i NotesItem) SetIsEncrypted(v Boolean) error {
	return i.Com().PutProperty("IsEncrypted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func (i NotesItem) IsNames() (Boolean, error) {
	val, err := i.Com().GetProperty("IsNames")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func (i NotesItem) SetIsNames(v Boolean) error {
	return i.Com().PutProperty("IsNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func (i NotesItem) IsProtected() (Boolean, error) {
	val, err := i.Com().GetProperty("IsProtected")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func (i NotesItem) SetIsProtected(v Boolean) error {
	return i.Com().PutProperty("IsProtected", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func (i NotesItem) IsReaders() (Boolean, error) {
	val, err := i.Com().GetProperty("IsReaders")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func (i NotesItem) SetIsReaders(v Boolean) error {
	return i.Com().PutProperty("IsReaders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func (i NotesItem) IsSigned() (Boolean, error) {
	val, err := i.Com().GetProperty("IsSigned")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func (i NotesItem) SetIsSigned(v Boolean) error {
	return i.Com().PutProperty("IsSigned", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func (i NotesItem) IsSummary() (Boolean, error) {
	val, err := i.Com().GetProperty("IsSummary")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func (i NotesItem) SetIsSummary(v Boolean) error {
	return i.Com().PutProperty("IsSummary", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_ITEM.html */
func (i NotesItem) LastModified() (Time, error) {
	val, err := i.Com().GetProperty("LastModified")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ITEM.html */
func (i NotesItem) Name() (String, error) {
	val, err := i.Com().GetProperty("Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ITEM.html */
func (i NotesItem) Parent() (NotesDocument, error) {
	dispatchPtr, err := i.Com().GetObjectProperty("Parent")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func (i NotesItem) SaveToDisk() (Boolean, error) {
	val, err := i.Com().GetProperty("SaveToDisk")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func (i NotesItem) SetSaveToDisk(v Boolean) error {
	return i.Com().PutProperty("SaveToDisk", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY.html */
func (i NotesItem) Text() (String, error) {
	val, err := i.Com().GetProperty("Text")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_ITEM.html */
func (i NotesItem) Type() (NotesItemDataType, error) {
	val, err := i.Com().GetProperty("Type")
	return helpers.CastValue[NotesItemDataType](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUELENGTH_PROPERTY.html */
func (i NotesItem) ValueLength() (Long, error) {
	val, err := i.Com().GetProperty("ValueLength")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUES_PROPERTY.html */
/* TODO: Handle different types. */
func (i NotesItem) Values() (any, error) {
	return i.Com().GetArrayProperty("Values")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUES_PROPERTY.html */
func (i NotesItem) SetValues(v any) error {
	/* TODO: Handle different types. */
	return i.Com().PutProperty("Values", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_METHOD.html */
func (i NotesItem) Abstract(maxAbstract Long, dropVowels Boolean, useDictionary Boolean) (String, error) {
	val, err := i.Com().CallMethod("Abstract", maxAbstract, dropVowels, useDictionary)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDTOTEXTLIST_METHOD.html */
func (i NotesItem) AppendToTextList(newValues []String) error {
	_, err := i.Com().CallMethod("AppendToTextList", newValues)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD.html */
func (i NotesItem) Contains(value any) (Boolean, error) {
	val, err := i.Com().CallMethod("Contains", value)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEMTODOCUMENT_METHOD.html */
func (i NotesItem) CopyItemToDocument(document NotesDocument, newName String) (NotesItem, error) {
	dispatchPtr, err := i.Com().CallObjectMethod("CopyItemToDocument", document.Com().Dispatch(), newName)
	return NewNotesItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func (i NotesItem) GetValueCustomDataBytes(dataTypeName String) ([]Byte, error) {
	vals, err := i.Com().CallArrayMethod("GetValueCustomDataBytes", dataTypeName)
	return helpers.CastSlice[Byte](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUEDATETIMEARRAY_METHOD.html */
/* TODO: Check documentation, return type can vary... */
func (i NotesItem) GetValueDateTimeArray() ([]NotesDateTime, error) {
	return com.GetObjectArrayProperty(i.Com(), NewNotesDateTime, "GetValueDateTimeArray")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_NOTESITEM.html */
func (i NotesItem) GetMIMEEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := i.Com().CallObjectMethod("GetMIMEEntity")
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ITEM.html */
func (i NotesItem) Remove() error {
	_, err := i.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func (i NotesItem) SetValueCustomDataBytes(dataTypeName String, byteArray []Byte) error {
	_, err := i.Com().CallMethod("SetValueCustomDataBytes", dataTypeName, byteArray)
	return err
}
