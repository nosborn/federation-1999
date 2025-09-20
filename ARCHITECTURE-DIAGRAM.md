# Federation: 1999 - System Architecture

This diagram illustrates how the various programs in the Federation: 1999 codebase work together to create a multi-user text-based space trading game.

```mermaid
graph TB
    %% External connections
    Player[👤 Player<br/>Telnet Client]
    Internet[🌐 Internet<br/>Port 23/992]

    %% Main programs
    telnetd[📡 telnetd<br/>Telnet Server<br/>Modem Emulation<br/>:8081 monitoring]
    login[🔐 login<br/>Authentication<br/>User Validation]
    perivale[🖥️ perivale<br/>Terminal Handler<br/>I/O Processing]
    fedtpd[🎮 fedtpd<br/>Game Server<br/>:8080 monitoring]
    workbench[🏗️ workbench<br/>Planet Editor<br/>Player Tool]

    %% Data storage
    billing_db[(💳 Billing Database<br/>SQLite<br/>User Accounts)]
    persona_db[(👥 Persona Database<br/>Binary File<br/>Player Data)]

    %% Internal systems
    subgraph "Game Engine Components"
        direction TB
        sessions[🎯 Session Manager<br/>Player Sessions]
        parser[📝 Command Parser<br/>Game Commands]
        world[🌍 Game World<br/>Locations & Objects]
        systems[⭐ Star Systems<br/>Sol, Snark, Arena, Horsell]
    end

    subgraph "Core Systems"
        direction TB
        monitoring[📊 Health & Metrics<br/>Prometheus]
        global_state[🔄 Global State<br/>Locks & Jobs]
        database[💾 Database Layer<br/>Player Persistence]
    end

    %% Communication flow
    Player -.->|Telnet Protocol| Internet
    Internet -->|Modem emulation<br/>V.32-V.90 speeds| telnetd
    telnetd -->|Spawns with PTY| login
    login -->|Validates credentials| billing_db
    login -->|Exec on success| perivale
    perivale -->|Unix socket| fedtpd

    %% Game server internals
    fedtpd --> sessions
    sessions --> parser
    parser --> world
    world --> systems

    %% Data persistence
    fedtpd --> database
    database --> persona_db
    sessions -.->|Player state| database

    %% Monitoring
    telnetd --> monitoring
    fedtpd --> monitoring

    %% Global coordination
    fedtpd --> global_state

    %% Planet editing and validation
    fedtpd -->|Runs in check mode| workbench
    workbench -->|Self-contained editor| world

    %% Styling for better aesthetics
    classDef program fill:#e1f5fe,stroke:#0277bd,stroke-width:2px,color:#000
    classDef database fill:#f3e5f5,stroke:#7b1fa2,stroke-width:2px,color:#000
    classDef external fill:#e8f5e8,stroke:#2e7d32,stroke-width:2px,color:#000
    classDef system fill:#fff3e0,stroke:#ef6c00,stroke-width:2px,color:#000

    class telnetd,login,perivale,fedtpd,workbench program
    class billing_db,persona_db database
    class Player,Internet external
    class sessions,parser,world,systems,monitoring,global_state,database system
```

## Program Descriptions

### **telnetd** 📡
- **Purpose**: Telnet server that accepts incoming connections
- **Features**: Handles PROXY protocol, PTY management, terminal settings, authentic 1990s modem emulation
- **Modem Speeds**: V.32 (9.6K), V.32bis (14.4K), V.32terbo (19.2K), V.34 (28.8K/33.6K), V.90 (56K)
- **Monitoring**: Health checks on port 8081
- **Key Role**: Entry point with period-appropriate connection experience

### **login** 🔐
- **Purpose**: Authentication system with security features
- **Features**: Password validation, login attempt limiting
- **Database**: Uses SQLite billing database for account management
- **Security**: Implements backoff delays and attempt limits

### **perivale** 🖥️
- **Purpose**: Terminal I/O handler and protocol processor
- **Features**: Non-blocking I/O, terminal formatting, spy/trace protocols
- **Communication**: Uses Unix domain sockets to connect to game server
- **Role**: Bridges between terminal display and game logic

### **fedtpd** 🎮
- **Purpose**: Main game server with multiplayer session management
- **Features**: Command processing, game world simulation, player persistence
- **Architecture**: Event-driven with global locks for consistency
- **Monitoring**: Health checks and Prometheus metrics on port 8080

### **workbench** 🏗️
- **Purpose**: Self-contained planet creation, editing, and validation tool
- **Features**: Interactive planet design with built-in validation logic
- **Integration**: Used by fedtpd in check-only mode to validate player planets before loading

## Data Flow

1. **Connection**: Player connects via telnet to port 23/992
2. **Authentication**: telnetd spawns login process for credential validation
3. **Session Setup**: login execs perivale which connects to fedtpd via Unix socket
4. **Game Loop**: Commands flow through perivale ↔ fedtpd, with responses formatted for terminal
5. **Persistence**: Player state automatically saved to binary persona database
6. **Monitoring**: Health checks and metrics exposed for operational visibility

## Key Features

- **Multi-user**: Concurrent player sessions with session management
- **Persistent World**: Binary database maintains game state between sessions
- **Security**: Rate limiting, and secure authentication
- **Monitoring**: Health checks and Prometheus metrics for observability
- **Single-server**: Designed for vertical scaling (bigger hardware) typical of 1999 era
