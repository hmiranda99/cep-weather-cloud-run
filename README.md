# CEP Weather - Cloud Run

## Endpoint
GET /weather?cep=01001000

## Run local
docker compose up --build

## Tests
go test ./...

## Deploy Cloud Run
gcloud builds submit --tag gcr.io/SEU_PROJETO/cep-weather
gcloud run deploy cep-weather --image gcr.io/SEU_PROJETO/cep-weather --platform managed --region us-central1 --allow-unauthenticated


## URL
Para criar a conta gratuita é necessário fazer um paragamento de 200 reais, no qual não disponho.

![alt text](image.png)
