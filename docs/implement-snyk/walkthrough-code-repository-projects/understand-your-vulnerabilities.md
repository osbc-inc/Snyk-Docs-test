# Understand your vulnerabilities

{% hint style="info" %}
**요약**\
[스캔된 프로젝트를 본 후 이해했습](view-your-first-snyk-projects.md)니 이제 해당 프로젝트의 취약성 세부 정보를 살펴볼 수 있습니다.
{% endhint %}

## **취약성 세부 정보 보기**

먼저, Snyk 프로젝트를 확인하려면 대상을 열어보세요:

<figure><img src="../../.gitbook/assets/image (354).png" alt="가져온 프로젝트 보기"><figcaption><p>가져온 프로젝트 보기</p></figcaption></figure>

다음으로, 해당 목록에서 프로젝트를 선택하여 해당 프로젝트에서 발견된 취약성 세부 정보를 확인하세요.

예를 들어, **코드 분석** 프로젝트를 로 스캔한 경우:

<figure><img src="../../.gitbook/assets/image (149) (1) (1) (2).png" alt="취약성 예시 - 코드 분석"><figcaption><p>취약성 예시 - 코드 분석</p></figcaption></figure>

자세한 내용은 [프로젝트 정보 보기](../../snyk-admin/snyk-projects/project-information.md)를 참조하세요.

## 이슈 카드 보기

이제 각 Snyk 프로젝트의 취약성 정보를 제공하는 이슈 카드를 살펴보세요:

<figure><img src="../../.gitbook/assets/image (101) (2).png" alt="취약성 세부 정보 이슈 카드"><figcaption><p>취약성 세부 정보 이슈 카드</p></figcaption></figure>

다시 말하지만 이해할 내용이 많으므로 취약성에 대한 이 모든 정보가 어떻게 관련되는지 이해하는 데 시간을 투자하여 어떤 조치를 취할지 결정하는 데 도움이 되도록 해야 합니다.

자세한 내용은 [이슈 카드 정보](../../snyk-admin/snyk-projects/issue-card-information.md)를 참조하세요.

## 추가 취약성 정보 액세스

Snyk는 취약성에 대한 자세한 정보를 위한 자원을 제공하며, 카드에서 직접 액세스할 수 있습니다:

* [**Snyk 취약성 데이터베이스**](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md): 특정 취약성에 대한 자세한 내용을 액세스할 수 있습니다.
* [**Snyk 학습**](../../getting-started/snyk-learn/): 해당 유형의 취약성에 대한 일반 정보에 액세스할 수 있습니다.

### Snyk 취약성 데이터베이스 액세스

오픈 소스 및 컨테이너 취약성의 경우, 취약성 식별자 옆에 있는 Snyk 취약성 식별자를 클릭하여 해당 취약성에 대한 세부 [Snyk 취약성 데이터베이스](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md) 정보에 액세스할 수 있습니다. 예를 들면:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FAqaPY3eauZpcQ02Rcye7%252Fimage.png%3Falt%3Dmedia%26token%3D79355b4a-e8d8-4bff-b629-58d23c906952&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=b5a32e9d&#x26;sv=2" alt="Snyk 취약성 데이터베이스에 액세스"><figcaption><p>Snyk 취약성 데이터베이스에 액세스</p></figcaption></figure>

위 예시에서는 Snyk 취약성 데이터베이스에서 Hibernate Core 및 해당 라이브러리가 SQL 인젝션에 취약하다는 것을 확인할 수 있습니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Fgit-blob-b4bd2161ca3811f4d9a0d5d02e0b3bf4197f8b8b%252Fimage%2520%28149%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%282%29.png%3Falt%3Dmedia&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=10a3cdec&#x26;sv=2" alt=""><figcaption><p>Snyk 취약점 데이터베이스 예시 항목</p></figcaption></figure>

{% hint style="info" %}
및이슈 카드에는 해당 영역에 대한 별도의 정보 세트가 있습니다.
{% endhint %}

### Snyk 학습 액세스

**이 유형의 취약성에 대한 학습**을 클릭하여 [Snyk Learn](https://learn.snyk.io/) 보안 교육 자료에 액세스할 수 있습니다:

<figure><img src="../../.gitbook/assets/image (119) (1).png" alt="취약성 카드에서 Snyk 학습 액세스"><figcaption><p>취약성 카드에서 Snyk 학습 액세스</p></figcaption></figure>

예를 들어, 이 유형의 취약성에 대한 자세한 내용을 알아보려면 [Snyk Learn SQL 인젝션](https://learn.snyk.io/lessons/sql-injection/javascript/)을 확인하세요.

{% hint style="info" %}
일부 카드에는 Snyk Learn 레슨을 이용할 수 없는 경우가 있습니다 - 그렇다면 링크가 표시되지 않습니다.
{% endhint %}

## Snyk 우선 순위 점수 이해

0부터 1,000까지의 [Snyk 우선 순위 점수](../../manage-risk/prioritize-issues-for-fixing/priority-score.md)는 취약성의 심각성을 평가한 것입니다. Snyk 우선 순위 점수에는 [CVSS](https://www.first.org/cvss/calculator/3.1) (공통 취약점 점수 계산 시스템) 정보뿐만 아니라 공격 복잡성 및 알려진 공격에 대한 요소 등이 포함됩니다. 예를 들어, **Hibernate** 취약성은 해당 취약성을 악용할 수 있는 알려진 공격 수단이 없는 것입니다.

다른 요소도 점수에 영향을 줍니다. 예를 들어, SQL 인젝션은 실행하기 쉽지만 (웹 브라우저를 열고 양식을 제출하기만 하면 됨) 실행 및 해당 공격에 대한 결과를 이해하고 악용해야 하므로 점수가 감소합니다.

## 오픈 소스 취약성: 수정 및 종속성 정보

Snyk 오픈 소스에 의한 오픈 소스 라이브러리 스캔의 경우 프로젝트 결과의 **수정 사항** 및 **종속성** 탭에서 수정 및 종속성 정보를 액세스할 수 있습니다.

### 수정 탭

Snyk는 프로젝트의 전이 종속성에 대한 지식을 토대로 **수정 사항** 탭에서 수정 조언을 제공할 수 있습니다:

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.57.07.png" alt="오픈 소스 취약성에 대한 수정 권장 사항"><figcaption><p>오픈 소스 취약성에 대한 수정 권장 사항</p></figcaption></figure>

자세한 내용은 [첫 번째 취약성 수정](fix-your-first-vulnerability.md)을 참조하세요.

### 종속성 탭

Snyk는 애플리케이션의 패키지 관리자를 사용하여 종속성 트리를 구축하고 이를 프로젝트 뷰의 **종속성** 탭에서 표시합니다:

<figure><img src="../../.gitbook/assets/image (119) (1) (1).png" alt="오픈 소스 취약성의 종속성"><figcaption><p>오픈 소스 취약성의 종속성</p></figcaption></figure>

파일 트리 아이콘을 클릭하여 (![](<../../.gitbook/assets/image (201) (1) (1) (1) (1) (1) (1) (2).png>)) 종속성 트리를 구축하여 취약성을 도입한 구성 요소가 어떤 것인지 확인할 수 있습니다. 이는 어플리케이션에 어떻게 종속성이 도입되었는지 이해하는 데 도움이 됩니다:

<figure><img src="../../.gitbook/assets/image23 (1) (1).png" alt="종속성 트리 세부사항"><figcaption><p>종속성 트리 세부사항</p></figcaption></figure>

예를 들어, 위 스크린샷은 직접 종속성인 **body-parser@ 1.9.0**에서 가져온 **qs@2.2.4**에 기초한 취약성을 보여줍니다.

이제 취약성 정보를 이해했으므로 어떻게 수정할지 결정할 수 있습니다.

[첫 번째 취약성 수정](fix-your-first-vulnerability.md)을 계속해주세요.
