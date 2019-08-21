# Armada fields

- **repository** : This field has the name of the repository we want to use. It also serves as the default value for the folder and the result file name. This field has no default value, it's a necessary field.
- **git** : The url of the git we are going to look for the package, gitlab, github, etc... The default value is https://github.com.
- **version** : The tag or branch we want to use for the package. Default value is `master`.
- **user** : The user field is here to target the repository owner. Default value is `armada` (This tool name).
- **folder** : If you want to make multiple deployment of the same app using different names, or just change the app folder name. Default value is set to the repository name value.
- **overlays** : The kustomize overlays you want to apply. Default value is `apply`.