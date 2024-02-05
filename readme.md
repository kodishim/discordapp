# ⚡ Discordapp

A simplistic wrapper for the Discord Application API.

## 🦺 Tests

Tests are located in the test directory. In order for the tests in test/integration to function correctly several environment variables are required. This can be done by creating a .env file with the below variables in the test/integration directory.

### ⚙️ Example .env in the test/integration directory

```
TOKEN=
SECRET=
GUILD=
MEMBER=
ACCESS_TOKEN=
AUTHORIZED_USER=
```

- TOKEN: The bot token of a Discord application.
- SECRET: The client secret of a Discord application.
- GUILD: The guild ID of a Discord guild that the bot with the test token is currently in.
- MEMBER: The user ID of a discord user who is currently in the test guild.
- ACCESS_TOKEN: Access token of an authorized user.
- AUTHORIZED_USER: The ID of the authorized user.
