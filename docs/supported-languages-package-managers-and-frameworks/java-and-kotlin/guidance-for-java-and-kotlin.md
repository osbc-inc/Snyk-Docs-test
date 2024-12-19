# 자바 및 코틀린 안내

이 안내서는 기술 스택에서 Snyk를 효과적으로 적용하는 데 도움이 되도록 제공됩니다.

## 패키지 레지스트리 통합 (Artifactory/Nexus) - Maven

Artifactory 및 Nexus 패키지 레지스트리 통합은 Snyk 엔터프라이즈 플랜 사용자에게 제공됩니다.

- Snyk 오픈 소스는 프라이빗 패키지를 통해 트랜지티브 종속성을 해결하기 위해 Artifactory나 Nexus를 사용합니다.
- Snyk는 Snyk 브로커를 사용하여 네트워크의 개인 서버나 사용자 이름과 비밀번호를 사용하여 공개적으로 이용 가능한 인스턴스에 연결할 수 있습니다.
- Snyk 오픈 소스는 보안 테스트를 위해 레지스트리와 상호 작용하는 Artifactory 및 Nexus와의 통합을 제공합니다. [Nexus Repository Manager 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) 및 [Artifactory 레지스트리 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/) 참조

{% hint style="info" %}
Snyk 엔터프라이즈 사용자가 아닌 경우 Artifactory나 Nexus를 사용하는 경우 CLI를 사용하여 분석하는 것이 가장 좋습니다. 빌드 시스템이 종속성을 검색하고 로컬에 존재할 것입니다.
{% endhint %}

Maven을 포함한 패키지 레지스트리 통합에 대한 자세한 내용은 다음을 참조하십시오.

- 패키지 레지스트리 통합: [Nexus Repository Manager 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) 및 [Artifactory 레지스트리 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)
- [Maven을 위한 Artifactory 레지스트리](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/artifactory-registry-for-maven.md)
- [Maven을 위한 Nexus 레지스트리](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/nexus-repository-manager-for-maven.md)
- Nexus 컨테이너 레지스트리: [Nexus 통합을 통한 컨테이너 보안](../../scan-with-snyk/snyk-container/container-registry-integrations/integrate-with-nexus-container-registry.md)
- 게이트키퍼 플러그인: [Artifactory 게이트키퍼 플러그인](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/artifactory-gatekeeper-plugin.md) 및 [Nexus Repository Manager 게이트키퍼 플러그인](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

## 언어 및 패키지 매니저 고려 사항

오픈 소스를 사용하는 경우 Maven이나 Gradle을 사용하려 한다면 Snyk를 설치하는 가장 좋은 방법에 영향을 줄 수 있습니다.

- **Maven 또는 gradle.lockfile을 사용하는 경우**:
Git 코드 저장소 통합은 Snyk를 사용하고 가시성을 얻는 좋은 방법이거나 CLI/IDE 또는 CI/CD 통합을 사용하여 테스트/게이트/모니터링을 할지 결정할 수 있습니다!
- **gradle.lockfile 없이 Gradle을 사용하는 경우**:
전체 종속성 트리가 명확하지 않을 수 있거나 아티팩트가 외부 리소스로부터 가져올 수 있기 때문에 로컬 스캔에 CLI/IDE 워크플로우 및 CI/CD가 분석하는 추천된 접근 방식입니다.

### Maven

Snyk는 Git 통합 또는 CLI를 통해 POM에서 종속성 트리를 생성할 수 있습니다.

- 지역적 및 CI/CD 사용: Snyk는 패키지 매니저와 상호 작용하여 종속성 목록을 생성합니다.
- Git 통합: 일반적으로 그 시간에 빌드된 것처럼 빌드를 근사적으로 실행합니다.

{% hint style="info" %}
개발자 종속성 (`scope=test`)은 제품에 푸시되지 않으며 일반적으로 노이즈로 간주되므로 무시됩니다. `--dev`를 추가하여 CLI에서 활성화할 수 있습니다.
{% endhint %}

### Gradle

- Snyk는 패키지 매니저와 상호 작용하여 종속성 목록을 생성할 것입니다.
- 일반적으로 Gradle은 빌드 프로세스 중에 코드 및 다른 작업을 실행할 수 있으므로 gradle.lockfile이 없는 경우 CLI 워크플로우가 권장됩니다.

### Kotlin

다음 manifest 파일이 지원됩니다.

- SCM 및 CLI를 위한 Groovy DSL인 build.gradle
- CLI 전용 Kotlin DSL인 build.gradle.kts

지원되는 기능에 대한 자세한 내용은 [Java 및 Kotlin](./#open-source-and-licensing) 페이지를 참조하십시오.

### **APIs**

고객이 고급 종속성 관리 전략을 개발하고 일반적으로 사용되는 표준 패키지 매니저를 사용하지 않을 수 있습니다. 이에 따라 Snyk는 테스트 API를 제공했습니다.

Snyk API를 사용하여 즉시 테스트하려면 [Test](../../snyk-api/reference/test-v1.md) 엔드포인트를 사용하십시오. [Maven) 공개 패키지의 문제에 대한 테스트](../../snyk-api/reference/test-v1.md#test-maven-groupid-artifactid-version) 및 [패키지에 대한 문제 목록](../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues)와 같은 예시가 있습니다.

## Snyk 통합 및 일반적으로 사용되는 패턴

Java는 개발자에게 강력한 유연성과 설정 옵션을 제공합니다. 특히 오픈 소스와 관련하여 테스트하는 데 고려할 사항이 있을 수 있습니다.

### gradle.lockfile을 사용하는 Maven 및 Gradle 프로젝트

일반적으로 빌드 시스템의 일부로 테스트를 설정하거나 프로세스의 일부로 lockfile을 채택하는 것이 가능합니다.

- 대규모 조직에서 주요 응용 프로그램에 대한 시작일부로 Git 통합을 통해 응용 프로그램을 모니터링하는 것은 꽤 일반적입니다.
- 개발자가 Snyk 기능에 익숙해지면 게이트 지원을 위해 PR 확인을 넓힙니다.
- CI/CD를 사용하여 수동으로 모니
기능별로 모니터링한 후 CI/CD 통합에서 snyk [제품] 테스트 및 모니터 명령을 사용하여 게이팅을 실행합니다.

### lock file이 없는 Gradle 프로젝트

- CI/CD를 사용하여 수동으로 모니터링한 후 CI/CD 통합에서 `snyk [제품] 테스트` 및 `monitor` 명령을 사용하여 게이팅을 실행합니다.
- 일반적으로 빌드를 실패시키는 게이팅은 시작할 때 하나의 프로젝트에 대해 켜져 있으며 모두가 프로세스를 익힘과 동시에 나머지 포트폴리오에 대해 수동 모니터링을 사용합니다.

## Snyk CLI의 팁과 트릭

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)

### 무엇을 테스트할 지

CLI에서 `--help` 옵션을 사용하여 Snyk CLI 명령의 자세한 내용을 확인할 수 있습니다.

:link: [CLI 명령 및 옵션 요약](../../snyk-cli/cli-commands-and-options-summary.md)

#### 자사 코드 테스트:

- 프레임워크 지원 - [Snyk Code - 지원되는 언어 및 프레임워크](../).
- 프로젝트 루트에서 `snyk code test` 명령을 사용하여 소스 코드 분석을 수행할 수 있습니다.
- `--scan-all-unmanaged --all-projects`를 사용하여 현재 작업 디렉토리 아래의 모든 jar 파일을 재귀적으로 찾을 수 있습니다.

#### 오픈 소스 라이브러리

**Maven**

`snyk test` 명령은 발견한 첫 manifest를 테스트하고 해당 개별 entry point를 스캔합니다. 모든 manifest를 스캔하려면 다음 지침을 따르십시오.

- 프로젝트를 집계하여 스캔하려면 `--maven-aggregate-project` 옵션을 사용합니다 (예: `snyk test --maven-aggregate-project`).
- 모든 프로젝트를 스캔하려면 `--all-projects` 옵션을 사용합니다: (즉, `snyk test --all-projects`)

Snyk는 기본적으로 활성화된 활성 프로필을 스캔합니다.

- 추가 Maven 인수를 전달할 수 있으며, 일반적으로 표준이 아닌 settings.xml 위치가 있습니다. `snyk test -- -s path/to/settings.xml`과 같은 방식으로 특정 설정을 스캔할 수 있습니다.
- 특정 Maven 프로필을 테스트하려면 `-P [이름]`을 사용하여 이름이 지정된 특정 Maven 프로필을 스캔할 수 있습니다. 예를 들어 **prod** 구성을 스캔하려면 `snyk test -- -P prod`를 사용합니다.

**Gradle**

기본적으로 Snyk CLI는 현재 폴더의 루트 프로젝트 또는 `--file=path/to/build.gradle`로 지정된 프로젝트만 스캔합니다.

- `--all-sub-projects`를 사용하여 한 번에 모든 프로젝트를 스캔하려면 추천됩니다 (즉, `snyk test --all-sub-projects`). 각각의 개별 sub-프로젝트는 웹 UI에서 개별적인 Snyk 프로젝트로 나타납니다.
- 특정 프로젝트(예: myapp)를 스캔하려면 `--sub-project=`를 사용합니다 (즉, `snyk test --sub-project=myapp`).

특정 구성을 테스트하려면 여기서 자세한 예제를 참조하십시오 [Java 및 Kotlin을 위한 Snyk](./).

Android Build 변형에 대한 자세한 내용은 [Java 및 Kotlin을 위한 Snyk](./)를 참조하십시오.