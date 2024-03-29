# Armada fields

## Command line option

- **-f** Path to the file with the armada configuration.
- **-o** Default overlays you are deploying, value can be overrided by the overlays field in armada file.
- **-a** Auto apply the newly generated configuration to your cluster.

## Armada file other fields

- **namespaces**: A list of string that will be used to create namespaces.

## Armada file package fields

- **name**: This is the name of the package going to be deployed.
- **deployment**: This field contains everything in relation to the deployment.
    - **folder** : If you want to make multiple deployment of the same app using different names, or just change the app folder name. Default value is set to the repository name value.
    - **overlays** : The kustomize overlays you want to apply. Default value is `apply`.
    - **wait** : This numerical value is used to wait for the pod to finish it's start in case we need something from the pod to deploy another app without error (CRD)
- **variables**: This field contains an array of variable that are used for the templating, name of the environment variables and of the template variable.