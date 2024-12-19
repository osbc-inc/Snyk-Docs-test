# 보고서 및 BI 통합: Snowflake Data Share

새로운 Snowflake Data Share 통합을 통해 데이터 과학, BI 및 AppSec 팀은 Snyk 보고서에서 사용 가능한 동일한 기본 데이터에 안전하게 액세스할 수 있지만, 자체 Snowflake 계정 내에서 사용하여 Snyk 데이터를 이해하고 시각화하는 강력한 새로운 분석 도구를 활용할 수 있습니다.

이 통합을 사용하여 팀이 Snowflake가 지원하는 도구 생태계를 활용하여 신속하게 탐색 및 사용자 정의 분석을 신속하게 구축할 수 있습니다. 고객은 Snyk 데이터를 PowerBI, Tableau 및 Looker Data Studio 같은 BI 도구에 연결하거나 사용자 정의 Streamlit 앱을 작성할 수 있습니다.

Snyk 데이터 집합을 직접 Snowflake 계정에 넣게되면 Snyk 데이터를 추가 데이터 소스와 결합할 수 있는 문을 열어주며 조직 수준의 통합된 AppSec 보기에 상당한 기여를 합니다.

## Snowflake Data Share란?<a href="#what-is-snowflake-data-share" id="what-is-snowflake-data-share"></a>

[Snowflake Data Share](https://docs.snowflake.com/en/user-guide/data-sharing-intro.html)는 기업이 데이터를 고객에게 안전하고 간단하게 제공할 수 있도록 하는 서비스입니다. Snowflake 데이터 공유를 통해 데이터는 계정 간에 교환되지 않고 데이터 컨슈머가 고객 자신의 Snowflake 계정을 통해 공유된 데이터베이스에 대한 읽기 전용 액세스 권한을 부여받습니다.

### Snowflake 비용 영향<a href="#main-use-cases" id="main-use-cases"></a>

* **저장소**\
  데이터 공유 중에 데이터가 전송되지 않기 때문에 Snowflake 계정에 추가 저장소 비용이 발생하지 않습니다.
* **컴퓨팅 리소스**\
  데이터 공유를 통해 쿼리하는 데 사용된 컴퓨팅 리소스에 대해 Snowflake가 요금을 청구할 것입니다.

## 주요 사용 사례<a href="#main-use-cases" id="main-use-cases"></a>

Snowflake Data Share는 다양한 사용 사례에 활용되며 수많은 보안 및 비즈니스 관련 질문에 대답할 수 있습니다. 다음은 일부 사용 사례입니다:

* **CISO 및 경영진을위한 AppSec 포지션 가시성 강화**\
  Snyk 데이터를 BI 플랫폼 및 기존 보안 대시보드로 최적화하고 성능 메트릭 및 KPI(예: MTTR, SLA 준수, 교정 트렌드 등)를 반영하세요.&#x20;
* **구체적인 질문에 대답하거나 독특한 통찰력 제공**\
  특정 위험 노출 트렌드를 더 잘 이해하라. 예를 들어, 특정 위험 점수 이상의 총 문제를 추적하다. 특정 프로젝트 컬렉션이나 모든 Snyk 그룹에서 특정한 태그에 영향을 미치거나 main 또는 master 브랜치에 대해서만 필터링하세요.\
  SLA에 대한 수정 동작의 성능을 측정하세요. 예를 들어, 고객 SLA 목표를 입력하고 해당 목표를 추적하세요.\
  좌측 이동 영향을 이해하기 위해 사용자 정의 예방 보고서를 작성하세요. 예를 들어, 모든 Snyk 그룹에서 특정 심각성 및 위험 점수로 필터링된 예방 가능한 오픈 소스 취약점의 트렌드를 확인합니다.

## 시작하기<a href="#getting-started" id="getting-started"></a>

### Snowflake Data Share 액세스 요청<a href="#request-a-snowflake-data-share-access" id="request-a-snowflake-data-share-access"></a>

다음 단계를 따라 Snowflake Data Share 액세스를 요청하세요:

1. Snyk 계정 담당자에게 연락하여 액세스를 요청하세요.
2. Snyk 연락 담당자에게 다음 Snowflake 계정 세부 정보를 제공하세요(자격 증명을 추적하기 위한 [여기](https://docs.snowflake.com/en/user-guide/admin-account-identifier#finding-the-organization-and-account-name-for-an-account) 가이드라인 참조):
   * 계정 이름.
   * 조직 이름.
   * 데이터 공유를 특정 Snyk 그룹에 제한하려는 경우 관련 그룹 ID를 언급하십시오(ID는 Snyk 그룹 설정에 있습니다).
3. Snyk가 Snowflake 계정 세부 정보를 받은 후, 팀은 데이터 공유를 준비합니다. 데이터가 24시간 이내에 표시되도록 응답해야 합니다.

### Snowflake Data Share 사용 준비<a href="#prepare-to-consume-snowflake-data-shares" id="prepare-to-consume-snowflake-data-shares"></a>

데이터 공유는 **비공개 공유 목록(Privately Shared Listing)** 형식으로 제공됩니다. Snowflake 데이터 공유를 처음 사용하는 경우, 다음 단계를 수행하십시오. 그렇지 않으면 [Snyk Data Share에서 데이터베이스 생성](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share)를 이어서 진행하십시오:

1. Snowflake 조직 관리자(ORGADMIN 역할 필요)는 [Snowflake Provider 및 Consumer 약관을 수락](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#accept-the-snowflake-provider-and-consumer-terms-of-service)해야 합니다.
2. [필요한 권한 설정](https://other-docs.snowflake.com/en/collaboration/consumer-becoming#set-up-required-privileges) (CREATE DATABASE 및 IMPORT SHARE 권한이 있는 ACCOUNTADMIN 역할 또는 다른 역할 필요).

### 데이터 공유로부터 데이터베이스 생성<a href="#create-a-database-from-the-data-share" id="create-a-database-from-the-data-share"></a>

데이터 공유에서 데이터에 쿼리할 수 있도록 액세스하고 싶다면, 다음 단계를 따라 데이터 공유로부터 데이터베이스를 생성하세요:

1. ACCOUNTADMIN 역할(또는 CREATE DATABASE 및 IMPORT SHARE 권한을 가진 역할)의 Snowflake 사용자는 [사용 가능한 공유 보기](https://docs.snowflake.com/en/user-guide/data-share-consumers#viewing-available-shares)하십시오.
2. [Snyk 데이터 공유로부터 데이터베이스 생성](https://docs.snowflake.com/en/user-guide/data-share-consumers#creating-a-database-from-a-share)하십시오.\
   **참고:** 클라우드 지역에 따라 데이터가 프로비저닝되어 사용 준비가되는 데 약 10분이 소요될 수 있습니다.
3. [공유된 데이터베이스에 권한 부여](https://docs.snowflake.com/en/user-guide/data-share-consumers#granting-privileges-on-a-shared-database).

### Snyk 데이터를 사용하여 대시보드 생성

Snyk 데이터로 대시보드를 작성하려면, Snyk가 [사용 사례 및 예제 쿼리](build-your-first-dashboard.md)를 제공하여 AppSec 목표에 관련된 주요 성능 메트릭을 시각화할 수 있습니다.

쿼리 작성에 익숙해지면, Snyk는 특정 요구 사항에 맞는 사용자 정의 쿼리를 작성할 것을 권장합니다.

{% hint style="info" %}
Snowflake 데이터 공유에서 쿼리를 사용하는 방법에 대한 자세한 내용은 [Snowflake에서 데이터 쿼리하기](https://docs.snowflake.com/en/guides-overview-queries)를 참조하세요.
{% endhint %}

## 데이터 정책<a href="#data-policy" id="data-policy"></a>

### 데이터 범위 및 접근성<a href="#data-freshness" id="data-freshness"></a>

- Snowflake 데이터 공유는 요청된 Snyk 그룹의 데이터 범위로 제한됩니다. 고객은 모든 Snyk 그룹 또는 특정 그룹에 대한 액세스를 요청할 수 있습니다.
- Snyk는 공유된 Snyk 데이터베이스에 있는 요청된 Snyk 그룹의 모든 사용 가능한 데이터를 공유하는데 있어 데이터 공유 자체에 추가적인 제한은 없습니다.
- 데이터 공유 자체는 읽기 전용 데이터베이스로 제공되며 Snowflake 표준 역할 기반 액세스 제어에 따라 접근할 수 있습니다.

### 어떤 데이터를 제공하나요? <a href="#supported-data" id="supported-data"></a>

Snyk의 데이터 공유 제공은 그룹, 조직, 프로젝트, 문제, 사용 이벤트 및 Jira 문제를 포함한 다양한 개체를 다룰 수 있도록 하며, 이를 통해 Snyk 데이터 위에 강력한 보고 능력을 제공합니다.

{% hint style="info" %}
Snyk의 제공에 관한 구체적인 내용은 [데이터 공유 사전](data-share-data-dictionary.md)을 참조하십시오.
{% endhint %}

### 데이터 신선도<a href="#data-freshness" id="data-freshness"></a>

데이터 공유에서 사용 가능한 데이터는 약 2시간마다 새로고침됩니다. 이 과정은 자동으로 발생하며 고객의 조치가 필요하지 않습니다.

### 변경 추적

Snowflake Change Tracking은 Snowflake Data Share에 포함된 모든 테이블을 활성화합니다. 이 기능은 이러한 테이블 내의 행의 변경 사항을 추적하여 워크플로우의 효율적인 증분 처리와 동기화를 가능하게 합니다.

* **유지 기간:** 변경 사항은 2일 동안 보존됩니다.
* **사용 방법:** Snowflake가 제공하는 CHANGES 절을 사용하여 변경 추적 메타데이터를 쿼리할 수 있습니다.

이 기능을 활용하는 방법에 대한 자세한 내용은 [Snowflake의 변경 추적 설명서](https://docs.snowflake.com/en/sql-reference/constructs/changes)를 참조하십시오.