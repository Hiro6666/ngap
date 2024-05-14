package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	UserLocationInformationWAGFPresentNothing int = iota /* No components present */
	UserLocationInformationWAGFPresentGlobalLineID
	UserLocationInformationWAGFPresentHFCNodeID
	UserLocationInformationWAGFPresentChoiceExtensions
)

type UserLocationInformationWAGF struct {
	TWAPID       TWAPID
	IPAddress    TransportLayerAddress
	PortNumber   *PortNumber                                                  `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationTWIFExtIEs `aper:"optional"`
}
