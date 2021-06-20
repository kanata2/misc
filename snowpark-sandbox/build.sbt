scalaVersion := "2.12.13"

name := "snowpark-sandbox"
organization := "com.github.kanata2.misc"
version := "0.1"

resolvers += "OSGeo Release Repository" at "https://repo.osgeo.org/repository/release/"
libraryDependencies += "com.snowflake" % "snowpark" % "0.6.0"
libraryDependencies += "org.scala-lang.modules" %% "scala-parser-combinators" % "1.1.2"

Compile/console/scalacOptions += "-Yrepl-class-based"
Compile/console/scalacOptions += "-Yrepl-outdir"
Compile/console/scalacOptions += "repl_classes"
