# 진부한 ToDo

---

### 누구나 한번쯤 만들어봤을 진부한 ToDo 앱을 개발합니다.

---

#### 사용스택:

- **[Go 1.21](https://go.dev/)**
- **[Echo](https://echo.labstack.com/)**
- **[GORM](https://gorm.io/ko_KR/docs/index.html)**
- **[Godotenv](https://github.com/joho/godotenv)**
- **[Golang-jwt](https://github.com/golang-jwt/jwt)**
- **[PostgreSQL 15.1](https://www.postgresql.org/)**
- **[Docker/Docker Compose](https://docs.docker.com/compose/)**

```shell
    # 모듈 생성
    go mod init trivial-todo-be
    # docker compose로 db 생성
    docker compose up -d
    # db 삭제
    docker compose rm -s -f -v
    # 로컬 개발환경에서 어플리케이션 실행
    GO_ENV=dev go run .
    # 로컬 개발환경에서 db 마이그레이션
    GO_ENV=dev go run migrate/migrate.go
```
    