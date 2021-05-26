# Information

This project abuses what I believe to be an unintentional use of the [local RPC](https://discord.com/developers/docs/topics/rpc) on port 6463.\
According to discord, this is not a problem that needs fixing.\
This Repository shows a way that an attacker can get a discord users token, email, phone number, payment information, and even keystrokes without ever reading a file or interacting with memory.

# How it works

This works by opening a websocket connection to `ws://127.0.0.1:6463/?v=1&encoding=json`, normally this connection would be rejected but when supplying the header `Origin: https://discord.com`, it seems to open perfectly fine.\
Then we listen for the `READY` event, when we get that event we write two payloads, `SUBSCRIBE` and `CONNECT`.\
After that we listen for an incoming payload that passes all these checks. `payload.Cmd == "DISPATCH" && payload.Data.Type == "DISPATCH" && payload.Data.PID == 4`

# Why is this important

I belive this is an important exploit because it bypasses any way an antivirus could reasonably detect a token grabber for discord.\
Usually with token grabbers a program will read files in `%appdata%\discord\Local Storage\leveldb` matching a regex for the users token.\
Firstly this can be unreliable, tokens are stored in there when the client closes so if the user had logged out and back in (or had a forced relog due to a VPN being used) then those files will not have an up-to-date token.\
Secondly a token logger that does this could be stopped relatively easily, if the `discord` folder isn't installed in the `%appdata%` folder then you'll be pretty safe for the most part.\
This exploit doesn't require any reading of those files and always gets a valid and current token every time.\
Another reason this is important is because, though all the information this program grathers is avalible in the `leveldb`, the only thing you can realistically get in wide-spread use is a token. You can then gather all the other information with that token later, this exploit gets all that information without ever having to use the token (possibly setting off some sort of security measure).

# Note

Please don't use this tool in any way discord wouldn't allow, this is simply a POC to show what **can** be done. Thanks for reading!