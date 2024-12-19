# IaC 규칙 푸시

## 사용법

**기능 가용성:** 이 기능은 초기 액세스 상태입니다.

`snyk iac rules push [<OPTIONS>]`

## 설명

`snyk iac rules push` 명령은 Rego로 작성된 규칙을 번들로 묶어서 변경 사항을 Snyk 플랫폼에 업로드합니다.

관련 명령어 목록을 보려면 `snyk iac --help`를 실행하세요.

## Snyk CLI 구성

Snyk API에 연결하는 데 환경 변수와 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.

## 옵션

### `--delete`

이전에 푸시된 번들을 조직에서 삭제합니다. 조직은 `--org` 플래그를 사용하여 수동으로 재정의할 수 있습니다.

## snyk iac rules push 명령어 예제

### 새 번들 번들링 및 업로드

```
$ snyk iac rules push
```

### 번들 삭제

```
$ snyk iac rules push --delete
```