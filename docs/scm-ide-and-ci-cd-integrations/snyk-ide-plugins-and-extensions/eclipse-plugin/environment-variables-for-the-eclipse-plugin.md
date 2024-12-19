# Eclipse 플러그인을 위한 환경 변수

오픈 소스 종속성 및 IAC 템플릿 파일을 분석하기 위해 플러그인은 Snyk CLI를 사용합니다. CLI는 다음 환경 변수가 필요합니다:

- `PATH`: Maven과 같은 필요한 이진 파일의 경로입니다. `PATH` 변수는 설정 대화 상자의 **`Path`** 필드를 사용하여 수동으로 조정할 수도 있습니다.
- `JAVA_HOME`: Java 종속성을 분석하는 데 사용하려는 JDK의 경로입니다.
- `http_proxy`와 `https_proxy`: `http://username:password@proxyhost:proxyport` 형식의 값으로 설정된 프록시입니다.\
  **참고:** 값의 `http://`는 `https://`로 바뀌지 않습니다. Eclipse에서 프록시 설정을 작성하면 Snyk 플러그인이 언어 서버 및 CLI를 위해 환경 변수를 자동으로 설정합니다.

이러한 변수를 셸 환경에서만 설정(예: `~/.bashrc`를 사용)하면 충분하지 않습니다. 이 때는 Eclipse를 명령줄에서 시작하지 않거나 셸 환경을 사용하여 Eclipse를 시작하는 스크립트 파일을 만들지 않는 한요한다.

- **Windows**에서는 GUI를 사용하거나 [setx](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/setx) 도구를 사용하여 명령 줄에서 변수를 설정합니다.
- **macOS**에서는 `launchd` 프로세스가 Finder에서 Eclipse을 직접 시작하기 위해 환경 변수를 알고 있어야 합니다. 이러한 환경 변수는 `launchctl setenv` 명령을 사용하여 설정합니다(예: 시작 시 또는 사용자 로그인 시에 실행하는 스크립트를 사용).\
  **참고:** macOS UI에 환경 변수를 제공하는 방법은 운영 체제 버전 간에 변경될 수 있으므로 Eclipse 앱을 시작하는 작은 셸 스크립트를 만드는 것이 더 쉬울 수 있습니다. 이는 `~/.bashrc`를 통해 정의할 수 있습니다.
- **Linux**에서는 `/etc/environment` 파일을 업데이트하여 환경 변수를 창 관리자와 UI에 전파할 수 있습니다.

{% hint style="info" %}
Snyk의 Eclipse 플러그인은 Eclipse의 프록시 설정을 사용하지만 환경 변수에서도 프록시 설정을 가져옵니다.
{% endhint %}