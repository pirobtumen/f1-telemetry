# f1-game-telemetry

Tool for visualizing and storing real-time telemetry data for the F1 Game (2021). 

> Early stage project.

![Telemetry dashboard](https://raw.githubusercontent.com/pirobtumen/f1-telemetry/main/dashboards/dashboard_example.PNG)

## Introduction

I've been playing the F1 2021 Game this year, and I saw it sends telemetry data to an UDP host.

I found that they explains the packet enconding, so you can parse all the data. So why not improve my coding skills in Go with a real time time-series tool!

This tool stores the telemtry info in a database (currently I've chosen InfluxDB for time series) and Go for the development language.

For the charts I've used the InfluxDB UI, that it's quite similar to Grafana. The thing is that both (Grafana and InfluxDB UI) limits the refresh rate up to 5 seconds... I would like real time updates, but it's OK for now. 

## Roadmap

- [x] Metrics: car speed, braking [0, 1], throttling [0, 1] and steering [-1, 1].
- [x] Last lap time.
- [x] Base dashboard with auto-updating (5s).
- [ ] Custom UI
- [ ] Real time UI updates.
- [ ] Client live reload for development (air?).
- [ ] Parse all the relevant packages.
- [ ] Improve local development environment.

I haven't thought anymore really, so I'll be adding more features as I come up with new ideas :D

## Setup

> You need docker/docker-compose installed.


> Windows: if you try to run this on Windows + WSL2, there are problems with UDP conexions so you need to run the project without Docker.

In one terminal run:

```
$ make start
```

You can check the logs running the command:

```
$ make logs
```

> You need to update your .env DB access config as explained previously.


### Database / Charts

Now access http://localhost:8086 form your browser, configure your user/password/company/bucket. Set those correct values in the `.env` file. 

Finally, navigate to the "token" section and copy the access token into  the `.env` file (DB_TOKEN).

#### Default dashboard

I've created a default dashboard that shows last lap time, current speed, speed over time, braking vs throttling over time, and steering over time.

Go to the dashboard section and click "New -> Import" and select the file `f1-game-telemetry/dashboards/influxdb_f1.json`.
