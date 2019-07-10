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


<H1 Implementing a char server in GOLANG />
This article was originally from <a href="" > </a> but the code in the original article is not working due to its
dependencies has changed. I have corrected the code and made it working with the version of dependencies that I 
maintained together with this article.

<br>
<H2>Using Socket.IO vs RestAPI - a case analysis </H2>
Use case scenrio:

 A messaging app that any subscribers can broadcast infrequent messagings to all subscribers.
 The message needs to be relayed without too long a delay that feels like real-time communication. 
 The subscriber numbers can be between thousands to hundreds of thousands of different clients.
 
If use REST API

  In order to attain real-time chat, each subscriber would have to poll the server once every second for 
  any new messages, then we will have 60 REST API calls from each subscriber. If we have a total of 100,000 subscribers, 
  then we will have a total of 6,000,000 REST API calls per minute.

If use Socket.IO

   Each client would maintain one solitary connection to the serve. 100,000 subscribers will have a 100,000 simultaneous 
   connections. But the subscribers would not need to poll for new messages. If someone posts a new message, only then would our 
   server push out an update to our 100,000 clients. 

Conclusion

   For above use cases, the Socket.IO approach is more scalable than the REST API approach.

Chat Server Implementation using GOLANG version of socket.IO 

  Preparation:  

      go get github.com/cndpost/go-socket.io

  Use in app:

      import "github.com/cndpost/go-socket.io"

  Chat server implementation (save the code in file main.go ):
  <PRE>
		package main

		import (
			"log"
			"net/http"

			socketio "github.com/googollee/go-socket.io"
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

   </PRE>

   To run above server, just type following command in the same folder as above code:
   
             go run main.go
			 
			 

   Chat client:
   
      Following HTML code will be a chat client. Save it in index.html and open it in 
	  a web browser:
	  
	 <PRE> 

        <!DOCTYPE html>
		<html lang="en">
		  <head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<meta http-equiv="X-UA-Compatible" content="ie=edge" />
			<title>Go WebSocket Tutorial</title>
		  </head>
		  <body>
			<h2>Hello World</h2>

			<script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/2.1.1/socket.io.js"></script>
			<script>
			  const socket = io("http://localhost:5000/socket.io/");
			</script>
		  </body>
		</html>


   /PRE>

   

