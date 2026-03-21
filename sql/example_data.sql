-- Comprehensive example data for testing the parts management system
-- This creates a rich dataset with multiple manufacturers, categories, parts, and objects

-- ============================================================================
-- Sample Manufacturers
-- ============================================================================
INSERT INTO manufacturers (name, country_of_origin) VALUES
    ('Arduino', 'ITA'),
    ('Raspberry Pi Foundation', 'GBR'),
    ('Texas Instruments', 'USA'),
    ('Murata Manufacturing', 'JPN'),
    ('SparkFun Electronics', 'USA'),
    ('Adafruit Industries', 'USA'),
    ('STMicroelectronics', 'CHE'),
    ('Bosch', 'DEU'),
    ('NXP Semiconductors', 'NLD'),
    ('Analog Devices', 'USA'),
    ('Seeed Studio', 'CHN'),
    ('DFRobot', 'CHN'),
    ('Pololu', 'USA'),
    ('3M', 'USA'),
    ('Samsung', 'KOR')
ON CONFLICT DO NOTHING;

-- ============================================================================
-- Sample Categories (Flat Structure)
-- ============================================================================
INSERT INTO categories (name, description) VALUES
    ('Electronics', 'Electronic components and devices'),
    ('Mechanical', 'Mechanical parts and hardware'),
    ('Tools', 'Tools and equipment'),
    ('Consumables', 'Consumable materials'),
    ('Microcontrollers', 'Microcontroller boards and modules'),
    ('Sensors', 'Various sensors and transducers'),
    ('Displays', 'LCD, OLED, LED displays'),
    ('Motors', 'DC motors, stepper motors, servos'),
    ('Power Supply', 'Power supplies, batteries, regulators'),
    ('Communication', 'WiFi, Bluetooth, RF modules'),
    ('Passive Components', 'Resistors, capacitors, inductors'),
    ('Fasteners', 'Screws, nuts, bolts'),
    ('Enclosures', 'Cases and housings'),
    ('Structural', 'Frames, brackets, mounts'),
    ('Wires & Cables', 'Wires, cables, connectors'),
    ('Adhesives', 'Tape, glue, epoxy')
ON CONFLICT DO NOTHING;

-- ============================================================================
-- Sample Parts Catalog (50+ parts)
-- ============================================================================
DO $$
DECLARE
    arduino_mfr_id UUID;
    rpi_mfr_id UUID;
    ti_mfr_id UUID;
    sparkfun_mfr_id UUID;
    adafruit_mfr_id UUID;
    st_mfr_id UUID;
    bosch_mfr_id UUID;
    samsung_mfr_id UUID;
    pololu_mfr_id UUID;
    seeed_mfr_id UUID;
    threeM_mfr_id UUID;

    microcontroller_cat_id UUID;
    sensor_cat_id UUID;
    display_cat_id UUID;
    motor_cat_id UUID;
    power_cat_id UUID;
    comm_cat_id UUID;
    wire_cat_id UUID;
BEGIN
    -- Get manufacturer IDs
    SELECT id INTO arduino_mfr_id FROM manufacturers WHERE name = 'Arduino';
    SELECT id INTO rpi_mfr_id FROM manufacturers WHERE name = 'Raspberry Pi Foundation';
    SELECT id INTO ti_mfr_id FROM manufacturers WHERE name = 'Texas Instruments';
    SELECT id INTO sparkfun_mfr_id FROM manufacturers WHERE name = 'SparkFun Electronics';
    SELECT id INTO adafruit_mfr_id FROM manufacturers WHERE name = 'Adafruit Industries';
    SELECT id INTO st_mfr_id FROM manufacturers WHERE name = 'STMicroelectronics';
    SELECT id INTO bosch_mfr_id FROM manufacturers WHERE name = 'Bosch';
    SELECT id INTO samsung_mfr_id FROM manufacturers WHERE name = 'Samsung';
    SELECT id INTO pololu_mfr_id FROM manufacturers WHERE name = 'Pololu';
    SELECT id INTO seeed_mfr_id FROM manufacturers WHERE name = 'Seeed Studio';
    SELECT id INTO threeM_mfr_id FROM manufacturers WHERE name = '3M';

    -- Get category IDs
    SELECT id INTO microcontroller_cat_id FROM categories WHERE name = 'Microcontrollers';
    SELECT id INTO sensor_cat_id FROM categories WHERE name = 'Sensors';
    SELECT id INTO display_cat_id FROM categories WHERE name = 'Displays';
    SELECT id INTO motor_cat_id FROM categories WHERE name = 'Motors';
    SELECT id INTO power_cat_id FROM categories WHERE name = 'Power Supply';
    SELECT id INTO comm_cat_id FROM categories WHERE name = 'Communication';
    SELECT id INTO wire_cat_id FROM categories WHERE name = 'Wires & Cables';

    -- Microcontrollers
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('Arduino Uno R3', 'A000066', arduino_mfr_id, 'Microcontroller board based on ATmega328P',
        '{"mcu": "ATmega328P", "voltage": "5V", "digital_io": 14, "analog_in": 6, "flash": "32KB", "sram": "2KB"}'::jsonb),
    ('Arduino Mega 2560', 'A000067', arduino_mfr_id, 'Microcontroller board with 54 digital I/O pins',
        '{"mcu": "ATmega2560", "voltage": "5V", "digital_io": 54, "analog_in": 16, "flash": "256KB", "sram": "8KB"}'::jsonb),
    ('Arduino Nano', 'A000005', arduino_mfr_id, 'Compact Arduino board',
        '{"mcu": "ATmega328P", "voltage": "5V", "digital_io": 14, "analog_in": 8, "flash": "32KB"}'::jsonb),
    ('Raspberry Pi 4 Model B 8GB', 'RPI4-MODBP-8GB', rpi_mfr_id, 'Single-board computer with 8GB RAM',
        '{"cpu": "Cortex-A72 1.5GHz", "ram": "8GB", "usb": 4, "ethernet": "1Gbps", "wifi": "802.11ac"}'::jsonb),
    ('Raspberry Pi Zero W', 'RPI-ZERO-W', rpi_mfr_id, 'Ultra-compact Pi with WiFi',
        '{"cpu": "1GHz single-core", "ram": "512MB", "wifi": "802.11n", "bluetooth": "4.1"}'::jsonb),
    ('ESP32 DevKit', 'ESP32-DEVKIT-V1', sparkfun_mfr_id, 'WiFi + Bluetooth development board',
        '{"mcu": "ESP32", "wifi": "802.11bgn", "bluetooth": "4.2", "flash": "4MB"}'::jsonb),
    ('ESP8266 NodeMCU', 'ESP8266-12E', sparkfun_mfr_id, 'WiFi development board',
        '{"mcu": "ESP8266", "wifi": "802.11bgn", "flash": "4MB", "gpio": 17}'::jsonb);

    -- Sensors
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('TMP36 Temperature Sensor', 'TMP36GT9Z', adafruit_mfr_id, 'Analog temperature sensor',
        '{"range": "-40°C to +125°C", "accuracy": "±2°C", "output": "10mV/°C"}'::jsonb),
    ('DHT22 Temperature & Humidity', 'DHT22', adafruit_mfr_id, 'Digital temperature and humidity sensor',
        '{"temp_range": "-40°C to +80°C", "humidity_range": "0-100%", "accuracy": "±0.5°C"}'::jsonb),
    ('BMP280 Pressure Sensor', 'BMP280', bosch_mfr_id, 'Barometric pressure sensor',
        '{"range": "300-1100hPa", "accuracy": "±1hPa", "interface": "I2C/SPI"}'::jsonb),
    ('MPU6050 IMU', 'MPU-6050', adafruit_mfr_id, '6-axis accelerometer and gyroscope',
        '{"gyro": "±250 to ±2000°/s", "accel": "±2g to ±16g", "interface": "I2C"}'::jsonb),
    ('HC-SR04 Ultrasonic', 'HC-SR04', sparkfun_mfr_id, 'Ultrasonic distance sensor',
        '{"range": "2cm to 400cm", "accuracy": "3mm", "voltage": "5V"}'::jsonb),
    ('PIR Motion Sensor', 'HC-SR501', sparkfun_mfr_id, 'Passive infrared motion detector',
        '{"range": "7m", "angle": "120°", "voltage": "5-20V"}'::jsonb),
    ('Light Sensor LDR', 'GL5528', sparkfun_mfr_id, 'Light dependent resistor',
        '{"dark_resistance": "1MΩ", "light_resistance": "10-20kΩ"}'::jsonb);

    -- Displays
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('LCD 16x2 Display', 'LCD1602', sparkfun_mfr_id, '16x2 character LCD display',
        '{"size": "16x2", "interface": "parallel", "backlight": "blue", "voltage": "5V"}'::jsonb),
    ('OLED 0.96" 128x64', 'SSD1306', adafruit_mfr_id, 'Small OLED display',
        '{"resolution": "128x64", "size": "0.96 inch", "interface": "I2C/SPI", "color": "white"}'::jsonb),
    ('TFT 2.8" Touch Display', 'ILI9341', adafruit_mfr_id, 'Color touch screen',
        '{"resolution": "240x320", "size": "2.8 inch", "touch": "resistive", "interface": "SPI"}'::jsonb),
    ('7-Segment Display', '7SEG-4DIGIT', sparkfun_mfr_id, '4-digit 7-segment display',
        '{"digits": 4, "height": "0.56 inch", "color": "red", "interface": "I2C"}'::jsonb);

    -- Motors
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('DC Motor 6V', 'DC-6V-130', pololu_mfr_id, 'Small DC motor',
        '{"voltage": "6V", "rpm": "11000", "current": "70mA"}'::jsonb),
    ('Servo Motor SG90', 'SG90', pololu_mfr_id, 'Micro servo motor',
        '{"voltage": "4.8-6V", "torque": "1.8kg-cm", "angle": "180°"}'::jsonb),
    ('Stepper Motor NEMA17', 'NEMA17-42', pololu_mfr_id, 'Stepper motor for 3D printers',
        '{"voltage": "12V", "steps": 200, "torque": "40Ncm", "current": "1.5A"}'::jsonb),
    ('Stepper Motor 28BYJ-48', '28BYJ-48', pololu_mfr_id, 'Small stepper motor with driver',
        '{"voltage": "5V", "steps": 2048, "reduction": "1/64"}'::jsonb);

    -- Power Supply
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('LM7805 Voltage Regulator', 'LM7805', ti_mfr_id, '5V voltage regulator',
        '{"output": "5V", "current": "1.5A", "input": "7-35V"}'::jsonb),
    ('AMS1117 3.3V Regulator', 'AMS1117-3.3', adafruit_mfr_id, '3.3V LDO regulator',
        '{"output": "3.3V", "current": "1A", "dropout": "1.3V"}'::jsonb),
    ('18650 Battery 3400mAh', 'NCR18650B', samsung_mfr_id, 'Lithium-ion battery',
        '{"capacity": "3400mAh", "voltage": "3.7V", "chemistry": "Li-ion"}'::jsonb),
    ('USB Power Bank 10000mAh', 'PB-10000', samsung_mfr_id, 'Portable USB power bank',
        '{"capacity": "10000mAh", "output": "5V/2A", "ports": 2}'::jsonb);

    -- Communication Modules
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('Bluetooth HC-05', 'HC-05', sparkfun_mfr_id, 'Bluetooth serial module',
        '{"range": "10m", "baud": "9600-115200", "voltage": "3.3-5V"}'::jsonb),
    ('WiFi ESP-01', 'ESP-01', sparkfun_mfr_id, 'Compact WiFi module',
        '{"wifi": "802.11bgn", "gpio": 2, "voltage": "3.3V"}'::jsonb),
    ('NRF24L01 RF Module', 'NRF24L01+', sparkfun_mfr_id, '2.4GHz wireless transceiver',
        '{"frequency": "2.4GHz", "range": "100m", "data_rate": "2Mbps"}'::jsonb),
    ('LoRa Module RFM95W', 'RFM95W', adafruit_mfr_id, 'Long range radio module',
        '{"frequency": "915MHz", "range": "10km", "power": "20dBm"}'::jsonb);

    -- Wires & Connectors
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('Jumper Wires M-M 20cm', 'JW-MM-20', sparkfun_mfr_id, 'Male-to-male jumper wires pack of 40',
        '{"length": "20cm", "count": 40, "type": "male-male"}'::jsonb),
    ('Jumper Wires M-F 20cm', 'JW-MF-20', sparkfun_mfr_id, 'Male-to-female jumper wires pack of 40',
        '{"length": "20cm", "count": 40, "type": "male-female"}'::jsonb),
    ('USB Cable Micro-B 1m', 'USB-MICRO-1M', sparkfun_mfr_id, 'USB 2.0 Micro-B cable',
        '{"length": "1m", "type": "USB 2.0 Micro-B"}'::jsonb),
    ('Breadboard 830 points', 'BB-830', sparkfun_mfr_id, 'Solderless breadboard',
        '{"points": 830, "size": "165x55mm"}'::jsonb),
    ('Dupont Connectors Kit', 'DUPONT-KIT', sparkfun_mfr_id, 'Assorted Dupont connectors',
        '{"pins": "620pcs", "pitch": "2.54mm"}'::jsonb);

    -- Misc
    INSERT INTO parts (name, part_number, manufacturer_id, description, specifications) VALUES
    ('LED Red 5mm', 'LED-RED-5MM', sparkfun_mfr_id, 'Red LED 5mm',
        '{"color": "red", "forward_voltage": "2.0V", "current": "20mA"}'::jsonb),
    ('LED RGB Common Cathode', 'LED-RGB-CC', sparkfun_mfr_id, 'RGB LED common cathode',
        '{"type": "RGB", "pins": 4, "voltage": "2.0-3.2V"}'::jsonb),
    ('Resistor 10kΩ Pack', 'RES-10K-50', sparkfun_mfr_id, 'Carbon film resistors pack of 50',
        '{"value": "10kΩ", "tolerance": "5%", "power": "0.25W", "count": 50}'::jsonb),
    ('Capacitor 100uF', 'CAP-100UF', sparkfun_mfr_id, 'Electrolytic capacitor',
        '{"value": "100µF", "voltage": "25V", "type": "electrolytic"}'::jsonb),
    ('Push Button Switch', 'BTN-6MM', sparkfun_mfr_id, 'Tactile push button',
        '{"type": "tactile", "size": "6x6mm", "force": "160gf"}'::jsonb),
    ('Rotary Encoder', 'ROT-ENC-EC11', sparkfun_mfr_id, 'Rotary encoder with push button',
        '{"pulses": 20, "switch": "yes"}'::jsonb),
    ('Heat Shrink Tubing', 'HST-ASSORTED', threeM_mfr_id, 'Assorted heat shrink tubing',
        '{"sizes": "2-13mm", "count": 328, "ratio": "2:1"}'::jsonb);

    -- Associate parts with categories
    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, microcontroller_cat_id FROM parts p
    WHERE p.part_number IN ('A000066', 'A000067', 'A000005', 'RPI4-MODBP-8GB', 'RPI-ZERO-W', 'ESP32-DEVKIT-V1', 'ESP8266-12E');

    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, sensor_cat_id FROM parts p
    WHERE p.part_number IN ('TMP36GT9Z', 'DHT22', 'BMP280', 'MPU-6050', 'HC-SR04', 'HC-SR501', 'GL5528');

    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, display_cat_id FROM parts p
    WHERE p.part_number IN ('LCD1602', 'SSD1306', 'ILI9341', '7SEG-4DIGIT');

    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, motor_cat_id FROM parts p
    WHERE p.part_number IN ('DC-6V-130', 'SG90', 'NEMA17-42', '28BYJ-48');

    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, power_cat_id FROM parts p
    WHERE p.part_number IN ('LM7805', 'AMS1117-3.3', 'NCR18650B', 'PB-10000');

    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, comm_cat_id FROM parts p
    WHERE p.part_number IN ('HC-05', 'ESP-01', 'NRF24L01+', 'RFM95W');

    INSERT INTO part_categories (part_id, category_id)
    SELECT p.id, wire_cat_id FROM parts p
    WHERE p.part_number IN ('JW-MM-20', 'JW-MF-20', 'USB-MICRO-1M', 'BB-830', 'DUPONT-KIT');
END $$;

-- ============================================================================
-- Parts Inventory (100+ items)
-- ============================================================================
DO $$
DECLARE
    part_record RECORD;
    i INT;
BEGIN
    -- Arduino Uno R3: 15 units
    SELECT id INTO part_record FROM parts WHERE part_number = 'A000066';
    FOR i IN 1..15 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available, notes)
        VALUES (
            part_record.id,
            'ARD-UNO-' || LPAD(i::TEXT, 3, '0'),
            i <= 10,  -- First 10 are available
            CASE WHEN i <= 10 THEN 'Available in stock' ELSE 'In use' END
        );
    END LOOP;

    -- Arduino Nano: 20 units (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'A000005';
    FOR i IN 1..20 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'ARD-NANO-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- Raspberry Pi 4: 8 units
    SELECT id INTO part_record FROM parts WHERE part_number = 'RPI4-MODBP-8GB';
    FOR i IN 1..8 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available, notes)
        VALUES (
            part_record.id,
            'RPI4-' || LPAD(i::TEXT, 3, '0'),
            i <= 5,  -- First 5 are available
            CASE WHEN i <= 5 THEN 'Available in stock' ELSE 'In use - Production' END
        );
    END LOOP;

    -- ESP32: 25 units
    SELECT id INTO part_record FROM parts WHERE part_number = 'ESP32-DEVKIT-V1';
    FOR i IN 1..25 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available, notes)
        VALUES (part_record.id, 'ESP32-' || LPAD(i::TEXT, 3, '0'),
                i <= 20,  -- First 20 are available
                CASE WHEN i > 20 THEN 'In use - IoT Project ' || i ELSE NULL END);
    END LOOP;

    -- Temperature sensors: 30 units (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'TMP36GT9Z';
    FOR i IN 1..30 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'TMP36-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- DHT22: 15 units (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'DHT22';
    FOR i IN 1..15 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'DHT22-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- OLED Displays: 12 units (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'SSD1306';
    FOR i IN 1..12 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'OLED-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- Servo Motors: 20 units (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'SG90';
    FOR i IN 1..20 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'SERVO-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- LEDs: 50 units (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'LED-RED-5MM';
    FOR i IN 1..50 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'LED-R-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- Jumper wires: 10 packs (all available)
    SELECT id INTO part_record FROM parts WHERE part_number = 'JW-MM-20';
    FOR i IN 1..10 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available)
        VALUES (part_record.id, 'JWMM-' || LPAD(i::TEXT, 3, '0'), true);
    END LOOP;

    -- Breadboards: 15 units
    SELECT id INTO part_record FROM parts WHERE part_number = 'BB-830';
    FOR i IN 1..15 LOOP
        INSERT INTO parts_inventory (part_id, serial_number, is_available, notes)
        VALUES (part_record.id, 'BB-' || LPAD(i::TEXT, 3, '0'),
                i <= 12,  -- First 12 are available
                CASE WHEN i > 12 THEN 'In use - Lab Kit ' || i ELSE NULL END);
    END LOOP;
END $$;

-- ============================================================================
-- Objects (10 different projects/products)
-- ============================================================================
DO $$
DECLARE
    obj_id UUID;

    -- Part IDs
    arduino_uno UUID;
    arduino_nano UUID;
    rpi4 UUID;
    esp32 UUID;
    tmp36 UUID;
    dht22 UUID;
    bmp280 UUID;
    oled UUID;
    lcd UUID;
    servo UUID;
    led UUID;
    jumper UUID;
    breadboard UUID;
BEGIN
    -- Get part IDs
    SELECT id INTO arduino_uno FROM parts WHERE part_number = 'A000066';
    SELECT id INTO arduino_nano FROM parts WHERE part_number = 'A000005';
    SELECT id INTO rpi4 FROM parts WHERE part_number = 'RPI4-MODBP-8GB';
    SELECT id INTO esp32 FROM parts WHERE part_number = 'ESP32-DEVKIT-V1';
    SELECT id INTO tmp36 FROM parts WHERE part_number = 'TMP36GT9Z';
    SELECT id INTO dht22 FROM parts WHERE part_number = 'DHT22';
    SELECT id INTO bmp280 FROM parts WHERE part_number = 'BMP280';
    SELECT id INTO oled FROM parts WHERE part_number = 'SSD1306';
    SELECT id INTO lcd FROM parts WHERE part_number = 'LCD1602';
    SELECT id INTO servo FROM parts WHERE part_number = 'SG90';
    SELECT id INTO led FROM parts WHERE part_number = 'LED-RED-5MM';
    SELECT id INTO jumper FROM parts WHERE part_number = 'JW-MM-20';
    SELECT id INTO breadboard FROM parts WHERE part_number = 'BB-830';

    -- Product 1: Arduino Starter Kit
    INSERT INTO products (name, description, version)
    VALUES ('Arduino Starter Kit', 'Complete Arduino starter kit for beginners', 'v1.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, arduino_uno, 1, 'Main board'),
        (obj_id, breadboard, 1, 'For prototyping'),
        (obj_id, jumper, 1, 'Wire pack'),
        (obj_id, led, 10, 'For basic projects'),
        (obj_id, tmp36, 1, 'Temperature sensor example');

    -- Product 2: IoT Weather Station
    INSERT INTO products (name, description, version)
    VALUES ('IoT Weather Station', 'WiFi-enabled weather monitoring station', 'v2.1')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, esp32, 1, 'Main controller with WiFi'),
        (obj_id, dht22, 1, 'Temperature and humidity'),
        (obj_id, bmp280, 1, 'Barometric pressure'),
        (obj_id, oled, 1, 'Display');

    -- Product 3: Home Automation Hub
    INSERT INTO products (name, description, version)
    VALUES ('Home Automation Hub', 'Raspberry Pi based home automation controller', 'v3.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, rpi4, 1, 'Main computer'),
        (obj_id, esp32, 3, 'Sensor nodes'),
        (obj_id, tmp36, 5, 'Room temperature sensors');

    -- Product 4: Robot Car
    INSERT INTO products (name, description, version)
    VALUES ('Robot Car', 'Arduino-based robot car with obstacle avoidance', 'v1.5')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, arduino_uno, 1, 'Main controller'),
        (obj_id, servo, 2, 'For steering and camera pan'),
        (obj_id, led, 4, 'Headlights and indicators');

    -- Product 5: Temperature Logger
    INSERT INTO products (name, description, version)
    VALUES ('Multi-Zone Temperature Logger', 'Data logger for multiple temperature zones', 'v1.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, arduino_nano, 1, 'Main controller'),
        (obj_id, tmp36, 8, 'Eight temperature zones'),
        (obj_id, lcd, 1, 'Display');

    -- Product 6: Smart Garden Monitor
    INSERT INTO products (name, description, version)
    VALUES ('Smart Garden Monitor', 'Monitor soil moisture and environmental conditions', 'v2.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, esp32, 1, 'Main controller'),
        (obj_id, dht22, 2, 'Air temp/humidity'),
        (obj_id, tmp36, 4, 'Soil sensors');

    -- Product 7: LED Matrix Display
    INSERT INTO products (name, description, version)
    VALUES ('LED Matrix Display', 'Scrolling message display', 'v1.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, arduino_uno, 1, 'Controller'),
        (obj_id, led, 64, 'LED matrix 8x8');

    -- Product 8: Education Lab Kit
    INSERT INTO products (name, description, version)
    VALUES ('Electronics Education Lab Kit', 'Complete kit for electronics education', 'v1.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, arduino_uno, 2, 'For experiments'),
        (obj_id, breadboard, 2, 'Prototyping'),
        (obj_id, jumper, 2, 'Wire packs'),
        (obj_id, led, 20, 'Various LEDs'),
        (obj_id, tmp36, 3, 'Sensor experiments'),
        (obj_id, servo, 2, 'Motor experiments');

    -- Product 9: Wireless Sensor Node
    INSERT INTO products (name, description, version)
    VALUES ('Wireless Sensor Node', 'Battery-powered wireless sensor', 'v1.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, esp32, 1, 'Controller with WiFi'),
        (obj_id, dht22, 1, 'Environmental sensor');

    -- Product 10: Display Module
    INSERT INTO products (name, description, version)
    VALUES ('Universal Display Module', 'Reusable display module with OLED', 'v1.0')
    RETURNING id INTO obj_id;

    INSERT INTO product_parts (product_id, part_id, quantity, notes) VALUES
        (obj_id, arduino_nano, 1, 'Compact controller'),
        (obj_id, oled, 1, 'Display');
END $$;

-- ============================================================================
-- Product Inventory (Built instances)
-- ============================================================================
DO $$
DECLARE
    obj_id UUID;
    obj_inv_id UUID;

    -- Parts inventory IDs
    part_inv_ids UUID[];
BEGIN
    -- Build 5 Weather Stations (some fully tracked, some not)
    SELECT id INTO obj_id FROM products WHERE name = 'IoT Weather Station';

    -- Weather Station #1 - Fully tracked
    INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
    VALUES (obj_id, 'WS-001', false, 'In use - Production unit - fully tracked')
    RETURNING id INTO obj_inv_id;

    INSERT INTO product_inventory_parts (product_inventory_id, part_inventory_id)
    SELECT obj_inv_id, pi.id FROM parts_inventory pi
    WHERE pi.serial_number IN ('ESP32-001', 'DHT22-001', 'OLED-001');

    -- Weather Station #2 - Partially tracked
    INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
    VALUES (obj_id, 'WS-002', false, 'In use - Only tracking main controller')
    RETURNING id INTO obj_inv_id;

    INSERT INTO product_inventory_parts (product_inventory_id, part_inventory_id)
    SELECT obj_inv_id, pi.id FROM parts_inventory pi WHERE pi.serial_number = 'ESP32-002';

    -- Weather Station #3 - Not tracked
    INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
    VALUES (obj_id, 'WS-003', true, 'Available - parts not tracked');

    -- Weather Station #4 & #5
    INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
    VALUES
        (obj_id, 'WS-004', false, 'In use - Customer Site 1'),
        (obj_id, 'WS-005', false, 'In use - Customer Site 2');

    -- Build 3 Robot Cars
    SELECT id INTO obj_id FROM products WHERE name = 'Robot Car';

    INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
    VALUES (obj_id, 'ROBOT-001', false, 'In use - Engineering Lab - Prototype v1')
    RETURNING id INTO obj_inv_id;

    INSERT INTO product_inventory_parts (product_inventory_id, part_inventory_id)
    SELECT obj_inv_id, pi.id FROM parts_inventory pi
    WHERE pi.serial_number IN ('ARD-UNO-011', 'SERVO-001', 'SERVO-002');

    INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
    VALUES
        (obj_id, 'ROBOT-002', false, 'In use - Demo Unit'),
        (obj_id, 'ROBOT-003', true, 'Available for sale');

    -- Build 10 Arduino Starter Kits
    SELECT id INTO obj_id FROM products WHERE name = 'Arduino Starter Kit';

    FOR i IN 1..10 LOOP
        INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
        VALUES (
            obj_id,
            'ASK-' || LPAD(i::TEXT, 3, '0'),
            i <= 3,  -- First 3 are available
            CASE
                WHEN i <= 3 THEN 'Available in stock'
                WHEN i <= 7 THEN 'In use - Student ' || i
                ELSE 'In use - Classroom ' || (i - 7)
            END
        ) RETURNING id INTO obj_inv_id;

        -- Track some components for the first 3 kits
        IF i <= 3 THEN
            INSERT INTO product_inventory_parts (product_inventory_id, part_inventory_id)
            SELECT obj_inv_id, pi.id FROM parts_inventory pi
            WHERE pi.serial_number = 'ARD-UNO-' || LPAD(i::TEXT, 3, '0');
        END IF;
    END LOOP;

    -- Build Education Lab Kits
    SELECT id INTO obj_id FROM products WHERE name = 'Electronics Education Lab Kit';

    FOR i IN 1..3 LOOP
        INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
        VALUES (
            obj_id,
            'EDU-KIT-' || LPAD(i::TEXT, 3, '0'),
            false,  -- All in use
            'In use - School Lab ' || CHR(64 + i) || ' - Deployed to education'
        ) RETURNING id INTO obj_inv_id;

        -- Track Arduinos and breadboards
        INSERT INTO product_inventory_parts (product_inventory_id, part_inventory_id)
        SELECT obj_inv_id, pi.id FROM parts_inventory pi
        WHERE pi.serial_number IN ('BB-' || LPAD((13+i)::TEXT, 3, '0'));
    END LOOP;

    -- Build Temperature Loggers
    SELECT id INTO obj_id FROM products WHERE name = 'Multi-Zone Temperature Logger';

    FOR i IN 1..4 LOOP
        INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
        VALUES (
            obj_id,
            'TEMPLOG-' || LPAD(i::TEXT, 3, '0'),
            i > 2,  -- Last 2 are available
            CASE WHEN i <= 2 THEN 'In use - Data Center ' || i || ' - Monitoring active' ELSE 'Ready for deployment' END
        );
    END LOOP;

    -- Build Wireless Sensor Nodes
    SELECT id INTO obj_id FROM products WHERE name = 'Wireless Sensor Node';

    FOR i IN 1..8 LOOP
        INSERT INTO product_inventory (product_id, serial_number, is_available, notes)
        VALUES (
            obj_id,
            'WSN-' || LPAD(i::TEXT, 3, '0'),
            i > 5,  -- Last 3 are available
            CASE WHEN i <= 5 THEN 'In use - Building ' || CHR(64 + i) || ' Floor 1 - Deployed and active' ELSE 'Spare unit' END
        ) RETURNING id INTO obj_inv_id;

        -- Track ESP32 for deployed units
        IF i <= 5 THEN
            INSERT INTO product_inventory_parts (product_inventory_id, part_inventory_id)
            SELECT obj_inv_id, pi.id FROM parts_inventory pi
            WHERE pi.serial_number = 'ESP32-' || LPAD((20+i)::TEXT, 3, '0');
        END IF;
    END LOOP;
END $$;
