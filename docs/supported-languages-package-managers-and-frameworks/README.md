# 지원되는 언어, 패키지 관리자 및 프레임워크

Snyk는 다양한 언어를 지원하며, 사용 중인 Snyk 제품에 따라 사용자 정의가 가능합니다. 이 페이지들은  및 에 초점을 맞춥니다.

의 언어 지원 정보는 [지원되는 작업 부하, 컨테이너 레지스트리, 언어 및 운영 체제](../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/supported-workloads-container-registries-languages-and-operating-systems.md) 및 [에서 지원하는 운영 체제 배포](../scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md)를 참조하세요.

IaC 언어 지원에 대한 정보는 [지원되는 IaC 언어, 클라우드 제공업체 및 클라우드 리소스](../scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/)를 참조하세요.

## 지원되는 언어

다음 표는 지원되는 언어 및 각 언어를 SCM 통합 및 Snyk CLI, IDE 및 CI/CD를 사용하는 데 지원이 가능한지를 나타냅니다. 자세한 내용은 각 언어 페이지로 이동하세요.

<table><thead><tr><th width="270">언어</th><th width="225"></th><th width="210"></th><th data-hidden>SCM 지원</th><th data-hidden>Snyk CLI, IDE, CI/CD</th></tr></thead><tbody><tr><td><a href="apex.md">Apex</a></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="c-c++/">C/C++</a></td><td>✔️</td><td>✔️</td><td>For </td><td>✔️</td></tr><tr><td><a href="dart-and-flutter.md">Dart 및 Flutter</a></td><td>✔️</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td></tr><tr><td><a href="elixir.md">Elixir</a></td><td>✔️</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>✔️</td></tr><tr><td><a href="go/">Go</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="java-and-kotlin/">Java 및 Kotlin</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="javascript/">JavaScript</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href=".net/">.NET</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="php/">PHP</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="python/">Python</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="ruby/">Ruby</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="rust.md">Rust</a></td><td>✔️</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td></tr><tr><td><a href="scala/">Scala</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="swift-and-objective-c/">Swift 및 Objective-C</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="typescript.md">TypeScript</a></td><td>✔️</td><td>✔️</td><td>✔️</td><td>✔️</td></tr><tr><td><a href="vb.net.md">VB NET</a></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>✔️</td><td>✔️</td><td>✔️</td></tr></tbody></table>

## 패키지 관리자 및 프레임워크

<table data-full-width="false"><thead><tr><th width="198">언어</th><th width="276"></th><th></th><th data-hidden>프레임워크</th><th data-hidden>지원되는 기능</th></tr></thead><tbody><tr><td><a href="apex.md">Apex</a></td><td> N/A</td><td><p><strong>패키지 관리자</strong>: N/A</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>.cls</code></li><li><code>.trigger</code> </li><li><code>.tgr</code></li></ul></td><td><a href="apex.md#frameworks-and-libraries">Apex</a> 페이지에서 전체 목록 확인</td><td><ul><li>Interfile 분석 지원</li><li>사용자 정의 규칙</li><li>보고서</li><li>Interfile 분석</li></ul></td></tr><tr><td><a href="c-c++/">C/C++</a></td><td><p><strong>패키지 관리자</strong>: N/A<br><strong>지원되는 파일</strong>: N/A</p><p></p></td><td><p><strong>패키지 관리자</strong>: N/A<br><strong>지원되는 파일</strong>:</p><ul><li><code>.c</code> </li><li><code>.cc</code> </li><li><code>.cpp</code> </li><li><code>.cxx</code> </li><li><code>.</code>h </li><li><code>.hpp</code> </li><li><code>.hxx</code></li></ul></td><td><a href="c-c++/#frameworks-and-libraries">C/C++</a> 페이지에서 전체 목록 확인</td><td><ul><li>라이선스 스캐닝 ()</li><li>보고서 ()</li><li>Interfile 분석 ()</li></ul></td></tr><tr><td><a href="dart-and-flutter.md">Dart 및 Flutter</a></td><td><strong>패키지 관리자</strong>: Pub<br><strong>지원되는 파일</strong>: N/A</td><td> N/A</td><td>N/A</td><td>N/A</td></tr><tr><td><a href="elixir.md">Elixir</a></td><td><strong>패키지 관리자</strong>: Mix/Hex<br><strong>지원되는 파일</strong>: N/A</td><td> N/A</td><td>N/A</td><td><ul><li>보고서</li></ul></td></tr><tr><td><a href="go/">Go</a></td><td><p><strong>패키지 관리자</strong>: Go Modules, dep<br><strong>지원되는 파일</strong>:</p><ul><li><code>go.mod</code> </li><li><code>gopkg.lock</code></li></ul></td><td><p><strong>패키지 관리자</strong>: Go Modules, dep</p><p><strong>지원되는 파일</strong>: <code>.go</code></p></td><td><a href="go/#frameworks-and-libraries">Go</a> 페이지에서 전체 목록 확인</td><td><ul><li>라이선스 스캐닝</li><li>보고서</li><li>사용자 정의 규칙</li><li>Interfile 분석</li></ul></td></tr><tr><td><a href="java-and-kotlin/">Java 및 Kotlin</a></td><td><p><strong>패키지 관리자</strong>: Maven, Gradle, Bazel<br></p><p><strong>지원되는 파일</strong>:</p><ul><li><code>pom.xml</code> </li><li><code>build.gradle</code></li><li><code>build.gradle.kts</code></li></ul></td><td><p><strong>패키지 관리자</strong>: Maven, Gradle, Bazel</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>.java</code> </li><li><code>.jsp</code> </li><li><code>.jspx</code> </li><li><code>.kt</code></li></ul></td><td><a href="java-and-kotlin/#frameworks-and-libraries">Java 및 Kotlin</a> 페이지에서 전체 목록 확인</td><td><ul><li>도구 PR (Maven) </li><li>라이선스 스캐닝 </li><li>보고서</li><li>사용자 정의 규칙 </li><li>Interfile 분석</li></ul></td></tr><tr><td><a href="javascript/">JavaScript</a></td><td><p><strong>패키지 관리자</strong>: npm, pnpm, Yarn</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>package.json</code> </li><li><code>package-lock.json</code> </li><li><code>pnpm-lock.yaml</code> </li><li><code>yarn.lock</code></li></ul></td><td><p><strong>패키지 관리자</strong>: npm, pnpm, Yarn</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>.ejs</code> </li><li><code>.es</code> </li><li><code>.es6</code> </li><li><code>.htm</code> </li><li><code>.html</code> </li><li><code>.js</code> </li><li><code>.jsx</code> </li><li><code>.ts</code> </li><li><code>.cts</code> </li><li><code>.mts</code> </li><li><code>.tsx</code> </li><li><code>.vue</code> </li><li><code>.mjs</code> </li><li><code>.cjs</code></li></ul></td><td><a href="javascript/#frameworks-and-libraries">JavaScript</a> 페이지에서 전체 목록 확인</td><td><ul><li>보고서 </li><li>사용자 정의 규칙 </li><li>Interfile 분석</li><li>도구 PRs </li><li>라이선스 스캐닝</li></ul></td></tr><tr><td><a href=".net/">.NET</a></td><td><p><strong>패키지 관리자</strong>: NuGet, Paket</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>project.assets.json</code> </li><li><code>*.sln</code> </li><li><code>packages.config</code> </li><li><code>project.json</code> </li><li><code>paket.dependencies</code></li><li> <code>paket.lock</code></li></ul></td><td><p></p><p><strong>패키지 관리자</strong>: NuGet, Paket<br><strong>지원되는 파일</strong>: N/A</p></td><td><a href=".net/#frameworks-and-libraries">.NET</a> 페이지에서 전체 목록 확인</td><td><ul><li>도구 PRs (NuGet) </li><li>라이선스 스캐닝 </li><li>보고서</li><li>사용자 정의 규칙 </li><li>Interfile 분석</li></ul></td></tr><tr><td><a href="php/">PHP</a></td><td><p><strong>패키지 관리자</strong>: Composer</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>composer.json</code>, <code>composer.lock</code></li></ul></td><td><p><strong>패키지 관리자</strong>: Composer</p><p><strong>지원되는 파일</strong>:</p><ul><li><code>.php</code> </li><li><code>.phtml</code> </li><li><code>.module</code> </li><li><code>.inc</code> </li><li><code>.install</code> </li><li><code>.theme</code> </li><li><code>.profile</code></li></ul></td><td><a href="php/#frameworks-and-libraries">PHP</a>에서 전체 목록 확인</td><td><ul><li>라이선스 스캐닝 </li><li>보고서</li><li>사용자 정의 규칙 </li><li>Interfile 분석</li></ul></td></tr><tr><td><a href="python/">Python</a></td><td><p><strong>패키지 관리자</strong>: Pip, Poetry, pipenv, setup.py<br><strong>지원되는 파일</strong>:</p><ul><li><code>pyproject.toml</code> </li><li><code>poetry.lock</code></li><li> <code>requirements.txt</code></li><li> <code>pipfile</code></li><li> <code>pipfile.lock</code></li><li> <code>setup.py</code></li></ul></td><td><strong>패키지 관리자</strong>: Pip, Poetry, pipenv, setup.py<br><strong>지원되는 파일</strong>: <code>.py</code>, </td><td><a href="python/### Python

- Fix PRs
- License scanning
- Reports
- Customer rules
- Interfile analysis

For more details, refer to the [Python](python.md#frameworks-and-libraries) page.

## Ruby

**Package manager**: Bundler  
**Supported files**:
- `gemfile`
- `gemfile.lock`

**Package manager**: Bundler  
**Supported files**:
- `.erb`
- `.haml`
- `.rb`
- `.rhtml`
- `.slm`

For more details, refer to the [Ruby](ruby/#frameworks-and-libraries) page.

- Fix PRs
- License scanning
- Reports
- Custom rules

## Rust

**Package manager**: Cargo  
**Supported files**: N/A

## Scala

**Package manager**: sbt  
**Supported files**: `build.sbt`

**Package manager**: sbt  
**Supported files**: `.scala`

For more details, refer to the [Scala](scala/#frameworks-and-libraries) page.

- License scanning
- Reports
- Interfile analysis

## Swift and Objective-C

**Package manager**: CocoaPods, Swift Package Manager  
**Supported files**:
- `podfile`
- `podfile.lock`
- `package.swift`

**Package manager**: CocoaPods, Swift Package Manager  
**Supported files**: `.swift`

For more details, refer to the [Swift and Objective-C](swift-and-objective-c/#frameworks-and-libraries) page.

- License scanning (CocoaPods)
- Reports
- Interfile analysis

## TypeScript

**Package manager**: npm, pnpm, Yarn  
**Supported files**:
- `package.json`
- `package-lock.json`
- `pnpm-lock.yaml`
- `yarn.lock`

**Package manager**: npm, pnpm, Yarn  
**Supported files**:
- `.ejs`
- `.es`
- `.es6`
- `.htm`
- `.html`
- `.js`
- `.jsx`
- `.ts`
- `.cts`
- `.mts`
- `.tsx`
- `.vue`
- `.mjs`
- `.cjs`

For more details, refer to the [TypeScript](typescript.md#frameworks-and-libraries).

- Reports
- Custom rules
- Interfile analysis
- Fix PRs
- License scanning

## VB NET

**Package manager**: N/A  
**Supported files**: `.vb`

For more details, refer to the [VB NET](vb.net.md#frameworks-and-libraries) page.

- Reports
- Interfile analysis

---
**Interfile analysis in Snyk Code is available for all languages supported except Ruby.**

**For Snyk Open Source, only official releases are tracked. Commits, including into the default branch, are not identified unless included in an official release or tag.**

**For Projects with a package manager, an official release of the package manager is required.**

**For Go and Unmanaged scans (C/C++), an official release or tag on the GitHub repository is required.**