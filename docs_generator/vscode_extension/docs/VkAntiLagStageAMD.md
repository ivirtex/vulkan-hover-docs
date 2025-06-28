# VkAntiLagStageAMD(3) Manual Page

## Name

VkAntiLagStageAMD - Report the application stage



## [](#_c_specification)C Specification

Possible values of [VkAntiLagPresentationInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagPresentationInfoAMD.html)::`stage`, specifying the current application stage, are:

```c++
// Provided by VK_AMD_anti_lag
typedef enum VkAntiLagStageAMD {
    VK_ANTI_LAG_STAGE_INPUT_AMD = 0,
    VK_ANTI_LAG_STAGE_PRESENT_AMD = 1,
} VkAntiLagStageAMD;
```

## [](#_description)Description

- `VK_ANTI_LAG_STAGE_INPUT_AMD` specifies the stage before processing input.
- `VK_ANTI_LAG_STAGE_PRESENT_AMD` specifies the stage before [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html).

## [](#_see_also)See Also

[VK\_AMD\_anti\_lag](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_anti_lag.html), [VkAntiLagPresentationInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagPresentationInfoAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAntiLagStageAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0