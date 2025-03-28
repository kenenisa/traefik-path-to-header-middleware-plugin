# Traefik Path to Header Middleware Plugin

This Traefik plugin adds the request path as a header to all requests. It's designed to work with K3s and uses Yaegi for runtime interpretation.

## Features

- Adds `X-Request-Path` header to all requests
- Compatible with K3s and Traefik
- Uses Yaegi for runtime interpretation (no compilation needed)

## Installation

1. Add the plugin to your Traefik configuration in K3s:

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

2. Restart Traefik to apply the changes:
```bash
kubectl get pod -n kube-system -l app=traefik -w
```

## Usage

The plugin will automatically add the `X-Request-Path` header to all requests. You can verify this by:

1. Setting up port forwarding:
```bash
kubectl port-forward -n kube-system pod/traefik-<pod-name> 8082:8082
```

2. Making a request to your service and checking the headers.

## Development

This plugin is built using Yaegi, which allows for runtime interpretation without compilation. The plugin is designed to be simple and efficient, adding the request path as a header to all requests.

## License

MIT License # traefik-path-to-header-middleware-plugin
