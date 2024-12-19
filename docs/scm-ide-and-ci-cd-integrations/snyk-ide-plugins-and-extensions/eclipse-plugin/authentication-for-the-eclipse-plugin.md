# Eclipse 플러그인을 위한 인증

프로젝트를 스캔하려면 Snyk와 인증해야 합니다. 

Snyk는 다음 프로토콜을 지원합니다:

- OAuth 2.0 (기본)
- Snyk API 토큰 (대체 옵션)

## OAuth 2.0 프로토콜을 사용하여 인증하는 단계

확장 프로그램이 설치된 후, 다음 단계에 따라 인증을 수행합니다:

1. 열리는 대화 상자에서 사용자 정의 멀티 테넌트 또는 단일 테넌트 설정을 위한 Snyk API 엔드포인트를 설정합니다. 기본값은 `https://api.snyk.io`입니다. 자세한 내용은 [IDE URLs](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)를 참조하세요.

<figure><img src="../../../.gitbook/assets/SCR-20240822-mgxw (1).png" alt="Snyk 테넌트 구성" width="563"><figcaption><p>Snyk 엔드포인트 구성</p></figcaption></figure>

2. 추가 정보가 포함된 다음 페이지에서 **완료**를 클릭합니다:

<figure><img src="../../../.gitbook/assets/SCR-20240822-mibb (1).png" alt="추가 정보 및 완료" width="563"><figcaption><p>추가 정보 및 완료</p></figcaption></figure>

3. 확장 프로그램이 기본 브라우저에서 새 페이지를 열고 Snyk 계정으로 로그인하라고 요청합니다:

<figure><img src="../../../.gitbook/assets/SCR-20240821-qogt.png" alt="Snyk 로그인" width="375"><figcaption><p>Snyk 로그인</p></figcaption></figure>

4. 다음 페이지에서 IDE가 귀하를 대신하여 작동할 수 있는 권한을 요청합니다. **앱 액세스 부여**를 클릭합니다.

<figure><img src="../../../.gitbook/assets/SCR-20240821-qnpy.png" alt="앱 액세스 부여" width="375"><figcaption><p>앱 액세스 부여</p></figcaption></figure>

5. 성공적으로 인증한 후 확인 메시지를 확인합니다.

<figure><img src="../../../.gitbook/assets/SCR-20240821-qrgp.png" alt="성공적인 인증" width="375"><figcaption><p>성공적인 인증</p></figcaption></figure>

6. IDE는 로컬 기계에 인증 토큰을 읽고 저장합니다. 브라우저 창을 닫고 IDE로 돌아갑니다.

분석이 자동으로 시작됩니다. 문제가 있는 경우 [OAuth 2.0 인증이 작동하지 않음](../troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md)을 참조하세요.

{% hint 스타일="info" %}
OAuth 2.0 토큰은 정적이 아니며 Snyk 계정 페이지에서 복사할 수 없습니다.
{% endhint %}

## Snyk API 토큰을 사용하여 인증하는 단계

{% hint 스타일="warning" %}
이 방법은 OAuth 방법에 비해 불안정합니다.
{% endhint %}

API 토큰을 사용하여 인증하려면 다음 단계를 따르세요:

1. **프레퍼런스** > **Snyk**로 이동합니다.
2. **토큰 인증 사용** 플래그를 설정합니다.
3. API 토큰을 복사합니다. 자세한 내용은 [Snyk API 토큰 얻기 및 사용](../../../getting-started/#obtain-and-use-your-snyk-api-token)를 참조하세요.
4. 개인 토큰을 사용해야 합니다. 토큰을 **토큰** 필드에 붙여넣거나 입력하세요.
5. **적용 및 닫기**를 클릭합니다. 분석이 자동으로 시작됩니다.

## 계정 전환 방법

다른 계정으로 다시 인증하려면 다음 단계를 따르세요:

1. **프레퍼런스** > **Snyk**로 이동합니다.
2. **토큰** 필드의 값을 지웁니다.
3. **적용 및 닫기**를 클릭합니다.
4. 로그아웃 후 처음부터 인증을 시작합니다.