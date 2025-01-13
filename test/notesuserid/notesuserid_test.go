/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESUSERID_CLASS.html */
package notesuserid_test

import (
	"errors"
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var userid domigo.NotesUserID

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error = errors.New("User ID could not be retrieved")
		//userid, err = /* TODO: Get variable. */
		defer userid.Release()

		if err != nil {
			return "NotesUserID could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENCRYPTIONKEYS_USERID.html */
func TestGetEncryptionKeys(t *testing.T) {
	_, err := userid.GetEncryptionKeys()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERNAME_USERID.html */
func TestGetUserName(t *testing.T) {
	_, err := userid.GetUserName()
	require.Nil(t, err)
}
