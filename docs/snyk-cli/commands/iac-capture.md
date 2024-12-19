# IaC 캡처

## 사용법

**기능 가용성:** 이 기능은 Snyk CLI 버전 v1.1117.0 이상에서 이용 가능합니다.

`snyk iac capture [<옵션>] [<경로>]`

## 설명

`snyk iac capture` 명령은 Terraform 상태 파일로부터 리소스 ID 및 이름과 같은 필수 정보를 포함하는 매핑 아티팩트를 생성하여 코드에서 클라우드로 리소스 매핑을 생성하는 데 필요한 최소한의 정보를 포함하며, 해당 매핑 아티팩트를 Snyk로 전송합니다.

Snyk는 이 정보를 사용하여 클라우드 문제와 해당 IaC 파일을 연결합니다. 이 링크는 Snyk 웹 UI에서 볼 수 있습니다.

자세한 정보는 [IaC에서 클라우드 문제 해결하기](https://docs.snyk.io/scan-using-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/fix-cloud-issues-in-iac)를 참조하십시오.

관련 명령어 목록은 [snyk iac 도움말](iac.md)을 참조하세요; `iac --help`

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공\
**1**: 하나 이상의 Terraform 상태를 캡처하는 데 실패함

## Snyk CLI 구성

Snyk API와 연결하기 위해 환경 변수 및 변수를 설정할 수 있습니다; [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.

## 옵션

### `--org=<ORG_ID>`

특정 Snyk 조직에 연결된 Snyk 명령을 실행할 `<ORG_ID>`를 지정합니다. 현재 기본적으로 설정된 기본 `<ORG_ID>` (이곳에서 설정하는 현재 선호하는 조직)을 재정의합니다. 자세한 내용은 [계정 설정](https://app.snyk.io/account)에서 확인할 수 있습니다.

`--org=<orgslugname>`도 사용할 수 있습니다. `ORG_ID`는 CLI와 API에서 모두 작동하며, 조직 슬러그 이름은 CLI에서만 작동하며 API에서는 작동하지 않습니다.

자세한 정보는 [CLI에서 사용할 조직 선택 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) 문서를 확인하십시오.

### `--stdin`

`PATH`에서 상태 파일을 찾는 대신 표준 입력에서 Terraform 상태로부터 매핑 아티팩트를 생성합니다.

```
$ terraform state pull | snyk iac capture --stdin
```

### `PATH`

`PATH`에있는 Terraform 상태 파일에서 매핑 아티팩트를 생성하는 선택적 인수입니다. 디렉토리 경로, 파일 경로 또는 글로브 패턴이 될 수 있습니다.

```
$ snyk iac capture /path/to/states/**/*.tfstate
```

## snyk iac capture 명령어 예제

### 현재 작업 디렉토리의 모든 상태에서 캡처

```
$ snyk iac capture
```

### 디렉토리에 있는 .tfstate로 끝나는 모든 파일에서 캡처

```
$ snyk iac capture /path/to/states/**/*.tfstate
```

### 단일 상태 파일에서 캡처

```
$ snyk iac capture /path/to/state.tfstate
```

### 표준 입력에서 Terraform으로 가져온 상태에서 캡처

```
$ terraform state pull | snyk iac capture --stdin
```