/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESITEM_CLASS.html */
package notesitem

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesdatetime"
	"domigo/domino/notesmimeentity"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type DataType = domino.Long

const (
	DATATYPE_ACTIONCD       DataType = 16
	DATATYPE_ASSISTANTINFO  DataType = 17
	DATATYPE_ATTACHMENT     DataType = 1084
	DATATYPE_AUTHORS        DataType = 1076
	DATATYPE_COLLATION      DataType = 2
	DATATYPE_DATETIMES      DataType = 1024
	DATATYPE_EMBEDDEDOBJECT DataType = 1090
	DATATYPE_ERRORITEM      DataType = 256
	DATATYPE_FORMULA        DataType = 1536
	DATATYPE_HTML           DataType = 21
	DATATYPE_ICON           DataType = 6
	DATATYPE_LSOBJECT       DataType = 20
	DATATYPE_MIME_PART      DataType = 25
	DATATYPE_NAMES          DataType = 1074
	DATATYPE_NOTELINKS      DataType = 7
	DATATYPE_NOTEREFS       DataType = 4
	DATATYPE_NUMBERS        DataType = 768
	DATATYPE_OTHEROBJECT    DataType = 1085
	DATATYPE_QUERYCD        DataType = 15
	DATATYPE_READERS        DataType = 1075
	DATATYPE_RFC822TEXT     DataType = 1282
	DATATYPE_RICHTEXT       DataType = 1
	DATATYPE_SIGNATURE      DataType = 8
	DATATYPE_TEXT           DataType = 1280
	DATATYPE_UNAVAILABLE    DataType = 512
	DATATYPE_UNKNOWN        DataType = 0
	DATATYPE_USERDATA       DataType = 14
	DATATYPE_USERID         DataType = 1792
	DATATYPE_VIEWMAPDATA    DataType = 18
	DATATYPE_VIEWMAPLAYOUT  DataType = 19
)

type NotesItem struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesItem {
	return NotesItem{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func (i NotesItem) DateTimeValue() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := i.Com().GetObjectProperty("DateTimeValue")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func (i NotesItem) SetDateTimeValue(v notesdatetime.NotesDateTime) error {
	return i.Com().PutProperty("DateTimeValue", v.Com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func (i NotesItem) IsAuthors() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsAuthors")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func (i NotesItem) SetIsAuthors(v domino.Boolean) error {
	return i.Com().PutProperty("IsAuthors", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func (i NotesItem) IsEncrypted() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsEncrypted")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func (i NotesItem) SetIsEncrypted(v domino.Boolean) error {
	return i.Com().PutProperty("IsEncrypted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func (i NotesItem) IsNames() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsNames")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func (i NotesItem) SetIsNames(v domino.Boolean) error {
	return i.Com().PutProperty("IsNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func (i NotesItem) IsProtected() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsProtected")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func (i NotesItem) SetIsProtected(v domino.Boolean) error {
	return i.Com().PutProperty("IsProtected", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func (i NotesItem) IsReaders() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsReaders")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func (i NotesItem) SetIsReaders(v domino.Boolean) error {
	return i.Com().PutProperty("IsReaders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func (i NotesItem) IsSigned() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsSigned")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func (i NotesItem) SetIsSigned(v domino.Boolean) error {
	return i.Com().PutProperty("IsSigned", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func (i NotesItem) IsSummary() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("IsSummary")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func (i NotesItem) SetIsSummary(v domino.Boolean) error {
	return i.Com().PutProperty("IsSummary", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_ITEM.html */
func (i NotesItem) LastModified() (domino.Time, error) {
	val, err := i.Com().GetProperty("LastModified")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ITEM.html */
func (i NotesItem) Name() (domino.String, error) {
	val, err := i.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func (i NotesItem) SaveToDisk() (domino.Boolean, error) {
	val, err := i.Com().GetProperty("SaveToDisk")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func (i NotesItem) SetSaveToDisk(v domino.Boolean) error {
	return i.Com().PutProperty("SaveToDisk", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY.html */
func (i NotesItem) Text() (domino.String, error) {
	val, err := i.Com().GetProperty("Text")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_ITEM.html */
func (i NotesItem) Type() (DataType, error) {
	val, err := i.Com().GetProperty("Type")
	return helpers.CastValue[DataType](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUELENGTH_PROPERTY.html */
func (i NotesItem) ValueLength() (domino.Long, error) {
	val, err := i.Com().GetProperty("ValueLength")
	return helpers.CastValue[domino.Long](val), err
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
func (i NotesItem) Abstract(maxAbstract domino.Long, dropVowels domino.Boolean, useDictionary domino.Boolean) (domino.String, error) {
	val, err := i.Com().CallMethod("Abstract", maxAbstract, dropVowels, useDictionary)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDTOTEXTLIST_METHOD.html */
func (i NotesItem) AppendToTextList(newValues []domino.String) error {
	_, err := i.Com().CallMethod("AppendToTextList", newValues)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD.html */
func (i NotesItem) Contains(value any) (domino.Boolean, error) {
	val, err := i.Com().CallMethod("Contains", value)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func (i NotesItem) GetValueCustomDataBytes(dataTypeName domino.String) ([]domino.Byte, error) {
	vals, err := i.Com().CallArrayMethod("GetValueCustomDataBytes", dataTypeName)
	return helpers.CastSlice[domino.Byte](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUEDATETIMEARRAY_METHOD.html */
/* TODO: Check documentation, return type can vary... */
func (i NotesItem) GetValueDateTimeArray() ([]notesdatetime.NotesDateTime, error) {
	return com.GetObjectArrayProperty(i.Com(), notesdatetime.New, "GetValueDateTimeArray")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_NOTESITEM.html */
func (i NotesItem) GetMIMEEntity() (notesmimeentity.NotesMIMEEntity, error) {
	dispatchPtr, err := i.Com().CallObjectMethod("GetMIMEEntity")
	return notesmimeentity.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ITEM.html */
func (i NotesItem) Remove() error {
	_, err := i.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func (i NotesItem) SetValueCustomDataBytes(dataTypeName domino.String, byteArray []domino.Byte) error {
	_, err := i.Com().CallMethod("SetValueCustomDataBytes", dataTypeName, byteArray)
	return err
}
