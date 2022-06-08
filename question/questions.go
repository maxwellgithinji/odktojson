package question

import (
	"encoding/xml"

	"github.com/maxwellgithinji/odktojson/domain"
)

// DecodeBody decodes the body of the XML file
// The body contains the information required to display a form
func DecodeBody(body string) (domain.Body, error) {
	var bodyObject domain.Body

	question := domain.SurveyQuestionnaire{
		Body: bodyObject,
	}

	err := xml.Unmarshal([]byte(body), &question)
	if err != nil {
		return domain.Body{}, err
	}
	return question.Body, err
}

// DecodeModel decodes the instances and bindings of the XML document
func DecodeBindings(model string) (domain.BindData, error) {
	var modelObject domain.BindData

	question := domain.SurveyQuestionnaire{
		Head: domain.Head{
			Model: domain.Model{
				BindData: modelObject,
			},
		},
	}

	err := xml.Unmarshal([]byte(model), &question)
	if err != nil {
		return modelObject, err
	}

	return question.Head.Model.BindData, nil

}

// DecodeInstance decodes the instance of the XML file
// The instance contains the information required to display a form
func DecodeInstances(instance string) (domain.InstanceData, error) {
	var instanceObject domain.InstanceData

	question := domain.SurveyQuestionnaire{
		Head: domain.Head{
			Model: domain.Model{
				InstanceData: instanceObject,
			},
		},
	}

	err := xml.Unmarshal([]byte(instance), &question)
	if err != nil {
		return instanceObject, err
	}
	return question.Head.Model.InstanceData, err
}
