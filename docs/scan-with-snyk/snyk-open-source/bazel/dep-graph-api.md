# Dep 그래프 API

{% hint style="info" %}
**기능 가용성**\
Snyk API는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는[플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Dep 그래프 API에는 추가 권한이 필요합니다. 액세스 권한을 요청하려면 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

[Bazel](./)에서 관리하는 종속성을 테스트하고 모니터링하는 경우, Snyk Dep 그래프 API 엔드포인트인 [테스트 Dep 그래프](../../../snyk-api/reference/test-v1.md#test-dep-graph) 및 [모니터 Dep 그래프](../../../snyk-api/reference/monitor-v1.md)를 사용하는 것이 권장됩니다. 모니터 기능을 사용하면 Snyk가 취약점을 모니터링할 트리를 제출할 수 있습니다. Bazel은 C++을 포함한 여러 언어에 사용될 수 있지만, **Dep 그래프 엔드포인트는 C++을 지원하지 않습니다**.

다음 기본 단계를 따르십시오:

1. Maven, Cocoapods 등 종속성 유형별로, [Dep 그래프 JSON 객체](https://github.com/snyk/dep-graph)를 생성하여 종속성 패키지와 버전을 나열하십시오. [Baszel을 위한 Snyk 예제](dep-graph-api.md#example-of-snyk-for-bazel)를 참조하십시오.
2.  Bazel 테스트 규칙의 일부로, Dep 그래프 JSON 객체를 [인증 토큰](../../../snyk-api/rest-api/authentication-for-api/)과 함께 엔드포인트 [테스트 Dep 그래프](../../../snyk-api/reference/test-v1.md#test-dep-graph)에 POST 요청으로 보내십시오. 다음은 예시 curl 요청입니다:

    ```
    curl -X POST 'https://api.snyk.io/v1/test/dep-graph' \
      -H 'Authorization: token ' \
      -H 'Content-Type: application/json; charset=utf-8' \
      -d @dep-graph.json
    ```
3. API 응답을 확인하여 통과/실패 상태 및 발생한 취약점을 확인하십시오.

### 테스트 Dep 그래프 API 동작 방식

테스트 Dep 그래프API는 일반 종속성 그래프를 취하고 해당 종속성에 대한 관련 취약점을 나타내는 보고서를 반환합니다.

지원되는 패키지 관리자 및 저장소 생태계는 [테스트 Dep 그래프](../../../snyk-api/reference/test-v1.md#test-dep-graph) 및 [모니터 Dep 그래프](../../../snyk-api/reference/monitor-v1.md) 문서에 나열되어 있습니다.

지원되는 생태계에서 사용 가능한 Bazel 종속성 중 어떤 것이든 Snyk API를 사용하여 테스트할 수 있습니다.

### Snyk Dep 그래프 JSON 구문

테스트 Dep 그래프 API는 루트 애플리케이션과 직접 및 전달 종속성의 그래프를 설명하는 [Snyk Dep 그래프](https://github.com/snyk/dep-graph) JSON 객체를 사용합니다.

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

Dep 그래프객체의 특정 구성 요소에 대한 추가 설명은 다음과 같습니다:

* `schemaVersion` - Dep 그래프 스키마의 버전. 이 값을 `1.2.0`으로 설정합니다.
* `pkgManager.name` - `deb` , `gomodules` , `gradle` , `maven` , `npm` , `nuget` , `paket` , `pip` , `rpm` , `rubygems`, 또는 `cocoapods` 중 하나일 수 있습니다.
* `pkgs` - Dep 그래프의 모든 패키지의 `id`, `name`, `version`을 포함하는 객체 배열입니다. `id`는 `name@version` 형식이어야 합니다. 이 배열에 종속성을 각각 나열하십시오.
* `graph.nodes` - `pkgs`의 항목 간 관계를 설명하는 객체 배열입니다. 이 배열은 일반적으로 Project 노드와 `deps.`의 직접 종속성으로 정의된 다른 모든 패키지를 평면 배열로 포함합니다.
* `graph.rootNodeId` - 그래프의 루트 노드로 사용할 `graph.nodes` 항목의 `id`를 지정합니다. 이를 Project 노드의 `nodeId`로 설정합니다.

### Snyk Dep 그래프  테스트 API 응답

테스트 Dep 그래프 API는 Dep 그래프프 종속성에서 발견된 문제(취약점 및 라이선스)을 나타내는 JSON 객체를 반환합니다.

단일 취약점이 포함된 예시 응답은 다음과 같습니다:

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

* `ok` - Snyk가 제공된 종속성에서 취약점을 발견했는지에 대한 요약된 부울 값입니다. 이 값은 빠른 통과 또는 실패 테스트에 사용할 수 있습니다.
* `issuesData` - 발견된 각 고유한 취약점의 해시입니다. 각 취약점에는 `title`, `description`, `identifiers`, `publicationTime`, `severity` 등 많은 유용한 속성이 포함됩니다.
* `issues` - `issuesData`의 취약점과 패키지 간의 매핑의 간단한 배열입니다. 취약점이 여러 패키지에 관련될 수 있으므로 이 매핑은 응답 길이를 가능한 한 짧게 유지하는 데 사용됩니다.
