(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-d18a9468"],{"615b":function(t,e,a){},b0af:function(t,e,a){"use strict";a("0481"),a("4069"),a("a9e3");var r=a("5530"),n=(a("615b"),a("10d2")),c=a("297c"),o=a("1c87"),s=a("58df");e["a"]=Object(s["a"])(c["a"],o["a"],n["a"]).extend({name:"v-card",props:{flat:Boolean,hover:Boolean,img:String,link:Boolean,loaderHeight:{type:[Number,String],default:4},outlined:Boolean,raised:Boolean,shaped:Boolean},computed:{classes:function(){return Object(r["a"])({"v-card":!0},o["a"].options.computed.classes.call(this),{"v-card--flat":this.flat,"v-card--hover":this.hover,"v-card--link":this.isClickable,"v-card--loading":this.loading,"v-card--disabled":this.disabled,"v-card--outlined":this.outlined,"v-card--raised":this.raised,"v-card--shaped":this.shaped},n["a"].options.computed.classes.call(this))},styles:function(){var t=Object(r["a"])({},n["a"].options.computed.styles.call(this));return this.img&&(t.background='url("'.concat(this.img,'") center center / cover no-repeat')),t}},methods:{genProgress:function(){var t=c["a"].options.methods.genProgress.call(this);return t?this.$createElement("div",{staticClass:"v-card__progress",key:"progress"},[t]):null}},render:function(t){var e=this.generateRouteLink(),a=e.tag,r=e.data;return r.style=this.styles,this.isClickable&&(r.attrs=r.attrs||{},r.attrs.tabindex=0),t(a,this.setBackgroundColor(this.color,r),[this.genProgress(),this.$slots.default])}})},f642:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-col",{attrs:{md:"4",lg:"3",cols:"6"}},[a("v-card",{attrs:{height:"100%"}},[a("AnimeImg",{attrs:{srcURL:t.animeImg,title:t.title}})],1)],1)},n=[],c=(a("d3b7"),a("d4ec")),o=a("99de"),s=a("7e84"),i=a("262e"),l=a("9ab4"),u=a("60a3"),d=function(){return a.e("chunk-4e1e40ab").then(a.bind(null,"122d"))},f=function(t){function e(){return Object(c["a"])(this,e),Object(o["a"])(this,Object(s["a"])(e).apply(this,arguments))}return Object(i["a"])(e,t),e}(u["c"]);Object(l["a"])([Object(u["b"])(String)],f.prototype,"animeImg",void 0),Object(l["a"])([Object(u["b"])(String)],f.prototype,"title",void 0),Object(l["a"])([Object(u["b"])(String)],f.prototype,"remoteImg",void 0),f=Object(l["a"])([Object(u["a"])({components:{AnimeImg:d}})],f);var b=f,p=b,g=a("2877"),h=a("6544"),v=a.n(h),m=a("b0af"),j=(a("4160"),a("caad"),a("13d5"),a("45fc"),a("4ec9"),a("a9e3"),a("b64b"),a("ac1f"),a("3ca3"),a("5319"),a("2ca0"),a("159b"),a("ddb0"),a("ade3")),O=a("5530"),y=(a("4b85"),a("2b0e")),S=a("d9f7"),k=a("80d2"),w=["sm","md","lg","xl"],B=function(){return w.reduce((function(t,e){return t[e]={type:[Boolean,String,Number],default:!1},t}),{})}(),C=function(){return w.reduce((function(t,e){return t["offset"+Object(k["s"])(e)]={type:[String,Number],default:null},t}),{})}(),N=function(){return w.reduce((function(t,e){return t["order"+Object(k["s"])(e)]={type:[String,Number],default:null},t}),{})}(),x={col:Object.keys(B),offset:Object.keys(C),order:Object.keys(N)};function I(t,e,a){var r=t;if(null!=a&&!1!==a){if(e){var n=e.replace(t,"");r+="-".concat(n)}return"col"!==t||""!==a&&!0!==a?(r+="-".concat(a),r.toLowerCase()):r.toLowerCase()}}var L=new Map,_=y["a"].extend({name:"v-col",functional:!0,props:Object(O["a"])({cols:{type:[Boolean,String,Number],default:!1}},B,{offset:{type:[String,Number],default:null}},C,{order:{type:[String,Number],default:null}},N,{alignSelf:{type:String,default:null,validator:function(t){return["auto","start","end","center","baseline","stretch"].includes(t)}},tag:{type:String,default:"div"}}),render:function(t,e){var a=e.props,r=e.data,n=e.children,c=(e.parent,"");for(var o in a)c+=String(a[o]);var s=L.get(c);return s||function(){var t,e;for(e in s=[],x)x[e].forEach((function(t){var r=a[t],n=I(e,t,r);n&&s.push(n)}));var r=s.some((function(t){return t.startsWith("col-")}));s.push((t={col:!r||!a.cols},Object(j["a"])(t,"col-".concat(a.cols),a.cols),Object(j["a"])(t,"offset-".concat(a.offset),a.offset),Object(j["a"])(t,"order-".concat(a.order),a.order),Object(j["a"])(t,"align-self-".concat(a.alignSelf),a.alignSelf),t)),L.set(c,s)}(),t(a.tag,Object(S["a"])(r,{class:s}),n)}}),E=Object(g["a"])(p,r,n,!1,null,null,null);e["default"]=E.exports;v()(E,{VCard:m["a"],VCol:_})}}]);
//# sourceMappingURL=chunk-d18a9468.af0e10d8.js.map