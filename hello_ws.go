package main

import (
    "code.google.com/p/go.net/websocket"
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "unicode/utf8"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}

func echo(ws *websocket.Conn) {
    uri := ws.Request().RequestURI
    fmt.Println("Connected!", uri)

    for {
        var data []byte
        //var reply string
        if err := websocket.Message.Receive(ws, &data); err != nil {
            fmt.Println("Can't receive!")
            break
        }
        fmt.Println("Test:", data)

        reply := string(data)
        fmt.Println("Received: " + reply)
        fmt.Println("Test:", len(reply), utf8.RuneCountInString(reply), []byte(reply))

        msg := "ReEcho: " + reply
        fmt.Println("Reply:", msg)
        if err := websocket.Message.Send(ws, data); err != nil {
            fmt.Println("Can't send")
            break
        }
    }
    fmt.Println("Closed!")
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:8080/ws";

        window.onload = function() {
            sock = new WebSocket(wsuri);

            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }

            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }

            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
                console.log(e.data.type, e.data.size, e.data.slice());
            }
        };

        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
        };
    </script>
    <h1>WebSocket Echo Test</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="Hello, world!">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>`)
}

func main() {
    http.Handle("/", http.HandlerFunc(index))
    http.Handle("/hello", http.HandlerFunc(hello))
    http.Handle("/ws", websocket.Handler(echo))

    //http.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
    //http.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
    //http.Handle("/debug/pprof/heap", pprof.Handler("heap"))
    //http.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error!")
    }
}
