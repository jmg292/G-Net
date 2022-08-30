# A Peer-to-Peer Distributed, Decentralized Zero Trust Toolkit

Once upon a time I set out to build a VR-native, Windows-based, bare metal malware analysis lab capable of:

1. Dynamically generating a force-directed, 3D Virtual Reality node map of every device on my network
2. Viewing a device's IP address and hostname by highlighting its corresponding node
3. Authenticating with a device and negotiating a secure connection by clicking on its corresponding node
4. Executing malicious binaries on a remote, bare-metal Windows box
5. Capturing any and all network traffic going to and from that device for later analysis
6. Capturing any and all unauthorized changes to that device for later analysis
7. Viewing the captured information in 3D Virtual Reality
8. Performing static binary analysis in 3D virtual reality

Since then, I've spent a lot of time conducting research on everything I'd need to build in order to realize my cyberpunk dreams.

If you'd like a glipse into why I started this project and where it's heading, check out the series "[The Many Whys of G-Net](https://www.gnzlabs.io/gnzlabs-blog/many-whys-g-net/)" on my blog! I have a new post planned to coinciding with each release over the next few weeks.

## Research Foundation

I learned a _ton_ while I was working on G-Net's design, specifically in the areas of:

* Windows Security
* Network Security
* Infrastructure Security
* Peer-to-peer networking
* Virtual private networking
* Mobile Device Management
* Cryptographic coprocessors and security tokens

If you'd like to learn more about G-Net's research foundation, check out the [Foundational Research, Inspiration, and Credits](https://github.com/jmg292/G-Net/wiki/Foundational-Research,-Inspiration,-and-Credits) page on the wiki!

## Design Overview

G-Net's design covers the G-Net framework itself, as well as a number of tools designed to leverage that toolkit. Every G-Net component is designed to be:

* Distributed & Decentralized
* Fully auditable
* Easy to Use
* Low maintenance
* Open & Compatible

An overview of each of these components is available in this section.

### A Zero Trust Security Framework

I didn't actually set out to get involved in the "Zero Trust" game at all. G-Net's security framework was designed as a means of:

* Auditing network traffic to and from my malware analysis lab
* Restricting the destination of network traffic eminating from a specific origin
* Irrefutably guaranteeing the origin of network traffic using a devices' TPM
* Authorizing network sessions using a devices' TPM

Searching for related work led me to [Google's BeyondCorp](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/43231.pdf) implementation, which ultimately led me to [NIST.SP.800-207 - Zero Trust](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-207.pdf). I was pleasantly surprised to find an industry standard solution to my specific problems and decided to adopt ZTNA tenets as a part of this project.

### Identity & Access Management

I thoroughly evaluated every IDP, PDP, and VPN available to me during the research phase of this project. During this evaluation I was unable to find simple way to identity and authenticate devices using a TPM.  All of the products I evaluated:

1. Either required a _ton_ of additional PKI to enable certificate-based authentication (see: [Okta Device Trust](https://help.okta.com/en-us/Content/Topics/Mobile/Okta_Mobile_Device_Trust_Windows-desktop.htm)), or
2. Offered no support for the TPM at all

[StrongSwan](https://www.strongswan.org/) is a notable exception.  It's an open source VPN solution that can be configured to use TPMs and smart cards for authentication, but is unable to support the introspection and discovery aspects of my project without significant modification.  The level of effort required to modify StrongSwan was about the same as the level of effort required to build an IDP.

I decided to build my own identity engine - one with a bit more flexibility.

#### Network Blockchain

At its heart, the G-Net uses a Proof of Authority blockchain (called a [Tracery](https://github.com/jmg292/G-Net/wiki/Traceries:-The-Network-Configuration-Blockchain)) as an identity database.  Other components use the Tracery for domain configuration, device management, authorization, and access control.

The Tracery is a core component of G-Net.  Incorporating the Tracery into the design means G-Net can store full copy of this tamper-proof database is stored on every node, which allows a node to act as its own IDP, PDP, and Certificate Authority.

This is only possible thanks to a PoA blockchain's natural tamper-resistance. Modifying a locally stored copy of the Tracery without delegated authority from the network's Founder will cause validation checks to _fail_. G-Net's core network services will cease to function if an invalid Tracery is detected, preventing the node from communicating with the rest of the network (unless the unauthorized modification can be reverted).

The Tracery's fail-secure tamper detection allowed me to design some pretty cool functionality, including: Device Management, Network Configuration, Directory Services and a full-blown Windows Active Directory Domain Controller. To reiterate: 

> Each node is capable of serving as its own fully privileged, centrally managed, federated Active Directory Domain Controller. 

And:

> Even in this configuration, a node is unable to make changes to the domain domain without authorization from a central Network Authority.

That opens up some pretty neat policy options for airgapped devices.

#### Distributed Identity Platform

G-Net's dIDP is designed to provide the identity verification and authentication functionality required to adopt the "Tenets of Zero Trust" described in [NIST.SP.800-207](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-207.pdf). It mandates the use of security tokens (such as the Yubikey) or commodity HSMs (such as the TPM) for:

* [[NIST.FIPS.201-3]](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.201-3.pdf) - Secure, reliable identification of users and devices
* Device identity verification and authentication
* PIV-compatible X509 certificate storage and management
* Certificate Authority operations related to PIV-compatible X509 certificate storage and management

An network administrator can use the dIDP to:

* Build and manage a directory of authorized users and device identities (Directory services)
* Provide cryptographically verifiable guarantees about managed identities (Attestation, Non-repudiation)
* Assign and manage permissions associated with managed identities (Access management)

The dIDP provides a message bus interface to other components.  This allows G-Net to integrate its identity database with existing services to support standard protocols (such as LDAP and OIDC).

### Virtual Private Networking

G-Net's VPN exists as a byproduct of the network and routing services available to and offered by participants (or nodes) in G-Net's peer-to-peer network. Nodes leverage NAT traversal capabilities and packet relays to communicate directly without knowledge of the underlying physical infrastructure. G-Net uses these logical data-links to organize nodes into a loosely defined mesh topology.

If you'd like to learn more about G-Net's VPN design, see: [Networking: Topology, Routing, and Packet Delivery](https://github.com/jmg292/G-Net/wiki/Networking:-Topology,-Routing,-and-Packet-Delivery)

### Additional Components

The project is under active development, with a few additional G-Net components still stuck in the design phase.

The project roadmap, planned development, and overall progress are available within [G-Net's Project Page](https://github.com/users/jmg292/projects/1)
