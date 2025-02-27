[![Main-Docker](https://github.com/aceberg/exercisediary/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/exercisediary/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/exercisediary)](https://goreportcard.com/report/github.com/aceberg/exercisediary)
[![Maintainability](https://api.codeclimate.com/v1/badges/e8f67994120fc7936aeb/maintainability)](https://codeclimate.com/github/aceberg/ExerciseDiary/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/exercisediary)

> This is a fork off of https://github.com/aceberg/exercisediary. Support the original repository.

<h1><a href="https://github.com/tnvmadhav/ExerciseDiary">
    <img src="https://raw.githubusercontent.com/aceberg/exercisediary/main/assets/logo.png" width="35" />
</a>Exercise Diary</h1>

Workout diary with GitHub-style year visualization

- [Quick start](https://github.com/tnvmadhav/ExerciseDiary#quick-start)
- [Config](https://github.com/tnvmadhav/ExerciseDiary#config)
- [Options](https://github.com/tnvmadhav/ExerciseDiary#options)
- [Local network only](https://github.com/tnvmadhav/ExerciseDiary#local-network-only)
- [Roadmap](https://github.com/tnvmadhav/ExerciseDiary/blob/main/docs/ROADMAP.md)
- [Thanks](https://github.com/tnvmadhav/ExerciseDiary#thanks)


![Screenshot](https://raw.githubusercontent.com/aceberg/ExerciseDiary/main/assets/Screenshot%202023-12-25%20at%2023-19-08%20Exercise%20Diary.png)

## Quick start

```sh
docker run --name exdiary \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/ExerciseDiary:/data/ExerciseDiary \
-p 8851:8851 \
aceberg/exercisediary
```
Or use [docker-compose.yml](docker-compose.yml)


## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8851 |
| THEME | Any theme name from https://bootswatch.com in lowcase or [additional](https://github.com/aceberg/aceberg-bootswatch-fork) (emerald, grass, sand)| grass |
| COLOR | Background color: light or dark | light |
| HEATCOLOR | HeatMap color | #03a70c |
| PAGESTEP | Items on one page | 10 |
| TZ | Set your timezone for correct time | "" |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -d | Path to config dir | /data/ExerciseDiary | 
| -n | Path to local JS and Themes ([node-bootstrap](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap)) | "" | 

## Local network only
By default, this app pulls themes, icons and fonts from the internet. But, in some cases, it may be useful to have an independent from global network setup. I created a separate [image](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap) with all necessary modules and fonts.    
```sh
docker run --name node-bootstrap       \
    -v ~/.dockerdata/icons:/app/icons  \ # For local images
    -p 8850:8850                       \
    aceberg/node-bootstrap
```
```sh
docker run --name exdiary \
    -v ~/.dockerdata/ExerciseDiary:/data/ExerciseDiary \
    -p 8851:8851 \
    aceberg/exercisediary -n "http://$YOUR_IP:8850"
```
Or use [docker-compose](docker-compose-local.yml)

## Roadmap
Moved to [docs/ROADMAP.md](docs/ROADMAP.md)

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/exercisediary/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- [Chart.js](https://github.com/chartjs/Chart.js) and [chartjs-chart-matrix](https://github.com/kurkle/chartjs-chart-matrix)
- Favicon and logo: [Flaticon](https://www.flaticon.com/icons/)