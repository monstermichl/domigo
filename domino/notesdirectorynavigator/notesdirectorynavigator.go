/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORYNAVIGATOR_CLASS.html */
package notesdirectorynavigator

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDirectoryNavigator struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDirectoryNavigator {
	return NotesDirectoryNavigator{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTITEM_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentItem() (domino.String, error) {
	val, err := d.Com().GetProperty("CurrentItem")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCH_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentMatch() (domino.Long, error) {
	val, err := d.Com().GetProperty("CurrentMatch")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCHES_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentMatches() (domino.Long, error) {
	val, err := d.Com().GetProperty("CurrentMatches")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTNAME_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentName() (domino.String, error) {
	val, err := d.Com().GetProperty("CurrentName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTVIEW_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentView() (domino.String, error) {
	val, err := d.Com().GetProperty("CurrentView")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MATCHLOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) MatchLocated() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("MatchLocated")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMELOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) NameLocated() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("NameLocated")
	return helpers.CastValue[domino.Boolean](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindFirstMatch() (domino.Boolean, error) {
	val, err := d.Com().CallMethod("FindFirstMatch")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindFirstName() (domino.Long, error) {
	val, err := d.Com().CallMethod("FindFirstName")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNextMatch() (domino.Boolean, error) {
	val, err := d.Com().CallMethod("FindNextMatch")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNextName() (domino.Long, error) {
	val, err := d.Com().CallMethod("FindNextName")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNthMatch(n domino.Integer) (domino.Boolean, error) {
	val, err := d.Com().CallMethod("FindNthMatch", n)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNthName(n domino.Integer) (domino.Long, error) {
	val, err := d.Com().CallMethod("FindNthName", n)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetFirstItemValue() (any, error) {
	return d.Com().CallMethod("GetFirstItemValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetNextItemValue() (any, error) {
	return d.Com().CallMethod("GetNextItemValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetNthItemValue(n domino.Integer) (any, error) {
	return d.Com().CallMethod("GetNthItemValue", n)
}
