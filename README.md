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

<img width="661" height="267" alt="image" src="https://github.com/user-attachments/assets/b709e97a-087c-4e14-a83e-33aaea0c7419" />

