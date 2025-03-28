# Traefik Path to Header Middleware Plugin

A simple Traefik plugin that adds the request path as a header to all requests. Built with Yaegi for K3s compatibility.

## Quick Start

1. Add to your Traefik configuration in K3s:

```yaml
apiVersion: helm.cattle.io/v1
kind: HelmChartConfig
metadata:
  name: traefik
  namespace: kube-system
spec:
  valuesContent: |-
    experimental:
      plugins:
        pathHeader:
          moduleName: "github.com/kenenisa/traefik-path-to-header-middleware-plugin"
          version: "v0.1.0"
```

2. Restart Traefik:
```bash
kubectl get pod -n kube-system -l app=traefik -w
```

## What it Does

- Adds `X-Request-Path` header containing the request URL path
- Works with K3s and Traefik
- Uses Yaegi for runtime interpretation (no compilation needed)

## Development

```bash
cd path-header
go mod init github.com/kenenisa/traefik-path-to-header-middleware-plugin
```

## License

MIT
