package appconfig

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	conf, err := GetConfig("config_test.yaml")
	if err != nil {
		t.Error(err)
	}
	require.NotNil(t, conf)
	require.Equal(t, conf.S3.Endpoint, "endpoint")
	require.Equal(t, conf.S3.AccessKeyID, "accessKeyID")
	require.Equal(t, conf.S3.SecretAccessKey, "secretAccessKey")
	require.Equal(t, conf.S3.UseSSL, false)
	require.Equal(t, conf.Database.Host, "host")
	require.Equal(t, conf.Database.Port, 1234)
	require.Equal(t, conf.Database.Username, "username")
	require.Equal(t, conf.Database.Password, "password")
	require.Equal(t, conf.Database.DatabaseName, "dbname")
}
