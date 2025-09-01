# VkDeviceFaultVendorBinaryHeaderVersionOneEXT(3) Manual Page

## Name

VkDeviceFaultVendorBinaryHeaderVersionOneEXT - Structure describing the layout of the vendor binary crash dump header



## [](#_c_specification)C Specification

Version one of the crash dump header is defined as:

```c++
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultVendorBinaryHeaderVersionOneEXT {
    uint32_t                                     headerSize;
    VkDeviceFaultVendorBinaryHeaderVersionEXT    headerVersion;
    uint32_t                                     vendorID;
    uint32_t                                     deviceID;
    uint32_t                                     driverVersion;
    uint8_t                                      pipelineCacheUUID[VK_UUID_SIZE];
    uint32_t                                     applicationNameOffset;
    uint32_t                                     applicationVersion;
    uint32_t                                     engineNameOffset;
    uint32_t                                     engineVersion;
    uint32_t                                     apiVersion;
} VkDeviceFaultVendorBinaryHeaderVersionOneEXT;
```

## [](#_members)Members

- `headerSize` is the length in bytes of the crash dump header.
- `headerVersion` is a [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html) enum value specifying the version of the header. A consumer of the crash dump **should** use the header version to interpret the remainder of the header. `headerVersion` **must** be written as exactly 4 bytes.
- `vendorID` is the `VkPhysicalDeviceProperties`::`vendorID` of the implementation.
- `deviceID` is the `VkPhysicalDeviceProperties`::`deviceID` of the implementation.
- `driverVersion` is the `VkPhysicalDeviceProperties`::`driverVersion` of the implementation.
- `pipelineCacheUUID` is an array of `VK_UUID_SIZE` `uint8_t` values matching the `VkPhysicalDeviceProperties`::`pipelineCacheUUID` property of the implementation.
- `applicationNameOffset` is zero, or an offset from the base address of the crash dump header to a null-terminated UTF-8 string containing the name of the application. If `applicationNameOffset` is non-zero, this string **must** match the application name specified via [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`pApplicationName` during instance creation.
- `applicationVersion` **must** be zero or the value specified by [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`applicationVersion` during instance creation.
- `engineNameOffset` is zero, or an offset from the base address of the crash dump header to a null-terminated UTF-8 string containing the name of the engine (if any) used to create the application. If `engineNameOffset` is non-zero, this string **must** match the engine name specified via [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`pEngineName` during instance creation.
- `engineVersion` **must** be zero or the value specified by [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`engineVersion` during instance creation.
- `apiVersion` **must** be zero or the value specified by [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` during instance creation.

## [](#_description)Description

Unlike most structures declared by the Vulkan API, all fields of this structure are written with the least significant byte first, regardless of host byte-order.

The C language specification does not define the packing of structure members. This layout assumes tight structure member packing, with members laid out in the order listed in the structure, and the intended size of the structure is 56 bytes. If a compiler produces code that diverges from that pattern, applications **must** employ another method to set values at the correct offsets.

Valid Usage

- [](#VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerSize-07340)VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerSize-07340  
  `headerSize` **must** be 56
- [](#VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-07341)VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-07341  
  `headerVersion` **must** be `VK_DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT`

Valid Usage (Implicit)

- [](#VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-parameter)VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-parameter  
  `headerVersion` **must** be a valid [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultVendorBinaryHeaderVersionOneEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0