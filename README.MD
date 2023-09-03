# Tesselite.App 


### version: v0.2.x


### test 

````bash
$ docker compose up -d
````

### database

restore database from .sql dump:
````bash
$ tessel mysql -U tesselite -P tesselite sync -b tesselite.sql -d tesselite
````