<div id="landing-page">
    <div id="tagline-content">
        <div>
            <h1>Crafting Connections<br>with Code:<br>Your Social Experience,<br>Our Expertise</h1>
            <h2><a href="#project-about">Learn More</a></h2>
        </div>
    </div>
    <div id="project-about">
        <h1>About the Project</h1>
        <p>
            Torch Social is a web development project that I'm particularly proud of, as it showcases my skills in web
            development. Specifically, it highlights my use of Svelte for the frontend and my proficiency in creating a
            custom API from scratch using Go with the Echo framework, as well as database management with MongoDB.
        </p>
    </div>
    <div id="project-stack">
        <div>
            <h2 style="grid-area: stack-title">Svelte-Powered Frontend</h2>
            <p style="grid-area: stack-body">
                At the core of Torch Social is a responsive and dynamic user interface built with Svelte. Svelte's
                simplicity and reactivity make it an ideal choice for creating a smooth and engaging user experience
                tailored for desktop use. With Svelte, I've crafted an intuitive and highly responsive platform that
                ensures a user-friendly experience.
            </p>
        </div>
        <div>
            <h2 style="grid-area: stack-title">Custom API using Echo</h2>
            <p style="grid-area: stack-body">
                The backend of Torch Social is powered by a custom API that I've developed from the ground up using the
                Go programming language, along with the Echo framework. This custom API allows for real-time
                interactions, user authentication, and various features essential for a social media platform. It
                ensures a seamless user experience, and because it's custom-built, it offers a high level of control and
                customization.
            </p>
        </div>
        <div>
            <h2 style="grid-area: stack-title">MongoDB Database</h2>
            <p style="grid-area: stack-body">
                Behind the scenes, Torch Social relies on MongoDB for robust and scalable database management. MongoDB,
                a NoSQL database, allows for the efficient storage and retrieval of user data. My database management
                skills guarantee data security and reliability, ensuring that user interactions are both smooth and
                responsive.
            </p>
        </div>
    </div>
</div>

<style>
    #landing-page {
        margin: 0 300px;
    }
    #landing-page p {
        font-size: 18px;
    }
    #tagline-content {
        /*display: grid;*/
        /*grid-template-columns: 50% 50%;*/
        /*grid-template-rows: auto;*/
        /*grid-template-areas:*/
        /*    "tagline learn-more";*/
        /*align-items: center;*/
        /*justify-items: center;*/
        height: 100vh;
        font-size: 32px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    #tagline-content h1 {
        padding-bottom: 20px;
    }

    #tagline-content h2 a {
        background-color: #E9724C;
        text-decoration: none;
        padding: 15px;
        border-radius: 10px;
        font-size: 24px;

    }
    #project-about {
        padding-bottom: 30px;
        text-align: center;
    }
    #project-about h1 {
        color: #E9724C;
        padding-bottom: 10px;
    }
    #project-stack {
        display: grid;
        grid-template-columns: auto auto auto;
        grid-template-rows: 30px auto;
        grid-template-areas:
            "stack-title stack-title stack-title"
            "stack-body stack-body stack-body";
        text-align: center;
    }
    #project-stack h2 {
        color: #FFC857;
        padding: 0 10px 10px 10px;
    }
    #project-stack p {
        padding: 0 25px;
        margin-bottom: 400px;
    }


</style>