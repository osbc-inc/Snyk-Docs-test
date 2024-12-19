---
description: GitLab 옵션 목록 및 예제
---

# GitLab - 예시

`snyk-scm-contributors-count gitlab` 명령어에 사용 가능한 옵션 목록은 다음과 같습니다:

```
  --version                 버전 번호 표시                        [부울]
  --help                    도움말 표시                                 [부울]
  --token                   GitLab 토큰                               [필수]
  --url                     [선택 사항] GitLab 호스트 사용자 지정 URL. 호스트를 제공하지 않은 경우 기본값은 https://gitlab.com/
  --groups                  [선택 사항] 기여자 수를 계산할 Gitlab 그룹 이름
                            *참고* 하위 수준 그룹을 위해서는 최하위 그룹 이름을 제공하세요                                           
  --project                 [선택 사항] 기여자 수를 계산할 네임스페이스가 있는 GitLab 프로젝트 경로
  --exclusionFilePath       [선택 사항] 제외 목록 파일 경로
  --json                    [선택 사항] JSON 출력, "consolidateResults" 명령어를 사용할 때 필요
```

## **명령 실행 전**
GitLab 토큰을 얻거나 새로 생성하려면 [가이드](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html)를 참조하세요.

**참고:** 토큰이 리포지토리의 읽기 액세스 권한이 있는지 확인하세요.

## 명령 실행

다음은 사용 수준 및 옵션을 고려하세요:

### 사용 수준

* GitLab의 모든 그룹과 프로젝트에 대한 커밋을 가져오려면 GitLab 토큰을 제공(및 GitLab Enterprise의 서버 URL을 제공하려면):

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --url URL
    ```
* GitLab의 일부 그룹 및 그들의 프로젝트에 대한 커밋을 가져오려면 GitLab 토큰과 쉼표로 구분된 그룹 이름을 제공하세요:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP1,GROUP2
    ```

{% hint style="info" %}
중첩된 그룹의 경우 가장 낮은 수준 그룹 이름을 제공해야 합니다. 예를 들어, `최상위그룹/중간수준그룹/최하위그룹`의 경우 "최하위그룹"만 `--groups` 플래그와 함께 제공하세요.
{% endhint %}

* GitLab의 특정 프로젝트에 대한 커밋을 가져오려면 GitLab 토큰 및 **한 개** 그룹 이름 및 **한 개** 프로젝트 이름을 제공하세요:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --groups GROUP --project PROJECT
    ```

### 옵션

* 특정 기여자를 커밋 계산에서 제외하려면 무시할 이메일을 포함한 제외 파일을 추가하고 `--exclusionFilePath`를 파일 경로와 함께 적용하세요:

```
snyk-scm-contributors-count gitlab --token TOKEN --projectKeys ID1,ID2,Path1/Namespace1 --exclusionFilePath PATH_TO_FILE
```

* JSON 형식으로 출력을 설정하려면 `--json` 플래그를 추가하세요:

    ```
    snyk-scm-contributors-count gitlab --token TOKEN --json
    ```
* 상세한 출력을 위해 디버그 모드에서 실행하려면 `DEBUG=snyk*`를 접두사로 붙이세요:

```
DEBUG=snyk* snyk-scm-contributors-count gitlab --token TOKEN --url URL --exclusionFilePath PATH_TO_FILE --json
```