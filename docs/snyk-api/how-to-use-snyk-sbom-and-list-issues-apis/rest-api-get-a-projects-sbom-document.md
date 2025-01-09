# 프로젝트의 SBOM 문서 가져오기

{% hint style="info" %}
**사전 준비 사항**\
이 기능은 Snyk 엔터프라이즈 플랜을 가입한 고객에게 제공됩니다.
{% endhint %}

Snyk은 Open Source 및 Container 프로젝트의 SBOM 문서를 생성하기 위한 [프로젝트의 SBOM 문서 가져오기](../reference/sbom.md) 엔드포인트를 제공합니다. 이 프로젝트들은 지속적으로 문제를 모니터링하고 있습니다.

SBOM 문서는 프로젝트의 종속성과 이들 간의 관계의 최신 상태를 나타냅니다.

SBOM 문서는 [CycloneDX](https://cyclonedx.org/) v1.4, v1.5, v1.6 (JSON, XML) 및 [SPDX](https://spdx.dev/) v2.3 (JSON) 형식으로 생성할 수 있습니다.

## 프로젝트용 SBOM 생성 방법

1. Snyk 웹 UI에서 `organization ID` (UUID 형식), `project ID` (UUID) 및 API 키를 검색합니다.\
   이러한 값들을 찾는 데 도움이 필요한 경우 [그룹 및 조직 간 전환](../../snyk-admin/groups-and-organizations/switch-between-groups-and-organizations.md), [프로젝트 설정 보기 및 편집](../../snyk-admin/snyk-projects/view-and-edit-project-settings.md), [API용 인증](../rest-api/authentication-for-api/) 문서를 참조하세요.
2. 생성할 SBOM의 형식을 결정합니다.\
   사용 가능한 옵션: CycloneDX 1.4 JSON (`cyclonedx1.4+json`), CycloneDX 1.4 XML (`cyclonedx1.4+xml`), CycloneDX 1.5 JSON (`cyclonedx1.5+json`), CycloneDX 1.5 XML (`cyclonedx1.5+xml`), CycloneDX 1.6 JSON (`cyclonedx1.6+json`), CycloneDX 1.6 XML (`cyclonedx1.6+xml`) 또는 SPDX v2.3 JSON (`spdx2.3+json`).
3. HTTP client인 Postman이나 `curl`과 같은 도구를 사용하여 엔드포인트에 요청을 보냅니다.\
   `format` 매개변수가 URL 인코딩되어 있어야 합니다.\
   예: CycloneDX 1.4 JSON 문서를 검색하려면 쿼리에 `format=cyclonedx1.4%2Bjson`을 설정합니다.

```bash
$ curl --location 'https://api.snyk.io/rest/orgs/<ORG_ID>/projects/<PROJECT_ID>/sbom?version=2024-03-12&format=<SBOM_FORMAT>' \
--header 'Authorization: token <SNYK_API_TOKEN>'
```

## 사용자 지정 CycloneDX 속성

Snyk에 의해 생성된 SBOM 문서에는 내보낸 내용에 대한 Snyk 특정 메타데이터가 포함됩니다. 이는 CycloneDX로 내보낼 때 문서의 `metadata.properties` 섹션에 포함됩니다.

<table><thead><tr><th width="240">속성 이름</th><th>설명</th></tr></thead><tbody><tr><td><code>snyk:org_id</code></td><td>소속된 조직 ID (UUID), 해당하는 경우</td></tr><tr><td><code>snyk:collection_id</code></td><td>프로젝트 컬렉션 ID (UUID), 해당하는 경우</td></tr><tr><td><code>snyk:project_id</code></td><td>프로젝트 ID (UUID), 해당하는 경우</td></tr><tr><td><code>snyk:target_id</code></td><td>대상 ID (UUID), 해당하는 경우</td></tr></tbody></table>

## 프로젝트의 SBOM 문서 가져오기 엔드포인트에 대한 문제 해결

다음 응답 코드는 성공을 나타냅니다.

**200 OK**

SBOM 문서가 성공적으로 생성되었습니다. 응답 본문에 요청된 형식의 문서가 포함되어 있습니다.

다음은 API를 사용할 때 받을 수 있는 **오류 상태**입니다. 여기에서 다루지 않은 문제가 발생하거나 이를 해결하기 어려울 경우 솔루션 엔지니어 또는 기술 성공 매니저에게 문의하거나 [Snyk 지원팀](https://support.snyk.io)에 요청을 제출하세요.

**401 Unauthorized**

인증 방법, Bearer 토큰을 위한 API 토큰이 잘못되었습니다. Authorization 헤더를 올바르게 설정했는지 확인하세요.

**403 Forbidden**

요청을 수행하기에 필요한 권한이 없습니다. 요청된 조직의 일부가 아니거나, 조직이 Snyk API를 사용할 자격이 없거나, 요청된 프로젝트에 대한 충분한 읽기 액세스 권한이 없는 경우 발생할 수 있습니다.

**429 Too Many Requests**

Snyk API는 요청 제한이 있으므로 지나치게 많은 요청은 거부될 수 있습니다. 추가 요청을 하기 전에 잠시 기다리세요.

**500 Internal Server Error**

서비스가 내부 시스템 오류를 만났습니다.
