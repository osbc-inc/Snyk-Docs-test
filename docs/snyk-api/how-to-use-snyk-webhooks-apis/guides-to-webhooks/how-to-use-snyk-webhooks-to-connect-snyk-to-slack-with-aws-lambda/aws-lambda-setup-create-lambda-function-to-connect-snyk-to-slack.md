# AWS Lambda 설정: Lambda 함수를 생성하여 Snyk을 Slack에 연결하는 방법

AWS Lambda 함수는 새로운 Snyk 취약점이 발견될 때와 같이 이벤트에 의해 트리거된 코드를 실행할 수 있는 저렴하고 효율적인 방법으로 Snyk을 Slack에 연결하기 위해 사용됩니다.

**참고:** Lambda 함수를 API Gateway를 통해 배포하는 경우 두 가지를 동일한 지역에 구성해야 합니다. 이를 AWS 콘솔의 오른쪽 상단에서 확인할 수 있습니다.

먼저 Lambda 함수와 필요한 종속성을 포함하는 zip 파일을 생성하십시오.

1. 다음 두 코드 블록을 package.json 및 index.js로 저장하세요.
   1.  package.json (코드에 필요한 다른 종속성이 있을 경우 수정하여 조정하십시오. 필요한 종속성은 **axios** 및 **crypto**입니다)

       ```json
       {
         "name": "snyk-webhook-handler",
         "version": "1.0.0",
         "description": "Snyk to Slack Webhook Integration",
         "main": "index.js",
         "scripts": {
           "test": "echo \"Error: no test specified\" && exit 1"
         },
         "author": "",
         "license": "ISC",
         "dependencies": {
           "axios": "^1.1.3",
           "crypto": "^1.0.1"
         }
       }
       ```
   2.  index.js

       ```javascript
       const crypto = require('crypto')
       const axios = require('axios')

       let slackWebhookUrl = '<your_slackWebhookUrl_here>' // 조정

       // 필요에 따라 수정할 Slack로 사용자 지정된 메시지
       async function messageSlack(
         message,
         snykProjectUrl,
         snykProjectName,
         snykIssuePackage,
         snykIssueUrl,
         snykIssueId,
         severity,
         snykIssuePriority
       ) {
         // 오류를 피하기 위해 문자열 수정
         snykProjectUrl = snykProjectUrl.replace(/['"]+/g, '')
         snykProjectName = snykProjectName.replace(/['"]+/g, '')
         snykIssueUrl = snykIssueUrl.replace(/['"]+/g, '')
         snykIssueId = snykIssueId.replace(/['"]+/g, '')
         snykIssuePackage = snykIssuePackage.replace(/['"]+/g, '')
         severity = severity.replace(/['"]+/g, '')

         // 메시지 구성
         let payload = {
           text: `${message}`,
           blocks: [
             {
               type: 'header',
               text: {
                 type: 'plain_text',
                 text: `${message}`,
               },
             },
             {
               type: 'section',
               text: {
                 type: 'mrkdwn',
                 text: `Snyk has found a new vulnerability in the project:\n*<${snykProjectUrl}|${snykProjectName}>*`,
               },
             },
             {
               type: 'divider',
             },
             {
               type: 'section',
               fields: [
                 {
                   type: 'mrkdwn',
                   text: `*Package name:*\n${snykIssuePackage}`,
                 },
                 {
                   type: 'mrkdwn',
                   text: `*Vulnerability:*\n<${snykIssueUrl}|${snykIssueId}>`,
                 },
                 {
                   type: 'mrkdwn',
                   text: `*Severity:*\n${severity}`,
                 },
                 {
                   type: 'mrkdwn',
                   text: `*Priority Score:*\n${snykIssuePriority}`,
                 },
               ],
             },
             {
               type: 'actions',
               elements: [
                 {
                   type: 'button',
                   text: {
                     type: 'plain_text',
                     text: 'View in Snyk',
                   },
                   style: 'primary',
                   url: snykProjectUrl,
                   value: 'browseUrl',
                 },
               ],
             },
           ],
         }

         // 메시지 전송
         const res = await axios.post(slackWebhookUrl, payload)
         const data = res.data
       }

       exports.handler = async (event) => {
         // 페이로드 무결성 보호, 인증을 위해 별도의 Lambda 함수로 이동 가능
         let response

         const {hmac_verification, severity_threshold} = process.env
         const hmac = crypto.createHmac('sha256', hmac_verification)
         const buffer = JSON.stringify(event.body)
         hmac.update(buffer, 'utf8')
         const stored_signature = `sha256=${hmac.digest('hex')}`

         let sent_signature = event.headers['x-hub-signature']

         if (stored_signature !== sent_signature) {
           console.log('요청의 무결성이 손상되었습니다. 중지')
           response = {
             statusCode: 403,
             body: JSON.stringify('잘못된 요청'),
           }
           return response
         }

         // 무결성이 확인되면, 웹훅에 프로젝트 개체가 실제로 포함되어 있는지 확인하고 필터링
         if (buffer.indexOf('project') !== -1 && buffer.indexOf('newIssues') !== -1) {
           // 새로운 이슈 반복
           var len = buffer['newIssues'] ? buffer['newIssues'].length : 0

           for (let x = 0; x < len; x++) {
             // 심각도 가져오기
             let severity = JSON.stringify(buffer['newIssues'][x]['issueData']['severity'])
             // 필터링
             if (severity.includes('high') || severity.includes('critical')) {
               let snykProjectName = JSON.stringify(buffer['project'].name)
               let snykProjectUrl = JSON.stringify(buffer['project'].browseUrl)
               let snykIssueUrl = JSON.stringify(buffer['newIssues'][x]['issueData'].url)
               let snykIssueId = JSON.stringify(buffer['newIssues'][x].id)
               let snykIssuePackage = JSON.stringify(buffer['newIssues'][x].pkgName)
               let snykIssuePriority = JSON.stringify(buffer['newIssues'][x]['priority'].score)
               let message = 'New Snyk Vulnerability'

               // 결과 Slack로 전송
               await messageSlack(
                 message,
                 snykProjectUrl,
                 snykProjectName,
                 snykIssuePackage,
                 snykIssueUrl,
                 snykIssueId,
                 severity,
                 snykIssuePriority
               )
             }
           }
         }
         // 아무것도 안 함 또는 필요한 작업에 맞게 수정
         else {
           console.log('유효한 웹훅이지만 프로젝트가 누락되었거나 비어 있습니다.')
         }

         // Snyk에 응답
         response = {
           statusCode: 200,
           body: JSON.stringify('성공'),
         }

         return response
       }

       ```
2. 터미널에서 다음 명령어 사용: - `cd /snyk/폴더/경로` (두 파일을 저장한 폴더로 이동)\
   \- `npm install` (package.json 파일 설정)\
   \- `cd ..` (이전 폴더로 이동)\
   \- `zip -r snyk.zip /snyk/폴더/경로` (AWS Lambda에 업로드할 수 있는 파일로 snyk 압축)

AWS Lambda 함수를 생성하려면 다음 단계를 따르십시오:

1. AWS 콘솔로 이동합니다.
2. Lambda로 이동합니다.
3. **함수 생성**을 클릭합니다.
4. **런타임**에 **Node.js 16.x**를 선택합니다.
5. **아키텍처**에 **X86\_64**를 선택합니다.
6. Lambda 함수를 API Gateway를 통해 게시하는 경우 AWS API Gateway와 상호 작용하도록 기본 정책 **AmazonAPIGatewayInvokeFullAccess**을 추가하거나 생성합니다.
7.  AWS 콘솔 화면에 다음 항목이 표시되는지 확인합니다:

    ![Lambda 함수 생성 항목이 있는 AWS 콘솔](https://lh6.googleusercontent.com/xzJzGjfuzj0U27-pxcaIcrU-wBj8DTuEiQpivJZAnqRAO3rEPccx48l8KSZ5AE01BYJDwjJwkiFMR-Oj3ozWyG-CI20bwFtK_yjY9HKEoY0-4V4pa8l351JqrYdkK29va1x7BdlPoQ7N12SROjDQy3CmUQsDTtQ5lYOw3QvwoG1c1nDms-EFiQSElA)
8. **함수 생성**을 클릭하고 함수가 준비되면 **.zip 파일에서 업로드**를 클릭합니다.
9.  입력한 코드가 **Code Source**에 표시되는지 확인합니다.

    ![AWS 코드 소스 표시](https://lh3.googleusercontent.com/97qnO6V9xBXaf6dyO0hg41Y2vmeB1-0aPK-qskqTI-L2WII3d75zb4XsK6Mg5ljJUEdS7AGYJ5sQ5IoDHvzofkfK_gPId9e-XuBqEGkuWNxlIyL4IHu7-S8hrbGKnuyOehU2fjScDi0jazvuhWkADyFDGkkdAdzQGSEfWO30YGPJ9x4ocfwFXS5LfQ)
10. 코드에서 `slackWebhookUrl`을 사용자의 Slack 웹훅 URL과 일치하도록 수정합니다.
11. 붙여 넣은 스크립트에 대한 자세한 내용은 [AWS Lambda 스크립트 구성](configure-the-aws-lambda-script.md)를 확인하십시오.
