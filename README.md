# Local dev env with Skaffold and Kind

## Requirements

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
  Our local k8s cluster of choice

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
  To interact with our local cluster

- [helm](https://helm.sh/docs/intro/install/)
  To render our manifests templates

- [skaffold](https://skaffold.dev/docs/install/)
  To manage our dev environment (talk to local cluster, render templates, sync files)

## Before start

- Create Kind cluster
kind create cluster --name skaffold-workshop --config kind.conf.yaml

## Examples

### Getting Started

Simple go application deployed using a k8s pod.

### Helm deployment

You can deploy multiple releases with skaffold, each will need a chartPath, a values file, and namespace.
Skaffold can inject intermediate build tags in the the values map in the skaffold.yaml.

### Typescript

Using express server with typescript

### Helm Workshop

Code used in previous helm workshop