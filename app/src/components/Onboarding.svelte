<script>
    const billingEndpoint = "/v1/shopify/billing/create";

    const Loading = window["app-bridge"].actions.Loading;
    const loading = Loading.create(window["app"]);

    const Redirect = window["app-bridge"].actions.Redirect;
    const redirect = Redirect.create(window["app"]);

    const handleClick = () => {
        console.log(loading, Loading, Loading.Action.START);
        loading.dispatch(Loading.Action.START);

        fetch(`${billingEndpoint}?shop=${window["shop"]}`)
            .then((response) => response.json())
            .then((json) => {
                redirect.dispatch(Redirect.Action.REMOTE, json.return_url);
            })
            .catch((e) => console.error(e)); // TODO handle loading, error state
        //.finally(() => );
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
