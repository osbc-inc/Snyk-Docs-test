# TypeScript

## 적용 가능성

Snyk은 Snyk Open Source 및 Snyk Code를 위해 TypeScript를 지원합니다. Snyk Open Source에서 TypeScript는 [JavaScript](javascript/javascript-for-open-source.md)와 정확히 동일한 방식으로 지원됩니다.

Snyk 제품을 사용하여 가져올 수 있는 언어의 가용성을 확인하고 애플리케이션으로 가져오거나 테스트하거나 모니터링할 수 있습니다.&#x20;

{% hint style="info" %}
**TypeScript 지원 버전**

4.2 이하의 버전을 사용할 수 있습니다.
{% endhint %}

사용 가능한 기능:

- SCM 가져오기:  및 에서 사용 가능합니다.
- CLI 및 IDE를 통한 앱 테스트 또는 모니터링:  및 에서 사용 가능합니다.
- 앱의 SBOM 테스트: `pkg:npm`을 사용합니다.&#x20;
- 패키지 테스트: `pkg:npm`을 사용합니다.

## 패키지 관리자 및 지원되는 파일

TypeScript에 대해 Snyk은 다음과 같은 패키지 관리자인 npm, pnpm, Yarn을 지원하며 다음과 같은 버전을 지원합니다.&#x20;

- npm: `Lockfile 1`, `Lockfile 2`, `Lockfile 3, 7.*`&#x20;
- pnpm: `pnpm 7`, `pnpm 8`, `pnpm 9`&#x20;
- Yarn: `Yarn 1`, `Yarn 2`, `Yarn 3`

패키지 레지스트리로는 [npmjs.org](https://www.npmjs.org/)를 지원합니다.

TypeScript에 대해 Snyk은 다음 파일 형식을 지원합니다:

- :&#x20;
  - npm의 경우: `package.json`, `package-lock.json`
  - pnpm의 경우: `pnpm-lock.yaml`,&#x20;
  - Yarn의 경우: `yarn.lock`
- : `.ejs`, `.es`, `.es6`, `.htm`, `.html`, `.js`, `.jsx`, `.ts`, `.cts`, `.mts`, `.tsx`, `.vue`, `.mjs`, `.cjs`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리는 TypeScript에 대해 Snyk에서 지원됩니다:&#x20;

- 모든 [JavaScript 프레임워크 및 라이브러리](javascript/#frameworks-and-libraries)
- @Google Drive/generative-ai - 포괄적인 지원&#x20;
- @anthropic-ai/sdk - 포괄적인 지원&#x20;
- @huggingface/inference - 포괄적인 지원&#x20;
- @mistralai/mistralai - 포괄적인 지원&#x20;
- axios - 포괄적인 지원&#x20;
- Angular - 부분적인 지원&#x20;
- apollo-server - 부분적인 지원&#x20;
- bcrypt-nodejs - 포괄적인 지원&#x20;
- cross-spawn - 포괄적인 지원&#x20;
- crypto-js - 포괄적인 지원&#x20;
- date-fns - 포괄적인 지원&#x20;
- dayjs - 포괄적인 지원&#x20;
- dompurify - 포괄적인 지원&#x20;
- electron - 부분적인 지원&#x20;
- ejs - 부분적인 지원&#x20;
- execa - 포괄적인 지원&#x20;
- express - 포괄적인 지원&#x20;
- express-graphql - 부분적인 지원&#x20;
- express-jwt - 부분적인 지원&#x20;
- fs - 포괄적인 지원&#x20;
- fs-extra - 포괄적인 지원&#x20;
- fs-plus - 포괄적인 지원&#x20;
- graceful-fs - 포괄적인 지원&#x20;
- graphql-js - 부분적인 지원&#x20;
- jQuery - 포괄적인 지원&#x20;
- js-yaml - 포괄적인 지원&#x20;
- jzip - 포괄적인 지원&#x20;
- koa - 포괄적인 지원&#x20;
- koa-graphql - 포괄적인 지원&#x20;
- libxml - 포괄적인 지원&#x20;
- libxmljs - 포괄적인 지원&#x20;
- lodash - 포괄적인 지원&#x20;
- luxon - 포괄적인 지원&#x20;
- minimongo - 포괄적인 지원&#x20;
- minimist - 포괄적인 지원&#x20;
- mongodb - 포괄적인 지원&#x20;
- Mongoose - 포괄적인 지원&#x20;
- mercurius - 부분적인 지원&#x20;
- Nestjs - 부분적인 지원&#x20;
- Node Crypto - 포괄적인 지원&#x20;
- node-buffer - 부분적인 지원&#x20;
- node-cmd - 포괄적인 지원&#x20;
- Node Crypto - 포괄적인 지원&#x20;
- node-dir - 포괄적인 지원&#x20;
- node-forge - 포괄적인 지원&#x20;
- node-pty - 포괄적인 지원&#x20;
- node-serialize - 포괄적인 지원&#x20;
- octokit - 포괄적인 지원&#x20;
- openai - 포괄적인 지원&#x20;
- pg - 포괄적인 지원&#x20;
- pg-promise - 포괄적인 지원&#x20;
- React - 부분적인 지원&#x20;
- request-promise - 포괄적인 지원&#x20;
- restler - 부분적인 지원&#x20;
- rimraf - 포괄적인 지원&#x20;
- sanitize-html - 포괄적인 지원&#x20;
- shelljs - 포괄적인 지원&#x20;
- Stanford JS Crypto - 포괄적인 지원&#x20;
- superagent - 포괄적인 지원&#x20;
- tar-stream - 포괄적인 지원&#x20;
- unirest - 포괄적인 지원&#x20;
- unzip - 포괄적인 지원&#x20;
- underscore - 포괄적인 지원&#x20;
- url - 포괄적인 지원&#x20;
- vm - 포괄적인 지원&#x20;
- webstomp-client - 부분적인 지원&#x20;
- WebCryptoAPI - 포괄적인 지원&#x20;
- xpath - 포괄적인 지원&#x20;
- yargs - 포괄적인 지원

## 기능

다음 기능은 TypeScript에 대해 Snyk에서 지원됩니다:

|                                     |                                                                  |
| --------------------------------------------------- | ------------------------------------------------------------------------- |
| <ul><li>라이센스 스캔 </li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙</li><li>파일 간 분석</li></ul> |