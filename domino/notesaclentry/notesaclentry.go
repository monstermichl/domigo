/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACLENTRY_CLASS.html */
package notesaclentry

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesname"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesACLEntry struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesACLEntry {
	return NotesACLEntry{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORNOCREATE_PROPERTY.html */
func (a NotesACLEntry) CanCreateDocuments() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanCreateDocuments")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORNOCREATE_PROPERTY.html */
func (a NotesACLEntry) SetCanCreateDocuments(v domino.Boolean) error {
	return a.Com().PutProperty("CanCreateDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATELSORJAVAAGENT_PROPERTY_6760.html */
func (a NotesACLEntry) CanCreateLSOrJavaAgent() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanCreateLSOrJavaAgent")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATELSORJAVAAGENT_PROPERTY_6760.html */
func (a NotesACLEntry) SetCanCreateLSOrJavaAgent(v domino.Boolean) error {
	return a.Com().PutProperty("CanCreateLSOrJavaAgent", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEAGENT_PROPERTY.html */
func (a NotesACLEntry) CanCreatePersonalAgent() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanCreatePersonalAgent")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEAGENT_PROPERTY.html */
func (a NotesACLEntry) SetCanCreatePersonalAgent(v domino.Boolean) error {
	return a.Com().PutProperty("CanCreatePersonalAgent", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEFOLDER_PROPERTY.html */
func (a NotesACLEntry) CanCreatePersonalFolder() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanCreatePersonalFolder")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEFOLDER_PROPERTY.html */
func (a NotesACLEntry) SetCanCreatePersonalFolder(v domino.Boolean) error {
	return a.Com().PutProperty("CanCreatePersonalFolder", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATESHAREDFOLDER_PROPERTY_7622.html */
func (a NotesACLEntry) CanCreateSharedFolder() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanCreateSharedFolder")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATESHAREDFOLDER_PROPERTY_7622.html */
func (a NotesACLEntry) SetCanCreateSharedFolder(v domino.Boolean) error {
	return a.Com().PutProperty("CanCreateSharedFolder", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNODELETE_PROPERTY.html */
func (a NotesACLEntry) CanDeleteDocuments() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanDeleteDocuments")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNODELETE_PROPERTY.html */
func (a NotesACLEntry) SetCanDeleteDocuments(v domino.Boolean) error {
	return a.Com().PutProperty("CanDeleteDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANREPLICATEORCOPYDOCUMENTS_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) CanReplicateOrCopyDocuments() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("CanReplicateOrCopyDocuments")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANREPLICATEORCOPYDOCUMENTS_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) SetCanReplicateOrCopyDocuments(v domino.Boolean) error {
	return a.Com().PutProperty("CanReplicateOrCopyDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_88.html */
func (a NotesACLEntry) IsAdminReaderAuthor() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsAdminReaderAuthor")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_88.html */
func (a NotesACLEntry) SetIsAdminReaderAuthor(v domino.Boolean) error {
	return a.Com().PutProperty("IsAdminReaderAuthor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINSERVER_PROPERTY_2946.html */
func (a NotesACLEntry) IsAdminServer() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsAdminServer")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINSERVER_PROPERTY_2946.html */
func (a NotesACLEntry) SetIsAdminServer(v domino.Boolean) error {
	return a.Com().PutProperty("IsAdminServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISGROUP_PROPERTY_4731.html */
func (a NotesACLEntry) IsGroup() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsGroup")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISGROUP_PROPERTY_4731.html */
func (a NotesACLEntry) SetIsGroup(v domino.Boolean) error {
	return a.Com().PutProperty("IsGroup", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPERSON_PROPERTY_651.html */
func (a NotesACLEntry) IsPerson() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsPerson")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPERSON_PROPERTY_651.html */
func (a NotesACLEntry) SetIsPerson(v domino.Boolean) error {
	return a.Com().PutProperty("IsPerson", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICREADER_PROPERTY.html */
func (a NotesACLEntry) IsPublicReader() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsPublicReader")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICREADER_PROPERTY.html */
func (a NotesACLEntry) SetIsPublicReader(v domino.Boolean) error {
	return a.Com().PutProperty("IsPublicReader", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICWRITER_PROPERTY.html */
func (a NotesACLEntry) IsPublicWriter() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsPublicWriter")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICWRITER_PROPERTY.html */
func (a NotesACLEntry) SetIsPublicWriter(v domino.Boolean) error {
	return a.Com().PutProperty("IsPublicWriter", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSERVER_PROPERTY_1888.html */
func (a NotesACLEntry) IsServer() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsServer")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSERVER_PROPERTY_1888.html */
func (a NotesACLEntry) SetIsServer(v domino.Boolean) error {
	return a.Com().PutProperty("IsServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY.html */
func (a NotesACLEntry) Level() (domino.Long, error) {
	val, err := a.Com().GetProperty("Level")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY.html */
func (a NotesACLEntry) SetLevel(v domino.Long) error {
	return a.Com().PutProperty("Level", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) Name() (domino.String, error) {
	val, err := a.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) SetName(v domino.String) error {
	return a.Com().PutProperty("Name", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOBJECT_PROPERTY.html */
func (a NotesACLEntry) NameObject() (notesname.NotesName, error) {
	dispatchPtr, err := a.Com().GetObjectProperty("NameObject")
	return notesname.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) Roles() ([]domino.String, error) {
	vals, err := a.Com().GetArrayProperty("Roles")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERTYPE_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) UserType() (domino.Long, error) {
	val, err := a.Com().GetProperty("UserType")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERTYPE_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) SetUserType(v domino.Long) error {
	return a.Com().PutProperty("UserType", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLEROLE_METHOD.html */
func (a NotesACLEntry) DisableRole(name domino.String) error {
	_, err := a.Com().CallMethod("DisableRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEROLE_METHOD.html */
func (a NotesACLEntry) EnableRole(name domino.String) error {
	_, err := a.Com().CallMethod("EnableRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROLEENABLED_METHOD.html */
func (a NotesACLEntry) IsRoleEnabled(name domino.String) (domino.Boolean, error) {
	val, err := a.Com().CallMethod("IsRoleEnabled", name)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ACLENTRY.html */
func (a NotesACLEntry) Remove() error {
	_, err := a.Com().CallMethod("Remove")
	return err
}
