import{S as h,o as D,e as j,h as v,i as S,c as l,j as x,U as u,g as o,k as T,l as A,m as E,q as K}from"./jrd53t_p.js";function y(s,P=null,L){if(typeof s!="object"||s===null||h in s)return s;const I=E(s);if(I!==D&&I!==j)return s;var f=new Map,b=K(s),w=v(0);b&&f.set("length",v(s.length));var g;return new Proxy(s,{defineProperty(i,e,t){(!("value"in t)||t.configurable===!1||t.enumerable===!1||t.writable===!1)&&S();var n=f.get(e);return n===void 0?(n=v(t.value),f.set(e,n)):l(n,y(t.value,g)),!0},deleteProperty(i,e){var t=f.get(e);if(t===void 0)e in i&&f.set(e,v(u));else{if(b&&typeof e=="string"){var n=f.get("length"),r=Number(e);Number.isInteger(r)&&r<n.v&&l(n,r)}l(t,u),R(w)}return!0},get(i,e,t){var d;if(e===h)return s;var n=f.get(e),r=e in i;if(n===void 0&&(!r||(d=x(i,e))!=null&&d.writable)&&(n=v(y(r?i[e]:u,g)),f.set(e,n)),n!==void 0){var a=o(n);return a===u?void 0:a}return Reflect.get(i,e,t)},getOwnPropertyDescriptor(i,e){var t=Reflect.getOwnPropertyDescriptor(i,e);if(t&&"value"in t){var n=f.get(e);n&&(t.value=o(n))}else if(t===void 0){var r=f.get(e),a=r==null?void 0:r.v;if(r!==void 0&&a!==u)return{enumerable:!0,configurable:!0,value:a,writable:!0}}return t},has(i,e){var a;if(e===h)return!0;var t=f.get(e),n=t!==void 0&&t.v!==u||Reflect.has(i,e);if(t!==void 0||T!==null&&(!n||(a=x(i,e))!=null&&a.writable)){t===void 0&&(t=v(n?y(i[e],g):u),f.set(e,t));var r=o(t);if(r===u)return!1}return n},set(i,e,t,n){var O;var r=f.get(e),a=e in i;if(b&&e==="length")for(var d=t;d<r.v;d+=1){var c=f.get(d+"");c!==void 0?l(c,u):d in i&&(c=v(u),f.set(d+"",c))}r===void 0?(!a||(O=x(i,e))!=null&&O.writable)&&(r=v(void 0),l(r,y(t,g)),f.set(e,r)):(a=r.v!==u,l(r,y(t,g)));var _=Reflect.getOwnPropertyDescriptor(i,e);if(_!=null&&_.set&&_.set.call(n,t),!a){if(b&&typeof e=="string"){var N=f.get("length"),m=Number(e);Number.isInteger(m)&&m>=N.v&&l(N,m+1)}R(w)}return!0},ownKeys(i){o(w);var e=Reflect.ownKeys(i).filter(r=>{var a=f.get(r);return a===void 0||a.v!==u});for(var[t,n]of f)n.v!==u&&!(t in i)&&e.push(t);return e},setPrototypeOf(){A()}})}function R(s,P=1){l(s,s.v+P)}export{y as p};
