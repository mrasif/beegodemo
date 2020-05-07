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
