/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORY_CLASS.html */
package notesdirectory_test

import (
	"domigo/domino/notesdirectory"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/require"
)

var directory notesdirectory.NotesDirectory

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	directory, _ = session.GetDirectory()

	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLEITEMS_PROPERTY_DIRECTORY.html */
func TestAvailableItems(t *testing.T) {
	_, err := directory.AvailableItems()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLENAMES_PROPERTY_DIRECTORY.html */
func TestAvailableNames(t *testing.T) {
	_, err := directory.AvailableNames()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AVAILABLEVIEW_PROPERTY_DIRECTORY.html */
func TestAvailableView(t *testing.T) {
	_, err := directory.AvailableView()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPAUTHORIZATIONONLY_PROPERTY_DIRECTORY.html */
func TestGroupAuthorizationOnly(t *testing.T) {
	_, err := directory.GroupAuthorizationOnly()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPAUTHORIZATIONONLY_PROPERTY_DIRECTORY.html */
func TestSetGroupAuthorizationOnly(t *testing.T) {
	s, err := directory.GroupAuthorizationOnly()
	require.Nil(t, err)

	err = directory.SetGroupAuthorizationOnly(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITMATCHES_PROPERTY_DIRECTORY.html */
func TestLimitMatches(t *testing.T) {
	_, err := directory.LimitMatches()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LIMITMATCHES_PROPERTY_DIRECTORY.html */
func TestSetLimitMatches(t *testing.T) {
	s, err := directory.LimitMatches()
	require.Nil(t, err)

	err = directory.SetLimitMatches(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARTIALMATCHES_PROPERTY_DIRECTORY.html */
func TestPartialMatches(t *testing.T) {
	_, err := directory.PartialMatches()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCHALLDIRECTORIES_PROPERTY_DIRECTORY.html */
func TestSearchAllDirectories(t *testing.T) {
	_, err := directory.SearchAllDirectories()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SEARCHALLDIRECTORIES_PROPERTY_DIRECTORY.html */
func TestSetSearchAllDirectories(t *testing.T) {
	s, err := directory.SearchAllDirectories()
	require.Nil(t, err)

	err = directory.SetSearchAllDirectories(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVER_PROPERTY_DIRECTORY.html */
func TestServer(t *testing.T) {
	_, err := directory.Server()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUSTEDONLY_PROPERTY_DIRECTORY.html */
func TestTrustedOnly(t *testing.T) {
	_, err := directory.TrustedOnly()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUSTEDONLY_PROPERTY_DIRECTORY.html */
func TestSetTrustedOnly(t *testing.T) {
	s, err := directory.TrustedOnly()
	require.Nil(t, err)

	err = directory.SetTrustedOnly(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECONTEXTSERVER_PROPERTY_DIRECTORY.html */
func TestUseContextServer(t *testing.T) {
	_, err := directory.UseContextServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECONTEXTSERVER_PROPERTY_DIRECTORY.html */
func TestSetUseContextServer(t *testing.T) {
	s, err := directory.UseContextServer()
	require.Nil(t, err)

	err = directory.SetUseContextServer(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENAVIGATOR_METHOD_DIRECTORY.html */
func TestCreateNavigator(t *testing.T) {
	_, err := directory.CreateNavigator()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREELOOKUPBUFFER_METHOD_DIRECTORY.html */
func TestFreeLookupBuffer(t *testing.T) {
	err := directory.FreeLookupBuffer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMAILINFO_METHOD_DIRECTORY.html */
// func TestGetMailInfo(t *testing.T) {
// 	_, err := directory.GetMailInfo( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOOKUPALLNAMES_METHOD_DIRECTORY.html */
// func TestLookupAllNames(t *testing.T) {
// 	_, err := directory.LookupAllNames( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOOKUPNAMES_METHOD_DIRECTORY.html */
// func TestLookupNames(t *testing.T) {
// 	_, err := directory.LookupNames( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }
