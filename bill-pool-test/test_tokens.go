package bill_pool

var(
	USDT_CONTRACT = "5fc56b05322e302e306a00527ac4681953797374656d2e53746f726167652e476574436f6e746578746a51527ac404555344546a52527ac404555344546a53527ac4566a54527ac40340420f6a55527ac4224158646d647a62796633575a4b517a5274724e5177415239315a784d556668586b7468204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a56527ac4070080c6a47e8d036a57527ac401016a58527ac401026a59527ac40b546f74616c537570706c796a5a527ac46c0125c56b6a00527ac46a51527ac46a52527ac46a51c304696e69747d9c7c75641500006a00c3068403000000006e6c75666203006a51c3046e616d657d9c7c75641500006a00c306b204000000006e6c75666203006a51c30673796d626f6c7d9c7c75641500006a00c306ca04000000006e6c75666203006a51c308646563696d616c737d9c7c75641500006a00c306e204000000006e6c75666203006a51c30b746f74616c537570706c797d9c7c75641500006a00c306fc04000000006e6c75666203006a51c30962616c616e63654f667d9c7c756435006a52c3c0517d9e7c75640a00006c75666203006a52c300c36a54527ac46a54c3516a00c3062f05000000006e6c75666203006a51c3087472616e736665727d9c7c75644f006a52c3c0537d9e7c75640a00006c75666239006a52c300c36a55527ac46a52c351c36a56527ac46a52c352c36a57527ac46a57c36a56c36a55c3536a00c3069105000000006e6c75666203006a51c30d7472616e736665724d756c74697d9c7c756418006a52c3516a00c3063707000000006e6c75666203006a51c30c7472616e7366657246726f6d7d9c7c75645c006a52c3c0547d9e7c75640a00006c75666203006a52c300c36a58527ac46a52c351c36a55527ac46a52c352c36a56527ac46a52c353c36a57527ac46a57c36a56c36a55c36a58c3546a00c3060309000000006e6c75666203006a51c307617070726f76657d9c7c75644f006a52c3c0537d9e7c75640a00006c75666203006a52c300c36a59527ac46a52c351c36a58527ac46a52c352c36a57527ac46a57c36a58c36a59c3536a00c306f307000000006e6c75666203006a51c309616c6c6f77616e63657d9c7c756442006a52c3c0527d9e7c75640a00006c75666203006a52c300c36a59527ac46a52c351c36a58527ac46a58c36a59c3526a00c3064b0b000000006e6c7566620300006c75665ac56b6a00527ac46a51527ac46203006a00c356c3c001147d9e7c756434000e4f776e657220696c6c6567616c2151c176c9681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a00c35ac36a00c351c3681253797374656d2e53746f726167652e47657464360014416c726561647920696e697469616c697a656421681553797374656d2e52756e74696d652e4e6f74696679006c7566628a006a00c357c36a00c355c3956a52527ac46a52c36a00c35ac36a00c351c3681253797374656d2e53746f726167652e5075746a52c36a00c358c36a00c356c37e6a00c351c3681253797374656d2e53746f726167652e5075746a52c36a00c356c300087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75666c756656c56b6a00527ac46a51527ac46203006a00c352c36c756656c56b6a00527ac46a51527ac46203006a00c353c36c756656c56b6a00527ac46a51527ac46203006a00c354c300936c756656c56b6a00527ac46a51527ac46203006a00c35ac36a00c351c3681253797374656d2e53746f726167652e47657400936c756658c56b6a00527ac46a51527ac46a52527ac46203006a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a00c358c36a52c37e6a00c351c3681253797374656d2e53746f726167652e47657400936c75665fc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c3c001147d9e7c7576630e00756a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c7576630c00756a54c3007d9f7c75640a00006c75666203006a00c358c36a52c37e6a55527ac46a55c36a00c351c3681253797374656d2e53746f726167652e4765746a56527ac46a54c36a56c37da07c75640a00006c75666203006a54c36a56c37d9c7c756425006a55c36a00c351c3681553797374656d2e53746f726167652e44656c6574656226006a56c36a54c3946a55c36a00c351c3681253797374656d2e53746f726167652e5075746a00c358c36a53c37e6a57527ac46a57c36a00c351c3681253797374656d2e53746f726167652e4765746a58527ac46a58c36a54c3936a57c36a00c351c3681253797374656d2e53746f726167652e5075746a54c36a53c36a52c3087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac4620300006a53527ac46a52c36a54527ac46a54c3c06a55527ac46a53c36a55c39f6485006a54c36a53c3c36a56527ac46a53c351936a53527ac46a56c3c0537d9e7c756423001b7472616e736665724d756c746920706172616d73206572726f722ef06203006a56c352c36a56c351c36a56c300c3536a00c3069105000000006e007d9c7c75641d00157472616e736665724d756c7469206661696c65642ef06203006277ff516c75665dc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c3c001147d9e7c7576630e00756a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c75640a00006c75666203006a54c36a52c3516a00c3062f05000000006e7da07c7576630c00756a54c3007d9f7c75640a00006c75666203006a00c359c36a52c37e6a53c37e6a56527ac46a54c36a56c36a00c351c3681253797374656d2e53746f726167652e5075746a54c36a53c36a52c308617070726f76616c54c1681553797374656d2e52756e74696d652e4e6f74696679516c75660114c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46203006a52c3c001147d9e7c7576631d00756a53c3c001147d9e7c7576630e00756a54c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c75640a00006c75666203006a00c358c36a53c37e6a56527ac46a56c36a00c351c3681253797374656d2e53746f726167652e4765746a57527ac46a55c36a57c37da07c7576630c00756a55c3007d9f7c75640a00006c75666203006a00c359c36a53c37e6a52c37e6a58527ac46a58c36a00c351c3681253797374656d2e53746f726167652e4765746a59527ac46a00c358c36a54c37e6a5a527ac46a55c36a59c37da07c75640a00006c7566629b006a55c36a59c37d9c7c756448006a58c36a00c351c3681553797374656d2e53746f726167652e44656c6574656a57c36a55c3946a56c36a00c351c3681253797374656d2e53746f726167652e5075746249006a59c36a55c3946a58c36a00c351c3681253797374656d2e53746f726167652e5075746a57c36a55c3946a56c36a00c351c3681253797374656d2e53746f726167652e5075746a5ac36a00c351c3681253797374656d2e53746f726167652e4765746a5b527ac46a5bc36a55c3936a5ac36a00c351c3681253797374656d2e53746f726167652e5075746a55c36a54c36a53c3087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75665ac56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a00c359c36a52c37e6a53c37e6a54527ac46a54c36a00c351c3681253797374656d2e53746f726167652e47657400936c7566"
	BTC_CONTRACT = "5fc56b05322e302e306a00527ac4681953797374656d2e53746f726167652e476574436f6e746578746a51527ac4034254436a52527ac4034254436a53527ac4596a54527ac40400ca9a3b6a55527ac4224158646d647a62796633575a4b517a5274724e5177415239315a784d556668586b7468204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a56527ac40400ca9a3b6a57527ac401016a58527ac401026a59527ac40b546f74616c537570706c796a5a527ac46c0125c56b6a00527ac46a51527ac46a52527ac46a51c304696e69747d9c7c75641500006a00c3068003000000006e6c75666203006a51c3046e616d657d9c7c75641500006a00c306ae04000000006e6c75666203006a51c30673796d626f6c7d9c7c75641500006a00c306c604000000006e6c75666203006a51c308646563696d616c737d9c7c75641500006a00c306de04000000006e6c75666203006a51c30b746f74616c537570706c797d9c7c75641500006a00c306f804000000006e6c75666203006a51c30962616c616e63654f667d9c7c756435006a52c3c0517d9e7c75640a00006c75666203006a52c300c36a54527ac46a54c3516a00c3062b05000000006e6c75666203006a51c3087472616e736665727d9c7c75644f006a52c3c0537d9e7c75640a00006c75666239006a52c300c36a55527ac46a52c351c36a56527ac46a52c352c36a57527ac46a57c36a56c36a55c3536a00c3068d05000000006e6c75666203006a51c30d7472616e736665724d756c74697d9c7c756418006a52c3516a00c3063307000000006e6c75666203006a51c30c7472616e7366657246726f6d7d9c7c75645c006a52c3c0547d9e7c75640a00006c75666203006a52c300c36a58527ac46a52c351c36a55527ac46a52c352c36a56527ac46a52c353c36a57527ac46a57c36a56c36a55c36a58c3546a00c306ff08000000006e6c75666203006a51c307617070726f76657d9c7c75644f006a52c3c0537d9e7c75640a00006c75666203006a52c300c36a59527ac46a52c351c36a58527ac46a52c352c36a57527ac46a57c36a58c36a59c3536a00c306ef07000000006e6c75666203006a51c309616c6c6f77616e63657d9c7c756442006a52c3c0527d9e7c75640a00006c75666203006a52c300c36a59527ac46a52c351c36a58527ac46a58c36a59c3526a00c306470b000000006e6c7566620300006c75665ac56b6a00527ac46a51527ac46203006a00c356c3c001147d9e7c756434000e4f776e657220696c6c6567616c2151c176c9681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a00c35ac36a00c351c3681253797374656d2e53746f726167652e47657464360014416c726561647920696e697469616c697a656421681553797374656d2e52756e74696d652e4e6f74696679006c7566628a006a00c357c36a00c355c3956a52527ac46a52c36a00c35ac36a00c351c3681253797374656d2e53746f726167652e5075746a52c36a00c358c36a00c356c37e6a00c351c3681253797374656d2e53746f726167652e5075746a52c36a00c356c300087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75666c756656c56b6a00527ac46a51527ac46203006a00c352c36c756656c56b6a00527ac46a51527ac46203006a00c353c36c756656c56b6a00527ac46a51527ac46203006a00c354c300936c756656c56b6a00527ac46a51527ac46203006a00c35ac36a00c351c3681253797374656d2e53746f726167652e47657400936c756658c56b6a00527ac46a51527ac46a52527ac46203006a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a00c358c36a52c37e6a00c351c3681253797374656d2e53746f726167652e47657400936c75665fc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c3c001147d9e7c7576630e00756a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c7576630c00756a54c3007d9f7c75640a00006c75666203006a00c358c36a52c37e6a55527ac46a55c36a00c351c3681253797374656d2e53746f726167652e4765746a56527ac46a54c36a56c37da07c75640a00006c75666203006a54c36a56c37d9c7c756425006a55c36a00c351c3681553797374656d2e53746f726167652e44656c6574656226006a56c36a54c3946a55c36a00c351c3681253797374656d2e53746f726167652e5075746a00c358c36a53c37e6a57527ac46a57c36a00c351c3681253797374656d2e53746f726167652e4765746a58527ac46a58c36a54c3936a57c36a00c351c3681253797374656d2e53746f726167652e5075746a54c36a53c36a52c3087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac4620300006a53527ac46a52c36a54527ac46a54c3c06a55527ac46a53c36a55c39f6485006a54c36a53c3c36a56527ac46a53c351936a53527ac46a56c3c0537d9e7c756423001b7472616e736665724d756c746920706172616d73206572726f722ef06203006a56c352c36a56c351c36a56c300c3536a00c3068d05000000006e007d9c7c75641d00157472616e736665724d756c7469206661696c65642ef06203006277ff516c75665dc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c3c001147d9e7c7576630e00756a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c75640a00006c75666203006a54c36a52c3516a00c3062b05000000006e7da07c7576630c00756a54c3007d9f7c75640a00006c75666203006a00c359c36a52c37e6a53c37e6a56527ac46a54c36a56c36a00c351c3681253797374656d2e53746f726167652e5075746a54c36a53c36a52c308617070726f76616c54c1681553797374656d2e52756e74696d652e4e6f74696679516c75660114c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46203006a52c3c001147d9e7c7576631d00756a53c3c001147d9e7c7576630e00756a54c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c75640a00006c75666203006a00c358c36a53c37e6a56527ac46a56c36a00c351c3681253797374656d2e53746f726167652e4765746a57527ac46a55c36a57c37da07c7576630c00756a55c3007d9f7c75640a00006c75666203006a00c359c36a53c37e6a52c37e6a58527ac46a58c36a00c351c3681253797374656d2e53746f726167652e4765746a59527ac46a00c358c36a54c37e6a5a527ac46a55c36a59c37da07c75640a00006c7566629b006a55c36a59c37d9c7c756448006a58c36a00c351c3681553797374656d2e53746f726167652e44656c6574656a57c36a55c3946a56c36a00c351c3681253797374656d2e53746f726167652e5075746249006a59c36a55c3946a58c36a00c351c3681253797374656d2e53746f726167652e5075746a57c36a55c3946a56c36a00c351c3681253797374656d2e53746f726167652e5075746a5ac36a00c351c3681253797374656d2e53746f726167652e4765746a5b527ac46a5bc36a55c3936a5ac36a00c351c3681253797374656d2e53746f726167652e5075746a55c36a54c36a53c3087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75665ac56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a00c359c36a52c37e6a53c37e6a54527ac46a54c36a00c351c3681253797374656d2e53746f726167652e47657400936c7566"
	ETH_CONTRACT = "5fc56b05322e302e306a00527ac4681953797374656d2e53746f726167652e476574436f6e746578746a51527ac4034554486a52527ac4034554486a53527ac401126a54527ac408000064a7b3b6e00d6a55527ac4224158646d647a62796633575a4b517a5274724e5177415239315a784d556668586b7468204f6e746f6c6f67792e52756e74696d652e426173653538546f416464726573736a56527ac40400ca9a3b6a57527ac401016a58527ac401026a59527ac40b546f74616c537570706c796a5a527ac46c0125c56b6a00527ac46a51527ac46a52527ac46a51c304696e69747d9c7c75641500006a00c3068503000000006e6c75666203006a51c3046e616d657d9c7c75641500006a00c306b304000000006e6c75666203006a51c30673796d626f6c7d9c7c75641500006a00c306cb04000000006e6c75666203006a51c308646563696d616c737d9c7c75641500006a00c306e304000000006e6c75666203006a51c30b746f74616c537570706c797d9c7c75641500006a00c306fd04000000006e6c75666203006a51c30962616c616e63654f667d9c7c756435006a52c3c0517d9e7c75640a00006c75666203006a52c300c36a54527ac46a54c3516a00c3063005000000006e6c75666203006a51c3087472616e736665727d9c7c75644f006a52c3c0537d9e7c75640a00006c75666239006a52c300c36a55527ac46a52c351c36a56527ac46a52c352c36a57527ac46a57c36a56c36a55c3536a00c3069205000000006e6c75666203006a51c30d7472616e736665724d756c74697d9c7c756418006a52c3516a00c3063807000000006e6c75666203006a51c30c7472616e7366657246726f6d7d9c7c75645c006a52c3c0547d9e7c75640a00006c75666203006a52c300c36a58527ac46a52c351c36a55527ac46a52c352c36a56527ac46a52c353c36a57527ac46a57c36a56c36a55c36a58c3546a00c3060409000000006e6c75666203006a51c307617070726f76657d9c7c75644f006a52c3c0537d9e7c75640a00006c75666203006a52c300c36a59527ac46a52c351c36a58527ac46a52c352c36a57527ac46a57c36a58c36a59c3536a00c306f407000000006e6c75666203006a51c309616c6c6f77616e63657d9c7c756442006a52c3c0527d9e7c75640a00006c75666203006a52c300c36a59527ac46a52c351c36a58527ac46a58c36a59c3526a00c3064c0b000000006e6c7566620300006c75665ac56b6a00527ac46a51527ac46203006a00c356c3c001147d9e7c756434000e4f776e657220696c6c6567616c2151c176c9681553797374656d2e52756e74696d652e4e6f74696679006c75666203006a00c35ac36a00c351c3681253797374656d2e53746f726167652e47657464360014416c726561647920696e697469616c697a656421681553797374656d2e52756e74696d652e4e6f74696679006c7566628a006a00c357c36a00c355c3956a52527ac46a52c36a00c35ac36a00c351c3681253797374656d2e53746f726167652e5075746a52c36a00c358c36a00c356c37e6a00c351c3681253797374656d2e53746f726167652e5075746a52c36a00c356c300087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75666c756656c56b6a00527ac46a51527ac46203006a00c352c36c756656c56b6a00527ac46a51527ac46203006a00c353c36c756656c56b6a00527ac46a51527ac46203006a00c354c300936c756656c56b6a00527ac46a51527ac46203006a00c35ac36a00c351c3681253797374656d2e53746f726167652e47657400936c756658c56b6a00527ac46a51527ac46a52527ac46203006a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a00c358c36a52c37e6a00c351c3681253797374656d2e53746f726167652e47657400936c75665fc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c3c001147d9e7c7576630e00756a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c7576630c00756a54c3007d9f7c75640a00006c75666203006a00c358c36a52c37e6a55527ac46a55c36a00c351c3681253797374656d2e53746f726167652e4765746a56527ac46a54c36a56c37da07c75640a00006c75666203006a54c36a56c37d9c7c756425006a55c36a00c351c3681553797374656d2e53746f726167652e44656c6574656226006a56c36a54c3946a55c36a00c351c3681253797374656d2e53746f726167652e5075746a00c358c36a53c37e6a57527ac46a57c36a00c351c3681253797374656d2e53746f726167652e4765746a58527ac46a58c36a54c3936a57c36a00c351c3681253797374656d2e53746f726167652e5075746a54c36a53c36a52c3087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75665bc56b6a00527ac46a51527ac46a52527ac4620300006a53527ac46a52c36a54527ac46a54c3c06a55527ac46a53c36a55c39f6485006a54c36a53c3c36a56527ac46a53c351936a53527ac46a56c3c0537d9e7c756423001b7472616e736665724d756c746920706172616d73206572726f722ef06203006a56c352c36a56c351c36a56c300c3536a00c3069205000000006e007d9c7c75641d00157472616e736665724d756c7469206661696c65642ef06203006277ff516c75665dc56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46203006a53c3c001147d9e7c7576630e00756a52c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c75640a00006c75666203006a54c36a52c3516a00c3063005000000006e7da07c7576630c00756a54c3007d9f7c75640a00006c75666203006a00c359c36a52c37e6a53c37e6a56527ac46a54c36a56c36a00c351c3681253797374656d2e53746f726167652e5075746a54c36a53c36a52c308617070726f76616c54c1681553797374656d2e52756e74696d652e4e6f74696679516c75660114c56b6a00527ac46a51527ac46a52527ac46a53527ac46a54527ac46a55527ac46203006a52c3c001147d9e7c7576631d00756a53c3c001147d9e7c7576630e00756a54c3c001147d9e7c75641c001461646472657373206c656e677468206572726f72f06203006a52c3681b53797374656d2e52756e74696d652e436865636b5769746e657373007d9c7c75640a00006c75666203006a00c358c36a53c37e6a56527ac46a56c36a00c351c3681253797374656d2e53746f726167652e4765746a57527ac46a55c36a57c37da07c7576630c00756a55c3007d9f7c75640a00006c75666203006a00c359c36a53c37e6a52c37e6a58527ac46a58c36a00c351c3681253797374656d2e53746f726167652e4765746a59527ac46a00c358c36a54c37e6a5a527ac46a55c36a59c37da07c75640a00006c7566629b006a55c36a59c37d9c7c756448006a58c36a00c351c3681553797374656d2e53746f726167652e44656c6574656a57c36a55c3946a56c36a00c351c3681253797374656d2e53746f726167652e5075746249006a59c36a55c3946a58c36a00c351c3681253797374656d2e53746f726167652e5075746a57c36a55c3946a56c36a00c351c3681253797374656d2e53746f726167652e5075746a5ac36a00c351c3681253797374656d2e53746f726167652e4765746a5b527ac46a5bc36a55c3936a5ac36a00c351c3681253797374656d2e53746f726167652e5075746a55c36a54c36a53c3087472616e7366657254c1681553797374656d2e52756e74696d652e4e6f74696679516c75665ac56b6a00527ac46a51527ac46a52527ac46a53527ac46203006a00c359c36a52c37e6a53c37e6a54527ac46a54c36a00c351c3681253797374656d2e53746f726167652e47657400936c7566"
)