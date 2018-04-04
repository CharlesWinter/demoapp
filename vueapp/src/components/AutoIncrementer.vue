<template>
  <div>
    <p>My current count is {{ currentCount }}</p>
  </div>
</template>


<script>
export default {
  name: "AutoInc",
  data() {
    return {
      currentCount : 0
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
