# npm Teams와 npm Enterprise 통합

{% hint style="info" %}
**기능 가용성**\
이 기능은 엔터프라이즈 요금제에서만 사용할 수 있습니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans/)을 참조하십시오.

이 안내서는 Snyk 웹 UI 통합에 관련이 있습니다. Snyk CLI는 이미 개인 npm Teams 및 npm Enterprise 레지스트리를 지원하는 Yarn 및 npm 프로젝트를 지원합니다.
{% endhint %}

Snyk은 npm 및 Yarn 프로젝트에서 사용자 정의 npm Teams 및 npm Enterprise 저장소를 사용할 수 있습니다.

이를 통해 Snyk은 사용자 정의 레지스트리에 호스팅되는 패키지의 모든 직접 및 간접 종속성을 해결하고 더 완전하고 정확한 종속성 그래프와 관련된 취약점을 계산할 수 있습니다.

구성한 후에 Snyk은 또한 이 정보를 사용하여 풀/병합 요청을 만들 때 개인 종속성에 액세스하고, npm과 yarn이 해당 종속성에 도달하여 lockfile을 다시 생성할 수 있도록 허용합니다.

당신은 Snyk에게 개인 npm Teams 및 npm Enterprise Node.js 패키지가 호스팅되고 어떤 범위 아래에 있는지 알려주는 구성을 추가할 수 있습니다.

이는 보통 `.yarnrc` 또는 `.npmrc`에 추가하는 정보와 동일합니다.

## JavaScript 언어 설정

1. **설정 > 언어 > JavaScript**로 이동하여 프로젝트 유형에 따라 npm 또는 Yarn 설정을 선택합니다. Yarn은 스크린샷에서 표시됩니다.
2. 이전에 npm Teams 또는 npm Enterprise에 연결하지 않은 경우, 먼저 통합을 구성하라는 메시지가 표시됩니다. [npm Teams 및 npm Enterprise 레지스트리 설정](npm-teams-and-npm-enterprise-integration.md#npm-teams-and-npm-enterprise-registry-settings)를 참조하십시오.
3. 통합을 설정한 후 **레지스트리 구성 추가**를 선택합니다.
   1. **Package 소스**로 **npm**을 선택합니다.
   2. **기본 레지스트리 URL**로 이 레지스트리를 설정하려면 **범위**를 비워 둡니다.
   3. **이 레지스트리를 사용하려면 범위에만 구성**하려면 **범위**를 추가합니다.
   4. **기본 레지스트리 URL 및 범위 패키지를 추가**하려면, 기본 및 각 범위당 하나의 구성을 추가합니다.
4. 원하는 모든 레지스트리와 범위를 추가했을 때 **설정 업데이트**를 클릭합니다.
5. 통합을 테스트하려면, 이전에 생성되지 않았던 Snyk 픽스  요청에 업데이트된 lockfile을 포함한 개인 종속성이 포함된 프로젝트에서 풀/병합 요청을 엽니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Fgit-blob-ccdc82531288c5468d46e0d49f926b65c1fe24be%252Fimage%2520%2834%29.png%3Falt%3Dmedia%26token%3De8e028ad-0767-4bfc-8fa1-e2c9928eb334&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=5c77f962&#x26;sv=2" alt="npm 통합 테스트"><figcaption><p>npm 통합 테스트</p></figcaption></figure>

## npm Teams 및 npm Enterprise 레지스트리 설정

npm Teams 및 npm Enterprise 통합에 대해 토큰 기반 인증을 구성할 수 있습니다.

1. **설정 > 통합 > 패키지 저장소 > npm**으로 이동합니다.\
   처음에 **자격 증명** 화면이 표시됩니다.
2. **공개 URL** 및 **토큰** 값을 입력합니다.
3. **저장**을 클릭합니다.

<figure><img src="../../../.gitbook/assets/image (35) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="npm 자격 증명 화면"><figcaption><p>npm 자격 증명 화면</p></figcaption></figure>
