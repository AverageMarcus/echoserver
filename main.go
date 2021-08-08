package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type echoResponse struct {
	RequestPath string              `json:"requestPath"`
	RequestHost string              `json:"requestHost"`
	Protocol    string              `json:"protocol"`
	Method      string              `json:"method"`
	Headers     map[string][]string `json:"headers"`
	RemoteAddr  string              `json:"remoteAddr"`

	ServerHostname string    `json:"serverHostname"`
	ServerIP       string    `json:"serverIP"`
	DateTime       time.Time `json:"dateTime"`
}

func main() {
	hostname, _ := os.Hostname()
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ipAddress := conn.LocalAddr().String()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		resp := echoResponse{
			RequestPath:    r.URL.String(),
			RequestHost:    r.Host,
			Protocol:       r.Proto,
			Method:         r.Method,
			Headers:        r.Header,
			RemoteAddr:     r.RemoteAddr,
			ServerHostname: hostname,
			ServerIP:       ipAddress,
			DateTime:       time.Now(),
		}

		b, _ := json.Marshal(resp)

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(b)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
