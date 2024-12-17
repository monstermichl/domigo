/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOMINOQUERY_CLASS.html#reference_fzc_syh_cgb */
package notesdominoquery

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdatetime"
	"github.com/monstermichl/domigo/domino/notesdocumentcollection"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDominoQuery struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDominoQuery {
	return NotesDominoQuery{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANDOCS_PROPERTY_NDQ.html#reference_a4g_r13_cgb */
func (d NotesDominoQuery) MaxScanDocs() (domino.Long, error) {
	val, err := d.Com().GetProperty("MaxScanDocs")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANDOCS_PROPERTY_NDQ.html#reference_a4g_r13_cgb */
func (d NotesDominoQuery) SetMaxScanDocs(v domino.Long) error {
	return d.Com().PutProperty("MaxScanDocs", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANENTRIES_PROPERTY_NDQ.html#reference_yyx_cd3_cgb */
func (d NotesDominoQuery) MaxScanEntries() (domino.Long, error) {
	val, err := d.Com().GetProperty("MaxScanEntries")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANENTRIES_PROPERTY_NDQ.html#reference_yyx_cd3_cgb */
func (d NotesDominoQuery) SetMaxScanEntries(v domino.Long) error {
	return d.Com().PutProperty("MaxScanEntries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOVIEWS_PROPERTY_NDQ.html#reference_o3c_kf3_cgb */
func (d NotesDominoQuery) NoViews() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("NoViews")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOVIEWS_PROPERTY_NDQ.html#reference_o3c_kf3_cgb */
func (d NotesDominoQuery) SetNoViews(v domino.Boolean) error {
	return d.Com().PutProperty("NoViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REBUILDDESIGNCATALOG_PROPERTY_NDQ.html#reference_pdc_hyy_ljb */
func (d NotesDominoQuery) RebuildDesignCatalog() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("RebuildDesignCatalog")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REBUILDDESIGNCATALOG_PROPERTY_NDQ.html#reference_pdc_hyy_ljb */
func (d NotesDominoQuery) SetRebuildDesignCatalog(v domino.Boolean) error {
	return d.Com().PutProperty("RebuildDesignCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHDESIGNCATALOG_PROPERTY_NDQ.html#reference_vvz_3zy_ljb */
func (d NotesDominoQuery) RefreshDesignCatalog() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("RefreshDesignCatalog")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHDESIGNCATALOG_PROPERTY_NDQ.html#reference_vvz_3zy_ljb */
func (d NotesDominoQuery) SetRefreshDesignCatalog(v domino.Boolean) error {
	return d.Com().PutProperty("RefreshDesignCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHFULLTEXT_PROPERTY_NDQ.html#reference_wzg_d1z_ljb */
func (d NotesDominoQuery) RefreshFullText() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("RefreshFullText")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHFULLTEXT_PROPERTY_NDQ.html#reference_wzg_d1z_ljb */
func (d NotesDominoQuery) SetRefreshFullText(v domino.Boolean) error {
	return d.Com().PutProperty("RefreshFullText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHVIEWS_PROPERTY_NDQ.html#reference_enn_rx3_cgb */
func (d NotesDominoQuery) RefreshViews() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("RefreshViews")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHVIEWS_PROPERTY_NDQ.html#reference_enn_rx3_cgb */
func (d NotesDominoQuery) SetRefreshViews(v domino.Boolean) error {
	return d.Com().PutProperty("RefreshViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNDOCUMENTSONLY_PROPERTY_NDQ.html */
func (d NotesDominoQuery) DesignDocumentsOnly() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("DesignDocumentsOnly")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNDOCUMENTSONLY_PROPERTY_NDQ.html */
func (d NotesDominoQuery) SetDesignDocumentsOnly(v domino.Boolean) error {
	return d.Com().PutProperty("DesignDocumentsOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSEC_PROPERTY_NDQ.html#reference_qwv_123_cgb */
func (d NotesDominoQuery) TimeoutSec() (domino.Integer, error) {
	val, err := d.Com().GetProperty("TimeoutSec")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSEC_PROPERTY_NDQ.html#reference_qwv_123_cgb */
func (d NotesDominoQuery) SetTimeoutSec(v domino.Integer) error {
	return d.Com().PutProperty("TimeoutSec", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTE_METHOD_NDQ.html#reference_ivd_gz3_cgb */
type executeParams struct {
	resultName  *domino.String
	replace     *domino.Boolean
	expireHours *domino.Long
}

type executeParam func(*executeParams)

func WithExecuteResultName(resultName domino.String) executeParam {
	return func(c *executeParams) {
		c.resultName = &resultName
	}
}

func WithExecuteReplace(replace domino.Boolean) executeParam {
	return func(c *executeParams) {
		c.replace = &replace
	}
}

func WithExecuteExpireHours(expireHours domino.Long) executeParam {
	return func(c *executeParams) {
		c.expireHours = &expireHours
	}
}

func (d NotesDominoQuery) Execute(query domino.String, params ...executeParam) (notesdocumentcollection.NotesDocumentCollection, error) {
	paramsStruct := &executeParams{}
	paramsOrdered := []interface{}{query}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.resultName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.resultName)
		if paramsStruct.replace != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.replace)
			if paramsStruct.expireHours != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.expireHours)
			}
		}
	}

	dispatchPtr, err := d.Com().CallObjectMethod("Execute", paramsOrdered...)
	return notesdocumentcollection.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPLAIN_METHOD_NDQ.html#reference_jkw_wcj_cgb */
func (d NotesDominoQuery) Explain(query domino.String) (domino.String, error) {
	val, err := d.Com().CallMethod("Explain", query)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARSE_METHOD_NDQ.html#reference_rqs_m2j_cgb */
func (d NotesDominoQuery) Parse(query domino.String) (domino.String, error) {
	val, err := d.Com().CallMethod("Parse", query)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESETNAMEDVARIABLES_METHOD_NDQ.html#reference_rz2_thj_cgb */
func (d NotesDominoQuery) ResetNamedVariables() error {
	_, err := d.Com().CallMethod("ResetNamedVariables")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
func (d NotesDominoQuery) SetNamedVariableString(variableName domino.String, value domino.String) error {
	_, err := d.Com().CallMethod("SetNamedVariable", variableName, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
func (d NotesDominoQuery) SetNamedVariableNumber(variableName domino.String, value domino.Long) error {
	_, err := d.Com().CallMethod("SetNamedVariable", variableName, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
func (d NotesDominoQuery) SetNamedVariableDateTime(variableName domino.String, value notesdatetime.NotesDateTime) error {
	_, err := d.Com().CallMethod("SetNamedVariable", variableName, value.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEINDEX_METHOD_NDQ.html#reference_mvq_brw_ljb */
type createIndexParams struct {
	isVisible *domino.Boolean
	noBuild   *domino.Boolean
}

type createIndexParam func(*createIndexParams)

func WithCreateIndexIsVisible(isVisible domino.Boolean) createIndexParam {
	return func(c *createIndexParams) {
		c.isVisible = &isVisible
	}
}

func WithCreateIndexNoBuild(noBuild domino.Boolean) createIndexParam {
	return func(c *createIndexParams) {
		c.noBuild = &noBuild
	}
}

func (d NotesDominoQuery) CreateIndex(indexName domino.String, field []domino.String, params ...createIndexParam) error {
	paramsStruct := &createIndexParams{}
	paramsOrdered := []interface{}{indexName, field}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.isVisible != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.isVisible)
		if paramsStruct.noBuild != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.noBuild)
		}
	}
	_, err := d.Com().CallMethod("CreateIndex", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDEXES_METHOD_NDQ.html#reference_rkj_2rw_ljb */
/* TODO: Re-add later. JSON modules have to many circular dependencies for now. */
// func (d NotesDominoQuery) ListIndexes() (notesjsonnavigator.NotesJSONNavigator, error) {
// 	dispatchPtr, err := d.Com().CallObjectMethod("ListIndexes")
// 	return notesjsonnavigator.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEINDEX_METHOD_NDQ.html#reference_uqw_grw_ljb */
func (d NotesDominoQuery) RemoveIndex(indexName domino.String) error {
	_, err := d.Com().CallMethod("RemoveIndex", indexName)
	return err
}
