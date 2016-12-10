#![feature(convert)]
use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::Regex;

struct Program { pc: u32, registers: (u32, u32), }
#[derive(Clone)]
enum InstructionType { Hlf, Tpl, Inc, Jmp, Jie, Jio, }
#[derive(Clone)]
struct Instruction {
    register: Option<String>,
    argument: Option<i32>,
    instr: InstructionType,
}

impl Instruction {
    fn eval(&self, program: Program) -> Program {
        match self.instr {
            InstructionType::Hlf => self.eval_hlf(program),
            InstructionType::Tpl => self.eval_tpl(program),
            InstructionType::Inc => self.eval_inc(program),
            InstructionType::Jmp => self.eval_jmp(program),
            InstructionType::Jie => self.eval_jie(program),
            InstructionType::Jio => self.eval_jio(program),
        }
    }

    fn eval_hlf(&self, program: Program) -> Program {
        let register = self.clone().register.unwrap();
        let registers = match register.as_str() {
            "a" => (program.registers.0 / 2, program.registers.1),
            "b" => (program.registers.0, program.registers.1 / 2),
            _ => panic!(""),
        };
        let pc = program.pc + 1;
        Program { pc: pc, registers: registers }
    }

    fn eval_tpl(&self, program: Program) -> Program {
        let register = self.clone().register.unwrap();
        let registers = match register.as_str() {
            "a" => (program.registers.0 * 3, program.registers.1),
            "b" => (program.registers.0, program.registers.1 * 3),
            _ => panic!(""),
        };
        let pc = program.pc + 1;
        Program { pc: pc, registers: registers }
    }

    fn eval_inc(&self, program: Program) -> Program {
        let register = self.clone().register.unwrap();
        let registers = match register.as_str() {
            "a" => (program.registers.0 + 1, program.registers.1),
            "b" => (program.registers.0, program.registers.1 + 1),
            _ => panic!(""),
        };
        let pc = program.pc + 1;
        Program { pc: pc, registers: registers }
    }

    fn eval_jmp(&self, program: Program) -> Program {
        let jump = self.clone().argument.unwrap();
        let pc = self.compute_jump(program.pc, jump);
        Program { pc: pc, registers: program.registers, }
    }

    fn eval_jie(&self, program: Program) -> Program {
        let instr = self.clone();
        let (register, jump) = (instr.register.unwrap(), instr.argument.unwrap());
        let value = match register.as_str() {
            "a" => program.registers.0,
            "b" => program.registers.1,
            _   => panic!(""),
        };
        let pc = if value % 2 == 0 {
            self.compute_jump(program.pc, jump)
        } else {
            program.pc + 1
        };
        Program { pc: pc, registers: program.registers, }
    }

    fn eval_jio(&self, program: Program) -> Program {
        let instr = self.clone();
        let (register, jump) = (instr.register.unwrap(), instr.argument.unwrap());
        let value = match register.as_str() {
            "a" => program.registers.0,
            "b" => program.registers.1,
            _   => panic!(""),
        };
        let pc = if value == 1 {
            self.compute_jump(program.pc, jump)
        } else {
            program.pc + 1
        };
        Program { pc: pc, registers: program.registers, }
    }

    fn compute_jump(&self, pc: u32, jump: i32) -> u32 {
        if jump < 0 {
            pc - ((-jump) as u32)
        } else {
            pc + (jump as u32)
        }
    }
}

fn parse_jump(jump: String) -> i32 {
    let first = &jump[0..1];
    let rest = &jump[1..];
    let value = i32::from_str_radix(rest, 10).unwrap();
    if first == "-" {
        -value
    } else {
        value
    }
}

fn parse_input(line: &String) -> Instruction {
    let re = Regex::new(r"(\w+) ([^, ]+),? ?(.*)").unwrap();
    let caps = match re.captures(line) {
        Some(c) => c,
        None    => panic!("could not parse input"),
    };
    let instr_str = caps.at(1).unwrap();
    let (register, argument) = match instr_str {
        "hlf" | "tpl" | "inc" => (Some(caps.at(2).unwrap().to_string()), None),
        "jmp" => (None,
                  Some(parse_jump(caps.at(2).unwrap().to_string()))),
        "jie" | "jio" => (Some(caps.at(2).unwrap().to_string()),
                          Some(parse_jump(caps.at(3).unwrap().to_string()))),
        _ => panic!("could not parse input"),
    };
    let instr = match instr_str {
        "hlf" => InstructionType::Hlf,
        "tpl" => InstructionType::Tpl,
        "inc" => InstructionType::Inc,
        "jmp" => InstructionType::Jmp,
        "jie" => InstructionType::Jie,
        "jio" => InstructionType::Jio,
        _ => panic!("could not parse input"),
    };
    Instruction {
        register: register,
        argument: argument,
        instr: instr,
    }
}

fn main() {
    let stdin = io::stdin();
    let instructions: Vec<Instruction> =
        stdin
        .lock()
        .lines()
        .map(|l| parse_input(&l.unwrap()))
        .collect();
    let mut program = Program { pc: 0, registers: (1, 0), };
    loop {
        println!(
            "pc: {}, registers: ({}, {})",
            program.pc,
            program.registers.0,
            program.registers.1,
        );
        let pc = program.pc;
        if pc >= instructions.len() as u32 {
            break;
        }

        let instr = instructions.get(pc as usize).unwrap().clone();
        program = instr.eval(program);
    }
}
