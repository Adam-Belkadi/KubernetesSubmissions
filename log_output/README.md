
# Log Output

A Go application developed for the [DevOps with Kubernetes course](https://courses.mooc.fi/org/uh-cs/courses/devops-with-kubernetes-2026).

The application generates a random UUID at startup and prints it to standard output every five seconds, alongside the current timestamp.

## Requirements

- Go
- Docker
- kubectl
- k3d

## Build

Build the Docker image:

```bash
docker build -t log-output:latest .
```

## Deploy to Kubernetes

Create a local k3d cluster if you don't already have one:

```bash
k3d cluster create dev
```

Import the image into the cluster:

```bash
k3d image import log-output:latest -c dev
```

Apply the Kubernetes Deployment:

```bash
kubectl apply -f manifests/deployment.yaml
```

## Verify

Check the Pod status:

```bash
kubectl get pods
```

View the application logs:

```bash
kubectl logs -f deployment/log-output
```

Expected output:

```text
2026-11-08T12:55:39Z: <generated-uuid>
2026-11-08T12:55:44Z: <same-generated-uuid>
```

## Cleanup

Delete the Deployment:

```bash
kubectl delete -f manifests/deployment.yaml
```

Delete the local cluster:

```bash
k3d cluster delete dev
```

## Course Exercise

- Chapter: 2 — First deploy
- Exercise: 1.1 — Getting started
