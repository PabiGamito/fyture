/**
 * App.js
 */

import Vue from "vue"
import Vuex from "vuex"
import VModal from "vue-js-modal"
import 'babel-polyfill'

import App from "./components/App.vue"

// Vue setup
if (process.env.NODE_ENV == "production") {

	Vue.config.silent = true
	Vue.config.devtools = false
	Vue.config.productionTip = false
}

let init = () => {

	// Store
	Vue.use(Vuex)

	// Modal
	Vue.use(VModal, { dialog: true, dynamic: true });

	const store = {
		state: {},
		mutations: {}
	}

	// New app instance
	new Vue({

		el: '#app',
		store: new Vuex.Store(store),
		render: h => h(App)
	})
}
init()
