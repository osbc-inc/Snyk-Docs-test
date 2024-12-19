# Snyk 웹훅 연결을 테스트합니다

Snyk 웹훅은 새 취약점이 발견될 때만 업데이트되지만, Postman을 사용하여 구현을 테스트할 수 있습니다.

만약 Postman이 없다면 무료로 [설치](https://www.postman.com/downloads/)할 수 있습니다.

연결을 테스트하려면 AWS API Gateway에 샘플 페이로드를 포함한 POST 요청을 보내는 방법이 있습니다. 이때 API 토큰으로 보안되어야 합니다.

연결을 테스트하기 위해 다음 단계를 따르세요:

1. Postman을 설치합니다.
2. **컬렉션**을 만듭니다.

    <figure><img src="https://lh3.googleusercontent.com/j7ab9JGG5IAmqb6xuA7AjwPcF6cmUIhIzrn6p1f7CUQTkwQqHm7P5fVHxDx8I6tysjM93uqu5whBFq_qI1Q5h5y_KK0uR3Hv--uYhcDXJehU5ZCc68Fvdv79S8z7yqCp0CNbLYXOaxwc9cTR0ueQ9lYuydCDhyJmpA5TGNJ08wexCGIpeDMX0fO4Tw" alt=""><figcaption><p>새 Postman 컬렉션 만들기</p></figcaption></figure>
3. API 토큰(비밀 문자열)을 컬렉션의 환경 변수로 추가합니다.\
    변수 이름을 `x-hub-signature`로 지정하여 페이로드의 무결성을 확인하는 데 변수를 사용할 수 있습니다.\
    API 토큰을 찾는 방법은 [Snyk 웹훅 설정](set-up-the-snyk-webhook.md)을 참조하세요.

    <figure><img src="https://lh5.googleusercontent.com/QiPKevkpzyOwSscKxGu9BbzhbfU53bCQKF7y5CaXaImlQFA2VQKuwW5I2TSeKCis1fTDYkyJHaBa8koNDZ1izAHTE1fPWUo2S9bLETght4jPaaKujS8TZKyjOLpk4lUMyeBdSvg5wJvQ553VgK-p_eBJdDyM1St6pXadh9FaVdElZRFh14WBEMLGZA" alt=""><figcaption><p>새 Postman 컬렉션 만들기</p></figcaption></figure>
4. 컬렉션 내에 새 **HTTP 요청**을 생성합니다.

    <figure><img src="https://lh3.googleusercontent.com/j7ab9JGG5IAmqb6xuA7AjwPcF6cmUIhIzrn6p1f7CUQTkwQqHm7P5fVHxDx8I6tysjM93uqu5whBFq_qI1Q5h5y_KK0uR3Hv--uYhcDXJehU5ZCc68Fvdv79S8z7yqCp0CNbLYXOaxwc9cTR0ueQ9lYuydCDhyJmpA5TGNJ08wexCGIpeDMX0fO4Tw" alt=""><figcaption><p>Postman API 요청 생성</p></figcaption></figure>
5. 방법을 **POST**로 변경하고 API Gateway URL 또는 기능 URL을 추가합니다. URL을 찾는 방법은 [AWS API Gateway: Slack에 Snyk를 연결하는 POST 방법 추가](aws-lambda-setup-set-up-the-trigger/with-api-gateway/aws-api-gateway-add-the-post-method-to-connect-snyk-to-slack.md)을 참조하세요.

    <figure><img src="https://lh4.googleusercontent.com/5QxR-05QtK6FNpoyuPW06L_vyVAl6cCxMnph7euIKafc-YyGIgjaiA74KSNO93uTMGFGxNQnzwyfiZ5Oi3e1y0GA0P2INodvIbamhe6lpwwf1Kc7bCajYUPG0RcfedUOKMqI0l4mmuq1jECRHUiUtnsel7PiBxiIvddcCnplxwVDY9r0FDcYNKZPag" alt=""><figcaption><p>Postman POST 방법에 AWS API Gateway URL 추가</p></figcaption></figure>
6. 사전 요청 스크립트를 다음 코드처럼 설정합니다. 이 스크립트는 API 토큰을 가져와 Lambda 함수가 페이로드를 해독할 수 있도록 보안합니다.\
    아래의 코드를 따라 작성하되, '`your-secret-string-here'` 부분에 당신의 API 토큰을 입력하세요.\
    \
    `postman.setEnvironmentVariable('x-hub-signature', CryptoJS.HmacSHA256(request.data, 'your-secret-string-here').toString(CryptoJS.digest)); postman.setEnvironmentVariable('x-hub-signature', 'sha256='+ postman.getEnvironmentVariable('x-hub-signature'));`\

    <figure><img src="https://lh4.googleusercontent.com/imlrHdNQOJQVExPXvHiwNSR0zerKrR4qUJKeeXmJsfW-UTarEZtB9S3uW5K0xY4EarI5zft8PqUKEE5AS3TPWIWE5hTMNrLA5iCmv8f9Nv5onoTzPRsS8lXUTOQt4Fl-SFyFMvyTfLs3FBhcu_PCwjfB0zLvFXqGPFjYPw3b6ctorVVZ3YsVMQeVpg" alt=""><figcaption><p>Postman 사전 요청 스크립트</p></figcaption></figure>
7. **헤더**를 추가하세요:\
    **Content-Type**: **application/json**\
    **x-hub-signature** `{{x-hub-signature}}`: (단계 3에서 만든 API 토큰(비밀 문자열) 환경 변수입니다.

    <figure><img src="https://lh5.googleusercontent.com/SLs1bStNsB5yEBMRpie_PseTXwZuj5qYp_w5CIboLgNcrAJks87wVzoJuwI0TVa71kbXSS-k0zHbrEVSXaKp3j33S3Jn3Fy5dH21Yla8iNqFFSFqHQDf6ArhjbxUheFAaZbPFYoLuyhxoHlsKDNJkdoSk2L7v0vDGUrN4_-Bcf7S91PgvvT6wtZt9w" alt=""><figcaption><p>Postman POST 요청 헤더</p></figcaption></figure>
8. **본문**에 유효한 페이로드를 추가합니다. 아래 예제와 같이 작성하되 보기 좋게 포매팅하지 않으세요. 읽기 쉬울수록 동작하지 않습니다:\
    `{"project":{"id":"bc75a806-0893-4ccf-84c5-8fde48a88df8","name":"snyk/juice-shop:frontend/package.json","created":"2022-06-17T06:54:21.326Z","origin":"github","type":"npm","readOnly":false,"testFrequency":"daily","totalDependencies":1216,"issueCountsBySeverity":{"low":2,"high":16,"medium":17,"critical":0},"imageTag":"12.3.0","imagePlatform":"","lastTestedDate":"2022-06-29T05:45:12.569Z","browseUrl":"https://app.snyk.io/org/api-test/project/bc75a806-0893-4ccf-84c5-8fde48a88df7","importingUser":null,"owner":null,"tags":[],"isMonitored":true,"attributes":{"criticality":[],"lifecycle":[],"environment":[]},"branch":"master"},"org":{"id":"570a1e02-8774-4697-80fc-129f5c5195a1","name":"API","slug":"api-quc","url":"https://app.snyk.io/org/api-test","group":null,"created":"2022-05-25T06:29:29.833Z"},"newIssues":[{"id":"SNYK-JS-SCSSTOKENIZER-2339884","issueType":"vuln","pkgName":"scss-tokenizer","pkgVersions":["0.2.3"],"priorityScore":336,"priority":{"score":336,"factors":[{"name":"isFresh","description":"Recently disclosed"},{"name":"cvssScore","description":"CVSS 5.3"}]},"issueData":{"id":"SNYK-JS-SCSSTOKENIZER-2339884","title":"Regular Expression Denial of Service (ReDoS)","severity":"high","url":"https://snyk.io/vuln/SNYK-JS-SCSSTOKENIZER-2339884","description":"Long description","identifiers":{"CWE":["CWE-1333"],"CVE":["CVE-2022-25758"]},"credit":["Paul Bastide"],"exploitMaturity":"no-known-exploit","semver":{"vulnerable":["*"]},"publicationTime":"2022-06-29T10:29:38Z","disclosureTime":"2022-01-13T16:29:34Z","CVSSv3":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L","cvssScore":5.3,"functions":[],"language":"js","patches":[],"nearestFixedInVersion":"","isMaliciousPackage":false},"isPatched":false,"isIgnored":false,"fixInfo":{"isUpgradable":false,"isPinnable":false,"isPatchable":false,"isFixable":false,"isPartiallyFixable":false,"nearestFixedInVersion":"","fixedIn":[]}}],"removedIssues":[]}`

    <figure><img src="https://lh6.googleusercontent.com/vi_Mt44ag0EzWi9bn9ruwnzBcF-cYxGqajF-F6jQF2nwJEEvNa6KW45ZgszlekP17zLQwRH-z9iar-oTvkOKXdAWEb-ewCJVujrj-pzkHlKftd4Y1GmPyaguELBtbKP-m3RLAN9-R6PxzO1psWDY_KoW7iHwLc3oQax7gcQArwMtf2oxSlmvHUxzWA" alt=""><figcaption><p>유효한 페이로드가 포함된 Postman POST 요청 본문</p></figcaption></figure>
9. URL 옆의 **Send** 버튼을 클릭합니다.
10. Postman 하단에 성공 메시지(`statusCode 200`)가 표시되는지 확인합니다.

    <figure><img src="https://lh4.googleusercontent.com/YHelnzIIPgeL7ZkbVOy67hMiaVe6_lz3VvFjhNg8vkeRm4EtevSypMR_PsSRCfzkZcob76KSSgdvrPoqhVwEBL8FwT2LXiIn9u9hv5-bVrF_zh7sK3lB0rJM3lBmqc5w6miUx7hD7ROlLrXROIbAgUWWqCnYpvZ6C8TJcKI_kSTYG5LMaYg2lRm3RA" alt=""><figcaption><p>Postman 성공 메시지</p></figcaption></figure>
11. Slack에 새로운 알림이 있는지 확인합니다: **새로운 Snyk 취약점** 경로, **패키지 이름**, **scss-tokenizer**, **심각도 레벨**, **취약점 이름**, 그리고 **우선 순위 점수**가 포함된 알림이 출력됩니다.

    <figure><img src="https://lh5.googleusercontent.com/1nvqWOgUaA6P6kc7MTObqXxfEXrFaP1DKXqHKy7wQhPxpWIA9HyMHV7dwOHd2HGQiJuL9rwn9aVQlhvlg-rBcHTggXh6nhRWB8T7PAtfM4S73bTL1ytUK3ZaKtzbCnofDUg9ER22zcMI84PXv1byQnN9BUToJk49qiOcq6627VLFlDvUBrXpL1Atjg" alt=""><figcaption><p>취약점 Slack 알림<br><br></p></figcaption></figure>

    다음에 Snyk가 새로운 취약점을 발견하면 Slack에서 Snyk로부터 알림을 받게 될 것입니다.