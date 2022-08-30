package extensions

import "encoding/asn1"

var (
	OIDSecurityServiceConfidentiality = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 7, 1}
	OIDSecurityServiceIntegrity       = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 7, 2}
	OIDSecurityServiceAuthentication  = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 7, 3}
	OIDSecurityServiceNonRepudiation  = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 7, 4}
	OIDSecurityServiceAccessControl   = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 7, 5}
)

var (
	OIDEncrypt = append(OIDSecurityServiceConfidentiality, 1)
)

var (
	OIDAuthenticate         = append(OIDSecurityServiceAuthentication, 1)
	OIDVerifyAuthentication = append(OIDSecurityServiceAuthentication, 2)
)

var (
	OIDProofOfOrigin       = append(OIDSecurityServiceNonRepudiation, 1)
	OIDVerifyProofOfOrigin = append(OIDSecurityServiceNonRepudiation, 2)
)

var (
	OIDApplyAccessControl  = append(OIDSecurityServiceAccessControl, 1)
	OIDVerifyAccessControl = append(OIDSecurityServiceAccessControl, 2)
)
