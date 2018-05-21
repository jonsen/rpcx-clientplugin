package clientPlugin

import (
	//"context"
	//"fmt"
	//"github.com/smallnest/rpcx/protocol"
	"net"
	//"reflect"
	//"runtime"
)

type ClientPlugin struct {
	enable  bool
	clients AuthClients
}

func NewClientPlugin(enable bool, ipaddress []string) (p *ClientPlugin) {
	ips := LoadAuthClients(ipaddress)
	return &ClientPlugin{enable: enable, clients: ips}
}

func (p *ClientPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	ipAddr := conn.RemoteAddr().(*net.TCPAddr).IP

	//fmt.Printf("HandleConnAccept conn (%s) %#v\n", ipAddr, len(p.clients))
	if !p.clients.ClientAuthor(ipAddr) {
		return conn, false
	}
	return conn, true
}

/*
func (p *ClientPlugin) Register(name string, rcvr interface{}, metadata string) error {

	fmt.Printf("Register %s: %T\n", name, rcvr)
	return nil
}

func (p *ClientPlugin) RegisterFunction(name string, fn interface{}, metadata string) error {
	fmt.Printf("RegisterFunction %s: %T\n", name, GetFunctionName(fn))
	return nil
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (p *ClientPlugin) PostReadRequest(ctx context.Context, r *protocol.Message, e error) error {

	return nil
}

func (p *ClientPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, err error) error {

	return nil
}
*/
