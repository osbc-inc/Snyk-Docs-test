# C++ 규칙

각 규칙은 다음 정보를 포함합니다.

* **Rule Name**: Snyk의 규칙 이름.
* **CWE(s):** 이 규칙에 의해 다루어지는 [CWE 번호들](https://cwe.mitre.org/).
* **보안 카테고리**: 이 규칙이 속하는 경우, [OWASP Top 10](https://owasp.org/Top10/) 카테고리 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함된 경우.
* **자동 수정 가능**: DeepCode AI Fix에 의해 자동 수정 가능한 보안 규칙입니다. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| Rule Name                                                                | CWE(s)           | Security Categories    | Autofixable |
| ------------------------------------------------------------------------ | ---------------- | ---------------------- | ----------- |
| Memory Allocation Of String Length                                       | CWE-170          | 없음                     | 예           |
| Insecure Anonymous LDAP Binding                                          | CWE-287          | Sans Top 25, OWASP:A07 | 아니오         |
| Buffer Overflow                                                          | CWE-122          | 없음                     | 예           |
| Division By Zero                                                         | CWE-369          | 없음ne                   | 아니오         |
| Missing Release of File Descriptor or Handle after Effective Lifetime    | CWE-775          | 없음                     | 예           |
| Command Injection                                                        | CWE-78           | Sans Top 25, OWASP:A03 | 아니오         |
| Dereference of a NULL Pointer                                            | CWE-476          | Sans Top 25            | 아니오         |
| Double Free                                                              | CWE-415          | 없음                     | 예           |
| Use of Externally-Controlled Format String                               | CWE-134          | 없음                     | 예           |
| Use of Hardcoded Cryptographic Key                                       | CWE-321          | OWASP:A02              | 아니오         |
| Improper Null Termination                                                | CWE-170          | 없음                     | 아니오         |
| Use of Password Hash With Insufficient Computational Effort              | CWE-916          | OWASP:A02              | 예           |
| Integer Overflow                                                         | CWE-190          | Sans Top 25            | 아니오         |
| LDAP Injection                                                           | CWE-90           | OWASP:A03              | 아니오         |
| Missing Release of Memory after Effective Lifetime                       | CWE-401          | 없음                     | 예           |
| An optimizing compiler may remove memset non-zero leaving data in memory | CWE-1330         | 없음                     | No아니오       |
| Potential Negative Number Used as Index                                  | CWE-125, CWE-787 | Sans Top 25            | 아니오         |
| Path Traversal                                                           | CWE-23           | OWASP:A01              | 아니오         |
| Exposure of Private Personal Information to an Unauthorized Actor        | CWE-359          | OWASP:A01              | 아니오         |
| Size Used as Index                                                       | CWE-125, CWE-787 | Sans Top 25            | 예           |
| SQL Injection                                                            | CWE-89           | Sans Top 25, OWASP:A03 | 아니오         |
| Server-Side Request Forgery (SSRF)                                       | CWE-918          | Sans Top 25, OWASP:A10 | 아니오         |
| Inadequate Encryption Strength                                           | CWE-326          | OWASP:A02              | 예           |
| Potential buffer overflow from usage of unsafe function                  | CWE-122          | 없음                     | 예           |
| Use of Expired File Descriptor                                           | CWE-910          | 없음                     | 아니오         |
| Use After Free                                                           | CWE-416          | Sans Top 25            | 아니오         |
| User Controlled Pointer                                                  | CWE-1285         | 없음                     | 아니오         |
| Authentication Bypass by Spoofing                                        | CWE-290          | OWASP:A07              | 아니오         |
| Cross-site Scripting (XSS)                                               | CWE-79           | Sans Top 25, OWASP:A03 | 아니오         |
| XML External Entity (XXE) Injection                                      | CWE-611          | OWASP:A05              | 아니오         |
| XPath Injection                                                          | CWE-643          | OWASP:A03              | 아니오         |
