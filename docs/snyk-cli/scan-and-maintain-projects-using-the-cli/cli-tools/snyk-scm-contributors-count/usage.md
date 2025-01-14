---
description: SCM-Contributors-Count 모드 및 레벨
---

# 사용법

## 일반 명령어

```
snyk-scm-contributors-count <command> <command-options>
```

**`<command>`**&#xB294; 다음 중 하나입니다:

* `azure-devops`
* `bitbucket-cloud`
* `bitbucket-server`
* `github`
* `github-enterprise`
* `gitlab`
* `consolidateResults`

**`<command-options>`**: SCM별 페이지(예시 페이지)의 [scripts](scripts-for-scm-contributors-count/) 섹션을 참조하십시오.

## 모드

### 온보딩 이전 사용 범위 설정

이 모드는 Bitbucket 및 Azure에서만 작동합니다.

예시로 `skipSnykMonitoredRepos` 플래그를 적용하십시오:

```
snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD --skipSnykMonitoredRepos
```

### Snyk 라이선스 소비

이 모드는 Bitbucket 및 Azure에서만 작동합니다.

`SNYK_TOKEN`을 내보내기하십시오. 예시로:

```
export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password PASSWORD
```

## 레벨

### 최상위 레벨

이 사용 레벨에서 도구는 SCM의 맨 위부터 시작하여 Orgs/그룹을 가져온 후, 모든 repo를 검색하고, 지난 90일간의 커밋 수를 계산합니다.

이 레벨을 사용하려면 자격 증명을 제공하고, 해당하는 경우 호스트 또는 URL을 제공하면 됩니다. 도구는 모든 orgs/groups 및 그들의 모든 repo에 대한 기여자 수를 가져옵니다. 예시로:

```
snyk-scm-contributors-count github --token TOKEN
```

### 중간 레벨

이 사용 레벨에서 도구는 사용자가 제공한 Orgs/그룹에서 시작하여 repo 수준까지 내려가서 지난 90일간의 커밋 수를 계산합니다.

이 레벨을 사용하려면 자격 증명과 가져올 repo 및 해당하는 기여자 수를 원하는 그룹 또는 orgs의 쉼표로 구분된 목록을 제공하면 됩니다. 예시로:

```
snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
```

### 최하위 레벨

이 사용 레벨에서 도구는 기여자 수를 가져올 하나의 repo에 집중합니다.

이 레벨을 사용하려면 자격 증명(적용 가능한 경우 호스트/URL), 한 가지 org/group, 그리고 하나의 repo를 제공하면 됩니다. 예시로:

```
snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO
```

## 디버그 모드

명령어의 시작 부분에 `DEBUG=snyk*`를 추가하십시오. 예시로:

```
DEBUG=snyk* snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
```

## 추가 옵션 플래그

명령어에 추가적인 플래그를 설정할 수 있습니다:

* `snyk-api-import` 도구와 사용자의 `my Snyk 계정`으로 가져올 수 없는 repo 데이터가 포함된 **import 파일**을 생성합니다. Bitbucket 및 Azure에서만 작동합니다. 사용 가능한 폴더 경로를 가리키는 `importConfDir` 플래그를 적용하십시오. 이 플래그는 `importFileRepoType` 플래그와 관련이 있습니다.
* 가져온 파일에 추가할 **repo 유형을 선택**합니다. Bitbucket 및 Azure에서만 작동합니다. `importFileRepoType` 플래그를 `all,` `private`, 또는 `public` 중 하나로 설정하십시오.
* **계수에서 제외할 커미터를 제외**하십시오. 제외하려는 커미터의 이메일을 포함하는 텍스트 파일 경로를 가리키는 `exclusionFilePath` 플래그를 명령어에 적용하십시오.
* 결과를 요약하고 Json 형식으로 출력하십시오. 명령어에 `json` 플래그를 적용하십시오.

## consolidateResults 명령어

여러 SCMs에서 여러 명령어의 결과를 하나의 파일로 통합하여 고유한 기여자 수를 계산하는 데 사용됩니다. [명령 페이지](consolidate-results.md)를 참조하십시오.
