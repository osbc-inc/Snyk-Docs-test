# 가시성 설정 및 조직 템플릿 구성

하나의 조직을 만들거나 여러 조직을 생성하기 위한 템플릿을 만들고 싶다면, 초기 설정을 구성해야 합니다.

{% hint style="info" %}
프로젝트를 가져오려면 [프로젝트 가져오기](../../phase-3-gain-visibility/import-projects.md) 및 [전개](../../phase-5-initial-rollout-to-team/) 논의를 참조하십시오.
{% endhint %}

## 템플릿을 사용하여 조직 구조 생성

새 조직을 만들 때, 기존 조직을 설정 및 통합을 위한 모델로 사용할 수 있습니다. Snyk 조직 구조를 만들기 전에 템플릿 조직을 구성하는 것을 권장합니다.

템플릿 기반 접근 방식은 모든 조직마다 통합을 수동으로 구성하는 번거로움을 피할 수 있도록 해줍니다.

Snyk에서 별도의 템플릿 기능은 없습니다. 권장하는 프로세스는 **템플릿**이라는 조직을 생성하고, 이 조직을 완전히 구성한 후, 더 많은 조직을 만들 때 **템플릿** 조직에서 설정을 복제할 수 있는 옵션을 사용하는 것입니다.

{% hint style="info" %}
**API를 이용한 템플릿 생성**\
Snyk를 사용하여 조직을 만드는 경우, 예를 들어 GitHub 조직과 같은 기존 소스에서 조직을 복제하려는 경우에는 [snyk-api-import](../../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/) 도구를 사용하거나, [새 조직 생성](../../../../snyk-api/reference/organizations-v1.md#org) 엔드포인트를 이용하여 `sourceOrgId`를 제공하여 API를 사용하여 템플릿을 생성할 수 있습니다.
{% endhint %}

## 템플릿 조직 설정 구성

템플릿 조직에서는 전체 조직 구조를 만들 때 복사할 수 있는 여러 설정을 구성할 수 있습니다:

* 모든 관련 통합, 예를 들어 GitHub Enterprise, Docker Hub.\
  참고: 온프레미스 소스 코드 관리 도구를 사용하는 경우, 통합을 활성화하기 위해 [Snyk Broker](../../../../enterprise-setup/snyk-broker/)를 구성하고 실행해야 합니다.
* 통합 설정, 예를 들어 Snyk가 PR에서 테스트를 실행할지 여부를 구성하는 것 등.
  * 새 Git 리포지토리 통합의 기본 설정에는 Snyk가 새로 생성된 PR에서 테스트를 실행하고, 새 취약점을 발견하면 자동으로 PR을 생성할 수 있는 옵션이 포함되어 있습니다. Snyk는이러한 설정을 초기에 비활성화하고, [예방 단계](../../phase-6-rolling-out-the-prevention-stage/)에서 이러한 기능을 도입할 준비가 되었을 때 이를 활성화하는 것을 권장합니다.
  * 다음 [통합](configure-integrations.md) 섹션에서 복사하기 전에 템플릿에 추가할 수 있는 통합에 대해 설명합니다.
* 제품 설정(예: Snyk Code 활성화)을 예로 들 수 있습니다.
  * 새 조직에는 기본적으로 가 비활성화되어 있습니다.
  * Git 리포지토리 통합을 통해 가져온 저장소에서 SAST 스캔을 활성화하려면, 프로젝트를 가져오기 전에 를 활성화하도록 토글을 사용하십시오.
  * **프로젝트를 가져오기 전에 활성화하는 것을 잊어버린 경우** Snyk Code를 활성화하고 **Git에서 코드를 다시 가져와야** 합니다. [Snyk Code 활성화](enable-snyk-code.md)를 참조하십시오.
* 알림 설정 (이메일 알림)
  * Snyk는 사용자가 프로젝트를 가져오는 중에 많은 수의 알림을 받지 않도록하기 위해 초기에 모든 그룹 및 조직 이메일 알림을 비활성화하는 것을 권장합니다.
  * 새 조직의 경우 그룹 수준에서 알림을 비활성화하고, 기존 조직의 경우 조직 수준에서 알림을 비활성화하십시오.

다음 표는 웹 인터페이스나 API를 사용하여 새 조직을 만들 때 템플릿 조직에서 새 조직으로 복사되는 내용을 보여줍니다.

<table><thead><tr><th width="394">모든 통합과 그 설정은 복사됩니다</th><th width="466">다음 항목은 복사되지 않습니다</th></tr></thead><tbody><tr><td>Source control 통합</td><td>Snyk 서비스 계정</td></tr><tr><td>Container registry 통합</td><td>멤버</td></tr><tr><td>Container orchestrators 통합 (Kubernetes)</td><td>프로젝트</td></tr><tr><td>PaaS 및 Serverless 통합</td><td>알림 선호도</td></tr><tr><td>알림 통합 (Slack 및 Jira)</td><td></td></tr><tr><td>정책</td><td></td></tr><tr><td>Ignore 설정</td><td></td></tr><tr><td>언어 설정</td><td></td></tr><tr><td>Infrastructure as code (IaC) 설정</td><td></td></tr><tr><td>Snyk Code 설정</td><td></td></tr></tbody></table>

## 템플릿 조직 복제

조직이 생성되고 구성되면, 나머지 구조를 만들기 위해 새 조직을 만들 때 이를 템플릿으로 사용할 수 있습니다.
