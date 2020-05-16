# Wine Manager

Author: lapostoj

Contact: jerome.lapostolet@gmail.com

## Description

Online application to manage the reserve of a wine cellar.
Built for a personal use so some aspects might be specific.
The project is made with an upload on Google appengine in mind.

## Technology

- Go 1.14 (with modules)
=======

## Development

### Build the app

```bash
docker build -t winemanager .
```

### Serve the app

```bash
docker run -p 8080:8080 --env FRONTEND_FOLDER="/frontend/" winemanager
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
