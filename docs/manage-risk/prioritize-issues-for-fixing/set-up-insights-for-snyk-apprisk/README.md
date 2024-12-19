# Snyk AppRisk의 인사이트 설정

{% hint style="info" %}
인사이트 기능은 Snyk AppRisk Pro에서만 이용할 수 있습니다.
{% endhint %}

## Snyk AppRisk을 위한 우선순위 설정 전제조건

[{{ Snyk }} Container](../../../scan-with-snyk/snyk-container/)을 사용하여 이미지를 스캔하는 응용 프로그램과 Set up Insights 옵션을 사용하여 Snyk AppRisk 우선순위를 사용자화할 수 있습니다. 또한 [{{ Snyk }} Open Source](../../../scan-with-snyk/snyk-open-source/)와 [{{ Snyk }} Code](../../../scan-with-snyk/snyk-code/)로 오픈 소스 종속성 및 소스 코드를 스캔함으로써 추가 가치를 얻을 수 있습니다.

### 필요한 위험 요소는 무엇인가요?

{% hint style="info" %}
위험 요소는 Snyk AppRisk Pro에서만 이용할 수 있습니다.
{% endhint %}

{{ Snyk }} AppRisk Insights 제품은 취약점에 대한 다음과 같은 위험 요소를 제공하여 작동합니다:

- [**이미 배포됨**](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md): 내 코드와 컨테이너 이미지가 어디에 배포되었는가?
- [**로드된 패키지**](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md): 이미지의 종속성인 서드 파티 패키지가 로드되었는가?
- [**OS 상태**](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-os-condition.md): 이 취약점이 운영 체제에 적용되는가?
- [**공개된**](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-public-facing.md): 내 컨테이너에 인터넷 노출이 있는가?

### 위험 요인 사용 기준

이 네 가지 위험 요소에 대한 데이터를 얻으려면 다음 조건을 충족해야 합니다:

#### **로드된 패키지 위험 요인**

로드된 패키지 위험 요인을 사용하려면 다음 조건을 충족해야 합니다:

- 자주 로드되는 패키지가 다른 패키지보다 애플리케이션에 더 높은 위험을 야기합니다.
- 인사이트를 통해 실행시 우선순위를 적용하기 위한 최소 요구 사항은 다음과 같습니다:
  - Snyk AppRisk에 Dynatrace 또는 Sysdig 통합 또는 [Snyk 런타임 센서](../../snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)을 설정해야 합니다. [실행시 타사 통합](../../snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md) 페이지에서 자세한 내용을 찾을 수 있습니다.

#### **OS 상태 위험 요인**

OS 상태 위험 요인을 사용하려면 다음 조건을 충족해야 합니다:

- 소스 코드 및 종속성이 컨테이너 이미지로 빌드되고 {{Snyk Container}}를 사용하여 스캔되어야 합니다. 이 우선 순위를 위해 값이 발생하는 데 필요한 최소한의 요구 사항입니다.

<figure><img src="../../../.gitbook/assets/Example OS condition.png" alt="컨테이너 이미지로 빌드된 소스 코드와 종속성"><figcaption><p>컨테이너 이미지로 빌드된 소스 코드와 종속성</p></figcaption></figure>

#### **배포 및 공개된 위험 요인**

배포 및 공개된 위험 요인을 사용하려면 다음 조건을 충족해야 합니다:

- 두 위험 요인 모두 쿠버네티스 클러스터에 배포된 컨테이너 이미지가 있어야 합니다. 여기서 [쿠버네티스 커넥터](set-up-insights-kubernetes-connector.md)를 배포할 수 있습니다.

이미지를 스캔하는 중에 코드의 네 가지 위험 요소에 대한 데이터를 수집하려면 이러한 요구 사항을 충족해야 합니다.

### 인사이트를 최대화하세요

Snyk은 인사이트를 최대로 활용하기 위해 다음 단계를 수행하는 것을 권장합니다:

- [Snyk Open Source](../../../getting-started/snyk-web-ui.md)를 사용하여 써드 파티 종속성을 스캔합니다.
- [Snyk Code](../../../getting-started/snyk-web-ui.md)를 사용하여 소스 코드를 스캔합니다.
- 하나의 응용 프로그램부터 시작하고 이어서 확장합니다.

소스 코드와 써드 파티 종속성을 함께 스캔함으로써 애플리케이션 맥락을 제공하는 위험 요인 데이터를 얻을 수 있으며, 열린 문제를 더 잘 우선순위를 지정할 수 있습니다.

## 우선순위 설정 개요

{% hint style="info" %}
Snyk는 가장 효과적인 통합을 달성하고 끊임없이 확장되는 기능에 액세스하기 위해 [Snyk 런타임 센서](../../snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)를 설치하는 것을 권장합니다.
{% endhint %}

이슈 설정을 위한 주요 단계는 다음과 같습니다:

1. 사용자에게 그룹 뷰어 역할 또는 조직 협력자 역할을 부여합니다. [우선순위 설정: 사용자 권한](set-up-insights-user-permissions.md)을 참조하십시오.
2. 필요한 조직, 역할, 권한을 생성하고 에이전트를 배포합니다. [우선순위 설정: 쿠버네티스 커넥터](set-up-insights-kubernetes-connector.md)을 참조하십시오.

{% hint style="info" %}
쿠버네티스 커넥터는 쿠버네티스 컨트롤러인 Snyk-Monitor와 다릅니다.
{% endhint %}

3. 이미지를 적절하게 스캔하여 Snyk가 올바른 데이터에 액세스할 수 있도록 합니다. [우선순위 설정: 이미지 스캔](set-up-insights-image-scanning.md)을 참조하십시오.
4. 우선순위를 사용하려는 응용 프로그램을 위해 필요한 링크를 설정합니다. [우선순위 설정: {{Snyk Open Source}}, Code, Container 프로젝트 연결](set-up-insights-associating-snyk-open-source-code-and-container-projects.md)를 참조하십시오.
5. 타사 런타임 통합 또는 Snyk 런타임 센서를 설정하여 더 많은 런타임 데이터를 얻습니다.
6. 우선순위 기능을 올바르게 설정했는지 확인하려면 **이슈** 페이지의 **Snyk AppRisk 설정** 탭으로 이동하고 Snyk가 액세스한 데이터를 확인합니다.\
   진행 상황의 세부 정보를 확인하려면 그룹별로 관련 섹션을 필터링할 수 있습니다.

{% hint style="info" %}
인사이트 페이지에서 **인사이트 설정** 탭은 Snyk AppRisk Pro 사용자에게만 제공됩니다.
{% endhint %}

## 인사이트 UI 설정 설정

인사이트 설정은 아래 세 가지 주요 범주로 구성됩니다:

- [위험 요소](./#risk-factors): 식별된 문제의 위험을 계산하는데 관련된 위험 요소를 활성화 또는 비활성화할 수 있습니다.
- [공급자 선택](./#provider-selection): 쿠버네티스 런타임 공급자를 구성할 수 있습니다.
- [쿠버네티스 클러스터 매핑](./#kubernetes-cluster-mapping): 프로바이더에 의해 식별된 클러스터를 선호하는 이름과 매치할 수 있게 해줍니다.

위 세 가지 설정은 Snyk 웹 UI에서 그룹 설정 아래의 설정 옵션에서 찾을 수 있습니다.

### 위험 요소

사용 가능한 위험 요소를 활성화하거나 비활성화할 수 있습니다: [배포됨](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md), [로드된 패키지](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md), [OS 상태](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-os-condition.md), [공개된](../assets-and-risk-factors-for-snyk-apprisk/risk-factor-public-facing.md). 위험 요소가 비활성화되면 문제 계산에 사용되지 않습니다.

위험 요소를 [{{ Snyk }} Web UI](../../../getting-started/snyk-web-ui.md), Group Settings, Settings 옵션, Risk factors에서 활성화 또는 비활성화할 수 있습니다.

### 공급자 선택

여러 쿠버네티스 런타임 공급자를 설정하여 현재 통합에서 관련 런타임 위험 요소를 수집할 수 있습니다. 여러 쿠버네티스 런타임 공급자가 동일한 리소스를 보고할 때, 로드된 패키지와 같은 일부 세부 정보는 한 곳에서만 사용할 수 있습니다.

동일한 배포가 식별되면 어느 것이 우선 순위를 가져야 하는지 선택할 수 있습니다. 기본 공급자 설정이 선택되지 않은 경우 가장 먼저 제공된 데이터가 사용됩니다.

개별 공급자도 여기에서 사용하지 않도록 비활성화할 수 있습니다.

### 쿠버네티스 클러스터 매핑

런타임 공급자는 서로 다른 이름으로 식별되는 클러스터에 대한 다른 이름을 보고할 수 있습니다. Insight가 두 데이터 세트에서 리소스를 연관시킬 수 있도록 클러스터 이름 매핑을 추가할 수 있습니다.

예를 들어, Snyk가 `dev`로 클러스터 이름을 보고하고 통합이 동일한 클러스터 이름을 `dev-foo`로 보고하는 경우, 통합용 클러스터 이름을 소스 이름 `dev-foo` 및 대상 이름 `dev`로 추가할 수 있습니다.

{% hint style="info" %}
각 소스 이름이 하나의 클러스터 매핑에만 할당되도록 하십시오.
{% endhint %}

## 통합 우선순위 설정

문제를 우선순위를 지정할 때 사용 가능한 통합 옵션과 관련된 위험 요소를 이해하는 것이 중요합니다.

문제 우선순위 설정을 선택할 때 선택할 수 있는 다음과 같은 통합 옵션이 있습니다. 이 설정을 사용자 지정하려면 [Snyk Web UI](../../../getting-started/snyk-web-ui.md)에서 그룹 수준 탐색하여 설정 메뉴로 이동합니다.

- [**Snyk 런타임 센서**](../../snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md): 실행 중 분석을 위한 사용되며, 실행 중 애플리케이션의 실제 사용 및 잠재적 취약점에 대한 자세한 통찰력을 제공합니다. 이 센서는 실시간 애플리케이션 동작을 기반으로 실시간 애플리케이션 행동에 따라 런타임 취약점을 식별하고 잠재적 위험을 평가하는 데 도움이 됩니다.
- **쿠버네티스 커넥터**: 쿠버네티스 배포에서 포괄적인 모니터링을 제공합니다. 이 통합은 쿠버네티스 클러스터 내에서 취약점을 식별하고 워크로드 취약점, 인프라 설정 오류 및 잠재적 악의적 활동에 대한 데이터를 제공합니다.
- [**타사 통합**](../../snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md): 클라우드 제공업체 또는 CI/CD 도구와 같은 이러한 통합을 사용하여 더 나은 취약성 평가를 위해 추가 컨텍스트 및 데이터 원본을 제공합니다. 이 통합은 구성 오류, 노출 지점 및 특정 통합 취약성을 식별하는 데 도움이 됩니다.

### 통합 옵션에 매핑된 위험 요소

각 통합 옵션은 다른 위험 요소를 사용합니다:

- **Snyk 런타임 센서**: 배포됨 및 로드된 패키지
- **쿠버네티스 커넥터**: 배포됨, 공개된
- **타사 통합**: 배포됨 및 로드된 패키지

이 통합 옵션을 활용하여 안전성 위험의 포괄적인 커버리지와 정확한 우선순위를 지정할 수 있습니다.

**배포됨** 및 **로드된 패키지** 위험 요소는 실제 애플리케이션의 상태와 동작에 대한 구체적이고 실질적인 통찰력을 제공하여 우선순위를 지정하고 문제 해결 노력을 개선할 수 있습니다.

**공개된** 위험 요소