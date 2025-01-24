/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDXLIMPORTER_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
	"github.com/monstermichl/domigo/internal/helpers"
)

type NotesDXLImporter struct {
	NotesStruct
}

func newNotesDXLImporter(dispatchPtr *ole.IDispatch) NotesDXLImporter {
	return NotesDXLImporter{newNotesStruct(dispatchPtr)}
}

func (d NotesDXLImporter) checkImportTypes(val any) error {
	return helpers.CheckTypeNames(val, []string{"string", "NotesRichTextItem", "NotesStream"})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLIMPORTOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) ACLImportOption() (Long, error) {
	return getComProperty[Long](d, "ACLImportOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLIMPORTOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetACLImportOption(v Long) error {
	return putComProperty(d, "ACLImportOption", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPILELOTUSSCRIPT_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) CompileLotusScript() (Boolean, error) {
	return getComProperty[Boolean](d, "CompileLotusScript")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPILELOTUSSCRIPT_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetCompileLotusScript(v Boolean) error {
	return putComProperty(d, "CompileLotusScript", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) CreateFTIndex() (Boolean, error) {
	return getComProperty[Boolean](d, "CreateFTIndex")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetCreateFTIndex(v Boolean) error {
	return putComProperty(d, "CreateFTIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNIMPORTOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) DesignImportOption() (Long, error) {
	return getComProperty[Long](d, "DesignImportOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNIMPORTOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetDesignImportOption(v Long) error {
	return putComProperty(d, "DesignImportOption", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTIMPORTOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) DocumentImportOption() (Long, error) {
	return getComProperty[Long](d, "DocumentImportOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTIMPORTOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetDocumentImportOption(v Long) error {
	return putComProperty(d, "DocumentImportOption", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_DXLIMPORTER.html */
func (d NotesDXLImporter) ImportedNoteCount() (Long, error) {
	return getComProperty[Long](d, "ImportedNoteCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INPUTVALIDATIONOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) InputValidationOption() (Long, error) {
	return getComProperty[Long](d, "InputValidationOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INPUTVALIDATIONOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetInputValidationOption(v Long) error {
	return putComProperty(d, "InputValidationOption", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEDBPROPERTIES_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) ReplaceDbProperties() (Boolean, error) {
	return getComProperty[Boolean](d, "ReplaceDbProperties")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEDBPROPERTIES_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetReplaceDbProperties(v Boolean) error {
	return putComProperty(d, "ReplaceDbProperties", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAREQUIREDFORREPLACEORUPDATE_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) ReplicaRequiredForReplaceOrUpdate() (Boolean, error) {
	return getComProperty[Boolean](d, "ReplicaRequiredForReplaceOrUpdate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAREQUIREDFORREPLACEORUPDATE_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetReplicaRequiredForReplaceOrUpdate(v Boolean) error {
	return putComProperty(d, "ReplicaRequiredForReplaceOrUpdate", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNKNOWNTOKENLOGOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) UnknownTokenLogOption() (Long, error) {
	return getComProperty[Long](d, "UnknownTokenLogOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNKNOWNTOKENLOGOPTION_PROPERTY_IMPORTER.html */
func (d NotesDXLImporter) SetUnknownTokenLogOption(v Long) error {
	return putComProperty(d, "UnknownTokenLogOption", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTNOTEID_METHOD_DXLIMPORTER.html */
func (d NotesDXLImporter) GetFirstImportedNoteID() (String, error) {
	return callComMethod[String](d, "GetFirstImportedNoteID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_DXLIMPORTER.html */
func (d NotesDXLImporter) GetNextImportedNoteID(noteID String) (String, error) {
	return callComMethod[String](d, "GetNextImportedNoteID", noteID)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMPORT_METHOD_IMPORTER.html */
func (d NotesDXLImporter) Import(input any, importDb NotesDatabase) error {
	err := d.checkImportTypes(input)

	if err != nil {
		return err
	}
	return callComVoidMethod(d, "Import", input, importDb)
}
