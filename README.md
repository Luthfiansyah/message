# Warpin Message API

## How to Install ?

Do the following command
```
mkdir -p your_path/warpin-message
cd your_path/warpin-message
git clone https://github.com/Luthfiansyah/warpin-message.git
```

And you should now find the source code under `/warpin-message` directory.

## Directory Structure

```
  + your_path/
  |
  +--+ src/
  |  |
  |  +--+ github.com/
  |     |
  |     +--+ warpin-message/
  |         |
  |         +--+ main.go
  |            + routing.go
  |            + app/
  |            |
  |            +--+ + handlers/
  |                 + models/
  |                 + types/
  |                 + repositories/
  |            + config/
  |            + database/
  |            + ... any other source code
  |            + ... executable file .warpin-message
```

Get into the `warpin-message` folder by

```
cd warpin-message/
```

Under that folder you need to run below command to install required library dependency file go.mod

```
go mod init github.com/Luthfiansyah/warpin-message
go get 
```
Add new dependency
```
go get -u github.com/lorem/ipsum
```

## How to compile and run the code ?
And then following this command :

Just run this command
```
docker-compose build
docker-compose up
```

Shutting down app Just run this command
```
docker-compose down
```

The configuration is described in the config.toml file.
In particular, in the [server] section, you will find:
- mode: set it to development, production, localserver
- port: the port the server will bind to
- debug: the debug flag keeps greater amount of logging enabled, including gin logging

## How do we know this application server is running?
You can try calling the root context of url
for the example, the base url of this server is

http://localhost:8888

Just call it in normal browser and you will get message something like
```
{
    "message":"Warpin Message Run On DEV Mode",
    "start_time":"2020-03-17 19:41:13"
}

```

## How to test a code ?

to run unit testing cannot be done in the docker-compose, rabbitmq must be installed on premise on the computer, and change RABBITMQ_HOST rabbitmq from rabbitmq -> localhost in the config.toml file

To run all the test in root folder
```
go test -v
```

To run specific test file (for the example only `message_test.go` file)
```
go test message_test.go
```

## How to add new service API in app ?

Follow the following structure

```
1. At `app/handlers/{handlername}.go` add new handlers
2. At `app/models/{modelname}.go` add new model
3. At `app/types/{typename}.go` add new type
4. At `app/repositories/{repositoryname}.go` add new repository

```
## How to commit and push the change into repository ?

### Branch Conventions

There are 4 kinds of branch
- features-{your-feature-name}
- enhance-{your-enhance-name}
- hotfix-{your-hotfix-name}
- bugfix-{your-bugfix-name}

### How To Create New Tag

- Checkout to master branch
- Run command to create new tag, like the following example
    ```
    git tag -a v1.0.0.0 -m "add basic trans and internal routes"
    ```
- Push your tags with command below
    ```
    git push origin --tags

## How To Deploy

- Make sure your changes/branch already in production
- Create a tag with following format v1.0.0.0

    ```v{MajorRelease}.{MinorRelease}.{MajorBugfix}.{MinorBugfix} and related message```
- Inform in Deployment group

### How to build a Docker Image ?

Just run this command
```
docker login registry.github.com
docker build -t registry.gitlab.com/Luthfiansyah/warpin-message .
docker push registry.gitlab.com/Luthfiansyah/warpin-message
docker pull registry.gitlab.com/Luthfiansyah/warpin-message
docker run --name=warpin-message -p 8888:8888 warpin-message_app:latest >> /var/www/warpin-message/warpin-message.log
docker run -d --name=warpin-message -p 8888:8888 warpin-message_app:latest 
```

### ENDPOINT LIST 
- GET http://localhost:8888
- POST http://localhost:8888/v1/message 
- GET http://localhost:8888/v1/message

### Test Starting API

PING SERVER
```
PING GET http://localhost:8888
=====
{
    "message":"Warpin Message Run On DEV Mode",
    "start_time":"2020-03-17 19:41:13"
}
```

ADD NEW MESSAGE

- REQUEST
```
curl -i -X POST -H "Content-Type: application/json" -d '{"text":"Hola test message !!"}' http://localhost:8888/v1/message
```

- RESPONSE
```
{
   "general_response":{
      "response_status":true,
      "response_code":0,
      "response_message":"Success",
      "response_timestamp":"2020-03-17 19:59:16.601810909 +0700 WIB"
   },
   "result":{
      "text":"SentHola 456 !!"
   }
}
```

GET ALL MESSAGE 

- REQUEST
```
curl http://localhost:8888/v1/message | json_pp
```

- RESPONSE
```
{
   "result" : [
      {
         "text" : "Hola test message !!"
      }
   ],
   "general_response" : {
      "response_message" : "Success",
      "response_timestamp" : "2020-03-17 20:00:32.77725705 +0700 WIB",
      "response_status" : true,
      "response_code" : 0
   }
}

```

## Redis Get Key Example

exec redis server from another host using port 6377 

```
$redis-cli -h localhost -p 6377 

localhost:6379>keys "*"
1) "29b1de684be26d084376d02695c1a2fa"

localhost:6379>get 29b1de684be26d084376d02695c1a2fa
"2020-03-17 19:41:13"
```

## Access Rabbitmq Admin

```
open browser and enter http://localhost:15672
username : guest
password : guest
```


====================================================================================

## warpin-message 

The server is at: warpin-message

```
DEVELOPMENT:
   host   : http://localhost:8888
   branch : dev
   mode   : development (local)
   folder : warpin-message

STAGING   :
   host   : http://staging.api.lorem-ipsum.com
   branch : staging
   mode   : staging (staging)
   folder : warpin-message

PRODUCTION:
   host   : https://api.lorem-ipsum.com
   branch : master
   mode   : production (production)
   folder : warpin-message
```

Author Moh Reza Luthfiansyah