# 패키지를 위한 문제 목록

Snyk REST API 엔드포인트인 [패키지를 위한 문제 목록](../reference/issues.md#orgs-org_id-packages-purl-issues)은 `purl`을 사용하여 패키지에 대한 모든 직접(비전달)적인 취약점을 가져올 수 있습니다. 이 `purl`은 [패키지 URL 명세](https://github.com/package-url/purl-spec)에 정의된 것처럼, 생태계를 가로질러 소프트웨어 패키지를 식별하는 일관된 방법입니다.

`purl`을 엔드포인트에 전달하면 Snyk은 해당 패키지에 대해 알려진 모든 취약점을 찾아 응답 본문의 일부로 반환합니다.

이 API는 패키지 버전의 취약점 목록을 검색하려는 경우 유용합니다.

{% hint style="info" %}
이 페이지의 예시는 [HTTPie](https://httpie.io/)를 사용하고 있지만, Snyk REST API에 액세스하기 위해 모든 HTTP 클라이언트를 사용할 수 있습니다.
{% endhint %}

## 지원되는 purl 유형

현재 릴리스에서는 다음 `purl` 유형을 지원합니다: `apk`, `cargo`, `cocoapods`, `composer`, `deb`, `gem`, `generic`, `golang`, `hex`, `npm`, `nuget`, `pub`, `pypi`, `rpm`, `swift` 및 `maven`.

추가적인 생태계 지원에 관심이 있다면, [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

## 패키지 엔드포인트를 위한 문제 목록 요청

API 엔드포인트를 호출하기 위해 **다음 HTTP 요청을 사용하십시오**:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/{purl}/issues" \
  "Authorization: token $API_TOKEN" \
  version==<snyk-api-version>
```

**purl은 URL로 인코딩**되어야 합니다.

다음은 유효한 URL로 인코딩된 purl을 사용한 예제입니다:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/pkg%3Amaven%2fcom.fasterxml.woodstox%2fwoodstox-core%405.0.0/issues" \
  "Authorization: token $API_TOKEN" \
  version==2024-06-26
```

운영 체제 패키지의 경우, 네임스페이스 부분에 벤더를 지정하고 `distro` 한정자를 지정해야 합니다. 지원되는 벤더는 다음과 같습니다: `debian`, `alpine`, `rhel`, `ubuntu`, `amzn`, `centos`, `oracle`, `rocky`, `sles`.

유효한 URL로 인코딩된 운영 체제 purl을 사용한 예시는 다음과 같습니다:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/pkg%3Adeb%2Fdebian%2Fcurl%3Fdistro%3Dbullseye/issues" \
  "Authorization: token $API_TOKEN" \
  version==2024-06-26
```

Snyk REST API는 페이지네이션을 지원합니다. 이는 기본 페이지 제한이 **1000**이며 기본 오프셋이 **0**입니다. 현재, 다음 및 이전 페이지는 응답에 링크로 반환됩니다. 다음 매개변수를 쿼리 매개변수로 제공할 수 있습니다: `offset`, `limit`.

페이지네이션된 요청의 예시는 다음과 같습니다:

```bash
$ http \
  "https://api.snyk.io/rest/orgs/{org_id}/packages/pkg%3Amaven%2fcom.fasterxml.woodstox%2fwoodstox-core%405.0.0/issues" \
  "Authorization: token $API_TOKEN" \
  version==2024-06-26 \
  limit==100 \
  offset==0
```

## 패키지 엔드포인트를 위한 문제 목록 응답

기대되는 출력은 패키지와 관련된 취약점을 식별하는 [JSON API](https://jsonapi.org/format/) 응답을 제공합니다.

다음 예제는 `pypi` 패키지 [django](https://security.snyk.io/package/pip/django)에 대한 응답을 제공합니다.

응답은 요청의 purl로 식별된 패키지에 대해 발견된 취약점 목록을 제공합니다. 응답은 취약점에 대한 설명으로 시작합니다:

### **패키지 개요**

이 패키지의 영향을 받는 버전은 `urlize()`와 `urlizetrunc()` 템플릿 필터의 특정 문자열 순서를 사용하여 아주 큰 입력으로 인한 서비스 거부 (DoS)에 취약합니다.

### **수정 방법**

이 취약점을 수정하려면 패키지 버전을 4.2.15,5.0.8로 업그레이드하십시오.

### 응답 세부사항

{% hint style="info" %}
응답은 연속적이며, 여기서는 설명을 위해 분할했습니다.
{% endhint %}

**각 취약점에 대해**, 응답은 다음을 제공합니다:

* Snyk 이슈 ID 및 이슈 유형:

    ```json
    "id": "SNYK-PYTHON-DJANGO-7642790",
    "type": "issue",
    ```
* 취약점에 대한 일반 메타데이터, 제목, 게시 및 공개 시간과 같은 취약점에 관련된 타임스탬프, 설명을 포함합니다:

    ```json
    "title": "서비스 거부 (DoS)",
    "type": "package_vulnerability",
    "created_at": "2024-08-07T08:13:29.424951Z",
    "updated_at": "2024-08-08T13:36:35.964359Z",
    "description": ...
    ```
* CVE 및 CWE 식별자:

    ```json
    "problems": [
        {
            "id": "CVE-2024-41990",
            "source": "CVE"
        },
        {
            "id": "CWE-400",
            "source": "CWE"
        }
    ],
    ```
* 취약점의 심각도:

    ```json
    "severities": [
        {
            "type": "primary",
            "source": "Snyk",
            "level": "medium",
            "score": 6.9,
            "version": "4.0",
            "vector": "CVSS:4.0/AV:N/AC:L/AT:N/PR:N/UI:N/VC:N/VI:N/VA:L/SC:N/SI:N/SA:N"
        },
        {
            "type": "secondary",
            "source": "Snyk",
            "level": "medium",
            "score": 5.3,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L"
        },
        {
            "type": "secondary",
            "source": "NVD",
            "level": "high",
            "score": 7.5,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H"
        },
        {
            "type": "secondary",
            "source": "Red Hat",
            "level": "high",
            "score": 7.5,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H"
        },
        {
            "type": "secondary",
            "source": "SUSE",
            "level": "high",
            "score": 7.5,
            "version": "3.1",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H"
        }
    ],
    ```

{% hint style="info" %}
2024년 6월부터 {{Snyk 오픈 소스}}가 식별하는 모든 새로운 알림은 CVSS v4.0 및 CVSS v3.1 심각성이 함께 제공됩니다. 가장 정확한 심각성 평가를 위해, Snyk은 사용 가능한 경우 CVSS v4.0을 사용할 것을 권장합니다.
{% endhint %}

* 해당 취약점에 대한 사용 가능한 수정 사항 및 취약 버전 표현:

    ```json
    "coordinates": [
        {
            "remedies": [
                {
                    "type": "indeterminate",
                    "description": "이 취약점을 수정하려면 패키지 버전을 4.2.15,5.0.8로 업그레이드하십시오",
                    "details": {
                        "upgrade_package": "4.2.15,5.0.8"
                    }
                }
            ],
            "representations": [
                {
                    "resource_path": "[,4.2.15),[5.0,5.0.8)"
                },
                {
                    "package": {
                        "name": "django",
                        "version": "4.2.14",
                        "type": "pypi",
                        "url": "pkg:pypi/django@4.2.14"
                    }
                }
            ]
        }
    ],
    ```
* 취약점에 대한 추가 정보를 가진 외부 리소스로의 링크:

```json
"references": [
 {
    "url": "https://www.djangoproject.com/weblog/2024/aug/06/security-releases/",
    "title": "Django 보안 릴리스"
},
    ... 
```

**패키지 메타데이터**는 다음을 포함합니다:

* 패키지 이름
* 패키지 유형
* 패키지 URL 명세
* 패키지 버전

```json
"meta": {
    "package": {
        "name": "django",
        "type": "pypi",
        "url": "pkg:pypi/django@4.2.14",
        "version": "4.2.14"
    }
}
```

적용 가능한 경우 결과에 대한 **페이지네이션 링크**가 다음과 같이 포함됩니다:

* 다음 경로 (적용 가능한 경우)
* 이전 경로 (적용 가능한 경우)
* 현재 경로

```json
"links": {
    "prev": "/orgs/<org-id>/packages/{purl}/issues?version=<api-version>&limit=1000&offset=0",
    "self": "/orgs/<org-id>/packages/{purl}/issues?version=<api-version>&limit=1000&offset=1"
},
```

## 패키지 엔드포인트를 위한 문제 목록 문제 해결

다음은 API를 사용할 때 받을 수 있는 **오류 상태**입니다. 여기서 다루지 않는 문제가 발생하거나 이 문제를 해결하는 데 문제가 있는 경우, 솔루션 엔지니어, 기술 성공 담당자에게 문의하거나 [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

**잘못된 PURL**\
400\
제공한 purl 사양이 유효한 purl인지 확인하십시오. 자세한 내용은 [패키지 URL 명세](https://github.com/package-url/purl-spec)를 참조하십시오.

**지원되지 않는 생태계**\
400\
패키지 유형이 [지원되는 purl 유형](list-issues-for-a-package.md#supported-purl-types) 중 하나인지 확인하십시오.

**네임스페이스 없이 요청된 패키지**\
400\
패키지 URL에서 네임스페이스를 지정하고 다시 시도하십시오. 자세한 내용은 [패키지 URL 명세](https://github.com/package-url/purl-spec)를 참조하십시오.

**지원되지 않는 PURL 구성요소**\
400\
지원되지 않는 구성요소를 제거하고 요청을 다시 보내십시오. 엔드포인트는 특정 구성요소만 허용합니다. 자세한 내용은 [패키지 URL 명세](https://github.com/package-url/purl-spec)를 참조하십시오.

**귀하의 조직은 이 작업을 수행할 권한이 없습니다.**\
403\
액세스 권한을 받으려면 솔루션 엔지니어, 기술 성공 담당자 또는 팀 관리자에게 문의하십시오.

**요청 속도 제한 초과**\
429\
이 API 엔드포인트에서는 사용자 당 분당 180개의 요청이 허용됩니다. 이 한도를 초과하면 429 오류 응답 코드를 받게 됩니다.

**잘못된 페이지네이션 매개변수**\
400\
제공된 limit 및 offset 쿼리 매개변수는 다음과 같아야 합니다:

* Limit > 0 및 <= 1000
* Offset >= 0

이 요청에 대한 기본 매개변수는 limit = 1000이며 offset >= 0입니다.

**권한 요청 실패**\
500\
이 문제는 예상치 못했으며 서비스가 빠르게 복구해야 합니다. 그렇지 않은 경우, [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

**내부 서버 오류**\
500\
이 문제는 예상치 못했으며 서비스가 빠르게 복구해야 합니다. 그렇지 않은 경우, [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

**취약점 서비스 사용 불가**\
503\
이 문제는 예상치 못했으며 서비스가 빠르게 복구해야 합니다. 그렇지 않은 경우, [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

**취약점 서비스 오류**\
500\
이 문제는 예상치 못했으며 서비스가 빠르게 복구해야 합니다. 그렇지 않은 경우, [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

다음은 이러한 오류 응답의 예시입니다:

```json
"jsonapi": {
    "version": "1.0"
},
"errors": [{
    "id": "8139dce7-eeb4-404b-be0e-5e4c15004972",
    "detail": "지원되지 않는 생태계",
    "status": "400"
}]
```