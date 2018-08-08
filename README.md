## Golang+Vue and Docker quick start

### The application consists of:
* golang 1.10
** glide
** httprouter
** godotenv
* postgres 10.4
* redis
* nginx

### Getting started:
* install [docker](https://docs.docker.com/install/) if not already installed
* install [docker-compose](https://docs.docker.com/compose/install/) if not already installed

## Install GIT
```bash
apt-get update && apt-get install git
```

Exec init command (it will create .env files)
```bash
./init.sh
```

### Run docker container

```bash
docker-compose up
```

To get container name run follow command
```bash
docker-compose ps | grep <service_name> | awk '{ print $1 }'
```

To exec commands inside containers use `docker exec -i -t <container_name> bash`

For example to install new API dependency:
```bash
docker exec -i -t jilexandr_api bash
glide get <package_name>
```

- FRONTEND host [http://www.jilexandr.local](http://www.jilexandr.local)
- API host [http://www.api.jilexandr.local](http://www.api.jilexandr.local)

To get access via domain add to /etc/hosts
```
127.0.0.1    jilexandr.local www.jilexandr.local
127.0.0.1    api.jilexandr.local www.api.jilexandr.local
```
