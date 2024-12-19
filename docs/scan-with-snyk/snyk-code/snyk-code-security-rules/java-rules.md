# 자바 규칙

각 규칙에는 다음 정보가 포함됩니다.

* **Rule Name**: Snyk 이름의 규칙.
* **CWE(s):** 이 규칙에 해당하는 [CWE 번호](https://cwe.mitre.org/).
* **보안 카테고리**: 이 규칙이 속한 [OWASP Top 10 ](https://owasp.org/Top10/) 카테고리(2021년판) 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되어 있는 경우.
* **자동 수정 가능 여부**: DeepCode AI Fix로 자동 수정 가능한 보안 규칙에 대한 정보입니다. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Android World Writeable/Readable File Permission Found       | CWE-732          | None                   | Yes         |
| Use of Potentially Dangerous Function                        | CWE-676          | None                   | No          |
| Cleartext Storage of Sensitive Information in a Cookie       | CWE-315          | OWASP:A05              | No          |
| Code Injection                                               | CWE-94           | Sans Top 25, OWASP:A03 | No          |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| Deserialization of Untrusted Data                            | CWE-502          | Sans Top 25, OWASP:A08 | Yes         |
| Cross-Site Request Forgery (CSRF)                            | CWE-352          | Sans Top 25, OWASP:A01 | Yes         |
| Information Exposure                                         | CWE-200          | OWASP:A01              | No          |
| Cleartext Transmission of Sensitive Information              | CWE-319          | OWASP:A02              | No          |
| Indirect Command Injection via User Controlled Environment   | CWE-78           | Sans Top 25, OWASP:A03 | No          |
| External Control of System or Configuration Setting          | CWE-15           | OWASP:A05              | No          |
| Process Control                                              | CWE-114          | None                   | No          |
| File Access Enabled                                          | CWE-200          | OWASP:A01              | Yes         |
| Android Fragment Injection                                   | CWE-470          | OWASP:A03              | Yes         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | No          |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | No          |
| Improper Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | No          |
| Disabled Neutralization of CRLF Sequences in HTTP Headers    | CWE-113          | OWASP:A03              | No          |
| Inadequate Padding for AES encryption                        | CWE-326          | OWASP:A02              | Yes         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | No          |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | Yes         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | Yes         |
| Android Intent Forwarding                                    | CWE-940          | OWASP:A07              | No          |
| Improper Validation of Certificate with Host Mismatch        | CWE-297          | OWASP:A07              | No          |
| JavaScript Enabled                                           | CWE-79           | Sans Top 25, OWASP:A03 | Yes         |
| Java Naming and Directory Interface (JNDI) Injection         | CWE-074          | None                   | No          |
| JWT Signature Verification Bypass                            | CWE-347          | OWASP:A02              | No          |
| Improper Authentication                                      | CWE-287          | Sans Top 25, OWASP:A07 | No          |
| LDAP Injection                                               | CWE-90           | OWASP:A03              | No          |
| Use of Hardcoded Credentials                                 | CWE-798          | Sans Top 25, OWASP:A07 | No          |
| The cipher text is equal to the provided input plain text    | CWE-311          | OWASP:A04              | No          |
| NoSQL Injection                                              | CWE-943          | None                   | No          |
| Use of Sticky broadcasts                                     | CWE-265          | None                   | No          |
| Use of Hardcoded, Security-relevant Constants                | CWE-547          | OWASP:A05              | No          |
| Open Redirect                                                | CWE-601          | OWASP:A01              | No          |
| Path Traversal                                               | CWE-23           | OWASP:A01              | Yes         |
| Privacy Leak                                                 | CWE-532          | OWASP:A09              | No          |
| Unsafe Reflection                                            | CWE-470          | OWASP:A03              | No          |
| Regular expression injection                                 | CWE-400, CWE-730 | None                   | No          |
| Unprotected Storage of Credentials                           | CWE-256          | OWASP:A04              | No          |
| Incorrect Permission Assignment                              | CWE-732          | None                   | No          |
| Server Information Exposure                                  | CWE-209          | OWASP:A04              | Yes         |
| Improper Handling of Insufficient Permissions or Privileges  | CWE-280          | OWASP:A04              | No          |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | Yes         |
| Unrestricted Android Broadcast                               | CWE-862          | Sans Top 25, OWASP:A01 | No          |
| Spring Cross-Site Request Forgery (CSRF)                     | CWE-352          | Sans Top 25, OWASP:A01 | No          |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | Yes         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | No          |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | Yes         |
| Code Execution via Third Party Package Context               | CWE-94           | Sans Top 25, OWASP:A03 | No          |
| Code Execution via Third Party Package Installation          | CWE-940          | OWASP:A07              | No          |
| Observable Timing Discrepancy (Timing Attack)                | CWE-208          | None                   | Yes         |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | Yes         |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07              | No          |
| Cryptographic Issues                                         | CWE-310          | OWASP:A02              | Yes         |
| Trust Boundary Violation                                     | CWE-501          | OWASP:A04              | No          |
| Unauthorized File Access                                     | CWE-79           | Sans Top 25, OWASP:A03 | No          |
| Android Uri Permission Manipulation                          | CWE-266          | OWASP:A04              | No          |
| Use of Externally-Controlled Format String                   | CWE-134          | None                   | Yes         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | Yes         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | Yes         |
| Insufficient Session Expiration                              | CWE-613          | OWASP:A07              | No          |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | Yes         |
| XPath Injection                                              | CWE-643          | OWASP:A03              | No          |