# Variables list and description

{% tabs %}
{% tab title="API 커스텀 PR 템플릿" %}
다음 변수들을 템플릿에서 사용할 수 있습니다.

### <mark style="color:purple;">`jira_ids: string[]`</mark>

PR에 포함된 이슈와 연관된 Jira 티켓 목록입니다. 프로젝트 또는 저장소에 Snyk Jira 통합이 활성화되어 있고 Snyk 이슈가 Jira 티켓과 연결되어 있는지 확인하십시오.

관련 PR에 Jira를 자동으로 연결하려면 커밋 메시지에 연관된 Jira 티켓 목록을 포함하세요.

#### 입력

```json
{
    "data": {
        "attributes": {
            "commit_message": "This pull request is from Snyk and relates to {{ jira_ids }}" 
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 커밋 메시지는 다음과 같습니다:

```json
This pull request is from Snyk and relates to JIRA-1,JIRA-2,JIRA-3
```

이 출력은 제안된 해결책이 세 가지 문제를 성공적으로 해결했음을 나타냅니다. 또한, 각 Jira 티켓에 대한 링크도 포함됩니다.

### <mark style="color:purple;">`snyk_project_url: string`</mark>

이는 Snyk 프로젝트 URL이며, Snyk 프로젝트 페이지에 연결하는 데 사용될 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "자세한 내용을 보려면 Snyk 프로젝트 {{ snyk_project_url }}를 참조하십시오."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
자세한 내용을 보려면 Snyk 프로젝트 https://app.snyk.io/org/my-org/project/xx-xxx-xx-xx
```

이 출력에서 `my-org`는 Snyk 조직 이름이고, `xx-xxx-xx-xx-xxxx`는 프로젝트 또는 리포지토리의 공개 ID입니다.

### <mark style="color:purple;">`snyk_project_name: string`</mark>

이것은 Snyk 프로젝트 이름입니다. Snyk 프로젝트 이름을 설명에 추가할 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "프로젝트 {{ snyk_project_name }}에 수정이 적용되었습니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
프로젝트 my-org/project:filename에 수정이 적용되었습니다.
```

### <mark style="color:purple;">`snyk_org_name: string`</mark>

이것은 Snyk 조직 이름입니다. Snyk 조직 이름을 설명에 추가할 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ snyk_org_name }}에서 수정이 적용되었습니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
my-org에서 수정이 적용되었습니다.
```

### <mark style="color:purple;">`package_name: string`</mark>

이것은 수정되거나 업그레이드되는 패키지의 이름입니다. 여러 개의 패키지가 변경될 경우, 이 변수는 기본적으로 첫 번째 패키지를 나타냅니다.

다음 예시를 따라 PR에서 수정되는 첫 번째 의존성의 패키지 이름을 설명에 표시할 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ package_name }} 수정"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
adm-zip 수정
```

### <mark style="color:purple;">`package_from: string`</mark>

이것은 수정되거나 업그레이드되는 패키지의 버전입니다. 여러 패키지가 변경되는 경우, 이 변수는 첫 번째 패키지의 `from` 버전을 기본값으로 사용합니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ package_from }}에서 이동하여 수정이 적용됩니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
0.4.7에서 이동하여 수정이 적용됩니다.
```

### <mark style="color:purple;">`package_to: string`</mark>

이 패키지는 특정 버전으로 이동하고 있습니다. 여러 패키지가 변경되는 경우, 이 변수는 첫 번째 패키지의 `to` 버전을 기본값으로 사용합니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ package_to }}로 이동하여 수정이 적용됩니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
0.5.2로 이동하여 수정이 적용됩니다.
```

### <mark style="color:purple;">`issue_count: number`</mark>

이것은 PR에서 해결하는 프로젝트나 리포지토리의 문제 수입니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "PR은 {{ issue_count }}개의 문제를 수정합니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
PR은 98개의 문제를 수정합니다.
```

### <mark style="color:purple;">`product_is_container: boolean`</mark>

이 변수는 PR이 Snyk 컨테이너 제품인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 현재 Snyk에는 PR을 열 수 있는 두 가지 다른 제품 유형(Snyk Open Source PR과 Snyk Container PR)이 있습니다. 이 변수를 사용하면 두 가지를 구별할 수 있도록 템플릿을 사용자화하는 데 도움이 됩니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "{{ #product_is_container }} 이 컨테이너 PR은 프로젝트의 취약점을 수정하기 위해 열렸습니다. {{ /product_is_container }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 PR이 오픈 소스 제품이라면, PR의 설명은 다음과 같습니다:

```
이 컨테이너 PR은 프로젝트의 취약점을 수정하기 위해 열렸습니다.
```

### <mark style="color:purple;">`product_is_open_source: boolean`</mark>

이 변수는 PR이 오픈 소스 제품인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 현재 Snyk에는 PR을 열 수 있는 두 가지 다른 제품 유형(오픈 소스 PR과 컨테이너 PR)이 있습니다. 이 변수를 사용하면 두 가지를 구별할 수 있도록 템플릿을 사용자화하는 데 도움이 됩니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "{{ #product_is_open_source }} 이 오픈 소스 PR은 프로젝트의 취약점을 수정하기 위해 열렸습니다. {{ /product_is_open_source }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 PR이 오픈 소스 제품이라면, PR의 설명은 다음과 같습니다:

```
이 오픈 소스 PR은 프로젝트의 취약점을 수정하기 위해 열렸습니다.
```

### <mark style="color:purple;">`is_fix_pr: boolean`</mark>

이 변수는 PR이 백로그 PR인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 예를 들어, 최신 스캔에서 프로젝트나 리포지토리에 새로 도입된 취약점을 수정하기 위해 열린 PR에 해당합니다. 아래 예시에서는 PR이 수정된 PR인 경우에만 설명이 표시되는 것을 볼 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ #is_fix_pr }} 이 PR은 프로젝트의 취약점을 수정하기 위해 열렸습니다. {{ /is_fix_pr }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 PR이 수정된 PR이라면, PR의 설명은 다음과 같습니다:

```yaml
이 PR은 프로젝트의 취약점을 수정하기 위해 열렸습니다.
```

### <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

이 변수는 PR이 백로그 PR인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 예를 들어, 프로젝트나 리포지토리에 이미 존재하는 알려진 취약점을 수정하기 위해 열린 PR에 해당합니다. 아래 예시에서는 PR이 백로그 PR인 경우에만 설명이 표시되는 것을 볼 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ #is_backlog_pr }} 이 PR은 알려진 취약점을 수정하기 위해 열렸습니다. 이 취약점들은 프로젝트의 백로그에서 검색된 것입니다. {{ /is_backlog_pr }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 PR이 백로그 PR이라면, PR의 설명은 다음과 같습니다:

```yaml
이 PR은 알려진 취약점을 수정하기 위해 열렸습니다. 이 취약점들은 프로젝트의 백로그에서 검색된 것입니다.
```

### <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

이 변수는 PR이 업그레이드 PR인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 즉, 취약점과 관계없이 의존성을 최신 버전으로 업그레이드하기 위한 PR입니다. 아래 예시에서는 PR이 업그레이드 PR인 경우에만 설명이 표시되는 것을 볼 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "description": "{{ #is_upgrade_pr }} 이 PR은 우리의 리포지토리가 최신 상태로 유지되도록 열렸습니다. 이 PR은 {{ package_name }}를 버전 {{ package_from }}에서 버전 {{ package_to }}로 업데이트합니다. 가능한 파괴적인 변경 사항에 대한 관련 문서를 검토하십시오. {{ /is_upgrade_pr }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 PR이 업그레이드 PR이라면, PR의 설명은 다음과 같습니다:

```json
이 PR은 우리의 리포지토리가 최신 상태로 유지되도록 열렸습니다. 이 PR은 package-x를 버전 1.0.0에서 버전 2.0.0으로 업데이트합니다. 가능한 파괴적인 변경 사항에 대한 관련 문서를 검토하십시오.
```

### <mark style="color:purple;">`files_changed`</mark>

이 변수는 템플릿에서 PR의 일환으로 변경된 파일을 나열하는 데 사용할 수 있습니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "이 PR에 포함된 변경 사항: {{ files_changed }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 PR이 Maven 프로젝트를 위한 것이고 변경 사항이 `pom.xml` 파일에 있었다면, PR의 설명은 다음과 같을 것입니다.

```
이 PR에 포함된 변경 사항: pom.xml
```

### <mark style="color:purple;">`container.recommended_base_image_name`</mark>

이 변수는 컨테이너 프로젝트에만 해당합니다. 이 PR에서 적용된 권장 기본 이미지의 이름을 표시하는 데 사용할 수 있습니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "우리는 {{ container.recommended_base_image_name }}로 업그레이드를 권장합니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```
우리는 node:xx.xx.x로 업그레이드를 권장합니다.
```

### <mark style="color:purple;">`container.current_base_image_name`</mark>

이 변수는 컨테이너 프로젝트에만 해당합니다. 현재 기본 이미지의 이름을 표시하는 데 사용할 수 있습니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "현재 기본 이미지는: {{ container.current_base_image_name }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```
현재 기본 이미지는: node:xx.xx.x
```

### <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

이 변수는 프로젝트나 리포지토리의 PR 유형(prType)입니다. 이를 사용하여 PR 설명에서 PR 유형을 표시할 수 있습니다.

#### 입력

```json
{
    "data": {
        "attributes": {
            "commit_message": "{{ snyk_pull_request_type}}: for {{ package_name }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 Fix PR을 열었다면, PR의 커밋 메시지는 다음과 같습니다:

```yaml
fix: for package-x
```
{% endtab %}

{% tab title="YAML 파일 사용자화 PR 템플릿" %}
템플릿에서 다음 변수를 사용할 수 있습니다. 이 변수들은 사용자화 가능한 모든 PR 속성에서 사용될 수 있습니다.

### <mark style="color:purple;">`jira_ids: string[]`</mark>

이 변수는 PR에 포함된 이슈와 관련된 Jira 티켓 목록입니다. Snyk Jira 통합이 프로젝트나 리포지토리에 활성화되어 있고, Snyk 이슈가 Jira 티켓에 연결되어 있는지 확인하세요.

Jira를 관련 PR에 자동으로 연결하려면, 커밋 메시지에 관련 Jira 티켓 목록을 포함해야 합니다.

#### 입력

```yaml
commitMessage: |
  이 풀 리퀘스트는 Snyk에서 발생하며 {{ jira_ids }}와 관련이 있습니다.
```

#### 출력

PR의 커밋 메시지는 다음과 같습니다:

```yaml
이 풀 리퀘스트는 Snyk에서 발생하며 JIRA-1,JIRA-2,JIRA-3와 관련이 있습니다.
```

이 출력은 제안된 솔루션이 세 가지 문제를 성공적으로 해결했음을 나타냅니다. 또한 각 Jira 티켓에 대한 링크도 포함됩니다.

### <mark style="color:purple;">`snyk_project_url: string`</mark>

이 변수는 Snyk 프로젝트 URL로, Snyk 프로젝트 페이지에 링크하는 데 사용할 수 있습니다.

#### 입력

```yaml
description: |
  더 많은 세부사항은 Snyk 프로젝트 {{ snyk_project_url }}에서 확인하세요.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
더 많은 세부사항은 Snyk 프로젝트 https://app.snyk.io/org/my-org/project/xx-xxx-xx-xx에서 확인하세요.
```

이 출력에서 `my-org`는 Snyk 조직 이름이며, `xx-xxx-xx-xx-xxxx`는 프로젝트 또는 리포지토리의 공개 ID입니다.

### <mark style="color:purple;">`snyk_project_name: string`</mark>

이 변수는 Snyk 프로젝트 이름입니다. Snyk 프로젝트 이름을 설명에 추가할 수 있습니다.

#### 입력

```yaml
description: |
  프로젝트 {{ snyk_project_name }}에 수정이 적용되었습니다.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
프로젝트 my-org/project:filename에 수정이 적용되었습니다.
```

### <mark style="color:purple;">`snyk_org_name: string`</mark>

이 변수는 Snyk 조직 이름입니다. Snyk 조직 이름을 설명에 추가할 수 있습니다.

#### 입력

```yaml
description: |
  {{ snyk_org_name }}에 의해 수정이 적용되었습니다.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
my-org에 의해 수정이 적용되었습니다.
```

### <mark style="color:purple;">`package_name: string`</mark>

이 변수는 수정되거나 업그레이드되는 패키지의 이름입니다. 여러 패키지가 변경되는 경우, 이 변수는 첫 번째 패키지로 기본 설정됩니다.

다음 예시처럼 PR에서 수정되는 첫 번째 종속성의 패키지 이름을 설명에 표시할 수 있습니다.

#### 입력

```yaml
description: |
  수정된 패키지 {{ package_name }}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
수정된 패키지 adm-zip
```

### <mark style="color:purple;">`package_from: string`</mark>

이 변수는 수정되거나 업그레이드되는 패키지의 버전입니다. 여러 패키지가 변경되는 경우, 이 변수는 첫 번째 패키지의 `from` 버전으로 기본 설정됩니다.

#### 입력

```yaml
description: |
  수정은 {{ package_from }} 버전에서 이동하여 적용됩니다.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
수정은 0.4.7 버전에서 이동하여 적용됩니다.
```

### <mark style="color:purple;">`package_to: string`</mark>

이 변수는 패키지가 특정 버전으로 이동하는 것을 나타냅니다. 여러 패키지가 변경되는 경우, 이 변수는 첫 번째 패키지의 `to` 버전으로 기본 설정됩니다.

#### 입력

```yaml
description: |
  수정은 {{ package_to }} 버전으로 이동하여 적용됩니다.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
수정은 0.5.2 버전으로 이동하여 적용됩니다.
```

### <mark style="color:purple;">`issue_count: number`</mark>

이 변수는 PR에 의해 해결되는 프로젝트 또는 리포지토리의 문제 수입니다.

#### 입력

```yaml
description: |
   이 PR은 {{ issue_count }}개의 문제를 해결합니다.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
이 PR은 98개의 문제를 해결합니다.
```

### <mark style="color:purple;">`product_is_container: boolean`</mark>

이 변수는 PR이 컨테이너 제품인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 현재 Snyk에는 PR을 열 수 있는 두 가지 제품 유형(오픈 소스 PR과 컨테이너 PR)이 있습니다. 이 변수를 사용하면 두 가지 유형을 구분하는 템플릿을 사용자화하는 데 도움이 됩니다.

#### 입력

```
description: |
  {{ #product_is_container }}
  이 컨테이너 PR은 우리의 리포지토리가 최신 상태로 유지되도록 열렸습니다.
  이 PR은 {{ package_name }}을(를) {{ package_from }} 버전에서 {{ package_to }} 버전으로 업데이트합니다.
  가능한 브레이킹 변경 사항에 대한 관련 문서를 검토하십시오.
  {{ /product_is_container }}
```

#### 출력

프로젝트가 컨테이너 프로젝트인 경우, 설명은 다음과 같습니다:

```
  이 컨테이너 PR은 우리의 리포지토리가 최신 상태로 유지되도록 열렸습니다.
  이 PR은 package x을(를) 1 버전에서 2 버전으로 업데이트합니다.
  가능한 브레이킹 변경 사항에 대한 관련 문서를 검토하십시오.
```

### <mark style="color:purple;">`product_is_open_source: boolean`</mark>

이 변수는 PR이 오픈 소스 제품인지 여부에 따라 속성을 사용자화하는 데 사용할 수 있습니다. 현재 Snyk에는 PR을 열 수 있는 두 가지 제품 유형(오픈 소스 PR과 컨테이너 PR)이 있습니다. 이 변수를 사용하면 두 가지 유형을 구분하는 템플릿을 사용자화하는 데 도움이 됩니다.

#### 입력

```
description: |
  {{ #product_is_open_source }}
  이 오픈 소스 PR은 우리의 리포지토리가 최신 상태로 유지되도록 열렸습니다.
  이 PR은 {{ package_name }}을(를) {{ package_from }} 버전에서 {{ package_to }} 버전으로 업데이트합니다.
  가능한 브레이킹 변경 사항에 대한 관련 문서를 검토하십시오.
  {{ /product_is_open_source }}
```

#### 출력

프로젝트가 오픈 소스 프로젝트인 경우, 설명은 다음과 같습니다:

```
  이 오픈 소스 PR은 우리의 리포지토리가 최신 상태로 유지되도록 열렸습니다.
  이 PR은 package x을(를) 1 버전에서 2 버전으로 업데이트합니다.
  가능한 브레이킹 변경 사항에 대한 관련 문서를 검토하십시오.
```

### <mark style="color:purple;">`is_fix_pr: boolean`</mark>

이 변수는 풀 리퀘스트가 수정 PR인지 여부를 확인하는 데 사용됩니다. 예를 들어, 최신 스캔에서 프로젝트나 리포지토리에 새로 도입된 취약점을 수정하기 위해 열린 PR인지 확인하는 데 사용됩니다.

#### 입력

```yaml
description: |
  이 PR은 수정 PR인가요? {{ is_fix_pr }}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
이 PR은 수정 PR인가요? true
```

### <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

이 변수는 풀 리퀘스트가 백로그 PR인지 여부를 확인하는 데 사용됩니다. 예를 들어, 이미 프로젝트나 리포지토리에 존재하는 알려진 취약점을 수정하기 위해 열린 PR인지 확인하는 데 사용됩니다.

#### 입력

```yaml
description: |
  이 PR은 백로그 PR인가요? {{ is_backlog_pr }}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
이 PR은 백로그 PR인가요? false
```

### <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

이 변수는 풀 리퀘스트가 업그레이드 PR인지 여부를 확인하는 데 사용됩니다. 예를 들어, 취약점과 상관없이 의존성을 더 최신 버전으로 업그레이드하기 위해 열린 PR인지 확인하는 데 사용됩니다.

#### 입력

```yaml
description: |
  이 PR은 업그레이드 PR인가요? {{ is_upgrade_pr }}
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
이 PR은 업그레이드 PR인가요? false
```

### <mark style="color:purple;">`files_changed`</mark>

이 변수는 템플릿에서 풀 리퀘스트의 일부로 변경된 파일을 나열하는 데 사용될 수 있습니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "이 PR에 포함된 변경 사항: {{ files_changed }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

만약 풀 리퀘스트가 Maven 프로젝트에 대한 것이고, 변경 사항이 `pom.xml` 파일에 있었다면, PR의 설명은 다음과 같습니다:

```
이 PR에 포함된 변경 사항: pom.xml
```

### <mark style="color:purple;">`container.recommended_base_image_name`</mark>

이 변수는 컨테이너 프로젝트에서만 사용됩니다. 이 변수는 PR에 적용된 추천 기본 이미지의 이름을 표시하는 데 사용할 수 있습니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "우리는 {{ container.recommended_base_image_name }}로 업그레이드를 권장합니다."
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```
우리는 node:xx.xx.x로 업그레이드를 권장합니다.
```

### <mark style="color:purple;">`container.current_base_image_name`</mark>

이 변수는 컨테이너 프로젝트에서만 사용됩니다. 이 변수는 현재 기본 이미지를 표시하는 데 사용할 수 있습니다.

#### 입력

```
{
    "data": {
        "attributes": {
            "description": "현재 기본 이미지는: {{ container.current_base_image_name }}"
        },
        "type": "pull_request_template"
    }
}
```

#### 출력

PR의 설명은 다음과 같습니다:

```
현재 기본 이미지는: node:xx.xx.x
```

### <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

이 변수는 프로젝트나 리포지토리의 prType을 나타냅니다. 이를 사용하여 풀 리퀘스트 설명에서 PR 유형을 표시할 수 있습니다.

#### 입력

```yaml
description: |
  이것은 {{ snyk_pull_request_type }} 풀 리퀘스트입니다.
```

#### 출력

PR의 설명은 다음과 같습니다:

```yaml
이것은 fix 풀 리퀘스트입니다.
```
{% endtab %}
{% endtabs %}
