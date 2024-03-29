extern crate num_cpus;

use std::thread;
use std::time::SystemTime;

fn loops_loop(len: u32, tag: u32) {
    //获取最大循环
    let maxloops: u32 = 2u32.pow(len);
    //获取CPU数量
    let cpus: u32 = num_cpus::get() as u32;
    //获取每个CPU均分的数量
    let perloop: u32 = maxloops / cpus;
    //获取剩余数量
    let addloop: u32 = maxloops % cpus;
    println!("Max Loop:{maxloops}\nHave CPUs:{cpus}\nPerLoop:{perloop}\nAddLoop:{addloop}\n");

    let mut start: u32 = 0;
    let mut handles = vec![];
    for cpu in 0..cpus {
        let mut end = (cpu + 1) * perloop;
        if cpu + 1 == cpus {
            end += addloop;
        }
        let handle = thread::spawn(move || {
            println!("Cpu:{cpu} Start:{start} End:{end}");
            // here change to loop !
            //for cusor in start..end {
            let mut cusor = start;
            loop {
                let mut n = cusor;
                let mut sum = 0;
                let mut nmb = 1;
                while n > 0 {
                    if n & 1 == 1 {
                        sum += nmb;
                    }
                    n >>= 1;
                    nmb += 1;
                }
                if sum == tag {
                    println!("Got Number is :{:?}", sum);
                    break;
                }
                cusor += 1;
                if cusor & end == end {
                    break;
                }
            }
        });
        start += perloop;
        handles.push(handle);
    }
    for handle in handles {
        handle.join().unwrap();
    }
}

fn loops(len: u32, tag: u32) {
    //获取最大循环
    let maxloops: u32 = 2u32.pow(len);
    //获取CPU数量
    let cpus: u32 = num_cpus::get() as u32;
    //获取每个CPU均分的数量
    let perloop: u32 = maxloops / cpus;
    //获取剩余数量
    let addloop: u32 = maxloops % cpus;
    println!("Max Loop:{maxloops}\nHave CPUs:{cpus}\nPerLoop:{perloop}\nAddLoop:{addloop}");

    let mut start: u32 = 0;
    let mut handles = vec![];
    for cpu in 0..cpus {
        let mut end = (cpu + 1) * perloop;
        if cpu + 1 == cpus {
            end += addloop;
        }
        let handle = thread::spawn(move || {
            println!("Cpu:{cpu} Start:{start} End:{end}");
            for cusor in start..end {
                let mut n = cusor;
                let mut sum = 0;
                let mut nmb = 1;
                while n > 0 {
                    if n & 1 == 1 {
                        sum += nmb;
                    }
                    n >>= 1;
                    nmb += 1;
                }
                if sum == tag {
                    println!("Got {sum}");
                    break;
                }
            }
        });
        start += perloop;
        handles.push(handle);
    }
    for handle in handles {
        handle.join().unwrap();
    }
}

fn empty_loop() {
    let max = 2u32.pow(20);
    let mut _v = 0;
    for i in 0..max {
        _v = i;
    }
}

fn main() {
    let sy_time = SystemTime::now();

    println!("函数运行开始！");
    //loops_loop(30, 900);
    loops(30, 900);
    //empty_loop();
    let mut t: f64 = sy_time.elapsed().unwrap().as_micros() as f64;
    t = t / 1000000.0;
    println!("函数运行时间: {:.}秒", t);
}
