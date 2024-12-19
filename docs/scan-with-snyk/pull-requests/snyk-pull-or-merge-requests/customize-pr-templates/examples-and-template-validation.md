# 사용 예 및 템플릿 유효성 검사

## 사용자 정의 PR 템플릿 예시

다음 예제 템플릿은 PR 템플릿에서 변수를 사용하는 방법을 보여줍니다.

### API 사용자 정의 PR 템플릿 예시

```json
{
    "data": {
        "attributes": {
            "title": "[{{ snyk_pull_request_type }}] for {{ package_name }}",
            "commit_message": "{{ snyk_pull_request_type}}: for {{ package_name }}",
            "description": "Moving package {{ package_name }} from {{ package_from }} to {{ package_to }}\nFixes {{ issue_count }} issues\nFor more details see {{ snyk_project_url }}\nProject {{ snyk_project_name }}\nOrg {{ snyk_org_name }}"
        },
        "type": "pull_request_template"
    }
}
```

### YAML 사용자 정의 PR 템플릿 예시

```yaml
title: This PR fixes {{ issue_count }} issues
commitMessage: "fix: {{ issue_count }} Snyk issues"
description: |
  {{ #is_upgrade_pr }}
  This PR has been opened to make sure our repositories are kept up-to-date.
  It updates {{ package_name }} from version {{ package_from }} to version {{ package_to }}.
  Review relevant docs for possible breaking changes.
  {{ /is_upgrade_pr }}
  
  **Tickets**
  {{ #jira_ids }}
  - Fixes {{ . }}
  {{ /jira_ids }}
  
  To find more details, see the Snyk project [{{ snyk_project_name }}]({{ snyk_project_url }})
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