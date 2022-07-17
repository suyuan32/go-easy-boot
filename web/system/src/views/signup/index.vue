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
                  :model="signupFrom"
                  name="normal_login"
                  class="login-form"
                  :onsubmit="signUp"
              >
                <a-form-item
                    :label="$t('noun.mail')"
                    name="email"
                    :labelCol="labelCol"
                    label-align="left"
                    :rules="[{ required: true, message:  $t('message.input_mail') }]"
                >
                  <a-input v-model:value="signupFrom.mail">
                    <template #prefix>
                      <mail-outlined class="site-form-item-icon" />
                    </template>
                  </a-input>
                </a-form-item>

                <a-form-item
                    :label="$t('noun.username')"
                    name="username"
                    :labelCol="labelCol"
                    label-align="left"
                    :rules="[{ required: true, message:  $t('message.input_username') }]"
                >
                  <a-input v-model:value="signupFrom.username">
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
                  <a-input-password v-model:value="signupFrom.password" >
                    <template #prefix>
                      <LockOutlined class="site-form-item-icon" />
                    </template>
                  </a-input-password>
                </a-form-item>

                <a-form-item
                    :label="$t('common.confirm')"
                    name="confirm"
                    :labelCol="labelCol"
                    label-align="left"
                    :rules="[{ required: true, message: $t('message.input_pass') }]"
                >
                  <a-input-password v-model:value="signupFrom.confirm" >
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
                        <a-input v-model:value="signupFrom.captcha" :placeholder="$t('noun.captcha')"/>
                      </a-col>
                    </a-row>
                </a-form-item>

                <a-form-item>
                  <a-button type="primary" class="signup-button" html-type="submit">
                    {{$t('operation.signup')}}
                  </a-button>
                  <a-divider>{{$t('common.or')}}</a-divider>
                  <a-button type="primary" class="login-form-button" @click="redirectToLogin">
                    {{$t('operation.login')}}
                  </a-button>
                </a-form-item>
              </a-form>
            </a-row>
          </a-card>
        </a-col>
    </a-row>
    <a-alert message="$t('message.signup_success')" v-show="isSignupSuccess" type="success" />
  </div>
</template>

<script>
import { Row, Card, Col, RadioButton, notification } from 'ant-design-vue';
import { UserOutlined, LockOutlined, MailOutlined } from '@ant-design/icons-vue';
import { getCaptcha } from '@/apis/captcha';
import { signup } from "@/apis/user";
import { reactive } from 'vue';
import router from "@/router";


export default {
  name: "sign-up",
  components: {
    ARow: Row,
    ACard: Card,
    ACol: Col,
    ARadioButton: RadioButton,
    UserOutlined,
    LockOutlined,
    MailOutlined
  },
  data(){
    return {
      locale: "en-US",
      imgPath: "",
      labelCol: {
        span: 8
      },
      isSignupSuccess: false
    }
  },
  methods: {
    changeLanguage(val){
      localStorage.setItem('lang', val);
      this.$emit('switchLocale', val);
    },
    redirectToLogin(){
      router.push('/user/login')
    },
    signUp(){
      if (this.signupFrom.password !== this.signupFrom.confirm) {
        this.openNotification('danger', this.$t('noun.error'), this.$t('message.pass_diff'))
        return
      }
      let signUpData = {
        username: this.signupFrom.username,
        password: this.signupFrom.password,
        captcha: this.signupFrom.captcha,
        captchaId: this.signupFrom.captchaId
      }
      signup(signUpData).then(data => {
        if (data.code === 200) {
          this.openNotification('success', this.$t('operation.signup'), this.$t('message.signup_success'))
        }
      })
    },
  },
  emits: ['switchLocale'],
  setup() {
     const signupFrom = reactive({
      username: "",
      password: "",
      captcha: "",
      captchaId: "",
      confirm: "",
      mail: "",
    });

    const openNotification = (type, title, content) => {
      notification[type]({
        message:title,
        description: content,
      });
    };

    return {
      signupFrom,
      openNotification
    };
  },
  mounted(){
    // initialize languages
    this.locale = localStorage.getItem('lang')
    this.$i18n.locale = this.locale

    getCaptcha().then((data) => {
      this.signupFrom.captchaId = data.data.captchaId;
      this.imgPath = data.data.imgPath;
    })
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
  margin-top: 20vh;
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