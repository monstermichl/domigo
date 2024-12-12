/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORY_CLASS.html */
package notesdirectory

import (
	"domigo/domino"
	"domigo/domino/notesdirectorynavigator"
	"domigo/domino/notesview"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDirectory struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDirectory {
	return NotesDirectory{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLEITEMS_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) AvailableItems() ([]domino.Variant, error) {
	vals, err := d.Com().GetArrayProperty("AvailableItems")
	return helpers.CastSlice[domino.Variant](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLENAMES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) AvailableNames() ([]domino.Variant, error) {
	vals, err := d.Com().GetArrayProperty("AvailableNames")
	return helpers.CastSlice[domino.Variant](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLEVIEW_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) AvailableView() (domino.String, error) {
	val, err := d.Com().GetProperty("AvailableView")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPAUTHORIZATIONONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) GroupAuthorizationOnly() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("GroupAuthorizationOnly")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPAUTHORIZATIONONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetGroupAuthorizationOnly(v domino.Boolean) error {
	return d.Com().PutProperty("GroupAuthorizationOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITMATCHES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) LimitMatches() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("LimitMatches")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITMATCHES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetLimitMatches(v domino.Boolean) error {
	return d.Com().PutProperty("LimitMatches", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARTIALMATCHES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) PartialMatches() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("PartialMatches")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCHALLDIRECTORIES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SearchAllDirectories() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("SearchAllDirectories")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCHALLDIRECTORIES_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetSearchAllDirectories(v domino.Boolean) error {
	return d.Com().PutProperty("SearchAllDirectories", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) Server() (domino.String, error) {
	val, err := d.Com().GetProperty("Server")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUSTEDONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) TrustedOnly() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("TrustedOnly")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUSTEDONLY_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetTrustedOnly(v domino.Boolean) error {
	return d.Com().PutProperty("TrustedOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECONTEXTSERVER_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) UseContextServer() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("UseContextServer")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECONTEXTSERVER_PROPERTY_DIRECTORY.html */
func (d NotesDirectory) SetUseContextServer(v domino.Boolean) error {
	return d.Com().PutProperty("UseContextServer", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENAVIGATOR_METHOD_DIRECTORY.html */
func (d NotesDirectory) CreateNavigator() (notesdirectorynavigator.NotesDirectoryNavigator, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("CreateNavigator")
	return notesdirectorynavigator.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREELOOKUPBUFFER_METHOD_DIRECTORY.html */
func (d NotesDirectory) FreeLookupBuffer() error {
	_, err := d.Com().CallMethod("FreeLookupBuffer")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMAILINFO_METHOD_DIRECTORY.html */
type getMailInfoParams struct {
	getver                 *domino.Boolean
	errorOnMultipleMatches *domino.Boolean
}

type getMailInfoParam func(*getMailInfoParams)

func WithGetMailInfoGetver(getver domino.Boolean) getMailInfoParam {
	return func(c *getMailInfoParams) {
		c.getver = &getver
	}
}

func WithGetMailInfoErrorOnMultipleMatches(errorOnMultipleMatches domino.Boolean) getMailInfoParam {
	return func(c *getMailInfoParams) {
		c.errorOnMultipleMatches = &errorOnMultipleMatches
	}
}

func (d NotesDirectory) GetMailInfo(username domino.String, params ...getMailInfoParam) ([]domino.String, error) {
	paramsStruct := &getMailInfoParams{}
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
	vals, err := d.Com().CallArrayMethod("GetMailInfo", paramsOrdered...)
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOOKUPALLNAMES_METHOD_DIRECTORY.html */
func (d NotesDirectory) LookupAllNames(view notesview.NotesView, items domino.Variant) (notesdirectorynavigator.NotesDirectoryNavigator, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("LookupAllNames", view.Com().Dispatch(), items)
	return notesdirectorynavigator.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOOKUPNAMES_METHOD_DIRECTORY.html */
type lookupNamesParams struct {
	partialmatches *domino.Boolean
}

type lookupNamesParam func(*lookupNamesParams)

func WithLookupNamesPartialmatches(partialmatches domino.Boolean) lookupNamesParam {
	return func(c *lookupNamesParams) {
		c.partialmatches = &partialmatches
	}
}

func (d NotesDirectory) LookupNames(view notesview.NotesView, names []domino.Variant, items []domino.Variant, params ...lookupNamesParam) (notesdirectorynavigator.NotesDirectoryNavigator, error) {
	paramsStruct := &lookupNamesParams{}
	paramsOrdered := []interface{}{view, names, items}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.partialmatches != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.partialmatches)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("LookupNames", paramsOrdered...)
	return notesdirectorynavigator.New(dispatchPtr), err
}
