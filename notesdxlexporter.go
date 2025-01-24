/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDXLEXPORTER_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
	"github.com/monstermichl/domigo/internal/helpers"
)

type NotesDXLExporter struct {
	NotesStruct
}

func newNotesDXLExporter(dispatchPtr *ole.IDispatch) NotesDXLExporter {
	return NotesDXLExporter{newNotesStruct(dispatchPtr)}
}

func (d NotesDXLExporter) checkImportTypes(val any) error {
	return helpers.CheckTypeNames(val, []string{"NotesDatabase", "NotesDocument", "NotesDocumentCollection", "NotesNoteCollection"})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHMENTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) AttachmentOmittedText() (String, error) {
	return getComProperty[String](d, "AttachmentOmittedText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHMENTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetAttachmentOmittedText(v String) error {
	return putComProperty(d, "AttachmentOmittedText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTNOTESBITMAPSTOGIF_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) ConvertNotesBitmapsToGIF() (Boolean, error) {
	return getComProperty[Boolean](d, "ConvertNotesBitmapsToGIF")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTNOTESBITMAPSTOGIF_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetConvertNotesBitmapsToGIF(v Boolean) error {
	return putComProperty(d, "ConvertNotesBitmapsToGIF", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCTYPESYSTEM_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) DoctypeSYSTEM() (String, error) {
	return getComProperty[String](d, "DoctypeSYSTEM")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCTYPESYSTEM_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetDoctypeSYSTEM(v String) error {
	return putComProperty(d, "DoctypeSYSTEM", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORCENOTEFORMAT_PROPERTY_DXLEXPORTER.html */
func (d NotesDXLExporter) ForceNoteFormat() (Boolean, error) {
	return getComProperty[Boolean](d, "ForceNoteFormat")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORCENOTEFORMAT_PROPERTY_DXLEXPORTER.html */
func (d NotesDXLExporter) SetForceNoteFormat(v Boolean) error {
	return putComProperty(d, "ForceNoteFormat", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MIMEOPTION_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) MIMEOption() (Long, error) {
	return getComProperty[Long](d, "MIMEOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MIMEOPTION_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetMIMEOption(v Long) error {
	return putComProperty(d, "MIMEOption", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OLEOBJECTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OLEObjectOmittedText() (String, error) {
	return getComProperty[String](d, "OLEObjectOmittedText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OLEOBJECTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOLEObjectOmittedText(v String) error {
	return putComProperty(d, "OLEObjectOmittedText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITITEMNAMES_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OmitItemNames() ([]String, error) {
	return getComArrayProperty[String](d, "OmitItemNames")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITITEMNAMES_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOmitItemNames(v []String) error {
	return putComProperty(d, "OmitItemNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITMISCFILEOBJECTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OmitMiscFileObjects() (Boolean, error) {
	return getComProperty[Boolean](d, "OmitMiscFileObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITMISCFILEOBJECTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOmitMiscFileObjects(v Boolean) error {
	return putComProperty(d, "OmitMiscFileObjects", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITOLEOBJECTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OmitOLEObjects() (Boolean, error) {
	return getComProperty[Boolean](d, "OmitOLEObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITOLEOBJECTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOmitOLEObjects(v Boolean) error {
	return putComProperty(d, "OmitOLEObjects", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTATTACHMENTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OmitRichtextAttachments() (Boolean, error) {
	return getComProperty[Boolean](d, "OmitRichtextAttachments")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTATTACHMENTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOmitRichtextAttachments(v Boolean) error {
	return putComProperty(d, "OmitRichtextAttachments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTPICTURES_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OmitRichtextPictures() (Boolean, error) {
	return getComProperty[Boolean](d, "OmitRichtextPictures")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTPICTURES_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOmitRichtextPictures(v Boolean) error {
	return putComProperty(d, "OmitRichtextPictures", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OUTPUTDOCTYPE_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) OutputDOCTYPE() (Boolean, error) {
	return getComProperty[Boolean](d, "OutputDOCTYPE")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OUTPUTDOCTYPE_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetOutputDOCTYPE(v Boolean) error {
	return putComProperty(d, "OutputDOCTYPE", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PICTUREOMITTEDTEXT_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) PictureOmittedText() (String, error) {
	return getComProperty[String](d, "PictureOmittedText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PICTUREOMITTEDTEXT_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetPictureOmittedText(v String) error {
	return putComProperty(d, "PictureOmittedText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESTRICTTOITEMNAMES_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) RestrictToItemNames() ([]String, error) {
	return getComArrayProperty[String](d, "RestrictToItemNames")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESTRICTTOITEMNAMES_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetRestrictToItemNames(v []String) error {
	return putComProperty(d, "RestrictToItemNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTOPTION_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) RichTextOption() (Long, error) {
	return getComProperty[Long](d, "RichTextOption")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTOPTION_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetRichTextOption(v Long) error {
	return putComProperty(d, "RichTextOption", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SCHEMALOCATION_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SchemaLocation() (String, error) {
	return getComProperty[String](d, "SchemaLocation")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SCHEMALOCATION_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetSchemaLocation(v String) error {
	return putComProperty(d, "SchemaLocation", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNCOMPRESSATTACHMENTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) UncompressAttachments() (Boolean, error) {
	return getComProperty[Boolean](d, "UncompressAttachments")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNCOMPRESSATTACHMENTS_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetUncompressAttachments(v Boolean) error {
	return putComProperty(d, "UncompressAttachments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALIDATIONSTYLE_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) ValidationStyle() (Integer, error) {
	return getComProperty[Integer](d, "ValidationStyle")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALIDATIONSTYLE_PROPERTY_EXPORTER.html */
func (d NotesDXLExporter) SetValidationStyle(v Integer) error {
	return putComProperty(d, "ValidationStyle", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPORT_METHOD_EXPORTER.html */
func (d NotesDXLExporter) Export(input any) (String, error) {
	err := d.checkImportTypes(input)

	if err != nil {
		return "", err
	}
	return callComMethod[String](d, "Export", input)
}
