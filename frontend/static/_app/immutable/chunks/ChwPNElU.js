import{H as G,I as J,J as ae,k as O,_ as ue,z as _e,a6 as re,D as A,Q as M,R as oe,L as ce,g as K,a7 as de,a8 as Q,P as R,F as x,T as he,a9 as C,aa as ne,A as se,ab as Ee,Y as fe,ac as ye,ad as q,ae as U,af as V,ag as ge,X as Ae,ah as ke,C as me,ai as Te,aj as pe,ak as Ie,h as W,q as Ne,al as we,N as xe,am as He,an as Ce,m as Le,ao as Se,ap as Me,y as Re,aq as De}from"./CtZ02a2i.js";import{a as be,t as Oe}from"./ZUxJqvzW.js";import"./rDCIrboE.js";let Z=!1;function le(){Z||(Z=!0,document.addEventListener("reset",e=>{Promise.resolve().then(()=>{var a;if(!e.defaultPrevented)for(const r of e.target.elements)(a=r.__on_r)==null||a.call(r)})},{capture:!0}))}function qe(e){var a=ae,r=O;G(null),J(null);try{return e()}finally{G(a),J(r)}}function Ve(e,a,r,f=r){e.addEventListener(a,()=>qe(r));const n=e.__on_r;n?e.__on_r=()=>{n(),f(!0)}:e.__on_r=()=>f(!0),le()}function Ke(e,a){return a}function Ye(e,a,r,f){for(var n=[],t=a.length,u=0;u<t;u++)ge(a[u].e,n,!0);var c=t>0&&n.length===0&&r!==null;if(c){var d=r.parentNode;Ae(d),d.append(r),f.clear(),p(e,a[0].prev,a[t-1].next)}ke(n,()=>{for(var k=0;k<t;k++){var o=a[k];c||(f.delete(o.k),p(e,o.prev,o.next)),me(o.e,!c)}})}function Qe(e,a,r,f,n,t=null){var u=e,c={flags:a,items:new Map,first:null},d=(a&re)!==0;if(d){var k=e;u=A?M(ce(k)):k.appendChild(ue())}A&&oe();var o=null,m=!1,I=Te(()=>{var s=r();return Ne(s)?s:s==null?[]:fe(s)});_e(()=>{var s=K(I),i=s.length;if(m&&i===0)return;m=i===0;let _=!1;if(A){var N=u.data===de;N!==(i===0)&&(u=Q(),M(u),R(!1),_=!0)}if(A){for(var h=null,E,y=0;y<i;y++){if(x.nodeType===8&&x.data===he){u=x,_=!0,R(!1);break}var g=s[y],l=f(g,y);E=te(x,c,h,null,g,l,y,n,a),c.items.set(l,E),h=E}i>0&&M(Q())}if(!A){var v=ae;Pe(s,c,u,n,a,(v.f&C)!==0,f)}t!==null&&(i===0?o?ne(o):o=se(()=>t(u)):o!==null&&Ee(o,()=>{o=null})),_&&R(!0),K(I)}),A&&(u=x)}function Pe(e,a,r,f,n,t,u,c){var P,B,X,z;var d=(n&pe)!==0,k=(n&(q|V))!==0,o=e.length,m=a.items,I=a.first,s=I,i,_=null,N,h=[],E=[],y,g,l,v;if(d)for(v=0;v<o;v+=1)y=e[v],g=u(y,v),l=m.get(g),l!==void 0&&((P=l.a)==null||P.measure(),(N??(N=new Set)).add(l));for(v=0;v<o;v+=1){if(y=e[v],g=u(y,v),l=m.get(g),l===void 0){var ie=s?s.e.nodes_start:r;_=te(ie,a,_,_===null?a.first:_.next,y,g,v,f,n),m.set(g,_),h=[],E=[],s=_.next;continue}if(k&&Be(l,y,v,n),l.e.f&C&&(ne(l.e),d&&((B=l.a)==null||B.unfix(),(N??(N=new Set)).delete(l))),l!==s){if(i!==void 0&&i.has(l)){if(h.length<E.length){var H=E[0],T;_=H.prev;var Y=h[0],L=h[h.length-1];for(T=0;T<h.length;T+=1)$(h[T],H,r);for(T=0;T<E.length;T+=1)i.delete(E[T]);p(a,Y.prev,L.next),p(a,_,Y),p(a,L,H),s=H,_=L,v-=1,h=[],E=[]}else i.delete(l),$(l,s,r),p(a,l.prev,l.next),p(a,l,_===null?a.first:_.next),p(a,_,l),_=l;continue}for(h=[],E=[];s!==null&&s.k!==g;)(t||!(s.e.f&C))&&(i??(i=new Set)).add(s),E.push(s),s=s.next;if(s===null)continue;l=s}h.push(l),_=l,s=l.next}if(s!==null||i!==void 0){for(var w=i===void 0?[]:fe(i);s!==null;)(t||!(s.e.f&C))&&w.push(s),s=s.next;var S=w.length;if(S>0){var ve=n&re&&o===0?r:null;if(d){for(v=0;v<S;v+=1)(X=w[v].a)==null||X.measure();for(v=0;v<S;v+=1)(z=w[v].a)==null||z.fix()}Ye(a,w,ve,m)}}d&&ye(()=>{var F;if(N!==void 0)for(l of N)(F=l.a)==null||F.apply()}),O.first=a.first&&a.first.e,O.last=_&&_.e}function Be(e,a,r,f){f&q&&U(e.v,a),f&V?U(e.i,r):e.i=r}function te(e,a,r,f,n,t,u,c,d,k){var o=(d&q)!==0,m=(d&we)===0,I=o?m?Ie(n):W(n):n,s=d&V?W(u):u,i={i:s,v:I,k:t,a:null,e:null,prev:r,next:f};try{return i.e=se(()=>c(e,I,s),A),i.e.prev=r&&r.e,i.e.next=f&&f.e,r===null?a.first=i:(r.next=i,r.e.next=i.e),f!==null&&(f.prev=i,f.e.prev=i.e),i}finally{}}function $(e,a,r){for(var f=e.next?e.next.e.nodes_start:r,n=a?a.e.nodes_start:r,t=e.e.nodes_start;t!==f;){var u=xe(t);n.before(t),t=u}}function p(e,a,r){a===null?e.first=r:(a.next=r,a.e.next=r&&r.e),r!==null&&(r.prev=a,r.e.prev=a&&a.e)}function Ue(e){if(A){var a=!1,r=()=>{if(!a){if(a=!0,e.hasAttribute("value")){var f=e.value;j(e,"value",null),e.value=f}if(e.hasAttribute("checked")){var n=e.checked;j(e,"checked",null),e.checked=n}}};e.__on_r=r,He(r),le()}}function j(e,a,r,f){var n=e.__attributes??(e.__attributes={});A&&(n[a]=e.getAttribute(a),a==="src"||a==="srcset"||a==="href"&&e.nodeName==="LINK")||n[a]!==(n[a]=r)&&(a==="style"&&"__styles"in e&&(e.__styles={}),a==="loading"&&(e[Ce]=r),r==null?e.removeAttribute(a):typeof r!="string"&&Xe(e).includes(a)?e[a]=r:e.setAttribute(a,r))}var ee=new Map;function Xe(e){var a=ee.get(e.nodeName);if(a)return a;ee.set(e.nodeName,a=[]);for(var r,f=e,n=Element.prototype;n!==f;){r=Se(f);for(var t in r)r[t].set&&a.push(t);f=Le(f)}return a}function We(e,a,r=a){var f=Me();Ve(e,"input",n=>{var t=n?e.defaultValue:e.value;if(t=D(e)?b(t):t,r(t),f&&t!==(t=a())){var u=e.selectionStart,c=e.selectionEnd;e.value=t??"",c!==null&&(e.selectionStart=u,e.selectionEnd=Math.min(c,e.value.length))}}),(A&&e.defaultValue!==e.value||Re(a)==null&&e.value)&&r(D(e)?b(e.value):e.value),De(()=>{var n=a();D(e)&&n===b(e.value)||e.type==="date"&&!n&&!e.value||n!==e.value&&(e.value=n??"")})}function D(e){var a=e.type;return a==="number"||a==="range"}function b(e){return e===""?null:+e}var ze=Oe('<header class="svelte-1k5yyk9"><div class="svelte-1k5yyk9"><h1 class="svelte-1k5yyk9">Header</h1></div></header>');function Ze(e){var a=ze();be(e,a)}export{Ze as H,We as b,Qe as e,Ke as i,Ue as r,j as s};
