# Maven 및 Gradle와의 SCM 통합

## Maven 및 Gradle 프로젝트용 SCM 통합

### Maven 프로젝트

Maven 애플리케이션을 스캔할 때 Snyk는 각 `pom.xml` 파일에 대해 프로젝트를 생성합니다. 이 프로젝트에는 해당 파일과 관련된 모든 직접 및 간접 종속성이 포함됩니다.

프로젝트는 `compile`, `provided`, `runtime` 범위의 프로덕션 종속성만 포함합니다.

### Gradle 프로젝트

가져올 프로젝트를 선택한 후, Snyk는 `build.gradle` 파일과 (선택 사항으로) `gradle.lockfile`에 기반하여 종속성 트리를 빌드합니다.

Gradle 프로젝트(그루비 및 코틀린 DSL 포함)용 개선된 스캔이 이제 [이 페이지](git-repositories-with-maven-and-gradle.md#improved-gradle-scm-scanning-early-access)에 설명된 대로 얼리 액세스로 제공됩니다.

`api`, `compile`, `classpath`, `implementation`, `runtime`, `runtimeOnly` 구성에서만 프로덕션 종속성이 포함됩니다.

가능하다면, 애플리케이션에서 [Gradle lockfiles](https://docs.gradle.org/current/userguide/dependency_locking.html)을 활성화하세요. lockfile이 존재하면, Snyk는 프로젝트에서 사용되는 종속성의 최종 버전을 보다 정확하게 해결할 수 있습니다.

lockfile이 없는 Gradle 프로젝트의 경우, 가장 정확한 결과를 얻기 위해 Snyk CLI 사용을 권장합니다.

## 개선된 Gradle SCM 스캐닝

{% hint style="info" %}
**릴리스 상태**

개선된 Gradle SCM 스캐닝은 얼리 액세스 상태에 있습니다. [Snyk 미리보기](../../snyk-admin/snyk-preview.md)를 사용하여 기능을 활성화할 수 있습니다.
{% endhint %}

개선된 Gradle SCM 스캐닝을 사용하여 Git 통합을 통해 가져온 Gradle 프로젝트에 대해 더 정확한 결과를 얻을 수 있습니다.

### 지원되는 Gradle 기능

다음은 주요 지원되는 Gradle 기능 목록입니다:

* [Groovy](https://docs.gradle.org/current/userguide/groovy_build_script_primer.html) 및 [Kotlin](https://docs.gradle.org/current/userguide/kotlin_dsl.html) DSL - `build.gradle(.kts)` 및 `settings.gradle(.kts)`
* 기본 및 사용자 지정 패키지 저장소 ([Artifactory](https://docs.gradle.org/current/userguide/declaring_repositories.html#sec:declaring_public_repository), [Nexus](https://docs.gradle.org/current/userguide/declaring_repositories.html#sec:declaring_custom_repository) 예시)
* 내장 객체 [`ext`](https://docs.gradle.org/current/dsl/org.gradle.api.plugins.ExtraPropertiesExtension.html), [`project`](https://docs.gradle.org/current/dsl/org.gradle.api.Project.html), `rootProject`, 및 [`settings`](https://docs.gradle.org/current/dsl/org.gradle.api.initialization.Settings.html)
* 로컬 및 전역 변수, 맵, 문자열 보간
* Gradle [lockfiles](https://docs.gradle.org/current/userguide/dependency_locking.html)
* [Gradle 속성 및 시스템 속성](https://docs.gradle.org/current/userguide/build_environment.html#sec:gradle_system_properties) - `gradle.properties`
* [종속성 제외](https://docs.gradle.org/current/userguide/dependency_downgrade_and_exclude.html#sec:excluding-transitive-deps)
* [Gradle](https://docs.gradle.org/current/userguide/platforms.html#sub:version-catalog-declaration) 및 [TOML 파일](https://docs.gradle.org/current/userguide/platforms.html#sub::toml-dependencies-format)에서 선언된 버전 카탈로그 - `gradle/libs.versions.toml`
* 다중 프로젝트 빌드, 프로젝트 이름, 프로젝트 참조
* [Spring의 `mavenBom`](https://docs.spring.io/dependency-management-plugin/docs/current/reference/html/#dependency-management-configuration-bom-import)
* [`platform`](https://docs.gradle.org/current/userguide/platforms.html#sub:using-platform-to-control-transitive-deps) 종속성으로서의 Maven BOM

일부 Gradle 기능은 지원되지 않을 수 있으며, 이는 스캔 결과에 영향을 줄 수 있습니다. 이러한 Gradle 기능에는 다음이 포함됩니다:

* [buildSrc](https://docs.gradle.org/current/userguide/organizing_gradle_projects.html#sec:build_sources) 디렉토리의 사용자 정의 구성
* [플러그인](https://docs.gradle.org/current/userguide/plugins.html)을 통해 도입된 종속성

이 얼리 액세스 기능에서 예상치 못한 결과가 나타나는 경우에는 [Snyk 지원팀](https://support.snyk.io)에 문의하세요.

### 개선된 Gradle SCM 스캐닝을 활성화하는 방법

{% hint style="warning" %}
개선된 Gradle 스캐닝은 Git 저장소 당 최대 5,000개의 `build.gradle(.kts)` 파일을 가져올 수 있습니다. 5,000개 이상의 Gradle 빌드 파일이 있는 저장소를 가져오려고 하면 실패할 것입니다.
{% endhint %}

이 기능을 활성화하려면 다음과 같이 진행하세요(귀하의 Snyk 조직에 대해):

1. [패키지 저장소 통합](../../scan-with-snyk/snyk-open-source/package-repository-integrations/)을 구성하십시오 (Artifactory 또는 Nexus를 사용하는 경우 [아래](git-repositories-with-maven-and-gradle.md#package-repository-integrations)를 참조하세요).
2. Snyk 미리보기에서 [SCM 통합을 위한 워크스페이스](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/introduction-to-git-repository-integrations/workspaces-for-scm-integrations.md)를 활성화하십시오.
3. Snyk 미리보기에서 **개선된 Gradle 스캐닝**을 활성화하십시오.

개선된 Gradle SCM 스캐닝이 활성화된 후:

* 이전에 가져온 Git 저장소는 수동 또는 주기적 테스트하는 다음에 자동으로 업데이트됩니다.
* Gradle Kotlin DSL 프로젝트의 결과를 보려면 저장소를 다시 가져와야 합니다.

## Java에 대한 Snyk의 언어 설정 구성

조직 수준에서 오픈 소스 및 라이선스에 대한 언어 설정을 구성하십시오. 구성 설정은 해당 조직의 모든 프로젝트에 적용됩니다.

1. Snyk Web UI를 열고 **Settings >** **Languages** 섹션으로 이동하십시오.
2. **Languages** 아래의 **Java**로 이동하고 **Edit settings**를 선택하십시오.
3. **Maven** 구성을 수행하십시오.
4. 변경 사항을 저장하려면 **Update Settings**를 클릭하십시오.

## 패키지 저장소 통합

{% hint style="info" %}
응용 프로그램 빌드에 프라이빗 패키지 저장소를 사용하는 경우 정확한 결과를 얻으려면 관련 Snyk 통합을 구성해야 합니다.
{% endhint %}

{% hint style="info" %}
개선된 Gradle 스캐닝 얼리 액세스 기능을 사용하기 위해 패키지 저장소 통합을 원하는 경우 Maven의 구성 지침 및 설정을 사용하십시오.

이러한 구성은 개선된 Gradle 스캔에 감지되고 사용될 것입니다.
{% endhint %}

Java 언어 설정에서 Snyk를 프라이빗 패키지 저장소(Artifactory 또는 Nexus 등)와 통합할 수 있습니다.

이를 통해 Snyk는 프라이빗 패키지를 참조하는 Maven 또는 Gradle (얼리 액세스) 프로젝트를 스캔할 때 완전한 종속성 트리를 작성할 수 있습니다.

자세한 내용은 [Maven을 위한 Artifactory 레지스트리](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/artifactory-registry-for-maven.md)에서 확인하십시오. [패키지 저장소 통합](../../scan-with-snyk/snyk-open-source/package-repository-integrations/)에서 찾을 수 있습니다.