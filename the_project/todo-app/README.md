# Todo App

A Node.js web server developed for the [DevOps with Kubernetes course](https://courses.mooc.fi/org/uh-cs/courses/devops-with-kubernetes-2026).

The application starts an HTTP server and uses the `PORT` environment variable to determine which port to listen on. If no port is specified, it defaults to port 3000.

## Requirements

- Node.js
- Docker
- kubectl
- k3d

## Build

Build the Docker image:

```bash
docker build -t todo-app:latest .
```

## Deploy to Kubernetes

Create a local k3d cluster if you don't already have one:

```bash
k3d cluster create dev
```

Import the image into the cluster:

```bash
k3d image import todo-app:latest -c dev
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
kubectl logs -f deployment/todo-app-dep
```

Expected output:

```text
Server started in port 3000
```

## Cleanup

Delete the Deployment:

```bash
kubectl delete -f manifests/deployment.yaml
```

Delete the local cluster (only if you no longer need it for other exercises):

```bash
k3d cluster delete dev
```

## Course Exercise

- Chapter: 2 — First deploy
- Exercise: 1.2 — The project