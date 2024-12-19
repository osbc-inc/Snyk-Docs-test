# Snyk CLI 설치 또는 업데이트

이 페이지에서 설명된 방법을 사용하여 [Snyk CLI](../)를 설치하거나 업데이트할 수 있습니다.

Snyk CLI를 설치한 후에는 [인증](../commands/auth.md)해야 하며, 그 후에 [시작](../getting-started-with-the-snyk-cli.md)해서 설치를 테스트하고 취약점을 수정할 수 있습니다.

{% hint style="info" %}
Snyk에서는 항상 CLI 설치를 최신 버전으로 유지하는 것을 권장합니다. 설치된 Snyk CLI의 버전을 확인하려면 `snyk --version`을 실행하면 됩니다.
{% endhint %}

IDE용 CLI를 설치하는 방법에 대한 자세한 정보는 [IDE 문서](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)를 참조하십시오. 또한 CI/CD 환경에 CLI를 설치할 수 있습니다. 자세한 내용은 [CI/CD 문서](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)를 참조하십시오.

## 독립 실행 파일로 설치

{% hint style="info" %}
이 방법을 사용할 때는 Snyk CLI를 수동으로 최신 상태로 유지해야 합니다.
{% endhint %}

원하는 플랫폼용 (macOS, Linux, Windows) Snyk CLI의 독립 실행 파일을 다운로드하려면 [GitHub 릴리스](https://github.com/snyk/snyk/releases)를 사용합니다.

Snyk는 또한 Snyk 콘텐츠 전송 네트워크(CDN)에서 이러한 독립 실행 파일을 제공합니다. 최신 `release.json` [파일](https://downloads.snyk.io/cli/stable/release.json)에서 다운로드 링크를 확인하십시오. 특정 버전 또는 플랫폼에 대한 예시는 다음과 같습니다:

* [https://downloads.snyk.io/cli/v1.666.0/release.json](https://downloads.snyk.io/cli/v1.666.0/release.json)
* [https://downloads.snyk.io/cli/stable/snyk-macos](https://downloads.snyk.io/cli/stable/snyk-macos)

예를 들어, macOS에서 최신 Snyk CLI를 다운로드하고 실행하려면 다음을 실행할 수 있습니다:

```bash
curl --compressed https://downloads.snyk.io/cli/stable/snyk-macos -o snyk
chmod +x ./snyk
mv ./snyk /usr/local/bin/
```

이러한 직접 다운로드 링크를 사용하여 실행 파일을 다운로드할 수도 있습니다:

| OS\Architecture | amd64                                                                | arm64                                                                       |
| --------------- | -------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| macOS           | [snyk-macos](https://downloads.snyk.io/cli/stable/snyk-macos)        | [snyk-macos-arm64](https://downloads.snyk.io/cli/stable/snyk-macos-arm64)   |
| Linux           | [snyk-linux](https://downloads.snyk.io/cli/stable/snyk-linux)        | [snyk-linux-arm64](https://downloads.snyk.io/cli/stable/snyk-linux-arm64)   |
| Alpine          | [snyk-alpine](https://downloads.snyk.io/cli/stable/snyk-alpine)      | [snyk-alpine-arm64](https://downloads.snyk.io/cli/stable/snyk-alpine-arm64) |
| Windows         | [snyk-win.exe](https://downloads.snyk.io/cli/v1.1294.0/snyk-win.exe) | -                                                                           |

예를 들어, snyk-win.exe 파일을 snyk.exe로 이름을 변경하여 설치하여 snyk test와 같이 명령을 실행할 수 있습니다.

{% hint style="info" %}
Apple M1 또는 M2 기기 (darwin/arm64)에서 1.1230.0 버전 이전의 CLI 릴리스를 사용하려면 [Apple M1 또는 M2 기기에서 1.1230.0 버전 이전의 CLI 릴리스 사용하기](using-cli-releases-before-version-1.1230.0-on-an-apple-m1-or-m2-machine.md)를 참조하십시오.

Alpine 리눅스에서 CLI를 사용하려면 [알파인 리눅스 운영 체제에서 CLI 및 Jenkins 플러그인의 사전 필요 조건](prerequisites-for-cli-and-jenkins-plugin-on-alpine-linux-operating-system.md)을 참조하십시오.

더 많은 정보는 [CLI 독립 실행 파일 확인](verifying-cli-standalone-binaries.md)을 참조하십시오.
{% endhint %}

## Homebrew로 설치하기 (macOS, Linux)

다음을 실행하여 [Homebrew](https://brew.sh)를 사용하여 [Snyk의 tap](https://github.com/snyk/homebrew-tap)로부터 Snyk CLI를 설치하십시오. 이 tap은 매일 최신 Snyk CLI 릴리스로 업데이트됩니다.

```bash
brew tap snyk/tap
brew install snyk
```

{% hint style="warning" %}
Apple M1 또는 M2 (darwin/arm64)를 사용하려면 다음을 참조하십시오 : [Apple M1 또는 M2 기기에서 1.1230.0 버전 이전의 CLI 릴리스 사용하기](using-cli-releases-before-version-1.1230.0-on-an-apple-m1-or-m2-machine.md).
{% endhint %}

## Scoop를 사용하여 설치하기 (Windows)

다음을 실행하여 [Scoop](https://scoop.sh)를 사용하여 [Snyk's bucket](https://github.com/snyk/scoop-snyk)로부터 Snyk CLI를 설치하십시오. 이 bucket은 매일 최신 Snyk CLI 릴리스로 업데이트됩니다.

```
scoop bucket add snyk https://github.com/snyk/scoop-snyk
scoop install snyk
```

## npm 또는 Yarn으로 Snyk CLI 설치

npm를 사용하여 Snyk CLI를 설치하기 전에 다음 **전제 조건**이 충족되어 있는지 확인하십시오:

* Node.js 버전 12 이상을 사용하여 로컬 환경에 npm의 최신 버전을 설치하십시오. Node.js를 업데이트하는 방법에 대한 정보는 [Snyk CLI에 필요한 Node.js 버전 설치 또는 업그레이드](install-or-upgrade-to-version-of-node.js-required-for-snyk-cli.md)를 참조하십시오.
* Alpine Linux에서 Snyk를 실행하려면 먼저 libstdc++를 설치하십시오.\
  자세한 내용은 [알파인 리눅스 운영 체제에서 CLI 및 Jenkins 플러그인의 사전 필요 조건](prerequisites-for-cli-and-jenkins-plugin-on-alpine-linux-operating-system.md)을 참조하십시오.

그런 다음 다음 **단계를 따라 npm 또는 Yarn으로 설치**하십시오:

[Snyk CLI는 npm 패키지로 제공](https://www.npmjs.com/package/snyk)됩니다. 로컬에 Node.js가 설치되어 있다면 `npm install snyk -g` 명령을 실행하여 npm 패키지를 **설치**할 수 있습니다.

Yarn을 사용하는 경우 `yarn global add snyk` 명령을 실행하여 **설치**할 수 있습니다.

추가 정보는 [npm을 통해 Snyk CLI를 바이너리로 설치](installing-snyk-cli-as-a-binary-via-npm.md)를 참조하십시오.

## Docker 이미지에서 Snyk CLI

Snyk CLI는 Docker 이미지에서 실행할 수도 있습니다. Snyk는 [Docker Hub의 snyk/snyk](https://hub.docker.com/r/snyk/snyk)에서 여러 Docker 이미지를 제공합니다. 상세 내용은 [GitHub의 snyk/snyk-images](https://github.com/snyk/snyk-images)를 참조하십시오.

이러한 이미지는 Snyk CLI를 둘러싸고 있으며 Tag에 따라 다른 프로젝트에 대한 관련 도구가 제공됩니다. `snyk/snyk`를 사용하여 Gradle 프로젝트를 스캔하는 예시는 다음과 같습니다:

```bash
docker run -it \
    -e "SNYK_TOKEN=<TOKEN>" \
    -v "<PROJECT_DIRECTORY>:/project" \
    -v "/home/user/.gradle:/home/node/.gradle" \
  snyk/snyk:gradle:6.4 test --org=my-org-name
```

`Snyk/snyk`를 사용하여 Maven 프로젝트를 스캔하는 예시는 다음과 같습니다:

```
docker run --rm \
-e SNYK_TOKEN=<YOUR_SNYK_TOKEN> \
-v <PROJECT_DIRECTORY>:/app \
-v <PROJECT_DIRECTORY>/settings.xml:/root/.m2/settings.xml \
snyk/snyk:maven snyk monitor \
--all-projects=true \
--maven-aggregate-project
```