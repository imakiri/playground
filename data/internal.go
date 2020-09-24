package data

import "sync"

func (i *internal) GetSalt() string {
	return i.salt
}

func (i *internal) GetUser(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = i.main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE login = ?", login)
	if re.Check(err, c) {
		return
	}

	c <- re
	close(c)
}

func (i *internal) GetUserById(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = i.main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE id = ?", id)
	if re.Check(err, c) {
		return
	}

	c <- re
	close(c)
}

func (i *internal) GetUserPassHash(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = i.main.Get(&re.User, "SELECT passHash FROM main.users WHERE login = ?", login)
	if re.Check(err, c) {
		return
	}

	c <- re
	close(c)
}

func (i *internal) GetData(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}

func (i *internal) AddPic(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}

func (i *internal) CreateUser(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}

func (i *internal) UpdateUser(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}

func (i *internal) DeleteUser(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}
