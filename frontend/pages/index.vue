<template>
  <div class="bg-white rounded-3xl">
    <div class="container">
      <h1 class="text-3xl font-bold text-right">หน้าหลัก</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: "IndexPage",
  data: () => ({
    // posts: [],
  }),
  async asyncData({ params, $axios }) {
    const post = await $axios.$get(`https://www.mecallapi.com/api/users`);
    return { post };
  },
};
</script>
<style>
body {
  overflow-y: hidden;
}
</style>
