# Container test

## 사용법

`snyk container test [<OPTIONS>] [<IMAGE>]`

## 설명

`snyk container test` 명령은 알려진 취약점을 가진 컨테이너 이미지를 테스트합니다.

## 종료 코드

가능한 종료 코드와 그 의미는 다음과 같습니다:

**0**: 성공 (스캔 완료), 취약점 없음\
**1**: 조치 필요 (스캔 완료), 취약점 발견\
**2**: 실패, 명령을 다시 실행하십시오. 디버그 로그를 출력하려면 `-d`를 사용하십시오.\
**3**: 실패, 지원되는 프로젝트를 감지할 수 없음

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와 연결하는 데 필요한 변수를 설정할 수 있습니다.

컨테이너 명령에 적용되는 환경 변수가 있습니다. [Snyk CLI 구성](https://docs.snyk.io/features/snyk-cli/configure-the-snyk-cli) 를 참조하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 옵션

### `--print-deps`

분석을 보내기 전에 종속성 트리를 출력합니다.

### `--org=<ORG_ID>`

특정 Snyk 조직에 속한 Snyk 명령을 실행하려면 `<ORG_ID>`를 지정하십시오. `<ORG_ID>`는 일부 기능의 가용성 및 개인 테스트 한도에 영향을 미칩니다.

여러 조직을 보유하고 있는 경우, 다음을 사용하여 CLI에서 기본값을 설정할 수 있습니다:

`$ snyk config set org=<ORG_ID>`

기본값을 설정하여 새로 테스트하고 모니터링하는 프로젝트가 기본 조직에서 테스트 및 모니터링되도록 합니다. 기본값을 재정의해야 하는 경우 `--org=<ORG_ID>` 옵션을 사용할 수 있습니다.

기본값: 계정 설정에서 현재 기본 조직인 `<ORG_ID>`

**참고:** `--org=<orgslugname>` 또한 사용할 수 있습니다. `ORG_ID`는 CLI 및 API에서 모두 작동합니다. 조직 슬러그 이름은 CLI에서는 작동하지만 API에서는 작동하지 않습니다.

`orgslugname`은 Snyk UI의 URL에 표시된 조직 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`. orgname은 작동하지 않습니다.

자세한 정보는 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) 문서를 참조하세요.

### `--file=<FILE_PATH>`

이미지에 대한 Dockerfile 경로를 포함하여 더 상세한 정보를 얻으려면 지정하십시오.

### `--project-name=<PROJECT_NAME>`

사용자 정의 Snyk 프로젝트 이름을 지정합니다.

### `--policy-path=<PATH_TO_POLICY_FILE>`

`.snyk` 정책 파일의 경로를 수동으로 전달합니다.

### `--json`

콘솔에 결과를 JSON 데이터 구조로 출력합니다.

예: `$ snyk container test --json`

### `--json-file-output=<OUTPUT_FILE_PATH>`

`--json` 옵션을 사용하든 말든, 결과를 JSON 데이터 구조로 지정된 파일에 직접 저장합니다.

JSON 데이터 구조를 파일에 저장함과 동시에 인간이 읽을 수 있는 테스트 결과를 stdout에 표시하려면 사용합니다.

예: `$ snyk container test --json-file-output=vuln.json`

### `--sarif`

SARIF 형식으로 결과를 반환합니다.

### `--sarif-file-output=<OUTPUT_FILE_PATH>`

`--sarif` 옵션을 사용하든 말든, 결과를 SARIF 형식으로 지정된 `<OUTPUT_FILE_PATH>` 파일에 직접 저장합니다.

stdout를 통해 인간이 읽을 수 있는 테스트 결과를 표시하고 동시에 SARIF 형식 결과를 파일에 저장하려면 사용합니다.

### `--severity-threshold=<low|medium|high|critical>`

지정된 수준 이상의 취약점만 보고합니다.

### `--fail-on=<all|upgradable>`

수정 가능한 취약점이 있는 경우에만 실패합니다.

* `all`: 업그레이드나 패치할 수 있는 취약점이 하나 이상 있는 경우 실패합니다.
* `upgradable`: Snyk이 계산된 복구 가능한 원순결을 가지고 있는 취약점이 하나 이상 있는 경우 실패합니다.

Snyk이 `snyk test`로 준수하지 않을 수 있는 메타데이터 제약을 확인하면 고쳐도 코드가 깨지지 않도록 회피하려고 합니다. 수동으로 고칠 수 있는 곳을 식별하고 적용할 수 있을 수 있습니다.

### `--app-vulns`

컨테이너 이미지의 애플리케이션 종속성과 운영 체제에서 취약점을 모두 한 번에 스캔합니다.

1.1090.0 (2023-01-24) 버전 이상의 CLI 버전에서는 Snyk이 이미지의 애플리케이션 종속성을 기본적으로 스캔합니다. `--app-vulns` 플래그를 지정할 필요가 없습니다.

1.962.0부터 v1.1089.0까지의 CLI 버전에서는 운영 체제 및 애플리케이션 취약점을 JSON 형식으로 결과에서 볼 수 있게 하려면 `--app-vulns` 옵션을 사용하세요.

### `--exclude-app-vulns`

애플리케이션 취약점 스캔을 비활성화합니다; 1.1090.0 (2023-01-24) 버전 이상의 경우 `app-vulns`가 기본적으로 활성화됩니다.

이전 릴리스에서는 `app-vulns`와 함께 사용할 수 없습니다.

### `--nested-jars-depth`

`app-vulns`를 활성화했을 때, `--nested-jars-depth=n` 옵션을 사용하여 Snyk이 얼마나 많은 수준의 중첩된 jar를 풀어야 하는지 설정합니다. 깊이는 숫자여야 합니다.

### `--exclude-base-image-vulns`

베이스 이미지에 의해 도입된 취약점만 표시하지 않습니다. `snyk container test`에서만 작동합니다. 운영 체제 패키지에만 적용됩니다.

### `--platform=<PLATFORM>`

다중 아키텍처 이미지의 경우 테스트할 플랫폼을 지정합니다.

지원되는 플랫폼: `linux/amd64`, `linux/arm64`, `linux/riscv64`, `linux/ppc64le`, `linux/s390x`, `linux/386`, `linux/arm/v7`, 또는 `linux/arm/v6`

### `--username=<CONTAINER_REGISTRY_USERNAME>`

컨테이너 레지스트리에 연결할 때 사용할 사용자 이름을 지정합니다. Docker가 설치되어 있는 경우 로컬 Docker 이진 자격 증명을 우선시합니다.

### `--password=<CONTAINER_REGISTRY_PASSWORD>`

컨테이너 레지스트리에 연결할 때 사용할 암호를 지정합니다. Docker가 설치되어 있는 경우 로컬 Docker 이진 자격 증명을 우선시합니다.

## 컨테이너 테스트 명령 예제

### Docker 이미지 스캔

`$ snyk container test <image>`

### 베이스 이미지 복구를 포함한 자세한 정보를 얻기 위한 옵션

`--file=path/to/Dockerfile`

### Dockerfile을 사용하여 생성된 Docker 이미지와 지정된 정책 경로를 사용하여 Docker 이미지 스캔

`$ snyk container test app:latest --file=Dockerfile`

`$ snyk container test app:latest --file=Dockerfile --policy-path=path/to/.snyk`

### 다이제스트를 사용하여 컨테이너 이미지 참조

`$ snyk container test app@sha256:17cb37098f0efb819c075eea4ff2a495be909a396e86ece317a6e3a8968e025c --file=Dockerfile`

더 많은 정보 및 예제를 보려면 [고급 CLI 사용법](https://docs.snyk.io/snyk-container/snyk-cli-for-container-security/advanced-snyk-container-cli-usage)를 참조하세요.

또한 [컨테이너 이미지에서 애플리케이션 취약점 감지하기](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/detect-application-vulnerabilities-in-container-images)도 참조하세요.
