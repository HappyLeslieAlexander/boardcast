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
		body{margin:0;padding:10px;height:100vh;display:flex;flex-direction:column;font-family:system-ui,sans-serif;background:#f5f5f5;transition:all .5s}
		.header{display:flex;justify-content:space-between;align-items:center;margin-bottom:15px}
		.logo{display:flex;align-items:center;gap:10px}
		.logo-text-1{font-size:18px;font-weight:600;color:#8fbffa;margin-right:-10px}
		.logo-text-2{font-size:18px;font-weight:600;color:#2859c5}
		.auth-form{display:flex;align-items:center;gap:10px}
		.btn{width:32px;height:32px;border-radius:4px;cursor:pointer;display:flex;align-items:center;justify-content:center;background:#f0f0f0;border:1px solid #ddd;transition:all .5s}
		.btn:hover{background:#e0e0e0}
		.btn.disabled{background:#ccc;cursor:not-allowed;opacity:0.5}
		.btn svg{width:16px;height:16px;fill:#666}
		#password{width:64px;padding:8px;border:1px solid #ddd;border-radius:4px;background:#f0f0f0;transition:all .5s}
		#password.status-disconnected{background:rgba(255,107,129,.5)}
		#password.status-connecting{background:rgba(255,193,7,.5)}
		#password.status-connected{background:rgba(34,197,94,.5)}
		#whiteboard,.placeholder{flex:1;width:100%%;background:#fff;border:1px solid #ddd;border-radius:4px;box-shadow:0 1px 3px rgba(0,0,0,.1);transition:all .5s}
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
	
/* --- Added styles for markdown preview placeholder, dark mode, responsive and draggable divider --- */
.editor-container{display:flex;flex-direction:column;height:60vh;min-height:300px;border-radius:8px;overflow:hidden}
#whiteboard{flex:1;min-height:120px;padding:10px;font-family:inherit;border:1px solid #ddd;border-bottom:none;resize:none;background:transparent}
#divider{height:8px;cursor:row-resize;display:block;background:linear-gradient(90deg, rgba(0,0,0,0.06), rgba(0,0,0,0.12));align-items:center;justify-content:center}
#divider .bar{width:60px;height:4px;border-radius:3px;margin:2px auto;opacity:.6}
.preview-wrapper{flex:1;overflow:auto;border:1px solid #ddd;padding:10px;background:var(--preview-bg,#fff);color:var(--preview-color,#111);min-height:80px}
.preview-placeholder{opacity:0.6;font-style:italic;color:var(--placeholder-color,#666)}
/* dark mode */
body.dark{background:#111;color:#eee}
body.dark .header{background:transparent}
body.dark .preview-wrapper{--preview-bg:#111;--preview-color:#f5f5f5;--placeholder-color:#bbb;border-color:#333}
/* mobile tweaks */
@media (max-width:768px){
  .editor-container{height:50vh}
  #whiteboard{font-size:14px;padding:8px}
  .preview-wrapper{font-size:14px;padding:8px}
}
</style>
  <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/dompurify@3.0.6/dist/purify.min.js"></script>
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
					<path d="M9 2c-1.05 0-2.05.16-3 .46 4.06 1.27 7 5.06 7 9.54 0 4.48-2.94 8.27-7 9.54.95.3 1.95.46 3 .46 5.52 0 10-4.48 10-10S14.52 2 9 2z"/>
				</svg>
			</button>
		</div>
	</div>
	<div class="placeholder" id="placeholder">Enter password to access BoardCast</div>
	<div class="editor-container">
    <textarea id="whiteboard" placeholder="Start typing markdown here..."></textarea>
    <div id="divider"><div class="bar"></div></div>
    <div id="preview" class="preview-wrapper"><div id="preview-inner"></div></div>
  </div>
	<footer style="text-align:center;font-size:10px;color:#999;margin-top:8px">
		BoardCast %s | Licensed under BSD 3-Clause | <a href="https://github.com/yosebyte/boardcast" target="_blank" style="color:#999;text-decoration:none;">View on GitHub</a>
	</footer>
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
			
			updateButtons=()=>{
				const canConnect=auth||p.value.trim();
				sb.disabled=!auth;rb.disabled=!auth;a.disabled=!canConnect;
				sb.classList.toggle('disabled',!auth);rb.classList.toggle('disabled',!auth);a.classList.toggle('disabled',!canConnect)
			},

			connect=()=>{
				if(!auth)return;
				status('connecting');
				s=new WebSocket((location.protocol==='https:'?'wss:':'ws:')+'//'+location.host+'/ws');
				s.onopen=()=>{status('connected');timer&&(clearTimeout(timer),timer=null);fetch('/content',{credentials:'include'}).then(r=>r.text()).then(c=>w.value=c).catch(()=>{})};
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
				connect();updateButtons()
			}).catch(()=>{p.value='';updateButtons()}),
			
			disconnect=()=>fetch('/logout',{method:'POST',credentials:'include'}).finally(()=>{
				timer&&(clearTimeout(timer),timer=null);s?.close();auth=false;p.value='';p.disabled=false;
				w.style.display='none';h.style.display='flex';w.value='';
				a.querySelector('path').setAttribute('d',icons.connect);status('disconnected');updateButtons()
			}),
			
			init=()=>fetch('/content',{credentials:'include'}).then(r=>{
				if(r.ok)return r.text();throw new Error('Not authenticated')
			}).then(c=>{
				auth=true;p.disabled=true;p.value='';w.style.display='block';h.style.display='none';
				a.querySelector('path').setAttribute('d',icons.disconnect);w.value=c;connect();updateButtons()
			}).catch(()=>{status('disconnected');updateButtons()}),
			
			snap=(u)=>auth&&fetch(u,{method:'POST',credentials:'include'}).catch(()=>{});
		
		load();
		t.onclick=()=>{document.body.classList.toggle('dark');icon();save()};
		a.onclick=()=>auth?disconnect():authenticate();
		sb.onclick=()=>snap('/save');
		rb.onclick=()=>snap('/restore');
		p.addEventListener('keypress',e=>e.key==='Enter'&&a.click());
		p.addEventListener('input',updateButtons);
		init()
	</script>
  <script>
    function updatePreview() {
      var raw = document.getElementById('whiteboard').value || '';
      var html = raw.trim() ? DOMPurify.sanitize(marked.parse(raw)) : '';
      var inner = document.getElementById('preview-inner');
      if(raw.trim()==='') {
        inner.innerHTML = '<div class="preview-placeholder">预览区：暂无内容 — 开始输入 Markdown 或粘贴文本以查看渲染结果。</div>';
      } else {
        inner.innerHTML = html;
      }
    }
    document.getElementById('whiteboard').addEventListener('input', updatePreview);
    window.addEventListener('load', updatePreview);
  </script>
<script>
// Divider drag to resize editor and preview (non-invasive)
(function(){
  var divider = document.getElementById('divider');
  var editor = document.getElementById('whiteboard');
  var preview = document.getElementById('preview');
  if(!divider||!editor||!preview) return;
  var dragging = false;
  var startY, startEditorH, startPreviewH;
  divider.addEventListener('mousedown', function(e){
    dragging = true;
    startY = e.clientY;
    startEditorH = editor.offsetHeight;
    startPreviewH = preview.offsetHeight;
    document.body.style.userSelect = 'none';
  });
  document.addEventListener('mousemove', function(e){
    if(!dragging) return;
    var dy = e.clientY - startY;
    var newEditorH = Math.max(60, startEditorH + dy);
    var newPreviewH = Math.max(60, startPreviewH - dy);
    editor.style.height = newEditorH + 'px';
    preview.style.height = newPreviewH + 'px';
  });
  document.addEventListener('mouseup', function(){
    if(dragging){ dragging=false; document.body.style.userSelect = ''; }
  });
  // Touch support
  divider.addEventListener('touchstart', function(e){ startY = e.touches[0].clientY; dragging=true; startEditorH = editor.offsetHeight; startPreviewH = preview.offsetHeight; });
  document.addEventListener('touchmove', function(e){ if(!dragging) return; var dy = e.touches[0].clientY - startY; editor.style.height = Math.max(60, startEditorH + dy) + 'px'; preview.style.height = Math.max(60, startPreviewH - dy) + 'px'; });
  document.addEventListener('touchend', function(){ dragging=false; });
})();
</script>
</body>
</html>`
