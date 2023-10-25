<script>
    import '../../static/css/main.css'
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
        grid-template-columns: 300px auto 120px;
        grid-template-rows: 80px;
        grid-template-areas:
        "title . navigation";
        align-content: center;
    }
    #navbar #logo {
        grid-area: title;
        display: flex;
        align-items: center;
    }

    #logo a {
        color: #E9724C;
    }

    #logo img{
        width: 30px;
        margin: 0 10px 0 10px;
    }

    #navbar ul {
        grid-area: navigation;
        display: flex;
        list-style-type: none;
        align-items: center;

    }

    #navbar li {
        padding-right: 10px;
    }

    #navbar li a:hover {
        color: #E9724C;
    }

    #navbar a{
        text-decoration: none;
    }
</style>