# SCM 통합을 위한 애플리케이션 컨텍스트

이는 응용프로그램 컨텍스트를 설정할 수 있는 사용 가능한 통합 사항입니다:

* [백스테이지 파일](./#backstage-file-for-scm-integrations)
* [ServiceNow CMDB](./#servicenow-cmdb-for-scm-integrations)
* [Atlassian Compass](./#atlassian-compass)
* [Harness](./#harness)
* [OpsLevel](./#opslevel)
* [Datadog Service Catalog](./#datadog-service-catalog)

{% hint style="info" %}
이 페이지의 응용프로그램 컨텍스트 통합은 AppRisk SCM 통합을 통해 발견된 자산과 협력하여 작동합니다. 통합 페이지의 그룹 레벨에서 Snyk AppRisk SCM 통합이 구성되어 있지 않으면 이러한 통합에서 데이터가 채워지지 않습니다.
{% endhint %}

## SCM 통합을 위한 백스테이지 파일

{% hint style="info" %}
**릴리스 상태**

백스테이지 파일 통합은 초기 액세스 상태이며, Snyk AppRisk Essentials 및 Snyk AppRisk Pro 플랜에서 이용할 수 있습니다.
{% endhint %}

백스테이지는 사용자들이 리포지토리에 메타데이터 또는 주석을 추가하여 사용 가능한 리소스를 더 쉽게 탐색하고 이해하기 위해 정리하고 분류하는 서비스 카탈로그입니다. SCM 통합을 활용하여 백스테이지 카탈로그 파일과 관련된 메타데이터를 Snyk AppRisk로 가져올 수 있습니다.

백스테이지 카탈로그 파일은 GitHub, GitLab, Azure DevOps, BitBucket Cloud 및 BitBucket 온프레미스 SCM 통합에 사용할 수 있습니다.

### 백스테이지 파일의 필수 매개변수

* 구성된 SCM 통합.
* 프로젝트의 `catalog-info.yaml` 파일.

### 백스테이지 파일을 위한 통합 허브 설정

1. **인테그레이션 허브** 메뉴를 엽니다.
2. SCM 통합을 선택합니다.
3. SCM 통합의 **설정** 옵션을 클릭합니다.
4. **백스테이지 카탈로그 추가** 옵션을 활성화합니다.
5. 선택 사항 - 리포지토리의 백스테이지 카탈로그 파일 이름이 `catalog-info.yaml`가 아닌 경우, 백스테이지 카탈로그 파일 이름 필드에서 기본값을 변경할 수 있습니다.
6. Snyk AppRisk에 추가할 적어도 하나 이상의 속성을 선택합니다.

{% hint style="info" %}
Snyk AppRisk는 감지된 파일의 필드를 기본 필드 이름을 사용하여 구문 분석하지만 대체 필드 이름이 지정되지 않으면 설정할 필요가 있습니다.
{% endhint %}

7. **완료** 버튼을 클릭합니다.

백스테이지 카탈로그를 구성한 후에는 Snyk AppRisk가 백스테이지 카탈로그 .yaml 파일에서 발견된 데이터로 리포지토리 자산을 풍부하게 만들기 시작합니다.

{% hint style="warning" %}
카탈로그 속성을 설정할 때 특정 서비스 수준 속성을 사용해야 합니다. 예를 들어 `attribute.name.`
{% endhint %}

## SCM 통합을 위한 ServiceNow CMDB

{% hint style="info" %}
**릴리스 상태**

ServiceNow CMDB 통합은 초기 액세스 상태이며, Snyk AppRisk Essentials 및 Snyk AppRisk Pro 플랜에서 이용할 수 있습니다.
{% endhint %}

### ServiceNow CMDB를 위한 필수 매개변수 <a href="#servicenow-cmdb-required-parameters" id="servicenow-cmdb-required-parameters"></a>

1. ServiceNow CMDB 인스턴스를 위한 **프로필 이름**을 추가합니다.
2. ServiceNow CMDB의 **CMDB 인스턴스**를 설정합니다. 예를 들어 이 예제를 따라서 사용하여 ServiceNow CMDB용 CMDB 인스턴스를 설정합니다: `https://<INSTANCE_NAME>.service-now.com`.
3. **사용자 이름** 및 **비밀번호** - ServiceNow CMDB 인스턴스에 대한 자격 증명입니다.
4. CMDB 구성 항목 클래스를 위한 **테이블 이름**을 추가합니다. 전체 이름 목록에 대한 자세한 내용은 [ServiceNow CMDB 테이블 세부사항](https://docs.servicenow.com/bundle/washingtondc-servicenow-platform/page/product/configuration-management/reference/cmdb-tables-details.html) 페이지를 참조하십시오.
5. **Repo URL을 매핑할 CMDB 필드 추가** - 리포지토리의 URL을 추가합니다.

{% hint style="info" %}
* Snyk이 ServiceNow CMDB로부터 수집한 데이터는 리포지토리 자산과 연관시킬 것입니다.
* ServiceNow CMDB 통합은 기본 인증을 사용하며 서비스 계정의 "웹 서비스 액세스만" 옵션을 활성화하는 것을 제안합니다.
{% endhint %}

### ServiceNow CMDB를 위한 통합 허브 설정 <a href="#servicenow-cmdb-integration-hub-setup" id="servicenow-cmdb-integration-hub-setup"></a>

* **인테그레이션 허브** 메뉴를 엽니다.
* **App Context** 태그를 선택하고 ServiceNow CMDB를 검색합니다.
* **추가** 버튼을 클릭합니다.
* **프로필 이름**을 추가합니다 - 이는 ServiceNow CMDB 프로필의 이름입니다.
* **CMDB 인스턴스**를 추가합니다 - 이는 ServiceNow 인스턴스입니다. 이 형식을 사용하십시오: `https://<INSTANCE_NAME>.service-now.com`
* **사용자 이름** 및 **비밀번호** - ServiceNow CMDB 인스턴스에 액세스하는 사용자 이름 및 비밀번호를 추가합니다.
* **테이블 이름**을 추가합니다 - Snyk AppRisk가 적용해야 할 구성 항목 클래스를 선택합니다. 이 형식을 사용하십시오: `cmdb_ci_<class>`
* **Repo URL을 매핑할 CMDB 필드 추가** - ServiceNow CMDB 레코드에서 참조되는 특정 URL을 추가합니다.
* 리포지토리 자산과 관련된 하나 이상의 속성을 선택하고 Snyk AppRisk가 이 속성을 ServiceNow CMDB에서 어디에서 가져올 수 있는지 구성합니다. 예시:
  * 카테고리: application\_type
  * 소유자: business\_unit
* **완료** 버튼을 클릭합니다.
* 연결이 설정되면 ServiceNow CMDB 통합의 상태가 **연결됨**으로 변경되며, Snyk AppRisk는 ServiceNow CMDB에서 발견된 데이터로 리포지토리 자산을 풍부하게 만들기 시작합니다.

{% hint style="warning" %}
카탈로그 속성을 설정할 때 속성의 이름을 사용자 정의할 수 있지만 동일한 이름이 카탈로그 및 인테그레이션 허브 설정에서 사용됨이 보장되어야 합니다.
{% endhint %}

다음 비디오는 **Integration Hub**에서 ServiceNow CMDB 옵션에 대한 개요와 사용 가능한 속성에 대한 간단한 설명을 제공합니다:

{% embed url="https://youtu.be/I4EyhOeQGT0" %}

비디오가 마음에 들었나요? [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)에서 나머지 코스를 확인하세요!

## Atlassian Compass

{% hint style="info" %}
**릴리스 상태**

Atlassian Compass 통합은 초기 액세스 상태이며, Snyk AppRisk Essentials 및 Snyk AppRisk Pro 플랜에서 이용할 수 있습니다.
{% endhint %}

### Atlassian Compass를 위한 필수 매개변수

1. Atlassian Compass **프로필 이름**을 추가합니다.
2. Atlassian Compass **인스턴스 URL**을 추가합니다. 다음 형식을 사용할 수 있습니다: `https://<YOUR ORGANIZATION>.atlassian.net`.
3. Atlassian Compass **사용자 이름**을 추가합니다.
4. Atlassian Compass 인스턴스 **토큰**을 추가합니다. Atlassian API 토큰을 만드는 자세한 내용은 [Atlassian 계정의 API 토큰 관리](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/) 페이지를 참조하십시오.

{% hint style="info" %}
Atlassian Compass로부터 수집한 데이터는 리포지토리 자산과 연관시킬 것입니다.

이 기능은 Atlassian Compass와의 통합에 대해서만 사용할 수 있습니다.
{% endhint %}

### Atlassian Compass를 위한 통합 허브 설정

1. **인테그레이션 허브** 메뉴를 엽니다.
2. **App Context** 태그를 선택하고 Atlassian Compass를 검색합니다.
3. **추가** 버튼을 클릭합니다.
4. **프로필 이름**을 추가합니다 - 이는 Atlassian Compass 프로필의 이름입니다.
5. **인스턴스 URL**을 추가합니다 - 이는 Atlassian Compass 인스턴스의 URL입니다. 이 형식을 사용할 수 있습니다: `https://<YOUR ORGANIZATION>.atlassian.net`
6. **사용자 이름**을 추가합니다 - 이는 Atlassian Compass 인스턴스에 액세스하는 사용자 이름입니다.
7. **토큰**을 추가합니다 - 이는 Atlassian Compass 인스턴스에 액세스하는 API 토큰입니다.
8. 리포지토리 자산과 관련된 하나 이상의 속성을 선택하고 Snyk AppRisk가 Atlassian Compass에서 가져올 수 있는 속성에 따라 구성합니다:
   * **카탈로그 이름** - 이름과 일치합니다.
   * **카테고리** - '`fields.definition.name`'이 tier일 때 식별됩니다.
   * **라이프사이클** - '`fields.definition.name`'이 라이프사이클일 때 식별됩니다.
   * **소유자** - `ownerId` (ownerId에서 소유자 이름 찾기).
   * **애플리케이션** - `typeId` (모든 구성 요소 유형, 애플리케이션, 서비스, 라이브러리 등에 ID가 할당됩니다).
9. **완료** 버튼을 클릭합니다.
10. 연결이 설정되면 Atlassian Compass 통합의 상태가 **연결됨**으로 변경되며, Snyk AppRisk는 Atlassian Compass에서 발견된 데이터로 리포지토리 자산을 풍부하게 만들기 시작합니다.

{% hint style="warning" %}
카탈로그 속성을 설정할 때 특정 서비스 수준 속성을 사용해야 합니다. 예를 들어 `attribute.name.`
{% endhint %}

## Harness

{% hint style="info" %}
**릴리스 상태**

Harness 통합은 초기 액세스 상태이며, Snyk AppRisk Essentials 및 Snyk AppRisk Pro 플랜에서 이용할 수 있습니다.
{% endhint %}

### Harness를 위한 필수 매개변수

1. Harness **프로필 이름**을 추가합니다.
2. Harness 계정의 **호스트 URL**을 추가합니다. 다음 형식을 사용할 수 있습니다: `https://<YOUR ORGANIZATION>.harness.io`
3. Harness 인스턴스의 **API 키**를 추가합니다. API 키를 관리하려면 Harness [API 키 추가 및 관리](https://developer.harness.io/docs/platform/automation/api/add-and-manage-api-keys/) 문서를 참조하십시오.

{% hint style="info" %}
이 통합은 [Harness의](https://developer.harness.io/docs/internal-developer-portal/catalog/software-catalog/) 서비스 카탈로그 모듈에 집중됩니다. 이는 백스테이지 카탈로그를 백업합니다.
{% endhint %}

### Harness를 위한 통합 허브 설정

1. **인테그레이션 허브** 메뉴를 엽니다.
2. **App Context** 태그를 선택하고 Harness를 검색합니다.
3. **추가** 버튼을 클릭합니다.
4. **프로필 이름**을 추가합니다 - 이는 Harness 인스턴스의 이름입니다.
5. Harness 계정의 **호스트 URL**을 추가합니다.
6. Harness 인스턴스의 **API 키**를 추가합니다.
7. 최소한 하나의 Harness 소프트웨어 카탈로그 [메타데이터](https://developer.harness.io/docs/internal-developer-portal/catalog/software-catalog#component-definition-yaml)를 선택합니다:
   * 카탈로그 이름 - 이 메타데이터를 선택하면 **카탈로그 이름 키**를 필수적으로 추가해야 합니다.
   * 제목 - 이 메타데이터를 선택하면 **제목 키**를 필수적으로 추가해야 합니다.
   * 카테고리 - 이 메타데이터를 선택하면 **카테고리 키**를 필수적으로 추가해야 합니다.
   * 라이프사이클 - 이 메타데이터를 선택하면 **라이프사이클 키**를 필수적으로 추가해야 합니다.
   * 소유자 - 이 메타데이터를 선택하면 **소유자 키**를 필수적으로 추가해야 합니다.
   * 애플리케이션 - 이 메타데이터를 선택하면 **애플리케이션 키**를 필수적으로 추가해야 합니다.
8. **완료** 버튼을 클릭합니다.
9. 연결이 설정되면 Harness 통합의 상태가 **연결됨**으로 변경되며, Snyk AppRisk는 Harness에서 발견된 데이터로 리포지토리 자산을 풍부하게 만들기 시작합니다.

{% hint style="warning" %}
카탈로그 속성을 설정할 때 속성의 이름을 사용자 정의할 수 있지만 동일한 이름이 카탈로그 및 인테그레이션 허브 설정에서 사용됨이 보장되어야 합니다.
{% endhint %}

## OpsLevel

{% hint style="info" %}
**릴리스 상태**

OpsLevel 통합은 초기 액세스 단계에 있으며, Snyk AppRisk Essentials와 Snyk AppRisk Pro 플랜에서 사용할 수 있습니다.
{% endhint %}

### OpsLevel에 필요한 매개변수

1. OpsLevel **프로필 이름**을 추가하세요.
2. OpsLevel 계정의 **인스턴스 URL**을 추가하세요. 이 형식 유형을 사용할 수 있습니다: `https://<YOUR Organizer>.opslevel.com`
3. OpsLevel 인스턴스의 **API 토큰**을 추가하세요. OpsLevel 계정에서 API 토큰을 생성하려면, OpsLevel [API 토큰 생성](https://docs.opslevel.com/docs/graphql#1-create-an-api-token) 문서 페이지를 참조하세요.

### OpsLevel을 위한 통합 허브 설정

1. **통합 허브** 메뉴를 엽니다.
2. **앱 컨텍스트** 태그를 선택하고 OpsLevel을 검색합니다.
3. **추가** 버튼을 클릭합니다.
4. **프로필 이름**을 추가합니다. 이는 OpsLevel 인스턴스의 이름입니다.
5. OpsLevel 계정의 **인스턴스 URL**을 추가합니다.
6. OpsLevel 인스턴스의 **API 토큰**을 추가합니다.
7. Snyk AppRisk가 OpsLevel에서 가져올 수 있는 리포지토리 자산과 관련된 하나 이상의 속성을 선택할 수 있습니다. 다음 매핑을 확인하세요:
   * 카탈로그 이름 - OpsLevel에서 `name`으로 식별됩니다.
   * 카테고리 - OpsLevel에서 `tier.name`으로 식별됩니다.
   * 라이프사이클 - OpsLevel에서 `lifecycle.name`으로 식별됩니다.
   * 소유자 - OpsLevel에서 `owner.name`으로 식별됩니다.
   * 애플리케이션 - OpsLevel에서 `product`로 식별됩니다.
8. **완료** 버튼을 클릭합니다.
9. 연결이 설정되면 OpsLevel 통합 상태가 **연결됨**으로 변경되며, Snyk AppRisk는 OpsLevel에서 찾은 데이터를 사용하여 리포지토리 자산을 강화하기 시작합니다.

{% hint style="warning" %}
카탈로그 속성을 설정할 때는 반드시 특정 서비스 수준 속성을 사용해야 합니다. 예를 들어 `attribute.name`과 같이 사용해야 합니다.
{% endhint %}

## Datadog 서비스 카탈로그

{% hint style="info" %}
**릴리스 상태**

Datadog 서비스 카탈로그 통합은 초기 액세스 단계에 있으며, Snyk AppRisk Essentials와 Snyk AppRisk Pro 플랜에서 사용할 수 있습니다.
{% endhint %}

### Datadog 서비스 카탈로그에 필요한 매개변수

1. Datadog **프로필 이름**을 추가하세요.
2. Datadog 인스턴스의 **API 키**를 추가하세요. 토큰은 다음 범위 권한을 가져야 합니다: `apm_service_catalog_read`.
3. Datadog의 **애플리케이션 키**를 추가하세요. 애플리케이션 키는 사용자가 Datadog의 프로그래밍 API에 액세스할 수 있도록 합니다. 자세한 내용은 [Datadog API 및 애플리케이션 키](https://docs.datadoghq.com/account_management/api-app-keys/) 문서 페이지를 참조하세요.

### Datadog 서비스 카탈로그를 위한 통합 허브 설정

1. **통합 허브** 메뉴를 엽니다.
2. **앱 컨텍스트** 태그를 선택하고 **Datadog 서비스 카탈로그**를 검색합니다.
3. **추가** 버튼을 클릭합니다.
4. **프로필 이름**을 추가합니다. 이는 Datadog 인스턴스의 이름입니다.
5. Datadog 인스턴스의 **API 키**를 추가합니다.
6. Datadog 인스턴스의 **애플리케이션 키**를 추가합니다.
7. **Datadog 사이트**에 대한 세부 정보를 추가합니다.
8. Snyk AppRisk가 Datadog 서비스 카탈로그에서 가져올 수 있는 리포지토리 자산과 관련된 하나 이상의 속성을 선택할 수 있습니다. 다음 매핑을 확인하세요:
   * 카탈로그 이름 - 이 메타데이터를 선택하면 **카탈로그 이름 키**를 추가해야 합니다.
   * 제목 - 이 메타데이터를 선택하면 **제목 키**를 추가해야 합니다.
   * 카테고리 - 이 메타데이터를 선택하면 **카테고리 키**를 추가해야 합니다.
   * 라이프사이클 - 이 메타데이터를 선택하면 **라이프사이클 키**를 추가해야 합니다.
   * 소유자 - 이 메타데이터를 선택하면 **소유자 키**를 추가해야 합니다.
   * 애플리케이션 - 이 메타데이터를 선택하면 **애플리케이션 키**를 추가해야 합니다.
9. **완료** 버튼을 클릭합니다.
10. 연결이 설정되면 Datadog 서비스 카탈로그 통합 상태가 **연결됨**으로 변경되며, Snyk AppRisk는 Snyk AppRisk SCM 통합에서 수집된 리포지토리 자산을 Datadog 서비스 카탈로그에서 찾은 데이터로 강화하기 시작합니다.

{% hint style="warning" %}
카탈로그 속성을 설정할 때 속성의 이름을 사용자 정의할 수 있지만, 카탈로그와 통합 허브 설정에서 동일한 이름을 사용해야 합니다.
{% endhint %}
