DROP TABLE IF EXISTS `LocationInfo`;
CREATE TABLE `LocationInfo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `province` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `district` varchar(255) NOT NULL,
  `code` varchar(13) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '北京市',    '110000000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '市辖区',    '110100000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '直辖县',    '110200000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '东城区',    '110101000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '西城区',    '110102000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '朝阳区',    '110105000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '丰台区',    '110106000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '石景山区',	 '110107000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '海淀区',	 '110108000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '门头沟区',	 '110109000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '房山区',	 '110111000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '通州区',	 '110112000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '顺义区',	 '110113000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '昌平区',	 '110114000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '大兴区',	 '110115000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '怀柔区',	 '110116000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '平谷区',	 '110117000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '密云县',	 '110228000000');
insert into `LocationInfo` (province, city, district, code) values ('北京市', '北京市', '延庆县',	 '110229000000');

DROP TABLE IF EXISTS `Frequency`;
CREATE TABLE `Frequency` (
  `channelID` int(11) NOT NULL,
  `channel` Char(4) NOT NULL,
  `video` float NOT NULL,
  `audio` float NOT NULL,
  `center` float NOT NULL,
  `low` float NOT NULL,
  `high` float NOT NULL,
  PRIMARY KEY (`channelID`),
  UNIQUE KEY `channelID` (`channelID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

insert into Frequency (channelID, channel, video, audio, center, low, high) values (16, 'DS16', 495.25, 501.75, 498, 494, 502);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (17, 'DS17', 503.25, 509.75, 506, 502, 510);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (18, 'DS18', 511.25, 517.75, 514, 510, 518);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (19, 'DS19', 519.25, 225.75, 522, 518, 526);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (20, 'DS20', 527.25, 533.75, 530, 526, 534);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (21, 'DS21', 535.25, 541.75, 538, 534, 542);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (22, 'DS22', 543.25, 549.75, 546, 542, 550);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (23, 'DS23', 551.25, 557.75, 554, 550, 558);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (24, 'DS24', 559.25, 565.75, 562, 558, 566);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (25, 'DS25', 607.25, 613.75, 610, 606, 614);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (26, 'DS26', 615.25, 621.75, 618, 614, 622);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (27, 'DS27', 623.25, 629.75, 626, 622, 630);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (28, 'DS28', 631.25, 637.75, 634, 630, 638);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (29, 'DS29', 639.25, 645.75, 642, 638, 646);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (30, 'DS30', 647.25, 653.75, 650, 646, 654);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (31, 'DS31', 655.25, 661.75, 658, 654, 662);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (32, 'DS32', 663.25, 669.75, 666, 662, 670);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (33, 'DS33', 671.25, 677.75, 674, 670, 678);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (34, 'DS34', 679.25, 685.75, 682, 678, 686);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (35, 'DS35', 687.25, 693.75, 690, 686, 694);
insert into Frequency (channelID, channel, video, audio, center, low, high) values (36, 'DS36', 695.25, 701.75, 698, 694, 702);


DROP TABLE IF EXISTS `CMMB`;
CREATE TABLE `CMMB` (
  `id` int NOT NULL AUTO_INCREMENT,
  `districtcode` varchar(12) NOT NULL,
  `channel` int NOT NULL,
  `power` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

insert into CMMB (districtcode, channel, power) values ('110105000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110105000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110108000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110105000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110106000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110102000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110107000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110108000000', 20, 1);
insert into CMMB (districtcode, channel, power) values ('110105000000', 20, 1);

DROP TABLE IF EXISTS `DTMB`;
CREATE TABLE `DTMB` (
  `id` int NOT NULL AUTO_INCREMENT,
  `districtcode` varchar(12) NOT NULL,
  `channel` int NOT NULL,
  `power` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
insert into DTMB (districtcode, channel, power) values ('110108000000', 32, 1);
insert into DTMB (districtcode, channel, power) values ('110108000000', 33, 3);
insert into DTMB (districtcode, channel, power) values ('110105000000', 32, 1);
insert into DTMB (districtcode, channel, power) values ('110114000000', 32, 1);
insert into DTMB (districtcode, channel, power) values ('110105000000', 32, 1);
insert into DTMB (districtcode, channel, power) values ('110108000000', 14, 1);
insert into DTMB (districtcode, channel, power) values ('110108000000', 22, 1);
insert into DTMB (districtcode, channel, power) values ('110105000000', 22, 1);
insert into DTMB (districtcode, channel, power) values ('110105000000', 22, 1);
insert into DTMB (districtcode, channel, power) values ('110105000000', 22, 1);

DROP TABLE IF EXISTS `TV`;
CREATE TABLE `TV` (
  `id` int NOT NULL AUTO_INCREMENT,
  `districtcode` varchar(12) NOT NULL,
  `channel` int NOT NULL,
  `power` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
insert into TV (districtcode, channel, power) values ('110108000000', 8, 30);
insert into TV (districtcode, channel, power) values ('110108000000', 15, 30);
insert into TV (districtcode, channel, power) values ('110108000000', 44, 3);
insert into TV (districtcode, channel, power) values ('110111000000', 23, 0.3);
insert into TV (districtcode, channel, power) values ('110111000000', 17, 0.3);
insert into TV (districtcode, channel, power) values ('110116000000', 13, 1);
insert into TV (districtcode, channel, power) values ('110116000000', 47, 1);
insert into TV (districtcode, channel, power) values ('110117000000', 47, 1);
insert into TV (districtcode, channel, power) values ('110228000000', 38, 0.3);
insert into TV (districtcode, channel, power) values ('110228000000', 23, 1);
insert into TV (districtcode, channel, power) values ('110228000000', 46, 1);
insert into TV (districtcode, channel, power) values ('110228000000', 34, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 10, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 41, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 24, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 25, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 30, 1);
insert into TV (districtcode, channel, power) values ('110116000000', 12, 1);
insert into TV (districtcode, channel, power) values ('110112000000', 39, 1);
insert into TV (districtcode, channel, power) values ('110116000000', 29, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 47, 1);
insert into TV (districtcode, channel, power) values ('110229000000', 4, 0.3);
insert into TV (districtcode, channel, power) values ('110228000000', 31, 0.3);
insert into TV (districtcode, channel, power) values ('110111000000', 40, 3);
insert into TV (districtcode, channel, power) values ('110113000000', 18, 3);
insert into TV (districtcode, channel, power) values ('110114000000', 13, 3);
insert into TV (districtcode, channel, power) values ('110117000000', 4, 1);
insert into TV (districtcode, channel, power) values ('110228000000', 4, 1);
insert into TV (districtcode, channel, power) values ('110115000000', 47, 3);
insert into TV (districtcode, channel, power) values ('110116000000', 12, 0.3);
insert into TV (districtcode, channel, power) values ('110108000000', 2, 30);
insert into TV (districtcode, channel, power) values ('110108000000', 6, 10);
insert into TV (districtcode, channel, power) values ('110108000000', 21, 30);
insert into TV (districtcode, channel, power) values ('110108000000', 27, 30);

DROP TABLE IF EXISTS `FreqUsing`;
CREATE TABLE `FreqUsing` (
  `id` int NOT NULL AUTO_INCREMENT,
  `districtcode` varchar(12) NOT NULL,
  `latitude` float NOT NULL,
  `longtitude` float NOT NULL,
  `channel` int NOT NULL,
  `power` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

