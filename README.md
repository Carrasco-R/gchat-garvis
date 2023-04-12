# G-Chat Chatbot 🤖

Only route available is `/recognize`

```sh
curl --request POST \
  --url ${CONTAINER_IP}/recognize \
  --header 'content-type: application/json' \
  --data '{"person":"My good friend","reason":"For being a genius"}'
```

Runs on docker!
