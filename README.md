# Armada

Kustomize packaging command line tool.

It allows you to git clone a packaged kustomize base and call it with the help of a config file.

## Usage

Create a armada.yaml file in each one of your kustomize infra folder.

Use armada to generate the app.yaml file.
```bash
armada armada.yaml
```

The generated file can be used directly to deploy to a kubernetes cluster using `kubectl -f app.yaml` or stored on a git repo or any other storage like (S3,GCS...) to be stored and versionned.

## Community

## ToDo
