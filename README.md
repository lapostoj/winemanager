# Wine Manager

Author: lapostoj

Contact: jerome.lapostolet@gmail.com

## Description

Online application to manage the reserve of a wine cellar.
Built for a personal use so some aspects might be specific.
The project is made with an upload on Google appengine in mind.

## Technology

- Go 1.12 (with modules)
=======

## Development

### Serve the app

```bash
FRONTEND_FOLDER=./app/frontend/ go run app/main.go
```

### Run the tests

```bash
go test ./...
```

## Deploy

The application expects the following in the `app` folder:

- a static frontend part to live in an `frontend` folder to serve (with `index.html` and `notfound.html`)

```bash
gcloud app deploy
```
