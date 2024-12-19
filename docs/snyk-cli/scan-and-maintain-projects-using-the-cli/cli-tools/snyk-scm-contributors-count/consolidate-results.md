---
description: consolidateResults 명령어 사용 방법
---

# 결과 통합

## consolidateResults 명령어

SCM-Contributors-Count 도구를 사용할 때 두 개 이상의 소스 제어 관리자(SCM)와 작업할 수 있습니다. 각 SCM에 대해 기여자 수를 얻으려면 각 SCM에 대해 별도의 명령어를 실행해야 합니다.

예를 들어, GitHub 리포지토리와 Bitbucket 프로젝트에 커밋하는 기여자가 있다면, 해당 기여자의 세부 정보가 두 SCM 모두의 출력에 나타날 것입니다.

모든 SCM을 통틀어 전체 기여자를 확인하려면 `consolidateResults` 명령어를 사용하세요.

이 명령어를 사용하면 각 SCM에 대한 `snyk-scm-contributors-count` 명령의 여러(`json`) 출력을 하나의 파일로 통합할 수 있습니다. 이 파일에는 고유한 기여자 수와 모든 SCM의 리포지토리 총합이 포함됩니다.

`consolidateResults` 명령어에 대한 다음 옵션을 사용할 수 있습니다:

```
  --version                 버전 번호 표시                             [boolean]
  --help                    도움말 표시                                   [boolean]
  --folderPath              json 출력을 포함하는 폴더의 경로                    [필수]
```

## 명령어 실행하기

* `--json` 플래그와 함께 각 리포지토리에 대해 `snyk-scm-contributors-count` 명령어를 실행하고 출력을 지정된 폴더로 보내세요. 예를 들면:

```
snyk-scm-contributors-count github --token 토큰 --json > 경로/파일명
snyk-scm-contributors-count github-enterprise --token 토큰 --json > 경로/다른파일명
```

* `consolidateResults` 명령어를 실행하고 `--folderPath` 플래그와 각 SCM 결과가 들어 있는 읽기/쓰기 가능한 폴더의 경로를 지정하세요.

```
snyk-scm-contributors-count consolidateResults --folderPath 경로
```

* 도구는 그런 다음 폴더에서 유효한 파일을 찾아 파일의 내용을 읽고, 읽힌 모든 파일의 합쳐진 고유한 결과를 포함하는 새 파일을 만들어 `consolidated-results.json`이라는 이름으로 저장합니다.