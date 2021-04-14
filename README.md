# Local dev env with Skaffold and Kind

## Requirements

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
  Our local k8s cluster of choice

- [ctlptl](https://github.com/tilt-dev/ctlptl/#how-do-i-install-it)
  Abstraction to manage local cluster

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
  To interact with our local cluster

- [helm](https://helm.sh/docs/intro/install/)
  To render our manifests templates

- [skaffold](https://skaffold.dev/docs/install/)
  To manage our dev environment (talk to local cluster, render templates, sync files)

## What is Skaffold?
