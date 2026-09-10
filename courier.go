package klikresi

// Supported courier codes. Use these with Tracking.Get and the Rates
// courier filter instead of hard-coding strings.
const (
	// CourierSPX is Shopee Express.
	CourierSPX = "spx"
	// CourierJNE is Jalur Nugraha Ekakurir.
	CourierJNE = "jne"
	// CourierJNT is J&T Express.
	CourierJNT = "jnt"
	// CourierSicepat is Sicepat Express.
	CourierSicepat = "sicepat"
	// CourierNinja is Ninja Express.
	CourierNinja = "ninja"
	// CourierPos is POS Indonesia.
	CourierPos = "pos"
	// CourierSAP is SAP Express.
	CourierSAP = "sap"
	// CourierLEX is Lazada Logistics.
	CourierLEX = "lex"
	// CourierLion is Lion Parcel.
	CourierLion = "lion"
	// CourierIDExpress is ID Express.
	CourierIDExpress = "ide"
	// CourierAnteraja is Anteraja.
	CourierAnteraja = "anteraja"
	// CourierWahana is Wahana Prestasi Logistik.
	CourierWahana = "wahana"
	// CourierTiki is TIKI.
	CourierTiki = "tiki"
)

// DeliveryStatus is the normalized shipment status returned by the
// tracking API.
type DeliveryStatus string

// Delivery statuses returned by the tracking API.
const (
	// StatusInfoReceived means the carrier has received the package info
	// and is about to pick up the package.
	StatusInfoReceived DeliveryStatus = "InfoReceived"
	// StatusInTransit means the package is in transit and has a good
	// transportation condition.
	StatusInTransit DeliveryStatus = "InTransit"
	// StatusOutForDelivery means the package has arrived at the local
	// point or is on the way to the recipient.
	StatusOutForDelivery DeliveryStatus = "OutForDelivery"
	// StatusFailedAttempt means the delivery of the package was attempted
	// but failed due to some reasons.
	StatusFailedAttempt DeliveryStatus = "FailedAttempt"
	// StatusDelivered means the package has been delivered.
	StatusDelivered DeliveryStatus = "Delivered"
	// StatusReturnToSender means the package is on the way back to the
	// sender.
	StatusReturnToSender DeliveryStatus = "ReturnToSender"
	// StatusException means the package was lost, damaged, on hold, etc.
	StatusException DeliveryStatus = "Exception"
	// StatusExpired means the last track of the package has not been
	// updated for 30 days.
	StatusExpired DeliveryStatus = "Expired"
	// StatusPending means no information yet as the package is pending to
	// track or the carrier is wrong.
	StatusPending DeliveryStatus = "Pending"
)
