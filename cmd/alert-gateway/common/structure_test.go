package common

import (
	"net/http"
	"reflect"
	"testing"
)

func TestHttpPost(t *testing.T) {
	type args struct {
		url     string
		params  map[string]string
		headers map[string]string
		body    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "test",

		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HttpPost(tt.args.url, tt.args.params, tt.args.headers, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpPost() got = %v, want %v", got, tt.want)
			}
		})
	}
}