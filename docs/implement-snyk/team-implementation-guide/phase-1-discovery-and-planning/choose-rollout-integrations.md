# 롤아웃 통합 선택

## **SDLC 통합 포인트**

Snyk는 SDLC 각 단계에 완벽하게 통합되는 많은 통합을 제공합니다.

많은 기업은 일반적으로 자동화된 솔루션을 먼저 도입한 후 점진적으로 개발자를 지원할 도구를 도입합니다. 또한 접근 기능은 방해를 최소화하기 위해 시간을 두고 점진적으로 활성화됩니다.

{% hint style="info" %}
여러 통합을 사용하는 것은 문제의 중복 보고로 이어질 수 있으므로 초기에는 두 가지 이상의 통합 유형을 도입할 필요가 없습니다. 예를 들어 Git 저장소로 모든 것을 가져온 후 CI/CD 뷰를 사용하여 미세한 상세 정보를 얻을 수 있습니다 (양쪽 뷰를 원하지 않는 경우 원본 제어 통합을 제거할 수 있음).
{% endhint %}

## 통합 유형

다음은 일반적인 초기 통합입니다.

### 소스 코드 관리 (SCM) 통합

GitHub, GitLab, Azure Repos 및 Bitbucket과 같은 인기있는 버전 관리 플랫폼과의 통합을 통해 Snyk 보안 검사를 코드 검토 프로세스에 완벽하게 통합합니다. 이를 통해 코드가 본 브랜치로 병합되기 전에 잠재적인 취약성이 식별되고 해결됩니다. 중요한 기능은 다음과 같습니다:

* 특정 브랜치 (일반적으로 "개발" 브랜치)의 매일 테스트/모니터링,
* (옵션) 어떤 브랜치에 대한 Pull Request/Merge Request 확인.
* (옵션) 풀 리퀘스트를 사용한 자동 종속성 업그레이드 및 자동 보안 수정 업그레이드.

장점:

* 저장소 보안 포지션의 가시성
* 코드 변경 시 자동 스캔
* 개발자에 대한 문제에 대한 즉각적인 피드백
* UI를 사용하여 저장소 추가 설정
* 팀 플랜의 클라우드 저장소 지원

자세한 내용은 [Git 저장소 (SCM)](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)를 참조하십시오.

{% hint style="info" %}
클라우드 미노출 및 자체 git SCM 인스턴스가있는 경우:

* Snyk이 리포지토리와 통신하도록 [Snyk Broker](https://docs.snyk.io/snyk-admin/snyk-broker)를 배포하는 것을 고려하십시오. 이는 Snyk 엔터프라이즈 플랜이 필요합니다.
* 엔터프라이즈 고객은 API를 통해 [Snyk Broker](https://docs.snyk.io/enterprise-setup/snyk-broker)를 활성화하고 관리할 수 있습니다.

브로커 배포를 지원하는 [유료 서비스](https://docs.snyk.io/more-info/snyk-terms-of-support-and-services-glossary)를 이용할 수 있습니다.
{% endhint %}

### 지속적 통합/지속적 배포 (CI/CD) 파이프라인 통합

Jenkins, Travis CI 또는 CircleCI와 같은 CI/CD 파이프라인에 Snyk를 통합함으로써 빌드 및 배포 프로세스 중 보안 검사를 자동화합니다. 이를 통해 소프트웨어 개발 라이프사이클 초기에 취약성이 조기에 감지되어 프로덕션으로 전파되는 것을 방지합니다. 일반적인 기능은 다음과 같습니다:

* (옵션) 빌드 중 결과를 수동으로 모니터링하고 Snyk에서 결과 확인 기능
* (옵션) 지정한 기준에 따라 테스트하고 빌드를 실패로 만들 수 있는 능력
* 특정 Marketplace 플러그인을 통해 또는 일반적으로 CLI와 함께 파이프라인 스크립트의 일부로 통합

장점:

* 로컬 코드 취약성 평가
* 테스트 제어권 (어떤 테스트를 실행할지, 빌드 스크립트의 어디에서 실행할지)
* CI/CD를 통한 자동화

자세한 내용은 [Snyk CI/CD 통합](../../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)을 참조하십시오.

### 통합 개발 환경 (IDE) 통합

Visual Studio Code, IntelliJ IDEA, Eclipse와 같은 IDE 통합을 통해 개발자는 코딩 환경에서 직접 Snyk 보안 기능에 액세스할 수 있습니다. 이를 통해 개발자가 코드를 작성하는 가장 초기 단계에서 실시간 스캔 및 문제 해결이 가능해집니다.

자세한 내용은 [IDE에서 Snyk 사용](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)을 참조하십시오.

## 가져오기 전략을 위한 고려사항

<table><thead><tr><th width="200">프로젝트 가져오기 전략</th><th>고려 사항</th><th>장점</th><th>단점</th></tr></thead><tbody><tr><td>CLI (자동화된 CI/CD)</td><td>CI/CD 내 각 애플리케이션에 대해 구성해야 합니다.</td><td><ul><li>무엇을 테스트할지 및 언제 실행할지 선택할 수 있음 (즉, 어떤 패키지 관리자, 프로세스의 어디에, 어떤 언어를 분석할 지).</li><li>통합을 위해 개발 작업이 필요할 수 있음.</li></ul></td><td>각 애플리케이션에 대한 구성이 필요합니다.</td></tr><tr><td>CLI (사용자가 로컬에서 실행)</td><td>사용자는 CLI를 사용하여 애플리케이션 작업 중 로컬에서 테스트 수행할 수 있으며 각 스캔 유형에 대해 매우 구성 가능합니다.</td><td>로컬 사용 사례</td><td>가시성이나 자동화를 위한 것이 아님. 빌드 가능한 코드나 종속성 설치가 필요할 수 있습니다 (예: Lockfile이 없는 Gradle, Scala).</td></tr><tr><td>Git 코드 저장소 통합</td><td>입사 및 일일 모니터링: 애플리케이션 포트폴리오 전체에서 빠른 취약성 평가.</td><td><ul><li>저장소의 지속적인 모니터링 (작업 중이 아닐 때도).</li><li>팀을 위한 중앙화된 가시성.</li><li>지정한 브랜치 모니터링</li><li>코드를 빌드할 필요가 없음</li></ul></td><td><ul><li>UI를 통해 시작할 수 있음</li><li>일부 언어/패키지 관리자는 CLI를 사용하여 해상도를 개선할 수 있음 (Lockfile이 없는 Gradle, Scala).</li></ul></td></tr><tr><td></td><td>Pull Request(PR)/Merge Request(MR) 스캔</td><td>리포지토리의 어떤 브랜치에 대한 PR/MR에서 소개된 문제에 대한 즉각적인 피드백 제공.</td><td>패스/페일을 위한 구성 가능한 규칙</td></tr></tbody></table>

## 추가 고려 사항

### **Infrastructure as Code**

Snyk의 경우, Terraform이나 yaml 구성 파일이 SCM에 보관되지만 별도의 영역이나 저장소에 있을 수 있습니다. 결과적으로 가져와야 할 다른 영역이 있는지 고려하십시오. "Terraform run" 프로세스 중에 Snyk 테스트를 가능하게 하기 위해 Terraform Cloud와 통합하려고 할 수도 있습니다.

복잡한 환경, 모듈 및 매우 템플릿화된 구현의 경우, Terraform Plan 파일에서 CLI를 활용하는 것이 최상의 결과를 제공할 수 있습니다.

### CR (컨테이너 레지스트리)

Snyk은 다양한 [컨테이너 레지스트리](../../../scan-with-snyk/snyk-container/container-registry-integrations/)와도 통합하여 취약성을 모니터링할 수 있게 하며 가져온 컨테이너에 대한 알려진 보안 취약성을 빈도를 제어하여 테스트합니다.

###
