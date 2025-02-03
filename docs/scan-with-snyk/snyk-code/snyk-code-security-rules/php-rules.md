# PHP 규칙

각 규칙은 다음 정보를 포함합니다.

* **Rule Name**: 규칙의 Snyk 이름
* **CWE(s)**: 이 규칙이 다루는 [CWE 번호](https://cwe.mitre.org/)
* **보안 카테고리**: 규칙이 속한 [OWASP Top 10](https://owasp.org/Top10/) (2021년판) 범주 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함된 경우
* **자동 수정 가능**: DeepCode AI Fix에 의해 자동으로 수정 가능한 보안 규칙. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| Rule Name                                                    | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Code Injection                                               | CWE-94           | Sans Top 25, OWASP:A03 | 아니오         |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | 아니오         |
| Deserialization of Untrusted Data                            | CWE-502          | Sans Top 25, OWASP:A08 | 아니오         |
| Improper Access Control: Email Content Injection             | CWE-284          | OWASP:A01              | 아니오         |
| File Inclusion                                               | CWE-98           | OWASP:A03              | 아니오         |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | 아니오         |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | 아니오         |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | 아니오         |
| Inadequate Padding for Public Key Encryption                 | CWE-326          | OWASP:A02              | 아니오         |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | 아니오         |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | 아니오         |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | 아니오         |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | 아니오         |
| Allocation of Resources Without Limits or Throttling         | CWE-770          | 없음                     | 아니오         |
| Open Redirect                                                | CWE-601          | OWASP:A01              | 아니오         |
| Path Traversal                                               | CWE-23           | OWASP:A01              | 아니오         |
| Information Exposure                                         | CWE-200          | OWASP:A01              | 아니오         |
| SQL Injection                                                | CWE-89           | Sans Top 25, OWASP:A03 | 아니오         |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | 아니오         |
| Origin Validation Error                                      | CWE-942, CWE-346 | OWASP:A05, OWASP:A07   | 아니오         |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021         | OWASP:A04              | 아니오         |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | 아니오         |
| Sensitive Cookie Without 'HttpOnly' Flag                     | CWE-1004         | OWASP:A05              | 아니오         |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | 아니오         |
| XPath Injection                                              | CWE-643          | OWASP:A03              | 아니오         |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | 아니오         |
| Arbitrary File Write via Archive Extraction (Zip Slip)       | CWE-22           | Sans Top 25, OWASP:A01 | 아니오         |
| Regular Expression Denial of Service (ReDoS)                 | CWE-400          | 없음                     | 아니오         |
