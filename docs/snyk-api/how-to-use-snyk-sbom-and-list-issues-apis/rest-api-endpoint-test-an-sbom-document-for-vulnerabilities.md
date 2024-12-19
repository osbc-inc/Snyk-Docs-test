# 취약점을 테스트하는 SBOM 문서

{% hint style="info" %}
**릴리스 상태 및 기능 가용성**

Snyk REST API는 엔터프라이즈 플랜에서만 사용 가능합니다. 자세한 정보는 [요금제 정보](https://snyk.io/plans)를 참조하십시오.

이러한 엔드포인트는 베타 API 버전입니다. 기능 중 일부가 변경될 수 있습니다. 자세한 정보는 REST API의 [버전 관리](../rest-api/about-the-rest-api.md#versioning) 정보를 참조하십시오.
{% endhint %}

Snyk은 [API 엔드포인트 모음](https://apidocs.snyk.io/?version=2024-09-03%7Ebeta#post-/orgs/-org_id-/sbom_tests)을 제공하여 소프트웨어 원료료록(SBOM) 문서를 비동기식으로 테스트할 수 있습니다. 이러한 엔드포인트를 사용하여 SBOM 및 해당 패키지에 영향을 주는 취약성에 대해 자세히 알아볼 수 있습니다.

{% hint style="info" %}
지원되는 SBOM 형식은 [CycloneDX](https://cyclonedx.org/) 1.4/1.5/1.6 JSON 및 [SPDX](https://spdx.dev/) 2.3 JSON입니다.
{% endhint %}

Snyk은 SBOM 내의 구성 요소를 [패키지 URL](https://github.com/package-url/purl-spec)(purl)에 따라 식별합니다. 구성 요소에 purl이 포함되어 있지 않거나 purl 유형이 지원되지 않는 경우 Snyk는 해당 구성 요소에 대해 취약성 분석을 건너 뜁니다. 지원되는 purl 유형은 다음과 같습니다: `apk`, `cargo`, `cocoapods`, `composer`, `deb`, `gem`, `golang`, `hex`, `maven`, `npm`, `nuget`, `pypi`, `rpm`, `swift`, `generic` (미관리 C/C++ 종속성용).

## SBOM 문서를 테스트하는 방법

[SBOM 엔드포인트](https://apidocs.snyk.io/?version=2024-09-03%7Ebeta#post-/orgs/-org_id-/sbom_tests)를 사용하여 SBOM 문서를 생성하고 상태를 확인하고 결과를 확인할 수 있습니다. 다음 단계를 따르세요:

1. Snyk에 SBOM을 보내어 테스트를 생성합니다. (rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md#Create-a-test-by-sending-an-sbom-to-snyk)
2. [테스트의 상태를 확인합니다](rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md#Check-the-status-of-the-test).
3. [테스트 결과를 확인합니다. 테스트가 완료된 경우](rest-api-endpoint-test-an-sbom-document-for-vulnerabilities.md#View-results-of-the-test).

### SBOM 문서를 Snyk에 전송하여 테스트 생성하기&#x20;

SBOM을 테스트하는 것은 시간이 오래 걸릴 수 있습니다. 테스트 결과가 준비될 때까지 기다리는 대신, 초기 요청 후에 Snyk는 SBOM을 보낸 후 `job_id`를 반환하고 요청을 비동기적으로 처리합니다.

다음 단계를 따라 SBOM을 테스트합니다:

1. Snyk 웹 UI에 로그인하고 조직 ID(UUID 형식), 프로젝트 ID(UUID) 및 API 키를 검색합니다.\
   이러한 값들을 찾는 데 도움이 필요한 경우 [조직 일반 설정](../../snyk-admin/groups-and-organizations/organizations/organization-general-settings.md), [프로젝트 설정 보기 및 편집](../../snyk-admin/snyk-projects/view-and-edit-project-settings.md) 및 [API 인증](../rest-api/authentication-for-api/authenticate-for-the-api.md)을 참조하십시오.
2. `curl` 또는 Postman과 같은 HTTP 클라이언트를 사용하여 [SBOM 테스트 실행 생성](https://apidocs.snyk.io/?version=2024-09-03%7Ebeta#post-/orgs/-org_id-/sbom_tests) 엔드포인트로 요청을 보냅니다.&#x20;

{% hint style="info" %}
요청 본문에 SBOM 문서가 JSON 객체의 형태로 포함됩니다. 이 요청은 SBOM 문서에 대한 테스트 실행을 생성합니다.
{% endhint %}

{% code title="HTTP 요청" %}
```bash
curl --request POST \
    -H "Authorization: token <SNYK_TOKEN>" \
    -H "Content-Type: application/vnd.api+json" \
    --data-binary '@request_body.json' \
    'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests?version=2023-089-03~beta'
```
{% endcode %}

{% code title="request_body.json" %}
```json
{
    "data": {
        "type": "sbom_test",
        "attributes":{ 
            "sbom": {
            <SBOM_CONTENTS>
            }
        }
    }
}
```
{% endcode %}

3. 응답에서 다음 단계에서 사용할 `job_id`를 얻습니다. \
   이는 SBOM 문서에 대한 수행 중인 테스트 실행의 고유 식별자입니다.

{% code title="JSON 응답 본문" %}
```json
{
    "data": {
        "id": "<JOB_ID>",
        "type": "sbom_tests"
    },
    "jsonapi": {
        "version": "1.0"
    },
    "links": {
        "self": "/rest/orgs/<ORG_ID>/sbom_tests?version=2023-08-31~beta",
        "related": "/rest/orgs/<ORG_ID>/sbom_tests/<JOB_ID>?version=2023-08-31~beta"
    }
}
```
{% endcode %}

### 테스트의 상태 확인하기 (옵션)

초기 요청 후 언제든지 테스트의 상태를 확인할 수 있습니다. &#x20;

1. [SBOM 테스트 실행 생성](https://apidocs.snyk.io/?version=2024-09-03%7Ebeta#post-/orgs/-org_id-/sbom_tests) 엔드포인트로부터 반환된 `job_id`를 사용하여 엔드포인트 [SBOM 테스트 실행 상태 가져오기](https://apidocs.snyk.io/?version=2024-09-03%7Ebeta#get-/orgs/-org_id-/sbom_tests/-job_id-)로 요청을 보냅니다.
2. 이 엔드포인트에 대한 성공적인 요청은 테스트의 상태인 `processing` 또는 `finished`를 반환합니다. 호출에 실패하면 오류가 반환됩니다.

```bash
  curl --get \
      -H "Authorization: token <SNYK_TOKEN>" \
      'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests/<TEST_ID>?version=2024-09-03~beta'
```

### 테스트 결과 보기

테스트가 완료되면 검사한 SBOM의 결과를 볼 수 있습니다.

1. 테스트의 상태가 `finished`로 반환된 경우 [SBOM 테스트 실행 결과 가져오기](https://apidocs.snyk.io/?version=2024-09-03%7Ebeta#get-/orgs/-org_id-/sbom_tests/-job_id-/results) 엔드포인트로 요청을 보냅니다.
2. 테스트된 SBOM에 대한 요약 수준 정보 및 자세한 결과를 포함하는 반환된 정보를 확인합니다.

```bash
curl --get \
    -H "Authorization: token <SNYK_TOKEN>" \
    'https://api.snyk.io/rest/orgs/<ORG_ID>/sbom_tests/<TEST_ID>/results?version=2023-08-31~beta'
```

## SBOM 테스트 실행 생성 엔드포인트의 문제 해결

다음 응답 코드는 성공을 나타냅니다.

**201 Created**

SBOM 테스트 실행이 성공적으로 생성되었습니다. 응답 본문에 테스트 실행의 작업 ID가 포함됩니다.

다음은 API를 사용할 때 받을 수 있는 오류 상태입니다. 여기에 다루지 않은 문제가 있거나 이러한 문제를 해결하는 데 문제가 있는 경우 솔루션 엔지니어나 기술 성공 관리자에 문의하거나 [Snyk 지원](https://support.snyk.io)으로 요청을 제출하십시오.

**400 Bad Request**

서버가 유효하지 않거나 손상된 데이터로 인해 요청을 처리할 수 없습니다. 요청을 검토하고 다시 시도하십시오.

**401 Unauthorized**

인증 방법, API 토큰 또는 Bearer 토큰이 잘못되었습니다. Authorization 헤더를 올바르게 설정했는지 확인해 주십시오.

**403 Forbidden**

요청을 만들 권한이 없습니다. 요구된 조직의 일부가 아니거나 귀하의 조직이 Snyk API를 사용할 자격이 없거나 요청된 프로젝트에 대한 충분한 읽기 액세스 권한이 없는 경우 발생할 수 있습니다.

**429 Too Many Requests**

Snyk API가 요청을 제한하므로 과도한 요청의 경우 일정 시간 후 요청이 거부됩니다. 추가 요청을 하기 전에 기다려야 합니다.

**500 Internal Server Error**

서비스가 내부 시스템 오류를 겪었습니다.