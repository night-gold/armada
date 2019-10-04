# Armada fields

## Command line option

- **-f** Path to the file with the armada configuration.
- **-o** Default overlays you are deploying, value can be overrided by the overlays field in armada file.
- **-a** Auto apply the newly generated configuration to your cluster.

## Armada file fields

- **name**: This is the name of the package going to be deployed.
- **git**: This field contains everything in relation to git.
    - **repository** : This field has the name of the repository we want to use. It also serves as the default value for the folder and the result file name. This field has no default value, it's a necessary field.
    - **git** : The url of the git we are going to look for the package, gitlab, github, etc... The default value is https://github.com. For private repository you should use the `git@github.com` version of the git service you are using.
    - **version** : The tag or branch we want to use for the package. Default value is `master`.
    - **user** : The user field is here to target the repository owner. Default value is `armada` (This tool name).
    - **private** : A bool differentiate the private repository and compose the git url right.
    - **basepath** : The path where the base is inside the repository.
- **deployment**: This field contains everything in relation to the deployment.
    - **folder** : If you want to make multiple deployment of the same app using different names, or just change the app folder name. Default value is set to the repository name value.
    - **overlays** : The kustomize overlays you want to apply. Default value is `apply`.
    - **wait** : This numerical value is used to wait for the pod to finish it's start in case we need something from the pod to deploy another app without error (CRD)