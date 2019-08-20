# Multiple folder deployment

The most advance usage, using one armada.yaml file at the root of the apps folder, you can deploy multiple apps from different folders. The results yaml file will all be created at the root folder with the armada file.

```bash
>/apps-folder
>├── armada.yaml
>├── app-1
>|   └── overlays
>│       └── apply
>│           └── kustomization.yaml
>└── app-2
>    └── overlays
>        └── apply
>            └── kustomization.yaml
```

We don't need to specify the overlays as the apps folder are differents.

The generated yaml file will each have the name of the folder.