# Go 규칙

각 규칙에는 다음 정보가 포함됩니다.

- **규칙 이름**: 규칙의 Snyk 이름입니다.
- **CWE(s)**: 이 규칙이 다루는 [CWE 번호](https://cwe.mitre.org/)입니다.
- **보안 범주**: 규칙이 속한 [OWASP Top 10](https://owasp.org/Top10/) (2021년 버전) 카테고리 및 해당 규칙이 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되어 있는 경우입니다.
- **자동 수정 가능**: DeepCode AI Fix에 의해 자동으로 수정 가능한 보안 규칙입니다. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| 규칙 이름                                                      | CWE(s)           | 보안 범주              | 자동 수정 가능 |
| -------------------------------------------------------------- | ---------------- | ---------------------- | ----------- |
| Clear Text Logging                                             | CWE-200, CWE-312 | OWASP:A01, OWASP:A04   | 아니오       |
| Command Injection                                              | CWE-78           | Sans Top 25, OWASP:A03 | 아니오       |
| Improper Access Control: Email Content Injection               | CWE-284          | OWASP:A01              | 아니오       |
| Generation of Error Message Containing Sensitive Information   | CWE-209          | OWASP:A04              | 아니오       |
| Hardcoded Secret                                               | CWE-547          | OWASP:A05              | 아니오       |
| Use of Hardcoded Passwords                                     | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | 아니오       |
| Use of a Broken or Risky Cryptographic Algorithm               | CWE-327          | OWASP:A02              | 네          |
| Use of Password Hash With Insufficient Computational Effort    | CWE-916          | OWASP:A02              | 네          |
| Insecure TLS Configuration                                     | CWE-327          | OWASP:A02              | 아니오       |
| Use of Insufficiently Random Values                            | CWE-330          | OWASP:A02              | 아니오       |
| Use of Hardcoded Credentials                                   | CWE-798          | Sans Top 25, OWASP:A07 | 아니오       |
| Open Redirect                                                  | CWE-601          | OWASP:A01              | 아니오       |
| Path Traversal                                                 | CWE-23           | OWASP:A01              | 아니오       |
| SQL Injection                                                  | CWE-89           | Sans Top 25, OWASP:A03 | 아니오       |
| Server-Side Request Forgery (SSRF)                             | CWE-918          | Sans Top 25, OWASP:A10 | 아니오       |
| Improper Neutralization of Directives in Statically Saved Code | CWE-96           | OWASP:A03              | 아니오       |
| Improper Certificate Validation                                | CWE-295          | OWASP:A07              | 아니오       |
| Inadequate Encryption Strength                                 | CWE-326          | OWASP:A02              | 네          |
| Sensitive Cookie Without 'HttpOnly' Flag                       | CWE-1004         | OWASP:A05              | 네          |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute   | CWE-614          | OWASP:A05              | 네          |
| Cross-site Scripting (XSS)                                     | CWE-79           | Sans Top 25, OWASP:A03 | 아니오       |
| XPath Injection                                                | CWE-643          | OWASP:A03              | 아니오       |