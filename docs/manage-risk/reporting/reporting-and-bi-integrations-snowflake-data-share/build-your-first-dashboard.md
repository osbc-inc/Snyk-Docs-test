# 대시보드 만들기

이 안내서는 주요 성과 지표(KPI) 및 관련 사용 사례를 기반으로 첫 번째 AppSec 대시보드를 구성하기 위한 예제 쿼리를 제공합니다. 이러한 사용 사례는 비즈니스 가치 및 구현 고려 사항 측면에서 설명되며 사용 사례에 따라 정리되어 있습니다. 제공된 쿼리는 시작점을 제공하지만 Snyk은 사용자가 특정 요구 사항에 맞게 사용자 정의할 것을 권장합니다.

다음 사용 사례에 대한 쿼리를 참조하십시오:

- [미해결 문제 누적량](build-your-first-dashboard.md#open-issues-backlog) - 주의가 필요한 현재 AppSec 위험 공개 &#x20;
- [노후화(Aging)](build-your-first-dashboard.md#aging) - 미해결 문제의 노출 창 추적
- [해결까지의 평균 시간(MTTR)](build-your-first-dashboard.md#mttr) - 엔지니어링 팀의 해결 속도 분석
- [서비스 수준 협약(SLA)](build-your-first-dashboard.md#sla) - 문제 해결이 귀하의 규정 요구 사항과 일치하는지 확인
- [IDE 및 CLI 테스트 비율](build-your-first-dashboard.md#developers-ide-and-cli-test-usage-and-adoption) - 개발 단계에서의 AppSec 테스팅의 개발자 채택 측정
- [CI/CD 파이프라인 테스트 비율](build-your-first-dashboard.md#ci-cd-pipelines-test-usage-and-adoption) - CI/CD 파이프라인에서의 AppSec 테스팅 채택 측정

{% hint style="warning" %}
실행 전 제공된 예제 쿼리에서 데이터베이스 및 스키마 명을 업데이트해야 합니다.
{% endhint %}

## 미해결 문제 누적량

### 비즈니스 가치

AppSec 팀은 현재 위험에 노출됨을 이해해야 합니다. 이를 위해 기존 문제 누적량의 다양한 측면을 살펴봅니다:

- 미해결된 문제 수
- 높은 또는 위험도가 높은 미해결 문제 수
- 미해결된 문제에 대한 수정사항이 있는지 여부

보다 큰 맥락을 제공하기 위해, 이러한 수치는 공학 팀, 응용 프로그램 또는 결과를 보다 간결하고 실행 가능하게 하는 유의미한 비즈니스 구조로 분해됩니다. 다음 예제 쿼리가 이 조사를 가능하게 합니다.

### SCA를 위한 예제 쿼리

이 쿼리는 Snyk 조직별로 그룹화되고 수정 가능성에 따라 분할된 열린 SCA 문제 누적 카운터를 반환합니다.

결과는 다음을 기반으로 합니다:
- Snyk Open Source (SCA)에서 발견된 열린 위험 또는 위험도 높은 문제
- 소음 제거:
  - 모니터링된 프로젝트의 문제만 포함
  - 삭제된 문제가 없음

```sql
SELECT  o.DISPLAY_NAME AS organization_display_name, -- 원하는 집계에 따라 업데이트
        COUNT_IF(ISSUE_SEVERITY='Critical' AND COMPUTED_FIXABILITY='Fixable') AS fixable_critical_issues,
        COUNT_IF(ISSUE_SEVERITY='High'AND COMPUTED_FIXABILITY='Fixable') AS fixable_high_issues,
        COUNT_IF(ISSUE_SEVERITY='Critical' AND COMPUTED_FIXABILITY='Partially Fixable') AS partially_fixable_critical_issues,
        COUNT_IF(ISSUE_SEVERITY='High'AND COMPUTED_FIXABILITY='Partially Fixable') AS partially_fixable_high_issues,
        COUNT_IF(ISSUE_SEVERITY='Critical' AND COMPUTED_FIXABILITY='No Fix Supported') AS unfixable_critical_issues,
        COUNT_IF(ISSUE_SEVERITY='High'AND COMPUTED_FIXABILITY='No Fix Supported') AS unfixable_high_issues
FROM SNYK.SNYK.ISSUES__V_1_0 i
     INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
     INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID
WHERE p.IS_MONITORED = TRUE                      -- 모니터링된 프로젝트만 포함
     AND i.DELETED_AT IS NULL                    -- 삭제된 문제 제거
     AND ISSUE_STATUS = 'Open'                   -- 열린 문제만 포함
     AND i.PRODUCT_NAME = ''     -- Snyk Open Source만 포함
GROUP BY o.DISPLAY_NAME                          -- 원하는 집계에 따라 업데이트
ORDER BY fixable_critical_issues DESC, fixable_high_issues DESC, 
    partially_fixable_critical_issues DESC, partially_fixable_high_issues DESC; -- 원하는 순서에 따라 업데이트
```

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (2).png" alt="SCA 문제 누적 카운터에 대한 SQL 쿼리 출력"><figcaption><p>SCA 문제 누적 카운터에 대한 SQL 쿼리 출력</p></figcaption></figure>

### Code를 위한 예제 쿼리

이 쿼리는 Snyk 조직별로 그룹화되고 심각도별로 분할된 열린 Snyk Code 문제 누적 카운터를 반환합니다.

결과는 다음을 기반으로 합니다:
- Snyk Code에서 발견된 열린 문제
- 소음 제거:
  - 모니터링된 프로젝트의 문제만 포함
  - 삭제된 문제가 없음

```sql
SELECT o.DISPLAY_NAME AS organization_display_name, -- 원하는 집계에 따라 업데이트
        COUNT_IF(ISSUE_SEVERITY='High') AS high_issues,
        COUNT_IF(ISSUE_SEVERITY='Medium') AS medium_issues,
        COUNT_IF(ISSUE_SEVERITY='Low') AS low_issues
FROM SNYK.SNYK.ISSUES__V_1_0 i
     INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
     INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID
WHERE p.IS_MONITORED = TRUE                     -- 모니터링된 프로젝트만 포함
     AND i.DELETED_AT IS NULL                   -- 삭제된 문제 제거
     AND ISSUE_STATUS = 'Open'                  -- 열린 문제만 포함
     AND i.PRODUCT_NAME = ''           -- Snyk Open Source만 포함
GROUP BY o.DISPLAY_NAME                         -- 원하는 집계에 따라 업데이트
ORDER BY high_issues DESC, 
         medium_issues DESC, 
         low_issues DESC;                     -- 원하는 순서에 따라 업데이트
```

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (1) (18).png" alt="모든 Snyk Code 문제 누적 카운터를 위한 SQL 쿼리 출력 형식"><figcaption><p>모든 Snyk Code 문제 누적 카운터를 위한 SQL 쿼리 출력 형식</p></figcaption></figure>

## 노후화(Aging)

### 비즈니스 가치

문제 노후화는 문제가 발생한 시점부터 현재 날짜까지 경과된 시간을 의미합니다. 조직은 이 지표에 대해 관심이 있으며 노출 창이 연장될수록 이용 가능성이 증가한다고 합니다.

이 리스크를 완화하기 위해 AppSec 팀은 문제가 예상된 해결 시간을 초과했을 때를 지정하는 사전 정의된 SLA 기준을 모니터링합니다.

{% hint style="info" %}
문제가 재등장했을 때, 노후화는 마지막 등장일 기반으로 계산됩니다.
{% endhint %}

### 예제 쿼리

아래 쿼리는 Snyk 조직별로 크리티컬 문제의 평균 노후화(일 단위)를 반환합니다.\
결과는 다음을 기반으로 합니다:
- 열린 크리티컬 문제
- 소음 제거:
  - 모니터링된 프로젝트의 문제만 포함
  - 삭제된 문제가 없음

```sql
SELECT o.DISPLAY_NAME AS organization_display_name, 
    ROUND(AVG(
    CASE
        WHEN LAST_INTRODUCED IS NULL THEN DATEDIFF('DAY', TO_DATE(FIRST_INTRODUCED), CURRENT_DATE)
        WHEN TO_DATE(FIRST_INTRODUCED) <= TO_DATE(LAST_INTRODUCED) THEN DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED), CURRENT_DATE)
    END),0) AS open_issues_aging
FROM SNYK.SNYK.ISSUES__V_1_0 i
     INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
     INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID 
WHERE p.IS_MONITORED = TRUE               -- 모니터링된 프로젝트만 포함
     AND i.DELETED_AT IS NULL             -- 삭제된 문제 제거
     AND ISSUE_STATUS = 'Open'            -- 열린 문제만 포함
     AND ISSUE_SEVERITY IN ('Critical')   -- 크리티컬 문제만 포함
GROUP BY o.DISPLAY_NAME                   -- 원하는 집계에 따라 업데이트
ORDER BY open_issues_aging DESC;          -- 원하는 순서에 따라 업데이트
```

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (2) (19).png" alt="크리티컬 문제의 평균 노후화에 대한 SQL 쿼리 출력 형식"><figcaption><p>크리티컬 문제의 평균 노후화에 대한 SQL 쿼리 출력 형식</p></figcaption></figure>

## MTTR

### 비즈니스 가치

MTTR(Mean Time to Resolve: 평균 해결 시간) 메트릭은 보안 문제를 해결하는 데 걸리는 평균 시간을 추적합니다. 해결된 문제를 기반으로 계산되며 일반적으로 마지막 해결 날짜에 따라 고정 기간(주로 월단위, 분기별 또는 연간 단위)으로 측정됩니다.

MTTR 결과를 분석하면 엔지니어링 팀의 해결 속도에 대한 통찰력을 얻을 수 있습니다. 그러나 긴 기간 동안 열려 있는 문제는 해결되지 않는 한 MTTR 결과에 나타나지 않기 때문에 항상 MTTR 및 노후화를 함께 측정하는 것이 중요합니다.

### 예제 쿼리

아래 쿼리는 마지막 달의 문제 심각도별 MTTR 결과를 Snyk 조직별로 반환합니다.\
결과는 다음을 기반으로 합니다:
- 지난 달에 해결된 문제
- 소음 제거:
  - 모니터링된 프로젝트의 문제만 포함
  - 삭제된 문제가 없음

```sql
SELECT 
    o.DISPLAY_NAME AS organization_display_name,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'Critical' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS critical_mttr,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'High' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS high_mttr,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'Medium' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS medium_mttr,
    ROUND(AVG(CASE WHEN ISSUE_SEVERITY = 'Low' THEN 
        DATEDIFF('DAY', TO_DATE(LAST_INTRODUCED),TO_DATE(LAST_RESOLVED)) ELSE NULL END),2) AS low_mttr
FROM SNYK.SNYK.ISSUES__V_1_0 i
    INNER JOIN SNYK.SNYK.PROJECTS__V_1_0 p ON i.PROJECT_PUBLIC_ID = p.PUBLIC_ID
    INNER JOIN SNYK.SNYK.ORGS__V_1_0 o ON i.ORG_PUBLIC_ID = o.PUBLIC_ID 
WHERE IS_MONITORED = TRUE                       -- 모니터링된 프로젝트만 포함
     AND i.DELETED_AT IS NULL                   -- 삭제된 문제 제거
     AND ISSUE_STATUS = 'Resolved'              -- 해결된 문제만 포함
     -- 지난 달에 해결된 문제
     AND TO_DATE(LAST_RESOLVED) >= DATE_TRUNC('MONTH', DATEADD('MONTH', -12, CURRENT_DATE))
     AND TO_DATE(LAST_RESOLVED) <= DATEADD('DAY', -1, DATE_TRUNC('MONTH', CURRENT_DATE))
GROUP BY organization_display_name
ORDER BY organization_display_name ASC;         -- 원하는 순서에 따라 업데이트
```

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (3).png" alt="시야별 문제 심각도별 MTTR에 대한 SQL 쿼리 출력 형식"><figcaption><p>시야별 문제 심각도별 MTTR에 대한 SQL 쿼리 출력 형식</p></figcaption></figure>

## SLA

### 비즈니스 가치

취약점을 해결하는 것은 중요한 실천 방법이지만 이는 기업의 비즈니스를 선도하는 제품 개발을 늦추게 합니다. 이 간단한 사실 때문에 엔지니어링 팀은 제품 개발 작업을 선호하여 미해결된 취약점을 무시할 수 있습니다.

취약점 교정에 대한 서비스 수준 협약(SLA)을 설정하여 균형을 유지하고 제품 개발을 전진하면서 진화하는 보안 리스크가 명확하고 투명한 정책에 따라 처리되고 있는지 확인합니다.

SLA 목표는 심각도, 자산의 비즈니스 중요성, 코드 소유권 또는 기타 위험 요인 등과 같은 요소를 기준으로 취약점에 대한 허용 가능한 노출 창을 정의합니다.

Snyk 이슈 데이터는 AppSec 팀이 문제 노후화를 추적하고 어떤 취약점이 SLA 목표를 초과했는지 식별하는 데 도움을# 문제
```sql
        AND i.DELETED_AT IS NULL    -- 삭제된 문제 제거
    )
SELECT
    SLA_STATUS,
    SUM(CASE WHEN ISSUE_SEVERITY = 'Critical' THEN 1 ELSE 0 END) AS critical,
    SUM(CASE WHEN ISSUE_SEVERITY = 'High' THEN 1 ELSE 0 END) AS high,
    SUM(CASE WHEN ISSUE_SEVERITY = 'Medium' THEN 1 ELSE 0 END) AS medium,
    SUM(CASE WHEN ISSUE_SEVERITY = 'Low' THEN 1 ELSE 0 END) AS low
FROM base
GROUP BY SLA_STATUS
ORDER BY SLA_STATUS
```

{% hint style="info" %}
예시 쿼리는 다양한 SLA 유스케이스를 지원하기 위해 확장할 수 있습니다. 예를 들어, Snyk 조직 또는 그룹별로 서로 다른 SLA 대상을 정의하거나, 위험에 처한 또는 위반된 문제를 자세히 살펴보고 그들의 문제 해결을 우선시하거나, 다른 비즈니스 부서에 대한 SLA 상태를 분석할 수 있습니다.
{% endhint %}

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (4).png" alt="SLA 상태별 개방된 문제 카운터를 위한 SQL 쿼리의 출력 형식"><figcaption><p>SLA 상태별 개방된 문제 카운터를 위한 SQL 쿼리의 출력 형식</p></figcaption></figure>

## 개발자들의 IDE 및 CLI 테스트 사용 및 채택

### 비즈니스 가치

이 섹션은 Snyk IDE 및 CLI 테스트의 채택을 개발자들로부터 발견하는 방법을 보여줍니다. 개발 단계에서 AppSec 테스트를 구현하는 것은 새로운 보안 위험 사례가 프로덕션에 도달하는 것을 방지하는 가장 비용 효율적인 방법 중 하나로 간주됩니다. 개발자들은 이미 SDLC에서 코드가 더 진행되기 전에 문제에 대처할 적절한 문맥에 있기 때문에 더 효율적입니다. 나중 단계에서 문제를 감지하는 것은 개발자들에게 문제를 다시 방문하고 다른 문맥으로 전환해야 하는 데에 있어 덜 효율적이고 더 많은 시간이 소요될 수 있습니다.

### 예시 쿼리

아래 쿼리는 환경 및 Snyk 제품 별 고유한 개발자 및 총 스캔 횟수를 반환합니다.

결과는 다음에 기반합니다:

* 실행된 테스트
* CI/CD 단계 중 수행된 테스트 제외

```sql
SELECT
    ENVIRONMENT_DISPLAY_NAME AS IDE,
    PRODUCT_DISPLAY_NAME AS PRODUCT,
    COUNT(DISTINCT USER_EMAIL) AS 고유_DEVELOPERS,
    COUNT(1) AS TOTAL_SCANS
from SNYK.SNYK.USAGE_EVENTS__V_1_0
WHERE (RUNTIME_APPLICATION_DATA_SCHEMA_VERSION = 'v2' 
AND ARRAY_CONTAINS('test'::VARIANT, INTERACTION_CATEGORIES) 
AND INTERACTION_STAGE IN ('dev')
		OR RUNTIME_APPLICATION_DATA_SCHEMA_VERSION = 'v1'
	 )
GROUP BY IDE, PRODUCT
```

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (5).png" alt="Snyk 환경별 스캔 수에 대한 SQL 쿼리의 출력 형식"><figcaption><p>Snyk 환경별 스캔 수에 대한 SQL 쿼리의 출력 형식</p></figcaption></figure>

## CI/CD 파이프라인 테스트 사용 및 채택

### 비즈니스 가치

프로덕션에 도달하는 취약점을 방지하기 위해서는 소프트웨어 개발 라이프사이클(SDLC) 전반에 걸쳐 보안 게이트를 배치해야 합니다. 가장 일반적인 게이트 중 하나는 CI/CD 파이프라인 내부에 있으며, 이를 통해 이전 단계에서 놓친 어떤 취약점이라도 빌드 프로세스 중에 잡히고 차단될 수 있습니다.

Snyk Data Share를 활용하면 CI/CD 파이프라인 내에서의 테스트 및 보안 게이트의 현재 채택 상황을 평가할 수 있습니다.

### 예시 쿼리

아래 쿼리는 지난 3개월간 CI/CD 단계에서 실행된 테스트를 기반으로, 각 Snyk 제품별로 테스트된 저장소 수, 총 테스트 수 및 테스트 성공률을 반환합니다.

```sql
SELECT
    PRODUCT_DISPLAY_NAME AS PRODUCT,
    COUNT(DISTINCT INTERACTION_TARGET_ID) AS "테스트된 REPOS",
    COUNT(1) AS "총 SCANS",
    ROUND(((SUM(CASE WHEN INTERACTION_EXIT_CODE=0 THEN 1 ELSE 0 END))/
    (NULLIF(SUM(CASE WHEN INTERACTION_EXIT_CODE IN (0,1) THEN 1 ELSE 0 END),0))
    *100),0) AS "성공률"
FROM SNYK.SNYK.USAGE_EVENTS__V_1_0
WHERE INTERACTION_STAGE != 'cicd'
    AND ARRAY_CONTAINS('test'::VARIANT, INTERACTION_CATEGORIES) 
    AND INTERACTION_TIMESTAMP >= DATE_TRUNC('MONTH', DATEADD('MONTH', -3, CURRENT_DATE))
GROUP BY PRODUCT
```

#### **출력 형식:**

<figure><img src="../../../.gitbook/assets/image (6).png" alt="Snyk 제품별로 테스트된 저장소 수, 총 테스트 및 테스트 성공률에 대한 SQL 쿼리의 출력 형식"><figcaption><p>Snyk 제품별로 테스트된 저장소 수, 총 테스트 및 테스트 성공률에 대한 SQL 쿼리의 출력 형식</p></figcaption></figure>