# 코드 테스트

## 사용법

`snyk code test [<OPTIONS>] [<PATH>]`

## 설명

`snyk code test` 명령은 소스 코드를 알려진 보안 취약점(정적 응용 프로그램 보안 테스팅)에 대해 테스트합니다.

## 종료 코드

가능한 종료 코드와 그 의미:

**0**: 성공 (스캔 완료), 취약점 발견되지 않음\
**1**: 조치 필요 (스캔 완료), 취약점 발견됨\
**2**: 실패, 명령을 다시 실행하십시오. 디버그 로그를 출력하려면 `-d`를 사용하십시오.\
**3**: 실패, 지원되는 프로젝트가 감지되지 않음

## Snyk CLI 구성

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 옵션

### `--org=<ORG_ID>`

`<ORG_ID>`를 지정하여 특정 Snyk 조직과 관련된 Snyk 명령을 실행합니다. `<ORG_ID>`는 개인 테스트 제한에 영향을 미칩니다.

여러 조직이 있는 경우 CLI에서 다음을 사용하여 기본값을 설정할 수 있습니다:

`$ snyk config set org=<ORG_ID>`

기본값을 설정하여 모든 새로 테스트된 프로젝트가 기본 조직을 사용하여 테스트되도록 합니다. 기본값을 재정의해야 하는 경우 `--org=<ORG_ID>` 옵션을 사용하십시오.

기본값: 현재 [계정 설정](https://app.snyk.io/account)에서 현재 선호하는 조직인 `<ORG_ID>`

**참고:** `--org=<orgslugname>`을 사용할 수도 있습니다.`ORG_ID`는 CLI와 API에서 모두 작동하며, 조직 슬러그 이름은 CLI에서는 작동하지만 API에서는 작동하지 않습니다.

`orgslugname`은 Snyk UI의 URL에 표시되는 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`. 조직 이름은 작동하지 않습니다.

자세한 내용은 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) 문서를 참조하십시오.

### `--json`

콘솔에 결과를 JSON 데이터 구조로 출력합니다.

예: `$ snyk code test --json`

### `--json-file-output=<OUTPUT_FILE_PATH>`

`--json` 옵션을 사용하든 말든 지정된 파일로 테스트 출력을 JSON 데이터 구조로 직접 저장합니다.

인간이 읽을 수 있는 테스트 출력을 stdout에서 표시하고 동시에 JSON 데이터 구조를 파일에 저장하는 데 사용합니다.

SAST의 경우 Snyk에서 문제가 발견되지 않으면 `json` 파일을 생성하지 않습니다. 반면에 오픈 소스의 경우 문제가 발견되더라도 Snyk은 파일을 생성합니다.

예: `$ snyk code test --json-file-output=vuln.json`

### `--sarif`

결과를 SARIF 형식으로 반환합니다.

예: `$ snyk code test --sarif`

### `--sarif-file-output=<OUTPUT_FILE_PATH>`

`--sarif` 옵션을 사용하든 말든 지정된 파일로 SARIF 형식으로 테스트 출력을 직접 저장합니다.

인간이 읽을 수 있는 테스트 출력을 stdout에서 표시하고 동시에 SARIF 형식 출력을 파일에 저장하는 데 사용합니다.

### `--severity-threshold=<low|medium|high>`

지정된 수준 또는 그 이상의 취약점만 보고합니다.

**참고:**  구성 문제에서는 `critical` 심각도 수준을 사용하지 않습니다.