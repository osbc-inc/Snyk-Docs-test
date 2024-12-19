# Docker를 사용하여 인증 방법 변경

각 통합은 서비스에 따라 다양한 방법으로 기본 인증 방법이 설정되어 있습니다.

예를 들어, BitBucket Server 및 Data Center는 `accept.json` 파일에서 사용자 이름과 비밀번호를 사용하는 Basic Auth를 사용합니다.

```json
{
  "private": [
    {
      ...,
      "auth": {
         "scheme": "basic",
         "username": "${BITBUCKET_USERNAME}",
         "password": "${BITBUCKET_PASSWORD}"
      }
    },
    ...
  ]
}
```

Artifactory의 경우, 인증 방법은 기본적으로 `.env` 파일에 구성됩니다.

```shell
# 당신의 Artifactory URL
# 기본 인증을 사용하지 않는 경우 "<yourdomain.artifactory.com>/artifactory"로만 구성해야 합니다.
ARTIFACTORY_URL=<username>:<password>@<yourdomain.artifactory.com>/artifactory
```

GitHub의 경우, 인증 방법은 `accept.json` 파일의 `origin` 필드의 일부입니다.

```json
{
  "private": [
    {
      ...,
      "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
    },
    ...
  ]
}
```

인증 방법을 재정의할 수 있습니다. `scheme`에 대한 유효한 값은 각각 `bearer`, `token`, 및 `basic`이며, 이는 각각 헤더를 `Bearer`, `Token`, 및 `Basic`로 설정합니다. Bearer 토큰을 선호하는 경우, `accept.json` 파일을 다음과 같이 구성할 수 있습니다.

```json
{
  "private": [
    {
      ...,
      "auth": {
        "scheme": "bearer",
        "token": "${BEARER_TOKEN}"
      }
    },
    ...
  ]
}
```

`private` 배열의 각 개별 객체에 대해 이것을 설정해야 합니다.

`scheme`가 `bearer` 또는 `token`인 경우, 반드시 `token`을 제공해야 하며, `scheme`가 `basic`인 경우 `username` 및 `password`를 제공해야 합니다.

이는 `origin` 필드나 `.env` 파일에 토큰을 설정하는 등 다른 구성된 인증 방법을 재정의합니다.

{% hint style="info" %}
Snyk Broker는 mTLS 방법을 사용한 인증을 지원하지 않습니다. &#x20;
{% endhint %}