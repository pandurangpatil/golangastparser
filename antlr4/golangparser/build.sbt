name := "golangparser"

scalaVersion       := "2.13.8"
crossScalaVersions := Seq("2.13.8", "3.2.2")

libraryDependencies ++= Seq(
  "org.scalatest" %% "scalatest"         % "3.2.15" % Test,
  "org.antlr"      % "antlr4-runtime"    % "4.7"
)

scalacOptions ++= Seq(
  "-deprecation" // Emit warning and location for usages of deprecated APIs.
)

// enablePlugins(JavaAppPackaging)
enablePlugins(JavaAppPackaging, LauncherJarPlugin, Antlr4Plugin)

 Antlr4 / antlr4PackageName := Some("io.joern.gosrc2cpg.parser")
 Antlr4 / antlr4Version     := "4.7"
 Antlr4 / antlr4GenVisitor  := true
 Antlr4 / javaSource        := (Compile / sourceManaged).value
