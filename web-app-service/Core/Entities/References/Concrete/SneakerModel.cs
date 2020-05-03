﻿using System;
using System.ComponentModel.DataAnnotations;
using System.Text.RegularExpressions;
using Core.Reference;

namespace Core.Entities.References
{
	public class SneakerModel : IModel
	{
		[Key]
		public string UniqueID { get; }

		public string Name { get; set; }

		public SneakerBrand Brand { get; set; }

		public SneakerModel BaseModel { get; set; }

		public string Technology { get; set; }

		public string Description { get; set; }

		public string HeroPath { get; set; }

		public static implicit operator SneakerModel(string field) => new SneakerModel(field);

		public static implicit operator string(SneakerModel property) => property.Name;

		public SneakerModel(string name)
		{
			Name = name;
			UniqueID = new Regex("[\\n\\t-;,.\\s()\\/]").Replace(Convert.ToString(name), "_").ToLower();
		}

		public SneakerModel(string name, SneakerBrand brand) : this(name)
		{
			Brand = brand;
			var predictedBase = name.Split(new[] {' ', '_', '-'}, StringSplitOptions.RemoveEmptyEntries)[0];
			if (predictedBase != name) BaseModel = new SneakerModel(predictedBase, brand);
		}

		public bool Equals(SneakerModel other) => other != null && UniqueID == other.UniqueID;
	}
}