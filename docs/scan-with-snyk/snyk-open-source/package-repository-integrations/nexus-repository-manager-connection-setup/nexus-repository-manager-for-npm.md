# npm용 Nexus 저장소 관리자

{% hint style="info" %}
**기능 가용성**\
패키지 저장소 통합은 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 내용은 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.

이 가이드는 Snyk 웹 UI 통합에 관련되며, Snyk CLI는 비공개 Nexus Repository Manager 레지스트리를 지원하는 Yarn 및 npm 프로젝트를 지원합니다.
{% endhint %}

Snyk은 Git에서 가져온 npm 및 Yarn 프로젝트에서 Nexus Repository Manager를 사용할 수 있습니다.

이를 통해 Snyk은 Pull/Merge Requests를 생성할 때 올바른 URL로 lock 파일을 다시 생성할 수 있습니다.

Snyk에게 어디에 개인 Nexus Repository Manager Node.js 패키지가 호스팅되었는지와 어떤 스코프 하에 있는지 알려주는 구성을 추가할 수 있습니다.

이는 일반적으로 `.yarnrc` 또는 `.npmrc`에 추가하는 정보와 동일합니다.

## JavaScript 언어 설정

**Settings > Languages > JavaScript**로 이동하여 프로젝트 유형에 따라 npm 또는 Yarn 설정 중 하나를 선택하십시오.

이전에 Nexus Repository Manager에 연결하지 않았다면 먼저 통합을 구성하라는 메시지가 표시됩니다. [Nexus Repository Manager 연결 설정](./)을 참조하십시오.

아래 탭을 따라서 당신의 Nexus 버전에 맞게 단계를 따르십시오.

![Nexus 레지스트리 구성](<../../../../.gitbook/assets/Screenshot 2022-07-15 at 14.18.43.png>)

{% tabs %}
{% tab title="Nexus 3" %}
1. **레지스트리 구성 추가**를 선택합니다.
2. 패키지 소스로 **Nexus**를 선택합니다.
3. 이 레지스트리를 **기본 레지스트리 URL**로 구성하려면 **범위**를 비워 둡니다.
4. **이 레지스트리를 사용하는 패키지만 구성**하려면 **범위**를 추가합니다.
5. **기본 레지스트리 URL** 및 **범위 패키지**를 추가하려면 기본값을 하나 추가하고 각 스코프당 하나의 구성을 추가합니다.
6. **리포지토리** 섹션은 내부 리포지토리 URL에서 `repository/` 뒤에 오는 항목으로 설정해야 합니다.\
   예를 들어, URL이 `http://nexus.company.io/repository/npm-group`인 경우, **리포지토리**를 `npm-group`으로 설정해야 합니다.
7. 원하는 모든 레지스트리와 범위를 추가한 후 **설정 업데이트**를 클릭합니다.
{% endtab %}

{% tab title="Nexus 2" %}
1. **레지스트리 구성 추가**를 선택합니다.
2. 패키지 소스로 **Nexus**를 선택합니다.
3. 이 레지스트리를 **기본 레지스트리 URL**로 구성하려면 **범위**를 비워 둡니다.
4. **이 레지스트리를 사용하는 패키지만 구성**하려면 **범위**를 추가합니다.
5. **기본 레지스트리 URL** 및 **범위 패키지**를 추가하려면 기본값을 하나 추가하고 각 스코프당 하나의 구성을 추가합니다.
6. **리포지토리** 섹션은 내부 리포지토리 URL에서 `nexus/content/` 뒤에 오는 항목으로 설정해야 합니다.\
   예를 들어, URL이 `http://nexus.company.io/nexus/content/groups/npm-group`인 경우, **리포지토리**를 `groups/npm-group`으로 설정해야 합니다.\
   또는 `http://nexus.company.io/nexus/content/repositories/npm-hosted`인 경우, **리포지토리**는 `repositories/npm-hosted`로 설정되어야 합니다.
7. 원하는 모든 레지스트리와 범위를 추가한 후 **설정 업데이트**를 클릭합니다.
{% endtab %}
{% endtabs %}

## 테스트해보기

Nexus에 호스팅된 비공개 종속성이 포함된 프로젝트에서 Pull/Merge Request를 열어보세요. 그러면 **레지스트리로 올바른 URL이 포함된 lock 파일이 Snyk Fix Pull Request에 업데이트되어 표시됩니다.**

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 14.22.59.png" alt="Nexus 통합 테스트를 위한 Pull Request"><figcaption><p>Nexus 통합 테스트를 위한 Pull Request</p></figcaption></figure>