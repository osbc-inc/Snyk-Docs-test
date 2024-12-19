# 모든 관리되지 않은 JAR 파일 스캔

[Snyk](https://snyk.io/) CLI는 [Java 애플리케이션](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/#open-source-and-licensing)에서 관리되지 않은 JAR 파일을 스캔하여 어떤 오픈 소스 패키지를 포함하고 있는지 식별합니다.

CLI는 패키지 이름, 버전 및 취약점을 식별하는데, 이는 패키지가 Maven Central에서 사용 가능하고 JAR 파일 해시가 Maven Central의 해시와 일치하는 경우에만 이루어집니다.

**전제 조건:** 관리되지 않은 JAR 파일을 스캔하려면 [지원되는 버전](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/#supported-frameworks-and-package-managers)의 Maven을 설치해야 합니다.

## 단일 폴더의 모든 JAR 파일 스캔 및 각 JAR 파일 개별적으로 스캔

개별적으로 각 JAR 파일을 스캔하려면 다음 명령을 사용합니다:
```
snyk test --scan-unmanaged --file=/path/to/file
```

각 JAR 파일을 개별적으로 테스트할 때, 스캔된 JAR 파일의 이름이 Snyk 웹 UI에 나타납니다.

**WAR 파일 지원**: Maven Central에 게시된 개별 WAR 파일들을 스캔할 수 있습니다. 오픈 소스 종속성 JAR를 직접 스캔하려면 다른 WAR 파일이나 기타 JAR 파일을 포함하는 모든 다른 WAR 파일을 추출(압축 해제)해야 합니다.

## 모든 하위 폴더들을 재귀적으로 스캔

{% hint style="warning" %}
**여기서 설명된 방법은 사용이 중단되었습니다.**

Snyk CLI 1.1176.0부터 기본적으로 `--scan-all-unmanaged` 옵션을 사용하여 모든 하위 폴더를 스캔할 수 있습니다. --`scan-all-unmanaged`를 사용하여 스캔하면 파일 이름 대신 패키지 이름이 나타납니다.
{% endhint %}

Java 앱은 종종 애플리케이션 내 다양한 폴더에 JAR 파일을 보유합니다.

다음은 현재 폴더부터 시작하여 모든 하위 폴더를 재귀적으로 거치면서 발견된 각 JAR 파일을 테스트하는 Snyk CLI 버전 1.1176.0 이전에 사용된 Linux/Mac Bash 스크립트입니다.

`REMOTE_REPO_URL` 변수에 값을 설정하는 것이 중요하며, 이는 `--remote-repo-url` 매개변수를 사용하여 UI에서 단일 Snyk 프로젝트 아래 모든 스캔 결과를 결합하는 데 사용됩니다.

```bash
#!/bin/bash

SNYK_CLI_BINARY_NAME=snyk-cli
SNYK_CLI_BINARY_LOCATION=https://github.com/snyk/cli/releases/latest/download/
REMOTE_REPO_URL= #UI에서 원하는 프로젝트명 입력

detected_jars=""
undetected_jars=""
detected_count=0
undetected_count=0

[[ -z "$REMOTE_REPO_URL" ]] && { echo "REMOTE_REPO_URL가 비어 있습니다. REMOTE_REPO_URL(6행)을 입력하고 스크립트를 다시 실행하십시오." ; exit 1; }

#OS에 특화된 Snyk 바이너리 다운로드 (MacOS 또는 Linux)
case "$(uname -s)" in
   Darwin)
     curl -L -O $SNYK_CLI_BINARY_LOCATION/snyk-macos
     mv snyk-macos snyk-cli
     ;;
   Linux)
     curl -L -O $SNYK_CLI_BINARY_LOCATION/snyk-linux
     mv snyk-linux snyk-cli
     ;;
esac

chmod +x $SNYK_CLI_BINARY_NAME

#모든 .jar 파일을 재귀적으로 찾아 루프 실행
#주의: 이름에 공백이 있는 파일이나 폴더에서 오류 발생 가능
for file in $(find . -type f -name '*.jar' | uniq)
do
echo ""
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-="    
echo $file
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=" 

#각 .jar에 대해 Snyk 모니터 실행
if (./$SNYK_CLI_BINARY_NAME monitor --scan-unmanaged --file=$file --project-name=$file --remote-repo-url=$REMOTE_REPO_URL) then
  detected_jars+=$file'\n'
  let detected_count++
else
  undetected_jars+=$file'\n'
  let undetected_count++
fi

done

#콘솔에 메트릭 출력
echo ""
echo ""
echo ""
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=" 
echo "찾은 JAR 파일 ($detected_count) - 전달된 종속성은 포함되지 않음:"
echo ""
if [ ${detected_count} -gt 0 ]; then
   printf "${detected_jars}"
fi
echo ""
echo ""
echo ""
echo "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=" 
echo "찾지 못한 JAR 파일 ($undetected_count) - Maven Central에 없음:"
echo ""
if [ ${undetected_count} -gt 0 ]; then
   printf "${undetected_jars}"
fi
```

다음은 모든 하위 폴더의 JAR를 스캔하는 Windows 배치 스크립트입니다. `scanjar.bat` 파일에서 실행합니다.

이 스크립트를 사용하려면 Snyk CLI를 설치해야 합니다.

```batch
REM 사용법:
REM 예: scanjar.bat "C:\workspace\app" "myapp"
SET WORKSPACE=%1
SET REMOTE_REPO_URL=%2
for /R %WORKSPACE% %%f in (*.jar) do cmd /c snyk monitor --scan-unmanaged --remote-repo-url=%REMOTE_REPO_URL% --file=%%f --project-name=%%f
```

"REMOTE_REPO_URL"이 "econnect"으로 설정된 스크립트를 사용한 후 Snyk UI에 표시된 결과의 예는 다음과 같습니다.

<figure><img src="../../.gitbook/assets/untitled (1) (1) (1).png" alt="관리되지 않은 JAR 파일 스캔 결과"><figcaption><p>관리되지 않은 JAR 파일 스캔 결과</p></figcaption></figure>