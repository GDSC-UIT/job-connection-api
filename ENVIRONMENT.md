## JC_FIREBASE_SERVICE_ACCOUNT

[Docs](https://firebase.google.com/docs/admin/setup?authuser=0#go)

Create firebase project

Add provider `Email/Password` and `Google`
![provider](https://i.imgur.com/lxCAzKX.png)
Set up `Storage` in `test mod`
![Imgur](https://i.imgur.com/y69th6l.png)
![Imgur](https://i.imgur.com/qIXHB9w.png)
Generate new private key
![Imgur](https://i.imgur.com/LzG14RX.png)
Remove `\n` character in private key
![Imgur](https://i.imgur.com/wBbL3V5.png)

Update `.env` variable

```bash
JC_FIREBASE_SERVICE_ACCOUNT = {"type":"service_account","project_id":"job-connection-b4340","private_key_id":"______","private_key":"-----BEGIN PRIVATE KEY-----.......}
```

## DATABASE_URL

[install Docker](https://docs.docker.com/get-started/)

Run database container

```bash
docker compose up -d
```

Update `.env` variable

```bash
DATABASE_URL = postgres://postgres:postgres@localhost:5432/postgres
```
