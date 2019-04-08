//** This file was code generated by got. DO NOT EDIT. ***

package welcome

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/pkg/resource"
)

func init() {
	resource.RegisterPath("/goradd/configure.html",
		func(ctx context.Context, buf *bytes.Buffer) (headers map[string]string, err error) {

			buf.WriteString(`<!DOCTYPE html>
<head>
<meta charset="utf-8"/>
<title>`)

			buf.WriteString(`Configuring the Database`)

			buf.WriteString(`</title>
</head>
<body>
`)

			buf.WriteString(`
<h1>Configuring the Database</h1>
<p>
	Goradd relies on your database(s) not only to get to data, but to understand its structure. For
	sql databases, it does this by reading the database schema. When NoSQL databases are supported,
	the structure will come from a configuration file.
</p>
<p>
	Using its knowledge of the structure of your database, the goradd code generator will
	create data models, data nodes, forms, and more to get your a working application to start customizing.
</p>

<p>
	Goradd currently only supports MySQL databases. Adapters to other databases are easy to write, so
	if you would like to see another database supported and are willing to help, open an issue
	at our github site.
</p>
<p>
	To configure your database, edit the goradd-project/config/db.go file. You will see directions there.
	If you want to run the examples code and tutorial, you should create a local copy of the goradd
	examples database. To do that, import the sql code found in the goradd/web/examples/db directory,
	which is reprinted below for your convenience.
</p>
<code>
`)

			buf.WriteString(`<p>-- phpMyAdmin SQL Dump<br>
-- version 4.7.6<br>
-- https://www.phpmyadmin.net/<br>
--<br>
-- Host: db<br>
-- Generation Time: Jan 24, 2019 at 11:47 AM<br>
-- Server version: 5.6.34<br>
-- PHP Version: 7.1.9</p><br>
<p>SET FOREIGN_KEY_CHECKS=0;<br>
SET SQL_MODE = &#34;NO_AUTO_VALUE_ON_ZERO&#34;;<br>
SET AUTOCOMMIT = 0;<br>
START TRANSACTION;<br>
SET time_zone = &#34;+00:00&#34;;</p><br>
<p>--<br>
-- Database: ` + "`" + `goradd` + "`" + `<br>
--</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `address` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `address` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `person_id` + "`" + ` int(10) UNSIGNED DEFAULT NULL,<br>
  ` + "`" + `street` + "`" + ` varchar(100) NOT NULL,<br>
  ` + "`" + `city` + "`" + ` varchar(100) DEFAULT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `address` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `address` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `person_id` + "`" + `, ` + "`" + `street` + "`" + `, ` + "`" + `city` + "`" + `) VALUES<br>
(1, 1, &#39;1 Love Drive&#39;, &#39;Phoenix&#39;),<br>
(2, 2, &#39;2 Doves and a Pine Cone Dr.&#39;, &#39;Dallas&#39;),<br>
(3, 3, &#39;3 Gold Fish Pl.&#39;, &#39;New York&#39;),<br>
(4, 3, &#39;323 W QCubed&#39;, &#39;New York&#39;),<br>
(5, 5, &#39;22 Elm St&#39;, &#39;Palo Alto&#39;),<br>
(6, 7, &#39;1 Pine St&#39;, &#39;San Jose&#39;),<br>
(7, 7, &#39;421 Central Expw&#39;, &#39;Mountain View&#39;);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `login` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `login` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `person_id` + "`" + ` int(10) UNSIGNED DEFAULT NULL,<br>
  ` + "`" + `username` + "`" + ` varchar(20) NOT NULL,<br>
  ` + "`" + `password` + "`" + ` varchar(20) DEFAULT NULL,<br>
  ` + "`" + `is_enabled` + "`" + ` tinyint(1) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `login` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `login` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `person_id` + "`" + `, ` + "`" + `username` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `is_enabled` + "`" + `) VALUES<br>
(1, 1, &#39;jdoe&#39;, &#39;p@$$.w0rd&#39;, 0),<br>
(2, 3, &#39;brobinson&#39;, &#39;p@$$.w0rd&#39;, 1),<br>
(3, 4, &#39;mho&#39;, &#39;p@$$.w0rd&#39;, 1),<br>
(4, 7, &#39;kwolfe&#39;, &#39;p@$$.w0rd&#39;, 0),<br>
(5, NULL, &#39;system&#39;, &#39;p@$$.w0rd&#39;, 1);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `milestone` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `milestone` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `project_id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `name` + "`" + ` varchar(50) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `milestone` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `milestone` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `project_id` + "`" + `, ` + "`" + `name` + "`" + `) VALUES<br>
(1, 1, &#39;Milestone A&#39;),<br>
(2, 1, &#39;Milestone B&#39;),<br>
(3, 1, &#39;Milestone C&#39;),<br>
(4, 2, &#39;Milestone D&#39;),<br>
(5, 2, &#39;Milestone E&#39;),<br>
(6, 3, &#39;Milestone F&#39;),<br>
(7, 4, &#39;Milestone G&#39;),<br>
(8, 4, &#39;Milestone H&#39;),<br>
(9, 4, &#39;Milestone I&#39;),<br>
(10, 4, &#39;Milestone J&#39;);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `person` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `person` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `first_name` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `last_name` + "`" + ` varchar(50) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `person` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + `) VALUES<br>
(1, &#39;John&#39;, &#39;Doe&#39;),<br>
(2, &#39;Kendall&#39;, &#39;Public&#39;),<br>
(3, &#39;Ben&#39;, &#39;Robinson&#39;),<br>
(4, &#39;Mike&#39;, &#39;Ho&#39;),<br>
(5, &#39;Alex&#39;, &#39;Smith&#39;),<br>
(6, &#39;Wendy&#39;, &#39;Smith&#39;),<br>
(7, &#39;Karen&#39;, &#39;Wolfe&#39;),<br>
(8, &#39;Samantha&#39;, &#39;Jones&#39;),<br>
(9, &#39;Linda&#39;, &#39;Brady&#39;),<br>
(10, &#39;Jennifer&#39;, &#39;Smith&#39;),<br>
(11, &#39;Brett&#39;, &#39;Carlisle&#39;),<br>
(12, &#39;Jacob&#39;, &#39;Pratt&#39;);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `person_persontype_assn` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `person_persontype_assn` + "`" + ` (<br>
  ` + "`" + `person_id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `person_type_id` + "`" + ` int(10) UNSIGNED NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `person_persontype_assn` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `person_persontype_assn` + "`" + ` (` + "`" + `person_id` + "`" + `, ` + "`" + `person_type_id` + "`" + `) VALUES<br>
(3, 1),<br>
(10, 1),<br>
(1, 2),<br>
(3, 2),<br>
(7, 2),<br>
(1, 3),<br>
(3, 3),<br>
(9, 3),<br>
(2, 4),<br>
(7, 4),<br>
(2, 5),<br>
(5, 5);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `person_type` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `person_type` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `name` + "`" + ` varchar(50) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `person_type` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `person_type` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `) VALUES<br>
(4, &#39;Company Car&#39;),<br>
(1, &#39;Contractor&#39;),<br>
(3, &#39;Inactive&#39;),<br>
(2, &#39;Manager&#39;),<br>
(5, &#39;Works From Home&#39;);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `person_with_lock` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `person_with_lock` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `first_name` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `last_name` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `sys_timestamp` + "`" + ` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `person_with_lock` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `person_with_lock` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + `, ` + "`" + `sys_timestamp` + "`" + `) VALUES<br>
(1, &#39;John&#39;, &#39;Doe&#39;, NULL),<br>
(2, &#39;Kendall&#39;, &#39;Public&#39;, NULL),<br>
(3, &#39;Ben&#39;, &#39;Robinson&#39;, NULL),<br>
(4, &#39;Mike&#39;, &#39;Ho&#39;, NULL),<br>
(5, &#39;Alfred&#39;, &#39;Newman&#39;, NULL),<br>
(6, &#39;Wendy&#39;, &#39;Johnson&#39;, NULL),<br>
(7, &#39;Karen&#39;, &#39;Wolfe&#39;, NULL),<br>
(8, &#39;Samantha&#39;, &#39;Jones&#39;, NULL),<br>
(9, &#39;Linda&#39;, &#39;Brady&#39;, NULL),<br>
(10, &#39;Jennifer&#39;, &#39;Smith&#39;, NULL),<br>
(11, &#39;Brett&#39;, &#39;Carlisle&#39;, NULL),<br>
(12, &#39;Jacob&#39;, &#39;Pratt&#39;, NULL);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `project` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `project` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `num` + "`" + ` int(11) NOT NULL COMMENT &#39;To simplify checking test results and as a non pk id test&#39;,<br>
  ` + "`" + `project_status_type_id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `manager_id` + "`" + ` int(10) UNSIGNED DEFAULT NULL,<br>
  ` + "`" + `name` + "`" + ` varchar(100) NOT NULL,<br>
  ` + "`" + `description` + "`" + ` text,<br>
  ` + "`" + `start_date` + "`" + ` date DEFAULT NULL,<br>
  ` + "`" + `end_date` + "`" + ` date DEFAULT NULL,<br>
  ` + "`" + `budget` + "`" + ` decimal(12,2) DEFAULT NULL,<br>
  ` + "`" + `spent` + "`" + ` decimal(12,2) DEFAULT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `project` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `num` + "`" + `, ` + "`" + `project_status_type_id` + "`" + `, ` + "`" + `manager_id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `start_date` + "`" + `, ` + "`" + `end_date` + "`" + `, ` + "`" + `budget` + "`" + `, ` + "`" + `spent` + "`" + `) VALUES<br>
(1, 1, 3, 7, &#39;ACME Website Redesign&#39;, &#39;The redesign of the main website for ACME Incorporated&#39;, &#39;2004-03-01&#39;, &#39;2004-07-01&#39;, &#39;9560.25&#39;, &#39;10250.75&#39;),<br>
(2, 2, 1, 4, &#39;State College HR System&#39;, &#39;Implementation of a back-office Human Resources system for State College&#39;, &#39;2006-02-15&#39;, NULL, &#39;80500.00&#39;, &#39;73200.00&#39;),<br>
(3, 3, 1, 1, &#39;Blueman Industrial Site Architecture&#39;, &#39;Main website architecture for the Blueman Industrial Group&#39;, &#39;2006-03-01&#39;, &#39;2006-04-15&#39;, &#39;2500.00&#39;, &#39;4200.50&#39;),<br>
(4, 4, 2, 7, &#39;ACME Payment System&#39;, &#39;Accounts Payable payment system for ACME Incorporated&#39;, &#39;2005-08-15&#39;, &#39;2005-10-20&#39;, &#39;5124.67&#39;, &#39;5175.30&#39;);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `project_status_type` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `project_status_type` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `name` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `description` + "`" + ` text,<br>
  ` + "`" + `guidelines` + "`" + ` text,<br>
  ` + "`" + `is_active` + "`" + ` tinyint(1) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `project_status_type` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `project_status_type` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `guidelines` + "`" + `, ` + "`" + `is_active` + "`" + `) VALUES<br>
(1, &#39;Open&#39;, &#39;The project is currently active&#39;, &#39;All projects that we are working on should be in this state&#39;, 1),<br>
(2, &#39;Cancelled&#39;, &#39;The project has been canned&#39;, NULL, 1),<br>
(3, &#39;Completed&#39;, &#39;The project has been completed successfully&#39;, &#39;Celebrate successes!&#39;, 1);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `qc_watchers` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `qc_watchers` + "`" + ` (<br>
  ` + "`" + `table_key` + "`" + ` varchar(200) NOT NULL,<br>
  ` + "`" + `ts` + "`" + ` varchar(40) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT=&#39;{&#34;NoCodegen&#34;:true}&#39;;</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `related_project_assn` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `related_project_assn` + "`" + ` (<br>
  ` + "`" + `parent_id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `child_id` + "`" + ` int(10) UNSIGNED NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `related_project_assn` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `related_project_assn` + "`" + ` (` + "`" + `parent_id` + "`" + `, ` + "`" + `child_id` + "`" + `) VALUES<br>
(4, 1),<br>
(1, 3),<br>
(1, 4);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `team_member_project_assn` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `team_member_project_assn` + "`" + ` (<br>
  ` + "`" + `team_member_id` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `project_id` + "`" + ` int(10) UNSIGNED NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `team_member_project_assn` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `team_member_project_assn` + "`" + ` (` + "`" + `team_member_id` + "`" + `, ` + "`" + `project_id` + "`" + `) VALUES<br>
(2, 1),<br>
(5, 1),<br>
(6, 1),<br>
(7, 1),<br>
(8, 1),<br>
(2, 2),<br>
(4, 2),<br>
(5, 2),<br>
(7, 2),<br>
(9, 2),<br>
(10, 2),<br>
(1, 3),<br>
(4, 3),<br>
(6, 3),<br>
(8, 3),<br>
(10, 3),<br>
(1, 4),<br>
(2, 4),<br>
(3, 4),<br>
(5, 4),<br>
(8, 4),<br>
(11, 4),<br>
(12, 4);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `tmp` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `tmp` + "`" + ` (<br>
  ` + "`" + `d` + "`" + ` decimal(5,2) NOT NULL,<br>
  ` + "`" + `i` + "`" + ` smallint(5) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `tmp` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `tmp` + "`" + ` (` + "`" + `d` + "`" + `, ` + "`" + `i` + "`" + `) VALUES<br>
(&#39;-123.46&#39;, 12345);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `two_key` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `two_key` + "`" + ` (<br>
  ` + "`" + `server` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `directory` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `file_name` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `person_id` + "`" + ` int(11) UNSIGNED NOT NULL,<br>
  ` + "`" + `project_id` + "`" + ` int(11) UNSIGNED DEFAULT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `two_key` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `two_key` + "`" + ` (` + "`" + `server` + "`" + `, ` + "`" + `directory` + "`" + `, ` + "`" + `file_name` + "`" + `, ` + "`" + `person_id` + "`" + `, ` + "`" + `project_id` + "`" + `) VALUES<br>
(&#39;cnn.com&#39;, &#39;us&#39;, &#39;news&#39;, 1, 1),<br>
(&#39;google.com&#39;, &#39;drive&#39;, &#39;&#39;, 2, 2),<br>
(&#39;google.com&#39;, &#39;mail&#39;, &#39;mail.html&#39;, 3, 2),<br>
(&#39;google.com&#39;, &#39;news&#39;, &#39;news.php&#39;, 4, 3),<br>
(&#39;mail.google.com&#39;, &#39;mail&#39;, &#39;inbox&#39;, 5, NULL),<br>
(&#39;yahoo.com&#39;, &#39;&#39;, &#39;&#39;, 6, NULL);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `type_test` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `type_test` + "`" + ` (<br>
  ` + "`" + `id` + "`" + ` int(11) NOT NULL,<br>
  ` + "`" + `date` + "`" + ` date DEFAULT NULL,<br>
  ` + "`" + `time` + "`" + ` time DEFAULT NULL,<br>
  ` + "`" + `date_time` + "`" + ` datetime DEFAULT NULL,<br>
  ` + "`" + `ts` + "`" + ` timestamp NULL DEFAULT CURRENT_TIMESTAMP,<br>
  ` + "`" + `test_int` + "`" + ` int(11) DEFAULT &#39;5&#39;,<br>
  ` + "`" + `test_float` + "`" + ` float DEFAULT NULL,<br>
  ` + "`" + `test_text` + "`" + ` text,<br>
  ` + "`" + `test_bit` + "`" + ` tinyint(1) DEFAULT NULL,<br>
  ` + "`" + `test_varchar` + "`" + ` varchar(10) DEFAULT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Dumping data for table ` + "`" + `type_test` + "`" + `<br>
--</p><br>
<p>INSERT INTO ` + "`" + `type_test` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `date` + "`" + `, ` + "`" + `time` + "`" + `, ` + "`" + `date_time` + "`" + `, ` + "`" + `ts` + "`" + `, ` + "`" + `test_int` + "`" + `, ` + "`" + `test_float` + "`" + `, ` + "`" + `test_text` + "`" + `, ` + "`" + `test_bit` + "`" + `, ` + "`" + `test_varchar` + "`" + `) VALUES<br>
(1, &#39;2019-01-02&#39;, &#39;06:17:28&#39;, &#39;2019-01-02 06:17:28&#39;, &#39;2019-01-23 08:52:06&#39;, 5, 1.2, &#39;Sample&#39;, 1, &#39;Sample&#39;);</p><br>
<p>-- --------------------------------------------------------</p><br>
<p>--<br>
-- Table structure for table ` + "`" + `unsupported_types` + "`" + `<br>
--</p><br>
<p>CREATE TABLE ` + "`" + `unsupported_types` + "`" + ` (<br>
  ` + "`" + `type_set` + "`" + ` set(&#39;a&#39;,&#39;b&#39;,&#39;c&#39;) NOT NULL,<br>
  ` + "`" + `type_enum` + "`" + ` enum(&#39;a&#39;,&#39;b&#39;,&#39;c&#39;) NOT NULL,<br>
  ` + "`" + `type_decimal` + "`" + ` decimal(10,4) NOT NULL,<br>
  ` + "`" + `type_double` + "`" + ` double NOT NULL,<br>
  ` + "`" + `type_geo` + "`" + ` geometry NOT NULL,<br>
  ` + "`" + `type_tiny_blob` + "`" + ` tinyblob NOT NULL,<br>
  ` + "`" + `type_medium_blob` + "`" + ` mediumblob NOT NULL,<br>
  ` + "`" + `type_varbinary` + "`" + ` varbinary(10) NOT NULL,<br>
  ` + "`" + `type_longtext` + "`" + ` longtext NOT NULL,<br>
  ` + "`" + `type_binary` + "`" + ` binary(10) NOT NULL,<br>
  ` + "`" + `type_small` + "`" + ` smallint(6) NOT NULL,<br>
  ` + "`" + `type_medium` + "`" + ` mediumint(9) NOT NULL,<br>
  ` + "`" + `type_big` + "`" + ` bigint(20) NOT NULL,<br>
  ` + "`" + `type_polygon` + "`" + ` polygon NOT NULL,<br>
  ` + "`" + `type_serial` + "`" + ` bigint(20) UNSIGNED NOT NULL,<br>
  ` + "`" + `type_unsigned` + "`" + ` int(10) UNSIGNED NOT NULL,<br>
  ` + "`" + `type_multFk1` + "`" + ` varchar(50) NOT NULL,<br>
  ` + "`" + `type_multiFk2` + "`" + ` varchar(50) NOT NULL<br>
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p><br>
<p>--<br>
-- Indexes for dumped tables<br>
--</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `address` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `address` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD KEY ` + "`" + `IDX_address_1` + "`" + ` (` + "`" + `person_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `login` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `login` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD UNIQUE KEY ` + "`" + `IDX_login_2` + "`" + ` (` + "`" + `username` + "`" + `),<br>
  ADD UNIQUE KEY ` + "`" + `IDX_login_1` + "`" + ` (` + "`" + `person_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `milestone` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `milestone` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD KEY ` + "`" + `IDX_milestoneproj_1` + "`" + ` (` + "`" + `project_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `person` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD KEY ` + "`" + `IDX_person_1` + "`" + ` (` + "`" + `last_name` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `person_persontype_assn` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person_persontype_assn` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `person_id` + "`" + `,` + "`" + `person_type_id` + "`" + `),<br>
  ADD KEY ` + "`" + `person_type_id` + "`" + ` (` + "`" + `person_type_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `person_type` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person_type` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD UNIQUE KEY ` + "`" + `name` + "`" + ` (` + "`" + `name` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `person_with_lock` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person_with_lock` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `project` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `project` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD UNIQUE KEY ` + "`" + `num` + "`" + ` (` + "`" + `num` + "`" + `),<br>
  ADD KEY ` + "`" + `IDX_project_1` + "`" + ` (` + "`" + `project_status_type_id` + "`" + `),<br>
  ADD KEY ` + "`" + `IDX_project_2` + "`" + ` (` + "`" + `manager_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `project_status_type` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `project_status_type` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `),<br>
  ADD UNIQUE KEY ` + "`" + `IDX_projectstatustype_1` + "`" + ` (` + "`" + `name` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `qc_watchers` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `qc_watchers` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `table_key` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `related_project_assn` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `related_project_assn` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `parent_id` + "`" + `,` + "`" + `child_id` + "`" + `),<br>
  ADD KEY ` + "`" + `IDX_relatedprojectassn_2` + "`" + ` (` + "`" + `child_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `team_member_project_assn` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `team_member_project_assn` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `team_member_id` + "`" + `,` + "`" + `project_id` + "`" + `) USING BTREE,<br>
  ADD KEY ` + "`" + `IDX_teammemberprojectassn_2` + "`" + ` (` + "`" + `project_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `tmp` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `tmp` + "`" + `<br>
  ADD UNIQUE KEY ` + "`" + `d` + "`" + ` (` + "`" + `d` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `two_key` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `two_key` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `server` + "`" + `,` + "`" + `directory` + "`" + `),<br>
  ADD KEY ` + "`" + `person_id` + "`" + ` (` + "`" + `person_id` + "`" + `),<br>
  ADD KEY ` + "`" + `project_id` + "`" + ` (` + "`" + `project_id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `type_test` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `type_test` + "`" + `<br>
  ADD PRIMARY KEY (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Indexes for table ` + "`" + `unsupported_types` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `unsupported_types` + "`" + `<br>
  ADD UNIQUE KEY ` + "`" + `type_serial` + "`" + ` (` + "`" + `type_serial` + "`" + `),<br>
  ADD KEY ` + "`" + `type_multFk1` + "`" + ` (` + "`" + `type_multFk1` + "`" + `,` + "`" + `type_multiFk2` + "`" + `);</p><br>
<p>--<br>
-- AUTO_INCREMENT for dumped tables<br>
--</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `address` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `address` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `login` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `login` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `milestone` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `milestone` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `person` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=52;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `person_type` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person_type` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `person_with_lock` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person_with_lock` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `project` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `project` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `project_status_type` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `project_status_type` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `type_test` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `type_test` + "`" + `<br>
  MODIFY ` + "`" + `id` + "`" + ` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;</p><br>
<p>--<br>
-- AUTO_INCREMENT for table ` + "`" + `unsupported_types` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `unsupported_types` + "`" + `<br>
  MODIFY ` + "`" + `type_serial` + "`" + ` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;</p><br>
<p>--<br>
-- Constraints for dumped tables<br>
--</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `address` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `address` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `person_address` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `login` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `login` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `person_login` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `milestone` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `milestone` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `project_milestone` + "`" + ` FOREIGN KEY (` + "`" + `project_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `person_persontype_assn` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `person_persontype_assn` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `person_persontype_assn_1` + "`" + ` FOREIGN KEY (` + "`" + `person_type_id` + "`" + `) REFERENCES ` + "`" + `person_type` + "`" + ` (` + "`" + `id` + "`" + `),<br>
  ADD CONSTRAINT ` + "`" + `person_persontype_assn_2` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `project` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `project` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `person_project` + "`" + ` FOREIGN KEY (` + "`" + `manager_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `),<br>
  ADD CONSTRAINT ` + "`" + `project_status_type_project` + "`" + ` FOREIGN KEY (` + "`" + `project_status_type_id` + "`" + `) REFERENCES ` + "`" + `project_status_type` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `related_project_assn` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `related_project_assn` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `related_project_assn_1` + "`" + ` FOREIGN KEY (` + "`" + `parent_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `),<br>
  ADD CONSTRAINT ` + "`" + `related_project_assn_2` + "`" + ` FOREIGN KEY (` + "`" + `child_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `team_member_project_assn` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `team_member_project_assn` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `person_team_member_project_assn` + "`" + ` FOREIGN KEY (` + "`" + `team_member_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `),<br>
  ADD CONSTRAINT ` + "`" + `project_team_member_project_assn` + "`" + ` FOREIGN KEY (` + "`" + `project_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `);</p><br>
<p>--<br>
-- Constraints for table ` + "`" + `two_key` + "`" + `<br>
--<br>
ALTER TABLE ` + "`" + `two_key` + "`" + `<br>
  ADD CONSTRAINT ` + "`" + `two_key_person` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `),<br>
  ADD CONSTRAINT ` + "`" + `two_key_project` + "`" + ` FOREIGN KEY (` + "`" + `project_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `);<br>
SET FOREIGN_KEY_CHECKS=1;<br>
COMMIT;<br>
</p>
`)

			buf.WriteString(`
</code>
`)

			buf.WriteString(`
</body>
</html>
`)

			return

		})
}
