# IT College Schedule Relay

Used for receiving HTTP schedule events and resending it to specified servers

## Config

```jsonc
{
  // working address
  "address": ":8080",

  // array of URLs - recipients
  "clients": [
    "http://1.1.1.1/blablabla",
    "http://1.1.1.1/blabla"
  ]
}
```

## Run using Docker

```sh
$ docker run \
  -p 8080:8080 \
  -v ./config.json:/relay/config.json \
  eliva1e/itc_schedule_relay
```
