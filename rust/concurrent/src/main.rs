use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let mut x: i32 = 1;
    for _ in 0..10 {
        thread::spawn(|| {
            x = x + 1;
        });
    }
    println!("The value of x is: {}", x);
}