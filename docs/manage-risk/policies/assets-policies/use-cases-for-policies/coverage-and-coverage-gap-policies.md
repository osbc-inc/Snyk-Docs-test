# Coverage 및 Coverage 갭 정책

Coverage 필터를 사용하여 자산이 제품에 의해 테스트되었는지 확인하고 Coverage 갭 필터를 사용하여 자산이 Set coverage control 정책에서 설정된 커버리지 요구 사항을 충족하는지 확인할 수 있습니다.

{% hint style="info" %}
Coverage 갭 필터는 Coverage 필터의 반대인 것이 아닙니다.

자산이 커버되었을 수 있지만(한 달 전에 스캔되었다면), 동시에 요구 사항이 매일 스캔인 경우 커버리지 갭을 여전히 가질 수 있습니다.
{% endhint %}

다음 비디오는 Snyk 웹 UI에서 커버리지 및 커버리지 갭 필터의 개요를 제시합니다.

{% embed url="https://youtu.be/ozdmnb2cXVY" %}

비디오가 마음에 드셨나요? [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 나머지 코스를 확인해보세요!

## Coverage

### 선택된 모든 제품 포함

Snyk Code와 Snyk OS에서 동시에 스캔된 자산을 식별합니다. 이러한 자산을 식별하는 것은 커버리지 요구 사항을 충족한다는 것을 의미하지 않습니다.

\

<figure><img src="https://lh7-us.googleusercontent.com/1aKKSl4O03NT8YL3qR0K1vpcfEMtlCw9pLYrKJ3Q2OdtVYTqdMbsbtWr7Jq32TzMBKEo1t53c7gaEndbiFVqLObxPcUcw7vmmaaSHO5K7UsgtjVu6FO3kLCp6cT_-CX1CzX5Anst0acYqVom89K9y14" alt="Set Coverage filters"><figcaption><p>커버리지 필터 설정</p></figcaption></figure>

### 선택된 제품 중 최소 하나 포함

Snyk Code 또는 Snyk OS에서 스캔된 자산을 식별합니다. 이러한 자산을 식별하는 것은 커버리지 요구 사항을 충족한다는 것을 의미하지 않습니다.

<figure><img src="https://lh7-us.googleusercontent.com/V9uzAQdi6GRne6GXxQ5cQLYXrMD6BD-HMcDIX5ebRk6OWpgxgkU7JSWf49CsNwciu2WZtCoKY7Eg4gk_7mQOXtsGRRns-Z0z96L4aDQQzT_CD17RVEVr57TJK-mMgYiCZW64z4EK71BjvldkWF8iLe4" alt="최소 하나의 선택된 제품을 포함하도록 커버리지 필터 설정"><figcaption><p>최소 하나의 선택된 제품을 포함하도록 커버리지 필터 설정</p></figcaption></figure>

## Coverage 갭

### 모든 제품 제외

정책에 속하지 않는 모든 자산을 식별합니다. 이를 위해 사용 가능한 모든 제품을 커버하는 커버리지 갭 필터를 생성해야 합니다.

<figure><img src="https://lh7-us.googleusercontent.com/RcfoCkR_1a6-L44Bf55ed7xSX8Loyr57KKyI4oX4yh0j6ce3Oj4fu0XL67v9Ij1XKTES-uwTMgqJBFicBtLwaHKilj1orTi_LU0_dEllCvUE2jhfpJimlXIfRON8-0_DF_Qe__tmFLuKmSTOJoFOxCk" alt="커버리지 갭 필터 설정"><figcaption><p>커버리지 갭 필터 설정</p></figcaption></figure>

### 컴플라이언트 자산 필터링

Snyk Code 또는 Snyk OS 또는 둘 다의 커버리지 요구 사항을 충족하지 않는 자산을 식별합니다.

<figure><img src="https://lh7-us.googleusercontent.com/guCzWVv9SP7H1h6WYSFGwHEVvW3c0DVvg26mHAdxkorPgZI3gYCIH93QN0fXNl71ZDZxucfpROkkjruxuQ_vu5QCjS7_ImROEZlBTYIh-hxZnsM3comPaQpQbsy7s_3MDuFVEiljw2G8szWddXjqPgQ" alt="선택된 제품 중 최소 하나를 포함하도록 커버리지 갭 필터 설정"><figcaption><p>선택된 제품 중 최소 하나를 포함하도록 커버리지 갭 필터 설정</p></figcaption></figure>

\
Snyk Code와 Snyk OS의 커버리지 요구 사항을 모두 충족하는 자산을 식별합니다.&#x20;

\

<figure><img src="https://lh7-us.googleusercontent.com/-Ys7HZ5UcthgyDyPbBNG572CTM04RJ_Tcc1JTa9GrltfSVUM5gvFLrxpNRlV6ZNqRJQOw5hL0QFworAAOVbGYCMM4J-H4z9G8L3BiU3-PEU79GqxAalKB5UvdXxKUIgNEszwH0jUN_7kpos8HLSXvo8" alt="선택된 제품을 제외하도록 커버리지 필터 설정"><figcaption><p>선택된 제품을 제외하도록 커버리지 필터 설정</p></figcaption></figure>

{% hint style="info" %}
커버리지 및 커버리지 갭 필터를 사용하여 이미지 자산의 커버리지와 커버리지 갭을 모니터링할 수 있습니다. 이로써 자산의 위험을 더 완전하게 파악하고 필요할 때 조치를 취할 수 있습니다.
{% endhint %}  