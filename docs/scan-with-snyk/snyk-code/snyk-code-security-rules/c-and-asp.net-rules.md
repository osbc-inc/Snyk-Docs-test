# C# 및 ASP.NET 규칙

각 규칙은 다음 정보를 포함합니다.

- **규칙 이름**: Snyk의 규칙 이름.
- **CWE(s)**: 이 규칙이 해당하는 [CWE 번호들](https://cwe.mitre.org/).
- **보안 카테고리**: 해당 규칙이 속하는 [OWASP Top 10](https://owasp.org/Top10/) (2021 버전) 카테고리, 있는 경우 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되어 있는지.
- **자동 수정 가능 여부**: DeepCode AI Fix에 의해 자동으로 수정 가능한 보안 규칙입니다. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| 규칙 이름                                                          | CWE(s)            | Security Categories    | 자동 수정 가능 여부 |
| ------------------------------------------------------------------- | ----------------- | ---------------------- | ------------------- |
| Anti-forgery token validation disabled                             | CWE-352           | Sans Top 25, OWASP:A01 | 예                  |
| Debug Features Enabled                                             | CWE-215           | 없음                   | 예                  |
| Usage of BinaryFormatter                                           | CWE-502           | Sans Top 25, OWASP:A08 | 아니요               |
| Cleartext Storage of Sensitive Information in a Cookie             | CWE-315           | OWASP:A05              | 아니요               |
| Code Injection                                                     | CWE-94            | Sans Top 25, OWASP:A03 | 아니요               |
| Command Injection                                                  | CWE-78            | Sans Top 25, OWASP:A03 | 아니요               |
| Deserialization of Untrusted Data                                  | CWE-502           | Sans Top 25, OWASP:A08 | 아니요               |
| Hardcoded Secret                                                   | CWE-547           | OWASP:A05              | 아니요               |
| Improper Neutralization of CRLF Sequences in HTTP Headers          | CWE-113           | OWASP:A03              | 아니요               |
| Use of a Broken or Risky Cryptographic Algorithm                   | CWE-327           | OWASP:A02              | 아니요               |
| Use of Password Hash With Insufficient Computational Effort        | CWE-916           | OWASP:A02              | 아니요               |
| Use of Insufficiently Random Values                                | CWE-330           | OWASP:A02              | 아니요               |
| Insecure Data Transmission                                         | CWE-319           | OWASP:A02              | 아니요               |
| LDAP Injection                                                     | CWE-90            | OWASP:A03              | 아니요               |
| Log Forging                                                        | CWE-117           | OWASP:A09              | 아니요               |
| Use of Hardcoded Credentials                                       | CWE-798           | Sans Top 25, OWASP:A07 | 아니요               |
| Open Redirect                                                      | CWE-601           | OWASP:A01              | 아니요               |
| Path Traversal                                                     | CWE-23            | OWASP:A01              | 아니요               |
| Exposure of Private Personal Information to an Unauthorized Actor  | CWE-359           | OWASP:A01              | 아니요               |
| Regular expression injection                                       | CWE-400, CWE-730  | 없음                   | 아니요               |
| Request Validation Disabled                                        | CWE-554           | 없음                   | 예                  |
| Information Exposure                                               | CWE-200           | OWASP:A01              | 아니요               |
| SQL Injection                                                      | CWE-89            | Sans Top 25, OWASP:A03 | 아니요               |
| Server-Side Request Forgery (SSRF)                                 | CWE-918           | Sans Top 25, OWASP:A10 | 아니요               |
| Inadequate Encryption Strength                                     | CWE-326           | OWASP:A02              | 아니요               |
| Sensitive Cookie Without 'HttpOnly' Flag                           | CWE-1004          | OWASP:A05              | 아니요               |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute       | CWE-614           | OWASP:A05              | 아니요               |
| Cross-site Scripting (XSS)                                         | CWE-79            | Sans Top 25, OWASP:A03 | 아니요               |
| XML External Entity (XXE) Injection                                | CWE-611           | OWASP:A05              | 아니요               |
| XAML Injection                                                     | CWE-611           | OWASP:A05              | 아니요               |
| XML Injection                                                      | CWE-91            | OWASP:A03              | 아니요               |
| XPath Injection                                                    | CWE-643           | OWASP:A03              | 아니요               |
| Arbitrary File Write via Archive Extraction (Zip Slip)             | CWE-22            | Sans Top 25, OWASP:A01 | 아니요               |