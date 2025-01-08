# Snyk API

{% hint style="info" %}
Snyk API는 엔터프라이즈 플랜에서만 이용 가능합니다.&#x20;

자세한 정보는 [Plans and pricing](https://snyk.io/plans)를 참조하십시오.
{% endhint %}

Snyk API를 사용하면 엔터프라이즈 고객이 Snyk을 프로그래밍 방식으로 통합할 수 있습니다.

Snyk API는 개발자가 특정 워크플로우를 수행하기 위해 Snyk 프로세스를 자동화할 수 있도록 해줍니다. 이를 통해 개발자 경험과 플랫폼 거버넌스의 일관성을 보장할 수 있습니다. [Snyk REST API](rest-api/about-the-rest-api.md)와 [V1 API](v1-api.md)는 CLI나 통합 대신 API를 사용할 결정을 내렸을 때 사용할 수 있습니다. 두 API는 Snyk API [Reference](reference/)에서 이용 가능합니다. 또한 [OAuth2 API 레퍼런스](oauth2-api.md)에서 추가 엔드포인트를 이용할 수 있습니다.

Snyk 프로세스를 **맞춤화, 통합 및 자동화**하고자 할 때 API를 사용하십시오.

API, CLI 및 통합의 **출력에는 차이가 있을 수 있습니다**.

예를 들어 많은 패키지 관리자의 경우 API를 사용하면 빌드 파이프 또는 로컬 패키지에서 Snyk CLI를 실행하는 것보다 정확성이 떨어질 수 있습니다. 패키지 요구 사항에 여러 버전의 패키지가 맞을 수 있습니다. CLI를 로컬에서 실행하면 실제로 배포된 코드를 테스트하고 사용 중인 의존성 버전에 대한 정확한 스냅샷을 만들 수 있습니다. API는 정확성이 더 낮은 스냅샷을 추론합니다. Snyk CLI는 기계 판독 가능한 JSON(`snyk test --json`)을 출력할 수 있습니다.

Snyk 통합을 사용하여 개발 플로에 Snyk 액세스할 수 있습니다. 이점은 모든 새 풀 리퀘스트를 모니터하고 수정 사항을 제안하여 새 풀 리퀘스트를 열 수 있다는 것입니다. Snyk을 소스 코드 관리(SCM) 도구에 직접 통합하거나 Broker를 사용하여 보다 높은 보안과 감사 기능을 이용할 수 있습니다.