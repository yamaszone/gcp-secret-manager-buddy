package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecretIDList(t *testing.T) {
	secretList, err := GetSecretIDList("../fixtures/secrets.json")
	assert.Nil(t, err)
	assert.NotNil(t, secretList)
}
