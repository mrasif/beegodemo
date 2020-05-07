# beegodemo
Beego Demo GoLang

## Step to RUN projects.

**1. Fetch Dependency:**

```
go mod tidy
```

**2. Environment Settings:**

Copy `.env.sample` to `.env` and write environment details. Then export using this command.
```
export $(cat .env | xargs)
```

**3. Build Project:**

```
go build
```

**4. Run Application:**

```
./beegodemo
```
Now your api is running on 8080 port.

**5. Routes:**

| Requesst Method | Endpoint | Path Parameters | Body (application/json) | Query Parameters |
|-|-|-|-|-|
| GET   | http://localhost:8080 | | | |
| GET   | http://localhost:8080/articles | | | |
| POST  | http://localhost:8080/articles | | { "title": "your title", "description": "your description" } | |
| GET   | http://localhost:8080/articles/:id | id | | |
| DELETE   | http://localhost:8080/articles/:id | id | | |
