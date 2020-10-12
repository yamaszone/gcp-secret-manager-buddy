package reader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecretIDListWithValidData(t *testing.T) {
	secretList, err := GetSecretIDList("../fixtures/secrets.json")
	assert.Nil(t, err)
	assert.NotNil(t, secretList)
}

func TestGetSecretIDListWithInvalidData(t *testing.T) {
	secretList, err := GetSecretIDList("../fixtures/invalid.json")
	assert.NotNil(t, err)
	assert.Nil(t, secretList)
}

func TestGetSecretsWithMockedData(t *testing.T) {
	err := os.Setenv("GSM_IS_STUB", "yes")
	err = GetSecrets("../fixtures/secrets.json", "dummy-project", "latest")
	assert.Nil(t, err)
	_ = os.Setenv("GSM_IS_STUB", "no")
}

func TestGetSecretsWithMockedInvalidData(t *testing.T) {
	err := os.Setenv("GSM_IS_STUB", "yes")
	err = GetSecrets("../fixtures/invalid.json", "dummy-project", "latest")
	assert.NotNil(t, err)
	_ = os.Setenv("GSM_IS_STUB", "no")
}
