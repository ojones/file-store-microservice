>Go microservice that allows registered users to store files.

# File Store Microservice

   * [Good practices](#good-practices)
   * [Usage](#usage)
   * [Run](#run)
   * [Test](#test)
   * [Before Production](#before-production)
   * [Personal Note](#personal-note)
  
# Good practices
- Use store interface which allows for easy mock testing
- Seperate struct for configs used to instantiate service
- Vendored dependencies
- JWT session validation
- Authorization data set in request context

# Usage
- Register
- Login
- List files
- Put file
- Get file
- Delete file

# Run
From root folder:
```
go build && ./simple-storage-server
```
Localhost address is http://localhost:9999/

Expected file form field is "file"

# Test
```
go test
```

# Before Production
- Handle space limitations
- More test coverage
- Store registered users to disk
- Create and cleanup folders on startup
- Logger middleware
- Integration tests
- Read configs from text file
- Docker image
- Publish API

# Personal Note
What's the big deal with microservices? Honestly, it's just the next logical step in the evolution of service oriented architecture (SOA).

SOA allowed us to escape the single point of failure of having everything run on one server. Back then, network bottlenecks prevented us from over using services but that is much less of problem now.

With clouds, we can freely decouple data and the logic that goes with them across infrastructure. Just imagine your working code stays up. When something goes wrong, only one service goes down. It's a beauatiful dream. The cost is the heavy lifting of configuring, monitoring, and complicating your system.