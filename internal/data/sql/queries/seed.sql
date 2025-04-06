insert into types (id, name) values
('bd122736-5450-4aa5-9afa-26b6ad0b765d', 'Manga'),
('0db1cfd5-ab77-45a0-bd4f-56e2a668fe03', 'Light Novel');

insert into publishers (id, name) values
('00000000-0000-0000-0000-000000000000', 'None'),
('01960c5a-c064-7083-ad2f-ebc70397ae1d', 'Big Comic Spirits'),
('01960c5b-c2be-7042-9129-05cd830272ea', 'Shounen Sunday'),
('01960c5c-446d-7c6a-87c5-c6dbfc34bb93', 'Magazine pocket'),
('01960c5c-446d-7fde-b0a6-6407ad3e14d6', 'Afternoon');

insert into authors (id, full_name) VALUES
('01960c5c-446d-7629-aefc-36fdf23053f1', 'Asato Asato'),
('01960c5c-446d-763f-afce-6a095ba45fdb', 'Kazuki Miya'),
('01960c5d-bf63-78e2-abdb-814d5ea19038', 'Asano Inio'),
('01960c5e-1a89-74c0-b421-4cd7fac46647', 'Yukimura Makoto'),
('01960c5e-d7f8-7a1b-93cd-176140084c7a', 'Mikami Saka'),
('01960c60-09f0-71d0-b0b0-7d8e72a4c543', 'Yamada Kaneshito');

insert into books (id, name, score, publication_date, type_id, author_id, publisher_id) values
('01960c60-09f0-7de6-9179-0d1407f929a8', 'Dead Dead Demons Dededede Destruction', 8.20, '2014-04-28', 'bd122736-5450-4aa5-9afa-26b6ad0b765d', '01960c5d-bf63-78e2-abdb-814d5ea19038', '01960c5a-c064-7083-ad2f-ebc70397ae1d'),
('01960c60-09f0-7b06-8fc5-e4fb3798b01d', 'Oyasumi Punpun', 8.99, '2007-03-15', 'bd122736-5450-4aa5-9afa-26b6ad0b765d', '01960c5d-bf63-78e2-abdb-814d5ea19038', '01960c5a-c064-7083-ad2f-ebc70397ae1d'),
('01960c60-09f0-7b9a-86c2-70263a499f15', 'Sousou no Frieren', 8.84, '2020-04-28', 'bd122736-5450-4aa5-9afa-26b6ad0b765d', '01960c60-09f0-71d0-b0b0-7d8e72a4c543', '01960c5b-c2be-7042-9129-05cd830272ea'),
('01960c60-09f0-782e-b04a-c4faed77a643', 'Kaoru Hana wa Rin to Saku', 8.82, '2021-10-21', 'bd122736-5450-4aa5-9afa-26b6ad0b765d', '01960c5e-d7f8-7a1b-93cd-176140084c7a', '01960c5c-446d-7c6a-87c5-c6dbfc34bb93'),
('01960c60-09f0-74ed-8a8e-373003e5911f', 'Vinland Saga', 9.08, '2005-04-13', 'bd122736-5450-4aa5-9afa-26b6ad0b765d', '01960c5e-1a89-74c0-b421-4cd7fac46647', '01960c5c-446d-7fde-b0a6-6407ad3e14d6'),
('01960c60-09f0-7bbc-a5c3-3590cff866a0', '86', 8.83, '2017-02-10', '0db1cfd5-ab77-45a0-bd4f-56e2a668fe03', '01960c5c-446d-7629-aefc-36fdf23053f1', '00000000-0000-0000-0000-000000000000'),
('01960c7e-abbc-7802-9634-faa863fb5d3c', 'Ascendance of a Bookworm: I''ll Do Anything to Become a Librarian!', 8.78, '2015-01-25', '0db1cfd5-ab77-45a0-bd4f-56e2a668fe03', '01960c5c-446d-763f-afce-6a095ba45fdb', '00000000-0000-0000-0000-000000000000');
