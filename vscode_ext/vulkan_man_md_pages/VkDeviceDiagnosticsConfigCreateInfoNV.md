# VkDeviceDiagnosticsConfigCreateInfoNV(3) Manual Page

## Name

VkDeviceDiagnosticsConfigCreateInfoNV - Specify diagnostics config for a
Vulkan device



## <a href="#_c_specification" class="anchor"></a>C Specification

When using the Nsight<sup>™</sup> Aftermath SDK, to configure how device
crash dumps are created, add a
[VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html)
structure to the `pNext` chain of the
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure.

``` c
// Provided by VK_NV_device_diagnostics_config
typedef struct VkDeviceDiagnosticsConfigCreateInfoNV {
    VkStructureType                     sType;
    const void*                         pNext;
    VkDeviceDiagnosticsConfigFlagsNV    flags;
} VkDeviceDiagnosticsConfigCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkDeviceDiagnosticsConfigFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagBitsNV.html)
  specifying additional parameters for configuring diagnostic tools.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceDiagnosticsConfigCreateInfoNV-sType-sType"
  id="VUID-VkDeviceDiagnosticsConfigCreateInfoNV-sType-sType"></a>
  VUID-VkDeviceDiagnosticsConfigCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_DIAGNOSTICS_CONFIG_CREATE_INFO_NV`

- <a href="#VUID-VkDeviceDiagnosticsConfigCreateInfoNV-flags-parameter"
  id="VUID-VkDeviceDiagnosticsConfigCreateInfoNV-flags-parameter"></a>
  VUID-VkDeviceDiagnosticsConfigCreateInfoNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDeviceDiagnosticsConfigFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagBitsNV.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_diagnostics_config](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostics_config.html),
[VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceDiagnosticsConfigCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
