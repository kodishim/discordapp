# ⚡ Gocord

A minimalistic package that allows the user to interact with the Discord API using a Discord Application/Bot.

## 🦺 Tests

Tests are located in the test directory. In order for the tests in test/integration to function correctly several environment variables are required. This can be done by creating a .env file with the below variables in the test/integration directory.

### ⚙️ Example .env in the test/integration directory

```
TEST_TOKEN=
TEST_SECRET=
TEST_GUILD=
TEST_MEMBER=
```

- TEST_TOKEN: The bot token of a Discord application.
- TEST_SECRET: The client secret of a Discord application.
- TEST_GUILD: The guild ID of a Discord guild that the bot with the test token is currently in.
- TEST_MEMBER: The user ID of a discord user who is currently in the test guild.
