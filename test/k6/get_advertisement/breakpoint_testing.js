import http from "k6/http";
import { sleep,check } from "k6";

export const options={
    stages: [
        { duration: "1h", target: 100000 }
    ]
}

function getRandomAge(){
    const ages = [
        {name: Array.from({ length: 24 }, (_, i) => 1+i), distribution: 19.13},
        {name: Array.from({ length: 10 }, (_, i) => 25+i), distribution: 30.44},
        {name: Array.from({ length: 10 }, (_, i) => 35+i), distribution: 22.93},
        {name: Array.from({ length: 10 }, (_, i) => 45+i), distribution: 14.14},
        {name: Array.from({ length: 10 }, (_, i) => 55+i), distribution: 8.3},
        {name: Array.from({ length: 36 }, (_, i) => 65+i), distribution: 5.05},
    ];
    const randomNum = Math.random() * 100;
    let sum = 0;

    for (const age of ages) {
        sum += age.distribution;
        if (randomNum <= sum) {
            return age.name[Math.floor(Math.random() * age.name.length)]
        }
    }
    return ages[ages.length-1].name[Math.floor(Math.random() * ages[ages.length-1].name.length)]
}
function getRandomCountry(){
    const countries = [
        {name: 'TW', distribution: 92.88},
        {name: 'HK', distribution: 3.11},
        {name: 'US', distribution: 1.79},
        {name: 'MY', distribution: 0.39},
        {name: 'AU', distribution: 0.25},
        {name: 'OTHER', distribution: 1.58},
    ];
    const otherCountries=["AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM", "CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "SZ", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MG", "MW", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MK", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SE", "CH", "SY", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB" , "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"]
    const randomNum = Math.random() * 100;
    let sum = 0;

    for (const country of countries) {
        sum += country.distribution;
        if (randomNum <= sum) {
            if (country.name==="OTHER"){
                return otherCountries[Math.floor(Math.random() * otherCountries.length)]
            }
            return country.name;
        }
    }
}
function getRandomGender(){
    const genders = [
        {name: 'M', distribution: 57.77},
        {name: 'F', distribution: 42.23},
    ];
    const randomNum = Math.random() * 100;
    let sum = 0;

    for (const gender of genders) {
        sum += gender.distribution;
        if (randomNum <= sum) {
            return gender.name;
        }
    }
}
function getRandomPlatform(){
    const platforms = [
        {name: 'web', distribution: 44.11},
        {name: 'ios', distribution: 33},
        {name: 'android', distribution: 22.89}
    ];
    const randomNum = Math.random() * 100;
    let sum = 0;

    for (const platform of platforms) {
        sum += platform.distribution;
        if (randomNum <= sum) {
            return platform.name;
        }
    }
}
function getRandomOffset(){
    return Math.floor(Math.random() * 1000)
}
function getRandomLimit(){
    return Math.floor(Math.random() * 100)+1
}

export default function () {
    var requestUrl = __ENV.REQUEST_URL;

    if (requestUrl===undefined){
        requestUrl="http://127.0.0.1"
    }

    const url = requestUrl+"/api/v1/ad"+"?"
        +"age="+getRandomAge()+"&"
        +"country="+getRandomCountry()+"&"
        +"gender="+getRandomGender()+"&"
        +"platform="+getRandomPlatform()+"&"
        +"offset="+getRandomOffset()+"&"
        +"limit="+getRandomLimit()

    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };

    const response = http.get(url, params);

    check(response, {
        "Status is 200": (res) => res.status === 200,
    });

    sleep(1);
}