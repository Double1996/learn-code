package org.javacore.io;

import java.io.BufferedReader;
import java.io.InputStreamReader;


public class OSExecuteT {
    public static void command(String command) {
        boolean err = false;

        try {
            // 创建操作系统进程
            Process process = new ProcessBuilder(command.split(" ")).start();

            // 读取进程的输入流
            BufferedReader results = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String s;
            while ((s = results.readLine()) != null)
                System.out.println(s);

            // 读取进程的错误流
            BufferedReader errors = new BufferedReader(new InputStreamReader(process.getErrorStream()));
            while ((s = errors.readLine()) != null) {
                System.err.println(s);
                if (!err)
                    err = true;
            }

        } catch (Exception e) {
            if (!command.startsWith("CMD /C"))
                command("CMD /C " + command);
            else
                throw new RuntimeException(e);
        }

        if (err)
            throw new OSExecuteException("Errors Executing " + command);
    }

    public static void main(String[] args) {
        command("java -version");
    }
}

class OSExecuteException extends RuntimeException {
    private static final long serialVersionUID = -3254218368058055219L;

    public OSExecuteException(String msg) {
        super(msg);
    }
}
