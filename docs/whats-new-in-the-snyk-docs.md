---
cover: .gitbook/assets/Snyk General Banner.webp
coverY: 0
---

# Snyk 문서의 새로운 내용

가장 최근의 업데이트에는 사용자 문서에 중요한 변경 사항이 포함되어 있습니다. 추가 또는 제거된 기능, 관련 정보를 검색하는 방법에 영향을 주는 구조 변경 및 Snyk 지식 베이스와 상호 작용을 향상시키기 위한 기타 개선 사항이 포함되어 있습니다.

## 2024년 11월

### **Snyk AppRisk**

* [자산 대시보드](manage-issues/reporting/available-snyk-reports.md#asset-dashboard)가 재설계되어 보고서 섹션에 포함되었습니다. 새로운 글로벌 필터 바, 새로운 데이터 위젯 및 대시보드를 PDF로 내보내는 옵션 등 여러 개선 사항이 추가되었습니다.
* [Snyk Broker - AppRisk](enterprise-setup/snyk-broker/snyk-broker-apprisk.md#scm-integrations) 문서가 업데이트되어 각 SCM 통합을 구성하는 데 필요한 구체적인 단계를 설명하도록 되었습니다.

### **Snyk Container**

* [Snyk Container에서 지원하는 운영 체제 배포](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md) 목록이 업데이트되어 Ubuntu 24.10 - Oracular Oriole 및 Ubuntu 24.04 - Noble Numbat 04가 포함되었습니다.
* [Snyk Container 작동 방식](scan-with-snyk/snyk-container/how-snyk-container-works/)이 업데이트되어 Snyk가 공개 베이스 이미지 추천을 제공할 때 적용하는 로직에 대한 세부 정보가 추가되었습니다.

### **Snyk 통합**

* Snyk AppRisk 수준 통합의 GitLab 페이지가 업데이트되었습니다. [PAT 생성](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab#gitlab-access-tokens)을 위해 GitLab 그룹을 생성하는 사용자는 최소한의 GitLab 권한 수준인 'Guest' 권한을 가져야 합니다.

### **기타 업데이트**

* Pull Request Checks 섹션이 업데이트되어 새로운 [Pull Request 경험을 위한 PR 체크](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks/pull-request-experience)가 포함되었습니다.
* [지원되는 언어](supported-languages-package-managers-and-frameworks/) 페이지가 재구성되어 각 Snyk 제품에 대한 언어 가용성에 대한 자세한 정보를 제공합니다. 또한, 각 지원 언어에 대한 패키지 관리자, 프레임워크 및 기능 목록을 제공합니다.
* OAuth 2.0을 사용하는 서비스 계정을 [Snyk 웹 UI를 통해 생성할 수 있습니다](enterprise-setup/service-accounts/service-accounts-using-oauth-2.0.md#create-oauth-service-accounts-through-the-ui).
* [API 인덱스](snyk-api/api-endpoints-index-and-tips/)가 업데이트되어 Snyk 사용자 문서에 언급된 각 API 엔드포인트에 대한 항목이 포함되었습니다.
* [개발자 IDE 및 CLI 사용](manage-issues/reporting/available-snyk-reports.md#developer-ide-and-cli-usage) 보고서가 추가 기능인 **개발자 이메일 주소** 및 **PDF 내보내기** 기능으로 강화되었습니다.
* [취약점 세부 정보](manage-issues/reporting/available-snyk-reports.md#vulnerabilities-detail-report) 보고서가 **대상 표시** 및 **열 선택기**와 같은 추가 기능으로 강화되었습니다.

## 2024년 10월

### **Snyk API**

* Snyk 사용자 문서 사이트 전반에 걸쳐 API 링크가 업데이트되었습니다.
* [API](snyk-api/) 페이지가 API 게시 방식의 변화를 반영하기 위해 대대적으로 업데이트되었습니다.
* [API 인증](snyk-api/rest-api/authentication-for-api/) 페이지가 개인 토큰을 사용할 때와 서비스 계정을 사용할 때를 명확히 구분하고, 정보 흐름을 개선하기 위해 업데이트되었습니다.

### **Snyk AppRisk**

* [자산의 위험 요소](manage-assets/#inventory-overview)가 이제 [EA](getting-started/snyk-release-process.md#early-access) 버전으로 Snyk AppRisk Pro 고객에게 제공됩니다.
* [자산 인벤토리 구성 요소](manage-assets/assets-inventory-components.md#clusters)가 클러스터에 대한 세부 정보를 포함하도록 업데이트되었습니다.
* Snyk Broker - AppRisk는 이제 [SonarQube 설치 단계](enterprise-setup/snyk-broker/snyk-broker-apprisk.md#sonarqube-sast-integration)를 포함하도록 업데이트되었습니다.

### **Snyk CLI 및 IDE**

* [CLI 인증 페이지](snyk-cli/authenticate-to-use-the-cli.md)가 OAuth 2.0 프로토콜에 맞게 업데이트되었습니다.
* [Snyk CLI 디버깅](snyk-cli/debugging-the-snyk-cli.md) 페이지가 추가되었습니다.
* [CLI 독립 실행 파일](snyk-cli/install-or-update-the-snyk-cli/#install-with-standalone-executables)이 Alpine Arm64을 포함하도록 업데이트되었습니다.
* IDE Eclipse [플러그인](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/)과 [JetBrains 플러그인](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/) 문서 페이지가 업데이트되었습니다.
* 모든 IDE에 대한 [인증 정보](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)가 업데이트되었습니다.

### **Snyk 통합**

* [Snowflake 데이터 공유](manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/)가 이제 [GA](getting-started/snyk-release-process.md#general-availability) 버전으로 제공됩니다.
* [Snyk SCM 통합](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)이 리포지토리 검색 및 초기 구성 후 권한 또는 범위 수정에 대한 추가 공지가 포함되도록 업데이트되었습니다.
* **GitHub 및 GitHub Enterprise 권한 요구 사항** 섹션 제목이 [SCM을 위한 개인 액세스 토큰(PAT) 범위 및 리포지토리 권한 요구 사항](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#personal-access-token-pat-scopes-and-repository-permission-requirements-for-scms)으로 변경되어 GitHub Cloud App을 포함한 여러 SCM 통합에 대해 PAT 및 리포지토리 권한을 다루고 있습니다.
* GitHub Cloud App이 Fix, Backlog 및 Upgrade PRs의 기능 지원 공지에 추가되었습니다.
* Snyk SCM 통합이 [GitHub Cloud App에 필요한 권한 및 범위](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#github-cloud-app-permission-requirements)를 상세히 설명하는 표를 포함하도록 업데이트되었습니다.

### **기타 업데이트**

* [시작하기](getting-started/) 페이지가 업데이트되어 Snyk 사용 전에 알아야 할 모든 내용을 중앙화했습니다.
* [Snyk Jumpstart 서비스 설명](working-with-snyk/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-services-description.md)과 [고객 전제 조건](working-with-snyk/snyk-terms-of-support-and-services-glossary/snyk-jumpstart-customer-prerequisites.md)이 AppRisk Pro에 맞게 업데이트되었습니다.
* [Dart 및 Flutter](supported-languages-package-managers-and-frameworks/dart-and-flutter.md) 언어에 대한 스캔 방법이 추가되었습니다.

## 2024년 9월

### **Snyk API**

* 문서 사이트의 대부분의 API 링크가 업데이트되었습니다.
* [API 문서 인덱스 페이지](snyk-api/api-endpoints-index-and-tips/)에 추가 항목이 추가되었습니다.
* [API의 지역 URL](snyk-api/rest-api/about-the-rest-api.md#api-url)이 업데이트되었습니다.

### **Snyk AppRisk**

* [GitHub 통합](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-enterprise.md#prerequisites)에서 그룹 수준에 대한 전제 조건 섹션이 추가되었으며, [개인 리포지토리 풀 옵션](scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-enterprise.md#github-integrate-using-snyk-apprisk)에 대한 세부 사항이 추가되었습니다.
* [Snyk AppRisk에 대한 통찰력 설정](manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/) 섹션이 각 통합 옵션에 대한 위험 요소 가용성을 강조하도록 업데이트되었습니다.
* [Snyk 런타임 센서](manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)가 이를 채택하는 것이 가장 효과적인 통합을 달성하고 지속적으로 확장되는 기능 세트를 활용하기 위한 중요성을 반영하도록 업데이트되었습니다.

### Snyk Broker

Universal Broker 기능이 이제 Early Access로 제공됩니다. Universal Broker는 배포와 컨테이너 문제를 연결 문제에서 분리합니다. 이는 여러 유형의 연결을 지원하기 위해 더 작거나 단일 배포를 가능하게 합니다.

### **Snyk CLI**

* [CLI 명령어 및 옵션 요약](snyk-cli/cli-commands-and-options-summary.md)이 업데이트되었습니다.
* [인증](snyk-cli/authenticate-to-use-the-cli.md)이 업데이트되었습니다.
* 구성에 대한 업데이트: Snyk CLI의 환경 변수, [`snyk config`](snyk-cli/commands/config.md) 도움말, [`snyk config environment`](snyk-cli/commands/config-environment.md) 도움말.

### **Snyk 통합**

Snowflake 데이터 공유 섹션이 [데이터 공유 사전 정의](manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/data-share-data-dictionary.md)를 포함하도록 업데이트되어 데이터셋을 탐색하고 구축하는 데 도움이 됩니다.

### **기타 업데이트**

* 업데이트된 [지역 호스팅 및 데이터 거주지](working-with-snyk/regional-hosting-and-data-residency.md) 페이지가 게시되었습니다.
* [용어집](getting-started/glossary.md) 용어가 SCA, SAST, DAST, IAST 및 소프트웨어 구성 분석에 대해 업데이트되었습니다.
* [Early Access](getting-started/snyk-release-process.md#early-access) 출시 상태 알림이 업데이트되었습니다.

## 2024년 8월

### **Snyk API**

* API 참조 문서의 링크가 업데이트되었습니다.
* [API 엔드포인트 인덱스 및 노트](snyk-api/api-endpoints-index-and-tips/)가 업데이트되었습니다.

### **Snyk AppRisk**

* **Manage Risk > Analytics** 페이지가 두 가지 Snyk 제공 사항을 더 잘 반영하도록 통합되었습니다:
  * [문제 분석](manage-risk/enterprise-analytics/issues-analytics.md) - 현재 Snyk Enterprise 고객을 위한 Early Access 버전.
  * [애플리케이션 분석](manage-risk/enterprise-analytics/application-analytics.md) - Snyk AppRisk Pro에 대한 Early Access 버전.
