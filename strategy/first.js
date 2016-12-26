function log(string){
    console.log(string)
}


log("from global of first context")
function onTick() {
    // 策略采用轮询而非事件驱动是因为作者喜欢对代码100%的掌控力.
}

function main() {
    // console.log("Working in First JS Strategy.");
    // for (i = 0; i < exchanges.length; i++) { 
    //     log("Exchange Name:",exchanges[0].GetName());
    //     log("Acct:" ,exchange.GetEnabledCurrencies())
    // }
    log(exchange.GetName())

    var ticker = exchange.GetTicker()
    log(ticker.High)


    // for(i = 0;i<100;i++){
    //     console.log("-------  :" + i);
    // }

    var moneys = exchange.GetEnabledPair();
    console.log(moneys)


    console.log("Finish  main First JS Strategy.");
}

function onExit(){
    log("Exiting first strategy.")
}

function onError(){
    
}