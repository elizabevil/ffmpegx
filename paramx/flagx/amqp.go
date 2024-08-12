package flagx

type DeliveryMode = Flag

const (
	//Delivery mode set to "persistent" (2).This is the default value.Messages may be written to the brokers disk depending on its setup.

	DeliveryMode_persistent DeliveryMode = "persistent"
	//Delivery mode set to "non-persistent" (1).Messages will stay in brokers memory unless the broker is under memory pressure.

	DeliveryMode_non_persistent DeliveryMode = "non-persistent"
)
