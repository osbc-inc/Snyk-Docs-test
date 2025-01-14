# AWS Lambda 스크립트 구성하기

제공된 예제 코드를 다양한 방법으로 구성하여 Slack에 원하는 정보를 얻을 수 있습니다.

핵심 영역은 페이로드(Slack 메시지가 구성되는 곳) 및 필터링(Snyk 정보가 처리되는 곳)입니다.

이외에도 암호화 및 비밀 검증을 구성할 수 있습니다. 이러한 구성은 이 지침의 범위를 넘어섭니다.

## Snyk 페이로드 필터링

다음 코드는 Snyk 페이로드를 필터링합니다.

다음과 같이 코드를 추가하세요:

```javascript
if(snykbody.indexOf("project") !== -1 && snykbody.indexOf("newIssues") !== -1){
    // 새 문제를 반복함
    var len = event.body['newIssues'].length;
    
    for(let x=0;x<len;x++){    
        // 심각성 얻기
        let severity = JSON.stringify(event.body['newIssues'][x]['issueData']['severity']);
        // 필터링
        if(severity.includes("high") || severity.includes("critical")){
            
            let snykProjectName = JSON.stringify(event.body['project'].name);
            let snykProjectUrl = JSON.stringify(event.body['project'].browseUrl);
            let snykIssueUrl = JSON.stringify(event.body['newIssues'][x]['issueData'].url);
            let snykIssueId = JSON.stringify(event.body['newIssues'][x].id);
            let snykIssuePackage = JSON.stringify(event.body['newIssues'][x].pkgName);
            let snykIssuePriority = JSON.stringify(event.body['newIssues'][x]['priority'].score);
            let message = "New Snyk Vulnerability";
            
            // 결과를 Slack으로 전송
            const result = await messageSlack(
                message,snykProjectUrl,snykProjectName,snykIssuePackage,snykIssueUrl,snykIssueId,severity,snykIssuePriority
            );
        } 
    }
```

이 코드는 다음을 실행합니다:

* 유효한 프로젝트 및 문제 확인
* 문제를 반복하면서 심각성이 높거나 심각한지 확인
* 지정한 모든 정보를 Slack 메시지로 전달. 원하는 정보를 지정하세요.

필터를 수정하여 예를 들어 특정 CWE만 처리하거나 모든 취약점을 허용하도록 할 수 있습니다.

## Slack 메시지 형식

Slack 페이로드는 다음과 같이 messageSlack 함수에서 형식화됩니다.

```javascript
async function messageSlack(message,snykProjectUrl,snykProjectName,snykIssuePackage,snykIssueUrl,snykIssueId,severity,snykIssuePriority) {
    
    // Axios/Slack 에러를 피하기 위해 문자열 수정
    snykProjectUrl = snykProjectUrl.replace(/['"]+/g, '')
    snykProjectName = snykProjectName.replace(/['"]+/g, '')
    snykIssueUrl = snykIssueUrl.replace(/['"]+/g, '')
    snykIssueId = snykIssueId.replace(/['"]+/g, '')
    snykIssuePackage = snykIssuePackage.replace(/['"]+/g, '')
    severity = severity.replace(/['"]+/g, '')
    
    // 메시지 구성
    let payload = { "text": `${message}`,
                    "blocks": [
		                {
                			"type": "header",
                			"text": {
                				"type": "plain_text",
                				"text": `${message}`,
                			}
                		},{
                			"type": "section",
                			"text": {
                				"type": "mrkdwn",
                				"text": "Snyk has found a new vulnerability in the project:\n*<"+snykProjectUrl+"|"+snykProjectName+">*"
                			}
                		},
                		{
                			"type": "divider"
                		},
                		{
                			"type": "section",
                			"fields": [
                				{
                					"type": "mrkdwn",
                					"text": "*Package name:*\n"+snykIssuePackage
                				},
                				{
                					"type": "mrkdwn",
                					"text": "*Vulnerability:*\n<"+snykIssueUrl+"|"+snykIssueId+">"
                				},
                				{
                					"type": "mrkdwn",
                					"text": "*Severity:*\n"+severity
                				},
                				{
                					"type": "mrkdwn",
                					"text": "*Priority Score:*\n"+snykIssuePriority
                				}
                			]
                		},
                		{
                			"type": "actions",
                			"elements": [
                				{
                					"type": "button",
                					"text": {
                						"type": "plain_text",
                						"text": "View in Snyk"
                					},
                					"style": "primary",
                					"url": snykProjectUrl,
                					"value": "browseUrl"
                				}
                			]
                		}
	               ]};
    
    // 메시지 전송
    const res = await axios.post(slackWebhookUrl, payload);
    const data = res.data;
    console.log(data);
}
```

Snyk는 Slack의 내장 블록 빌더를 사용하여 페이로드를 원하는 형식으로 디자인했습니다. 블록 빌더를 사용하여 페이로드의 JSON을 구성하여 CWE와 같은 더 많은 정보를 표시하거나 상호작용을 추가하고 Snyk의 ignore API를 사용하여 무시 버튼을 생성할 수 있습니다.

블록 빌더에 대한 자세한 정보는 Slack 웹사이트의 [Block Kit](https://api.slack.com/block-kit)에서 찾을 수 있습니다.
