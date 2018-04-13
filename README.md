# simple-storage-server
Simple storage server homework assignment for interview. Go microservice for managing multipler user accounts for personal file storage. 

Completed all instructions. Included some best practices but not everywhere due to time. Should give a general idea of my capabilities.

Suggestions welcome.

### Build and run
From within repo:
```
go build && ./simple-storage-server
```

### Test
```
go test
```
For coverage:
```
go test -cover
```

### TODO if this were a production project
More test coverage - gave just a few examples<br>
Store registered users to disk - currently an in memory map, <br>
Create and cleanup folders on startup - folders included in project<br>
Logger middleware instead - currenlty just does validation<br>
Integration tests<br>
Read configs from text file<br>
Docker image<br>