<template>
  <a-config-provider :locale="locale" :getPopupContainer="getPopupContainer">
    <router-view @switchLocale="switchLanguage" v-if="isRouterAlive"></router-view>
  </a-config-provider>
</template>

<script>
import {ConfigProvider} from 'ant-design-vue'
import antZhCN from 'ant-design-vue/es/locale/zh_CN'
import antEnUS from 'ant-design-vue/es/locale/en_US'

export default {
  name: 'App',
  components: {
    AConfigProvider: ConfigProvider
  },
  data(){
    return {
      locale: antEnUS,
      isRouterAlive: true,
    }
  },
  methods: {
    getPopupContainer(el, dialogContext) {
      if (dialogContext) {
        return dialogContext.getDialogWrap();
      } else {
        return document.body;
      }
    },
    switchLanguage(value) {
      switch (value) {
        case 'zh_CN':
          localStorage.setItem('lang', 'zh_CN');
          this.locale = antZhCN
          this.$i18n.locale = 'zh_CN';
          break;
        case 'en_US':
          localStorage.setItem('lang', 'en_US');
          this.locale = antEnUS
          this.$i18n.locale = 'en_US';
          break;
      }

      // refresh page
      this.isRouterAlive = false; 
      this.$nextTick(function() {
        this.isRouterAlive = true; 
      });
    }
  },
  mounted() {
    localStorage.setItem('lang', 'en_US');
    this.$i18n.locale = 'en_US';
  },
}
</script>

