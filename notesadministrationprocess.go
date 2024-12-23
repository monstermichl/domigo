/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESADMINISTRATIONPROCESS_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesAdministrationProcess struct {
	NotesStruct
}

func NewNotesAdministrationProcess(dispatchPtr *ole.IDispatch) NotesAdministrationProcess {
	return NotesAdministrationProcess{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEAUTHORITYORG_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertificateAuthorityOrg() (String, error) {
	val, err := getComProperty(a, "CertificateAuthorityOrg")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEAUTHORITYORG_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertificateAuthorityOrg(v String) error {
	return putComProperty(a, "CertificateAuthorityOrg", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEEXPIRATION_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertificateExpiration() (NotesDateTime, error) {
	dispatchPtr, err := getComObjectProperty(a, "CertificateExpiration")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEEXPIRATION_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertificateExpiration(v NotesDateTime) error {
	return putComProperty(a, "CertificateExpiration", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERFILE_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertifierFile() (String, error) {
	val, err := getComProperty(a, "CertifierFile")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERFILE_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertifierFile(v String) error {
	return putComProperty(a, "CertifierFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERPASSWORD_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) CertifierPassword() (String, error) {
	val, err := getComProperty(a, "CertifierPassword")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERPASSWORD_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetCertifierPassword(v String) error {
	return putComProperty(a, "CertifierPassword", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCERTIFICATEAUTHORITYAVAILABLE_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) IsCertificateAuthorityAvailable() (Boolean, error) {
	val, err := getComProperty(a, "IsCertificateAuthorityAvailable")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) UseCertificateAuthority() (Boolean, error) {
	val, err := getComProperty(a, "UseCertificateAuthority")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_ADMINP.html */
func (a NotesAdministrationProcess) SetUseCertificateAuthority(v Boolean) error {
	return putComProperty(a, "UseCertificateAuthority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDGROUPMEMBERS_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) AddGroupMembers(group String, members []String) (String, error) {
	val, err := callComMethod(a, "AddGroupMembers", group, members)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDINTERNETCERTIFICATETOUSER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) AddInternetCertificateToUser(user String, keyringfile String, keyringpassword String, expiration NotesDateTime) (String, error) {
	val, err := callComMethod(a, "AddInternetCertificateToUser", user, keyringfile, keyringpassword, expiration.com().Dispatch())
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDSERVERTOCLUSTER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) AddServerToCluster(server String, cluster String) (String, error) {
	val, err := callComMethod(a, "AddServerToCluster", server, cluster)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDELETEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveDeletePersonInDirectory(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveDeletePersonInDirectory", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDELETEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveDeleteServerInDirectory(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveDeleteServerInDirectory", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDESIGNELEMENTDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveDesignElementDeletion(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveDesignElementDeletion", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEHOSTEDORGSTORAGEDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveHostedOrgStorageDeletion(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveHostedOrgStorageDeletion", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEMAILFILEDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveMailFileDeletion(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveMailFileDeletion", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEMOVEDREPLICADELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveMovedReplicaDeletion(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveMovedReplicaDeletion", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVENAMECHANGERETRACTION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveNameChangeRetraction(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveNameChangeRetraction", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERENAMEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveRenamePersonInDirectory(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveRenamePersonInDirectory", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERENAMEPERSONINDIRECTORY_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveRenameServerInDirectory(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveRenameServerInDirectory", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEREPLICADELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveReplicaDeletion(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveReplicaDeletion", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERESOURCEDELETION_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ApproveResourceDeletion(noteID String) (String, error) {
	val, err := callComMethod(a, "ApproveResourceDeletion", noteID)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHANGEHTTPPASSWORD_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) ChangeHTTPPassword(username String, oldpassword String, newpassword String) (String, error) {
	val, err := callComMethod(a, "ChangeHTTPPassword", username, oldpassword, newpassword)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONFIGUREMAILAGENT_METHOD_ADMINP.html */
type notesAdministrationProcessConfigureMailAgentParams struct {
	activatable *Boolean
	enable      *Boolean
}

type notesAdministrationProcessConfigureMailAgentParam func(*notesAdministrationProcessConfigureMailAgentParams)

func WithNotesAdministrationProcessConfigureMailAgentActivatable(activatable Boolean) notesAdministrationProcessConfigureMailAgentParam {
	return func(c *notesAdministrationProcessConfigureMailAgentParams) {
		c.activatable = &activatable
	}
}

func WithNotesAdministrationProcessConfigureMailAgentEnable(enable Boolean) notesAdministrationProcessConfigureMailAgentParam {
	return func(c *notesAdministrationProcessConfigureMailAgentParams) {
		c.enable = &enable
	}
}

func (a NotesAdministrationProcess) ConfigureMailAgent(username String, agentname String, params ...notesAdministrationProcessConfigureMailAgentParam) (String, error) {
	paramsStruct := &notesAdministrationProcessConfigureMailAgentParams{}
	paramsOrdered := []interface{}{username, agentname}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.activatable != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.activatable)
		if paramsStruct.enable != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.enable)
		}
	}
	val, err := callComMethod(a, "ConfigureMailAgent", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD_ADMINP.html */
type notesAdministrationProcessCreateReplicaParams struct {
	destdbfile    *String
	copyacl       *Boolean
	createftindex *Boolean
}

type notesAdministrationProcessCreateReplicaParam func(*notesAdministrationProcessCreateReplicaParams)

func WithNotesAdministrationProcessCreateReplicaDestdbfile(destdbfile String) notesAdministrationProcessCreateReplicaParam {
	return func(c *notesAdministrationProcessCreateReplicaParams) {
		c.destdbfile = &destdbfile
	}
}

func WithNotesAdministrationProcessCreateReplicaCopyacl(copyacl Boolean) notesAdministrationProcessCreateReplicaParam {
	return func(c *notesAdministrationProcessCreateReplicaParams) {
		c.copyacl = &copyacl
	}
}

func WithNotesAdministrationProcessCreateReplicaCreateftindex(createftindex Boolean) notesAdministrationProcessCreateReplicaParam {
	return func(c *notesAdministrationProcessCreateReplicaParams) {
		c.createftindex = &createftindex
	}
}

func (a NotesAdministrationProcess) CreateReplica(sourceserver String, sourcedbfile String, destserver String, params ...notesAdministrationProcessCreateReplicaParam) (String, error) {
	paramsStruct := &notesAdministrationProcessCreateReplicaParams{}
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
	val, err := callComMethod(a, "CreateReplica", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATEMAILFILE_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) Delegatemailfile(Mailfileowner String, Publicreaders []String, Publicwriters []String, Otherreaders Variant, Otherwriters []String, Othereditors []String, Otherdeleters []String, Removefromacl Variant, Mailfilename String, Mailserver String) (String, error) {
	val, err := callComMethod(a, "Delegatemailfile", Mailfileowner, Publicreaders, Publicwriters, Otherreaders, Otherwriters, Othereditors, Otherdeleters, Mailfileowner, Mailfileowner, Removefromacl, Mailfilename, Mailserver)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEGROUP_METHOD_ADMINP.html */
type notesAdministrationProcessDeleteGroupParams struct {
	deletewindowsgroup *Boolean
}

type notesAdministrationProcessDeleteGroupParam func(*notesAdministrationProcessDeleteGroupParams)

func WithNotesAdministrationProcessDeleteGroupDeletewindowsgroup(deletewindowsgroup Boolean) notesAdministrationProcessDeleteGroupParam {
	return func(c *notesAdministrationProcessDeleteGroupParams) {
		c.deletewindowsgroup = &deletewindowsgroup
	}
}

func (a NotesAdministrationProcess) DeleteGroup(groupname String, immediate Boolean, params ...notesAdministrationProcessDeleteGroupParam) error {
	paramsStruct := &notesAdministrationProcessDeleteGroupParams{}
	paramsOrdered := []interface{}{groupname, immediate}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.deletewindowsgroup != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.deletewindowsgroup)
	}
	_, err := callComMethod(a, "DeleteGroup", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEREPLICAS_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) DeleteReplicas(servername String, filename String) (String, error) {
	val, err := callComMethod(a, "DeleteReplicas", servername, filename)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETESERVER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) DeleteServer(servername String, immediate Boolean) (String, error) {
	val, err := callComMethod(a, "DeleteServer", servername, immediate)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEUSER_METHOD_ADMINP.html */
type notesAdministrationProcessDeleteUserParams struct {
	idvaultaction     *Integer
	deletewindowsuser *Boolean
}

type notesAdministrationProcessDeleteUserParam func(*notesAdministrationProcessDeleteUserParams)

func WithNotesAdministrationProcessDeleteUserIdvaultaction(idvaultaction Integer) notesAdministrationProcessDeleteUserParam {
	return func(c *notesAdministrationProcessDeleteUserParams) {
		c.idvaultaction = &idvaultaction
	}
}

func WithNotesAdministrationProcessDeleteUserDeletewindowsuser(deletewindowsuser Boolean) notesAdministrationProcessDeleteUserParam {
	return func(c *notesAdministrationProcessDeleteUserParams) {
		c.deletewindowsuser = &deletewindowsuser
	}
}

func (a NotesAdministrationProcess) DeleteUser(username String, immediate Boolean, mailfileaction Integer, denygroup String, params ...notesAdministrationProcessDeleteUserParam) (String, error) {
	paramsStruct := &notesAdministrationProcessDeleteUserParams{}
	paramsOrdered := []interface{}{username, immediate, mailfileaction, denygroup}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.idvaultaction != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.idvaultaction)
		if paramsStruct.deletewindowsuser != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.deletewindowsuser)
		}
	}
	val, err := callComMethod(a, "DeleteUser", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDGROUPINDOMAIN_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) FindGroupInDomain(group String) (String, error) {
	val, err := callComMethod(a, "FindGroupInDomain", group)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDSERVERINDOMAIN_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) FindServerInDomain(server String) (String, error) {
	val, err := callComMethod(a, "FindServerInDomain", server)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDUSERINDOMAIN_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) FindUserInDomain(user String) (String, error) {
	val, err := callComMethod(a, "FindUserInDomain", user)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEMAILUSER_METHOD_ADMINP.html */
type notesAdministrationProcessMoveMailUserParams struct {
	usescos                  *Boolean
	newclusterreplicaarray   *[]String
	deleteoldclusterreplicas *Boolean
}

type notesAdministrationProcessMoveMailUserParam func(*notesAdministrationProcessMoveMailUserParams)

func WithNotesAdministrationProcessMoveMailUserUsescos(usescos Boolean) notesAdministrationProcessMoveMailUserParam {
	return func(c *notesAdministrationProcessMoveMailUserParams) {
		c.usescos = &usescos
	}
}

func WithNotesAdministrationProcessMoveMailUserNewclusterreplicaarray(newclusterreplicaarray []String) notesAdministrationProcessMoveMailUserParam {
	return func(c *notesAdministrationProcessMoveMailUserParams) {
		c.newclusterreplicaarray = &newclusterreplicaarray
	}
}

func WithNotesAdministrationProcessMoveMailUserDeleteoldclusterreplicas(deleteoldclusterreplicas Boolean) notesAdministrationProcessMoveMailUserParam {
	return func(c *notesAdministrationProcessMoveMailUserParams) {
		c.deleteoldclusterreplicas = &deleteoldclusterreplicas
	}
}

func (a NotesAdministrationProcess) MoveMailUser(username String, newhomeserver String, newhomeservermailpath String, params ...notesAdministrationProcessMoveMailUserParam) (String, error) {
	paramsStruct := &notesAdministrationProcessMoveMailUserParams{}
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
	val, err := callComMethod(a, "MoveMailUser", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEREPLICA_METHOD_ADMINP.html */
type notesAdministrationProcessMoveReplicaParams struct {
	destdbfile    *String
	copyacl       *Boolean
	createftindex *Boolean
}

type notesAdministrationProcessMoveReplicaParam func(*notesAdministrationProcessMoveReplicaParams)

func WithNotesAdministrationProcessMoveReplicaDestdbfile(destdbfile String) notesAdministrationProcessMoveReplicaParam {
	return func(c *notesAdministrationProcessMoveReplicaParams) {
		c.destdbfile = &destdbfile
	}
}

func WithNotesAdministrationProcessMoveReplicaCopyacl(copyacl Boolean) notesAdministrationProcessMoveReplicaParam {
	return func(c *notesAdministrationProcessMoveReplicaParams) {
		c.copyacl = &copyacl
	}
}

func WithNotesAdministrationProcessMoveReplicaCreateftindex(createftindex Boolean) notesAdministrationProcessMoveReplicaParam {
	return func(c *notesAdministrationProcessMoveReplicaParams) {
		c.createftindex = &createftindex
	}
}

func (a NotesAdministrationProcess) MoveReplica(sourceserver String, sourcedbfile String, destserver String, params ...notesAdministrationProcessMoveReplicaParam) (String, error) {
	paramsStruct := &notesAdministrationProcessMoveReplicaParams{}
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
	val, err := callComMethod(a, "MoveReplica", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEROAMINGUSER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) MoveRoamingUser(username String, destserver String, destserverpath String) (String, error) {
	val, err := callComMethod(a, "MoveRoamingUser", username, destserver, destserverpath)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEUSERINHIERARCHYCOMPLETE_METHOD_ADMINP.html */
type notesAdministrationProcessMoveUserInHierarchyCompleteParams struct {
	lastname          *String
	firstname         *String
	middleinitial     *String
	orgunit           *String
	altcommonname     *String
	altorgunit        *String
	altlanguage       *String
	renamewindowsuser *Boolean
}

type notesAdministrationProcessMoveUserInHierarchyCompleteParam func(*notesAdministrationProcessMoveUserInHierarchyCompleteParams)

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteLastname(lastname String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.lastname = &lastname
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteFirstname(firstname String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.firstname = &firstname
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteMiddleinitial(middleinitial String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.middleinitial = &middleinitial
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteOrgunit(orgunit String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.orgunit = &orgunit
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteAltcommonname(altcommonname String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.altcommonname = &altcommonname
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteAltorgunit(altorgunit String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.altorgunit = &altorgunit
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteAltlanguage(altlanguage String) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.altlanguage = &altlanguage
	}
}

func WithNotesAdministrationProcessMoveUserInHierarchyCompleteRenamewindowsuser(renamewindowsuser Boolean) notesAdministrationProcessMoveUserInHierarchyCompleteParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyCompleteParams) {
		c.renamewindowsuser = &renamewindowsuser
	}
}

func (a NotesAdministrationProcess) MoveUserInHierarchyComplete(requestnoteid String, params ...notesAdministrationProcessMoveUserInHierarchyCompleteParam) (String, error) {
	paramsStruct := &notesAdministrationProcessMoveUserInHierarchyCompleteParams{}
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
	val, err := callComMethod(a, "MoveUserInHierarchyComplete", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEUSERINHIERARCHYREQUEST_METHOD_ADMINP.html */
type notesAdministrationProcessMoveUserInHierarchyRequestParams struct {
	allowprimarynamechange *Boolean
}

type notesAdministrationProcessMoveUserInHierarchyRequestParam func(*notesAdministrationProcessMoveUserInHierarchyRequestParams)

func WithNotesAdministrationProcessMoveUserInHierarchyRequestAllowprimarynamechange(allowprimarynamechange Boolean) notesAdministrationProcessMoveUserInHierarchyRequestParam {
	return func(c *notesAdministrationProcessMoveUserInHierarchyRequestParams) {
		c.allowprimarynamechange = &allowprimarynamechange
	}
}

func (a NotesAdministrationProcess) MoveUserInHierarchyRequest(username String, targetcertifier String, params ...notesAdministrationProcessMoveUserInHierarchyRequestParam) (String, error) {
	paramsStruct := &notesAdministrationProcessMoveUserInHierarchyRequestParams{}
	paramsOrdered := []interface{}{username, targetcertifier}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.allowprimarynamechange != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.allowprimarynamechange)
	}
	val, err := callComMethod(a, "MoveUserInHierarchyRequest", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFYSERVER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RecertifyServer(server String) (String, error) {
	val, err := callComMethod(a, "RecertifyServer", server)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFYUSER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RecertifyUser(username String) (String, error) {
	val, err := callComMethod(a, "RecertifyUser", username)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVESERVERFROMCLUSTER_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RemoveServerFromCluster(server String) error {
	_, err := callComMethod(a, "RemoveServerFromCluster", server)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEGROUP_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) RenameGroup(group String, newgroup String) (String, error) {
	val, err := callComMethod(a, "RenameGroup", group, newgroup)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMENOTESUSER_METHOD_ADMINP.html */
type notesAdministrationProcessRenameNotesUserParams struct {
	lastname          *String
	firstname         *String
	middleinitial     *String
	orgunit           *String
	altcommonname     *String
	altorgunit        *String
	altlanguage       *String
	renamewindowsuser *Boolean
}

type notesAdministrationProcessRenameNotesUserParam func(*notesAdministrationProcessRenameNotesUserParams)

func WithNotesAdministrationProcessRenameNotesUserLastname(lastname String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.lastname = &lastname
	}
}

func WithNotesAdministrationProcessRenameNotesUserFirstname(firstname String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.firstname = &firstname
	}
}

func WithNotesAdministrationProcessRenameNotesUserMiddleinitial(middleinitial String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.middleinitial = &middleinitial
	}
}

func WithNotesAdministrationProcessRenameNotesUserOrgunit(orgunit String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.orgunit = &orgunit
	}
}

func WithNotesAdministrationProcessRenameNotesUserAltcommonname(altcommonname String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.altcommonname = &altcommonname
	}
}

func WithNotesAdministrationProcessRenameNotesUserAltorgunit(altorgunit String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.altorgunit = &altorgunit
	}
}

func WithNotesAdministrationProcessRenameNotesUserAltlanguage(altlanguage String) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.altlanguage = &altlanguage
	}
}

func WithNotesAdministrationProcessRenameNotesUserRenamewindowsuser(renamewindowsuser Boolean) notesAdministrationProcessRenameNotesUserParam {
	return func(c *notesAdministrationProcessRenameNotesUserParams) {
		c.renamewindowsuser = &renamewindowsuser
	}
}

func (a NotesAdministrationProcess) RenameNotesUser(username String, params ...notesAdministrationProcessRenameNotesUserParam) (String, error) {
	paramsStruct := &notesAdministrationProcessRenameNotesUserParams{}
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
	val, err := callComMethod(a, "RenameNotesUser", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEWEBUSER_METHOD_ADMINP.html */
type notesAdministrationProcessRenameWebUserParams struct {
	newlastname        *String
	newfirstname       *String
	newmiddleinitial   *String
	newshortname       *String
	newinternetaddress *String
}

type notesAdministrationProcessRenameWebUserParam func(*notesAdministrationProcessRenameWebUserParams)

func WithNotesAdministrationProcessRenameWebUserNewlastname(newlastname String) notesAdministrationProcessRenameWebUserParam {
	return func(c *notesAdministrationProcessRenameWebUserParams) {
		c.newlastname = &newlastname
	}
}

func WithNotesAdministrationProcessRenameWebUserNewfirstname(newfirstname String) notesAdministrationProcessRenameWebUserParam {
	return func(c *notesAdministrationProcessRenameWebUserParams) {
		c.newfirstname = &newfirstname
	}
}

func WithNotesAdministrationProcessRenameWebUserNewmiddleinitial(newmiddleinitial String) notesAdministrationProcessRenameWebUserParam {
	return func(c *notesAdministrationProcessRenameWebUserParams) {
		c.newmiddleinitial = &newmiddleinitial
	}
}

func WithNotesAdministrationProcessRenameWebUserNewshortname(newshortname String) notesAdministrationProcessRenameWebUserParam {
	return func(c *notesAdministrationProcessRenameWebUserParams) {
		c.newshortname = &newshortname
	}
}

func WithNotesAdministrationProcessRenameWebUserNewinternetaddress(newinternetaddress String) notesAdministrationProcessRenameWebUserParam {
	return func(c *notesAdministrationProcessRenameWebUserParams) {
		c.newinternetaddress = &newinternetaddress
	}
}

func (a NotesAdministrationProcess) RenameWebUser(username String, newfullname String, params ...notesAdministrationProcessRenameWebUserParam) (String, error) {
	paramsStruct := &notesAdministrationProcessRenameWebUserParams{}
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
	val, err := callComMethod(a, "RenameWebUser", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETENABLEOUTLOOKSUPPORT_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) SetEnableOutlookSupport(Servername String) (String, error) {
	val, err := callComMethod(a, "SetEnableOutlookSupport", Servername)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSERVERDIRECTORYASSISTANCESETTINGS_METHOD_ADMINP.html */
func (a NotesAdministrationProcess) SetServerDirectoryAssistanceSettings(server String, dbfile String) (String, error) {
	val, err := callComMethod(a, "SetServerDirectoryAssistanceSettings", server, dbfile)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERPASSWORDSETTINGS_METHOD_ADMINP.html */
type notesAdministrationProcessSetUserPasswordSettingsParams struct {
	notespasswordchecksetting   *Integer
	notespasswordchangeinterval *Integer
	notespasswordgraceperiod    *Integer
	internetpasswordforcechange *Boolean
}

type notesAdministrationProcessSetUserPasswordSettingsParam func(*notesAdministrationProcessSetUserPasswordSettingsParams)

func WithNotesAdministrationProcessSetUserPasswordSettingsNotespasswordchecksetting(notespasswordchecksetting Integer) notesAdministrationProcessSetUserPasswordSettingsParam {
	return func(c *notesAdministrationProcessSetUserPasswordSettingsParams) {
		c.notespasswordchecksetting = &notespasswordchecksetting
	}
}

func WithNotesAdministrationProcessSetUserPasswordSettingsNotespasswordchangeinterval(notespasswordchangeinterval Integer) notesAdministrationProcessSetUserPasswordSettingsParam {
	return func(c *notesAdministrationProcessSetUserPasswordSettingsParams) {
		c.notespasswordchangeinterval = &notespasswordchangeinterval
	}
}

func WithNotesAdministrationProcessSetUserPasswordSettingsNotespasswordgraceperiod(notespasswordgraceperiod Integer) notesAdministrationProcessSetUserPasswordSettingsParam {
	return func(c *notesAdministrationProcessSetUserPasswordSettingsParams) {
		c.notespasswordgraceperiod = &notespasswordgraceperiod
	}
}

func WithNotesAdministrationProcessSetUserPasswordSettingsInternetpasswordforcechange(internetpasswordforcechange Boolean) notesAdministrationProcessSetUserPasswordSettingsParam {
	return func(c *notesAdministrationProcessSetUserPasswordSettingsParams) {
		c.internetpasswordforcechange = &internetpasswordforcechange
	}
}

func (a NotesAdministrationProcess) SetUserPasswordSettings(username String, params ...notesAdministrationProcessSetUserPasswordSettingsParam) (String, error) {
	paramsStruct := &notesAdministrationProcessSetUserPasswordSettingsParams{}
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
	val, err := callComMethod(a, "SetUserPasswordSettings", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNDATABASEWITHSERVERID_METHOD_ADMINP.html */
type notesAdministrationProcessSignDatabaseWithServerIDParams struct {
	updateonly *Boolean
}

type notesAdministrationProcessSignDatabaseWithServerIDParam func(*notesAdministrationProcessSignDatabaseWithServerIDParams)

func WithNotesAdministrationProcessSignDatabaseWithServerIDUpdateonly(updateonly Boolean) notesAdministrationProcessSignDatabaseWithServerIDParam {
	return func(c *notesAdministrationProcessSignDatabaseWithServerIDParams) {
		c.updateonly = &updateonly
	}
}

func (a NotesAdministrationProcess) SignDatabaseWithServerID(server String, dbfile String, params ...notesAdministrationProcessSignDatabaseWithServerIDParam) (String, error) {
	paramsStruct := &notesAdministrationProcessSignDatabaseWithServerIDParams{}
	paramsOrdered := []interface{}{server, dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.updateonly != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.updateonly)
	}
	val, err := callComMethod(a, "SignDatabaseWithServerID", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPGRADEUSERTOHIERARCHICAL_METHOD_ADMINP.html */
type notesAdministrationProcessUpgradeUserToHierarchicalParams struct {
	orgunit       *String
	altcommonname *String
	altorgunit    *String
	altlanguage   *String
}

type notesAdministrationProcessUpgradeUserToHierarchicalParam func(*notesAdministrationProcessUpgradeUserToHierarchicalParams)

func WithNotesAdministrationProcessUpgradeUserToHierarchicalOrgunit(orgunit String) notesAdministrationProcessUpgradeUserToHierarchicalParam {
	return func(c *notesAdministrationProcessUpgradeUserToHierarchicalParams) {
		c.orgunit = &orgunit
	}
}

func WithNotesAdministrationProcessUpgradeUserToHierarchicalAltcommonname(altcommonname String) notesAdministrationProcessUpgradeUserToHierarchicalParam {
	return func(c *notesAdministrationProcessUpgradeUserToHierarchicalParams) {
		c.altcommonname = &altcommonname
	}
}

func WithNotesAdministrationProcessUpgradeUserToHierarchicalAltorgunit(altorgunit String) notesAdministrationProcessUpgradeUserToHierarchicalParam {
	return func(c *notesAdministrationProcessUpgradeUserToHierarchicalParams) {
		c.altorgunit = &altorgunit
	}
}

func WithNotesAdministrationProcessUpgradeUserToHierarchicalAltlanguage(altlanguage String) notesAdministrationProcessUpgradeUserToHierarchicalParam {
	return func(c *notesAdministrationProcessUpgradeUserToHierarchicalParams) {
		c.altlanguage = &altlanguage
	}
}

func (a NotesAdministrationProcess) UpgradeUserToHierarchical(username String, params ...notesAdministrationProcessUpgradeUserToHierarchicalParam) (String, error) {
	paramsStruct := &notesAdministrationProcessUpgradeUserToHierarchicalParams{}
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
	val, err := callComMethod(a, "UpgradeUserToHierarchical", paramsOrdered...)
	return helpers.CastValue[String](val), err
}
