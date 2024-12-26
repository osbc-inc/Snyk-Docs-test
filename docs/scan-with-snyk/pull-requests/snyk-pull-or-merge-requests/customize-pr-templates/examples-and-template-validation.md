# 사용 예 및 템플릿 유효성 검사

## 사용자 정의 PR 템플릿 예시

다음 예제 템플릿은 PR 템플릿에서 변수를 사용하는 방법을 보여줍니다.

### API 사용자 정의 PR 템플릿 예시

```json
{
    "data": {
        "attributes": {
            "title": "[] for ",
            "commit_message": ": for ",
            "description": "Moving package  from  to \nFixes  issues\nFor more details see \nProject \nOrg "
        },
        "type": "pull_request_template"
    }
}
```

### YAML 사용자 정의 PR 템플릿 예시

```yaml
title: This PR fixes  issues
commitMessage: "fix:  Snyk issues"
description: |
  
  This PR has been opened to make sure our repositories are kept up-to-date.
  It updates  from version  to version .
  Review relevant docs for possible breaking changes.
  
  
  **Tickets**
  
  - Fixes 
  
  
  To find more details, see the Snyk project []()
```

## 사용자 정의 PR 템플릿 유효성 검사

각 섹션에 대한 설명을 따라 템플릿의 정확성을 검사할 수 있습니다.

### API 템플릿 유효성 검사

API 구성에 대한 템플릿의 정확성을 다음 단계에 따라 확인할 수 있습니다:

1. 에러 응답이 반환되면 사용자 정의 내용에 문제가 있습니다. 문제를 해결하고 API 성공 응답을 받을 때까지 기다립니다.&#x20;
2. PR을 열고 사용자 지정 입력이 사용되는지 확인합니다.&#x20;

### YAML 템플릿 유효성 검사

다음 단계에 따라 템플릿의 정확성을 확인할 수 있습니다:

1. 템플릿 파일을 프로젝트(저장소)에 추가합니다.&#x20;
2. PR을 열고 사용자 지정 입력이 사용되는지 확인합니다.&#x20;