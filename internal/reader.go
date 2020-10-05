package reader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	viper "github.com/spf13/viper"
)

type SecretIDList map[string]string
type SecretsPayload map[string]interface{}

func GetSecrets(path string, filename string, group string, projectId string, version string) {

	sp := SecretsPayload{}
	secretIDList, _ := GetSecretIDList(path, filename, group)

	// Iterate serially to avoid exceeding secret manager rate limits
	for k, v := range secretIDList {
		res, _ := GetSecret(v, projectId, version)
		sp[strings.ToUpper(k)] = res
	}

	secretsPayload, err := json.Marshal(sp)
	if err != nil {
			fmt.Println(err.Error())
			return
	}

	jsonStr := string(secretsPayload)
	fmt.Println(jsonStr)
}

func GetSecret(name string, projectId string, version string) (string, error) {

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf("Error: failed to create secretmanager client: %v", err)
		return "", nil
	}
	//fmt.Println(name)

	if version == "" {
		version = "latest"
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, name, version),
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Printf("Error: failed to get secret: %v", err)
		return "", err
	}
	p := result.GetPayload()

	return string(p.Data), nil
}

func GetSecretIDList(path string, filename string, group string) (SecretIDList, error) {

	var response = SecretIDList{}

	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return response, errors.New("Failed to read input JSON file content.")
	}
	response = viper.GetStringMapString(group)
	//fmt.Println(response)

	return response, nil
}
