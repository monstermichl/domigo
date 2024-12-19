/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEW_CLASS.html */
package domigo

import (
	"errors"

	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesView struct {
	NotesStruct
}

func NewNotesView(dispatchPtr *ole.IDispatch) NotesView {
	return NotesView{NewNotesStruct(dispatchPtr)}
}

func (v NotesView) checkKey(key any) (any, error) {
	if helpers.CheckTypeNames(key, []string{"string, number"}) != nil || helpers.CheckSliceTypeNames(key, []string{"string", "number"}) != nil {
		if casted, ok := key.(NotesDateTime); ok {
			key = casted.com().Dispatch()
		} else if casted, ok := key.(NotesDateRange); ok {
			key = casted.com().Dispatch()
		} else {
			return key, errors.New("key must be string, integer, long, double, string slice, number slice, NotesDateTime slice or NotesDateRange slice")
		}
	}
	return key, nil
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func (v NotesView) Aliases() ([]String, error) {
	vals, err := v.com().GetArrayProperty("Aliases")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func (v NotesView) SetAliases(aliases []String) error {
	return v.com().PutProperty("Aliases", aliases)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIES_PROPERTY_6084.html */
func (v NotesView) AllEntries() (NotesViewEntryCollection, error) {
	dispatchPtr, err := v.com().GetObjectProperty("AllEntries")
	return NewNotesViewEntryCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func (v NotesView) AutoUpdate() (Boolean, error) {
	val, err := v.com().GetProperty("AutoUpdate")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func (v NotesView) SetAutoUpdate(val Boolean) error {
	return v.com().PutProperty("AutoUpdate", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func (v NotesView) BackgroundColor() (Long, error) {
	val, err := v.com().GetProperty("BackgroundColor")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func (v NotesView) SetBackgroundColor(val Long) error {
	return v.com().PutProperty("BackgroundColor", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNCOUNT_PROPERTY_7753.html */
func (v NotesView) ColumnCount() (Long, error) {
	val, err := v.com().GetProperty("ColumnCount")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNNAMES_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) ColumnNames() ([]String, error) {
	vals, err := v.com().GetArrayProperty("ColumnNames")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNS_PROPERTY.html */
func (v NotesView) Columns() (NotesViewColumn, error) {
	dispatchPtr, err := v.com().GetObjectProperty("Columns")
	return NewNotesViewColumn(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_VIEW.html */
func (v NotesView) Created() (Time, error) {
	val, err := v.com().GetProperty("Created")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCOUNT_PROPERTY_VIEW.html */
func (v NotesView) EntryCount() (Long, error) {
	val, err := v.com().GetProperty("EntryCount")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERLINES_PROPERTY_1209.html */
func (v NotesView) HeaderLines() (Long, error) {
	val, err := v.com().GetProperty("HeaderLines")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) HttpURL() (String, error) {
	val, err := v.com().GetProperty("HttpURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCALENDAR_PROPERTY.html */
func (v NotesView) IsCalendar() (Boolean, error) {
	val, err := v.com().GetProperty("IsCalendar")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORIZED_PROPERTY_2814.html */
func (v NotesView) IsCategorized() (Boolean, error) {
	val, err := v.com().GetProperty("IsCategorized")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1242.html */
func (v NotesView) IsConflict() (Boolean, error) {
	val, err := v.com().GetProperty("IsConflict")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func (v NotesView) IsDefaultView() (Boolean, error) {
	val, err := v.com().GetProperty("IsDefaultView")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func (v NotesView) SetIsDefaultView(val Boolean) error {
	return v.com().PutProperty("IsDefaultView", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFOLDER_PROPERTY.html */
func (v NotesView) IsFolder() (Boolean, error) {
	val, err := v.com().GetProperty("IsFolder")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIERARCHICAL_PROPERTY_7038.html */
func (v NotesView) IsHierarchical() (Boolean, error) {
	val, err := v.com().GetProperty("IsHierarchical")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMODIFIED_PROPERTY_6416.html */
func (v NotesView) IsModified() (Boolean, error) {
	val, err := v.com().GetProperty("IsModified")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_VIEW.html */
func (v NotesView) IsPrivate() (Boolean, error) {
	val, err := v.com().GetProperty("IsPrivate")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func (v NotesView) IsProhibitDesignRefresh() (Boolean, error) {
	val, err := v.com().GetProperty("IsProhibitDesignRefresh")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func (v NotesView) SetIsProhibitDesignRefresh(val Boolean) error {
	return v.com().PutProperty("IsProhibitDesignRefresh", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_VIEW.html */
func (v NotesView) LastModified() (Time, error) {
	val, err := v.com().GetProperty("LastModified")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_VIEW.html */
func (v NotesView) LockHolders() (String, error) {
	val, err := v.com().GetProperty("LockHolders")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func (v NotesView) Name() (String, error) {
	val, err := v.com().GetProperty("Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func (v NotesView) SetName(val String) error {
	return v.com().PutProperty("Name", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) NotesURL() (String, error) {
	val, err := v.com().GetProperty("NotesURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_VIEW.html */
func (v NotesView) Parent() (NotesDatabase, error) {
	dispatchPtr, err := v.com().GetObjectProperty("Parent")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func (v NotesView) ProtectReaders() (Boolean, error) {
	val, err := v.com().GetProperty("ProtectReaders")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func (v NotesView) SetProtectReaders(val Boolean) error {
	return v.com().PutProperty("ProtectReaders", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func (v NotesView) Readers() ([]String, error) {
	vals, err := v.com().GetArrayProperty("Readers")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func (v NotesView) SetReaders(readers []String) error {
	return v.com().PutProperty("Readers", readers)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROWLINES_PROPERTY_4578.html */
func (v NotesView) RowLines() (Long, error) {
	val, err := v.com().GetProperty("RowLines")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func (v NotesView) SelectionFormula() (String, error) {
	val, err := v.com().GetProperty("SelectionFormula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func (v NotesView) SetSelectionFormula(formula String) error {
	return v.com().PutProperty("SelectionFormula", formula)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func (v NotesView) Spacing() (Long, error) {
	val, err := v.com().GetProperty("Spacing")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func (v NotesView) SetSpacing(val Long) error {
	return v.com().PutProperty("Spacing", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TOPLEVELENTRYCOUNT_PROPERTY_6487.html */
func (v NotesView) TopLevelEntryCount() (Long, error) {
	val, err := v.com().GetProperty("TopLevelEntryCount")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_VIEW.html */
func (v NotesView) UniversalID() (String, error) {
	val, err := v.com().GetProperty("UniversalID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWINHERITEDNAME_PROPERTY_VIEW.html */
func (v NotesView) ViewInheritedName() (String, error) {
	val, err := v.com().GetProperty("ViewInheritedName")
	return helpers.CastValue[String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_VIEW.html */
func (v NotesView) Clear() error {
	_, err := v.com().CallMethod("Clear")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COPYCOLUMN_METHOD_VIEW.html */
type notesViewCopyColumnParams struct {
	destinationIndex *Long
}

type notesViewCopyColumnParam func(*notesViewCopyColumnParams)

func WithNotesViewCopyColumnDestinationIndex(destinationIndex Long) notesViewCopyColumnParam {
	return func(c *notesViewCopyColumnParams) {
		c.destinationIndex = &destinationIndex
	}
}

/* TODO: Handle different types. */
func (v NotesView) CopyColumn(sourceColumn Integer, params ...notesViewCopyColumnParam) (NotesViewColumn, error) {
	paramsStruct := &notesViewCopyColumnParams{}
	paramsOrdered := []interface{}{sourceColumn}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.destinationIndex != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.destinationIndex)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CopyColumn", paramsOrdered...)
	return NewNotesViewColumn(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLUMN_METHOD_VIEW.html */
type notesViewCreateColumnParams struct {
	position   *Long
	columnName *String
	formula    *String
}

type notesViewCreateColumnParam func(*notesViewCreateColumnParams)

func WithNotesViewCreateColumnPosition(position Long) notesViewCreateColumnParam {
	return func(c *notesViewCreateColumnParams) {
		c.position = &position
	}
}

func WithNotesViewCreateColumnColumnName(columnName String) notesViewCreateColumnParam {
	return func(c *notesViewCreateColumnParams) {
		c.columnName = &columnName
	}
}

func WithNotesViewCreateColumnFormula(formula String) notesViewCreateColumnParam {
	return func(c *notesViewCreateColumnParams) {
		c.formula = &formula
	}
}

func (v NotesView) CreateColumn(params ...notesViewCreateColumnParam) (NotesViewColumn, error) {
	paramsStruct := &notesViewCreateColumnParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.position != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.position)
		if paramsStruct.columnName != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.columnName)
			if paramsStruct.formula != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.formula)
			}
		}
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateColumn", paramsOrdered...)
	return NewNotesViewColumn(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAV_METHOD_1631.html */
type notesViewCreateViewNavParams struct {
	cacheSize *Long
}

type notesViewCreateViewNavParam func(*notesViewCreateViewNavParams)

func WithNotesViewCreateViewNavCacheSize(cacheSize Long) notesViewCreateViewNavParam {
	return func(c *notesViewCreateViewNavParams) {
		c.cacheSize = &cacheSize
	}
}

func (v NotesView) CreateViewNav(params ...notesViewCreateViewNavParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNav", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROM_METHOD_5742.html */
type notesViewCreateViewNavFromParams struct {
	cacheSize *Long
}

type notesViewCreateViewNavFromParam func(*notesViewCreateViewNavFromParams)

func WithNotesViewCreateViewNavFromCacheSize(cacheSize Long) notesViewCreateViewNavFromParam {
	return func(c *notesViewCreateViewNavFromParams) {
		c.cacheSize = &cacheSize
	}
}

func (v NotesView) CreateViewNavFrom(navigatorObject Variant, params ...notesViewCreateViewNavFromParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavFromParams{}
	paramsOrdered := []interface{}{navigatorObject}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNavFrom", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMALLUNREAD.html */
type notesViewCreateViewNavFromAllUnreadParams struct {
	username *String
}

type notesViewCreateViewNavFromAllUnreadParam func(*notesViewCreateViewNavFromAllUnreadParams)

func WithNotesViewCreateViewNavFromAllUnreadUsername(username String) notesViewCreateViewNavFromAllUnreadParam {
	return func(c *notesViewCreateViewNavFromAllUnreadParams) {
		c.username = &username
	}
}

func (v NotesView) CreateViewNavFromAllUnread(params ...notesViewCreateViewNavFromAllUnreadParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavFromAllUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNavFromAllUnread", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMCATEGORY_METHOD_3595.html */
type notesViewCreateViewNavFromCategoryParams struct {
	cacheSize *Long
}

type notesViewCreateViewNavFromCategoryParam func(*notesViewCreateViewNavFromCategoryParams)

func WithNotesViewCreateViewNavFromCategoryCacheSize(cacheSize Long) notesViewCreateViewNavFromCategoryParam {
	return func(c *notesViewCreateViewNavFromCategoryParams) {
		c.cacheSize = &cacheSize
	}
}

func (v NotesView) CreateViewNavFromCategory(category String, params ...notesViewCreateViewNavFromCategoryParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavFromCategoryParams{}
	paramsOrdered := []interface{}{category}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNavFromCategory", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMCHILDREN_METHOD_9100.html */
type notesViewCreateViewNavFromChildrenParams struct {
	cacheSize *Long
}

type notesViewCreateViewNavFromChildrenParam func(*notesViewCreateViewNavFromChildrenParams)

func WithNotesViewCreateViewNavFromChildrenCacheSize(cacheSize Long) notesViewCreateViewNavFromChildrenParam {
	return func(c *notesViewCreateViewNavFromChildrenParams) {
		c.cacheSize = &cacheSize
	}
}

func (v NotesView) CreateViewNavFromChildren(navigatorObject Variant, params ...notesViewCreateViewNavFromChildrenParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavFromChildrenParams{}
	paramsOrdered := []interface{}{navigatorObject}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNavFromChildren", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMDESCENDANTS_METHOD_2893.html */
type notesViewCreateViewNavFromDescendantsParams struct {
	cacheSize *Long
}

type notesViewCreateViewNavFromDescendantsParam func(*notesViewCreateViewNavFromDescendantsParams)

func WithNotesViewCreateViewNavFromDescendantsCacheSize(cacheSize Long) notesViewCreateViewNavFromDescendantsParam {
	return func(c *notesViewCreateViewNavFromDescendantsParams) {
		c.cacheSize = &cacheSize
	}
}

func (v NotesView) CreateViewNavFromDescendants(navigatorObject Variant, params ...notesViewCreateViewNavFromDescendantsParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavFromDescendantsParams{}
	paramsOrdered := []interface{}{navigatorObject}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNavFromDescendants", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVMAXLEVEL_METHOD_NOTESVIEW_CLASS.html */
type notesViewCreateViewNavMaxLevelParams struct {
	cacheSize *Long
}

type notesViewCreateViewNavMaxLevelParam func(*notesViewCreateViewNavMaxLevelParams)

func WithNotesViewCreateViewNavMaxLevelCacheSize(cacheSize Long) notesViewCreateViewNavMaxLevelParam {
	return func(c *notesViewCreateViewNavMaxLevelParams) {
		c.cacheSize = &cacheSize
	}
}

func (v NotesView) CreateViewNavMaxLevel(level Long, params ...notesViewCreateViewNavMaxLevelParam) (NotesViewNavigator, error) {
	paramsStruct := &notesViewCreateViewNavMaxLevelParams{}
	paramsOrdered := []interface{}{level}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.com().CallObjectMethod("CreateViewNavMaxLevel", paramsOrdered...)
	return NewNotesViewNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_VIEW.html */
func (v NotesView) FTSearch(query String, maxDocs Integer) (Long, error) {
	val, err := v.com().CallMethod("FTSearch", query, maxDocs)
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLDOCUMENTSBYKEY_METHOD.html */
type notesViewGetAllDocumentsByKeyParams struct {
	exactMatch *Boolean
}

type notesViewGetAllDocumentsByKeyParam func(*notesViewGetAllDocumentsByKeyParams)

func WithNotesViewGetAllDocumentsByKeyExactMatch(exactMatch Boolean) notesViewGetAllDocumentsByKeyParam {
	return func(c *notesViewGetAllDocumentsByKeyParams) {
		c.exactMatch = &exactMatch
	}
}

func (v NotesView) GetAllDocumentsByKey(key any, params ...notesViewGetAllDocumentsByKeyParam) (NotesDocumentCollection, error) {
	key, err := v.checkKey(key)

	if err != nil {
		return NotesDocumentCollection{}, err
	}
	paramsStruct := &notesViewGetAllDocumentsByKeyParams{}
	paramsOrdered := []interface{}{key}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.exactMatch != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.exactMatch)
	}
	dispatchPtr, err := v.com().CallObjectMethod("GetAllDocumentsByKey", paramsOrdered...)
	return NewNotesDocumentCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIESBYKEY_METHOD_9837.html */
type notesViewGetAllEntriesByKeyParams struct {
	exactMatch *Boolean
}

type notesViewGetAllEntriesByKeyParam func(*notesViewGetAllEntriesByKeyParams)

func WithNotesViewGetAllEntriesByKeyExactMatch(exactMatch Boolean) notesViewGetAllEntriesByKeyParam {
	return func(c *notesViewGetAllEntriesByKeyParams) {
		c.exactMatch = &exactMatch
	}
}

func (v NotesView) GetAllEntriesByKey(key any, params ...notesViewGetAllEntriesByKeyParam) (NotesViewEntryCollection, error) {
	key, err := v.checkKey(key)

	if err != nil {
		return NotesViewEntryCollection{}, err
	}
	paramsStruct := &notesViewGetAllEntriesByKeyParams{}
	paramsOrdered := []interface{}{key}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.exactMatch != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.exactMatch)
	}
	dispatchPtr, err := v.com().CallObjectMethod("GetAllEntriesByKey", paramsOrdered...)
	return NewNotesViewEntryCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLREADENTRIES.html */
type notesViewGetAllReadEntriesParams struct {
	username *String
}

type notesViewGetAllReadEntriesParam func(*notesViewGetAllReadEntriesParams)

func WithNotesViewGetAllReadEntriesUsername(username String) notesViewGetAllReadEntriesParam {
	return func(c *notesViewGetAllReadEntriesParams) {
		c.username = &username
	}
}

func (v NotesView) GetAllReadEntries(params ...notesViewGetAllReadEntriesParam) (NotesViewEntryCollection, error) {
	paramsStruct := &notesViewGetAllReadEntriesParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	dispatchPtr, err := v.com().CallObjectMethod("GetAllReadEntries", paramsOrdered...)
	return NewNotesViewEntryCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLUNREADENTRIES.html */
type notesViewGetAllUnreadEntriesParams struct {
	username *String
}

type notesViewGetAllUnreadEntriesParam func(*notesViewGetAllUnreadEntriesParams)

func WithNotesViewGetAllUnreadEntriesUsername(username String) notesViewGetAllUnreadEntriesParam {
	return func(c *notesViewGetAllUnreadEntriesParams) {
		c.username = &username
	}
}

func (v NotesView) GetAllUnreadEntries(params ...notesViewGetAllUnreadEntriesParam) (NotesViewEntryCollection, error) {
	paramsStruct := &notesViewGetAllUnreadEntriesParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	dispatchPtr, err := v.com().CallObjectMethod("GetAllUnreadEntries", paramsOrdered...)
	return NewNotesViewEntryCollection(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD.html */
func (v NotesView) GetChild(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetChild", document.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCOLUMN_METHOD_NOTESVIEW_CLASS.html */
func (v NotesView) GetColumn(columnNumber Long) (NotesViewColumn, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetColumn", columnNumber)
	return NewNotesViewColumn(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENTBYKEY_METHOD.html */
type notesViewGetDocumentByKeyParams struct {
	exactMatch *Boolean
}

type notesViewGetDocumentByKeyParam func(*notesViewGetDocumentByKeyParams)

func WithNotesViewGetDocumentByKeyExactMatch(exactMatch Boolean) notesViewGetDocumentByKeyParam {
	return func(c *notesViewGetDocumentByKeyParams) {
		c.exactMatch = &exactMatch
	}
}

func (v NotesView) GetDocumentByKey(key any, params ...notesViewGetDocumentByKeyParam) (NotesDocument, error) {
	key, err := v.checkKey(key)

	if err != nil {
		return NotesDocument{}, err
	}
	paramsStruct := &notesViewGetDocumentByKeyParams{}
	paramsOrdered := []interface{}{key}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.exactMatch != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.exactMatch)
	}
	dispatchPtr, err := v.com().CallObjectMethod("GetDocumentByKey", paramsOrdered...)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYKEY_METHOD_3846.html */
type notesViewGetEntryByKeyParams struct {
	exactMatch *Boolean
}

type notesViewGetEntryByKeyParam func(*notesViewGetEntryByKeyParams)

func WithNotesViewGetEntryByKeyExactMatch(exactMatch Boolean) notesViewGetEntryByKeyParam {
	return func(c *notesViewGetEntryByKeyParams) {
		c.exactMatch = &exactMatch
	}
}

func (v NotesView) GetEntryByKey(key any, params ...notesViewGetEntryByKeyParam) (NotesViewEntry, error) {
	key, err := v.checkKey(key)

	if err != nil {
		return NotesViewEntry{}, err
	}
	paramsStruct := &notesViewGetEntryByKeyParams{}
	paramsOrdered := []interface{}{key}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.exactMatch != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.exactMatch)
	}
	dispatchPtr, err := v.com().CallObjectMethod("GetEntryByKey", paramsOrdered...)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetFirstDocument() (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetFirstDocument")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetLastDocument() (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetLastDocument")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetNextDocument(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetNextDocument", document.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD.html */
func (v NotesView) GetNextSibling(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetNextSibling", document.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetNthDocument(index Long) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetNthDocument", index)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTDOCUMENT_METHOD.html */
func (v NotesView) GetParentDocument(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetParentDocument", document.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetPrevDocument(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetPrevDocument", document.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD.html */
func (v NotesView) GetPrevSibling(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetPrevSibling", document.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_VIEW.html */
type notesViewLockParams struct {
	name          *[]String
	provisionalOK *Boolean
}

type notesViewLockParam func(*notesViewLockParams)

func WithNotesViewLockName(name []String) notesViewLockParam {
	return func(c *notesViewLockParams) {
		c.name = &name
	}
}

func WithNotesViewLockProvisionalOK(provisionalOK Boolean) notesViewLockParam {
	return func(c *notesViewLockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (v NotesView) Lock(params ...notesViewLockParam) (Boolean, error) {
	paramsStruct := &notesViewLockParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
		if paramsStruct.provisionalOK != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.provisionalOK)
		}
	}
	val, err := v.com().CallMethod("Lock", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_VIEW.html */
type notesViewLockProvisionalParams struct {
	name *[]String
}

type notesViewLockProvisionalParam func(*notesViewLockProvisionalParams)

func WithNotesViewLockProvisionalName(name []String) notesViewLockProvisionalParam {
	return func(c *notesViewLockProvisionalParams) {
		c.name = &name
	}
}

func (v NotesView) LockProvisional(params ...notesViewLockProvisionalParam) (Boolean, error) {
	paramsStruct := &notesViewLockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	val, err := v.com().CallMethod("LockProvisional", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEW.html */
type notesViewMarkAllReadParams struct {
	username *String
}

type notesViewMarkAllReadParam func(*notesViewMarkAllReadParams)

func WithNotesViewMarkAllReadUsername(username String) notesViewMarkAllReadParam {
	return func(c *notesViewMarkAllReadParams) {
		c.username = &username
	}
}

func (v NotesView) MarkAllRead(params ...notesViewMarkAllReadParam) error {
	paramsStruct := &notesViewMarkAllReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := v.com().CallMethod("MarkAllRead", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEW.html */
type notesViewMarkAllUnreadParams struct {
	username *String
}

type notesViewMarkAllUnreadParam func(*notesViewMarkAllUnreadParams)

func WithNotesViewMarkAllUnreadUsername(username String) notesViewMarkAllUnreadParam {
	return func(c *notesViewMarkAllUnreadParams) {
		c.username = &username
	}
}

func (v NotesView) MarkAllUnread(params ...notesViewMarkAllUnreadParam) error {
	paramsStruct := &notesViewMarkAllUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := v.com().CallMethod("MarkAllUnread", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESH_METHOD_VIEW.html */
func (v NotesView) Refresh() error {
	_, err := v.com().CallMethod("Refresh")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_VIEW.html */
func (v NotesView) Remove() error {
	_, err := v.com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTVIEW_METHOD_VIEW.html */
type notesViewResortViewParams struct {
	columnName    *String
	ascendingFlag *Boolean
}

type notesViewResortViewParam func(*notesViewResortViewParams)

func WithNotesViewResortViewColumnName(columnName String) notesViewResortViewParam {
	return func(c *notesViewResortViewParams) {
		c.columnName = &columnName
	}
}

func WithNotesViewResortViewAscendingFlag(ascendingFlag Boolean) notesViewResortViewParam {
	return func(c *notesViewResortViewParams) {
		c.ascendingFlag = &ascendingFlag
	}
}

func (v NotesView) ResortView(params ...notesViewResortViewParam) error {
	paramsStruct := &notesViewResortViewParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.columnName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.columnName)
		if paramsStruct.ascendingFlag != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.ascendingFlag)
		}
	}
	_, err := v.com().CallMethod("ResortView", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVECOLUMN_METHOD_VIEW.html */
type notesViewRemoveColumnParams struct {
	columnNameOrIndex *Long
}

type notesViewRemoveColumnParam func(*notesViewRemoveColumnParams)

func WithNotesViewRemoveColumnColumnNameOrIndex(columnNameOrIndex Long) notesViewRemoveColumnParam {
	return func(c *notesViewRemoveColumnParams) {
		c.columnNameOrIndex = &columnNameOrIndex
	}
}

func (v NotesView) RemoveColumn(params ...notesViewRemoveColumnParam) error {
	paramsStruct := &notesViewRemoveColumnParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.columnNameOrIndex != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.columnNameOrIndex)
	} else {
		paramsOrdered = append(paramsOrdered, nil)
	}
	_, err := v.com().CallMethod("RemoveColumn", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_VIEW.html */
func (v NotesView) UnLock() error {
	_, err := v.com().CallMethod("UnLock")
	return err
}
