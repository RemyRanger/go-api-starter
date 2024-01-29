package models

const SelectProgramPropertiesIndex = `
SELECT 
	p.id,
	p.ubiflow_id,
	p.developer_id,
	d.name AS developer_name,
	p.program_id,
	pr.title AS program_name,
	p.description,
	p.reference,
	p.flow_id,
	pr.ubiflow_id AS program_ubiflow_id,
	p.flow_origin,
	p.area,
	p.lot_availability AS availability,
	DATE_FORMAT(p.integration_date, '%Y-%m-%d') AS created_at,
	p.update_at AS updated_at,
	p.title,
	p.nb_rooms,
	p.price,
	p.rent,
	p.service,
	(CASE 
		WHEN p_linked.id IS NOT NULL THEN
			GROUP_CONCAT( DISTINCT(p_linked.certifications) SEPARATOR ',')
		ELSE
			NULL
	END) AS label,
	(CASE 
		WHEN p_linked.id IS NOT NULL THEN
			GROUP_CONCAT( DISTINCT(p_linked.tax_exemption) SEPARATOR ',')
		ELSE
			pr.tax_exemption
	END) AS tax_exemption,
	p.type_parent,
	p.owner_type,
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
		WHEN p_linked.latitude IS NOT NULL AND p_linked.longitude IS NOT NULL THEN 
			CONCAT(p_linked.latitude ,',', p_linked.longitude)
		WHEN p.latitude IS NOT NULL OR p.longitude IS NOT NULL THEN 
			CONCAT(p.latitude ,',', p.longitude)
		ELSE
			CONCAT(0,',',0) 
	END) AS location,
	(CASE 
		WHEN p.program_id IS NOT NULL AND GROUP_CONCAT( DISTINCT(CONCAT('"',mp.path, '"')) SEPARATOR ',') IS NOT NULL THEN
			CONCAT('[', GROUP_CONCAT( DISTINCT(CONCAT('"',mp.path, '"')) SEPARATOR ','), ']')
		WHEN p.program_id IS NULL AND GROUP_CONCAT( DISTINCT(CONCAT('"',m.path, '"')) SEPARATOR ',') IS NOT NULL THEN
			CONCAT('[', GROUP_CONCAT( DISTINCT(CONCAT('"',m.path, '"')) SEPARATOR ','), ']')
		ELSE
			'[]'
	END) AS image_paths,
	(CASE
		WHEN p_linked.id IS NOT NULL AND pr.delivery_date IS NULL THEN
			DATE_FORMAT(p_linked.delivery_date, '%Y-%m-%d')
		WHEN pr.delivery_date != '0000-00-00 00:00:00' AND pr.delivery_date IS NOT NULL THEN
			DATE_FORMAT(pr.delivery_date, '%Y-%m-%d')
		ELSE
			NULL
	END) AS delivery,
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
FROM cdv_property AS p
LEFT JOIN cdv_program pr
	ON p.program_id = pr.id
LEFT JOIN cdv_developer d
	ON p.developer_id = d.id
LEFT JOIN cdv_media m
	ON p.id = m.property_id
LEFT JOIN cdv_media mp
	ON pr.id = mp.program_id
LEFT JOIN cdv_agency a
	ON p.agency_id = a.id
LEFT JOIN cdv_program_visibility v_linked
	ON (p.program_id = v_linked.program_link_id)
LEFT JOIN cdv_program p_linked
	ON (p_linked.id = v_linked.program_id)
LEFT JOIN cdv_program_visibility v_enabled
	ON (p.program_id = v_enabled.program_id AND v_enabled.status = 1)
LEFT JOIN cdv_program_visibility v_shared
	ON (p.program_id = v_shared.program_id AND v_shared.status = 2)
LEFT JOIN cdv_program_visibility v_suspended
	ON (p.program_id = v_suspended.program_id AND v_suspended.status = 3)
LEFT JOIN cdv_developer_visibility dv
	ON (d.id = dv.developer_id AND (dv.status = 1 OR dv.status = 2))
WHERE p.program_id = ? AND p.enabled
GROUP BY p.id
`
