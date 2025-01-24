/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESQUERYRESULTSPROCESSOR_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesQueryResultsProcessor struct {
	NotesStruct
}

func newNotesQueryResultsProcessor(dispatchPtr *ole.IDispatch) NotesQueryResultsProcessor {
	return NotesQueryResultsProcessor{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func (q NotesQueryResultsProcessor) MaxEntries() (Long, error) {
	return getComProperty[Long](q, "MaxEntries")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func (q NotesQueryResultsProcessor) SetMaxEntries(v Long) error {
	return putComProperty(q, "MaxEntries", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSECQRP_PROPERTY.html */
func (q NotesQueryResultsProcessor) TimeOutSec() (Integer, error) {
	return getComProperty[Integer](q, "TimeOutSec")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSECQRP_PROPERTY.html */
func (q NotesQueryResultsProcessor) SetTimeOutSec(v Integer) error {
	return putComProperty(q, "TimeOutSec", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCOLLECTION_METHOD.html */
func (q NotesQueryResultsProcessor) AddCollection(Collection Variant, ReferenceName String) error {
	return callComVoidMethod(q, "AddCollection", Collection, ReferenceName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCOLUMN_METHOD.html */
type notesQueryResultsProcessorAddColumnParams struct {
	title       *String
	formula     *String
	sortOrder   *Integer
	hidden      *Boolean
	categorized *Boolean
}

type notesQueryResultsProcessorAddColumnParam func(*notesQueryResultsProcessorAddColumnParams)

func WithNotesQueryResultsProcessorAddColumnTitle(title String) notesQueryResultsProcessorAddColumnParam {
	return func(c *notesQueryResultsProcessorAddColumnParams) {
		c.title = &title
	}
}

func WithNotesQueryResultsProcessorAddColumnFormula(formula String) notesQueryResultsProcessorAddColumnParam {
	return func(c *notesQueryResultsProcessorAddColumnParams) {
		c.formula = &formula
	}
}

func WithNotesQueryResultsProcessorAddColumnSortOrder(sortOrder Integer) notesQueryResultsProcessorAddColumnParam {
	return func(c *notesQueryResultsProcessorAddColumnParams) {
		c.sortOrder = &sortOrder
	}
}

func WithNotesQueryResultsProcessorAddColumnHidden(hidden Boolean) notesQueryResultsProcessorAddColumnParam {
	return func(c *notesQueryResultsProcessorAddColumnParams) {
		c.hidden = &hidden
	}
}

func WithNotesQueryResultsProcessorAddColumnCategorized(categorized Boolean) notesQueryResultsProcessorAddColumnParam {
	return func(c *notesQueryResultsProcessorAddColumnParams) {
		c.categorized = &categorized
	}
}

func (q NotesQueryResultsProcessor) AddColumn(name String, params ...notesQueryResultsProcessorAddColumnParam) error {
	paramsStruct := &notesQueryResultsProcessorAddColumnParams{}
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
	return callComVoidMethod(q, "AddColumn", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOMINOQUERY_METHOD.html */
func (q NotesQueryResultsProcessor) AddDominoQuery(query NotesDominoQuery, queryString String, referenceName String) error {
	return callComVoidMethod(q, "AddDominoQuery", query, queryString, referenceName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDFORMULA_METHOD.html */
func (q NotesQueryResultsProcessor) AddFormula(formula String, columnName String, referenceName String) error {
	return callComVoidMethod(q, "AddFormula", formula, columnName, referenceName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTETOJSON_METHOD.html */
func (q NotesQueryResultsProcessor) ExecuteToJSON() (String, error) {
	return callComMethod[String](q, "ExecuteToJSON")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTETOVIEW_METHOD.html */
type notesQueryResultsProcessorExecuteToViewParams struct {
	expireHours       *Long
	readers           *[]String
	designSrcDB       *NotesDatabase
	designSrcViewName *String
}

type notesQueryResultsProcessorExecuteToViewParam func(*notesQueryResultsProcessorExecuteToViewParams)

func WithNotesQueryResultsProcessorExecuteToViewExpireHours(expireHours Long) notesQueryResultsProcessorExecuteToViewParam {
	return func(c *notesQueryResultsProcessorExecuteToViewParams) {
		c.expireHours = &expireHours
	}
}

func WithNotesQueryResultsProcessorExecuteToViewReaders(readers []String) notesQueryResultsProcessorExecuteToViewParam {
	return func(c *notesQueryResultsProcessorExecuteToViewParams) {
		c.readers = &readers
	}
}

func WithNotesQueryResultsProcessorExecuteToViewDesignSrcDB(designSrcDB NotesDatabase) notesQueryResultsProcessorExecuteToViewParam {
	return func(c *notesQueryResultsProcessorExecuteToViewParams) {
		c.designSrcDB = &designSrcDB
	}
}

func WithNotesQueryResultsProcessorExecuteToViewDesignSrcViewName(designSrcViewName String) notesQueryResultsProcessorExecuteToViewParam {
	return func(c *notesQueryResultsProcessorExecuteToViewParams) {
		c.designSrcViewName = &designSrcViewName
	}
}

func (q NotesQueryResultsProcessor) ExecuteToView(name String, params ...notesQueryResultsProcessorExecuteToViewParam) error {
	paramsStruct := &notesQueryResultsProcessorExecuteToViewParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.expireHours != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.expireHours)

		if paramsStruct.readers != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.readers)
		} else {
			paramsOrdered = append(paramsOrdered, nil) /* According to the examples in the documentation, it's allowed to pass null/nil. */
		}

		if paramsStruct.designSrcDB != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.designSrcDB)
			if paramsStruct.designSrcViewName != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.designSrcViewName)
			}
		}
	}
	return callComVoidMethod(q, "ExecuteToView", paramsOrdered...)
}
