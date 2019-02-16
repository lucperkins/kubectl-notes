# kubectl-notes

This is an extremely simple and useless extension for [kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/) that lets you add notes to Kubernetes Pods. It's intended to illustrate how you might build your own kubectl extension.

## Usage

Here's an example:

```bash
# Add a note to a Pod
kubectl notes add --namespace my-ns --pod my-pod "This Pod has been acting shady"

# Fetch a note from a Pod
kubectl notes get --namespace my-ns --pod my-pod
This Pod has been acting shady

# Delete a note
kubectl notes delete --namespace my-ns --pod my-pod
```

You can substitute `-n` for `--namespace` and `-p` for `--pod`.

The `default` namespace will be used if none is provided.

## Installation

You can install it using `go install`:

```bash
go install github.com/lucperkins/kubectl-notes
```

If `$GOPATH/bin` is on your `PATH`, you can invoke it directly from `kubectl` as in the example [above](#usage).

## How it works

`kubectl-notes` simply adds an annotation to the Pod with the key `note`.
