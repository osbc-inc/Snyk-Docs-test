# 리프레시 토큰 교환 설정

`access_token`이 짧은 시간 내에 만료될 것이므로, 앱은 여전히 유효한 `refresh_token`을 사용하여 자주 새로운 토큰을 요청해야 합니다.

새로운 `access_token`을 가져오기 위해 토큰 엔드포인트로 POST 요청을 수행하십시오:

```
https://api.snyk.io/oauth2/token
```

다음 속성을 가진 x-www-form-urlencoded 형식의 요청 본문으로 요청을 만드십시오:

```
grant_type=refresh_token
&refresh_token=(이전 단계의 refresh 토큰)
&client_id=(앱 생성 시 clientId)
&client_secret=(앱 생성 시 clientSecret)
```

이 호출에 대한 응답은 각각 새 `access_token`, `refresh_token`, 그리고 만료 시간을 제공합니다.

이제 이전 `refresh_token` 대신 새 `refresh_token`을 저장해야 합니다. 이 엔드포인트로 동시 호출을 하면 refresh 토큰이 예기치 않게 무효화될 수 있으므로, 한 번에 주어진 refresh 토큰당 새로운 리프레시 토큰 쌍을 한 번만 요청하도록 해야 합니다. 이 프로세스가 실패하면 사용자는 권한 플로우를 다시 수행해야 합니다.