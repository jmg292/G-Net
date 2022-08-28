# Decentralize Zero Trust

G-Net provides the core identity management, verification, and access control funtionality that defines Zero Trust<sup>[1](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/43231.pdf)</sup> with an approach that strives to be:

* Distributed & Decentralized
* Fully auditable
* Easy to Use
* Low maintenance
* Open & Compatible

It's designed to use this core IAM framework to expose domain management, device administration, and peer-to-peer network access functionality. Its target goal is to reduce the overall barrier to entry into modern enterprise-grade security by reducing the complexity zero trust network architecture and eliminating third party dependencies.

There are a number of core components in various stages of active and planned development, including:

* peer-to-peer distributed device administration tools,
* peer-to-peer distributed network connectivity tools, and
* interfaces between G-Net components and the existing standards in security and technology

The project roadmap, planned development, and overall progress are available within [G-Net's Project Page](https://github.com/users/jmg292/projects/1)


## Distributed Identity Platform

Existing zero trust network solutions rely on centralized identity providers to handle user authentication and access control. This has created a gap between free solutions capable of authenticating a personal user (e.g. `Log In With [GitHub | Google | Facebook]`) and expensive enterprise IAM solutions (e.g. Active Directory, Okta).  There are plenty of open source identity platforms available to fill the gap, but these solutions often fall short of enterprise products when it comes to device identity management and attestation.

G-Net's Distributed Identity Provider is designed to allow a network administrator to:

* Build and manage a directory of authorized users and device identities (Directory services)
* Provide cryptographically verifiable guarantees about managed identities (Attestation, Non-repudiation)
* Assign and manage permissions associated with managed identities (Access management)

To do this while achieving its design goals, G-Net's IDP:

* Leverages widely available, low-cost commodity hardware security modules ([TPM](https://trustedcomputinggroup.org/resource/tpm-library-specification/), [Yubikey](https://www.yubico.com/), etc.) to perform secure, reliable identification of users and devices<sup>[2](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.201-3.pdf)</sup>
* Uses a Proof of Authority blockchain to securely decentralize and distribute its identity database
* Provides a standard interface to integrate its identity database with existing industry standards such as LDAP and OIDC

### References

1. [Google | BeyondCorp - A New Approach to Enterprise security](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/43231.pdf)
1. [NIST FIPS 201-3 | Announcing the Standard for Personal Identity Verification of Federal Employees and Contractors](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.201-3.pdf)
