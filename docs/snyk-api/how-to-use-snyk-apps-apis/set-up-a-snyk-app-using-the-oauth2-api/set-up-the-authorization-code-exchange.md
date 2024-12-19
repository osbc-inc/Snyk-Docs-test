# 권한 코드 교환 설정

권한 **코드**를 받은 후에는 액세스 토큰으로 교환해야 합니다.

액세스 토큰을 요청하려면 토큰 엔드포인트에 POST 요청을 보냅니다:

```
https://api.snyk.io/oauth2/token
```

다음 속성을 포함한 x-www-form-urlencoded 서식의 요청 본문으로 POST 요청을 보냅니다:

```
grant_type=authorization_code
&code=(이전 단계의 코드)
&redirect_uri=(이전 단계의 리다이렉트 URI)
&client_id=(앱 생성 시 clientId)
&client_secret=(앱 생성 시 clientSecret)
&code_verifier=(이전 단계에서 생성된 코드 확인자)
```

응답에는 Snyk API와 통신하는 데 필요한 세부 정보인 `access_token`과 `refresh_token` 및 각각의 만료 시간이 포함됩니다.

두 토큰 모두 안전한 데이터 저장소에 저장되어야 합니다. 저장하기 전에 값들을 암호화하는 것이 매우 권장됩니다.

`access_token`은 사용자와 조직을 대표하여 미래의 API 호출에 사용됩니다. `access_token`의 수명은 `refresh_token`보다 훨씬 짧습니다.

`refresh_token`은 만료되면 새 `access_token`을 얻기 위해 한 번만 사용할 수 있습니다. 즉, `refresh_token`의 만료 시간이 지나거나 `access_token`를 새로 고칠 때 사용되면 더 이상 사용할 수 없게됩니다.