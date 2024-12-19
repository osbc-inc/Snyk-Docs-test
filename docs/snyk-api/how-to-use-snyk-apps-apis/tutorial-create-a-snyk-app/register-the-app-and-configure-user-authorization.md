# 앱 등록 및 사용자 권한 부여 구성

이 튜토리얼의 이전 섹션에서 TypeScript 프로젝트를 설정하고 Express 서버를 추가하고 기본 라우팅을 구성했습니다. 이전 섹션에서 만든 프로젝트를 기반으로 작업할 것입니다. 이미 완료하지 않았다면 이 튜토리얼의 이전 부분을 완료하는 것이 매우 권장됩니다.

## Snyk 앱 만들고 등록하기

우리는 지금까지 TypeScript 애플리케이션에 대한 어느 정도의 진전을 보았지만, 현재는 단순히 TypeScript 애플리케이션에 불과합니다. 이를 진정한 Snyk 앱으로 변환하려면 Snyk API를 사용하여 프로젝트를 새로운 앱으로 등록해야 합니다.

### 전제 조건

* API 권한이 있는 Snyk 계정
* [Snyk API 토큰](https://docs.snyk.io/features/snyk-api-info/authentication-for-api)
* 앱 소유자로서 등록할 Snyk 조직의 `orgid`

### `orgid` 얻기

```
https://app.snyk.io/org/{your-org-name}/manage/settings
```

{% hint style="info" %}
API 사용자가 사용하는 Snyk API Token은 API로 보내지지 않고 안전하게 보관됩니다.
{% endhint %}

### Snyk 앱 및 Snyk API에 대해

Snyk 앱은 유료로 앱을 설치한 사용자와 관계없이 API에 대한 풀 액세스 권한이 있습니다. 이 기능을 활용하려면 앱이 API에 접근할 때 `https://api.snyk.io/` 도메인의 API 엔드포인트를 사용해야 합니다.

### Snyk에 앱 등록

새로운 Snyk 앱을 등록하려면 간단한 POST 요청을 사용하여 Snyk의 API를 통해 수행됩니다. 이 튜토리얼 동안 코드를 설명하는 대신 `curl`을 사용하여 직접 요청을 보냅니다. 요청의 본문에는 다음 세부 정보가 필요합니다.

* `name`: Snyk 앱의 이름
* `redirectUris`: 최종 사용자 인증 중 허용되는 콜백 위치
* `scopes`: 사용자에게 부여할 Snyk 앱의 계정 권한

스코프에 대한 주의: 등록된 이후에는 Snyk 앱의 스코프를 현재 변경할 수 없습니다. 유일한 대응 방법은 [Delete App](https://snykv3.docs.apiary.io/#reference/apps/single-app-management/delete-app) API 엔드포인트를 사용하여 Snyk 앱을 삭제하고 새로운 Snyk 앱으로 다시 등록하는 것뿐입니다.

현재 작성 시점에서 **Snyk 앱은 여전히 베타 버전**입니다. 현재 사용 가능한 스코프는 **`apps:beta` 하나뿐**입니다. 이 스코프는 앱이 프로젝트를 테스트 및 모니터링하고 Snyk 조직, 기존 프로젝트, 문제, 보고서 정보를 읽을 수 있게 합니다.

**Snyk 앱 베타의 제한 중 하나는 Snyk 앱이 등록된 조직에 관리자 액세스 권한이 있는 사용자에게만 사용자 권한을 부여받을 수 있다는 점**입니다. 

API 토큰과 `orgid`를 사용하여 터미널에서 다음 명령을 수행하십시오. 이 튜토리얼에서는 `redirectUris` 값으로 `http://localhost:3000/callback`을 사용합니다.

```bash
curl --include \
     --request POST \
     --header "Content-Type: application/vnd.api+json" \
     --header "Authorization: token <API_TOKEN>" \
     --data-binary "{
       \"name\": \"My Awesome App\",
       \"redirectUris\": [ \"http://localhost:3000/callback\" ],
       \"scopes\": [ \"apps:beta\" ]
       }" \
     'https://api.snyk.io/rest/orgs/<ORG_ID>/apps?version='
```

Snyk로부터의 응답에는 Snyk 앱 통합을 완료하는 데 필요한 `clientId`와 `clientSecret` 두 가지 중요한 값이 포함되어 있습니다. 이러한 값을 안전하게 보관하십시오. 이는 Snyk의 `clientSecret`를 볼 수 있는 유일한 시기이며, **`clientSecret`를 공개적으로 공유해서는 안 됩니다**. 이 값은 Snyk과의 앱 인증에 사용됩니다.

이제 Snyk 앱으로 앱을 등록했기 때문에 TypeScript 프로젝트를 조정하여 사용자가 권한을 부여할 수 있도록 시작할 수 있습니다.```korean
const installs = db.data?.installs || [];

const index = installs.findIndex((install) => install.date === oldData.date);
if (index === -1) return false;
installs[index] = newData;
// 기존 설치를 새로운 설치로 대체함
db.data.installs = installs;
await db.write();
return true;
}
```

### API 호출을 위한 준비

이전에 API 호출을 처리하기 위해 인기 있는 `axios` 패키지를 설치했습니다. 동일한 API에 반복적으로 호출해야 하는 상황이 있기 때문에 프로젝트 전체에서 코드를 쉽게 재사용할 수 있도록 몇 가지 도우미 함수를 추상화합니다. `util` 디렉터리에 `APIHelpers.ts` 파일을 생성합니다.

그 내용을 작성하기 전에 Snyk의 API를 일관적으로 호출할 것이므로, Snyk API v1에서 Snyk REST API로의 마이그레이션 중인 엔드포인트 상태에 따라 API 버전을 여러 개 사용해야 할 가능성이 있습니다. 이를 처리하는 한 가지 방법은 TypeScript Enum을 정의하고 해당 Enum의 가능한 값과 인수를 비교하여 필요한 쿼리 매개변수를 교체하는 것입니다.

새 파일에 다음 내용을 추가하거나 원하는 경우 `APIHelpers.ts`에 추가합니다. 후에 사용하기 위해 반드시 내보내도록 하십시오.

```typescript
// ./interfaces/API.ts
export const enum APIVersion {
  V1 = "v1",
  REST = "rest",
}
```

우리는 Snyk API를 쉽게 호출하기 위한 단일 함수를 추가하여 시작합니다. 이 함수는 `tokenType` (bearer 또는 token), `token` 그리고 `APIVersion` (우리가 방금 정의한 Enum에 편리하게 해당함)을 취합니다.

```typescript
// ./src/util/APIHelpers.ts

import qs from "qs";
import axios, { AxiosInstance } from "axios";
import { API_BASE, CLIENT_ID, CLIENT_SECRET, TOKEN_URL } from "../app";
import { AuthData } from "../interfaces/DB";

export function callSnykApi(tokenType: string, token: string, version: APIVersion): AxiosInstance {
  const contentType = version === APIVersion.V1 ? "application/json": "application/vnd.api+json";

  const axiosInstance = axios.create({
    baseURL: `${API_BASE}/${version}`,
    headers: {
      "Content-Type": contentType,
      Authorization: `${tokenType} ${token}`,
    },
  });

  return axiosInstance;
}
```

이 함수가 `AxiosInstance` 이기 때문에 이와 같은 객체에 사용 가능한 `.get()`, `.post()` 또는 기타 메서드를 호출하여 쉽게 API의 다른 엔드포인트와 통신할 수 있습니다.

두 번째 비동기 함수를 정의하여 Snyk 앱의 조직 ID를 검색하는 것을 통해 이것이 작동하는 것을 확인하겠습니다.

```typescript
// ./src/util/APIHelpers.ts

...

export async function getAppOrgID(tokenType: string, accessToken: string): Promise<{ orgId: string }> {
  try {
    const clientId = CLIENT_ID;
    const result = await callSnykApi(tokenType, accessToken, APIVersion.REST)({
      method: "GET",
      url: `/apps/${clientId}/orgs?version=2021-08-11~experimental`,
    });
    // 첫 번째 조직 가져오기
    const org = result.data.data[0];
    return {
      orgId: org.id,
    };
  } catch (error) {
    console.error("Error fetching org info: " + error);
    throw error;
  }
}
```

### 암호화 / 복호화를 쉽게 만들기

API에서 가져올 데이터를 암호화하는 것은 좋은 아이디어입니다. 이를 수행하는 데 사용할 작은 클래스를 정의합니다. 이 클래스에는 두 가지 멤버가 있습니다.

1. `secret`: 데이터를 암호화하는 데 사용되는 키
2. `cryptr`: Cryptr 라이브러리의 인스턴스

```typescript
// ./src/util/encrypt-decrypt.ts

import Cryptr from "cryptr";

export class EncryptDecrypt {
  private secret: string;
  private cryptr: Cryptr;

  constructor(secret: string) {
    // 전달된 비밀 값 사용
    this.secret = secret as string;
    // 비밀 값으로 Cryptr 인스턴스 초기화
    this.cryptr = new Cryptr(this.secret);
  }

  /**
   * 데이터를 암호화하는 데 사용되는 함수
   * @param {String} 암호화할 메시지
   * @returns {String} 암호화된 메시지
   */
  public encryptString(message: string): string {
    return this.cryptr.encrypt(message);
  }
  /**
   * 데이터를 복호화하는 데 사용되는 함수
   * @param  {String} 복호화할 암호화된 문자열
   * @returns {String} 복호화된 문자열
   */
  public decryptString(encryptedString: string): string {
    return this.cryptr.decrypt(encryptedString);
  }
}
```

### Passport.js 및 Snyk-OAuth2 전략 구성

이제 준비 작업이 완료되었으니 이제 작업을 시작하는 시간입니다.

이전 섹션에서 설명한대로, 앱은 특정 토큰 URL로 인증자를 전송해야합니다. Snyk App의 `/auth` 경로를 추가하고 Express에 인증 미들웨어를 추가할 것입니다. 이를 위해 훌륭한 [passportjs](https://www.passportjs.org), [passport-oauth2](https://https/www.passportjs.org/packages/passport-oauth2) 인증 전략 및 Snyk의 [@snyk/passport-snyk-oauth2](https://www.npmjs.com/package/@snyk/passport-snyk-oauth2)을 사용할 것입니다. 

`passport` 및 해당 모듈은 일반적으로 긴 과정이 되는 인증 프로세스의 대부분을 처리합니다.

"passport"가 엔캡슐레이션 철학을 엄격히 준수하기 때문에 사용할 전략의 패스포트 인스턴스를 설정해야 합니다. 앞서 만든 데이터베이스 도우미를 여기에 활용하여 성공적인 인증을 받으면 데이터베이스에 항목을 추가합니다.

파일의 작업 내용을 이해했는지 확인하도록 시간을 들이시기 바랍니다.

```typescript
// ./util/OAuth2Strategy.ts

import type { Request } from "express";
import axios, { AxiosResponse, AxiosInstance } from "axios";
import OAuth2Strategy, { VerifyCallback } from "passport-oauth2";
import SnykOAuth2Strategy, { ProfileFunc } from "@snyk/passport-snyk-oauth2";
import { v4 as uuid4 } from "uuid";
import jwt_decode from "jwt-decode";
import { EncryptDecrypt } from "../util/encrypt-decrypt";
import { writeToDb } from "../util/db";
import { AuthData } from "../interfaces/db";
import { getAppOrgID } from "../util/APIHelpers";

// 이것은 간단히 공략의 앱 구성을 묶습니다.
// 직접 각 구성 변수를 작성하는 것을 피하고자 합니다.
// 보통 환경 변수를 구문 분석하고자 할 것입니다.
import * as config from "../app";

// 우리의 인증과 함께 사용할 매개변수에 대해 새로운 유형 정의합니다.
type Params = {
  expires_in: number;
  scope: string;
  token_type: string;
};

export function getOAuth2(): SnykOAuth2Strategy {
  // 프로필을 가져오려면 사용자가 프로필을 가져오는 방법에 대한 구현을 제공할 수 있습니다.
  // Snyk OAuth2 전략
  // 요청과 연결된 프로필을 가져오는 프로필Func를 제공합니다.
  const profileFunc: ProfileFunc = function(accessToken: string) {
    return axios.get("https://api.snyk.io/v1/user/me", {
      headers: {
        "Content-Type": "application/json; charset=utf-8",
        Authorization: `bearer ${accessToken}`,
      },
    });
  };

  // 버전 값이 수동으로 추가되었음에 유의하세요
  return new SnykOAuth2Strategy(
    {
      authorizationURL: `${config.APP_BASE}${config.AUTHORIZATION_URL}`,
      tokenURL: `${config.API_BASE}${config.TOKEN_URL}`,
      clientID: config.CLIENT_ID,
      clientSecret: config.CLIENT_SECRET,
      callbackURL: "http://localhost:3000/callback",
      state: true,
      pkce: true,
      passReqToCallback: true,
      profileFunc,
    },
    async function (
      req: Request,
      access_token: string,
      refresh_token: string,
      params: Params,
      profile: AxiosResponse,
      done: VerifyCallback
    ) {
      try {
        // 작업이 모두 완료되었음을 지정하기 위해 passport에 알림
        const userId = profile.data.id;
        const { expires_in, scope, token_type } = params;

        const { orgId } = await getAppOrgID(token_type, access_token);
        const ed = new EncryptDecrypt(config.ENCRYPTION_SECRET as string);

        await writeToDb({
          date: new Date(),
          userId,
          orgId,
          access_token: ed.encryptString(access_token),
          expires_in,
          scope,
          token_type,
          refresh_token: ed.encryptString(refresh_token),
        } as AuthData);
      } catch (error) {
        return done(error as Error, false);
      }
      return done(null, { nonce });
    }
  );
}
```

### Express 미들웨어 업데이트

Passport 전략을 구현한 후, `app.ts`를 수정하여 다음 코드 블록에 표시된 대로 `passport` 미들웨어를 설정합니다. 직접 호출하는 대신 `initGlobalMiddlewares()`라는 함수를 생성하여 동시에 몇 가지 다른 미들웨어를 설정할 수 있습니다.

* `express.json()`: JSON 요청 처리를 위한 Express 미들웨어
* `express.urlencoded()`: URL-encoded 호출을 처리하기 위한 Express 미들웨어
* `expressSession`: `express-session` 미들웨어 패키지, 이것은 `passport`에서 확장됩니다.
* `setupPassport`: `passport` 설정을 초기화하기 위해

```typescript
// ./src/app.ts

...

import passport from "passport";
import { getOAuth2 } from "./util/OAuth2Strategy";

...

constructor(controllers: Controller[], port: number) {
  ...
  this.initDatabaseFile();
  this.initGlobalMiddlewares();
  this.initRoutes(controllers);
  ...
}

...

private setupPassport() {
  passport.use(getOAuth2());
  this.app.use(passport.initialize());
  this.app.use(passport.session());
  passport.serializeUser((user: Express.User, done) => {
    done(null, user);
  });
  passport.deserializeUser((user: Express.User, done) => {
    done(null, user);
  });
}

private initGlobalMiddlewares() {
  this.app.use(express.json());
  this.app.use(express.urlencoded({ extended: true }));
  this.app.use(
    expressSession({
      secret: uuid4(),
      resave: false,
      saveUninitialized: true,
    })
  );
  this.setupPassport();
}

...
```

### 인가 및 콜백 경로 처리

인가 및 콜백 컨트롤러는 비교적 간단합니다. 다음 두 가지 새 컨트롤러 파일을 만듭니다.

```bash
mkdir -p ./src/routes/auth;
mkdir -p ./src/routes/callback;
touch ./src/routes/auth/authController.ts
touch ./src/routes/callback/callbackController.ts
```

`AuthController`는 이전에 설명한 인가 흐름을 통해 앱의 로그인을 처리합니다. 이것은 `passport` 설정의 세 번째 단계입니다. 모든 컨트롤러 클래스는 두 가지 멤버인 경로와 라우터를 포함하는 컨트롤러 인터페이스를 구현합니다.

이 컨트롤러는 사용자를 (passport의 도움을 받아) Snyk 웹 사이트로 보내기 위해 사용할 `/auth` 경로를 처리합니다.

```typescript
// ./src/routes/auth/authController.ts

import type { Controller } from "../../interfaces/Controller";
import { Router } from "express";
import passport from "passport";

class AuthController implements Controller {
  // 이 컨트롤러의 기본 URL 경로
  public path = "/auth";
  // 이 컨트롤러를 위한 Express 라우터
  public router = Router();

  constructor() {
    this.initRoutes();
  }

  /**
   * /auth 경로는 passportjs authenticate 방법을 사용하여
   * 앱의 인증을 처리합니다.
   */
  private initRoutes() {
    this.router.get(`${this.path}`, passport.authenticate("snyk-oauth2"));
  }
}

export default AuthController;
```

사용자가 Snyk 웹 사이트를 통해 저희의 Snyk 앱에 대한 인가를 승인하면, 사용자는 로컬 앱의 콜백 경로인 `/callback`로 다시 전송됩니다. 우리는 이 경로를 유사하게 처리하여 passport를 다시 활성화할 것입니다. 이는 사용자 인가의 마지막 단계입니다.

`CallbackController`는 `/callback`에서 요청을 수락할 뿐만 아니라 `/callback/success` 및 `/callback/failure`의 두 하위 경로를 생성하여 Snyk로부터 받을 수 있는 다양한 결과를 처리합니다.

```typescript
// ./src/routes/callback/callbackController.ts

import type { Controller } from '../../interfaces/Controller';
import type { NextFunction, Request, Response } from 'express';
import { Router } from 'express';

export class CallbackController implements Controller {
  public path = '/callback';
  public router: Router = Router();

  constructor() {
    this.initRoutes();
  }

  private initRoutes() {
    // 인증 흐름 결과나 콜백/리디렉트 URI를 처리하는 경로
    this.router.get(`${this.path}`, this.passportAuthenticate());
    // 성공을 처리하는 경로, passport로 전달한 것과 동일한
    this.router.get(`${this.path}/success`, this.success);
    // 실패 처리하는 경로, passport로 전달한 것과 동일한
    this.router.get(`${this.path}/failure`, this.failure);
  }
  
  private passportAuthenticate() {
    return passport.authenticate('snyk-oauth2', {
      successRedirect: '/callback/success',
      failureRedirect: '/callback/failure',
    });
  }

  private success(req: Request, res: Response, next: NextFunction) {
    // return res.render('callback');
    return res.send('SUCCESS!');
  }

  private failure(req: Request, res: Response, next: NextFunction) {
    // return next(new HttpException(401, 'Authentication failed'));
  }
}

export default CallbackController;
```

마무리 전에 새 컨트롤러들에 대한 참조를 `index.ts`에 추가해야 합니다.

```typescript
// ./src/index.ts

import IndexController from "./routes/index/indexController";
import AuthController from "./routes/auth/authController";
import CallbackController from "./routes/callback/callbackController";
import App from "./app";

new App([
   new IndexController(),
   new AuthController(),
   new CallbackController()],
  3000
);
```

### 토큰 갱신 관리

현재 이 시점에서 Snyk 앱을 빌드하고 실행하면 `/auth` 경로를 사용하여 Snyk 인가 포털로 성공적으로 이동하고, 인가를 확인하면 우리 로컬 앱의 콜백 경로인 `/callback`로 되돌아오게 됩니다. 매우 간단한 사용 사례가 있다면 여기서 끝낼 수 있지만, 사용자의 인가를 계속 유지하는 경우에는 마지막 조각이 있습니다; 토큰 만료.

어플리케이션을 테스트하기 위해 어플리[Axios interceptors](https://axios-http.com/docs/interceptors)를 활용하여 Snyk 앱 내에서이 프로세스를 자동화 할 수 있습니다. 이를 통해 만드는 요청들을 가로채서 최신 `access_token`을 보장합니다.

먼저, 필요한 모든 패키지, 클래스 등을 상단에 import하여 `./src/util/interceptors.ts` 파일을 생성하세요.

```typescript
// ./src/util/interceptors.ts

import type { AxiosRequestConfig } from "axios";
import { AuthData } from "../interfaces/DB";
import { DateTime } from "luxon";
import { readFromDb, updateDb } from "./DB";
import { mostRecent } from "../routes/projects/projectsController";
import { EncryptDecrypt } from "./encrypt-decrypt";
import { refreshAuthToken } from "../util/APIHelpers";
import { ENCRYPTION_SECRET } from "../app";
import axios from "axios";
```

총 세 개의 인터셉터를 추가할 것입니다.

첫 번째로, `refreshTokenReqInterceptor`는 `auth_token`이 만료 될 때 `refresh_token`을 사용하여 `auth_token`을 새로고침합니다. `AxiosRequestConfig` 요청을 인자로 사용하여 추가적인 확인에 사용할 수 있습니다.

```typescript
// ./src/util/interceptors.ts

...

export async function refreshTokenReqInterceptor(request: AxiosRequestConfig): Promise<AxiosRequestConfig> {
  // 최신 데이터(인증 토큰, 리프레시 토큰 및 만료) 읽기
  const db = await readFromDb();
  const data = mostRecent(db.installs);
  // 데이터가 없으면 요청 계속
  if (!data) return request;
  // 만료 시간 계산을 위해 사용되는 데이터
  const expiresIn = data.expires_in;
  const createdDate = data.date;
  // 만료일 계산
  const parsedCreateDate = DateTime.fromISO(createdDate.toString());
  const expirationDate = parsedCreateDate.plus({ seconds: expiresIn });
  // 만료되었는지 확인
  if (expirationDate < DateTime.now()) {
    await refreshAndUpdateDb(data);
  }
  return request;
}
```

`refreshTokenRespInterceptor`은 요청 응답 중에 사용됩니다. 받는 응답이 401 Unauthorized 인 경우에만 토큰을 새로 고치고 다시 시도합니다. 이는 문제가 발생했을 때 Passport가 반환하는 것입니다.

```typescript
// ./src/util/interceptors.ts

...

export async function refreshTokenRespInterceptor(error: AxiosError): Promise<AxiosError> {
  const status = error.response ? error.response.status: null;

  // Access-token이 만료되기 전에 무효화되는 경우에만 401 Unauthorized 상태에서 토큰을 새로 고칩니다.
  if (status === 401) {
    // 최신 데이터(인증 토큰, 리프레시 토큰 및 만료) 읽기
    const db = await readFromDb();
    const data = mostRecent(db.installs);
    // 데이터가 없을 경우 재시도 실패
    if (!data) return Promise.reject(error);

    const newAccessToken = await refreshAndUpdateDb(data);

    // 새로 받은 액세스 토큰을 사용하여 실패한 요청을 다시 시도
    error.config.headers['Authorization'] = `${data.token_type} ${newAccessToken}`;
    return axios.request(error.config);
  }

  return Promise.reject(error);
}
```

마지막으로, `refreshAndUpdateDb`는 주어진 데이터베이스 레코드의 액세스 토큰을 새로 고치고 다시 반환하기 전에 데이터베이스를 업데이트합니다.

```typescript
// ./src/util/interceptors.ts

...

async function refreshAndUpdateDb(data: AuthData): Promise<string> {
  // 암호화 및 복호화를위한 인스턴스 생성
  const eD = new EncryptDecrypt(process.env[Envars.EncryptionSecret] as string);
  // 토큰 새로 고치기 요청
  const { access_token, expires_in, refresh_token, scope, token_type } = await refreshAuthToken(
    eD.decryptString(data.refresh_token),
  );
  // 새로 가져온 액세스 및 리프레시 토큰을 사용하여 데이터베이스 업데이트
  // 만료 및 다른 필요한 정보와 함께
  await updateDb(data, {
    ...data,
    access_token: eD.encryptString(access_token),
    expires_in,
    refresh_token: eD.encryptString(refresh_token),
    token_type,
    scope,
    date: new Date(),
  });

  return access_token;
}
```

인터셉터를 정의했으므로, 할 일은 `callSnykApi` 함수를 업데이트하여 그것들을 사용하는 것뿐입니다. 인터셉터는 `axiosInstance` 객체의 메서드이므로 `axios.create()` 호출 이후 함수의 `return` 이전에 추가하면 됩니다.

```typescript
// ./src/util/APIHelpers.ts
...

import {
  refreshTokenReqInterceptor,
  refreshTokenRespInterceptor,
} from "./interceptors";

...

export function callSnykApi(tokenType: string, token: string, version: APIVersion): AxiosInstance {

  ...

  axiosInstance.interceptors.request.use(
    refreshTokenReqInterceptor,
    Promise.reject
  );
  axiosInstance.interceptors.response.use(
    (response) => response,
    refreshTokenRespInterceptor
  );

  ...

}

...
```

## 요약

지금까지 읽어주셔서 축하드립니다! Snyk와 등록, 권한 흐름 구성, `auth_token`의 유효 기간이 지나지 않도록 유지하고 TypeScript를 사용한 훌륭한 시작점을 설정하는 방법을 배웠습니다.

이 튜토리얼의 다음 모듈에서는 템플릿 시스템을 추가하고 Snyk의 모든 프로젝트를 사용자에게 보여주도록 앱을 구성할 것입니다.