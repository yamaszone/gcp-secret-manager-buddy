package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/kelseyhightower/envconfig"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type Config struct {
	GsmIsStub string `envconfig:"GSM_IS_STUB" default:"no"`
	Timeout   string `envconfig:"GSM_TIMEOUT" default:"10"`
}

type GsmResponse struct {
	secretKey   string
	secretValue string
	err         error
}

type SecretIDList map[string]string

type SecretsPayload map[string]interface{}

var (
	EnvConfig = InitConfig()
)

func InitConfig() Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error: failed to load config: %v", err)
	}
	return cfg
}

func printPayload(sp SecretsPayload) error {
	secretsPayload, err := json.Marshal(sp)
	if err != nil {
		return err
	}

	jsonStr := string(secretsPayload)
	fmt.Println(jsonStr)
	return nil
}

func GetSecrets(filename string, projectId string, version string) error {

	sp := SecretsPayload{}

	secretIDList, err := GetSecretIDList(filename)
	if err != nil {
		return err
	}

	if EnvConfig.GsmIsStub == "yes" {
		for k, v := range secretIDList {
			sp[k] = v
		}
		err := printPayload(sp)
		if err != nil {
			return err
		}
		return nil
	}

	ch := make(chan *GsmResponse, len(secretIDList))

	for k, v := range secretIDList {
		go func(k string, v string) {
			res, err := GetSecret(v, projectId, version)
			ch <- &GsmResponse{k, res, err}
		}(k, v)
	}

	for {
		select {
		case r := <-ch:
			//fmt.Printf("Secret Key = %s was fetched, Secret Value = %s\n", r.secretKey, r.secretValue)
			sp[r.secretKey] = r.secretValue
			if len(sp) == len(secretIDList) {
				err := printPayload(sp)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	return nil
}

func GetSecret(name string, projectId string, version string) (string, error) {

	timeout, _ := strconv.Atoi(EnvConfig.Timeout)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf("Error: failed to create secretmanager client: %v", err)
		return "", err
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

func GetSecretIDList(filename string) (SecretIDList, error) {
	var response SecretIDList

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Printf("Error: failed to open %s. %s", filename, err)
		return response, err
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Error: failed to read %s. %s", filename, err)
		return response, err
	}

	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		log.Printf("Error: failed to parse JSON file %s. %s", filename, err)
		return response, err
	}
	//fmt.Println(response)

	return response, nil
}
