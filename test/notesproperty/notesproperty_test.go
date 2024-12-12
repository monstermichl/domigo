/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLASS.html */
package notesproperty_test

import (
	"domigo/domino/notesproperty"
	"domigo/domino/notessession"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var property notesproperty.NotesProperty

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	defer session.Release()

	broker, err := session.GetPropertyBroker()

	if err != nil {
		fmt.Println("Broker could not be retrieved. Canceling tests...")
		return
	}
	property, _ = broker.GetProperty("Subject")

	m.Run()
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
