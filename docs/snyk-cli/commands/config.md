# 구성

## 사용법

`snyk config <하위명령> [<옵션>]`

## 설명

`snyk config` 명령은 로컬 Snyk CLI 구성 파일을 관리합니다. 이 파일은 `$XDG_CONFIG_HOME` 또는 `~/.config` 다음에 `configstore/snyk.json`에 위치한 JSON 파일입니다.

예: `~/.config/configstore/snyk.json`

이 명령은 프로젝트의 일부인 `.snyk` 파일을 관리하지 않습니다. [`snyk policy` ](policy.md) 및 [`snyk ignore`](ignore.md) 명령을 참조하십시오.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 하위명령

### `get <키>`

구성 값 출력.

### `set <키>=<값>`

새 구성 값 생성.

### `unset <키>`

구성 값 제거.

### `clear`

모든 구성 값 제거.

### `environment`

사용할 엔드포인트 변경. `config environment --help`을 실행하거나 [Config environment 도움말 페이지](https://docs.snyk.io/snyk-cli/commands/config-environment)를 참조하십시오.

## 지원되는 `<키>` 값

### `api`

Snyk API 호출 시 사용할 API 토큰.

### `endpoint`

사용할 API 엔드포인트 정의.

### `disable-analytics`

분석 보고를 끕니다.

### `org`

특정 Snyk 조직에 연결된 Snyk 명령을 실행하기 위한 `<ORG_ID>` 지정.

### `oci-registry-url`

사용자 정의 규칙으로 IaC 스캔 시 사용하는 OCI 레지스트리 구성.

### `oci-registry-username`

사용자 정의 규칙으로 IaC 스캔 시 사용하는 OCI 레지스트리의 사용자 이름 구성.

### `oci-registry-password`

사용자 정의 규칙으로 IaC 스캔 시 사용하는 OCI 레지스트리의 암호 구성.