# Terraform module - cw sns to slack Lambda


## Testing

Prerequisites:
- SAM cli installed
- You replaced "replace_with_your_webhook_url" with your slack webhook url


### SAM local with alarm message
```
sam local invoke "CloudWatchSNSToSlackLambda" --event events/alarm.json
```

### SAM local with ok message
```
sam local invoke "CloudWatchSNSToSlackLambda" --event events/ok.json
```
