# justlog [![Build Status](https://github.com/gempir/justlog/workflows/ci/badge.svg)](https://github.com/gempir/justlog/actions?query=workflow%3Aci)

### What is this?
justlog is a bot I maintain for a couple of channels. Its features differ from other bots in that it doesn't support
commands, etc. yet, it focuses on logging and providing an api for the logs.

### Commands
- `!justlog status` will respond with uptime
- `!justlog join gempir,pajlada` will join the channels and append them to the config
- `!justlog messageType gempir 1,2` will set the recorded message types to 1 and 2 in channel gempir (will fetch the userid on its own) 
- `!justlog messageType gempir reset` will reset to default 

### Config

```
{
    "admins": ["gempir"], // will only respond to commands executed by these users
    "logsDirectory": "./logs", // the directory to log into
    "adminAPIKey": "noshot", // your secret api key to access the admin api (documented below)
    "username": "gempbot", // bot username (can be justinfan123123 if you don't want to use an account)
    "oauth": "oauthtokenforchat", // bot token can be anything if justinfan123123
    "clientID": "mytwitchclientid", // your client ID, needed for fetching userids or usernames etc
    "clientSecret": "mysecret", // your twitch client secret
    "logLevel": "info", // the log level, keep this to info probably
    "channels": ["77829817", "11148817"], // the channels (userids) you want to log
    "archive": true, // probably keep to true, will disable gzipping of old logs if false, useful if you setup compression on your own
    "channelConfigs": {
        "77829817": { // only applies to this channel
            "messageTypes": [1] // will only log messages of type 1 (check go-twitch-irc MessageType https://github.com/gempir/go-twitch-irc/blob/master/message.go#L17), default is all 
        }
    }
}
```

### Admin API

`POST` `/admin/channelConfigs/{channelID}`

Will set the messageTypes logged for a channel

###### Headers
```
X-Api-Key: yourcoolapikey
```

###### Body
```
{
    "messageTypes": [1]
}
```

---

`DELETE` `/admin/channelConfigs/{channelID}`

Will reset the messageTypes logged for a channel

###### Headers
```
X-Api-Key: yourcoolapikey
```

###### Body
```
{
    "messageTypes": true
}
```

---


`POST` `/admin/channels`

Will add the channels to log

###### Headers
```
X-Api-Key: yourcoolapikey
```

###### Body
```
{
    "channels": ["77829817"]
}
```

---

`DELETE` `/admin/channels`

Will remove the channels to log

###### Headers
```
X-Api-Key: yourcoolapikey
```

###### Body
```
{
    "channels": ["77829817"]
}
```
