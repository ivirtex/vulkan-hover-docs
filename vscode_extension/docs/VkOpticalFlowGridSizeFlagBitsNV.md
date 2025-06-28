# VkOpticalFlowGridSizeFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowGridSizeFlagBitsNV - Bits specifying grid sizes for optical flow operations



## [](#_c_specification)C Specification

Optical flow vectors are generated block-wise, one vector for each block of NxN pixels (referred to as grid).

Bits which **can** be set in [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html)::`outputGridSize` and [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html)::`hintGridSize`, or which are returned in [VkPhysicalDeviceOpticalFlowPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpticalFlowPropertiesNV.html)::`supportedOutputGridSizes` and [VkPhysicalDeviceOpticalFlowPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpticalFlowPropertiesNV.html)::`supportedHintGridSizes` controlling optical flow grid sizes, are:

```c++
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowGridSizeFlagBitsNV {
    VK_OPTICAL_FLOW_GRID_SIZE_UNKNOWN_NV = 0,
    VK_OPTICAL_FLOW_GRID_SIZE_1X1_BIT_NV = 0x00000001,
    VK_OPTICAL_FLOW_GRID_SIZE_2X2_BIT_NV = 0x00000002,
    VK_OPTICAL_FLOW_GRID_SIZE_4X4_BIT_NV = 0x00000004,
    VK_OPTICAL_FLOW_GRID_SIZE_8X8_BIT_NV = 0x00000008,
} VkOpticalFlowGridSizeFlagBitsNV;
```

## [](#_description)Description

- `VK_OPTICAL_FLOW_GRID_SIZE_1X1_BIT_NV` specifies that grid is 1x1 pixel.
- `VK_OPTICAL_FLOW_GRID_SIZE_2X2_BIT_NV` specifies that grid is 2x2 pixel.
- `VK_OPTICAL_FLOW_GRID_SIZE_4X4_BIT_NV` specifies that grid is 4x4 pixel.
- `VK_OPTICAL_FLOW_GRID_SIZE_8X8_BIT_NV` specifies that grid is 8x8 pixel.

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowGridSizeFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0