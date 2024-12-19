# 기존 통합 업데이트 엔드포인트에 대한 예시

[Integrations (v1)](../reference/integrations-v1.md) API의 [기존 통합 업데이트](../reference/integrations-v1.md#org-orgid-integrations-integrationid) 엔드포인트에 대한 예시가 있습니다.

## 기존 통합을 위한 브로커 설정

### 명령어

```bash
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"github\",
    \"broker\": { \"enabled\": true }
}" \
'https://api.snyk.io/v1/org/{orgId}/integrations/{integrationId}'
```

### 응답

```json
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b",
  "brokerToken": "4a18d42f-0706-4ad0-b127-24078731fbed"
}
```

### Type에 대한 가능한 값 (통합 유형, enum)

`acr`, `artifactory-cr`, `azure-repos`, `bitbucket-cloud`, `bitbucket-server`, `digitalocean-cr`, `docker-hub`, `ecr,` `gcr`, `github`, `github-cr`, `github-enterprise`, `gitlab`, `gitlab-cr`, `google-artifact-cr`, `harbor-cr`, `nexus-cr`, `quay-cr`

### 업데이트하는 통합에 필요한 자격 증명

AcrCredentials: object\
username: 필수, string\
password: 필수, string\
registryBase: 필수, string, 예시: name.azurecr.io

...

### Snyk에서 필요한 권한

조직보기\
통합 보기\
통합 편집

## 기존 브로커되지 않은 통합의 자격 증명 업데이트

### 명령어

```bash
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"gitlab\",
    \"credentials\": { \"token\": \"GITLAB_TOKEN\" }
}" \
'https://api.snyk.io/v1/org/{orgId}/integrations/{integrationId}'
```

### 응답

```json
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b"
}
```

### Type에 대한 가능한 값 (통합 유형, enum)

[기존 통합에 대한 브로커 설정의 가능한 값](examples-for-the-update-existing-integration-endpoint.md#possible-values-for-type-integration-type-enum)과 동일합니다.

### 업데이트하는 통합에 필요한 자격 증명

[기존 통합에 대한 브로커 설정에 필요한 자격 증명과 동일](examples-for-the-update-existing-integration-endpoint.md#credentials-needed-for-the-integration-you-are-updating).

### Snyk에서 필요한 권한

조직보기\
통합 보기\
통합 편집

## 기존 통합에 대한 브로커 비활성화

### 명령어

```bash
curl --include \
     --request PUT \
     --header "Content-Type: application/json; charset=utf-8" \
     --header "Authorization: token API_KEY" \
     --data-binary "{
    \"type\": \"github\",
    \"broker\": { \"enabled\": false },
    \"credentials\": { \"token\": \"GITHUB_TOKEN\" }
}" \
'https://api.snyk.io/v1/org/{orgId}/integrations/{integrationId}'
```

### 응답

```json
{
  "id": "9a3e5d90-b782-468a-a042-9a2073736f0b"
}
```