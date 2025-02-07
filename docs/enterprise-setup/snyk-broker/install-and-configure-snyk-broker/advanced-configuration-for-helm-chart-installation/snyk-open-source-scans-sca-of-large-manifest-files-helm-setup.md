# 대용량 매니페스트 파일의 Snyk 오픈 소스 스캔(SCA), Helm 설정

큰 매니페스트 파일이 Snyk에 의해 감지되면 때로는 파일을 검색하기 위해 다른 방법 (다른 SCM 엔드포인트)을 사용해야 할 수 있습니다. 대체 방법을 사용하려면 더 융통성 있는 규칙이 필요하므로 기본적으로 비활성화됩니다.

{% hint style="info" %}
이는 Github 및 Github Enterprise Broker 통합에만 해당됩니다.
{% endhint %}

브로커 클라이언트가 다른 엔드포인트를 사용하여 더 큰 매니페스트 파일을 검색할 수 있도록 규칙을 쉽게 추가하려면 다음 환경 변수를 추가하세요.

```console
--set 'env[0].name=ACCEPT_LARGE_MANIFESTS' \
--set 'env[0].value=true'
```

{% hint style="info" %}
다른 사용자 정의 환경 변수를 사용 중이면이 코드에서 0 인덱스를 조정하여 환경 변수가 서로 덮어쓰이지 않도록해야 할 수 있습니다.
{% endhint %}

ACCEPT 환경 변수 대신 [사용자 정의 accept.json](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/adding-custom-accept.json-for-helm-installation)을 사용하는 경우 accept.json의 private 섹션에 다음 내용을 추가하세요.

```json
{
    "//": "used to get given manifest file",
    "method": "GET",
    "path": "/repos/:owner/:repo/git/blobs/:sha",
    "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
}
```

{% hint style="info" %}
가능한 최대 보안을 보장하기 위해 Snyk는 해당 규칙을 기본적으로 활성화하지 않습니다. 이 엔드포인트의 사용은 경로에 구체적으로 허용된 파일 이름이 포함되어 있지 않으므로 이 엔드포인트를 통해 이 리포지토리의 모든 파일에 대해 이론적으로 Snyk 플랫폼이 접근할 수 있음을 의미합니다.
{% endhint %}
