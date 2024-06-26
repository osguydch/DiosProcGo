package main
import (
    "context"
    "fmt"
    "log"
    "net"
    // importing generated stubs
    gen "github.com/osguydch/DiosProcGo/proto/Device/V3"
    "google.golang.org/grpc"
)
// DeviceServerImpl will implement the service defined in protocol buffer definitions
type DeviceServerImpl struct {
    gen.UnimplementedDeviceServer
}

// Open is the implementation of RPC call defined in protocol definitions.
// This will take OpenRequest message and return OpenReply
func (g *DeviceServerImpl) Open(ctx context.Context, request *gen.OpenRequest) (*gen.OpenReply, error) {
    return &gen.OpenReply{
        Message: fmt.Sprintf("hello %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Close( ctx context.Context, request *gen.OpenRequest) (*gen.CloseReply, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("hello %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Get( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("Get %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Update( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("Update %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Add( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("Add %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Del( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("Del %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Action( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("Action %s",request.Name),
    },nil
}

func (g *DeviceServerImpl) Message( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error)
{
    return &gen.OpenReply{
        Message: fmt.Sprintf("Message %s",request.Name),
    },nil
}



func main() {
    // create new gRPC server
    server := grpc.NewServer()
    // register the GreeterServerImpl on the gRPC server
    gen.RegisterDeviceServer(server, &DeviceServerImpl{})
    // start listening on port :8080 for a tcp connection
    if l, err := net.Listen("tcp", ":8080"); err != nil {
        log.Fatal("error in listening on port :8080", err)
    } else {
        // the gRPC server
        if err:=server.Serve(l);err!=nil {
            log.Fatal("unable to start server",err)
        }
    }
}