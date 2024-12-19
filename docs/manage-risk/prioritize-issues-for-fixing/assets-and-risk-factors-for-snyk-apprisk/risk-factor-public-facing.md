# 위험 요인: Public facing

{% hint style="info" %}
Public facing 위험 요인은 Snyk AppRisk Pro 사용자만 이용할 수 있습니다.
{% endhint %}

코드가 배포된 사실을 알면, 우려하는 결함을 누군가 악용할 수 있는 가능성이 있다는 것을 알 수 있습니다. 해당 누군가는 귀하의 조직 내에서 신뢰할 만한 사람이거나 완전히 알 수 없는 외부 엔터티가 될 수 있습니다. Snyk는 패키지나 이미지가 외부 트래픽에 노출되도록 구성되었는지 여부를 결정하여 가능성을 좁히는 데 도움을 줍니다.

AWS, Kubernetes, GCP와 같은 클라우드 네이티브 솔루션은 워크로드 실행 및 워크로드에 대한 네트워크 연결성 제공을 포함하여 다양한 목적으로 사용됩니다. 이는 서비스, 인그레스, 로드 밸런서와 같은 핵심 네트워킹 기본 요소 또는 Gloo, Istio, Envoy 등 다양한 네트워킹 솔루션의 조합을 사용하여 수행될 수 있습니다.

## 네트워크 연결 경로 결정을 위한 Snyk 분석

Snyk 플랫폼은 다양한 데이터 소스를 분석하여 네트워크 연결 경로를 계산합니다. 이 정보는 패키지와 이미지가 외부 트래픽에 노출될 수 있는지 여부를 결정하는 데 Snyk AppRisk에서 사용됩니다.&#x20;

클라우드 네이티브 솔루션은 네트워크 연결성이 구성되는 방식이 결정론적입니다. Snyk는 이러한 지식을 활용하여 사용 가능한 정보를 기반으로 답을 계산합니다. 예를 들어, 네트워크 연결성을 구성하는 방법을 이해하려면 Kubernetes의 [Services](https://kubernetes.io/docs/concepts/services-networking/service/)와 [Ingresses](https://kubernetes.io/docs/concepts/services-networking/ingress/) 문서를 참조하십시오.&#x20;

{% hint style="info" %}
Snyk AppRisk Insights는 현재 다음 구성을 지원합니다: Kubernetes services 및 ingress, 그리고 Gloo.
{% endhint %}

Kubernetes Connector는 이미지를 인그레스 구성에 대해 확인합니다. 감지되지 않으면 해당 이미지는 public facing으로 간주됩니다.

## Kubernetes Connector 통합

Public facing 위험 요인은 Kubernetes Connector 통합에 적용할 수 있습니다.&#x20;

Public facing 위험 요인은 Kubernetes Connector 통합에 중대한 영향을 미칩니다. 취약점 및 잠재적인 공격 경로가 Kubernetes 환경 내에서 어떻게 우선 순위가 매겨지고 관리되는지에 영향을 줍니다. Kubernetes Connector에 의한 Kubernetes 이벤트의 지속적 모니터링을 통해 어떠한 변경사항이나 잠재적 위험이 즉시 평가되고 Snyk 플랫폼으로 전달됩니다. 이 실시간 데이터는 보안 정책의 동적 조정 및 선행적 위험 완화를 위한 계속적인 모니터링을 가능하게 하여 클라우드 네이티브 인프라의 무결성과 보안을 보장합니다.

## Public-facing 위험 요인에 대한 기술적 세부 정보

Kubernetes Connector는 Kubernetes 이벤트를 지속적으로 모니터링합니다. 이 이벤트는 Snyk 플랫폼에 지속적으로 스트리밍됩니다.&#x20;

매 시간마다, 데이터 파이프라인은 클러스터의 상태를 조정하여 스냅샷을 생성합니다. 이 스냅샷은 서비스에서 파드로, 인그레스에서 서비스로의 네트워크 관계를 계산하는 데 사용됩니다. 동시에 해당 스냅샷은 해당 시점에 실행 중인 이미지를 추정하는 데 사용됩니다.

분석은 현재 연결 사양에 따라 수행됩니다. 분석의 세분화는 `port` 및 `http path` 수준에서 수행됩니다. 구성된 경로에는 네트워크 정책, 보안 그룹, 방화벽 등이 적용될 수 있습니다. Snyk는 현재 해당 제약을 계산에 포함시키지 않습니다.&#x20;

동일한 간격으로, 데이터 파이프라인은 모든 Snyk 프로젝트 및 데이터 소스의 스냅샷을 취하고 이미지와 패키지를 추정합니다. 이 스냅샷은 주어진 고객의 경우 Snyk에 알려진 이미지 및 패키지를 결정하는 데 사용됩니다.&#x20;

두 스냅샷이 비교되고 해당 시점의 public-facing 요인을 결정하기 위해 증거 그래프가 생성됩니다.