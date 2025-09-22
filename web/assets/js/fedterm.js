(() => {
  const WS_URL =
    (location.protocol === "https:" ? "wss:" : "ws:") +
    "//" +
    location.host +
    "/play";

  const IAC = 255;
  const DONT = 254;
  const DO = 253;
  const WONT = 252;
  const WILL = 251;
  const SB = 250;
  const SE = 240;

  const TELOPT_ECHO = 1;
  const TELOPT_SGA = 3;
  const TELOPT_LINEMODE = 34;

  const NAWS = 31;
  const TSPEED = 32;

  const term = new Terminal({
    allowTransparency: true,
    cursorInactiveStyle: "none",
    disableStdin: false,
    fontFamily: "monospace",
    fontSize: parseInt(getComputedStyle(document.body).fontSize),
    rightClickSelectsWord: false,
    windowsMode: true,
  });
  term.open(document.getElementById("terminal"));
  term.clear();

  const statusEl = document.getElementById("status");
  const input = document.getElementById("command");

  // history
  const history = [];
  let historyIdx = -1;

  // websocket + telnet logic
  let socket = null;

  // telnet state machine
  const STATE_DATA = 0,
    STATE_IAC = 1,
    STATE_CMD = 2,
    STATE_SB = 3,
    STATE_SB_IAC = 4;
  let telnetState = STATE_DATA;
  let telnetCmd = 0;
  let telnetBuffer = [];

  function setStatus(s) {
    statusEl.textContent = s;

    // Enable/disable input based on connection status
    const isConnected = s === "Connected";
    input.disabled = !isConnected;

    if (isConnected) {
      input.focus();
    }
  }

  function connect() {
    setStatus("Connecting...");
    try {
      socket = new WebSocket(WS_URL);
    } catch (e) {
      setStatus("Error");
      return;
    }
    socket.binaryType = "arraybuffer";

    socket.addEventListener("open", () => {
      setStatus("Connected");
      sendNAWS();
    });

    socket.addEventListener("message", (evt) => {
      const raw = new Uint8Array(evt.data);
      const text = handleTelnetFrame(raw);
      if (text) term.write(text);
    });

    socket.addEventListener("close", (evt) => {
      setStatus("Disconnected");
      socket = null;
    });

    socket.addEventListener("error", (evt) => {
      setStatus("Error");
    });
  }

  // send IAC bytes helper
  function sendIAC(bytes) {
    if (!socket || socket.readyState !== WebSocket.OPEN) return;
    socket.send(new Uint8Array(bytes));
  }

  // send NAWS (cols, rows)
  function sendNAWS() {
    if (!socket || socket.readyState !== WebSocket.OPEN) return;
    const cols = term.cols || 80;
    const rows = term.rows || 24;
    const payload = new Uint8Array([
      IAC,
      SB,
      NAWS,
      (cols >> 8) & 0xff,
      cols & 0xff,
      (rows >> 8) & 0xff,
      rows & 0xff,
      IAC,
      SE,
    ]);
    socket.send(payload);
  }

  // send TSPEED (terminal speed)
  function sendTSPEED() {
    if (!socket || socket.readyState !== WebSocket.OPEN) return;
    const speedStr = "9600,9600";
    const speedBytes = new TextEncoder().encode(speedStr);
    const payload = new Uint8Array([
      IAC,
      SB,
      TSPEED,
      0,
      ...speedBytes,
      IAC,
      SE,
    ]);
    socket.send(payload);
  }

  // Fit terminal to container and notify server
  function fitTerminal() {
    const container = document.getElementById('terminal');
    const rect = container.getBoundingClientRect();
    const fontSize = parseInt(getComputedStyle(document.body).fontSize);

    // Approximate character dimensions for monospace font
    const charWidth = fontSize * 0.6;
    const lineHeight = fontSize * 1.2;

    const cols = Math.floor(rect.width / charWidth);
    const rows = Math.floor(rect.height / lineHeight);

    if (cols > 0 && rows > 0) {
      term.resize(cols, rows);
      // sendNAWS will be called automatically by the onResize handler
    }
  }

  // Initial fit after DOM is ready
  setTimeout(fitTerminal, 0);

  // Fit on window resize
  window.addEventListener('resize', fitTerminal);

  term.onResize(({ cols, rows }) => {
    sendNAWS();
  });

  // Telnet parsing & negotiation:
  // - Accept WILL/DO by replying DO/WILL respectively (agree),
  // - Convert IAC IAC -> single 0xFF byte in output,
  // - Skip subnegotiations (SB ... SE), but if server asks DO NAWS we WILL and send size,
  // - Return the text to pass to terminal (decoded as UTF-8).
  function handleTelnetFrame(buf) {
    const outBytes = [];

    for (let i = 0; i < buf.length; i++) {
      const b = buf[i];

      switch (telnetState) {
        case STATE_DATA:
          if (b === IAC) {
            telnetState = STATE_IAC;
          } else {
            outBytes.push(b);
          }
          break;

        case STATE_IAC:
          if (b === IAC) {
            outBytes.push(IAC);
            telnetState = STATE_DATA;
          } else if (b === SB) {
            telnetState = STATE_SB;
            telnetBuffer = [];
          } else if (b >= WILL && b <= DONT) {
            telnetCmd = b;
            telnetState = STATE_CMD;
          } else {
            telnetState = STATE_DATA;
          }
          break;

        case STATE_CMD:
          handleOption(telnetCmd, b);
          telnetState = STATE_DATA;
          break;

        case STATE_SB:
          if (b === IAC) {
            telnetState = STATE_SB_IAC;
          } else {
            telnetBuffer.push(b);
          }
          break;

        case STATE_SB_IAC:
          if (b === SE) {
            if (telnetBuffer.length > 0) {
              handleSubnegotiation(telnetBuffer[0], telnetBuffer.slice(1));
            }
            telnetState = STATE_DATA;
          } else if (b === IAC) {
            telnetBuffer.push(IAC);
            telnetState = STATE_SB;
          } else {
            telnetState = STATE_SB;
          }
          break;
      }
    }

    if (outBytes.length === 0) return "";
    return String.fromCharCode(...outBytes);
  }

  function handleOption(cmd, opt) {
    switch (cmd) {
      case WILL:
        switch (opt) {
          case TELOPT_ECHO:
          case TELOPT_SGA:
            sendIAC([IAC, DO, opt]);
            break;
          default:
            sendIAC([IAC, DONT, opt]);
            break;
        }
        break;
      case WONT:
        sendIAC([IAC, DONT, opt]);
        break;
      case DO:
        if (opt === NAWS) {
          sendIAC([IAC, WILL, opt]);
          sendNAWS();
        } else if (opt === TSPEED) {
          sendIAC([IAC, WILL, opt]);
          sendTSPEED();
        } else {
          sendIAC([IAC, WILL, opt]);
        }
        break;
      case DONT:
        sendIAC([IAC, WONT, opt]);
        break;
    }
  }

  function handleSubnegotiation(opt, data) {
    // No server subnegotiations to handle currently
  }

  // Input handling
  input.addEventListener("keydown", (e) => {
    if (e.key === "Enter") {
      const cmd = input.value;
      if (!cmd && socket && socket.readyState === WebSocket.OPEN) {
        // send CRLF if empty line as well (some prompts expect)
        socket.send(new Uint8Array([13, 10]));
        input.value = "";
        historyIdx = -1;
        return;
      }
      if (socket && socket.readyState === WebSocket.OPEN) {
        // Send CRLF (telnet expects CRLF)
        const payload = new TextEncoder().encode(cmd + "\r\n");
        socket.send(payload);
      } else {
        term.writeln(
          "\r\n*** Not connected. Press Enter with empty input to try reconnect.",
        );
      }
      // push into history (ignore empty)
      if (cmd.trim() !== "") {
        history.unshift(cmd);
        if (history.length > 200) history.pop();
      }
      historyIdx = -1;
      input.value = "";
    } else if (e.key === "ArrowUp") {
      // history navigation
      e.preventDefault();
      if (history.length === 0) return;
      historyIdx = Math.min(historyIdx + 1, history.length - 1);
      input.value = history[historyIdx];
    } else if (e.key === "ArrowDown") {
      e.preventDefault();
      if (history.length === 0) return;
      historyIdx = historyIdx <= 0 ? -1 : historyIdx - 1;
      input.value = historyIdx === -1 ? "" : history[historyIdx];
    }
  });

  // connect on load
  connect();

  // allow reconnect by clicking status
  statusEl.addEventListener("click", () => {
    if (!socket) connect();
  });

  // initial message - ProComm-style FEDTERM logo
  term.writeln("");

  // set focus to input box
  // input.focus();
})();
