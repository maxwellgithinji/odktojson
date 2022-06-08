package domain

import "encoding/xml"

// Body contains the information required to display a form.
// The <body> contains the information required to display a question to a user,
// including the type of prompt, the appearance of the prompt (widget), the labels,
// the hints and the choice options.
type Body struct {
	// Body properties
	XMLName       xml.Name   `xml:"body"`
	XMLAttributes []xml.Attr `xml:"body,attr"`

	// Body children nodes
	// Group contains the questions that will be displayed to the user in groups
	Group []Group `xml:"group"`

	// Select1 contains the questions that will be displayed to the user in a single select
	Select1 []Select1 `xml:"select1"`
}

/*
Group combines elements together. If it has a child <label> element, the group is considered a
presentation group and will be displayed as a visually distinct group.
A <group> may or may not contain a ref attribute. If it does, the group
is considered a logical group. A logical group has a corresponding element in the primary
instance and usually a corresponding <bind> element. A logical group’s ref is
used as the context node for the relative ref paths of its descendants.
A group can be both a logical and a presentation group.
Groups may be nested to provide different levels of structure.
Apart from providing structure, a logA <group> combines elements together.
If it has a child <label> element, the group is considered a presentation group and will be
displayed as a visually distinct group.
A <group> may or may not contain a ref attribute. If it does, the group is
considered a logical group. A logical group has a corresponding element in the primary
instance and usually a corresponding <bind> element. A logical group’s ref is used as the context
node for the relative ref paths of its descendants.
A group can be both a logical and a presentation group.
Groups may be nested to provide different levels of structure.
Apart from providing structure, a logical group can also contain a relevant
attribute on its <bind> element, offering a powerful way to keep form logic maintainable (see bind attributes).
The sample below includes both the body and corresponding instance. The respondent
group is a logical group and the context group is both a logical and a presentation group.
The context group will only be shown if both first name and last name are filled
in.ical group can also contain a relevant attribute on its <bind> element, offering a powerful
way to keep form logic maintainable (see bind attributes).
The sample below includes both the body and corresponding instance. The respondent group is a
logical group and the context group is both a logical and a presentation group.
The context group will only be shown if both first name and last name are filled in.
*/
type Group struct {
	Input []Input `xml:"input"`

	// Include all question types
	Select1 []Select1 `xml:"select1"`
}

//  Input contains the information required to display a question to a user,
type Input struct {
	Label Label `xml:"label"`
}

// Label contains the information required to display a question to a user,
type Label struct {
	Text string `xml:",chardata"`
}

// Select1 contains questions of type option selection
type Select1 struct {
	Ref          string  `xml:"ref,attr"`
	Label        string  `xml:"label"`
	Hint         string  `xml:"hint"`
	Trigger      Trigger `xml:"trigger"`
	Items        []Item  `xml:"item"`
	QuestionType string
}

// Trigger contains the information required to display a question to a user,
type Trigger struct {
	Label string `xml:"label"`
}

// Item contains the selections options
type Item struct {
	Label string `xml:"label"`
	Value string `xml:"value"`
}
