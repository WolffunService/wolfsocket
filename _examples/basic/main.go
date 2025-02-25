package main

import (
	"encoding/json"
)

/*
	Read the README.md
*/

const (
	addr      = "localhost:8081"
	endpoint  = "/echo"
	namespace = "default"
	// false if client sends a join request.
	serverJoinRoom = false
	// if the above is true then this field should be filled, it's the room name that server force-joins a namespace connection.
	serverRoomName = "room1"
)

func main() {

}

// userMessage implements the `MessageBodyUnmarshaler` and `MessageBodyMarshaler`.
type userMessage struct {
	From string `json:"from"`
	Text string `json:"text"`
}

// Defaults to `DefaultUnmarshaler & DefaultMarshaler` that are calling the json.Unmarshal & json.Marshal respectfully
// if the instance's Marshal and Unmarshal methods are missing.
func (u *userMessage) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

func (u *userMessage) Unmarshal(b []byte) error {
	return json.Unmarshal(b, u)
}

//
//var serverAndClientEvents = wolfsocket.Namespaces{
//	namespace: wolfsocket.Events{
//		wolfsocket.OnNamespaceConnected: func(c *wolfsocket.NSConn, msg wolfsocket.Message) error {
//			log.Printf("[%s] connected to namespace [%s].", c, msg.Namespace)
//
//			if !c.Conn.IsClient() && serverJoinRoom {
//				c.JoinRoom(nil, serverRoomName)
//			}
//
//			return nil
//		},
//		wolfsocket.OnNamespaceDisconnect: func(c *wolfsocket.NSConn, msg wolfsocket.Message) error {
//			log.Printf("[%s] disconnected from namespace [%s].", c, msg.Namespace)
//			return nil
//		},
//		wolfsocket.OnRoomJoined: func(c *wolfsocket.NSConn, msg wolfsocket.Message) error {
//			text := fmt.Sprintf("[%s] joined to room [%s].", c, msg.Room)
//			log.Println(text)
//
//			// notify others.
//			if !c.Conn.IsClient() {
//				c.Conn.Server().Broadcast(c, wolfsocket.Message{
//					Namespace: msg.Namespace,
//					Room:      msg.Room,
//					Event:     "notify",
//					Body:      []byte(text),
//				})
//			}
//
//			return nil
//		},
//		wolfsocket.OnRoomLeft: func(c *wolfsocket.NSConn, msg wolfsocket.Message) error {
//			text := fmt.Sprintf("[%s] left from room [%s].", c, msg.Room)
//			log.Println(text)
//
//			// notify others.
//			if !c.Conn.IsClient() {
//				c.Conn.Server().Broadcast(c, wolfsocket.Message{
//					Namespace: msg.Namespace,
//					Room:      msg.Room,
//					Event:     "notify",
//					Body:      []byte(text),
//				})
//			}
//
//			return nil
//		},
//		"chat": func(c *wolfsocket.NSConn, msg wolfsocket.Message) error {
//			if !c.Conn.IsClient() {
//				c.Conn.Server().Broadcast(c, msg)
//			} else {
//				var userMsg userMessage
//				err := msg.Unmarshal(&userMsg)
//				if err != nil {
//					log.Fatal(err)
//				}
//				fmt.Printf("%s >> [%s] says: %s\n", msg.Room, userMsg.From, userMsg.Text)
//			}
//			return nil
//		},
//		// client-side only event to catch any server messages comes from the custom "notify" event.
//		"notify": func(c *wolfsocket.NSConn, msg wolfsocket.Message) error {
//			if !c.Conn.IsClient() {
//				return nil
//			}
//
//			fmt.Println(string(msg.Body))
//			return nil
//		},
//	},
//}
//
//func main() {
//	wolfsocket.EnableDebug(nil)
//
//	args := os.Args[1:]
//	if len(args) == 0 {
//		log.Fatalf("expected program to start with 'server' or 'client' argument")
//	}
//	side := args[0]
//
//	switch side {
//	case "server":
//		startServer()
//	case "client":
//		startClient()
//	default:
//		log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
//	}
//}
//
//func startServer() {
//	server := wolfsocket.New(gobwas.DefaultUpgrader, serverAndClientEvents)
//	server.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {
//		if userID := r.Header.Get("X-Username"); userID != "" {
//			return userID
//		}
//
//		return wolfsocket.DefaultIDGenerator(w, r)
//	}
//	redisConf := redis.Config{
//		Addr:     "localhost:6379",
//		Password: "",
//	}
//	redisExc, _ := redis.NewStackExchange(redisConf,
//		"thetan-websocket",
//	)
//
//	server.UseStackExchange(redisExc)
//	server.OnUpgradeError = func(err error) {
//		log.Printf("ERROR: %v", err)
//	}
//	server.OnConnect = func(c *wolfsocket.Conn) error {
//
//		if c.WasReconnected() {
//			log.Printf("[%s] connection is a result of a client-side re-connection, with tries: %d", c.ID(), c.ReconnectTries)
//		}
//
//		log.Printf("[%s] connected to the server.", c)
//
//		// If you want to close the connection immediately
//		// from server's OnConnect event then you should
//		// set the `FireDisconnectAlways` option to true.
//		// ws.FireDisconnectAlways = true:
//		//
//		// return fmt.Errorf("custome rror")
//		// c.Close()
//
//		// if returns non-nil error then it refuses the client to connect to the server.
//		return nil
//	}
//
//	server.OnDisconnect = func(c *wolfsocket.Conn) {
//		log.Printf("[%s] disconnected from the server.", c)
//	}
//
//	log.Printf("Listening on: %s\nPress CTRL/CMD+C to interrupt.", addr)
//	http.Handle("/", http.FileServer(http.Dir("./browser")))
//	http.Handle(endpoint, server)
//	log.Fatal(http.ListenAndServe(addr, nil))
//}
//
//func startClient() {
//	var username string
//	fmt.Print("Please specify a unique name: ")
//	fmt.Scanf("%s", &username)
//
//	// init a gobwas(could use a gorilla one instead) Dialer.
//	options := gobwas.Options{
//		Header: gobwas.Header{
//			"X-Username": []string{username},
//		},
//	}
//	dialer := gobwas.Dialer(options)
//
//	// init the websocket connection by dialing the server.
//	client, err := wolfsocket.Dial(
//		// Optional context cancelation and deadline for dialing.
//		context.TODO(),
//		// The underline dialer, can be also a gobwas.Dialer/DefautlDialer or a gorilla.Dialer/DefaultDialer.
//		// Here we wrap a custom gobwas dialer in order to send the username among, on the handshake state,
//		// see `startServer().server.IDGenerator`.
//		dialer,
//		// The endpoint, i.e ws://localhost:8080/path.
//		addr+endpoint,
//		// The namespaces and events, can be optionally shared with the server's.
//		serverAndClientEvents,
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer client.Close()
//
//	go func() {
//		<-client.NotifyClose
//		os.Exit(0)
//	}()
//
//	// connect to the "default" namespace.
//	c, err := client.Connect(nil, namespace)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var room *wolfsocket.Room
//
//	scanner := bufio.NewScanner(os.Stdin)
//
//askRoom:
//	if !serverJoinRoom {
//		fmt.Print("Please specify a room to join, i.e room1: ")
//		if !scanner.Scan() {
//			log.Fatal(scanner.Err())
//		}
//		roomToJoin := scanner.Text()
//
//		room, err = c.JoinRoom(nil, roomToJoin)
//		if err != nil {
//			log.Fatal(err)
//		}
//	} else {
//		room = c.Room(serverRoomName)
//		if room == nil {
//			log.Fatalf("room %s is nil", serverRoomName)
//		}
//	}
//
//	fmt.Fprint(os.Stdout, ">> ")
//
//	for {
//		if !scanner.Scan() {
//			log.Printf("ERROR: %v", scanner.Err())
//			break
//		}
//
//		text := scanner.Text()
//
//		if text == "exit" {
//			break
//		}
//
//		if text == "leave" {
//			room.Leave(nil)
//			if !serverJoinRoom {
//				goto askRoom
//			}
//		}
//
//		// username is the connection's ID ==
//		// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
//		// which generated by server-side via `Server#IDGenerator`.
//		userMsg := userMessage{From: username, Text: text}
//		ok := room.Emit("chat", wolfsocket.Marshal(userMsg))
//		if !ok {
//			log.Fatal("Emit failed")
//		}
//		fmt.Fprint(os.Stdout, ">> ")
//	}
//}
