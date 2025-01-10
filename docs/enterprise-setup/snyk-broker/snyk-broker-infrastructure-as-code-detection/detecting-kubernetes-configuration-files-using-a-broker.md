# Snyk Broker(사용자 정의)를 사용하여 Kubernetes 구성 파일 감지하기

## Snyk 브로커에서 Kubernetes 구성 파일 감지

기본적으로 인프라스트럭처-애즈-코드 (IaC)에서 사용되는 일부 파일 유형은 비활성화되어 있습니다. 예를 들어 브로커가 리포지토리에 IaC 파일에 액세스하도록 허용하려면 Kubernetes 구성 파일과 같은 파일에 환경 변수인 `ACCEPT_IAC`를 추가할 수 있습니다. `tf,yaml,yml,json,tpl`의 어떤 조합도 사용할 수 있습니다.

예시:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=비밀-브로커-토큰 \
           -e GITHUB_TOKEN=비밀-깃헙-토큰 \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl
       snyk/broker:github-com
```

그렇지 않으면 `accept.json`을 편집하고 관련 IaC 특정 규칙을 추가하여 사용자 지정 허용 파일을 컨테이너에 로드할 수 있습니다. 사용자 지정 허용 파일(다른 폴더에서 가져온)을 사용하는 경우(`ACCEPT` 환경 변수를 사용함), `ACCEPT_IAC` 메커니즘을 사용할 수 없음에 주의하시기 바랍니다.

이러한 지침은 사용자 지정 허용 목록이 필요하며 Snyk이 스캔할 수 있는 파일에 Kubernetes 구성 파일을 추가하려는 경우입니다.

## 구성 작성

리포지토리의 특정 파일에 브로커에 액세스 권한을 부여해야 합니다. 이는 특정 API 권한이 필요합니다. 사용하는 소스 컨트롤 시스템에 따라이 API 권한은 약간 다를 수 있습니다. 다음 구성은 `.yaml`, `.yml`, `.json` 파일 확장자를 위한 것입니다. 이를 통해 브로커가 Kubernetes 및 CloudFormation 파일에 액세스할 수 있지만 필요에 따라 구성을 조정할 수 있습니다. 예를 들어 Terraform HCL 파일을 스캔하기 위해 `.tf` 파일에 대한 구성을 추가할 수 있습니다.

1. 소스 컨트롤 시스템에 적합한 `accept.json` 샘플 파일을 [브로커 리포지토리](https://github.com/snyk/broker/tree/master/client-templates)에서 다운로드합니다.
2. 파일 이름을 `accept.json`로 변경하고 JSON 파일의 **private** 배열에 다음 규칙을 적절하게 추가합니다.
3. [브로커 구성](detecting-kubernetes-configuration-files-using-a-broker.md#configuring-broker)을 위한 지침을 따릅니다.

### GitHub 및 GitHub Enterprise 규칙

```
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.yml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.yml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.json",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.json",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.tpl",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.tpl",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
```

### Bitbucket 규칙

```
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
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
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
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
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
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
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
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
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
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
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.json",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.tpl",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.tpl",
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
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.yml",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.yml",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.json",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.json",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.tpl",
  "origin": "https://${GITLAB}"
},
{
  "//": "Infrastructure as Code 문제를 결정하는 데 사용됨",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.tpl",
  "origin": "https://${GITLAB}"
},
```

### Azure Repo 규칙

```````
{
  "public": [
    {
      "//": "Azure에서 웹훅을 푸시하기 위해 사용됨",
      "method": "POST",
      "path": "/webhook/azure-repos/:webhookId"
    }
  ],
  "private": [
    {
      "//": "지정된 조직에 대한 프로젝트 목록을 가져옴",
      "method": "GET",
      "path": "/_apis/projects",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "지정된 조직에 대한 특정 저장소를 가져옴",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "지정된 조직에 대한 저장소 목록을 가져옴",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "ref 목록을 가져옴",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/refs",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "지정된 조직의 저장소를 검색함",
      "method": "GET",
      "path": "_apis/git/repositories",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "훅 생성함",
      "method": "POST",
      "path": "/_apis/hooks/subscriptions",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "훅 삭제함",
      "method": "DELETE",
      "path": "/_apis/hooks/subscriptions/:subscriptionId",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "파일 컨텐츠를 가져옴. 파일 유형으로 제한됨",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/items",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "valid": [
        {
          "queryParam": "path",
          "values": [
            // 여기에 파일 타입 추가
          ]
        },
        {
          "queryParam": "recursionLevel",
          "values": ["none"]
        },
        {
          "queryParam": "download",
          "values": ["true"]
        },
        {
          "queryParam": "includeContent",
          "values": ["true"]
        }
      ],
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "지정된 저장소 파일 목록을 가져옴",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/items",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
  ],
}
``````json
{
  "valid": [
    {
      "queryParam": "recursionLevel",
      "values": ["full"]
    },
    {
      "queryParam": "download",
      "values": ["false"]
    },
    {
      "queryParam": "includeContent",
      "values": ["false"]
    }
  ],
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "get list of commits for given repository",
  "method": "GET",
  "path": "/:owner/_apis/git/repositories/:repo/commits",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "update status of given commit",
  "method": "POST",
  "path": "/:owner/_apis/git/repositories/:repo/commits/:commitId/statuses",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "update status of given pull request",
  "method": "POST",
  "path": "/:owner/_apis/git/repositories/:repo/pullRequests/:pullRef/statuses",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "find PR for given repository",
  "method": "GET",
  "path": "/:owner/_apis/git/repositories/:repo/pullrequests",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "create new PR in given repository",
  "method": "POST",
  "path": "/:owner/_apis/git/repositories/:repo/pullrequests",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "update existing PR in given repository",
  "method": "PATCH",
  "path": "/:owner/_apis/git/repositories/:repo/pullrequests/:pullRef",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "push new commit in given repository",
  "method": "POST",
  "path": "/:owner/_apis/git/repositories/:repo/pushes",
  "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
  "auth": {
    "scheme": "basic",
    "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
  }
},
{
  "//": "used to redirect requests to snyk git client",
  "method": "any",
  "path": "/snykgit/*",
  "origin": "${GIT_CLIENT_URL}"
}
]
}
```````

## Broker 구성

Broker는 ACCEPT 환경 변수에 적용 가능한 규칙이 추가된 accept.json 파일의 경로를 취합니다. 다음은 해당 변수를 GitHub Broker에 전달하는 예제를 제공합니다.

```bash
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e GITHUB_TOKEN=secret-github-token \
  -e PORT=8000 \
  -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
  -e ACCEPT=/private/accept.json \
  -v /local/path/to/private:/private \
  snyk/broker:github-com
```

주의: 이를 통해 Snyk이 `.yaml`, `.yml` 또는 `.json` 파일에 대해 쿼리할 수 있게 됩니다. 특정 프로젝트나 파일 레이아웃에 대해 더 제한적으로 설정하려면 앞의 예제에서 경로를 더 제한적으로 변경할 수 있습니다.

```
```
