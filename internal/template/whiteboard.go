// Package template provides HTML templates for the boardcast application.
package template

// WhiteboardHTML contains the complete HTML template for the whiteboard interface.
const WhiteboardHTML = `<!DOCTYPE html>
<html>
<head>
    <title>BoardCast</title>
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <style>
        * { box-sizing: border-box; }
        
        body { 
            margin: 0; 
            padding: 20px; 
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif; 
            background: #f5f5f5; 
            height: 100vh; 
            display: flex; 
            flex-direction: column; 
        }
        
        .header { 
            display: flex; 
            justify-content: space-between; 
            align-items: center;
            margin-bottom: 15px; 
        }
        
        .logo { 
            display: flex; 
            align-items: center; 
            gap: 10px; 
        }
        
        .logo-text-1 { 
            font-size: 18px; 
            font-weight: 600; 
            color: #8fbffa;
            margin-right: -10px;
        }

        .logo-text-2 { 
            font-size: 18px; 
            font-weight: 600; 
            color: #2859c5; 
        }

        .auth-form { 
            display: flex; 
            align-items: center; 
            gap: 10px; 
        }
        
        #password { 
            width: 120px;
            padding: 8px; 
            border: 1px solid #ddd; 
            border-radius: 4px; 
        }
        
        #authBtn { 
            width: 32px; 
            height: 32px; 
            background: #8fbffa; 
            border: none; 
            border-radius: 4px; 
            cursor: pointer; 
            display: flex; 
            align-items: center; 
            justify-content: center;
        }
        
        #authBtn:hover { background: #2859c5; }
        
        #authBtn svg { 
            width: 16px; 
            height: 16px; 
            fill: white; 
        }
        
        #whiteboard { 
            flex: 1; 
            width: 100%; 
            background: white; 
            border: 1px solid #ddd; 
            border-radius: 4px; 
            padding: 20px; 
            font-size: 16px; 
            line-height: 1.5; 
            resize: none; 
            box-shadow: 0 1px 3px rgba(0,0,0,0.1); 
            font-family: inherit;
            display: none;
        }
        
        #whiteboard:focus { 
            outline: none; 
            border-color: #8fbffa; 
            box-shadow: 0 0 0 2px rgba(143,191,250,0.2); 
        }
        
        .placeholder {
            flex: 1;
            display: flex;
            align-items: center;
            justify-content: center;
            background: white;
            border: 1px solid #ddd;
            border-radius: 4px;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
            color: #999;
        }
        
        @media (max-width: 768px) { 
            body { padding: 10px; } 
            #whiteboard { padding: 15px; } 
        }
    </style>
</head>
<body>
    <div class="header">
        <div class="logo">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 14 14">
                <g fill="none" fill-rule="evenodd" clip-rule="evenodd">
                    <path fill="#8fbffa" d="M.58 1.961A1.92 1.92 0 0 1 1.937 1.4H12.06a1.92 1.92 0 0 1 1.92 1.92v7.362a1.92 1.92 0 0 1-1.92 1.92H6.995A6.283 6.283 0 0 0 .017 5.524V3.32c0-.51.202-.998.563-1.358Z"/>
                    <path fill="#2859c5" d="M.768 6.73a.75.75 0 1 0 0 1.5a3.533 3.533 0 0 1 3.533 3.532a.75.75 0 0 0 1.5 0A5.033 5.033 0 0 0 .768 6.73m0 2.676a.75.75 0 0 0 0 1.5a.856.856 0 0 1 .856.856a.75.75 0 1 0 1.5 0A2.356 2.356 0 0 0 .768 9.406"/>
                </g>
            </svg>
            <span class="logo-text-1">Board</span>
            <span class="logo-text-2">Cast</span>
        </div>
        <div class="auth-form">
            <input type="password" id="password" placeholder="Password" required>
            <button id="authBtn" title="Connect">
                <svg id="authIcon" viewBox="0 0 24 24">
                    <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                </svg>
            </button>
        </div>
    </div>
    <div class="placeholder" id="placeholder">Enter password to access BoardCast</div>
    <textarea id="whiteboard"></textarea>
    
    <script>
        const whiteboard = document.getElementById('whiteboard');
        const passwordInput = document.getElementById('password');
        const authButton = document.getElementById('authBtn');
        const authIcon = document.getElementById('authIcon');
        const placeholder = document.getElementById('placeholder');
        
        let websocket = null;
        let authenticated = false;
        let isUpdatingFromServer = false;
        
        const connectIcon = '<path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>';
        const disconnectIcon = '<path d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10 10-4.47 10-10S17.53 2 12 2zm5 13.59L15.59 17 12 13.41 8.41 17 7 15.59 10.59 12 7 8.41 8.41 7 12 10.59 15.59 7 17 8.41 13.41 12 17 15.59z"/>';
        
        function authenticate() {
            fetch('/auth', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify({ password: passwordInput.value })
            })
            .then(response => response.ok ? response.text() : Promise.reject())
            .then(() => {
                authenticated = true;
                passwordInput.disabled = true;
                whiteboard.style.display = 'block';
                placeholder.style.display = 'none';
                authIcon.innerHTML = disconnectIcon;
                authButton.title = 'Disconnect';
                
                fetch('/content', { credentials: 'include' }).then(r => r.text()).then(content => whiteboard.value = content);
                
                websocket = new WebSocket((location.protocol === 'https:' ? 'wss:' : 'ws:') + '//' + location.host + '/ws');
                websocket.onmessage = e => {
                    if (isUpdatingFromServer) return;
                    const cursorStart = whiteboard.selectionStart;
                    const cursorEnd = whiteboard.selectionEnd;
                    whiteboard.value = e.data;
                    whiteboard.setSelectionRange(cursorStart, cursorEnd);
                };
                whiteboard.oninput = () => {
                    if (websocket?.readyState === 1) {
                        isUpdatingFromServer = true;
                        websocket.send(whiteboard.value);
                        setTimeout(() => isUpdatingFromServer = false, 50);
                    }
                };
            })
            .catch(() => passwordInput.value = '');
        }
        
        function disconnect() {
            fetch('/logout', { method: 'POST', credentials: 'include' })
            .finally(() => {
                websocket?.close();
                authenticated = false;
                passwordInput.value = '';
                passwordInput.disabled = false;
                whiteboard.style.display = 'none';
                placeholder.style.display = 'flex';
                whiteboard.value = '';
                authIcon.innerHTML = connectIcon;
                authButton.title = 'Connect';
            });
        }
        
        authButton.onclick = () => authenticated ? disconnect() : authenticate();
        passwordInput.addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                authButton.click();
            }
        });
    </script>
</body>
</html>`
