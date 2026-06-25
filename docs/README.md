# log-logstash — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package logstash ships togo's slog logs as JSON lines to a Logstash TCP input,
in addition to the app's existing log output. Install alongside togo-framework/log;
blank-import registers it.

Env: LOGSTASH_ADDR (host:port of a `tcp { codec => json_lines }` input — required,
no-op when empty).

## Install

```bash
togo install togo-framework/log-logstash
```

Set `LOG_DRIVER=logstash`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `LOGSTASH_ADDR` | _see provider docs_ |

## Usage

```go
// Structured logs/errors are forwarded to the configured sink automatically
// once this driver is installed and its env is set.
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/log-logstash
- README: ../README.md
