import { createI18n } from 'vue-i18n'
import messages from '@/languages/lang/index'

const newI18n = createI18n({
    locale: 'en_US', // set locale
    fallbackLocale: 'zh_CN', // set fallback locale
    messages,
});

export default newI18n;
