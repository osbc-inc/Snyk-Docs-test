# IaC 규칙 테스트

## 사용법

**기능 가용성:** 이 기능은 초기 액세스 상태입니다.

`snyk iac rules test [<OPTIONS>]`

## 설명

`snyk iac rules test` 명령은 Rego로 작성된 모든 테스트를 실행합니다.

관련 명령어 목록은 `snyk iac --help`를 실행합니다.

## Snyk CLI 구성

환경 변수 및 Snyk API와 연결하기 위해 변수를 설정할 수 있습니다. Snyk CLI 구성은 [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)에서 확인할 수 있습니다.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.

## 옵션

특정 빌드 환경, 패키지 관리자, 언어에 대한 옵션 및 마지막으로 지정하는 `[<CONTEXT-SPECIFIC OPTIONS>]`에 대한 옵션은 이후 섹션도 참조하세요.

### `--update-expected`

테스트가 올바르게 작성되었다고 가정하여 다음 실행에서 테스트가 통과하도록 기대되는 스펙 파일을 덮어씁니다.

## snyk iac rules test 명령의 예시

**모든 규칙에 대한 테스트 실행**

```
snyk iac rules test
```