# beegodemo

Beego Demo GoLang

## Step to RUN projects

**1. Fetch Dependency:**

```sh
go mod tidy
```

**2. Environment Settings:**

Copy `docker/.env.sample` to `docker/.env` and write environment details. Then export using this command.

```sh
export $(cat docker/.env | xargs)
```

**3. Build Project:**

```sh
go build
```

**4. Run Application:**

```sh
./beegodemo
```

Or run with docker:
```sh
cd docker
docker-compose up --build
```

Now your api is running on 3000 port (if you use docker then port is 3001).

**5. Routes:**

| Request Method | Endpoint                             | Path Parameters | Body (application/json)                                        | Query Parameters |
| -------------- | ------------------------------------ | --------------- | -------------------------------------------------------------- | ---------------- |
| `GET`          | `http://localhost:3000`              |                 |                                                                |                  |
| `GET`          | `http://localhost:3000/docs`              |                 |                                                                |                  |
| `GET`          | `http://localhost:3000/articles`     |                 |                                                                |                  |
| `POST`         | `http://localhost:3000/articles`     |                 | `{ "title": "your title", "description": "your description" }` |                  |
| `GET`          | `http://localhost:3000/articles/:id` | `id`            |                                                                |                  |
| `DELETE`       | `http://localhost:3000/articles/:id` | `id`            |                                                                |                  |

## Deploying to Heroku

```sh
heroku create
git push heroku master
heroku open
```

Or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
