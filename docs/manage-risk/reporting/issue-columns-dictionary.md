# 이슈 열 사전

Snyk 보고서에는 수십 가지 필터 및 열(columns)이 포함되어 있어 사용자가 데이터의 정교한 표시를 개발하고 쉽게 필요한 통찰을 얻을 수 있도록 합니다. 따라서 정확한 결론을 도출하려면 사용된 열(columns)과 필터의 의미를 이해해야 합니다. 이 사전은 Snyk 이슈 상세 보고서의 이슈 열(columns) 뒤에 숨겨진 의미를 설명합니다.

## 이슈 특성 <a href="#issue-characteristics" id="issue-characteristics"></a>

이슈의 주요 특성에 대해 설명합니다.

* **COMPUTED FIXABILITY** - 취약성 개선 경로를 기반으로 이슈를 해결할 수 있는지 나타냅니다. 자세한 내용은 [Computed Fixability 필터](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/vulnerability-fix-types.md#computed-fixability-filters)를 참조하십시오.
  * **Fixable:** 모든 식별된 이슈에 대한 해결책이 있으며 모든 상세 경로에 개선이 있음을 의미합니다.
  * **Partially fixable:** 이슈는 업그레이드 가능한 경로를 가지고 있지만 모든 상세 경로에 개선이 없습니다.
  * **No supported fix**: 이슈에는 업그레이드 가능한 경로가 없습니다.
* **INTRODUCTION CATEGORY** - 이슈가 어떻게 소개되었는지를 분류합니다:
  * **Baseline Issue** - 프로젝트가 모니터링되기 시작한 직후 식별된 이슈입니다.
  * **Preventable Issue** - Snyk에서 관련 문제를 감지하기 전 최소 일곱 일 동안 발생한 이슈입니다.
  * **Non Preventable Issue** - 새로운 취약성 게시 때문 등 외부 요인으로 인해 발생된 이슈입니다.
  * **Other New Issue** - Snyk이 그들의 예방 가능성을 분류 할 수 없는 이슈들입니다. 자세한 내용은 [위험이 소개되는 방법의 구분](../enterprise-analytics/issues-analytics.md#delineation-of-how-risk-is-introduced)을 참조하십시오.
* **ISSUE** - 다음을 포함하는:
  * **Problem Title**: Snyk 취약점 이름.
  * **Issue Type:** 이슈가 취약점, 라이선스 또는 구성과 관련이 있는지를 나타냅니다.
* **ISSUE STATUS** - 이슈가 개방, 해결되었거나 무시된 상태인지 나타냅니다.
* **PRODUCT NAME** - Snyk 제품 이름.
* **SEVERITY** - 특정 Snyk 제품의 분석에 따른 이슈 심각성을 나타냅니다.
* **점수** - 분석 모델을 기반으로 한 점수입니다. Priority score는 일반적으로 사용 가능하며 Risk Score는 초기 액세스 상태입니다. 자세한 내용은 [Priority Score vs Risk Score](../prioritize-issues-for-fixing/priority-score-vs-risk-score.md)를 참조하십시오.
* **REACHABILITY** - 이슈의 접근성은 애플리케이션에서 호출되는 기능과 관련된지 여부에 따라이슈의 폭로 가능성이 더 큽니다.\
  허용되는 값:
  * **Reachable** - 애플리케이션에서 취약한 코드로의 직접 또는 간접 경로가 발견되었습니다.
  * **No path found** - 애플리케이션에서 취약한 코드로의 경로를 찾을 수 없습니다.
  * **Not applicable** - 특정 이슈에 대해 접근성이 관련이 없습니다.

## 이슈 취약점 세부정보 <a href="#issue-vulnerability-details" id="issue-vulnerability-details"></a>

취약점 세부 정보는 Snyk, Mitre, NVD 또는 기타 신뢰할 수 있는 보안 기관에서 정의한 다양한 이슈 속성을 참조합니다.

* **ATTACK VECTOR -** 취약성이 공격되기 가능한 맥락을 나타냅니다. 공격 벡터 및 그 값(Network, Adjacent, Local, Physical)에 대한 자세한 내용은 [사양 문서](https://www.first.org/cvss/specification-document)를 참조하십시오.
* **CVE** - Mitre CVE ID
* **CWE** - Mitre CWE ID
* **EPSS SCORE** - 앞으로 30일 동안 야생에서의 공격 발생 가능성입니다.
* **EPSS PERCENTILE** - 동일하거나 낮은 EPSS 점수를 가진 모든 취약성의 비율입니다.
* **EXPLOIT MATURITY** - Snyk에서 확인한 공개 취약점의 존재 및 성숙도를 나타냅니다. 자세한 내용은 [프로젝트에서 exploits 보기](../prioritize-issues-for-fixing/view-exploits.md#view-exploits-in-projects)를 참조하십시오. 허용되는 값은 다음과 같습니다:
  * **Mature:** 이 취약점에 대한 Snyk의 공개 코드 exploit이 있습니다.
  * **Proof of concept:** Snyk이 이러한 취약성을 어떻게 exploit 하는지에 대한 concept 또는 자세한 설명이 있습니다. Proof of concept 취약점 패치는 비활성화되지 않으며 찾은 곳에서 fix PR에 나타날 것입니다.
  * **No known exploit:** Snyk이 이 취약성에 대한 concept 또는 공개 exploit을 찾지 못했습니다.
  * **No data**: 이슈는 취약점이 아니라 라이선스 이슈 또는 취약점 알림입니다.
* **FIXED IN AVAILABLE** - 취약성에 대한 해결책을 포함한 새 패키지 버전이 있는지 나타냅니다.
* **FIXED IN VERSION** - 취약성에 대한 해결책이 제공된 패키지 버전을 나타냅니다.
* **HAS JIRA ISSUE(S) ASSIGNED** - 적어도 하나의 Jira 이슈가 할당되어 있을 때 `true`를 표시하고, 그렇지 않으면 `false`를 표시합니다.
* **JIRA ISSUES LIST** - 모든 첨부된 Jira 이슈 키의 목록입니다.
* **LATEST JIRA ISSUE** - 프로젝트 페이지의 이슈 카드에 링크된 최신 첨부된 Jira 이슈 키입니다.
* **NVD SCORE** - NVD에서 계산한 취약성 점수입니다.
* **NVD SEVERITY** - NVD에서 평가한 취약성 심각도입니다.
* **PACKAGE NAME AND VERSION** - 관련된 패키지 이름과 버전입니다.
* **PROBLEM ID** - 취약성을 고유하게 식별하는 Snyk Vuln DB ID입니다.
* **PROBLEM TITLE** - Snyk이 설명하는 취약성 이름입니다.
* **SEMVER VULNERABLE RANGE** - 패키지 버전의 취약 범위(semantic versioning을 기반으로 함)입니다.
* **SNYK CVSS SCORE -** Snyk의 CVSS (공통 취약점 평가 시스템) 점수입니다.
* **VULNERABILITY PUBLICATION DATE** - 취약점이 Snyk에서 발행된 날짜를 나타냅니다.

## 이슈 컨텍스트 열(columns) <a href="#issue-context-columns" id="issue-context-columns"></a>

컨텍스트 열(columns)은 이슈의 영향과 위험을 연관된 프로젝트, 대상, 조직 등과 함께 이해하는 데 도움을 줍니다.

* **ORG NAME** - Snyk 조직 이름입니다.
* **PROJECT COLLECTION** - 프로젝트 컬렉션은 프로젝트의 정적 컬렉션입니다.
* **PROJECT CRITICALITY** - 프로젝트의 비즈니스 중요도. 자세한 내용은 [프로젝트 속성](../../snyk-admin/snyk-projects/project-attributes.md)을 참조하십시오.
* **PROJECT ENVIRONMENT** - 프로젝트의 환경입니다. 자세한 내용은 [프로젝트 속성](../../snyk-admin/snyk-projects/project-attributes.md)을 참조하십시오.
* **PROJECT LIFECYCLE** - 프로젝트의 수명주기입니다. 자세한 내용은 [프로젝트 속성](../../snyk-admin/snyk-projects/project-attributes.md)을 참조하십시오.
* **PROJECT NAME** - 프로젝트 이름입니다.
* **PROJECT ORIGIN** - 프로젝트의 통합 소스입니다; 원래 SCM, 컨테이너 레지스트리 이름 등이 될 수 있습니다.
* **PROJECT OWNER** - 프로젝트 소유자로 정의된 사용자입니다.
* **PROJECT TAGS** - 프로젝트와 관련된 태그입니다. 자세한 내용은 [프로젝트 태그](../../snyk-admin/introduction-to-snyk-projects/project-tags.md)를 참조하십시오.
* **PROJECT TARGET** - 대상 이름입니다.
* **PROJECT TARGET REFERENCE** - 이 프로젝트를 구별하는 참조를 지정합니다. 예를 들어 브랜치 이름이나 버전일 수 있습니다. 자세한 내용은 [그룹 프로젝트를 모니터링하기 위해 브랜치나 버전으로 그룹화](../../snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring.md)를 참조하십시오.
* **PROJECT TYPE** - 프로젝트의 패키지 관리자입니다.

## 이슈 날짜 열(columns) <a href="#issue-date-columns" id="issue-date-columns"></a>

날짜 열(columns)은 다양한 계산 및 트렌드 분석에 사용되는 중요한 이벤트를 나타냅니다.

* **INTRODUCED** - 이 문제를 식별한 첫 번째 스캔의 타임스탬프입니다.
* **LAST IGNORED** - 이 문제가 마지막으로 무시된 날짜를 나타내는 타임스탬프입니다.
* **LAST INTRODUCED** - 이 문제를 식별한 최근 스캔의 타임스탬프입니다.
* **LAST RESOLVED** - 이 문제가 테스트에서 발견되지 않은 마지막 타임스탬프입니다.
