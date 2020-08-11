## gomock

**Reference**
- https://github.com/golang/mock
- https://blog.51cto.com/9291927/2346777

```
go get github.com/golang/mock/gomock
```

```
mkdir mock
mockgen -source=./infra/db.go -destination=./mock/mock_repository.go -package=mock

```