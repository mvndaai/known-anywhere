var Bn=Array.isArray,ln=Array.prototype.indexOf,Un=Array.from,Vn=Object.defineProperty,Et=Object.getOwnPropertyDescriptor,sn=Object.getOwnPropertyDescriptors,Gn=Object.prototype,Kn=Array.prototype,an=Object.getPrototypeOf;const $n=()=>{};function Zn(t){return t()}function gt(t){for(var n=0;n<t.length;n++)t[n]()}const y=2,At=4,V=8,_t=16,I=32,G=64,Z=128,C=256,z=512,p=1024,x=2048,N=4096,F=8192,L=16384,un=32768,kt=65536,zn=1<<17,on=1<<19,It=1<<20,yt=Symbol("$state"),Jn=Symbol("legacy props"),Wn=Symbol("");function xt(t){return t===this.v}function fn(t,n){return t!=t?n==n:t!==n||t!==null&&typeof t=="object"||typeof t=="function"}function Dt(t){return!fn(t,this.v)}function _n(t){throw new Error("https://svelte.dev/e/effect_in_teardown")}function cn(){throw new Error("https://svelte.dev/e/effect_in_unowned_derived")}function vn(t){throw new Error("https://svelte.dev/e/effect_orphan")}function pn(){throw new Error("https://svelte.dev/e/effect_update_depth_exceeded")}function Xn(){throw new Error("https://svelte.dev/e/hydration_failed")}function Qn(t){throw new Error("https://svelte.dev/e/props_invalid_value")}function tr(){throw new Error("https://svelte.dev/e/state_descriptors_fixed")}function nr(){throw new Error("https://svelte.dev/e/state_prototype_fixed")}function hn(){throw new Error("https://svelte.dev/e/state_unsafe_local_read")}function dn(){throw new Error("https://svelte.dev/e/state_unsafe_mutation")}let nt=!1;function rr(){nt=!0}const er=1,lr=2,sr=4,ar=8,ur=16,or=1,ir=2,fr=4,_r=8,cr=16,vr=1,pr=2,En="[",yn="[!",wn="]",Rt={},hr=Symbol();function ct(t,n){var r={f:0,v:t,reactions:null,equals:xt,rv:0,wv:0};return r}function dr(t){return St(ct(t))}function Tn(t,n=!1){var e;const r=ct(t);return n||(r.equals=Dt),nt&&i!==null&&i.l!==null&&((e=i.l).s??(e.s=[])).push(r),r}function Er(t,n=!1){return St(Tn(t,n))}function St(t){return u!==null&&!k&&u.f&y&&(g===null?bn([t]):g.push(t)),t}function mn(t,n){return u!==null&&!k&&rt()&&u.f&(y|_t)&&(g===null||!g.includes(t))&&dn(),gn(t,n)}function gn(t,n){return t.equals(n)||(t.v,t.v=n,t.wv=Jt(),Ot(t,x),rt()&&f!==null&&f.f&p&&!(f.f&(I|G))&&(A===null?qn([t]):A.push(t))),n}function Ot(t,n){var r=t.reactions;if(r!==null)for(var e=rt(),l=r.length,s=0;s<l;s++){var a=r[s],o=a.f;o&x||!e&&a===f||(w(a,n),o&(p|C)&&(o&y?Ot(a,N):st(a)))}}function Ct(t){console.warn("https://svelte.dev/e/hydration_mismatch")}let S=!1;function yr(t){S=t}let T;function H(t){if(t===null)throw Ct(),Rt;return T=t}function wr(){return H(b(T))}function Tr(t){if(S){if(b(T)!==null)throw Ct(),Rt;T=t}}function mr(t=1){if(S){for(var n=t,r=T;n--;)r=b(r);T=r}}function gr(){for(var t=0,n=T;;){if(n.nodeType===8){var r=n.data;if(r===wn){if(t===0)return n;t-=1}else(r===En||r===yn)&&(t+=1)}var e=b(n);n.remove(),n=e}}var wt,An,Nt,bt;function Ar(){if(wt===void 0){wt=window,An=document;var t=Element.prototype,n=Node.prototype;Nt=Et(n,"firstChild").get,bt=Et(n,"nextSibling").get,t.__click=void 0,t.__className="",t.__attributes=null,t.__styles=null,t.__e=void 0,Text.prototype.__t=void 0}}function at(t=""){return document.createTextNode(t)}function ut(t){return Nt.call(t)}function b(t){return bt.call(t)}function kr(t,n){if(!S)return ut(t);var r=ut(T);if(r===null)r=T.appendChild(at());else if(n&&r.nodeType!==3){var e=at();return r==null||r.before(e),H(e),e}return H(r),r}function Ir(t,n){if(!S){var r=ut(t);return r instanceof Comment&&r.data===""?b(r):r}return T}function xr(t,n=1,r=!1){let e=S?T:t;for(var l;n--;)l=e,e=b(e);if(!S)return e;var s=e==null?void 0:e.nodeType;if(r&&s!==3){var a=at();return e===null?l==null||l.after(a):e.before(a),H(a),a}return H(e),e}function Dr(t){t.textContent=""}function qt(t){var n=y|x;f===null?n|=C:f.f|=It;var r=u!==null&&u.f&y?u:null;const e={children:null,ctx:i,deps:null,equals:xt,f:n,fn:t,reactions:null,rv:0,v:null,wv:0,parent:r??f};return r!==null&&(r.children??(r.children=[])).push(e),e}function Rr(t){const n=qt(t);return n.equals=Dt,n}function Pt(t){var n=t.children;if(n!==null){t.children=null;for(var r=0;r<n.length;r+=1){var e=n[r];e.f&y?vt(e):O(e)}}}function kn(t){for(var n=t.parent;n!==null;){if(!(n.f&y))return n;n=n.parent}return null}function Ft(t){var n,r=f;Q(kn(t));try{Pt(t),n=Xt(t)}finally{Q(r)}return n}function Lt(t){var n=Ft(t),r=(R||t.f&C)&&t.deps!==null?N:p;w(t,r),t.equals(n)||(t.v=n,t.wv=Jt())}function vt(t){Pt(t),U(t,0),w(t,L),t.v=t.children=t.deps=t.ctx=t.reactions=null}function Mt(t){f===null&&u===null&&vn(),u!==null&&u.f&C&&cn(),ht&&_n()}function In(t,n){var r=n.last;r===null?n.last=n.first=t:(r.next=t,t.prev=r,n.last=t)}function M(t,n,r,e=!0){var l=(t&G)!==0,s=f,a={ctx:i,deps:null,deriveds:null,nodes_start:null,nodes_end:null,f:t|x,first:null,fn:n,last:null,next:null,parent:l?null:s,prev:null,teardown:null,transitions:null,wv:0};if(r){var o=q;try{Tt(!0),lt(a),a.f|=un}catch(m){throw O(a),m}finally{Tt(o)}}else n!==null&&st(a);var _=r&&a.deps===null&&a.first===null&&a.nodes_start===null&&a.teardown===null&&(a.f&(It|Z))===0;if(!_&&!l&&e&&(s!==null&&In(a,s),u!==null&&u.f&y)){var c=u;(c.children??(c.children=[])).push(a)}return a}function Sr(t){const n=M(V,null,!1);return w(n,p),n.teardown=t,n}function Or(t){Mt();var n=f!==null&&(f.f&I)!==0&&i!==null&&!i.m;if(n){var r=i;(r.e??(r.e=[])).push({fn:t,effect:f,reaction:u})}else{var e=Yt(t);return e}}function Cr(t){return Mt(),pt(t)}function Nr(t){const n=M(G,t,!0);return(r={})=>new Promise(e=>{r.outro?Rn(n,()=>{O(n),e(void 0)}):(O(n),e(void 0))})}function Yt(t){return M(At,t,!1)}function br(t,n){var r=i,e={effect:null,ran:!1};r.l.r1.push(e),e.effect=pt(()=>{t(),!e.ran&&(e.ran=!0,mn(r.l.r2,!0),Hn(n))})}function qr(){var t=i;pt(()=>{if(en(t.l.r2)){for(var n of t.l.r1){var r=n.effect;r.f&p&&w(r,N),Y(r)&&lt(r),n.ran=!1}t.l.r2.v=!1}})}function pt(t){return M(V,t,!0)}function Pr(t,n=[],r=qt){const e=n.map(r);return xn(()=>t(...e.map(en)))}function xn(t,n=0){return M(V|_t|n,t,!0)}function Fr(t,n=!0){return M(V|I,t,!0,n)}function Ht(t){var n=t.teardown;if(n!==null){const r=ht,e=u;mt(!0),X(null);try{n.call(null)}finally{mt(r),X(e)}}}function jt(t){var n=t.deriveds;if(n!==null){t.deriveds=null;for(var r=0;r<n.length;r+=1)vt(n[r])}}function Bt(t,n=!1){var r=t.first;for(t.first=t.last=null;r!==null;){var e=r.next;O(r,n),r=e}}function Dn(t){for(var n=t.first;n!==null;){var r=n.next;n.f&I||O(n),n=r}}function O(t,n=!0){var r=!1;if((n||t.f&on)&&t.nodes_start!==null){for(var e=t.nodes_start,l=t.nodes_end;e!==null;){var s=e===l?null:b(e);e.remove(),e=s}r=!0}Bt(t,n&&!r),jt(t),U(t,0),w(t,L);var a=t.transitions;if(a!==null)for(const _ of a)_.stop();Ht(t);var o=t.parent;o!==null&&o.first!==null&&Ut(t),t.next=t.prev=t.teardown=t.ctx=t.deps=t.fn=t.nodes_start=t.nodes_end=null}function Ut(t){var n=t.parent,r=t.prev,e=t.next;r!==null&&(r.next=e),e!==null&&(e.prev=r),n!==null&&(n.first===t&&(n.first=e),n.last===t&&(n.last=r))}function Rn(t,n){var r=[];Vt(t,r,!0),Sn(r,()=>{O(t),n&&n()})}function Sn(t,n){var r=t.length;if(r>0){var e=()=>--r||n();for(var l of t)l.out(e)}else n()}function Vt(t,n,r){if(!(t.f&F)){if(t.f^=F,t.transitions!==null)for(const a of t.transitions)(a.is_global||r)&&n.push(a);for(var e=t.first;e!==null;){var l=e.next,s=(e.f&kt)!==0||(e.f&I)!==0;Vt(e,n,s?r:!1),e=l}}}function Lr(t){Gt(t,!0)}function Gt(t,n){if(t.f&F){t.f^=F,t.f&p||(t.f^=p),Y(t)&&(w(t,x),st(t));for(var r=t.first;r!==null;){var e=r.next,l=(r.f&kt)!==0||(r.f&I)!==0;Gt(r,l?n:!1),r=e}if(t.transitions!==null)for(const s of t.transitions)(s.is_global||n)&&s.in()}}const On=typeof requestIdleCallback>"u"?t=>setTimeout(t,1):requestIdleCallback;let J=!1,W=!1,ot=[],it=[];function Kt(){J=!1;const t=ot.slice();ot=[],gt(t)}function $t(){W=!1;const t=it.slice();it=[],gt(t)}function Mr(t){J||(J=!0,queueMicrotask(Kt)),ot.push(t)}function Yr(t){W||(W=!0,On($t)),it.push(t)}function Cn(){J&&Kt(),W&&$t()}const Zt=0,Nn=1;let K=!1,$=Zt,j=!1,B=null,q=!1,ht=!1;function Tt(t){q=t}function mt(t){ht=t}let D=[],P=0;let u=null,k=!1;function X(t){u=t}let f=null;function Q(t){f=t}let g=null;function bn(t){g=t}let h=null,E=0,A=null;function qn(t){A=t}let zt=1,tt=0,R=!1,i=null;function Jt(){return++zt}function rt(){return!nt||i!==null&&i.l===null}function Y(t){var c;var n=t.f;if(n&x)return!0;if(n&N){var r=t.deps,e=(n&C)!==0;if(r!==null){var l,s,a=(n&z)!==0,o=e&&f!==null&&!R,_=r.length;if(a||o){for(l=0;l<_;l++)s=r[l],(a||!((c=s==null?void 0:s.reactions)!=null&&c.includes(t)))&&(s.reactions??(s.reactions=[])).push(t);a&&(t.f^=z)}for(l=0;l<_;l++)if(s=r[l],Y(s)&&Lt(s),s.wv>t.wv)return!0}(!e||f!==null&&!R)&&w(t,p)}return!1}function Pn(t,n){for(var r=n;r!==null;){if(r.f&Z)try{r.fn(t);return}catch{r.f^=Z}r=r.parent}throw K=!1,t}function Fn(t){return(t.f&L)===0&&(t.parent===null||(t.parent.f&Z)===0)}function et(t,n,r,e){if(K){if(r===null&&(K=!1),Fn(n))throw t;return}r!==null&&(K=!0);{Pn(t,n);return}}function Wt(t,n,r=0){var e=t.reactions;if(e!==null)for(var l=0;l<e.length;l++){var s=e[l];s.f&y?Wt(s,n,r+1):n===s&&(r===0?w(s,x):s.f&p&&w(s,N),st(s))}}function Xt(t){var dt;var n=h,r=E,e=A,l=u,s=R,a=g,o=i,_=k,c=t.f;h=null,E=0,A=null,u=c&(I|G)?null:t,R=!q&&(c&C)!==0,g=null,i=t.ctx,k=!1,tt++;try{var m=(0,t.fn)(),v=t.deps;if(h!==null){var d;if(U(t,E),v!==null&&E>0)for(v.length=E+h.length,d=0;d<h.length;d++)v[E+d]=h[d];else t.deps=v=h;if(!R)for(d=E;d<v.length;d++)((dt=v[d]).reactions??(dt.reactions=[])).push(t)}else v!==null&&E<v.length&&(U(t,E),v.length=E);if(rt()&&A!==null&&!(t.f&(y|N|x)))for(d=0;d<A.length;d++)Wt(A[d],t);return l!==null&&tt++,m}finally{h=n,E=r,A=e,u=l,R=s,g=a,i=o,k=_}}function Ln(t,n){let r=n.reactions;if(r!==null){var e=ln.call(r,t);if(e!==-1){var l=r.length-1;l===0?r=n.reactions=null:(r[e]=r[l],r.pop())}}r===null&&n.f&y&&(h===null||!h.includes(n))&&(w(n,N),n.f&(C|z)||(n.f^=z),U(n,0))}function U(t,n){var r=t.deps;if(r!==null)for(var e=n;e<r.length;e++)Ln(t,r[e])}function lt(t){var n=t.f;if(!(n&L)){w(t,p);var r=f,e=i;f=t;try{n&_t?Dn(t):Bt(t),jt(t),Ht(t);var l=Xt(t);t.teardown=typeof l=="function"?l:null,t.wv=zt;var s=t.deps,a}catch(o){et(o,t,r,e||t.ctx)}finally{f=r}}}function Qt(){if(P>1e3){P=0;try{pn()}catch(t){if(B!==null)et(t,B,null);else throw t}}P++}function tn(t){var n=t.length;if(n!==0){Qt();var r=q;q=!0;try{for(var e=0;e<n;e++){var l=t[e];l.f&p||(l.f^=p);var s=[];nn(l,s),Mn(s)}}finally{q=r}}}function Mn(t){var n=t.length;if(n!==0)for(var r=0;r<n;r++){var e=t[r];if(!(e.f&(L|F)))try{Y(e)&&(lt(e),e.deps===null&&e.first===null&&e.nodes_start===null&&(e.teardown===null?Ut(e):e.fn=null))}catch(l){et(l,e,null,e.ctx)}}}function Yn(){if(j=!1,P>1001)return;const t=D;D=[],tn(t),j||(P=0,B=null)}function st(t){$===Zt&&(j||(j=!0,queueMicrotask(Yn))),B=t;for(var n=t;n.parent!==null;){n=n.parent;var r=n.f;if(r&(G|I)){if(!(r&p))return;n.f^=p}}D.push(n)}function nn(t,n){var r=t.first,e=[];t:for(;r!==null;){var l=r.f,s=(l&I)!==0,a=s&&(l&p)!==0,o=r.next;if(!a&&!(l&F))if(l&V){if(s)r.f^=p;else try{Y(r)&&lt(r)}catch(v){et(v,r,null,r.ctx)}var _=r.first;if(_!==null){r=_;continue}}else l&At&&e.push(r);if(o===null){let v=r.parent;for(;v!==null;){if(t===v)break t;var c=v.next;if(c!==null){r=c;continue t}v=v.parent}}r=o}for(var m=0;m<e.length;m++)_=e[m],n.push(_),nn(_,n)}function rn(t){var n=$,r=D;try{Qt();const l=[];$=Nn,D=l,j=!1,tn(r);var e=t==null?void 0:t();return Cn(),(D.length>0||l.length>0)&&rn(),P=0,B=null,e}finally{$=n,D=r}}async function Hr(){await Promise.resolve(),rn()}function en(t){var m;var n=t.f,r=(n&y)!==0;if(r&&n&L){var e=Ft(t);return vt(t),e}if(u!==null&&!k){g!==null&&g.includes(t)&&hn();var l=u.deps;t.rv<tt&&(t.rv=tt,h===null&&l!==null&&l[E]===t?E++:h===null?h=[t]:h.push(t))}else if(r&&t.deps===null)for(var s=t,a=s.parent,o=s;a!==null;)if(a.f&y){var _=a;o=_,a=_.parent}else{var c=a;(m=c.deriveds)!=null&&m.includes(o)||(c.deriveds??(c.deriveds=[])).push(o);break}return r&&(s=t,Y(s)&&Lt(s)),t.v}function Hn(t){var n=k;try{return k=!0,t()}finally{k=n}}const jn=-7169;function w(t,n){t.f=t.f&jn|n}function jr(t,n=!1,r){i={p:i,c:null,e:null,m:!1,s:t,x:null,l:null},nt&&!n&&(i.l={s:null,u:null,r1:[],r2:ct(!1)})}function Br(t){const n=i;if(n!==null){const a=n.e;if(a!==null){var r=f,e=u;n.e=null;try{for(var l=0;l<a.length;l++){var s=a[l];Q(s.effect),X(s.reaction),Yt(s.fn)}}finally{Q(r),X(e)}}i=n.p,n.m=!0}return{}}function Ur(t){if(!(typeof t!="object"||!t||t instanceof EventTarget)){if(yt in t)ft(t);else if(!Array.isArray(t))for(let n in t){const r=t[n];typeof r=="object"&&r&&yt in r&&ft(r)}}}function ft(t,n=new Set){if(typeof t=="object"&&t!==null&&!(t instanceof EventTarget)&&!n.has(t)){n.add(t),t instanceof Date&&t.getTime();for(let e in t)try{ft(t[e],n)}catch{}const r=an(t);if(r!==Object.prototype&&r!==Array.prototype&&r!==Map.prototype&&r!==Set.prototype&&r!==Date.prototype){const e=sn(r);for(let l in e){const s=e[l].get;if(s)try{s.call(t)}catch{}}}}}export{An as $,at as A,ut as B,f as C,pr as D,kt as E,H as F,wr as G,Rn as H,rn as I,Vn as J,Tn as K,Jn as L,Hr as M,Gn as N,Kn as O,ct as P,tr as Q,Et as R,yt as S,vr as T,hr as U,nr as V,an as W,Bn as X,yn as Y,gr as Z,yr as _,dr as a,Lr as a0,Yt as a1,pt as a2,Mr as a3,Qn as a4,zn as a5,fr as a6,Dt as a7,I as a8,G as a9,Sn as aA,ar as aB,ur as aC,Yr as aD,Wn as aE,sn as aF,rt as aG,br as aH,qr as aI,Er as aJ,mr as aK,fn as aL,Q as aa,or as ab,ir as ac,_r as ad,Rr as ae,cr as af,X as ag,u as ah,Sr as ai,on as aj,En as ak,b as al,Ar as am,Rt as an,wn as ao,Ct as ap,Xn as aq,Dr as ar,Un as as,Nr as at,sr as au,F as av,er as aw,gn as ax,lr as ay,Vt as az,mn as b,kr as c,Br as d,Or as e,Ir as f,en as g,i as h,Hn as i,Zn as j,gt as k,Ur as l,qt as m,rr as n,nt as o,jr as p,xn as q,Tr as r,xr as s,Pr as t,Cr as u,Fr as v,$n as w,O as x,S as y,T as z};
