import org.antlr.v4.runtime.{CharStreams, CommonTokenStream}

object Main{
  def main(args: Array[String]): Unit = {
    println("Hello world")
    val charStream = CharStreams.fromFileName("/Users/pandurang/projects/golang/helloworld/hello.go")
    val lexer = new GoLexer(charStream)
    val tokenStream = new CommonTokenStream(lexer)
    val parser = new GoParser(tokenStream)
    val programCtx = parser.sourceFile()
    println(AstPrinter.print(programCtx))
  }
}