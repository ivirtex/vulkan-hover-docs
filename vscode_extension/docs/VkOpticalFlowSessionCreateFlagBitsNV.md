# VkOpticalFlowSessionCreateFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowSessionCreateFlagBitsNV - Bits specifying flags for optical flow session



## [](#_c_specification)C Specification

Bits which **can** be set in [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html)::`flags`, controlling optical flow session operations, are:

```c++
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowSessionCreateFlagBitsNV {
    VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_HINT_BIT_NV = 0x00000001,
    VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_COST_BIT_NV = 0x00000002,
    VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_GLOBAL_FLOW_BIT_NV = 0x00000004,
    VK_OPTICAL_FLOW_SESSION_CREATE_ALLOW_REGIONS_BIT_NV = 0x00000008,
    VK_OPTICAL_FLOW_SESSION_CREATE_BOTH_DIRECTIONS_BIT_NV = 0x00000010,
} VkOpticalFlowSessionCreateFlagBitsNV;
```

## [](#_description)Description

- `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_HINT_BIT_NV` specifies that a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with external flow vectors will be used as hints in performing the motion search and **must** be bound to `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_HINT_NV`.
- `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_COST_BIT_NV` specifies that the cost for the forward flow is generated in a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) which **must** be bound to `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_COST_NV`. Additionally, if `VK_OPTICAL_FLOW_SESSION_CREATE_BOTH_DIRECTIONS_BIT_NV` is also set, the cost for backward flow is generated in a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) which **must** be bound to `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_BACKWARD_COST_NV`. The cost is the confidence level of the flow vector for each grid in the frame. The Cost implies how (in)accurate the flow vector is. Higher cost value implies the flow vector to be less accurate and vice-versa.
- `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_GLOBAL_FLOW_BIT_NV` specifies that a global flow vector is estimated from forward flow in a single pixel [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) which **must** be bound to `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_GLOBAL_FLOW_NV`.
- `VK_OPTICAL_FLOW_SESSION_CREATE_ALLOW_REGIONS_BIT_NV` specifies that regions of interest **can** be specified in [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html).
- `VK_OPTICAL_FLOW_SESSION_CREATE_BOTH_DIRECTIONS_BIT_NV` specifies that backward flow is generated in addition to forward flow which is always generated.

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowSessionCreateFlagBitsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0