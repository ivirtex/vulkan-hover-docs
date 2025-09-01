# VkAntiLagModeAMD(3) Manual Page

## Name

VkAntiLagModeAMD - Set the status of the anti-lag feature



## [](#_c_specification)C Specification

Possible values of [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html)::`mode`, specifying the anti-lag status, are:

```c++
// Provided by VK_AMD_anti_lag
typedef enum VkAntiLagModeAMD {
    VK_ANTI_LAG_MODE_DRIVER_CONTROL_AMD = 0,
    VK_ANTI_LAG_MODE_ON_AMD = 1,
    VK_ANTI_LAG_MODE_OFF_AMD = 2,
} VkAntiLagModeAMD;
```

## [](#_description)Description

- `VK_ANTI_LAG_MODE_DRIVER_CONTROL_AMD` specifies that anti-lag will be enabled or disabled depending on driver settings.
- `VK_ANTI_LAG_MODE_ON_AMD` specifies that anti-lag will be enabled.
- `VK_ANTI_LAG_MODE_OFF_AMD` specifies that anti-lag will be disabled.

## [](#_see_also)See Also

[VK\_AMD\_anti\_lag](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_anti_lag.html), [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAntiLagModeAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0