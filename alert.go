package alert

import (
	"context"
	"errors"
	pub "github.com/edunx/rock-public-go"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"time"
)

func (self *Alert) Start() error {

	if self.C.Url == "null" {
		return errors.New("got nil url")
	}

	//重置DNS功能
	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: 5000 * time.Millisecond,
				}
				return d.DialContext(ctx, "udp", self.C.Resolver )
			},
		},
	}

	self.Client = resty.New()
	self.Client.GetClient().Transport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.DialContext(ctx, network, addr)
	}

	return nil
}

func (self *Alert) Do( severity string , alertType string ,
	alertObject string,  alertAttribute string , subject string , body string , tags string) {

	r , e := self.Client.R().
		SetBody(map[string]string{
			"origin_name": self.C.Origin,
			"alert_type": alertType,
			"alert_object": alertObject,
			"alert_attribute": alertAttribute,
			"severity": severity,
			"tags": tags,
			"notifier": self.C.Notifier,
			"subject": subject,
			"body": body,

	}).Post(self.C.Url)
	if e != nil {
		pub.Out.Err("send alert fail , err: %v" , e)
		return
	}

	if r.StatusCode() != 200 {
		pub.Out.Err("send alert fail , body: %s" , r.Body())
		return
	}
	pub.Out.Err("send alert ok, body: %s" , r.Body())
}
