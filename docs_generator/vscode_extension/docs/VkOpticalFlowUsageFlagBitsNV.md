# VkOpticalFlowUsageFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowUsageFlagBitsNV - Bits specifying usage for optical flow operations



## [](#_c_specification)C Specification

Bits which **can** be set in [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html)::`usage`, controlling optical flow usage, are:

```c++
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowUsageFlagBitsNV {
    VK_OPTICAL_FLOW_USAGE_UNKNOWN_NV = 0,
    VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV = 0x00000001,
    VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV = 0x00000002,
    VK_OPTICAL_FLOW_USAGE_HINT_BIT_NV = 0x00000004,
    VK_OPTICAL_FLOW_USAGE_COST_BIT_NV = 0x00000008,
    VK_OPTICAL_FLOW_USAGE_GLOBAL_FLOW_BIT_NV = 0x00000010,
} VkOpticalFlowUsageFlagBitsNV;
```

## [](#_description)Description

- `VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV` specifies that the image **can** be used as input or reference frame for an optical flow operation.
- `VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV` specifies that the image **can** be used as output flow vector map for an optical flow operation.
- `VK_OPTICAL_FLOW_USAGE_HINT_BIT_NV` specifies that the image **can** be used as hint flow vector map for an optical flow operation.
- `VK_OPTICAL_FLOW_USAGE_COST_BIT_NV` specifies that the image **can** be used as output cost map for an optical flow operation.
- `VK_OPTICAL_FLOW_USAGE_GLOBAL_FLOW_BIT_NV` specifies that the image **can** be used as global flow vector for an optical flow operation.

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowUsageFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowUsageFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0