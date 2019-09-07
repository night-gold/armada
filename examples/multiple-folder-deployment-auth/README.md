# Multiple folder deployment with authenticathion

Using armada to deploy more that one repository, one with authentication along with a public repository. The results yaml file will all be created at the root folder with the armada file.

To be able to clone the private repository, you will have to have previously setted the ssh-agent with a valid ssh-key authorized on the repository you want to clone.

```bash
eval `ssh-agent -s`
ssh-add PATH/TO/SSH-KEY
```

>```
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
>```

We don't need to specify the overlays as the apps folder are differents.

The generated yaml file will each have the name of the folder.