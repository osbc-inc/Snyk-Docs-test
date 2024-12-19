# IaC update-exclude-policy

## 사용법

`snyk iac update-exclude-policy [<OPTIONS>]`

## 설명

`snyk iac update-exclude-policy`는 `snyk iac describe`에서 사용할 exclude 정책 규칙을 생성합니다.

관련 명령어 목록은 [snyk iac 도움말](iac.md)을 참조하세요; `iac --help`

자세한 내용은 [리소스 무시](https://docs.snyk.io/products/snyk-infrastructure-as-code/detect-drift-and-manually-created-resources/ignore-resources)를 참조하세요.

## 종료 코드

가능한 종료 코드와 그 의미:

**0**: 성공, exclude 규칙이 성공적으로 생성됨\
**1**: 오류, exclude 규칙 생성 중에 문제 발생

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와 연결할 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.

## 옵션

### `--exclude-missing`

누락된 리소스를 제외합니다.

### `--exclude-unmanaged`

IaC에서 관리되지 않는 리소스를 제외합니다.

## 예제

```
$ snyk iac describe --json | snyk iac update-exclude-policy
```