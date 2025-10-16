# crawler

### 1) Prerequisites:
1. `Together AI API key`

### 2) Command:
Run this command to crawl website. make sure the API Key have been provided.
```
go run main.go <<website_name>>
```

### Build (Optional) :
Run this command to create binary file
```
go build -o <<file_name>>
```

### Settings (Optional) :
Refer `settings.go` file in `src/controller` folder for configureable settings.
| setting | Value |
| --- | --- |
| instruction | prompt based on what you need |
| AIModel | Refer Together AI website for AI model available |
| Cost | Refer Together AI website for model cost |

### Together AI available models
https://api.together.ai/models
