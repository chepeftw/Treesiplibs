package treesiplibs

// Possible future improvements
// https://developers.google.com/protocol-buffers/docs/gotutorial
// https://github.com/golang/protobuf
// http://msgpack.org/index.html
// Basically changing JSON for ProtocolBuffers or MsgPack
// Of course I should check first IF its really an improvement in the packet size


import (
    "net"
    // "github.com/op/go-logging"
)


// +++++++++ Constants
const (
    StartType = iota // 0

    TimeoutType
    QueryType // 2

    AggregateType
    AggregateFwdType //4

    AggregateRteType
    HelloType // 6

    HelloTimeoutType
    HelloReplyType // 8

    RouteByTableType
    RouteByGossipType  //10
)


// +++++++++ Packet structure
type Packet struct {
    ID           string     `json:"id,omitempty"`
    Type         int        `json:"tp"`

    Parent       net.IP     `json:"prnt,omitempty"`
    Source       net.IP     `json:"src,omitempty"`
    Destination  net.IP     `json:"dst,omitempty"`
    Gateway      net.IP     `json:"gw,omitempty"`
    Port         int        `json:"prt,omitempty"`

    Timeout      int        `json:"tmo,omitempty"`
    Query        *Query     `json:"qry,omitempty"`
    Aggregate    *Aggregate `json:"agt,omitempty"`

    Timestamp    string     `json:"ts,omitempty"`
    TimeToLive   int        `json:"ttl,omitempty"`
    Hops         int        `json:"hps,omitempty"`
    Level        int        `json:"lvl,omitempty"`
}

type Query struct {
    Function  string    `json:"fct,omitempty"`
    RelaySet  []*net.IP `json:"rSt,omitempty"`
}

type Aggregate struct {
    Outcome      float32 `json:"otc,omitempty"`
    Observations int    `json:"obs,omitempty"`
}


// Function to calculate the RelaySet
// the idea is to always have 3 elements. Seen as a tree, the closer 3 parents.
func calculateRelaySet( newItem net.IP, receivedRelaySet []*net.IP ) []*net.IP {
    slice := []*net.IP{&newItem}

    if len(receivedRelaySet) < 3 {
        return append(slice, receivedRelaySet...)
    }

    return append(slice, receivedRelaySet[:len(receivedRelaySet)-1]...)
}


func AssembleTimeout() Packet {
    payload := Packet{
        Type: TimeoutType,
    }

    return payload
}

func AssembleTimeoutHello(stamp string) Packet {
    payload := Packet{
        Type: HelloTimeoutType,
        Timestamp: stamp,
    }

    return payload
}

func AssembleAggregate(dest net.IP, out float32, obs int, dad net.IP, me net.IP, tmo int, stamp string, port int) Packet {
    aggregate := Aggregate{
            Outcome: out,
            Observations: obs,
        }

    payload := Packet{
        Type: AggregateType,
        Parent: dad,
        Source: me,
        Destination: dest,
        Port: port,
        Timeout: tmo,
        Timestamp: stamp,
        Aggregate: &aggregate,
    }

    return payload
}

func AssembleQuery(payloadIn Packet, dad net.IP, me net.IP) Packet {
	relaySet := []*net.IP{}
	if payloadIn.Type != StartType {
        relaySet = calculateRelaySet(payloadIn.Source, payloadIn.Query.RelaySet)
    }

    query := Query{
            Function: payloadIn.Query.Function,
            RelaySet: relaySet,
        }

    payload := Packet{
        Type: QueryType,
        Parent: dad,
        Source: me,
        Port: payloadIn.Port,
        Timeout: payloadIn.Timeout,
        Level: payloadIn.Level+1,
        Query: &query,
    }

    return payload
}


func AssembleHello(me net.IP, stamp string) Packet {
    payload := Packet{
        Type: HelloType,
        Source: me,
        Timestamp: stamp,
    }

    return payload
}

func AssembleHelloReply(payloadIn Packet, me net.IP) Packet {
    payload := Packet{
        Type: HelloReplyType,
        Source: me,
        Destination: payloadIn.Source,
        Timestamp: payloadIn.Timestamp,
    }

    return payload
}

func AssembleRoute(gw net.IP, payloadIn Packet) Packet {
    payload := payloadIn

    payload.Type = RouteByGossipType
    payload.Gateway = gw
    // payload.TimeToLive = payloadIn.TimeToLive-1
    payload.Hops = payloadIn.Hops+1

    return payload
}