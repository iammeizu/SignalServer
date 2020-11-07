# SignalServer

- Signal server for Webrtc by Go

---
### Architecture

- signal server is used to exchange sdp and candidate between Peer A and Peer B by websocket.
- sdp: session description protocol, media information including codec and so on. 
- candidate: net information including ip port and so on.

```

----------   sdp/candidate   ----------    sdp/candidate   ----------
| peer A |  -------------->  | signal |   -------------->  | peer B |
|        |  <-------------   | server |   <-------------   |        |
----------    (websocket)    ----------     (websocket)    ----------
```
---

### Attachment

[webrtc/pion](https://github.com/pion/webrtc.git)

[websocket](https://github.com/gorilla/websocket.git)

### Document

[webrtc](https://www.bookstack.cn/read/webrtc-book-cn/README.md)