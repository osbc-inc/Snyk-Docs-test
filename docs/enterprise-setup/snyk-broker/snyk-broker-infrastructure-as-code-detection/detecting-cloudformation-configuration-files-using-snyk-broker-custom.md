# Snyk 브로커(Custom)를 사용하여 CloudFormation 구성 파일 검색

## Snyk 브로커에서의 CloudFormation

기본적으로, 인프라스트럭처-애스-코드 (IaC)에서 사용되는 일부 파일 유형은 활성화되어 있지 않습니다. 예를 들어 리포지토리에서 IaC 파일에 브로커 액세스 권한을 부여하려면 CloudFormation과 같이, 환경 변수인 `ACCEPT_IAC`를 추가할 수 있습니다. `tf,yaml,yml,json,tpl`의 어떤 조합도 사용할 수 있습니다.

예시:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl
       snyk/broker:github-com
```

그 외에도, `accept.json`을 편집하여 관련 IaC 특정 규칙을 추가하고 컨테이너에 사용자 지정 accept 파일을 로드할 수 있습니다. 사용자 정의 accept 파일(별도의 폴더에서 사용)을 사용하는 경우(`ACCEPT` 환경 변수 사용) `ACCEPT_IAC` 메커니즘을 사용할 수 없음을 참고해야 합니다.

CloudFormation 파일을 Snyk가 스캔할 수 있도록 포함하려는 경우 사용하는 사용자 지정 허용 목록에 관한 요구 사항은 다음과 같습니다.

## 구성 작성

CloudFormation 스캔 기능은 리포지토리의 YAML 또는 JSON 파일에 액세스할 수 있는 권한이 필요합니다. 이는 특정 API 권한이 필요합니다. 이러한 API 권한은 소스 컨트롤 시스템에 따라 약간 다를 수 있습니다.

1. 올바른 소스 컨트롤 시스템용 accept.json 샘플 파일을 브로커 리포지토리에서 다운로드하고 [다음 링크](https://github.com/snyk/broker/tree/master/client-templates)로 이동합니다.
2. 파일을 `accept.json`로 이름을 바꾸고 JSON 파일의 **private** 배열에 필요에 맞는 규칙을 추가합니다.
3. [브로커 구성 지침](detecting-cloudformation-configuration-files-using-snyk-broker-custom.md#configuring-broker)을 따릅니다.

### GitHub 및 GitHub Enterprise 규칙

```
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.yml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.yml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.json",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.json",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
```

### Bitbucket 규칙

```
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.yaml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.yaml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.yml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.yml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.json",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.json",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
```

### GitLab 규칙

```
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.yml",
  "origin": "https://${GITLAB}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.yml",
  "origin": "https://${GITLAB}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.json",
  "origin": "https://${GITLAB}"
},
{
  "//": " 이슈를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.json",
  "origin": "https://${GITLAB}"
},
```

## 브로커 구성

브로커는 사용자가 추가한 규칙이 포함된 accept.json 파일 경로를 `ACCEPT` 환경 변수로 가져옵니다. 다음은 해당 변수를 GitHub 브로커에 전달하는 예시입니다.

```
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e GITHUB_TOKEN=secret-github-token \
  -e PORT=8000 \
  -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
  -e ACCEPT=/private/accept.json
  -v /local/path/to/private:/private \
  snyk/broker:github-com
```

이러면 Snyk가 `.yaml`, `.yml`, 또는 `.json` 파일에 대한 쿼리 수행이 가능해집니다. 특정 프로젝트 또는 파일 레이아웃에 대해 더 엄격하게 하려는 경우, 앞의 예시 경로를 수정하여 특정 프로젝트나 파일 레이아웃에 대해 더 제한적으로 할 수 있습니다.