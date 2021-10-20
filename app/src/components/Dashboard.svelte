<script>
    const updateStatusEndpoint = "/v1/should-render";

    export let extensionEnabled;

    const handleStatusToggle = () => {
        fetch(updateStatusEndpoint, {
            method: "PATCH",
            body: JSON.stringify({
                shopURL: window['shop'],
                shouldRender: !extensionEnabled
            }),
            headers: {
                "Content-type": "application/json",
            },
        })
            .then((response) => {
                extensionEnabled = !extensionEnabled
            })
            .catch(e => console.error(e))
    };
</script>

<main class="container">
    <div class="row no-gutters">
        <div class="col-md-4 card-annotation">
            <h3>Extension Status</h3>
        </div>
        <div class="col-md-8">
            <div class="card">
                <div class="card-body">
                    <div class="statusInnerContainer">
                        <p>
                            The extension <strong
                                >is{#if !extensionEnabled}&nbsp;not{/if} showing</strong
                            > in your checkout pages
                        </p>
                        <div>
                            <a
                                href="#"
                                on:click|preventDefault={handleStatusToggle}
                                class="btn btn-primary"
                                >{#if extensionEnabled}Disable{:else}Enable{/if}</a
                            >
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="row no-gutters">
        <div class="col-md-4 card-annotation">
            <h3>KPI 1</h3>
            <p>KPI 1 desc</p>
        </div>
        <div class="col-md-8">
            <div class="card">
                <div class="card-body">
                    <h3>KPI 1</h3>
                    <p>KPI 1</p>
                </div>
            </div>
        </div>
    </div>
    <div class="row no-gutters">
        <div class="col-md-4 card-annotation">
            <h3>KPI 2</h3>
            <p>KPI 2 desc</p>
        </div>
        <div class="col-md-8">
            <div class="card">
                <div class="card-body">
                    <h3>KPI 2</h3>
                    <p>KPI 2</p>
                </div>
            </div>
        </div>
    </div>
</main>

<style>
    .statusInnerContainer {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .statusInnerContainer p {
        margin-bottom: 0px;
    }
</style>
