# Development environment constructions

1. Install Docker Desktop, npm, go, golang-migrate

2. Install dependencies

```bash
  go mod tidy

  cd web
  npm i
```

3. Create .env file and input env variables

```bash
  cp .env.template .env
  cd web
  cp .env.template .env
```

4. Run

Start database
```bash
  docker compose up
```

Run api
```bash
  make run/api
```

Run frontend
```bash
  cd web
  npm run start
```

Run database migrations
```bash
  make db/migrations/up
```
