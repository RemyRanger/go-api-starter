package models

const SelectProgramIndex = `
SELECT 
	p.id AS id,
	v_enabled.program_link_id AS link_id,
	d.id AS developer_id,
	d.name AS developer_name,
	p.title,
	p.description,
	p.commercialisation_date AS agenda_commercialisation_date,
	(CASE 
		WHEN p_linked.id IS NOT NULL AND p.delivery_date IS NULL THEN
			p_linked.delivery_date
		WHEN p.delivery_date != '0000-00-00 00:00:00' THEN
			p.delivery_date
		ELSE
			NULL
	END) AS agenda_delivery_date,
	p.offer_start_date AS agenda_start_launch,
	p.integration_date AS created_at,
	p.update_at AS updated_at,
	p.flow_origin AS flow_origin,
	p.ubiflow_id AS ubiflow_id,
	p.flow_id AS flow_id,
	p.nb_properties AS nb_properties,
	p.nb_available_properties AS nb_available_properties,
	p.nb_rooms_min AS nb_rooms_min,
	p.nb_rooms_max AS nb_rooms_max,
	p.area_min AS area_min,
	p.area_max AS area_max,
	p.price_min AS price_min,
	p.price_max AS price_max,
	(CASE 
		WHEN p_linked.id IS NOT NULL THEN
			GROUP_CONCAT( DISTINCT(p_linked.certifications) SEPARATOR ',')
		ELSE
			p.certifications
	END) AS label,
	(CASE 
		WHEN p_linked.id IS NOT NULL THEN
			GROUP_CONCAT( DISTINCT(p_linked.tax_exemption) SEPARATOR ',')
		ELSE
			p.tax_exemption
	END) AS tax_exemption,
	GROUP_CONCAT( DISTINCT(pr.type_parent) SEPARATOR ',') AS type_parent,
	p.city AS address_city,
	p.country AS address_country,
	p.street AS address_street,
	p.zip AS address_zip,
	(CASE 
		WHEN p_linked.latitude IS NOT NULL THEN
			p_linked.latitude
		ELSE
			p.latitude
	END) AS latitude,
	(CASE 
		WHEN p_linked.longitude IS NOT NULL THEN
			p_linked.longitude
		ELSE
			p.longitude
	END) AS longitude,
	(CASE 
		WHEN GROUP_CONCAT( DISTINCT(CONCAT('"',m.path, '"')) SEPARATOR ',') IS NULL THEN
			'[]'
		ELSE
			CONCAT('[', GROUP_CONCAT( DISTINCT(CONCAT('"',m.path, '"')) SEPARATOR ','), ']')
	END) AS image_paths,
	(CASE 
		WHEN p_linked.latitude IS NOT NULL AND p_linked.longitude IS NOT NULL THEN 
			CONCAT(p_linked.latitude ,',', p_linked.longitude)
		ELSE
			CONCAT(p.latitude ,',', p.longitude)
	END) AS location,
	(CASE 
		WHEN p.nb_rooms IS NOT NULL THEN 
			p.nb_rooms
		ELSE
			GROUP_CONCAT( DISTINCT(pr.nb_rooms) SEPARATOR ',') 
	END) AS nb_room,
	(CASE 
		WHEN p_linked.id IS NOT NULL AND p.delivery_date IS NULL THEN
			DATE_FORMAT(p_linked.delivery_date, '%Y-%m-%d')
		WHEN p.delivery_date != '0000-00-00 00:00:00' AND p.delivery_date IS NOT NULL THEN
			DATE_FORMAT(p.delivery_date, '%Y-%m-%d')
		ELSE
			NULL
	END) AS delivery,
	(CASE 
		WHEN dv.developer_id IS NOT NULL THEN 
			GROUP_CONCAT( DISTINCT(v_linked.subscription_id) SEPARATOR ' ')
	END) AS linked_for_subscriptions_ids,
	(CASE 
		WHEN dv.developer_id IS NOT NULL THEN 
			GROUP_CONCAT( DISTINCT(v_enabled.partner_id) SEPARATOR ' ')
	END) AS visible_for_keys_ids,
	(CASE 
		WHEN dv.developer_id IS NOT NULL THEN 
			GROUP_CONCAT( DISTINCT(v_shared.subscription_id) SEPARATOR ' ')
	END) AS visible_for_subscriptions_ids,
	(CASE 
		WHEN dv.developer_id IS NOT NULL THEN 
			GROUP_CONCAT( DISTINCT(v_suspended.subscription_id) SEPARATOR ' ')
	END) AS suspended_for_subscriptions_ids
FROM cdv_program AS p
LEFT JOIN cdv_program_visibility v_linked
	ON (p.id = v_linked.program_link_id)
LEFT JOIN cdv_program p_linked
	ON (p_linked.id = v_linked.program_id)
LEFT JOIN cdv_program_visibility v_enabled
	ON (p.id = v_enabled.program_id AND v_enabled.status = 1)
LEFT JOIN cdv_program_visibility v_shared
	ON (p.id = v_shared.program_id AND v_shared.status = 2)
LEFT JOIN cdv_program_visibility v_suspended
	ON (p.id = v_suspended.program_id AND v_suspended.status = 3)
LEFT JOIN cdv_media m
	ON p.id = m.program_id
LEFT JOIN cdv_developer d
	ON p.developer_id = d.id
LEFT JOIN cdv_developer_visibility dv
	ON (d.id = dv.developer_id AND (dv.status = 1 OR dv.status = 2))
LEFT JOIN cdv_property pr
	ON (p.id = pr.program_id AND pr.enabled)
WHERE p.id = ? AND p.enabled AND (pr.id IS NOT NULL OR p.ubiflow_id IS NULL)
GROUP BY p.id
`
