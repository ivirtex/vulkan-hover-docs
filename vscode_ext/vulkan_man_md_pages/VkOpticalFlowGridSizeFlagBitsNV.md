# VkOpticalFlowGridSizeFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowGridSizeFlagBitsNV - Bits specifying grid sizes for optical
flow operations



## <a href="#_c_specification" class="anchor"></a>C Specification

Optical flow vectors are generated block-wise, one vector for each block
of NxN pixels (referred to as grid).

Bits which **can** be set in
[VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html)::`outputGridSize`
and
[VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html)::`hintGridSize`,
or which are returned in
[VkPhysicalDeviceOpticalFlowPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpticalFlowPropertiesNV.html)::`supportedOutputGridSizes`
and
[VkPhysicalDeviceOpticalFlowPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpticalFlowPropertiesNV.html)::`supportedHintGridSizes`
controlling optical flow grid sizes, are:

``` c
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowGridSizeFlagBitsNV {
    VK_OPTICAL_FLOW_GRID_SIZE_UNKNOWN_NV = 0,
    VK_OPTICAL_FLOW_GRID_SIZE_1X1_BIT_NV = 0x00000001,
    VK_OPTICAL_FLOW_GRID_SIZE_2X2_BIT_NV = 0x00000002,
    VK_OPTICAL_FLOW_GRID_SIZE_4X4_BIT_NV = 0x00000004,
    VK_OPTICAL_FLOW_GRID_SIZE_8X8_BIT_NV = 0x00000008,
} VkOpticalFlowGridSizeFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_OPTICAL_FLOW_GRID_SIZE_1X1_BIT_NV` specifies that grid is 1x1
  pixel.

- `VK_OPTICAL_FLOW_GRID_SIZE_2X2_BIT_NV` specifies that grid is 2x2
  pixel.

- `VK_OPTICAL_FLOW_GRID_SIZE_4X4_BIT_NV` specifies that grid is 4x4
  pixel.

- `VK_OPTICAL_FLOW_GRID_SIZE_8X8_BIT_NV` specifies that grid is 8x8
  pixel.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowGridSizeFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
