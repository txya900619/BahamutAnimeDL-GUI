(function(e){function t(t){for(var a,s,c=t[0],o=t[1],l=t[2],u=0,g=[];u<c.length;u++)s=c[u],Object.prototype.hasOwnProperty.call(i,s)&&i[s]&&g.push(i[s][0]),i[s]=0;for(a in o)Object.prototype.hasOwnProperty.call(o,a)&&(e[a]=o[a]);A&&A(t);while(g.length)g.shift()();return r.push.apply(r,l||[]),n()}function n(){for(var e,t=0;t<r.length;t++){for(var n=r[t],a=!0,c=1;c<n.length;c++){var o=n[c];0!==i[o]&&(a=!1)}a&&(r.splice(t--,1),e=s(s.s=n[0]))}return e}var a={},i={app:0},r=[];function s(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,s),n.l=!0,n.exports}s.m=e,s.c=a,s.d=function(e,t,n){s.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},s.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},s.t=function(e,t){if(1&t&&(e=s(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(s.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)s.d(n,a,function(t){return e[t]}.bind(null,a));return n},s.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return s.d(t,"a",t),t},s.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},s.p="/";var c=window["webpackJsonp"]=window["webpackJsonp"]||[],o=c.push.bind(c);c.push=t,c=c.slice();for(var l=0;l<c.length;l++)t(c[l]);var A=o;r.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("cd49")},1427:function(e,t){e.exports="data:image/png;base64,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"},cd49:function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("2b0e"),i=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("router-view")},r=[],s=n("2877"),c={},o=Object(s["a"])(c,i,r,!1,null,null,null),l=o.exports,A=n("8c4f"),u=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-app",{attrs:{id:"inspire"}},[n("Bar"),n("v-content",[n("v-container",[n("router-view",{directives:[{name:"show",rawName:"v-show",value:e.showPage,expression:"showPage"}]}),n("SearchPage",{directives:[{name:"show",rawName:"v-show",value:!e.showPage,expression:"!showPage"}]})],1),n("FloatButton"),n("v-progress-circular",{staticStyle:{position:"fixed",right:"17px",bottom:"25px"},attrs:{size:"62",value:"90",color:""}})],1)],1)},g=[],h=(n("ac1f"),n("841c"),n("d4ec")),d=n("bee2"),m=n("99de"),I=n("7e84"),p=n("262e"),E=n("9ab4"),C=n("60a3"),f=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("v-navigation-drawer",{staticStyle:{"z-index":"5","padding-top":"64px"},attrs:{app:"",clipped:"",temporary:""},model:{value:e.drawer,callback:function(t){e.drawer=t},expression:"drawer"}},[n("v-list",{attrs:{nav:""}},[n("v-list-item",{attrs:{"active-class":"deep-purple--text text--accent-4",to:"/"}},[n("v-list-item-title",[e._v("本季新番")])],1),n("v-list-item",{attrs:{"active-class":"deep-purple--text text--accent-4",to:"/Animations"}},[n("v-list-item-title",[e._v("所有動畫")])],1)],1)],1),n("v-app-bar",{attrs:{"clipped-left":"",app:"",color:e.selectMode?"indigo ":void 0}},[e.selectMode?e._e():n("v-app-bar-nav-icon",{on:{click:function(t){e.drawer=!e.drawer}}}),n("v-spacer"),e.selectMode?e._e():n("v-text-field",{attrs:{"hide-details":"","prepend-inner-icon":"mdi-magnify",label:"search"},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}}),n("v-spacer"),e.selectMode?n("v-btn",{attrs:{icon:""}},[n("v-icon",{attrs:{color:"white"}},[e._v("mdi-arrow-collapse-down")])],1):e._e(),e.selectMode?n("v-btn",{attrs:{icon:""},on:{click:e.cancelSelectMode}},[n("v-icon",{attrs:{color:"white"}},[e._v("mdi-close")])],1):e._e()],1)],1)},b=[],B=function(e){function t(){var e;return Object(h["a"])(this,t),e=Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments)),e.drawer=!1,e.search="",e}return Object(p["a"])(t,e),Object(d["a"])(t,[{key:"mounted",value:function(){this.$store.commit("getNewAnime")}},{key:"cancelSelectMode",value:function(){this.$store.commit("unSelectMode")}},{key:"whenSearchChange",value:function(){this.$store.commit("changeSearch",this.search)}},{key:"whenDrawerChange",value:function(){this.search&&(this.drawer=!1)}},{key:"selectMode",get:function(){return this.$store.getters.selectMode}}]),t}(C["d"]);Object(E["a"])([Object(C["e"])("search")],B.prototype,"whenSearchChange",null),Object(E["a"])([Object(C["e"])("drawer")],B.prototype,"whenDrawerChange",null),B=Object(E["a"])([C["a"]],B);var Q=B,v=Q,w=n("6544"),S=n.n(w),O=n("40dc"),j=n("5bc1"),x=n("8336"),M=n("132d"),y=n("8860"),k=n("da13"),D=n("5d23"),G=n("f774"),R=n("2fa4"),Y=n("8654"),P=Object(s["a"])(v,f,b,!1,null,null,null),_=P.exports;S()(P,{VAppBar:O["a"],VAppBarNavIcon:j["a"],VBtn:x["a"],VIcon:M["a"],VList:y["a"],VListItem:k["a"],VListItemTitle:D["a"],VNavigationDrawer:G["a"],VSpacer:R["a"],VTextField:Y["a"]});var N=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",e._l(e.animeData,(function(t,a){return n("v-col",{key:t.ref,attrs:{md:"4",lg:"2",cols:"6"}},[n("AnimeCard",{attrs:{animeImg:t.img,animeTitle:t.title,animeRef:t.ref,indexInPage:a,isSelected:e.animeSelectStatus[a]},on:{"click-on-select-mode":e.clickOnSelectMode}})],1)})),1)},J=[],V=(n("4de4"),n("d3b7"),n("96cf"),n("1da1")),T=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-card",{attrs:{height:"100%"},on:{mousedown:e.onMouseDown,mouseout:e.cancelTimer,click:e.onClick}},[a("v-img",{staticClass:"white--text align-end",attrs:{src:e.src,height:"100%",gradient:e.isSelected?"rgba(0,0,0,0) 1%,#42A5F5":"to bottom,rgba(0,0,0,0),rgba(0,0,0,.5)","lazy-src":n("1427")}},[a("v-card-title",{staticClass:"d-inline-block text-truncate",staticStyle:{"font-size":"20px","max-width":"100%"},domProps:{textContent:e._s(e.animeTitle)}}),a("v-dialog",{model:{value:e.dialog,callback:function(t){e.dialog=t},expression:"dialog"}},[a("AnimeDialog",{attrs:{sns:e.dialogSns,title:e.animeTitle}})],1)],1)],1)},$=[],z=(n("a9e3"),function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-card",[n("v-container",{attrs:{fluid:""}},[n("v-card-title",[e._v(e._s(e.title)+" ")]),Object.keys(e.sns).length<2?n("v-row",{attrs:{justify:"start"}},e._l(e.sns[""],(function(t){return n("v-col",{key:t.sn,attrs:{md:"1"}},[n("v-row",{attrs:{justify:"center"}},[n("v-btn",[e._v(e._s(t.number))])],1)],1)})),1):n("div",[n("div",{staticStyle:{"font-size":"18px","font-weight":"bold"}},[e._v("本篇")]),n("v-row",{attrs:{justify:"start"}},e._l(e.sns["本篇"],(function(t){return n("v-col",{key:t.sn,attrs:{md:"1"}},[n("v-row",{attrs:{justify:"center"}},[n("v-btn",[e._v(e._s(t.number))])],1)],1)})),1),n("div",{staticStyle:{"font-size":"18px","font-weight":"bold"}},[e._v("特別篇")]),n("v-row",{attrs:{justify:"start"}},e._l(e.sns["特別篇"],(function(t){return n("v-col",{key:t.sn,attrs:{md:"1"}},[n("v-row",{attrs:{justify:"center"}},[n("v-btn",[e._v(e._s(t.number))])],1)],1)})),1)],1)],1)],1)}),U=[],F=function(e){function t(){return Object(h["a"])(this,t),Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments))}return Object(p["a"])(t,e),t}(C["d"]);Object(E["a"])([Object(C["c"])()],F.prototype,"sns",void 0),Object(E["a"])([Object(C["c"])(String)],F.prototype,"title",void 0),F=Object(E["a"])([C["a"]],F);var L=F,K=L,W=n("b0af"),Z=n("99d9"),H=n("62ad"),X=n("a523"),q=n("0fd9"),ee=Object(s["a"])(K,z,U,!1,null,null,null),te=ee.exports;S()(ee,{VBtn:x["a"],VCard:W["a"],VCardTitle:Z["a"],VCol:H["a"],VContainer:X["a"],VRow:q["a"]});var ne=function(e){function t(){var e;return Object(h["a"])(this,t),e=Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments)),e.src=e.animeImg,e.timer=null,e.firstSelected=!1,e.dialog=!1,e.dialogSns={1:[{sn:"0",number:"1"}]},e}return Object(p["a"])(t,e),Object(d["a"])(t,[{key:"cancelTimer",value:function(){null!==this.timer&&(clearTimeout(this.timer),this.timer=null)}},{key:"onClick",value:function(){var e=Object(V["a"])(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(this.cancelTimer(),!this.firstSelected){e.next=4;break}return this.firstSelected=!1,e.abrupt("return");case 4:if(!this.selectMode){e.next=8;break}this.clickOnSelectMode(),e.next=24;break;case 8:if(!this.animeSn){e.next=14;break}return e.t0=JSON,e.next=12,window.getAnimeAllSn(this.animeSn);case 12:e.t1=e.sent,this.dialogSns=e.t0.parse.call(e.t0,e.t1);case 14:if(!this.animeRef){e.next=23;break}return e.next=17,window.getRealSn(this.animeRef);case 17:return t=e.sent,e.t2=JSON,e.next=21,window.getAnimeAllSn(t);case 21:e.t3=e.sent,this.dialogSns=e.t2.parse.call(e.t2,e.t3);case 23:this.dialog=!0;case 24:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"onMouseDown",value:function(){this.selectMode||(this.timer=setTimeout(this.clickOnSelectMode,500))}},{key:"clickOnSelectMode",value:function(){return this.selectMode||(this.firstSelected=!0,this.$store.commit("toSelectMode")),this.animeRef?{index:this.indexInPage,ref:this.animeRef}:{index:this.indexInPage,sn:this.animeSn}}},{key:"selectMode",get:function(){return this.$store.getters.selectMode}}]),t}(C["d"]);Object(E["a"])([Object(C["c"])(String)],ne.prototype,"animeImg",void 0),Object(E["a"])([Object(C["c"])(String)],ne.prototype,"animeTitle",void 0),Object(E["a"])([Object(C["c"])(String)],ne.prototype,"animeSn",void 0),Object(E["a"])([Object(C["c"])(String)],ne.prototype,"animeRef",void 0),Object(E["a"])([Object(C["c"])(Number)],ne.prototype,"indexInPage",void 0),Object(E["a"])([Object(C["c"])()],ne.prototype,"isSelected",void 0),Object(E["a"])([Object(C["b"])("click-on-select-mode")],ne.prototype,"clickOnSelectMode",null),ne=Object(E["a"])([Object(C["a"])({components:{AnimeDialog:te}})],ne);var ae=ne,ie=ae,re=n("169a"),se=n("adda"),ce=Object(s["a"])(ie,T,$,!1,null,null,null),oe=ce.exports;S()(ce,{VCard:W["a"],VCardTitle:Z["a"],VDialog:re["a"],VImg:se["a"]});var le=function(e){function t(){var e;return Object(h["a"])(this,t),e=Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments)),e.animeData=[],e.searchUpdateTime=0,e.selectedAnimes=[],e.animeSelectStatus=[],e}return Object(p["a"])(t,e),Object(d["a"])(t,[{key:"sleep",value:function(){var e=Object(V["a"])(regeneratorRuntime.mark((function e(t){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.abrupt("return",new Promise((function(e){return setTimeout(e,t)})));case 1:case"end":return e.stop()}}),e)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"onAnimeData",value:function(){for(var e=[],t=0;t<this.animeData.length;t++)e.push(!1);this.animeSelectStatus=e}},{key:"OnSearchChange",value:function(){var e=Object(V["a"])(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(""==this.search){e.next=14;break}return this.searchUpdateTime=(new Date).valueOf(),e.next=4,this.sleep(500);case 4:if(t=(new Date).valueOf(),!(t-this.searchUpdateTime<500)){e.next=7;break}return e.abrupt("return");case 7:return e.t0=JSON,e.next=10,window.getAnimesByFilter(this.search);case 10:e.t1=e.sent,this.animeData=e.t0.parse.call(e.t0,e.t1),e.next=16;break;case 14:this.searchUpdateTime=(new Date).valueOf(),this.animeData=[];case 16:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"onSelectModeChange",value:function(){if(!this.selectMode){for(var e=[],t=0;t<this.animeData.length;t++)e.push(!1);this.animeSelectStatus=e,this.selectedAnimes=[]}}},{key:"clickOnSelectMode",value:function(){var e=Object(V["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.$set(this.animeSelectStatus,t.index,!this.animeSelectStatus[t.index]),e.next=3,window.getRealSn(t.ref);case 3:n=e.sent,this.animeSelectStatus[t.index]?this.selectedAnimes.push(n):(this.selectedAnimes=this.selectedAnimes.filter((function(e){return e!==n})),this.selectedAnimes.length<=0&&this.$store.commit("unSelectMode"));case 5:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"search",get:function(){return this.$store.getters.search}},{key:"selectMode",get:function(){return this.$store.getters.selectMode}}]),t}(C["d"]);Object(E["a"])([Object(C["e"])("animeData")],le.prototype,"onAnimeData",null),Object(E["a"])([Object(C["e"])("search")],le.prototype,"OnSearchChange",null),Object(E["a"])([Object(C["e"])("selectMode")],le.prototype,"onSelectModeChange",null),le=Object(E["a"])([Object(C["a"])({components:{AnimeCard:oe}})],le);var Ae=le,ue=Ae,ge=Object(s["a"])(ue,N,J,!1,null,null,null),he=ge.exports;S()(ge,{VCol:H["a"],VRow:q["a"]});var de=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-fab-transition",[n("v-btn",{staticClass:"mr-1 mb-3",attrs:{fab:"",right:"",bottom:"",fixed:"",color:"blue darken-3"}},[n("v-icon",{attrs:{"x-large":"",color:"grey lighten-3"}},[e._v(" mdi-arrow-down-bold ")])],1)],1)},me=[],Ie=function(e){function t(){return Object(h["a"])(this,t),Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments))}return Object(p["a"])(t,e),t}(C["d"]);Ie=Object(E["a"])([C["a"]],Ie);var pe=Ie,Ee=pe,Ce=n("0789"),fe=Object(s["a"])(Ee,de,me,!1,null,null,null),be=fe.exports;S()(fe,{VBtn:x["a"],VFabTransition:Ce["b"],VIcon:M["a"]});var Be=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-progress-circular",{staticStyle:{position:"fixed",right:"17px",bottom:"25px"},attrs:{size:"62",value:"90",color:""}})},Qe=[],ve=function(e){function t(){return Object(h["a"])(this,t),Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments))}return Object(p["a"])(t,e),t}(C["d"]);ve=Object(E["a"])([C["a"]],ve);var we=ve,Se=we,Oe=n("490a"),je=Object(s["a"])(Se,Be,Qe,!1,null,null,null),xe=je.exports;S()(je,{VProgressCircular:Oe["a"]});var Me=n("2f62");a["a"].use(Me["a"]);var ye=new Me["a"].Store({state:{newAnime:[],search:"",selectMode:!1},mutations:{getNewAnime:function(e){window.getNewAnimeList().then((function(t){e.newAnime=JSON.parse(t)}))},changeSearch:function(e,t){e.search=t},toSelectMode:function(e){e.selectMode=!0},unSelectMode:function(e){e.selectMode=!1}},actions:{},modules:{},getters:{newAnime:function(e){return e.newAnime},search:function(e){return e.search},selectMode:function(e){return e.selectMode}}}),ke=function(e){function t(){var e;return Object(h["a"])(this,t),e=Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments)),e.showPage=!0,e}return Object(p["a"])(t,e),Object(d["a"])(t,[{key:"OnSearchChange",value:function(){""===this.search?this.showPage=!0:this.showPage=!1}},{key:"search",get:function(){return ye.getters.search}}]),t}(C["d"]);Object(E["a"])([Object(C["e"])("search")],ke.prototype,"OnSearchChange",null),ke=Object(E["a"])([Object(C["a"])({components:{Bar:_,SearchPage:he,FloatButton:be,ProgressCircle:xe}})],ke);var De=ke,Ge=De,Re=n("7496"),Ye=n("a75b"),Pe=Object(s["a"])(Ge,u,g,!1,null,null,null),_e=Pe.exports;S()(Pe,{VApp:Re["a"],VContainer:X["a"],VContent:Ye["a"],VProgressCircular:Oe["a"]});var Ne=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-row",e._l(e.animeData,(function(t,a){return n("v-col",{key:t.sn,attrs:{md:"4",lg:"3",cols:"6"}},[n("AnimeCard",{attrs:{animeImg:t.img,animeTitle:t.title,animeSn:t.sn,indexInPage:a,isSelected:e.animeSelectStatus[a]},on:{"click-on-select-mode":e.clickOnSelectMode}})],1)})),1)},Je=[],Ve=function(e){function t(){var e;return Object(h["a"])(this,t),e=Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments)),e.selectedAnimes=[],e.animeSelectStatus=[],e}return Object(p["a"])(t,e),Object(d["a"])(t,[{key:"onAnimeData",value:function(){for(var e=[],t=0;t<this.animeData.length;t++)e.push(!1);this.animeSelectStatus=e}},{key:"onSelectModeChange",value:function(){if(!this.selectMode){for(var e=[],t=0;t<this.animeData.length;t++)e.push(!1);this.animeSelectStatus=e,this.selectedAnimes=[]}}},{key:"clickOnSelectMode",value:function(e){this.$set(this.animeSelectStatus,e.index,!this.animeSelectStatus[e.index]),this.animeSelectStatus[e.index]?this.selectedAnimes.push(e.sn):(this.selectedAnimes=this.selectedAnimes.filter((function(t){return t!==e.sn})),this.selectedAnimes.length<=0&&this.$store.commit("unSelectMode"))}},{key:"animeData",get:function(){return this.$store.getters.newAnime}},{key:"selectMode",get:function(){return this.$store.getters.selectMode}}]),t}(C["d"]);Object(E["a"])([Object(C["e"])("animeData")],Ve.prototype,"onAnimeData",null),Object(E["a"])([Object(C["e"])("selectMode")],Ve.prototype,"onSelectModeChange",null),Ve=Object(E["a"])([Object(C["a"])({components:{AnimeCard:oe}})],Ve);var Te=Ve,$e=Te,ze=Object(s["a"])($e,Ne,Je,!1,null,null,null),Ue=ze.exports;S()(ze,{VCol:H["a"],VRow:q["a"]});var Fe=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("v-row",e._l(e.animeData,(function(t,a){return n("v-col",{key:t.ref,attrs:{md:"4",lg:"2",cols:"6"}},[n("AnimeCard",{attrs:{indexInPage:a,animeImg:t.img,animeTitle:t.title,animeRef:t.ref,isSelected:e.selectedIndexInPages[e.page-1][a]},on:{"click-on-select-mode":e.clickOnSelectMode}})],1)})),1),n("v-pagination",{directives:[{name:"show",rawName:"v-show",value:e.showPagination,expression:"showPagination"}],attrs:{length:e.maxPage,"total-visible":"10"},model:{value:e.page,callback:function(t){e.page=t},expression:"page"}})],1)},Le=[],Ke=function(e){function t(){var e;return Object(h["a"])(this,t),e=Object(m["a"])(this,Object(I["a"])(t).apply(this,arguments)),e.animeData=[],e.page=1,e.maxPage=0,e.showPagination=!0,e.selectedAnimes=[],e.selectedIndexInPages=[],e}return Object(p["a"])(t,e),Object(d["a"])(t,[{key:"mounted",value:function(){var e=Object(V["a"])(regeneratorRuntime.mark((function e(){var t,n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,window.getMaxPage();case 2:for(this.maxPage=e.sent,t=0;t<this.maxPage;t++){for(n=[],a=0;a<18;a++)n.push(!1);this.selectedIndexInPages.push(n)}return e.t0=JSON,e.next=7,window.getAnimesByPage(1);case 7:e.t1=e.sent,this.animeData=e.t0.parse.call(e.t0,e.t1);case 9:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"OnPageChange",value:function(){var e=this;window.getAnimesByPage(this.page).then((function(t){e.animeData=JSON.parse(t)}))}},{key:"onSelectModeChange",value:function(){if(!this.selectMode){for(var e=[],t=0;t<this.maxPage;t++){for(var n=[],a=0;a<18;a++)n.push(!1);e.push(n),this.selectedIndexInPages=e}this.selectedAnimes=[]}}},{key:"clickOnSelectMode",value:function(){var e=Object(V["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.selectedIndexInPages[this.page-1][t.index]=!this.selectedIndexInPages[this.page-1][t.index],this.$set(this.selectedIndexInPages,this.page-1,this.selectedIndexInPages[this.page-1]),e.next=4,window.getRealSn(t.ref);case 4:n=e.sent,this.selectedIndexInPages[this.page-1][t.index]?this.selectedAnimes.push(n):(this.selectedAnimes=this.selectedAnimes.filter((function(e){return e!==n})),this.selectedAnimes.length<=0&&this.$store.commit("unSelectMode")),console.log(this.selectedAnimes);case 7:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"selectMode",get:function(){return this.$store.getters.selectMode}}]),t}(C["d"]);Object(E["a"])([Object(C["e"])("page")],Ke.prototype,"OnPageChange",null),Object(E["a"])([Object(C["e"])("selectMode")],Ke.prototype,"onSelectModeChange",null),Ke=Object(E["a"])([Object(C["a"])({components:{AnimeCard:oe}})],Ke);var We=Ke,Ze=We,He=n("891e"),Xe=Object(s["a"])(Ze,Fe,Le,!1,null,null,null),qe=Xe.exports;S()(Xe,{VCol:H["a"],VPagination:He["a"],VRow:q["a"]}),a["a"].use(A["a"]);var et=[{path:"/",name:"Main",component:_e,children:[{path:"/",name:"NewAnimePage",component:Ue},{path:"/Animations",name:"AllAnimePage",component:qe}]}],tt=new A["a"]({mode:"history",base:"/",routes:et}),nt=tt,at=n("f309");a["a"].use(at["a"]);var it=new at["a"]({});a["a"].config.productionTip=!1,new a["a"]({router:nt,store:ye,vuetify:it,render:function(e){return e(l)}}).$mount("#app")}});
//# sourceMappingURL=app.a5e749e4.js.map