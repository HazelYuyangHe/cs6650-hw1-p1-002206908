# Gin Albums REST API (Go + Gin + Google Cloud Run)

A simple RESTful API built with [Go](https://go.dev/) and the [Gin](https://github.com/gin-gonic/gin) framework.
The service provides endpoints to list, retrieve, add, and search albums.

---

## Project Info

- **Project ID**: `psychic-order-472603-i2`
- **Region**: `us-central1`
- **Public Service URL**:  
  **https://gin-albums-450812111044.us-central1.run.app**

---

## Local Development

### Prerequisites
- Go 1.21+ installed  
- `go mod tidy` to fetch dependencies

### Run locally
```bash
go run .
Default port is 8080 (set by $PORT env variable).

Test locally:

bash
Copy code
curl http://localhost:8080/albums
curl http://localhost:8080/albums/2
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"Betty Carter","artist":"Betty Carter","price":49.99}'
Deploy to Google Cloud Run
Enable APIs (Cloud Run, Cloud Build, Artifact Registry)

Deploy

bash
Copy code
gcloud run deploy gin-albums \
  --source . \
  --region us-central1 \
  --allow-unauthenticated
Cloud Run automatically:

Builds the container image with Buildpacks

Deploys and serves traffic on the given Service URL

API Endpoints
Method	Path	Description
GET	/albums	List all albums
GET	/albums/:id	Retrieve album by ID
POST	/albums	Add a new album (JSON body)
GET	/albums/search?title	Search albums by title substring

Quick Tests (Cloud Run)
Replace <URL> with the Service URL above:

bash
Copy code
URL=https://gin-albums-450812111044.us-central1.run.app

curl "$URL/albums"
curl "$URL/albums/2"
curl -X POST "$URL/albums" \
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"Betty Carter","artist":"Betty Carter","price":49.99}'
curl "$URL/albums/search?title=blue"
Screenshots
Place screenshots in docs/screenshots/ and reference them here.

Main.go in Cloud Shell Editor

Successful Cloud Run deployment output

cURL test output

Browser test of public URL

Repository Contents
bash
Copy code
.
├── cmd/server/main.go   # application entry point (if using cmd/)
├── internal/...         # handlers, models, router
├── go.mod
├── go.sum
├── test.http            # REST Client tests
└── docs/screenshots/    # deployment & test screenshots
