<template>
  <div id="app">
    <img src="./assets/logo.png">
    <HelloWorld/>
    <ClickButton/>
    <ClickButton/>
    <BlogPost title="This is a new post"/>
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld";
import ClickButton from "./components/ClickButton";
import BlogPost from "./components/PostTitle";

export default {
  name: "App",
  components: {
    HelloWorld,
    ClickButton,
    BlogPost
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
