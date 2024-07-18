# VkDeviceFaultVendorBinaryHeaderVersionOneEXT(3) Manual Page

## Name

VkDeviceFaultVendorBinaryHeaderVersionOneEXT - Structure describing the
layout of the vendor binary crash dump header



## <a href="#_c_specification" class="anchor"></a>C Specification

Version one of the crash dump header is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `headerSize` is the length in bytes of the crash dump header.

- `headerVersion` is a
  [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html)
  enum value specifying the version of the header. A consumer of the
  crash dump **should** use the header version to interpret the
  remainder of the header.

- `vendorID` is the `VkPhysicalDeviceProperties`::`vendorID` of the
  implementation.

- `deviceID` is the `VkPhysicalDeviceProperties`::`deviceID` of the
  implementation.

- `driverVersion` is the `VkPhysicalDeviceProperties`::`driverVersion`
  of the implementation.

- `pipelineCacheUUID` is an array of `VK_UUID_SIZE` `uint8_t` values
  matching the `VkPhysicalDeviceProperties`::`pipelineCacheUUID`
  property of the implementation.

- `applicationNameOffset` is zero, or an offset from the base address of
  the crash dump header to a null-terminated UTF-8 string containing the
  name of the application. If `applicationNameOffset` is non-zero, this
  string **must** match the application name specified via
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`pApplicationName` during
  instance creation.

- `applicationVersion` **must** be zero or the value specified by
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`applicationVersion`
  during instance creation.

- `engineNameOffset` is zero, or an offset from the base address of the
  crash dump header to a null-terminated UTF-8 string containing the
  name of the engine (if any) used to create the application. If
  `engineNameOffset` is non-zero, this string **must** match the engine
  name specified via
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`pEngineName` during
  instance creation.

- `engineVersion` **must** be zero or the value specified by
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`engineVersion` during
  instance creation.

- `apiVersion` **must** be zero or the value specified by
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` during
  instance creation.

## <a href="#_description" class="anchor"></a>Description

Unlike most structures declared by the Vulkan API, all fields of this
structure are written with the least significant byte first, regardless
of host byte-order.

The C language specification does not define the packing of structure
members. This layout assumes tight structure member packing, with
members laid out in the order listed in the structure, and the intended
size of the structure is 56 bytes. If a compiler produces code that
diverges from that pattern, applications **must** employ another method
to set values at the correct offsets.

Valid Usage

- <a
  href="#VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerSize-07340"
  id="VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerSize-07340"></a>
  VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerSize-07340  
  `headerSize` **must** be 56

- <a
  href="#VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-07341"
  id="VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-07341"></a>
  VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-07341  
  `headerVersion` **must** be
  `VK_DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT`

Valid Usage (Implicit)

- <a
  href="#VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-parameter"
  id="VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-parameter"></a>
  VUID-VkDeviceFaultVendorBinaryHeaderVersionOneEXT-headerVersion-parameter  
  `headerVersion` **must** be a valid
  [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultVendorBinaryHeaderVersionOneEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
