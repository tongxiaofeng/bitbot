function log(string){
    console.log(string)
}

function main() {
	log("Running second strategy")
    var  before =  Date.now();

    for(i = 0;i<100;i++){
        console.log("+++++++  :" + i);
    }
    var after =  Date.now();

    var diff = after - before;
    log("Time used:" + diff + " ms")

    var version = Version();

    log("Docker Version:");
    log(version);
    // log("Second strategy started sleeping");
    // var  before =  Date.now();
    // Sleep(3000);
    // var after =  Date.now();

    // var diff = after - before;
    // log("Time slept:" + diff + " ms")

    _.each([1, 2, 3,4,5,6,7], log);
    Log("this is a test Log in second strategy")
    //Log("Sin:", Sin([1.5,2.0]))
    log("finish main in second strategy")
}


function onTick() {
    // 策略采用轮询而非事件驱动是因为作者喜欢对代码100%的掌控力.
}

function onExit(){
    log("exiting second strategy from js");
}

function onError(){
    
}