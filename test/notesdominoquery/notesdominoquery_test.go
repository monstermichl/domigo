/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOMINOQUERY_CLASS.html#reference_fzc_syh_cgb */
package notesdominoquery_test

import (
	"testing"

	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdominoquery"
	"github.com/monstermichl/domigo/domino/notessession"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var session notessession.NotesSession
var dominoquery notesdominoquery.NotesDominoQuery

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.Initialize()
	db, _ := testhelpers.CreateTestDatabase(session)
	dominoquery, _ := db.CreateDominoQuery()

	defer dominoquery.Release()
	defer db.Release()
	defer db.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* TODO: Access type for MaxScanDocs could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANDOCS_PROPERTY_NDQ.html#reference_a4g_r13_cgb */
func TestMaxScanDocs(t *testing.T) {
	_, err := dominoquery.MaxScanDocs()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANDOCS_PROPERTY_NDQ.html#reference_a4g_r13_cgb */
func TestSetMaxScanDocs(t *testing.T) {
	s, err := dominoquery.MaxScanDocs()
	require.Nil(t, err)

	err = dominoquery.SetMaxScanDocs(s)
	require.Nil(t, err)
}

/* TODO: Access type for MaxScanEntries could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANENTRIES_PROPERTY_NDQ.html#reference_yyx_cd3_cgb */
func TestMaxScanEntries(t *testing.T) {
	_, err := dominoquery.MaxScanEntries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANENTRIES_PROPERTY_NDQ.html#reference_yyx_cd3_cgb */
func TestSetMaxScanEntries(t *testing.T) {
	s, err := dominoquery.MaxScanEntries()
	require.Nil(t, err)

	err = dominoquery.SetMaxScanEntries(s)
	require.Nil(t, err)
}

/* TODO: Access type for NoViews could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOVIEWS_PROPERTY_NDQ.html#reference_o3c_kf3_cgb */
func TestNoViews(t *testing.T) {
	_, err := dominoquery.NoViews()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOVIEWS_PROPERTY_NDQ.html#reference_o3c_kf3_cgb */
func TestSetNoViews(t *testing.T) {
	s, err := dominoquery.NoViews()
	require.Nil(t, err)

	err = dominoquery.SetNoViews(s)
	require.Nil(t, err)
}

/* TODO: Access type for RebuildDesignCatalog could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REBUILDDESIGNCATALOG_PROPERTY_NDQ.html#reference_pdc_hyy_ljb */
func TestRebuildDesignCatalog(t *testing.T) {
	_, err := dominoquery.RebuildDesignCatalog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REBUILDDESIGNCATALOG_PROPERTY_NDQ.html#reference_pdc_hyy_ljb */
func TestSetRebuildDesignCatalog(t *testing.T) {
	s, err := dominoquery.RebuildDesignCatalog()
	require.Nil(t, err)

	err = dominoquery.SetRebuildDesignCatalog(s)
	require.Nil(t, err)
}

/* TODO: Access type for RefreshDesignCatalog could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHDESIGNCATALOG_PROPERTY_NDQ.html#reference_vvz_3zy_ljb */
func TestRefreshDesignCatalog(t *testing.T) {
	_, err := dominoquery.RefreshDesignCatalog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHDESIGNCATALOG_PROPERTY_NDQ.html#reference_vvz_3zy_ljb */
func TestSetRefreshDesignCatalog(t *testing.T) {
	s, err := dominoquery.RefreshDesignCatalog()
	require.Nil(t, err)

	err = dominoquery.SetRefreshDesignCatalog(s)
	require.Nil(t, err)
}

/* TODO: Access type for RefreshFullText could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHFULLTEXT_PROPERTY_NDQ.html#reference_wzg_d1z_ljb */
func TestRefreshFullText(t *testing.T) {
	_, err := dominoquery.RefreshFullText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHFULLTEXT_PROPERTY_NDQ.html#reference_wzg_d1z_ljb */
func TestSetRefreshFullText(t *testing.T) {
	s, err := dominoquery.RefreshFullText()
	require.Nil(t, err)

	err = dominoquery.SetRefreshFullText(s)
	require.Nil(t, err)
}

/* TODO: Access type for RefreshViews could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHVIEWS_PROPERTY_NDQ.html#reference_enn_rx3_cgb */
func TestRefreshViews(t *testing.T) {
	_, err := dominoquery.RefreshViews()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHVIEWS_PROPERTY_NDQ.html#reference_enn_rx3_cgb */
func TestSetRefreshViews(t *testing.T) {
	s, err := dominoquery.RefreshViews()
	require.Nil(t, err)

	err = dominoquery.SetRefreshViews(s)
	require.Nil(t, err)
}

/* TODO: Access type for DesignDocumentsOnly could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNDOCUMENTSONLY_PROPERTY_NDQ.html */
func TestDesignDocumentsOnly(t *testing.T) {
	_, err := dominoquery.DesignDocumentsOnly()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNDOCUMENTSONLY_PROPERTY_NDQ.html */
func TestSetDesignDocumentsOnly(t *testing.T) {
	s, err := dominoquery.DesignDocumentsOnly()
	require.Nil(t, err)

	err = dominoquery.SetDesignDocumentsOnly(s)
	require.Nil(t, err)
}

/* TODO: Access type for TimeoutSec could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSEC_PROPERTY_NDQ.html#reference_qwv_123_cgb */
func TestTimeoutSec(t *testing.T) {
	_, err := dominoquery.TimeoutSec()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSEC_PROPERTY_NDQ.html#reference_qwv_123_cgb */
func TestSetTimeoutSec(t *testing.T) {
	s, err := dominoquery.TimeoutSec()
	require.Nil(t, err)

	err = dominoquery.SetTimeoutSec(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTE_METHOD_NDQ.html#reference_ivd_gz3_cgb */
func TestExecute(t *testing.T) {
	_, err := dominoquery.Execute("lastName = 'Herold'")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPLAIN_METHOD_NDQ.html#reference_jkw_wcj_cgb */
func TestExplain(t *testing.T) {
	_, err := dominoquery.Explain("lastName = 'Herold'")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARSE_METHOD_NDQ.html#reference_rqs_m2j_cgb */
func TestParse(t *testing.T) {
	_, err := dominoquery.Parse("lastName = 'Herold'")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESETNAMEDVARIABLES_METHOD_NDQ.html#reference_rz2_thj_cgb */
func TestResetNamedVariables(t *testing.T) {
	err := dominoquery.ResetNamedVariables()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
func TestSetNamedVariableString(t *testing.T) {
	err := dominoquery.SetNamedVariableString("variable", "value")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
func TestSetNamedVariableNumber(t *testing.T) {
	err := dominoquery.SetNamedVariableNumber("variable", 0)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
func TestSetNamedVariableDateTime(t *testing.T) {
	dt, err := session.CreateDateTime("Today")
	require.Nil(t, err)

	err = dominoquery.SetNamedVariableDateTime("variable", dt)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEINDEX_METHOD_NDQ.html#reference_mvq_brw_ljb */
func TestCreateIndex(t *testing.T) {
	err := dominoquery.CreateIndex("testIndex", []domino.String{})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDEXES_METHOD_NDQ.html#reference_rkj_2rw_ljb */
// func TestListIndexes(t *testing.T) {
// 	_, err := dominoquery.ListIndexes()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEINDEX_METHOD_NDQ.html#reference_uqw_grw_ljb */
func TestRemoveIndex(t *testing.T) {
	err := dominoquery.RemoveIndex("testIndex")
	require.Nil(t, err)
}
