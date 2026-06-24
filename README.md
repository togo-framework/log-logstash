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
