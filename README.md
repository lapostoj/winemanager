# Wine Manager

Author: lapostoj

Contact: jerome.lapostolet@gmail.com

## Description

Online application to manage the reserve of a wine cellar. Built for a personal use so some aspects might be specific.
The project is made with an upload on Google appengine in mind. This is why the "locals" imports are possible and done.
Without them there is a double import when running goapp.

## Technology

- Go (backend)

## Development

### Serve the app

```bash
dev_appserver.py cmd/app.yaml
```

The application expects the following in the `cmd` folder:

- an appengine config `app.yaml`.
- a static frontend part to live in an `app` folder to serve (with `index.html` and `notfound.html`)

### Run the tests

```bash
go test ./...
```

## Deploy

```bash
gcloud app deploy
```
