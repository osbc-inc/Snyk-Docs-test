# Snyk 웹 UI 탐색

Snyk 웹 UI에서는 프로젝트를 관리하고 보안 취약점을 확인하고 해결하며, 종속성을 모니터링하고, 코드의 상태를 확인할 수 있습니다. 또한 계정 설정 구성, API 및 인증 토큰 관리, 애플리케이션 인증, 조직 환경 설정, 그리고 이메일 알림을 사용자 지정할 수 있습니다.

그룹 또는 조직 수준에서 이름을 클릭하여 정보를 시각화할 수 있습니다. 모든 레벨 유형에 대해 보고서, 문제, 종속성, 멤버, 설정, 도움말 및 설정과 같은 일반 정보가 제공됩니다. 자세한 정보는 [그룹 및 조직](../snyk-admin/groups-and-organizations/)을 참조하세요.

## 그룹 수준

다음과 같은 Snyk 기능이 웹 UI에서 그룹 수준에서 사용 가능합니다:

* [선택한 그룹에 대한 사용 가능한 조직](snyk-web-ui.md#organizations-available-for-the-selected-group)
* [자산 대시보드보기](snyk-web-ui.md#view-the-assets-dashboard)
* [자산 재고보기 및 관리](snyk-web-ui.md#view-and-manage-your-assets-inventory)
* [정책 관리 및 사용자 정의](snyk-web-ui.md#manage-and-customize-your-policies)
* [제 3자 공급 업체의 자산 발견, 자산 범위 및 문제 관리를위한 통합 관리](snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors)

다음 동영상에서는 Snyk 웹 UI의 Snyk AppRisk Essentials 개요를 제공합니다.

{% embed url="https://youtu.be/zlR0YqtqDbo" %}
비디오가 마음에 드셨나요? [Snyk 배우기](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 코스의 나머지를 확인해보세요!
{% endembed %}

### 선택한 그룹에 대한 사용 가능한 조직

그룹 수준으로 이동하고 조직 페이지를 선택하면 해당 그룹에서 액세스 권한이 있는 모든 조직과 각 사용 가능한 조직의 조직 역할을 볼 수 있습니다.

### 자산 대시보드보기

{% hint style="info" %}
자산 대시 보드는 Snyk AppRisk Essentials 사용자에게만 제공됩니다. Snyk AppRisk Pro를 사용하는 경우 [애플리케이션 분석](../manage-risk/enterprise-analytics/application-analytics.md)을 참조하세요.
{% endhint %}

Snyk AppRisk 자산 대시 보드 보고 페이지는 응용 프로그램과 관련된 보안 제어에 대한 포괄적인 개요를 제공합니다. 스캔 범위 및 자산 클래스, 소스 등과 관련된 다양한 정보를 카테고리별로 세분화 된 자산을 자세히 파악하는데 필요한 중요한 메트릭 및 데이터를 제시합니다. 또한 대시 보드는 응용 프로그램과 관련된 컨텍스트 데이터를 사용하여 특정 응용 프로그램 및 소유자를 기준으로 결과를 필터링할 수있는 확장된 전역 필터 옵션을 포함합니다.

자세한 내용은 [자산 대시 보드](../manage-issues/reporting/available-snyk-reports.md#asset-dashboard) 문서 섹션을 참조하십시오.

### 자산 재고보기 및 관리

{% hint style="info" %}
재고는 Snyk AppRisk 사용자에게만 제공됩니다.
{% endhint %}

[재고](../manage-assets/)는 Snyk AppRisk를 사용하는 경우에만 제공됩니다. **재고** 페이지를 사용하여 저장소 자산을 구성하여 SCM 도구에서 모든 저장소 자산을 시각화하고, 제품 제어 범위를 추적하고, 비즈니스 영향에 따라 범위 완화를 우선순위로 정할 수 있습니다.

재고의 각 행은 저장소 자산 또는 식별 정보가 부족한 저장소와 유사할 수 있는 Snyk에서 스캔 된 품목을 나타냅니다. 정책을 통해 스캔 된 품목은 지원되지 않습니다.

### 정책 관리 및 사용자 정의

{% hint style="info" %}
정책은 Snyk AppRisk 사용자에게만 제공됩니다.
{% endhint %}

업무 컨텍스트를 자동화 및 알림 수신하기 위한 정보는 [정책](../manage-risk/policies/assets-policies/)을 참조하십시오.

### 제 3자 공급 업체의 자산 발견, 자산 범위 및 문제 관리를 위한 통합 관리

{% hint style="info" %}
그룹 수준의 통합은 Snyk AppRisk 사용자에게만 제공됩니다.
{% endhint %}

**통합** 페이지에는 존재하는 모든 통합과, <SCM](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 또는, [제 3자](../integrate-with-snyk/#integrations-for-snyk-apprisk)의 통합과 관련된 모든 데이터가 자동으로 동기화된 Snyk 조직, 통합 허브에 대한 액세스가 제공되며, SCM 통합을 추가하고, 제 3자 통합을 연결하고, SCM 통합에 애플리케이션 컨텍스트를 추가하거나 Snyk 런타임 센서를 사용할 수 있습니다.

모든 통합에 대한 개요를 Snyk 웹 UI **통합** 페이지에서 찾을 수 있습니다. 통합을 활성화 또는 비활성화하고, 수정하거나 구성에서 삭제할 수 있습니다.

사용 가능한 통합에 대한 자세한 내용은 [Snyk AppRisk SCM 통합](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 및 [Snyk와 통합](../integrate-with-snyk/#integrations-for-snyk-apprisk)을 참조하십시오.

#### 통합 활성화 또는 비활성화

통합을 연결하거나 일시 중단할 수 있습니다. 통합을 활성화하거나 비활성화하려면 재생 또는 일시 중단을 클릭하세요.

#### 통합을 위한 새 프로필 추가

각 통합은 하나 이상의 프로필에서 실행할 수 있습니다. 동일한 소스 내의 여러 인스턴스에서 데이터를 검색할 때 유용합니다.

새 프로필 추가:

1. 이미 사용 가능한 통합 프로필에서 **설정** 아이콘을 클릭합니다.
2. **프로필 추가**를 클릭합니다.
3. 구성 필드를 작성하고 **완료**를 클릭합니다.

#### 통합 편집

기존 통합을 편집하려면 통합 카드에서 **설정을 클릭**한 다음 해당 통합에 대해 추가한 조직에서 다시 **설정**을 클릭합니다.

<figure><img src="../../assets/image.png" alt=""><figcaption><p>통합 허브에서 기존 통합 편집</p></figcaption></figure>

{% hint style="info" %}
보안상의 이유로 기존 통합의 설정을 열 때 모든 자격 증명이 익명화됩니다.
{% endhint %}

#### 통합 제거

환경에서 기존 통합을 제거하려면 통합을 선택하고 **삭제**를 클릭합니다.

{% hint style="info" %}
이미 삭제된 통합은 복원할 수 없습니다. 다시 활성화해야 합니다.
{% endhint %}

## 조직 수준

다음과 같은 Snyk 기능이 웹 UI에서 조직 수준에서 사용 가능합니다:

* [대시 보드 탐색](snyk-web-ui.md#explore-the-dashboard)
* [프로젝트 관리](snyk-web-ui.md#manage-your-projects)
* [보고서 보기](snyk-web-ui.md#view-reports)
* [종속성 및 라이선스 보기](snyk-web-ui.md#view-dependencies-and-licenses)
* [통합 관리](snyk-web-ui.md#manage-your-integrations)
* [문제보기 및 우선 순위 설정](snyk-web-ui.md#view-and-prioritize-issues)
* [조직 또는 그룹 멤버 관리](snyk-web-ui.md#manage-organization-or-group-members)
* [Snyk 조직 및 그룹 설정 설정](snyk-web-ui.md#snyk-organization-or-group-settings)
* [도움이 될 자료 보기](snyk-web-ui.md#view-helpful-resources)
* [계정 설정 및 선호 설정 관리](snyk-web-ui.md#manage-account-preferences-and-settings)

{% hint style="info" %}
또한 [Snyk CLI](../snyk-cli/), [IDE 내](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/) 및 [Snyk API](../snyk-api/)를 통해 Snyk 기능을 사용할 수 있습니다.
{% endhint %}

### 대시 보드 탐색

기존 계정으로 로그인하고 조직을 선택하면 웹 UI에서 해당 조직의 대시 보드가 열립니다. 최상위 대기 중인 작업 및 취약한 프로젝트를 볼 수 있으며 새 프로젝트를 추가할 수 있습니다.

#### 최상위 대기 중인 작업

**대기 중인 작업** 섹션에는 Snyk 조직의 프로젝트에서 처리해야 할 다음 할 일이 표시됩니다.

#### 프로젝트 보기

대시 보드에서 프로젝트에 대한 링크를 사용하여 프로젝트의 메타데이터, 다시 테스트 및 수정 옵션을 탐색하고 관리할 수 있습니다. 각 링크는 프로젝트 세부 정보 페이지를 열어줍니다. 여기서 프로젝트 **개요**를 볼 수 있거나 **기록** 및 **설정** 탭으로 전환할 수 있습니다.

자세한 정보는 [Snyk 프로젝트](../snyk-admin/snyk-projects/)를 참조하십시오.

#### 취약점 수정

Snyk는 가장 취약한 프로젝트에서 PR (Pull Requests)를 추적하고 표시하며 다음을 포함합니다:

* 가장 취약한 프로젝트에서 취약성을 해결할 수 있는 PRs
* 이미 발생한 PRs 또는 Snyk를 통해 발생되었고 검토를 기다리는 오픈 PRs

**취약성 해결** 링크가있는 프로젝트의 경우 링크를 사용하여 수정 PR을 열 수 있는 프로젝트 세부 정보를 볼 수 있습니다. 자세한 내용은 [Snyk PR 또는 병합 요청 수정](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/)을 참조하십시오.

{% hint style="info" %}
Snyk는 GitHub, GitHub Enterprise 및 Bitbucket Cloud에서만 PR을 추적하고 표시하며 최상위 취약한 프로젝트에 대해서만 이를 지원합니다. 다른 SCM을 사용하는 경우 **대기 중인 작업** 섹션에 제기 될 수 있는 PR은 나타나지만 이미 제기된 PR은 나타나지 않습니다.
{% endhint %}

#### 최상위 취약한 프로젝트

또한 **최상위 취약 프로젝트** 섹션에서 최취약으로 판명 된 Snyk 프로젝트를 확인할 수 있고 **대기 중인 작업** 섹션과 유사한 기능이 제공됩니다.

#### 프로젝트 추가

Snyk 프로젝트를 추가하려면 대시보드에서 **프로젝트 추가** 링크를 클릭하여 프로젝트 추가 방법을 선택하십시오. 자세한 내용은 [프로젝트 가져오기](./#import-a-project-to-scan-and-identify-issues)를 확인하세요.

### **프로젝트 관리**

사이드 메뉴에서 **프로젝트** 링크를 선택하여 **프로젝트** 목록 페이지를 엽니다. 이 페이지에서 여러 작업을 수행할 수 있습니다:

* 프로젝트 추가. **프로젝트 추가** 드롭다운에서 프로젝트 추가 방법을 선택합니다.
* 프로젝트 필터링, 그룹화 및 정렬.
* 프로젝트에 대한 팁 및 최신 가져오기 기록 보기.
* 각 프로젝트 링크를 선택하여 프로젝트 세부 정보 페이지를 열어 요약 및 문제 정보를 볼 수 있습니다.
* 미분류된 **프로젝트** 목록에서 플러스 아이콘을 사용하거나 **대상**으로 프로젝트 분류 시 타겟에서 다른 대상에 대한 프로젝트를 그룹화 하도록 타겟 추가.
* 프로젝트 세부 정보 페이지에서 **셋팅** 탭 또는 **기록** 탭을 사용하여 알림, 프로젝트 테스트 및 PR(풀 리퀘스트) 빈도에 대한 일반 및 GitHub 통합 설정을 구성합니다. **셋팅** 탭에서 고유 프로젝트 ID를 검색하고 프로젝트를 비활성화 또는 삭제합니다.
* **역사** 탭에서 프로젝트 이력 보기.

### **통합 관리**

대시보드에서 사용 가능한 [통합](../integrate-with-snyk/) 페이지를 통해 Snyk와의 다양한 통합을 설정할 수 있습니다.

## 모든 수준 유형에 대한 일반 설정

다음과 같은 Snyk 기능이 웹 UI에서 모든 수준 유형에 사용 가능합니다:

* [보고서 보기](snyk-web-ui.md#view-reports)
* [## 계정

- 이메일 **알림**을 위한 계정 설정 관리 (왼쪽 내비게이션에 링크)로서, 이슈 이메일 알림, 주간 보고서 이메일, 사용량 경고뿐만 아니라 보고서가 이용 가능하게 되었을 때의 이메일 알림 그리고 판매 및 마케팅 통신에 대한 환경 설정도 관리합니다. 자세한 내용은 [알림 관리](../snyk-admin/manage-notifications.md)를 참조하세요.
- **친구와 공유**를 위한 추천 링크를 받아보세요. 이 링크는 계정 설정의 왼쪽 내비게이션에 있습니다.