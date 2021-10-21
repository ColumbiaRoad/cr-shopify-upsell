var app=function(){"use strict";function n(){}function t(n){return n()}function e(){return Object.create(null)}function o(n){n.forEach(t)}function c(n){return"function"==typeof n}function i(n,t){return n!=n?t==t:n!==t||n&&"object"==typeof n||"function"==typeof n}function r(n,t){n.appendChild(t)}function s(n,t,e){n.insertBefore(t,e||null)}function a(n){n.parentNode.removeChild(n)}function d(n){return document.createElement(n)}function l(n){return document.createTextNode(n)}function u(){return l(" ")}function p(n,t,e,o){return n.addEventListener(t,e,o),()=>n.removeEventListener(t,e,o)}function f(n){return function(t){return t.preventDefault(),n.call(this,t)}}function h(n,t,e){null==e?n.removeAttribute(t):n.getAttribute(t)!==e&&n.setAttribute(t,e)}let m;function b(n){m=n}const $=[],g=[],v=[],w=[],y=Promise.resolve();let x=!1;function E(n){v.push(n)}let _=!1;const T=new Set;function A(){if(!_){_=!0;do{for(let n=0;n<$.length;n+=1){const t=$[n];b(t),k(t.$$)}for(b(null),$.length=0;g.length;)g.pop()();for(let n=0;n<v.length;n+=1){const t=v[n];T.has(t)||(T.add(t),t())}v.length=0}while($.length);for(;w.length;)w.pop()();x=!1,_=!1,T.clear()}}function k(n){if(null!==n.fragment){n.update(),o(n.before_update);const t=n.dirty;n.dirty=[-1],n.fragment&&n.fragment.p(n.ctx,t),n.after_update.forEach(E)}}const P=new Set;let O;function S(n,t){n&&n.i&&(P.delete(n),n.i(t))}function C(n,t,e,o){if(n&&n.o){if(P.has(n))return;P.add(n),O.c.push((()=>{P.delete(n),o&&(e&&n.d(1),o())})),n.o(t)}}function I(n){n&&n.c()}function L(n,e,i,r){const{fragment:s,on_mount:a,on_destroy:d,after_update:l}=n.$$;s&&s.m(e,i),r||E((()=>{const e=a.map(t).filter(c);d?d.push(...e):o(e),n.$$.on_mount=[]})),l.forEach(E)}function K(n,t){const e=n.$$;null!==e.fragment&&(o(e.on_destroy),e.fragment&&e.fragment.d(t),e.on_destroy=e.fragment=null,e.ctx=[])}function H(n,t){-1===n.$$.dirty[0]&&($.push(n),x||(x=!0,y.then(A)),n.$$.dirty.fill(0)),n.$$.dirty[t/31|0]|=1<<t%31}function M(t,c,i,r,s,d,l,u=[-1]){const p=m;b(t);const f=t.$$={fragment:null,ctx:null,props:d,update:n,not_equal:s,bound:e(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(c.context||(p?p.$$.context:[])),callbacks:e(),dirty:u,skip_bound:!1,root:c.target||p.$$.root};l&&l(f.root);let h=!1;if(f.ctx=i?i(t,c.props||{},((n,e,...o)=>{const c=o.length?o[0]:e;return f.ctx&&s(f.ctx[n],f.ctx[n]=c)&&(!f.skip_bound&&f.bound[n]&&f.bound[n](c),h&&H(t,n)),e})):[],f.update(),h=!0,o(f.before_update),f.fragment=!!r&&r(f.ctx),c.target){if(c.hydrate){const n=function(n){return Array.from(n.childNodes)}(c.target);f.fragment&&f.fragment.l(n),n.forEach(a)}else f.fragment&&f.fragment.c();c.intro&&S(t.$$.fragment),L(t,c.target,c.anchor,c.customElement),A()}b(p)}class R{$destroy(){K(this,1),this.$destroy=n}$on(n,t){const e=this.$$.callbacks[n]||(this.$$.callbacks[n]=[]);return e.push(t),()=>{const n=e.indexOf(t);-1!==n&&e.splice(n,1)}}$set(n){var t;this.$$set&&(t=n,0!==Object.keys(t).length)&&(this.$$.skip_bound=!0,this.$$set(n),this.$$.skip_bound=!1)}}function j(t){let e,o,c,i,l,m,b,$,g,v;return{c(){e=d("main"),o=d("div"),c=d("div"),i=d("div"),l=d("div"),l.innerHTML="<h3>Welcome</h3> \n\n                    <p>This app is ....</p> \n                    <p>TODO</p> \n                    <p>We need a subscription because ..</p>",m=u(),b=d("div"),$=d("a"),$.textContent="Create Subscription",h(l,"class","card-body"),h($,"href","#"),h($,"class","btn btn-primary"),h(b,"class","card-action"),h(i,"class","card"),h(c,"class","col-md-12"),h(o,"class","row"),h(e,"class","container")},m(n,a){s(n,e,a),r(e,o),r(o,c),r(c,i),r(i,l),r(i,m),r(i,b),r(b,$),g||(v=p($,"click",f(t[0])),g=!0)},p:n,i:n,o:n,d(n){n&&a(e),g=!1,v()}}}function N(n){const t=window["app-bridge"].actions.Loading,e=t.create(window.app),o=window["app-bridge"].actions.Redirect,c=o.create(window.app),i=window["app-bridge"].actions.Toast,r=i.create(window.app,{message:"Couldn't redirect :(",duration:5e3});return[()=>{e.dispatch(t.Action.START),fetch(`/v1/shopify/billing/create?shop=${window.shop}`).then((n=>n.json())).then((n=>{c.dispatch(o.Action.REMOTE,n.return_url)})).catch((()=>r.dispatch(i.Action.SHOW))).finally((()=>e.dispatch(t.Action.STOP)))}]}class W extends R{constructor(n){super(),M(this,n,N,j,i,{})}}function D(n){let t;return{c(){t=l(" not")},m(n,e){s(n,t,e)},d(n){n&&a(t)}}}function z(n){let t;return{c(){t=l("Enable")},m(n,e){s(n,t,e)},d(n){n&&a(t)}}}function q(n){let t;return{c(){t=l("Disable")},m(n,e){s(n,t,e)},d(n){n&&a(t)}}}function B(t){let e,o,c,i,m,b,$,g,v,w,y,x,E,_,T,A,k,P,O,S,C,I,L,K=!t[0]&&D();function H(n,t){return n[0]?q:z}let M=H(t),R=M(t);return{c(){e=d("main"),o=d("div"),c=d("div"),c.innerHTML="<h3>Extension Status</h3>",i=u(),m=d("div"),b=d("div"),$=d("div"),g=d("div"),v=d("p"),w=l("The extension "),y=d("strong"),x=l("is"),K&&K.c(),E=l(" showing"),_=l(" in your checkout pages"),T=u(),A=d("div"),k=d("a"),R.c(),P=u(),O=d("div"),O.innerHTML='<div class="col-md-4 card-annotation"><h3>KPI 1</h3> \n            <p>KPI 1 desc</p></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><h3>KPI 1</h3> \n                    <p>KPI 1</p></div></div></div>',S=u(),C=d("div"),C.innerHTML='<div class="col-md-4 card-annotation"><h3>KPI 2</h3> \n            <p>KPI 2 desc</p></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><h3>KPI 2</h3> \n                    <p>KPI 2</p></div></div></div>',h(c,"class","col-md-4 card-annotation"),h(v,"class","svelte-1phfzfm"),h(k,"href","#"),h(k,"class","btn btn-primary"),h(g,"class","statusInnerContainer svelte-1phfzfm"),h($,"class","card-body"),h(b,"class","card"),h(m,"class","col-md-8"),h(o,"class","row no-gutters"),h(O,"class","row no-gutters"),h(C,"class","row no-gutters"),h(e,"class","container")},m(n,a){s(n,e,a),r(e,o),r(o,c),r(o,i),r(o,m),r(m,b),r(b,$),r($,g),r(g,v),r(v,w),r(v,y),r(y,x),K&&K.m(y,null),r(y,E),r(v,_),r(g,T),r(g,A),r(A,k),R.m(k,null),r(e,P),r(e,O),r(e,S),r(e,C),I||(L=p(k,"click",f(t[1])),I=!0)},p(n,[t]){n[0]?K&&(K.d(1),K=null):K||(K=D(),K.c(),K.m(y,E)),M!==(M=H(n))&&(R.d(1),R=M(n),R&&(R.c(),R.m(k,null)))},i:n,o:n,d(n){n&&a(e),K&&K.d(),R.d(),I=!1,L()}}}function J(n,t,e){let{extensionEnabled:o}=t;const c=window["app-bridge"].actions.Loading,i=c.create(window.app),r=window["app-bridge"].actions.Toast,s=r.create(window.app,{message:"Couldn't change the state :(",duration:5e3});return n.$$set=n=>{"extensionEnabled"in n&&e(0,o=n.extensionEnabled)},[o,()=>{i.dispatch(c.Action.START),fetch("/v1/should-render",{method:"PATCH",body:JSON.stringify({shopURL:window.shop,shouldRender:!o}),headers:{"Content-type":"application/json"}}).then((n=>{if(!n.ok)throw new Error;e(0,o=!o)})).catch((()=>s.dispatch(r.Action.SHOW))).finally((()=>i.dispatch(c.Action.STOP)))}]}class U extends R{constructor(n){super(),M(this,n,J,B,i,{extensionEnabled:0})}}function F(t){let e,o;return e=new W({}),{c(){I(e.$$.fragment)},m(n,t){L(e,n,t),o=!0},p:n,i(n){o||(S(e.$$.fragment,n),o=!0)},o(n){C(e.$$.fragment,n),o=!1},d(n){K(e,n)}}}function G(n){let t,e;return t=new U({props:{extensionEnabled:n[1]}}),{c(){I(t.$$.fragment)},m(n,o){L(t,n,o),e=!0},p(n,e){const o={};2&e&&(o.extensionEnabled=n[1]),t.$set(o)},i(n){e||(S(t.$$.fragment,n),e=!0)},o(n){C(t.$$.fragment,n),e=!1},d(n){K(t,n)}}}function Q(n){let t,e,c,i;const r=[G,F],d=[];function u(n,t){return n[0]?0:1}return t=u(n),e=d[t]=r[t](n),{c(){e.c(),c=l("")},m(n,e){d[t].m(n,e),s(n,c,e),i=!0},p(n,[i]){let s=t;t=u(n),t===s?d[t].p(n,i):(O={r:0,c:[],p:O},C(d[s],1,1,(()=>{d[s]=null})),O.r||o(O.c),O=O.p,e=d[t],e?e.p(n,i):(e=d[t]=r[t](n),e.c()),S(e,1),e.m(c.parentNode,c))},i(n){i||(S(e),i=!0)},o(n){C(e),i=!1},d(n){d[t].d(n),n&&a(c)}}}function V(n,t,e){let{subscribed:o}=t,{extensionEnabled:c}=t;return n.$$set=n=>{"subscribed"in n&&e(0,o=n.subscribed),"extensionEnabled"in n&&e(1,c=n.extensionEnabled)},[o,c]}return new class extends R{constructor(n){super(),M(this,n,V,Q,i,{subscribed:0,extensionEnabled:1})}}({target:document.body,props:{subscribed:window.subscribed,extensionEnabled:window.extensionEnabled}})}();
//# sourceMappingURL=bundle.js.map
