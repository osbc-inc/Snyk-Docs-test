# 도구: snyk-api-import

Snyk는 종속성에서 알려진 취약점을 찾아 수정하고 모니터링하는 데 도움을 줍니다. 이는 선택적으로 하되 메이저한 변동 사항과 빌드 시스템 내에서 일어나는 연속적인 통합 (CI)의 일부로 이뤄집니다.

## snyk-api-import

Snyk API 프로젝트 가져오기 도구인 `snyk-api-import`는 GitHub, GitLab, Bitbucket 등의 시스템에서 발생할 수 있는 속도 제약을 피하며 안정적인 가져오기를 제공하기 위해 사용 가능한 [Snyk APIs](../../../snyk-api/reference/)를 활용하여 프로젝트를 가져오는 데 도움을 주는 스크립트입니다. 이 스크립트는 가져오기를 일괄적으로 시작하고 완료를 기다린 다음 계속 진행합니다. 어떤 실패한 요청도 실수 처리되고 실패로 간주되며 로그에 기록되기 전에 재시도됩니다.

동시성을 조정해야 하는 경우, 스크립트를 중지하고 동시성 변수를 변경한 다음 다시 시작할 수 있습니다. 이 도구는 이미 가져오기가 요청된 이전 저장소 (Targets)를 건너뜁니다.

`snyk-api-import`를 사용하려면 미리 다음을 수행해야 합니다:

- 가져오기 전에 Snyk 조직 (Orgs)를 설정합니다.
- 해당 Snyk 조직을 SCM (GitHub, GitLab, Bitbucket 등)과 어떤 연결로 구성합니다. 가져오기 파일을 생성하기 위해 `integrationId`가 필요합니다.
- 가져오기 알림을 받지 않도록 설정하기 위해 [알림 설정](../../../snyk-api/reference/organizations-v1.md#org-orgid-notification-settings) API 엔드포인트를 사용하여 이메일 알림 등을 비활성화합니다 (권장됨).
- 가져오기가 완료될 때까지 SCMs (GitHub, GitLab, Bitbucket 등)로 추가 요청을 보내지 않도록 가져오기가 완료될 때까지 fix PRs와 PR 확인을 비활성화하기 위해 [업데이트](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-settings) (통합 설정) 엔드포인트를 사용합니다.

## 설치

`snyk-api-import` CLI는 여러 채널을 통해 설치할 수 있습니다.

### 독립 실행형 실행 파일 (macOS, Linux, Windows)

[GitHub 릴리스](https://github.com/snyk/snyk-api-import/releases)를 사용하여 해당 플랫폼용 `snyk-api-import` CLI의 독립 실행 파일을 다운로드하세요.

### npm 또는 Yarn으로 설치

[snyk-api-import CLI는 npm 패키지로 제공됩니다](https://www.npmjs.com/package/snyk-api-import). 로컬에 Node.js가 설치된 경우 다음을 실행하여 설치할 수 있습니다:

```
npm install snyk-api-import@latest -g
```

Yarn을 사용하는 경우 다음을 실행합니다:

```
yarn global add snyk-api-import
```

## 사용법

명령이 명시되지 않은 경우 `import` 명령이 기본적으로 실행됩니다.

- `import` - 가져오기 구성 파일에 정의된 기존 Snyk 조직으로 대상 저장소 (Targets)를 API를 활용하여 가져오기를 시작합니다. 모든 프로젝트 유형에 대한 지원은 [가져오기](../../../snyk-api/reference/import-projects-v1.md) API 엔드포인트, [가져오기 대상](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import) 및 [가져오기 작업 세부 정보 가져오기](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import-jobid)를 통해 제공됩니다. [가져오기 API (프로젝트 가져오기, 가져오기)](https://snyk.docs.apiary.io/#reference/import-projects).
- `help` - 도움말 및 사용 가능한 모든 명령 및 옵션 표시.
- `orgs:data` 유틸리티 - API를 사용하여 조직을 생성하는 데 필요한 데이터 생성에 사용합니다.
- `orgs:create` 유틸리티 - `orgs:data` 명령으로 생성된 데이터 파일을 기반으로 Snyk에 조직을 생성하는 데 사용합니다.
- `import:data` 유틸리티 - 가져오기를 시작하는 데 필요한 데이터를 생성하는 데 사용합니다. 기본적으로 아카이브된 저장소는 제외됩니다.
- `list:imported` 유틸리티 - 가져오기 중에 이전에 가져온 대상을 건너뛰는 데 도움을 주는 데이터를 생성하는 데 사용합니다.

로그는 [Bunyan CLI](http://trentm.com/node-bunyan/bunyan.1.html)를 사용하여 탐색할 수 있습니다.

## snyk-api-import 지침의 내용

- 유틸리티
  - [Snyk에 조직 생성](creating-organizations-in-snyk.md)
  - [가져오기 대상 및 가져오기 데이터 생성](creating-import-targets-data-for-import-command.md)
  - [Snyk에서 GitHub.com 및 GitHub Enterprise 조직 및 저장소 미러링](mirroring-github.com-and-github-enterprise-organizations-and-repos-in-snyk.md)
  - [Snyk에서 GitLab 조직 및 저장소 미러링](mirroring-gitlab-organizations-and-repos-in-snyk.md)
  - [Snyk에서 Bitbucket Server 조직 및 저장소 미러링](mirroring-bitbucket-server-organizations-and-repos-in-snyk.md)
  - [Snyk에서 Bitbucket Cloud 조직 및 저장소 미러링](mirroring-bitbucket-cloud-organizations-and-repos-in-snyk.md)
- [가져오기 시작](kicking-off-an-import.md)
- [snyk-api-import에 기여](contributing-to-snyk-api-import.md)
- [싱크: 모니터링된 저장소의 변경 감지 및 Snyk 프로젝트 업데이트](https://github.com/snyk/snyk-api-import/blob/master/docs/sync.md)
- 예제 워크플로: [AWS 자동화](https://github.com/snyk/snyk-api-import/blob/master/docs/example-workflows/aws-automation-example.md)

## FAQ

<details>

<summary><code>Error: ENFILE: file table overflow, open</code> 또는 <code>Error: EMFILE, too many open files</code></summary>

이러한 오류가 발생하는 경우 **ulimit**을 높여 더 많은 파일 작업을 허용해야 할 수 있습니다. 연산을 지연하지 않고 편리한 시점에 로그를 기록하기 때문에 도구는 루프의 끝까지 기다리지 않고 큰 데이터 구조를 기록합니다. 이는 설정한 동시 가져오기 수에 따라 시스템 기본 **ulimit**을 초과할 수 있음을 의미합니다.

다음 자료들이 **ulimit**을 높일 수 있도록 도와줍니다:

- [ss64.com](https://ss64.com/bash/ulimit.html)
- [StackOverflow](https://stackoverflow.com/questions/45004352/error-enfile-file-table-overflow-scandir-while-run-reaction-on-mac)
- [blog.mact.me](http://blog.mact.me/2014/10/22/yosemite-upgrade-changes-open-file-limit)

</details>

<details>

<summary><code>ERROR: HttpError: request to https://github.private.com failed, reason: self signed certificate in certificate chain</code></summary>

만약 GitHub, GitLab, Bitbucket, 또는 Azure 인스턴스가 자체 서명된 인증서를 사용 중이라면, `snyk-api-import`를 이러한 인증서를 사용하도록 구성할 수 있습니다.

`export NODE_EXTRA_CA_CERTS=./path-to-ca`

</details>

<details>

<summary>이 도구는 브로커롱된 통합과 함께 작동합니까?</summary>

네. Snyk는 가져오기를 수행하기 위해 SCM (Git) 저장소와의 기존 통합을 재사용하므로 브로커화된 연결이 구성된 경우 해당 연결이 사용됩니다.

</details>

<details>

<summary>가져오기 명령에서 무엇을 지원합니까?</summary>

`snyk-api-import`는 [가져오기 API](../../../snyk-api/reference/import-projects-v1.md) 문서에서 식별된 모든 통합 유형 및 프로젝트 소스를 지원합니다. 사용 사례에 대한 예제가 이 지침에 없는 경우 API 문서를 참조하세요.

</details>