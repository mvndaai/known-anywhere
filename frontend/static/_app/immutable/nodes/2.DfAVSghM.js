import{a as I,t as D}from"../chunks/ZUxJqvzW.js";import{p as q,s as h,f as H,a as F,g as a,b as l,c as o,d as u,r as c,n as L,t as K}from"../chunks/CtZ02a2i.js";import{d as Q,s as M}from"../chunks/CVDwAH_T.js";import{H as X,e as $,b as T,r as S,i as z,s as Y}from"../chunks/ChwPNElU.js";import{p as m}from"../chunks/-HZ5zasu.js";import{o as Z}from"../chunks/BxlYdRe1.js";const tt=async(f,e,n,r,t,i)=>{e==null||e.setItem("subject",a(n)),e==null||e.setItem("username",a(r));const p=await fetch(`${t}/test/jwt`,{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({sub:a(n),username:a(r),exp:Math.floor(new Date().setDate(new Date().getDate()+a(i))/1e3)})}),d=await p.json();p.status===200&&(e==null||e.setItem("jwt",d.data))};var et=async(f,e,n)=>{const r=e==null?void 0:e.getItem("jwt");console.log(r);const t=await fetch(`${n}/test/auth`,{headers:{Authorization:`Bearer ${r}`}});console.log(t);const i=await t.json();console.log(i)},at=D("<li><a> </a></li>"),ot=D("<p> </p> <ul></ul>",1),nt=D('<!> <h1>Known Socially</h1> <a href="./admin">Admin</a> <div><span>Create JWT</span> <input type="text" placeholder="Subject"> <input type="text" placeholder="Username"> <input type="number" placeholder="Expiration"> <button>Generate</button></div> <div><span>Test JWT auth</span> <button>Test Auth</button></div> <h2>Routes</h2> <!> <h2>Welcome to SvelteKit</h2> <p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>',1);function ut(f,e){q(e,!0);const n="http://localhost:8080";let r=h(m({}));Z(async function(){let v=await(await fetch(`${n}/api`)).json();l(r,m(v.data))});const t=typeof window<"u"?window.localStorage:null;let i=h(m((t==null?void 0:t.getItem("subject"))||"")),p=h(m((t==null?void 0:t.getItem("username"))||"")),d=h(m((t==null?void 0:t.getItem("days"))||5));var J=nt(),W=H(J);X(W);var _=o(W,6),b=o(u(_),2);S(b);var g=o(b,2);S(g);var w=o(g,2);S(w);var B=o(w,2);B.__click=[tt,t,i,p,n,d],c(_);var j=o(_,2),E=o(u(j),2);E.__click=[et,t,n],c(j);var G=o(j,4);$(G,17,()=>Object.entries(a(r)),z,(s,v)=>{let N=()=>a(v)[0],P=()=>a(v)[1];var A=ot(),y=H(A),R=u(y,!0);c(y);var O=o(y,2);$(O,21,P,z,(U,C)=>{var x=at(),k=u(x),V=u(k,!0);c(k),c(x),K(()=>{Y(k,"href",a(C)),M(V,a(C))}),I(U,x)}),c(O),K(()=>M(R,N())),I(s,A)}),L(4),T(b,()=>a(i),s=>l(i,s)),T(g,()=>a(p),s=>l(p,s)),T(w,()=>a(d),s=>l(d,s)),I(f,J),F()}Q(["click"]);export{ut as component};
