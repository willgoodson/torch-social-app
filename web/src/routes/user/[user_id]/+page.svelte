<script lang="ts">
    import { page } from '$app/stores'
    const user_id = parseInt($page.params.user_id);

</script>

<div>
    <h1>Hello {user_id}</h1>
</div>