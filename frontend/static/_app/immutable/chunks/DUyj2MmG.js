import{Y as h,Z as L,_ as Y,P as _,a0 as q,b as o,a1 as P,a2 as u,g as x,y as F,a3 as H,a4 as K,S as M,h as Z,i as R,k as B,a5 as C,H as U,l as z,j as G,m as O,q as A,u as S,v as j,o as J}from"./BHXzz-95.js";function b(i,g=null,T){if(typeof i!="object"||i===null||h in i)return i;const v=K(i);if(v!==L&&v!==Y)return i;var a=new Map,l=M(i),c=_(0);l&&a.set("length",_(i.length));var y;return new Proxy(i,{defineProperty(f,e,t){(!("value"in t)||t.configurable===!1||t.enumerable===!1||t.writable===!1)&&q();var r=a.get(e);return r===void 0?(r=_(t.value),a.set(e,r)):o(r,b(t.value,y)),!0},deleteProperty(f,e){var t=a.get(e);if(t===void 0)e in f&&a.set(e,_(u));else{if(l&&typeof e=="string"){var r=a.get("length"),n=Number(e);Number.isInteger(n)&&n<r.v&&o(r,n)}o(t,u),D(c)}return!0},get(f,e,t){var d;if(e===h)return i;var r=a.get(e),n=e in f;if(r===void 0&&(!n||(d=P(f,e))!=null&&d.writable)&&(r=_(b(n?f[e]:u,y)),a.set(e,r)),r!==void 0){var s=x(r);return s===u?void 0:s}return Reflect.get(f,e,t)},getOwnPropertyDescriptor(f,e){var t=Reflect.getOwnPropertyDescriptor(f,e);if(t&&"value"in t){var r=a.get(e);r&&(t.value=x(r))}else if(t===void 0){var n=a.get(e),s=n==null?void 0:n.v;if(n!==void 0&&s!==u)return{enumerable:!0,configurable:!0,value:s,writable:!0}}return t},has(f,e){var s;if(e===h)return!0;var t=a.get(e),r=t!==void 0&&t.v!==u||Reflect.has(f,e);if(t!==void 0||F!==null&&(!r||(s=P(f,e))!=null&&s.writable)){t===void 0&&(t=_(r?b(f[e],y):u),a.set(e,t));var n=x(t);if(n===u)return!1}return r},set(f,e,t,r){var I;var n=a.get(e),s=e in f;if(l&&e==="length")for(var d=t;d<n.v;d+=1){var m=a.get(d+"");m!==void 0?o(m,u):d in f&&(m=_(u),a.set(d+"",m))}n===void 0?(!s||(I=P(f,e))!=null&&I.writable)&&(n=_(void 0),o(n,b(t,y)),a.set(e,n)):(s=n.v!==u,o(n,b(t,y)));var w=Reflect.getOwnPropertyDescriptor(f,e);if(w!=null&&w.set&&w.set.call(r,t),!s){if(l&&typeof e=="string"){var E=a.get("length"),N=Number(e);Number.isInteger(N)&&N>=E.v&&o(E,N+1)}D(c)}return!0},ownKeys(f){x(c);var e=Reflect.ownKeys(f).filter(n=>{var s=a.get(n);return s===void 0||s.v!==u});for(var[t,r]of a)r.v!==u&&!(t in f)&&e.push(t);return e},setPrototypeOf(){H()}})}function D(i,g=1){o(i,i.v+g)}function k(i){return i!==null&&typeof i=="object"&&h in i?i[h]:i}function V(i,g){return Object.is(k(i),k(g))}function W(i,g,T=!1){R&&B();var v=i,a=null,l=null,c=u,y=T?C:0,f=!1;const e=(r,n=!0)=>{f=!0,t(n,r)},t=(r,n)=>{if(c===(c=r))return;let s=!1;if(R){const d=v.data===U;!!c===d&&(v=z(),G(v),O(!1),s=!0)}c?(a?A(a):n&&(a=S(()=>n(v))),l&&j(l,()=>{l=null})):(l?A(l):n&&(l=S(()=>n(v))),a&&j(a,()=>{a=null})),s&&O(!0)};Z(()=>{f=!1,g(e),f||t(null,null)},y),R&&(v=J)}export{V as a,W as i,b as p};
