/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDXLEXPORTER_CLASS.html */
package notesdxlexporter_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var db domigo.NotesDatabase
var dxlexporter domigo.NotesDXLExporter

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, dbTemp domigo.NotesDatabase) (string, error) {
		var err error

		db = dbTemp
		dxlexporter, err = session.CreateDXLExporter()
		defer dxlexporter.Release()

		if err != nil {
			return "NotesDXLExporter could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHMENTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func TestAttachmentOmittedText(t *testing.T) {
	_, err := dxlexporter.AttachmentOmittedText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHMENTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func TestSetAttachmentOmittedText(t *testing.T) {
	s, err := dxlexporter.AttachmentOmittedText()
	require.Nil(t, err)

	err = dxlexporter.SetAttachmentOmittedText(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTNOTESBITMAPSTOGIF_PROPERTY_EXPORTER.html */
func TestConvertNotesBitmapsToGIF(t *testing.T) {
	_, err := dxlexporter.ConvertNotesBitmapsToGIF()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTNOTESBITMAPSTOGIF_PROPERTY_EXPORTER.html */
func TestSetConvertNotesBitmapsToGIF(t *testing.T) {
	s, err := dxlexporter.ConvertNotesBitmapsToGIF()
	require.Nil(t, err)

	err = dxlexporter.SetConvertNotesBitmapsToGIF(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCTYPESYSTEM_PROPERTY_EXPORTER.html */
func TestDoctypeSYSTEM(t *testing.T) {
	_, err := dxlexporter.DoctypeSYSTEM()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCTYPESYSTEM_PROPERTY_EXPORTER.html */
func TestSetDoctypeSYSTEM(t *testing.T) {
	s, err := dxlexporter.DoctypeSYSTEM()
	require.Nil(t, err)

	err = dxlexporter.SetDoctypeSYSTEM(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORCENOTEFORMAT_PROPERTY_DXLEXPORTER.html */
func TestForceNoteFormat(t *testing.T) {
	_, err := dxlexporter.ForceNoteFormat()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORCENOTEFORMAT_PROPERTY_DXLEXPORTER.html */
func TestSetForceNoteFormat(t *testing.T) {
	s, err := dxlexporter.ForceNoteFormat()
	require.Nil(t, err)

	err = dxlexporter.SetForceNoteFormat(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MIMEOPTION_PROPERTY_EXPORTER.html */
func TestMIMEOption(t *testing.T) {
	_, err := dxlexporter.MIMEOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MIMEOPTION_PROPERTY_EXPORTER.html */
func TestSetMIMEOption(t *testing.T) {
	s, err := dxlexporter.MIMEOption()
	require.Nil(t, err)

	err = dxlexporter.SetMIMEOption(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OLEOBJECTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func TestOLEObjectOmittedText(t *testing.T) {
	_, err := dxlexporter.OLEObjectOmittedText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OLEOBJECTOMITTEDTEXT_PROPERTY_EXPORTER.html */
func TestSetOLEObjectOmittedText(t *testing.T) {
	s, err := dxlexporter.OLEObjectOmittedText()
	require.Nil(t, err)

	err = dxlexporter.SetOLEObjectOmittedText(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITITEMNAMES_PROPERTY_EXPORTER.html */
func TestOmitItemNames(t *testing.T) {
	_, err := dxlexporter.OmitItemNames()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITITEMNAMES_PROPERTY_EXPORTER.html */
func TestSetOmitItemNames(t *testing.T) {
	s, err := dxlexporter.OmitItemNames()
	require.Nil(t, err)

	err = dxlexporter.SetOmitItemNames(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITMISCFILEOBJECTS_PROPERTY_EXPORTER.html */
func TestOmitMiscFileObjects(t *testing.T) {
	_, err := dxlexporter.OmitMiscFileObjects()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITMISCFILEOBJECTS_PROPERTY_EXPORTER.html */
func TestSetOmitMiscFileObjects(t *testing.T) {
	s, err := dxlexporter.OmitMiscFileObjects()
	require.Nil(t, err)

	err = dxlexporter.SetOmitMiscFileObjects(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITOLEOBJECTS_PROPERTY_EXPORTER.html */
func TestOmitOLEObjects(t *testing.T) {
	_, err := dxlexporter.OmitOLEObjects()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITOLEOBJECTS_PROPERTY_EXPORTER.html */
func TestSetOmitOLEObjects(t *testing.T) {
	s, err := dxlexporter.OmitOLEObjects()
	require.Nil(t, err)

	err = dxlexporter.SetOmitOLEObjects(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTATTACHMENTS_PROPERTY_EXPORTER.html */
func TestOmitRichtextAttachments(t *testing.T) {
	_, err := dxlexporter.OmitRichtextAttachments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTATTACHMENTS_PROPERTY_EXPORTER.html */
func TestSetOmitRichtextAttachments(t *testing.T) {
	s, err := dxlexporter.OmitRichtextAttachments()
	require.Nil(t, err)

	err = dxlexporter.SetOmitRichtextAttachments(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTPICTURES_PROPERTY_EXPORTER.html */
func TestOmitRichtextPictures(t *testing.T) {
	_, err := dxlexporter.OmitRichtextPictures()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OMITRICHTEXTPICTURES_PROPERTY_EXPORTER.html */
func TestSetOmitRichtextPictures(t *testing.T) {
	s, err := dxlexporter.OmitRichtextPictures()
	require.Nil(t, err)

	err = dxlexporter.SetOmitRichtextPictures(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OUTPUTDOCTYPE_PROPERTY_EXPORTER.html */
func TestOutputDOCTYPE(t *testing.T) {
	_, err := dxlexporter.OutputDOCTYPE()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OUTPUTDOCTYPE_PROPERTY_EXPORTER.html */
func TestSetOutputDOCTYPE(t *testing.T) {
	s, err := dxlexporter.OutputDOCTYPE()
	require.Nil(t, err)

	err = dxlexporter.SetOutputDOCTYPE(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PICTUREOMITTEDTEXT_PROPERTY_EXPORTER.html */
func TestPictureOmittedText(t *testing.T) {
	_, err := dxlexporter.PictureOmittedText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PICTUREOMITTEDTEXT_PROPERTY_EXPORTER.html */
func TestSetPictureOmittedText(t *testing.T) {
	s, err := dxlexporter.PictureOmittedText()
	require.Nil(t, err)

	err = dxlexporter.SetPictureOmittedText(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESTRICTTOITEMNAMES_PROPERTY_EXPORTER.html */
func TestRestrictToItemNames(t *testing.T) {
	_, err := dxlexporter.RestrictToItemNames()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESTRICTTOITEMNAMES_PROPERTY_EXPORTER.html */
func TestSetRestrictToItemNames(t *testing.T) {
	s, err := dxlexporter.RestrictToItemNames()
	require.Nil(t, err)

	err = dxlexporter.SetRestrictToItemNames(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTOPTION_PROPERTY_EXPORTER.html */
func TestRichTextOption(t *testing.T) {
	_, err := dxlexporter.RichTextOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTOPTION_PROPERTY_EXPORTER.html */
func TestSetRichTextOption(t *testing.T) {
	s, err := dxlexporter.RichTextOption()
	require.Nil(t, err)

	err = dxlexporter.SetRichTextOption(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SCHEMALOCATION_PROPERTY_EXPORTER.html */
func TestSchemaLocation(t *testing.T) {
	_, err := dxlexporter.SchemaLocation()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SCHEMALOCATION_PROPERTY_EXPORTER.html */
func TestSetSchemaLocation(t *testing.T) {
	s, err := dxlexporter.SchemaLocation()
	require.Nil(t, err)

	err = dxlexporter.SetSchemaLocation(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNCOMPRESSATTACHMENTS_PROPERTY_EXPORTER.html */
func TestUncompressAttachments(t *testing.T) {
	_, err := dxlexporter.UncompressAttachments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNCOMPRESSATTACHMENTS_PROPERTY_EXPORTER.html */
func TestSetUncompressAttachments(t *testing.T) {
	s, err := dxlexporter.UncompressAttachments()
	require.Nil(t, err)

	err = dxlexporter.SetUncompressAttachments(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALIDATIONSTYLE_PROPERTY_EXPORTER.html */
func TestValidationStyle(t *testing.T) {
	_, err := dxlexporter.ValidationStyle()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALIDATIONSTYLE_PROPERTY_EXPORTER.html */
func TestSetValidationStyle(t *testing.T) {
	s, err := dxlexporter.ValidationStyle()
	require.Nil(t, err)

	err = dxlexporter.SetValidationStyle(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPORT_METHOD_EXPORTER.html */
func TestExport(t *testing.T) {
	_, err := dxlexporter.Export(db)
	require.Nil(t, err)
}
