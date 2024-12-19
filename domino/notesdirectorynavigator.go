/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORYNAVIGATOR_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDirectoryNavigator struct {
	NotesStruct
}

func NewNotesDirectoryNavigator(dispatchPtr *ole.IDispatch) NotesDirectoryNavigator {
	return NotesDirectoryNavigator{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTITEM_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentItem() (String, error) {
	val, err := d.com().GetProperty("CurrentItem")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCH_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentMatch() (Long, error) {
	val, err := d.com().GetProperty("CurrentMatch")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCHES_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentMatches() (Long, error) {
	val, err := d.com().GetProperty("CurrentMatches")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTNAME_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentName() (String, error) {
	val, err := d.com().GetProperty("CurrentName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTVIEW_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentView() (String, error) {
	val, err := d.com().GetProperty("CurrentView")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MATCHLOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) MatchLocated() (Boolean, error) {
	val, err := d.com().GetProperty("MatchLocated")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMELOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) NameLocated() (Boolean, error) {
	val, err := d.com().GetProperty("NameLocated")
	return helpers.CastValue[Boolean](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindFirstMatch() (Boolean, error) {
	val, err := d.com().CallMethod("FindFirstMatch")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindFirstName() (Long, error) {
	val, err := d.com().CallMethod("FindFirstName")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNextMatch() (Boolean, error) {
	val, err := d.com().CallMethod("FindNextMatch")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNextName() (Long, error) {
	val, err := d.com().CallMethod("FindNextName")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNthMatch(n Integer) (Boolean, error) {
	val, err := d.com().CallMethod("FindNthMatch", n)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNthName(n Integer) (Long, error) {
	val, err := d.com().CallMethod("FindNthName", n)
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetFirstItemValue() (any, error) {
	return d.com().CallMethod("GetFirstItemValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetNextItemValue() (any, error) {
	return d.com().CallMethod("GetNextItemValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetNthItemValue(n Integer) (any, error) {
	return d.com().CallMethod("GetNthItemValue", n)
}
