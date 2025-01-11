package mongodb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type builder struct {
	host     string
	port     string
	username string
	password string
	tls      bool
	ca       string
	params   map[string]string
	options  *options.ClientOptions
}

func (b *builder) uri() string {
	var params string
	uri := fmt.Sprintf("mongodb://%s:%s", b.host, b.port)
	if len(b.params) > 0 {
		for k, v := range b.params {
			params += fmt.Sprintf("&%s=%s", k, v)
		}
		uri = fmt.Sprintf("%s/?%s", uri, params[1:])
	}
	return uri
}

func (b *builder) setOptionURI() *builder {
	b.options = options.Client().ApplyURI(b.uri())
	return b
}

func (b *builder) setOptionAuth() *builder {
	if b.username != "" && b.password != "" {
		b.options = b.options.SetAuth(options.Credential{
			Username: b.username,
			Password: b.password,
		})
	}
	return b
}

func (b *builder) getTLSConfig() (*tls.Config, error) {
	caCert, err := base64.StdEncoding.DecodeString(b.ca)
	if err != nil {
		return nil, err
	}
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(caCert); !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	return &tls.Config{
		RootCAs: roots,
	}, nil
}

func (b *builder) setOptionTLS() (*builder, error) {
	tlsConfig, err := b.getTLSConfig()
	if err != nil {
		return nil, err
	}
	b.options = b.options.SetTLSConfig(tlsConfig)
	return b, nil
}

/*
NewBuilder returns a builder for a MongoDB connection.

	host: string - MongoDB host
	port: string - MongoDB port
	username: string - MongoDB username
	password: string - MongoDB password

	returns: *builder
*/
func NewBuilder(host, port, username, password string) *builder {
	return &builder{
		host:     host,
		port:     port,
		username: username,
		password: password,
		params:   make(map[string]string),
	}
}

/*
WithTLS sets the TLS configuration for the MongoDB connection.

	ca: string - base64 encoded CA certificate

	returns: *builder
*/
func (b *builder) WithTLS(ca string) *builder {
	b.tls = true
	b.ca = ca
	b.params["tls"] = "true"
	return b
}

/*
WithRetryWrites sets the retryWrites parameter for the MongoDB connection.

	value: string - The value should be a string representation of a boolean value (true or false).

	returns: *builder
*/
func (b *builder) WithRetryWrites(value string) *builder {
	b.params["retryWrites"] = value
	return b
}

/*
GetClient returns a new MongoDB client. The client is not connected to the MongoDB server.

	ctx: context.Context - The context for the client.

	returns: *mongo.Client, error
*/
func (b *builder) GetClient(ctx context.Context) (*mongo.Client, error) {
	b.setOptionURI().setOptionAuth()
	if b.tls {
		_, err := b.setOptionTLS()
		if err != nil {
			return nil, err
		}
	}
	return mongo.Connect(ctx, b.options)
}
