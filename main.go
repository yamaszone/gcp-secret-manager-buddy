/*
// Populate secrets
gcloud secrets create dev-foo --replication-policy=automatic --data-file="data.txt"
gcloud secrets versions add dev-foo --data-file="data.txt"

// Service Account Setup
project_id=prod
sa_name=secrets-manager-reader-blue
iam_account="${sa_name}@${project_id}.iam.gserviceaccount.com"
gcloud iam service-accounts create "$sa_name" --display-name "$sa_name"
gcloud projects add-iam-policy-binding "$project_id" --member "serviceAccount:${iam_account}" --role "roles/secretmanager.viewer"
gcloud iam service-accounts keys create --iam-account "$iam_account" ~/${sa_name}-key.json
export GOOGLE_APPLICATION_CREDENTIALS=~/${sa_name}-key.json

gcloud projects add-iam-policy-binding "$project_id" --member "serviceAccount:${iam_account}" --role "roles/secretmanager.secretAccessor"

*/

package main

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {
	res, _ := getSecret("dev-foo", "prod", "latest")
	fmt.Println(res)
}

func getSecret(name string, projectId string, version string) (*secretmanagerpb.SecretPayload, error) {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf("failed to create secretmanager client: %v", err)
		return nil, nil
	}
	//fmt.Println(name)
	if version == "" {
		version = "latest"
	}
	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, name, version),
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}

	return result.Payload, nil
}
