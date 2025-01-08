# Java 및 Kotlin용 Snyk CLI

`Snyk` CLI를 사용하여 다음과 같이 Maven 및 Gradle 프로젝트를 테스트하려면 `snyk test` 명령을 사용하세요:

* **Gradle와 함께 Snyk CLI**: 의존성 그래프를 빌드하기 위해 Snyk는 Gradle과 통합되어 도구에서 보고된 의존성을 검사합니다. 다음 매니페스트 파일이 지원됩니다: `build.gradle` (Groovy DSL) 및 `build.gradle.kts` (Kotlin DSL).
* **Maven과 함께 Snyk CLI**: 의존성 트리를 빌드하기 위해 Snyk는 Maven과 통합되어 도구에서 보고된 의존성을 검사합니다. 다음 매니페스트 파일이 지원됩니다: `pom.xml`.

이 페이지는 Maven 및 Gradle 프로젝트에서 Snyk CLI 도움말을 사용하는 방법에 대한 세부 정보를 제공하며 [ant 및 ivy에 대한 해결책](snyk-cli-for-java-and-kotlin.md#workaround-for-ant-and-ivy)도 제공합니다.

다음 표는 의존성을 스캔하기 위한 옵션 목록을 나열합니다. `snyk test` 및 `snyk monitor` 명령을 다룹니다. 이러한 명령을 위한 모든 옵션 목록을 보려면 [CLI 명령 및 옵션 요약](../../snyk-cli/cli-commands-and-options-summary.md)을 참조하세요.

| 패키지 관리자 및 환경   | 테스트 도움말                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | 모니터 도움말                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Maven          | <p><code>--maven-aggregate-project</code><br>더 많은 세부 정보를 보려면 <a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-maven-projects">Maven 프로젝트 옵션</a>을 참조하세요.<br><br>통합 프로젝트를 위한 예시:<br><code>snyk test --maven-aggregate-project</code><br><br>통합 프로젝트가 아닌 경우: <code>snyk test --all-projects</code><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> 루트 pom.xml 파일과 동일한 디렉터리에서 옵션을 실행해야 합니다.</p>                                                                                                                                                                                                                                                           | <p><code>--maven-aggregate-project</code><br>더 많은 세부 정보를 보려면 <a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-maven-projects">Maven 프로젝트 옵션</a>을 참조하세요.<br><br>통합 프로젝트를 위한 예시:<br><code>snyk monitor --maven-aggregate-project</code><br><br>통합 프로젝트가 아닌 경우: <code>snyk monitor --all-projects</code><br><br><span data-gb-custom-inline data-tag="emoji" data-code="2139">ℹ️</span> 루트 pom.xml 파일과 동일한 디렉터리에서 옵션을 실행해야 합니다.</p>                                                                                                                                                                                                                                       |
| Gradle         | <p><code>--sub-project=&#x3C;이름></code>, <code>--gradle-sub-project=&#x3C;이름></code> - 특정 Gradle 하위 프로젝트를 테스트합니다.</p><p><code>--all-sub-projects</code> - 모든 Gradle 하위 프로젝트를 테스트합니다.</p><p><code>--all-projects</code> - 모든 Gradle 프로젝트를 테스트합니다.</p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - 지정한 Java 정규 표현식과 일치하는 첫 번째 구성을 사용하여 의존성을 해결합니다.</p><p><code>--configuration-attributes=&#x3C;속성>[,&#x3C;속성>]...</code> - 설치 및 의존성 해결을 위해 구성 속성의 특정 값을 선택합니다.</p><p><code>--init-script=&#x3C;파일></code> - Gradle 초기화 스크립트가 있는 프로젝트에 사용됩니다.</p><p><a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-gradle-projects">Gradle 프로젝트 옵션</a>에서 자세한 내용을 확인하세요.</p> | <p><code>--sub-project=&#x3C;이름></code>, <code>--gradle-sub-project=&#x3C;이름></code> - 특정 Gradle 하위 프로젝트를 모니터링합니다.</p><p><code>--all-sub-projects</code> - 모든 Gradle 하위 프로젝트를 모니터링합니다.</p><p><code>--all-projects</code> - 모든 Gradle 프로젝트를 모니터링합니다.</p><p><code>--configuration-matching=&#x3C;CONFIGURATION_REGEX></code> - 지정한 Java 정규 표현식과 일치하는 첫 번째 구성을 사용하여 의존성을 해결합니다.</p><p><code>--configuration-attributes=[,]...</code> - 설치 및 의존성 해결을 위해 구성 속성의 특정 값을 선택합니다.</p><p><code>--init-script=&#x3C;파일></code> - Gradle 초기화 스크립트가 있는 프로젝트에 사용됩니다.<br><br><a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-gradle-projects">Gradle 프로젝트 옵션</a>에서 자세한 내용을 확인하세요.</p> |
| 빌드 도구          | <p><code>snyk test -- [&#x3C;특정_컨텍스트_옵션>]</code><br><br><a href="https://docs.snyk.io/snyk-cli/commands/test#options-for-build-tools">빌드 도구 옵션</a>에서 자세한 내용을 확인하세요.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <p><code>snyk monitor -- [&#x3C;특정_컨텍스트_옵션>]</code><br></p><p><a href="https://docs.snyk.io/snyk-cli/commands/monitor#options-for-build-tools">빌드 도구 옵션</a>에서 자세한 내용을 확인하세요.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| 관리되지 않는 JAR 파일 | <p><code>--scan-unmanaged</code> - 관리되지 않는 파일 테스트<br></p><p><code>--scan-unmanaged --file=&#x3C;JAR_FILE_NAME></code> - 개별 JAR, WAR 및 AAR 파일 테스트<br></p><p><code>--scan-all-unmanaged</code> - 현재 폴더에서 Maven, JAR, WAR 및 AAR 파일을 재귀적으로 자동 감지합니다.<br><br><a href="../../snyk-cli/commands/test.md#scan-all-unmanaged">관리되지 않는 JAR 파일에 대한 옵션</a>에서 자세한 내용을 확인하세요.</p>                                                                                                                                                                                                                                                                                                                                  | <p><br><br><code>--scan-unmanaged</code> - 관리되지 않는 파일 테스트<br></p><p><code>--scan-unmanaged --file=&#x3C;JAR_FILE_NAME></code> - 개별 JAR, WAR 및 AAR 파일 테스트<br></p><p><code>--scan-all-unmanaged</code> - 현재 폴더에서 Maven, JAR, WAR 및 AAR 파일을 재귀적으로 자동 감지합니다.</p><p><a href="../../snyk-cli/commands/monitor.md#scan-all-unmanaged">관리되지 않는 JAR 파일에 대한 옵션</a>에서 자세한 내용을 확인하세요.</p>                                                                                                                                                                                                                                                                                                             |

## Maven 프로젝트용 CLI 도움말

**Maven 집계 프로젝트**는 모듈과 상속을 사용하는 프로젝트입니다.

이러한 유형의 프로젝트를 스캔할 때 Snyk은 Maven 리액터에 의해 수정 가능한 모든 모듈을 보장하기 위해 컴파일을 수행합니다.

*   **집계 프로젝트를 스캔하려면**, `--maven-aggregate-project` 옵션을 사용하세요:

    ```
    snyk test --maven-aggregate-project
    ```
*   **집계 프로젝트가 아닌 프로젝트를 스캔하려면**, `--all-projects` 옵션을 사용하세요:

    ```
    snyk test --all-projects
    ```

동일한 옵션은 `snyk monitor`에서도 사용할 수 있습니다.

루트 `pom.xml` 파일이 있는 동일한 디렉터리에서 옵션을 실행해야 합니다.

개별 하위 프로젝트는 웹 UI에서 별도의 Snyk 프로젝트로 나타납니다.

Maven 특정 옵션들을 Snyk CLI와 함께 사용하는 예시는 다음과 같습니다.

1. "prod"라는 특정 Maven 프로필을 테스트합니다.

```
snyk test -- -prod
```

2. pom.xml 파일에서 시스템 속성을 추가합니다. 예를 들어, pom.xml에 나타나는 패키지 버전은 다음과 같습니다:

```
${pkg_version}
```

3. 다음과 같이 시스템 속성을 정의합니다:

```
snyk test -- -Dpkg_version=1.4
```

## Gradle 프로젝트용 CLI 도움말

Gradle 빌드는 여러 하위 프로젝트로 구성될 수 있으며 각 하위 프로젝트가 자체 build.gradle을 갖고 있습니다. 루트 프로젝트는 `settings.gradle` 파일을 포함한 현재 폴더의 프로젝트이며, 하위 프로젝트는 루트 프로젝트에 의존하지만 다르게 구성할 수 있습니다.

기본적으로 Snyk CLI는 현재 프로젝트, 현재 폴더의 루트 프로젝트 또는 `--file=path/to/build.gradle`을 지정한 프로젝트만 스캔합니다.

*   한꺼번에 모든 프로젝트를 스캔하려면 (권장), `--all-sub-projects` 옵션을 사용하세요:

    ```
    snyk test --all-sub-projects
    ```

각각의 개별 하위 프로젝트는 웹 UI에서 별도의 Snyk 프로젝트로 나타납니다.

*   특정 프로젝트를 스캔하려면 (예: myapp):

    ```
    snyk test --sub-project=myapp
    ```

### Gradle 구성

Gradle 의존성은 특정 범위에 대해 선언됩니다. 각 범위는 Gradle에 의해 [구성](https://docs.gradle.org/current/userguide/declaring_dependencies.html#sec:what-are-dependency-configurations)으로 표현됩니다. 예를 들어:

* **implementation**: 런타임에서 필요한 종속성에 대한 구성으로 컴파일 타임 및 런타임에 필요하지만 소비자에게 노출되지 않습니다.
* **api**: 소비자에게 노출되는 컴파일 및 런타임에서 필요한 종속성을 위한 구성.
* **compileOnly**: 컴파일 타임에만 필요한 종속성을 위한 구성.
* **runtimeOnly**: 런타임에만 필요한 종속성을 위한 구성.
* **compileClasspath**: 컴파일 타임에 필요한 종속성을 위한 구성.

대부분의 경우, Snyk은 **compileClasspath** 구성에 모든 종속성을 포함하나 일부 경우에는 다를 수 있습니다.

특정 구성을 테스트하려면:

* [Java 정규 표현식](https://docs.oracle.com/javase/tutorial/essential/regex/) (대소문자를 구분하지 않는)을 매개변수로 사용하여 `--configuration-matching` 옵션을 사용하십시오. 일치하는 첫 번째 구성만 테스트됩니다.
* 서로 다른 하위 프로젝트가 서로 다른 구성을 포함하는 경우 각 하위 프로젝트를 별도로 스캔합니다.

\--configuration-matching 옵션을 사용하는 방법 예시

* `--configuration-matching=compile`은 compile, testCompile, compileOnly 등을 일치시킵니다.
* `--configuration-matching=^compile$`은 compile만 일치시킵니다.
* `--configuration-matching='^(debug|release)compile$'`는 debugCompile 및 releaseCompile을 일치시킵니다.
* `--configuration-matching='^(?!.*test).*$'`은 "test" 문자열이 포함된 구성을 제외한 모든 구성을 일치시킵니다.

### Android 빌드 변형

Android Gradle은 [빌드 변형](https://developer.android.com/studio/build/build-variants)을 구성하여 앱의 다른 버전을 생성할 수 있습니다.

Snyk의 기본 동작은 사용 가능한 모든 구성을 병합하는 것이기 때문에, 반복되는 변형은 병합할 수 없는 구성 충돌을 일으킵니다.

이러한 상황에서 Gradle에서 오류가 발생하면 다음 메시지 중 하나가 포함된 Gradle 오류로 인해 Snyk 스캔이 실패할 수 있습니다:

* `project :mymodulewithvariants`의 다음 구성을 선택할 수 없습니다
* `project :mymodulewithvariants`의 다음 변형을 선택할 수 없습니다
* 후보 값에서 값을 선택할 수 없습니다

이러한 충돌을 피하려면:

*   **특정 구성 사용:** 특정 빌드 구성을 알고 있는 경우 필요한 모든 속성이 포함된 빌드 구성이 모든 테스트 대상 하위 프로젝트에 동일한 경우 해당 구성을 지정합니다.\
    예시:

    ```
    --configuration-matching=prodReleaseRuntimeClasspath
    ```
*   **명시적으로 의존성 구성 지정:** build.gradle 파일의 intra-project 의존성을 수정하여 특정 구성을 사용합니다.

    ```
      dependencies {
          implementation project(path: ':mymodulewithvariants', configuration: 'default')
      }
    ```
* **구성 속성 제안:** 명령을 실행할 때 오류가 발생하면 오류에서 사용 가능한 속성 값을 나타내며 Gradle의 오류 세부 사항에서 어떤 종속성 변형이 어\`\`\`xml

\`\`\`

이러한 종속성 파일은 일반적으로 `build.xml`에 정의된 `ant` 작업을 사용하여 평가됩니다:

```xml
<target name="resolve-dependencies" depends="init">
    <ivy:retrieve pattern="${lib.dir}/[artifact]-[revision].[ext]"/>
</target>
```

`ant resolve-dependencies` 명령을 사용하면 의존성이 Maven Central에서 다운로드되며 일반적인 Maven 의존성처럼 작동합니다.

의존성 트리를 Snyk에 알리려면 먼저 Maven POM 형식으로 변환해야 합니다. `build.xml`에서 새로운 `makepom` 작업을 구성하여 시작하십시오.

```xml
<target name="makepom" depends="resolve-dependencies">
    <ivy:makepom ivyfile="${basedir}/ivy.xml" pomfile="${basedir}/pom.xml" conf="default,runtime">
        <mapping conf="default" scope="compile"/>
        <mapping conf="runtime" scope="runtime"/>
    </ivy:makepom>
</target>
```

이제 다음 명령을 실행할 수 있습니다:

```
ant makepom
snyk test --file=pom.xml
```

`pom.xml` 파일은 체크인할 필요가 없으며 `snyk`를 사용하여 테스트가 완료된 후 삭제할 수 있습니다. 또한 의존성 트리는 다음을 사용하여 모니터링할 수 있습니다:

```
snyk monitor --file=pom.xml
```
