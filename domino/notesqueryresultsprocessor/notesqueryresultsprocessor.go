/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESQUERYRESULTSPROCESSOR_CLASS.html */
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
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func (q NotesQueryResultsProcessor) MaxEntries() (domino.Long, error) {
	val, err := q.Com().GetProperty("MaxEntries")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func (q NotesQueryResultsProcessor) SetMaxEntries(v domino.Long) error {
	return q.Com().PutProperty("MaxEntries", v)
}

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
type addColumnParams struct {
	title       *domino.String
	formula     *domino.String
	sortOrder   *domino.Integer
	hidden      *domino.Boolean
	categorized *[]domino.Boolean
}

type addColumnParam func(*addColumnParams)

func WithAddColumnTitle(title domino.String) addColumnParam {
	return func(c *addColumnParams) {
		c.title = &title
	}
}

func WithAddColumnFormula(formula domino.String) addColumnParam {
	return func(c *addColumnParams) {
		c.formula = &formula
	}
}

func WithAddColumnSortOrder(sortOrder domino.Integer) addColumnParam {
	return func(c *addColumnParams) {
		c.sortOrder = &sortOrder
	}
}

func WithAddColumnHidden(hidden domino.Boolean) addColumnParam {
	return func(c *addColumnParams) {
		c.hidden = &hidden
	}
}

func WithAddColumnCategorized(categorized []domino.Boolean) addColumnParam {
	return func(c *addColumnParams) {
		c.categorized = &categorized
	}
}

func (q NotesQueryResultsProcessor) AddColumn(name domino.String, params ...addColumnParam) error {
	paramsStruct := &addColumnParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.title != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.title)
		if paramsStruct.formula != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.formula)
			if paramsStruct.sortOrder != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.sortOrder)
				if paramsStruct.hidden != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.hidden)
					if paramsStruct.categorized != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.categorized)
					}
				}
			}
		}
	}
	_, err := q.Com().CallMethod("AddColumn", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOMINOQUERY_METHOD.html */
func (q NotesQueryResultsProcessor) AddDominoQuery(query notesdominoquery.NotesDominoQuery, queryString domino.String, referenceName domino.String) error {
	_, err := q.Com().CallMethod("AddDominoQuery", query.Com().Dispatch(), queryString, referenceName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDFORMULA_METHOD.html */
func (q NotesQueryResultsProcessor) AddFormula(formula domino.String, columnName domino.String, referenceName domino.String) error {
	_, err := q.Com().CallMethod("AddFormula", formula, columnName, referenceName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTETOJSON_METHOD.html */
func (q NotesQueryResultsProcessor) ExecuteToJSON() (domino.String, error) {
	val, err := q.Com().CallMethod("ExecuteToJSON")
	return helpers.CastValue[domino.String](val), err
}
