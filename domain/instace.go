package domain

import "encoding/xml"

// InstanceData contains a list of instances
type InstanceData struct {
	Instances []Instance `xml:"instance"`
}

// UnmarshalXML intercepts the xml unmarshaling to add custom unmarshaling behavior
func (i *InstanceData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var bd Instance
	if err := d.DecodeElement(&bd, &start); err != nil {
		return err
	}
	i.Instances = append(i.Instances, bd)
	return nil
}

// Instance can be a primary or secondary instance
// The primary instance (Required) is the data structure of the record that will
// be created and submitted with the form.
// The primary instance is the first instance defined by the form and should contain a single childnode
// The primary instance’s single child is the document root that XPath expressions are evaluated on
// The secondary instances (Optional) are used to pre-load read-only data inside a form.
// This data is searchable in XPath. At the moment the key use case is in designing
// so-called cascading selections where the available options of a multiple-choice question can be
// filtered based on an earlier answer.
// A secondary instance should get a unique id attribute on the <instance> node.
// This allows apps to query the data (which is outside the root, ie. the primary instance,
// and would normally not be reachable).
// It uses the instance('instanceID')/root/item[childItemElement='childItemElementValue'] syntax to do this.
type Instance struct {
	ID string `xml:"id,attr"`
	// This is a bold assumption that the primary instance
	// will have a child element called `data`
	Data Data `xml:"data"`
	// all instances that have the child tag `root` are secondary instances
	Root Root `xml:"root"`
}

// Data contains the child of the primary instance
// This will be shared with the response forms
// The response form shares a similar data structure with the primary instance child element
type Data struct {
	ID      string   `xml:"id,attr"`
	XMLName xml.Name `xml:"data"`
	// DataItems map[string]interface{} `xml:",any,attr"`
}

// Root contains the child instance of the secondary elements
type Root struct {
	ID      string   `xml:"id,attr"`
	XMLName xml.Name `xml:"root"`
	// RootItems map[string]interface{} `xml:",any,attr"`
}
