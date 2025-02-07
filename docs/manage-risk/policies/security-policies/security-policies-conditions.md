# 보안 정책 조건

조건은 규칙을 설정하는 변수입니다.

조건 카테고리를 선택한 후, **포함** 또는 **포함하지 않음** 및 원하는 조건을 선택하도록 요구됩니다. 예를 들어, 성숙, CWE-1234를 선택할 수 있습니다.

보안 정책은 와 프로젝트에 적용됩니다.

**AND** 함수를 사용하여 동일한 규칙에 여러 조건을 중첩할 수 있습니다.

<table data-header-hidden><thead><tr><th width="210"></th><th></th></tr></thead><tbody><tr><td><strong>조건 카테고리</strong></td><td><strong>조건 정의 및 변수</strong></td></tr><tr><td>공격 성숙도</td><td>지정된 공격 성숙도 값을 포함하는 모든 문제와 일치합니다. 여러 값을 선택할 때, 조건은 선택된 값 중 어떤 값을 포함하는 모든 문제에 적용됩니다. 값은 <code>Mature</code>, <code>Proof of Concept</code>, <code>No known exploit,</code> <code>No data available</code>을 포함합니다.</td></tr><tr><td>CWE</td><td>지정된 CWE를 포함하는 모든 문제와 일치합니다. 여러 CWE를 지원합니다.</td></tr><tr><td>CVE</td><td>지정된 CVE를 포함하는 모든 문제와 일치합니다. 여러 CVE를 지원합니다.</td></tr><tr><td>Snyk ID</td><td>지정된 Snyk ID를 포함하는 모든 문제와 일치합니다. 여러 Snyk ID를 지원합니다. 모든 문제에 CVE가 있는 것은 아니므로, 이 방법은 그런 문제들을 지정하는 좋은 방법입니다.</td></tr></tbody></table>
