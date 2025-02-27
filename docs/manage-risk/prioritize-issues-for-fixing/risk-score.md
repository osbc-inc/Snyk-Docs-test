# Risk Score(위험 점수)

{% hint style="info" %}
**릴리스 상태**

위험 점수는 초기 액세스로 및 에서 사용할 수 있으며, Snyk 엔터프라이즈 및 Snyk 무료 요금제에서 사용할 수 있습니다. 귀하의 그룹에 설정하려면 Snyk 계정 팀과 연락하십시오.

새로운 위험 점수로 및 문제를 위한 우선순위 점수를 대체하려면 [Snyk 미리보기](https://docs.snyk.io/snyk-admin/manage-settings/snyk-preview)를 사용하세요.
{% endhint %}

Snyk 위험 점수는 취약점 유형 모든 문제에 대해 자동 위험 분석에 의해 적용된 단일값입니다. 라이선스 문제는 지원되지 않습니다. 위험 점수는 잠재적 영향 및 악용 가능성을 기반으로 합니다. 0부터 1,000까지의 범위로, 점수는 환경에 부과되는 위험을 나타내며 위험 중심 우선순위 접근 방식을 가능하게 합니다.

위험 점수는 기여 요인이 변경되지 않는 한 시간이 지나도 동일합니다. 그러나 악용 예측 점수 시스템 (EPSS)과 같은 일부 요인은 자주 변경될 수 있습니다. 취약점이 처음 발표된 이후의 일수도 요인이며, 일수가 1년 이상이 되고 가능성 하위 점수가 감소할 때에만 점수가 변경됩니다.

실제 위험이 드물기 때문에, Project 점수 분포의 중요한 변동이 예상되며, 예시로 볼 수 있는 Project 점수 분포는 다음과 같습니다:

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (2) (1) (1) (1).png" alt="예제 프로젝트 점수 분포"><figcaption><p>예제 프로젝트 점수 분포</p></figcaption></figure></div>

위험 점수는 우선순위 점수를 직접 대체합니다. Risk Score 를 상호 작용하는 방법은 [우선순위 점수 문서](priority-score.md)를 참조하십시오. Risk Score 가 활성화되면 UI, API 및 보고서에서 Risk Score 가 도입된 부분을 확인할 수 있습니다.

CLI에서는 Risk Score를 사용할 수 없습니다.

{% hint style="info" %}
위험 점수는 및 프로젝트가 다시 테스트된 후에만 우선 순위 점수를 대체합니다.
{% endhint %}

{% hint style="warning" %}
API에서 해당 필드는 여전히 `priority.`로 이름 지어져 있습니다. Risk Score가 활성화되면 이러한 필드에 채워진 점수와 요소가 초기 액세스 단계의 Risk Score 모델에 따라 지정됩니다.
{% endhint %}

## 문제 별 Risk Score 탐색

Issue 카드 정보를 보다 보면, 점수 위에 마우스를 가져가면 현재 표시 중인 점수 유형 (우선 순위 또는 위험 점수)를 볼 수 있습니다. 위험 점수 툴팁은 서브스코어와 위험 점수에 기여하는 위험 요소에 대한 정보를 제공합니다.

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (118) (2).png" alt="Risk Score tooltip" width="563"><figcaption><p>영향 및 가능성을 나타내는 Risk Score 툴팁</p></figcaption></figure></div>

## Risk Score 모델 소개

<figure><img src="../../.gitbook/assets/matrix (2).png" alt="he Snyk Risk Score Model"><figcaption><p>Snyk Risk Score 모델</p></figcaption></figure>

Risk Score를 적용하는 모델은 잠재적 영향 및 악용 가능성에 기초하여 각 보안 문제에 대해 자동 위험 분석을 적용합니다.

{% hint style="info" %}
Risk 모델은 Snyk 보안 데이터 과학팀과 숙련된 보안 연구원들이 수행한 폭넓은 연구를 통해 도출되었습니다. 이 모델은 [Snyk 취약점 데이터베이스](https://security.snyk.io/) 개발 과정에서 얻은 전문 지식을 기반으로 합니다.
{% endhint %}

### 영향 서브스코어

객관적인 영향 요인은 CVSS 영향 메트릭스, 가용성, 기밀성, 무결성 및 범위를 기준으로 하여 CVSS 영향 서브스코어를 기반으로 계산됩니다. 컨테이너 문제의 경우, 공급자 긴급도도 고려됩니다.

사업 비전요 속성은 상황 영향 요인으로 고려되어 영향 서브스코어를 증가 또는 감소시킵니다. 자세한 내용은 [프로젝트 속성](../../snyk-admin/snyk-projects/project-attributes.md)을 참조하십시오.

### 가능성 서브스코어

객관적 가능성 요인이 고려되며 다음이 포함됩니다:

* 악용 성숙도
* 악용 예측 점수 시스템 (EPSS)
* 공고의 연령
* CVSS 악용 가능성 메트릭스: 공격 벡터, 필요한 권한, 사용자 상호작용 및 범위
* 소셜 트렌드
* 해로운 패키지
* 공급자 긴급도 (Snyk Container)
* 패키지 인기도 (Snyk Open Source)

그런 다음 맥락적 가능성 요소가 가능성 서브스코어를 증가 또는 감소시킵니다:

* 도달 가능성 (Snyk Open Source Java, JavaScript)
* 서술적 깊이

{% hint style="info" %}
Fixability는 점수 계산의 일부로 고려되지 않으며, 보안 문제를 해결하는 데 필요한 노력이 위험에 미치는 영향은 없습니다. 실행 가능한 문제에 중점을 두기 위해 Fixability 필터를 사용한 다음 Risk Score로 가장 위험한 문제부터 시작하십시오.
{% endhint %}

## 위험 요소 자세히 보기

### 객관적 영향 위험 요인

#### 기밀성

고객의 데이터 기밀성에 미치는 영향을 나타내며 [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-3-1-Confidentiality-C)에 기초합니다. **가능한 입력 값:** `None`, `Low`, `High`

#### 무결성

고객의 데이터 무결성에 미치는 영향을 나타내며 [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-3-2-Integrity-I)에 기초합니다. **가능한 입력 값:** `None`, `Low`, `High`

#### 가용성

고객의 응용프로그램 가용성에 미치는 영향을 나타내며 [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-3-3-Availability-A)에 기초합니다. **가능한 입력 값:** `None`, `Low`, `High`

#### 범위

취약점이 대상의 보안 범위 외의 구성요소에 영향을 줄 수 있는지 여부를 나타냅니다. [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-2-Scope-S)를 기초로 합니다. 객관적 영향 서브스코어는 CVSS 영향 서브스코어를 기반으로 계산됩니다. 자세한 내용은 위에서 언급한 CVSS 정의와 [서브 스코어 방정식](https://www.first.org/cvss/v3.1/specification-document#7-1-Base-Metrics-Equations)을 참조하십시오.

| 가능한 입력 값    | 점수 영향                 |
| ----------- | --------------------- |
| `Unchanged` | 영향 서브스코어에 영향을 미치지 않음. |
| `Changed`   | 영향 서브스코어가 영향을 받음.     |

#### 제공자 긴급도 (Snyk Container)

해당 운영 체제 배포 보안 팀에 의해 제공되는 긴급도 등급입니다. 자세한 내용은 [상대적 중요도를 위한 외부 정보 소스](../../scan-with-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md#external-information-sources-for-relative-importance)를 참조하십시오.

| 가능한 입력 값   | 점수 영향             |
| ---------- | ----------------- |
| `Critical` | 영향 서브스코어가 크게 증가함. |
| `High`     | 영향 서브스코어가 증가함.    |
| `Medium`   | 영향 서브스코어가 크게 감소함. |
| `Low`      | 영향 서브스코어가 크게 감소함. |

{% hint style="info" %}
제공자 긴급도는 가능성 서브스코어에 영향을 미칩니다.
{% endhint %}

### 맥락적 영향 리스크 요소

### **사업 중요성**

해당 응용프로그램의 주관적 비즈니스 영향을 나타내는 사용자 정의 프로젝트 속성입니다. 자세한 내용은 [프로젝트 속성](../../snyk-admin/snyk-projects/project-attributes.md)을 참조하십시오.

| 가능한 입력 값   | 점수 영향                |
| ---------- | -------------------- |
| `Critical` | 영향 서브스코어가 증가함.       |
| `High`     | 영향 서브스코어가 영향을 받지 않음. |
| `Medium`   | 영향 서브스코어가 감소함.       |
| `Low`      | 영향 서브스코어가 크게 감소함.    |

{% hint style="info" %}
프로젝트에 사업 비전요 속성을 적용하면 새 데이터가 포함된 위험 점수를 반영하기 위해 리트스트가 필요합니다. 사업 비전요가 할당되지 않은 경우 영향 서브스코어는 영향을 받지 않습니다.
{% endhint %}

프로젝트의 사업 비전요가 구성되지 않은 경우, 하위 점수가 영향을 받지 않도록 기본값인 `고`가 사용됩니다.

### 객관적 가능성 리스크 요소

### 악용 성숙도 (Exploit maturity)

Snyk이 검색하고 확인한 공개 악용의 존재 및 성숙도를 나타냅니다. 자세한 내용은 [악용 보기, 악용 확인 방법](view-exploits.md#how-exploits-are-determined)을 참조하십시오.

| 가능한 입력 값           | 점수 영향             |
| ------------------ | ----------------- |
| `No Known Exploit` | 영향 서브스코어가 크게 감소함. |
| `Proof of Concept` | 영향 서브스코어가 약간 감소함. |
| `Functional`       | 영향 서브스코어가 증가함.    |
| `High`             | 영향 서브스코어가 크게 증가함. |

### EPSS 점수

악용 예측 점수 시스템 (EPSS)은 FIRST 조직이 개발하고 소유한 모델에 기초하여 CVE가 야생에서 악용될 것인지를 예측합니다. 이 확률은 EPSS 모델의 직접적 결과이며, 야생에서 악용 위협의 전반적인 느낌을 전달합니다. 이 데이터는 최신 EPSS 모델 버전을 기반으로 매일 업데이트됩니다. 자세한 내용은 EPSS [문서](https://www.first.org/epss/articles/prob_percentile_bins)를 참조하십시오. **가능한 입력 값:** `EPSS 점수 [0.00-1.00]`

{% hint style="info" %}
노출 점수에 따라 가능성 서브스코어가 크게 증가합니다.
{% endhint %}

취약점이 악성으로 판명되면 EPSS 값은 `1`로 설정됩니다. EPSS에 관한 정보가 없는 경우 기본값은 `0.01055`입니다.

### 공격 벡터 (Attack vector)

취약점이 악용 가능한 맥락을 나타냅니다. [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-1-1-Attack-Vector-AV)를 기반으로 합니다.

| 가능한 입력 값   | 점수 영향                                              |
| ---------- | -------------------------------------------------- |
| `Network`  | 가능성 하위 점수가 증가합니다.                                  |
| `Adjacent` | 취약점을 익스플로잇하는 데 필요한 원격 액세스 수준에 따라 가능성 하위 점수가 감소합니다. |
| `Local`    | 취약점을 익스플로잇하는 데 필요한 원격 액세스 수준에 따라 가능성 하위 점수가 감소합니다. |
| `Physical` | 취약점을 익스플로잇하는 데 필요한 원격 액세스 수준에 따라 가능성 하위 점수가 감소합니다. |

### 공격 복잡도 (Attack complexity)

취약점을 악용하기 위해 존재해야 하는 조건의 복잡도를 나타내며, [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-1-2-Attack-Complexity-AC)를 기반으로 합니다.

| 가능한 입력 값 | 점수 영향             |
| -------- | ----------------- |
| `High`   | 가능성 하위 점수가 감소합니다. |
| `Low`    | 가능성 하위 점수가 증가합니다. |

### 필요한 권한 (Privileges required)

공격자가 취약점을 성공적으로 악용하기 전에 보유해야 하는 권한 수준을 나타내며, [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-1-3-Privileges-Required-PR)를 기반으로 합니다.

| 가능한 입력 값 | 점수 영향                            |
| -------- | -------------------------------- |
| `High`   | 요구되는 권한 수준에 따라 가능성 하위 점수가 감소합니다. |
| `Low`    | 요구되는 권한 수준에 따라 가능성 하위 점수가 감소합니다. |
| `None`   | 가능성 하위 점수가 증가합니다.                |

### 사용자 상호작용 (User interaction)

취약점 악용 과정에서 사용자의 행동이 필요한지 여부를 나타내며, [CVSS 정의](https://www.first.org/cvss/v3.1/specification-document#2-1-4-User-Interaction-UI)를 기반으로 합니다.

| 가능한 입력 값   | 점수 영향             |
| ---------- | ----------------- |
| `Required` | 가능성 하위 점수가 감소합니다. |
| `None`     | 가능성 하위 점수가 증가합니다. |

### 사회적 동향 (Social trends)

이 취약점에 대한 소셜 미디어 트래픽을 나타냅니다. Snyk 연구에 따르면, 소셜 미디어 상호작용이 증가하면 향후 악용될 가능성이 높거나 이미 악용되고 있을 가능성이 있습니다. 자세한 내용은 [사회적 동향이 있는 취약점](vulnerabilities-with-social-trends.md)을 참조하세요.

| 가능한 입력 값       | 점수 영향                   |
| -------------- | ----------------------- |
| `Trending`     | 가능성 하위 점수가 증가합니다.       |
| `Not trending` | 가능성 하위 점수에 영향을 주지 않습니다. |

### 악성 패키지 (Malicious package)

공급망 종속성으로 배포된 악성 코드는 높은 악용 가능성을 가집니다.

| 가능한 입력 값 | 점수 영향                           |
| -------- | ------------------------------- |
| `True`   | 악성 패키지에 대한 가능성 하위 점수가 크게 증가합니다. |
| `False`  | 가능성 하위 점수에 영향을 주지 않습니다.         |

### 취약점의 연령 (Age of vulnerability)

새로운 취약점(최대 1년 미만)은 오래된 취약점(출시 후 1년 이상)보다 악용될 가능성이 높습니다.

| 가능한 입력 값           | 점수 영향                                                            |
| ------------------ | ---------------------------------------------------------------- |
| 취약점이 처음 공개된 이후의 일수 | <p>1년 미만 - 가능성 하위 점수가 증가합니다.</p><p>1년 이상 - 가능성 하위 점수가 감소합니다.</p> |

### 패키지 인기 (Snyk Open Source)

생태계에서 상대적으로 더 인기 있는 패키지는 해커가 더 넓은 표적 그룹을 활용할 수 있기 때문에 악용될 가능성이 높습니다.

| 가능한 입력 값 | 점수 영향                   |
| -------- | ----------------------- |
| `High`   | 가능성 하위 점수가 증가합니다.       |
| `Medium` | 가능성 하위 점수에 영향을 주지 않습니다. |
| `Low`    | 가능성 하위 점수가 감소합니다.       |

### 공급자 긴급성 (Snyk Container)

해당 운영 체제 배포 보안 팀에서 제공하는 중요도 평가를 나타냅니다. 자세한 내용은 [상대적 중요도에 대한 외부 정보 소스](../../scan-with-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md#external-information-sources-for-relative-importance)를 참조하세요.

| 가능한 입력 값   | 점수 영향               |
| ---------- | ------------------- |
| `Critical` | 영향 하위 점수가 크게 증가합니다. |
| `High`     | 영향 하위 점수가 증가합니다.    |
| `Medium`   | 영향 하위 점수가 감소합니다.    |
| `Low`      | 영향 하위 점수가 크게 감소합니다. |

{% hint style="info" %}
CVSS 또는 중요도 평가가 제공되지 않는 경우, 공급자 긴급성은 기본적으로 `Low`로 설정됩니다. 공급자 긴급성은 또한 영향 하위 점수에 영향을 줍니다.
{% endhint %}

### 맥락적 가능성 위험 요소 (Contextual likelihood risk factors)

### 전이적 깊이 (Transitive depth)

[이전 연구](https://arxiv.org/pdf/2301.07972.pdf)를 기반으로 Snyk 연구에 따르면, 취약점이 직접적으로가 아니라 전이적으로 프로젝트에 도입된 경우, 악용 가능한 함수 경로가 존재할 가능성이 낮아집니다.

| 가능한 입력 값              | 점수 영향                   |
| --------------------- | ----------------------- |
| `Direct dependency`   | 가능성 하위 점수에 영향을 주지 않습니다. |
| `Indirect dependency` | 가능성 하위 점수가 감소합니다.       |

### 도달 가능성 (Reachability)

Snyk 정적 코드 분석은 취약한 메서드가 호출되는지를 결정합니다. 이는 Java 및 JavaScript에 대해 지원됩니다. 자세한 내용은 [도달 가능성 분석](reachability-analysis.md) 페이지를 참조하세요.\
도달 가능성이 활성화되지 않은 경우, 가능성 하위 점수는 변경되지 않으며, 해당 요소는 표시되지 않습니다.

| 가능한 입력 값        | 점수 영향                               |
| --------------- | ----------------------------------- |
| `Reachable`     | 가능성 하위 점수가 증가하며, 전이적 깊이는 고려되지 않습니다. |
| `No path found` | 가능성 하위 점수에 영향을 주지 않습니다.             |
