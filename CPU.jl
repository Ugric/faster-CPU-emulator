function CPU(regSize, memorySize, outputend=false)
    registers::Matrix{Int64} = fill(0, (1,regSize))
    memory::Matrix{Int64} = fill(0, (1,ceil(Integer, memorySize)))
    PC = 1
    function exec()
        if memory[PC] == 1
            # add registers
            registers[memory[PC+3]+1] = registers[memory[PC+1]+1] + registers[memory[PC+2]+1]
            PC += 4
        elseif memory[PC] == 2
            # sub registers
            registers[memory[PC+3]+1] = registers[memory[PC+1]+1] - registers[memory[PC+2]+1]
            PC += 4
        elseif memory[PC] == 3
            # multiply registers
            registers[memory[PC+3]+1] = registers[memory[PC+1]+1] * registers[memory[PC+2]+1]
            PC += 4
        elseif memory[PC] == 4
            # divide registers
            registers[memory[PC+3]+1] = registers[memory[PC+1]+1] / registers[memory[PC+2]+1]
            PC += 4
        elseif memory[PC] == 5
            # output unicode character from register
            print(Char(registers[memory[PC+1]+1]))
            PC += 2
        elseif memory[PC] == 6
            # set memory from register
            memory[memory[PC+1]+1] = registers[memory[PC+2]+1]
            PC += 3
        elseif memory[PC] == 7
            # set register
            registers[memory[PC+1]+1] = memory[PC+2]
            PC += 3
        elseif memory[PC] == 8
            # set register from memory
            registers[memory[PC+1]+1] = memory[memory[PC+2]+1]
            PC += 3
        elseif memory[PC] == 9
            # sleep from register
            sleep(registers[memory[PC+1]+1]/1000)
            PC += 2
        elseif memory[PC] == 10
            # jump if less than
            if registers[memory[PC+1]+1] < registers[memory[PC+2]+1]
                PC = registers[memory[PC+3]+1]+1
            else
                PC += 4
            end
        elseif memory[PC] == 11
            # output number from register
            print(registers[memory[PC+1]+1])
            PC += 2
        elseif memory[PC] == 12
            # set memory (index in a register) from a register
            memory[registers[memory[PC+1]+1]+1] = registers[memory[PC+2]+1]
            PC += 3
        end
    end
    function run()
        if outputend
            println("----- START -----")
        end
       while true
            if memory[PC] == 0
                if outputend
                    println("-----  END  -----")
                    println("registers:", myCPU["registers"])
                    println("memory:", myCPU["memory"])
                end
                break
            end
            exec()
       end
    end
    function load(array::Array)
        for i in 1:length(array)
            memory[i] = array[i]
        end
    end
    return Dict("run" => run, "load" => load, "registers" => registers, "memory" => memory)
end

myCPU = CPU(8, 1000, true)
script = split(replace(replace((open(f->read(f, String), ARGS[1])), "\n"=>" "), "\r"=>""), " ")
code = Array{Int, 1}()

for i in 1:length(script)
    push!(code, parse(Int64, script[i]))
end

myCPU["load"](code)
myCPU["run"]()