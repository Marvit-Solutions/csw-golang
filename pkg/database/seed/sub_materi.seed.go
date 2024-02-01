package seed

// import (
// 	ds "csw-golang/internal/domain/entity/datastruct"
// )

// func CreateSubMateri() []*ds.SubMateri {
// 	subMateri := []*ds.SubMateri{
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e121",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Pengertian Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e122",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Asal Usul Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Sejarah Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Lambang dan Nilai Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Dimensi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e126",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Rumusan Kesatuan Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e127",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Implementasi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e128",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Pancasila Dalam Ketatanegaraan NKRI",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e129",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
// 			Name:     "Kedudukan dan Fungsi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},

// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e130",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e125",
// 			Name:     "Pengertian Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e131",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e125",
// 			Name:     "Asal Usul Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e132",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e125",
// 			Name:     "Sejarah Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e133",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e125",
// 			Name:     "Lambang dan Nilai Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e134",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e125",
// 			Name:     "Dimensi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e135",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e126",
// 			Name:     "Rumusan Kesatuan Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e136",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e126",
// 			Name:     "Implementasi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e137",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e126",
// 			Name:     "Pancasila Dalam Ketatanegaraan NKRI",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e138",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e126",
// 			Name:     "Kedudukan dan Fungsi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e139",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e127",
// 			Name:     "Implementasi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e140",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e127",
// 			Name:     "Pancasila Dalam Ketatanegaraan NKRI",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e141",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e127",
// 			Name:     "Kedudukan dan Fungsi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e142",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e128",
// 			Name:     "Kedudukan dan Fungsi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e143",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e128",
// 			Name:     "Implementasi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e144",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e128",
// 			Name:     "Pancasila Dalam Ketatanegaraan NKRI",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 		{
// 			ID:       "2daa6bd9-f39f-4ca8-5e8e-e692f687e145",
// 			MateriID: "2daa6bd9-f39g-4ca8-5e8e-e692f687e128",
// 			Name:     "Kedudukan dan Fungsi Pancasila",
// 			Konten:   "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed elementum tempus egestas sed. Aliquam malesuada bibendum arcu vitae. Id diam maecenas ultricies mi eget mauris pharetra et ultrices. Congue quisque egestas diam in arcu cursus. Lacus luctus accumsan tortor posuere ac ut consequat. Nulla at volutpat diam ut venenatis tellus in. Consequat nisl vel pretium lectus quam id leo in. Maecenas pharetra convallis posuere morbi. Eget arcu dictum varius duis at. Lorem sed risus ultricies tristique. In nibh mauris cursus mattis molestie a. Vitae congue eu consequat ac felis donec et odio. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in. Enim diam vulputate ut pharetra sit amet aliquam id. At lectus urna duis convallis. Est velit egestas dui id ornare arcu odio ut. Tempor id eu nisl nunc. Tristique senectus et netus et malesuada fames. Netus et malesuada fames ac.</p>",
// 		},
// 	}
// 	return subMateri
// }
