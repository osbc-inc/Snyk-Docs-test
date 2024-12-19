# Snyk 릴리스 프로세스

Snyk 기능은 다음 유형의 릴리스로 사용자에게 제공됩니다.

{% hint style="info" %}
모든 기능이 모든 단계를 따르는 것은 아니며, 각 기능이 단계를 이동하는 타임라인은 기능에 따라 다양합니다.
{% endhint %}

## 알파

<table><thead><tr><th width="240">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>내부 릴리스만.</td><td>Snyk 내부 사용자, 일부 디자인 파트너</td><td>제한됨</td><td>문서 제공 안 함.</td></tr></tbody></table>

## 폐쇄 베타

<table><thead><tr><th width="243">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>기능을 처음 소비자에게 전개.</td><td>미리 선택된 사용자 그룹</td><td>Snyk 초대에 따름</td><td>제공되지만 공개되지 않을 수 있음.</td></tr></tbody></table>

**Snyk 폐쇄 베타의 기능**

- [PR 확인 구성](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

## 조기 액세스

<table><thead><tr><th width="246">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>기능이 테스트되어 사용할 준비는 되었으나 기본적으로 제공되지는 않음</td><td>Snyk 계정 팀을 통한 요청 또는 Snyk 미리보기를 통한 선택적인 기반의 모든 사용자</td><td>Opt-in: 요청에 따라 Snyk 계정 팀 또는 Snyk 미리보기를 통해</td><td>공개 문서.</td></tr></tbody></table>

**Snyk 조기 액세스의 기능**

- [Snyk GitHub 클라우드 앱](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app.md)
-   프로젝트
    - [자동 생성된 프로젝트 컬렉션](../snyk-admin/introduction-to-snyk-projects/automatically-created-project-collections.md)
- Snyk Code
    - [Snyk Code 사용자 정의 규칙](../scan-with-snyk/snyk-code/snyk-code-custom-rules/)
    - [코드 취약점 자동으로 수정](../scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md)
- 리스크 관리
    - [리스크 점수](../manage-risk/prioritize-issues-for-fixing/risk-score.md)
    - [도달 가능성 분석](../manage-risk/prioritize-issues-for-fixing/reachability-analysis.md)
- Snyk 브로커
    - [Snyk 브로커 커밋 서명](../enterprise-setup/snyk-broker/snyk-broker-commit-signing.md)
    - [Snyk Code - 도커용 브로커와 함께 복제 능력](../enterprise-setup/snyk-broker/git-clone-through-broker.md)
    - [유니버설 브로커](../enterprise-setup/snyk-broker/universal-broker/)
- 언어 지원
    - [.NET 스캔 개선](../supported-languages-package-managers-and-frameworks/.net/improved-.net-scanning.md)
    - [Snyk CLI pnpm 지원](../supported-languages-package-managers-and-frameworks/javascript/javascript-for-open-source.md#pnpm)
    - [개선된 Gradle SCM 스캔](../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning-early-access)
- Snyk 앱 리스크용 타사 통합:
    - [Veracode](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#veracode-setup-guide)
    - [Checkmarx](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#checkmarx-setup-guide)
    - [SonarQube](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#sonarqube-setup-guide)
    - [GitGuardian](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#gitguardian-setup-guide)
    - [Nightfall](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#nightfall-setup-guide)
    - [Dynatrace](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#dynatrace-setup-guide)
    - [Sysdig](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#sysdig-setup-guide)
- [이슈 분석](../manage-risk/enterprise-analytics/issues-analytics.md)
- [애플리케이션 분석](../manage-risk/enterprise-analytics/application-analytics.md)
- [Snyk 런타임 센서](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)

## 일반 공개

<table><thead><tr><th width="249">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>기능이 완전히 활성화됨.</td><td>모든 사용자, 표준 기능 이용 가능.</td><td>기본적으로 이용 가능.</td><td>전체 공개 문서.</td></tr></tbody></table>

## 폐기됨

<table><thead><tr><th width="256">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>기능은 사용 가능하지만 사용은 권장되지 않음.</td><td>현재 사용자만</td><td>기본적으로 이용 가능.</td><td>전체 공개 문서, 적절히 표시됨.</td></tr></tbody></table>

## 유지보수 모드

<table><thead><tr><th width="256">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>기능에 대해 새로운 개발이나 업데이트가 이루어지지 않음.</td><td>현재 사용자만</td><td>기본적으로 이용 가능.</td><td>전체 공개 문서, 적절히 표시됨.</td></tr></tbody></table>

## 지원 종료

<table><thead><tr><th width="256">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>새로운 지원 티켓은 처리되지 않음.</td><td>현재 사용자만</td><td>기본적으로 이용 가능.</td><td>전체 공개 문서, 적절히 표시됨.</td></tr></tbody></table>

## 지원 종료

<table><thead><tr><th width="256">설명</th><th>사용 가능 대상</th><th>접속</th><th>문서</th></tr></thead><tbody><tr><td>해당 기능은 더 이상 사용할 수 없음.</td><td>사용자 없음</td><td>제공되지 않음</td><td>문서 없음</td></tr></tbody></table>