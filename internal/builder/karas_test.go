package builder

import (
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"
)

const (
	seleniumPath     = "./drivers/selenium-server.jar"
	chromePath       = "./drivers/chrome-linux/chrome"
	chromeDriverPath = "./drivers/chromedriver"
	port             = 8080
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestBuildFromConfig(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name     string
		args     args
		wantArgs args
		wantErr  bool
	}{
		{name: "test case 1", args: args{filePath: "./misc/test/complex/normal/Karasfile"}, wantArgs: args{filePath: "./misc/test/complex/normal/karas.json"}, wantErr: false},
		{name: "config load error", args: args{filePath: "./misc/test/complex/exception/config_load_error/Karasfile"}, wantArgs: args{filePath: "./misc/test/complex/exception/config_load_error/karas.json"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildFromConfig(tt.args.filePath, seleniumPath, chromePath, chromeDriverPath, port)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want, _ := BuildFromKarasJSON(tt.wantArgs.filePath, seleniumPath, chromePath, chromeDriverPath, port)

			if !tt.wantErr {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("BuildFromConfig() got = %v, want %v", got, want)
				}
			}
		})
	}

}

func TestBuildFromKarasJSON(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name     string
		args     args
		wantArgs args
		wantErr  bool
	}{
		{name: "test case 1", args: args{filePath: "./misc/test/complex/normal/karas.json"}, wantArgs: args{filePath: "./misc/test/complex/normal/Karasfile"}, wantErr: false},
		{name: "json load error", args: args{filePath: "./misc/test/complex/exception/json_load_error/karas.json"}, wantArgs: args{filePath: "./misc/test/complex/exception/Karasfile"}, wantErr: true},
		{name: "json format error", args: args{filePath: "./misc/test/complex/exception/json_format_error/karas.json"}, wantArgs: args{filePath: "./misc/test/complex/exception/json_format_error/Karasfile"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildFromKarasJSON(tt.args.filePath, seleniumPath, chromePath, chromeDriverPath, port)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want, _ := BuildFromConfig(tt.wantArgs.filePath, seleniumPath, chromePath, chromeDriverPath, port)

			if !tt.wantErr {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("BuildFromConfig() got = %v, want %v", got, want)
				}
			}
		})

	}
}
