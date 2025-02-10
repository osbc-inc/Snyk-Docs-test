# Snyk 오픈 소스별 CI/CD 전략

이러한 전략은 Snyk SCA ([소프트웨어 구성 분석](https://snyk.io/blog/what-is-software-composition-analysis-sca-and-does-my-company-need-it/)) 테스트 기능을 사용하는 팀에 유용합니다.

빌드 작업의 실패 상태를 사용자 정의하기 위해 `--fail-on` 및 `--severity-threshold`와 같은 CLI 플래그를 사용하세요. 좀 더 고급 사용을 위해서는 `--json` 옵션을 사용하여 전체 취약점 보고서를 포함하는 JSON 파일을 생성하고 JSON 데이터에 기반하여 자체 빌드 실패 상태를 설정할 수 있습니다.

## Gradle 및 Scala

* "Multi-project" 구성에서 모든 하위 프로젝트를 테스트합니다. `monitor` 또는 `test` 명령어와 함께 이 옵션을 사용하세요: `--all-sub-projects`.
* 특정 구성을 스캔하려면 구성 속성의 특정 값들을 선택하여 종속성을 해결하세요. `test` 또는 `monitor` 명령어와 함께 이 옵션을 사용하세요: `--configuration-attributes=`.

## Python

* Snyk은 Python을 사용하여 종속성을 스캔하고 찾습니다. Snyk은 스캔을 시작하기 위해 Python 버전이 필요하며 기본값은 `python`입니다. 여러 Python 버전을 사용하는 경우, 실행에 올바른 Python 명령을 지정하기 위해 `test` 또는 `monitor` 명령어와 함께 `--command=` 옵션을 사용하세요. 다음은 예시입니다:\
  `snyk test --command=python3`
* `setup.py` 파일을 대상으로 설정해야 합니다. 명령어를 사용하세요 `snyk test --file=setup.py`
* Pip 프로젝트를 스캔하고 사용하는 경우 표준 `requirements.txt`가 아닌 경우 `--file=` 옵션을 사용해야 하므로 다음 옵션을 사용하여 패키지 매니저로 Pip을 지정해야 합니다 `--package-manager=pip.`

## .NET

`.sln` 파일을 사용하는 경우, 파일의 경로를 지정하여 Snyk이 저장소의 일부인 모든 하위 프로젝트를 스캔합니다. 예를 들어:

```
snyk test --file=sln/.sln
```

## Yarn

Yarn 작업 영역의 경우, `--yarn-workspaces` 옵션을 사용하여 패키지를 테스트하고 모니터링하세요. 루트 lockfile은 모든 패키지의 스캔에 참조됩니다. 기본적으로 자동으로 발견되지 않는 하위 폴더를 찾기 위해 `--detection-depth` 옵션을 사용하세요.

{% hint style="info" %}
Yarn 작업 영역을 위한 지원은 `snyk test` 및 `snyk monitor` 명령어에만 사용할 수 있습니다.
{% endhint %}

현재 디렉토리 및 5개의 하위 디렉토리까지 발견된 작업 영역에 속한 패키지만 스캔하는 예시 명령어가 다음과 같습니다.

```
snyk test --yarn-workspaces --detection-depth=6
```

모든 검색된 작업 영역에 적용할 제왐 및 패치를 유지하기 위해 공통 [`.snyk` 정책 파일](../../../manage-risk/policies/the-.snyk-file.md)을 사용할 수 있습니다. 이것을 위해 정책 경로를 다음과 같이 제공하세요:

```
snyk test --yarn-workspaces --policy-path=src/.snyk
```

## Monorepo

일부 고객은 단일 저장소에 여러 언어, 패키지 매니저 및 프로젝트가 있는 복잡한 프로젝트를 가지고 있습니다. 이를 간단히하기 위해 다음과 같은 접근 방식을 취할 수 있습니다:

*   **각 프로젝트 및 언어를 빌드할 때**, 특정 프로젝트 파일을 타겟팅하는 `snyk test`를 실행하는 지시문을 추가하세요. 예를 들어:

    ```
    snyk test --file=package.json
    ```
* 각 프로젝트의 종속성을 설치한 후, `pom.xml`과 같은 특정 아티팩트를 가리키는 유사한 호출을 만들 수 있습니다. 이 방법은 빠르고 효율적이지만 프로젝트에 익숙하지 않으면 확장하기 어려울 수 있습니다.
* 대부분의 **Gradle 프로젝트**의 경우, `--all-projects`를 사용하면 Gradle 특정 옵션을 내부적으로 호출하므로 `--all-projects` 검색의 일부로 빌드 파일을 발견하면 `snyk test --file=build.gradle --all-sub-projects`과 같은 형태로 동작합니다.
* **Gradle**은 추가 구성 매개변수를 필요로 할 수 있습니다. 그렇다면, 다른 언어 및 패키지 매니저의 각 manifest에 대해 `--file=`을 사용하여 다른 아티팩트를 타겟팅하세요. 그런 다음 복잡한 Gradle 프로젝트를 스캔하기 위해 `--all-sub-projects` 및 필요에 따라 `--configuration-matching`을 사용해야 합니다.

더 많은 정보를 확인하려면 [Java 및 Kotlin](../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/)을 확인하세요.

##
