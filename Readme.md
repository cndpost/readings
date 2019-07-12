This is a collection of interesting links to different technologies and libraries, as well as their comparisons

<H1>Websocket </H1>
<a href="https://www.educba.com/websocket-vs-socket-io/" >Websocket vs Socket.IO </a>
  This comparison website also has many other comparisons of similar or closely related technologies

 <br>
 <br>

 Following are language implementations of websocket and socket.IO 

 <ul>

 <li><a href="https://github.com/zaphoyd/websocketpp" >  Websocket in C++ </a>  </li>
 <li><a href="https://github.com/cndpost/websocket">  Websocket in GOLANG </a> </li>
 <li><a href="https://github.com/cndpost/socket.io">  Socket.IO in JavaScript</a> </li>
 <li><a href="https://github.com/cndpost/go-socket.io">  Socket.IO in GOLANG </a>  </li>
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


<H1> Implementing a chat server in GOLANG </H1>
<H2> -- A case argument of using Socket.IO vs using Rest API </H2>

<H3>Use case scenario: </H3>

 A messaging app that any subscribers can broadcast infrequent messagings to all subscribers.
 The message needs to be relayed without too long a delay that feels like real-time communication. 
 The subscriber numbers can be between thousands to hundreds of thousands of different clients.
 
<H3>If use REST API </H3>

  In order to attain real-time chat, each subscriber would have to poll the server once every second for 
  any new messages, then we will have 60 REST API calls from each subscriber. If we have a total of 100,000 subscribers, 
  then we will have a total of 6,000,000 REST API calls per minute.

<H3> If use Socket.IO </H3>

   Each client would maintain one solitary connection to the serve. 100,000 subscribers will have a 100,000 simultaneous 
   connections. But the subscribers would not need to poll for new messages. If someone posts a new message, only then would our 
   server push out an update to our 100,000 clients. 

<h3> Conclusion </H3>

   For above use cases, the Socket.IO approach is more scalable than the REST API approach.

<h2>
Chat Server Implementation using GOLANG version of socket.IO 
</h2>
<br>

The live demo of this website is at <a href="webportfolio.n2k.net:5000"> Chat-App in GOLANG using go-socket.io </a>
The full source code is at the /chat-app subfolder of this current repo you are reading now.



  Preparation of the code for the site:

```
      go get github.com/cndpost/go-socket.io
      go get github.com/cndpost/go-engine.io
```
 <br>
   

  We actually just use the v1.0 version of the two packages in their v1.0 folder. 
  
  To use in app:

      import "github.com/cndpost/go-socket.io/v1.0"

  Chat server implementation (save the code in file main.go ):
  
  ```
		package main

		import (
			"log"
			"net/http"

			socketio "github.com/cndpost/go-socket.io/v1.0"
		)

		func main() {

			server, err := socketio.NewServer(nil)
			if err != nil {
				log.Fatal(err)
			}

			server.On("connection", func(so socketio.Socket) {

				log.Println("on connection")

				so.Join("chat")

				so.On("chat message", func(msg string) {
					log.Println("emit:", so.Emit("chat message", msg))
					so.BroadcastTo("chat", "chat message", msg)
				})

				so.On("disconnection", func() {
					log.Println("on disconnect")
				})
			})

			server.On("error", func(so socketio.Socket, err error) {
				log.Println("error:", err)
			})

			http.Handle("/socket.io/", server)

			fs := http.FileServer(http.Dir("static"))
			http.Handle("/", fs)

			log.Println("Serving at localhost:5000...")
			log.Fatal(http.ListenAndServe(":5000", nil))
		}

   ```

   To run above server, just type following command in the same folder as above code:
   
             go run main.go
			  

   Chat UI page in javascript:
   
      Following HTML and javascript code implemented the UI part of the chat-app. 
	  
   
```
<!doctype html>
<html>
  <head>
    <title>Socket.IO chat</title>
    <style>
      * { margin: 0; padding: 0; box-sizing: border-box; }
      body { font: 13px Helvetica, Arial; }
      form { background: #000; padding: 3px; position: fixed; bottom: 0; width: 100%; }
      form input { border: 0; padding: 10px; width: 90%; margin-right: .5%; }
      form button { width: 9%; background: rgb(130, 224, 255); border: none; padding: 10px; }
      #messages { list-style-type: none; margin: 0; padding: 0; }
      #messages li { padding: 5px 10px; }
      #messages li:nth-child(odd) { background: #eee; }
    </style>
  </head>
  <body>
    <ul id="messages"></ul>
    <form action="">
      <input id="m" autocomplete="off" /><button>Send</button>
    </form>
    <script src="/socket.io.js"></script>
    <script src="/jquery.js"></script>
    <script>
      var socket = io();
      $('form').submit(function(){
        socket.emit('chat message', $('#m').val());
        $('#m').val('');
        return false;
      });
      socket.on('chat message', function(msg){
        $('#messages').append($('<li>').text(msg));
      });
    </script>
  </body>
</html>
   
```

<br>

The live demo of this website is at <a href="webportfolio.n2k.net:5000"> Chat-App in GOLANG using go-socket.io </a>
The full source code is at the /chat-app subfolder of this current repo you are reading now.

<br>

<H2> Acknowledgement </H2>

This article was originally from <a href="https://tutorialedge.net/golang/golang-websockets-tutorial/" >Here </a> but the code in the original article is not working due to its code dependencies has changed. I have corrected the code and made it working with the version of dependencies that I 
maintained together with this article.

The corrections I have made are as follows: 

<ol>
<li> Moved the branch v1.0 of the code https://github.com/googollee/go-socket.io to https://github.com/cndpost/go-socket.io/v1.0/go-socket.io  </li>
<li> Moved the branch v1.0 of the code https://github.com/googollee/go-engine.io to https://github.com/cndpost/go-engine/v1.0/go-socket.io </li>
<li>Use the index.html examples from https://github.com/douglasmakey/go-socket.io/tree/master/example/asset/index.html,  
   I have modified the index.html to use the current jquery.js and socket.io.js in my folder of .asset, but you can use
   the original index.html as well but then you need to copy its referenced dependencies of particular versions of
   jquery and socket.io
</li>
</ol>

The reason we have to fork from the upstream repo https://github.com/googollee/go-socket.io is that the upstream repo has v1.0 in its branch v1.0, but the "go get " commands does not recognize the branch syntax and can only get the main branch.
And all the later branches have incompatible APIs and cause the tutorial code to be broken.
