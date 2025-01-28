import java.io.UnsupportedEncodingException;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;

public class Main {
    public static void main(String[] args) throws UnsupportedEncodingException {

        String str = "\u0412\u043e\u043b\u044c\u043d\u0435\u043d\u0441\u043a\u043e\u0435";
        System.out.println(str);
        //convert Unicode to UTF8 format
        byte[] utf8Bytes = str.getBytes(StandardCharsets.UTF_8);;

        String converted = new String(utf8Bytes, StandardCharsets.UTF_8);

        System.out.println(str.equals(converted));
    }
}
