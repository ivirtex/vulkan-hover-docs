# VkDeviceDiagnosticsConfigFlagBitsNV(3) Manual Page

## Name

VkDeviceDiagnosticsConfigFlagBitsNV - Bitmask specifying diagnostics
flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html)::`flags`
include:

``` c
// Provided by VK_NV_device_diagnostics_config
typedef enum VkDeviceDiagnosticsConfigFlagBitsNV {
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV = 0x00000001,
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV = 0x00000002,
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV = 0x00000004,
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV = 0x00000008,
} VkDeviceDiagnosticsConfigFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV` enables
  the generation of debug information for shaders.

- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV` enables
  driver side tracking of resources (images, buffers, etc.) used to
  augment the device fault information.

- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV`
  enables automatic insertion of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-diagnostic-checkpoints"
  target="_blank" rel="noopener">diagnostic checkpoints</a> for draw
  calls, dispatches, trace rays, and copies. The CPU call stack at the
  time of the command will be associated as the marker data for the
  automatically inserted checkpoints.

- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV`
  enables shader error reporting.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_diagnostics_config](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostics_config.html),
[VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceDiagnosticsConfigFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
