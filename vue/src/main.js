import Vue from "vue";
import "normalize.css/normalize.css"; // a modern alternative to CSS resets

import Element from "element-ui";
// import "./styles/element-variables.scss";
import 'element-ui/lib/theme-chalk/index.css';

import "@/styles/index.scss"; // global css

import App from "./App";
import store from "./store";
import router from "./router";
import "./icons"; // icon
import "./permission"; // permission control
import mavonEditor from "mavon-editor";
import "mavon-editor/dist/css/index.css";
// use

Vue.use(Element);

Vue.use(mavonEditor);

Vue.filter("parseTime", function(time, cFormat) {
    if (arguments.length === 0 || !time) {
        return null;
    }
    const format = cFormat || "{y}-{m}-{d} {h}:{i}:{s}";
    let date;
    if (typeof time === "object") {
        date = time;
    } else {
        if (typeof time === "string") {
            if (/^[0-9]+$/.test(time)) {
                time = parseInt(time);
            } else {
                // https://stackoverflow.com/questions/4310953/invalid-date-in-safari
                time = time.replace(new RegExp(/-/gm), "/");
            }
        }

        if (typeof time === "number" && time.toString().length === 10) {
            time = time * 1000;
        }
        date = new Date(time);
    }
    const formatObj = {
        y: date.getFullYear(),
        m: date.getMonth() + 1,
        d: date.getDate(),
        h: date.getHours(),
        i: date.getMinutes(),
        s: date.getSeconds(),
        a: date.getDay()
    };
    const time_str = format.replace(/{([ymdhisa])+}/g, (result, key) => {
        const value = formatObj[key];
        // Note: getDay() returns 0 on Sunday
        if (key === "a") {
            return ["日", "一", "二", "三", "四", "五", "六"][value];
        }
        return value.toString().padStart(2, "0");
    });
    return time_str;
});

Vue.config.productionTip = false;
new Vue({
    el: "#app",
    router,
    store,
    render: h => h(App)
});
