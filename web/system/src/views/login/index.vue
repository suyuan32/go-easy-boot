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
                  :model="formState"
                  name="normal_login"
                  class="login-form"
                  @finish="onFinish"
                  @finishFailed="onFinishFailed"
              >
                <a-form-item
                    :label="$t('noun.username')"
                    name="username"
                    :rules="[{ required: true, message:  $t('message.input_username') }]"
                >
                  <a-input v-model:value="formState.username">
                    <template #prefix>
                      <UserOutlined class="site-form-item-icon" />
                    </template>
                  </a-input>
                </a-form-item>

                <a-form-item
                    :label="$t('noun.password')"
                    name="password"
                    :rules="[{ required: true, message: $t('message.input_pass') }]"
                >
                  <a-input-password v-model:value="formState.password">
                    <template #prefix>
                      <LockOutlined class="site-form-item-icon" />
                    </template>
                  </a-input-password>
                </a-form-item>

                <div class="login-form-wrap">
                  <a-form-item name="remember" no-style>
                    <a-checkbox v-model:checked="formState.remember">{{$t('operation.remember_me')}}</a-checkbox>
                  </a-form-item>
                  <a class="login-form-forgot" href="">{{$t('operation.forgot_pass')}}</a>
                </div>

                <a-form-item>
                  <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button">
                    {{$t('operation.login')}}
                  </a-button>
                  Or
                  <a href="">{{$t('operation.signup')}}</a>
                </a-form-item>
              </a-form>
            </a-row>
          </a-card>
        </a-col>
    </a-row>
  </div>
</template>

<script>
import { Row, Card, Col } from 'ant-design-vue';
import { reactive, computed } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';


export default {
  name: "log-in",
  components: {
    ARow: Row,
    ACard: Card,
    ACol: Col,
    UserOutlined,
    LockOutlined,
  },
  data(){
    return {
      locale: "en-US",
    }
  },
  methods: {
    changeLanguage(val){
      localStorage.setItem('lang', val);
      console.log(this.locale);
      this.$emit('switchLocale', val);
    }
  },
  emits: ['switchLocale'],
  setup() {
    const formState = reactive({
      username: '',
      password: '',
      remember: true,
    });

    const onFinish = values => {
      console.log('Success:', values);
    };

    const onFinishFailed = errorInfo => {
      console.log('Failed:', errorInfo);
    };

    const disabled = computed(() => {
      return !(formState.username && formState.password);
    });

    return {
      formState,
      onFinish,
      onFinishFailed,
      disabled,
    };
  },
  mounted(){
    // initialize languages
    this.locale = localStorage.getItem('lang')
    this.$i18n.locale = this.locale
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

.login-container {
  width: 100%;
  height: 100vh;
  background: -webkit-linear-gradient(top, #007aff, #23a9ff, #9aebfe);
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
  margin-bottom: 24px;
}

.login-form-button {
  width: 100%;
}
</style>