import{a as B,t as J}from"../chunks/ZUxJqvzW.js";import{s as c,f as it,g as e,b as l,c as t,d as n,r as a,t as et}from"../chunks/CtZ02a2i.js";import{d as pt,s as m}from"../chunks/CVDwAH_T.js";import{H as dt,e as at,b as y,r as b,i as rt}from"../chunks/ChwPNElU.js";import{p as w}from"../chunks/-HZ5zasu.js";var ct=async(x,s,o,p,d,i)=>{const j=await(await fetch(`${s}/api/protected/domain`,{method:"POST",headers:{"Content-Type":"application/json",Authorization:`Bearer ${o==null?void 0:o.getItem("jwt")}`},body:JSON.stringify({display_name:e(p),description:e(d),notes:e(i)})})).json();console.log(j)},lt=async(x,s,o,p)=>{const i=await(await fetch(`${s}/api/domain${e(o)}`,{headers:{"Content-Type":"application/json"}})).json();console.log(i),l(p,w(i.data))},_t=J("<tr><td> </td><td> </td><td> </td><td> </td></tr>"),ht=async(x,s,o,p,d)=>{const g=await(await fetch(`${s}/api/protected/user`,{method:"POST",headers:{"Content-Type":"application/json",Authorization:`Bearer ${o==null?void 0:o.getItem("jwt")}`},body:JSON.stringify({username:e(p),display_name:e(d)})})).json();console.log(g)},vt=async(x,s,o,p)=>{const i=await(await fetch(`${s}/api/user${e(o)}`,{headers:{"Content-Type":"application/json"}})).json();console.log(i),l(p,w(i.data))},ut=J("<tr><td> </td><td> </td><td> </td></tr>"),mt=J('<!> <h1>Testing page</h1> <a href="../">Home</a> <div><h2>Domain</h2> <span><h3>Create</h3> <input type="text" placeholder="Display Name"> <input type="text" placeholder="Description"> <input type="text" placeholder="Notes"> <button>Create</button></span> <span><h3>List</h3> <button>List</button> <input type="text" placeholder="Query Params"> <table><thead><tr><th>Display Name</th><th>Description</th><th>Notes</th><th>ID</th></tr></thead><tbody></tbody></table></span></div> <div><h2>User</h2> <span><h3>Create</h3> <input type="text" placeholder="Username"> <input type="text" placeholder="Display Name"> <button>Create</button></span> <span><h3>List</h3> <button>List</button> <input type="text" placeholder="Query Params"> <table><thead><tr><th>Username</th><th>Display Name</th><th>ID</th></tr></thead><tbody></tbody></table></span></div>',1);function jt(x){const s="http://localhost:8080",o=typeof window<"u"?window.localStorage:null;let p=c(""),d=c(""),i=c(""),g=c(w([])),j=c(""),D=c(""),N=c(""),q=c(w([])),k=c("");var E=mt(),F=it(E);dt(F);var C=t(F,6),$=t(n(C),2),T=t(n($),2);b(T);var L=t(T,2);b(L);var P=t(L,2);b(P);var nt=t(P,2);nt.__click=[ct,s,o,p,d,i],a($);var G=t($,2),K=t(n(G),2);K.__click=[lt,s,j,g];var S=t(K,2);b(S);var M=t(S,2),R=t(n(M));at(R,21,()=>e(g),rt,(r,_)=>{var h=_t(),v=n(h),H=n(v,!0);a(v);var u=t(v),z=n(u,!0);a(u);var f=t(u),A=n(f,!0);a(f);var tt=t(f),st=n(tt,!0);a(tt),a(h),et(()=>{m(H,e(_).display_name),m(z,e(_).description),m(A,e(_).notes),m(st,e(_).id)}),B(r,h)}),a(R),a(M),a(G),a(C);var V=t(C,2),I=t(n(V),2),O=t(n(I),2);b(O);var Q=t(O,2);b(Q);var ot=t(Q,2);ot.__click=[ht,s,o,D,N],a(I);var W=t(I,2),X=t(n(W),2);X.__click=[vt,s,k,q];var U=t(X,2);b(U);var Y=t(U,2),Z=t(n(Y));at(Z,21,()=>e(q),rt,(r,_)=>{var h=ut(),v=n(h),H=n(v,!0);a(v);var u=t(v),z=n(u,!0);a(u);var f=t(u),A=n(f,!0);a(f),a(h),et(()=>{m(H,e(_).username),m(z,e(_).display_name),m(A,e(_).id)}),B(r,h)}),a(Z),a(Y),a(W),a(V),y(T,()=>e(p),r=>l(p,r)),y(L,()=>e(d),r=>l(d,r)),y(P,()=>e(i),r=>l(i,r)),y(S,()=>e(j),r=>l(j,r)),y(O,()=>e(D),r=>l(D,r)),y(Q,()=>e(N),r=>l(N,r)),y(U,()=>e(k),r=>l(k,r)),B(x,E)}pt(["click"]);export{jt as component};
