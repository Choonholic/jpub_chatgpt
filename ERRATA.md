# 오탈자, 보충 설명, 보완 주석 등
출판 후 발견된 오류, 보충 설명, 보완 주석 등을 이 문서에서 확인하실 수 있습니다.

## 정오표 (1쇄)

### 96쪽 9행:
공수를 할애하**는** 못합니다. → 공수를 할애하**지** 못합니다.

### 100쪽 3.1.3절 1행:
작성한**는** → 작성한
 
### 301쪽 '3 지원 실시' 코드: 
```python
from langchain.prompts import (
    ChatPromptTemplate,
    MessagesPlaceholder,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate
)
from langchain.chains import ConversationChain
from langchain.chat_models import ChatOpenAI
from langchain.memory import ConversationBufferMemory

from langchain.prompts import (
    ChatPromptTemplate,
    MessagesPlaceholder,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate
)
from langchain.chains import ConversationChain
from langchain.chat_models import ChatOpenAI
from langchain.memory import ConversationBufferMemory
```
불필요한 코드가 중복되어 있으므로, 다음과 같이 수정하시면 됩니다.
```python
from langchain.prompts import (
    ChatPromptTemplate,
    MessagesPlaceholder,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate
)
from langchain.chains import ConversationChain
from langchain.chat_models import ChatOpenAI
from langchain.memory import ConversationBufferMemory
```

## 보충 설명
