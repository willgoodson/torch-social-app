<script>
    
</script>

<style>
    #register {
        display: grid;
        grid-template-columns: auto 500px auto;
        grid-template-rows:  50px auto;
        grid-template-areas:
            ". title ."
            ". register .";
        justify-content: center;
        align-items: center;
    }

    #register form {
        grid-area: register;
    }

    #register h1 {
        grid-area: title;
        text-align: center;
        color: coral;
    }
</style>

<div id="register">
    <h1>Register</h1>
    <form action="">
        <label for="f_name">First Name</label>
        <input type="text">
        <label for="l_name">Last Name</label>
        <input type="text">
        <label for="email">Email</label>
        <input type="email">
        <label for="password">Password</label>
        <input type="password">
        <button type="submit">Register</button>
    </form>
</div>