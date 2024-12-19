# JavaScript 및 TypeScript 규칙

각 규칙에는 다음 정보가 포함됩니다.

* **규칙 이름**: 규칙의 Snyk 이름입니다.
* **CWE(s)**: 이 규칙에서 다루는 [CWE 번호](https://cwe.mitre.org/)입니다.
* **보안 카테고리**: 해당 규칙이 속한 [OWASP Top 10](https://owasp.org/Top10/) (2021 에디션) 카테고리 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되는 경우, 이를 나타냅니다.
* **자동 수정 가능**: DeepCode AI Fix에 의해 자동 수정 가능한 보안 규칙입니다. 이 정보는 지원되는 프로그래밍 언어에 대해서만 포함됩니다.

| 규칙 이름                                                                                                         | CWE(s)                  | 보안 카테고리       | 자동 수정 가능 |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------- | ------------------ | ------------- |
| 엄격한 문맥적 이스케이핑(SCE) 비활성화는 크로스사이트 스크립팅(XSS)에 대한 추가 공격 표면을 제공할 수 있습니다.           | CWE-79                  | Sans Top 25, OWASP:A03 | 예         |
| 손상된 또는 위험한 암호화 알고리즘 사용                                                                                  | CWE-327                 | OWASP:A02              | 아니오       |
| 평문으로 민감한 데이터 저장                                                                                         | CWE-200, CWE-312        | OWASP:A01, OWASP:A04   | 아니오       |
| 코드 삽입                                                                                                          | CWE-94                  | Sans Top 25, OWASP:A03 | 예         |
| 명령 삽입                                                                                                         | CWE-78                  | Sans Top 25, OWASP:A03 | 예         |
| 크로스사이트 스크립팅 (XSS)                                                                                            | CWE-79                  | Sans Top 25, OWASP:A03 | 예         |
| 신뢰할 수 없는 데이터의 역직렬화                                                                                     | CWE-502                 | Sans Top 25, OWASP:A08 | 아니오       |
| 정보 노출                                                                                                          | CWE-200                 | OWASP:A01              | 예         |
| Electron 보안 경고 비활성화                                                                                       | CWE-16                  | OWASP:A05              | 아니오       |
| Electron 보안이 취약한 웹 환경설정                                                                               | CWE-16                  | OWASP:A05              | 예         |
| Electron 보안이 취약한 콘텐츠 로드                                                                                 | CWE-16                  | OWASP:A05              | 예         |
| 외부로 제어되는 포맷 문자열 사용                                                                              | CWE-134                 | 없음                   | 예         |
| GraphQL 삽입                                                                                                      | CWE-89                  | Sans Top 25, OWASP:A03 | 아니오       |
| 부적절한 유형 유효성 검사                                                                                         | CWE-1287                | 없음                   | 예         |
| 하드코딩된 비밀                                                                      | CWE-547                 | OWASP:A05              | 예         |
| 민감한 정보의 평문 전송                                                                                           | CWE-319                 | OWASP:A02              | 예         |
| 부적절한 코드 세정화                                                                                             | CWE-94, CWE-79, CWE-116 | Sans Top 25, OWASP:A03 | 아니오       |
| 계산 비용 부족으로 한 패스워드 해시 사용                                                                      | CWE-916                 | OWASP:A02              | 예         |
| 충분히 랜덤한 값 사용                                                                                          | CWE-330                 | OWASP:A02              | 아니오       |
| 보안이 취약한 TLS 구성                                                                                           | CWE-327                 | OWASP:A02              | 예         |
| 미검증된 postMessage 유효성                                                                                 | CWE-20                  | Sans Top 25, OWASP:A03 | 예         |
| Introspection 활성화                                                                                           | CWE-200                 | OWASP:A01              | 아니오       |
| 보안이 취약한 JWT 확인 방법                                                                               | CWE-347                 | OWASP:A02              | 아니오       |
| JWT 서명 확인 방법 비활성화                                                                       | CWE-347                 | OWASP:A02              | 아니오     |
| JWT 'none' 알고리즘 사용 가능                                                                    | CWE-347                 | OWASP:A02              | 아니오       |
| 중첩된 GraphQL 쿼리로 인한 서비스 거부 (DoS)                                                               | CWE-400                 | 없음                   | 예         |
| 루프 조건에 대한 입력 미확인                                                                               | CWE-400, CWE-606        | 없음                   | 아니오       |
| 관측 가능한 타이밍 불일치 (타이밍 공격)                                                                     | CWE-208                 | 없음                   | 아니오       |
| 하드코딩된 자격 증명 사용                                                                                    | CWE-798                 | Sans Top 25, OWASP:A07 | 예         |
| 하드코딩된 비밀번호 사용                                                                                      | CWE-798, CWE-259        | Sans Top 25, OWASP:A07 | 예         |
| 제한 없이 또는 쓰로틀링 없이 자원 할당                                                                  | CWE-770                 | 없음                   | 예         |
| NoSQL 삽입                                                                                                   | CWE-943                 | 없음                   | 아니오       |
| 버퍼 오버리드                                                                                                  | CWE-126                 | 없음                   | 아니오       |
| 열린 리디렉트                                                                                                    | CWE-601                 | OWASP:A01              | 예         |
| 경로 탐색                                                                                                      | CWE-23                  | OWASP:A01              | 예         |
| Prototype Pollution                                                                                               | CWE-1321                | 없음                   | 아니오       |
| XSS 위험 처리를 위해 dangerouslySetInnerHTML 사용                                                          | CWE-79                  | Sans Top 25, OWASP:A03 | 예         |
| 잊어버린 비밀번호를 위한 취약한 비밀번호 복구 메커니즘                                                    | CWE-640                 | OWASP:A07              | 아니오       |
| SQL 삽입                                                                                                         | CWE-89                  | Sans Top 25, OWASP:A03 | 예         |
| 서버측 요청 위조 (SSRF)                                                                                   | CWE-918                 | Sans Top 25, OWASP:A10 | 아니오       |
| 정적으로 저장된 코드의 지시문의 부적절한 중화                                  | CWE-96                  | OWASP:A03              | 아니오       |
| 출처 확인 오류                                                                                           | CWE-942, CWE-346        | OWASP:A05, OWASP:A07   | 예         |
| 권한 있는 Cross-domain 정책                                                                      | CWE-942                 | OWASP:A05              | 예         |
| 렌더링된 UI 레이어 또는 프레임의 부적절한 제한                                                             | CWE-1021                | OWASP:A04              | 아니오       |
| 암호화 문제                                                                                               | CWE-310                 | OWASP:A02              | 예         |
| 안전하지 않은 JQuery 플러그인                                                                           | CWE-79, CWE-116         | Sans Top 25, OWASP:A03 | 아니오       |
| 크로스사이트 요청 위조 (CSRF)                                                                       | CWE-352                 | Sans Top 25, OWASP:A01 | 예         |
| 'HttpOnly' 플래그가없는 민감한 쿠키                                                            | CWE-1004                | OWASP:A05              | 예         |
| HTTPS 세션에서 'Secure' 속성이없는 민감한 쿠키                                                   | CWE-614                 | OWASP:A05              | 예         |
| XML 외부 엔티티 (XXE) 삽입                                                                             | CWE-611                 | OWASP:A05              | 아니오       |
| XPath 삽입                                                                                                   | CWE-643                 | OWASP:A03              | 아니오       |
| 압축을 통한 임의 파일 쓰기 (Zip Slip)                                                             | CWE-22                  | Sans Top 25, OWASP:A01 | 아니오       |
| 정규식 거부 서비스 (ReDoS)                                                                            | CWE-400                 | 없음                   | 예         |