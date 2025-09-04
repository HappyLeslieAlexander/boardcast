package template

// WhiteboardHTML 是返回给客户端的 HTML 模板。
// 注意：本文件做了最小侵入性的改动，加入 Markdown 预览功能 (marked + DOMPurify)。
// 保存/恢复/实时同步逻辑保持不变：仍然使用 /save /restore /ws /content 等后端接口。
const WhiteboardHTML = `<!doctype html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>BoardCast — 文字白板（支持 Markdown Preview）</title>
  <style>
    /* 简洁的样式：不改变原有布局（若你的项目已有样式，请保留） */
    html,body { height:100%; margin:0; font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial; }
    .app { display:flex; flex-direction:column; height:100vh; }
    header { padding:10px; display:flex; gap:8px; align-items:center; background:#f8f9fa; border-bottom:1px solid #e5e5e5; }
    header .title { font-weight:600; font-size:18px; }
    main { flex:1; display:flex; gap:8px; padding:12px; box-sizing:border-box; }
    /* editor / preview 列布局 */
    .pane { flex:1; display:flex; flex-direction:column; min-width:0; }
    textarea#editor { flex:1; width:100%; resize:none; padding:12px; box-sizing:border-box; font-family: monospace; font-size:14px; line-height:1.5; border:1px solid #ddd; border-radius:6px; }
    .preview { flex:1; overflow:auto; padding:12px; border:1px solid #ddd; border-radius:6px; background:#fff; }
    .controls { display:flex; gap:8px; }
    .small { padding:6px 10px; font-size:13px; }
    /* Markdown preview rendered elements */
    .preview h1, .preview h2, .preview h3 { margin:12px 0 6px; }
    .preview p { margin:8px 0; }
    .bottom-bar { padding:8px 12px; border-top:1px solid #eee; background:#fafafa; }
    .muted { color:#666; font-size:13px; }
    @media (max-width:900px) {
      main { flex-direction:column; }
      .pane { min-height:150px; }
    }
  </style>
</head>
<body>
  <div class="app">
    <header>
      <div class="title">BoardCast — 文字白板</div>
      <div class="controls" style="margin-left:auto;">
        <button id="togglePreviewBtn" class="small">显示 Markdown 预览</button>
        <button id="saveBtn" class="small">保存快照</button>
        <button id="restoreBtn" class="small">恢复快照</button>
        <button id="logoutBtn" class="small">退出</button>
      </div>
    </header>

    <main>
      <section class="pane" aria-label="editor-pane">
        <label class="muted">编辑（文本/Markdown）</label>
        <textarea id="editor" placeholder="在这里输入文本或 Markdown..." aria-label="editor"></textarea>
        <div class="bottom-bar muted">保存的内容仍为纯文本（如果你输入 Markdown，它会被原样保存）。</div>
      </section>

      <section class="pane" aria-label="preview-pane">
        <label class="muted">渲染预览</label>
        <!-- preview 区域，初始隐藏提示 -->
        <div id="preview" class="preview" aria-live="polite">切换“显示 Markdown 预览”以查看渲染结果。</div>
      </section>
    </main>
  </div>

  <!-- 外部库：marked（Markdown -> HTML），DOMPurify（HTML 消毒）-->
  <!-- CDN：若需要离线托管，请将脚本放到静态目录并修改引用路径 -->
  <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js" integrity="" crossorigin="anonymous"></script>
  <script src="https://cdn.jsdelivr.net/npm/dompurify/dist/purify.min.js" integrity="" crossorigin="anonymous"></script>

  <script>
  (function(){
    'use strict';

    // DOM 元素
    const editor = document.getElementById('editor');
    const preview = document.getElementById('preview');
    const togglePreviewBtn = document.getElementById('togglePreviewBtn');
    const saveBtn = document.getElementById('saveBtn');
    const restoreBtn = document.getElementById('restoreBtn');
    const logoutBtn = document.getElementById('logoutBtn');

    // 预览状态：false -> 编辑，true -> 显示渲染
    let previewEnabled = false;

    // 初始化：从后端获取当前内容（如果后端已有 /content）
    async function loadContentOnce() {
      try {
        const resp = await fetch('/content', { credentials: 'same-origin' });
        if (resp.ok) {
          const text = await resp.text();
          editor.value = text || '';
          renderPreviewIfNeeded();
        } else {
          console.warn('/content 返回非 200 状态', resp.status);
        }
      } catch (err) {
        console.warn('加载初始内容失败：', err);
      }
    }

    // WebSocket 连接与消息处理（尽量兼容原实现的消息格式）
    function setupWebSocket() {
      try {
        const scheme = (location.protocol === 'https:') ? 'wss' : 'ws';
        const wsUrl = scheme + '://' + location.host + '/ws';
        const ws = new WebSocket(wsUrl);

        ws.addEventListener('open', () => {
          console.info('WebSocket 已连接');
        });

        ws.addEventListener('message', (ev) => {
          // 假设消息为纯文本（最新全文）或 JSON {type, payload}
          try {
            let data = ev.data;
            // 如果是 JSON 格式并包含 payload.text，则采用其 text
            try {
              const parsed = JSON.parse(data);
              if (parsed && typeof parsed.payload === 'string') {
                data = parsed.payload;
              } else if (parsed && parsed.payload && typeof parsed.payload.text === 'string') {
                data = parsed.payload.text;
              }
            } catch (_) {
              // not json, keep raw
            }
            // 将接收到的文本放回 editor（保持行为和原项目一致：接收远端更新直接覆盖本地）
            const cur = editor.value;
            if (cur !== data) {
              editor.value = data;
              renderPreviewIfNeeded();
            }
          } catch (err) {
            console.error('处理 WebSocket 消息时出错：', err);
          }
        });

        ws.addEventListener('close', () => {
          console.info('WebSocket 断开，稍后重试连接');
          // 简单的自动重连策略（指数退避可另行增强）
          setTimeout(setupWebSocket, 2000);
        });

        ws.addEventListener('error', (e) => {
          console.warn('WebSocket 错误', e);
          ws.close();
        });
      } catch (err) {
        console.error('无法建立 WebSocket：', err);
      }
    }

    // 渲染预览（仅当 previewEnabled 为 true）
    function renderPreviewIfNeeded() {
      if (!previewEnabled) return;
      try {
        // marked -> HTML，再用 DOMPurify 清洗 -> 注入
        const raw = editor.value || '';
        // marked.parse for marked >=4, fallback to marked() for older
        const html = (typeof marked.parse === 'function') ? marked.parse(raw) : marked(raw);
        const clean = DOMPurify.sanitize(html);
        preview.innerHTML = clean;
      } catch (err) {
        preview.innerText = '渲染失败：' + err.message;
      }
    }

    // 切换预览按钮行为
    togglePreviewBtn.addEventListener('click', () => {
      previewEnabled = !previewEnabled;
      if (previewEnabled) {
        togglePreviewBtn.innerText = '隐藏 Markdown 预览';
        renderPreviewIfNeeded();
      } else {
        togglePreviewBtn.innerText = '显示 Markdown 预览';
        preview.innerText = '切换“显示 Markdown 预览”以查看渲染结果。';
      }
      // 为了无缝切换，不改变编辑区内容或焦点（用户体验优化）
    });

    // 编辑器输入时实时更新预览（防抖）
    let renderTimer = null;
    editor.addEventListener('input', () => {
      if (renderTimer) clearTimeout(renderTimer);
      renderTimer = setTimeout(() => {
        renderPreviewIfNeeded();
        renderTimer = null;
      }, 150); // 150ms 防抖
    });

    // 保存按钮：POST 到 /save（发送纯文本 body）
    saveBtn.addEventListener('click', async () => {
      try {
        saveBtn.disabled = true;
        const resp = await fetch('/save', {
          method: 'POST',
          credentials: 'same-origin',
          headers: { 'Content-Type': 'text/plain; charset=utf-8' },
          body: editor.value
        });
        if (!resp.ok) {
          alert('保存失败：' + resp.statusText);
        } else {
          // 若后端返回信息，可显示提示（保持兼容）
          // const txt = await resp.text();
          // console.info('保存结果', txt);
        }
      } catch (err) {
        alert('保存时出错：' + err.message);
      } finally {
        saveBtn.disabled = false;
      }
    });

    // 恢复按钮：POST /restore （后端会广播恢复后的内容）
    restoreBtn.addEventListener('click', async () => {
      if (!confirm('确定要从快照恢复？当前编辑器内容将被覆盖。')) return;
      try {
        restoreBtn.disabled = true;
        const resp = await fetch('/restore', {
          method: 'POST',
          credentials: 'same-origin'
        });
        if (!resp.ok) {
          alert('恢复失败：' + resp.statusText);
        } else {
          // 恢复成功后，后端应该会通过 WebSocket 广播新内容；但仍尝试获取 /content 以立即更新
          setTimeout(loadContentOnce, 300);
        }
      } catch (err) {
        alert('恢复时出错：' + err.message);
      } finally {
        restoreBtn.disabled = false;
      }
    });

    // 退出按钮：调用 /logout POST
    logoutBtn.addEventListener('click', async () => {
      try {
        const resp = await fetch('/logout', {
          method: 'POST',
          credentials: 'same-origin'
        });
        // 尝试跳回登录页（或根路径）
        location.reload();
      } catch (err) {
        console.warn('注销失败', err);
        location.reload();
      }
    });

    // 页面加载 -> 初始化内容 + websocket
    document.addEventListener('DOMContentLoaded', () => {
      loadContentOnce();
      setupWebSocket();
    });

    // 安全提示：若浏览器未加载 marked 或 DOMPurify，禁用预览
    if (typeof marked === 'undefined' || typeof DOMPurify === 'undefined') {
      togglePreviewBtn.disabled = true;
      togglePreviewBtn.title = '无法加载 Markdown 渲染库（marked / DOMPurify）。';
      console.warn('marked 或 DOMPurify 未加载，Markdown 预览不可用。');
    }
  })();
  </script>
</body>
</html>
`
