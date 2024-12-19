# Maven 플러그인 Snyk 통합

Snyk은 [Snyk CLI](https://docs.snyk.io/snyk-cli/cli-reference)를 기반으로 한 [Maven 플러그인](https://github.com/snyk/snyk-maven-plugin)을 제공합니다. 이 플러그인을 사용하면 Maven 종속성을 취약점에 대해 스캔하고 모니터링할 수 있습니다.

[Maven 중앙 저장소](https://search.maven.org/artifact/io.snyk/snyk-maven-plugin)에서 모든 릴리스를 확인하세요.

## Maven 플러그인 설치

1. [Snyk API 토큰](https://docs.snyk.io/snyk-api-info/authentication-for-api)을 얻으세요.
2. `pom.xml`에 Snyk Maven 플러그인을 추가하고 필요한 대로 구성하세요.

```xml
<!-- 예시 플러그인 구성 -->
<build>
  <plugins>
    <plugin>
      <groupId>io.snyk</groupId>
      <artifactId>snyk-maven-plugin</artifactId>
      <version>2.2.0</version>
      <inherited>false</inherited>
      <executions>
        <execution>
          <id>snyk-test</id>
          <goals>
            <goal>test</goal>
          </goals>
        </execution>
        <execution>
          <id>snyk-monitor</id>
          <goals>
            <goal>monitor</goal>
          </goals>
        </execution>
      </executions>
      <configuration>
        <apiToken>${env.SNYK_TOKEN}</apiToken>
        <args>
          <arg>--all-projects</arg>
        </args>
      </configuration>
    </plugin>
  </plugins>
</build>
```

## 지원되는 버전

* Java 8 이상.
* Maven 3.2.5 이상.

## 목표

### `code-test` (실험적)

기본 단계: `test`

프로젝트의 소스 코드를 정적으로 분석하고 발견된 취약점 목록을 제공합니다.

### `container-test` (실험적)

기본 단계: `install`

컨테이너 이미지의 레이어를 분석합니다. 스캔할 이미지의 태그를 인자로 제공해야 합니다.

```xml
<!-- 스캔할 이미지의 태그 지정 예시 -->
<configuration>
  <args>
    <arg>--print-deps</arg>
    <arg>nginx:1.9.5</arg>
  </args>
</configuration>
```

### `test`

기본 단계: `test`

프로젝트의 종속성을 스캔하고 발견된 취약점 목록을 제공합니다.

### `monitor`

기본 단계: `install`

프로젝트의 종속성 트리 스냅샷을 가져와 [snyk.io](https://snyk.io)에서 모니터링합니다. 새로운 관련 취약점, 업데이트 또는 패치가 공개될 때 알림을 받게 됩니다.

## Maven 플러그인 구성

`<configuration>` 섹션 내부에서 다음 매개변수를 구성할 수 있습니다. 모든 매개변수는 선택 사항입니다.

### `apiToken` \[문자열]

**절대로** API 토큰을 `pom.xml`에 직접 포함하지 마세요. 대신 변수를 사용하세요.

Snyk의 서비스에 액세스하기 위해 Snyk API 토큰을 제공해야 합니다. 다음 중 하나를 통해 수행할 수 있습니다:

* 구성에서 `apiToken`을 변수를 사용하여 제공.
* `SNYK_TOKEN` 환경 변수를 제공.
* 이 플러그인을 사용하기 전에 Snyk CLI를 통해 `snyk auth`로 인증.

### `skip` \[부울]

기본값: `false`

이 실행을 완전히 건너뛰기.

`mvn`을 실행할 때 `-Dsnyk.skip`을 사용하여 이 동작을 활성화할 수도 있습니다.

### `failOnIssues` \[부울]

기본값: `true`

`Snyk` CLI 도구가 보안 문제를 해결해야 한다고 지적할 때 Maven 빌드가 실패로 간주되어야 하는지 여부입니다. `false`로 설정하면 작업이 필요한 경우에도 빌드가 계속됩니다.

### `args` \[문자열 배열]

이 플러그인은 [Snyk CLI](https://github.com/snyk/snyk)를 사용하므로 `<args>`를 사용하여 지원되는 인수를 전달할 수 있습니다. 아래 예시를 참조하세요.

지원되는 CLI 옵션 목록은 [Snyk CLI 명령 및 옵션 요약](https://docs.snyk.io/snyk-cli/cli-reference)을 참조하세요.

```xml
<!-- 예시 인수 구성 -->
<configuration>
  <args>
    <arg>--severity-threshold=high</arg>
    <arg>--scan-all-unmanaged</arg>
    <arg>--json</arg>
  </args>
</configuration>
```

### `cli` \[객체]

이 플러그인에서 사용하는 Snyk CLI를 구성할 수 있습니다.

기본적으로 CLI는 자동으로 다운로드되고 업데이트됩니다.

다음에 이어지는 CLI 구성 섹션을 참조하세요.

## CLI 구성

대부분의 경우 `<cli>` 옵션을 설정할 필요가 없습니다.

세 가지 다른 모드에서 CLI를 구성할 수 있습니다:

* 자동 다운로드 및 업데이트 (기본값)
* 사용자 정의 CLI 실행 파일
* 특정 CLI 버전

각 모드에 대한 매개변수를 확인하려면 해당 링크를 따르세요.

```xml
<!-- 예시 CLI 구성 -->
<configuration>
  <cli>
    <updatePolicy>daily</updatePolicy>
  </cli>
</configuration>
```

### 자동 다운로드 및 업데이트

#### **`updatePolicy` \[문자열]**

기본값: `daily`

최신 CLI 릴리스를 다운로드하는 빈도입니다. Snyk은 항상 CLI 설치를 최신 버전으로 유지하는 것을 권장합니다.

* `daily` - 하루에 한 번
* `always` - 실행할 때마다
* `never` - 최초 다운로드 이후에 업데이트하지 않음
* `interval:<분>` - 마지막 업데이트 이후 `<분>`이 경과한 후 실행됨. 예를 들어 `interval:60`은 1시간 후에 업데이트됨

#### **`downloadDestination` \[문자열]**

기본값: OS별

다운로드한 실행 파일을 놓을 위치입니다. 기본적으로 다음과 같습니다:

* Linux - `$XDG_DATA_HOME/snyk/snyk-linux` 또는 `~/.local/share/snyk/snyk-linux`
* macOS - `~/Library/Application Support/Snyk/snyk-macos`
* Windows - `%APPDATA%\Snyk\snyk-win.exe`

### 사용자 정의 CLI 실행 파일

#### **`executable` \[문자열]**

예시: `~/.local/share/snyk/snyk-linux`

사전에 설치된 Snyk CLI 실행 파일의 경로입니다. 실행 파일은 [Snyk CLI 릴리스 페이지](https://github.com/snyk/snyk/releases)에서 확인할 수 있습니다.

### 특정 CLI 버전

#### **`version` \[문자열]**

예시: `1.542.0`

특정 버전을 사용하려면 이 옵션을 지정합니다. 버전은 [Snyk CLI 릴리스 페이지](https://github.com/snyk/snyk/releases)에서 확인할 수 있습니다.

이 옵션을 설정하면 실행될 때마다 CLI가 다운로드됩니다.

## 데모

이 플러그인을 사용해보려면 [데모 프로젝트](https://github.com/snyk/demo-snyk-maven-plugin)를 참조하세요.

## Snyk Maven 플러그인 v1에서 v2로 마이그레이션

v1의 모든 플러그인 매개변수는 `<args>` 객체로 이동하여 CLI 사용과 일치하도록 유지해야 합니다. 예를 들어:

* `org` => `<arg>--org=my-org-name</arg>`
* `failOnSeverity` => `<arg>--severity-threshold=low|medium|high</arg>`
* `failOnAuthError` => 플러그인 실행을 건너뛰도록 `<skip>true</skip>` 사용
* `includeProvidedDependencies` => `provided` 종속성은 항상 포함됩니다.

지원되는 인수 목록은 구성을 확인하세요.

***