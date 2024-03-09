import http from "k6/http";
import {check, sleep} from "k6";

export const options = {
    vus: 100,
    iterations:2000,
};

function generateAdvertisement(vu,iter){
    return {
        title: "Title vu : " + vu.toString() + " iter :" + iter.toString(),
        startAt: new Date().toISOString(),
        endAt: new Date(new Date().getTime() + 60 * 60 * 1000).toISOString(),
        condition: {
            ageStart: 20,
            ageEnd: 30,
            country: ["TW", "JP"],
            platform: ["android", "ios"],
            gender: ["M", "F"]
        }
    }
}

export default function () {
    let requestUrl = __ENV.REQUEST_URL;

    if (requestUrl===undefined){
        requestUrl="http://127.0.0.1";
    }

    const url = requestUrl+"/api/v1/ad"

    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };


    const response = http.post(url, JSON.stringify(generateAdvertisement(__VU,__ITER)), params);


    check(response, {
        "Status is 200": (res) => res.status === 200,
    });
    console.log(response.body)

    sleep(1);
}