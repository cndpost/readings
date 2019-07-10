This is a collection of interesting links to different technologies and libraries, as well as their comparisons

<H1>Websocket </H1>
<a href="https://www.educba.com/websocket-vs-socket-io/" >Websocket vs Socket.IO </a>
  This comparison website also has many other comparisons of similar or closely related technologies

 <br>

 Following are language implementations of websocket and socket.IO 

 <ul>

 <li><a href="https://github.com/zaphoyd/websocketpp" >  Websocket in C++ </a>  </li>
 <li><a href="https://github.com/cndpost/websocket">  Websocket in GOLANG </a> </li>
 <li><a href="https://github.com/cndpost/socket.io">  Socket.IO in JavaScript</a> </li>
 <li><a href="github.com/googollee/go-socket.io">  Socket.IO in GOLANG </a>  </li>
 <li><a href="https://github.com/socketio/socket.io-client-cpp" >  Socket.IO in C++ </a> </li>
 <li><a href="https://github.com/socketio/socket.io-client-swift" > Socket.IO in Swift </a> </li>

</ul>

Some histories regarding the websocket:


I initially get involved in websocket in 2019 at Microsoft when I ported the websocket C++ library from Windows to Sony Playstation PS4
(A subset of FreeBSD OS plus extensive audio and GPU library for game graphics ) as part of the Bumbleion project of Microsoft.

Bumblelion is a cross platform port (ported to Android, IOS, Nintendo Switch, Sony Playstation PS4, Windows, XBox, Linux)
of the P2P messaging library for games called <a href="https://github.com/playfab" > PlayFabSDK </a>.  

The reason we used websocket is solely because PlayFabSDK/Bumblelion used Azure's Cognition Servives to provide the Text To Sound
(TTS) and Sound To Text (STT) service. This Azure cloud API used websocket, so we have to include websocket as our dependecies. This
forced us to port the websocket to all the platforms that PlayFabSDK/Bumblerlion will need to go.

The TTS and STT enables one player to do voice or text chatting to other team players, and other players see the chat message in text, even
it is muted.  Or, they can listen to others chatting in sound without looking at the text and without moving attention away from the
current gaming activities. This capability is powerful for future use cases such as drivers in a moving vehicle or a soldier in a 
wearable AR helmet.

The TTS and STT as a service from Azure cloud makes the client code much lighter because the application coder do not need to have any
voice recognition training data at all. This is especially important for a battery powered mobile devices. 

The TTS and STT, when piped / combined with natural language translation services, can be used for real time natural language communciations between players who speaks totally different languages.   

Because TTS and STT have real-time, low-latency needs, and have to be in bidirectional push mode, and the client pulling / server response mode of HTTP REST API is not efficient. So Microsoft made TTS and STT API only available using the websocket protocols.










