# 변수 목록과 설명

{% tabs %}
{% tab title="API 사용자 정의 PR 템플릿" %}


템플릿에서 다음 변수를 사용할 수 있습니다.

## <mark style="color:purple;">`jira_ids: string[]`</mark>

풀 리퀘스트에 포함된 이슈에 연결된 Jira 티켓의 목록입니다. 풀 리퀘스트 내에서 추천된 문제가 성공적으로 해결된 경우를 나타냅니다. 모든 Jira 티켓에 대한 링크가 포함되어 있습니다.

## <mark style="color:purple;">`snyk_project_url: string`</mark>

이것은 Snyk 프로젝트 URL이며 Snyk 프로젝트 페이지로 연결할 수 있습니다.

## <mark style="color:purple;">`snyk_project_name: string`</mark>

이것은 Snyk 프로젝트 이름입니다. 설명에 Snyk 프로젝트 이름을 추가할 수 있습니다.

## <mark style="color:purple;">`snyk_org_name: string`</mark>

이것은 Snyk 조직 이름입니다. 설명에 Snyk 조직 이름을 추가할 수 있습니다.

## <mark style="color:purple;">`package_name: string`</mark>

수정되거나 업그레이드되는 패키지의 이름입니다.

## <mark style="color:purple;">`package_from: string`</mark>

수정되거나 업그레이드되는 패키지의 버전입니다.

## <mark style="color:purple;">`package_to: string`</mark>

패키지가 이 특정 버전으로 전환 중입니다.

## <mark style="color:purple;">`issue_count: number`</mark>

PR이 해결할 문제의 수입니다.

## <mark style="color:purple;">`product_is_container: boolean`</mark>

이 변수는 PR이 Snyk Container 제품인지를 기반으로 속성을 사용자 정의하는 데 사용할 수 있습니다.

## <mark style="color:purple;">`product_is_open_source: boolean`</mark>

이 변수는 PR이 오픈 소스 제품인지에 따라 속성을 사용자 정의하는 데 사용할 수 있습니다.

## <mark style="color:purple;">`is_fix_pr: boolean`</mark>

이 변수는 PR이 새로운 취약점을 수정하기 위해 열린 백로그 PR인지에 따라 속성을 사용자 정의하는 데 사용할 수 있습니다.

## <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

이 변수는 PR이 알려진 취약점을 수정하기 위해 열린 백로그 PR인지에 따라 속성을 사용자 정의하는 데 사용할 수 있습니다.

## <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

이 변수는 PR이 업그레이드 PR인지에 따라 속성을 사용자 정의하는 데 사용할 수 있습니다.

## <mark style="color:purple;">`files_changed`</mark>

풀 리퀘스트로 변경된 파일을 나열하는 데 사용할 수 있는 변수입니다.

## <mark style="color:purple;">`container.recommended_base_image_name`</mark>

컨테이너 프로젝트 전용입니다. 이 변수는 PR에서 적용된 권장 기본 이미지의 이름을 표시하는 데 사용할 수 있습니다.

## <mark style="color:purple;">`container.current_base_image_name`</mark>

컨테이너 프로젝트 전용입니다. 현재 기본 이미지를 표시하기 위해 사용할 수 있습니다.

## <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

프로젝트 또는 리포지토리의 prType입니다. 풀 리퀘스트 설명에서 PR 유형을 표시하기 위해 사용할 수 있습니다.
{% endtab %}

{% tab title="YAML 파일 사용자 정의 PR 템플릿" %}


템플릿에서 다음 변수를 사용할 수 있습니다. 이러한 변수는 사용자 정의 가능한 PR 속성 중 어디에서든 사용할 수 있습니다.

## <mark style="color:purple;">`jira_ids: string[]`</mark>

풀 리퀘스트에 포함된 이슈에 연결된 Jira 티켓의 목록입니다.

## <mark style="color:purple;">`snyk_project_url: string`</mark>

이것은 Snyk 프로젝트 URL이며 Snyk 프로젝트 페이지로 연결할 수 있습니다.

## <mark style="color:purple;">`snyk_project_name: string`</mark>

이것은 Snyk 프로젝트 이름입니다. 설명에 Snyk 프로젝트 이름을 추가할 수 있습니다.

## <mark style="color:purple;">`snyk_org_name: string`</mark>

이것은 Snyk 조직 이름입니다. 설명에 Snyk 조직 이름을 추가할 수 있습니다.
{% endtab %}
{% endtabs %}  ### 입력

```yaml
description: |
  Snyk 조직 이름에 의해 적용된 수정
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
수정된 부분: my-org
```

## <mark style="color:purple;">`package_name: string`</mark>

이것은 수정되거나 업그레이드되는 패키지의 이름입니다. 여러 패키지가 변경된 경우, 이 변수는 기본적으로 첫 번째 패키지로 설정됩니다.

다음 예를 따라 PR에서 수정되는 첫 번째 종속성의 패키지 이름을 설명에 표시하십시오.&#x20;

### 입력

```yaml
description: |
  수정: {{ package_name }}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
수정: adm-zip
```

## <mark style="color:purple;">`package_from: string`</mark>

이것은 수정되거나 업그레이드되는 패키지의 버전입니다. 여러 패키지가 변경된 경우, 이 변수는 기본적으로 첫 번째 패키지의 `from` 버전으로 설정됩니다.

### 입력

```yaml
description: |
  수정된 부분: {{ package_from}}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
수정된 부분: 0.4.7
```

## <mark style="color:purple;">`package_to: string`</mark>

이 패키지는 특정 버전으로 전환됩니다. 여러 패키지가 변경된 경우, 이 변수는 기본적으로 첫 번째 패키지의 `to` 버전으로 설정됩니다.

### 입력

```yaml
description: |
  수정된 부분: {{ package_to}}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
수정된 부분: 0.5.2
```

## <mark style="color:purple;">`issue_count: number`</mark>

이것은 PR에서 처리되는 프로젝트나 저장소의 이슈 수입니다.&#x20;

### 입력

```yaml
description: |
   해당 PR은 {{ issue_count }}개의 이슈를 해결합니다.
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
해당 PR은 98개의 이슈를 해결합니다.
```

## <mark style="color:purple;">`product_is_container: boolean`</mark>

이 변수는 PR이 컨테이너 제품인지 여부에 따라 속성을 맞춤 설정하는 데 사용될 수 있습니다. 현재 Snyk에서 두 가지 다른 제품 유형이 PR을 열 수 있습니다(오픈 소스 PR 및 컨테이너 PR). 이 변수를 사용하여 두 가지 제품을 구분할 수 있습니다.

### 입력

```
description: |
  {{ #product_is_container }}
  이 컨테이너 PR은 우리의 저장소를 최신 상태로 유지하기 위해 열었습니다.
  패키지 x를 버전 {{ package_from }}에서 버전 {{ package_to }}로 업데이트합니다.
  가능한 호환성 문제에 대한 문서를 검토하십시오.
  {{ /product_is_container }}
```

### 출력

귀하의 프로젝트가 컨테이너 프로젝트인 경우, 설명은 다음과 같을 것입니다:

```
  이 컨테이너 PR은 우리의 저장소를 최신 상태로 유지하기 위해 열었습니다.
  패키지 x를 버전 1에서 버전 2로 업데이트합니다.
  가능한 호환성 문제에 대한 문서를 검토하십시오.
```

## <mark style="color:purple;">`product_is_open_source: boolean`</mark>

이 변수는 PR이 오픈 소스 제품인지 여부에 따라 속성을 맞춤 설정하는 데 사용될 수 있습니다. 현재 Snyk에서 두 가지 다른 제품 유형이 PR을 열 수 있습니다(오픈 소스 PR 및 컨테이너 PR). 이 변수를 사용하여 두 가지 제품을 구분할 수 있습니다.

### 입력

```
description: |
  {{ #product_is_open_source }}
  이 오픈 소스 PR은 우리의 저장소를 최신 상태로 유지하기 위해 열었습니다.
  패키지 x를 버전 {{ package_from }}에서 버전 {{ package_to }}로 업데이트합니다.
  가능한 호환성 문제에 대한 문서를 검토하십시오.
  {{ /product_is_open_source }}
```

### 출력

귀하의 프로젝트가 오픈 소스 프로젝트인 경우, 설명은 다음과 같을 것입니다:

```
  이 오픈 소스 PR은 우리의 저장소를 최신 상태로 유지하기 위해 열었습니다.
  패키지 x를 버전 1에서 버전 2로 업데이트합니다.
  가능한 호환성 문제에 대한 문서를 검토하십시오.
```

## <mark style="color:purple;">`is_fix_pr: boolean`</mark>

이는 PR이 새로운 취약성을 수정하기 위해 열리는 등 프로젝트나 저장소에 도입된 새로운 취약성을 수정하기 위해 열린 수정 PR인지 여부를 확인합니다.

### 입력

```yaml
description: |
  이 PR은 수정 PR입니까? {{ is_fix_pr }}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
이것은 수정 PR입니까? true
```

## <mark style="color:purple;">`is_backlog_pr: boolean`</mark>

이는 PR이 이미 프로젝트나 저장소에 있는 알려진 취약성을 수정하기 위해 열린 백로그 PR인지 여부를 확인합니다.

### 입력

```yaml
description: |
  이 PR은 백로그 PR입니까? {{ is_backlog_pr }}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
이것은 백로그 PR입니까? false
```

## <mark style="color:purple;">`is_upgrade_pr: boolean`</mark>

이는 의존성을 보완적으로 업그레이드하기 위해 열리는 등 새로운 취약성과 무관하게 의존성을 더 높은 버전으로 업그레이드하기 위해 열린 업그레이드 PR인지 여부를 확인합니다.

### 입력

```yaml
description: |
  이 PR은 업그레이드 PR입니까? {{ is_upgrade_pr }}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
이것은 업그레이드 PR입니까? false
```



## <mark style="color:purple;">`files_changed`</mark>

이 변수를 사용하여 PR의 일부로 변경된 파일을 나열할 수 있습니다.&#x20;

### 입력

```
{
    "data": {
        "attributes": {
            "description": "변경된 내용: {{ files_changed }}"
            
        },
        "type": "pull_request_template"
    }
}
```



### 출력

Maven 프로젝트용 PR이며 변경된 사항이 pom.xml 파일인 경우, PR의 설명은 다음과 같을 것입니다.&#x20;

```
변경된 내용: pom.xml
```

## <mark style="color:purple;">`container.recommended_base_image_name`</mark>

이 변수는 컨테이너 프로젝트 전용입니다. PR에 적용된 권장 기본 이미지의 이름을 표시하는 데 사용할 수 있습니다.

### 입력

```
{
    "data": {
        "attributes": {
            "description": "다음을 권장합니다: {{ container.recommended_base_image_name }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```
다음을 권장합니다: node:xx.xx.x
```

## <mark style="color:purple;">`container.current_base_image_name`</mark>

이 변수는 컨테이너 프로젝트 전용입니다. 현재 기본 이미지를 표시하는 데 사용할 수 있습니다.

### 입력

```
{
    "data": {
        "attributes": {
            "description": "현재 기본 이미지: {{ container.current_base_image_name }}"
            
        },
        "type": "pull_request_template"
    }
}
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```
현재 기본 이미지: node:xx.xx.x
```

## <mark style="color:purple;">`snyk_pull_request_type: prType (fix, upgrade, backlog, unknown)`</mark>

이것은 귀하의 프로젝트나 저장소의 prType입니다. PR 설명에서 PR 유형을 표시하는 데 사용할 수 있습니다.

### 입력

```yaml
description: |
  이것은 {{ snyk_pull_request_type }} PR입니다
```

### 출력

당신의 PR 설명은 다음과 같을 것입니다:&#x20;

```yaml
이것은 수정 PR입니다
```