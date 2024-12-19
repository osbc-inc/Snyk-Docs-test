# Insights 설정: 이미지 스캔

{% hint style="info" %}
Insights 기능은 Snyk AppRisk Pro에서만 사용 가능합니다.
{% endhint %}

코드, 오픈 소스 및 컨테이너 문제의 위험 요소를 결정하고 우선 순위를 정하려면 컨테이너 이미지를 스캔해야 합니다. [Snyk Container](../../../scan-with-snyk/snyk-container/)을 사용하여 컨테이너 이미지를 스캔합니다.&#x20;

컨테이너 이미지는 Snyk AppRisk를 구동하는 애플리케이션 모델의 중심에 있습니다. 컨테이너 이미지에는 소스 코드와 의존성이 포함되어 있으며, 이는 실행 중인 환경에 배포되어 Snyk AppRisk가 개발 및 배포 상태를 연결하는 데 사용됩니다.

Snyk AppRisk는 배포된 컨테이너 이미지를 [Kubernetes Connector](set-up-insights-kubernetes-connector.md)를 이용하여 식별하고, Snyk Container를 사용하여 스캔한 이미지 목록과 배포된 컨테이너 이미지를 비교합니다.&#x20;

{% hint style="info" %}
Snyk은 각 이미지를 적어도 하나의 Snyk Container 통합을 사용하여 스캔하는 것을 권장합니다.
{% endhint %}

## Snyk Container CLI를 사용한 스캔

이미지 이름이 일치하는지 확인하려면 Kubernetes 배포에서 사용되는 이미지의 전체 이름을 명시하십시오.&#x20;

예:

`snyk container monitor gcr.io/my-company/my-image:latest`

자세한 내용은 [Snyk 컨테이너 보안을 위한 CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)를 참조하십시오.

## Snyk Container 레지스트리 스캔

이름이 일치하는 경우 Snyk에서 동일한 컨테이너 레지스트리에서 가져온 이미지를 Kubernetes 배포에서 참조하는 경우 일치합니다.

자세한 내용은 [Snyk Container - 통합](../../../scan-with-snyk/snyk-container/container-registry-integrations/)을 참조하십시오.

## Kubernetes 통합을 통한 Snyk Container 스캔

컨테이너 이미지의 이름이 일치할 것입니다. 배포된 이미지가 Snyk에 의해 스캔되고 프로젝트로 생성되기 때문입니다.

{% hint style="info" %}
Kubernetes Connector를 올바르게 설정했는지 확인하려면 **Issues** 페이지의 **Set up Insights** 탭으로 이동하여 **Image coverage** 섹션을 확인하여 Insights에서 액세스할 수 있는 데이터를 볼 수 있습니다.
{% endhint %}

자세한 내용은 [Kubernetes 통합](../../../scan-with-snyk/snyk-container/kubernetes-integration/)을 참조하십시오.