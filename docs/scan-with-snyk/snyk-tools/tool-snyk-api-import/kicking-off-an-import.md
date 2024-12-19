# Import 시작

`snyk-api-import`은 Snyk API를 사용하여 가져올 수 있는 동일한 프로젝트 소스를 지원합니다: Git 리포지토리, Docker 이미지, 컨테이너, 구성 파일 등 다양한 소스. Snyk 조직 설정 페이지의 통합 설정을 사용하여 통합을 구성할 수 있습니다. 자세한 내용은 [Snyk Projects 문서 페이지](../../../snyk-admin/snyk-projects/#target)에서 **Target**의 정의를 참조하십시오.

로그는 `SNYK_LOG_PATH` 디렉토리에 생성됩니다.

다음은 가져오기를 시작하는 단계입니다.

## `import-projects.json` 파일 생성

파일은 **필수** `targets` 최상위 키를 가져야합니다. 이 키는 가져오려는 대상의 배열입니다.

```json
{
  targets: [
    {..},
    {..}
  ],
}
```

각 **가져오기 대상**은 다음과 같은 키를 가지고 있습니다:

```json
{
  // 필수
  "orgId": "<public_snyk_org_id>",
  "integrationId": "<public_snyk_integration_id>",
  "target": {..} // 프로젝트가 위치하는 곳의 식별자 (예: Github의 브랜치, 리포지토리 이름 및 소유자)

   // 옵션
  "files": [ { path: "full/path/to/file1"} , { path: "full/path/to/file2" }],
  "exclusionGlobs": "fixtures, tests, __tests__, node_modules",
}
```

* `orgId` - 귀하의 조직 설정 페이지에서 찾을 수 있습니다.
* `integrationId` - 귀하의 조직 설정 페이지에서 각 SCM에 대해 통합 메뉴에서 찾을 수 있습니다.
* `target`, `files`, `exclusionGlobs` - 자세한 내용은 [Snyk Import API 문서](../../../snyk-api/reference/import-projects-v1.md)를 참조하십시오.
  * `exclusionGlobs` - 스캔에서 제외할 최대 열 개의 폴더 이름으로 구성된 쉼표로 구분된 목록 (각 폴더 이름은 100자를 초과하지 않아야 함). 지정하지 않으면 기본값은 "fixtures, tests, **tests**, node\_modules"입니다. 빈 문자열이 제공되면 어떤 폴더도 제외되지 않습니다.
  * `files` - 객체 배열입니다. 각 경로는 Target의 루트로부터 파일까지의 전체 상대 경로여야 합니다. 해당 위치에서 발견된 파일만 가져옵니다.

리포지토리마다 200개 이상의 매니페스트 파일이 있는 경우, Snyk는 특정 파일을 지정하여 가져오기를 여러 번 분할하는 것을 권장합니다. 단일 리포지토리에서 한 번에 수백 개의 파일을 가져 오는 것은 가져오기가 일부 오류나 실패로 이어질 수 있습니다.

가져오기를 분할하여 일부 파일이나 일부 폴더만 가져올 경우, 재시도를 이용할 수 있으며 사용 중인 소스 제어 관리 시스템에 대한 부하를 줄일 수 있습니다. 이를 위해 가져오기 JSON에서 `files` 속성을 채웁니다.

무시해야 할 테스트 또는 fixtures가 있는 경우 `exclusionGLobs` 속성을 설정합니다:

> 스캔에서 제외해야 하는 최대 열 개의 폴더 이름으로 구성된 쉼표로 구분된 목록입니다. 지정하지 않으면 "fixtures, tests, **tests**, node\_modules"로 기본 설정됩니다. 빈 문자열이 제공되면 어떤 폴더도 제외되지 않습니다.

### **예: GitLab**

```json
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "id": 13,
        "branch": "master"
      },
    },
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "id": 2,
        "branch": "master"
      },
    },
  ]
}
```

### **예: Bitbucket Server**

```json
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "repoSlug": "api-import-circle-test",
        "name": "Snyk api-import-circle-test",
        "projectKey": "PROJECT"
      },
      "files": [
        {
          "path": "package.json"
        },
        {
          "path": "package/package.json"
        },
        {
          "path": "Gemfile.lock"
        }
      ],
      "exclusionGlobs": "fixtures, test"
    },
  ]
}
```

### **예: GitHub.com, GitHub Enterprise, dev.azure.com, Hosted Azure Repos**

```json
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "name": "shallow-goof-policy",
        "owner": "api-import-circle-test",
        "branch": "master"
      },
      "exclusionGlobs": "fixtures, test"
    }
  ]
}
```

### **예: Google Container Registry**

```json
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "name": "projectId/repository:tag"
      },
    }
  ]
}
```

### **예: Azure Container Registry, Elastic Container Registry, Artifactory Container Registry**

```json
{
  "targets": [
    {
      "orgId": "******",
      "integrationId": "******",
      "target": {
        "name": "repository:tag"
      },
    }
  ]
}
```

## 환경 변수 설정

* `SNYK_IMPORT_PATH`- 가져올 파일의 경로 또는 `--file` 매개변수 사용
* `SNYK_TOKEN` - 귀하의 [Snyk API 토큰](https://app.snyk.io/account)
* `SNYK_LOG_PATH` - 모든 로그가 저장될 폴더의 경로. Snyk는 실행 중인 각 가져오기에 대한 전용 로그 폴더를 생성하는 것을 권장합니다. 참고: 모든 로그가 추가됩니다.
* `CONCURRENT_IMPORTS` (옵션) - 한 번에 권장되는 최대 15개 리포지토리에 기본값으로 설정되며, 한 리포지토리만 많은 프로젝트를 가질 수 있으므로 사용자의 SCM 인스턴스로부터 많은 파일을 요청할 수 있습니다. 일부는 요청 속도 제한을 가질 수 있습니다. 이 스크립트는 요청 제한에 도달할 위험을 줄이기 위해 만들어졌습니다.
* `SNYK_API` (옵션) 기본값은 `https://api.snyk.io/v1`

## 다운로드 및 실행

[릴리스 페이지](https://github.com/snyk/snyk-api-import/releases)에서 바이너리를 다운로드하고 `DEBUG=snyk* snyk-api-import-macos import --file=path/to/imported-targets.json`로 실행하십시오.

### **이전에 가져온 대상 모두 건너뛰기**

이 유틸리티는 이전에 가져온 대상 (리포지토리)을 건너뛰도록 사용할 수 있어서 나머지 대상만 가져오게 됩니다.

해당 유틸리티는 특정 Snyk 그룹에 이미 있는 프로젝트를 분석하여 `imported-targets.log` 파일을 생성합니다. 로깅 경로에 있는 경우, 해당 파일은 가져오기 중 건너뛸 대상을 조회하는 데 사용됩니다.

### 예시

* 모든 GitHub 리포지토리는 초기 온보딩 중에 해당하는 Snyk 그룹에 가져와 졌습니다.
* 이후 새로운 GitHub 리포지토리가 추가되고 이를 Snyk에 추가해야 합니다.
* 모든 것을 다시 가져오지 않으려면, 이 유틸리티를 사용하여 "새로운" GitHub 리포지토리만 가져올 수 있습니다. 이렇게 하면 훨씬 빠르고 불필요한 Snyk 및 GitHub에 대한 파일 가져오기 및 모든 것을 다시 가져오는 호출이 제거됩니다.

### 동일한 대상 가져오기

* 다른 조직에 가져온 동일한 대상을 가져올 수 있습니다.
* 다른 소스에서 가져온 동일한 대상을 가져올 수 있습니다; 예를 들어 GitHub에 있는 동일한 리포지토리를 GitHub Enterprise를 통해 동일한 조직에 가져올 수 있습니다.

### 실행할 명령어

* 이전에 모든 조직에 가져온 것을 모두 건너뛰려면: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
* 특정 조직에 이전에 가져온 것을 모두 건너뛰려면: `snyk-api-import-macos list:imported --integrationType=<integration-type> --orgId=<snyk_org_id>`
* 단일 통합 및 프로젝트 소스를 가져오려면: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>`
* 여러 통합 및 프로젝트 소스를 가져오려면: `snyk-api-import-macos list:imported --integrationType=<integration-type> --integrationType=<integration-type> --orgId=<snyk_org_id>`

### 지원되는 통합 유형

* GitHub.com `github`
* GitHub Enterprise `github-enterprise`
* GitLab `gitlab`
* Bitbucket Cloud `bitbucket-cloud`
* Google Cloud Registry `gcr`
* DockerHub registry `docker-hub`
* Azure repos `azure-repos`