# ddd_challenge
[詳細](https://chatwork.connpass.com/event/263334)

# ユビキタス言語 
- user  
サービスを利用するユーザー
- chat  
メッセージのやり取りを行うチャット
- memberIDs
chatに所属するuserのid
- message  
送信済みで他のユーザーから閲覧可能なメッセージ
- scheduled message  
まだチャットには送信されておらず、指定の時刻になったらスケジューラーによって送信されるメッセージ
- scheduled sending time  
送信予定日時
- schedule sending message  
メッセージを送信予定時刻に送信するよう登録すること


# ドメインモデル図
![ddd drawio (1)](https://user-images.githubusercontent.com/52554440/202906727-26037f00-3c5b-4b57-a66c-c1dcc18a9bb6.png)
[diagramsのリンク](https://drive.google.com/file/d/1RRtEIXxGXNl_dkC8ilsaZBkRqGEwHCfz/view?usp=sharing)

# ドメインオブジェクトの型定義
※ メソッドに関して各フィールドのGetterは省略しています。
### user オブジェクト  
**プロパティ**  
id   uint64  
name string  

### chat オブジェクト
**プロパティ**  
name      string  
memberIDs []uint64  

**メソッド**  
EditName(n string) error  
AppendMembers(userIDs []uint64)  
DeleteMembers(userIDs []uint64)  

### message オブジェクト
**プロパティ**  
id        uint64  
text      text  
chatID    uint64  
userID    uint64  
createdAt time.Time  

### scheduledMessage オブジェクト
**プロパティ**  
id        uint64  
text      text  
chatID    uint64  
userID    uint64  
sendTime  sendTime  
createdAt time.Time 

**メソッド**  
ChangeSendTime(scheduledSendingTime time.Time) error  
ChangeText(text string) error
