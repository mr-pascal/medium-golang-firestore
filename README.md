# medium-golang-firestore

Small example project to use GCP Firestore with the Go programming language.

## Install & start the GCP Firestore emulator



```sh
## Starts the Firestore emulator on local adress 'localhost:8090'
gcloud beta emulators firestore start --host-port=localhost:8090
```

Checkout [GCP Firestore Emulator](https://cloud.google.com/sdk/gcloud/reference/beta/emulators/firestore) for documentation.


## Start Development

Before starting the below code for local development make sure that the GCP Firestore emulator is up & running in a separate terminal window.

```sh
## Navigating to source folder
cd src

## Starting go application with env variable to local Firestore emulator
export FIRESTORE_EMULATOR_HOST=localhost:8090 && go run .

```