# JetBrains 플러그인용 인증

프로젝트를 스캔하려면 Snyk으로 인증해야 합니다.

{% hint style="warning" %}
인증하기 전에 지역을 올바르게 설정했는지 확인하세요. 자세한 내용은 [지역 호스팅 및 데이터 저장](../../../working-with-snyk/regional-hosting-and-data-residency.md)를 참조하십시오. [지역별 URL 목록](../../../working-with-snyk/regional-hosting-and-data-residency.md#regional-urls)도 확인할 수 있습니다.
{% endhint %}

Snyk은 다음과 같은 프로토콜을 지원하여 인증합니다:

* OAuth 2.0 (기본)
* Snyk API 토큰 (대안 옵션)

## OAuth 2.0 프로토콜을 사용하여 인증하는 단계

다음 단계에 따라 인증하세요:

1. 확장 프로그램을 설치한 후, 내비게이션 바의 **Snyk 아이콘**을 클릭한 다음 **프로젝트 신뢰 및 스캔**을 클릭합니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252F3Jbz4DhVMaexKQpCZeCm%252FSCR-20240821-twbu.png%3Falt%3Dmedia%26token%3D7ca02316-3500-4418-a413-93938b69f225&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=fe4685e0&#x26;sv=2" alt="Snyk 아이콘 및 연결 및 신뢰"><figcaption><p>Snyk 아이콘 및 연결 및 신뢰</p></figcaption></figure>

2. 확장 프로그램이 기본 브라우저에서 새 페이지를 열고 Snyk 계정에 로그인하도록 요청합니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FuJIIqXr8i3KqhoEO15rl%252FSCR-20240821-qogt.png%3Falt%3Dmedia%26token%3Dabb4fa92-53ea-4d10-bad5-0b6f46eee9e1&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=1e960e53&#x26;sv=2" alt="Snyk 로그인" width="375"><figcaption><p>Snyk 로그인</p></figcaption></figure>

3. 다음 페이지에서 IDE가 대신할 수 있는지를 승인하기 위해 권한을 요청합니다. **앱 액세스 부여**를 클릭합니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FMJFAK5Z7El4iwNgTpeC3%252FSCR-20240821-qnpy.png%3Falt%3Dmedia%26token%3Dfc7df895-8bf4-4a5a-8682-2355df1cd921&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=d245441b&#x26;sv=2" alt="앱 액세스 부여" width="375"><figcaption><p>앱 액세스 부여</p></figcaption></figure>

4. 성공적으로 인증했을 때 확인 메시지가 표시됩니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FyPOisvpzyLdvxhFsnZfm%252FSCR-20240821-qrgp.png%3Falt%3Dmedia%26token%3D05c46314-750c-45d8-9c8e-1ad8906f6eba&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=629f155a&#x26;sv=2" alt="성공적 인증" width="375"><figcaption><p>성공적 인증</p></figcaption></figure>

5. IDE는 로컬 기계에 인증 정보를 읽고 저장합니다. 브라우저 창을 닫고 IDE로 돌아갑니다.

분석이 자동으로 시작됩니다. 문제가 있는 경우 [OAuth 2.0 인증이 작동하지 않음](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md)을 참조하세요.

{% hint style="info" %}
OAuth 2.0 토큰은 정적이 아니며 Snyk 계정 페이지에서 복사할 수 없습니다.
{% endhint %}

## Snyk API 토큰을 사용하여 인증하는 단계

{% hint style="warning" %}
이 방법은 OAuth 방법보다 뒤떨어진 방법입니다.
{% endhint %}

다음 단계에 따라 인증하세요:

1. JetBrains 플러그인에서 **Preferences** > **Tools** > **Snyk**으로 이동합니다.
2. **Authentication Method**를 찾아 **Token authentication**으로 변경합니다.
3. API 토큰을 복사합니다. 자세한 내용은 [Snyk API 토큰을 가져오고 사용](../../../getting-started/#obtain-and-use-your-snyk-api-token)을 참조하세요.
4. **Token** 필드에 붙여넣거나 입력합니다.
5. **Apply** 또는 **OK**를 클릭합니다. 분석이 자동으로 시작됩니다.

## 계정 전환 방법

다른 계정으로 다시 인증하려면 다음 단계를 따르세요:

1. JetBrains 플러그인에서 **Preferences** > **Tools** > **Snyk**으로 이동합니다.
2. **Token** 필드의 값을 지우세요.
3. **Apply** 또는 **OK**를 클릭합니다.
4. 로그아웃되면 처음부터 인증을 시작하세요.

## Linux 및 Unix 요구 사항

Snyk와 인증할 때 Linux 및 Unix 사용자는 자신의 클립보드에 인증 URL을 복사할 수 있습니다.

Linux 및 Unix 사용자는 `xclip` 또는 `xsel` 유틸리티가 설치되어 있어야 합니다.
