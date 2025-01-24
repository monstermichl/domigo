/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORYNAVIGATOR_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesDirectoryNavigator struct {
	NotesStruct
}

func newNotesDirectoryNavigator(dispatchPtr *ole.IDispatch) NotesDirectoryNavigator {
	return NotesDirectoryNavigator{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTITEM_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentItem() (String, error) {
	return getComProperty[String](d, "CurrentItem")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCH_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentMatch() (Long, error) {
	return getComProperty[Long](d, "CurrentMatch")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCHES_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentMatches() (Long, error) {
	return getComProperty[Long](d, "CurrentMatches")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTNAME_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentName() (String, error) {
	return getComProperty[String](d, "CurrentName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTVIEW_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) CurrentView() (String, error) {
	return getComProperty[String](d, "CurrentView")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MATCHLOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) MatchLocated() (Boolean, error) {
	return getComProperty[Boolean](d, "MatchLocated")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMELOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) NameLocated() (Boolean, error) {
	return getComProperty[Boolean](d, "NameLocated")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindFirstMatch() (Boolean, error) {
	return callComMethod[Boolean](d, "FindFirstMatch")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindFirstName() (Long, error) {
	return callComMethod[Long](d, "FindFirstName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNextMatch() (Boolean, error) {
	return callComMethod[Boolean](d, "FindNextMatch")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNextName() (Long, error) {
	return callComMethod[Long](d, "FindNextName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNthMatch(n Integer) (Boolean, error) {
	return callComMethod[Boolean](d, "FindNthMatch", n)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHNAME_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) FindNthName(n Integer) (Long, error) {
	return callComMethod[Long](d, "FindNthName", n)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetFirstItemValue() (any, error) {
	return callComAnyMethod(d, "GetFirstItemValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetNextItemValue() (any, error) {
	return callComAnyMethod(d, "GetNextItemValue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func (d NotesDirectoryNavigator) GetNthItemValue(n Integer) (any, error) {
	return callComAnyMethod(d, "GetNthItemValue", n)
}
