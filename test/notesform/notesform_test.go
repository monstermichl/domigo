/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFORM_CLASS.html */
package notesform_test

import (
	"fmt"
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var form domigo.NotesForm

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		const TEST_FORM = "TestForm"
		forms, err := db.Forms()

		if err != nil {
			return "Forms could not be retrieved", err
		}

		for _, f := range forms {
			defer f.Release()
		}
		err = db.SetIsDesignLockingEnabled(true)

		if err != nil {
			return "Design locking could not be enabled", err
		}

		for _, f := range forms {
			name, err := f.Name()

			if err != nil {
				return "Name could not be retrieved", err
			}

			if name == TEST_FORM {
				form = f
			}
		}

		if !form.IsReady() {
			return "", fmt.Errorf("%s could not be found", TEST_FORM)
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_FORM.html */
func TestAliases(t *testing.T) {
	_, err := form.Aliases()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIELDS_PROPERTY.html */
func TestFields(t *testing.T) {
	_, err := form.Fields()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMUSERS_PROPERTY.html */
func TestFormUsers(t *testing.T) {
	_, err := form.FormUsers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMUSERS_PROPERTY.html */
func TestSetFormUsers(t *testing.T) {
	s, err := form.FormUsers()
	require.Nil(t, err)

	err = form.SetFormUsers(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_FORM_COM.html */
func TestHttpURL(t *testing.T) {
	_, err := form.HttpURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUBFORM_PROPERTY.html */
func TestIsSubForm(t *testing.T) {
	_, err := form.IsSubForm()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_FORM.html */
func TestLockHolders(t *testing.T) {
	_, err := form.LockHolders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_FORM.html */
func TestName(t *testing.T) {
	_, err := form.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_FORM_COM.html */
func TestNotesURL(t *testing.T) {
	_, err := form.NotesURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY.html */
func TestProtectReaders(t *testing.T) {
	_, err := form.ProtectReaders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY.html */
func TestSetProtectReaders(t *testing.T) {
	s, err := form.ProtectReaders()
	require.Nil(t, err)

	err = form.SetProtectReaders(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTUSERS_PROPERTY.html */
func TestProtectUsers(t *testing.T) {
	_, err := form.ProtectUsers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTUSERS_PROPERTY.html */
func TestSetProtectUsers(t *testing.T) {
	s, err := form.ProtectUsers()
	require.Nil(t, err)

	err = form.SetProtectUsers(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY.html */
func TestReaders(t *testing.T) {
	_, err := form.Readers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY.html */
func TestSetReaders(t *testing.T) {
	s, err := form.Readers()
	require.Nil(t, err)

	err = form.SetReaders(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIELDTYPE_METHOD_FORM.html */
func TestGetFieldType(t *testing.T) {
	_, err := form.GetFieldType("testField")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_FORM.html */
func TestLock(t *testing.T) {
	_, err := form.Lock()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_FORM.html */
func TestUnLock(t *testing.T) {
	err := form.UnLock()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_FORM.html */
func TestLockProvisional(t *testing.T) {
	_, err := form.LockProvisional()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_FORM.html */
func TestRemove(t *testing.T) {
	err := form.Remove()
	require.Nil(t, err)
}
