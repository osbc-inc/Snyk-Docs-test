# Apex 규칙

각 규칙은 다음 정보를 포함합니다.

- **Rule Name**: 규칙의 Snyk 이름.
- **CWE(s)**: 해당 규칙이 다루는 [CWE 번호](https://cwe.mitre.org/).
- **보안 카테고리**: 해당 규칙이 속한 [OWASP Top 10](https://owasp.org/Top10/) (2021년 판) 카테고리, 그리고 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함된 경우.
- **자동 수정 가능 여부**: DeepCode AI Fix에 의해 자동 수정 가능한 보안 규칙인지 여부. 이 정보는 지원되는 프로그래밍 언어에만 포함됩니다.

| Rule Name                                                    | CWE(s)           | 보안 카테고리         | 자동 수정 가능 |
| ------------------------------------------------------------ | ---------------- | --------------------- | ------------- |
| Access Violation                                             | CWE-284, CWE-285 | OWASP:A01             | 예           |
| Clear Text Sensitive Storage                                 | CWE-200, CWE-312 | OWASP:A01, OWASP:A04  | 아니요       |
| Command Injection                                            | CWE-78           | Sans Top 25, OWASP:A03 | 아니요       |
| Improper Access Control: Email Content Injection             | CWE-284          | OWASP:A01             | 아니요       |
| Use of Hardcoded Credentials                                 | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | 아니요       |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259 | Sans Top 25, OWASP:A07 | 아니요       |
| Hardcoded Secret                                             | CWE-547          | OWASP:A05             | 아니요       |
| Use of Password Hash With Insufficient Computational Effort  | CWE-916          | OWASP:A02             | 예           |
| Insecure Data Transmission                                   | CWE-319          | OWASP:A02             | 아니요       |
| Open Redirect                                                | CWE-601          | OWASP:A01             | 아니요       |
| Cross-site Scripting (XSS)                                   | CWE-79           | Sans Top 25, OWASP:A03 | 아니요       |
| Regular expression injection                                 | CWE-400, CWE-730 | 없음                  | 아니요       |
| SOQL Injection                                               | CWE-89           | Sans Top 25, OWASP:A03 | 아니요       |
| SOSL Injection                                               | CWE-89           | Sans Top 25, OWASP:A03 | 아니요       |
| Server-Side Request Forgery (SSRF)                           | CWE-918          | Sans Top 25, OWASP:A10 | 아니요       |
| Unverified Password Change                                   | CWE-620          | OWASP:A07             | 아니요       |
| Unsafe SOQL Concatenation                                    | CWE-89           | Sans Top 25, OWASP:A03 | 아니요       |
| Unsafe SOSL Concatenation                                    | CWE-89           | Sans Top 25, OWASP:A03 | 아니요       |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | CWE-614          | OWASP:A05             | 예           |
| XML Injection                                                | CWE-91           | OWASP:A03             | 아니요       |