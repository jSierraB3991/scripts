fn main() {
    println!("Hello, world, with using cargo!");
    
    const SCORE_LIMIT :u32 = 100;
    let unmtable_variable = 5;
    let mut mutable_variable = 5;
    let shadowing_variable = 32;
    let shadowing_variable = shadowing_variable + 32;

    println!("mutable variable 1: {}", mutable_variable);
    mutable_variable = 6;
    println!("umutable: {}", unmtable_variable);
    println!("The score limit is: {}", SCORE_LIMIT);
    println!("mutable variable 2: {}", mutable_variable);
    println!("shadowing variable {}", shadowing_variable);

    // data types
    // numbers (i32)     : i8, i16, i32, i64, i128, isize <> u8, u16, u32, u64, u128 usize
    // decimal (f32)     : f32, f64
    // booleans (bool)   : true, false
    // characters (char) : Represents letter, single quotes
    //
    // Arrays ([type, length]),     let array: [u32, 3] = [ 1, 2, 3 ];
    // Tuples (type...),            let tuple: (bool, u16, u32) = ( true, 2, 3 ); let data = tuple.0;

    // Operators
    // mathematics      : +, -, *, /, %
}
