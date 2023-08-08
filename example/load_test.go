package example

import (
	configUtil "github.com/spf13/viper"
	confUtil "github.com/wegoteam/wepkg/config"
	"reflect"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	type fields struct {
		Config   *configUtil.Viper
		Name     string
		Type     string
		Path     []string
		Profiles string
		IsLoad   bool
	}
	type args struct {
		prefix string
		data   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				Config:   configUtil.New(),
				Name:     "config",
				Type:     "yaml",
				Path:     []string{"./config"},
				Profiles: "dev",
				IsLoad:   false,
			},
			args: args{
				prefix: "mysql",
				data:   &confUtil.MySQL{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &confUtil.Config{
				Config:   tt.fields.Config,
				Name:     tt.fields.Name,
				Type:     tt.fields.Type,
				Path:     tt.fields.Path,
				Profiles: tt.fields.Profiles,
				IsLoad:   tt.fields.IsLoad,
			}
			if err := config.Load(tt.args.prefix, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_Parse(t *testing.T) {
	type fields struct {
		Config   *configUtil.Viper
		Name     string
		Type     string
		Path     []string
		Profiles string
		IsLoad   bool
	}
	type args struct {
		prefix string
		data   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &confUtil.Config{
				Config:   tt.fields.Config,
				Name:     tt.fields.Name,
				Type:     tt.fields.Type,
				Path:     tt.fields.Path,
				Profiles: tt.fields.Profiles,
				IsLoad:   tt.fields.IsLoad,
			}
			if err := config.Parse(tt.args.prefix, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name string
		want *confUtil.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confUtil.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	type args struct {
		configName string
		configType string
		profiles   string
		confPaths  []string
	}
	tests := []struct {
		name string
		args args
		want *confUtil.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confUtil.NewConfig(tt.args.configName, tt.args.configType, tt.args.profiles, tt.args.confPaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead(t *testing.T) {
	type args struct {
		config *confUtil.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := confUtil.Read(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetConfig(t *testing.T) {
	type args struct {
		configName string
		configType string
		profiles   string
		confPaths  []string
	}
	tests := []struct {
		name string
		args args
		want *confUtil.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confUtil.SetConfig(tt.args.configName, tt.args.configType, tt.args.profiles, tt.args.confPaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
