# OneLogin SAML 애플리케이션 설정

이 예제는 OneLogin에서 SAML 애플리케이션을 설정하고 이를 Snyk에 연결하여 SSO를 용이하게 하는 방법을 보여줍니다. OneLogin을 구성하여 SSO를 사용하려면 먼저 Snyk에서 엔티티 ID 및 응답 URL(즉, Assertion Consumer Service URL)을 가져와야 합니다.

1\. 왼쪽 상단의 드롭다운에서 **GROUP OVERVIEW**를 선택한 다음, 그룹 설정으로 가기 위해 **cog**(오른쪽 상단) 아이콘을 클릭합니다.

<figure><img src="https://lh5.googleusercontent.com/nHeI8z3TliigfUaI1lTr46yVvgIYd18vjAf9kVwMgVgcV_X4S6bBJDCNjiOppGQVstJ-XtDD6ZK0ErVzMIj8yXZafaJk4Tu8JKoilGAOuddSRHsIKdpDasnviWAYK50NWFrAU9GTGMVqD_gGSe1pTOI" alt="Select group overview"><figcaption><p>그룹 개요 선택</p></figcaption></figure>

2\. **SSO**를 클릭하고 **엔티티 ID**와 **ACS URL** 아래의 값을 복사하거나 웹 브라우저 탭을 열어 둡니다.

3\. [OneLogin](https://www.onelogin.com/)으로 이동하여 **applications** 메뉴로 이동한 후, **Applications**을 클릭하고 우측 상단의 **Add App** 버튼을 클릭합니다.

<figure><img src="https://lh4.googleusercontent.com/eWStu1dJQcV618MFMWswLT-88RtDQU4XV-dR25IxjMi_lZpvmgQ97FmF3wJlbWHSVG-kNYCfI7Nis0mB050nXeQJKvsw34irMC7fB_XYYu3GivpfmN-d775-3p64qcBSY0Q5ZfsDahcu_YLHuvem5XM" alt="Add app"><figcaption><p>앱 추가</p></figcaption></figure>

4\. "saml"로 필터링하고 **SAML Custom Connector (Advanced)** 앱을 선택합니다.

<figure><img src="https://lh5.googleusercontent.com/NcVS2ScxD3_3l464zhgBhVuxC6hpJLyJy7y5c5uyoYv0cfyY5izIiMnmYQIlrerUusud7bbIpFJjQeSHnDHH7v5CbnVhzBwm8qpoO9ryfpCC8WGo4sw3OpDU1SwZWXHaPtSR1-sGX103CoaugXPEI1w" alt="Select SAML Cusotm Connector (Advanced) app"><figcaption><p>SAML Cusotm Connector (Advanced) app 선택</p></figcaption></figure>

5\. 앱에 대한 **Display Name**을 입력합니다. 예를 들어, Snyk입니다. 옵션으로 아이콘을 선택할 수 있습니다. Snyk 로고는 Snyk의 [press kit](https://snyk.io/press-kit/) 페이지에서 찾을 수 있습니다.

<figure><img src="https://lh6.googleusercontent.com/Ar8VZnNLeqHKP0wgAZYFT4jNo87CTiiNkc4driJsI-ipg8vy13uN_z3CsFGmtnaxbJbpWciw7VH88nzLch68f-jiJOUqbPaiLHJxYZN7F6MZ374IJqzJC7Jj-_ijJefZ3zbvmPtOikZRzHpbln8EtZg" alt="Enter a display name and pick an icon"><figcaption><p>디스플레이 이름 입력 및 아이콘 선택</p></figcaption></figure>

6\. 앱이 저장된 후, **configuration** 페이지로 이동하여 **Audience** (EntityID), **ACS URL**, ACS URL Validator를 입력합니다. ACS URL 값도 여기에 사용합니다. 단계 2에서 엔티티 ID 및 ACS URL을 복사합니다.

<figure><img src="https://lh4.googleusercontent.com/S11TB8rvOOs7abB3bOugmDB041wHIfyFzX9gByH6I12oDLiyiba7ZptPkheT_1wc2hR-QPhiCJgYd4swA_x4zqf1IW-zf2MF7Y4ClvDbgyyX42u12e77_VbQqOow8DPHRVoSFYcecFaHfBj8S3_MKxw" alt="Enter application details"><figcaption><p>애플리케이션 세부 정보 입력</p></figcaption></figure>

7\. **파라미터** 섹션으로 이동하여 **새로운 매개변수**를 추가합니다. Snyk가 SAML 응답의 `email` 속성으로 이메일 주소를 필요로 하기 때문입니다. **+** 아이콘을 클릭합니다.

<figure><img src="https://lh3.googleusercontent.com/XcsNQ0cEhNE-UTJHK2fOMBEM01KIxR3BHc8Y5M6dQnKHMQQuzJEQ6zuRARY3mXzyw6SPo9miw89pxr2bOPk3NuyMqVZAiIiMxibB0jQlH3kDRuWdkBZmKUKAd_8rdPVgB3Bs1T24HQ--3yRIEKAO_sY" alt="Add a new parameter"><figcaption><p>새로운 매개변수 추가</p></figcaption></figure>

**필드 이름**으로 **email**을 사용합니다. 확인란 **Include in SAML assertion**이 선택되어 있는지 확인합니다. **저장**합니다.

<figure><img src="https://lh6.googleusercontent.com/nuR-C1_nGoY87m_fsQUiDhC5dV2nGjyaoyuz_K4uRonw3PB8gWWI3YIvsn0Yp67F2L_yhue-PlaBEYPEsDLjnkvR_hTok-BE4rA4a5xgYWW7Bgu-f44p6J5dSbTVCqZ5lTMHzo2Bpt71Wvt-DCYnpJM" alt="Enter email for the field name"><figcaption><p>필드 이름에 이메일 입력</p></figcaption></figure>

다음 화면에서 **Value** 드롭다운 목록에서 **Email**을 선택합니다. **저장**합니다.

<figure><img src="https://lh5.googleusercontent.com/IgUtsnagxiK8GIFB-FomTnlNWoymq-PWpRnsKqeHJebcjiOi9pK6mAdmW7JG-DRQSuzu2-oxjy90SQVJnDLjFE0nZ9Fo0x_lNLsVwceArXqzK2QlRBrTw9xzVsx7URFHeiw4jAzIYqzq9mK0HcIfReY" alt="Select Email form the value dropdown"><figcaption><p>Value 드롭다운에서 이메일 선택</p></figcaption></figure>

**name** 속성에 대해 동일한 단계를 반복합니다.

<figure><img src="https://lh6.googleusercontent.com/mdS5fhCGEhI1CzJyUVhyv_Wdp3MiWJb33ImkBrcIparoO9FutqssO0668iiov12--VwevXmpVw8HT0cfMuq2P2Jg6aYX1o-d7ODqajSKLCPY-bI2LEt-lAzytx9u_tejJrJZbRE38lhr1H6lTWWXDfk" alt="Enter name for the field name"><figcaption><p>필드 이름에 이름 입력</p></figcaption></figure>

<figure><img src="https://lh6.googleusercontent.com/mqFRW8bqzSEqpNFoHBSXbLsDvTVo0cSbb-B5AjiHd6MaMF6TyKcv1VDIxLMYUbk7CDFGoTzIuNrhssluwVycCV6GLNGAcn8fGRtBE8VSGXQpshmm2L8CrcMm8o1Ve9xPMQ__tSnC9QXBJt3bhxoA0rk" alt="Add the name parameter"><figcaption><p>이름 매개변수 추가</p></figcaption></figure>

8\. **SSO** 섹션으로 이동하여 **certificate**를 다운로드하고 **SAML 2.0 Endpoint (HTTP)** 값을 메모해 둡니다.

<figure><img src="https://lh5.googleusercontent.com/qp6ACOk2bxhJiV8PG0XZIHsC_nUIKTCSu6fhPIybQ9FGI4JPWg6gwv72o00Xj1HEfDcQVNRe9jkrtuK0Bzvserc_NVgl0gVFyFozknHJ34dDyqHIceT3xH-iY753ZP7VeDGTS80baRwalnJFFBgKhbE" alt="SAML 2.o Endpoint (HTTP) value"><figcaption><p>SAML 2.o Endpoint (HTTP) 값</p></figcaption></figure>

9\. 사용자들을 위해 앱을 활성화하려면 **Users > Roles** 섹션으로 이동하여 Snyk라는 새 역할을 만듭니다.\
활성화된 역할 앱으로 Snyk SAML 앱을 선택하십시오. **저장**합니다.\
그런 다음, Snyk를 사용할 수 있도록 하는 사용자에게 이 역할을 지정하세요.

<figure><img src="https://lh4.googleusercontent.com/jZL7kElRSz3PX4LmKkCH1k5vYNCgj2BHqlGHU3dNmJRPIJwQjyMFchWSc6et-m7qeVv2QELr_OWH0IJok0Xwn8OifxWjdfkYqiD2YYs1ubmLBQL2ZM8XAOiPKadNfMSLYoOfMEQ4-JsVCQ0wo0YW4b8" alt="Create Snyk role"><figcaption><p>Snyk 역할 생성</p></figcaption></figure>

10\. Snyk 포털로 돌아가서 2단계로 스크롤하여 단계 8에서의 세부사항을 메모하고, SSO 연결을 통해 사용할 도메인을 포함해 입력하세요.\
**IdP-initiated workflow**를 활성화할지 여부를 확인하고, **Create Auth0 connection**을 클릭하세요.

<figure><img src="https://lh6.googleusercontent.com/N_sEZ9IrkaSDpmkYVGhHTiSUf1kVL3P1VWBjBhIJfZgraVdifO8zFfS9Y6yQYjNlc5ic9mSimYGfw07-cm7LsweGdlywAAv99LqSz5964wne9EOjB_PvPuE8yhyLf3kvmKhRU6vQKhVsKxiGNR9Mb_E" alt="Enter SAML attributes in the Snyk portal"><figcaption><p>Snyk 포털에서 SAML 속성 입력</p></figcaption></figure>

11\. 3단계로 스크롤하여 사용자가 로그인할 때 새로운 사용자가 어떻게 처리되어야 하는지 결정합니다.\
사용하고자 하는 옵션을 선택하세요: **Group member, Org collaborator,** 또는 **Org admin**.\
마지막으로, OneLogin에서 구성한 **프로필 속성**을 입력하고, **저장**을 클릭하여 로그인이 가능한지 확인하세요. 직접 URL 상단에 있는 단계 3의 URL로 또는[일반 SSO 로그인](https://app.snyk.io/login/sso)으로 이동합니다.

<figure><img src="https://lh4.googleusercontent.com/OIEztWL9xGSkLQ1yu2jS8IzU1dLWVuX7YJgfTyHYt3aV_pUn53WWc7qOCZvgK0b2M28SmNsTUDtJJZMdQhhA-5kNA2je71LM-AwHwvyd8UyBtPhfHFEnn0rlCmBEM4tppxVXsiLY78KOLJihIMids0E" alt="Enter profile attributes"><figcaption><p>프로필 속성 입력</p></figcaption></figure>