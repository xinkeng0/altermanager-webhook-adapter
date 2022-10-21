# webhook-adapter
Simple AlterManager WeWork Robot Webhook adapter

Support Webhooks
- [x] WeWork Robot Webhook 
# how to use
With GO
1. Run the main.go

With Docker
1. docker build . altermanager-webhook-adapter:0.1
2. docker run -p 8080:8080 -d altermanager-webhook-adapter:0.1

then send request to the url

POST http://127.0.0.1:8080/?url={your webhook url}
