# VkDeviceDiagnosticsConfigFlagBitsNV(3) Manual Page

## Name

VkDeviceDiagnosticsConfigFlagBitsNV - Bitmask specifying diagnostics flags



## [](#_c_specification)C Specification

Bits which **can** be set in [VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html)::`flags` include:

```c++
// Provided by VK_NV_device_diagnostics_config
typedef enum VkDeviceDiagnosticsConfigFlagBitsNV {
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV = 0x00000001,
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV = 0x00000002,
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV = 0x00000004,
    VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV = 0x00000008,
} VkDeviceDiagnosticsConfigFlagBitsNV;
```

## [](#_description)Description

- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV` enables the generation of debug information for shaders.
- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV` enables driver side tracking of resources (images, buffers, etc.) used to augment the device fault information.
- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV` enables automatic insertion of [diagnostic checkpoints](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#device-diagnostic-checkpoints) for draw calls, dispatches, trace rays, and copies. The CPU call stack at the time of the command will be associated as the marker data for the automatically inserted checkpoints.
- `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV` enables shader error reporting.

## [](#_see_also)See Also

[VK\_NV\_device\_diagnostics\_config](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostics_config.html), [VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceDiagnosticsConfigFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0