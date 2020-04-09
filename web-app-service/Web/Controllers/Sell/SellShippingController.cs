﻿using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		public ActionResult Shipping(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(3, model);
			return this.ViewStep(5, model);
		}
	}
}