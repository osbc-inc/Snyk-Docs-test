# 이클립스 플러그인에 대한 문제 해결

{% hint style="warning" %}
Snyk 플러그인은 유통업체의 End Of Life (EOL)에 도달한 모든 운영 체제를 지원하지 않습니다.&#x20;
{% endhint %}

## 일반적인 문제 해결

일반적인 문제 해결 방법은 [IDE 문제 해결 페이지](../troubleshooting-ides/)를 참조하십시오.

## 로그

플러그인 로그의 저장 위치를 확인하려면 **Preferences** > **Language Servers** > **Logs**로 이동하고 Eclipse에서 **Snyk Language Server** 행을 찾으십시오. Language Server를 비활성화할 수 있으므로 로그를 검색하려면 활성화해야 할 수도 있습니다. 로그는 설정에 따라 콘솔이나 파일에 저장됩니다.

추가 플러그인 오류 로그를 보려면:

1. **창** > **보기 보기** > **기타...**로 이동하십시오.
2. **텍스트 필터**에 **오류 로그**를 검색하십시오.
3. **열기**를 클릭하여 오류 로그 탭을 확인하십시오. 탭 보기를 플러그인 별로 그룹화하면 (오른쪽 상단의 세 점 메뉴 > **그룹화 기준** > **플러그인**), `io.snyk.eclipse.plugin` 행을 통해 플러그인 오류를 확인할 수 있습니다.

로그 수준은 환경 변수 `SNYK_LOG_LEVEL`을 사용하여 Eclipse 설정에서 설정할 수 있으며, 예를 들어`SNYK_LOG_LEVEL=debug`로 설정할 수 있습니다.

Language Server의 설치 경로를 확인하고, 플러그인 설정에서 표시되는 경로를 찾으십시오. Language Server 로그는 같은 디렉토리에 `snyk-ls.log`라는 파일 이름으로 있습니다.

터미널에서 작업할 때는 **폴더 경로에 있는 공백을 주의 깊게 처리**해야 합니다.

## 프록시 설정

Snyk [Eclipse 문서](./)를 읽었는지 확인하십시오.

### &#x20;해결할 사항

* 문제가 IDE 외의 CLI 터미널에서 발생하는지 확인하십시오.
* 가능하다면, 최신 Snyk CLI 버전을 사용하십시오.
* 사용자의 `snyk test` 및 `snyk monitor` 출력을 가져오기 위해 디버그 옵션을 사용합니다.
* 명령줄에서 프록시 변수를 설정하십시오: `set http_proxy=<http….>`

### &#x20;프록시 설정 확인

Eclipse에서 **Windows → Preferences → General → Network Connections**로 이동하십시오.

구성된 프록시 설정이 CLI 터미널에서 설정된 것과 동일한지 확인하십시오.

프록시 설정은 Snyk 환경 추가 환경을 통해 설정을 추가하여 재정의할 수 있습니다:

`https_proxy=http://your-proxy.com:8080`

## 이클립스 플러그인 설치 문제

이클립스 플러그인 설치 중 문제가 발생하는 경우, 이는 이클립스가 사용하는 JDK 버전의 호환성 문제일 수 있습니다. 이는 JDK 17 이상 버전으로 이클립스를 실행함으로써 해결할 수 있습니다. 다음 단계를 따라 진행하십시오:

1. JDK 17 이상 버전을 다운로드하고 설치하십시오.
2. 공식 Eclipse 웹사이트에서 Eclipse IDE를 다운로드하고 설치하십시오.
3. Eclipse IDE가 설치된 후, Eclipse 설치 디렉터리로 이동하고 `eclipse.ini` 파일을 찾으십시오.
4.  `eclipse.ini` 파일을 열고 다음을 추가하십시오:

    `vm {JDK 17 이상 버전 경로}\bin`

    `{JDK 17 이상 버전 경로}`를 시스템에 설치된 실제 JDK 17 이상 버전 경로로 대체하십시오.
5. `eclipse.ini` 파일을 저장하고 닫으십시오.
6. 일반적인 방법으로 Eclipse IDE를 실행하십시오.
7. Eclipse를 실행한 후 **Window > Preferences > Java > Installed JREs**로 이동하십시오.
8. **추가**를 클릭하고 시스템에 설치된 **JDK 17 이상 버전 경로**를 선택하십시오.
9. **OK**를 클릭하고 환경 대화 상자를 닫으십시오.

이제 선택한 모든 Eclipse 플러그인을 설치할 준비가되었습니다.