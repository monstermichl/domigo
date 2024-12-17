/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESADMINISTRATIONPROCESS_CLASS.html */
package notesadministrationprocess

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdatetime"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesAdministrationProcess struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesAdministrationProcess {
	return NotesAdministrationProcess{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEAUTHORITYORG_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertificateAuthorityOrg() (domino.String, error) {
	val, err := a.Com().GetProperty("CertificateAuthorityOrg")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEAUTHORITYORG_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertificateAuthorityOrg(v domino.String) error {
	return a.Com().PutProperty("CertificateAuthorityOrg", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEEXPIRATION_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertificateExpiration() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := a.Com().GetObjectProperty("CertificateExpiration")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEEXPIRATION_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertificateExpiration(v notesdatetime.NotesDateTime) error {
	return a.Com().PutProperty("CertificateExpiration", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERFILE_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertifierFile() (domino.String, error) {
	val, err := a.Com().GetProperty("CertifierFile")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERFILE_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertifierFile(v domino.String) error {
	return a.Com().PutProperty("CertifierFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERPASSWORD_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertifierPassword() (domino.String, error) {
	val, err := a.Com().GetProperty("CertifierPassword")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERPASSWORD_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertifierPassword(v domino.String) error {
	return a.Com().PutProperty("CertifierPassword", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCERTIFICATEAUTHORITYAVAILABLE_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) IsCertificateAuthorityAvailable() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsCertificateAuthorityAvailable")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) UseCertificateAuthority() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("UseCertificateAuthority")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetUseCertificateAuthority(v domino.Boolean) error {
	return a.Com().PutProperty("UseCertificateAuthority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDGROUPMEMBERS_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) AddGroupMembers(group domino.String, members []domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("AddGroupMembers", group, members)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDINTERNETCERTIFICATETOUSER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) AddInternetCertificateToUser(user domino.String, keyringfile domino.String, keyringpassword domino.String, expiration notesdatetime.NotesDateTime) (domino.String, error) {
	val, err := a.Com().CallMethod("AddInternetCertificateToUser", user, keyringfile, keyringpassword, expiration.Com().Dispatch())
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDSERVERTOCLUSTER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) AddServerToCluster(server domino.String, cluster domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("AddServerToCluster", server, cluster)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDELETEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveDeletePersonInDirectory(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveDeletePersonInDirectory", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDELETEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveDeleteServerInDirectory(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveDeleteServerInDirectory", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDESIGNELEMENTDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveDesignElementDeletion(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveDesignElementDeletion", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEHOSTEDORGSTORAGEDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveHostedOrgStorageDeletion(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveHostedOrgStorageDeletion", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEMAILFILEDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveMailFileDeletion(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveMailFileDeletion", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEMOVEDREPLICADELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveMovedReplicaDeletion(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveMovedReplicaDeletion", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVENAMECHANGERETRACTION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveNameChangeRetraction(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveNameChangeRetraction", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERENAMEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveRenamePersonInDirectory(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveRenamePersonInDirectory", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERENAMEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveRenameServerInDirectory(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveRenameServerInDirectory", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEREPLICADELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveReplicaDeletion(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveReplicaDeletion", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERESOURCEDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveResourceDeletion(noteID domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ApproveResourceDeletion", noteID)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHANGEHTTPPASSWORD_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ChangeHTTPPassword(username domino.String, oldpassword domino.String, newpassword domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("ChangeHTTPPassword", username, oldpassword, newpassword)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONFIGUREMAILAGENT_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ConfigureMailAgent(username domino.String, agentname domino.String, activatable domino.Boolean, enable domino.Boolean) (domino.String, error) {
	val, err := a.Com().CallMethod("ConfigureMailAgent", username, agentname, activatable, enable)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD_ADMINP.html */
type createReplicaParams struct {
	destdbfile    *domino.String
	copyacl       *domino.Boolean
	createftindex *domino.Boolean
}

type createReplicaParam func(*createReplicaParams)

func WithCreateReplicaDestdbfile(destdbfile domino.String) createReplicaParam {
	return func(c *createReplicaParams) {
		c.destdbfile = &destdbfile
	}
}

func WithCreateReplicaCopyacl(copyacl domino.Boolean) createReplicaParam {
	return func(c *createReplicaParams) {
		c.copyacl = &copyacl
	}
}

func WithCreateReplicaCreateftindex(createftindex domino.Boolean) createReplicaParam {
	return func(c *createReplicaParams) {
		c.createftindex = &createftindex
	}
}

func (a NotesAdministrationProcess) CreateReplica(sourceserver domino.String, sourcedbfile domino.String, destserver domino.String, params ...createReplicaParam) (domino.String, error) {
	paramsStruct := &createReplicaParams{}
	paramsOrdered := []interface{}{sourceserver, sourcedbfile, destserver}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.destdbfile != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.destdbfile)
		if paramsStruct.copyacl != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.copyacl)
			if paramsStruct.createftindex != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.createftindex)
			}
		}
	}
	val, err := a.Com().CallMethod("CreateReplica", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATEMAILFILE_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) DelegateMailFile(mailFileOwner domino.String, publicReaders []domino.String, publicWriters []domino.String, otherReaders []domino.String, otherWriters []domino.String, otherEditors []domino.String, otherDeleters []domino.String, RemoveFromACL domino.String, mailFileName domino.String, mailServer domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("DelegateMailFile", mailFileOwner, publicReaders, publicWriters, otherReaders, otherWriters, otherEditors, otherDeleters, RemoveFromACL, mailFileName, mailServer)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEGROUP_METHOD_ADMINP.html */
type deleteGroupParams struct {
	deleteWindowsGroup *domino.Boolean
}

type deleteGroupParam func(*deleteGroupParams)

func WithDeleteGroupDeleteWindowsGroup(deleteWindowsGroup domino.Boolean) deleteGroupParam {
	return func(c *deleteGroupParams) {
		c.deleteWindowsGroup = &deleteWindowsGroup
	}
}

func (a NotesAdministrationProcess) DeleteGroup(groupName domino.String, immediate domino.Boolean, params ...deleteGroupParam) error {
	paramsStruct := &deleteGroupParams{}
	paramsOrdered := []interface{}{groupName, immediate}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.deleteWindowsGroup != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.deleteWindowsGroup)
	}
	_, err := a.Com().CallMethod("DeleteGroup", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEREPLICAS_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) DeleteReplicas(servername domino.String, filename domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("DeleteReplicas", servername, filename)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETESERVER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) DeleteServer(serverName domino.String, immediate domino.Boolean) (domino.String, error) {
	val, err := a.Com().CallMethod("DeleteServer", serverName, immediate)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEUSER_METHOD_ADMINP.html */
type deleteUserParams struct {
	deleteWindowsUser *domino.Boolean
}

type deleteUserParam func(*deleteUserParams)

func WithDeleteUserDeletewindowsuser(deleteWindowsUser domino.Boolean) deleteUserParam {
	return func(c *deleteUserParams) {
		c.deleteWindowsUser = &deleteWindowsUser
	}
}

func (a NotesAdministrationProcess) DeleteUser(username domino.String, immediate domino.Boolean, mailFileAction domino.Integer, idVaultAction domino.Integer, denyGroup domino.String, params ...deleteUserParam) (domino.String, error) {
	paramsStruct := &deleteUserParams{}
	paramsOrdered := []interface{}{username, immediate, mailFileAction, idVaultAction, denyGroup}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.deleteWindowsUser != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.deleteWindowsUser)
	}
	val, err := a.Com().CallMethod("DeleteUser", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDGROUPINDOMAIN_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) FindGroupInDomain(group domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("FindGroupInDomain", group)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDSERVERINDOMAIN_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) FindServerInDomain(server domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("FindServerInDomain", server)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDUSERINDOMAIN_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) FindUserInDomain(user domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("FindUserInDomain", user)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEMAILUSER_METHOD_ADMINP.html */
type moveMailUserParams struct {
	usescos                  *domino.Boolean
	newclusterreplicaarray   *[]domino.String
	deleteoldclusterreplicas *domino.Boolean
}

type moveMailUserParam func(*moveMailUserParams)

func WithMoveMailUserUsescos(usescos domino.Boolean) moveMailUserParam {
	return func(c *moveMailUserParams) {
		c.usescos = &usescos
	}
}

func WithMoveMailUserNewclusterreplicaarray(newclusterreplicaarray []domino.String) moveMailUserParam {
	return func(c *moveMailUserParams) {
		c.newclusterreplicaarray = &newclusterreplicaarray
	}
}

func WithMoveMailUserDeleteoldclusterreplicas(deleteoldclusterreplicas domino.Boolean) moveMailUserParam {
	return func(c *moveMailUserParams) {
		c.deleteoldclusterreplicas = &deleteoldclusterreplicas
	}
}

func (a NotesAdministrationProcess) MoveMailUser(username domino.String, newhomeserver domino.String, newhomeservermailpath domino.String, params ...moveMailUserParam) (domino.String, error) {
	paramsStruct := &moveMailUserParams{}
	paramsOrdered := []interface{}{username, newhomeserver, newhomeservermailpath}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.usescos != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.usescos)
		if paramsStruct.newclusterreplicaarray != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.newclusterreplicaarray)
			if paramsStruct.deleteoldclusterreplicas != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.deleteoldclusterreplicas)
			}
		}
	}
	val, err := a.Com().CallMethod("MoveMailUser", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEREPLICA_METHOD_ADMINP.html */
type moveReplicaParams struct {
	destdbfile    *domino.String
	copyacl       *domino.Boolean
	createftindex *domino.Boolean
}

type moveReplicaParam func(*moveReplicaParams)

func WithMoveReplicaDestdbfile(destdbfile domino.String) moveReplicaParam {
	return func(c *moveReplicaParams) {
		c.destdbfile = &destdbfile
	}
}

func WithMoveReplicaCopyacl(copyacl domino.Boolean) moveReplicaParam {
	return func(c *moveReplicaParams) {
		c.copyacl = &copyacl
	}
}

func WithMoveReplicaCreateftindex(createftindex domino.Boolean) moveReplicaParam {
	return func(c *moveReplicaParams) {
		c.createftindex = &createftindex
	}
}

func (a NotesAdministrationProcess) MoveReplica(sourceserver domino.String, sourcedbfile domino.String, destserver domino.String, params ...moveReplicaParam) (domino.String, error) {
	paramsStruct := &moveReplicaParams{}
	paramsOrdered := []interface{}{sourceserver, sourcedbfile, destserver}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.destdbfile != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.destdbfile)
		if paramsStruct.copyacl != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.copyacl)
			if paramsStruct.createftindex != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.createftindex)
			}
		}
	}
	val, err := a.Com().CallMethod("MoveReplica", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEROAMINGUSER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) MoveRoamingUser(username domino.String, destserver domino.String, destserverpath domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("MoveRoamingUser", username, destserver, destserverpath)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEUSERINHIERARCHYCOMPLETE_METHOD_ADMINP.html */
type moveUserInHierarchyCompleteParams struct {
	lastname          *domino.String
	firstname         *domino.String
	middleinitial     *domino.String
	orgunit           *domino.String
	altcommonname     *domino.String
	altorgunit        *domino.String
	altlanguage       *domino.String
	renamewindowsuser *domino.Boolean
}

type moveUserInHierarchyCompleteParam func(*moveUserInHierarchyCompleteParams)

func WithMoveUserInHierarchyCompleteLastname(lastname domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.lastname = &lastname
	}
}

func WithMoveUserInHierarchyCompleteFirstname(firstname domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.firstname = &firstname
	}
}

func WithMoveUserInHierarchyCompleteMiddleinitial(middleinitial domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.middleinitial = &middleinitial
	}
}

func WithMoveUserInHierarchyCompleteOrgunit(orgunit domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.orgunit = &orgunit
	}
}

func WithMoveUserInHierarchyCompleteAltcommonname(altcommonname domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.altcommonname = &altcommonname
	}
}

func WithMoveUserInHierarchyCompleteAltorgunit(altorgunit domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.altorgunit = &altorgunit
	}
}

func WithMoveUserInHierarchyCompleteAltlanguage(altlanguage domino.String) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.altlanguage = &altlanguage
	}
}

func WithMoveUserInHierarchyCompleteRenamewindowsuser(renamewindowsuser domino.Boolean) moveUserInHierarchyCompleteParam {
	return func(c *moveUserInHierarchyCompleteParams) {
		c.renamewindowsuser = &renamewindowsuser
	}
}

func (a NotesAdministrationProcess) MoveUserInHierarchyComplete(requestnoteid domino.String, params ...moveUserInHierarchyCompleteParam) (domino.String, error) {
	paramsStruct := &moveUserInHierarchyCompleteParams{}
	paramsOrdered := []interface{}{requestnoteid}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.lastname != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.lastname)
		if paramsStruct.firstname != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.firstname)
			if paramsStruct.middleinitial != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.middleinitial)
				if paramsStruct.orgunit != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.orgunit)
					if paramsStruct.altcommonname != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.altcommonname)
						if paramsStruct.altorgunit != nil {
							paramsOrdered = append(paramsOrdered, *paramsStruct.altorgunit)
							if paramsStruct.altlanguage != nil {
								paramsOrdered = append(paramsOrdered, *paramsStruct.altlanguage)
								if paramsStruct.renamewindowsuser != nil {
									paramsOrdered = append(paramsOrdered, *paramsStruct.renamewindowsuser)
								}
							}
						}
					}
				}
			}
		}
	}
	val, err := a.Com().CallMethod("MoveUserInHierarchyComplete", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEUSERINHIERARCHYREQUEST_METHOD_ADMINP.html */
type moveUserInHierarchyRequestParams struct {
	allowprimarynamechange *domino.Boolean
}

type moveUserInHierarchyRequestParam func(*moveUserInHierarchyRequestParams)

func WithMoveUserInHierarchyRequestAllowprimarynamechange(allowprimarynamechange domino.Boolean) moveUserInHierarchyRequestParam {
	return func(c *moveUserInHierarchyRequestParams) {
		c.allowprimarynamechange = &allowprimarynamechange
	}
}

func (a NotesAdministrationProcess) MoveUserInHierarchyRequest(username domino.String, targetcertifier domino.String, params ...moveUserInHierarchyRequestParam) (domino.String, error) {
	paramsStruct := &moveUserInHierarchyRequestParams{}
	paramsOrdered := []interface{}{username, targetcertifier}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.allowprimarynamechange != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.allowprimarynamechange)
	}
	val, err := a.Com().CallMethod("MoveUserInHierarchyRequest", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFYSERVER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RecertifyServer(server domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("RecertifyServer", server)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFYUSER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RecertifyUser(username domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("RecertifyUser", username)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVESERVERFROMCLUSTER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RemoveServerFromCluster(server domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("RemoveServerFromCluster", server)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEGROUP_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RenameGroup(group domino.String, newgroup domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("RenameGroup", group, newgroup)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMENOTESUSER_METHOD_ADMINP.html */
type renameNotesUserParams struct {
	lastname          *domino.String
	firstname         *domino.String
	middleinitial     *domino.String
	orgunit           *domino.String
	altcommonname     *domino.String
	altorgunit        *domino.String
	altlanguage       *domino.String
	renamewindowsuser *domino.Boolean
}

type renameNotesUserParam func(*renameNotesUserParams)

func WithRenameNotesUserLastname(lastname domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.lastname = &lastname
	}
}

func WithRenameNotesUserFirstname(firstname domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.firstname = &firstname
	}
}

func WithRenameNotesUserMiddleinitial(middleinitial domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.middleinitial = &middleinitial
	}
}

func WithRenameNotesUserOrgunit(orgunit domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.orgunit = &orgunit
	}
}

func WithRenameNotesUserAltcommonname(altcommonname domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.altcommonname = &altcommonname
	}
}

func WithRenameNotesUserAltorgunit(altorgunit domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.altorgunit = &altorgunit
	}
}

func WithRenameNotesUserAltlanguage(altlanguage domino.String) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.altlanguage = &altlanguage
	}
}

func WithRenameNotesUserRenamewindowsuser(renamewindowsuser domino.Boolean) renameNotesUserParam {
	return func(c *renameNotesUserParams) {
		c.renamewindowsuser = &renamewindowsuser
	}
}

func (a NotesAdministrationProcess) RenameNotesUser(username domino.String, params ...renameNotesUserParam) (domino.String, error) {
	paramsStruct := &renameNotesUserParams{}
	paramsOrdered := []interface{}{username}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.lastname != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.lastname)
		if paramsStruct.firstname != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.firstname)
			if paramsStruct.middleinitial != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.middleinitial)
				if paramsStruct.orgunit != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.orgunit)
					if paramsStruct.altcommonname != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.altcommonname)
						if paramsStruct.altorgunit != nil {
							paramsOrdered = append(paramsOrdered, *paramsStruct.altorgunit)
							if paramsStruct.altlanguage != nil {
								paramsOrdered = append(paramsOrdered, *paramsStruct.altlanguage)
								if paramsStruct.renamewindowsuser != nil {
									paramsOrdered = append(paramsOrdered, *paramsStruct.renamewindowsuser)
								}
							}
						}
					}
				}
			}
		}
	}
	val, err := a.Com().CallMethod("RenameNotesUser", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEWEBUSER_METHOD_ADMINP.html */
type renameWebUserParams struct {
	newlastname        *domino.String
	newfirstname       *domino.String
	newmiddleinitial   *domino.String
	newshortname       *domino.String
	newinternetaddress *domino.String
}

type renameWebUserParam func(*renameWebUserParams)

func WithRenameWebUserNewlastname(newlastname domino.String) renameWebUserParam {
	return func(c *renameWebUserParams) {
		c.newlastname = &newlastname
	}
}

func WithRenameWebUserNewfirstname(newfirstname domino.String) renameWebUserParam {
	return func(c *renameWebUserParams) {
		c.newfirstname = &newfirstname
	}
}

func WithRenameWebUserNewmiddleinitial(newmiddleinitial domino.String) renameWebUserParam {
	return func(c *renameWebUserParams) {
		c.newmiddleinitial = &newmiddleinitial
	}
}

func WithRenameWebUserNewshortname(newshortname domino.String) renameWebUserParam {
	return func(c *renameWebUserParams) {
		c.newshortname = &newshortname
	}
}

func WithRenameWebUserNewinternetaddress(newinternetaddress domino.String) renameWebUserParam {
	return func(c *renameWebUserParams) {
		c.newinternetaddress = &newinternetaddress
	}
}

func (a NotesAdministrationProcess) RenameWebUser(username domino.String, newfullname domino.String, params ...renameWebUserParam) (domino.String, error) {
	paramsStruct := &renameWebUserParams{}
	paramsOrdered := []interface{}{username, newfullname}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.newlastname != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.newlastname)
		if paramsStruct.newfirstname != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.newfirstname)
			if paramsStruct.newmiddleinitial != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.newmiddleinitial)
				if paramsStruct.newshortname != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.newshortname)
					if paramsStruct.newinternetaddress != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.newinternetaddress)
					}
				}
			}
		}
	}
	val, err := a.Com().CallMethod("RenameWebUser", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETENABLEOUTLOOKSUPPORT_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) SetEnableOutlookSupport(Servername domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("SetEnableOutlookSupport", Servername)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSERVERDIRECTORYASSISTANCESETTINGS_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) SetServerDirectoryAssistanceSettings(server domino.String, dbfile domino.String) (domino.String, error) {
	val, err := a.Com().CallMethod("SetServerDirectoryAssistanceSettings", server, dbfile)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERPASSWORDSETTINGS_METHOD_ADMINP.html */
type setUserPasswordSettingsParams struct {
	notespasswordchecksetting   *domino.Integer
	notespasswordchangeinterval *domino.Integer
	notespasswordgraceperiod    *domino.Integer
	internetpasswordforcechange *domino.Boolean
}

type setUserPasswordSettingsParam func(*setUserPasswordSettingsParams)

func WithSetUserPasswordSettingsNotespasswordchecksetting(notespasswordchecksetting domino.Integer) setUserPasswordSettingsParam {
	return func(c *setUserPasswordSettingsParams) {
		c.notespasswordchecksetting = &notespasswordchecksetting
	}
}

func WithSetUserPasswordSettingsNotespasswordchangeinterval(notespasswordchangeinterval domino.Integer) setUserPasswordSettingsParam {
	return func(c *setUserPasswordSettingsParams) {
		c.notespasswordchangeinterval = &notespasswordchangeinterval
	}
}

func WithSetUserPasswordSettingsNotespasswordgraceperiod(notespasswordgraceperiod domino.Integer) setUserPasswordSettingsParam {
	return func(c *setUserPasswordSettingsParams) {
		c.notespasswordgraceperiod = &notespasswordgraceperiod
	}
}

func WithSetUserPasswordSettingsInternetpasswordforcechange(internetpasswordforcechange domino.Boolean) setUserPasswordSettingsParam {
	return func(c *setUserPasswordSettingsParams) {
		c.internetpasswordforcechange = &internetpasswordforcechange
	}
}

func (a NotesAdministrationProcess) SetUserPasswordSettings(username domino.String, params ...setUserPasswordSettingsParam) (domino.String, error) {
	paramsStruct := &setUserPasswordSettingsParams{}
	paramsOrdered := []interface{}{username}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.notespasswordchecksetting != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.notespasswordchecksetting)
		if paramsStruct.notespasswordchangeinterval != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.notespasswordchangeinterval)
			if paramsStruct.notespasswordgraceperiod != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.notespasswordgraceperiod)
				if paramsStruct.internetpasswordforcechange != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.internetpasswordforcechange)
				}
			}
		}
	}
	val, err := a.Com().CallMethod("SetUserPasswordSettings", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNDATABASEWITHSERVERID_METHOD_ADMINP.html */
type signDatabaseWithServerIDParams struct {
	updateonly *domino.Boolean
}

type signDatabaseWithServerIDParam func(*signDatabaseWithServerIDParams)

func WithSignDatabaseWithServerIDUpdateonly(updateonly domino.Boolean) signDatabaseWithServerIDParam {
	return func(c *signDatabaseWithServerIDParams) {
		c.updateonly = &updateonly
	}
}

func (a NotesAdministrationProcess) SignDatabaseWithServerID(server domino.String, dbfile domino.String, params ...signDatabaseWithServerIDParam) (domino.String, error) {
	paramsStruct := &signDatabaseWithServerIDParams{}
	paramsOrdered := []interface{}{server, dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.updateonly != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.updateonly)
	}
	val, err := a.Com().CallMethod("SignDatabaseWithServerID", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPGRADEUSERTOHIERARCHICAL_METHOD_ADMINP.html */
type upgradeUserToHierarchicalParams struct {
	orgunit       *domino.String
	altcommonname *domino.String
	altorgunit    *domino.String
	altlanguage   *domino.String
}

type upgradeUserToHierarchicalParam func(*upgradeUserToHierarchicalParams)

func WithUpgradeUserToHierarchicalOrgunit(orgunit domino.String) upgradeUserToHierarchicalParam {
	return func(c *upgradeUserToHierarchicalParams) {
		c.orgunit = &orgunit
	}
}

func WithUpgradeUserToHierarchicalAltcommonname(altcommonname domino.String) upgradeUserToHierarchicalParam {
	return func(c *upgradeUserToHierarchicalParams) {
		c.altcommonname = &altcommonname
	}
}

func WithUpgradeUserToHierarchicalAltorgunit(altorgunit domino.String) upgradeUserToHierarchicalParam {
	return func(c *upgradeUserToHierarchicalParams) {
		c.altorgunit = &altorgunit
	}
}

func WithUpgradeUserToHierarchicalAltlanguage(altlanguage domino.String) upgradeUserToHierarchicalParam {
	return func(c *upgradeUserToHierarchicalParams) {
		c.altlanguage = &altlanguage
	}
}

func (a NotesAdministrationProcess) UpgradeUserToHierarchical(username domino.String, params ...upgradeUserToHierarchicalParam) (domino.String, error) {
	paramsStruct := &upgradeUserToHierarchicalParams{}
	paramsOrdered := []interface{}{username}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.orgunit != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.orgunit)
		if paramsStruct.altcommonname != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.altcommonname)
			if paramsStruct.altorgunit != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.altorgunit)
				if paramsStruct.altlanguage != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.altlanguage)
				}
			}
		}
	}
	val, err := a.Com().CallMethod("UpgradeUserToHierarchical", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}
