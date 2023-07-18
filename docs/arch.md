## Achitecture

<hr/>
Implementing Hexagon Achitecture for software pattern. This pattern isolate external systems and other external code such as user interfaces and databases from the core application. Hexagon achitecture separate the code into three large formalized areas: user-side (interface), business logic (core) and server-side (infrastructure).

### Hexagon archicture pattern

<img src="hexa_arch.jpeg">

### User-Side (Interface)

User-facing interface that take the input data from user and repackage it in from that is convinient for this use cases, than return data back in a form that is convinient for displaying it back for user (HTTP, HTML, JSON, CLI etc).

### Business Logic (Core)

The core part of a service that will isolate the business logic from external systems (user-side or server-side). The communication between core and external systems must be through a port.

### Server-Side (Infrastructure)

Contain technology tools (like repository, access to external API/services, message brokers, frameworks, etc) and adapt its input or output to a port, which fits the application core needs.

### Implementation

```sh
app/                               # directory for application initialization
    api/                           # app that use restapi
        extl/                      # external directory
            main.go                # external main file
        intl/                      # internal directory
            main.go                # internal main file
    worker/                        # the app worker
    migration/                     # migration

docs/                              # documentation
    img/                           # image docs file
    readme/                        # readme docs file
    swagger/                       # directory for store the api documentation

interface/                         # interface directory
    worker/                        # interface worker
    api/                           # interface rest-api
        extl/                      # external interface rest-api
            v1/                    # api version
                ingredient/             # ingredient domain
                    response/      # api response needs (struct response etc)
                    request/       # api request needs (struct request etc)
                    handler.go     # api handler
                routes/            # routes directory
                    middleware/    # middleware for routing
                    routing.go     # declare api routes
        intl/                      # internal interface rest-api
            v1/                    # api version
        common/                    # store common reusable code
        utils/                     # store reusable tools

core/                              # business logic directory
    model/                         # model direcotry
        base.go                    # the base model
    port/                          # port directory
        ingredient/                     # ingredient port directory
            service.go             # port/interface service
            repostiroy.go          # port/interface repository
    ingredient/                         # directory for ingredient service
        service.go                 # ingredient service
        service_test.go            # ingredient test service

infrastructure/                    # directory for technology tools
    repository/                    # repository directory
        mysql/                     # database specification
            ingredient/                 # ingredient domain
                repository.go      # store the ingredient repository
        redis/
        mongodb/
        mock/                      # mock repository directory
          ingredient/                   # ingredient domain
                repository.go      # store the ingredient mock repository
    adapter/

utils/                             # resable tools or code
    config/                        # config file directory
        mongo/
            mongo.go               # mongo configuration
    logger/                        # logger
    helper/                        # helper
```
