# snyk-delta

이 도구는 두 Snyk Open Source 스냅숏 간의 차이를 구하는 수단을 제공합니다. 특히 CLI 기반의 스캔(로컬 환경, 깃훅 등에서 실행될 때)을 진행할 때 유용합니다.

`snyk-delta`는 스냅숏을 비교하여 다음과 같은 세부 정보를 제공합니다:
- 기준 스냅숏에 발견되지 않은 새로운 취약점
- 기준 스냅숏에 발견되지 않은 새로운 라이선스 이슈
- 두 스냅숏 간의 의존성 변화:
  - 직접 의존성 추가 및 제거
  - 간접 의존성 추가 및 제거
  - 새로운 취약점을 포함하고 있는 플래그 경로

## 전제 조건

- Snyk Enterprise 요금제 (Snyk API가 필요합니다)
- 모니터링되어야 하는 프로젝트

## 설치

`npm i -g snyk-delta`

또는

[릴리스 페이지](https://github.com/snyk-tech-services/snyk-delta/releases)에서 선택한 바이너리를 다운로드합니다.

## 사용법

이 도구를 인라인 또는 독립형 명령어로 사용할 수 있습니다.

### 인라인 작업

`snyk test --json --print-deps | snyk-delta`를 사용합니다.

* 특정 스냅숏을 가리키려면 org+project 좌표를 지정합니다:
  `snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx`
* `snyk-prevent_commit_status`와 함께 사용되고 프로젝트가 모니터링되지 않은 경우, `--setPassIfNoBaseline`를 사용합니다. 이렇게 하면 `snyk-prevent_commit_status`가 실패하지 않습니다.
  - `setPassIfNoBaseline` 기본값은 false입니다.
  - `snyk test --json --print-deps | snyk-delta --baselineOrg xxx --baselineProject xxx --setPassIfNoBaseline true`

{% hint style="info" %}
**BaselineProject** 값은 이름이 아닌 UUID여야 합니다.
그 UUID는 Snyk 웹 UI나 API에서 검색할 수 있습니다.
{% endhint %}

### 독립형 작업

`snyk-delta --baselineOrg xxx --baselineProject xxx --currentOrg xxx --currentProject xxx --setPassIfNoBaseline false`를 사용합니다.

## 모듈로 사용

```
import { getDelta } from 'snyk-delta'

const jsonResultsFromSnykTest = 파일에서 읽거나 snyk test 명령어로 받은 JSON 결과

const result = await getDelta(jsonResultsFromSnykTest);
```

**결과**는 다음과 같은 숫자입니다:
- 0: 새로운 문제 없음
- 1: 새 문제(들) 또는 StrictMode를 사용하고 비모니터링 프로젝트에 문제가 있는 경우 (자세한 내용은 [StrictMode](snyk-delta.md#strictmode) 참조)
- 2: 잘못된 인증과 같은 오류

문제에 대한 상세 정보는 표준 출력에 나열됩니다.

## snyk-delta 도움말

도움말을 표시하려면 `-h`를 사용합니다.

### StrictMode

`snyk-delta`는 테스트 결과를 비교할 때 Snyk 플랫폼에서 모니터링되고 있는 동일한 프로젝트를 찾으려고 합니다. 모니터링된 프로젝트를 찾을 수 없는 경우, `snyk-delta`는 CLI 스캔에서 발견한 모든 문제를 반환하여 통과로 작동합니다.

문제가 없는 경우 반환 코드는 0이며, 문제가 발견되면 1입니다.

### 주의 사항

모듈로 사용 시 Snyk CLI에서 가져온 이슈 목록이 필요합니다.
`snyk-delta`는 Snyk API에서 직접 데이터를 가져오는 것과 호환되지 않습니다.

### all-projects

`snyk-delta`는 `--all-projects` 옵션을 지원하지 않지만, 그렇게 될 때까지 해결책으로 `snyk_delta_all_projects.sh`를 사용해 볼 수 있습니다.

## snyk-delta의 문제 해결

문제가 발생한 경우 다음을 시도해 볼 수 있습니다:
- Snyk `test -d` 단계를 먼저 실행하고 정상 작동하는지 확인합니다.
- `delta allprojects` 스크립트를 사용하는 경우, 해당 스크립트를 제거하고 프로젝트를 개별적으로 테스트합니다.
- 기준 스냅숏을 찾을 수 없는 경우, 먼저 모니터링되는 기존 프로젝트가 있는지 확인하고 SCM에서 모니터링되는 프로젝트와 일치시키려는 경우 `.git` 메타데이터를 확인합니다.

도움이 필요한 경우 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.