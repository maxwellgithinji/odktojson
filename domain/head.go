package domain

// Head contains the title of the HTML document and also the meta model
// This will mainly include the `model` of the ODK XML
type Head struct {
	Model Model `xml:"model"`
}
