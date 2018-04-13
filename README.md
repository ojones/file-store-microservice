# simple-storage-server
Simple storage server homework assignment for interview. Go microservice for managing multipler user accounts for personal file storage. 

Completed all instructions. Included best practices but not absolutely everywhere due to time. Should give a general idea of my capabilities.

Suggestions welcome.

### Build and run
From within repo:
```
go build && ./simple-storage-server
```
Localhost address is http://localhost:9999/

File upload form field is "file"

### Test
```
go test
```

### TODO if this were a production project
Handle space limitations<br>
More test coverage - gave just a few examples<br>
Store registered users to disk - currently an in memory map, <br>
Create and cleanup folders on startup - folders included in project<br>
Logger middleware instead - currenlty just does validation<br>
Integration tests<br>
Read configs from text file<br>
Docker image<br>