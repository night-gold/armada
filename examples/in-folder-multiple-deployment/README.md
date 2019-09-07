# In folder multiple deployment

You can use armada to deploy the same app multiple time using multiple overlays, for now this function works only if the namespace is different for each deployment.

The folder should be organised as follow:

>```
>/app
>├── armada.yaml
>└── overlays
>    ├── apply
>    │   └── kustomization.yaml
>    └── apply2
>        └── kustomization.yaml
>```

For the deployment of the first overlays no need to specify the overlay we need to use as it's the default value, but to deploy the second one we must add the overlays with the name of the folder.

The results yaml file will have a name like this: app-apply-2.yaml for the non default overlays in the example.