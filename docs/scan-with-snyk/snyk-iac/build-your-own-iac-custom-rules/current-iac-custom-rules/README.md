# 현재 IaC 사용자 정의 규칙

{% hint style="info" %}
**기능 가용성**

IaC 사용자 정의 규칙은 엔터프라이즈 요금제에서만 사용할 수 있습니다. 자세한 정보는 [요금제 및 가격 책정](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

{{Snyk IaC}}에는 AWS, Azure, GCP 및 Kubernetes를 다루는 보안 규칙의 포괄적인 목록이 포함되어 있습니다. 이러한 규칙은 보안 연구, 모범 사례, 인정된 표준 및 벤치마크에 기반하고 있습니다. 이러한 규칙은 Snyk의 보안 엔지니어링 팀에 의해 적극적으로 유지보수되며, 새로운 규칙이 정기적으로 출시됩니다.

이러한 규칙은 첫 번째 스캔에서 대부분의 요구 사항을 충족하기 위해 의도되었지만 시스템에 대해 추가적인 보안 규칙을 강제 적용해야 하는 경우가 있을 수 있습니다.

## 추가 {{Snyk IaC}} 사용자 정의 규칙 작성

IaC SDK를 사용하여 보안 팀이 직접 규칙을 정의하여 [Snyk CLI](../../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/)에서 실행되고 개발자에게 피드백을 제공할 수 있습니다.

이 SDK를 사용하면 Snyk IaC에 직접 추가적인 규칙을 추가하여 제공된 표준 규칙과 함께 실행할 수 있으므로 개발 팀에 종합적인 보안 피드백을 제공할 수 있습니다.

이 섹션은 Snyk Infrastructure as Code (IaC) SDK를 사용하는 데 도움이 되는 초기 지침을 제공합니다:

- [SDK 설치](install-the-sdk.md)
- [SDK로 시작하기](writing-rules-using-the-sdk/)
- [Snyk CLI와 함께 IaC 사용자 정의 규칙 사용](use-iac-custom-rules-with-cli/)
- [파이프라인 내에서 사용자 정의 규칙 통합](iac-custom-rules-within-a-pipeline.md)
- [SDK 참조](sdk-reference.md)

<figure><img src="../../../../.gitbook/assets/image (159) (1) (1) (2).png" alt="자체 사용자 정의 규칙을 작성, 배포하고 Snyk CLI를 사용하여 파일을 스캔하는 종단 간 흐름"><figcaption><p>자체 사용자 정의 규칙을 작성, 배포하고 Snyk CLI를 사용하여 파일을 스캔하는 종단 간 흐름</p></figcaption></figure>

## Snyk 플랫폼 정책 및 {{Snyk IaC}} 사용자 정의 규칙

{% hint style="info" %}
요약:

- Snyk 플랫폼 정책: 이슈 관리
- {{Snyk IaC}} 사용자 정의 규칙: 이슈 생성
{% endhint %}

Snyk 플랫폼을 통해 Snyk가 스캔 중 식별하는 이슈를 어떻게 우선 순위를 정하고 처리할지를 관리하는 [정책](../../../../manage-risk/policies/)을 직접 만들 수 있습니다. 예를 들어, 특정 속성을 가진 이슈의 우선 순위를 중에서 고로 변경하거나 특정 기준을 충족하는 경우 이슈를 대량으로 무시하는 정책을 정의할 수 있습니다.

{Snyk IaC}} 사용자 정의 규칙 기능을 사용하면 강제하고자 하는 잘못된 구성을 위한 자체 규칙을 정의할 수 있습니다. 구성 파일에서 사용자 정의 규칙 실패의 결과는 이슈 생성입니다.