package utilities

var MapRoleToRoleType = map[int32]string{
	int32(NoRole):       string(NoRole),
	int32(Barber):       string(Staff),
	int32(Receptionist): string(Staff),
	int32(Manager):      string(Staff),
	int32(Admin):        string(Administrator),
	int32(SuperAdmin):   string(Administrator),
	int32(OtherStaff):   string(Staff),
}
