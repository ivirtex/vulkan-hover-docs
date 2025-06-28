# VkOpticalFlowExecuteFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowExecuteFlagBitsNV - Bits specifying flags for an optical flow vector calculation



## [](#_c_specification)C Specification

Bits which **can** be set in [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html)::`flags`, controlling optical flow execution, are:

```c++
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowExecuteFlagBitsNV {
    VK_OPTICAL_FLOW_EXECUTE_DISABLE_TEMPORAL_HINTS_BIT_NV = 0x00000001,
} VkOpticalFlowExecuteFlagBitsNV;
```

## [](#_description)Description

- `VK_OPTICAL_FLOW_EXECUTE_DISABLE_TEMPORAL_HINTS_BIT_NV` specifies that temporal hints from previously generated flow vectors are not used. If temporal hints are enabled, optical flow vectors from previous [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdOpticalFlowExecuteNV.html) call are automatically used as hints for the current [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdOpticalFlowExecuteNV.html) call, to take advantage of temporal correlation in a video sequence. Temporal hints should be disabled if there is a-priori knowledge of no temporal correlation (e.g. a scene change, independent successive frame pairs).

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowExecuteFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0