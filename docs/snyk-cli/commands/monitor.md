# 모니터

## 사용법

`snyk monitor [<OPTIONS>]`

## 설명

`snyk monitor` 명령은 오픈 소스 취약점과 라이선스 문제를 지속적으로 모니터링하는 프로젝트를 Snyk 계정에 생성하여 그 결과를 [snyk.io](https://snyk.io) 로 전송합니다.

프로덕션에 프로젝트를 통합하기 전에 `monitor` 명령을 사용하여 코드 스냅샷을 취하고, 취약점을 프로덕션으로 전달하는 것을 피하기 위해 모니터링할 코드의 스냅샷을 취하십시오. 원하는 경우 기본값인 매일마다 변경하려면 설정에서 테스트 빈도를 선택하십시오.

PR 확인도 테스트를 수행합니다.

`snyk monitor` 명령을 실행한 후에 Snyk 웹사이트에 로그인하여 프로젝트를 보고 모니터를 확인하십시오.

코드를 변경하는 경우 `monitor` 명령을 다시 실행해야 합니다.

Snyk 의 경우, [`snyk container` 도움말](https://docs.snyk.io/snyk-cli/commands/container) 을 참조하십시오

`monitor` 명령은 Snyk 에서 지원되지 않습니다.

Snyk 를 사용하는 경우, [Snyk CLI for IaC](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac)의 "정기적으로 IaC 파일 테스트하기" 지침을 따르십시오.

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공, 스냅샷 생성됨\
**2**: 실패, 명령을 다시 실행하십시오. 디버그 로그를 출력하려면 `-d` 를 사용하십시오.\
**3**: 실패, 지원되는 프로젝트가 감지되지 않음

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와 연결할 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하십시오.

## 코드 실행 경고

코드를 스캔하기 전에 Snyk CLI에 대한 [코드 실행 경고](https://docs.snyk.io/snyk-cli/code-execution-warning-for-snyk-cli)를 검토하십시오.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 옵션

특정 빌드 환경, 패키지 관리자, 언어 및 `[<CONTEXT-SPECIFIC OPTIONS>]`에 대한 옵션을 확인하려면 이어지는 섹션을 참조하십시오.

### `--all-projects`

Yarn 작업 영역을 포함하여 작업 디렉토리의 모든 프로젝트를 자동으로 감지합니다.

자세한 정보는 [Snyk CLI가 monorepo 또는 여러 manifest 파일을 지원하는가?](https://support.snyk.io/s/article/Does-the-Snyk-CLI-support-monorepos-or-multiple-manifest-files) 문서를 참조하십시오.

### `--fail-fast`

`--all-projects`와 함께 사용하면 오류 발생 시 스캔이 중단되고 이러한 오류가 사용자에게 보고됩니다.

종료 코드는 2이며 스캔이 종료됩니다. 오류가 발생하지 않은 프로젝트에 대해 취약성 정보가 보고되지 않습니다.

스캔을 수행하려면 오류를 해결하고 다시 스캔하십시오.

**참고**: `--fail-fast`를 사용하지 않으면 Snyk는 모든 프로젝트를 스캔하지만 구성 오류 또는 기타 오류로 인해 스캔할 수 없는 프로젝트에 대해 어떠한 취약점도 보고하지 않습니다.

### `--detection-depth=<DEPTH>`

`--all-projects` 또는 `--yarn-workspaces`와 함께 사용하여 검색할 서브디렉터리의 깊이를 지정합니다. `DEPTH`는 숫자여야 하며 1 이상이어야 합니다. 0은 현재 디렉터리입니다.

기본값: 제한 없음

예: `--detection-depth=3`은 지정된 디렉터리(또는 `<PATH>`가 지정되지 않은 경우 현재 디렉터리) 및 서브디렉터리 세 단계에 대해 검색을 제한합니다. 0은 현재 디렉터리입니다.

### `--exclude=<NAME>[,<NAME>]...>`

`--all-projects` 및 `--yarn-workspaces`와 함께 사용하여 제외할 디렉터리 이름 및 파일 이름을 나열합니다. 쉼표로 구분하며 경로를 포함할 수 없습니다.

예: `$ snyk test --all-projects --exclude=dir1,file2`

이렇게 하면 `./dir1`, `./src/dir1`, `./file2`, `./src/file2` 등과 같은 프로젝트 매니페스트 파일을 스캔할 때 `dir1` 및 `file2`라는 디렉터리 및 파일이 제외됩니다.

**참고**: `--exclude=dir1`은 `./dir1` 및 `./src/dir1`을 찾습니다.\
그러나 `--exclude=./src/dir1`은 경로를 포함하므로 오류가 발생합니다.

### `--prune-repeated-subdependencies`, `-p`

중복 서브 의존성을 제거하여 종속성 트리를 정리합니다.

모든 취약성은 계속해서 찾지만 모든 취약 경로를 찾지 못할 수 있습니다.

대형 프로젝트가 테스트에 실패하는 경우 이 옵션을 사용하십시오.

기본값: false

### `--print-deps`

분석 전에 종속성 트리를 출력합니다.

### `--remote-repo-url=<URL>`

모니터링하려는 리포지토리를 설정하거나 오버라이드합니다.

단일 대상 아래에서 모든 찾은 프로젝트를 그룹화합니다.

### `--dev`

개발 전용 종속성을 포함합니다. 일부 패키지 관리자(예: npm의 devDependencies 또는 Gemfile의 :development 종속성)에만 해당됩니다.

**참고**: 이 옵션은 Maven, npm 및 Yarn 프로젝트에서만 사용할 수 있습니다.

기본값: false, 프로덕션 종속성만 스캔함.

### `--org=<ORG_ID>`

특정 Snyk 조직과 관련된 Snyk 명령을 실행하기 위해 `<ORG_ID>`를 지정합니다. `<ORG_ID>`는 일부 기능 가용성 및 개인 테스트 제한에 영향을 미칩니다.

여러 조직이 있는 경우 CLI에서 기본값을 설정할 수 있습니다:

`$ snyk config set org=<ORG_ID>`

기본값을 설정하여 모든 새롭게 모니터링된 프로젝트가 기본 조직 하에 생성되도록 합니다. 기본값을 무시해야 하는 경우 `--org=<ORG_ID>` 옵션을 사용하십시오.

기본값: [계정 설정](https://app.snyk.io/account)에 나타난 현재 우선 조직의 `<ORG_ID>`

**참고:** `--org=<orgslugname>.` `ORG_ID`는 CLI와 API 모두에서 작동합니다. 조직 슬러그 이름은 CLI에서는 작동하지만 API에서는 작동하지 않습니다.

`orgslugname`은 UI의 조직 URL에 표시된 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`. 조직 이름은 작동하지 않습니다.

자세한 내용은 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) 문서를 참조하십시오.

### `--file=<FILE>`

패키지 파일을 지정합니다.

로컬에서 테스트하거나 프로젝트를 모니터링할 때 Snyk가 패키지 정보를 검사해야 하는 파일을 지정할 수 있습니다. 파일이 지정되지 않으면 Snyk는 프로젝트에 적합한 파일을 자동으로 감지하려고 시도합니다.

Python 프로젝트에 대한 자세한 내용은 [Python 프로젝트용 옵션](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-python-projects) 섹션을 참조하십시오. 

### `--package-manager=<PACKAGE_MANAGER_NAME>`

`--file=<FILE>` 옵션과 함께 기본이 아닌 파일을 지정할 때 패키지 관리자의 이름을 지정합니다. 이렇게 하면 Snyk가 파일을 찾을 수 있습니다.

예: `$ snyk monitor --file=req.txt --package-manager=pip`

자세한 내용은 [Python 프로젝트용 옵션](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-python-projects) 문서를 참조하십시오.

### `--unmanaged`

C++ 전용, 알려진 오픈 소스 종속성을 가진 모든 파일을 스캔합니다.

`--unmanaged`로 사용할 수 있는 옵션에 대한 자세한 내용은 [`--unmanaged`를 사용하여 스캔하는 옵션](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-scanning-using-unmanaged)을 참조하십시오.

### `--ignore-policy`

모든 정책 설정, `.snyk` 파일의 현재 정책, 조직 레벨 무시 및 snyk.io의 프로젝트 정책을 무시합니다.

### `--trust-policies`

의존성에서 무시 규칙을 적용 및 사용하십시오; 그렇지 않으면 의존성의 무시 규칙이 제안으로만 표시됩니다.

### `--project-name=<PROJECT_NAME>`

사용자 정의 Snyk 프로젝트 이름을 지정합니다.

예: `$ snyk monitor --project-name=my-project`

### `--target-reference=<TARGET_REFERENCE>`

이 프로젝트를 구분하는 참조를 지정합니다. 예를 들어 브랜치 이름이나 버전입니다. 동일한 참조를 가진 프로젝트는 해당 참조를 기준으로 그룹화됩니다. Snyk Open Source에서 지원하고 `--unmanaged` 와 함께 사용합니다.

프로젝트의 그룹화에 대해 자세히 알아보려면 [모니터링을 위해 브랜치 또는 버전으로 프로젝트 그룹화](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring) 문서를 참조하십시오.

### `--policy-path=<PATH_TO_POLICY_FILE>`

`.snyk` 정책 파일의 경로를 수동으로 전달합니다.

### `--json`

콘솔에 결과를 JSON 데이터 구조로 출력합니다.

**참고**: 프로젝트 속성을 설정하는 옵션을 사용하고 권한이 없는 역할을 가진 경우 `monitor` 명령이 실패합니다. 계속하기 위한 지침은 [Snyk CLI에서 프로젝트 속성을 편집하려면 필요한 권한(역할)](https://docs.snyk.io/snyk-admin/manage-permissions-and-roles/manage-member-roles#permissions-role-required-to-edit-project-attributes-from-the-snyk-cli) 문서를 참조하십시오.

### `--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`

프로젝트 환경 프로젝트 속성을 하나 이상의 값으로 설정할 수 있습니다(쉼표로 구분). 프로젝트 환경을 지우려면 `--project-environment=`을 설정합니다.

허용되는 값: `frontend, backend, internal, external, mobile, saas, onprem, hosted, distributed`

자세한 내용은 [프로젝트 속성](https://docs.snyk.io/getting-started/introduction-to-snyk-projects/view-project-information/project-attributes) 문서를 참조하십시오.

### `--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...>`

프로젝트 수명주기 프로젝트 속성을 하나 이상의 값으로 설정할 수 있습니다(쉼표로 구분). 프로젝트 수명주기를 지우려면 `--project-lifecycle=`를 설정합니다.

허용되는 값: `production, development, sandbox`

자세한 내용은 [프로젝트 속성](https://docs.snyk.io/snyk-admin/snyk-projects/project-tags) 문서를 참조하십시오.

### `--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`

프로젝트 중요도 프로젝트 속성을 하나 이상의 값으로 설정할 수 있습니다(쉼표로 구분). 프로젝트 중요도를 지우려면 `--project-business-criticality=`를 설정합니다.

허용되는 값: `critical, high, medium, low`

자세한 내용은 [프로젝트 속성](https://docs.snyk.io/snyk-admin/snyk-projects/project-tags) 문서를 참조하십시오.

### `--project-tags=<TAG>[,<TAG>]...>`

프로젝트 태그를 하나 이상의 값(키 값 쌍으로 쉼표로 구분)으로 설정합니다.

예: `--project-tags=department=finance,team=alpha`

프로젝트 태그를 지우려면 `--project-tags=`를 설정합니다.

허용되는 문자 및 자세한 내용은 [프로젝트 태그](https://docs.snyk.io/snyk-admin/snyk-projects/project-tags) 문서를 참조하십시오.

### `--tags=<TAG>[,<TAG>]...>`

`--project-tags`의 별칭입니다.

## Maven 프로젝트용 옵션

**참고**: `--dev` 옵션은 Maven 프로젝트와 함께 사용할 수 있습니다. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#dev)도 확인하십시오.

### `--maven-aggregate-project`

모듈 및 상속을 사용# 옵션의 스캔 대상

이 옵션과 함께 지정된 대상 Framework는 표준 [명명 규칙](https://learn.microsoft.com/en-us/dotnet/standard/frameworks#supported-target-frameworks)을 따라 정의되어야 합니다.

예: `snyk test --dotnet-runtime-resolution --dotnet-target-framework=net6.0`

## npm 프로젝트용 옵션

**참고**: npm 프로젝트에 다음 옵션을 사용할 수 있습니다.

`--dev`. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#dev) 참조

`--all-projects`. 디렉토리의 모든 npm 프로젝트 및 다른 모든 프로젝트를 스캔 및 감지합니다. [`--all-projects` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#all-projects) 참조

`--prune-repeated-subdependencies, -p`. [--prune-repeated subdependencies 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#prune-repeated-subdependencies-p) 참조

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 락 파일 모니터링을 제어합니다.

기본값: true

## pnpm 프로젝트용 옵션

**Snyk CLI pnpm 지원은 초기 액세스입니다**. 활성화하려면 Snyk 계정에서 설정으로 이동한 다음 Snyk 미리 보기를 선택하고 CLI 버전 1.1293.0 이상을 설치하십시오.

**참고**: pnpm 프로젝트에 다음 옵션을 사용할 수 있습니다.

`--dev`. [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#dev) 참조

`--all-projects`. 디렉토리의 모든 pnpm 프로젝트 및 다른 모든 프로젝트를 스캔 및 감지합니다. [`--all-projects` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#all-projects) 참조

`--prune-repeated-subdependencies, -p`. [--prune-repeated subdependencies 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#prune-repeated-subdependencies-p) 참조

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 락 파일 모니터링을 제어합니다.

기본값: true

## Yarn 프로젝트용 옵션

**참고**: Yarn 프로젝트에 다음 옵션을 사용할 수 있습니다.

`--dev.` [`--dev` 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#dev) 참조

`--prune-repeated-subdependencies, -p.` [--prune-repeated subdependencies 옵션 도움말](https://docs.snyk.io/snyk-cli/commands/monitor#prune-repeated-subdependencies-p) 참조

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 락 파일 모니터링을 제어합니다.

기본값: true

### `--yarn-workspaces`

루트에 락 파일이있을 때 Yarn Workspaces만 감지하고 스캔합니다.

`--detection-depth`를 사용하여 검색할 하위 디렉토리 수를 지정할 수 있습니다.

`--exclude`를 사용하여 디렉토리 및 파일을 제외할 수 있습니다.

기본 값 : `--all-projects`는 다른 프로젝트와 함께 Yarn Workspaces를 자동으로 감지하고 스캔합니다.

## CocoaPods 프로젝트용 옵션

### `--strict-out-of-sync=true|false`

동기화가 맞지 않는 락 파일 모니터링을 제어합니다.

기본값: false

## Python 프로젝트용 옵션

### `--command=<COMMAND>`

Python 버전을 기반으로 사용할 특정 Python 명령을 지정합니다.

Snyk는 종속성을 찾기 위해 Python을 사용합니다. 여러 Python 버전을 사용하는 경우 이 매개변수를 사용하여 실행에 올바른 Python 명령을 지정하십시오.

기본값: `python`은 기본 python 버전을 실행합니다. 기본 버전을 확인하려면 `python -V`를 실행하십시오.

예: `snyk monitor --command=python3`

### `--skip-unresolved=true|false`

환경에서 찾을 수 없는 패키지를 건너뛰십시오. 예를 들어 스캔을 실행하는 기계에서 액세스할 수 없는 개인 패키지를 건너뛰세요.

### `--file=<Python_용_FILE>`

Python 프로젝트에서 모니터링할 특정 파일을 지정하십시오.

기본값: Snyk는 프로젝트의 최상위에 있는 requirements.txt 파일을 스캔합니다.

**중요:** 기본 파일이 아닌 `--file` 매개변수 값을 지정할 때는 `--package-manager=pip` 옵션을 포함해야합니다. 이 매개변수가 없으면 테스트가 실패합니다.

사용자 지정 `--file` 값을 사용할 때 항상 `pip` 값을 사용하여 매개변수를 지정하십시오. 예:

```bash
snyk test --file=requirements-dev.txt --package-manager=pip
```

이렇게 하면 Snyk가 지정한 manifest 파일을 올바르게 인식하고 스캔할 수 있으며, 이름을 `requirements-dev.txt`로 변경한 경우와 같이 파일을 올바르게 인식합니다.

### `--package-manager=pip`

파일 이름이 `requirements.txt`가 아닌 경우 명령에`--package-manager=pip`를 추가하십시오.

기본 `requirements.txt` 파일이 아닌 값을 `--file` 매개변수로 지정하는 경우 이 옵션은 필수입니다. 이 매개변수가 없으면 테스트가 실패합니다. 이 매개변수를`pip` 값과 함께 지정하십시오.

명령의 자세한 정보는 [`--package-manager=<PACKAGE_MANAGER_NAME>`](https://docs.snyk.io/snyk-cli/commands/monitor#package-manager-less-than-package_manager_name-greater-than)를 참조하십시오.

## `--unmanaged`를 사용한 스캔 옵션

다음 `snyk monitor` 옵션은 이 도움말에 기록된대로`--unmanaged`와 함께 사용할 수 있습니다.

`--org=<ORG_ID>`

`--json`

`--remote-repo-url=<URL>`

`--target-reference=<TARGET_REFERENCE>`

`--project-name=<c-project>`

특별한 옵션들도 있습니다.

### `--max-depth`

아카이브 추출의 최대 수준을 지정합니다.

사용법: `--max-depth=1`

아카이브 추출을 완전히 비활성화하려면 0 (제로, 기본값)을 사용하십시오.

### `--print-dep-paths`

의존성을 표시합니다.

각 식별된 의존성에 기여한 파일을 확인하려면 이 옵션을 사용하십시오.

식별된 의존성 및 해당 버전에 대한 Snyk의 신뢰도를 확인하려면 `--print-deps` 또는 `--print-dep-paths` 옵션을 사용하십시오.

## 빌드 도구용 옵션

### `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

빌드 도구(예: Gradle 또는 Maven)에 직접 전달할 옵션(인수, 플래그)을 전달하려면 완전한 Snyk 명령 뒤에 이중 대시(`--`)를 사용하십시오.

형식은 `snyk <command> -- [<context-specific_options>]`입니다.

예: `snyk monitor -- --build-cache`

**참고:** 모든 `-- [<context-specific_options>]`에 이중 인용 부호를 사용하지 마십시오.

예: `snyk monitor --org=myorg -- -s settings.xml`를 `snyk monitor --org=myorg -- "-s settings.xml"` 대신 사용합니다.