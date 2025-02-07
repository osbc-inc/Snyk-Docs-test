# Maven용 Nexus 리포지토리 관리자

{% hint style="info" %}
**기능 사용 가능성**\
패키지 저장소 통합은 엔터프라이즈 플랜에서 사용할 수 있습니다. 자세한 정보는 [요금제 및 가격정책](https://snyk.io/plans/)을 참조하세요.

**지원되는 프로젝트**\
Nexus 리포지토리 관리자 통합은 [Node.js](../../../../supported-languages-package-managers-and-frameworks/javascript/#supported-frameworks-and-package-managers) (npm 및 Yarn) 및 [Maven](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-frameworks-and-package-managers) 프로젝트를 지원합니다. [개선된 Gradle SCM 스캐닝](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning-early-access)을 위해 Maven 설정을 사용하세요.
{% endhint %}

Snyk은 Maven 프로젝트와 함께 Nexus 리포지토리 관리자를 사용할 수 있습니다.

이를 통해 Snyk은 사용자 지정 레지스트리에 호스팅된 패키지의 모든 직접 및 간접 종속성을 해결하고 더 완전하고 정확한 종속성 그래프 및 관련된 취약점을 계산할 수 있습니다.

Maven 프로젝트를 구성하여 모든 요청을 사용자 지정 패키지 저장소를 통해 미러링하거나 Maven Central과 함께 사용할 추가 저장소를 지정할 수 있습니다.

## **사용자 지정 Maven 패키지 레지스트리 설정**

Nexus 레지스트리에 접근하기 위해 인증이 필요한 경우 먼저 Nexus 리포지토리 관리자 통합을 구성해야 합니다. [Nexus 리포지토리 관리자 설정](./)을 참조하세요.

Nexus를 미러 또는 추가 리포지토리로 사용할 수 있습니다.

이러한 설정은 `~/.m2/settings.xml`에 있는 것과 매우 유사합니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 15.10.52.png" alt="미러 설정"><figcaption><p>미러 설정</p></figcaption></figure>

`Type`에 값을 선택하십시오. **직접(Direct)** 또는 인증을 사용하는 경우 **통합(Integration)** 중 하나를 선택하십시오.

**직접(Direct)**&#xB97C; 사용하는 경우 **URL**, **저장소 이름** 및 **거울 대상**을 완료해야 합니다.

**거울 대상(Mirror Of)** 값은 모든 것을 미러링하기 위해 `*`가 될 수 있거나 예를 들어 `central`과 같은 값을 입력할 수 있습니다.

**통합(Integration)** 유형을 사용하는 경우 Nexus 통합 유형을 선택하고 **저장소 이름** 및 **거울 대상** 세부 정보를 제공해야 합니다.

**저장소 이름(Repository Name)**&#xC744; Nexus 버전에 따라 설정하십시오.

{% tabs %}
{% tab title="Nexus 3" %}
내부 저장소 URL에서 `repository/` 뒤에 나오는 내용을 **저장소 이름**으로 설정하십시오.

예를 들어, URL이 `http://nexus.company.io/repository/libs-release`인 경우, 저장소 이름은 `libs-release`여야 합니다.
{% endtab %}

{% tab title="Nexus 2" %}
내부 저장소 URL에서 `nexus/content/` 뒤에 나오는 내용을 **저장소 이름**으로 설정하십시오.

예를 들어, URL이 `http://nexus.company.io/nexus/content/groups/public`인 경우, 저장소 이름은 `groups/public`이어야 합니다.

URL이 `http://nexus.company.io/nexus/content/repositories/releases`인 경우, 저장소 이름은 `repositories/releases`여야 합니다.
{% endtab %}
{% endtabs %}

또는 아티팩트를 확인할 추가 위치로 사용될 저장소를 구성할 수 있습니다.

저장소는 [미러](nexus-repository-manager-for-maven.md#mirrors)와 동일한 방식으로 구성되지만 **거울 대상**이 필요하지 않습니다.
