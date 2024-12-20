/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATERANGE_CLASS.html */
package notesdaterange_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var daterange domigo.NotesDateRange

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		daterange, err = session.CreateDateRange()
		defer daterange.Release()

		if err != nil {
			return "Date range could not be created", err
		}
		start, err := session.CreateDateTime("Yesterday")

		if err != nil {
			return "Start could not be created", err
		}
		stop, err := session.CreateDateTime("Today")

		if err != nil {
			return "Stop could not be created", err
		}
		err = daterange.SetStartDateTime(start)

		if err != nil {
			return "Start could not be set", err
		}
		err = daterange.SetEndDateTime(stop)

		if err != nil {
			return "End could not be set", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func TestEndDateTime(t *testing.T) {
	_, err := daterange.EndDateTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func TestSetEndDateTime(t *testing.T) {
	s, err := daterange.EndDateTime()
	require.Nil(t, err)

	err = daterange.SetEndDateTime(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func TestStartDateTime(t *testing.T) {
	_, err := daterange.StartDateTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func TestSetStartDateTime(t *testing.T) {
	s, err := daterange.StartDateTime()
	require.Nil(t, err)

	err = daterange.SetStartDateTime(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func TestText(t *testing.T) {
	_, err := daterange.Text()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func TestSetText(t *testing.T) {
	s, err := daterange.Text()
	require.Nil(t, err)

	err = daterange.SetText(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
