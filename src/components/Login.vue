<template>
  <div class="register">
    <!-- 向上距离5个边距 -->
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <!-- 外框线 -->
        <b-card title="登陆">
          <b-form>
            <b-form-group label="手机号">
              <b-input
                type="number"
                placeholder="输入手机号"
                v-model="$v.user.telephone.$model"
                :state="validationState('telephone')"
              ></b-input>
              <b-form-invalid-feedback :state="validationState('telephone')">
                手机号必须是11位
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validationState('telephone')">
                手机号符合11位
              </b-form-valid-feedback>
              <!-- <b-form-input
                v-model="user.telephone"
                type="number"
                placeholder="请输入你的手机号(选填)"
              ></b-form-input> -->
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="请输入你的密码"
                :state="validationState('password')"
              ></b-form-input>
              <b-form-valid-feedback :state="validationState('password')">
                <!-- 留空占位 -->
              </b-form-valid-feedback>
              <b-form-invalid-feedback :state="validationState('password')">
                密码必须是6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button @click="login" variant="outline-primary" block
                >登陆</b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators'
// const telephoneValidator = (value)=>/^1[3|5|7|9]\d{9}$/.test(value);
import customValidator from '@/helper/validator'
import { mapActions } from 'vuex';
export default {
  name: "Login",
  data() {
    return {
      user: {
       
        telephone: "",
        password: "",
      },
      validation:null,
    };
  },
  methods:{
    ...mapActions('userModule', { userLogin: 'login' }),
    validationState(name){
      //es6解构赋值
    const {$dirty,$error}=this.$v.user[name];
    return $dirty ? !$error:null;
    },
    login(){
      /* if(this.user.telephone.length!==11){
        this.validation=false;
        return
      }
      this.validation=true; */
     this.$v.user.$touch();
      // 请求api
      this.userLogin(this.user).then(() => {
        // 跳转首页
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        this.$bvToast.toast(err.data.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
      console.log("login")
    }
  },
  validations: {
    user: {
    
      telephone:{
        required,
        telephone:customValidator.telephoneValidator,
      },
      password:{
        required,
        minLength:minLength(6),
      },
    },
  }
};
</script>

<style>
</style>