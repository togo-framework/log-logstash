<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/log-logstash</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/log-logstash"><img src="https://pkg.go.dev/badge/github.com/togo-framework/log-logstash.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/log-logstash
```

<!-- /togo-header -->

<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">log-logstash</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

**Logstash** log shipping for togo. Streams your app's `slog` logs as JSON lines to
a [Logstash](https://www.elastic.co/logstash) TCP input (e.g. into the ELK stack),
in addition to the existing local/stdout output.

```bash
togo install togo-framework/log-logstash
```

Install alongside `togo-framework/log`. Blank-importing the plugin registers it.

## Env

| Var | Required | Description |
|---|---|---|
| `LOGSTASH_ADDR` | yes | `host:port` of a `tcp { codec => json_lines }` input. No-op when empty. |

Example Logstash input:

```ruby
input { tcp { port => 5000 codec => json_lines } }
```

## How it works

On boot (after the `log` plugin) it wraps the kernel logger so every record is
**also** written as a JSON line over a reconnecting TCP connection to Logstash.
Connection failures are swallowed so logging never breaks the app.

MIT © togo-framework

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
