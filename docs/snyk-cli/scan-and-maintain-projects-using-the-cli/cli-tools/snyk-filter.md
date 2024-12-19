# `snyk-filter`

`Snyk` CLI 출력에 대한 **사용자 정의 필터링을 제공**하는 `snyk-filter` 도구입니다. `snyk-filter`는 예를 들어 `snyk test --json`과 같이 [Snyk CLI](../../)에서 제공하는 JSON 형식의 출력을 사용하여 결과물의 사용자 정의 필터링 및 빌드 실패 옵션을 적용합니다.

## `snyk-filter` 클론 및 설치

먼저 [repo](https://github.com/snyk-labs/snyk-filter)를 클론한 다음 다음 명령을 실행하십시오:

`npm install -g`

**`snyk-filter`는 `node-jq` 라이브러리를 사용**하며 이는 [`jq`](https://stedolan.github.io/jq/) 이 설치되어 있어야 합니다. 대부분의 경우 `npm install -g`를 통해 이 과정이 자동으로 이루어지지만, 일부 시스템에서 `jq`가 올바르게 로컬에 설치되지 않을 수 있습니다. `node-jq`와 관련하여 설치 후 오류가 발생하는 경우에는 해당 오류를 피하기 위해 `jq`를 수동으로 설치하십시오.

```
# jq 미리 설치하기 (ubuntu 예시)
sudo apt-get install -y jq

# node-jq에 직접 설치하지 말라고 알림
export NODE_JQ_SKIP_INSTALL_BINARY=true

# node-jq에 기존 jq 이진 파일 위치 알리기
export JQ_PATH=$(which jq)

# 마지막으로 snyk-filter 설치하기 (node 버전 > 12에서 작동하지 않음)
sudo npm install -g
```

## 사용법

1. `.snyk-filter/snyk.yml` 파일에서 현재 작업 디렉토리에 상대적인 사용자 정의 `jq` 필터를 구현하십시오. `snyk test`를 실행하는 위치에서 [sample-filters](https://github.com/snyk-tech-services/snyk-filter/tree/develop/sample-filters)를 참고하여 필터를 조정하고 여기서부터 시작하십시오. [JQPlay](https://jqplay.org/)를 사용하십시오.
2. 그런 다음 `snyk test --json` 출력을 `snyk-filter`로 파이프하거나, 자체 필터가 포함된 yml 파일을 가리키기 위해 `-f` 인수를 사용하십시오(기본 위치(.snyk-filter/snyk.yml)를 사용하지 않는 경우).
3. `snyk-filter`의 종료 코드는 0 (문제 없음)이면 통과이고, 1 (문제 발견)이면 실패입니다.

{% hint style="info" %}
`snyk test --json > result-opensource.json`과 같은 다단계 접근을 사용하고 결과를 플러그인에 전달할 때, [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands)는 빌드 시스템에서 프로세스가 중단되거나 중단될 수 있습니다. 이 경우 빌드 도구의 기능에 따라 여러 옵션이 있습니다:\
\
1\) 코드를 매개변수로 캡처하여 프로세스에 반환되지 않도록하고 오류 상태를 확인합니다.\
2\) `||true` 또는 논리 형태의 논리를 사용하여 [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands)가 프로세스를 종료하지 않도록 지정합니다.\
이렇게하면 네트워크 또는 `Snyk` 플랫폼 문제 또는 다른 스캔 결과 문제를 나타내는 오류 코드와 같이 모든 반환 코드가 무시되므로 JSON을 사용하는 다음 단계는 실패할 가능성이 높습니다. 스크립트의 다음 단계로 넘어가기 전에 종료 코드를 검토하는 것이 좋습니다.\
3\) 가능한 경우 `실패시 계속` 옵션을 설정합니다.
{% endhint %}

### Snyk CLI 사용 예시 (기본적으로 .snyk-filter/snyk.yml 사용)

`snyk test --json | snyk-filter`

### Snyk CLI 및 사용자 지정 yml 파일 위치 사용 예시

`snyk test --json | snyk-filter -f /path/to/example-cvss-9-or-above.yml`

### JSON 파일 입력 예시

`snyk test --json-file-output=results-opensource.json`

`snyk-filter -i results-opensource.json`

### 사용자 지정 yml 파일 위치 사용 예시

`snyk-filter -i snyk_results.json -f /path/to/example-high-upgradeable-vulns.yml`

## 옵션

`--json`은 JSON 출력 옵션입니다.

## 라이선스

라이선스: Apache License, Version 2.0