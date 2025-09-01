# VkDeviceFaultVendorInfoEXT(3) Manual Page

## Name

VkDeviceFaultVendorInfoEXT - Structure specifying vendor-specific fault information



## [](#_c_specification)C Specification

The `VkDeviceFaultVendorInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultVendorInfoEXT {
    char        description[VK_MAX_DESCRIPTION_SIZE];
    uint64_t    vendorFaultCode;
    uint64_t    vendorFaultData;
} VkDeviceFaultVendorInfoEXT;
```

## [](#_members)Members

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a null-terminated UTF-8 string which is a human readable description of the fault.
- `vendorFaultCode` is the vendor-specific fault code for this fault.
- `vendorFaultData` is the vendor-specific fault data associated with this fault.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultVendorInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0