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
