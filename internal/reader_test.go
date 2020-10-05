package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataDir = "../fixtures/"

func TestGetSecretIDList(t *testing.T) {
	secretList, err := GetSecretIDList(dataDir, "secrets", "dev")
	assert.Nil(t, err)
	assert.NotNil(t, secretList)
}
