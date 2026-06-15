# telebot Helm chart

A Helm chart for deploying the **telebot** Telegram bot on a Kubernetes cluster.

## Prerequisites

- Kubernetes cluster + `kubectl` configured
- Helm 3+ installed
- A Telegram bot token (`TELE_TOKEN`) from [@BotFather](https://t.me/BotFather)

## Installation

### 1. Download the packaged chart from the GitHub release

```bash
helm install telebot \
  https://github.com/vladyslav-ops/telebot/releases/download/v1.0.0/telebot-1.0.0.tgz \
  --set secret.token=<YOUR_TELE_TOKEN>
```

The chart creates a Kubernetes `Secret` from `secret.token` and injects it into
the container as the `TELE_TOKEN` environment variable.

### Alternative: use a pre-created Secret

If you prefer to manage the token outside Helm, create the Secret yourself and
point the chart at it:

```bash
kubectl create secret generic telebot --from-literal=token=<YOUR_TELE_TOKEN>

helm install telebot \
  https://github.com/vladyslav-ops/telebot/releases/download/v1.0.0/telebot-1.0.0.tgz \
  --set secret.name=telebot
```

## Configuration

| Parameter           | Description                                   | Default               |
|---------------------|-----------------------------------------------|-----------------------|
| `image.repository`  | Container image repository                    | `quay.io/vladyslav-ops` |
| `image.tag`         | Image tag                                     | `v1.0.0`              |
| `image.arch`        | Architecture appended to the tag              | `amd64`               |
| `secret.token`      | Telegram bot token (creates a Secret)         | `""`                  |
| `secret.name`       | Name of an existing Secret to use instead     | chart fullname        |
| `replicaCount`      | Number of replicas                            | `1`                   |

The resulting image reference is built as:
`{{ image.repository }}/telebot:{{ image.tag }}-{{ image.arch }}` →
`quay.io/vladyslav-ops/telebot:v1.0.0-amd64`.

## Verify the deployment

```bash
helm list
kubectl get pods -l app.kubernetes.io/name=telebot
kubectl logs -l app.kubernetes.io/name=telebot
```

## Uninstall

```bash
helm uninstall telebot
```
