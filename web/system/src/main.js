import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import App from './App';
import 'ant-design-vue/dist/antd.css';
import router from './router'
import newI18n from './languages/i18n';

const app = createApp(App);
app.use(router).use(newI18n).use(Antd).mount('#app');
