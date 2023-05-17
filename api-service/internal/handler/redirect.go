package handler

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

func (h *Handler) getAddr() string {
	return h.storageAddr[0]
}

func (h *Handler) redirect(c *gin.Context) {
	req := c.Request
	proxy, err := url.Parse(h.getAddr())
	if err != nil {
		log.Printf("error in parse addr: %v", err)
		c.String(500, "error")
		return
	}
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host

	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(req)
	if err != nil {
		log.Printf("error in roundtrip: %v", err)
		c.String(500, "error")
		return
	}

	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(c.Writer)
	return
}
