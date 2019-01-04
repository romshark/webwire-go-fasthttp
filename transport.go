package fasthttp

import (
	"crypto/tls"
	"log"
	"net/url"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/qbeon/webwire-go"
	"github.com/valyala/fasthttp"
)

// TLS represents TLS configuration
type TLS struct {
	CertFilePath       string
	PrivateKeyFilePath string
	Config             *tls.Config
}

// Transport implements the webwire transport layer with fasthttp
type Transport struct {
	// Host defines the address of the host
	Host string

	// OnOptions is invoked when the websocket endpoint is examined by the
	// client using the HTTP OPTION method.
	OnOptions func(*fasthttp.RequestCtx)

	// BeforeUpgrade is invoked right before the upgrade of the connection of an
	// incoming HTTP request to a WebSocket connection and can be used to
	// intercept, configure or prevent incoming connections. BeforeUpgrade must
	// return the connection options to be applied or set options.Connection to
	// webwire.Refuse to refuse the incoming connection
	BeforeUpgrade func(ctx *fasthttp.RequestCtx) webwire.ConnectionOptions

	// WarnLog defines the warn logging output target
	WarnLog *log.Logger

	// ErrorLog defines the error logging output target
	ErrorLog *log.Logger

	// KeepAlive enables the keep-alive option if set to a duration above -1.
	// KeepAlive is automatically set to 30 seconds when it's set to 0
	KeepAlive time.Duration

	// Upgrader specifies the websocket connection upgrader
	Upgrader *websocket.FastHTTPUpgrader

	// HTTPServer specifies the FastHTTP server
	HTTPServer *fasthttp.Server

	// TLS enables TLS encryption if specified
	TLS *TLS

	listener        *tcpKeepAliveListener
	addr            url.URL
	readTimeout     time.Duration
	isShuttingdown  webwire.IsShuttingDown
	onNewConnection webwire.OnNewConnection
}

// Address returns the URL address the server is listening on
func (srv *Transport) Address() url.URL {
	return srv.addr
}
