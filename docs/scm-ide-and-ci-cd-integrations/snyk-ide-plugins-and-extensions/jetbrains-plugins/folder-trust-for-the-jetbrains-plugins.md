# JetBrains 플러그인에 대한 폴더 신뢰

은 분석을 위해 추가 데이터를 얻기 위해 자동으로 귀하의 컴퓨터에서 코드를 실행할 수 있습니다. 이는 의존성 정보를 얻기 위해 패키지 관리자(예: pip, Gradle, Maven, Yarn, npm 등)를 호출하는 것을 포함합니다. 악성 구성을 가진 신뢰할 수 없는 코드에서 이러한 프로그램을 호출하면 시스템이 악성 코드 실행과 악용에 노출될 수 있습니다.

신뢰할 수 없는 프로젝트에서 확장 기능을 사용하는 것을 방지하기 위해 Snyk 확장 프로그램은 코드를 스캔하기 전에 프로젝트 신뢰를 요구합니다. 의심스러울 때는 스캔을 진행하지 마십시오.

<figure><img src="../../../.gitbook/assets/modal-dialog copy.png" alt="Snyk 확장 프로그램에서 프로젝트를 신뢰할 것인지를 알리는 프롬프트"><figcaption><p>Snyk 확장 프로그램에서 프로젝트를 신뢰할 것인지를 알리는 프롬프트</p></figcaption></figure>

한 번 프로젝트를 신뢰하면, Snyk은 해당 프로젝트 폴더 및 하위 폴더에 대한 신뢰를 다시 요청하지 않습니다. 기존 폴더 신뢰를 철회하려면, [JetBrains IDE 구성 옵션 디렉토리](https://www.jetbrains.com/help/idea/directories-used-by-the-ide-to-store-settings-caches-plugins-and-logs.html#config-directory)에 위치한 `snyk.settings.xml`에서 `TRUSTED_PATHS` 옵션을 수동으로 편집하십시오.

Snyk은 추가적인 보안층으로 IntelliJ 플랫폼 버전 2021.2.4/2021.3.1 이상에서 제공되는 기본 [Trusted Projects](https://plugins.jetbrains.com/docs/intellij/trusted-projects.html) 사용을 권장합니다.
