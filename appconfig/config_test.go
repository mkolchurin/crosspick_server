package appconfig

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	conf, err := GetConfig()
	if err != nil {
		t.Error(err)
	}
	if conf == nil {
		t.Errorf("Conf is nil")
	}
	t.Log(conf.S3)
	t.Log(conf.Database)
}
