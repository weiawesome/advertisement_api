import http from "k6/http";
import {check, sleep} from "k6";
import { Counter } from 'k6/metrics';

const CounterError = new Counter('Response_Day_Limit_Error');
const CounterSuccess = new Counter('Response_Success');

export const options = {
    vus: 100,
    iterations:3500,
};

function getRandomChoice(){
    return Math.random()>0.5
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
        return age.name[0],age.name[age.name.length-1]
    }
    let lastAge=ages[ages.length-1]
    return lastAge.name[0],lastAge.name[lastAge.name.length-1]
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

function generateAdvertisement(vu,iter){
    let condition={};
    let map={};
    let val="";
    let ageStart=0,ageEnd=0;

    if (getRandomChoice()){
        ageStart,ageEnd=getRandomAge()
        condition["age_start"]=ageStart
        condition["age_end"]=ageEnd
        while (getRandomChoice() && !map[ageStart] && !map[ageEnd]){
            condition["age_start"]=Math.min(ageStart,condition["age_start"])
            condition["age_end"]=Math.min(ageEnd,condition["age_end"])
            map[ageStart]=true
            map[ageEnd]=true
            ageStart,ageEnd=getRandomAge()
        }
    }
    if (getRandomChoice()){
        val=getRandomCountry()
        condition["country"]=val
        while (getRandomChoice() && !map[val]){
            condition["country"]=val
            map[val]=true
            val=getRandomCountry()
        }
    }
    if (getRandomChoice()){
        val=getRandomGender()
        condition["gender"]=val
        while (getRandomChoice() && !map[val]){
            condition["gender"]=val
            map[val]=true
            val=getRandomGender()
        }
    }
    if (getRandomChoice()){
        val=getRandomPlatform()
        condition["platform"]=val
        while (getRandomChoice() && !map[val]){
            condition["platform"]=val
            map[val]=true
            val=getRandomPlatform()
        }
    }

    return {
        "title": "Title vu : " + vu.toString() + " iter :" + iter.toString(),
        "startAt": new Date(new Date().getTime()+(iter%500)*(60 * 60 * 1000)).toISOString(),
        "endAt": new Date(new Date().getTime()+(iter%500)*(60 * 60 * 1000)+(30 * 60 * 1000)).toISOString(),
        "condition": condition
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
        "Status is 200": (res) => res.status===200
    });

    if(response.status === 200) {
        CounterSuccess.add(1);
    } else if(response.status === 429) {
        CounterError.add(1);
    }

    sleep(1);
}