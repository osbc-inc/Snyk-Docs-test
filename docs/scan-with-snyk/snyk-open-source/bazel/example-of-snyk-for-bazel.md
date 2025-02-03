# Bazel을 위한 Snyk의 예

{% hint style="info" %}
전체 Bazel Java 프로젝트 및 해당 Snyk Dep 그래프 객체의 예제를 보려면 [Bazel Java 프로젝트에서 수동으로 Dep 그래프 생성하기](https://github.com/snyk/bazel-simple-app)를 참조하십시오.
{% endhint %}

Maven 패키지에 대한 단일 종속성이 있는 간단한 Bazel 프로젝트의 경우, 다음과 같이 종속성을 지정할 수 있습니다:

```
maven_jar(
    name = "logback-core",
    artifact = "ch.qos.logback:logback-core:1.0.13",
    sha1 = "dc6e6ce937347bd4d990fc89f4ceb469db53e45e",
)
```

이를 통해 다음과 같은 Dep 그래프 JSON 객체를 작성할 수 있습니다:

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

이 특정 패키지인 (`ch.qos.logback:logback-core@1.0.13`)은 결과 JSON 응답 객체에서 자세히 설명된 취약점이 포함되어 있습니다.
