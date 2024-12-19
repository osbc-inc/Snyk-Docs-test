# Snyk 브로커(Custom)를 사용하여 Terraform 설정 파일 감지

## **Snyk 브로커에서의 Terraform**

기본적으로, 인프라스트럭처-애즈-코드(IaC)에서 사용되는 일부 파일 유형은 활성화되어 있지 않습니다. 예를 들어 Terraform과 같이 브로커가 리포지토리의 IaC 파일에 액세스하도록 허용하려면 환경 변수 `ACCEPT_IAC`를 추가할 수 있습니다. 가능한 조합으로 `tf, yaml, yml, json, tpl`을 사용하세요.

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

그렇지 않으면 `accept.json`을 편집하고 관련 IaC 특정 규칙을 추가하여 사용자 정의 accept 파일을 컨테이너로 로드할 수 있습니다. 별도 폴더의 사용자 지정 accept 파일이 사용된 경우(`ACCEPT` 환경 변수 사용),`ACCEPT_IAC` 메커니즘을 사용할 수 없습니다.

Terraform 파일을 Snyk가 스캔할 수 있는 파일로 추가하고자 하는 경우 사용자 정의 허용 목록이 필요한 경우 다음 지침을 따르세요.

## 설정 작성

Terraform 스캐닝 기능은 리포지토리의 `.tf` 파일에 액세스해야 합니다. 이를 위해 특정 API 권한이 필요합니다. 이 API 권한은 사용 중인 소스 컨트롤 시스템에 따라 약간 다를 수 있습니다.

1. 소스 컨트롤 시스템에 맞는 accept.json 예시 파일을 찾아 [브로커 리포지토리](https://github.com/snyk/broker/tree/master/client-templates)에서 다운로드하세요.
2. 파일 이름을 `accept.json`으로 변경하고 JSON 파일의 **private** 배열에 해당 규칙을 추가하세요.
3. [브로커 구성](detecting-terraform-configuration-files-using-snyk-broker-custom.md#configuring-broker)에 대한 지침을 따르세요.

### GitHub 규칙

```
{
  "//": "Terraform 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.tf",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Terraform 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.tf",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
```

### Bitbucket 규칙

```
{
  "//": "Terraform 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.tf",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "Terraform 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.tf",
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
  "//": "Terraform 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.tf",
  "origin": "https://${GITLAB}"
},
{
  "//": "Terraform 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.tf",
  "origin": "https://${GITLAB}"
},
```

### Azure Repos

다음 파일 확장 목록을 복사하세요:

```
            "**/*.yaml",
            "**%2F*.yaml",
            "**/*.yml",
            "**%2F*.yml",
            "**/*.json",
            "**%2F*.json",
            "**/*.tf",
            "**%2F*.tf",
```

[accept.json](https://github.com/snyk/broker/blob/master/client-templates/azure-repos/accept.json.sample)에서 `values` 배열에 확장을 추가하세요:

* `"//": "파일 유형으로 제한하여 파일 내용 가져오기"`
* `"//": "파일 유형으로 제한하여 파일 존재 확인하기"`

## 브로커 구성

브로커는 ACCEPT 환경 변수에 추가된 적용 가능한 규칙이 있는 accept.json 파일 경로를 취합니다. 다음은 해당 변수를 GitHub 브로커에 전달하는 예제입니다.

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

이를 통해 Snyk에게 `.tf` 파일을 조회할 수 있는 능력을 제공합니다. 더 엄격하게 하려면 앞의 예제에서 경로를 변경하여 특정 프로젝트나 파일 레이아웃에 대해 더 강력한 제한을 설정할 수 있습니다.