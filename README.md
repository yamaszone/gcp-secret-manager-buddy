# GCP Secret Manager Buddy (gsm-buddy)
`gsm-buddy` can be used to fetch secrets from [GCP Secret Manager](https://cloud.google.com/secret-manager/docs) as a group which is not currently supported by GCP Secret Manager.

## Use Case
- Fetch secrets for an app prior to it's deployment
- Run `gsm-buddy` as a sidecar of an application to feed secrets periodically

## Input
`cat input.json`
```
{
	"KEY1":"secret-name1",
	"KEY2":"secret-name2"
}
```

## Execute
`gsm-buddy get -i input.json -p my-gcp-project`

## Output
```
{
	"KEY1":"secret-value1",
	"KEY2":"secret-value2"
}
```

## Installation
### Linux/macOS
```
curl -sSL https://github.com/yamaszone/gcp-secret-manager-buddy/releases/download/v0.1.0/gcp-secret-manager-buddy-v0.1.0-$(
    bash -c '[[ $OSTYPE == darwin* ]] && echo darwin || echo linux'
  )-amd64 -o gsm-buddy && chmod a+x gsm-buddy && sudo mv gsm-buddy /usr/local/bin/
```
### Windows
Download executable from [releases page](https://github.com/yamaszone/gcp-secret-manager-buddy/releases/tag/v0.1.0)

## GCP Service Account Setup
```
project_id=my-gcp-project-id
sa_name=secrets-manager-reader-foo
iam_account="${sa_name}@${project_id}.iam.gserviceaccount.com"
gcloud iam service-accounts create "$sa_name" --display-name "$sa_name"
gcloud projects add-iam-policy-binding "$project_id" --member "serviceAccount:${iam_account}" --role "roles/secretmanager.viewer"
gcloud projects add-iam-policy-binding "$project_id" --member "serviceAccount:${iam_account}" --role "roles/secretmanager.secretAccessor"
gcloud iam service-accounts keys create --iam-account "$iam_account" ~/${sa_name}-key.json
export GOOGLE_APPLICATION_CREDENTIALS=~/${sa_name}-key.json
```
