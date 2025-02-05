# TypeScript

## 적용 가능성

Snyk는 Snyk Open Source 및 Snyk Code를 위해 TypeScript를 지원합니다. Snyk Open Source에서 TypeScript는 [JavaScript](javascript/javascript-for-open-source.md)와 정확히 동일한 방식으로 지원됩니다.

Snyk 제품을 사용하여 가져올 수 있는 언어의 가용성을 확인하고 애플리케이션으로 가져오거나 테스트하거나 모니터링할 수 있습니다.

{% hint style="info" %}
**TypeScript 지원 버전**

4.2 이하의 버전을 사용할 수 있습니다.
{% endhint %}

사용 가능한 기능:

* SCM 가져오기: 및 에서 사용 가능합니다.
* CLI 및 IDE를 통한 앱 테스트 또는 모니터링: 및 에서 사용 가능합니다.
* 앱의 SBOM 테스트: `pkg:npm`을 사용합니다.
* 패키지 테스트: `pkg:npm`을 사용합니다.

## 패키지 관리자 및 지원되는 파일

TypeScript에 대해 Snyk는 다음과 같은 패키지 관리자인 npm, pnpm, Yarn을 지원하며 다음과 같은 버전을 지원합니다.

* npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`
* pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`
* Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

패키지 레지스트리로는 [npmjs.org](https://www.npmjs.org/)를 지원합니다.

TypeScript에 대해 Snyk는 다음 파일 형식을 지원합니다:

* :
  * npm의 경우: `package.json`, `package-lock.json`
  * pnpm의 경우: `pnpm-lock.yaml`,
  * Yarn의 경우: `yarn.lock`
* : `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리는 TypeScript에 대해 Snyk에서 지원됩니다:

* 모든 [JavaScript 프레임워크 및 라이브러리](javascript/#frameworks-and-libraries)
* @Google Drive/generative-ai - 포괄적인 지원
* @anthropic-ai/sdk - 포괄적인 지원
* @huggingface/inference - 포괄적인 지원
* @mistralai/mistralai - 포괄적인 지원
* axios - 포괄적인 지원
* Angular - 부분적인 지원
* apollo-server - 부분적인 지원
* bcrypt-nodejs - 포괄적인 지원
* cross-spawn - 포괄적인 지원
* crypto-js - 포괄적인 지원
* date-fns - 포괄적인 지원
* dayjs - 포괄적인 지원
* dompurify - 포괄적인 지원
* electron - 부분적인 지원
* ejs - 부분적인 지원
* execa - 포괄적인 지원
* express - 포괄적인 지원
* express-graphql - 부분적인 지원
* express-jwt - 부분적인 지원
* fs - 포괄적인 지원
* fs-extra - 포괄적인 지원
* fs-plus - 포괄적인 지원
* graceful-fs - 포괄적인 지원
* graphql-js - 부분적인 지원
* jQuery - 포괄적인 지원
* js-yaml - 포괄적인 지원
* jzip - 포괄적인 지원
* koa - 포괄적인 지원
* koa-graphql - 포괄적인 지원
* libxml - 포괄적인 지원
* libxmljs - 포괄적인 지원
* lodash - 포괄적인 지원
* luxon - 포괄적인 지원
* minimongo - 포괄적인 지원
* minimist - 포괄적인 지원
* mongodb - 포괄적인 지원
* Mongoose - 포괄적인 지원
* mercurius - 부분적인 지원
* Nestjs - 부분적인 지원
* Node Crypto - 포괄적인 지원
* node-buffer - 부분적인 지원
* node-cmd - 포괄적인 지원
* Node Crypto - 포괄적인 지원
* node-dir - 포괄적인 지원
* node-forge - 포괄적인 지원
* node-pty - 포괄적인 지원
* node-serialize - 포괄적인 지원
* octokit - 포괄적인 지원
* openai - 포괄적인 지원
* pg - 포괄적인 지원
* pg-promise - 포괄적인 지원
* React - 부분적인 지원
* request-promise - 포괄적인 지원
* restler - 부분적인 지원
* rimraf - 포괄적인 지원
* sanitize-html - 포괄적인 지원
* shelljs - 포괄적인 지원
* Stanford JS Crypto - 포괄적인 지원
* superagent - 포괄적인 지원
* tar-stream - 포괄적인 지원
* unirest - 포괄적인 지원
* unzip - 포괄적인 지원
* underscore - 포괄적인 지원
* url - 포괄적인 지원
* vm - 포괄적인 지원
* webstomp-client - 부분적인 지원
* WebCryptoAPI - 포괄적인 지원
* xpath - 포괄적인 지원
* yargs - 포괄적인 지원

## 기능

다음 기능은 TypeScript에 대해 Snyk에서 지원됩니다:

|                                       |                                                         |
| ------------------------------------- | ------------------------------------------------------- |
| <ul><li>라이센스 스캔</li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙</li><li>파일 간 분석</li></ul> |
