# pods_management

Migration Docker usage
```
docker run -v /home/emumba/learning/go/pods_management/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/pods_management1 up
```

Separate container to mount migration
```
cd migrator
docker build -t migrator .
docker run --network host migrator -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/pods_management1 up
```