package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestLanguagesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		pages *PageParams
	}
	tests := []struct {
		name    string
		args    args
		want    *PageSummary
		want1   []*Language
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   12,
				PerPage: 2,
				Total:   23,
			},
			want1: []*Language{
				{
					Code:       "ar",
					Name:       "Arabic",
					NativeForm: "عربي",
				},
				{
					Code:       "bg",
					Name:       "Bulgarian ",
					NativeForm: "Български",
				},
			},
			wantErr: false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lang := &LanguagesService{
				client: client,
			}

			got, got1, err := lang.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("LanguagesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LanguagesService.List() got = %v, want %v", got, tt.want)
			}
			for i := range got1 {
				if !reflect.DeepEqual(got1[i], tt.want1[i]) {
					t.Errorf("LanguagesService.List() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
				}
			}
		})
	}
}

func TestLanguagesService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		languageCode string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Language
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				languageCode: testutils.TestDefaultLanguageCode,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Language{
				Code:       testutils.TestDefaultLanguageCode,
				Name:       "Japanese ",
				NativeForm: "日本語",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because languageCode is invalid",
			args: args{
				languageCode: testutils.TestInvalidLanguageCode,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/languages/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidLanguageCode,
				),
				Code: 200,
				Message: []ErrorMessage{
					{
						ID:    "150",
						Key:   "Language is not yet supported in the API",
						Value: "Response requested in an unsupported language.",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LanguagesService{
				client: client,
			}
			got, got1, err := c.Get(tt.args.languageCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("LanguagesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("LanguagesService.Get() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LanguagesService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LanguagesService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
