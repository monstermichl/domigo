/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESINTERNATIONAL_CLASS.html */
package notesinternational_test

import (
    "github.com/monstermichl/domigo"
    testhelpers "github.com/monstermichl/domigo/test/helpers"
    "testing"

    "github.com/stretchr/testify/require"
)

var international domigo.NotesInternational

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
    testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
        var err error
        international, err = session.International()
        defer international.Release()

        if err != nil {
            return "NotesInternational could not be created", err
        }
        m.Run()
        return "", nil
    })
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AMSTRING_PROPERTY.html */
func TestAMString(t *testing.T) {
    _, err := international.AMString()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENCYDIGITS_PROPERTY.html */
func TestCurrencyDigits(t *testing.T) {
    _, err := international.CurrencyDigits()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENCYSYMBOL_PROPERTY.html */
func TestCurrencySymbol(t *testing.T) {
    _, err := international.CurrencySymbol()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATESEP_PROPERTY.html */
func TestDateSep(t *testing.T) {
    _, err := international.DateSep()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECIMALSEP_PROPERTY.html */
func TestDecimalSep(t *testing.T) {
    _, err := international.DecimalSep()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENCYSPACE_PROPERTY.html */
func TestIsCurrencySpace(t *testing.T) {
    _, err := international.IsCurrencySpace()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENCYSUFFIX_PROPERTY.html */
func TestIsCurrencySuffix(t *testing.T) {
    _, err := international.IsCurrencySuffix()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENCYZERO_PROPERTY.html */
func TestIsCurrencyZero(t *testing.T) {
    _, err := international.IsCurrencyZero()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDATEDMY_PROPERTY.html */
func TestIsDateDMY(t *testing.T) {
    _, err := international.IsDateDMY()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDATEMDY_PROPERTY.html */
func TestIsDateMDY(t *testing.T) {
    _, err := international.IsDateMDY()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDATEYMD_PROPERTY.html */
func TestIsDateYMD(t *testing.T) {
    _, err := international.IsDateYMD()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDST_PROPERTY_INTL.html */
func TestIsDST(t *testing.T) {
    _, err := international.IsDST()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISTIME24HOUR_PROPERTY.html */
func TestIsTime24Hour(t *testing.T) {
    _, err := international.IsTime24Hour()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_INTERNATIONAL_COM.html */
func TestParent(t *testing.T) {
    _, err := international.Parent()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PMSTRING_PROPERTY.html */
func TestPMString(t *testing.T) {
    _, err := international.PMString()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_THOUSANDSSEP_PROPERTY.html */
func TestThousandsSep(t *testing.T) {
    _, err := international.ThousandsSep()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMESEP_PROPERTY.html */
func TestTimeSep(t *testing.T) {
    _, err := international.TimeSep()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONE_PROPERTY_INTL.html */
func TestTimeZone(t *testing.T) {
    _, err := international.TimeZone()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TODAY_PROPERTY.html */
func TestToday(t *testing.T) {
    _, err := international.Today()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TOMORROW_PROPERTY.html */
func TestTomorrow(t *testing.T) {
    _, err := international.Tomorrow()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_YESTERDAY_PROPERTY.html */
func TestYesterday(t *testing.T) {
    _, err := international.Yesterday()
    require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */