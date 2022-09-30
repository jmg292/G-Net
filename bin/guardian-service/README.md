# G-Net Guardian

The Guardian service is responsible for maintaining the security of (almost) all network traffic going to or from a G-Net peer. The Guardian service accomplishes this by inserting a driver into the NDIS filter driver stack that is configured to capture and divert all packets traversing the stack.

These packets are tagged with metadata that irrefutably declares the packets origin.  This metadata is checked against the peer's configuration to provide a user with unprecedented control of and visibility into their network traffic.
