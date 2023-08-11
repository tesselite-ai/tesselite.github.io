# Dockerized Website

packages Tesselite website into old-fashioned Html with static resources.

### philosophy of this version

*build everything from scratch. build exactly what you want.* 


### deploy loop

To update version online, do this: 

- build [docker](.build/dockerfile) image with GitHub pipeline.


- deploy into [Google Cloud Platform](.deploy/dev) with Terraform.


### dev loop

- edit html [file](index.html) and static [sources](static).


- check in a navigator


- check with docker 

````bash
## build go container
$ docker compose -f .deploy/local/docker-compose.yml build
## run the container
$ docker compose -f .deploy/local/docker-compose.yml up
````

### change Log

[here](changeLog.MD)


