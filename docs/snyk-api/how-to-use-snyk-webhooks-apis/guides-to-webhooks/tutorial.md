# Snyk 웹훅을 Zapier와 함께 사용하는 방법

{% hint style="info" %}
Snyk API v1 문서는 [여기](https://snyk.docs.apiary.io/)에서 확인할 수 있습니다.
{% endhint %}

## 통합 예시

우선 [Zapier](https://zapier.com)에서 새로운 Zap을 만들어야 합니다.

### 트리거

요청 헤더에 액세스하기 위해 **"Catch Raw Hook"** 트리거를 생성해야 합니다. 이 방식은 요청 페이로드가 문자열로 제공되어 JSON으로 파싱해야 한다는 단점이 있습니다.

![](<../../../.gitbook/assets/untitled (1) (1) (1) (1) (1) (1) (1).png>)

이렇게 하면 요청을 보낼 수 있는 Webhook URL을 받게 됩니다:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-1%20\(1\).png)

이제 제공된 URL로 Snyk에 API를 통해 Webhook을 생성해야 합니다.

```
POST /api/v1/org/{orgId}/webhooks HTTP/2
> Host: snyk.io
> Authorization: token {authToken}
> Content-Type: application/json
| {
|   "url": "https://hooks.zapier.com/hooks/catch/9002958/oemlgkz/",
|   "secret": "my-secret-string"
| }
```

API는 거의 완성된 Webhook으로 응답할 것입니다.

```
< HTTP/2 200 
< Content-Type: application/json
| {
|   "id": "{webhookId}",
|   "url": "https://hooks.zapier.com/hooks/catch/9002958/oemlgkz/",
| }
```

이제 Zapier의 트리거를 테스트하기 위해 Webhook에 핑을 보낼 수 있습니다.

```
> POST /api/v1/org/{orgId}/webhooks/{webhookId}/ping HTTP/2
> Host: snyk.io
> Authorization: token {authToken}
> Content-Type: application/json
```

핑 요청을 선택하고 필드를 매핑할 수 있게 됩니다.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-2%20\(1\).png)

### 액션 (페이로드 유효성 검증)

페이로드를 유효성 검사하려면 JS 액션을 만들어야 합니다:

**"Code by Zapier" → "Run Javascript"**

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-3%20\(1\).png)

`headers['X-Hub-Signature']` 및 페이로드 문자열을 스니펫 변수에 매핑해야 합니다.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-4%20\(1\).png)

다음의 스니펫은 `isValid: boolean` 변수를 Zap의 필드에 도입할 것입니다.

{% hint style="info" %}
`my-secret-string` 문자열을 웹훅의 비밀 문자열로 교체하세요.
{% endhint %}

```javascript
const crypto = require('crypto');
const secret = "my-secret-string";

function makeSignature(body, secret) {
  const hmac = crypto.createHmac('sha256', secret);
  hmac.update(body, 'utf8');

  return `sha256=${hmac.digest('hex')}`;
}

try {
  const body = JSON.parse(inputData.body);
  const { project, org, group, newIssues } = body;

  output = { 
    isValid: inputData.signature === makeSignature(inputData.body, secret)
  };
} catch (err) {
  output = { isValid: false, err: err.message };
}
```

스니펫을 테스트하여 `isValid === true`임을 확인하세요.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-5%20\(1\).png)

### 액션 (페이로드 파싱)

이제 문자열에서 Zapier가 이해할 수 있는 형식으로 페이로드를 파싱하기 위해 또 다른 액션을 생성해야 합니다.

동일한 JS 액션을 만들어야 합니다:

**"Code by Zapier" → "Run Javascript"**, 필드 매핑은 다음과 같이 수행됩니다:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-6%20\(1\).png)

다음의 JS 스니펫을 사용합니다:

```
try {
  output = JSON.parse(inputData.body);
} catch (err) {
  output = { err: err.message };
}
```

이렇게 하면 요청 페이로드가 파싱되어 Zap의 변수로 매핑됩니다.

### 액션 (이슈 형식화)

새로운 이슈는 객체 목록으로 제공됩니다. 이 형식은 Zapier가 잘 이해하지 못하므로, 대신 문자열 목록을 선호합니다. 따라서 `newIssues`를 `string[]`로 형식화해야 합니다.

하나의 추가 JS 액션을 더 만들어야 합니다:

**"Code by Zapier" → "Run Javascript"**, 다음 스니펫을 붙여넣으세요:

```
function formatIssue({ pkgName, pkgVersions, issueData }) {
  return `
  <a href="${issueData.url}">${issueData.title}</a><br/>
  Vulnerability in ${pkgName} (${pkgVersions.join(', ')}). ${issueData.severity} severity.
`;
}

try {
  const { newIssues, ...body } = JSON.parse(inputData.body);

  output = { ...body, newIssues: newIssues.map(formatIssue) };
} catch (err) {
  output = { newIssues: [], err: err.message };
}
```

### 액션 (필터)

모든 필드가 제공되었으므로, 이벤트에 대해 무언가를 수행할지 여부를 결정할 수 있습니다.

필터링하려면 **"Filter by Zapier"** 앱을 만들어야 합니다:

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-7%20\(1\).png)

이제 어떻게 필터링할지 선택할 수 있게 됩니다.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-8%20\(1\).png)

### 액션 (알림 보내기)

위의 작업을 통해 필요한 모든 필드에 액세스할 수 있으므로, 알림 템플릿을 작성할 수 있습니다. 저의 경우, 이메일을 보내기로 선택했지만, 다른 작업을 선택할 수도 있습니다.

![](https://partner-workshop-assets.s3.us-east-2.amazonaws.com/untitled-9%20\(1\).png)