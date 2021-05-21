# Templating

The concept of template may seem to go against the kustomize view of apps management but it can easily work with it an facilitate it's usage.

For templating we choose simple go templating, and the value can be fetched from environment variable with the same name.

# Warning

- The templating is not here to replace kustomize basic usage, just to ease the use of it and prevent unwanted duplication of resources that we really want to be able to variabilise.
- Templating can't be used with remote base as the base would be called by the kubectl command.

# Usage

For example, we will use it for namespace separation of different environment, you would still be deploying the same dev overlays but to ease the CI deployment and prevent the need to have a manual operation on the folder (normal kustomize flow) we would template some files:
- namespace.yaml
- deployment.yaml, daemonset.yaml or statefulset.yaml
- kustomization.yaml

To create a template of those file just add a .tmpl extension to the files and then use go templating:
```
---
bases:
- "../../base"
namespace: {{or .NAMESPACE "default"}}
patches:
- service.yaml
```

In the example you can see that we templated the kustomization file from an overlays and variabilized the namespace to allow us to be able to create multiple deployment in different namespace and if we did not define anything it would deploy in the default namespace.