<template>
  <div class="login-container-image">
    <a-row>
        <a-col :xs="0" :md="0" :sm="12" :lg="14" :xl="16"></a-col>
        <a-col :xs="24" :sm="24" :md="12" :lg="10" :xl="6">
          <a-card :title="$t('operation.login')" :bordered="false" class="login-card">
            <template #extra>
              <a-radio-group v-model:value="locale">
                <a-radio-button @click="changeLanguage('zh_CN')" value="zh_CN">{{$t('lang.cn')}}</a-radio-button>
                <a-radio-button @click="changeLanguage('en_US')" value="en_US">{{$t('lang.en')}}</a-radio-button>
              </a-radio-group>
            </template> 
            <a-row>
              <a-form
                  :model="loginForm"
                  name="normal_login"
                  class="login-form"
                  @submit="loginSubmit"
                  @finish="onFinish"
                  @finishFailed="onFinishFailed"
              >
                <a-form-item
                    :label="$t('noun.username')"
                    name="username"
                    :labelCol="labelCol"
                    label-align="left"
                    :rules="[{ required: true, message:  $t('message.input_username') }]"
                >
                  <a-input v-model:value="loginForm.username">
                    <template #prefix>
                      <UserOutlined class="site-form-item-icon" />
                    </template>
                  </a-input>
                </a-form-item>

                <a-form-item
                    :label="$t('noun.password')"
                    name="password"
                    :labelCol="labelCol"
                    label-align="left"
                    :rules="[{ required: true, message: $t('message.input_pass') }]"
                >
                  <a-input-password v-model:value="loginForm.password" >
                    <template #prefix>
                      <LockOutlined class="site-form-item-icon" />
                    </template>
                  </a-input-password>
                </a-form-item>

                <a-form-item
                    :label="$t('noun.captcha')"
                    name="captcha"
                    :rules="[{ required: true, message: $t('message.input_captcha') }]"
                >
                    <a-row>
                      <a-col :span="14">
                        <img :src="imgPath" style="width: 90%"/>
                      </a-col>
                      <a-col :span="10">
                        <a-input v-model:value="loginForm.captcha" :placeholder="$t('noun.captcha')"/>
                      </a-col>
                    </a-row>
                </a-form-item>

                <a-form-item
                    v-show="false"
                    name="captchaId"
                >
                  <a-input v-model:value="loginForm.captchaId" />
                </a-form-item>

                <div class="login-form-wrap">
                  <a-form-item name="remember" no-style>
                    <a-checkbox v-model:checked="loginForm.isRemember">{{$t('operation.remember_me')}}</a-checkbox>
                  </a-form-item>
                  <a class="login-form-forgot" href="">{{$t('operation.forgot_pass')}}?</a>
                </div>

                <a-form-item>
                  <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button">
                    {{$t('operation.login')}}
                  </a-button>
                  <a-divider>{{$t('common.or')}}</a-divider>
                  <a-button type="primary" class="signup-button" @click="redirectToSignUp">
                    {{$t('operation.signup')}}
                  </a-button>
                </a-form-item>
              </a-form>
            </a-row>
          </a-card>
        </a-col>
    </a-row>
  </div>
</template>

<script>
import {Row, Card, Col, RadioButton, notification} from 'ant-design-vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import { getCaptcha } from '@/apis/captcha';
import { login } from "@/apis/user";
import { reactive, computed } from 'vue';
import router from '@/router/index'


export default {
  name: "log-in",
  components: {
    ARow: Row,
    ACard: Card,
    ACol: Col,
    ARadioButton: RadioButton,
    UserOutlined,
    LockOutlined,
  },
  data(){
    return {
      locale: "en-US",
      imgPath: "",
      labelCol: {
        span: 8
      },
    }
  },
  methods: {
    changeLanguage(val){
      localStorage.setItem('lang', val);
      this.$emit('switchLocale', val);
    },
    redirectToSignUp(){
      router.push('/user/signup')
    },
    refreshCaptcha(){
      getCaptcha().then((data) => {
        this.loginForm.captchaId = data.data.captchaId;
        this.imgPath = data.data.imgPath;
      })
    },
    loginSubmit(){
      let loginData = {
        username: this.loginForm.username,
        password: this.loginForm.password,
        captcha: this.loginForm.captcha,
        captchaId: this.loginForm.captchaId
      }

      login(loginData).then(data => {
        if (data['code'] === 200) {
          this.openNotification('success', this.$t('operation.login'), this.$t('message.login_success'))
          let d = data['data']
          localStorage.setItem('userId', d['userId'])
          localStorage.setItem('avatar', d['avatar'])
          localStorage.setItem('username', d['username'])
          localStorage.setItem('accessExpire', d['accessExpire'])
          localStorage.setItem('token', data['data']['accessToken'])
          setTimeout(() => {
            router.push('/dashboard')
          }, 3000)
        } else if (data['code'] === 400) {
          this.openNotification('error', this.$t('operation.signup'), this.$t('message.wrong_captcha'))
          this.refreshCaptcha()
        }
      })
    }
  },
  emits: ['switchLocale'],
  setup() {
     const loginForm = reactive({
      username: "",
      password: "",
      captcha: "",
      captchaId: "",
      isRemember: true,
    });

    const onFinish = data => {
      console.log(data)
    };

    const onFinishFailed = errorInfo => {
      console.log('Failed:', errorInfo);
    };

    const disabled = computed(() => {
      return !(loginForm.username && loginForm.password);
    });

    const openNotification = (type, title, content) => {
      notification[type]({
        message:title,
        description: content,
      });
    };

    return {
      loginForm,
      onFinish,
      onFinishFailed,
      disabled,
      openNotification
    };
  },
  mounted(){
    // initialize languages
    this.locale = localStorage.getItem('lang')
    this.$i18n.locale = this.locale

    this.refreshCaptcha()
  },
};
</script>

<style scoped>
.login-container-image {
    width: 100%;
    height: 100vh;
    background: url('~@/assets/images/login_background.jpg');
    background-size: cover;
}

/* if you like gradient color, select this one  */
.login-container {
  width: 100%;
  height: 100vh;
  background: -webkit-linear-gradient(top, #007aff, #23a9ff, #9aebfe);
}

.signup-button{
  background-color: deepskyblue;
  width: 100%;
  border: #23a9ff 1px solid;
}


.login-card {
  margin-top: 30vh;
}

.login-form {
  max-width: 300px;
}

.login-form-wrap {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.login-form-forgot {
  margin-bottom: 2px;
}

.login-form-button {
  width: 100%;
}
</style>