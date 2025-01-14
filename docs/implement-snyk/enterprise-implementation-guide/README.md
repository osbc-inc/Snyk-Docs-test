# Enterprise 구현 가이드

각 비즈니스와 환경은 다릅니다. 그것을 염두에 두고, 본 안내서는 기업이 Snyk을 구현하는 데 도움을 주려는 목적으로 작성되었습니다. 이 안내서는 대규모 롤아웃을 구현하는 데 필요한 단계에 중점을 두며 이상적인 롤아웃에 도달하는 데 필요한 권장사항을 제공합니다.

이 안내서는 대부분의 비즈니스가 다음을 인식하고 시작한다고 가정합니다:

* 기존 소프트웨어에서의 문제 누적이 있습니다.
* 지속적으로 새로운 소프트웨어를 만들고 새 코드를 보호해야 합니다.

{% hint style="info" %}
비즈니스의 규모와 범위에 따라 구현에 대한 **일반적인 일정**이 있습니다.

규모가 작고 민첩한 비즈니스라면 Snyk 구현은 며칠 안에 완료할 수 있습니다. 구매 후 곧바로 Git 통합과 [API 가져오기 도구](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/)를 사용하여 자주 Snyk로 스캔을 시작할 수 있습니다. 이 유형의 프로세스에 대한 자세한 내용은 [시작하기](../../getting-started/) 및 [스캔 시작](../../scan-with-snyk/start-scanning.md) 섹션을 참조하십시오.

그러나 더 크고 프로세스 지향적인 기업의 경우 구현 과정은 몇 주 또는 몇 달이 걸리며 성공을 위해 더 자세한 계획이 필요합니다.
{% endhint %}

Snyk 앱위험 기본 플랜은 Snyk 엔터프라이즈 플랜에 포함되어 있어 다음 기능을 이용할 수 있습니다:

* [SCM의 커버리지 제어](../../manage-risk/policies/assets-policies/use-cases-for-policies/coverage-and-coverage-gap-policies.md).
* 특정 조치를 자동으로 트리거하기 위한 [정책](../../manage-risk/policies/assets-policies/) 생성.
* SCM 통합을 위한 [Backstage 파일](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/) 사용자 정의.
* [맞춤형 분석](../../manage-risk/enterprise-analytics/) 및 응용 프로그램에 관한 보고서 제공.

{% hint style="info" %}
Snyk 앱위험 프로로 플랜 업그레이드를 원하시면 영업 담당자에게 문의하십시오.\
[Snyk AppRisk Essentials vs Snyk AppRisk Pro ](../../scan-with-snyk/snyk-apprisk/snyk-apprisk-essentials-vs-snyk-apprisk-pro.md)페이지에서 두 제공품 간의 사용 가능한 기능에 대한 자세한 내용을 확인할 수 있습니다.
{% endhint %}

## 구현 전략 개요

본 안내서는 세 가지 목표와 일치하는 일련의 조치를 개요로 한 여러 단계로 구성되어 있습니다:

* [가시성 달성](./#achieve-visibility)
* [예방 달성 및 개발자 매력성 증진](./#achieve-prevention-and-drive-developer-adoption)
* [백로그 수정 및 문제 처리](./#fix-the-backlog-and-triage-issues)

### 가시성 달성

대규모 기업의 경우 Snyk은 보안 문제에 대한 명확한 인식을 얻는 것을 우선 권장하며 항상 이를 해결할 필요는 없다고합니다.

{% hint style="info" %}
이는 Snyk을 사용하여 문제를 해결하는 것을 방해하지 않습니다. 초기에 문제를 해결할 수 있지만 초반에 개발을 방해하지 않고 신뢰를 구축하고 나중에 일반적으로 예방 단계에서 천천히 차폐를 도입하는 것이 중요합니다.
{% endhint %}

가시성은 응용 프로그램 포트폴리오 전체의 보안을 넓게 파악하여 Snyk 스캔이 개발 프로세스에 미치는 영향을 최소화합니다.

이 가시성은 Snyk 롤아웃 과정 중 신뢰를 쌓는 데 도움이 됩니다.

### 예방 달성 및 개발자 매력성 증진

다음은 예방 단계입니다. 응용 프로그램에 새로운 보안 문제가 추가되는 것을 막습니다. 이 단계에서 개발자가 PR(풀 리퀘스트)/MR(마지막 리퀘스트) 확인 및 파이프라인에서 확인을 통해 문제를 볼 수 있도록 제어를 마련할 수 있습니다.

이 프로세스의 일환으로 개발자는 IDE 플러그인 및 [Snyk Advisor](https://snyk.io/advisor)와 같은 다른 도구를 사용하여 안전한 패키지를 선택하고 [Snyk Learn](https://learn.snyk.io/)를 통해 안전한 코딩, 보안 및 제품에 대해 교육을 받을 수 있습니다.

### 백로그 수정 및 문제 처리

마지막으로 보안 문제의 백로그 수정에 초점을 맞출 수 있습니다. 이는 여러 형태로 구성될 수 있습니다:

* 초기 롤아웃 일환으로 보안팀 또는 초기 이해관계자가 기존 포트폴리오의 초기 결과를 분류하고 조사하거나 처리해야 할 우선 순위 항목에 대한 티켓을 만들거나 팀이 주간 문제 처리 프로세스의 일환으로 자신의 응용 프로그램에 대해 그 작업을 수행할 수 있습니다.
* 가시성을 얻고 예방을 달성한 후에는 문제 백로그를 살펴볼 수 있습니다. 예를 들어, 핵심 이해관계자들과 함께 하는 주간 문제 처리 프로세스는 팀에게 어떤 문제를 해결해야 하는지 안내합니다.

## Snyk와 함께 향상된 리소스 사용

Snyk는 개발자를 염두에 두고 제작되어 다음과 같은 기능을 제공합니다:

* IDE, Git 및 CI/CD 통합을 위한 도구를 사용하여 안전한 응용 프로그램 생성.
* [Snyk Advisor](https://snyk.io/advisor) 및 의사 결정을 지원하는 다른 도구.
* 제품, 코드 보호 및 최상의 실행 사례에 대한 교육 자료인 [Snyk Learn](https://learn.snyk.io).
* 보안 및 규정 준수 팀이 방향을 제공할 수 있도록 하는 [정책](../../manage-risk/policies/).

**bold** and _italic_
