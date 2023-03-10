package builder

import (
	"reflect"
	"testing"
)

func 
TestBuildFromConfig(t *testing.T) {
	type args struct {
		configFilePath string
		selenium       string
		chrome         string
		chromeDriver   string
		port           int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Karas
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildFromConfig(tt.args.configFilePath, tt.args.selenium, tt.args.chrome, tt.args.chromeDriver, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildFromConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildFromKarasJSON(t *testing.T) {
	type args struct {
		jsonFilePath string
		selenium     string
		chrome       string
		chromeDriver string
		port         int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Karas
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildFromKarasJSON(tt.args.jsonFilePath, tt.args.selenium, tt.args.chrome, tt.args.chromeDriver, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildFromKarasJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildFromKarasJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
