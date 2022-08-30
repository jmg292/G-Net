package extensions

import "encoding/asn1"

var (
	OIDYubikeyFirmwareVersion = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 41482, 3, 3}
	OIDYubikeySerialNumber    = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 41482, 3, 7}
	OIDYubikeyPinTouchPolicy  = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 41482, 3, 8}
	OIDYubikeyFormfactor      = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 41482, 3, 9}
)
