import{a as V,t as U}from"../chunks/DlIn1gWU.js";import{z as re,A as te,B as me,e as Y,m as Te,r as A,K as ke,d as E,a4 as be,a5 as ne,J as q,I as J,t as C,L as Se,a6 as R,a7 as ge,n as we,a8 as Ne,P as Ee,a9 as se,aa as De,O as Ce,ab as Me,q as Le,D as Oe,R as Re,ac as He,ad as je,ae as z,af as F,ag as Ve,s as ie,ah as ye,i as qe,ai as Je,aj as Pe,F as Be,ak as Ke,al as We,h as Ye,am as Ge,an as Ue,k as ze,ao as Fe,v as Xe,ap as P,a1 as S,w as fe,x as Qe,c as B,a2 as L,a3 as M,a0 as oe}from"../chunks/Rts0-ZlK.js";import{d as Ze,s as le}from"../chunks/eG6FP_pp.js";import{p as O}from"../chunks/0HNuQjLK.js";import{o as $e}from"../chunks/CZuw5MZq.js";let ue=!1;function xe(){ue||(ue=!0,document.addEventListener("reset",e=>{Promise.resolve().then(()=>{var a;if(!e.defaultPrevented)for(const r of e.target.elements)(a=r.__on_r)==null||a.call(r)})},{capture:!0}))}function ea(e){var a=me,r=Y;re(null),te(null);try{return e()}finally{re(a),te(r)}}function aa(e,a,r,s=r){e.addEventListener(a,()=>ea(r));const t=e.__on_r;t?e.__on_r=()=>{t(),s(!0)}:e.__on_r=()=>s(!0),xe()}function ve(e,a){return a}function ra(e,a,r,s){for(var t=[],i=a.length,l=0;l<i;l++)De(a[l].e,t,!0);var _=i>0&&t.length===0&&r!==null;if(_){var d=r.parentNode;Ce(d),d.append(r),s.clear(),b(e,a[0].prev,a[i-1].next)}Me(t,()=>{for(var y=0;y<i;y++){var c=a[y];_||(s.delete(c.k),b(e,c.prev,c.next)),Le(c.e,!_)}})}function ce(e,a,r,s,t,i=null){var l=e,_={flags:a,items:new Map,first:null},d=(a&ye)!==0;if(d){var y=e;l=A?q(Oe(y)):y.appendChild(Re())}A&&ke();var c=null,g=!1,I=He(()=>{var n=r();return qe(n)?n:n==null?[]:Ee(n)});Te(()=>{var n=E(I),o=n.length;if(g&&o===0)return;g=o===0;let v=!1;if(A){var T=l.data===be;T!==(o===0)&&(l=ne(),q(l),J(!1),v=!0)}if(A){for(var h=null,p,m=0;m<o;m++){if(C.nodeType===8&&C.data===Se){l=C,v=!0,J(!1);break}var w=n[m],f=s(w,m);p=Ae(C,_,h,null,w,f,m,t,a),_.items.set(f,p),h=p}o>0&&q(ne())}if(!A){var u=me;ta(n,_,l,t,a,(u.f&R)!==0,s)}i!==null&&(o===0?c?ge(c):c=we(()=>i(l)):c!==null&&Ne(c,()=>{c=null})),v&&J(!0),E(I)}),A&&(l=C)}function ta(e,a,r,s,t,i,l,_){var Q,Z,$,ee;var d=(t&Je)!==0,y=(t&(z|F))!==0,c=e.length,g=a.items,I=a.first,n=I,o,v=null,T,h=[],p=[],m,w,f,u;if(d)for(u=0;u<c;u+=1)m=e[u],w=l(m,u),f=g.get(w),f!==void 0&&((Q=f.a)==null||Q.measure(),(T??(T=new Set)).add(f));for(u=0;u<c;u+=1){if(m=e[u],w=l(m,u),f=g.get(w),f===void 0){var N=n?n.e.nodes_start:r;v=Ae(N,a,v,v===null?a.first:v.next,m,w,u,s,t),g.set(w,v),h=[],p=[],n=v.next;continue}if(y&&na(f,m,u,t),f.e.f&R&&(ge(f.e),d&&((Z=f.a)==null||Z.unfix(),(T??(T=new Set)).delete(f))),f!==n){if(o!==void 0&&o.has(f)){if(h.length<p.length){var k=p[0],x;v=k.prev;var X=h[0],H=h[h.length-1];for(x=0;x<h.length;x+=1)_e(h[x],k,r);for(x=0;x<p.length;x+=1)o.delete(p[x]);b(a,X.prev,H.next),b(a,v,X),b(a,H,k),n=k,v=H,u-=1,h=[],p=[]}else o.delete(f),_e(f,n,r),b(a,f.prev,f.next),b(a,f,v===null?a.first:v.next),b(a,v,f),v=f;continue}for(h=[],p=[];n!==null&&n.k!==w;)(i||!(n.e.f&R))&&(o??(o=new Set)).add(n),p.push(n),n=n.next;if(n===null)continue;f=n}h.push(f),v=f,n=f.next}if(n!==null||o!==void 0){for(var D=o===void 0?[]:Ee(o);n!==null;)(i||!(n.e.f&R))&&D.push(n),n=n.next;var j=D.length;if(j>0){var Ie=t&ye&&c===0?r:null;if(d){for(u=0;u<j;u+=1)($=D[u].a)==null||$.measure();for(u=0;u<j;u+=1)(ee=D[u].a)==null||ee.fix()}ra(a,D,Ie,g)}}d&&je(()=>{var ae;if(T!==void 0)for(f of T)(ae=f.a)==null||ae.apply()}),Y.first=a.first&&a.first.e,Y.last=v&&v.e}function na(e,a,r,s){s&z&&se(e.v,a),s&F?se(e.i,r):e.i=r}function Ae(e,a,r,s,t,i,l,_,d,y){var c=(d&z)!==0,g=(d&Pe)===0,I=c?g?Ve(t):ie(t):t,n=d&F?ie(l):l,o={i:n,v:I,k:i,a:null,e:null,prev:r,next:s};try{return o.e=we(()=>_(e,I,n),A),o.e.prev=r&&r.e,o.e.next=s&&s.e,r===null?a.first=o:(r.next=o,r.e.next=o.e),s!==null&&(s.prev=o,s.e.prev=o.e),o}finally{}}function _e(e,a,r){for(var s=e.next?e.next.e.nodes_start:r,t=a?a.e.nodes_start:r,i=e.e.nodes_start;i!==s;){var l=Be(i);t.before(i),i=l}}function b(e,a,r){a===null?e.first=r:(a.next=r,a.e.next=r&&r.e),r!==null&&(r.prev=a,r.e.prev=a&&a.e)}function de(e){if(A){var a=!1,r=()=>{if(!a){if(a=!0,e.hasAttribute("value")){var s=e.value;G(e,"value",null),e.value=s}if(e.hasAttribute("checked")){var t=e.checked;G(e,"checked",null),e.checked=t}}};e.__on_r=r,Ke(r),xe()}}function G(e,a,r,s){var t=e.__attributes??(e.__attributes={});A&&(t[a]=e.getAttribute(a),a==="src"||a==="srcset"||a==="href"&&e.nodeName==="LINK")||t[a]!==(t[a]=r)&&(a==="style"&&"__styles"in e&&(e.__styles={}),a==="loading"&&(e[We]=r),r==null?e.removeAttribute(a):typeof r!="string"&&sa(e).includes(a)?e[a]=r:e.setAttribute(a,r))}var he=new Map;function sa(e){var a=he.get(e.nodeName);if(a)return a;he.set(e.nodeName,a=[]);for(var r,s=e,t=Element.prototype;t!==s;){r=Ge(s);for(var i in r)r[i].set&&a.push(i);s=Ye(s)}return a}function pe(e,a,r=a){var s=Ue();aa(e,"input",t=>{var i=t?e.defaultValue:e.value;if(i=K(e)?W(i):i,r(i),s&&i!==(i=a())){var l=e.selectionStart,_=e.selectionEnd;e.value=i??"",_!==null&&(e.selectionStart=l,e.selectionEnd=Math.min(_,e.value.length))}}),(A&&e.defaultValue!==e.value||ze(a)==null&&e.value)&&r(K(e)?W(e.value):e.value),Fe(()=>{var t=a();K(e)&&t===W(e.value)||e.type==="date"&&!t&&!e.value||t!==e.value&&(e.value=t??"")})}function K(e){var a=e.type;return a==="number"||a==="range"}function W(e){return e===""?null:+e}new TextEncoder;const ia=async(e,a,r,s,t)=>{let l=await(await fetch(`${a}/test/jwt`,{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({sub:E(r),username:E(s),exp:Math.floor(new Date().setDate(new Date().getDate()+5)/1e3)})})).json();t==null||t.setItem("subject",E(r)),t==null||t.setItem("username",E(s)),t==null||t.setItem("jwt",l.data)};var fa=U("<li><a> </a></li>"),oa=U("<p> </p> <ul></ul>",1),la=U('<h1>Welcome to SvelteKit</h1> <p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p> <h1>Routes</h1> <!> <div><span>Create JWT</span> <input type="text" placeholder="Subject"> <input type="text" placeholder="Username"> <button>Generate</button></div>',1);function ha(e,a){Xe(a,!0);const r="http://localhost:8080";let s=P(O({}));$e(async function(){let o=await(await fetch(`${r}/api`)).json();B(s,O(o.data))});const t=typeof window<"u"?window.localStorage:null;let i=P(O((t==null?void 0:t.getItem("subject"))||"")),l=P(O((t==null?void 0:t.getItem("username"))||""));var _=la(),d=S(fe(_),6);ce(d,17,()=>Object.entries(E(s)),ve,(n,o)=>{let v=()=>E(o)[0],T=()=>E(o)[1];var h=oa(),p=fe(h),m=L(p,!0);M(p);var w=S(p,2);ce(w,21,T,ve,(f,u)=>{var N=fa(),k=L(N),x=L(k,!0);M(k),M(N),oe(()=>{G(k,"href",E(u)),le(x,E(u))}),V(f,N)}),M(w),oe(()=>le(m,v())),V(n,h)});var y=S(d,2),c=S(L(y),2);de(c);var g=S(c,2);de(g);var I=S(g,2);I.__click=[ia,r,i,l,t],M(y),pe(c,()=>E(i),n=>B(i,n)),pe(g,()=>E(l),n=>B(l,n)),V(e,_),Qe()}Ze(["click"]);export{ha as component};
