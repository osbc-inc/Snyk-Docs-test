# Nexus 리포지토리 관리자 연결 설정

{% hint style="info" %}
**기능 가용성**\
패키지 저장소 통합 기능은 엔터프라이즈 플랜에서 이용 가능합니다. 자세한 정보는 [요금제 및 가격정책](https://snyk.io/plans/)을 참조하십시오.

**지원되는 프로젝트**\
Nexus 리포지토리 관리자 통합은 [Node.js](../../../../supported-languages-package-managers-and-frameworks/javascript/#supported-frameworks-and-package-managers) (npm 및 Yarn) 및 [Maven](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-frameworks-and-package-managers) 프로젝트를 지원합니다. [개선된 Gradle SCM 스캐닝](../../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning-early-access)을 위해 Maven 설정을 사용하십시오.
{% endhint %}

Nexus 리포지토리 관리자에 연결하면 Snyk이 Nexus 레지스트리에 호스팅된 패키지의 모든 직접 및 간접 종속성을 해결하고 더 완전하고 정확한 종속성 그래프 및 관련 취약점을 계산할 수 있습니다.

다음 유형의 Nexus 리포지토리 관리자를 구성할 수 있습니다:

* 기본 인증으로 보호된 공개 접근 가능 인스턴스
* Snyk Broker를 사용하여 사설 네트워크에 있는 인스턴스 (기본 인증 사용 또는 미사용)

{% hint style="info" %}
**지원되는 버전**

* Nexus 리포지토리 관리자 버전 3.0+
* Nexus 리포지토리 관리자 버전 2.15+
{% endhint %}

이 지침은 공개 접근 가능 인스턴스를 구성하는 데 해당됩니다. 브로커 인스턴스를 구성하는 지침은 [Nexus 리포지토리 관리자용 Snyk Broker 설치 및 구성](../../../../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/)을 참조하십시오.

## 공개 접근 가능 인스턴스 설정

1. **설정** > **통합 > 패키지 저장소 > Nexus**로 이동합니다.
2. Nexus를 구성할 수 있는 화면이 표시되는지 확인합니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 15.15.11.png" alt="Nexus 구성"><figcaption><p>Nexus 구성</p></figcaption></figure>

Nexus를 구성할 페이지에서 사용 중인 버전에 대한 정보를 입력합니다.

{% tabs %}
{% tab title="Nexus 3" %}
* Nexus 인스턴스의 URL을 입력합니다; 이는 `/repository`로 끝나야 합니다.
* 사용자 이름을 입력합니다.
* 암호를 입력합니다.
* **저장**을 클릭합니다.
{% endtab %}

{% tab title="Nexus 2" %}
* Nexus 인스턴스의 URL을 입력합니다; 이는 `/nexus/content`로 끝나야 합니다.
* 사용자 이름을 입력합니다.
* 암호를 입력합니다.
* **저장**을 클릭합니다.
{% endtab %}
{% endtabs %}

## 역방향 프록시를 사용한 Nexus

Nexus 서버가 역방향 프록시(예: Nginx) 뒤에서 실행 중인 경우, URL은 기본 `/repository` (Nexus 3) 또는 `/nexus/content` (Nexus 2)로 끝나지 않을 수 있습니다. 이는 역방향 프록시에서 구성된 경로에 따라 달라집니다. 이 경우, 역방향 프록시에서 구성된 URL을 사용해야 합니다.

예: Nexus 3의 경우, `http://nexus.company.io/repository`가 `http://nexus.company.io/my-company/my-repository`으로 매핑된 경우 `http://nexus.company.io/my-company/my-repository`를 사용합니다.

예: Nexus 2의 경우, `http://nexus.company.io/nexus/content`가 `http://nexus.company.io/my-nexus-content`로 매핑된 경우 `http://nexus.company.io/my-nexus-content`를 사용합니다.

{% hint style="info" %}
Snyk이 저장소에 연락할 수 있는 경우 녹색 성공 메시지가 표시됩니다.

노란색 경고 메시지가 표시되면 URL 및 자격 증명을 확인하고 다시 시도하십시오.
{% endhint %}

## Nexus 사용자 권한

Nexus 사용자에게 다음 권한이 필요합니다(Role의 일부 또는 개별적으로 추가됨):

{% tabs %}
{% tab title="Nexus 3" %}
* `nx-metrics-all` ([시스템 상태 확인 엔드포인트](https://support.sonatype.com/hc/en-us/articles/226254487-System-Status-and-Metrics-REST-API)를 위한)
* `nx-repository-view-[*-* | <ecosystem-repo-name>]-read`
* `nx-repository-view-[*-* | <ecosystem-repo-name>]-browse`
{% endtab %}

{% tab title="Nexus 2" %}
* `Status - Read`
* `All [<ecosystem>] Repositories - (read)`
* `[All Repositories | <repoName>] - (view)`
{% endtab %}
{% endtabs %}
