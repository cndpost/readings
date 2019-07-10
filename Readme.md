This is a collection of interesting links to different technologies and libraries, as well as their comparisons

<a href="https://www.educba.com/websocket-vs-socket-io/" >Websocket vs Socket.IO </a>
  This comparison website also has many other comparisons of similar or closely related technologies

 <br>

 Following are language implementations of websocket and socket.IO 

 
<a href="https://github.com/zaphoyd/websocketpp" >  Websocket in C++ </a>
<a href="https://github.com/cndpost/websocket">  Websocket in GOLANG </a>
<a href="https://github.com/cndpost/socket.io">  Socket.IO in JavaScript</a>
<a href="github.com/googollee/go-socket.io">  Socket.IO in GOLANG </a>
<a href="https://github.com/socketio/socket.io-client-cpp" >  Socket.IO in C++ </a>
<a href="https://github.com/socketio/socket.io-client-swift" > Socket.IO in Swift </a>

Some histories regarding the websocket:


I initially get involved in websocket in 2019 at Microsoft when I ported the websocket C++ library from Windows to Sony Playstation PS4
(A subset of FreeBSD OS plus extensive audio and GPU library for game graphics ) as part of the Bumbleion project of Microsoft.

Bumblelion is a cross platform port (ported to Android, IOS, Nintendo Switch, Sony Playstation PS4, Windows, XBox, Linux)
of the P2P messaging library for games called <a href="https://github.com/playfab" > PlayFabSDK </a>.  

The reason we used websocket is solely because PlatFabSDK/Bumblelion used Azure's Cognition Servives to provide the Text To Sound
(TTS) and Sound To Text (STT) service. This Azure cloud API used websocket, so we have to include websocket as our dependecies. 

The TTS and STT enables one player to do voice or text chatting to other team players, and other players see the chat message in text, even
it is muted.  Or, they can listen to others chatting in sound without looking at the screen and without moving attention away from the
current gaming activities. 

The TTS and STT as a service from Azure cloud makes the client code much lighter because the application coder do not need to have any
voice recognition training data at all.

The TTS and STT, when piped / combined with natural language translation services, can be used for real time natural language communciations between players who speaks totally different languages.   

Because TTS and STT have real-time, low-;atency needs, and have to be in bidirectional push mode, and the client pulling / server response mode of HTTP REST API is not efficient. So Microsoft made TTS and STT API only available using the websocket protocols.










