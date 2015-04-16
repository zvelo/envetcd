package envetcd

import (
	"runtime"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRoute(t *testing.T) {
	Convey("Getting default gateway", t, func() {
		switch runtime.GOOS {
		case "linux":
			ip, err := getDefaultRouteGateway()
			So(ip, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		default:
			ip, err := getDefaultRouteGateway()
			So(ip, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "not attempting to determine default gateway on non-linux OS")
		}
	})
}

func TestGetRouteCorrectly(t *testing.T) {
	Convey("Given a sample route file", t, func() {
		input := strings.NewReader(
			`Iface	Destination	Gateway 	Flags	RefCnt	Use	Metric	Mask		MTU	Window	IRTT                                                       
wlp6s0	00000000	010C150A	0003	0	0	600	00000000	0	0	0                                                                           
wlp6s0	000C150A	00000000	0001	0	0	0	00FFFFFF	0	0	0                                                                             
wlp6s0	000C150A	00000000	0001	0	0	600	00FFFFFF	0	0	0                                                                           
virbr1	000811AC	00000000	0001	0	0	0	00FFFFFF	0	0	0                                                                             
virbr0	007CA8C0	00000000	0001	0	0	0	00FFFFFF	0	0	0  `,
		)
		ip, err := scanRouteFile(input)
		So(ip, ShouldNotBeEmpty)
		So(err, ShouldBeNil)
		Convey("Gateway IP should be the IP from the file", func() {
			So(ip.String(), ShouldEqual, "10.21.12.1")
		})
	})
}
