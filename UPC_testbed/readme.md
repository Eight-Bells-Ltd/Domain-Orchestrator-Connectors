# UPC Testbed - Orchestrator and Mitigation Actions

## Overview

This document provides details on the interface protocol, authentication, and format for mitigation actions in the UPC testbed. It also provides an example of how to send a mitigation action to the testbed.

## 1. Protocol Used

The UPC testbed uses a **REST API** over **HTTP** to apply orchestration policies (e.g., mitigation actions) on the testbed.

## 2. Interface Endpoint

The IP address for the API is **10.19.2.1**, and the endpoints follow the format shown below:

- `http://10.19.2.1/dns_rate_limiting`
- `http://10.19.2.1/dns_service_disable`
- `http://10.19.2.1/dns_service_handover`
- `http://10.19.2.1/anycast_blackhole`
- `http://10.19.2.1/dns_firewall_spoofing_detection`

## 3. Authentication Method

The interface uses **JWT (JSON Web Tokens)** for authentication.

## 4. Format of Mitigation Action

The mitigation actions are enforced using **JSON** format.

## 5. Demonstration (Optional)

While optional, a demonstration is encouraged. UPC testbed owners are working on preparing videos to showcase how to send mitigation actions (e.g., rate limiting a specific IP) to their testbed and confirm that the actions are successfully enforced.

---

*Note: Updates will be provided once the demonstration videos are available.*
