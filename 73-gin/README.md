## To run postgres container


```bash
docker run -d --name pgdb -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=usersdb --network demo2-network postgres
```
```bash
docker run -d --name pgui --network demo2-network -p 8089:8080 adminer
```
