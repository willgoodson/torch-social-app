<script>
    import '$lib/css/main.css'
    import icon from '$lib/assets/torch-icon-color.svg'
</script>
<svelte:head>
    <title>Torch Social</title>
</svelte:head>

<div id="navbar">
    <div id="logo">
        <img src={icon} alt="torch">
        <h1><a href="/">Torch Social</a></h1>
    </div>
    <ul>
        <li><a href="/register">Register</a></li>
        <li><a href="/login">Login</a></li>
    </ul>
</div>
<slot />


<style>
    /* NavBar Styling */
    #navbar {
        display: grid;
        grid-template-columns: 300px auto 300px;
        grid-template-rows: 80px;
        grid-template-areas:
        "title . navigation";
        align-content: center;
    }
    #navbar #logo {
        grid-area: title;
        display: flex;
        align-items: center;
        margin-left: 20px;
    }

    #logo a {
        color: #E9724C;
    }

    #logo img{
        width: 30px;
        padding-right: 10px;
    }

    #navbar ul {
        grid-area: navigation;
        list-style-type: none;
        padding: 0;
        text-align: right;
        display: flex;
        justify-content: flex-end;
        align-items: center;
        margin-right: 20px;
    }

    #navbar li:not(:last-child) {
        padding-right: 10px;
    }

    #navbar li a:hover {
        color: #E9724C;
      text-align: right;
    }

    #navbar a{
      text-align: right;
        text-decoration: none;
    }
</style>
