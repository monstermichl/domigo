/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATETIME_CLASS.html */
package notesdatetime_test

import (
	"testing"

	"github.com/monstermichl/domigo/domino/notesdatetime"
	"github.com/monstermichl/domigo/domino/notessession"

	"github.com/stretchr/testify/require"
)

var session notessession.NotesSession
var datetime notesdatetime.NotesDateTime

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.Initialize()
	datetime, _ = session.CreateDateTime("Today")

	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEONLY_PROPERTY.html */
func TestDateOnly(t *testing.T) {
	_, err := datetime.DateOnly()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GMTTIME_PROPERTY.html */
func TestGMTTime(t *testing.T) {
	_, err := datetime.GMTTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDST_PROPERTY.html */
func TestIsDST(t *testing.T) {
	_, err := datetime.IsDST()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALIDDATE_PROPERTY_DATETIME.html */
func TestIsValidDate(t *testing.T) {
	_, err := datetime.IsValidDate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCALTIME_PROPERTY.html */
func TestLocalTime(t *testing.T) {
	_, err := datetime.LocalTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCALTIME_PROPERTY.html */
func TestSetLocalTime(t *testing.T) {
	s, err := datetime.LocalTime()
	require.Nil(t, err)

	err = datetime.SetLocalTime(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSGMTTIME_PROPERTY.html */
func TestLSGMTTime(t *testing.T) {
	_, err := datetime.LSGMTTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSLOCALTIME_PROPERTY.html */
func TestLSLocalTime(t *testing.T) {
	_, err := datetime.LSLocalTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSLOCALTIME_PROPERTY.html */
func TestSetLSLocalTime(t *testing.T) {
	s, err := datetime.LSLocalTime()
	require.Nil(t, err)

	err = datetime.SetLSLocalTime(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEONLY_PROPERTY.html */
func TestTimeOnly(t *testing.T) {
	_, err := datetime.TimeOnly()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONE_PROPERTY.html */
func TestTimeZone(t *testing.T) {
	_, err := datetime.TimeZone()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ZONETIME_PROPERTY.html */
func TestZoneTime(t *testing.T) {
	_, err := datetime.ZoneTime()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTDAY_METHOD.html */
func TestAdjustDay(t *testing.T) {
	err := datetime.AdjustDay(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTHOUR_METHOD.html */
func TestAdjustHour(t *testing.T) {
	err := datetime.AdjustHour(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTMINUTE_METHOD.html */
func TestAdjustMinute(t *testing.T) {
	err := datetime.AdjustMinute(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTMONTH_METHOD.html */
func TestAdjustMonth(t *testing.T) {
	err := datetime.AdjustMonth(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTSECOND_METHOD.html */
func TestAdjustSecond(t *testing.T) {
	err := datetime.AdjustSecond(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTYEAR_METHOD.html */
func TestAdjustYear(t *testing.T) {
	err := datetime.AdjustYear(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTTOZONE_METHOD.html */
func TestConvertToZone(t *testing.T) {
	err := datetime.ConvertToZone(1, false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETANYDATE_METHOD.html */
func TestSetAnyDate(t *testing.T) {
	err := datetime.SetAnyDate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETANYTIME_METHOD.html */
func TestSetAnyTime(t *testing.T) {
	err := datetime.SetAnyTime()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNOW_METHOD.html */
func TestSetNow(t *testing.T) {
	err := datetime.SetNow()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDIFFERENCE_METHOD.html */
func TestTimeDifference(t *testing.T) {
	yesterday, _ := session.CreateDateTime("Yesterday")

	_, err := datetime.TimeDifference(yesterday)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDIFFERENCEDOUBLE_METHOD.html */
func TestTimeDifferenceDouble(t *testing.T) {
	yesterday, _ := session.CreateDateTime("Yesterday")

	_, err := datetime.TimeDifferenceDouble(yesterday)
	require.Nil(t, err)
}
