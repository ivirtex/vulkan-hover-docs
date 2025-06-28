# VkDeviceDiagnosticsConfigCreateInfoNV(3) Manual Page

## Name

VkDeviceDiagnosticsConfigCreateInfoNV - Specify diagnostics config for a Vulkan device



## [](#_c_specification)C Specification

When using the Nsight™ Aftermath SDK, to configure how device crash dumps are created, add a [VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html) structure to the `pNext` chain of the [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) structure.

```c++
// Provided by VK_NV_device_diagnostics_config
typedef struct VkDeviceDiagnosticsConfigCreateInfoNV {
    VkStructureType                     sType;
    const void*                         pNext;
    VkDeviceDiagnosticsConfigFlagsNV    flags;
} VkDeviceDiagnosticsConfigCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkDeviceDiagnosticsConfigFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigFlagBitsNV.html) specifying additional parameters for configuring diagnostic tools.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDeviceDiagnosticsConfigCreateInfoNV-sType-sType)VUID-VkDeviceDiagnosticsConfigCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_DIAGNOSTICS_CONFIG_CREATE_INFO_NV`
- [](#VUID-VkDeviceDiagnosticsConfigCreateInfoNV-flags-parameter)VUID-VkDeviceDiagnosticsConfigCreateInfoNV-flags-parameter  
  `flags` **must** be a valid combination of [VkDeviceDiagnosticsConfigFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_device\_diagnostics\_config](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostics_config.html), [VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceDiagnosticsConfigCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0