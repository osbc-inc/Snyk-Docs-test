# OAuth2 API

Snyk은 [Snyk Apps](how-to-use-snyk-apps-apis/)와 주로 사용하기 위한 OAuth2 API를 제공합니다. 이는 RFC 6749를 준수합니다.

대부분의 엔드포인트는 Snyk API 서브도메인에서 제공됩니다(예: https://api.snyk.io), 하나의 예외는 `/oauth2/authorize`이며 이는 주요 앱 서브도메인에서 제공됩니다(예: https://app.snyk.io).

{% swagger src="../.gitbook/assets/oauth-app-spec.yaml" path="/oauth2/authorize" method="get" %}
[oauth-app-spec.yaml](../.gitbook/assets/oauth-app-spec.yaml)
{% endswagger %}

{% swagger src="../.gitbook/assets/oauth-api-spec.yaml" path="/token" method="post" %}
[oauth-api-spec.yaml](../.gitbook/assets/oauth-api-spec.yaml)
{% endswagger %}

{% swagger src="../.gitbook/assets/oauth-api-spec.yaml" path="/revoke" method="post" %}
[oauth-api-spec.yaml](../.gitbook/assets/oauth-api-spec.yaml)
{% endswagger %}