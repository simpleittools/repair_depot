import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import '@/index.css'
import './assets/tailwind.css'
// import '@/assets/vendor/bulma/css/bulma.min.css'

createApp(App).use(router).mount('#app')
