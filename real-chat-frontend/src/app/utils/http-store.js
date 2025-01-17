let area = ""
let baseURL = ""


function CheckENV() {
    return process.env.NEXT_PUBLIC_ENV_AREA ?? "dev"
}


function SetUpUrlStore() {
    if (area === "") {
        area = CheckENV()
    }
    console.log("from setting up URL Store", area)
    switch (area) {
        case "production":
            return "http://backend:80/"
        default:
            return "http://127.0.0.1:80/" // dev envirement on local machine
    }
}

export function GetURL(path) {
    if (baseURL === "") {
        baseURL = SetUpUrlStore()
    }

    if (path && typeof path === "string") {
        return `${baseURL}${path}`
    }

    return baseURL    
}



