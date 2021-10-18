var app=function(){"use strict";function t(){}function n(t){return t()}function e(){return Object.create(null)}function o(t){t.forEach(n)}function c(t){return"function"==typeof t}function r(t,n){return t!=t?n==n:t!==n||t&&"object"==typeof t||"function"==typeof t}function s(t,n){t.appendChild(n)}function i(t,n,e){t.insertBefore(n,e||null)}function a(t){t.parentNode.removeChild(t)}function d(t){return document.createElement(t)}function u(t){return document.createTextNode(t)}function l(t,n,e){null==e?t.removeAttribute(n):t.getAttribute(n)!==e&&t.setAttribute(n,e)}let f;function p(t){f=t}const h=[],m=[],$=[],v=[],g=Promise.resolve();let b=!1;function y(t){$.push(t)}let w=!1;const x=new Set;function _(){if(!w){w=!0;do{for(let t=0;t<h.length;t+=1){const n=h[t];p(n),E(n.$$)}for(p(null),h.length=0;m.length;)m.pop()();for(let t=0;t<$.length;t+=1){const n=$[t];x.has(n)||(x.add(n),n())}$.length=0}while(h.length);for(;v.length;)v.pop()();b=!1,w=!1,x.clear()}}function E(t){if(null!==t.fragment){t.update(),o(t.before_update);const n=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,n),t.after_update.forEach(y)}}const T=new Set;let k;function A(t,n){t&&t.i&&(T.delete(t),t.i(n))}function I(t,n,e,o){if(t&&t.o){if(T.has(t))return;T.add(t),k.c.push((()=>{T.delete(t),o&&(e&&t.d(1),o())})),t.o(n)}}function P(t){t&&t.c()}function K(t,e,r,s){const{fragment:i,on_mount:a,on_destroy:d,after_update:u}=t.$$;i&&i.m(e,r),s||y((()=>{const e=a.map(n).filter(c);d?d.push(...e):o(e),t.$$.on_mount=[]})),u.forEach(y)}function O(t,n){const e=t.$$;null!==e.fragment&&(o(e.on_destroy),e.fragment&&e.fragment.d(n),e.on_destroy=e.fragment=null,e.ctx=[])}function S(t,n){-1===t.$$.dirty[0]&&(h.push(t),b||(b=!0,g.then(_)),t.$$.dirty.fill(0)),t.$$.dirty[n/31|0]|=1<<n%31}function C(n,c,r,s,i,d,u,l=[-1]){const h=f;p(n);const m=n.$$={fragment:null,ctx:null,props:d,update:t,not_equal:i,bound:e(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(c.context||(h?h.$$.context:[])),callbacks:e(),dirty:l,skip_bound:!1,root:c.target||h.$$.root};u&&u(m.root);let $=!1;if(m.ctx=r?r(n,c.props||{},((t,e,...o)=>{const c=o.length?o[0]:e;return m.ctx&&i(m.ctx[t],m.ctx[t]=c)&&(!m.skip_bound&&m.bound[t]&&m.bound[t](c),$&&S(n,t)),e})):[],m.update(),$=!0,o(m.before_update),m.fragment=!!s&&s(m.ctx),c.target){if(c.hydrate){const t=function(t){return Array.from(t.childNodes)}(c.target);m.fragment&&m.fragment.l(t),t.forEach(a)}else m.fragment&&m.fragment.c();c.intro&&A(n.$$.fragment),K(n,c.target,c.anchor,c.customElement),_()}p(h)}class L{$destroy(){O(this,1),this.$destroy=t}$on(t,n){const e=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return e.push(n),()=>{const t=e.indexOf(n);-1!==t&&e.splice(t,1)}}$set(t){var n;this.$$set&&(n=t,0!==Object.keys(n).length)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}function M(n){let e,o,c,r,f,p,h,m,$,v;return{c(){e=d("main"),o=d("div"),c=d("div"),r=d("div"),f=d("div"),f.innerHTML="<h3>Welcome</h3> \n\n                    <p>This app is ....</p> \n                    <p>TODO</p> \n                    <p>We need a subscription because ..</p>",p=u(" "),h=d("div"),m=d("a"),m.textContent="Create Subscription",l(f,"class","card-body"),l(m,"href","#"),l(m,"class","btn btn-primary"),l(h,"class","card-action"),l(r,"class","card"),l(c,"class","col-md-12"),l(o,"class","row"),l(e,"class","container")},m(t,a){var d,u,l,g,b;i(t,e,a),s(e,o),s(o,c),s(c,r),s(r,f),s(r,p),s(r,h),s(h,m),$||(d=m,u="click",b=n[0],l=function(t){return t.preventDefault(),b.call(this,t)},d.addEventListener(u,l,g),v=()=>d.removeEventListener(u,l,g),$=!0)},p:t,i:t,o:t,d(t){t&&a(e),$=!1,v()}}}function N(t){const n=window["app-bridge"].actions.Loading,e=n.create(window.app),o=window["app-bridge"].actions.Redirect,c=o.create(window.app);return[()=>{console.log(e,n,n.Action.START),e.dispatch(n.Action.START),c.dispatch(o.Action.REMOTE,"http://example.com")}]}class R extends L{constructor(t){super(),C(this,t,N,M,r,{})}}function j(n){let e;return{c(){e=d("main"),e.innerHTML='<div class="row no-gutters"><div class="col-md-4 card-annotation"><h3>Extension Status</h3></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><div class="statusInnerContainer svelte-1p7lfiq"><p class="svelte-1p7lfiq">The extension <strong>is showing</strong> in your checkout\n                            pages</p> \n                        <div><a href="#" class="btn btn-primary">Disable</a></div></div></div></div></div></div> \n    <div class="row no-gutters"><div class="col-md-4 card-annotation"><h3>KPI 1</h3> \n            <p>KPI 1 desc</p></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><h3>KPI 1</h3> \n                    <p>KPI 1</p></div></div></div></div> \n    <div class="row no-gutters"><div class="col-md-4 card-annotation"><h3>KPI 2</h3> \n            <p>KPI 2 desc</p></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><h3>KPI 2</h3> \n                    <p>KPI 2</p></div></div></div></div>',l(e,"class","container")},m(t,n){i(t,e,n)},p:t,i:t,o:t,d(t){t&&a(e)}}}class q extends L{constructor(t){super(),C(this,t,null,j,r,{})}}function D(t){let n,e;return n=new R({}),{c(){P(n.$$.fragment)},m(t,o){K(n,t,o),e=!0},i(t){e||(A(n.$$.fragment,t),e=!0)},o(t){I(n.$$.fragment,t),e=!1},d(t){O(n,t)}}}function H(t){let n,e;return n=new q({}),{c(){P(n.$$.fragment)},m(t,o){K(n,t,o),e=!0},i(t){e||(A(n.$$.fragment,t),e=!0)},o(t){I(n.$$.fragment,t),e=!1},d(t){O(n,t)}}}function W(t){let n,e,c,r;const s=[H,D],d=[];function l(t,n){return t[0]?0:1}return n=l(t),e=d[n]=s[n](t),{c(){e.c(),c=u("")},m(t,e){d[n].m(t,e),i(t,c,e),r=!0},p(t,[r]){let i=n;n=l(t),n!==i&&(k={r:0,c:[],p:k},I(d[i],1,1,(()=>{d[i]=null})),k.r||o(k.c),k=k.p,e=d[n],e||(e=d[n]=s[n](t),e.c()),A(e,1),e.m(c.parentNode,c))},i(t){r||(A(e),r=!0)},o(t){I(e),r=!1},d(t){d[n].d(t),t&&a(c)}}}function B(t,n,e){let{subscribed:o}=n;return t.$$set=t=>{"subscribed"in t&&e(0,o=t.subscribed)},[o]}return new class extends L{constructor(t){super(),C(this,t,B,W,r,{subscribed:0})}}({target:document.body,props:{subscribed:window.subscribed}})}();
//# sourceMappingURL=bundle.js.map
