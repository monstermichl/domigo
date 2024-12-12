package notesqueryresultsprocessor

import (
	"domigo/domino"
	"domigo/domino/notesdominoquery"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesQueryResultsProcessor struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesQueryResultsProcessor {
	return NotesQueryResultsProcessor{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* TODO: Access type for MaxEntries could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func (q NotesQueryResultsProcessor) MaxEntries() (domino.Long, error) {
	val, err := q.Com().GetProperty("MaxEntries")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func (q NotesQueryResultsProcessor) SetMaxEntries(v domino.Long) error {
	return q.Com().PutProperty("MaxEntries", v)
}

/* TODO: Access type for TimeOutSec could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSECQRP_PROPERTY.html */
func (q NotesQueryResultsProcessor) TimeOutSec() (domino.Integer, error) {
	val, err := q.Com().GetProperty("TimeOutSec")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSECQRP_PROPERTY.html */
func (q NotesQueryResultsProcessor) SetTimeOutSec(v domino.Integer) error {
	return q.Com().PutProperty("TimeOutSec", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCOLLECTION_METHOD.html */
func (q NotesQueryResultsProcessor) AddCollection(Collection domino.Variant, ReferenceName domino.String) error {
	_, err := q.Com().CallMethod("AddCollection", Collection, ReferenceName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCOLUMN_METHOD.html */
/* TODO: Some parameters are optional. Make sure to handle them correctly. */
func (q NotesQueryResultsProcessor) AddColumn(name domino.String, title domino.String, formula domino.String, sortOrder domino.Integer, hidden domino.Boolean, categorized domino.Boolean) error {
	_, err := q.Com().CallMethod("AddColumn", name, title, formula, sortOrder, hidden, categorized)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOMINOQUERY_METHOD.html */
func (q NotesQueryResultsProcessor) AddDominoQuery(query notesdominoquery.NotesDominoQuery, QueryString domino.String, ReferenceName domino.String) error {
	_, err := q.Com().CallMethod("AddDominoQuery", query.Com().Dispatch(), QueryString, ReferenceName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDFORMULA_METHOD.html */
func (q NotesQueryResultsProcessor) AddFormula(Formula domino.String, ColumnName domino.String, ReferenceName domino.String) error {
	_, err := q.Com().CallMethod("AddFormula", Formula, ColumnName, ReferenceName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTETOJSON_METHOD.html */
func (q NotesQueryResultsProcessor) ExecuteToJSON() (domino.String, error) {
	val, err := q.Com().CallMethod("ExecuteToJSON")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTETOVIEW_METHOD.html */
func (q NotesQueryResultsProcessor) ExecuteToView() error {
	_, err := q.Com().CallMethod("ExecuteToView")
	return err
}
