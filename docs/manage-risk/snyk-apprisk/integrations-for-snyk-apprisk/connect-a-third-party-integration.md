# Snyk 앱 리스크를 위한 서드 파티 통합

Integrations 페이지는 자동으로 동기화된 기존 Snyk 조직 데이터를 포함하여 모든 활성 통합을 보여주며, Integration Hub에 액세스할 수 있습니다.

{% hint style="info" %}
로드된 패키지 리스크 팩터는 데비안 패키지와 같은 운영 체제 패키지에 대해 Snyk로 지원되지 않습니다. 이 기능은 npm, Maven 또는 PyPI와 같은 패키지 관리자에 호스팅된 패키지에만 적용됩니다.
{% endhint %}

AppRisk 통합을 **Integration Hub**에서 사용자 정의할 수 있으며, 다음 통합이 제공됩니다:

* [Veracode SAST](connect-a-third-party-integration.md#veracode-setup-guide)
* [Checkmarx SAST](connect-a-third-party-integration.md#checkmarx-setup-guide)
* [SonarQube](connect-a-third-party-integration.md#sonarqube-setup-guide)
* [Nightfall](connect-a-third-party-integration.md#nightfall-setup-guide)
* [GitGuardian](connect-a-third-party-integration.md#gitguardian-setup-guide)
* [Dynatrace](connect-a-third-party-integration.md#dynatrace-setup-guide)
* [Sysdig](connect-a-third-party-integration.md#sysdig-setup-guide)

{% hint style="info" %}
새 통합 설정으로부터 **Connected** 상태를 받은 후 최대 2시간이 소요될 수 있습니다.
{% endhint %}

통합을 설정한 후에는 연결 상태가 표시됩니다.

<figure><img src="../../../.gitbook/assets/image (633).png" alt=""><figcaption><p>서드 파티 통합 - 연결된 상태</p></figcaption></figure>

## Veracode 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk Pro용 Veracode는 조기 액세스 단계이며, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

### 전제 조건 <a href="#verocode-prerequisites" id="verocode-prerequisites"></a>

Veracode 애플리케이션 컨셉은 Snyk AppRisk 저장소 자산과 일치합니다. [Veracode API](https://app.swaggerhub.com/apis/Veracode/veracode-applications_api_specification/1.0#/Application%20information%20API/updateApplicationUsingPUT)를 사용하여 Veracode 사용자 정의 필드를 만들고 활용해야 합니다. 자세한 내용은 [Veracode 사용자 정의 메타데이터 필드](https://docs.veracode.com/r/t_create_custom_metadata)에서 확인하십시오.

다음과 같이 repoURL이라고 하는 사용자 정의 필드를 추가해야 합니다:

```
{
"name": "repoURL", 
"value": <YOUR GITHUB URL>
}

```

### 필수 매개변수 <a href="#veracode-required-parameters" id="veracode-required-parameters"></a>

* API ID와 API Key - 사용자 계정과 관련된 API 자격 증명. 자세한 내용은 [Veracode API 자격 증명](https://help.veracode.com/r/c_api_credentials3) 링크를 참조하십시오.

### Integration Hub 설정 <a href="#veracode-integration-hub-setup" id="veracode-integration-hub-setup"></a>

1. **Integration Hub** 메뉴를 엽니다.
2. **SAST** 태그를 선택하고 Veracode를 검색합니다.
3. **Add** 버튼을 클릭합니다.
4. 이 통합을 위한 프로필 이름을 추가하세요.
5. Veracode 계정에서 **API ID**를 추가합니다.
6. Veracode 계정에서 **API 키**를 추가합니다.
7. **Done** 버튼을 클릭합니다.
8. 연결이 설정되면 Veracode 통합의 상태가 **Connected**로 변경됩니다.

<figure><img src="../../../.gitbook/assets/image (589).png" alt=""><figcaption><p>Veracode - 설정 화면</p></figcaption></figure>

## Checkmarx 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk Pro용 Checkmarx는 조기 액세스 단계이며, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

Checkmarx SAST 통합을 설정하는 데 다음 지침을 사용합니다. Checkmarx SAST 통합은 Checkmarx SAST에만 작동하며, Checkmarx One은 아직 지원되지 않습니다.

{% hint style="warning" %}
Snyk AppRisk Pro는 Checkmarx One 통합을 지원하지 않습니다.
{% endhint %}

### 전제 조건 <a href="#checkmarx-prerequisites" id="checkmarx-prerequisites"></a>

* [Snyk Broker](../../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md#checkmarx-sast-integration) 연결을 설치하고 구성하세요.
* Checkmarx 프로젝트에 Git 설정을 올바르게 사용하고 있어야 합니다. 자세한 내용은 Checkmarx [프로젝트의 원격 소스 설정을 GIT로 설정](https://checkmarx.stoplight.io/docs/checkmarx-sast-api-reference-guide/8312d35369b9b-set-project-s-remote-source-settings-as-git) 문서 페이지를 참조하십시오.

### 필수 매개변수 <a href="#checkmarx-required-parameters" id="checkmarx-required-parameters"></a>

1. API URL - 예: `checkmarx.customer.com`과 같은 Checkmarx API의 URL입니다.
2. 사용자 이름과 비밀번호 - Checkmarx SAST 액세스 권한이 있는 사용자 계정의 자격 증명입니다.
  
### Integration Hub 설정 <a href="#checkmarx-integration-hub-setup" id="checkmarx-integration-hub-setup"></a>

Snyk Broker를 AppRisk용으로 설치하고 구성한 후 Checkmarx SAST를 위한 연결을 성공적으로 설정한 경우, Integration Hub에서 해당 통합을 구성해야 합니다.

1. **Integration Hub** 메뉴를 엽니다.
2. **SAST** 태그를 선택하고 Checkmarx를 검색합니다.
3. **Add** 버튼을 클릭합니다.
4. 이 통합을 위한 프로필 이름을 추가하세요.
5. Snyk AppRisk Checkmarx 통합을 위한 Broker 토큰을 추가하세요.
6. Checkmarx 호스트를 추가하세요. 예: `checkmarx.customer.com`
7. **Done** 버튼을 클릭합니다.
8. 연결이 설정되면 Checkmarx 통합의 상태가 **Connected**로 변경됩니다.

<figure><img src="../../../.gitbook/assets/image (590).png" alt=""><figcaption><p>Checkmarx - 설정 화면</p></figcaption></figure>

## SonarQube 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk Pro용 SonarQube는 조기 액세스 단계이며, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

### 필수 매개변수 <a href="#sonarqube-required-parameters" id="sonarqube-required-parameters"></a>

* API Key. SonarQube API 키에 대한 자세한 내용은 [여기](https://docs.sonarsource.com/sonarqube/latest/user-guide/user-account/generating-and-using-tokens/)에서 확인하십시오.

### Integration Hub 설정 <a href="#sonarqube-integration-hub-setup" id="sonarqube-integration-hub-setup"></a>

* **Integration Hub** 메뉴를 엽니다.
* **SAST** 태그를 선택하고 SonarQube를 검색합니다.
* **Add** 버튼을 클릭합니다.
* 이 통합을 위한 **프로필 이름**을 추가하세요.
* 이 통합을 위한 **호스트 URL**을 추가하세요.
* **API 토큰**을 추가하세요. SonarQube 계정에 액세스하여 사용자, 내 계정, 보안을 선택한 후 사용자 토큰을 선택합니다. SonarQube API 키에 대한 자세한 내용은 SonarQube의 [토큰 생성 및 사용](https://docs.sonarsource.com/sonarqube/latest/user-guide/user-account/generating-and-using-tokens/) 문서 페이지에서 확인하십시오.
* **Done** 버튼을 클릭합니다.
* 연결이 설정되면 SonarQube 통합의 상태가 **Connected**로 변경됩니다.

<figure><img src="../../../.gitbook/assets/image (591).png" alt=""><figcaption><p>SonarQube - 설정 화면</p></figcaption></figure>

## Nightfall 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk Pro용 Nightfall는 조기 액세스 단계이며, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

### 필수 매개변수 <a href="#nightfall-required-parameters" id="nightfall-required-parameters"></a>

* API Key. Nightfall API 키를 생성하는 방법에 대한 자세한 내용은 Nightfall [API 키 생성](https://help.nightfall.ai/nightfall-firewall-for-ai/key-concepts/setting-up-nightfall/creating-api-key) 문서 페이지에서 확인하십시오.

### Integration Hub 설정 <a href="#nightfall-integration-hub-setup" id="nightfall-integration-hub-setup"></a>

1. **Integration Hub** 메뉴를 엽니다.
2. **Secrets** 태그를 선택하고 Nightfall을 검색합니다.
3. **Add** 버튼을 클릭합니다.
4. 이 통합을 위한 **프로필 이름**을 추가하세요.
5. 이 통합을 위한 **기본 API URL**을 추가하세요.
6. 이 통합을 위한 **API 키**를 추가하세요.
7. **Done** 버튼을 클릭합니다.
8. 연결이 설정되면 Nightfall 통합의 상태가 **Connected**로 변경됩니다.

다음 비디오는 Integration Hub에서의 Nightfall 구성 개요를 제공합니다:

{% embed url="https://www.youtube.com/watch?v=FJ5fAyMYSUs" %}
비디오가 도움이 되었나요? [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training&topics=AppRisk)에서 나머지 코스를 확인하세요!
{% endembed %}

Integration Hub을 사용하여 Nightfall 통합을 설정한 후에는 시크릿 감지 범위를 확인할 수 있습니다.

{% embed url="https://www.youtube.com/watch?v=o6TqPMSq1rk" %}
비디오가 도움이 되었나요? [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training&topics=AppRisk)에서 나머지 코스를 확인하세요!
{% endembed %}

## GitGuardian 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk Pro용 GitGuardian는 조기 액세스 단계이며, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

### 필수 매개변수 <a href="#gitguardian-required-parameters" id="gitguardian-required-parameters"></a>

* API 키. GitGuardian [인증](https://docs.gitguardian.com/api-docs/authentication) 문서 페이지에서 GitGuardian API 키를 생성하는 방법에 대한 자세한 내용을 확인하십시오.

GitGuardian API 키를 생성할 때, 해당 API 키가 서비스 계정 및 개인 액세스 토큰 모두에 대해 작동함을 기억하세요.

다음 권한이 **READ**로 설정되어 있는지 확인하세요:

* Incident (`의무 사항`)
* Teams (유료 계정의 경우 **권장**)

### Integration Hub 설정 <a href="#gitguardian-integration-hub-setup" id="gitguardian-integration-hub-setup"></a>

1. **Integration Hub** 메뉴를 엽니다.
2. **Secrets** 태그를 선택하고 GitGuardian을 검색합니다.
3. **Add** 버튼을 클릭합니다.
4. 이 통합을 위한 **프로필 이름**을 추가하세요.
5. 이 통합을 위한 **API 토큰**을 추가하세요.
6. **Done** 버튼을 클릭합니다.
7. 연결이 설정되면 GitGuardian 통합의 상태가 **Connected**로 변경됩니다.

다음 비디오는 Integration Hub에서의 GitGuardian 구성 개요를 제공합니다:

{% embed url="https://www.youtube.com/watch?v=4u4QrJBZTkI" %}
비디오가 도움이 되었나요? [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training&topics=AppRisk)에서 나머지 코스를 확인하세요!
{% endembed %}

Integration Hub을 사용하여 GitGuardian 통합을 설정한 후에는 시크릿 감지 범위를 확인할 수 있습니다:

{% embed url="https://www.youtube.com/watch?v=zh4c5f_vv1k" %}
비디오가 도움이 되었나요? [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training&topics=AppRisk)에서 나머지 코스를 확인하세요!
{% endembed %}

## Dynatrace 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk Pro용 Dynatrace는 조기 액세스 단계이며, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

Dynatrace 런타임 통합에서 보고되는 다음 리스크 팩터는 [배포됨](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md) 및 [로드된 패키지](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md)입니다.

{% hint style="info" %}
Dynatrace 통합에서 보고되는 로드된 패키지 리스크 팩터를 위한 지원되는 언어에 대한 자세한 내용은 Dynatrace [지원되는 기술](https://docs.dynatrace.com/docs/platform-modules/application-security/getting-started/get-started-with-application-security#tech) 페이지에서 확인하십시오.
{% endhint %}

### 전제 조건 <a href="#dynatrace-prerequisites" id="dynatrace-prerequisites"></a>

* DPS 라이선스 모델에서 Dynatrace SaaS를 사용하세요.
* Dynatrace [Kubernetes 앱](https://docs.dynatrace.com/docs/platform-modules/infrastructure-monitoring/container-platform-monitoring/kubernetes-app/overview)은 적어도 하나의 클러스터를 모니터링하도록 구성되어 있어야 합니다.
* 사용자가 엔티티 모델을 쿼리할 수 있는## Dynatrace 설정 페이지

이후, `Identity & access management`으로 이동하십시오. `OAuth clients`를 선택하고 `Create client`를 클릭하십시오. 세부 정보를 입력하고 다음 권한을 확인한 후 `Create client`를 클릭하십시오:

```yaml
storage:entities:read
```

5. 나중을 위해 Client ID와 Client secret를 저장하고 `Finish`를 클릭하십시오.

### 필요한 매개변수

1. **계정 UUID** - Dynatrace 계정의 `account-uuid`.
2. **환경 ID** - Dynatrace에서 모니터링되는 환경의 ID.
3. **OAuth 클라이언트 ID** - 프리퀘시트에서 생성된 OAuth 클라이언트의 ID.
4. **OAuth 클라이언트 시크릿** - 프리퀘시트에서 생성된 OAuth 클라이언트의 시크릿.

### 통합 허브 설정

* **Integration Hub** 메뉴를 엽니다.
* **Runtime** 태그를 선택하고 **Dynatrace**를 검색합니다.
* **Add** 버튼을 클릭합니다.
* 통합 프로필 이름을 편집합니다.
* **계정 UUID**를 입력합니다.
* **환경 ID**를 입력합니다.
* **OAuth 클라이언트 ID**를 입력합니다.
* **OAuth 클라이언트 시크릿**를 입력합니다.
* **Done** 버튼을 클릭합니다.
* 연결이 설정되면 **Dynatrace** 통합 상태가 **Connected**로 변경됩니다.

{% hint style="info" %}
런타임 통합에서 Dynatrace 런타임 데이터가 제공되면, Snyk AppRisk에 몇 시간 내에 표시됩니다.
{% endhint %}

![Dynatrace - 설정 화면](../../../.gitbook/assets/image (592).png)

## Sysdig 설정 가이드

{% hint style="info" %}
**릴리스 상태**

Sysdig for Snyk AppRisk Pro는 이른 액세스 상태로, Snyk Enterprise 플랜 및 Snyk AppRisk Pro와 함께만 사용할 수 있습니다. 그룹에서 설정하려면, Snyk 계정 팀에 연락하십시오.
{% endhint %}

Sysdig 런타임 통합으로부터 보고된 다음 [위험 요인](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk#risk-factors): [배포](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed.md) 및 [로드된 패키지](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md).

{% hint style="info" %}
로드된 패키지 위험 요인에 대한 지원되는 언어는 다음과 같습니다: Go, Java, JavaScript/TypeScript 및 Python.
{% endhint %}

### 전제 조건

* 계정은 Sysdig Secure 제품에 액세스해야 합니다.
* 사용 중인 패키지 기능 플래그를 활성화하려면 Sysdig 담당자에게 문의하십시오.

### 필수 매개변수

* **서비스 계정 API 토큰** - API 토큰을 얻으려면 Sysdig 서비스 계정을 생성하는 방법에 대한 [서비스 계정 설정 지침 페이지](https://docs.sysdig.com/en/docs/administration/administration-settings/access-and-secrets/user-and-team-administration/manage-teams-and-roles/#service-accounts)로 이동하십시오.
  * 이 서비스 계정에 **View Only**를 롤로 설정합니다.
  * 서비스 계정에 **만료 날짜**를 설정합니다. 서비스 계정의 만료 후에는 새로운 서비스 계정으로 업데이트되기 전까지 Sysdig 통합이 정보를 가져올 수 없습니다.

{% hint style="info" %}
생성된 **서비스 계정**은 **Sysdig Monitor**가 아닌 **Sysdig Secure** 하에 있어야 합니다.
{% endhint %}

* **지역 -** Sysdig [SAAS 지역 및 IP 범위](https://docs.sysdig.com/en/docs/administration/saas-regions-and-ip-ranges) 페이지로 이동하여 Sysdig 지역 URL에 대한 자세한 내용을 확인하십시오.

### 알려진 제한 사항

* 만약 클러스터의 모든 노드에 Sysdig Agent가 배포되지 않은 경우, 이 통합에서 제공되는 런타임 데이터가 불완전할 수 있습니다.
* 다양한 Sysdig 스캔이 다른 간격으로 실행되므로, 클러스터 내 리소스에 변경을 적용하고 이 정보를 통합을 통해 보고하는 사이에 지연이 발생할 수 있습니다.

### 통합 허브 설정

* **Integration Hub** 메뉴를 엽니다.
* **Runtime** 태그를 선택하고 Sysdig를 검색합니다.
* **Add** 버튼을 클릭합니다.
* 이 통합을 위한 **프로필 이름**을 추가합니다.
* **계정 API 토큰**을 추가합니다.
* **Sysdig 지역**을 설정합니다.
* **Done** 버튼을 클릭합니다.
* 연결이 설정되면 Sysdig 통합 상태가 **Connected**로 변경됩니다.


{% hint style="info" %}
Sysdig 런타임 데이터가 런타임 통합을 통해 사용 가능해지면 몇 시간 내에 Snyk AppRisk에 표시됩니다.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (593).png" alt=""><figcaption><p>Sysdig - 설정 화면</p></figcaption></figure>