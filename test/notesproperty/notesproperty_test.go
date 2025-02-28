/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLASS.html */
package notesproperty_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"
	"github.com/stretchr/testify/require"
)

var property domigo.NotesProperty

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		broker, err := session.GetPropertyBroker()
		defer broker.Release()

		if err != nil {
			return "Broker could not be retrieved. Canceling tests...", err
		}
		property, err = broker.GetProperty("Subject")
		defer property.Release()

		if err != nil {
			return "Property could not be retrieved", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_DESCRIPTION_PROPERTY.html */
func TestDescription(t *testing.T) {
	_, err := property.Description()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_ISINPUT_PROPERTY.html */
func TestIsInput(t *testing.T) {
	_, err := property.IsInput()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAME_PROPERTY.html */
func TestName(t *testing.T) {
	_, err := property.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAMESPACE_PROPERTY.html */
func TestNameSpace(t *testing.T) {
	_, err := property.NameSpace()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TITLE_PROPERTY.html */
func TestTitle(t *testing.T) {
	_, err := property.Title()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TYPENAME_PROPERTY.html */
func TestTypename(t *testing.T) {
	_, err := property.Typename()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func TestValues(t *testing.T) {
	_, err := property.Values()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func TestSetValues(t *testing.T) {
	s, err := property.Values()
	require.Nil(t, err)

	err = property.SetValues(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLEAR_METHOD.html */
func TestClear(t *testing.T) {
	err := property.Clear()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_PUBLISH_METHOD.html */
func TestPublish(t *testing.T) {
	err := property.Publish()
	require.Nil(t, err)
}
