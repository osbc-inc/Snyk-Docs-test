# XML 규칙

각 규칙은 다음 정보를 포함합니다.

- **Rule Name**: 규칙의 Snyk 이름
- **CWE(s)**: 이 규칙이 해당하는 [CWE 번호](https://cwe.mitre.org/)
- **보안 카테고리**: 규칙이 속한 [OWASP Top 10](https://owasp.org/Top10/) (2021 버전) 카테고리 및 [SANS 25](https://www.sans.org/top25-software-errors/)에 포함되어 있는 경우
- **자동 수정 가능**: 지원되는 프로그래밍 언어에 대해서 DeepCode AI Fix에 의해 자동으로 수정 가능한 보안 규칙에 관한 정보

| 규칙 이름                                                     | CWE(s)                | Security Categories        | 자동 수정 가능 |
| ------------------------------------------------------------ | --------------------- | -------------------------- | ----------- |
| Android Debug Mode Enabled                                   | CWE-489               | 없음                       | 아니요       |
| Debug Features Enabled                                       | CWE-215               | 없음                       | 아니요       |
| Generation of Error Message Containing Sensitive Information | CWE-209               | OWASP:A04                  | 아니요       |
| Improper Restriction of Rendered UI Layers or Frames         | CWE-1021              | OWASP:A04                  | 아니요       |
| ASP SSL Disabled                                             | CWE-319               | OWASP:A02                  | 아니요       |
| Use of Hardcoded Passwords                                   | CWE-798, CWE-259      | Sans Top 25, OWASP:A07     | 아니요       |
| Request Validation Disabled                                  | CWE-554               | 없음                       | 아니요       |
| Struts Development Mode Enabled                              | CWE-489               | 없음                       | 아니요       |