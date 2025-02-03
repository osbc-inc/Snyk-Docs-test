# Swift 규칙

각 규칙은 다음 정보를 포함합니다.

* **규칙 이름**: 규칙의 Snyk 이름.
* **CWE(s)**: 이 규칙이 해당하는 [CWE 번호들](https://cwe.mitre.org/).
* **보안 범주**: 규칙이 속하는 [OWASP Top 10](https://owasp.org/Top10/) (2021 판) 범주 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되는 경우, 그 정보.
* **자동 수정 가능**: DeepCode AI Fix에 의해 자동으로 수정 가능한 보안 규칙. 이 정보는 지원되는 프로그래밍 언어에 대해만 포함됩니다.

| 규칙 이름                                                        | CWE(s)           | 보안 범주                  | 자동 수정 가능 |
| ------------------------------------------------------------ | ---------------- | ---------------------- | -------- |
| Clear Text Logging                                           | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | 아니오      |
| Code Injection                                               | CWE-94           | SANS Top 25, OWASP:A03 | 아니오      |
| Command Injection                                            | CWE-78           | SANS Top 25, OWASP:A03 | 아니오      |
| Device Authentication Bypass                                 | CWE-287          | SANS Top 25, OWASP:A07 | 아니오      |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05              | 아니오      |
| Use of a Broken or Risky Cryptographic Algorithm             | CWE-327          | OWASP:A02              | 아니오      |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02              | 아니오      |
| Information Exposure                                         | CWE-200          | OWASP:A01              | 아니오      |
| Use of Insufficiently Random Values                          | CWE-330          | OWASP:A02              | 아니오      |
| Insecure Data Storage                                        | CWE-922          | OWASP:A01              | 아니오      |
| Memory Corruption                                            | CWE-822          | 없음                     | 아니오      |
| Use of Hardcoded Credentials                                 | CWE-798          | SANS Top 25, OWASP:A07 | 아니오      |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | SANS Top 25, OWASP:A07 | 아니오      |
| Path Traversal                                               | CWE-23           | OWASP:A01              | 아니오      |
| Improper Certificate Validation                              | CWE-295          | OWASP:A07              | 아니오      |
| SQL Injection                                                | CWE-89           | SANS Top 25, OWASP:A03 | 아니오      |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | SANS Top 25, OWASP:A10 | 아니오      |
| Insecure Deserialization                                     | CWE-502          | SANS Top 25, OWASP:A08 | 아니오      |
| Inadequate Encryption Strength                               | CWE-326          | OWASP:A02              | 아니오      |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05              | N아니오     |
| Cross-site Scripting (XSS)                                   | CWE-79           | SANS Top 25, OWASP:A03 | 아니오      |
| XML External Entity (XXE) Injection                          | CWE-611          | OWASP:A05              | 아니오      |
