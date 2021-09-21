<template>

    <nav>
        <router-link to="/" class="navitem">Fuel</router-link>
        <router-link to="/History" class="navitem">History</router-link>
        <router-link to="/Optimise" class="navitem">Optimise</router-link>
        <router-view></router-view>
    </nav>

</template>

<script>

export default {
    name: 'Navbar'
}
</script>

<style scoped>
    nav {
        justify-content: space-between;
        align-items: left;
        margin-bottom: 20px;
        padding-top: 3%;
        padding-left: 1%;
        padding-bottom: 3%;
    }
    .navitem {
        padding-top: 20px;
        padding-bottom: 20px;
        padding-right: 15px;
        padding-left: 15px;
        text-decoration: none;
        font-size: 20px;
        border-right: 0.5px solid #001833;
        border-top: 0.5px solid #001833;
        border-radius:10px 10px 0px 0px;
    }
    .router-link-exact-active{
        background: #000913;
        z-index:-1;
        color: #00849b;
        box-shadow:0 0 20px rgba(0, 0, 0, 0.8);
        top:10px;
        bottom:0;
        left:0;
        right:10px;

        
    }

</style>