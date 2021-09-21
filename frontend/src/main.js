import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue'

//Routing
import VueRoute from 'vue-router'
Vue.use(VueRoute)
import router from './router'

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
Vue.component(Wails)

// import VueApexCharts from 'vue-apexcharts'
// Vue.use(VueApexCharts)
// Vue.component('apexchart', VueApexCharts)

Wails.Init(() => {
	new Vue({
        router,
        render: h => h(App)
    }).$mount('#app');
});
