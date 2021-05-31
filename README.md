# etrackings-linebot

### สร้างไฟล์ .env

```sh# GO Service

TZ=Asia/Bangkok
FIBER_MODE=LOCAL

# LINE

LINE_CHANNEL_ID=<You LINE_CHANNEL_ID>
LINE_CHANNEL_SECRET=<You LINE_CHANNEL_SECRET>
LINE_CHANNEL_TOKEN=<You LINE_CHANNEL_TOKEN>

# ETrackings API

ETRACKINGS_API_KEY=<You ETRACKINGS_API_KEY>
ETRACKINGS_KEY_SECRET=<You ETRACKINGS_KEY_SECRET>
```

### Load go lib
```sh
go mod vendor
```

### Run

```sh
go run main.go
```
