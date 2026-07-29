# JWT Secret Setup

`utils/jwt.go` reads the signing secret from the `JWT_SECRET` environment
variable — it is never hardcoded in source. The app will refuse to start
(`log.Fatal`) if `JWT_SECRET` is missing or empty.

## Local development

1. Copy the example file:
   ```bash
   cp .env.example .env
   ```
2. `.env` contains a **placeholder** secret so you can run the project
   immediately. Replace it with your own before doing anything real:
   ```bash
   openssl rand -base64 32
   ```
   Paste the output as the value of `JWT_SECRET` in `.env`.
3. Go does not read `.env` files by itself. Either export the variable into
   your shell before running the app:
   ```bash
   export $(grep -v '^#' .env | xargs)
   go run main.go
   ```
   or add a small loader (e.g. `github.com/joho/godotenv`) to `main.go` if
   you want `.env` picked up automatically — ask if you want this wired up.

`.env` is listed in `.gitignore` and must never be committed. `.env.example`
has no real secret value once you've swapped it — it's just a template for
teammates.

## Production

Don't use a `.env` file in production. Set `JWT_SECRET` through whatever
your deployment platform provides for secrets (Docker/Kubernetes secret,
systemd `EnvironmentFile`, cloud provider secrets manager, CI/CD secret
store, etc.), so it's encrypted at rest and access-controlled separately
from the code.

## Rotating the secret

Changing `JWT_SECRET` invalidates every JWT issued with the old value —
all logged-in users will need to log in again. Rotate it if you suspect the
secret has leaked (e.g. it was ever committed to git).
