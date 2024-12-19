# Slack 설정을 통한 Snyk 및 AWS Lambda 연결

Snyk가 Slack과 통신할 수 있도록 하려면 먼저 Slack 앱을 통해 수신 웹훅을 설정해야 합니다. 이 웹훅은 개발자가 Slack과 통신할 수 있도록 Slack에서 제공합니다.

웹훅을 설정하려면 다음 단계를 따르세요:

1. [https://api.slack.com/apps](https://api.slack.com/apps)로 이동합니다.
2. **새 앱 만들기**를 선택합니다.

    <figure><img src="https://lh5.googleusercontent.com/qw51g6soQ6IjBf95JM0hhIsON0RqAhwuDbd7p3FA_AoGatQWx_0VcefI7RhEoUuKkuDNmXQNSIxw9aD7T7uhG4YPxvRIsAhDnHtVCGT_PtGuAD3fZDO4Qlye45iz94j7xZb0Ze0g8h16xMNtE-3zhmsw8wmq-m_K6OI1UD8mN-CKbNZEJCynuOHEBg" alt=""><figcaption><p>Slack의 앱</p></figcaption></figure>
3. **From scratch**를 선택합니다.

    <figure><img src="https://lh4.googleusercontent.com/uDE4iWxfnHF0KvGGZlwZwAp39zNvG1Vav8yOSxak5DhOIdOl983GS8Xmr-YJ9WpfiS6WJD2b5yhgbUAxMpm7rDwpQkEH2W2zOSyNQZdDAqDvBFpBMP7uYZwDtPGE3OGt0-g-JW09Dx2RB2wcfghEpc8J47A-DH7fejMkupKPnhrspesfPt45duXivg" alt=""><figcaption><p>Scratch에서 Slack 앱 생성하기</p></figcaption></figure>
4. 앱에 "Snyk"와 같은 이름을 지정합니다.
5. 로고를 설정하고 적절한 배경을 원하면 Snyk 로고 [여기](https://snyk.io/press-kit/)에서 다운로드할 수 있으며, 배경 색상은 #1d1848를 사용합니다.
6. 워크스페이스를 선택합니다.
7. Slack 앱이 생성되면 **기능 추가**를 클릭합니다.
8. **수신 웹훅**을 선택합니다.

    <figure><img src="https://lh3.googleusercontent.com/yc2jyH0npATioGnzPLv5WEmI762OIYoefYVztKfvfAS9iV6yHNudbralS8VfLE0NT2x9TqM7lDCVLfV_27cC6Z82P5qprCIu4FKnVco1FfzsDJb3t6_V5BowDpBYw8GrNEaW8TZGbb1hmXsQflr1eeCTNAhKNpbE-AbUJGnxT65Uu67niA_HdCklQg" alt=""><figcaption><p>Slack 앱에 기능 및 기능 추가하기, 수신 웹훅</p></figcaption></figure>
9. 해당 페이지에서 수신 웹훅을 활성화합니다.

    <figure><img src="../../../../.gitbook/assets/image (1) (4) (1) (1) (1) (1).png" alt=""><figcaption><p>수신 웹훅 활성화</p></figcaption></figure>
10. **워크스페이스에 새 웹훅 추가**를 클릭하여 원하는 채널에 대한 웹훅 URL을 생성합니다.
11. Snyk가 게시할 채널을 선택합니다. 아직 채널을 만들지 않은 경우 [채널을 생성](https://slack.com/intl/en-gb/help/articles/201402297-Create-a-channel)합니다.
12. 웹훅이 생성되면 웹훅 URL을 복사하여 AWS의 다음 단계에서 사용하도록 저장합니다.

<figure><img src="https://lh3.googleusercontent.com/av55N4Y2DyLFYmbrhC2gEjU9CINSP4DWUYfkhJju65Q9mpI-MqkKKsf5H8af2TMVy8f-jP6m-6Y-TAaaFsgf6dJ6LbtgGxfYM-vqAkUU5zYVYSoV8u8jKbFeBI9wWWpi9CFrSYPTM-ee2m7DJYDo1p4uBIf-IxqZSLpkJ4kQhp34lT6-6RQ9QLqIEQ" alt=""><figcaption><p>웹훅 URL이 포함된 생성된 웹훅</p></figcaption></figure>