# 워크스페이스 신뢰

Snyk은 취약점을 조사하기 위해 코드베이스를 검사하는 일환으로 귀하의 컴퓨터에서 코드를 자동으로 실행하여 추가 데이터를 획들할 수 있습니다. 예를 들어, 이는 Snyk 오픈 소스용 의존성 정보를 얻기 위해 패키지 관리자(예: pip, gradle, maven, yarn, npm 등)를 호출하는 것을 포함합니다. 악의적인 구성을 갖는 신뢰할 수 없는 코드에서 이러한 프로그램을 호출하는 것은 시스템을 악성 코드 실행과 악용에 노출시킬 수 있습니다.

신뢰할 수 없는 폴더에서 확장 프로그램을 사용하는 것에 대비하여, Snyk 확장 프로그램은 코드에 대한 스캔을 실행하기 전에 해당 폴더의 신뢰를 요청할 것입니다. 의심스러울 경우 스캔을 진행하지 마십시오.

<figure><img src="https://lh4.googleusercontent.com/zeg6dMXmw1BUVMTeCviRJ-b2ct0CcRXFdfm1AZJFfgzRfro6TzvFv2_4DwZbqcRRHUl20poHvThC-8TdJniBQ12zgFxgZCY7b5NLZcG4XOu2qd25avtIEZM9Hzq-LvqKmqDS4N3uAFydWTT99x4HMNi0g_2LwFH2LzKwD98KAt1gYxZHitFK3PwbpB7pFA" alt="폴더를 신뢰하는 프롬프트"><figcaption><p>폴더를 신뢰하는 프롬프트</p></figcaption></figure>

한 번의 폴더 신뢰가 부여된 후에는 Snyk은 열린 폴더 및 해당 하위 폴더에서 다시 신뢰를 요청하지 않을 것입니다. 기존의 폴더 신뢰를 취소하려면 다음 경로에 위치한 **settings.json** 파일에 있는 **trustedFolders** JSON 요소를 수동으로 편집하십시오:

`%LocalAppData%\Microsoft\VisualStudio\<Visual Studio version>\Extensions\Snyk\Snyk Security\<Snyk extension version>`