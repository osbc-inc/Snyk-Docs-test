# 예시: OneLogin에 사용자 정의 매핑 설정

이 예시는 Snyk에 대한 [OneLogin SSO를 구성한 후](../../set-up-snyk-single-sign-on-sso.md) 사용자 역할을 구성하는 방법을 보여줍니다.

OneLogin은 **그룹**과 **역할**의 개념을 갖고 있지만, 사용자에게 여러 그룹을 할당하는 것은 지원하지 않습니다.

따라서 사용자에게 역할을 직접 할당하고 간접적으로 그룹을 통해 할당하지 않습니다.

1. OneLogin에서 **사용자**로 이동한 다음 **역할** 섹션으로 이동하여 [사용자 지정 매핑](../)에 대한 명명 규칙을 따라 역할을 생성합니다. 각 역할은 역할 앱으로 Snyk SAML 앱이 활성화되어 있어야 합니다. 사용자를 필요에 따라 역할에 할당합니다.

<figure><img src="../../../../.gitbook/assets/image (379).png" alt="OneLogin 역할 섹션"><figcaption><p>OneLogin 역할 섹션</p></figcaption></figure>

2. SAML 어설션에서 사용자 역할을 Snyk로 전송하려면 **Applications**로 이동하여 Snyk SAML 앱을 선택한 다음 왼쪽에 있는 **Parameters** 섹션을 선택합니다.

<figure><img src="https://lh6.googleusercontent.com/zseB83vGEsQBiQ2_Rc6zOgkKHkv_KN6S-uLHbZc9k_US_aEzFX1AJUJkEgJpucRtdWYgx0mpUhpHiAhCVTsp3xj2o8hVEB0ArnuMmAVYQ9mw44zULICe57XRZDYxkKHpvpnk6o-TXrqYQHN3MuYMyjA" alt="OneLogin 애플리케이션 Parameters"><figcaption><p>OneLogin Applications Parameters</p></figcaption></figure>

3. **roles**라는 **새 매개변수**를 생성하고 **SAML 어설션에 포함**과 **다중 값 매개변수** 모두 선택된 상태로 확인란을 설정합니다. **저장**합니다.

4. 다음 화면에서 **기본 값**으로 **사용자 역할**을 선택하고 **세미콜론으로 구분된 출력 (다중 값 출력)**을 선택합니다. **SAML 어설션에 포함** 확인란이 선택되어 있는지 확인합니다. **저장**합니다.

<figure><img src="https://lh3.googleusercontent.com/fnsu9d998jEzxyzuIfHl3JSZHBh5iXsPATUj9jL_SZsFoFPFvvus_JyyY3YAeey5ZMtC9oCuhtjrmSMKAVlY8Tq_Sjf9plgDWagoFuLBQX2U0vbFPU76fNvpjSkpJdgL0JsPhXwq3ngBlgJvdsidoyM" alt="OneLogin roles 필드 편집"><figcaption><p>OneLogin Edit Field roles</p></figcaption></figure>