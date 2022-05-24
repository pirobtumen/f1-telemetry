# f1-game-telemetry

Tool for visualizing and storing real-time telemetry data for the F1 Game (2021). 

[You can read about the project in my blog post](https://www.pirobits.com/post/f1-telemetria-tiempo-real-golang-influxdb/).

![Telemetry dashboard](https://raw.githubusercontent.com/pirobtumen/f1-telemetry/main/dashboards/dashboard_example.PNG)

## Introduction

I've been playing the F1 2021 Game this year, and I saw it sends telemetry data to an UDP host.

I found that they explains the packet enconding, so you can parse all the data. So why not improve my coding skills in Go with a real time time-series tool!

This tool stores the telemtry info in a database (currently I've chosen InfluxDB for time series) and Go for the development language.

For the charts I've used the InfluxDB UI, that it's quite similar to Grafana. The thing is that both (Grafana and InfluxDB UI) limits the refresh rate up to 5 seconds... I would like real time updates, but it's OK for now. 

## Roadmap

- [ ] Full telemetry data support.
- [ ] Parse all the relevant packages.
- [ ] Custom UI
- [x] Metrics: car speed, braking [0, 1], throttling [0, 1] and steering [-1, 1].
- [x] Last lap time.
- [x] Base dashboard with auto-updating (5s).
- [X] Real time UI updates. => Chronograf now supports 1s update interval.
- [x] Client live reload for development.
- [x] Improve local development environment.

I haven't thought anymore really, so I'll be adding more features as I come up with new ideas :D

## Setup

> You need docker/docker-compose installed.


> Docker + UDP doesn't work properly. There are problems with UDP conexions so you need to run the server without Docker.

In one terminal run this to initialize the environment:

```
$ make start
```

Run server in development mode:

> I use air for live reloading `go install github.com/cosmtrek/air@latest`.

> You need to update your .env DB access config as explained in the next section.

```
$ make dev
```

### Database / Charts

Now access http://localhost:8086 form your browser, configure your user/password/company/bucket. Set those correct values in the `.env` file. 

Finally, navigate to the "token" section and copy the access token into  the `.env` file (DB_TOKEN).

#### Default dashboard

I've created a default dashboard that shows last lap time, current speed, speed over time, braking vs throttling over time, and steering over time.

Go to the dashboard section and click "New -> Import" and select the file `f1-game-telemetry/dashboards/influxdb_f1.json`.
