# 단계 4: 수정 전략 만들기

통합 설정을 완료하고 조직을 만들고 프로젝트를 가져온 후, 비즈니스의 현재 취약점 백로그를 볼 수 있습니다.

## 중요한 집중 영역 결정

구체적인 취약성에 뛰어들기 전에, 조직 및 저장소("Snyk"에서의 "타겟")를 고려하여 특정 중요 영역을 고려해보세요. 단계 1에서 수행한 비즈니스 크리티컬 애플리케이션의 발견은이 단계에서 우선순위를 결정하고 지원할 수 있습니다. 예를 들어, 조직이 다른 제품과 일치하는 경우, 초기에는 사용자가 가장 많은 제품이나 보안이 가장 중요한 제품에 초점을 맞출 수 있습니다.

조직 내에서 애플리케이션을 구성하는 저장소를 고려할 수 있습니다. 민감한 데이터를 처리하는 영역이나 공개 기능을 하는 영역이 더 중요할 수 있으므로, 이것은 검토할 프로젝트 목록을 좁히는 또 다른 방법입니다.

{% hint style="info" %}
**속성**을 사용하여 프로젝트에 메타데이터를 추가한 경우, 고려 중인 프로젝트의 수를 줄일 수 있는 좋은 방법입니다.
{% endhint %}

## 개발팀별 그룹 작업

우선순위를 매기기 위해 줄인 프로젝트 집합을 갖고 나면, 이를 서로 다른 개발팀 사이에 분배하고 싶을 것입니다. 예를 들어, 오픈 소스와 1차 코드를 수정하는 하나의 개발팀과 컨테이너 및 베이스 이미지 취약점을 담당하는 별도의 데브옵스 팀이 있을 수 있습니다.

## 우선순위 설정 방법

보다 급한 문제를 먼저 수정해야할 필요성을 판단하는 데 도움이 되는 필터가 있습니다. 우선순위 계획을 수립할 때 가장 일반적으로 사용되는 다음 검색 기준은 분석 결과를 분석할 때 반복적으로 사용하거나 혹은 결합하여 사용할 수 있습니다.

* 심각도 (먼저 **높음**과 **중요**부터 시작). 심각한 심각도로 필터링하는 것이 일반적입니다. 하지만 를 사용 중이라면, Snyk 코드는 높은 단계까지만 가기 때문에 결과를 분석할 때 High 단계에서 시작하세요.
* [Exploit Maturity](https://snyk.io/blog/whats-so-wild-about-exploits-in-the-wild-and-how-can-we-prioritize-accordingly/) (바로 사용 가능한 **Mature** 또는 **Proof of Concept** 문제가 더 취약함). 이 필터를 선택하면 결과를 자동으로 오픈 소스로 필터링합니다.
* 수정 가능 여부 (패키지를 업그레이드하면 쉽게 수정할 수 있는 경우, 수정이 더 빨라집니다). 
* 오픈 소스 취약점의 CVSS 점수
* [우선 순위 점수](https://docs.snyk.io/manage-risk/priorities-for-fixing-issues/priority-score) (우리는 이 점수를 계산하는 데 위의 값들을 사용합니다). 하나의 전략은 점수가 900-1000인 취약성을 제거한 다음, 800-900 점수의 취약성으로 이동하는 것입니다.

수리 전략을 계획할 때 어떤 메트릭을 사용할지 결정하고, 타임라인을 구체적으로 설정하세요. 예를 들어, 심각도로 수정하는 경우, 각 심각도 당 취약성을 해결하는 데 걸리는 시간을 추정하세요. 수정 전략을 구체적으로 설정하는 것이 권장됩니다.

**예시**

만약 50개의 중대한 심각성 문제와 100개의 높은 심각성 문제가 있다면, 팀 규모와 업무 부담을 고려하여 중대한 취약점을 수정하는 데 2주, 그 후 높은 심각성을 수정하는 데 4주 정도가 걸릴 것입니다.

또 다른 방법으로 문제 유형별로 수정할 수 있습니다.

## 문제 유형별 수정

특정 문제 유형 (예: 오픈 소스 취약점)에 초점을 맞추는 것이 일반적입니다. 이렇게 하면 다른 프로젝트 간에 문제를 더 쉽게 비교할 수 있습니다.

다른 유형의 팀에 의해 이끌어지는 수정 작업에 기반한 프로세스의 예시입니다.

### **예시: 개발자 중심 우선순위**

라이선스 위험을 최소화하기 위해 임원으로부터 지침을 받은 개발자 중심 구현:
1. Snyk Open Source에서 식별된 모든 라이선스 문제 해결.
2. Snyk Open Source에서 수정 가능한 중요하거나 높은 취약점 다루기.
3. 높은 심각성 문제부터 시작하여 Snyk Code를 사용하는 코드 분석 프로젝트에 집중.
4. 애플리케이션/환경을 실행하는 데 사용하는 컨테이너 및 IaC 파일을 스캔.

### **예시: DevSecOps 중심 우선순위**

사용자 정의 이미지와 환경을 보호하는 데 중점을 둔 DevSecOps 중심 구현:
1. 개발팀이 다운로드하는 사용자 정의 베이스 이미지를 스캔하고 보호.
2. 컨테이너 레지스트리와 통합하고, 개발팀에 제공하는 이미지를 스캔.
   1. 내부 툴을 추가하고 이미지 매개변수를 표준화한 후 선택 및 작성한 이미지를 스캔.
   2. 개발자가 컨테이너를 배포하기 전에 자신의 사용자 정의 도구/패키지를 추가할 때 스캔하여 환경이 보안 상태를 유지하도록 하세요. 이 스캔은 또한 애플리케이션 취약점을 검출할 것입니다.
3. 컨테이너를 보호한 후, 보안 위반으로 이어질 수 있는 미구성 구성 파일과 클라우드 환경을 적극적으로 스캔하십시오.
4. 개발팀에 Snyk를 소개하여 오픈 소스 및 코드를 스캔하여 애플리케이션 취약점을 줄이세요.

## 대상 취약점 캠페인

개발 프로세스에서 보안 테스트를 운영화하는 동안, 취약성 유형을 제거하기 위한 수정 전략 중 하나로 캠페인을 진행할 수 있습니다. 예를 들어 SQL 삽입처럼 취약성 유형을 제거할 수 있습니다. 검색 필터에서 CWE를 사용하면 문제를 식별하고 기록하는 데 매우 유용할 수 있습니다.

## 타임라인 업데이트

수정 전략을 작성한 후, 7단계의 타임라인을 업데이트하세요.