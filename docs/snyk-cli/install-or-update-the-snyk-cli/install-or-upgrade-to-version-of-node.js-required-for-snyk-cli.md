# Snyk CLI에 필요한 Node.js 버전 설치 또는 업그레이드

Snyk CLI 버전 1.853.0 이상에는 Node.js v12 이상이 필요합니다. Snyk는 완전한 호환성을 보장하기 위해 최신 버전의 Node.js 및 npm을 사용하는 것을 권장합니다.

### Node.js 업그레이드 방법

Node.js는 npm이 미리 설치되어 있지만, 이 매니저는 Node.js보다 더 자주 업데이트됩니다.

Node.js를 업데이트하려면, npm [n 모듈](https://www.npmjs.com/package/n)이 필요합니다.

다음 코드를 실행하여 npm 캐시를 지우고 `n`을 설치하고 최신 안정 버전의 Node.js를 설치합니다:

```bash
sudo npm cache clean -f
sudo npm install -g n
sudo n stable
```

최신 릴리스를 설치하려면 `n latest`를 사용합니다. 또는 특정 Node.js 버전을 얻으려면 `n #.#.#`를 실행할 수 있습니다.

Node.js를 업그레이드하려면 다음 단계를 따르십시오:

1. 현재 버전을 확인하려면 `npm -v`를 실행합니다.
2. 새로운 npm 업데이트를 설치하려면 `npm install npm@latest -g`를 실행합니다.
3. npm이 정확하게 업데이트되었는지 확인하려면 다시 `npm -v`를 실행합니다.

### npm 업그레이드 방법

npm을 업그레이드하려면 다음을 실행합니다:

```bash
npm install -g npm
```