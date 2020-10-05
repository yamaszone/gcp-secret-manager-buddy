/*
// Populate secrets
echo -n "foo" | gcloud secrets create dev-foo --replication-policy=automatic --labels="env=dev" --data-file=-
echo -n "bar" | gcloud secrets create dev-bar --replication-policy=automatic --labels="env=dev" --data-file=-
echo -n "foo" | gcloud secrets versions add dev-foo --data-file=-
echo -n "foo" | gcloud secrets versions add dev-foo --data-file=-

gcloud secrets update dev-foo --update-labels=env=dev

// Service Account Setup
project_id=tntprod
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
	reader "github.com/yamaszone/secret-manager-reader/internal"
)

// TODO: Add CLI interface
func main() {
	reader.GetSecrets(".", "secrets", "dev", "tntprod", "latest")
}
