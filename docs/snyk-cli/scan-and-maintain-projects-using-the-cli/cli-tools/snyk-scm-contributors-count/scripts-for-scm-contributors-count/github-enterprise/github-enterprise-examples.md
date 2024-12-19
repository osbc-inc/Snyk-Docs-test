---
description: 옵션 목록 및 일부 예시
---

# Github Enterprise - 예시

다음은 `snyk-scm-contributors-count github-enterprise` 명령에 사용 가능한 옵션 목록입니다:

```
  --version                 버전 번호 표시                         [boolean]
  --help                    도움말 표시                             [boolean]
  --token                   Github Enterprise 토큰                  [필수]
  --url                     사용자의 GitHub 호스트 맞춤 URL, 
                            예: https://ghe.prod.company.org/     [필수]
  --orgs                    [선택적] 쉼표로 구분된 GitHub Enterprise 조직 목록
                            리포지토리의 기여자를 가져오고 계산하는데 사용됨              
  --repo                    [선택적] 특정 리포지토리만 계산하는 데 사용
  --fetchAllOrgs            [선택적] 활성화되면 토큰이 접근할 수 있는 모든 조직을 가져옴
                            권한이 있는 조직만 가져오는 대신.
  --exclusionFilePath       [선택적] 제외 목록 파일 경로
  --json                    [선택적] JSON 출력, "consolidateResults" 명령을 사용할 때 필요
```

## **명령 실행 전**

GitHub Enterprise 토큰을 가져오거나 [가이드](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)를 통해 새로운 토큰을 생성합니다.

**참고:** 토큰이 리포지토리에 대한 읽기 액세스를 가져야 합니다.

## 명령 실행

다음과 같은 사용 수준 및 옵션을 고려하십시오:

### 사용 수준

*   GitHub Enterprise의 모든 org의 모든 리포지토리의 커밋을 가져오려면 GitHub Enterprise 토큰을 제공합니다:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL
    ```
*   GitHub Enterprise에서 특정 org 및 해당 리포지토리의 커밋을 가져오려면 GitHub Enterprise 토큰 및 쉼표로 구분된 org 이름을 제공합니다:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG_ONE,ORG_TWO,ORG_THREE
    ```
*   GitHub Enterprise에서 한 개의 리포지토리의 커밋만 가져오려면 GitHub Enterprise 토큰, org 이름 한 개, 리포지토리 이름 한 개를 제공합니다:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO
    ```

### 옵션

* GitHub Enterprise의 모든 org를 매핑하고 권한이 있는 것뿐만 아니라 추가하려면 `--fetchAllOrgs` 플래그를 추가합니다:

```
snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --fetchAllOrgs
```

*   커밋에서 제외할 기여자가 있는 경우 제외할 이메일을 포함하는 제외 파일을 추가하고 그 파일 경로와 함께 `--exclusionFilePath`를 적용합니다:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG_ONE,ORG_TWO --exclusionFilePath PATH_TO_FILE
    ```
*   출력을 json 형식으로 설정하려면 `--json` 플래그를 추가합니다:

    ```
    snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --json
    ```
* 자세한 출력을 위해 디버그 모드로 실행하려면 `DEBUG=snyk*`를 접두사로 붙입니다 :

```
DEBUG=snyk* snyk-scm-contributors-count github-enterprise --token TOKEN --url HOST_URL --orgs ORG --repo REPO --exclusionFilePath PATH_TO_FILE --json
```