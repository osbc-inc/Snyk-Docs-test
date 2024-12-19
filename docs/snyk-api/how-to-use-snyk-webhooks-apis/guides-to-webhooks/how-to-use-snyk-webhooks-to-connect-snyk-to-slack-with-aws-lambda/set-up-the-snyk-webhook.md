# Snyk 웹훅 설정

[Snyk API v1](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook)을 사용하여 내장 콘솔을 포함하여 Snyk 웹훅을 설정합니다.

Snyk 웹훅을 설정하려면 다음 단계를 따르세요:

1. Snyk 웹 UI의 **Organization** 설정에서 **Organization ID**를 복사합니다.

    <figure><img src="https://lh3.googleusercontent.com/n5_nk9_s2lIb982FQV8LwIQzgYWxC6xeDKEiZMnN_TvrAuDV5oWvCR2RO15XMzyhvVpQwpg1IcL97ljvhis1Q3hfynm91EEqRQvaA7mdkeholt_JvmKPeq1eVmgmnQu5Iaahmdl4UC_8oPP4A6kSGUBO7iz0YPrBca4hbhXOLndO_DLK0NkPPK4dmQ" alt=""><figcaption><p>Snyk 웹 UI에서 Organization ID 복사</p></figcaption></figure>
    
2. Snyk 웹 UI에서 Organization admin **API Token**을 Service Account 또는 본인 계정에서 얻습니다.

    <figure><img src="../../../../.gitbook/assets/account-settings-general-auth-token.png" alt=""><figcaption><p>Snyk 웹 UI에서 API Token 가져오기</p></figcaption></figure>
    
3. Snyk API v1에서 **Console**로 전환하고 조직 ID를 매개변수로 추가합니다.

    <figure><img src="https://lh3.googleusercontent.com/-sXMkOgM3GdCYP-15KqxtZ5DhxZlV3coqUZLYNdNnpVSdCFMH7wZApPhJAr9_8JxzAqyZOFGdIpqjT1t5Jpj570jQ67ykj_L3db4Gph3s74QOXdXjTwEJdRHRfWW0jpY14_lBAOinKC4x1An7yIIfHI-lk-cMULUosb8uDxC_z9mleGNkbdwUC3zVA" alt=""><figcaption><p>Snyk API v1에 orgId 추가하여 webhook 생성 요청을 POST</p></figcaption></figure>
    
4. 헤더 섹션에 Snyk API 키를 **Authorization**에 추가합니다.

    <figure><img src="https://lh6.googleusercontent.com/nhlX0u7hJZSTue4rK01FLvComCMVmEQc1uE_z0nsnQ2_uK0ew5TFryBrTBkL24AKj03NjwKZvK5DsoN6j3fdKu0K9lX2a6SN2JP30m5-ST_Fj-IlMYO4Nu6PwDaDMeQH0ZPzyCF7__zc77iIaHRxxV2_57JDmgv7NbCeJi3Ti3LwP5K9UyYpkrma1A" alt=""><figcaption><p>Snyk API v1에 API 키를 Authorization에 추가</p></figcaption></figure>
    
5. 본문 섹션에 값을 추가합니다\
    `{`\
    `“url”: “당신의 공개 URL 값”,`\
    `“secret”: “당신의 람다 시크릿 환경 변수 값”`\
    `}`

    <figure><img src="https://lh5.googleusercontent.com/VXsSM6NFIWtWa_D4t_pJsWMUm3jHLMxSTEH8N7uLmb7IX98oxfm80_nPg0F6SGd-ffqth-iH3a2afcRQvE58hl5YoAP0NfvfaSPeUP6osRYdnPiPd1-ZOGUajvFk3vvOfXye_khV6lOylFC-T-47nLjclQD7ls8soL-EbWa8KAznWZJeLtj05eshSQ" alt=""><figcaption><p>Snyk API v1 POST 본문</p></figcaption></figure>
    
6. **Call Resource**를 클릭합니다.

이 요청을 완료하면 Snyk에서 Slack으로의 연결이 완료됩니다. 새로운 취약점이 발생할 때마다 알림을 받게 됩니다.