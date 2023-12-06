# 진부한 ToDo

---

### 누구나 한번쯤 만들어봤을 진부한 ToDo 앱을 개발합니다.

---

#### 사용스택:

- **Go 1.21**
- **Echo**
- **React 18**

```shell
  # create module
    go mod init go-rest-api
    # start db
    docker compose up -d
    # remove db
    docker compose rm -s -f -v
    # start app
    GO_ENV=dev go run .
    # run migrate
    GO_ENV=dev go run migrate/migrate.go
```
    