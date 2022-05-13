<template>
  <div class="register">
    <!-- 向上距离5个边距 -->
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <!-- 外框线 -->
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="请输入你的名称(选填)"
              ></b-form-input>
            </b-form-group>
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
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button @click="register" variant="outline-primary" block
                >注册</b-button
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

//导入缓存服务
// import storageServie from "@/service/storageService.js";

// import userService from "@/service/userService.js"


// import { mapMutations } from 'vuex'

import { mapActions } from 'vuex';
export default {
  name: "Register",
  data() {
    return {
      user: {
        name: "",
        telephone: "",
        password: "",
      },
      validation:null,
    };
  },
  validations: {
    user: {
      name:{

      },
      telephone:{
        required,
        telephone:customValidator.telephoneValidator,
      },
      password:{
        required,
        minLength:minLength(6),
      },
    },
  },
  methods:{
    // ...mapMutations('userModule',['SET_TOKEN','SET_USERINFO']),
    ...mapActions('userModule', { userRegister: 'register' }),

    validationState(name){
      //es6解构赋值
    const {$dirty,$error}=this.$v.user[name];
    return $dirty ? !$error:null;
    },
    register(){
      /* if(this.user.telephone.length!==11){
        this.validation=false;
        return
      }
      this.validation=true; */
      // 验证数据
      this.$v.user.$touch();
      if(this.$v.user.$anyError){
        return;
      }
      // 请求api
      this.userRegister(this.user).then(() => {
        // 跳转首页
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        this.$bvToast.toast(err.data.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
      // const api ="http://localhost:8080/api/auth/register"
      // ,{Headers:{authorations:"token.."}}
      
      // this.axios.post(api,{...this.user}).then(res=>{
     /*  userService.register(this.user).then((res)=>{
      //保存token
      this.SET_TOKEN(res.data.data.token)
      // this.$store.commit('userModle/SET_TOKEN',res.data.data.token)
      // storageServie.set(storageServie.USER_TOKEN,res.data.data.token)
      return userService.info()
      }).then(response=>{
      //保存用户信息   序列化JSON.stringify
      this.SET_USERINFO(response.data.data.user)
      // this.$store.commit('userModle/SET_USERINFO',response.data.data.user)
      // storageServie.set(storageServie.USER_INFO,JSON.stringify(response.data.data.user))
        
      // localStorage.setItem('token',res.data.data.token)
      //跳转主页
      this.$router.replace({name:'home'})
      
      }).catch(err=>{
        // console.log('err',err.response.data.msg)
      //前台提示
        this.$bvToast.toast(err.response.data.msg, {
          title:  '数据验证错误',
          variant: "danger",
          solid: true
        })
      
      }) */
      console.log("register")
    }
  }
};
</script>

<style>
</style>