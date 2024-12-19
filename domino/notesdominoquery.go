/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOMINOQUERY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDominoQuery struct {
	NotesStruct
}

func NewNotesDominoQuery(dispatchPtr *ole.IDispatch) NotesDominoQuery {
	return NotesDominoQuery{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANDOCS_PROPERTY_NDQ.html#reference_a4g_r13_cgb */
func (d NotesDominoQuery) MaxScanDocs() (Long, error) {
	val, err := d.com().GetProperty("MaxScanDocs")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANDOCS_PROPERTY_NDQ.html#reference_a4g_r13_cgb */
func (d NotesDominoQuery) SetMaxScanDocs(v Long) error {
	return d.com().PutProperty("MaxScanDocs", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANENTRIES_PROPERTY_NDQ.html#reference_yyx_cd3_cgb */
func (d NotesDominoQuery) MaxScanEntries() (Long, error) {
	val, err := d.com().GetProperty("MaxScanEntries")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXSCANENTRIES_PROPERTY_NDQ.html#reference_yyx_cd3_cgb */
func (d NotesDominoQuery) SetMaxScanEntries(v Long) error {
	return d.com().PutProperty("MaxScanEntries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOVIEWS_PROPERTY_NDQ.html#reference_o3c_kf3_cgb */
func (d NotesDominoQuery) NoViews() (Boolean, error) {
	val, err := d.com().GetProperty("NoViews")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOVIEWS_PROPERTY_NDQ.html#reference_o3c_kf3_cgb */
func (d NotesDominoQuery) SetNoViews(v Boolean) error {
	return d.com().PutProperty("NoViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REBUILDDESIGNCATALOG_PROPERTY_NDQ.html#reference_pdc_hyy_ljb */
func (d NotesDominoQuery) RebuildDesignCatalog() (Boolean, error) {
	val, err := d.com().GetProperty("RebuildDesignCatalog")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REBUILDDESIGNCATALOG_PROPERTY_NDQ.html#reference_pdc_hyy_ljb */
func (d NotesDominoQuery) SetRebuildDesignCatalog(v Boolean) error {
	return d.com().PutProperty("RebuildDesignCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHDESIGNCATALOG_PROPERTY_NDQ.html#reference_vvz_3zy_ljb */
func (d NotesDominoQuery) RefreshDesignCatalog() (Boolean, error) {
	val, err := d.com().GetProperty("RefreshDesignCatalog")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHDESIGNCATALOG_PROPERTY_NDQ.html#reference_vvz_3zy_ljb */
func (d NotesDominoQuery) SetRefreshDesignCatalog(v Boolean) error {
	return d.com().PutProperty("RefreshDesignCatalog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHFULLTEXT_PROPERTY_NDQ.html#reference_wzg_d1z_ljb */
func (d NotesDominoQuery) RefreshFullText() (Boolean, error) {
	val, err := d.com().GetProperty("RefreshFullText")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHFULLTEXT_PROPERTY_NDQ.html#reference_wzg_d1z_ljb */
func (d NotesDominoQuery) SetRefreshFullText(v Boolean) error {
	return d.com().PutProperty("RefreshFullText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHVIEWS_PROPERTY_NDQ.html#reference_enn_rx3_cgb */
func (d NotesDominoQuery) RefreshViews() (Boolean, error) {
	val, err := d.com().GetProperty("RefreshViews")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESHVIEWS_PROPERTY_NDQ.html#reference_enn_rx3_cgb */
func (d NotesDominoQuery) SetRefreshViews(v Boolean) error {
	return d.com().PutProperty("RefreshViews", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNDOCUMENTSONLY_PROPERTY_NDQ.html */
func (d NotesDominoQuery) DesignDocumentsOnly() (Boolean, error) {
	val, err := d.com().GetProperty("DesignDocumentsOnly")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNDOCUMENTSONLY_PROPERTY_NDQ.html */
func (d NotesDominoQuery) SetDesignDocumentsOnly(v Boolean) error {
	return d.com().PutProperty("DesignDocumentsOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSEC_PROPERTY_NDQ.html#reference_qwv_123_cgb */
func (d NotesDominoQuery) TimeoutSec() (Integer, error) {
	val, err := d.com().GetProperty("TimeoutSec")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSEC_PROPERTY_NDQ.html#reference_qwv_123_cgb */
func (d NotesDominoQuery) SetTimeoutSec(v Integer) error {
	return d.com().PutProperty("TimeoutSec", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTE_METHOD_NDQ.html#reference_ivd_gz3_cgb */
type notesDominoQueryExecuteParams struct {
	resultName  *String
	replace     *Boolean
	expireHours *Long
}

type notesDominoQueryExecuteParam func(*notesDominoQueryExecuteParams)

func WithNotesDominoQueryExecuteResultName(resultName String) notesDominoQueryExecuteParam {
	return func(c *notesDominoQueryExecuteParams) {
		c.resultName = &resultName
	}
}

func WithNotesDominoQueryExecuteReplace(replace Boolean) notesDominoQueryExecuteParam {
	return func(c *notesDominoQueryExecuteParams) {
		c.replace = &replace
	}
}

func WithNotesDominoQueryExecuteExpireHours(expireHours Long) notesDominoQueryExecuteParam {
	return func(c *notesDominoQueryExecuteParams) {
		c.expireHours = &expireHours
	}
}

func (d NotesDominoQuery) Execute(query String, params ...notesDominoQueryExecuteParam) (NotesDocumentCollection, error) {
	paramsStruct := &notesDominoQueryExecuteParams{}
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

	dispatchPtr, err := d.com().CallObjectMethod("Execute", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPLAIN_METHOD_NDQ.html#reference_jkw_wcj_cgb */
func (d NotesDominoQuery) Explain(query String) (String, error) {
	val, err := d.com().CallMethod("Explain", query)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARSE_METHOD_NDQ.html#reference_rqs_m2j_cgb */
func (d NotesDominoQuery) Parse(query String) (String, error) {
	val, err := d.com().CallMethod("Parse", query)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESETNAMEDVARIABLES_METHOD_NDQ.html#reference_rz2_thj_cgb */
func (d NotesDominoQuery) ResetNamedVariables() error {
	_, err := d.com().CallMethod("ResetNamedVariables")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDVARIABLE_METHOD_NDQ.html#reference_h2j_fgj_cgb */
/* TODO: Handle different input types (see documentation). */
func (d NotesDominoQuery) SetNamedVariable(variableName String, value any) error {
	err := helpers.CheckTypeNames(value, []string{"string", "number", "NotesDateTime"})

	if err != nil {

	}
	_, err = d.com().CallMethod("SetNamedVariable", variableName, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEINDEX_METHOD_NDQ.html#reference_mvq_brw_ljb */
type notesDominoQueryCreateIndexParams struct {
	isVisible *Boolean
	noBuild   *Boolean
}

type notesDominoQueryCreateIndexParam func(*notesDominoQueryCreateIndexParams)

func WithNotesDominoQueryCreateIndexIsVisible(isVisible Boolean) notesDominoQueryCreateIndexParam {
	return func(c *notesDominoQueryCreateIndexParams) {
		c.isVisible = &isVisible
	}
}

func WithNotesDominoQueryCreateIndexNoBuild(noBuild Boolean) notesDominoQueryCreateIndexParam {
	return func(c *notesDominoQueryCreateIndexParams) {
		c.noBuild = &noBuild
	}
}

func (d NotesDominoQuery) CreateIndex(indexName String, field []String, params ...notesDominoQueryCreateIndexParam) error {
	paramsStruct := &notesDominoQueryCreateIndexParams{}
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
	_, err := d.com().CallMethod("CreateIndex", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTINDEXES_METHOD_NDQ.html#reference_rkj_2rw_ljb */
/* TODO: Re-add later. JSON modules have to many circular dependencies for now. */
// func (d NotesDominoQuery) ListIndexes() (NotesJSONNavigator, error) {
// 	dispatchPtr, err := d.com().CallObjectMethod("ListIndexes")
// 	return NewNotesJSONNavigator(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEINDEX_METHOD_NDQ.html#reference_uqw_grw_ljb */
func (d NotesDominoQuery) RemoveIndex(indexName String) error {
	_, err := d.com().CallMethod("RemoveIndex", indexName)
	return err
}
