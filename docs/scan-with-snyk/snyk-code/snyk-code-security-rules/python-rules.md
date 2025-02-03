# Python 규칙

각 규칙은 다음의 정보를 포함합니다.

* **Rule Name**: 규칙의 Snyk 이름
* **CWE(s)**: 이 규칙에 해당하는 [CWE 번호](https://cwe.mitre.org/)
* **보안 카테고리**: 이 규칙이 어떤 [OWASP Top 10](https://owasp.org/Top10/) (2021년 판) 카테고리에 속하는지, 그리고 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되어 있는지 여부
* **자동 수정 가능 여부**: DeepCode AI Fix에 의해 자동으로 수정될 수 있는 보안 규칙입니다. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| Rule Name                                                                  | CWE(s)           | Security Categories    | Autofixable |
| -------------------------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Authentication over HTTP                                                   | CWE-319          | OWASP:A02              | 예           |
| Binding to all network interfaces may open service to unintended traffic   | CWE-284          | OWASP:A01              | 예           |
| Broken User Authentication                                                 | CWE-287          | Sans Top 25, OWASP:A07 | 아니오         |
| Code Injection                                                             | CWE-94           | Sans Top 25, OWASP:A03 | 예           |
| Command Injection                                                          | CWE-78           | Sans Top 25, OWASP:A03 | 예           |
| Deserialization of Untrusted Data                                          | CWE-502          | Sans Top 25, OWASP:A08 | 예           |
| Cross-Site Request Forgery (CSRF)                                          | CWE-352          | Sans Top 25, OWASP:A01 | 예           |
| Password Requirements Not Enforced in Django Application                   | CWE-521          | OWASP:A07              | 아니오         |
| Use of Hardcoded Cryptographic Initialization Value                        | CWE-329          | OWASP:A02              | 아니오         |
| Use of Hardcoded Cryptographic Key                                         | CWE-321          | OWASP:A02              | 아니오         |
| Hardcoded Secret                                                           | CWE-547          | OWASP:A05              | 예           |
| Use of a Broken or Risky Cryptographic Algorithm                           | CWE-327          | OWASP:A02              | 아니오         |
| Insecure default value                                                     | CWE-453          | 없음                     | 아니오         |
| Insecure File Permissions                                                  | CWE-732          | 없음                     | 예           |
| Use of Password Hash With Insufficient Computational Effort                | CWE-916          | OWASP:A02              | 예           |
| Insecure Temporary File                                                    | CWE-377          | OWASP:A01              | 예           |
| Insecure Xml Parser                                                        | CWE-611          | OWASP:A05              | 아니오         |
| Jinja auto-escape is set to false.                                         | CWE-79           | Sans Top 25, OWASP:A03 | 예           |
| LDAP Injection                                                             | CWE-90           | OWASP:A03              | 아니오         |
| Improper Handling of Insufficient Permissions or Privileges                | CWE-280          | OWASP:A04              | 아니오         |
| Use of Hardcoded Credentials                                               | CWE-798          | Sans Top 25, OWASP:A07 | 아니오         |
| Use of Hardcoded Passwords                                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | 아니오         |
| NoSQL Injection                                                            | CWE-943          | 없음                     | 아니오         |
| Open Redirect                                                              | CWE-601          | OWASP:A01              | 아니오         |
| Path Traversal                                                             | CWE-23           | OWASP:A01              | 예           |
| Debug Mode Enabled                                                         | CWE-489          | 없음                     | 예           |
| Improper Certificate Validation                                            | CWE-295          | OWASP:A07              | 아니오         |
| Server Information Exposure                                                | CWE-209          | OWASP:A04              | 아니오         |
| SQL Injection                                                              | CWE-89           | Sans Top 25, OWASP:A03 | 예           |
| Server-Side Request Forgery (SSRF)                                         | CWE-918          | Sans Top 25, OWASP:A10 | 아니오         |
| Improper Neutralization of Directives in Statically Saved Code             | CWE-96           | OWASP:A03              | 아니오         |
| Inadequate Encryption Strength                                             | CWE-326          | OWASP:A02              | 예           |
| Arbitrary File Write via Archive Extraction (Tar Slip)                     | CWE-22           | Sans Top 25, OWASP:A01 | 아니오         |
| Origin Validation Error                                                    | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | 아니오         |
| Cryptographic Issues                                                       | CWE-310          | OWASP:A02              | 아니오         |
| Python 2 source code                                                       | CWE-1104         | OWASP:A06              | 아니오         |
| Selection of Less-Secure Algorithm During Negotiation (SSL instead of TLS) | CWE-757          | OWASP:A02              | 아니오         |
| Sensitive Cookie Without 'HttpOnly' Flag                                   | CWE-1004         | OWASP:A05              | 예           |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute               | CWE-614          | OWASP:A05              | 예           |
| Cross-site Scripting (XSS)                                                 | CWE-79           | Sans Top 25, OWASP:A03 | 예           |
| XPath Injection                                                            | CWE-643          | OWASP:A03              | 아니오         |
| Regular Expression Denial of Service (ReDoS)                               | CWE-400          | 없음                     | 아니오         |

