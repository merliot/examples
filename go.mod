module github.com/merliot/examples

go 1.19

require (
	github.com/CosmWasm/tinyjson v0.9.0
	github.com/go-daq/canbus v0.2.0
	github.com/merliot/merle v0.0.49
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	gobot.io/x/gobot v1.16.0
	tinygo.org/x/drivers v0.22.0
)

require (
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/msteinert/pam v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sigurn/crc8 v0.0.0-20160107002456-e55481d6f45c // indirect
	github.com/sigurn/utils v0.0.0-20190728110027-e1fefb11a144 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
	golang.org/x/text v0.3.6 // indirect
	periph.io/x/periph v3.6.2+incompatible // indirect
)

replace github.com/merliot/merle => /home/merle/merle

replace tinygo.org/x/drivers => /home/merle/tinygo-drivers
