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

func newNotesView(dispatchPtr *ole.IDispatch) NotesView {
	return NotesView{newNotesStruct(dispatchPtr)}
}

func (v NotesView) checkKey(key any) (any, error) {
	if helpers.CheckTypeNames(key, []string{"string, number"}) != nil || helpers.CheckSliceTypeNames(key, []string{"string", "number"}) != nil {
		if casted, ok := key.(NotesDateTime); ok {
			key = casted
		} else if casted, ok := key.(NotesDateRange); ok {
			key = casted
		} else {
			return key, errors.New("key must be string, integer, long, double, string slice, number slice, NotesDateTime slice or NotesDateRange slice")
		}
	}
	return key, nil
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func (v NotesView) Aliases() ([]String, error) {
	return getComArrayProperty[String](v, "Aliases")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_VIEW.html */
func (v NotesView) SetAliases(aliases []String) error {
	return putComProperty(v, "Aliases", aliases)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETALLENTRIES_PROPERTY_6084.html */
func (v NotesView) AllEntries() (NotesViewEntryCollection, error) {
	return getComObjectProperty(v, newNotesViewEntryCollection, "AllEntries")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func (v NotesView) AutoUpdate() (Boolean, error) {
	return getComProperty[Boolean](v, "AutoUpdate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AUTOUPDATE_PROPERTY.html */
func (v NotesView) SetAutoUpdate(val Boolean) error {
	return putComProperty(v, "AutoUpdate", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func (v NotesView) BackgroundColor() (Long, error) {
	return getComProperty[Long](v, "BackgroundColor")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BACKGROUNDCOLOR_PROPERTY_4608.html */
func (v NotesView) SetBackgroundColor(val Long) error {
	return putComProperty(v, "BackgroundColor", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNCOUNT_PROPERTY_7753.html */
func (v NotesView) ColumnCount() (Long, error) {
	return getComProperty[Long](v, "ColumnCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNNAMES_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) ColumnNames() ([]String, error) {
	return getComArrayProperty[String](v, "ColumnNames")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNS_PROPERTY.html */
func (v NotesView) Columns() (NotesViewColumn, error) {
	return getComObjectProperty(v, newNotesViewColumn, "Columns")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATED_PROPERTY_VIEW.html */
func (v NotesView) Created() (Time, error) {
	return getComProperty[Time](v, "Created")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCOUNT_PROPERTY_VIEW.html */
func (v NotesView) EntryCount() (Long, error) {
	return getComProperty[Long](v, "EntryCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERLINES_PROPERTY_1209.html */
func (v NotesView) HeaderLines() (Long, error) {
	return getComProperty[Long](v, "HeaderLines")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) HttpURL() (String, error) {
	return getComProperty[String](v, "HttpURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCALENDAR_PROPERTY.html */
func (v NotesView) IsCalendar() (Boolean, error) {
	return getComProperty[Boolean](v, "IsCalendar")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORIZED_PROPERTY_2814.html */
func (v NotesView) IsCategorized() (Boolean, error) {
	return getComProperty[Boolean](v, "IsCategorized")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1242.html */
func (v NotesView) IsConflict() (Boolean, error) {
	return getComProperty[Boolean](v, "IsConflict")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func (v NotesView) IsDefaultView() (Boolean, error) {
	return getComProperty[Boolean](v, "IsDefaultView")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULTVIEW_PROPERTY.html */
func (v NotesView) SetIsDefaultView(val Boolean) error {
	return putComProperty(v, "IsDefaultView", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFOLDER_PROPERTY.html */
func (v NotesView) IsFolder() (Boolean, error) {
	return getComProperty[Boolean](v, "IsFolder")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIERARCHICAL_PROPERTY_7038.html */
func (v NotesView) IsHierarchical() (Boolean, error) {
	return getComProperty[Boolean](v, "IsHierarchical")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISMODIFIED_PROPERTY_6416.html */
func (v NotesView) IsModified() (Boolean, error) {
	return getComProperty[Boolean](v, "IsModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_VIEW.html */
func (v NotesView) IsPrivate() (Boolean, error) {
	return getComProperty[Boolean](v, "IsPrivate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func (v NotesView) IsProhibitDesignRefresh() (Boolean, error) {
	return getComProperty[Boolean](v, "IsProhibitDesignRefresh")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROHIBITDESIGNREFRESH_PROPERTY_VIEW.html */
func (v NotesView) SetIsProhibitDesignRefresh(val Boolean) error {
	return putComProperty(v, "IsProhibitDesignRefresh", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_VIEW.html */
func (v NotesView) LastModified() (Time, error) {
	return getComProperty[Time](v, "LastModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_VIEW.html */
func (v NotesView) LockHolders() (String, error) {
	return getComProperty[String](v, "LockHolders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func (v NotesView) Name() (String, error) {
	return getComProperty[String](v, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_VIEW.html */
func (v NotesView) SetName(val String) error {
	return putComProperty(v, "Name", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_NOTESVIEW_CLASS.html */
func (v NotesView) NotesURL() (String, error) {
	return getComProperty[String](v, "NotesURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_VIEW.html */
func (v NotesView) Parent() (NotesDatabase, error) {
	return getComObjectProperty(v, newNotesDatabase, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func (v NotesView) ProtectReaders() (Boolean, error) {
	return getComProperty[Boolean](v, "ProtectReaders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY_VIEW.html */
func (v NotesView) SetProtectReaders(val Boolean) error {
	return putComProperty(v, "ProtectReaders", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func (v NotesView) Readers() ([]String, error) {
	return getComArrayProperty[String](v, "Readers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY_VIEW.html */
func (v NotesView) SetReaders(readers []String) error {
	return putComProperty(v, "Readers", readers)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROWLINES_PROPERTY_4578.html */
func (v NotesView) RowLines() (Long, error) {
	return getComProperty[Long](v, "RowLines")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func (v NotesView) SelectionFormula() (String, error) {
	return getComProperty[String](v, "SelectionFormula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SELECTIONFORMULA_PROPERTY_VIEW.html */
func (v NotesView) SetSelectionFormula(formula String) error {
	return putComProperty(v, "SelectionFormula", formula)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func (v NotesView) Spacing() (Long, error) {
	return getComProperty[Long](v, "Spacing")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACING_PROPERTY_4165.html */
func (v NotesView) SetSpacing(val Long) error {
	return putComProperty(v, "Spacing", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TOPLEVELENTRYCOUNT_PROPERTY_6487.html */
func (v NotesView) TopLevelEntryCount() (Long, error) {
	return getComProperty[Long](v, "TopLevelEntryCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_VIEW.html */
func (v NotesView) UniversalID() (String, error) {
	return getComProperty[String](v, "UniversalID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWINHERITEDNAME_PROPERTY_VIEW.html */
func (v NotesView) ViewInheritedName() (String, error) {
	return getComProperty[String](v, "ViewInheritedName")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_VIEW.html */
func (v NotesView) Clear() error {
	return callComVoidMethod(v, "Clear")
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
	return callComObjectMethod(v, newNotesViewColumn, "CopyColumn", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewColumn, "CreateColumn", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNav", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNavFrom", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNavFromAllUnread", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNavFromCategory", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNavFromChildren", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNavFromDescendants", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewNavigator, "CreateViewNavMaxLevel", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_VIEW.html */
func (v NotesView) FTSearch(query String, maxDocs Integer) (Long, error) {
	return callComMethod[Long](v, "FTSearch", query, maxDocs)
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
	return callComObjectMethod(v, newNotesDocumentCollection, "GetAllDocumentsByKey", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewEntryCollection, "GetAllEntriesByKey", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewEntryCollection, "GetAllReadEntries", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewEntryCollection, "GetAllUnreadEntries", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD.html */
func (v NotesView) GetChild(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetChild", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCOLUMN_METHOD_NOTESVIEW_CLASS.html */
func (v NotesView) GetColumn(columnNumber Long) (NotesViewColumn, error) {
	return callComObjectMethod(v, newNotesViewColumn, "GetColumn", columnNumber)
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
	return callComObjectMethod(v, newNotesDocument, "GetDocumentByKey", paramsOrdered...)
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
	return callComObjectMethod(v, newNotesViewEntry, "GetEntryByKey", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetFirstDocument() (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetFirstDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetLastDocument() (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetLastDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetNextDocument(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetNextDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD.html */
func (v NotesView) GetNextSibling(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetNextSibling", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetNthDocument(index Long) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetNthDocument", index)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTDOCUMENT_METHOD.html */
func (v NotesView) GetParentDocument(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetParentDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEW.html */
func (v NotesView) GetPrevDocument(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetPrevDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD.html */
func (v NotesView) GetPrevSibling(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(v, newNotesDocument, "GetPrevSibling", document)
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
	return callComMethod[Boolean](v, "Lock", paramsOrdered...)
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
	return callComMethod[Boolean](v, "LockProvisional", paramsOrdered...)
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
	return callComVoidMethod(v, "MarkAllRead", paramsOrdered...)
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
	return callComVoidMethod(v, "MarkAllUnread", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REFRESH_METHOD_VIEW.html */
func (v NotesView) Refresh() error {
	return callComVoidMethod(v, "Refresh")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_VIEW.html */
func (v NotesView) Remove() error {
	return callComVoidMethod(v, "Remove")
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
	return callComVoidMethod(v, "ResortView", paramsOrdered...)
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
	return callComVoidMethod(v, "RemoveColumn", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_VIEW.html */
func (v NotesView) UnLock() error {
	return callComVoidMethod(v, "UnLock")
}
