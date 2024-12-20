/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDIRECTORYNAVIGATOR_CLASS.html */
package notesdirectorynavigator_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"
	"github.com/stretchr/testify/require"
)

var directorynavigator domigo.NotesDirectoryNavigator

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		directory, err := session.GetDirectory()
		defer directory.Release()

		if err != nil {
			return "Directory could not be created", err
		}
		directorynavigator, err = directory.CreateNavigator()

		if err != nil {
			return "Navigator could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTITEM_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestCurrentItem(t *testing.T) {
	_, err := directorynavigator.CurrentItem()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCH_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestCurrentMatch(t *testing.T) {
	_, err := directorynavigator.CurrentMatch()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTMATCHES_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestCurrentMatches(t *testing.T) {
	_, err := directorynavigator.CurrentMatches()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTNAME_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestCurrentName(t *testing.T) {
	_, err := directorynavigator.CurrentName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTVIEW_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestCurrentView(t *testing.T) {
	_, err := directorynavigator.CurrentView()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MATCHLOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestMatchLocated(t *testing.T) {
	_, err := directorynavigator.MatchLocated()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMELOCATED_PROPERTY_DIRECTORYNAVIGATOR.html */
func TestNameLocated(t *testing.T) {
	_, err := directorynavigator.NameLocated()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func TestFindFirstMatch(t *testing.T) {
	_, err := directorynavigator.FindFirstMatch()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func TestFindFirstName(t *testing.T) {
	_, err := directorynavigator.FindFirstName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func TestFindNextMatch(t *testing.T) {
	_, err := directorynavigator.FindNextMatch()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTNAME_METHOD_DIRECTORYNAVIGATOR.html */
func TestFindNextName(t *testing.T) {
	_, err := directorynavigator.FindNextName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHMATCH_METHOD_DIRECTORYNAVIGATOR.html */
func TestFindNthMatch(t *testing.T) {
	_, err := directorynavigator.FindNthMatch(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHNAME_METHOD_DIRECTORYNAVIGATOR.html */
func TestFindNthName(t *testing.T) {
	_, err := directorynavigator.FindNthName(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func TestGetFirstItemValue(t *testing.T) {
	_, err := directorynavigator.GetFirstItemValue()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func TestGetNextItemValue(t *testing.T) {
	_, err := directorynavigator.GetNextItemValue()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHITEMVALUE_METHOD_DIRECTORYNAVIGATOR.html */
func TestGetNthItemValue(t *testing.T) {
	_, err := directorynavigator.GetNthItemValue(1)
	require.Nil(t, err)
}
