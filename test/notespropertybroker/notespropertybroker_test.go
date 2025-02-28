/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTYBROKER_CLASS.html */
package notespropertybroker_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"
	"github.com/stretchr/testify/require"
)

var propertybroker domigo.NotesPropertyBroker

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		propertybroker, err = session.GetPropertyBroker()
		defer propertybroker.Release()

		if err != nil {
			return "Property broker could not be retrieved", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_INPUTPROPERTYCONTEXT.html */
func TestInputPropertyContext(t *testing.T) {
	_, err := propertybroker.InputPropertyContext()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_CLEARPROPERTY_METHOD.html */
// func TestClearProperty(t *testing.T) {
// 	err := propertybroker.ClearProperty( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTY_METHOD.html */
// func TestGetProperty(t *testing.T) {
// 	err := propertybroker.GetProperty( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTYVALUE_METHOD.html */
// func TestGetPropertyValue(t *testing.T) {
// 	err := propertybroker.GetPropertyValue( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_HASPROPERTY_METHOD.html */
// func TestHasProperty(t *testing.T) {
// 	_, err := propertybroker.HasProperty( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_PUBLISH_METHOD.html */
func TestPublish(t *testing.T) {
	err := propertybroker.Publish()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_SETPROPERTYVALUE_METHOD.html */
// func TestSetPropertyValue(t *testing.T) {
// 	err := propertybroker.SetPropertyValue( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }
