# Snyk 앱리스크 통합 구성하기

{% hint style="info" %}
Snyk 앱리스크 에센셜은 Snyk 엔터프라이즈에 포함되어 있습니다.
{% endhint %}

## Snyk 앱리스크 에센셜 사전 준비 사항

Snyk 앱리스크 에센셜 설정 전에 다음 사전 준비 사항을 충족시켜야 합니다:

* Snyk 앱리스크와 연관된 그룹의 그룹 관리자이거나 보기 그룹 및 앱리스크 편집 권한이있는 그룹 수준 역할이 할당되어 있어야 합니다.
* Snyk 앱리스크와 연관된 그룹에는 Snyk 응용프로그램 보안 제품을 도입한 조직이 포함되어 있어야 합니다.
* 클라우드 기반 SCM 도구(Azure DevOps, GitHub, GitLab 등)를 Snyk 앱리스크에 레포지토리 자산 발견을 위해 등록하기 위한 필요한 권한 및 권한을 가지고 있어야 합니다.

## Snyk 앱리스크 구성 및 SCM 통합 설정

Snyk 앱리스크를 도입하고 모든 인벤토리 기반 코드 자산을 식별하고 보안 제어가 설정된 자산을 감지하여 앱리스크를 온보딩하기 시작합니다.

## Snyk 앱리스크 접근

[Snyk 웹 UI](../../../getting-started/snyk-web-ui.md)에서 Snyk 앱리스크에 액세스할 수 있습니다.

* Snyk 그룹 수준에서 Snyk 그룹의 Snyk 앱리스크에 액세스합니다.
* 그룹 관리자 액세스 권한을 확인하십시오.
* Snyk 앱리스크가 활성화 된 그룹에 액세스하시기 바랍니다.

## 통합 설정 <a href="#setup-integrations" id="setup-integrations"></a>

Snyk 앱리스크에 올바르게 액세스할 수 있음을 확인한 후 인테그레이션을 설정하여 자산 인벤토리를 작성할 수 있습니다.

{% hint style="info" %}
모든 기능을 활성화한 후 2시간 이내에 자동으로 스캔된 정보가 가져오됩니다.
{% endhint %}

인테그레이션 보기에서 인테그레이션을 액세스하고 구성할 수 있습니다. 통합 허브 옵션을 선택하여 사용 가능한 모든 인테그레이션 목록을 확인할 수 있습니다. [인테그레이션 허브 사용](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#using-the-integration-hub) 섹션에서 통합 구성에 대한 자세한 내용을 찾을 수 있습니다.

인테그레이션 보기의 기본 표시에는 구성된 Snyk 인테그레이션이 포함됩니다. 각 인테그레이션의 상태 **연결됨** 또는 **연결되지 않음**은 Snyk로 가져오는 특정 콘텐츠에 따라 달라집니다.

인테그레이션 보기는 개인 설정을 적용할 수 있도록 구성할 수 있으므로 [기존 인테그레이션 사용자 지정](../../../getting-started/snyk-web-ui.md#edit-an-integration) 또는 [새 SCM 인테그레이션 연결](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-apprisk-integrations-ecosystem)이 가능합니다.

<figure><img src="../../../.gitbook/assets/image (357) (1).png" alt="Snyk 앱리스크 - 사용 가능한 통합 목록을 표시하는 통합 허브 옵션"><figcaption><p>Snyk 앱리스크 - 사용 가능한 통합 목록을 표시하는 통합 허브 옵션</p></figcaption></figure>

통합 허브를 클릭하면 사용 가능한 인테그레이션 목록이 표시됩니다. 각 인테그레이션에 대해 하나 이상의 프로필을 추가할 수 있습니다.

### SCM 인테그레이션

Snyk 앱리스크 인테그레이션 허브를 사용하여 SCM 인테그레이션을 구성할 수 있습니다.

{% hint style="warning" %}
통합 허브는 개별적으로 Snyk 앱리스크에 전념한 별도의 통합 인터페이스입니다.
{% endhint %}

일반적으로 유지 또는 개발을 위해 유지되거나 제한되는 조직 수준에서 Snyk 코드, 오픈 소스, 컨테이너 및 IaC 스캔을 촉진하는 조직 수준의 인테그레이션은 개발자에 의해 유지 또는 제한됩니다. 그룹 수준의 인테그레이션 화면에서 토큰을 설정하면 보안적인 종합적인 관점을 보유하게 됩니다. 이 토큰은 개발자가 액세스 권한을 가지지 않은 모든 저장소에 대한 넓은 액세스를 제공합니다. 일종의 보안 관점에서 자산에 대한 보안적인 관점이 개발 팀과 일치하는지 확인하는 것입니다.

지원되는 SCM 인테그레이션은 다음과 같습니다 :

* [GitHub](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#group-level-snyk-apprisk-integrations)
* [GitLab](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab.md#group-level-snyk-apprisk-integrations)
* [Azure DevOps (Azure Repos)](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/azure-repositories-tfs.md#group-level-snyk-apprisk-integrations)
* [BitBucket](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-cloud.md#group-level-snyk-apprisk-integrations)

자세한 내용은 [Snyk SCM 인테그레이션](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#snyk-apprisk-integrations-ecosystem) 페이지를 참조하십시오.

### 브로커가 있는 SCM 인테그레이션 <a href="#brokered-scm-integration" id="brokered-scm-integration"></a>

Snyk 브로커를 설정할 때 새 브로커를 서비스하거나 기존 [Snyk 브로커 연결](https://docs.snyk.io/enterprise-setup/snyk-broker)을 업데이트해야 할 몇 가지 질문이 있습니다:

* API 요율 제한 문제가 발생하고 있습니까?
* 모든 관련 SCM 저장소에 액세스 권한이 있는 사용자의 SCM 토큰을 업데이트해야 하는가?
* 1000개 이상의 저장소가 있습니까?

위 질문 중 하나에 대답이 "예"이면 Snyk 앱리스크 SCM 연결을 수용할 수 있도록 새 Snyk 브로커를 배포해야 합니다.

{% hint style="info" %}
Snyk는 Snyk 앱리스크 브로커를 위해 별도로 Snyk에서 새 조직을 생성하는 것을 권장합니다.
{% endhint %}

자세한 내용은 [Snyk 브로커 - 앱리스크](../../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md) 페이지를 참조하십시오.

### 제3자 인테그레이션

{% hint style="info" %}
제3자 인테그레이션은 주로 Snyk 앱리스크 Pro에만 사용 가능합니다. Snyk 앱리스크 에센셜 사용자로서 Snyk 앱리스크 에센셜에도 사용 가능한 제3자 인테그레이션을 사용할 수 있습니다.
{% endhint %}

제3자 인테그레이션을 설정하려면 [Snyk 앱리스크 인테그레이션 허브](../../../getting-started/snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors)를 사용할 수 있습니다. 각 Snyk 조직에서 관리자는 개발자가 활용하는 애플리케이션에 대한 제한된 액세스를 제공하는 토큰을 부여할 수 있습니다. Snyk 앱리스크 관점에서 토큰의 목적은 현재 자산을 Snyk로 가져오는 것과 비교하는 것입니다.

각 Snyk 조직에서 관리자는 개발자가 활용하는 애플리케이션에 대한 제한된 액세스를 부여할 수 있습니다.

Snyk 앱리스크에서 사용하는 토큰의 범위는 Snyk로 가져온 기존 자산을 설명하는 것입니다.

지원되는 제3자 인테그레이션은 다음과 같습니다 :

* Veracode
* Nightfall

### SCM 인테그레이션을 위한 Backstage 파일

Backstage는 사용자가 리포지토리에 메타데이터 또는 주석을 추가하여 사용 가능한 자원을 조직하고 분류하여 쉽게 탐색하고 이해할 수 있는 서비스 카탈로그입니다. SCM 통합을 활용하여 Backstage 카탈로그 파일과 연관된 메타데이터를 Snyk 앱리스크로 가져올 수 있습니다.

이 기능을 사용하는 방법에 대한 자세한 내용은 [SCM 인테그레이션을 위한 애플리케이션 컨텍스트](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/) 문서를 참조하십시오.

{% hint style="info" %}
모든 Snyk 앱리스크 인테그레이션을 설정한 후, 레포지토리 수에 따라 결과가 나타나는 데 최대 하루가 소요될 수 있습니다.
{% endhint %}

## 기능

Snyk 앱리스크 기능은 그룹 수준에서 여러 메뉴 옵션으로 분할되어 있습니다.

* [대시보드](../../../getting-started/snyk-web-ui.md#view-the-assets-dashboard)
* [인벤토리](../../../manage-assets/)
* [정책](../../../manage-risk/policies/assets-policies/)
* [SCM을 위한 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 및 제3자를 위한 통합]\(../../../integrate-with-snyk/#integrations-for-snyk-apprisk)
* [이슈](../../../manage-risk/prioritize-issues-for-fixing/)

#### 인벤토리 보기

인벤토리 기능은 특정 영역에 집중된 네 가지 섹션으로 구성됩니다:

* **모든 자산**: 자산 유형별로 그룹화된 모든 발견된 자산입니다.
* **자산 계층 구조**: 자산 계층 구조 레이아웃은 계층 구조로 자산을 표시합니다. 자산 목록은 이슈 카운트에 따라 정렬되며, 해당하는 경우 패키지 자산은 해당 리포지토리 바로 아래에 나열됩니다. 자산 계층 구조는 필터가 적용되지 않은 경우에만 표시됩니다. 자산 계층 뷰의 모든 옵션에 대한 자세한 내용은 [자산 인벤토리 구성 요소](../../../manage-assets/assets-inventory-components.md) 페이지를 참조하고, 필터 옵션 및 사용 방법에 대한 자세한 내용은 [필터 기능](../../../manage-assets/assets-inventory-features.md#filters-capabilities) 페이지를 참조하십시오.
* **팀**:SCM 리포지토리 자산은 팀별로 그룹화됩니다. 팀이 할당된 SCM 조직 및 리포지토리만이 이 레이아웃에 나타납니다.
* **기술**: SCM 리포지토리 자산은 Snyk 앱리스크가 감지하고 태그를 지정한 기술에 따라 그룹화됩니다.

Snyk 앱리스크를 처음 사용하는 경우 Coverage 필터를 사용하여 현재 Snyk를 구현한 위치를 확인하는 것을 권장합니다. 그런 다음, **Set coverage control** 정책에서 설정한 커버리지 요구 사항을 충족하지 못하는 자산을 식별하기 위해 Coverage Gap 필터를 사용할 수 있습니다.

Coverage Gap 필터를 사용하여 다음을 수행할 수 있습니다:

*   **세트 커버리지 제어 정책 요구 사항을 충족하지 않는 모든 자산을 찾습니다**:

    <figure><img src="../../../.gitbook/assets/image (1) (10).png" alt="Coverage gap 필터를 사용하여 정책에서 벗어난 자산을 찾습니다"><figcaption><p>Coverage gap - Use case 1</p></figcaption></figure>
*   **Snyk 오픈 소스 또는 Snyk 코드 또는 둘 다를 위한 커버리지 요구 사항을 충족하지 않는 모든 자산을 찾습니다**:

    <figure><img src="../../../.gitbook/assets/image (1) (10) (1).png" alt="Coverage gap 필터를 사용하여 이나 에 대한 정책에서 벗어난 모든 자산을 찾습니다"><figcaption><p>Coverage gap - Use case 2</p></figcaption></figure>

#### 태그 <a href="#hardbreak-tags" id="hardbreak-tags"></a>

자산을 분류하려면 태그를 사용할 수 있습니다. 태그는 여러 가지 방법으로 사용할 수 있습니다:

* 자동 태그: Snyk AppRisk는 리포지토리에서 사용된 기술(Python, Terraform 등)과 최신 리포지토리 업데이트에 대한 정보를 자동으로 리포지토리 자산에 태그합니다. 또한 정책을 사용하여 리포지토리 및 패키지 자산에 태그를 추가할 수 있습니다. GitHub과 GitLab의 주제도 리포지토리에서 가져와 Snyk AppRisk의 자산 태그로 적용할 수 있습니다.

{% hint style="info" %}
BitBucket은 리포지토리에서 사용된 언어를 자동으로 감지할 수 없습니다. Snyk AppRisk에서는 BitBucket에 대해 수동으로 추가된 언어 태그만 볼 수 있습니다. 자세한 내용은 BitBucket에서 제공하는 공식 문서를 참조하십시오.
{% endhint %}

* 사용자 정의 태그: 정책을 통해 시스템에서 생성된 태그를 넘어 자산을 분류할 수 있는 맞춤형 태그를 설정할 수 있습니다. 자세한 내용은 [정책 생성](../../../manage-risk/policies/assets-policies/create-policies.md) 페이지를 참조하십시오.

#### 대시보드

대시보드를 사용하여 애플리케이션 및 보안 제어에 대한 빠른 개요를 확인할 수 있습니다. 기본 위젯을 사용하고 필요에 따라 표시된 정보를 사용자화하거나 필요에 맞는 새로운 위젯을 추가할 수 있습니다. 자세한 내용은 [자산 대시보드](../../../getting-started/snyk-web-ui.md#view-the-assets-dashboard) 페이지를 참조하십시오.

다음은 사용 가능한 대시보드 위젯입니다:

* **SAST 범위**: Snyk Code 및 Snyk Infrastructure as Code로 커버되는 리포지토리와 그렇지 않은 리포지토리를 확인합니다.

{% hint style="info" %}
SAST 범위 위젯은 OR 조건을 사용합니다. 즉, 리포지토리가 Snyk Code 또는 Snyk Infrastructure as Code로 커버된다면 해당 리포지토리는 SAST로 커버됩니다.
{% endhint %}

* **SCA 범위**: Snyk Open Source 및 Snyk Container로 커버되는 리포지토리와 그렇지 않은 리포지토리를 확인합니다. Snyk Open Source 범위 또는 Snyk Container 범위 중 하나만 보고 싶다면 위젯을 편집할 수 있습니다.

{% hint style="info" %}
SCA 범위 위젯은 OR 조건을 사용합니다. 즉, 리포지토리가 Snyk Open Source 또는 Snyk Container로 커버된다면 해당 리포지토리는 SCA로 커버됩니다.
{% endhint %}

* **소스별 리포지토리 분류**: Snyk AppRisk가 SCM 통합(Azure DevOps(Azure Repos), Gitlab, GitHub, BitBucket)을 사용하여 발견한 리포지토리를 확인합니다. "기타" 카테고리는 Snyk에서 발견했지만 SCM 리포지토리와 상호 연관되지 않은 리포지토리입니다.
* **기술별 분류**: Snyk이 발견한 리포지토리의 주요 기술(언어) 태그를 확인합니다.
* **유형별 자산 분류**: 자산이 리포지토리인지 패키지인지 확인합니다.
* **리포지토리 활동**: 리포지토리가 활성화되어 있는지, 비활성화되어 있는지, 휴면 상태인지 확인합니다.
