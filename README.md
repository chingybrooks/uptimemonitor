# UPTIME MONITOR

![](/docs/uptimemonitor.png)

[![Go](https://github.com/airlabspl/uptimemonitor/actions/workflows/go.yml/badge.svg)](https://github.com/airlabspl/uptimemonitor/actions/workflows/go.yml)

[uptimemonitor.dev](https://uptimemonitor.dev)

© 2025 AIR Labs

## Yet another uptime monitor?

Yes, but with following constraints:

- 100% self-hosted
- Single binary when build
- 0 configuration
- Simple functionality without bloat

## Installation

1. Download the latest release to a Linux VPS (check the
   [Releases page](https://github.com/airlabspl/uptimemonitor/releases)):

```bash
wget https://github.com/airlabspl/uptimemonitor/releases/download/v0.0.2-alpha/uptimemonitor
chmod +x uptimemonitor
```

2. Run it

```bash
./uptimemonitor -addr=":3000"
```

You can also use provided `Dockerfile` and `docker-compose.yml` for reference to
run the app using Docker.

For example, to run a local copy without https enabled:

```bash
SECURE=false COOLIFY_VOLUME_APP=~/testdata docker-compose up
```

## First run

When you first run the application you will be asked to create a new account,
which you will then use to log in.

![](/docs/setup.png)

## Usage

To start monitoring a given url, add the url on a `/new` page. You can specify
the HTTP method the check will use (GET, POST, etc.) as well as add custom
headers and body to the request.

![](/docs/new.png)

## Incidents

When the check fails, the incident will be created and will stay open until the
upcoming check succeeds. You can dig into the failing request details on a
incident page.

![](/docs/incident.png)

## Webhooks

If you want to get notified when an incident happens, you can also configure a
webhook notification to a given url of choice (e.g. slack or google chat webhook
url). You can use `{{ .Url }}` and `{{ .StatusCode }}` variables inside a body
to pass those information to the webhook.

![](/docs/webhook.png)

## Https

Because the app uses secure cookies for authentication, it is required to use
some reverse proxy with https certificate configured.

If you want to test the app without secure cookies enabled, provide the
`-secure=false` flag.

## Backups

This app uses a sqlite database so to backup the data just copy the
`uptimemonitor.sqlite*` files to a new location.

## Pricing

The app is free to use but if you are using it commercially and can aford a
small donation- please use Github Sponsors.

The donations of $50 a month and above will be featured in a sponsors area
inside the application dashboard.

## Roadmap

- Monitor status badges
- Change password
- Manage users
- Timezones
- Reset password via cli
- Add "Test Webhook" button with fake incident
- Sort monitors option (with localstorage sync) by created_at/name(domain)
