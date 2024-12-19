# Snyk AppRisk을 위한 자산 및 위험 요인

Snyk AppRisk Issues 메뉴의 기능은 보다 나은 보안 문제 우선순위 결정을 돕기 위해 당신의 애플리케이션 컨텍스트를 이해하는 데 의존합니다. Snyk AppRisk Essentials 요금제의 경우 이를 통해 자산 및 문제의 트리지 및 우선 순위를 제공하며 Snyk AppRisk Pro 요금제의 경우 특정 [위험 요인](./#risk-factors) 및 [증거 그래프](../using-the-issues-ui-with-snyk-apprisk/evidence-graph.md) 정보도 추가됩니다.&#x20;

* [자산](./#자산)은 Snyk 인사이트를 사용하여 이미지, Kubernetes 자원 및 패키지를 중점적으로 분석하여 이러한 요소 간 상호작용을 이해합니다.
* [위험 요인](./#risk-factors)은 Snyk 인사이트를 사용하여 분석되며 다음 네 가지 주요 범주로 그룹화됩니다:
  * [배포됨](risk-factor-deployed.md)&#x20;
  * [로드된 패키지](risk-factor-loaded-package.md)
  * [운영 체제 상태](risk-factor-os-condition.md)&#x20;
  * [외부 노출](risk-factor-public-facing.md)

다음 비디오에서는 Snyk AppRisk Issues UI에서 찾을 수 있는 주요 기능을 소개합니다:

{% embed url="https://www.youtube.com/watch?v=MZZINnVmGL0" %}
비디오를 좋아했다면, [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)에서 해당 코스를 확인해보세요!
{% endembed %}

## 자산

Snyk AppRisk 인사이트는 이 페이지에 설명된 자산을 분석합니다.

### 이미지

이미지는 Docker 이미지를 나타내는 자산입니다. Snyk Container는 Docker 이미지에 대해 보안 검사를 수행합니다. 이미지는 Snyk Project와 일대다로 매핑될 수 있습니다. Docker 이미지는 SHA로 표시되는 자연 ID를 갖습니다. Snyk는 동일한 이미지를 이 자연 ID를 사용하여 매핑하여 사용합니다.

### Kubernetes 자원

Kubernetes 자원은 Kubernetes 객체를 나타내는 자산입니다. Kubernetes 커넥터는 Kubernetes 클러스터에서 자원 정보를 수집합니다.&#x20;

Kubernetes 자원은 Snyk 프로젝트에 매핑되지 않습니다. 이는 일부 위험 요인을 계산하는 데 사용되는 내부 entity입니다. 이러한 위험 요인은 패키지 및 이미지와 관련될 수 있습니다.

### 패키지

패키지는 소프트웨어 패키지를 나타내는 자산입니다. Snyk Open Source 및 Snyk Code 제품은 파일에 대해 보안 검사를 수행합니다. 이 파일은 소프트웨어 패키지의 패키지 관리자 선언 및 소스 코드를 나타냅니다. 패키지는 해당 소프트웨어 패키지의 표현입니다.

패키지는 Snyk Projects와 일대일로 매핑될 수 있습니다. Snyk Open Source 및 Snyk Code의 검사로 생성된 프로젝트에 의해 생성된 패키지입니다. 이러한 제품이 식별한 모든 문제 및 해당 프로젝트에 속한 것은 패키지 entity에 매핑됩니다.&#x20;

패키지라는 용어는 매우 일반적인 추상화입니다. 버전이 없습니다. 현재 시점의 소프트웨어 패키지의 상태를 나타냅니다. 시간은 Snyk 처리 파이프라인이 완료되고 해당 시간의 Snyk Projects의 상태에 의해 결정됩니다.&#x20;

Snyk Open Source는 패키지 의존성 manifest에 선언된 타사 의존성을 가리키는 용어로 패키지를 사용합니다. Snyk는 현재 이러한 타사 의존성의 세분성을 노출하지 않습니다. 그러나 우선 순위 데이터 모델의 관점에서는 타사 및 자사 패키지 사이에 구분이 없습니다. 해당 시점에서 이를 패키지 객체로 표시할 것입니다.

## 위험 요인

이미지, 패키지 및 Kubernetes 자원을 "애플리케이션 컨텍스트"로 이해함으로써 Snyk는 다음과 같은 위험 요소를 계산할 수 있습니다:

* [배포됨](risk-factor-deployed.md)
* [로드된 패키지](risk-factor-loaded-package.md)
* [운영 체제 상태](risk-factor-os-condition.md)
* [외부 노출](risk-factor-public-facing.md)

당신의 애플리케이션을 위한 활성화 및 비활성화할 수 있는 이러한 "애플리케이션 컨텍스트" 위험 요인은 **그룹 설정**의 **Insights** UI 탭에서 설정할 수 있습니다. 위험 요인을 비활성화하면 공급자 선택 또는 Kubernetes 클러스터 매핑이 제거되며 Snyk가 해당 요소를 더 이상 계산하지 않습니다.&#x20;

{% hint style="info" %}
위험 요소는 Snyk AppRisk Pro에서만 사용할 수 있습니다.

Snyk 웹 UI에서 Group Settings의 Insights 탭은 Snyk AppRisk Pro에만 사용할 수 있습니다.&#x20;
{% endhint %}

애플리케이션에 대해 활성화된 통합 옵션에 따라 위험 요인이 다르게 적용됩니다. Group settings에서 사용 가능한 Insights 옵션을 사용자화하여 [통합을 우선순위 지정](../set-up-insights-for-snyk-apprisk/#prioritize-your-integrations)할 수 있습니다.

<figure><img src="../../../.gitbook/assets/image (457).png" alt="Snyk AppRisk Pro - Group settings의 Insights 탭"><figcaption><p>Snyk AppRisk Pro - Group settings의 Insights 탭</p></figcaption></figure>

{% hint style="info" %}
위험 요인 설정은 최대 네 시간이 소요될 수 있습니다.
{% endhint %}