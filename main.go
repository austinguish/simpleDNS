package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	fmt.Println("Received DNS Request")

	for _, question := range r.Question {
		fmt.Printf("Question: %s %s\n", question.Name, dns.TypeToString[question.Qtype])
	}

	// Forward the request to Google's public DNS (8.8.8.8)
	c := new(dns.Client)
	resp, _, err := c.Exchange(r, "8.8.8.8:53")
	if err != nil {
		fmt.Printf("Error forwarding DNS request: %v\n", err)
		return
	}

	// Send the response back to the client
	w.WriteMsg(resp)
}

func main() {
	// Start a DNS server on port 53
	server := &dns.Server{Addr: ":53", Net: "udp"}

	// Set the handler function for incoming requests
	dns.HandleFunc(".", handleRequest)

	fmt.Println("DNS Proxy Server listening on :53")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting DNS server: %v\n", err)
	}
}
