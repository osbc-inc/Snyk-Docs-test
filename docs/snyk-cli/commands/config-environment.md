# 환경 설정

**참고:** 이 명령은 CLI 버전 1.1293.0에서 사용할 수 있게 됩니다.

시스템 기본 환경 SNYK-US-01이 아닌 경우 `snyk auth`를 실행하기 전에 환경을 설정하려면 `snyk config environment` 명령을 사용하십시오.

## 사용법

`snyk config environment <환경>`

## 설명

`snyk config environment` 명령은 CLI에서 사용하는 엔드포인트를 변경하는 것을 더 쉽고 안전하게 만들기 위해 제공됩니다.

결과는 `snyk config set endpoint=<URL>`과 거의 동일하지만 추가로 `snyk config environment` 명령은 다음을 수행합니다:

* 전체 URL 사용을 피하기 위해 환경의 별칭 지원
* 모호하거나 예기치 않은 구성을 피하기 위한 기본적인 확인 수행
* 환경별로 예상되는 인증 및 조직 설정을 지운다

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 옵션

`--no-check`

모호하거나 예기치 않은 구성에 대한 기본 확인을 건너뛰세요.

보고된 구성 문제가 의도적이고 무시할 수 있는 경우에만 사용하십시오. 그렇지 않으면 구성이 예상대로 작동하지 않을 수 있습니다.

## 지원되는 환경 URL 매핑

* default => https://api.snyk.io&#x20;
* SNYK-US-01 => https://api.snyk.io&#x20;
* SNYK-US-02 => https://api.us.snyk.io&#x20;
* SNYK-AU-01 => https://api.au.snyk.io&#x20;
* SNYK-EU-01 => https://api.eu.snyk.io&#x20;
* SNYK-GOV-01 => https://api.snykgov.io

## 예시

```
snyk config environment default
snyk config environment SNYK-EU-01
snyk config environment SNYK-AU-01
snyk config environment SNYK-AU-01 --no-check
snyk config environment https://api.eu.snyk.io
```