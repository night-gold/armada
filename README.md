# Armada

Kustomize packaging command line tool.

It allows you to git clone a packaged kustomize base and call it with the help of a config file.

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

The armada file configuration sould look like this.

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

### Code of conduct 

### Create a package

### ToDo

 - [ ] Write some test
 - [ ] Add overlays options (-o, default value and file option)
 - [ ] Add auto apply option (-a)
 - [ ] Make a better documentation
 - [ ] Allow default overlays generation with overlays value in file
 - [ ] Allow package multiple deployment inside a single namespace (naming of each resources)