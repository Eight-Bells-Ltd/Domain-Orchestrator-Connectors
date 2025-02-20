# UMU Testbed - Orchestrator and Mitigation Actions

## Overview

This document provides details on the interface protocol, authentication, and format for mitigation actions in the UMU testbed. It also provides an example of how to send a mitigation action to the testbed.

## 1. Protocol Used

The orchestrator uses a **REST API** to apply orchestration policies (e.g., mitigation actions) on the testbed. The interface can be interacted with using `curl`.

### Example:
```bash
curl -H "Cache-Control: no-cache" -H "Content-Type: application/xml" -X POST -d @policy.xml http://orchestrator-url:8002/meservice
