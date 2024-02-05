use std::ops::Mul;
use std::sync::{Arc, Mutex};
use std::thread;
use ark_bn254::{G1Projective as G1, G1Affine, Fr};
use ark_ec::{AffineRepr, CurveGroup};

fn main() {
    let g1: G1 = G1Affine::generator().into();
    let expected = "(21415159568991615317144600033915305503576371596506956373206836402282692989778, 8573070896319864868535933562264623076420652926303237982078693068147657243287)";
    let expected = Arc::new(expected.to_string());
    let num_threads = 8;
    let range_per_thread = 144000000 / num_threads;
    let result = Arc::new(Mutex::new(None));

    let mut handles = vec![];

    for thread_id in 0..num_threads {
        let g1_clone = g1.clone();
        let expected_clone = Arc::clone(&expected);
        let result_clone = Arc::clone(&result);
        let handle = thread::spawn(move || {
            let start = thread_id * range_per_thread;
            let end = start + range_per_thread;
            for i in start..end {
                let tau = Fr::from(i as u64);
                let tg = g1_clone.mul(tau).into_affine();
                if tg.to_string() == *expected_clone {
                    let mut result = result_clone.lock().unwrap();
                    *result = Some((tau, tg.to_string()));
                    break;
                }
            }
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    let result_lock = result.lock().unwrap();
    if let Some((tau, tg)) = &*result_lock {
        println!("Found tau value: {:?}", tau);
        println!("tg: {:?}", tg);
    } else {
        println!("No matching tau value found.");
    }
}
