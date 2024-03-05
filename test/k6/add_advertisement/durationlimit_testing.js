import http from "k6/http";
import { sleep,check } from "k6";

export const options = {
    scenarios: {
        contacts: {
            executor: "ramping-arrival-rate",
            startRate: 10,
            timeUnit: "1s",
            preAllocatedVUs: 10000,

            stages: [
                { target: 100, duration: "2s" },
                { target: 100, duration: "2m" },
                { target: 10, duration: "2s" },
            ],
        },
    },
};

function generateAdvertisement(){

}

export default function () {
    const requestUrl = __ENV.REQUEST_URL;

    const url = requestUrl+"/api/v1/ad"

    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };
    const payload = {
        title:"Title",
        start_at:"2023-12-10T03:00:00.000Z",
        end_at:"2025-12-31T16:00:00.000Z",
    };

    const response = http.post(url, JSON.stringify(payload), params);

    check(response, {
        "Status is 200": (res) => res.status === 200,
    });

    sleep(1);
}