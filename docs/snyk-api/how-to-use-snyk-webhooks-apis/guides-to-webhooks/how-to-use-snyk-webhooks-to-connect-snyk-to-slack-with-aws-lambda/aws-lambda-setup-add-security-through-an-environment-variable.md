# AWS Lambda 설정: 환경 변수를 통한 보안 추가

보안상의 이유로 작성한 스크립트는 환경 변수를 사용합니다: `hmac_verification`은 Snyk에서 오는 페이로드를 유효성을 검증하고 변경되지 않았음을 확인하는 공유 비밀입니다.

환경 변수를 통한 보안 추가를 위해 다음 단계를 따르세요:

1. AWS Lambda 함수의 **구성** 탭으로 이동합니다.
2. **환경 변수**를 클릭합니다.
3. 새 환경 변수를 추가합니다.
4. `hmac_verification`을 키로 사용합니다.
5. 선호하는 비밀을 킷 값으로 입력합니다. 이 비밀을 안전한 곳에 보관하여 나중에 다시 사용합니다.

    <figure><img src="https://lh4.googleusercontent.com/eXXBAsVL2kDNpr9fDt_PErj9x0z7nBa-KywuWXJ0nGpuwwEiBiu8p0wFJLMacewmkRnYfrWSMzXqzhHAhRjifx-uEJF_BZm5Y0SazSMw60zKq8JOsLiGpqb7Risfr5zVBoBI7uiOJyMp_7G_HCajTB_vpIEVJotV4u1cJ4yO_t2wEi1jEARxk2sLjQ" alt=""><figcaption><p>AWS Lambda 함수에 환경 변수 추가 인터페이스</p></figcaption></figure>
6. 추가적인 보안을 위해 Secrets Manager나 Parameter Store를 사용하여 공유 비밀을 고려해보십시오.