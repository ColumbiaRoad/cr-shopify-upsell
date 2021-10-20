var app=function(){"use strict";function n(){}function t(n){return n()}function e(){return Object.create(null)}function o(n){n.forEach(t)}function r(n){return"function"==typeof n}function c(n,t){return n!=n?t==t:n!==t||n&&"object"==typeof n||"function"==typeof n}function i(n,t){n.appendChild(t)}function s(n,t,e){n.insertBefore(t,e||null)}function a(n){n.parentNode.removeChild(n)}function d(n){return document.createElement(n)}function l(n){return document.createTextNode(n)}function u(){return l(" ")}function f(n,t,e,o){return n.addEventListener(t,e,o),()=>n.removeEventListener(t,e,o)}function p(n){return function(t){return t.preventDefault(),n.call(this,t)}}function h(n,t,e){null==e?n.removeAttribute(t):n.getAttribute(t)!==e&&n.setAttribute(t,e)}let m;function b(n){m=n}const $=[],g=[],v=[],x=[],y=Promise.resolve();let w=!1;function E(n){v.push(n)}let _=!1;const T=new Set;function k(){if(!_){_=!0;do{for(let n=0;n<$.length;n+=1){const t=$[n];b(t),A(t.$$)}for(b(null),$.length=0;g.length;)g.pop()();for(let n=0;n<v.length;n+=1){const t=v[n];T.has(t)||(T.add(t),t())}v.length=0}while($.length);for(;x.length;)x.pop()();w=!1,_=!1,T.clear()}}function A(n){if(null!==n.fragment){n.update(),o(n.before_update);const t=n.dirty;n.dirty=[-1],n.fragment&&n.fragment.p(n.ctx,t),n.after_update.forEach(E)}}const P=new Set;let I;function K(n,t){n&&n.i&&(P.delete(n),n.i(t))}function L(n,t,e,o){if(n&&n.o){if(P.has(n))return;P.add(n),I.c.push((()=>{P.delete(n),o&&(e&&n.d(1),o())})),n.o(t)}}function C(n){n&&n.c()}function O(n,e,c,i){const{fragment:s,on_mount:a,on_destroy:d,after_update:l}=n.$$;s&&s.m(e,c),i||E((()=>{const e=a.map(t).filter(r);d?d.push(...e):o(e),n.$$.on_mount=[]})),l.forEach(E)}function S(n,t){const e=n.$$;null!==e.fragment&&(o(e.on_destroy),e.fragment&&e.fragment.d(t),e.on_destroy=e.fragment=null,e.ctx=[])}function M(n,t){-1===n.$$.dirty[0]&&($.push(n),w||(w=!0,y.then(k)),n.$$.dirty.fill(0)),n.$$.dirty[t/31|0]|=1<<t%31}function R(t,r,c,i,s,d,l,u=[-1]){const f=m;b(t);const p=t.$$={fragment:null,ctx:null,props:d,update:n,not_equal:s,bound:e(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(r.context||(f?f.$$.context:[])),callbacks:e(),dirty:u,skip_bound:!1,root:r.target||f.$$.root};l&&l(p.root);let h=!1;if(p.ctx=c?c(t,r.props||{},((n,e,...o)=>{const r=o.length?o[0]:e;return p.ctx&&s(p.ctx[n],p.ctx[n]=r)&&(!p.skip_bound&&p.bound[n]&&p.bound[n](r),h&&M(t,n)),e})):[],p.update(),h=!0,o(p.before_update),p.fragment=!!i&&i(p.ctx),r.target){if(r.hydrate){const n=function(n){return Array.from(n.childNodes)}(r.target);p.fragment&&p.fragment.l(n),n.forEach(a)}else p.fragment&&p.fragment.c();r.intro&&K(t.$$.fragment),O(t,r.target,r.anchor,r.customElement),k()}b(f)}class j{$destroy(){S(this,1),this.$destroy=n}$on(n,t){const e=this.$$.callbacks[n]||(this.$$.callbacks[n]=[]);return e.push(t),()=>{const n=e.indexOf(t);-1!==n&&e.splice(n,1)}}$set(n){var t;this.$$set&&(t=n,0!==Object.keys(t).length)&&(this.$$.skip_bound=!0,this.$$set(n),this.$$.skip_bound=!1)}}function H(t){let e,o,r,c,l,m,b,$,g,v;return{c(){e=d("main"),o=d("div"),r=d("div"),c=d("div"),l=d("div"),l.innerHTML="<h3>Welcome</h3> \n\n                    <p>This app is ....</p> \n                    <p>TODO</p> \n                    <p>We need a subscription because ..</p>",m=u(),b=d("div"),$=d("a"),$.textContent="Create Subscription",h(l,"class","card-body"),h($,"href","#"),h($,"class","btn btn-primary"),h(b,"class","card-action"),h(c,"class","card"),h(r,"class","col-md-12"),h(o,"class","row"),h(e,"class","container")},m(n,a){s(n,e,a),i(e,o),i(o,r),i(r,c),i(c,l),i(c,m),i(c,b),i(b,$),g||(v=f($,"click",p(t[0])),g=!0)},p:n,i:n,o:n,d(n){n&&a(e),g=!1,v()}}}function N(n){const t=window["app-bridge"].actions.Loading,e=t.create(window.app),o=window["app-bridge"].actions.Redirect,r=o.create(window.app);return[()=>{console.log(e,t,t.Action.START),e.dispatch(t.Action.START),fetch(`/v1/shopify/billing/create?shop=${window.shop}`).then((n=>n.json())).then((n=>{r.dispatch(o.Action.REMOTE,n.return_url)})).catch((n=>console.error(n)))}]}class D extends j{constructor(n){super(),R(this,n,N,H,c,{})}}function z(n){let t;return{c(){t=l(" not")},m(n,e){s(n,t,e)},d(n){n&&a(t)}}}function W(n){let t;return{c(){t=l("Enable")},m(n,e){s(n,t,e)},d(n){n&&a(t)}}}function q(n){let t;return{c(){t=l("Disable")},m(n,e){s(n,t,e)},d(n){n&&a(t)}}}function B(t){let e,o,r,c,m,b,$,g,v,x,y,w,E,_,T,k,A,P,I,K,L,C,O,S=!t[0]&&z();function M(n,t){return n[0]?q:W}let R=M(t),j=R(t);return{c(){e=d("main"),o=d("div"),r=d("div"),r.innerHTML="<h3>Extension Status</h3>",c=u(),m=d("div"),b=d("div"),$=d("div"),g=d("div"),v=d("p"),x=l("The extension "),y=d("strong"),w=l("is"),S&&S.c(),E=l(" showing"),_=l(" in your checkout pages"),T=u(),k=d("div"),A=d("a"),j.c(),P=u(),I=d("div"),I.innerHTML='<div class="col-md-4 card-annotation"><h3>KPI 1</h3> \n            <p>KPI 1 desc</p></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><h3>KPI 1</h3> \n                    <p>KPI 1</p></div></div></div>',K=u(),L=d("div"),L.innerHTML='<div class="col-md-4 card-annotation"><h3>KPI 2</h3> \n            <p>KPI 2 desc</p></div> \n        <div class="col-md-8"><div class="card"><div class="card-body"><h3>KPI 2</h3> \n                    <p>KPI 2</p></div></div></div>',h(r,"class","col-md-4 card-annotation"),h(v,"class","svelte-1phfzfm"),h(A,"href","#"),h(A,"class","btn btn-primary"),h(g,"class","statusInnerContainer svelte-1phfzfm"),h($,"class","card-body"),h(b,"class","card"),h(m,"class","col-md-8"),h(o,"class","row no-gutters"),h(I,"class","row no-gutters"),h(L,"class","row no-gutters"),h(e,"class","container")},m(n,a){s(n,e,a),i(e,o),i(o,r),i(o,c),i(o,m),i(m,b),i(b,$),i($,g),i(g,v),i(v,x),i(v,y),i(y,w),S&&S.m(y,null),i(y,E),i(v,_),i(g,T),i(g,k),i(k,A),j.m(A,null),i(e,P),i(e,I),i(e,K),i(e,L),C||(O=f(A,"click",p(t[1])),C=!0)},p(n,[t]){n[0]?S&&(S.d(1),S=null):S||(S=z(),S.c(),S.m(y,E)),R!==(R=M(n))&&(j.d(1),j=R(n),j&&(j.c(),j.m(A,null)))},i:n,o:n,d(n){n&&a(e),S&&S.d(),j.d(),C=!1,O()}}}function J(n,t,e){let{extentionEnabled:o}=t;return n.$$set=n=>{"extentionEnabled"in n&&e(0,o=n.extentionEnabled)},[o,()=>{fetch("/v1/should-render",{method:"PATCH",body:JSON.stringify({shopURL:window.shop,shouldRender:!o}),headers:{"Content-type":"application/json"}}).then((n=>{console.log(n),e(0,o=!o)})).catch((n=>console.error(n)))}]}class U extends j{constructor(n){super(),R(this,n,J,B,c,{extentionEnabled:0})}}function F(t){let e,o;return e=new D({}),{c(){C(e.$$.fragment)},m(n,t){O(e,n,t),o=!0},p:n,i(n){o||(K(e.$$.fragment,n),o=!0)},o(n){L(e.$$.fragment,n),o=!1},d(n){S(e,n)}}}function G(n){let t,e;return t=new U({props:{extentionEnabled:n[1]}}),{c(){C(t.$$.fragment)},m(n,o){O(t,n,o),e=!0},p(n,e){const o={};2&e&&(o.extentionEnabled=n[1]),t.$set(o)},i(n){e||(K(t.$$.fragment,n),e=!0)},o(n){L(t.$$.fragment,n),e=!1},d(n){S(t,n)}}}function Q(n){let t,e,r,c;const i=[G,F],d=[];function u(n,t){return n[0]?0:1}return t=u(n),e=d[t]=i[t](n),{c(){e.c(),r=l("")},m(n,e){d[t].m(n,e),s(n,r,e),c=!0},p(n,[c]){let s=t;t=u(n),t===s?d[t].p(n,c):(I={r:0,c:[],p:I},L(d[s],1,1,(()=>{d[s]=null})),I.r||o(I.c),I=I.p,e=d[t],e?e.p(n,c):(e=d[t]=i[t](n),e.c()),K(e,1),e.m(r.parentNode,r))},i(n){c||(K(e),c=!0)},o(n){L(e),c=!1},d(n){d[t].d(n),n&&a(r)}}}function V(n,t,e){let{subscribed:o}=t,{extentionEnabled:r}=t;return n.$$set=n=>{"subscribed"in n&&e(0,o=n.subscribed),"extentionEnabled"in n&&e(1,r=n.extentionEnabled)},[o,r]}return new class extends j{constructor(n){super(),R(this,n,V,Q,c,{subscribed:0,extentionEnabled:1})}}({target:document.body,props:{subscribed:window.subscribed,extensionEnabled:window.extensionEnabled}})}();
//# sourceMappingURL=bundle.js.map
