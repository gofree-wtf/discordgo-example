discordgo-example
=================

Go로 된 Discord 클라이언트 API인 [discordgo](https://github.com/bwmarrin/discordgo) 를 테스트 해봅니다.

## How to run

### Step 1. 디스코드 앱 및 봇 생성

https://discord.com/developers/applications

### Step 2. 로컬에서 봇 실행

(INSERTHERE에 자신의 봇 토큰으로 대체)

```bash
$ make run BOT_TOKEN=INSERTHERE
```

### Step 3. 서버에 봇 초대

(INSERTHERE에 자신의 앱 클라이언트 ID로 대체)

https://discord.com/oauth2/authorize?client_id=INSERTHERE&scope=bot&permissions=19456

### Step 4. 채널에서 ping 메세지 보내보기

봇이 Pong! 메세지를 준다면 테스트 완료
