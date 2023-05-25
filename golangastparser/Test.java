import com.sun.jna.Library;
import com.sun.jna.Native;
public class Test {
    public interface Sample extends Library {
        String ExternallyCalled();
        int Add(int a,int b);
    }

    public static void main(String[] args){
        Sample sample = Native.load("/Users/pandurang/projects/golangastparser/golangastparser/lib-sample.dylib", Sample.class);

        System.out.println(sample.Add(10,20));
        String result = sample.ExternallyCalled();
        System.out.println("We are in java now...");
        System.out.println(result);
    }
}
