# VkOpticalFlowPerformanceLevelNV(3) Manual Page

## Name

VkOpticalFlowPerformanceLevelNV - Optical flow performance level types



## [](#_c_specification)C Specification

Optical flow exposes performance levels which the application **can** choose based on the desired performance and quality requirement.

The optical flow performance level types are defined with the following:

```c++
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowPerformanceLevelNV {
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_UNKNOWN_NV = 0,
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_SLOW_NV = 1,
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_MEDIUM_NV = 2,
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_FAST_NV = 3,
} VkOpticalFlowPerformanceLevelNV;
```

## [](#_description)Description

- `VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_SLOW_NV` is a level with slower performance but higher quality.
- `VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_MEDIUM_NV` is a level with medium performance and medium quality.
- `VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_FAST_NV` is a preset with higher performance but lower quality.

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowPerformanceLevelNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0