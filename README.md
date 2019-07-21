[![Go Report Card](https://goreportcard.com/badge/github.com/tritonmedia/discord-bot)](https://goreportcard.com/report/github.com/tritonmedia/discord-bot)
[![codebeat badge](https://codebeat.co/badges/d9f69da5-cfa9-4b3b-bfb7-c15a4398b5cd)](https://codebeat.co/projects/github-com-tritonmedia-discord-bot-master-d046ddd4-43e8-4e90-9717-5578d6d5697d)

# discord-bot

## What is this? 
A fully-functional Discord Bot that is capable of interacting with the Triton Media Platform to invoke, and track media jobs

## Installation / Hacking

### Discord Development Portal

0. Create a sandbox Discord server or use a `test` text channel within an existing server. Change the Discord Server Notification scheme to be only @mentions or this will spam people with notifications.

1. Go to the [Discord dev portal](https://discordapp.com/developers/applications/) and create a new application (e.g. TritonBot)

2. Copy your ClientID

3. Go to the [Discord Permissions Calculator](https://discordapi.com/permissions.html) and give your bot `Read Messages`, `Send Messages`, and `Manage Messages` permissions at a bare minimum (permissions can be managed later after the bot is set up).

4. Use the ClientID from Step #2 to generate a permissions deeplink URL and navigate to it. 

5. Auth into Discord, and you should be provided with a modal popup asking you which servers you'd like to invite the bot to.

6. The DiscordBot should exist within your server as an offline entity

### Local Development environment
1. Install it (setup a GOPATH)

```
git clone git@github.com:tritonmedia/discord-bot discord-bot; cd "discord-bot"
```

2. Copy your Token from the Discord Developer Portal and configure your config.yaml file in `./config` (see `./config/config.example.yaml`)

3. Run it

```bash
# in the discord-bot directory
make && ./bin/discord-bot
```

4. Go to Discord, and you should be able to communicate with the bot by invoking `!triton`.

5. Killing (Ctrl+C) the application will cause the bot to go offline in Discord
