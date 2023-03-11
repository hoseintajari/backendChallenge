package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	gotConfig, err := LoadConfig("./testdata")
	if err != nil {
		t.Errorf("Error is : %s \n", err.Error())
	}
	expectedServer := Config{
		ServerIp:   "127.0.0.1",
		ServerPort: 8081,
	}
	if !reflect.DeepEqual(gotConfig, expectedServer) {
		t.Errorf("got:%v,expected:%v\n", gotConfig, expectedServer)
	}
}
