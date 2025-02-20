#!/usr/bin/env python3

import json
import xml.etree.ElementTree as ET
from typing import Dict, List
import requests

def json_to_mspl(json_data: Dict) -> str:
    """
    Transform JSON mitigation action into MSPL XML format.
    """
    # Create the root element
    mspl = ET.Element("mspl")
    policy = ET.SubElement(mspl, "policy")

    # Map JSON fields to XML
    ET.SubElement(policy, "name").text = json_data.get("name", "dns_rate_limiting")
    ET.SubElement(policy, "intent_id").text = json_data.get("intent_id", "30001")
    ET.SubElement(policy, "action").text = "rate_limit"  # Assumed from "dns_rate_limiting"

    # Target section
    target = ET.SubElement(policy, "target")
    ET.SubElement(target, "service").text = "dns"  # Assumed for DNS rate limiting
    source_ips = json_data["fields"].get("source_ip_filter", [])
    if isinstance(source_ips, list) and source_ips:
        ET.SubElement(target, "source_ip_filter").text = ",".join(source_ips)
    else:
        ET.SubElement(target, "source_ip_filter").text = str(source_ips)

    # Parameters section
    parameters = ET.SubElement(policy, "parameters")
    ET.SubElement(parameters, "rate").text = str(json_data["fields"].get("rate", "20"))
    ET.SubElement(parameters, "duration").text = str(json_data["fields"].get("duration", "60"))

    # Convert to string with proper formatting
    rough_string = ET.tostring(mspl, encoding="unicode")
    # Add proper XML declaration and formatting (optional, for readability)
    return '<?xml version="1.0" encoding="UTF-8"?>\n' + rough_string

def send_mitigation_action(xml_payload: str, endpoint: str) -> requests.Response:
    """
    Send the XML payload to the UMU testbed endpoint via HTTP POST.
    """
    headers = {
        "Cache-Control": "no-cache",
        "Content-Type": "application/xml"
    }

    try:
        response = requests.post(endpoint, data=xml_payload, headers=headers, timeout=10)
        response.raise_for_status()  # Raise an exception for bad status codes
        return response
    except requests.RequestException as e:
        print(f"Error sending request: {e}")
        return None

def main():
    # Sample JSON mitigation action (replace with your actual data or read from file)
    sample_json = {
        "name": "dns_rate_limiting",
        "intent_id": "30001",
        "fields": {
            "rate": "20",
            "duration": "60",
            "source_ip_filter": ["malicious_ips"]
        }
    }

    # Transform JSON to MSPL XML
    xml_payload = json_to_mspl(sample_json)
    print("Generated MSPL XML:")
    print(xml_payload)

    # UMU testbed endpoint (replace with actual URL from UMU team)
    orchestrator_url = "http://10.20.30.40:8002/meservice"  # Placeholder—update this!
    
    # Send the request
    response = send_mitigation_action(xml_payload, orchestrator_url)
    
    if response:
        print(f"Request successful! Status code: {response.status_code}")
        print(f"Response: {response.text}")
    else:
        print("Request failed.")

if __name__ == "__main__":
    main()