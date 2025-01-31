# Snyk 도구

## Snyk 도구의 범위

Snyk 도구는 "통증 포인트"를 해결하지 못하는 특정 상황에 도움을 줍니다. 이는 웹 UI, CLI, API 또는 통합을 통해 Snyk을 사용하더라도 해당됩니다. Snyk 도구는 Snyk API와 CLI의 기능을 확장합니다.

{% hint style="info" %}
Snyk 도구를 사용하려면 [Snyk 계정](https://snyk.io/login?cta=sign-up\&loc=nav\&page=support_docs_page)이 있고 Projects가 채워져 있어야 합니다.
{% endhint %}

## 주요 Snyk 도구

다음과 같은 핵심 Snyk 도구에 대한 전체 설명서를 제공합니다:

* [snyk-api-import (docs)](tool-snyk-api-import/): 강력하고 안정적인 방법으로 Projects를 대량으로 Snyk에 가져올 수 있습니다. Repo: [snyk-api-import](https://github.com/snyk/snyk-api-import)
* [snyk-delta (docs)](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-delta.md): 두 Snyk 스냅샷 간의 차이를 얻습니다. Repo: [snyk-delta](https://github.com/snyk-tech-services/snyk-delta)
* [snyk-filter (docs)](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md): Snyk CLI의 JSON 출력을 가져와 결과에 사용자 지정 필터를 적용합니다. Repo: [snyk-filter](https://github.com/snyk-tech-services/snyk-filter)
* [snyk-to-html](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md): Snyk JSON을 HTML로 변환하여 `snyk test --json`의 JSON 출력을 가져와 발견된 취약점을 표시해주는 로컬 HTML 파일을 생성합니다. Repo: [snyk-to-html](https://github.com/snyk/snyk-to-html)
* [jira-tickets-for-new-vulns (docs)](tool-jira-tickets-for-new-vulns.md): Snyk에서 모니터링하는 Projects를 동기화하고 문제에 대한 JIRA 티켓을 자동으로 엽니다. Repo: [jira-tickets-for-new-vulns](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns)
* [snyk-scm-contributors-count (docs)](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-scm-contributors-count/): 마지막 90일 동안 커밋이 있는 SCM 저장소의 기여자 수를 계산합니다. Repo: [snyk-scm-contributors-count](https://github.com/snyk-tech-services/snyk-scm-contributors-count).

## 추가 Snyk 도구

다음 추가 Snyk 도구를 사용하는 방법에 대한 지침은 해당 저장소를 참조하십시오.

* [snyk-disallow](https://github.com/snyk-tech-services/snyk-disallow): Snyk 그룹의 뷰어 토큰을 받아 이를 CI 또는 유사한 시스템을 위한 읽기/테스트 전용 토큰으로 얻을 수 있습니다.
* [snyk-prevent-gh-commit-status](https://github.com/snyk-tech-services/snyk-prevent-gh-commit-status): CI에서 실행된 [snyk-delta](https://github.com/snyk-tech-services/snyk-delta)의 결과를 PR의 커밋 상태로 전송합니다.
* [snyk-cr-monitor](https://github.com/snyk-tech-services/snyk-cr-monitor): 테스트할 Docker 저장소를 수집한 다음 결과를 동시에 여러 작업을 실행하기 위해 반복합니다.
* [backstage-plugin-snyk](https://github.com/snyk-tech-services/backstage-plugin-snyk): Snyk에서 보안 세부 정보를 표시하는 플러그인.
* [snyk-api-ts-client](https://github.com/snyk-tech-services/snyk-api-ts-client): Snyk API Typescript 클라이언트.
* [snyk-transitive-ignore](https://github.com/snyk-tech-services/snyk-transitive-ignore): 제공된 패키지 목록을 기반으로 동적으로 Snyk 무시 정책을 생성합니다.
* [snyk-user-sync-tool](https://github.com/snyk-tech-services/snyk-user-sync-tool): 사용자 멤버십을 추가, 제거 및 동기화합니다.
* [snyk-licenses-texts](https://github.com/snyk-tech-services/snyk-licenses-texts): 사용된 조직 수준의 라이센스, 저작권(2024년 1월 8일까지), 및 의존성 데이터를 제공합니다.
* [snyk-request-manager](https://github.com/snyk-tech-services/snyk-request-manager): Snyk API와 상호 작용하기 위한 요청 관리자입니다.
* [snyk-repo-issue-tracker](https://github.com/snyk-tech-services/snyk-repo-issue-tracker): Snyk Project issues API와 실행 간의 문제 변경 내용 집합을 생성할 수 있는 파이썬 스크립트/모듈입니다.
*   [snyk-repo-diff:](https://github.com/snyk-tech-services/snyk-repo-diff) Snyk에 의해 모니터링되지 않는 저장소를 판별하는 데 도움을 줍니다.

    이 도구는 특정 Snyk 그룹 내의 모든 조직에 속한 모든 프로젝트 목록을 불러오고, 주어진 GitHub 조직에서 발견된 저장소 목록을 연결하는 방식으로 작동합니다.
* [snyk-issues-to-csv](https://github.com/snyk-tech-services/snyk-issues-to-csv): PySnyk 모듈과 Pandas 모듈을 사용하여 보고서 API에서 모든 문제를 수집하고 전체 그룹의 단일 CSV 파일로 결합하는 파이썬 스크립트입니다.
* [snyk-bulk](https://github.com/snyk-tech-services/snyk-bulk): 빌드 환경 외부에서 Snyk CLI를 사용하여 오픈 소스 취약점을 재귀적으로 스캔합니다.
* [snyk-bulk-action-scripts](https://github.com/snyk-tech-services/snyk-bulk-action-scripts): Snyk 그룹 내의 모든 조직에 대한 통합 설정을 편집하기 위한 스크립트 모음.
* [snyk-deps-to-csv](https://github.com/snyk-tech-services/snyk-deps-to-csv): 그룹 내의 모든 조직에서 모든 종속성을 수집하여 CSV 파일로 출력합니다.
* [add-ignore-reason-to-csv-report](https://github.com/snyk-labs/add-ignore-reason-to-csv-report)**:** 사용자 인터페이스에서 CSV 무시 보고서에 무시 이유와 사용자 데이터를 추가합니다.

## 도구 아이디어

도구에 대한 아이디어가 있으신가요? 있다면 [Snyk Apps](../../snyk-api/how-to-use-snyk-apps-apis/)를 확인해보세요. 이를 통해 Snyk 경험을 사용자의 특정 요구에 맞게 조정할 수 있습니다. 질문이 있으면 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.
