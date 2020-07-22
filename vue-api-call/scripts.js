const app = new Vue({
    el: '#main',

    data: {
        result: ` `,
        responseAvailable: false,
    },
    methods: {
        fetchAPIData() {
            this.responseAvailable = false;
            fetch("http://localhost:8080/", {
                "method": "GET"
            })
            .then(response => {
                if(response.ok) {
                    return response
                } else {
                    alert("Server returned " + response.status + " : " +
                    response.statusText);
                }
            })
            .then(response => {
                this.result = response.body
                this.responseAvailable = true;
            })
            .catch(err => {
                console.log(err);
            });
        }
    }
})