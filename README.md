# etrackings-linebot

### สร้างไฟล์ .env

```sh# GO Service

TZ=Asia/Bangkok
FIBER_MODE=LOCAL

# LINE

LINE_CHANNEL_ID=<Your LINE_CHANNEL_ID>
LINE_CHANNEL_SECRET=<Your LINE_CHANNEL_SECRET>
LINE_CHANNEL_TOKEN=<Your LINE_CHANNEL_TOKEN>

# ETrackings API

ETRACKINGS_API_KEY=<Your ETRACKINGS_API_KEY>
ETRACKINGS_KEY_SECRET=<Your ETRACKINGS_KEY_SECRET>
```

### Load go lib
```sh
go mod vendor
```

### Run

```sh
go run main.go
```

### Deploy to heroku

```sh
heroku git:remote -a <Your project name>
```

### Push code to deployment

```sh
git push heroku <Your Git branch>
```

ตัวอย่าง

```sh
git push heroku main
```

### View logs on heroku

```sh
heroku logs --tail
```