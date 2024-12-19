---
description: GitHub 옵션 목록 및 예시
---

# GitHub - 예시

사용 가능한 옵션:

```
  --version                 버전 번호 표시                                [boolean]
  --help                    도움말 표시                                   [boolean]
  --token                   GitHub 토큰                                  [required]
  --orgs                    [Optional] GitHub 조직 목록, 콤마로 구분하여 나열한
                            해당 조직의 저장소에 대한 기여자 수를 가져오기 위함
  --repo                    [Optional] 특정 저장소에 대해 카운트할 때 사용
  --exclusionFilePath       [Optional] 제외 목록 파일 경로
  --json                    [Optional] JSON 출력, "consolidateResults" 명령어를 사용할 때 필요
```

## 명령어 실행 전

GitHub 토큰을 얻거나 새로 생성하려면 [가이드](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)를 참조하세요.

**참고:** 토큰이 저장소에 대한 읽기 액세스 권한이 있는지 확인하세요.

## 명령어 실행

다음과 같은 사용 수준과 옵션을 고려해 주세요:

### 사용 수준

*   GitHub의 모든 조직의 모든 저장소에 대한 커밋을 얻으려면 GitHub 토큰을 제공하세요:

    ```
    snyk-scm-contributors-count github --token 토큰
    ```

*   GitHub에서 특정 조직 및 해당 저장소의 커밋을 얻으려면 GitHub 토큰과 조직 이름을 콤마로 구분하여 제공하세요:

    ```
    snyk-scm-contributors-count github --token 토큰 --orgs 조직_일,조직_이,조직_삼
    ```

*   GitHub에서 특정 저장소의 커밋을 얻으려면 GitHub 토큰과 한 개의 조직 이름 및 한 개의 저장소 이름을 제공하세요:

    ```
    snyk-scm-contributors-count github --token 토큰 --orgs 조직 --repo 저장소
    ```

### 옵션

* 커밋에서 일부 기여자를 제외하려면 제외할 이메일을 포함하는 제외 파일을 추가하고 해당 파일 경로를 `--exclusionFilePath`에 적용하세요:

```
snyk-scm-contributors-count github --token 토큰 --orgs 조직_일,조직_이 --exclusionFilePath 파일_경로
```

*   결과를 json 형식으로 설정하려면 `--json` 플래그를 추가하세요:
    
    ```
    snyk-scm-contributors-count github --token 토큰 --json
    ```

* 자세한 출력을 위해 디버그 모드로 실행하려면 `DEBUG=snyk*`를 접두어로 사용하세요:

```
DEBUG=snyk* snyk-scm-contributors-count github --token 토큰 --orgs 조직 --repo 저장소 --exclusionFilePath 파일_경로 --json
```