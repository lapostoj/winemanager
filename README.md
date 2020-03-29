# Wine Manager

Author: lapostoj

Contact: jerome.lapostolet@gmail.com

## Description

Online application to manage the reserve of a wine cellar.
Built for a personal use so some aspects might be specific.
The project is made with an upload on Google appengine in mind.

## Technology

- Go 1.13 (with modules)
=======

## Development

### Serve the app

```bash
FRONTEND_FOLDER=./app/frontend/ go run app/main.go
```

### Manage dependencies

See upgrades to be done

```bash
go list -u -m all
```

Update all direct and indirect dependencies

```bash
go get -u ./...
```

Prune dependencies from `go.mod`

```bash
go mod tidy
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
