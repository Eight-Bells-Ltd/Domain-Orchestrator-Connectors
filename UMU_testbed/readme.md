# UMU Testbed - Orchestrator and Mitigation Actions

## Overview

This document provides details on the interface protocol, authentication, and format for mitigation actions in the UMU testbed. It also provides an example of how to send a mitigation action to the testbed.

## 1. Protocol Used

The orchestrator uses a **REST API** to apply orchestration policies (e.g., mitigation actions) on the testbed. The interface can be interacted with using `curl`.

### Example:
```bash
curl -H "Cache-Control: no-cache" -H "Content-Type: application/xml" -X POST -d @policy.xml http://orchestrator-url:8002/meservice
```

## 2. Interface Endpoint
Both URIs and IPs can be used as endpoints for the interface.

## 3. Authentication Method
Currently, no authentication method is used (assumed to be in a controlled environment).

## 4. Format of Mitigation Action
The orchestrator uses a custom XML-based format called MSPL (Medium Security Policy Language) for defining orchestration actions, such as mitigation actions. These actions are applied directly to a previously defined infrastructure in a system model.

## 5. Demonstration (Optional)
While optional, a demonstration is encouraged. Testbed owners can provide a brief demonstration showing how a mitigation action (e.g., rate limiting a specific IP) is sent to their testbed and confirmed to be successfully enforced.

Example of Applying a Mitigation Action (Rate Limiting an IP):
Administrator or Other Module: Sends a rate limiting MSPL policy (XML) to the system.

Policy Interpreter Module: Translates the MSPL into a specific device configuration.

Policy Interpreter Module: Sends the new configuration to the Security Orchestrator Module.

Security Orchestrator Module: Queries the System Model for the IP address and credentials required to configure the DNS.

Security Orchestrator Module: Selects the appropriate driver for applying the configuration.

Driver: Provides specific instructions to communicate with the device and apply the configuration. Since bind9 does not provide an API, the driver connects to the remote server via SSH, applies the configuration, and restarts the service with the new rate limit.


