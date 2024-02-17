# Custom Volume Exporter

## Overview

A lightweight prometheus exporter for monitoring file storage size for a NFS storage class in Kubernetes.

Currently supported for AMD64 architecture for Linux and MacOS.

## Instructions

The binaries are published to the GitHub releases in this repository. Run the binary on your NFS server and point prometheus to it.

To view the metrics, call the API on port 8000.
```
http://localhost:8000
```

## Metrics

### `pvc_volume_used_bytes`

Persistent volume size in bytes.

This metric can be filtered by:
- namespace
- persistent volume name
- persistent volume claim name
