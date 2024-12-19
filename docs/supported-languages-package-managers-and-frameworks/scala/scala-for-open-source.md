# 오픈 소스용 Scala

## Scala의 오픈 소스를 위한 Snyk 지원

지원되는 패키지 관리자 및 기능에 대해 자세히 알아보려면 [Scala 세부 정보](./)를 참조하세요.

도움이 필요한 경우, [Snyk 지원팀에 문의하십시오](https://support.snyk.io).

## 오픈 소스 및 라이선스

다음은 sbt에 대한 Snyk 지원을 요약한 것입니다.

| 패키지 관리자 / 기능              | CLI 지원 | Git 지원 | 라이센스 스캔 | PR 수정 |
| --------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [sbt](https://www.scala-sbt.org/) | ✔︎          | ✔︎          | ✔︎               |         |

### Scala의 오픈 소스를 위한 CLI 지원

[Snyk CLI](../../snyk-cli/)는 [`sbt-dependency-graph`](https://github.com/sbt/sbt-dependency-graph) 플러그인을 사용하며, 이 플러그인은 `sbt` 1.4 이후부터 내장 플러그인으로 [포함](https://www.scala-sbt.org/1.x/docs/Combined+Pages.html#sbt-dependency-graph+is+in-sourced)되어 있습니다.

그러나 sbt 1.4+에서 권장하는 플러그인 호출 방법은 Snyk와 호환되지 않습니다. 대신, `addSbtPlugin()`을 사용하십시오. Snyk는 `sbt-dependency-graph`를 [글로벌 플러그인](https://www.scala-sbt.org/1.x/docs/Using-Plugins.html#Global+plugins)으로 설치하여 모든 `sbt` 프로젝트에서 사용할 수 있도록 권장합니다.

이를 위해 `sbt` 0.13의 경우 `~/.sbt/0.13/plugins/plugins.sbt`에 플러그인 종속성을 추가하거나, `sbt` 1.0+의 경우 `~/.sbt/1.0/plugins/plugins.sbt`에 추가하십시오.

단일 프로젝트에 플러그인을 추가하려면 해당 프로젝트의 `project/plugins.sbt`를 업데이트하십시오.

어떤 `sbt` 버전을 사용하든, 관련 `plugins.sbt` 파일에서 다음 명령을 사용해야 합니다:

`addSbtPlugin("net.virtual-void" % "sbt-dependency-graph" % "0.10.0-RC1")`

{% hint style="warning" %}
sbt 1.4+를 위해 `sbt-dependency-graph` 플러그인 문서에서 권장하는 `addDependencyTreePlugin` 명령은 Snyk CLI와 호환되지 않습니다. 위에 주어진대로 `addSbtPlugin()` 명령을 사용하십시오.
{% endhint %}

Snyk CLI와 함께 `sbt-dependency-graph`를 설치하는 자세한 정보는 [Snyk CLI로 Scala 프로젝트를 테스트하기 위해 SBT 의존성 그래프 플러그인을 설치하는 방법](https://support.snyk.io/s/article/How-to-install-the-SBT-dependency-graph-plugin-to-test-Scala-projects-with-Snyk-CLI) 문서를 참조하십시오.

### Scala의 오픈 소스를 위한 Git 리포지토리 통합 지원

Scala `sbt` 프로젝트는 Snyk이 [지원하는](../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/) Git 리포지토리에서 가져올 수 있습니다.

`build.sbt` 파일을 분석하여 `sbt`를 사용하는 Scala 프로젝트를 테스트하는 데에 Snyk이 사용됩니다.\
이를 올바르게 작동하려면 프로젝트를 가져오기 전에 리포지토리에이 파일이 있어야 합니다.

예를 들어 `Dependencies.scala`와 같이SCM(Software Configuration Management) 통합을 사용하여 Snyk이 접근할 수 없는 파일에 종속성의 버전을 선언할 수 없습니다.

* SCM 통합을 사용하여 프로젝트를 가져올 때 Scala 종속성이 감지되도록 하려면 `build.sbt`의 종속성을 Snyk이 감지할 수 있는 형식으로 선언해야 합니다, 예를 들어:\
  `"commons-io" % "commons-io" % "2.11.0"`.
* `build.sbt` 파일에 변수가 있을 경우 변수에 선언된 버전을 사용할 수 있습니다. 예를 들어:

```scala
  lazy val derbyVersion = "10.4.1.3"
  libraryDependencies ++= Seq(
    "org.apache.derby" % "derby" % derbyVersion
  ) 
```

더 많은 정보는 [환경 간 오픈 소스 취약점 수에 대한 차이점](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/differences-in-open-source-vulnerability-counts-across-environments.md)을 참조하십시오.