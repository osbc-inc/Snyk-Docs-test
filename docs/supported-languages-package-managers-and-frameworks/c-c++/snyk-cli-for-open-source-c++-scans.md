# Snyk CLI를 사용한 오픈 소스 C++ 스캔

## C/C++를 위한 Snyk CLI

C/C++에 대한 취약점을 탐색하려면 [Snyk Vuln DB](https://security.snyk.io)를 검색하세요. Snyk는 이 데이터베이스를 기반으로 코드를 테스트하며, 이 데이터베이스는 온라인 소스에서 주기적으로 업데이트됩니다. 자세한 내용은 [Snyk 취약점 데이터베이스](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md)를 참조하세요.

Snyk가 오픈 소스 프로젝트를 스캔하려면 의존성이 스캔된 디렉토리 내의 소스 코드로 사용 가능해야 합니다. 의존성이 다른 위치에 있는 경우 해당 위치도 스캔해야 합니다. 자세한 내용은 [소스 코드 의존성은 스캔된 폴더에 있어야 합니다](snyk-cli-for-open-source-c++-scans.md#source-code-dependencies-must-be-in-the-scanned-folder)을 참조하세요.

[`snyk test --unmanaged`](../../snyk-cli/commands/test.md#unmanaged) 명령을 실행할 때, Snyk는 다음을 수행합니다:

1. 현재 폴더의 모든 파일을 해시 목록으로 변환합니다.
2. 해시를 Snyk 스캔 서버로 전송하여 의존성 목록을 계산합니다.
3. 데이터베이스를 쿼리하여 일치 가능성이 있는 의존성 목록을 찾습니다.
4. 이러한 의존성을 알려진 취약점에 연결합니다.
5. 결과를 표시합니다.

## 아카이브 스캔

기본적으로 아카이브는 스캔되지 않습니다. 그러나 CLI를 사용하여 아카이브를 재귀적으로 추출하여 내부의 소스 코드를 분석할 수 있습니다.

아카이브 추출을 활성화하려면 `--max-depth` 옵션을 사용하여 추출 깊이를 지정하세요.

지원되는 아카이브 형식은 다음과 같습니다:

- zip과 유사한 아카이브
- tar 아카이브
- gzip 압축 알고리즘이 적용된 tar 아카이브

## **소스 코드 의존성은 스캔된 폴더에 있어야 합니다**

CLI가 소스 코드에서 의존성을 찾을 수 있도록 하려면, 스캔된 폴더에 전체 의존성 소스 코드의 충분한 부분이 있어야 합니다.

원본(변경되지 않은) 형태의 파일이 크게 포함되어야 하며, 이것이 의존성을 정확하게 식별하고 올바른 취약점 세트를 보고하는 데 중요합니다. 해당 소스 코드를 수정하면 스캔 엔진의 신뢰도가 줄어들어 결과가 덜 정확해집니다. 다른 잠재적인 문제로는 의존성이 식별되지 않거나 다른 버전 또는 심지어 다른 패키지로 잘못 식별될 수 있습니다.

다음 예제는 일반적으로 의존성이 나열된 패키지를 보여줍니다:

```
c-example
├── deps
│   ├── curl-7.58.0
│   │   ├── include
│   │   │   ├── Makefile.am
│   │   │   ├── Makefile.in
│   │   │   ├── README
│   │   │   └── curl
│   │   ├── install-sh
│   │   ├── lib
│   │   │   ├── asyn.h
│   │   │   ├── base64.c
│   │   │   ├── checksrc.pl
│   │   │   ├── config-amigaos.h
│   │   │   ├── conncache.c
│   │   │   ├── conncache.h
│   ├── src
│   │   ├── tool_binmode.c
│   │   ├── tool_binmode.h
│   │   ├── tool_bname.c
│   │   ├── tool_xattr.c
...
```

## **릴리스 지원**

공식 릴리스만 추적됩니다. 패키지 관리자를 갖는 프로젝트의 경우, 패키지 관리자로 릴리스되어야 합니다.

Go 및 Unmanaged 스캔(C/C++)의 경우, GitHub 레포지토리에 공식 릴리스 또는 태그가 있어야 합니다.

## 스캔 중 데이터 수집

C++ 프로젝트를 스캔할 때 다음 데이터가 수집되어 문제 해결에 사용될 수 있습니다:

**스캔된 파일의 해시:** 모든 파일이 변경할 수 없는 해시 목록으로 변환됩니다.

**스캔된 파일의 상대 경로:** 스캔 중인 디렉토리에 대해 상대적인 파일 경로가 포함되어 정확한 식별 및 매칭이 가능해집니다.

예시:\
`./project-name/vendor/bzip2-1.0.6/blocksort.c`

## **의존성 표시**

의존성을 표시하려면 `--print-deps` 옵션을 사용하세요:

```bash
$ snyk test --unmanaged --print-deps

Testing /Users/user/src/foo...


Dependencies:

  https://curl.se|curl@7.29.0
  purl: pkg:generic/curl@7.29.0?download_url=https:%2F%2Fcurl.se%2Fdownload%2Farcheology%2Fcurl-7.29.0.tar.gz
  confidence: 1.000

  https://github.com|nih-at/libzip@1.8.0
  purl: pkg:generic/libzip@1.8.0?download_url=https:%2F%2Fgithub.com%2Fnih-at%2Flibzip%2Farchive%2Fv1.8.0.tar.gz
  confidence: 1.000

  https://github.com|madler/zlib@1.2.11
  purl: pkg:generic/zlib@1.2.11?download_url=https:%2F%2Fzlib.net%2Ffossils%2Fzlib-1.2.11.tar.gz
  confidence: 1.000
```

각 의존성 식별에 어떤 파일이 기여했는지 알아보려면 `--print-dep-paths` 옵션을 사용하세요:

```bash
$ snyk test --unmanaged --print-dep-paths

Testing /Users/user/src/foo...


Dependencies:

  https://curl.se|curl@7.29.0
  purl: pkg:generic/curl@7.29.0?download_url=https:%2F%2Fcurl.se%2Fdownload%2Farcheology%2Fcurl-7.29.0.tar.gz
  confidence: 1.000
  matching files:
    - curl-7.29.0/Android.mk
    - curl-7.29.0/CHANGES
    - curl-7.29.0/CMake/CMakeConfigurableFile.in
    ... 그리고 1766개의 파일

  https://github.com|nih-at/libzip@1.8.0
  purl: pkg:generic/libzip@1.8.0?download_url=https:%2F%2Fgithub.com%2Fnih-at%2Flibzip%2Farchive%2Fv1.8.0.tar.gz
  confidence: 1.000
  matching files:
    - libzip-1.8.0/API-CHANGES.md
    - libzip-1.8.0/AUTHORS
    - libzip-1.8.0/CMakeLists.txt
    ... 그리고 780개의 파일

  https://github.com|madler/zlib@1.2.11
  purl: pkg:generic/zlib@1.2.11?download_url=https:%2F%2Fzlib.net%2Ffossils%2Fzlib-1.2.11.tar.gz
  confidence: 1.000
  matching files:
    - zlib-1.2.11/CMakeLists.txt
    - zlib-1.2.11/ChangeLog
    - zlib-1.2.11/FAQ
    ... 그리고 249개의 파일
```

이 출력은 Snyk가 식별한 의존성과 해당 버전에 대한 확신 수준을 보여줍니다. Snyk가 식별한 의존성과 해당 버전에 대한 확신 수준을 보려면 `--print-deps` 또는 `--print-dep-paths` 옵션을 사용할 수 있습니다.

## **확신 수준 이해**

확신 수준은 Snyk가 실제 의존성 식별에 대해 얼마나 확신하는지를 보여줍니다. 해당 숫자는 `0`에서 `1` 사이일 수 있으며, 숫자가 클수록 식별이 더 정확합니다. 확신 수준이 `1`인 경우 소스 트리의 모든 파일이 Snyk 데이터베이스의 모든 예상 파일과 완전히 일치합니다.

`curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0 confidence: 0.993`

소프트웨어에서 사용하는 의존성의 소스 코드를 변경해야 할 수 있습니다. Snyk는 파일 시그니처를 사용하여 오픈 소스 라이브러리와 가장 유사한 항목을 찾습니다. 변경 사항은 실제 라이브러리 식별의 정확성을 낮출 수 있습니다.

## **JSON 출력**

JSON에서 기계 판독 가능한 출력을 얻으려면 `--json` 옵션을 사용하세요:

```
$ snyk test --unmanaged --json
[
  {
    "issues": [
      {
        "pkgName": "curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz",
        "pkgVersion": "7.58.0",
        "issueId": "CVE-2019-5481",
        "fixInfo": {
          "isPatchable": false,
          "isPinnable": false
        }
      }
    ],
    "issuesData": {
      "CVE-2019-5481": {
        "severity": "high",
        "CVSSv3": "",
        "originalSeverity": "high",
        "severityWithCritical": "high",
        "type": "vuln",
        "alternativeIds": [
          ""
        ],
        "creationTime": "2019-09-16T19:15:00.000Z",
        "disclosureTime": "2019-09-16T19:15:00.000Z",
        "modificationTime": "2020-10-20T22:15:00.000Z",
        "publicationTime": "2019-09-16T19:15:00.000Z",
        "credit": [
          ""
        ],
        "id": "CVE-2019-5481",
        "packageManager": "cpp",
        "packageName": "curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz",
        "language": "cpp",
        "fixedIn": [
          ""
        ],
        "patches": [],
        "exploit": "No Data",
        "functions": [
          ""
        ],
        "semver": {
          "vulnerable": [
            "7.58.0"
          ],
          "vulnerableHashes": [
            ""
          ],
          "vulnerableByDistro": {}
        },
        "references": [
          {
            "title": "https://curl.haxx.se/docs/CVE-2019-5481.html",
            "url": "https://curl.haxx.se/docs/CVE-2019-5481.html"
          },
        ],
        "internal": {},
        "identifiers": {
          "CVE": [
            "CVE-2019-5481"
          ],
          "CWE": [],
          "ALTERNATIVE": [
            ""
          ]
        },
        "title": "CVE-2019-5481",
        "description": "",
        "license": "",
        "proprietary": true,
        "nearestFixedInVersion": ""
      }
    },
    "fileSignaturesDetails": {
      "https://curl.se|curl@7.58.0": {
        "artifact": "curl",
        "version": "7.58.0",
        "author": "curl",
        "path": "curl-7.58.0",
        "id": "59d80da8ba341aaff828662700000000",
        "url": "https://curl.se/download/curl-7.58.0.tar.gz",
        "purl": "pkg:generic/curl@7.58.0?download_url=https:%2F%2Fcurl.se%2Fdownload%2Fcurl-7.58.0.tar.gz",
        "score": 1,
        "confidence": 1,
        "filePaths": [
          "deps/curl-7.58.0/CHANGES",
          "c-example/deps/curl-7.58.0/CMake/CMakeConfigurableFile.in",
          "c-example/deps/curl-7.58.0/CMake/CurlSymbolHiding.cmake"
        ],
        "confidence": 1
      }
    }
  }
]
```

## **커맨드 라인 옵션**

다음 `snyk` 커맨드 라인 옵션을 `snyk test --unmanaged` 및 `snyk monitor --unmanaged` 명령과 함께 사용할 수 있습니다:

`--org=<ORG_ID>`\
`--json`\
`--json-file-output=<OUTPUT_FILE_PATH>` (`snyk test`에서만 사용 가능)\
`--remote-repo-url=<URL>`\
`--severity-threshold=<low|medium|high|critical>` (`snyk test`에서만 사용 가능)\
`--max-depth`\
`--print-dep-paths`\
`--target-reference=<TARGET_REFERENCE>` (`snyk monitor`에서만 사용 가능)\
`--project-name=<c-project>` (`snyk monitor`에서만 사용 가능)

명령 라인 옵션에 대한 자세한 정보는 Snyk 도움말 문서를 참조하세요: [`snyk test --unmanaged`의 스캔 옵션](https://docs.snyk.io/snyk-cli/commands/test#options-for-scanning-using-unmanaged) 또는 [`snyk monitor --unmanaged`](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-scanning-using-unmanaged).

Snyk CLI로 테스트 결과(문제 및 의존성)를 가져오려면 `snyk monitor --unmanaged` 명령을 실행하세요:

```
$ snyk monitor --unmanaged
Monitoring /c-example (c-example)...

이 스냅샷은 https://app.snyk.io/org/example-org/project/8ac0e233-d0f9-403e-b422-5970e7a37443/history/5de4616d-3967-485f-bf21-bbbe91068029에서 탐색할 수 있습니다.

해당 의존성과 관련된 새로 공개된 문제에 대한 알림은 이메일로 전송됩니다.
```

이렇게 하면 의존성 및 취약점의 스냅샷이 생성되어 Snyk 웹 UI로 가져오며, 문제를 검토하고 보고서에 포함하는 것이 가능합니다.

관리되지 않는 의존성을 가진 프로젝트를 가져오면 프로젝트 페이지에 나열된 새 프로젝트가 생성됩니다:

<figure><img src="../../.gitbook