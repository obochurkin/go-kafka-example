# UI obsidiandynamics/kafdrop:latest
url to have a kafka UI
```localhost:3000``` 

## how to send messages
to send messages ```go_server``` container should be UP

to send messages used endpoint ``` POST localhost:8080/api/v1/send-message```
example json payload: 
```
{
  "message": "Test message 1" 
}
```
## how to consume messages
to consume messages ```go_consumer``` container should be UP