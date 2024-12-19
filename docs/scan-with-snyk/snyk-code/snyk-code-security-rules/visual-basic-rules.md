# Visual Basic 규칙

각 규칙에는 다음 정보가 포함됩니다.

* **규칙 이름**: 규칙의 Snyk 이름.
* **CWE(s):** 해당 규칙이 다루는 [CWE 번호](https://cwe.mitre.org/).
* **보안 범주**: 해당 규칙이 속하는 [OWASP Top 10 ](https://owasp.org/Top10/) (2021년 판) 범주 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되어 있는지 여부.
* **자동 수정 가능**: DeepCode AI Fix에 의해 자동 수정 가능한 보안 규칙. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| 규칙 이름                                                              | CWE(s)           | 보안 범주             | 자동 수정 가능 |
| ---------------------------------------------------------------------- | ---------------- | ----------------------- | -------------- |
| 디버그 기능 활성화                                                      | CWE-215          | 없음                    | 아니요         |
| BinaryFormatter 사용                                                    | CWE-502          | SANS Top 25, OWASP:A08  | 아니요         |
| 코드 삽입                                                               | CWE-94           | SANS Top 25, OWASP:A03  | 아니요         |
| 명령 삽입                                                              | CWE-78           | SANS Top 25, OWASP:A03  | 아니요         |
| 신뢰되지 않는 데이터의 역직렬화                                         | CWE-502          | SANS Top 25, OWASP:A08  | 아니요         |
| 하드코딩된 비밀                                                        | CWE-547          | OWASP:A05               | 아니요         |
| HTTP 헤더의 CRLF 시퀀스 무효화 미흡                                    | CWE-113          | OWASP:A03               | 아니요         |
| 취약한 또는 위험한 암호 알고리즘 사용                                  | CWE-327          | OWASP:A02               | 아니요         |
| 계산 비용이 충분하지 않는 비밀 해시 사용                              | CWE-916          | OWASP:A02               | 아니요         |
| 충분히 랜덤하지 않은 값 사용                                          | CWE-330          | OWASP:A02               | 아니요         |
| 하드코딩된 자격 증명 사용                                             | CWE-798          | SANS Top 25, OWASP:A07  | 아니요         |
| 개방형 리디렉트                                                       | CWE-601          | OWASP:A01               | 아니요         |
| 경로 탐색                                                             | CWE-23           | OWASP:A01               | 아니요         |
| 정규 표현식 삽입                                                      | CWE-400, CWE-730 | 없음                    | 아니요         |
| 요청 유효성 검사 비활성화                                              | CWE-554          | 없음                    | 아니요         |
| SQL 삽입                                                              | CWE-89           | SANS Top 25, OWASP:A03  | 아니요         |
| 서버 측 요청 위조 (SSRF)                                               | CWE-918          | SANS Top 25, OWASP:A10  | 아니요         |
| 불충분한 암호화 강도                                                   | CWE-326          | OWASP:A02               | 아니요         |
| 'HttpOnly' 플래그가 없는 민감한 쿠키                                 | CWE-1004         | OWASP:A05               | 아니요         |
| 'Secure' 속성이 없는 HTTPS 세션 내 민감한 쿠키                      | CWE-614          | OWASP:A05               | 아니요         |
| Cross-site Scripting (XSS)                                            | CWE-79           | SANS Top 25, OWASP:A03  | 아니요         |
| XML 외부 엔터티 (XXE) 삽입                                           | CWE-611          | OWASP:A05               | 아니요         |
| XML 삽입                                                             | CWE-91           | OWASP:A03               | 아니요         |
| XPath 삽입                                                           | CWE-643          | OWASP:A03               | 아니요         |