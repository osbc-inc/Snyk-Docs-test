# Ruby 규칙

각 규칙은 다음 정보를 포함합니다.

* **규칙 이름**: 규칙의 Snyk 이름.
* **CWE(s)**: 이 규칙이 해당하는 [CWE 번호들](https://cwe.mitre.org/).
* **보안 카테고리**: 이 규칙이 속한 [OWASP Top 10 ](https://owasp.org/Top10/)(2021 버전) 카테고리, 그리고 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함된 경우.
* **자동 수정 가능**: DeepCode AI Fix에 의해 자동으로 수정 가능한 보안 규칙입니다. 이 정보는 지원되는 프로그래밍 언어에만 포함됩니다.

| 규칙 이름                             | CWE(s)                                                      | 보안 카테고리                                                 | 자동 수정 가능 |
| --------------------------------- | ----------------------------------------------------------- | ------------------------------------------------------- | -------- |
| 코드 삽입                             | CWE-94                                                      | Sans Top 25, OWASP:A03                                  | 아니요      |
| 명령 삽입                             | CWE-78                                                      | Sans Top 25, OWASP:A03                                  | 아니요      |
| 엔드포인트를 통한 원격 코드 실행                | CWE-94                                                      | Sans Top 25, OWASP:A03                                  | 아니요      |
| 신뢰할 수 없는 데이터의 역직렬화                | CWE-502                                                     | Sans Top 25, OWASP:A08                                  | 아니요      |
| 하드코딩된 자격 증명 사용                    | CWE-798, CWE-259                                            | Sans Top 25, OWASP:A07                                  | 아니요      |
| 하드코딩된 암호화 키 사용                    | CWE-321                                                     | OWASP:A02                                               | 아니요      |
| 하드코딩된 비밀 사용                       | CWE-547                                                     | OWASP:A05                                               | 아니요      |
| 하드코딩된 암호 사용                       | CWE-798, CWE-259                                            | Sans Top 25, OWASP:A07                                  | 아니요      |
| 깨진 또는 위험한 암호 알고리즘 사용              | CWE-327                                                     | OWASP:A02                                               | 아니요      |
| 계산 노력이 충분하지 않은 비밀번호 해시 사용         | CWE-916                                                     | OWASP:A02                                               | 아니요      |
| 충분히 무작위 값 사용                      | CWE-330                                                     | OWASP:A02                                               | 아니요      |
| Sinatra 보호 계층 비활성화                | CWE-16, CWE-352, CWE-79, CWE-693, CWE-1021, CWE-35, CWE-348 | Sans Top 25, OWASP:A05, OWASP:A03, OWASP:A01, OWASP:A04 | 아니요      |
| 안전하지 않은 데이터 전송                    | CWE-319                                                     | OWASP:A02                                               | 아니요      |
| 부적절한 입력 유효성 검사                    | CWE-20                                                      | Sans Top 25, OWASP:A03                                  | 아니요      |
| 동적으로 결정된 객체 속성의 부적절한 제어           | CWE-915                                                     | OWASP:A08                                               | 아니요      |
| 협상 중 덜 안전한 알고리즘 선택 (SSL 강제)       | CWE-311, CWE-757                                            | OWASP:A02, OWASP:A04                                    | 아니요      |
| 열린 리디렉트                           | CWE-601                                                     | OWASP:A01                                               | 아니요      |
| 경로 탐색                             | CWE-23                                                      | OWASP:A01                                               | 아니요      |
| 안전하지 않은 반사                        | CWE-470                                                     | OWASP:A03                                               | 아니요      |
| 부적절한 인증서 유효성 검사                   | CWE-295                                                     | OWASP:A07                                               | 아니요      |
| 정보 노출                             | CWE-200                                                     | OWASP:A01                                               | 아니요      |
| 세션 조작                             | CWE-285                                                     | OWASP:A01                                               | 아니요      |
| SQL 삽입                            | CWE-89                                                      | Sans Top 25, OWASP:A03                                  | 아니요      |
| 정적으로 저장된 코드의 지시문의 부적절한 중립화        | CWE-96                                                      | OWASP:A03                                               | 아니요      |
| 약한 암호 요구 사항 없음                    | CWE-521                                                     | OWASP:A07                                               | 아니요      |
| 'HttpOnly' 플래그가 없는 민감한 쿠키         | CWE-1004                                                    | OWASP:A05                                               | 아니요      |
| 'Secure' 속성이 없는 HTTPS 세션에서 민감한 쿠키 | CWE-614                                                     | OWASP:A05                                               | 아니요      |
| 값 유효성 검사를 위한 잘못된 정규 표현식           | CWE-1286                                                    | 없음                                                      | 아니요      |
| 크로스사이트 스크립팅 (XSS)                 | CWE-79                                                      | Sans Top 25, OWASP:A03                                  | 아니요      |
| XML 외부 엔터티 (XXE) 삽입               | CWE-611                                                     | OWASP:A05                                               | 아니요      |
| XPath 삽입                          | CWE-643                                                     | OWASP:A03                                               | 아니요      |
| 정규 표현식 거부 서비스 (ReDoS)             | CWE-400                                                     | 없음                                                      | 아니요      |
