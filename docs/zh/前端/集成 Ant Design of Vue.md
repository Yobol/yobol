# 集成 Ant Design of Vue

> [Ant Design of Vue](https://www.antdv.com/docs/vue/introduce)
>
> - [Get Started](https://www.antdv.com/docs/vue/getting-started)

本文介绍如何在 Vue 中引入 Ant Design 的 UI 组件库。

## 安装

```shell
# 使用 npm help install 查看 npm install 命令帮助文档
# --save 选项用于自动将依赖添加到 package.json 文件的 dependencies 中
# --save-dev 选项用于自动将依赖添加到 package.json 文件中的 devDependencies 中
# ant-design-vue 默认使用最新版本
npm install ant-design-vue@3.2.12 --save
```

## 组件

### 引入组件

在 Vue 中使用 AntD 组件之前，需要显式注册组件。AntD 官方提供了三种注册组件的方式。

#### 全局引入全部组件

> ui/src/main.ts

```TypeScript
import { createApp } from 'vue'
// 引入 AntD 所有组件
import Antd from 'ant-design-vue'
import App from './App'
// 引入 AntD 样式文件
import 'ant-design-vue/dist/antd.css'

const app = createApp(app)

app.use(Antd)

app.mount('#app')
```

目前项目还处于快速迭代中，因此使用该方式引入全部组件。待项目稳定后，可修改为引入所需组件。

#### 全局引入所需组件

> ui/src/main.ts

```TypeScript
import { createApp } from 'vue'
// 引入 AntD 所需组件
import { Button, message } from 'ant-design-vue'
import App from './App'

const app = createApp(app)

/* Automatically register components under Button, such as Button.Group */
app.use(Button)

app.mount('#app')
```

#### 局部引入所需组件

> ui/src/components/Test.vue

```Vue
<template>
  <a-button>Add</a-button>
</template>

<script>
  import { Button } from 'ant-design-vue'
  const ButtonGroup = Button.Group

  export default {
    components: {
        AButton: Button,
        AButtonGroup: ButtonGroup,
    },
  }
</script>
```

该方式只在当前组件中有效。

#### 选择组件

我们可以在 [Ant Design Compoments](https://www.antdv.com/components/overview/) 界面选择所需的组件，如 [Form Component](https://www.antdv.com/components/form/#Form-Component)，然后将官方提供的 `JavaScript` 或 `TypeScript` 示例代码拷贝到项目中。

以登录表单为例，官方提供的一种 `TypeScript` 代码实现：

```Vue
<template>
  <a-form
    :model="formState"
    name="normal_login"
    class="login-form"
    @finish="onFinish"
    @finishFailed="onFinishFailed"
  >
    <a-form-item
      label="Username"
      name="username"
      :rules="[{ required: true, message: 'Please input your username!' }]"
    >
      <a-input v-model:value="formState.username">
        <template #prefix>
          <UserOutlined class="site-form-item-icon" />
        </template>
      </a-input>
    </a-form-item>

    <a-form-item
      label="Password"
      name="password"
      :rules="[{ required: true, message: 'Please input your password!' }]"
    >
      <a-input-password v-model:value="formState.password">
        <template #prefix>
          <LockOutlined class="site-form-item-icon" />
        </template>
      </a-input-password>
    </a-form-item>

    <a-form-item>
      <a-form-item name="remember" no-style>
        <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
      </a-form-item>
      <a class="login-form-forgot" href="">Forgot password</a>
    </a-form-item>

    <a-form-item>
      <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button">
        Log in
      </a-button>
      Or
      <a href="">register now!</a>
    </a-form-item>
  </a-form>
</template>
<script lang="ts">
import { defineComponent, reactive, computed } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
interface FormState {
  username: string;
  password: string;
  remember: boolean;
}
export default defineComponent({
  components: {
    UserOutlined,
    LockOutlined,
  },
  setup() {
    const formState = reactive<FormState>({
      username: '',
      password: '',
      remember: true,
    });
    const onFinish = (values: any) => {
      console.log('Success:', values);
    };

    const onFinishFailed = (errorInfo: any) => {
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
});
</script>
<style>
#components-form-demo-normal-login .login-form {
  max-width: 300px;
}
#components-form-demo-normal-login .login-form-forgot {
  float: right;
}
#components-form-demo-normal-login .login-form-button {
  width: 100%;
}
</style>
```

我们将上述代码拷贝到 `ui/src/view/login/LoginView.vue` 文件中即可。

### 定制组件

## 主题

### 定制主题

### 动态主题

## 国际化