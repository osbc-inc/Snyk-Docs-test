# Eclipse 플러그인 문제 해결

{% hint style="warning" %}
Snyk 플러그인은 유통업체의 End Of Life (EOL)에 도달한 모든 운영 체제를 지원하지 않습니다.
{% endhint %}

## 일반적인 문제 해결

일반적인 문제 해결 방법은 [IDE 문제 해결 페이지](../troubleshooting-ides/)를 참조하십시오.

## 로그

{% hint style="warning" %}
`debug`를 활성화하면 코드가 IDE 로그 파일에 기록될 수 있습니다. 예를 들어 `idea.log` 파일에 기록될 수 있습니다.
{% endhint %}

플러그인 로그가 저장되는 위치를 확인하려면 **Preferences** > **Language Servers** > **Logs**로 이동하여 Eclipse에서 **Snyk Language Server** 항목을 찾으세요. Language Server는 비활성화할 수 있으므로 로그를 가져오기 위해 활성화해야 할 수도 있습니다. 로그는 설정된 기본값에 따라 콘솔이나 파일에서 확인할 수 있습니다.

추가적인 플러그인 오류 로그를 보려면:

1. **Window** > **Show View** > Others...로 이동합니다.
2. **type text filer**에 **Error Log**를 검색합니다.
3. **Open**을 클릭하여 오류 로그 탭을 엽니다. 탭 보기에서 플러그인별로 그룹화하려면 (오른쪽 상단의 세 개의 점 메뉴 > **Group By** > **Plug-in**) `io.snyk.eclipse.plugin` 행에서 플러그인 오류를 확인할 수 있습니다.

로그 레벨은 Eclipse 설정에서 환경 변수 `SNYK_LOG_LEVEL`을 사용하여 설정할 수 있습니다. 예를 들어 `SNYK_LOG_LEVEL=debug`로 설정할 수 있습니다.

Language Server의 설치 경로는 플러그인 설정에서 표시됩니다. Language Server 로그는 `snyk-ls.log`라는 파일 이름으로 동일한 디렉토리에서 찾을 수 있습니다.

터미널에서 작업할 때는 **폴더 경로에서 공백을 이스케이프하는 것에 주의하세요**.

## 프록시 설정

Snyk [Eclipse 문서](./)를 반드시 읽어보세요.

### 다룰 세부 사항

* 문제 발생 여부를 CLI 터미널에서 IDE 외부에서 확인하세요.
* 가능하면 최신 Snyk CLI 버전을 사용하세요.
* 디버그 옵션을 사용하여 사용자의 `snyk test`와 `snyk monitor` 출력 값을 얻으세요.
* 명령줄에서 프록시 변수를 설정하세요: `set http_proxy=<http….>`

### 프록시 설정 확인

Eclipse에서 **Windows → Preferences → General → Network Connections**로 이동합니다.

구성된 프록시 설정이 CLI 터미널에서 설정한 설정과 동일한지 확인하세요.

프록시 설정은 Snyk 환경 설정의 **Additional Environment**에서 추가하여 덮어쓸 수 있습니다:

`https_proxy=http://your-proxy.com:8080`

## Eclipse 플러그인 설치 문제

Eclipse 플러그인 설치에 문제가 발생하는 경우, 이는 Eclipse에서 사용하는 JDK 버전과의 호환성 문제일 수 있습니다. 이 문제는 Eclipse를 JDK 17 이상 버전으로 실행하여 해결할 수 있습니다. 다음 단계를 따르세요:

1. JDK 17 또는 더 높은 버전을 다운로드하여 설치합니다.
2. 공식 Eclipse 웹사이트에서 Eclipse IDE를 다운로드하여 설치합니다.
3. Eclipse IDE가 설치된 후, Eclipse 설치 디렉터리로 이동하여 `eclipse.ini` 파일을 찾습니다.
4.  `eclipse.ini` 파일을 열고 다음을 추가합니다:

    `vm {JDK 17 이상 버전의 경로}\bin`

    `{JDK 17 이상 버전의 경로}`를 JDK 17 이상 버전이 설치된 실제 경로로 교체합니다.
5. `eclipse.ini` 파일을 저장하고 닫습니다.
6. 일반적인 방법으로 Eclipse IDE를 실행합니다.
7. Eclipse가 실행된 후, **Window > Preferences > Java > Installed JREs**로 이동합니다.
8. **Add**를 클릭하고 시스템에 설치된 **JDK 17 이상 버전의 경로**를 선택합니다.
9. **OK**를 클릭하고 Preferences 대화상자를 닫습니다.

이제 원하는 Eclipse 플러그인을 설치할 준비가 되었습니다.

다음은 JDK 17 이상 버전의 경로를 추가한 후 `eclipse.ini` 파일이 어떻게 보일지에 대한 예시입니다. 마지막 줄이 관련된 변경 사항이며, 나머지 파일은 수정할 필요는 없습니다.

`-startup plugins/org.eclipse.equinox.launcher_1.6.200.v20210416-2027.jar --launcher.library plugins/org.eclipse.equinox.launcher.win32.win32.x86_64_1.2.200.v20210429-1609 -product org.eclipse.epp.package.jee.product -showsplash org.eclipse.epp.package.common --launcher.defaultAction openFile --launcher.appendVmargs -vmargs -Dosgi.requiredJavaVersion=11 -Xms256m -Xmx2048m --add-modules=ALL-SYSTEM --add-opens=java.base/java.util=ALL-UNNAMED --add-opens=java.base/java.lang=ALL-UNNAMED --add-opens=java.base/java.io=ALL-UNNAMED --add-opens=java.base/java.nio=ALL-UNNAMED --add-opens=java.base/sun.nio.ch=ALL-UNNAMED -vm C:\\Program Files\\Java\\jdk-17.0.1\\bin`

`C:\\\\Program Files\\\\Java\\\\jdk-17.0.1\\\\bin`을 시스템에 JDK 17 이상 버전이 설치된 실제 경로로 교체하세요.

## 애플리케이션 개발 및 JDK 버전 <a href="#application-development" id="application-development"></a>

Eclipse를 JDK 17 이상 버전으로 실행하는 것이 플러그인 설치와 관련된 문제를 해결할 수 있지만, 모든 Java 버전이 애플리케이션 개발에 호환되는 것은 아니므로 주의해야 합니다. 애플리케이션의 특정 요구 사항에 따라 JDK 8 또는 11을 사용해야 할 수 있습니다. Java 환경을 변경하기 전에 애플리케이션의 문서 및 요구 사항을 반드시 확인하세요.

또한, 여러 버전의 JDK를 시스템에 설치할 수 있지만 한 번에 하나의 버전만 사용할 수 있습니다. 따라서 프로젝트마다 다른 JDK 버전을 사용해야 하는 경우 Eclipse 설정을 적절히 업데이트해야 합니다.

특정 프로젝트에 대해 Eclipse에서 사용되는 JDK 버전을 변경하려면 다음 단계를 따르세요:

1. Eclipse에서 프로젝트를 엽니다.
2. 프로젝트를 마우스 오른쪽 버튼으로 클릭하고 **Properties**를 선택합니다.
3. **Java Build Path > Libraries**로 이동합니다.
4. **JRE System Library**를 찾고 확장합니다.
5. **Edit**를 클릭하고 설치된 JDK 버전 목록에서 원하는 JDK 버전을 선택합니다.
6. **Finish**를 클릭하여 변경 사항을 저장합니다.

이 단계를 따르면 Eclipse에서 서로 다른 프로젝트에 대해 JDK 버전을 쉽게 전환할 수 있습니다. 각 프로젝트에 대해 올바른 JDK 버전을 사용해야 애플리케이션의 호환성 및 기능에 영향을 미칠 수 있습니다.

## Windows Defender가 바이너리 실행 경고

Windows Defender는 때때로 Go 바이너리 실행을 차단하거나 경고를 발행할 수 있습니다. 이는 다양한 이유로 발생할 수 있으며, 예를 들어 바이너리가 인식되지 않거나 바이러스 백신 소프트웨어를 유발하는 동작을 가질 수 있습니다.

이 문제를 해결하려면 다음 해결 방법을 시도할 수 있습니다:

* 제외 추가: Go 바이너리 또는 바이너리가 포함된 디렉터리를 Windows Defender의 검사에서 제외시켜 false positive를 방지하고 바이너리가 방해 없이 실행되도록 할 수 있습니다.
* Windows Defender 일시 중지: 제외 추가로 문제가 해결되지 않으면 Go 바이너리를 실행할 때 Windows Defender를 일시적으로 비활성화할 수 있습니다. 그러나 바이러스 백신 소프트웨어를 비활성화할 때는 주의가 필요하며, 다른 보안 조치를 마련해야 합니다.
* Microsoft에 바이너리 제출: 이 감지가 false positive라고 생각되면 Microsoft에 보고할 수 있습니다. 그들은 파일을 검토하는 프로세스를 가지고 있으며, false positive로 확인되면 향후 바이러스 백신 정의에서 해당 감지가 업데이트될 수 있습니다.

## **JAR 서명 정보**

플러그인의 올바른 출처를 확인하려면, Eclipse 대화 상자에서 서명 세부 사항을 확인하여 이를 검증하세요.

<figure><img src="../../../.gitbook/assets/image (134) (2) (1) (1) (1) (1) (1) (1) (1).png" alt="The signing key details to verify the integrity and origin of the downloaded plugin"><figcaption><p>다운로드한 플러그인의 무결성과 출처를 확인하기 위한 서명 키 세부 정보</p></figcaption></figure>
