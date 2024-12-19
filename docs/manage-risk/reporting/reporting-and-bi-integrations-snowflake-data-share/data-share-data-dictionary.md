---
layout:
  title:
    visible: true
  description:
    visible: false
  tableOfContents:
    visible: true
  outline:
    visible: true
  pagination:
    visible: true
---

# 데이터 공유 데이터 사전

Snyk 데이터 공유는 다양한 데이터 기둥을 포함하여 다양한 용례를 지원하는 포괄적인 데이터 세트입니다. 이 데이터 세트를 사용하여 이슈 백로그, 에이징, MTTR, SLA 준수, 테스트 커버리지 등 주요 보안 지표를 제시하고, 리스크 점수, 심각성, CVSS, EPSS 등 다양한 요소를 기반으로 이슈를 우선순위에 따라 설정할 수 있습니다.

이 사전은 데이터 집합을 효율적으로 탐색하는 데 도움이 되도록 설계되었으며, 각 테이블의 목적과 각 열에 포함된 구체적인 데이터를 명확히 설명하여 귀하의 데이터 보고 요구를 충족시키기 위해 X 데이터를 활용할 수 있도록 지원합니다.

## 데이터 공유 테이블

<figure><img src="../../../.gitbook/assets/Screenshot 2024-08-12 at 1.43.10 PM.png" alt="데이터 사전에 나열된 개체를 정의하는 데이터베이스 다이어그램"><figcaption><p>데이터 사전에 나열된 개체를 정의하는 데이터베이스 다이어그램</p></figcaption></figure>

위 다이어그램은 데이터 사전에 나열된 개체를 데이터베이스 다이어그램으로 표현한 것입니다. 다음 테이블을 다루고 있습니다:

* [그룹](data-share-data-dictionary.md#groups)
* [기관](data-share-data-dictionary.md#orgs)
* [프로젝트](data-share-data-dictionary.md#projects)
* [이슈](data-share-data-dictionary.md#issues)
* [사용 이벤트](data-share-data-dictionary.md#usage-events)
* [지라 이슈](data-share-data-dictionary.md#issue-jira-issues)

### 그룹

> 현재 버전: v1.0

`그룹` 테이블에는 Snyk 그룹의 주요 속성이 포함되어 있습니다. 이 데이터는 그룹 수준에서 집계를 수행하거나 특정 그룹의 범위에 초점을 맞추는 데 사용할 수 있습니다.

| 열 이름         | 데이터 유형    | 설명                                                                                                 |
| -------------- | -------------- | ---------------------------------------------------------------------------------------- |
| `public_id`    | varchar        | 그룹의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                        |
| `display_name` | varchar        | 이 그룹에 설정된 표시 이름.                                                                 |
| `slug`         | varchar        | Snyk 내에서 그룹의 이름.              |
| `created`      | timestamp\_ntz | 이 레코드가 Snyk에 생성된 날짜. |
| `deleted`      | timestamp\_ntz | 이 레코드가 Snyk에서 삭제된 날짜. |
| `modified`     | timestamp\_ntz | 이 레코드가 Snyk 내에서 마지막으로 수정된 날짜. |
| `__updated_at` | timestamp\_ntz | 데이터 공유 데이터 변환이 이 레코드를 마지막으로 업데이트한 시간.    |

### 기관

> 현재 버전: v1.0

`기관` 테이블에는 Snyk 기관의 주요 속성이 포함되어 있습니다. 이 데이터는 조직 수준에서 집계를 수행하거나 특정 조직의 범위에 초점을 맞추는 데 사용할 수 있습니다.

{% hint style="info" %}
`group_public_id` 열을 사용하여 특정 그룹의 기관을 쿼리할 수 있습니다.
{% endhint %}

| 열 이름       | 데이터 유형    | 설명                                                                               |
| ----------------- | -------------- | --------------------------------------------------------------------------------------------- |
| `public_id`       | varchar        | 기관의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다. |
| `group_public_id` | varchar        | 그룹의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.       |
| `display_name`    | varchar        | 이 기관에 설정된 표시 이름.                                                 |
| `slug`            | varchar        | Snyk 내에서 기관의 이름.                                                  |
| `created`         | timestamp\_ntz | 이 레코드가 Snyk에 생성된 날짜.                                           |
| `deleted`         | timestamp\_ntz | 이 레코드가 Snyk에서 삭제된 날짜.                                         |
| `modified`        | timestamp\_ntz | 이 레코드가 Snyk 내에서 마지막으로 수정된 날짜.                             |
| `__updated_at`    | timestamp\_ntz | 데이터 공유 데이터 변환이 이 레코드를 마지막으로 업데이트한 시간.       |

### 프로젝트

> 현재 버전: v1.0

`프로젝트` 테이블에는 주요 Snyk 프로젝트의 속성과 관련 대상이 포함되어 있습니다. 해당 데이터는 프로젝트 또는 대상 수준에서 필터링이나 집계를 수행하는 데 사용할 수 있으며, 이는 프로젝트 컬렉션, 프로젝트 태그 또는 특정 리포지토리 브랜치를 기반으로합니다(`target_ref` 사용).

{% hint style="info" %}
Snyk 보고서는 삭제되지 않은 모니터링 중인 프로젝트만 표시합니다. Snyk 보고서와 결과를 일치시키려면 쿼리를 `IS_MONITORED = TRUE` 및 `DELETE IS NULL`로 필터링하십시오.
{% endhint %}

| 열 이름                        | 데이터 유형    | 설명                                                                                                                                                                                                                       |
| ---------------------------------- | -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `public_id`                        | varchar        | 프로젝트의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                                                                                                                                                   |
| `org_public_id`                    | varchar        | 조직의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                                                                                                                                                                |
| `group_public_id`                  | varchar        | 그룹의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                                                                                                                                                    |
| `name`                             | varchar        | Snyk에 추가 될 때이 프로젝트에 지정된 이름.                                                                                                                                                                                         |
| `is_monitored`                     | boolean        | 현재 이 프로젝트가 활발히 모니터링되도록 설정되어 있는지 여부.                                                                                                                                                                      |
| `project_type`                     | varchar        | 특정 프로젝트에 대해 사용할 스캔 방법, 예: Snyk 코드를 사용한 정적 어플리케이션 보안 테스트 (SAST) 또는 Maven을 사용한 Maven 프로젝트에 대한 {{Snyk 오픈 소스}}. 이는 스캔의 구성 일부입니다. |
| `project_type_display_name`        | varchar        | 내부 프로젝트 유형 값에 Snyk에서 할당한 표시 이름.                                                                                                                                                                                    |
| `test_frequency`                   | varchar        | 특정 프로젝트의 테스트 빈도. 예: 매일, 매주 등.                                                                                                                                                                                 |
| `origin`                           | varchar        | Origin은 CLI, GitHub 또는 Kubernetes와 같은 대상 생태계를 정의합니다. Origin은 대상의 속성입니다.                                                                                                            |
| `target_ref`                       | varchar        | 이 프로젝트를 구분하는 참조, 예: 브랜치 이름이나 버전. 동일한 참조를 가진 프로젝트는 해당 참조를 기준으로 그룹화될 수 있습니다.                                                                                    |
| ....
*중략*
....

### 이슈

> 현재 버전: v1.0

`이슈` 테이블에는 Snyk 이슈의 다양한 속성이 포함되어 있습니다. 이슈는 해당 프로젝트, 대상, 조직 또는 그룹과 쉽게 연관 짓을 수 있으며, 해당 ID 열을 활용할 수 있습니다. 이슈의 기본 속성뿐만 아니라 도입 날짜, 유형, 심각성, 점수 등에 관한 열 외에도 CVSS 점수, EPSS 점수, NVD 점수 등과 같은 취약성 속성에 대해 설명하는 열이 있습니다.

이슈 테이블을 쿼리하여 다음을 수행할 수 있습니다:

* 이슈 백로그, 에이징, MTTR 및 SLA 준수를 포함한 다양한 메트릭 및 KPI 도출
* 시간에 따른 확인된, 무시된, 해결된 이슈의 트렌드 시각화
* 다양한 요소와 고려 사항에 따라 이슈 우선순위 설정

{% hint style="info" %}
Snyk 보고서와 결과가 일치하도록 원하신다면:

* `DELETED_AT IS NULL`로 쿼리를 필터링하여 삭제된 이슈를 표시하지 않습니다.
* 프로젝트 테이블과 이슈 테이블을 조인하고 `IS_MONITORED = TRUE`로 필터링하여 비활성화된 프로젝트의 이슈를 표시하지 않습니다.
{% endhint %}

| 열 이름                      | 데이터 유형    | 설명                                                                                                                                                                                   |
| -------------------------------- | -------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                             | varchar        | 프로젝트의 특정 취약성의 고유 식별자를 나타내는 고유 식별자.                                                                                                                       |
| `problem_id`                     | varchar        | 취약성을 고유 식별하는 Snyk 취약성 데이터베이스 ID.                                                                                                                   |
| `project_public_id`              | varchar        | 프로젝트의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                                                                                                                      |
| `org_public_id`                  | varchar        | 조직의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                                                                                                                |
| `group_public_id`                | varchar        | 그룹의 고유 식별자로, 레코드 소스 데이터베이스에서 할당됩니다.                                                                                                            |
| ...
*중략*
....

### 사용 이벤트

> 현재 버전: v1.0

`사용 이벤트` 테이블에는 Snyk의 CLI 인터페이스(CLI, IDE 플러그인, CI/CD 파이프라인 도구)에서 수집된 CLI 상호 작용 데이터가 포함되어 있습니다. CLI 상호 작용 이벤트는 해당 실행 컨텍스트(대상, 조직 또는 그룹)와 연계할 수 있습니다.

`사용 이벤트` 테이블을 쿼리하여 다음을 측정할 수 있습니다:

* 개발자가 Snyk IDE 플러그인을 사용하는 정도 및 채택
* CI/CD 파이프라인에서의 Snyk 테스트
* 다양한 CLI 명령어별 Snyk CLI 활용| Column name                               | Data type      | Description                                                                                                                                                                                                                                                                                                                          |
| ----------------------------------------- | -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                      | varchar        | 상호 작용 이벤트의 클라이언트 생성 ID로, `urn:snyk:interaction:00000000-0000-0000-0000-000000000000` 형식입니다.                                                                                                                                                                                                          |
| `org_public_id`                           | varchar        | 레코드의 소스 데이터베이스에서 할당된 조직을 식별하는 고유 식별자입니다.                                                                                                                                                                                                                                       |
| `group_public_id`                         | varchar        | 레코드의 소스 데이터베이스에서 할당된 그룹을 식별하는 고유 식별자입니다.                                                                                                                                                                                                                                               |
| `product_display_name`                    | varchar        | 이 상호 작용 중에 사용된 Snyk 제품입니다. 예를 들어, {{Snyk Open Source}}, {{Snyk IaC}}, {{Snyk Code}}, {{Snyk Container}} 등이 있습니다.                                                                                                                                                                                                                   |
| `runtime_application_name`                | varchar        | snyk 상호 작용을 실행하는 응용 프로그램입니다. 예를 들어, PyCharm, Visual Studio, snyk-ls, snyk-cli 등이 있습니다.                                                                                                                                                                                                                          |
| `runtime_application_version`             | varchar        | 통합의 버전입니다.                                                                                                                                                                                                                                                                                                      |
| `runtime_application_data_schema_version` | varchar        | Snyk의 런타임 상호 작용의 데이터 스키마 버전입니다. 현재 버전(v2)은 2024년 제2분기에 출시되었습니다. 이전 버전의 데이터는 다르게 작동할 수 있습니다.                                                                                                                                                                               |
| `interaction_type`                        | varchar        | 상호 작용의 유형으로, "Scan done"이 될 수 있습니다. Scan Done은 CLI 또는 IDE에서 실행되었던 테스트를 나타냅니다. 다른 유형은 자유롭게 선택할 수 있습니다.                                                                                                                                                                   |
| `interaction_categories`                  | array          | 상호 작용을 자세히 설명하는 데 사용되는 범주 벡터입니다. "oss","test" 등이 있습니다.                                                                                                                                                                                                                                                        |
| `interaction_timestamp`                   | array          | 상호 작용이 UTC에서 시작된 시간입니다.                                                                                                                                                                                                                                                                                             |
| `interaction_status`                      | timestamp\_ntz | 상태는 "success" 또는 "failure"가 될 수 있으며, 성공은 동작이 실행되었음을 나타내고, 실패는 실행되지 않았음을 의미합니다.                                                                                                                                                                                                              |
| `interaction_stage`                       | varchar        | 상호 작용이 발생한 SDLC의 단계로, "dev"\|"cicd"\|"prchecks"\|"unknown" 등이 있습니다.                                                                                                                                                                                                                                  |
| `interaction_exit_code`                   | integer        | 실행 중인 프로세스에서 반환된 상호 작용의 종료 코드입니다. 각 상호 작용(테스트, 모니터링 등)의 종료 코드와 그 의미에 대한 자세한 정보는 해당 상호 작용에 대한 Snyk 문서에서 확인할 수 있습니다.                                                                                                                                             |
| `interaction_target_id`                   | varchar        | purl은 일곱 가지 구성 요소로 구성된 URL입니다. scheme:type/namespace/name@version?qualifiers#subpath purl 사양은 다음에서 제공됩니다: `https://github.com/package-url/purl-spec` 일부 purl 예시 `pkg:github/package-url/purl-spec@244fd47e07d1004f0aed9c` `pkg:npm/%40angular/animation@12.3.1` `pkg:pypi/django@1.11.1` |
| `environment_display_name`                | varchar        | 이 상호 작용 중에 사용된 환경입니다. 예를 들어: CLI, Eclipse, Jetbrains IDE, Visual Studio, Visual Studio Code 또는 기타                                                                                                                                                                                                  |
| `runtime_platform_os`                     | varchar        | 통합의 운영 체제입니다 (darwin, windows, linux 등).                                                                                                                                                                                                                                                              |
| `runtime_platform_arch`                   | varchar        | 통합의 아키텍처입니다 (AMD64, ARM64, 386, ALPINE).                                                                                                                                                                                                                                                                    |
| `runtime_environment_name`                | varchar        | 통합 환경입니다 (예: IntelliJ Ultimate, Pycharm).                                                                                                                                                                                                                                                              |
| `runtime_environment_version`             | varchar        | 통합 환경의 버전입니다 (예: 2023.3)                                                                                                                                                                                                                                                                             |
| `runtime_integration_name`                | varchar        | 통합의 이름으로, 플러그인이나 확장일 수 있습니다.                                                                                                                                                                                                                                                                         |
| `runtime_integration_version`             | varchar        | 통합의 버전입니다. 예: 2.3.4.                                                                                                                                                                                                                                                                                  |
| `runtime_performance_duration_ms`         | number         | 상호 작용의 밀리초 단위 실행 시간입니다.                                                                                                                                                                                                                                                                                      |
| `user_email`                              | varchar        | 상호 작용 중 사용자의 이메일 주소입니다.                                                                                                                                                                                                                                                                  |
| `user_name`                               | varchar        | 상호 작용 중 사용자의 이름입니다.                                                                                                                                                                                                                                                                   |
| `deleted`                                 | timestamp\_ntz | 이 레코드가 Snyk에서 삭제된 시간입니다.                                                                                                                                                                                                                                                                                              |
| `__updated_at`                            | timestamp\_ntz | 데이터 공유 데이터 변환에 의해 이 레코드가 마지막으로 업데이트된 시간입니다.                                                                                                                                                                                                                                                                    |

### 이슈 Jira 이슈

> 현재 버전: v1.0

`ISSUE_JIRA_ISSUES` 테이블을 사용하여 Snyk 이슈와 할당된 Jira 이슈 사이를 연관시킬 수 있습니다. Snyk는 여러 유형의 Jira 통합을 지원하기 때문에, 데이터셋에 있는 Jira 이슈는 [이](https://docs.snyk.io/integrate-with-snyk/jira-and-slack-integrations/jira-integration) 문서에서 설명된 Jira 통합에서 유래된 것임을 강조하는 것이 중요합니다.

| Object name            | Data type     | Description                                                                                    |
| ---------------------- | ------------- | ---------------------------------------------------------------------------------------------- |
| id                     | varchar       | 프로젝트의 주어진 취약성의 고유 인스턴스를 나타내는 고유 식별자입니다.                             |
| `problem_id`           | varchar       | 취약성을 고유하게 식별하는 Snyk 취약성 데이터베이스 ID입니다.                     |
| `project_public_id`    | varchar       | 레코드의 소스 데이터베이스에서 할당된 프로젝트를 식별하는 고유 식별자입니다.       |
| `org_public_id`        | varchar       | 레코드의 소스 데이터베이스에서 할당된 조직을 식별하는 고유 식별자입니다. |
| `group_public_id`      | varchar       | 레코드의 소스 데이터베이스에서 할당된 그룹을 식별하는 고유 식별자입니다.         |
| `jira_integration_uri` | varchar       | Snyk Jira 통합에 제공된 Jira 계정의 URL입니다.                             |
| `jira_issues`          | array         | 이 문제에 대해 생성된 모든 Jira 이슈의 배열입니다.                                       |
| `latest_jira_issue`    | varchar       | 이 문제에 대해 가장 최근에 생성된 Jira 이슈입니다.                                           |
| `__updated_at`         | tiestamp\_ntz | 데이터 공유 데이터 변환이 이 레코드를 마지막으로 업데이트한 시간입니다.                              |