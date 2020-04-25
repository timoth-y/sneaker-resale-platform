﻿using Core.Attributes;

namespace Core.Entities.Users
{
	[EntityService(Resource = "api/users")]
	public class User : IBaseEntity
	{
		public string UniqueID { get; set; }

		public string UserName { get; set; }

		public string PasswordHash { get; set; }

		public string FirstName { get; set; }

		public string LastName { get; set; }

		public bool Equals(User other) => other != null && UniqueID == other.UniqueID;
	}
}