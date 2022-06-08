package domain

// Model contains the meta data of the HTML document, instance of the model and bindings
// A <model> can have multiple instances as childnodes.
// The first and required <instance> is called the primary instance and represents the data structure of
// the record that will be created and submitted with the form.
// Additional instances are called secondary instances.
type Model struct {
	InstanceData InstanceData `xml:"instance"`
	BindData     BindData     `xml:"bind"`
}
