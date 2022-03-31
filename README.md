# Job Connection

## Run Locally

Clone the project

```bash
  git clone https://github.com/GDSC-UIT/job-connection-api.git
```

Go to the project directory

```bash
  cd job-connection-api
```

Config `.env` file [ENVIRONMENT.md](./ENVIRONMENT.md)

```bash
DATABASE_URL = YOUR_DATABASE_URL
JC_FIREBASE_SERVICE_ACCOUNT = YOUR_FIREBASE_SERVICE_ACCOUNT
```

Start the server

```bash
  go run .
```

## Demo

https://jobconnection.herokuapp.com/

## Tech Stack

Golang, Fiber, Firebase, PostgreSQL
