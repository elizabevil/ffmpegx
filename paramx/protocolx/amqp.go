package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 24.1 amqp
type Amqp struct {
	Exchange string `json:"exchange" glag:"-exchange"`
	//Sets the exchange to use on the broker.RabbitMQ has several predefined exchanges: "amq.direct" is the default exchange, where the publisher and subscriber must have a matching routing_key;
	//"amq.fanout" is the same as a broadcast operation (i.e.the data is forwarded to all queues on the fanout exchange independent of the routing_key);
	//and "amq.topic" is similar to "amq.direct", but allows for more complex pattern matching (refer to the RabbitMQ documentation).

	RoutingKey string `json:"routing_key" glag:"-routing_key"`
	//Sets the routing key.The default value is "amqp".The routing key is used on the "amq.direct" and "amq.topic" exchanges to decide whether packets are written to the queue of a subscriber.

	PktSize typex.Size `json:"pkt_size" glag:"-pkt_size"`
	//Maximum size of each packet sent/received to the broker.Default is 131072. Minimum is 4096 and max is any large value (representable by an int).When receiving packets, this sets an internal buffer size in FFmpeg.It should be equal to or greater than the size of the published packets to the broker.Otherwise the received message may be truncated causing decoding errors.

	ConnectionTimeout typex.SecondI `json:"connection_timeout" glag:"-connection_timeout"`
	//The timeout in seconds during the initial connection to the broker.The default value is rw_timeout, or 5 seconds if rw_timeout is not set.

	DeliveryMode flagx.DeliveryMode `json:"delivery_mode" glag:"-delivery_mode"`
	//Sets the delivery mode of each message sent to broker.The following values are accepted:

}
