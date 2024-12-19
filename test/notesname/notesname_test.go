package notesname_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	"github.com/stretchr/testify/require"
)

var notesName domigo.NotesName

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := domigo.Initialize()
	db, _ := session.GetDatabase("", "GoInterface.nsf")
	aclTmp, _ := db.ACL()
	aclEntry, _ := aclTmp.GetFirstEntry()
	notesName, _ = aclEntry.NameObject()

	m.Run()
	session.Release()
}
func TestAbbreviated(t *testing.T) {
	_, err := notesName.Abbreviated()
	require.Nil(t, err)
}

func TestAddr821(t *testing.T) {
	_, err := notesName.Addr821()
	require.Nil(t, err)
}

func TestAddr822Comment1(t *testing.T) {
	_, err := notesName.Addr822Comment1()
	require.Nil(t, err)
}

func TestAddr822Comment2(t *testing.T) {
	_, err := notesName.Addr822Comment2()
	require.Nil(t, err)
}

func TestAddr822Comment3(t *testing.T) {
	_, err := notesName.Addr822Comment3()
	require.Nil(t, err)
}

func TestAddr822LocalPart(t *testing.T) {
	_, err := notesName.Addr822LocalPart()
	require.Nil(t, err)
}

func TestAddr822Phrase(t *testing.T) {
	_, err := notesName.Addr822Phrase()
	require.Nil(t, err)
}

func TestADMD(t *testing.T) {
	_, err := notesName.ADMD()
	require.Nil(t, err)
}

func TestCanonical(t *testing.T) {
	_, err := notesName.Canonical()
	require.Nil(t, err)
}

func TestCommon(t *testing.T) {
	_, err := notesName.Common()
	require.Nil(t, err)
}

func TestCountry(t *testing.T) {
	_, err := notesName.Country()
	require.Nil(t, err)
}

func TestGeneration(t *testing.T) {
	_, err := notesName.Generation()
	require.Nil(t, err)
}

func TestGiven(t *testing.T) {
	_, err := notesName.Given()
	require.Nil(t, err)
}

func TestInitials(t *testing.T) {
	_, err := notesName.Initials()
	require.Nil(t, err)
}

func TestIsHierarchical(t *testing.T) {
	_, err := notesName.IsHierarchical()
	require.Nil(t, err)
}

func TestKeyword(t *testing.T) {
	_, err := notesName.Keyword()
	require.Nil(t, err)
}

func TestLanguage(t *testing.T) {
	_, err := notesName.Language()
	require.Nil(t, err)
}

func TestOrganization(t *testing.T) {
	_, err := notesName.Organization()
	require.Nil(t, err)
}

func TestOrgUnit1(t *testing.T) {
	_, err := notesName.OrgUnit1()
	require.Nil(t, err)
}

func TestOrgUnit2(t *testing.T) {
	_, err := notesName.OrgUnit2()
	require.Nil(t, err)
}

func TestOrgUnit3(t *testing.T) {
	_, err := notesName.OrgUnit3()
	require.Nil(t, err)
}

func TestOrgUnit4(t *testing.T) {
	_, err := notesName.OrgUnit4()
	require.Nil(t, err)
}

func TestPRMD(t *testing.T) {
	_, err := notesName.PRMD()
	require.Nil(t, err)
}

func TestSurname(t *testing.T) {
	_, err := notesName.Surname()
	require.Nil(t, err)
}
