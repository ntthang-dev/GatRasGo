### **1. Package `modbus_inverter`**  
**Mục đích**: Quản lý kết nối Modbus TCP đến các inverter, đọc/ghi dữ liệu theo bảng địa chỉ quy định.

#### **Cấu trúc thư mục**:
```
modbus_inverter/
├── client.go          // Core Modbus client
├── config.go          // Cấu hình kết nối
├── errors.go          // Custom errors
└── types.go           // Data structures
```

### **2. Package `iec104_evn`**  
**Mục đích**: Triển khai IEC 60870-5-104 Server để nhận lệnh từ EVN và phản hồi dữ liệu.

#### **Cấu trúc thư mục**:
```
iec104_evn/
├── server.go          // IEC 104 Server core
├── config.go          // Cấu hình kết nối
├── commands.go        // Xử lý lệnh từ EVN
└── types.go           // Data structures
```

### **Chức năng package**:
1. **modbus_inverter**:
   - Đóng gói kết nối Modbus TCP, xử lý chuyển đổi dữ liệu (float32 ↔ bytes).
   - Hỗ trợ retry khi mất kết nối.
   - Tách biệt config để dễ quản lý nhiều inverter.

2. **iec104_evn**:
   - Triển khai IEC 104 Server dùng thư viện `go-iec104`.
   - Xử lý lệnh từ EVN (ví dụ: T45) và trigger action đến inverter.
   - Tự động gửi dữ liệu giám sát (T13) theo chu kỳ.

3. **Tích hợp**:
   - Gateway đóng vai trò **Modbus Client** (đọc inverter) và **IEC 104 Server** (phục vụ EVN).
   - Dữ liệu được đồng bộ hai chiều: EVN điều khiển → inverter, inverter → EVN.

---