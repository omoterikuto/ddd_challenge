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
![ddd drawio](https://user-images.githubusercontent.com/52554440/202905880-ee0d8148-c2c5-4c25-87ce-583a1ba89853.png)
[diagramsのリンク](https://drive.google.com/file/d/1RRtEIXxGXNl_dkC8ilsaZBkRqGEwHCfz/view?usp=sharing)

# ドメインオブジェクトの型定義
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
EditSendTime(scheduledSendingTime time.Time) error  
EditText(text string) error
