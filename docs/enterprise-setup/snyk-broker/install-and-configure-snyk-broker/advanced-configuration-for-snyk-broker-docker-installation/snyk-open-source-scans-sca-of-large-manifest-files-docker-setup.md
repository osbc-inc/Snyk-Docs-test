# Snyk Open Source Scans (SCA) of large manifest files, Docker setup

대규모 매니페스트 파일이 감지되면, 때로는 파일을 검색하기 위해 다른 방법(다른 SCM 엔드포인트)을 사용해야 할 수 있습니다. 대체 방법은 더 관대한 규칙이 필요하므로 기본적으로 비활성화됩니다. &#x20;

{% hint style="info" %}
이는 Github 및 Github Enterprise Broker 통합에만 해당됩니다.&#x20;
{% endhint %}

브로커 클라이언트가 다른 엔드포인트를 사용하여 더 큰 매니페스트 파일을 검색할 수 있도록 하는 이 규칙을 쉽게 추가하려면 다음 환경 변수를 추가하십시오.

```console
-e ACCEPT_LARGE_MANIFESTS=true
```

ACCEPT 환경 변수 대신 [사용자 지정 accept.json](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/custom-approved-listing-filter-with-docker)을 사용하는 경우, 다음을 accept.json의 private 섹션에 추가하십시오.

```
{
    "//": "given manifest file을 가져오는 데 사용됩니다",
    "method": "GET",
    "path": "/repos/:owner/:repo/git/blobs/:sha",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

{% hint style="info" %}
최대한의 보안을 보장하기 위해, Snyk는 이 규칙을 기본적으로 활성화하지 않습니다. 이 엔드포인트의 사용은 경로에 특정 허용된 파일 이름이 포함되어 있지 않기 때문에 이 리포지토리의 모든 파일에 대한 이론상의 Snyk 플랫폼 액세스가 가능합니다.
{% endhint %}