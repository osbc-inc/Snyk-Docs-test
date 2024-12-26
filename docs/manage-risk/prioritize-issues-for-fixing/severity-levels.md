# 심각성 수준

[vulnerability assessment](https://snyk.io/learn/vulnerability-assessment/)를 수행하기 위해 심각성 수준을 사용하십시오. 심각성 수준은 **C**ritical(심각), **H**igh(높음), **M**edium(중간), **L**ow(낮음)로 평가된 위험 수준을 나타냅니다. Snyk는 많은 곳에서 심각성 수준의 취약점 수를 보고합니다. 디스플레이는 다양하며 전형적인 예제는 아래와 같습니다.

<img src="../../.gitbook/assets/Screenshot 2022-08-16 at 09.52.22.png" alt="Issues at each level of severity, C, H, M, and L" data-size="original">

{% hint style="info" %}
심각성 수준은 라이선스 문제에도 적용됩니다. [Licenses](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md)를 참조하세요.
{% endhint %}

다음 표에 심각성 수준이 정의되어 있습니다.

| 아이콘                                                                                                        | 수준         | 설명                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------- | ------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| <img src="../../.gitbook/assets/image (131) (1) (1) (1) (1) (1).png" alt="C" data-size="line">                | **C**ritical | 공격자가 민감한 데이터에 액세스하고 응용 프로그램에서 코드를 실행할 수 있을 수도 있습니다.                                                              |
| <img src="../../.gitbook/assets/image (103) (1) (1) (1) (1) (1) (1) (2) (1).png" alt="H" data-size="original"> | **High**     | 공격자가 응용 프로그램에서 민감한 데이터에 액세스할 수 있을 수 있습니다.                                                                           |
| ![M](<../../.gitbook/assets/image (133) (1).png>)                                                             | **M**edium   | 특정 조건에서 공격자가 응용 프로그램에서 민감한 데이터에 액세스할 수 있을 수도 있습니다.                                                    |
| ![L](<../../.gitbook/assets/image (422).png>)                                                                 | **L**ow      | 응용 프로그램이 다른 취약성으로 공격자에게 응용 프로그램을 공격하는 데 사용할 수 있는 데이터를 노출할 수 있습니다. |

## 심각성 수준과 우선 순위 점수

심각성 수준은 각 취약점의 Snyk 우선 순위 점수를 결정하는 데 사용되는 요소 중 하나입니다. 다른 요소에는 [Snyk Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) 및 [Reachable Vulnerabilities](https://snyk.io/blog/optimizing-prioritization-with-deep-application-level-context/) 정보가 포함됩니다.

자세한 내용은 [Snyk 우선 순위 점수](priority-score.md)를 참조하십시오.

## 심각성 수준 보기 방법

심각성 수준은 항상 이 정보를 보기 쉽게 유지하기 위해 Snyk에서 표시됩니다.

예를 들어, 심각성 수준은 대시보드의 **보류 중인 작업** 섹션에 표시됩니다:

<img src="../../.gitbook/assets/image (158) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Severity levels with Pending tasks" data-size="original">

심각성 수준은 [Snyk Projects](../../snyk-admin/snyk-projects/)와 연계하여 표시됩니다:

<figure><img src="../../.gitbook/assets/image (43) (2).png" alt="Severity levels assoicated with Projects"><figcaption><p>프로젝트와 연계된 심각성 수준</p></figcaption></figure>

각 심각성 수준의 문제 수는 이슈 카드의 왼쪽 사이드바에도 표시됩니다:

<figure><img src="../../.gitbook/assets/image (39) (1) (1).png" alt="Issue card; severity levels in sidebar"><figcaption><p>이슈 카드; 사이드바에 표시된 심각성 수준</p></figcaption></figure>

## Snyk가 심각성 수준을 결정하는 방법

### 심각성 수준 및 CVSS

공통 취약성 점수 평가 시스템 (CVSS)이 취약성의 심각성 수준을 결정합니다.

Snyk는 취약점의 특성과 심각성을 지정하기 위해 [CVSS 프레임워크 버전 4.0](https://www.first.org/cvss/v4-0/)를 지원하며, 이전 버전인 [CVSS 프레임워크 버전 3.1](https://www.first.org/cvss/v3-1/)도 지원하여 취약성의 특성과 심각성을 지정합니다.

&#x20;CVSS v4.0의 지원 이전에 발표된 취약점은 **3.1 버전의 CVSS를 기반으로 심각성을 정의**합니다.

| **수준** | **CVSS 점수** |
| --------- | -------------- |
| Critical  | 9.0 - 10.0     |
| High      | 7.0 - 8.9      |
| Medium    | 4.0 - 6.9      |
| Low       | 0.0 - 3.9      |

심각성 수준 및 점수는 [CVSS 베이스 점수](https://www.first.org/cvss/specification-document) 계산을 사용하여 베이스 메트릭을 기반으로 결정됩니다. 임시 점수는 임시 메트릭을 기반으로 결정됩니다.

자세한 내용은 [CVE에 대한 CVSS 소개: 보안 취약점 평가 점수 매기기](https://snyk.io/blog/scoring-security-vulnerabilities-101-introducing-cvss-for-cve/)를 참조하십시오.

{% hint style="info" %}
심각성 수준은 항상 CVSS 점수와 일치하지 않을 수 있습니다. 예를 들어, 의 Linux 취약점의 심각성 점수는 NVD 심각성 순위에 따라 다를 수 있으니 [Linux 취약성 심각성 이해](../../scan-with-snyk/snyk-container/how-snyk-container-works/severity-levels-of-detected-linux-vulnerabilities.md)를 참조하십시오.
{% endhint %}

### **같은 취약성에 대한 여러 CVSS 점수가 있는 이유**

같은 취약성에 대해 여러 CVSS 점수가 있는 이유는 다음과 같습니다:

* 취약성의 심각성을 평가할 때 한 가지 CVSS 벡터가 없다는 점을 기억해야 합니다. 여러 벤더에 의해 정의된 여러 CVSS 벡터가 있으며, [National Vulnerability Database](https://nvd.nist.gov/) (NVD) 등이 있습니다.
* Snyk가 발표하는 대부분의 취약성은 [소유자에게 의해 발견됩니다](https://security.snyk.io/disclosed-vulnerabilities), 공개 정보 소스 또는 제3자 통지를 통해 얻어집니다.\
  예를 들어, Snyk가 심각도가 높은 [Spring4Shell 취약성](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORK-2436751)을 발견했을 때, CVSS 벡터 분석이 포함된 공지가 2022년 3월 30일에 게시되었습니다. 이는 공식 CVE가 할당되기 전이며, NVD가 필터링을 한 후 2022년 4월 8일에 게시한 것입니다.
* CVSS 벡터에 약간의 차이가 있을 경우가 있으며, 이는 일반적이고 예상되는 것입니다. 특정 공격 벡터의 가능성에 대한 일부 불일치와 해당 애플리케이션 및 오픈 소스 소프트웨어 사용자의 사용 사례에 대한 판단이 포함됩니다.
* 취약성의 심각성은 "레드 팀" 각도 또는 "블루 팀" 각도에서 비롯할 수 있는 다양한 요소에 의해 영향을 받습니다. 객관적이고 실행 가능한 등급을 도출하기 위해 Snyk 분석가들은 벤더, 제보자, 공격자까지 모든 관련 데이터 범위를 검토합니다.
* 벤더가 취약성에 대한 추가 정보를 발견하는 경우 해당 취약성의 심각성에 영향을 미칠 수 있습니다. 사용자는 Snyk에서 정리한 정보와 참고를 통해 심각성을 결정하는 데 사용된 모든 관련 정보를 찾을 수 있습니다.

### 심각성 수준 및 CCSS

국가표준기술원(NIST)에서 개발한 공통 구성 점수 시스템 (CCSS)은 소프트웨어 보안 구성 문제의 심각성을 측정합니다.

Snyk는 IaC+ 취약점과 잘못된 구성에 대한 특성과 심각성을 지정하기 위해 [CCSS](https://www.nist.gov/publications/common-configuration-scoring-system-ccss-metrics-software-security-configuration)를 사용합니다.

| **수준** | **CCSS 점수** |
| --------- | -------------- |
| Critical  | 9.0 - 10.0     |
| High      | 7.0 - 8.9      |
| Medium    | 4.0 - 6.9      |
| Low       | 0.0 - 3.9      |

심각성 수준 및 점수는 베이스 메트릭을 사용하여 CCSS 베이스 점수 계산을 통해 결정됩니다.