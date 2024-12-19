# 문제 해결을 위한 우선순위 설정

우선 순위 설정은 Snyk에서 여러 이름으로 사용되며 Snyk 계획에 따라 다양한 사용자 정의가 있습니다. 다음 목록은 우선순위 설정을 찾고 사용할 수 있는 모든 위치에 대한 개요를 제시합니다. 출력물은 사용하는 우선순위 유형에 따라 다양할 수 있습니다.&#x20;

* [프로젝트 수준에서의 우선 순위 설정](./#prioritization-at-the-project-level)
* [보고서 내 우선 순위 설정](./#prioritization-within-reporting)
* [위험에 기초한 우선 순위 설정](./#prioritization-based-on-risk)
* [보안 프로그램 및 커버리지에 기초한 우선 순위 설정](./#prioritization-based-on-coverage-and-security-program)

특정 유형의 우선순위에만 집중할 수도 있고, 즉시 주의를 요하는 문제에 더 복잡하고 명확한 포커스를 맞출 수 있도록 이들을 결합할 수도 있습니다. 우선순위 설정은 취약성, 보안 프로그램 및 커버리지(보안 관점에서)에 대해 사용하거나 프로젝트의 Reporting 섹션에서 사용할 수 있습니다.\
\
예를 들어, 리포지토리의 프로젝트만 우선 순위를 설정하려고 한다면, 이슈 목록은 더 일반적일 것입니다. 그러나 Snyk AppRisk Pro용 Insights와 함께 우선순위를 사용하는 경우, 리스크 팩터 및 자산을 통합하여 분석된 이슈는 더 복잡하고 명확한 필터에 따라 우선 순위 목록이 생성되며 필요에 따라 우선순위 설정을 사용자 정의할 수 있게 됩니다.&#x20;

## 프로젝트 수준에서의 우선 순위 설정

### 프로젝트 우선 순위 설정

이슈의 심각성에 따라 프로젝트를 우선적으로 처리하고자 하시는 경우에는 [프로젝트 우선 순위 설정](../../snyk-admin/snyk-projects/project-information.md)을 통해 리포지토리에서 프로젝트를 우선적으로 처리할 수 있습니다.

이 유형의 우선 순위 설정을 사용하면 프로젝트를 그룹화, 필터링 및 정렬할 수 있습니다.

이러한 필터에 기초하여 프로젝트의 우선 순위를 설정할 수 있습니다:

- 이슈가 있는 경우: 이슈가 있는 프로젝트만 표시
- 이슈가 없는 경우: 이슈가 없는 프로젝트만 표시
- 통합: Snyk에 가져온 통합된 Git 리포지토리를 표시

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk Free/Team Plan
- Snyk Enterprise Plan
- Snyk AppRisk Essentials
- Snyk AppRisk Pro

Snyk 그룹 및 Snyk 조직 수준에서 프로젝트의 우선 순위를 설정할 수 있습니다.

### 프로젝트 내 이슈 우선순위 설정

프로젝트 내의 이슈를 보기 위해 [이슈 우선순위 설정](../../snyk-admin/snyk-projects/view-project-issues-fixes-and-dependencies.md)을 사용하여 다음과 같은 사용자 지정 필터 중 하나를 사용하십시오:&#x20;

- 이슈 유형
- 심각성&#x20;
- 수정 가능성
- 공격성 성숙도
- 상태

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk Free/Team Plan
- Snyk Enterprise Plan
- Snyk AppRisk Essentials
- Snyk AppRisk Pro

Snyk 조직 수준에서 프로젝트의 우선 순위를 설정할 수 있습니다.

## 보고서 내 우선순위 설정

### 우선 순위 점수 및 리스크 점수 설정

Snyk Web UI에서 Reporting 옵션을 사용하고 필터 목록에서 원하는 프로젝트를 선택할 때 [우선 순위 및 리스크 점수 설정](./#prioritization-within-reporting)을 사용할 수 있습니다.&#x20;

우선 순위 점수는 긴급성에 초점을 맞추며 CVSS 점수, 트렌드 있는 취약점, 도달 가능성, 악용 가능성 및 기타 요소 등 다양한 기준에 따라 보안 이슈를 순위별로 정렬합니다. 반면 리스크 점수는 취약점과 관련된 전반적인 위험을 평가하여 심각성과 응용프로그램의 맥락을 고려하여, 공격 가능성 및 시스템에 미칠 잠재적인 영향과 같은 요소를 고려합니다. 우선 순위 점수는 팀이 가장 긴요한 위협에 대응하는 데 도움을 주는 반면, 리스크 점수는 보안 포지션에 대해 종합적인 보기를 제공합니다.

#### 우선 순위 점수

[우선 순위 점수](priority-score.md)는 긴급성에 따라 보안 취약점을 순위별로 정렬하여 팀이 중요한 보안 취약점을 신속하게 식별하고 해결할 수 있도록 돕습니다.&#x20;

우선 순위 점수는 다음과 같은 여러 산업 표준 기준을 바탕으로 결정됩니다:

- CVSS 점수
- 트렌드 있는 취약성
- 도달 가능성
- 악용 가능성
- 기타 요소

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk Free/Team Plan
- Snyk Enterprise Plan
- Snyk AppRisk Essentials
- Snyk AppRisk Pro

Snyk 조직 수준에서 우선 순위 점수를 사용할 수 있습니다.

#### 리스크 점수

[리스크 점수](risk-score.md)는 취약점들의 잠재적인 영향을 평가하고 심각한 결과를 초점으로 하는 취약점들을 우선적으로 처리합니다. 리스크 점수를 사용하여 각 보안 이슈에 대한 잠재적인 영향과 악용 가능성을 고려하여 자동적인 리스크 분석을 수행할 수 있습니다.

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk Free/Team Plan
- Snyk Enterprise Plan
- Snyk AppRisk Essentials
- Snyk AppRisk Pro

Snyk 조직 수준에서 리스크 점수를 사용할 수 있습니다.

### 보고서 사용 시 이슈 우선 순위 설정

[보고서 사용 시 이슈를 우선 순위로 설정](../../manage-issues/reporting/available-snyk-reports.md)하여 조직 또는 그룹 전체에서 보고서를 생성할 수 있습니다. 

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk Enterprise Plan
- Snyk AppRisk Essentials
- Snyk AppRisk Pro

Snyk 그룹 또는 Snyk 조직 수준에서 리스크 점수를 사용할 수 있습니다.

## 위험에 기초한 우선순위 설정

Snyk AppRisk을 위한 [Insights를 활용한 우선순위 설정](prioritization-for-snyk-apprisk.md) - Snyk AppRisk는 애플리케이션 내부 전체적인 인텔리전스를 활용하여 컨테이너, 코드 및 오픈소스 문제를 식별하고 우선순위를 정하도록 도와줍니다. Snyk AppRisk 정책에서 정의된 자산 분류를 기반으로 문제에 우선순위를 부여할 수 있습니다.&#x20;

우선 순위 설정을 사용하여 애플리케이션에 대한 위험이 되는 컨테이너, 코드 및 오픈 소스 문제를 식별하고 우선순위를 정할 수 있습니다. 다음과 같은 위험 요인을 사용하여 이러한 문제를 우선순위로 설정할 수 있습니다:

- 배포됨
- 로드된 패키지
- OS 상태
- 공개

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk AppRisk Pro

Snyk 그룹 또는 Snyk 조직 수준에서 Insights를 사용하여 우선순위를 설정할 수 있습니다.

## 커버리지 및 보안 프로그램에 기초한 우선순위 설정

Snyk Enterprise Analytics는 테넌트 수준에서 전체 애플리케이션 보안 프로그램에 대한 포괄적인 개요를 제공하며, 강점, 약점 및 기회를 식별합니다. 반면, Snyk Application Analytics는 특정 애플리케이션의 보안 및 성능에 특히 초점을 두는데, 이 더 좁은 분석은 특정 앱의 취약점과 비효율성을 정확히 파악하여 타겟팅된 개선과 좀 더 자세한 통찰을 가능하게 합니다. 

### 이슈 분석

[이슈 분석](../enterprise-analytics/issues-analytics.md)은 AppSec 프로그램의 성능을 보여줍니다. Enterprise Analytics를 사용하여 프로그램의 강점과 약점을 더 잘 이해하고 성공적인 실천 방법이 어디에 식별되는지, 투자가 필요한 가장 큰 개선 기회를 발견할 수 있습니다.

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk Enterprise
- Snyk AppRisk Pro

Snyk 그룹 및 Snyk 조직 수준에서 Enterprise Analytics를 사용할 수 있습니다.

{% hint style="info" %}
만약 {% raw %}Snyk Enterprise{% endraw %}와 {% raw %}Snyk AppRisk Pro{% endraw %} 플랜을 가지고 계신다면, [Issues](../enterprise-analytics/issues-analytics.md) 탭 안에 있는 Analytics 옵션 하단에 Enterprise Analytics를 찾을 수 있습니다. 
{% endhint %}

### 애플리케이션 분석

Snyk AppRisk용 [애플리케이션 분석](../enterprise-analytics/application-analytics.md)은 상향식 방법으로부터 AppSec 프로그램을 평가할 수 있도록 돕습니다. 애플리케이션, 팀 및 자산 분류를 검토하고 구체적인 자산에 초점을 맞출 수 있습니다. 보안을 강화하기 위해 향상 영역, 위험 인지, 블라인드 스팟을 식별할 수 있습니다. 이 도구는 테넌트의 모든 그룹에서 데이터를 검색합니다.

표시된 데이터의 우선 순위는 사용 가능한 필터, 차원 뷰 및 특정 시간대를 사용하여 설정할 수 있습니다. \
\
필터:

- 그룹
- 이슈 심각성
- 자산 기반 필터

뷰:

- 자산 클래스
- 애플리케이션
- 소유자

이 유형의 우선 순위 설정은 다음 {% raw %}Snyk Plans{% endraw %}에서 사용할 수 있습니다:&#x20;

- Snyk AppRisk Pro

Snyk 그룹 및 Snyk 조직 수준에서 Enterprise Analytics를 사용할 수 있습니다.

## 우선순위 전략

Snyk에는 검색된 이슈 중에서 어떤 것이 가장 중요한 이슈이며 이를 해결하기 위한 우선순위와 순서를 결정하는 데 도움이 되는 여러 기능이 있습니다.

{% hint style="info" %}
이슈를 무시하고 제외하는 방법에 대한 정보는 [.snyk 파일](../policies/the-.snyk-file.md)을 참조하고 프로젝트에 정책을 생성하고 할당하는 방법을 설명하는 [Policies](../policies/) 페이지를 참조하십시오.
{% endhint %}

일부 도구는 심각도만을 사용하여 이슈의 우선순위를 설정하지만, 이는 수천 개의 결과물이 발생하여 이러한 이슈를 해결하는 명확한 시작점이 없을 수 있습니다.

특정 프로젝트를 보는 경우 프로젝트 수준에서 우선순위를 설정할 수 있습니다. 엔터프라이즈 고객은 모든 프로젝트를 대상으로 우선순위를 설정할 수 있습니다.

Snyk 우선 순위 점수 및 리스크 점수는 문제의 [심각성](severity-levels.md) 및 해당 문제를 해결해야 하는 긴급성을 순위별로 결정합니다