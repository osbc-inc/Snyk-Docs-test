# 롤아웃 통합 선택

## **SDLC 통합 지점**

Snyk은 SDLC의 각 단계에서 연속적으로 작업할 수 있도록 다양한 통합을 제공합니다.

많은 기업은 자동화된 솔루션을 먼저 도입한 후에, 점진적으로 개발자들이 도구를 활성화하는 단계로 전환합니다. 또한, 방화벽 기능은 시간이 지남에 따라 서서히 사용되도록 하여 중단을 최소화합니다.

{% hint style="info" %}
여러 통합을 사용하면 문제 보고의 중복이 발생할 수 있으므로 초기에는 하나의 통합 유형만 구현하는 것이 좋습니다. 예를 들어 Git 저장소로 모든 것을 가져온 후 나중에 세부 정보를 위해 CI/CD 뷰를 사용할 수 있습니다. 둘 다 필요하지 않다면 소스 컨트롤 통합을 제거할 수 있습니다.
{% endhint %}

## 통합 유형

다음은 일반적인 초기 통합입니다.

### 소스 코드 관리 (SCM) 통합

GitHub, GitLab, Azure Repos, Bitbucket과 같은 인기 있는 버전 관리 플랫폼과의 통합을 통해 Snyk 보안 확인을 코드 검토 과정에 원활하게 통합할 수 있습니다. 이를 통해 코드가 주 표시 분기로 병합되기 전에 잠재적인 취약점이 식별되고 해결됩니다. 중요한 기능은 다음과 같습니다:

* 지정된 분기(일반적으로 개발 분기)의 매일 테스트 및 모니터링
* (선택 사항) 요청/Merge 요청에서 저장소의 모든 분기에 대한 확인
* (선택 사항) 풀 리퀘스트를 통한 자동 종속성 업그레이드 및 자동 보안 수정 업그레이드

SCM 통합의 장점은 다음과 같습니다:

* 저장소 보안 상태에 대한 가시성
* 코드 변경 시 자동 검사
* 개발자에게 문제에 대한 즉각적인 피드백
* UI 또는 [API/API 가져오기 도구](https://docs.snyk.io/snyk-api-info/other-tools/tool-snyk-api-import)를 통한 저장소 구성 가능
* Snyk Enterprise 플랜에서 클라우드 및 개인 코드 저장소 지원

추가 세부 정보는 [Snyk SCM 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하십시오.

온프레미스 Git 저장소의 경우, Snyk이 저장소와 통신할 수 있도록 [Snyk Broker](https://docs.snyk.io/snyk-admin/snyk-broker)를 배포해야 합니다.

{% hint style="info" %}
기업 고객은 API를 사용하여 [Snyk Broker](../../../enterprise-setup/snyk-broker/)를 활성화하고 관리할 수 있습니다.

Snyk Broker 배포를 지원하기 위해 [유료 서비스](../../../working-with-snyk/snyk-terms-of-support-and-services-glossary/)를 활용할 수 있습니다.
{% endhint %}

### 지속적인 통합/지속적인 배포 (CI/CD) 파이프라인 통합

Jenkins, Travis CI, 또는 CircleCI와 같은 CI/CD 파이프라인에 Snyk를 통합하면 빌드 및 배포 과정 중 보안 확인이 자동화됩니다. 이를 통해 취약점이 소프트웨어 개발 수명주기 초기에 발견되고, 제품으로 전파되는 것을 방지합니다. 일반적인 기능은 다음과 같습니다:

* (선택 사항) 빌드 중 결과를 수동으로 모니터링하고 Snyk에서 결과를 확인할 수 있음
* (선택 사항) 지정한 기준에 따라 결과를 검사하고 빌드를 중단할 수 있음

CI/CD 통합의 장점은 다음과 같습니다:

* 로컬 코드 취약점 평가
* 실행할 테스트 및 빌드 스크립트에 대한 위치에 대한 완전한 제어
* 원하는 경우 CI/CD에 의한 자동화

추가 세부 정보는 [Snyk CI/CD 통합](../../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)을 참조하십시오.

### IDE 통합

Visual Studio Code, IntelliJ IDEA, Eclipse와 같은 통합 개발 환경 (IDE)을 사용하면 개발자가 코딩 환경에서 바로 Snyk 보안 기능에 액세스할 수 있습니다. 이를 통해 개발자가 코드를 작성하는 동안 실시간 스캔과 문제 해결이 가능해집니다.

추가 세부 정보는 [Snyk IDE 플러그인 및 확장 프로그램](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)을 참조하십시오.

## 가져오기 전략을 위한 고려 사항

<table><thead><tr><th width="200">프로젝트 가져오기 전략</th><th>고려 사항</th><th>장점</th><th>단점</th></tr></thead><tbody><tr><td>CLI (CI/CD로 자동화)</td><td>CI/CD 내에서 각 애플리케이션에 대해 구성해야 함</td><td><ul><li>테스트 및 실행할 때 선택 사항: 어떤 패키지 관리자, 프로세스 내 어디에 실행할 것, 어떤 언어를 분석할 것</li><li>통합을 위해 개발 노력이 필요할 수 있음</li></ul></td><td>애플리케이션당 구성이 필요함.</td></tr><tr><td>CLI (사용자가 로컬에서 실행)</td><td>개발자가 애플리케이션을 작업하는 동안 로컬에서 테스트 수행에 사용될 수 있음, 매우 다양한 스캔 유형에 대해 매우 구성 가능함.</td><td>로컬 사용 사례</td><td>가시성 또는 자동화를 위해 적합하지 않음. 예를 들어, Gradle 잠금 파일 없이 빌드 가능한 코드나 종속성이 설치된 코드가 필요할 수 있음</td></tr><tr><td>API</td><td><ul><li>일반적으로 고급 사용 사례를 위해 사용됨.</li><li>CI/CD 워크플로 또는 사용자 지정 도구에 통합.</li></ul></td><td>CI/CD 파이프라인으로의 자동 통합</td><td>API 익숙함 필요, Enterprise Plan을 통해 접근 가능함.</td></tr><tr><td>Git 코드 저장소 통합</td><td>온보딩 및 일일 모니터링을 위해 사용됨: 애플리케이션 포트폴리오 전체에 걸친 빠른 취약점 평가</td><td><ul><li>작업 중이 아닌 경우에도 저장소의 지속적 모니터링</li><li>팀을 위한 중앙 집중형 가시성</li><li>지정한 분기 모니터링</li><li>코드 빌드가 필요하지 않음</li></ul></td><td><ul><li>UI를 통해 시작할 수 있지만, 대규모 포트폴리오의 경우 [API 가져오기 도구](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-api-import)를 사용하여 저장소를 시작할 수 있음</li><li>어떤 언어/패키지 관리자는 CLI 사용을 통해 더 나은 해결이 가능함: Gradle 잠금 파일 없이 빌드 가능한 코드, Scala</li></ul></td></tr><tr><td></td><td><ul><li>풀 리퀘스트(PR)/머지 리퀘스트(MR) 스캔</li></ul></td><td><ul><li>저장소의 어떤 분기에서든 PR/MR에 도입된 문제에 대한 즉각적인 피드백</li></ul></td><td>클수 있음단일 규칙 구성</td></tr></tbody></table>

## 추가 통합에 대한 고려 사항

### Infrastructure as Code 통합

Snyk Infrastructure as Code에 대해, Terraform 또는 YAML 구성 파일이 SCM에 저장될 수 있지만 별도의 영역이나 저장소에 있을 수 있습니다. 따라서 가져오기가 필요한 다른 영역이 있는지 고려해야 합니다. 해당하는 경우 Terraform Cloud와 통합할지, `terraform run` 프로세스의 일부로 Snyk 테스트를 가능하게 하기 위해 통합할지 고려할 수도 있습니다.

복잡한 환경, 모듈, 그리고 매우 템플릿화된 실행은 Terraform Plan 파일에 CLI를 사용하는 것이 최상의 결과를 제공할 수 있습니다.

### 컨테이너 레지스트리(CR) 통합

[컨테이너 레지스트리](../../../scan-with-snyk/snyk-container/container-registry-integrations/)와 Snyk도 통합하여 취약점을 모니터링하고 가져온 컨테이너를 정기적으로 확인할 수 있습니다. Snyk은 가져온 컨테이너를 테스트하여 발견된 알려진 보안 취약점이 있는지 확인하며, 사용자가 제어하는 빈도로 진행됩니다.

### Kubernetes

Snyk을 구성하여 Kubernetes에 배포된 작업 부하를 모니터링할 수 있습니다. 컨트롤러를 구성하는 방법에 대한 자세한 내용은 [Kubernetes 통합 개요](../../../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/)에서 확인하십시오.
