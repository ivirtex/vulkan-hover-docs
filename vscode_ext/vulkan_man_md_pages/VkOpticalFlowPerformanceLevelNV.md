# VkOpticalFlowPerformanceLevelNV(3) Manual Page

## Name

VkOpticalFlowPerformanceLevelNV - Optical flow performance level types



## <a href="#_c_specification" class="anchor"></a>C Specification

Optical flow exposes performance levels which the user can choose based
on the desired performance and quality requirement.

The optical flow performance level types are defined with the following:

``` c
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowPerformanceLevelNV {
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_UNKNOWN_NV = 0,
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_SLOW_NV = 1,
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_MEDIUM_NV = 2,
    VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_FAST_NV = 3,
} VkOpticalFlowPerformanceLevelNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_SLOW_NV` is a level with slower
  performance but higher quality.

- `VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_MEDIUM_NV` is a level with medium
  performance and medium quality.

- `VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_FAST_NV` is a preset with higher
  performance but lower quality.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowPerformanceLevelNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
