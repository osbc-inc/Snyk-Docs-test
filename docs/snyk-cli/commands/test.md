# Test

## 사용법

`snyk test [<OPTIONS>]`

## 설명

`snyk test` 명령은 오픈소스 취약점 및 라이선스 문제를 확인하기 위해 프로젝트를 검사합니다. 이 테스트 명령은 종속성을 가진 지원되는 매니페스트 파일을 자동으로 감지하고 그들을 테스트합니다.

**참고:** Snyk Code, Container 및 IaC 스캔 방법에 대한 특정 `snyk test` 명령이 있습니다: `code test`, `container test`, `iac test`.

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공 (스캔 완료), 취약점 없음\
**1**: 필요한 동작 (스캔 완료), 취약점 발견\
**2**: 실패, 명령을 다시 실행하십시오. 디버그 로그를 출력하려면 `-d`를 사용하십시오.\
**3**: 실패, 지원되는 프로젝트가 감지되지 않음

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와 연결하기 위한 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하세요.

## 코드 실행 경고

코드를 스캔하기 전에 [Snyk CLI에 대한 코드 실행 경고](https://docs.snyk.io/snyk-cli/code-execution-warning-for-snyk-cli)를 검토하십시오.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 옵션

특정 빌드 환경, 패키지 매니저, 언어 및 마지막으로 지정한 `[<CONTEXT-SPECIFIC OPTIONS>]`에 대한 옵션에 대한 자세한 내용은 하위 섹션을 참조하십시오.

### `--all-projects`

Yarn 워크스페이스를 포함한 작업 디렉토리의 모든 프로젝트를 자동으로 감지합니다.

자세한 내용은 다음 문서 [Snyk CLI가 모노 리포지토리 또는 여러 매니페스트 파일을 지원합니까?](https://support.snyk.io/s/article/Does-the-Snyk-CLI-support-monorepos-or-multiple-manifest-files)를 참조하십시오.

무효한 문자열 길이 오류를 보는 경우, [프로젝트를 스캔할 때 잘못된 문자열 길이 오류](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/invalid-string-length-error-when-scanning-projects)를 참조하십시오.

### `--fail-fast`

`--all-projects`와 함께 사용하여 오류 발생 시 스캔을 중단하고 해당 오류를 사용자에게 보고하도록합니다.

종료 코드는 2이고 스캔이 종료됩니다. 오류가 발생하지 않은 프로젝트의 취약성 정보는 보고되지 않습니다.

스캔을 수행하려면 오류를 해결하고 다시 스캔하십시오.

**참고**: `--fail-fast`를 사용하지 않으면 Snyk는 모든 프로젝트를 스캔하지만 구성 오류 또는 다른 오류로 인해 스캔할 수 없는 프로젝트에 대해 취약성을 보고하지 않습니다.

### `--detection-depth=<DEPTH>`

`--all-projects` 또는 `--yarn-workspaces`와 함께 사용하여 검색할 하위 디렉토리의 수를 나타냅니다. `DEPTH`는 숫자여야하며 1 이상이어야 합니다. 0은 현재 디렉토리를 나타냅니다.

기본값: 제한 없음.

예: `--detection-depth=3`은 지정된 디렉토리(또는 지정된 `<PATH>`가 없는 경우 현재 디렉토리)와 하위 디렉터리 세 단계까지 검색을 제한합니다. 0은 현재 디렉토리를 나타냅니다.

### `--exclude=<NAME>[,<NAME>]...>`

`--all-projects` 및 `--yarn-workspaces`와 함께 사용하여 제외할 디렉터리 및 파일 이름을 나타냅니다. 콤마로 구분되어야하며 경로를 포함할 수 없습니다.

예: `$ snyk test --all-projects --exclude=dir1,file2`

이렇게 하면 `dir1` 및 `file2`라는 디렉토리와 파일이 프로젝트 매니페스트 파일 스캔 중 제외됩니다: `./dir1`, `./src/dir1`, `./file2`, `./src/file2` 등.

**참고**: `--exclude=dir1`은 `./dir1` 및 `./src/dir1`을 찾을 것입니다.\
그러나 `--exclude=./src/dir1`은 경로를 포함하기 때문에 오류가 발생합니다.

### `--prune-repeated-subdependencies`, `-p`

중복된 하위 종속성을 제거하여 종속성 트리를 정리합니다.

모든 취약점을 계속해서 찾지만 취약한 경로를 모두 찾지 못할 수 있습니다.

큰 프로젝트가 테스트에 실패하는 경우에 사용하십시오.

기본값: false

### `--print-deps`

분석을 위해 보내기 전에 종속성 트리를 출력합니다.

### `--remote-repo-url=<URL>`

모니터링할 리모트 URL을 설정하거나 재정의합니다.

단일 타겟 아래 발견된 모든 프로젝트를 그룹화합니다.

### `--dev`

개발 전용 종속성을 포함합니다. 일부 패키지 관리자에만 해당됩니다, 예를 들어 npm의 `devDependencies`나 Gemfile의 `:development` 종속성 등.

**참고**: 이 옵션은 Maven, npm 및 Yarn 프로젝트와 함께 사용할 수 있습니다.

기본값: false, 제품 종속성만 검색

### `--org=<ORG_ID>`

특정 Snyk 조직과 연관된 Snyk 명령을 실행할 `<ORG_ID>`를 지정합니다. `<ORG_ID>`는 모든 자주 사용되는 기능과 개인 테스트 제한에 영향을 줍니다.

여러 조직이 있는 경우 CLI를 통해 기본값을 설정할 수 있습니다:

`$ snyk config set org=<ORG_ID>`

기본값: 현재 계정 설정의 기본 조직인 `<ORG_ID>`

**참고:** `--org=<orgslugname>.`를 사용할 수도 있습니다. `ORG_ID`는 CLI와 API 모두에서 작동하지만 조직 slug 이름은 CLI에서만 작동하며 API에서는 작동하지 않습니다.

`orgslugname`은 Snyk UI의 org URL에 표시된 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`. org 이름은 작동하지 않습니다.

자세한 내용은 다음 문서 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli)를 참조하십시오.

### `--file=<FILE>`

패키지 파일을 지정합니다.

로컬에서 테스트하거나 프로젝트를 모니터링할 때 Snyk이 패키지 정보를 검사해야 하는 파일을 지정할 수 있습니다. 파일이 지정되지 않은 경우 Snyk는 프로젝트에 적합한 파일을 감지하려고 시도합니다.

더 자세한 내용은 [Python 프로젝트 옵션](https://docs.snyk.io/snyk-cli/commands/test#options-for-python-projects)부분을 참조하십시오.

### `--package-manager=<PACKAGE_MANAGER_NAME>`

`--file=<FILE>` 옵션과 함께 사용되는 패키지 관리자의 이름을 지정합니다. 이를 통해 Snyk가 파일을 찾을 수 있습니다.

예: `$ snyk test --file=req.txt --package-manager=pip`

자세한 내용은 [Python 프로젝트 옵션](https://docs.snyk.io/snyk-cli/commands/test#options-for-python-projects)을 참조하십시오.

### `--unmanaged`

C++ 전용, 알려진 오픈 소스 종속성을 모든 파일에서 스캔합니다.

`--unmanaged` 사용 가능한 옵션은 [Unmanaged을 사용한 스캔 옵션들](https://docs.snyk.io/snyk-cli/commands/test#options-for-scanning-using-unmanaged)을 참조하십시오.

### `--ignore-policy`

모든 설정 정책, .snyk 파일의 현재 정책, Org 수준 무시 및 snyk.io의 프로젝트 정책을 무시합니다.

### `--trust-policies`

의존성의 Snyk 정책에서 무시 규칙을 적용하고 사용합니다; 그렇지 않으면 의존성의 무시 규칙은 제안으로만 표시됩니다.

### `--show-vulnerable-paths=<none|some|all>`

취약한 패키지의 최상위 종속성부터 취약한 패키지까지의 종속성 경로를 표시합니다. `--json-file-output`과 호환되지 않습니다.

기본값: `some`, 몇 가지 예제 경로 표시. `false`는 `none`의 별칭입니다.

예: `--show-vulnerable-paths=none`

### `--project-name=<PROJECT_NAME>`

사용자 정의 Snyk 프로젝트 이름을 지정합니다.

### `--target-reference=<TARGET_REFERENCE>`

이 프로젝트를 구분하는 참조를 지정합니다. 예를 들어 분기 이름이나 버전입니다. 동일한 참조를 가진 프로젝트는 해당 참조를 기반으로 그룹화할 수 있습니다. `--unmanaged` 사용하지 않고 Snyk Open Source에서 지원됩니다.

자세한 내용은 [모니터링을 위해 브랜치 또는 버전별로 프로젝트 그룹화](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring)를 참조하십시오.

테스트 실행 시 `--target-reference=<TARGET_REFERENCE>`를 사용하여 모니터링 대상에 대한 동일한 무시 및 정책을 적용할 수 있습니다.

자세한 내용은 [특정 이슈 무시](https://docs.snyk.io/scan-using-snyk/find-and-manage-priority-issues/ignore-issues)를 참조하십시오.

### `--policy-path=<PATH_TO_POLICY_FILE>`

`.snyk` 정책 파일의 경로를 수동으로 전달합니다.

### `--json`

콘솔에서 결과를 JSON 데이터 구조로 인쇄합니다.

예: `$ snyk test --json`

무효한 문자열 길이 오류를 보는 경우, [프로젝트를 스캔할 때 잘못된 문자열 길이 오류](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/invalid-string-length-error-when-scanning-projects)를 참조하십시오.

### `--json-file-output=<OUTPUT_FILE_PATH>`

`--json` 옵션을 사용하든 사용하지 않든, 결과를 JSON 데이터 구조로 지정된 파일에 직접 저장합니다.

사용 가능한 `--json-file-output=vuln.json`의 예제와 같이 작동합니다.

오픈 소스의 경우, 문제가 발견되더라도 Snyk는 항상 파일을 만듭니다. 그에 반해 SAST에서 문제가 발견되지 않으면 `json` 파일을 만들지 않습니다.

예: `$ snyk test --json-file-output=vuln.json`

무효한 문자열 길이 오류를 보는 경우, [프로젝트를 스캔할 때 잘못된 문자열 길이 오류](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/invalid-string-length-error-when-scanning-projects)를 참조하십시오.

### `--sarif`

결과를 SARIF 형식으로 반환합니다.

### `--sarif-file-output=<OUTPUT_FILE_PATH>`

`--sarif` 옵션을 사용하든 사용하지 않든, 결과를 SARIF 형식으로 지정된 파일에 직접 저장합니다.

사용 가능한 `--sarif-file-output=<OUTPUT_FILE_PATH>`의 예제와 같이 작동합니다.

### `--severity-threshold=<low|medium|high|critical>`

지정된 수준 이상의 취약점만 보고합니다.

### `--fail-on=<all|upgradable|patchable>`

고취가능한 취약점이 있는 경우에만 실패합니다. 다음 값 중 하나를 사용하십시오:

* `all`: 업그레이드 또는 패치 가능한 취약점이 하나 이상 있을 때 실패합니다.
* `upgradable`: Snyk이 계산된 복구를 사용할 수 있는 취약점이 하나 이상 있을 때 실패합니다.
* `patchable`: 패치할 수 있는 취약점이 하나 이상 있을 때 실패합니다. `patchable`을 사용할 때, 테스트는 패치 가능한 취약점이 하나 이상 있고 다른 취약점에 대한 계산된 복구가 있는 경우에도 실패합니다.

Snyk-계산 가능한 취약점에서 실패하려면 `--fail-on` 옵션을 사용하지 않도록 하십시오. 취약성에 Snyk-계산된 수정이 없고 이 옵션을 사용하는 경우에는 테스트가 통과됩니다.

**참고**: `snyk test`로 준수할 수 없는 메타데이터로 제한된 코드를 테스트하는 경우, Snyk는 코드를 손상시키지 않도록 하기 위해 수정을 제안하지 않습니다. 수정을 식별하고 수동으로 적용할 수 있습니다.

## Maven 프로젝트 옵션

**참고**: `--dev` 옵션은 Maven 프로젝트에서 사용할 수 있습니다. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#dev)을 참조하십시오.

### `--maven-aggregate-project`

모듈 및 상속을 사용하는 프로젝트, 즉 Maven 집합 프로젝트를 스캔할만약 이 폴더에 고유한 이름을 할당했다면, Snyk은 커스텀 경로를 입력해야만 해당 폴더를 찾을 수 있습니다.

의존성이 있는 폴더의 이름을 포함한 절대 또는 상대 경로를 사용하세요.

예시:

유닉스 OS의 경우 `snyk test --packages-folder=../location/to/packages`

윈도우의 경우 `snyk test --packages-folder=..\location\to\packages`

### `--project-name-prefix=<PREFIX_STRING>`

.NET 프로젝트를 모니터링할 때, 이 옵션을 사용하여 프로젝트 내 파일 이름에 사용자 정의 접두사를 추가할 수 있습니다. 필요한 구분자와 함께 사용하세요.

예시: `snyk monitor --file=my-project.sln --project-name-prefix=my-group/`

이 옵션은 다른 `.sln` 파일에 동일한 이름을 갖는 여러 프로젝트가 있는 경우 유용합니다.

## .NET 프로젝트 옵션

### `--dotnet-runtime-resolution`

**참고:** 이 옵션은 아직 초기 접근 단계에 있으며 출시될 때까지 변경될 수 있습니다.

필수 항목입니다. [런타임 해상도 스캔](https://docs.snyk.io/getting-started/supported-languages-and-frameworks/.net/improved-.net-scanning)을 사용하여 .NET 프로젝트를 테스트할 때 이 옵션을 사용해야 합니다.

예시: `snyk test --dotnet-runtime-resolution`

### `--dotnet-target-framework`

**참고:** 이 옵션은 아직 초기 접근 단계에 있으며 출시될 때까지 변경될 수 있습니다.

선택 사항입니다. 솔루션에 여러 `<TargetFramework>` 지시문이 포함될 경우 이 옵션을 사용할 수 있습니다. `--dotnet-target-framework` 옵션을 지정하지 않으면 모든 지원되는 Target Framework가 스캔됩니다.

이 옵션으로 지정된 Target Framework는 표준 [이름 지정 규칙](https://learn.microsoft.com/en-us/dotnet/standard/frameworks#supported-target-frameworks)을 따라야 합니다.

예시: `snyk test --dotnet-runtime-resolution --dotnet-target-framework=net6.0`

## npm 프로젝트 옵션

**참고**: npm 프로젝트에서 다음 옵션을 사용할 수 있습니다.

`--dev`. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#dev) 참조

`--all-projects`: 디렉터리의 npm 프로젝트 및 다른 프로젝트를 스캔하고 감지합니다. [`--all-projects` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#all-projects) 참조

`--fail-on`: [--fail-on 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#fail-on-less-than-all-or-upgradable-or-patchable-greater-than) 참조

`--prune-repeated-subdependencies, -p`: [`--prune-repeated subdependencies` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p) 참조

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 lockfile을 테스트하는 것을 방지합니다.

프로젝트에 동기화가 맞지 않는 lockfile이 있다면, `--strict-out-of-sync=true`로 설정할 때 `test` 명령이 실패합니다.

기본: true

## pnpm 프로젝트 옵션

**Snyk CLI pnpm 지원은 아직 초기 접근 단계입니다**. 활성화하려면 Snyk 계정에 들어가서 Settings를 선택하고 Preview를 선택한 후 CLI v1.1293.0 이상을 설치하세요.

**참고**: pnpm 프로젝트에서 다음 옵션을 사용할 수 있습니다:

`--dev`. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#dev) 참조

`--all-projects`: 디렉터리의 pnpm 프로젝트 및 다른 프로젝트를 스캔하고 감지합니다. [`--all-projects` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#all-projects) 참조

`--fail-on`: [--fail-on 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#fail-on-less-than-all-or-upgradable-or-patchable-greater-than) 참조

`--prune-repeated-subdependencies, -p`: [`--prune-repeated subdependencies` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p) 참조

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 lockfile을 테스트하는 것을 방지합니다.

프로젝트에 동기화가 맞지 않는 lockfile이 있다면, `--strict-out-of-sync=true`로 설정할 때 `test` 명령이 실패합니다.

기본: true

## Yarn 프로젝트 옵션

**참고**: Yarn 프로젝트에서 다음 옵션을 사용할 수 있습니다:

`--dev`. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#dev) 참조

`--fail-on`: [--fail-on 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#fail-on-less-than-all-or-upgradable-or-patchable-greater-than) 참조

`--prune-repeated-subdependencies, -p`: [`--prune-repeated subdependencies` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p) 참조

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 lockfile을 테스트하는 것을 방지합니다.

프로젝트에 동기화가 맞지 않는 lockfile이 있다면, `--strict-out-of-sync=true`로 설정할 때 `test` 명령이 실패합니다.

기본: true

### `--yarn-workspaces`

루트에 lockfile이 있는 경우 Yarn Workspaces만 감지하고 스캔합니다.

`--detection-depth`를 사용하여 검색할 하위 디렉터리 수를 지정할 수 있습니다.

`--exclude`를 사용하여 디렉터리와 파일을 제외할 수 있습니다.

기본: `--all-projects`는 자동으로 Yarn Workspaces를 탐지하고 다른 프로젝트와 함께 스캔합니다.

## CocoaPods 프로젝트 옵션

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 lockfile을 테스트하는 것을 제어합니다.

기본: false

## Python 프로젝트 옵션

### `--command=<COMMAND>`

Python 버전을 기준으로 사용할 특정 Python 명령어를 표시합니다.

Snyk은 종속성을 찾기 위해 Python을 사용합니다. 여러 Python 버전을 사용하는 경우, 올바른 Python 명령어를 실행할 때 이 매개변수를 사용하세요.

기본: `python` 이는 기본 Python 버전을 실행합니다. 기본 버전이 무엇인지 확인하려면 `python -V`를 실행하세요.

예시: `snyk test --command=python3`

### `--skip-unresolved=true|false`

환경에서 찾을 수 없는 패키지를 건너뛰세요. 예를 들어, 스캔을 실행하는 기기로 액세스할 수 없는 비공개 패키지 등을 건너뜁니다.

### `--file=<FILE_for_Python>`

Python 프로젝트에서 테스트할 특정 파일을 지정하세요.

기본: Snyk은 프로젝트의 최상위에 있는 requirements.txt 파일을 스캔합니다.

**중요:** 기본 파일이 아닌 값을 `--file` 매개변수에 지정할 때는 반드시 `--package-manager=pip` 옵션을 함께 포함해야 합니다. 이 매개변수 없이는 테스트가 실패합니다.

사용자 지정 `--file` 값을 사용할 때는 항상 `--package-manager=pip` 값을 함께 명시하세요. 예를 들어:

```bash
snyk test --file=requirements-dev.txt --package-manager=pip
```

이는 Snyk가 명시된 매니페스트 파일을 올바르게 인식하고 스캔할 수 있도록 합니다. 예를 들어, 이름이 `requirements-dev.txt`로 변경된 경우에도 올바르게 스캔됩니다.

### `--package-manager=pip`

파일 이름이 `requirements.txt`가 아닌 경우 `--package-manager=pip`를 명령에 추가하세요.

`--file` 매개변수에 기본 `requirements.txt` 파일이 아닌 값을 지정하는 경우 이 옵션은 필수입니다. 이 매개변수 없이는 테스트가 실패합니다. `pip` 값을 사용하여 이 매개변수를 지정하세요.

명령에 대한 자세한 정보는 [`--package-manager=<PACKAGE_MANAGER_NAME>`](https://docs.snyk.io/snyk-cli/commands/test#package-manager-less-than-package_manager_name-greater-than)를 참조하세요.

## Go 프로젝트 옵션

다음 옵션은 지원되지 않습니다:

`--fail-on=<all|upgradable|patchable>`

## `--unmanaged`를 사용한 스캔 옵션

표준 `snyk test` 옵션은 이 도움말에 설명된 대로 `--unmanaged`와 함께 사용할 수 있습니다.

`--org=<ORG_ID>`

`--json`

`--json-file-output=<OUTPUT_FILE_PATH>`

`--remote-repo-url=<URL>`

`--severity-threshold=<low|medium|high|critical>`

다음과 같이 특별한 옵션이 있습니다.

### `--max-depth`

아카이브 추출의 최대 레벨을 지정하세요.

사용법: `--max-depth=1`

0(기본값인 제로)을 사용하여 아카이브 추출을 완전히 비활성화할 수 있습니다.

### `--print-dep-paths`

의존성을 표시합니다.

식별된 각 의존성이 어떤 파일에 기여했는지 보려면 이 옵션을 사용하세요.

식별된 의존성 및 해당 버전에 대한 Snyk의 신뢰도를 확인하려면 `--print-deps` 또는 `--print-dep-paths` 옵션을 사용하세요.

## 빌드 도구 옵션

### `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

빌드 도구(예: Gradle 또는 Maven)에 바로 전달할 추가 옵션(매개변수, 플래그)을 전달하려면 완전한 Snyk 명령 다음에 이중 대시(`--`)를 사용하세요.

형식은 `snyk <command> -- [<context-specific_options>]`입니다.

예시: `snyk test -- --build-cache`

**참고:** `-- [<context-specific_options>]`에서는 이중 인용 부호를 사용하지 마세요.

예시: `snyk test --org=myorg -- -s settings.xml`

`"s settings.xml"`을 사용하지 마세요.

## `snyk test` 명령에 대한 예시

현재 폴더에 있는 프로젝트를 알려진 취약점에 대해 테스트합니다:

`$ snyk test`

특정 의존성의 취약점을 테스트합니다 (**npm**만 해당):

`$ snyk test ionic@1.6.5`

최신 버전의 **npm** 패키지를 테스트합니다:

`$ snyk test lodash`

공개 GitHub 저장소를 테스트합니다:

`$ snyk test https://github.com/snyk-labs/nodejs-goof`
