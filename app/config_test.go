package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	pUser string = "user"
	pPass string = "secret"
	dName string = "test"
	dHost string = "1.2.3.4"
	dSslM string = "none"
)

func TestNewConfig(t *testing.T) {
	t.Run("Configuration success", func(t *testing.T) {
		t.Setenv("POSTGRES_USER", pUser)
		t.Setenv("POSTGRES_PASSWORD", pPass)
		t.Setenv("DATABASE_NAME", dName)
		t.Setenv("DATABASE_HOST", dHost)
		t.Setenv("DATABASE_SSL_MODE", dSslM)

		c := NewConfig()

		assert.Equal(t, pUser, c.DbUser)
		assert.Equal(t, pPass, c.DbPass)
		assert.Equal(t, dName, c.DbName)
		assert.Equal(t, dHost, c.DbHost)
		assert.Equal(t, dSslM, c.DbSSLM)
	})
}
