# Bitbucket Server - 예제

## Bitbucket Server - 예제

***

### description: `Bitbucket Server`를 위한 옵션 목록과 몇 가지 예제입니다.

## Bitbucket Server - 예제

다음은 `snyk-scm-contributors-count bitbucket-server` 명령에 사용 가능한 옵션입니다:

```
  --version                 버전 번호 표시                        [boolean]
  --help                    도움말 표시                                  [boolean]
  --token                   Bitbucket 서버 토큰                     [required]
  --url                     Bitbucket 서버 베이스 URL 예: (https://bitbucket.mycompany.com)         [required]
  --projectKeys             [옵션] 기여자 수를 계산할 Bitbucket 서버 프로젝트 키
  --repo                    [옵션] 특정 리포를 위해 카운트할 수 있음
  --exclusionFilePath       [옵션] 제외 목록 파일 경로
  --json                    [옵션] JSON 출력, "consolidateResults" 명령을 사용할 때 필요
  --skipSnykMonitoredRepos  [옵션] Snyk 감시 중인 리포는 건너뛰고 모든 리포의 기여자 수 계산
  --importConfDir           [옵션] 감시되지 않는 리포에 대한 가져오기 파일 생성: 생성된 가져오기 파일에 대한 유효한 폴더 경로
  --importFileRepoType      [옵션] importConfDir 플래그와 함께 사용: 가져오기 파일에 추가할 리포 유형 지정. 옵션: all/private/public. 기본값: all
```

### 명령 실행 전

1. Snyk이 이미 모니터링 중인 리포에 대한 기여자만 얻고 싶다면 SNYK\_TOKEN을 내보냅니다.:
   * 토큰이 그룹 수준 액세스 권한을 갖고 있는지 확인하거나 그룹 수준 액세스 권한을 갖고 있는 서비스 계정 토큰을 사용합니다. 서비스 계정을 만드는 방법에 대해서는 [서비스 계정 설정 방법](https://docs.snyk.io/features/integrations/managing-integrations/service-accounts#how-to-set-up-a-service-account)을 참조하십시오.
   * 토큰 값을 복사합니다.
   *   환경에 토큰을 내보냅니다:

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2. Bitbucket 서버 토큰과 URL을 가져옵니다:
   *   만약 존재하지 않는다면, 토큰을 생성하십시오. Bitbucket 서버 통합을 사용하여 [가이드](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver)를 참조합니다.

       **참고**: 토큰이 리포에 대한 읽기 액세스 권한을 갖고 있는지 확인하십시오.
   * URL은 Bitbucket 서버 인스턴스의 실제 URL이며, 예를 들어 http://bitbucket-server.mycompany.com입니다.

### 명령 실행하기

다음은 사용 수준과 옵션을 고려하십시오:

#### 사용 수준

*   Bitbucket 서버에서 모든 프로젝트 및 해당 리포에 대한 커밋을 얻으려면 Bitbucket 서버 토큰 및 URL을 제공하십시오:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL
    ```
*   Bitbucket 서버의 일부 프로젝트 및 해당 리포에 대한 커밋을 얻으려면 Bitbucket 서버 토큰, Bitbucket 서버 URL 및 프로젝트를 쉼표로 구분하여 지정하십시오:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1,Key2...
    ```
*   Bitbucket 서버의 특정 리포에 대한 커밋을 얻으려면 Bitbucket 서버 토큰, Bitbucket 서버 URL, 프로젝트 및 리포 이름을 제공하십시오:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1
    ```

#### 옵션

*   Snyk에서 이미 모니터링 중인 리포를 고려하지 않고 Bitbucket Server에서 모든 커밋을 가져오려면 `--skipSnykMonitoredRepos` 플래그를 추가하십시오.\
    Snyk에서 모니터링되지 않는 Bitbucket 서버의 리포가 있을 수 있습니다. 이 플래그를 사용하여 Snyk에서 모니터링되는 리포의 확인을 건너뛰고 커밋을 직접 Bitbucket 서버에서 가져옵니다.

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --skipSnykMonitoredRepos
    ```
*   특정 기여자를 커밋에서 제외하려면 무시할 이메일을 포함한 제외 파일을 추가하고 해당 파일 경로에 `--exclusionFilePath`를 적용하십시오:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1,Key2 --exclusionFilePath PATH_TO_FILE
    ```
*   출력을 JSON 형식으로 설정하려면 `--json` 플래그를 추가하십시오:

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --json
    ```
*   나의 감시되지 않는 리포에 대한 가져오기 파일을 생성하려면 `--importConfDir` 플래그를 추가하고 가져오기 파일이 저장될 유효한(쓰기 가능한) 폴더 경로를 지정하고 선택적으로 `--importFileRepoType` 플래그를 추가하여 파일에 추가할 리포 유형(`all`/`private`/`public`, 기본값: `all`)을 지정하십시오. 이러한 플래그는 `--repo` 플래그와 함께 설정할 수 없습니다.

    ```
    snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --importConfDir ValidPathToFolder --importFileRepoType private/public/all
    ```

    이러한 플래그에 대한 자세한 정보는 [가져오기 파일 생성 및 사용](../../creating-and-using-the-import-file.md)을 참조하십시오.
*   자세한 출력을 위해 디버그 모드에서 실행하려면 명령어 앞에 `DEBUG=snyk*`를 추가하십시오:

    ```
    DEBUG=snyk* snyk-scm-contributors-count bitbucket-server --token BITBUCKET-TOKEN --url BITBUCKET-URL --projectKeys Key1 --repo Repo1 --exclusionFilePath PATH_TO_FILE --skipSnykMonitoredRepos --json
    ```
