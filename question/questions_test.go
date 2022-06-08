package question

import (
	"reflect"
	"testing"

	"github.com/maxwellgithinji/odktojson/domain"
	testutils "github.com/maxwellgithinji/odktojson/testUtils"
)

func TestDecodeBody(t *testing.T) {
	question1, err := testutils.ReadFileContents("../testUtils/question/testQuestion1.xml")
	if err != nil {
		t.Error(err)
	}

	wantBody1, err := testutils.DecodeBodyToJson("../testUtils/question/body1.json")
	if err != nil {
		t.Error(err)
	}

	question2, err := testutils.ReadFileContents("../testUtils/question/testQuestion2.xml")
	if err != nil {
		t.Error(err)
	}

	wantBody2, err := testutils.DecodeBodyToJson("../testUtils/question/body2.json")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		body string
	}

	tests := []struct {
		name    string
		args    args
		want    domain.Body
		wantErr bool
	}{
		{
			name: "A form with one instance, ungrouped body",
			args: args{
				body: question1,
			},
			want:    wantBody1,
			wantErr: false,
		},
		{
			name: "A form with more than one instance, grouped body",
			args: args{
				body: question2,
			},
			want:    wantBody2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBody(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// TODO: Remove, this outputs the decoded xml to a file
			err = testutils.DecodeTestXML(got)
			if err != nil {
				t.Error(err, ">>>>>error decoding test xml")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeInstances(t *testing.T) {
	// question1, err := testutils.ReadFileContents("../testUtils/question/testQuestion1.xml")
	// if err != nil {
	// 	t.Error(err)
	// }

	// wantInstance1, err := testutils.DecodeInstancesToJSON("../testUtils/question/instance1.json")
	// if err != nil {
	// 	t.Error(err)
	// }

	question2, err := testutils.ReadFileContents("../testUtils/question/testQuestion2.xml")
	if err != nil {
		t.Error(err)
	}

	wantInstance2, err := testutils.DecodeInstancesToJSON("../testUtils/question/instance2.json")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		instance string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.InstanceData
		wantErr bool
	}{
		// {
		// 	name: "A form with one instance, ungrouped body",
		// 	args: args{
		// 		instance: question1,
		// 	},
		// 	want:    wantInstance1,
		// 	wantErr: false,
		// },
		{
			name: "A form with more than one instance, grouped body",
			args: args{
				instance: question2,
			},
			want:    wantInstance2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeInstances(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// TODO: Remove, this outputs the decoded xml to a file
			err = testutils.DecodeTestXML(got)
			if err != nil {
				t.Error(err, ">>>>>error decoding test xml")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBinding(t *testing.T) {
	question1, err := testutils.ReadFileContents("../testUtils/question/testQuestion1.xml")
	if err != nil {
		t.Error(err)
	}

	wantModel1, err := testutils.DecodeBindingToJSON("../testUtils/question/binding1.json")
	if err != nil {
		t.Error(err)
	}

	question2, err := testutils.ReadFileContents("../testUtils/question/testQuestion2.xml")
	if err != nil {
		t.Error(err)
	}

	wantModel2, err := testutils.DecodeBindingToJSON("../testUtils/question/binding2.json")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		model string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.BindData
		wantErr bool
	}{
		{
			name: "A form with one instance, ungrouped body",
			args: args{
				model: question1,
			},
			want:    wantModel1,
			wantErr: false,
		},
		{
			name: "A form with multiple instances, grouped body",
			args: args{
				model: question2,
			},
			want:    wantModel2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBindings(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBindings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// TODO: Remove, this outputs the decoded xml to a file
			err = testutils.DecodeTestXML(got)
			if err != nil {
				t.Error(err, ">>>>>error decoding test xml")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBindings() = %v, want %v", got, tt.want)
			}
		})
	}
}
