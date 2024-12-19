# Reachability analysis

{% hint style="info" %}
**릴리즈 상태**

일부 통합 및 언어에 대한 reachability analysis는 초기 액세스 상태입니다.

지원되는 통합 및 언어에 대한 기능을 활성화하는 방법에 대한 자세한 정보는 [Snyk 미리보기](../../snyk-admin/snyk-preview.md)를 참조하세요.
{% endhint %}

Snyk reachability analysis를 사용하면 응용 프로그램이 취약점과 관련 있는 코드 요소(예: 함수, 클래스, 모듈, 주석 등)를 호출하는지 식별하여 해당 취약점이 응용 프로그램의 문맥에서 악용될 가능성을 높일 수 있습니다.

Reachability analysis는 결정을 내리는 데 독립적인 신호로 사용하거나 Snyk 위험 점수를 활용한 더 광범위한 위험 중심 우선순위 접근 방식의 일부로 사용할 수 있습니다.

Snyk는 프로그램 분석과 다양한 AI 기술의 조합을 사용하여 특정 취약점의 reachability를 결정하며, 보안 연구 전문가에 의해 확인됩니다.

다음 지침은 reachability analysis를 설정하고 사용하는 방법과 Snyk에서 기본 분석이 작동하는 방법에 대해 자세히 설명합니다.

## Reachability analysis 설정

다음 단계를 따라 Reachability analysis를 활성화하고 도달 가능한 취약점을 분석하기 시작하십시오:

* 조직 **설정**에서 **{{Snyk 오픈 소스}}** 섹션으로 이동합니다.
* **일반** 섹션에서 **Reachability analysis**를 찾습니다.
* **Enable reachability analysis**를 활성화합니다.

<figure><img src="../../.gitbook/assets/image (573).png" alt=""><figcaption><p>Reachability 설정 활성화</p></figcaption></figure>

Reachability analysis가 활성화되면 분석이 프로젝트 스캔의 일부로 수행됩니다.

{% hint style="info" %}
기존 프로젝트에 reachability analysis를 적용하려면 [수동 테스트](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/#manual-pull-and-merge-requests-for-project-code)를 트리거하십시오.
{% endhint %}

## 지원되는 언어 및 통합

다음 언어와 패키지 관리자에 대해 Reachability analysis가 지원됩니다:

| 언어                                                                                                                                                                   | 패키지 관리자     | 릴리즈 상태       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------- | -------------------- |
| [Java](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/)                                                                                         | Maven, Gradle       | 일반 사용 가능 |
| [JavaScript](../../supported-languages-package-managers-and-frameworks/javascript/), [TypeScript](../../supported-languages-package-managers-and-frameworks/typescript.md) | npm, Yarn           | 초기 액세스         |
| [Python](../../supported-languages-package-managers-and-frameworks/python/)                                                                                                | pip, poetry, pipenv | 초기 액세스         |

Reachability analysis는 다음 통합에서 지원됩니다:

| 통합                                                                                                                                                              | 릴리즈 상태       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------- |
| [GitHub](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md)                                                                                           | 일반 사용 가능 |
| [GitHub Enterprise](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-scm-contributors-count/scripts-for-scm-contributors-count/github-enterprise/) | 일반 사용 가능 |
| [Bitbucket Cloud](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-cloud-app.md)                                                                     | 일반 사용 가능 |
| [Bitbucket Server](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-data-center-server.md)                                                           | 일반 사용 가능 |
| [GitLab](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab.md)                                                                                           | 일반 사용 가능 |
| [Azure Repos](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/azure-repositories-tfs.md)                                                                      | 일반 사용 가능 |
| [Brokered connections](../../enterprise-setup/snyk-broker/connections-with-snyk-broker.md)                                                                               | 일반 사용 가능 |

{% hint style="info" %}
Snyk CLI, IDE 또는 기타 통합을 사용한 Reachability analysis는 지원되지 않습니다.
{% endhint %}

## Brokered connections를 사용한 Reachability analysis 활성화

만약 SCM에 대해 Brokered 연결을 사용한다면, Broker를 구성하여 소스 파일에 액세스하도록 설정하세요. [Broker를 통한 Git Clone](../../enterprise-setup/snyk-broker/git-clone-through-broker.md)에서 Snyk Code와 함께 Broker를 사용하는 구성 세부 정보를 확인하세요.

## Reachability analysis 사용

### Reachability 상태

취약점이 식별된 후, 해당 취약점의 reachability 상태 중 하나가 할당됩니다:

* `REACHABLE` - 응용 프로그램에서 취약한 코드로부터 직접 또는 간접적인 경로가 발견되었습니다.
* `NO PATH FOUND` - 응용 프로그램에서 취약한 코드로의 경로를 찾을 수 없습니다.

{% hint style="info" %}
`NO PATH FOUND` 상태가 제공된 경우 취약점이 완전히 접근할 수 없거나 악용할 수 없다고 가정해서는 안 됩니다.
{% endhint %}

Reachability analysis 상태는 [프로젝트 페이지](reachability-analysis.md#on-the-project-page), [위험 점수 일부로](reachability-analysis.md#as-part-of-the-risk-score), [문제 상세 보고서](../../manage-issues/reporting/available-snyk-reports.md#issues-detail-report) 및 API 엔드포인트 [그룹 ID별 문제 가져오기](../../snyk-api/reference/issues.md#groups-group_id-issues)에서 확인할 수 있습니다.

### 프로젝트 페이지에서 보이는 Reachability analysis

Snyk UI를 사용하여 프로젝트를 가져오거나 테스트한 후, 프로젝트는 Snyk에 의해 모니터링되며 도달 가능한 취약점 분석 결과가 프로젝트 페이지에 다음 위치에 표시됩니다:

* 필터 - 도달 가능한 취약점을 우선적으로 확인하기 위해 도달 가능성에 따라 결과를 필터링할 수 있습니다.
* Reachability 뱃지 - 취약점의 도달 가능성 수준을 이해할 수 있습니다.
* 호출 경로 - 코드에서 취약한 코드 요소로의 경로를 확인하여 결과를 검증할 수 있습니다.

<figure><img src="../../.gitbook/assets/image (124) (1) (1) (1) (2) (1) (1) (1) (2) (2).png" alt="프로젝트 UI에서의 Reachability 필터, 뱃지 및 호출 경로"><figcaption><p>프로젝트 UI에서의 Reachability 필터, 뱃지 및 호출 경로</p></figcaption></figure>

### 위험 점수의 일부로서의 Reachability analysis

[Risk Score](risk-score.md)는 취약점 또는 응용 프로그램의 문맥과 관련된 여러 요소를 결합하는 전체적인 위험 중심 우선순위를 적용하는 데 도움이 되며, Reachability analysis는 전체 점수를 크게 높일 수 있는 이러한 문맥 요소 중 하나입니다.

Risk Score는 프로젝트 페이지에서 사용할 수 있으며 API 및 보고서를 통해 확인할 수 있습니다.

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (1) (7).png" alt="Risk Score의 일부로서의 Reachability"><figcaption><p>Risk Score의 일부로서의 Reachability</p></figcaption></figure></div>

{% hint style="info" %}
[우선 순위 점수](priority-score.md), Risk Score 이전의 레거시 모델은 도달 가능한 취약점을 고려합니다.
{% endhint %}

## 도달 가능한 취약점 분석 방법

Reachable vulnerability analysis가 작동하는 방법은 다음과 같습니다:

Snyk는 취약점을 결정하기 위해 보안 전문가 분석, 프로그램 분석 및 다양한 AI 기술의 조합을 사용하며 다음과 같은 분석 단계가 포함됩니다:

1. **취약점에 수정 패치 적용** - Snyk 취약점 정제 프로세스의 일환으로 Snyk는 유지 관리자가 적용한 수정 커밋을 참조합니다.
2. **관련 요소 분석** - 커밋 정정한 기준으로, Snyk는 DeepCode AI 프로그램 분석을 사용하여 취약점과 관련된 코드 요소와 다른 매개 변수를 분석합니다.
3. **근본원인 분석** - Snyk는 DeepCode AI 및 NLP 기술을 사용하여 관련 코드 요소를 자동으로 우선 순위로 매기고 취약점의 작동 가능성을 확인합니다.
4. **Reachability analysis** - Snyk 스캔에 의해 응용 프로그램에서 문제가 발견되면, 응용 프로그램의 호출 그래프를 분석하여 사용된 오픈 소스 종속성간의 호출 그래프를 분석하는 데 DeepCode 프로그램 분석 엔진을 사용합니다. 응용 프로그램과 루트 원인으로 순위 매겨진 코드 요소 사이의 경로는 "도달 가능" 취약점을 제공할 것입니다.
5. **보안 전문가 감독** - Snyk 보안 전문가는 요소를 루트 원인으로 표시하고 전체 분석을 시간이 지남에 따라 더 정확하게 만들기 위해 수동으로 확인합니다.

**잘못된 양성 및 잘못된 음성**과 관련된 다음 고려 사항이 Reachable vulnerability analysis에 적용됩니다.

프로그램 분석은 잠재적으로 악용 가능한 취약점을 피하기 위해 정확한 결과, 잘못된 양성을 최소화하고 회수율을 최대화하는 간극에서 무엇을 할지를 따져야 합니다.

이 간극을 용이하게 하기 위해, Snyk DeepCode 분석은 사용 환경에서 도달 가능한 경로가 찾아질 가능성 분석을 기반으로 실시간 의사 결정을 적용합니다.

예를 들어, 리플렉션 프로그래밍이 사용될 때 항상 정확한 답변을 제공할 수 있는 것은 아닙니다. 이 경우, 잘못된 양성 집합을 반환하든 "도달 불가"를 반환하든 적절하지 않습니다. Snyk Deep Code 분석은 특정 코드 구조에 대해 가능한 최정확하고 완전한 결과를 검색하도록 최적화합니다.