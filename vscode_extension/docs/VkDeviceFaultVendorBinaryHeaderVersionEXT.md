# VkDeviceFaultVendorBinaryHeaderVersionEXT(3) Manual Page

## Name

VkDeviceFaultVendorBinaryHeaderVersionEXT - Encode vendor binary crash dump version



## [](#_c_specification)C Specification

Possible values of the `headerVersion` value of the crash dump header are:

```c++
// Provided by VK_EXT_device_fault
typedef enum VkDeviceFaultVendorBinaryHeaderVersionEXT {
    VK_DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT = 1,
} VkDeviceFaultVendorBinaryHeaderVersionEXT;
```

## [](#_description)Description

- `VK_DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT` specifies version one of the binary crash dump header.

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceFaultVendorBinaryHeaderVersionOneEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorBinaryHeaderVersionOneEXT.html), [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultVendorBinaryHeaderVersionEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0