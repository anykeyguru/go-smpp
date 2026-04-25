# SMPP 3.4

[![Go Reference](https://pkg.go.dev/badge/github.com/fiorix/go-smpp/v2.svg)](https://pkg.go.dev/github.com/fiorix/go-smpp/v2) [![Go Report Card](https://goreportcard.com/badge/github.com/fiorix/go-smpp)](https://goreportcard.com/report/github.com/fiorix/go-smpp) [![CI](https://github.com/fiorix/go-smpp/actions/workflows/ci.yml/badge.svg)](https://github.com/fiorix/go-smpp/actions/workflows/ci.yml)

This is an implementation of SMPP 3.4 for Go, based on the original
[smpp34](https://github.com/CodeMonkeyKevin/smpp34) from Kevin Patel.

The API has been refactored to idiomatic Go code with more tests
and documentation. There are also quite a few new features, such
as a test server (see smpptest package) and text codecs for GSM
7-bit (packed and unpacked), Latin-1, ISO-8859-5 and UCS-2.

It is not fully compliant, there are some TODOs in the code.

## Usage

Following is an SMPP client transmitter wrapped by an HTTP server
that can send Short Messages (SMS):

```go
func main() {
    // make persistent connection
    tx := &smpp.Transmitter{
        Addr:   "localhost:2775",
        User:   "foobar",
        Passwd: "secret",
    }
    conn := tx.Bind()
    // check initial connection status
    var status smpp.ConnStatus
    if status = <-conn; status.Error() != nil {
        log.Fatalln("Unable to connect, aborting:", status.Error())
    }
    log.Println("Connection completed, status:", status.Status().String())
    // example of connection checker goroutine
    go func() {
        for c := range conn {
            log.Println("SMPP connection status:", c.Status())
        }
    }()
    // example of sender handler func
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        sm, err := tx.Submit(&smpp.ShortMessage{
            Src:      r.FormValue("src"),
            Dst:      r.FormValue("dst"),
            Text:     pdutext.Raw(r.FormValue("text")),
            Register: pdufield.NoDeliveryReceipt,
            TLVFields: pdutlv.Fields{
                pdutlv.TagReceiptedMessageID: pdutlv.CString(r.FormValue("msgId")),
            },
        })
        if err == smpp.ErrNotConnected {
            http.Error(w, "Oops.", http.StatusServiceUnavailable)
            return
        }
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        io.WriteString(w, sm.RespID())
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

You can test from the command line:

```bash
curl localhost:8080 -X GET -F src=bart -F dst=lisa -F text=hello
```

If you don't have an SMPP server to test, check out
[Selenium SMPPSim](http://www.seleniumsoftware.com/downloads.html).
It has been used for the development of this package.

## Tools

The `cmd/sms` directory provides a command line client for sending
short messages and querying message status against an SMSC.

For an HTTP API server built on top of this library, see
[sms-api-server](https://github.com/fiorix/sms-api-server).

## Supported PDUs

- [x] bind_transmitter
- [x] bind_transmitter_resp
- [x] bind_receiver
- [x] bind_receiver_resp
- [x] bind_transceiver
- [x] bind_transceiver_resp
- [ ] outbind
- [x] unbind
- [x] unbind_resp
- [x] submit_sm
- [x] submit_sm_resp
- [ ] submit_sm_multi
- [ ] submit_sm_multi_resp
- [x] data_sm
- [x] data_sm_resp
- [x] deliver_sm
- [x] deliver_sm_resp
- [x] query_sm
- [x] query_sm_resp
- [ ] cancel_sm
- [ ] cancel_sm_resp
- [ ] replace_sm
- [ ] replace_sm_resp
- [x] enquire_link
- [x] enquire_link_resp
- [ ] alert_notification
- [x] generic_nack
- [x] tag-length-value (TLV)

## Copyright

See the LICENSE file for details. Contributors are listed in the
git history; run `git shortlog -sn` for the current list.
