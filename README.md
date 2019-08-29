# Armada

Kustomize packaging command line tool.

It allows you to git clone a packaged kustomize base and call it with the help of a config file.

## Pre-requisite

For armada to works, you need to have a recent version of kubectl with kustomize integration as armada will make calls to the kubectl command line to build the yaml file using `kubectl -k` option.

## Usage

### 1) Create a kustomize package file struct

Create an armada.yaml file in each one of your kustomize infra folder where you have your overlays.

> ```
> /app
> ├── redis
> │   ├── armada.yaml
> │   └── overlays
> │       └── apply
> │           └── kustomization.yaml
> └── Microservice
>     ├── armada.yaml
>     └── overlays
>         └── apply
>             └── kustomization.yaml
> ```

### 2) Configuration of the armada file

The armada file configuration should look like this.

```
--- 
repo:
  - repository: armada-prom-op
    version: master
    user: night-gold
    folder: "."
    Git: https://github.com
    Overlays: apply
```

Use armada to generate the app.yaml file.
```bash
armada -f armada.yaml
```

The generated file can be used directly to deploy to a kubernetes cluster using `kubectl -f app.yaml` or stored on a git repo or any other storage like (S3,GCS...) to be stored and versionned.

## Community

Before filling a bug please take a look here: [bugs](docs/bugs.md)
Before making a merge request please take a look here: [merge request](docs/mr.md) 

### Code of conduct 

Please be polite and respectful when communicating, filling a bug, commenting or when updating the repo. In case of non respect of this simple rule the comment will be deleted.  
### Create a package

See specific documentation for [packages](docs/packages.md)

### ToDo

 - [ ] Write some test
 - [x] ~~Add overlays options (-o, default value and file option)~~
 - [ ] Add auto apply option (-a)
 - [ ] Make a better documentation
 - [x] ~~Add git authentication options~~
 - [ ] Allow default overlays generation with overlays value in file
 - [ ] Allow package multiple deployment inside a single namespace (naming of each resources)
 - [ ] Add installation option (brew...)
 - [ ] Add a cache option for package deployed multiple time
 - [ ] Add a sleep time if CRD creation for operator is a bit long (prometheus-operator) before creating new resources.
 - [ ] Allow non package restrictive repository to not store their base folder at the root of their repository.