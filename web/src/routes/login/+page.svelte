<script>

</script>

<style>
    #login {
        display: grid;
        grid-template-columns: auto 500px auto;
        grid-template-rows: 50px auto;
        grid-template-areas:
            ". title ."
            ". register .";
        justify-content: center;
        align-items: center;
    }

    #login form {
        grid-area: register;
    }

    #login h1 {
        grid-area: title;
        text-align: center;
        color: coral;
    }

</style>

<div id="login">
    <h1>Login</h1>
    <form action="">
        <label for="email">Email</label>
        <input type="email">
        <label for="password">Password</label>
        <input type="password">
        <button type="submit">Login</button>
    </form>
</div>