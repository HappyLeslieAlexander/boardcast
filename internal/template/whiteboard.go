// Package template provides HTML templates for the boardcast application.
package template

// WhiteboardHTML contains the complete HTML template for the whiteboard interface.
const WhiteboardHTML = `<!DOCTYPE html>
<html>
<head>
	<title>BoardCast</title>
	<meta name="viewport" content="width=device-width,initial-scale=1">
	<style>
		*{box-sizing:border-box}
		body{margin:0;padding:20px;height:100vh;display:flex;flex-direction:column;font-family:system-ui,sans-serif;background:#f5f5f5;transition:all .5s}
		.header{display:flex;justify-content:space-between;align-items:center;margin-bottom:15px}
		.logo{display:flex;align-items:center;gap:10px}
		.logo-text-1{font-size:18px;font-weight:600;color:#8fbffa;margin-right:-10px}
		.logo-text-2{font-size:18px;font-weight:600;color:#2859c5}
		.auth-form{display:flex;align-items:center;gap:10px}
		.btn{width:32px;height:32px;border-radius:4px;cursor:pointer;display:flex;align-items:center;justify-content:center;background:#f0f0f0;border:1px solid #ddd;transition:all .5s}
		.btn:hover{background:#e0e0e0}
		.btn svg{width:16px;height:16px;fill:#666}
		#password{width:64px;padding:8px;border:1px solid #ddd;border-radius:4px;background:#f0f0f0;transition:all .5s}
		#password.status-disconnected{background:rgba(255,107,129,.5)}
		#password.status-connecting{background:rgba(255,193,7,.5)}
		#password.status-connected{background:rgba(34,197,94,.5)}
		#whiteboard,.placeholder{flex:1;width:100%;background:#fff;border:1px solid #ddd;border-radius:4px;box-shadow:0 1px 3px rgba(0,0,0,.1);transition:all .5s}
		#whiteboard{padding:20px;font-size:16px;line-height:1.5;resize:none;font-family:inherit;display:none}
		#whiteboard:focus{outline:none;border-color:#8fbffa;box-shadow:0 0 0 2px rgba(143,191,250,.2)}
		.placeholder{display:flex;align-items:center;justify-content:center;color:#999}
		body.dark{background:#1a1a1a}
		body.dark #whiteboard,body.dark .placeholder{background:#2d2d2d;border-color:#444;color:#e0e0e0}
		body.dark .placeholder{color:#999}
		body.dark #password{background:#2d2d2d;border-color:#444;color:#e0e0e0}
		body.dark #password.status-disconnected{background:rgba(255,107,129,.5)}
		body.dark #password.status-connecting{background:rgba(255,193,7,.5)}
		body.dark #password.status-connected{background:rgba(34,197,94,.5)}
		body.dark .btn{background:#3d3d3d;border-color:#555}
		body.dark .btn:hover{background:#4d4d4d}
		body.dark .btn svg{fill:#ccc}
		@media(max-width:768px){body{padding:10px}#whiteboard{padding:15px}}
	</style>
</head>
<body>
	<div class="header">
		<div class="logo">
			<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 14 14">
				<path fill="#8fbffa" d="M.58 1.961A1.92 1.92 0 0 1 1.937 1.4H12.06a1.92 1.92 0 0 1 1.92 1.92v7.362a1.92 1.92 0 0 1-1.92 1.92H6.995A6.283 6.283 0 0 0 .017 5.524V3.32c0-.51.202-.998.563-1.358Z"/>
				<path fill="#2859c5" d="M.768 6.73a.75.75 0 1 0 0 1.5a3.533 3.533 0 0 1 3.533 3.532a.75.75 0 0 0 1.5 0A5.033 5.033 0 0 0 .768 6.73m0 2.676a.75.75 0 0 0 0 1.5a.856.856 0 0 1 .856.856a.75.75 0 1 0 1.5 0A2.356 2.356 0 0 0 .768 9.406"/>
			</svg>
			<span class="logo-text-1">Board</span><span class="logo-text-2">Cast</span>
		</div>
		<div class="auth-form">
			<input type="password" id="password">
			<button class="btn" id="authBtn">
				<svg viewBox="0 0 24 24">
					<path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
				</svg>
			</button>
			<button class="btn" id="saveBtn">
				<svg viewBox="0 0 24 24">
					<path d="M17 3H5c-1.11 0-2 .9-2 2v14c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V7l-4-4zm-5 16c-1.66 0-3-1.34-3-3s1.34-3 3-3 3 1.34 3 3-1.34 3-3 3zm3-10H5V5h10v4z"/>
				</svg>
			</button>
			<button class="btn" id="restoreBtn">
				<svg viewBox="0 0 24 24">
					<path d="M13 3c-4.97 0-9 4.03-9 9H1l3.89 3.89.07.14L9 12H6c0-3.87 3.13-7 7-7s7 3.13 7 7-3.13 7-7 7c-1.93 0-3.68-.79-4.94-2.06l-1.42 1.42C8.27 19.99 10.51 21 13 21c4.97 0 9-4.03 9-9s-4.03-9-9-9zm-1 5v5l4.28 2.54.72-1.21-3.5-2.08V8H12z"/>
				</svg>
			</button>
			<button class="btn" id="themeBtn">
				<svg viewBox="0 0 24 24">
					<path d="M12 7c-2.76 0-5 2.24-5 5s2.24 5 5 5 5-2.24 5-5-2.24-5-5-5z"/>
				</svg>
			</button>
			<button class="btn" onclick="window.open('https://github.com/yosebyte/boardcast', '_blank')">
				<svg viewBox="0 0 24 24">
					<path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
				</svg>
			</button>
		</div>
	</div>
	<div class="placeholder" id="placeholder">Enter password to access BoardCast</div>
	<textarea id="whiteboard"></textarea>
	
	<script>
		const w=document.getElementById('whiteboard'),
			p=document.getElementById('password'),
			a=document.getElementById('authBtn'),
			h=document.getElementById('placeholder'),
			t=document.getElementById('themeBtn'),
			sb=document.getElementById('saveBtn'),
			rb=document.getElementById('restoreBtn');
		
		let s=null,auth=false,updating=false,timer=null;
		
		const status=st=>p.className='status-'+st,
			icons={
				connect:'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z',
				disconnect:'M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10 10-4.47 10-10S17.53 2 12 2zm5 13.59L15.59 17 12 13.41 8.41 17 7 15.59 10.59 12 7 8.41 8.41 7 12 10.59 15.59 7 17 8.41 13.41 12 17 15.59z',
				light:'M12 7c-2.76 0-5 2.24-5 5s2.24 5 5 5 5-2.24 5-5-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1z',
				dark:'M9 2c-1.05 0-2.05.16-3 .46 4.06 1.27 7 5.06 7 9.54 0 4.48-2.94 8.27-7 9.54.95.3 1.95.46 3 .46 5.52 0 10-4.48 10-10S14.52 2 9 2z'
			},
			icon=()=>t.querySelector('path').setAttribute('d',document.body.classList.contains('dark')?icons.light:icons.dark),
			load=()=>{localStorage.getItem('theme')==='dark'&&document.body.classList.add('dark');icon()},
			save=()=>localStorage.setItem('theme',document.body.classList.contains('dark')?'dark':'light'),
			
			connect=()=>{
				if(!auth)return;
				status('connecting');
				s=new WebSocket((location.protocol==='https:'?'wss:':'ws:')+'//'+location.host+'/ws');
				s.onopen=()=>{status('connected');timer&&(clearTimeout(timer),timer=null)};
				s.onmessage=e=>{updating||(w.value=e.data,w.setSelectionRange(w.value.length,w.value.length))};
				s.onclose=()=>{status('disconnected');auth&&!timer&&(timer=setTimeout(()=>{timer=null;connect()},3000))};
				s.onerror=()=>status('disconnected');
				w.oninput=()=>{s?.readyState===1&&(updating=true,s.send(w.value),setTimeout(()=>updating=false,50))}
			},
			
			authenticate=()=>fetch('/auth',{
				method:'POST',headers:{'Content-Type':'application/json'},credentials:'include',
				body:JSON.stringify({password:p.value})
			}).then(r=>r.ok?r.text():Promise.reject()).then(()=>{
				auth=true;p.disabled=true;p.value='';w.style.display='block';h.style.display='none';
				a.querySelector('path').setAttribute('d',icons.disconnect);
				fetch('/content',{credentials:'include'}).then(r=>r.text()).then(c=>w.value=c);
				connect()
			}).catch(()=>p.value=''),
			
			disconnect=()=>fetch('/logout',{method:'POST',credentials:'include'}).finally(()=>{
				timer&&(clearTimeout(timer),timer=null);s?.close();auth=false;p.value='';p.disabled=false;
				w.style.display='none';h.style.display='flex';w.value='';
				a.querySelector('path').setAttribute('d',icons.connect);status('disconnected')
			}),
			
			init=()=>fetch('/content',{credentials:'include'}).then(r=>{
				if(r.ok)return r.text();throw new Error('Not authenticated')
			}).then(c=>{
				auth=true;p.disabled=true;p.value='';w.style.display='block';h.style.display='none';
				a.querySelector('path').setAttribute('d',icons.disconnect);w.value=c;connect()
			}).catch(()=>status('disconnected')),
			
			snap=(u)=>auth&&fetch(u,{method:'POST',credentials:'include'}).catch(()=>{});
		
		load();
		t.onclick=()=>{document.body.classList.toggle('dark');icon();save()};
		a.onclick=()=>auth?disconnect():authenticate();
		sb.onclick=()=>snap('/save');
		rb.onclick=()=>snap('/restore');
		p.addEventListener('keypress',e=>e.key==='Enter'&&a.click());
		init()
	</script>
</body>
</html>`
