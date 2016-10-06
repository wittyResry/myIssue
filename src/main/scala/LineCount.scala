import org.apache.spark.{SparkContext, SparkConf}
import org.apache.spark.SparkContext._
object LineCount {
    def main(args: Array[String]) {
        val conf = new SparkConf().setAppName("Line Count Program")
        val sc = new SparkContext(conf)
        val lines = sc.textFile("/etc/passwd")
        println(lines.count)
    }
}
