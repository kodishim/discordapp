package application

// Constants
const BaseDiscordAPIURL = "https://discord.com/api"

// Discord Scopes
const (
	// ScopeActivitiesRead allows your app to fetch data from a user's "Now Playing/Recently Played" list — not currently available for apps
	ScopeActivitiesRead = "activities.read"

	// ScopeActivitiesWrite allows your app to update a user's activity - not currently available for apps (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
	ScopeActivitiesWrite = "activities.write"

	// ScopeApplicationsBuildsRead allows your app to read build data for a user's applications
	ScopeApplicationsBuildsRead = "applications.builds.read"

	// ScopeApplicationsBuildsUpload allows your app to upload/update builds for a user's applications - requires Discord approval
	ScopeApplicationsBuildsUpload = "applications.builds.upload"

	// ScopeApplicationsCommands allows your app to add commands to a guild - included by default with the bot scope
	ScopeApplicationsCommands = "applications.commands"

	// ScopeApplicationsCommandsUpdate allows your app to update its commands using a Bearer token - client credentials grant only
	ScopeApplicationsCommandsUpdate = "applications.commands.update"

	// ScopeApplicationsCommandsPermissionsUpdate allows your app to update permissions for its commands in a guild a user has permissions to
	ScopeApplicationsCommandsPermissionsUpdate = "applications.commands.permissions.update"

	// ScopeApplicationsEntitlements allows your app to read entitlements for a user's applications
	ScopeApplicationsEntitlements = "applications.entitlements"

	// ScopeApplicationsStoreUpdate allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications
	ScopeApplicationsStoreUpdate = "applications.store.update"

	// ScopeBot for oauth2 bots, this puts the bot in the user's selected guild by default
	ScopeBot = "bot"

	// ScopeConnections allows /users/@me/connections to return linked third-party accounts
	ScopeConnections = "connections"

	// ScopeDMChannelsRead allows your app to see information about the user's DMs and group DMs - requires Discord approval
	ScopeDMChannelsRead = "dm_channels.read"

	// ScopeEmail enables /users/@me to return an email
	ScopeEmail = "email"

	// ScopeGDMJoin allows your app to join users to a group dm
	ScopeGDMJoin = "gdm.join"

	// ScopeGuilds allows /users/@me/guilds to return basic information about all of a user's guilds
	ScopeGuilds = "guilds"

	// ScopeGuildsJoin allows /guilds/{guild.id}/members/{user.id} to be used for joining users to a guild
	ScopeGuildsJoin = "guilds.join"

	// ScopeGuildsMembersRead allows /users/@me/guilds/{guild.id}/member to return a user's member information in a guild
	ScopeGuildsMembersRead = "guilds.members.read"

	// ScopeIdentify allows /users/@me without email
	ScopeIdentify = "identify"

	// ScopeMessagesRead for local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates)
	ScopeMessagesRead = "messages.read"

	// ScopeRelationshipsRead allows your app to know a user's friends and implicit relationships - requires Discord approval
	ScopeRelationshipsRead = "relationships.read"

	// ScopeRoleConnectionsWrite allows your app to update a user's connection and metadata for the app
	ScopeRoleConnectionsWrite = "role_connections.write"

	// ScopeRPC for local rpc server access, this allows you to control a user's local Discord client - requires Discord approval
	ScopeRPC = "rpc"

	// ScopeRPCActivitiesWrite for local rpc server access, this allows you to update a user's activity - requires Discord approval
	ScopeRPCActivitiesWrite = "rpc.activities.write"

	// ScopeRPCNotificationsRead for local rpc server access, this allows you to receive notifications pushed out to the user - requires Discord approval
	ScopeRPCNotificationsRead = "rpc.notifications.read"

	// ScopeRPCVoiceRead for local rpc server access, this allows you to read a user's voice settings and listen for voice events - requires Discord approval
	ScopeRPCVoiceRead = "rpc.voice.read"

	// ScopeRPCVoiceWrite for local rpc server access, this allows you to update a user's voice settings - requires Discord approval
	ScopeRPCVoiceWrite = "rpc.voice.write"

	// ScopeVoice allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval
	ScopeVoice = "voice"

	// ScopeWebhookIncoming this generates a webhook that is returned in the oauth token response for authorization code grants
	ScopeWebhookIncoming = "webhook.incoming"
)
