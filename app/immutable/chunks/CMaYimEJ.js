import{S as R,N as M,O as Q,P as m,Q as X,b as P,R as N,U as y,g as E,C as x,V as J,W as p,X as ee,q as re,y as q,G as te,E as ne,Y as ae,Z as ie,F as se,_ as Y,a0 as j,v as k,H as $,z as fe,a1 as ue,a2 as le,i as B,a3 as ce,a4 as de,a5 as oe,a6 as _e,a7 as ve,a8 as he,a9 as ye,aa as G,ab as be,o as ge,ac as Pe,ad as Ee,L as me,m as H,ae as Re,af as we,K as Se}from"./DWxV8Muf.js";function I(r,s=null,_){if(typeof r!="object"||r===null||R in r)return r;const v=p(r);if(v!==M&&v!==Q)return r;var a=new Map,l=ee(r),d=m(0);l&&a.set("length",m(r.length));var c;return new Proxy(r,{defineProperty(f,e,t){(!("value"in t)||t.configurable===!1||t.enumerable===!1||t.writable===!1)&&X();var i=a.get(e);return i===void 0?(i=m(t.value),a.set(e,i)):P(i,I(t.value,c)),!0},deleteProperty(f,e){var t=a.get(e);if(t===void 0)e in f&&a.set(e,m(y));else{if(l&&typeof e=="string"){var i=a.get("length"),n=Number(e);Number.isInteger(n)&&n<i.v&&P(i,n)}P(t,y),K(d)}return!0},get(f,e,t){var o;if(e===R)return r;var i=a.get(e),n=e in f;if(i===void 0&&(!n||(o=N(f,e))!=null&&o.writable)&&(i=m(I(n?f[e]:y,c)),a.set(e,i)),i!==void 0){var u=E(i);return u===y?void 0:u}return Reflect.get(f,e,t)},getOwnPropertyDescriptor(f,e){var t=Reflect.getOwnPropertyDescriptor(f,e);if(t&&"value"in t){var i=a.get(e);i&&(t.value=E(i))}else if(t===void 0){var n=a.get(e),u=n==null?void 0:n.v;if(n!==void 0&&u!==y)return{enumerable:!0,configurable:!0,value:u,writable:!0}}return t},has(f,e){var u;if(e===R)return!0;var t=a.get(e),i=t!==void 0&&t.v!==y||Reflect.has(f,e);if(t!==void 0||x!==null&&(!i||(u=N(f,e))!=null&&u.writable)){t===void 0&&(t=m(i?I(f[e],c):y),a.set(e,t));var n=E(t);if(n===y)return!1}return i},set(f,e,t,i){var w;var n=a.get(e),u=e in f;if(l&&e==="length")for(var o=t;o<n.v;o+=1){var b=a.get(o+"");b!==void 0?P(b,y):o in f&&(b=m(y),a.set(o+"",b))}n===void 0?(!u||(w=N(f,e))!=null&&w.writable)&&(n=m(void 0),P(n,I(t,c)),a.set(e,n)):(u=n.v!==y,P(n,I(t,c)));var g=Reflect.getOwnPropertyDescriptor(f,e);if(g!=null&&g.set&&g.set.call(i,t),!u){if(l&&typeof e=="string"){var O=a.get("length"),A=Number(e);Number.isInteger(A)&&A>=O.v&&P(O,A+1)}K(d)}return!0},ownKeys(f){E(d);var e=Reflect.ownKeys(f).filter(n=>{var u=a.get(n);return u===void 0||u.v!==y});for(var[t,i]of a)i.v!==y&&!(t in f)&&e.push(t);return e},setPrototypeOf(){J()}})}function K(r,s=1){P(r,r.v+s)}function Z(r){return r!==null&&typeof r=="object"&&R in r?r[R]:r}function Le(r,s){return Object.is(Z(r),Z(s))}function Ne(r,s,_=!1){q&&te();var v=r,a=null,l=null,d=y,c=_?ne:0,f=!1;const e=(i,n=!0)=>{f=!0,t(n,i)},t=(i,n)=>{if(d===(d=i))return;let u=!1;if(q){const o=v.data===ae;!!d===o&&(v=ie(),se(v),Y(!1),u=!0)}d?(a?j(a):n&&(a=k(()=>n(v))),l&&$(l,()=>{l=null})):(l?j(l):n&&(l=k(()=>n(v))),a&&$(a,()=>{a=null})),u&&Y(!0)};re(()=>{f=!1,s(e),f||t(null,null)},c),q&&(v=fe)}function z(r,s){return r===s||(r==null?void 0:r[R])===s}function Ce(r={},s,_,v){return ue(()=>{var a,l;return le(()=>{a=l,l=[],B(()=>{r!==_(...l)&&(s(r,...l),a&&z(_(...a),r)&&s(null,...a))})}),()=>{ce(()=>{l&&z(_(...l),r)&&s(null,...l)})}}),r}let L=!1;function Ie(r){var s=L;try{return L=!1,[r(),L]}finally{L=s}}function V(r){for(var s=x,_=x;s!==null&&!(s.f&(he|ye));)s=s.parent;try{return G(s),r()}finally{G(_)}}function De(r,s,_,v){var U;var a=(_&be)!==0,l=!ge||(_&Pe)!==0,d=(_&Ee)!==0,c=(_&we)!==0,f=!1,e;d?[e,f]=Ie(()=>r[s]):e=r[s];var t=R in r||me in r,i=d&&(((U=N(r,s))==null?void 0:U.set)??(t&&s in r&&(h=>r[s]=h)))||void 0,n=v,u=!0,o=!1,b=()=>(o=!0,u&&(u=!1,c?n=B(v):n=v),n);e===void 0&&v!==void 0&&(i&&l&&de(),e=b(),i&&i(e));var g;if(l)g=()=>{var h=r[s];return h===void 0?b():(u=!0,o=!1,h)};else{var O=V(()=>(a?H:Re)(()=>r[s]));O.f|=oe,g=()=>{var h=E(O);return h!==void 0&&(n=void 0),h===void 0?n:h}}if(!(_&_e))return g;if(i){var A=r.$$legacy;return function(h,S){return arguments.length>0?((!l||!S||A||f)&&i(S?g():h),h):g()}}var w=!1,F=!1,C=Se(e),T=V(()=>H(()=>{var h=g(),S=E(C);return w?(w=!1,F=!0,S):(F=!1,C.v=h)}));return a||(T.equals=ve),function(h,S){if(arguments.length>0){const D=S?E(T):l&&d?I(h):h;return T.equals(D)||(w=!0,P(C,D),o&&n!==void 0&&(n=D),B(()=>E(T))),h}return E(T)}}const Oe="modulepreload",Ae=function(r,s){return new URL(r,s).href},W={},qe=function(s,_,v){let a=Promise.resolve();if(_&&_.length>0){const d=document.getElementsByTagName("link"),c=document.querySelector("meta[property=csp-nonce]"),f=(c==null?void 0:c.nonce)||(c==null?void 0:c.getAttribute("nonce"));a=Promise.allSettled(_.map(e=>{if(e=Ae(e,v),e in W)return;W[e]=!0;const t=e.endsWith(".css"),i=t?'[rel="stylesheet"]':"";if(!!v)for(let o=d.length-1;o>=0;o--){const b=d[o];if(b.href===e&&(!t||b.rel==="stylesheet"))return}else if(document.querySelector(`link[href="${e}"]${i}`))return;const u=document.createElement("link");if(u.rel=t?"stylesheet":Oe,t||(u.as="script"),u.crossOrigin="",u.href=e,f&&u.setAttribute("nonce",f),document.head.appendChild(u),t)return new Promise((o,b)=>{u.addEventListener("load",o),u.addEventListener("error",()=>b(new Error(`Unable to preload CSS for ${e}`)))})}))}function l(d){const c=new Event("vite:preloadError",{cancelable:!0});if(c.payload=d,window.dispatchEvent(c),!c.defaultPrevented)throw d}return a.then(d=>{for(const c of d||[])c.status==="rejected"&&l(c.reason);return s().catch(l)})};export{qe as _,De as a,Ce as b,Le as c,Ne as i,I as p};
