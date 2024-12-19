---
cover: .gitbook/assets/Snyk General Banner.webp
coverY: 0

# Snyk 문서의 새로운 내용

가장 최근의 업데이트에는 사용자 문서에 중요한 변경 사항이 포함되어 있습니다. 추가 또는 제거된 기능, 관련 정보를 검색하는 방법에 영향을주는 구조 변경 및 Snyk 지식 베이스와 상호 작용을 향상시키기 위한 기타 개선 사항이 포함되어 있습니다.

## 2024년 11월

### **Snyk AppRisk**

- [자산 대시보드](manage-issues/reporting/available-snyk-reports.md#asset-dashboard)가 재설계되어 보고서 섹션에 포함되었습니다. 새로운 글로벌 필터 바, 새로운 데이터 위젯 및 대시보드를 PDF로 내보내는 옵션 등 여러 개선 사항이 추가되었습니다.
- [Snyk Broker - AppRisk](enterprise-setup/snyk-broker/snyk-broker-apprisk.md#scm-integrations) 문서가 업데이트되어 각 SCM 통합을 구성하는 데 필요한 구체적인 단계를 설명하도록 되었습니다.

### **{{Snyk Container}}**

- [{Snyk Container에서 지원하는 운영 체제 배포](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) 목록이 업데이트되어 Ubuntu 24.10 - Oracular Oriole 및 Ubuntu 24.04 - Noble Numbat 04가 포함되었습니다.
- [{Snyk Container 작동 방식](scan-with-snyk/snyk-container/how-snyk-container-works/)이 업데이트되어 Snyk가 공개 베이스 이미지 추천을 제공할 때 적용하는 로직에 대한 세부 정보가 추가되었습니다.

### **Snyk 통합**

- GitLab 페이지가 업데이트되어 Snyk AppRisk 레벨 통합 [PAT 생성](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab#gitlab-access-tokens)에 대해 GitLab 그룹을 생성하는 사용자가 최소한의 Guest GitLab 권한 수준이 필요함을 명시합니다.

### **기타 업데이트**

- PR Checks 섹션이 업데이트되어 새로운 [PR Checks용 PR 경험](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience)이 포함되었습니다.
- [지원되는 언어](supported-languages-package-managers-and-frameworks/) 페이지는 각 Snyk 제품에 대한 언어 가용성에 대한 자세한 정보를 제공하기 위해 재구성되었습니다. 또한, 각 지원되는 언어에 대한 패키지 관리자, 프레임워크 및 기능 목록을 제공합니다.
- OAuth 2.0을 사용하는 서비스 계정은 이제 [Snyk 웹 UI를 통해](enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md#create-oauth-service-accounts-through-the-ui) 생성할 수 있습니다.
- [API 색인](snyk-api/api-endpoints-index-and-tips/)에는 Snyk 사용자 문서에 언급된 각 엔드포인트에 대한 항목이 포함되어 있습니다.
- [개발자 IDE 및 CLI 사용](manage-issues/reporting/available-snyk-reports.md#developer-ide-and-cli-usage) 보고서에는 추가 기능인 **개발자 이메일 주소** 및 **PDF 내보내기**가 추가되었습니다.
- [취약점 자세한 내용](manage-issues/reporting/available-snyk-reports.md#vulnerabilities-detail-report) 보고서는 **대상 지시** 및 **열 선택기**와 같은 추가 기능이 추가되었습니다.

## 2024년 10월

### **Snyk API**

...
...생략...
...

## 2024년 8월

### **Snyk API**

...
...생략...
...

## 2024년 7월

### **Snyk API**

...
...생략...
...

### Snyk AppRisk

...
...생략...
...

### **Snyk CLI**

...
...생략...
...

### Snyk IDE

...
...생략...
...

### **Snyk 통합**

...
...생략...
...

### **기타 업데이트**

...
...생략...
...## 자산 관리

* 문제 - [수정 필요 사항 우선 순위 설정](manage-risk/prioritize-issues-for-fixing/#prioritization-based-on-risk) 섹션
* 정책 - [자산 정책](manage-risk/policies/assets-policies/) 섹션
* 통합 - [Snyk SCM 통합](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 및 [Snyk와 통합](integrate-with-snyk/#integrations-for-snyk-apprisk) 섹션
* Snyk AppRisk - [Snyk로 스캔하기](scan-with-snyk/snyk-apprisk/) 섹션
* 분석 - [위험 관리](manage-risk/enterprise-analytics/application-analytics.md) 섹션

### Snyk 통합

* GitHub 및 GitHub Enterprise 통합 기능 비교는 현재 [SCM, IDE 및 CI/CD 통합](scm-ide-and-ci-cd-integrations/#github-vs-github-enterprise) 페이지에 있습니다.
* GitHub 통합에서 GitHub Enterprise 통합으로 마이그레이션하는 단계는 이제 GitHub 통합 페이지에 있습니다. ([마이그레이션 단계](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#migrate-to-the-github-enterprise-integration))
* [Snyk SCM 통합](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations) 페이지에는 SDLC에서 성공적으로 이러한 통합을 사용하는 데 중요한 정보가 포함되어 있습니다. 이에는 다음이 포함됩니다:
  * Snyk AppRisk 기능을 따르는 [그룹](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 및 [조직](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#organization-level-snyk-scm-integrations) 수준으로 통합 조직화
  * [Git 저장소 복제](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-git-repository-cloning)에 대한 세부 정보
  * 엔터프라이즈 고객을 위한 [배포 권장 사항](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#deployment-order-recommendations)
  * 각 SCM 통합의 사용자 권한 및 액세스 범위 요구 사항
  * [Snyk Broker를 위한 통합 SCM 토큰 생성 방법](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#integrated-scm-tokens-for-snyk-broker)에 대한 지침

### **기타 업데이트**

* **Snyk 보고서:** [문제 열 사전](manage-risk/reporting/issue-columns-dictionary.md#issue-vulnerability-details)에는 Jira (JIRA ISSUES LIST, LATEST JIRA ISSUE) 및 EPSS (EPSS SCORE, EPSS PERCENTILE)에 대한 새로운 필터 및 열이 포함되어 있습니다. 이를 통해 Jira에서 작업을 관리하고 우선 순위 지정 단계에 EPSS를 포함할 수 있습니다.
* **Snyk 보안:** Snyk는 새로운 취약점에 대한 기본 평가로 [CVSS V4.0](manage-risk/prioritize-issues-for-fixing/severity-levels.md#severity-levels-and-cvss)를 채택함으로써 우선 순위 및 위험 평가 워크플로를 개선했습니다.
* **코드 취약점 자동으로 수정:** [DeepCode AI Fix](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#deepcode-ai-fix-language-support)은 이제 AWS 환경 및 JetBrains IDE에서 사용할 수 있습니다. AWS 멀티 테넌트 환경을 사용하는 경우, Snyk 미리보기 [Snyk 코드] 수정 제안](scan-with-snyk/snyk-code/manage-code-vulnerabilities/fix-code-vulnerabilities-automatically.md#enable-deepcode-ai-fix)를 켜고 IDE에서 Snyk로 다시 테스트하세요.