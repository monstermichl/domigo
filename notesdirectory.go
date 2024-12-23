/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORY_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesDirectory struct {
	NotesStruct
}

func NewNotesDirectory(dispatchPtr *ole.IDispatch) NotesDirectory {
	return NotesDirectory{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLEITEMS_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) AvailableItems() ([]Variant, error) {
	return getComArrayProperty[Variant](d, "AvailableItems")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLENAMES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) AvailableNames() ([]Variant, error) {
	return getComArrayProperty[Variant](d, "AvailableNames")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLEVIEW_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) AvailableView() (String, error) {
	return getComProperty[String](d, "AvailableView")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPAUTHORIZATIONONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) GroupAuthorizationOnly() (Boolean, error) {
	return getComProperty[Boolean](d, "GroupAuthorizationOnly")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPAUTHORIZATIONONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetGroupAuthorizationOnly(v Boolean) error {
	return putComProperty(d, "GroupAuthorizationOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITMATCHES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) LimitMatches() (Boolean, error) {
	return getComProperty[Boolean](d, "LimitMatches")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITMATCHES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetLimitMatches(v Boolean) error {
	return putComProperty(d, "LimitMatches", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARTIALMATCHES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) PartialMatches() (Boolean, error) {
	return getComProperty[Boolean](d, "PartialMatches")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCHALLDIRECTORIES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SearchAllDirectories() (Boolean, error) {
	return getComProperty[Boolean](d, "SearchAllDirectories")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCHALLDIRECTORIES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetSearchAllDirectories(v Boolean) error {
	return putComProperty(d, "SearchAllDirectories", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) Server() (String, error) {
	return getComProperty[String](d, "Server")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUSTEDONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) TrustedOnly() (Boolean, error) {
	return getComProperty[Boolean](d, "TrustedOnly")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUSTEDONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetTrustedOnly(v Boolean) error {
	return putComProperty(d, "TrustedOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECONTEXTSERVER_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) UseContextServer() (Boolean, error) {
	return getComProperty[Boolean](d, "UseContextServer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECONTEXTSERVER_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetUseContextServer(v Boolean) error {
	return putComProperty(d, "UseContextServer", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENAVIGATOR_METHOD_DIRECTORY.html */
func (d NotesDirectory) CreateNavigator() (NotesDirectoryNavigator, error) {
	return callComObjectMethod(d, NewNotesDirectoryNavigator, "CreateNavigator")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREELOOKUPBUFFER_METHOD_DIRECTORY.html */
func (d NotesDirectory) FreeLookupBuffer() error {
	return callComVoidMethod(d, "FreeLookupBuffer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMAILINFO_METHOD_DIRECTORY.html */
type notesDirectoryGetMailInfoParams struct {
	getver                 *Boolean
	errorOnMultipleMatches *Boolean
}

type notesDirectoryGetMailInfoParam func(*notesDirectoryGetMailInfoParams)

func WithNotesDirectoryGetMailInfoGetver(getver Boolean) notesDirectoryGetMailInfoParam {
	return func(c *notesDirectoryGetMailInfoParams) {
		c.getver = &getver
	}
}

func WithNotesDirectoryGetMailInfoErrorOnMultipleMatches(errorOnMultipleMatches Boolean) notesDirectoryGetMailInfoParam {
	return func(c *notesDirectoryGetMailInfoParams) {
		c.errorOnMultipleMatches = &errorOnMultipleMatches
	}
}

func (d NotesDirectory) GetMailInfo(username String, params ...notesDirectoryGetMailInfoParam) ([]String, error) {
	paramsStruct := &notesDirectoryGetMailInfoParams{}
	paramsOrdered := []interface{}{username}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.getver != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.getver)
		if paramsStruct.errorOnMultipleMatches != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.errorOnMultipleMatches)
		}
	}
	return callComArrayMethod[String](d, "GetMailInfo", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOOKUPALLNAMES_METHOD_DIRECTORY.html */
func (d NotesDirectory) LookupAllNames(view NotesView, items []String) (NotesDirectoryNavigator, error) {
	return callComObjectMethod(d, NewNotesDirectoryNavigator, "LookupAllNames", view.com().Dispatch(), items)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOOKUPNAMES_METHOD_DIRECTORY.html */
type notesDirectoryLookupNamesParams struct {
	partialmatches *Boolean
}

type notesDirectoryLookupNamesParam func(*notesDirectoryLookupNamesParams)

func WithNotesDirectoryLookupNamesPartialmatches(partialmatches Boolean) notesDirectoryLookupNamesParam {
	return func(c *notesDirectoryLookupNamesParams) {
		c.partialmatches = &partialmatches
	}
}

func (d NotesDirectory) LookupNames(view NotesView, names []String, items []String, params ...notesDirectoryLookupNamesParam) (NotesDirectoryNavigator, error) {
	paramsStruct := &notesDirectoryLookupNamesParams{}
	paramsOrdered := []interface{}{view, names, items}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.partialmatches != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.partialmatches)
	}
	return callComObjectMethod(d, NewNotesDirectoryNavigator, "LookupNames", paramsOrdered...)
}
