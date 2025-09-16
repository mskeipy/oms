package constants

const (
	WarehouseInbound  = "inbound"
	WarehouseOutbound = "outbound"

	WarehouseOrderStatusDraft     = "draft"
	WarehouseOrderStatusAwaiting  = "awaiting"
	WarehouseOrderStatusReceived  = "received"
	WarehouseOrderStatusCompleted = "completed"
	WarehouseOrderStatusCancelled = "cancelled"
)

var OrderCodePrefix = map[string]string{
	WarehouseInbound:  "IN",
	WarehouseOutbound: "OUT",
}
