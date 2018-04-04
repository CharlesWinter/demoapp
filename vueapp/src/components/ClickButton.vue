<template>
  <div>
    <button v-on:click="incrementCount">The count is {{count}}</button>
  </div>
</template>


<script>
export default {
  name: "ClickButton",
  data() {
    return {
      count : 0
    };
  },
  methods: {
    incrementCount : function() {
      this.count++ 

      

    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
