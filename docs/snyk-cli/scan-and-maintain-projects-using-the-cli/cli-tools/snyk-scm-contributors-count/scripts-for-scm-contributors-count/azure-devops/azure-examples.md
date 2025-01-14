---
description: Azure에 대한 옵션 목록과 예제
---

# Azure - 예제

`snyk-scm-contributors-count azure devops` 명령에 대한 다음 옵션들을 사용할 수 있습니다:

```
  --version                 버전 번호 표시                               [boolean]
  --help                    도움말 표시                                 [boolean]
  --token                   Azure DevOps 토큰                         [required]
  --org                     Azure DevOps의 당신의 Org 이름, 예를 들어, https://dev.azure.com/{OrgName}           [required]
  --projectKeys             [선택 사항] 카운트 대상인 Azure DevOps 프로젝트 키/이름
  --repo                    [선택 사항] 카운트하기 위해 특정 저장소
  --exclusionFilePath       [선택 사항] 제외 목록 파일 경로
  --json                    [선택 사항] JSON 출력, "consolidateResults" 명령을 사용할 때 필요
  --skipSnykMonitoredRepos  [선택 사항] Snyk 모니터링 중인 저장소 건너뛰고 모든 저장소의 기여자 수 계산
  --importConfDir           [선택 사항] 비모니터링 저장소용 가져오기 파일 생성: 생성된 가져오기 파일을 저장할 유효한 폴더 경로
  --importFileRepoType      [선택 사항] importConfDir 플래그와 함께 사용: 가져오기 파일에 추가할 리포지토리 유형 지정. 옵션: all/private/public. 기본값: all
```

## 명령 실행 전

1. `SNYK_TOKEN`을 내보냅니다 (`Snyk이 이미 모니터링하는 저장소의 기여자만 가져오려는 경우`):
   * 토큰이 그룹 수준 액세스 권한이 있는지 확인하거나 그룹 수준 액세스 권한이 있는 서비스 계정의 토큰을 사용합니다. 서비스 계정을 만드는 방법에 대해 자세히 알아보려면 [서비스 계정 설정 방법](https://docs.snyk.io/features/integrations/managing-integrations/service-accounts#how-to-set-up-a-service-account)을 참조하십시오.
   * 토큰 값 복사합니다.
   *   환경에 토큰 내보내기:

       ```
       export SNYK_TOKEN=<YOUR-SNYK-TOKEN>
       ```
2. Azure DevOps 토큰 및 Org 가져오기:
   *   가이드를 통해 토큰을 만듭니다 (없는 경우) [가이드](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page)를 참조하여 만듭니다.

       **참고**: 토큰이 리포지토리에 대한 읽기 액세스 권한이 있는지 확인하십시오.
   * Azure의 왼쪽 패널에 나열된 Azure의 Org 이름을 찾으십시오 [Azure DevOps 사이트](https://dev.azure.com).

## 명령 실행

다음 사용 수준 및 옵션을 고려하십시오:

### 사용 수준

*   Azure의 내 Org에 있는 모든 프로젝트 및 저장소의 커밋을 가져오려면 Azure 토큰 및 Azure Org만 제공하세요:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG
    ```
*   Azure의 내 Org에 있는 일부 프로젝트 및 저장소의 커밋을 가져오려면 Azure 토큰, Azure Org 및 쉼표로 구분된 프로젝트 키/들을 제공하세요:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1,Key2...
    ```
*   Azure의 내 Org에 있는 특정 저장소의 커밋을 가져오려면 Azure 토큰, Azure Org, 프로젝트 키 및 저장소 이름을 제공하세요:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1
    ```

### 옵션

*   `Snyk에서 이미 모니터링하는 저장소를 고려하지 않고 Azure에서 모든 커밋을 가져오려면` `--skipSnykMonitoredRepos` 플래그를 추가하세요.

    Snyk에서 모니터링 중이 아닌 Azure의 저장소가 있을 수 있습니다. 이 플래그를 사용하여 Snyk에서 모니터링 중인 저장소를 확인하지 않고 직접 Azure에서 커밋을 가져올 수 있습니다.

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --skipSnykMonitoredRepos
    ```
*   커밋에서 제외할 일부 기여자를 제외하려면 무시할 이메일 (새 줄로 구분)이 포함된 제외 파일을 추가하고 해당 파일 경로로 `--exclusionFilePath`를 적용하세요:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1,Key2 --exclusionFilePath 파일_경로
    ```
*   출력 형식을 json으로 설정하려면 `--json` 플래그를 추가하세요:

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1 --json
    ```
*   모니터링 중이 아닌 저장소를 위해 가져오기 파일을 생성하려면 유효한(쓰기 가능한) 폴더 경로를 사용하여 `--importConfDir` 플래그를 추가하고 가져오기 파일에 추가할 리포지토리 유형(`all`/`private`/`public`, 기본값: `all`)을 사용하여 `--importFileRepoType` 플래그를 추가하십시오. 이러한 플래그는 `--repo` 플래그와 **함께 설정할 수 없음**에 유의하십시오.

    ```
    snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --importConfDir 유효한_폴더_경로 --importFileRepoType private/public/all
    ```

    이러한 플래그에 대한 자세한 내용은 [가져오기 파일 생성 및 사용 페이지](../../creating-and-using-the-import-file.md)를 참조하세요.
*   자세한 출력을 위해 디버그 모드로 실행하려면 명령을`DEBUG=snyk*`로 시작하세요:

    ```
    DEBUG=snyk* snyk-scm-contributors-count azure-devops --token AZURE-TOKEN --org AZURE-ORG --projectKeys Key1 --repo Repo1 --exclusionFilePath 파일_경로 --skipSnykMonitoredRepos --json
    ```
