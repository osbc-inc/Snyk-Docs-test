# REST 문제 실험 API에서 GA API로의 이주 가이드

{% hint style="info" %}
**실험 API에 관한 중요한 정보**

실험적 엔드포인트는 불안정하며 기술 프리뷰로 간주해야 합니다. 실험 버전은 파괴적인 변경을 가져올 수 있으며 언제든지 중단될 수 있습니다.
{% endhint %}

## GA REST Issues API에서 무엇이 새로운가요?

{% hint style="info" %}
GA REST Issues API 문서 [`/groups/{group_id}/issues`](https://apidocs.snyk.io/#get-/groups/-group_id-/issues) 및 [`/orgs/{org_id}/issues`](https://apidocs.snyk.io/#get-/orgs/-org_id-/issues).
{% endhint %}

이 API 버전은 다음을 제공합니다:

* **일관성:** REST Issues API의 성능 및 신뢰성 향상
* **깊이:** 오픈 소스 패키지 및 수정에 대한 상세한 표현
* **유연성:** 맞춤형 API 응답을 위한 새로운 필터
* **사용 편의성:** API 상호 작용을 단순화하는 향상된 페이지네이션 및 응답 관리

Snyk은 새로운 API로의 이주가 중요한 작업일 수 있으며, 이 과정 전반에서 지원하고자 합니다. 이 포괄적인 이주 가이드는 단계별 지침, 코드 예제 및 새 API와의 원활한 통합을 돕기 위한 모범 사례를 제공하여 신속한 전환을 지원하는 것을 목적으로 합니다.

폐기된 엔드포인트를 사용 중이라면, Snyk은 본 이주 가이드를 검토하고 모든 자동화 기능을 전환하도록 권장합니다.

## 실험 및 GA API 비교

{% hint style="info" %}
**실험 API 문제를 GA API 문제로 매핑**

아래 표에서 볼 수 있는 주요 차이점 중 하나는 문제의 ID 형식이 실험 API의 URI 형식(키 및 scan_item.id로 구성됨)에서 GA API의 UUID로 변경된다는 것입니다. 실험 API 응답의 문제를 GA API 응답의 동일한 문제와 일치시키려면 **key** 및 **scan_item.id**를 사용할 수 있습니다. **scan_item**은 **relationships** 블록의 일부이고 **key**는 **attributes** 블록의 일부입니다.
{% endhint %}

| 필드                                    | 실험 API               | GA                          |
| -------------------------------------- | ------------------- | --------------------------- |
| classes                                | 존재                  | 변경 없음                   |
| coordinates                            | 클라우드 문제에만 사용 가능    | 클라우드 및 SCA 문제에 사용 가능 및 새로운 수정 가능 필드 포함    |
| coordinates.is_fixable_manually        | coordinates.is_fixable_snyk | coordinates.is_fixable_upstream<br>coordinates.is_patchable<br>coordinates.is_pinnable<br>coordinates.is_upgradeable   |             
| Not present                            | 새로 도입된 수정 가능 데이터    |
| coordinates.reachability                | 존재                  | 새로 도입됨                 |
| coordinates.remedies                    | 존재                  | 변경 없음                   |
| representations                        | 존재                  | 새로운 필드                  |
| representations.resourcePath            | 존재                  | 변경 없음                   |
| representations.dependencyChain         | 존재                  | representations.dependency 대신 제거 |
| representations.dependency              | 존재하지 않음            | 새로 도입됨 (representations.dependency 대체)    |
| representations.dependency.package_name | 존재하지 않음            | 새로 도입됨 (representations.dependency 일부)    |
| cloud_resource                         | 존재                  | 변경 없음                   |
| sourceLocation                         | 존재                  | 변경 없음                   |
| created_at                             | 존재                  | 변경 없음                   |
| description                            |                    | 변경 없음                   |
| effective_severity_level               | 존재                  | 변경 없음                   |
| ignored                                | 존재                  | 변경 없음                   |
| key                                    | 존재                  | 변경 없음                   |
| priority                                | 존재                  | 제거 및 risk로 대체                       |
| priority.factors                        | 존재                  | risk.factors로 대체                 |
| priority.score                        | 존재                  | risk.score로 대체                     |
| risk                                    | 존재하지 않음           | 새로 도입됨 - priority로 대체           |
| risk.factors                        | 존재하지 않음           | 새로 도입됨 - priority.factors로 대체   |
| risk.factors[i].included_in_score        | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].links                    | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].links.evidence            | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].links.evidence.href        | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].links.evidence.meta        | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].name                    | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].updated_at                | 존재하지 않음           | 새로 도입됨                         |
| risk.factors[i].value                    | 존재하지 않음           | 새로 도입됨                         |
| risk.score                            | 존재하지 않음          | 새로 도입됨 - priority.score 대체        |
| risk.score.model                        | 존재하지 않음           | 변경 없음                           |
| problems                                | 존재                  | 변경 없음                           |
| resolution                                | 클라우드 문제에만 사용 가능    | 모든 문제 유형에 대해 사용 가능                    |
| severities                            | 존재 - 값이 없음        | 값이 없음이므로 제거되었으며 그대로 유지되지 않을 것임     |
| status                                | 존재                  | 변경 없음                             |
| title                                | 존재                  | 변경 없음                             |
| tool                                    | 존재                  | 변경 없음                             |
| type                                    | 존재                  | 변경 없음                             |
| updated_at                            | 존재                  | 변경 없음                             |
| id                                    | URI 형식 (키 & scan_item.id로 구성) | UUID 형식                   |
| relationships.ignore                    | 존재                  | 변경 없음                             |
| relationships.organization                | 존재                  | 변경 없음                             |
| relationships.policies                    | 존재 - 값이 없음        | 값이 없음이므로 제거되었으며 그대로 유지되지 않을 것임     |
| relationships.previous                    | 존재 - 값이 없음        | 값이 없음이므로 제거되었으며 그대로 유지되지 않을 것임     |
| relationships.scan_item                | 존재                  | 변경 없음                             |
| relationships.test_execution            | 존재                  | 변경 없음                             |

## 실험 vs GA API 필터 비교

| 필터명                        | 목적                                                    | 실험 API | GA       |
| ---------------------------- | ------------------------------------------------------- | -------- | -------- |
| starting\_after              | 이전 커서 바로 다음의 결과 페이지를 반환                   | 존재       | 변경 없음 |
| ending\_before               | 이전 커서 바로 이전의 결과 페이지를 반환                   | 존재       | 변경 없음 |
| limit                        | 페이지당 반환할 결과 수                                     | 존재       | 변경 없음 |
| scan\_item.id                | 스캔 항목 관계를 통해 문제를 필터링                        | 존재       | 변경 없음 |
| scan\_item.type              | 스캔 항목 관계를 통해 문제를 필터링                        | 존재       | 변경 없음 |
| type                         | 문제 유형에 따른 필터링                                    | 존재       | 변경 없음 |
| updated\_after               | 이 날짜 이후에 업데이트된 문제 필터링                        | 존재       | 변경 없음 |
| updated\_before              | 이 날짜 이전에 업데이트된 문제 필터링                       | 존재 없음   | 새로 도입됨 |
| created\_before              | 이 날짜 이전에 생성된 문제 필터링                          | 존재 없음   | 새로 도입됨 |
| created\_after               | 이 날짜 이후에 생성된 문제 필터링                          | 존재 없음   | 새로 도입됨 |
| effective\_severity\_level   | 하나 이상의 유효한 심각도 수준에 따라 문제 필터링            | 존재하지 않음 | 새로 도입됨 |
| status                       | 문제 상태에 따른 필터링                                   | 존재하지 않음 | 새로 도입됨 |
| ignored                      | 무시된 상태의 문제 필터링                                  | 존재하지 않음 | 새로 도입됨 |

[^1]: 현재 문제 API에서의 유일한 위험 요인은 통찰 요인입니다.

[^2]: 위험 점수를 활성화한 고객의 경우 위험 점수여야 하지만 위험 요인은 없어야 합니다. 위험 점수를 활성화하지 않은 고객의 경우 이는 이전 우선 순위 점수일 것입니다.

[^3]: 위에서 언급한 대로 실험 API 응답의 문제를 GA API 응답의 동일한 문제와 일치시키려면 key 및 scan\_item.id를 사용할 수 있습니다. scan\_item은 relationships 블록의 일부이고 key는 attributes 블록의 일부입니다.