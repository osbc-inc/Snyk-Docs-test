# Bazel - 빌드 및 테스트 도구

{% hint style="info" %}
**기능 이용 가능성**\
Snyk는 {{Snyk 오픈 소스}}에서 지원하는 Bazel을 위한 빌드 및 테스트 도구입니다. 이 설명서의 지침은 Bazel v 7에만 해당됩니다.&#x20;
{% endhint %}

Bazel은 [bazel.build](https://docs.bazel.build/versions/master/bazel-overview.html)에서 다음과 같이 정의됩니다.

> _Bazel은 Make, Maven 및 Gradle과 유사한 오픈 소스 빌드 및 테스트 도구입니다. 사람이 읽을 수 있고 고수준의 빌드 언어를 사용합니다. Bazel은 여러 언어로 구성된 프로젝트와 다양한 플랫폼에 대한 출력물을 빌드합니다. Bazel은 여러 리포지토리에서 대규모 코드베이스 및 많은 사용자를 지원합니다._

## 적용 가능성

Snyk는 Bazel을 **Snyk 오픈 소스**에 대해서만 지원합니다.

Snyk는 Bazel로 관리되는 종속성을 가진 프로젝트를 테스트 지원합니다. 권장되는 방법은 [Dep Graph API](dep-graph-api.md)를 사용하여 테스트하고 모니터링하는 것입니다. C++를 비롯한 여러 언어에 Bazel을 사용할 수 있지만, **Dep Graph 엔드포인트는 C++을 지원하지 않습니다**.

## 패키지 관리자

Bazel은 npm 등의 패키지 관리자와 같은 종속성 매니페스트 파일이나 록 파일을 갖고 있지 않습니다. 대신, 빌드 구성은 [BUILD](https://docs.bazel.build/versions/master/build-ref.html#BUILD_files) 파일에서 관리되며, Python3을 기반으로 하는 도메인 특화 언어인 [Starlark](https://docs.bazel.build/versions/master/skylark/language.html)를 사용합니다.

Bazel은 npmjs.org나 Maven Central과 같은 패키지 레지스트리와의 네이티브 통합이 제한적입니다. 외부 레지스트리에서 종속성을 설치하는 데 도움이 되는 일부 Bazel 규칙을 추가할 수 있습니다. 예를 들어, [Maven에서](https://docs.bazel.build/versions/master/external.html#maven-artifacts-and-repositories) 가져오기.

Bazel 종속성은 BUILD 파일의 코드로 지정되므로 Snyk는 프로젝트에서 종속성을 쉽게 발견할 수 없습니다. Snyk를 사용하여 Bazel 프로젝트를 테스트하고 모니터링하는 정보는 [Dep Graph API](dep-graph-api.md) 페이지를 참조하십시오.

## 프레임워크 및 라이브러리

Snyk를 위한 Bazel의 사용 가능한 프레임워크 및 라이브러리가 없습니다.

## 기능

Snyk를 위한 Bazel의 사용 가능한 기능이 없습니다.

## Dep Graph API

{% hint style="info" %}
**기능 이용 가능성**\
Snyk API는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Dep Graph API에는 추가 권한이 필요합니다. [Snyk 지원팀에 문의](https://support.snyk.io)하여 액세스를 요청하십시오.

[Bazel](./)에서 관리되는 종속성을 테스트하고 모니터링하기 위해 Snyk Dep Graph API 엔드포인트 [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph) 및 [Monitor Dep Graph](../../../snyk-api/reference/monitor-v1.md)를 사용하는 것이 좋습니다. 모니터 기능을 통해 고객은 취약점을 모니터링할 트리를 Snyk에 제출할 수 있습니다. C++를 비롯한 여러 언어에 Bazel을 사용할 수 있지만, **Dep Graph 엔드포인트는 C++을 지원하지 않습니다**.

다음과 같은 기본적인 단계를 따릅니다:

1. 예를 들어 Maven, Cocoapods와 같은 각 종속성 유형에 대해 [Dep Graph JSON 객체](https://github.com/snyk/dep-graph)를 만들어 모든 종속성 패키지와 버전을 나열합니다. [Snyk를 위한 Baszel 예제](example-of-snyk-for-bazel.md)를 참조하십시오.
   
2. Bazel 테스트 규칙의 일부로 Dep Graph JSON 객체를 `POST` 요청으로 [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph) 엔드포인트에 보냅니다. 아래는 예시 curl 요청입니다:

    ```
    curl -X POST 'https://api.snyk.io/v1/test/dep-graph' \
      -H 'Authorization: token {{your token}}' \
      -H 'Content-Type: application/json; charset=utf-8' \
      -d @dep-graph.json
    ```

3. API 응답을 확인하여 성공/실패 상태 및 발생한 취약점을 확인합니다.

### Test Dep Graph API 작동 방식

Test Dep Graph API는 일반적인 종속성 그래프를 가져와 해당 종속성에 대한 관련 취약점을 포함하는 보고서를 반환합니다.

지원되는 패키지 관리자 및 저장소 생태계는 [Test Dep Graph](../../../snyk-api/reference/test-v1.md#test-dep-graph) 및 [Monitor Dep Graph](../../../snyk-api/reference/monitor-v1.md) 문서에 나열되어 있습니다.

지원되는 생태계 내에서 Bazel 종속성이 Snyk API를 사용하여 테스트될 수 있습니다.

### Snyk Dep Graph JSON 구문

Test Dep Graph API는 루트 응용 프로그램과 직접 및 간접 종속성의 그래프를 설명하는 [Snyk Dep Graph](https://github.com/snyk/dep-graph) JSON 객체를 가져옵니다.

이 형식에 대한 [스키마](https://github.com/snyk/dep-graph#depgraphdata)는 다음과 같습니다:

{% code overflow="wrap" fullWidth="false" %}
```java
export interface DepGraphData {
  schemaVersion: string;
  pkgManager: {
    name: string;
    version?: string;
    repositories?: Array<{
      alias: string;
    }>;
  };
  pkgs: Array<{
    id: string;
    info: {
      name: string;
      version?: string;
    };
  }>;
  graph: {
    rootNodeId: string;
    nodes: Array<{
      nodeId: string;
      pkgId: string;
      info?: {
        versionProvenance?: {
          type: string;
          location: string;
          property?: {
            name: string;
          };
        },
        labels?: {
          [key: string]: string | undefined;
        };
      };
      deps: Array<{
        nodeId: string;
      }>;
    }>;
  };
}
```
{% endcode %}

Dep Graph 객체의 특정 구성 요소에 대한 추가 설명은 다음과 같습니다:

* `schemaVersion` - Dep Graph 스키마의 버전입니다. 이를 `1.2.0`로 설정합니다.
* `pkgManager.name` - `deb`, `gomodules`, `gradle`, `maven`, `npm`, `nuget`, `paket`, `pip`, `rpm`, `rubygems`, 또는 `cocoapods` 중 하나일 수 있습니다.
* `pkgs` - Dep Graph의 모든 패키지의 `id`, `name`, `version`을 포함하는 객체 배열입니다. `id`는 `name@version` 형식이어야 합니다. 이 배열에 각 종속성을 나열합니다. 프로젝트 자체를 나타내는 항목도 포함합니다.
* `graph.nodes` - `pkgs`에 있는 항목 간 관계를 설명하는 객체 배열입니다. 이것은 일반적으로 프로젝트 노드로, `deps`에 정의된 직접 종속성의 평면 배열로 모든 다른 패키지가 정의됩니다.
* `graph.rootNodeId` - 그래프의 루트 노드로 사용할 `graph.nodes`의 항목의 `id`를 지정합니다. 이를 프로젝트 노드의 `nodeId`로 설정합니다.

### Snyk Dep Graph 테스트 API 응답

Test Dep Graph API는 Dep Graph 종속성에서 발견된 문제(취약점 및 라이선스)를 설명하는 JSON 객체를 반환합니다.

단일 취약점을 포함한 예시 응답은 다음과 같습니다:

{% code overflow="wrap" %}
```java
{
    "ok": false,
    "packageManager": "maven",
    "issuesData": {
        "SNYK-JAVA-CHQOSLOGBACK-30208": {
            "CVSSv3": "CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
            "alternativeIds": [],
            "creationTime": "2017-03-19T14:58:38Z",
            "credit": [
                "Unknown"
            ],
            "cvssScore": 9.8,
            "description": "## Overview\n[ch.qos.logback:logback-core](https://mvnrepository.com/artifact/ch.qos.logback/logback-core) is a logback-core module.\n\nAffected versions of this package are vulnerable to Arbitrary Code Execution. A configuration can be ...",
            "disclosureTime": "2017-03-13T06:59:00Z",
            "exploit": "Not Defined",
            "fixedIn": [
                "1.1.11"
            ],
            "functions": [],
            "id": "SNYK-JAVA-CHQOSLOGBACK-30208",
            "identifiers": {
                "CVE": [
                    "CVE-2017-5929"
                ],
                "CWE": [
                    "CWE-502"
                ]
            },
            "language": "java",
            "mavenModuleName": {
                "artifactId": "logback-core",
                "groupId": "ch.qos.logback"
            },
            "modificationTime": "2020-06-12T14:36:56.271247Z",
            "moduleName": "ch.qos.logback:logback-core",
            "packageManager": "maven",
            "packageName": "ch.qos.logback:logback-core",
            "patches": [],
            "proprietary": false,
            "publicationTime": "2017-03-21T15:30:44Z",
            "references": [
                {
                    "title": "GitHub Commit #1",
                    "url": "https://github.com/qos-ch/logback/commit/f46044b805bca91efe5fd6afe52257cd02f775f8"
                },
                {
                    "title": "GitHub Commit #2",
                    "url": "https://github.com/qos-ch/logback/commit/979b042cb1f0b4c1e5869ccc8912e68c39f769f9"
                },
                {
                    "title": "Logback News",
                    "url": "https://logback.qos.ch/news.html"
                },
                {
                    "title": "NVD",
                    "url": "https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2017-5929"
                },
                {
                    "title": "NVD",
                    "url": "https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2017-5929/"
                }
            ],
            "semver": {
                "vulnerable": [
                    "[, 1.1.11)"
                ]
            },
            "severity": "high",
            "title": "Arbitrary Code Execution"
        }
    },
    "issues": [
        {
            "pkgName": "ch.qos.logback:logback-core",
            "pkgVersion": "1.0.13",
            "issueId": "SNYK-JAVA-CHQOSLOGBACK-30208",
            "fixInfo": {}
        }
    ],
    "org": {
        "id": "3e5fe3fe-9181-4f0f-a231-39764485e73f",
        "name": "stephen.elson-xnf"
    }
}
```
{% endcode %}

응답 객체의 특정 구성 요소에 대한 추가 설명은 다음과 같습니다:

* `ok` - Snyk가 제공된 종속성에서 어떠한 취약점도 찾지 못했는지를 요약하는 부울 값입니다. 이를 사용하여 간편한 통과 또는 실패 테스트에 사용할 수 있습니다.
* `issuesData` - 찾은 각 고유 취약점에 대한 해시입니다. 각 취약점에는 `title`, `description`, `identifiers`, `publicationTime`, `severity` 등 많은 유용한 속성이 포함됩니다.
* `issues` - `issuesData`의 취약점에서 패키지로의 매핑에 대한 간단한 배열입니다. 취약점이 여러 패키지에 관련될 수 있으므로 이러한 매핑은 가능한 한 응답 길이를 짧게 유지하는 데 사용됩니다.

## Snyk를 위한 Bazel 예제

{% hint style="info" %}
Bazel Java 프로젝트 및 해당 Snyk Dep Graph 객체의 전체 예제인 [Bazel Java 프로젝트에서 Dep Graph 수동 생성](https://github.com/snyk/bazel-simple-app)를 참조하십시오.
{% endhint %}

Maven 패키지에 대한 단일 종속성이 있는 간단한 Bazel 프로젝트의 경우, 종속성을 다음과 같이 지정할 수 있습니다:

```
maven_jar(
    name = "logback-core",
    artifact = "ch.qos.logback:logback-core:1.0.13",
    sha1 = "dc6e6ce937347bd4d990fc89f4ceb469db53e45e",
)
```

이를 기반으로 다음과 같은 Dep Graph JSON 객체를 작성할 수 있습니다:

```
{
  "depGraph": {
    "schemaVersion": "1.2.0",
    "pkgManager": {
      "name": "maven"
    },
    "pkgs": [
      {
        "id": "app@1.0.0",
        "info": {
          "name": "app",
          "version": "1.0.0"
        }
      },
      {
        "id": "ch.qos.logback:logback-core@1.0.13",
        "info": {
          "name": "ch.qos.logback:logback-core",
          "version": "1.0.13"
        }
      }
    ],
    "graph": {
      "rootNodeId": "root-node",
      "nodes": [
        {
          "nodeId": "root-node",
          "pkgId": "app@1.0.0",
          "deps": [
            {
              "nodeId": "ch.qos.logback:logback-core@1.0.13"
            }
          ]
        },
        {
          "nodeId": "ch.qos.logback:logback-core@1.0.13",
          "pkgId": "ch.qos.logback:logback-core@1.0.13",
          "deps": []
        }
      ]
    }
  }
}
```

특정 패키지(`ch.qos.logback:logback-core@1.0.13`)에는 JSON 응답 객체에 자세히 설명된 취약점이 포함되어 있습니다.