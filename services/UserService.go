type UserService struct {}

fun(u *UserService)Create(user *User);

}

func (us *UserService)Update(u *User){}

func (us *UserService)Remove(u *User){}

func (us *UserService)RemoveById(id uint){}

func (us *UserService)FindByEmail(email string){}

func (us *UserService)GetAll(){}
