/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESITEM_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesItemDataType = Long

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
	return getComObjectProperty(i, NewNotesDateTime, "DateTimeValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func (i NotesItem) SetDateTimeValue(v NotesDateTime) error {
	return putComProperty(i, "DateTimeValue", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func (i NotesItem) IsAuthors() (Boolean, error) {
	return getComProperty[Boolean](i, "IsAuthors")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func (i NotesItem) SetIsAuthors(v Boolean) error {
	return putComProperty(i, "IsAuthors", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func (i NotesItem) IsEncrypted() (Boolean, error) {
	return getComProperty[Boolean](i, "IsEncrypted")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func (i NotesItem) SetIsEncrypted(v Boolean) error {
	return putComProperty(i, "IsEncrypted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func (i NotesItem) IsNames() (Boolean, error) {
	return getComProperty[Boolean](i, "IsNames")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func (i NotesItem) SetIsNames(v Boolean) error {
	return putComProperty(i, "IsNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func (i NotesItem) IsProtected() (Boolean, error) {
	return getComProperty[Boolean](i, "IsProtected")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func (i NotesItem) SetIsProtected(v Boolean) error {
	return putComProperty(i, "IsProtected", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func (i NotesItem) IsReaders() (Boolean, error) {
	return getComProperty[Boolean](i, "IsReaders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func (i NotesItem) SetIsReaders(v Boolean) error {
	return putComProperty(i, "IsReaders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func (i NotesItem) IsSigned() (Boolean, error) {
	return getComProperty[Boolean](i, "IsSigned")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func (i NotesItem) SetIsSigned(v Boolean) error {
	return putComProperty(i, "IsSigned", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func (i NotesItem) IsSummary() (Boolean, error) {
	return getComProperty[Boolean](i, "IsSummary")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func (i NotesItem) SetIsSummary(v Boolean) error {
	return putComProperty(i, "IsSummary", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_ITEM.html */
func (i NotesItem) LastModified() (Time, error) {
	return getComProperty[Time](i, "LastModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ITEM.html */
func (i NotesItem) Name() (String, error) {
	return getComProperty[String](i, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ITEM.html */
func (i NotesItem) Parent() (NotesDocument, error) {
	return getComObjectProperty(i, NewNotesDocument, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func (i NotesItem) SaveToDisk() (Boolean, error) {
	return getComProperty[Boolean](i, "SaveToDisk")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func (i NotesItem) SetSaveToDisk(v Boolean) error {
	return putComProperty(i, "SaveToDisk", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY.html */
func (i NotesItem) Text() (String, error) {
	return getComProperty[String](i, "Text")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_ITEM.html */
func (i NotesItem) Type() (NotesItemDataType, error) {
	return getComProperty[NotesItemDataType](i, "Type")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUELENGTH_PROPERTY.html */
func (i NotesItem) ValueLength() (Long, error) {
	return getComProperty[Long](i, "ValueLength")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUES_PROPERTY.html */
/* TODO: Handle different types. */
/* TODO: Wrap values into struct. */
func (i NotesItem) Values() (any, error) {
	return getComAnyArrayProperty(i, "Values")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUES_PROPERTY.html */
/* TODO: Handle different types. */
/* TODO: Wrap values into struct. */
func (i NotesItem) SetValues(v any) error {
	return putComProperty(i, "Values", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_METHOD.html */
func (i NotesItem) Abstract(maxAbstract Long, dropVowels Boolean, useDictionary Boolean) (String, error) {
	return callComMethod[String](i, "Abstract", maxAbstract, dropVowels, useDictionary)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDTOTEXTLIST_METHOD.html */
func (i NotesItem) AppendToTextList(newValues []String) error {
	return callComVoidMethod(i, "AppendToTextList", newValues)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD.html */
func (i NotesItem) Contains(value any) (Boolean, error) {
	return callComMethod[Boolean](i, "Contains", value)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYITEMTODOCUMENT_METHOD.html */
func (i NotesItem) CopyItemToDocument(document NotesDocument, newName String) (NotesItem, error) {
	return callComObjectMethod(i, NewNotesItem, "CopyItemToDocument", document, newName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func (i NotesItem) GetValueCustomDataBytes(dataTypeName String) ([]Byte, error) {
	return callComArrayMethod[Byte](i, "GetValueCustomDataBytes", dataTypeName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUEDATETIMEARRAY_METHOD.html */
/* TODO: Check documentation, return type can vary... */
func (i NotesItem) GetValueDateTimeArray() ([]NotesDateTime, error) {
	return com.CallObjectArrayMethod(i.com(), NewNotesDateTime, "GetValueDateTimeArray")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_NOTESITEM.html */
func (i NotesItem) GetMIMEEntity() (NotesMIMEEntity, error) {
	return callComObjectMethod(i, NewNotesMIMEEntity, "GetMIMEEntity")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ITEM.html */
func (i NotesItem) Remove() error {
	return callComVoidMethod(i, "Remove")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func (i NotesItem) SetValueCustomDataBytes(dataTypeName String, byteArray []Byte) error {
	return callComVoidMethod(i, "SetValueCustomDataBytes", dataTypeName, byteArray)
}
