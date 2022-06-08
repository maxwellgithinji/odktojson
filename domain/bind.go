package domain

import "encoding/xml"

// BindData contains a list of bind data
type BindData struct {
	Bindings []Bind `xml:"bind"`
}

// UnmarshalXML intercepts the xml unmarshaling to add custom unmarshaling behavior
func (b *BindData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var bd Bind
	if err := d.DecodeElement(&bd, &start); err != nil {
		return err
	}
	b.Bindings = append(b.Bindings, bd)
	return nil
}

// Bind contains the questions bindings
// A <bind> element wires together a primary instance node and the presentation of
// the corresponding question to the user.
// It is used to describe the datatype and various kinds of logic related to the data.
// A bind can refer to any node in the primary instance including repeated nodes.
// It may or may not have a corresponding presentation node in the body.
// An instance node does not require a corresponding <bind> node, regardless of whether
// it has a presentation node.
type Bind struct {
	Preload       string `xml:"preload,attr"`
	PreloadParams string `xml:"preloadParams,attr"`
	NodeSet       string `xml:"nodeset,attr"`
	Type          string `xml:"type,attr"`
	Required      string `xml:"required,attr"`
}
