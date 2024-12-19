---
description: Bitbucket Cloud에 대한 옵션 목록과 몇 가지 예제
---

# Bitbucket Cloud - 예시

`snyk-scm-contributors-count bitbucket-cloud` 명령에 사용 가능한 다음 옵션들이 있습니다:

```
  --version                 버전 번호 표시                              [boolean]
  --help                    도움말 표시                                  [boolean]
  --user                    Bitbucket 클라우드 사용자명                  [필수]
  --password                Bitbucket 클라우드 앱 비밀번호              [필수]
  --workspaces              [선택 사항] Contributor 수를 계산할 Bitbucket 클라우드 워크스페이스 이름/uuid
  --repo                    [선택 사항] 특정 리포지토리에 대해서만 계산
  --exclusionFilePath       [선택 사항] 제외 목록 파일 경로
  --json                    [선택 사항] JSON 출력, "consolidateResults" 명령을 사용할 때 필요
  --skipSnykMonitoredRepos  [선택 사항] Snyk에서 모니터링하는 리포지토리를 건너뛰고 모든 리포지토리에 대해 Contributor 수를 계산
  --importConfDir           [선택 사항] 모니터링되지 않는 리포지토리에 대한 가져오기 파일 생성: 생성된 가져오기 파일을 저장할 유효한 폴더 경로
  --importFileRepoType      [선택 사항] importConfDir 플래그와 함께 사용: 가져오기 파일에 추가 할 리포지토리 유형 지정. 옵션: all/private/public. 기본값: all
```

## 명령을 실행하기 전에

1. `SNYK_TOKEN`을 내보내기(이미 Snyk에서 모니터링하는 리포지토리에 대한 Contributor만 얻고 싶은 경우):
   * 토큰이 그룹 수준 액세스를 갖고 있는지 확인하거나 그룹 수준 액세스를 갖고 있는 서비스 계정 토큰을 사용합니다. 서비스 계정을 만드는 방법에 대해 자세히 알아보려면 [서비스 계정 설정 방법](https://docs.snyk.io/features/integrations/managing-integrations/service-accounts#how-to-set-up-a-service-account)을 참조하세요.
   * 토큰 값을 복사합니다.
   *   환경 변수에 토큰을 내보냅니다:

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2.  Bitbucket Cloud 사용자명(**이메일이 아님**)과 [앱 비밀번호](https://developer.atlassian.com/cloud/bitbucket/rest/intro/#authentication)를 가져옵니다.

    **참고**: 자격 증명이 리포지토리에 읽기 액세스 권한을 갖고 있는지 확인하세요.

## 명령 실행

다음과 같은 사용 수준과 옵션을 고려하세요:

### 사용 수준

*   Bitbucket Cloud의 모든 워크스페이스 및 리포지토리의 커밋을 가져오려면 Bitbucket Cloud 사용자와 앱 비밀번호를 제공하십시오:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD
    ```
*   Bitbucket Cloud의 일부 워크스페이스 및 리포지토리의 커밋을 가져오려면 Bitbucket Cloud 사용자, Bitbucket Cloud 앱 비밀번호 및 쉼표로 구분된 워크스페이스 목록을 제공하십시오:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1,Workspace2...
    ```
*   Bitbucket Cloud의 특정 리포지토리의 커밋을 가져오려면 Bitbucket Cloud 사용자, Bitbucket Cloud 앱 비밀번호, 워크스페이스 및 리포지토리 이름을 제공하십시오:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1 --repo Repo1
    ```

### 옵션

*   이미 Snyk에서 모니터링하고 있는 리포지토리를 고려하지 않고 Bitbucket Cloud에서 모든 커밋을 가져오려면 `--skipSnykMonitoredRepos` 플래그를 추가하십시오.\
    Snyk에서 모니터링하고 있지 않은 Bitbucket Cloud의 리포지토리가 있을 수 있습니다. Snyk 모니터링 리포지토리를 확인하지 않고 커밋을 가져 오려면 이 플래그를 사용하세요.

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --skipSnykMonitoredRepos
    ```
*   커밋에서 빼려는 일부 기여자가 있는 경우 이를 무시하는 이메일 목록이 있는 제외 파일을 추가하고 해당 파일의 경로를 `--exclusionFilePath`에 적용하십시오:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1,Workspace2 --exclusionFilePath PATH_TO_FILE
    ```
*   출력을 json 형식으로 설정하려면 `--json` 플래그를 추가하십시오:

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1 --repo Repo1 --json
    ```
*   모니터링되지 않는 리포지토리를 위해 가져오기 파일을 생성하려면 유효한(쓰기 가능한) 폴더 경로와 함께 `--importConfDir` 플래그를 추가하고 가져오기 파일에 추가할 리포지토리 유형과 함께 선택적으로 `--importFileRepoType` 플래그를 추가하십시오(`all`/`private`/`public`, 기본값은 `all`). 이러한 플래그들은 `--repo` 플래그와 함께 설정할 수 없음에 유의하세요.

    ```
    snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --importConfDir ValidPathToFolder --importFileRepoType private/public/all
    ```

    이러한 플래그에 대한 자세한 정보는 [생성 및 사용 페이지](../../creating-and-using-the-import-file.md)를 참조하세요.
*   자세한 출력을 위해 디버그 모드로 실행하려면 `DEBUG=snyk*`를 접두어로 사용하세요:

    ```
    DEBUG=snyk* snyk-scm-contributors-count bitbucket-cloud --user USERNAME --password APP_PASSWORD --workspaces Workspace1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    .
    ```