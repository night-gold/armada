# Packages

A package contains the minimal configuration to deploy an apps without specific configuration.

## Versioning

For the packages specific repositories please try to follow the `vMAJOR.MINOR.PATCH` (vM.M.P). 

If you have an app that may have multiple base deployment available, let's say traefik for example. You can create a branch specific and tag it using this format: `branch-vM.M.P` that would be for example: `daemonset-v.1.0.0` .

## Content

A package specific repository root folder should be organised like this:

> ```
> /base
> /docs
> /examples
> README.md
> ```

The **base** folder should contain a basic non specific (as much as possible) kustomize config.

The **docs** folder contains docs with specific config explanations. This folder is optional.

The **examples** folder contains overlays examples. It should at least contains one working example of a deployment on any plateform and with any type of ingress. It can contains specific documentation.

The **README.md** must contains a description of what is deployed if it's not in the docs folder and if it's a package specific repository a link to the official doc of the app.

## Package base folder restrictions

The minimal configuration thinking put some restriction on the resources that can be included inside the base folder, some of them should not exists within:

* namespace : You can't force people to use the same namespace as you, no valid base should contains a namespace (excluding private packages)
* ingress : For the best compatibility, do not create an Ingress file in the base folder as people may be using simple service LoadBalancer or service NodePort. Put ingress only inside overlays examples.
* volumes : There is as much PVC as the number of cloud providers out there (maybe more) having a full working configuration may be hard with this. Having version specifics packages may be a good idea.

## Skeleton

There is a [package skeleton](../examples/package-skeleton) in the examples folder