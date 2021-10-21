<script>
    const billingEndpoint = "/v1/shopify/billing/create";

    const Loading = window["app-bridge"].actions.Loading;
    const loading = Loading.create(window["app"]);

    const Redirect = window["app-bridge"].actions.Redirect;
    const redirect = Redirect.create(window["app"]);

    const Toast = window["app-bridge"].actions.Toast;
    const toastOptions = {
        message: "Couldn't redirect :(",
        duration: 5000,
    };
    const toast = Toast.create(window["app"], toastOptions);

    const handleClick = () => {
        loading.dispatch(Loading.Action.START);

        fetch(`${billingEndpoint}?shop=${window["shop"]}`)
            .then((response) => response.json())
            .then((json) => {
                redirect.dispatch(Redirect.Action.REMOTE, json.return_url);
            })
            .catch(() => toast.dispatch(Toast.Action.SHOW)) 
            .finally(() => loading.dispatch(Loading.Action.STOP));
    };
</script>

<main class="container">
    <div class="row">
        <div class="col-md-12">
            <div class="card">
                <div class="card-body">
                    <h3>Welcome</h3>

                    <p>This app is ....</p>
                    <p>TODO</p>
                    <p>We need a subscription because ..</p>
                </div>
                <div class="card-action">
                    <a
                        on:click|preventDefault={handleClick}
                        href="#"
                        class="btn btn-primary">Create Subscription</a
                    >
                </div>
            </div>
        </div>
    </div>
</main>
