# nixy

Author: Brian Tomlinson <darthlukan@gmail.com>

## Description

Nixy is yet another IRC bot. At least it will be once I get to writing
it.

## TODO

It should:

1. Open connection to an IRC server.
2. Initialise IRC connection (username, nick, modes).
3. Join configured channels(s).
4. Have a dynamic plugin system, which can load/reload/unload plugins,
   even when the bot was running. This can be done by command.
5. Process channel events, and pass to plugins if necessary.
6. Develop quote plugin, for storing and retrieving channel
   quotes. Stored in a database.
7. Develop posted URLs display and shortening plugin.
8. Develop search plugin based on user queries, and return results.
9. Develop plugin for notes to be left for other users. These notes
   display when the user is next active in the channel.
10. Be available on the Nixheads Discord as well.

## License

GPLv3, see [LICENSE](LICENSE).
