# Packages

A package contains the minimal configuration to deploy an apps without specific configuration.

## Content

A package should be organised like this:

> ```
> /base
> /docs
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

## Should not be includer inside a package base folder

The minimal configuration thinking put some restriction on the resources that can be included inside the base folder, some of them should not exists within:

* namespace (hard restriction): You can't force people to use the same namespace as you, no valid base should contains a namespace (excluding private packages)
* ingress (soft restriction): The number of ingress available is a lot, if you create an ingress file, do not put specific annotations inside it. If you are not sure of the neutrality of your ingress, put overlays examples with full ingress configuration for one or more ingress.
* volumes: There is as much PVC as the number of cloud providers out there (maybe more) having a full working configuration may be hard with this. Having version specifics packages may be a good idea.