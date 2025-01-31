# Snyk AppRisk 사용

## 사전 요구 사항

시작하기 전에 다음 사전 요구 사항을 충족하는지 확인하십시오:

* 당신은 Snyk 엔터프라이즈 고객이어야 합니다.
* 당신의 계정은 Snyk AppRisk Essentials 또는 Snyk AppRisk Pro에 대한 액세스 권한이 있어야 합니다.
* Snyk AppRisk와 관련된 그룹의 그룹 관리자이거나 View Group 및 Edit AppRisk 권한이 있는 그룹 수준 역할을 할당 받았어야 합니다.
* Snyk AppRisk와 관련된 그룹은 Snyk 애플리케이션 보안 제품을 도입한 조직을 포함해야 합니다.
* 클라우드 기반 SCM 도구(Azure DevOps, GitHub, GitLab 등)를 Snyk AppRisk에 저장소 자산 발견을 위해 도입할 권한과 권한이 있어야 합니다.

{% hint style="info" %}
Git 코드 저장소를 Snyk AppRisk에 통합할 때, Snyk에 가져온 것뿐만 아니라 코드 저장소의 넓고 완전한 전망을 갖기 위해 보조 토큰을 사용해야 합니다.\
보조 토큰을 사용하면 Snyk을 사용하여 도입된 모든 것에 대한 반대의 의견을 갖게 됩니다.\
보조 토큰을 사용하면 조직 수준 설정에서 제한된 토큰에서 블라인드 스팟을 도입할 가능성이 줄어듭니다.\
첫 번째 가져오기, 동기화는 최대 24시간이 소요될 수 있습니다.
{% endhint %}

다음 비디오는 Snyk AppRisk가 제공하는 내용을 개요로 제시합니다:

{% embed url="https://www.youtube.com/embed/23y6pEnTfqQ" %}

## 권한

다음에 설명된 그룹 수준 역할 권한 중 하나로 Snyk AppRisk에 액세스할 수 있습니다. 권한에 액세스하려면 **View groups**로 이동한 다음 **Snyk AppRisk permissions** 옵션을 선택하십시오.

1. View AppRisk - AppRisk에 대한 읽기 전용 액세스를 부여합니다.
2. Edit AppRisk - AppRisk를 편집하는 액세스를 부여합니다. 예를 들어, 정책 편집, 자산 분류 편집, 통합 추가 등입니다.

그룹 관리자는 기본적으로 **Edit AppRisk** 권한을 할당 받고, 그룹 뷰어는 기본적으로 **View AppRisk** 권한을 할당 받습니다.

{% hint style="info" %}
기본 사용자 역할 및 권한에 대한 자세한 정보는 [기본 사용자 역할](../../snyk-admin/user-roles/pre-defined-roles.md)을 참조하십시오.
{% endhint %}

## 로그인 및 인증

기존 메커니즘 (SSO, Google SAML 등)을 사용하여 Snyk에 로그인하고 인증하십시오.

## Snyk AppRisk 액세스

Snyk AppRisk 옵션에 액세스하려면 그룹 수준에 있는지 확인하십시오. 그룹 수준에서는 프로젝트의 보안을 강화하고 보안 절차를 간소화하는 중앙 집중식 보안 관리가 제공됩니다.

다음 동영상은 Snyk AppRisk 인터페이스의 개요를 제공합니다:

{% embed url="https://www.youtube.com/watch?v=OUoukT5518c" %}
비디오가 마음에 드시나요? [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 과정의 나머지 부분을 확인해보세요!
{% endembed %}

다음은 Snyk AppRisk의 Snyk 웹 UI에서 제공되는 기능입니다:

* [대시보드](../../getting-started/snyk-web-ui.md#view-the-assets-dashboard) - 애플리케이션 및 보안 제어의 개요를 표시하는 위젯을 제공합니다.
* [재고](../../manage-assets/) - 자산 재고에 대한 컨텍스트와 명확성을 더 잘 이해할 수 있도록 도와줍니다.
* [이슈](../../manage-risk/prioritize-issues-for-fixing/prioritization-for-snyk-apprisk.md) - 이슈 페이지에 표시된 통찰력은 Snyk에 의해 식별된 모든 이슈에 대한 중앙 집중식 보기와 추가 자산 컨텍스트를 제공합니다.
* [정책](../../manage-risk/policies/assets-policies/) - 비즈니스 컨텍스트를 추가하고 알림을 받는 프로세스를 자동화할 수 있습니다.
* [SCM 통합](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 및 [타사 통합](../../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md) - 모든 활성 통합에 대한 정보를 제공하고 새 통합을 설정할 수 있습니다.
* [분석](../../manage-risk/enterprise-analytics/application-analytics.md) - 상하 접근 방식에서 애플리케이션 보안 프로그램 상태와 결과를 검토하고 탐색할 수 있습니다.

## 주요 개념

**자산**: 의미 있는 실제 세계의 요소로, 의미 있는 것은 위험을 나르거나 다른 구성 요소의 위험을 집계하는 것을 의미합니다(예: 패키지를 포함한 저장소). 실제 세계는 Snyk 외부에 존재하는 개념을 의미합니다(예: 보편적인 용어인 저장소).

**제어**: 자산과 관련된 보안 제어. 보안 제어의 모든 가능한 상태를 볼 수 있는 [커버리지 제어](../../manage-risk/policies/assets-policies/use-cases-for-policies/coverage-control-policy-use-case.md) 섹션으로 이동하십시오.

**커버리지**: 세부적으로 어플리케이션 보안 프로그램을 위한 보안 도구(예: Snyk Open Source)에 의해 검사 및 테스트된 자산이 사용 가능한지를 평가합니다. 적용해야 할 제어를 지정하고 선택적으로 실행해야 할 빈도를 지정할 수 있는 정책의 유형입니다.

**태그**: 자산을 분류하는 방법. 상호 속성에 따라 자산을 인식하거나 처리하는 데 도움이 됩니다. 자산은 인벤토리에서 태그별로 필터링되거나 정책 규칙을 생성할 때 태그로 분류할 수 있습니다. 자동으로 자산에 태그를 지정할 수 있거나 작성한 정책에서 자산에 태그를 지정할 수 있습니다. GitHub 및 GitLab 주제는 자산 태그로 처리되며 정책 생성에 사용할 수 있습니다.

**클래스**: 자산에 비즈니스 컨텍스트를 할당하고 비즈니스 비상 중요성에 따라 자산을 분류하는 방법. 자산은 Class A(비즈니스 중요, 민감한 데이터 처리, 준수 사항 준수 등)가 가장 중요하며 Class D(테스트 앱, 샌드박스 환경 등)가 가장 중요하지 않습니다. 자산은 기본적으로 클래스 C로 할당됩니다. 클래스는 정책에서 사용하거나 정책에서 정의할 수 있습니다.

**정책**: 특정 조건에서 자동화 작업을 수행하는 방법으로, 비즈니스 컨텍스트를 가진 자산을 분류하고 태깅하는 등의 작업을 자동화할 수 있습니다. 정책 빌더 UI를 사용하여 메시지를 보내거나 커버리지 갭 제어를 설정하는 것과 같은 작업을 구성하는 데 정책을 사용할 수도 있습니다.

## 스캔 방법

웹 UI, CLI, API 또는 PR Checks에서 스캔을 시작할 수 있습니다. 자세한 내용은 [스캔 시작](../start-scanning.md)을 참조하십시오.

CLI를 사용하여 스캔을 시작하는 경우 다음 상황 중 하나를 마주할 수 있습니다:

1. CLI가 스캔하는 디렉토리에 `.git` 폴더가 있는 경우, `git remoteurl`이 Snyk Open Source, Snyk Container, Snyk IaC에 자동으로 가져오게 됩니다.

{% hint style="info" %}
CLI에 의해 스캔된 디렉토리에 `.git` 폴더가 있는 경우에도 Snyk Code는 `git remoteurl`을 자동으로 가져오지 않습니다.
{% endhint %}

2. CLI가 스캔하는 디렉토리에 `.git` 폴더가 없는 경우, 동일한 결과를 달성하기 위해 다른 테스트 또는 모니터 명령을 사용할 수 있습니다:
   * [`snyk monitor`](../../snyk-cli/commands/monitor.md#remote-repo-url-less-than-url-greater-than) - Snyk Open Source에 대해
   * [`snyk iac test`](../../snyk-cli/commands/iac-test.md#remote-repo-url-less-than-url-greater-than) - 또한 `--report` 명령이 필요합니다
   * `snyk container monitor` - 현재 옵션이 없음
   * `snyk code test` - 현재 옵션이 없음

{% hint style="info" %}
자산 대시보드 메뉴 옵션은 Snyk AppRisk Essentials 사용자에게만 사용 가능합니다.

Snyk AppRisk Pro를 사용 중이라면 [애플리케이션 분석](../../manage-risk/enterprise-analytics/application-analytics.md) 페이지로 이동하십시오.
{% endhint %}
