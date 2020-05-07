# beegodemo

Beego Demo GoLang

## Step to RUN projects

**1. Fetch Dependency:**

```sh
go mod tidy
```

**2. Environment Settings:**

Copy `.env.sample` to `.env` and write environment details. Then export using this command.

```sh
export $(cat .env | xargs)
```

**3. Build Project:**

```sh
go build
```

**4. Run Application:**

```sh
./beegodemo
```

Now your api is running on 8080 port.

**5. Routes:**

| Request Method | Endpoint                             | Path Parameters | Body (application/json)                                        | Query Parameters |
| -------------- | ------------------------------------ | --------------- | -------------------------------------------------------------- | ---------------- |
| `GET`          | `http://localhost:8080`              |                 |                                                                |                  |
| `GET`          | `http://localhost:8080/articles`     |                 |                                                                |                  |
| `POST`         | `http://localhost:8080/articles`     |                 | `{ "title": "your title", "description": "your description" }` |                  |
| `GET`          | `http://localhost:8080/articles/:id` | `id`            |                                                                |                  |
| `DELETE`       | `http://localhost:8080/articles/:id` | `id`            |                                                                |                  |

## Deploying to Heroku

```sh
heroku create
git push heroku master
heroku open
```

Or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
