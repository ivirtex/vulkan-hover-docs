# VkOpticalFlowExecuteFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowExecuteFlagBitsNV - Bits specifying flags for a optical
flow vector calculation



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteInfoNV.html)::`flags`,
controlling optical flow execution, are:

``` c
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowExecuteFlagBitsNV {
    VK_OPTICAL_FLOW_EXECUTE_DISABLE_TEMPORAL_HINTS_BIT_NV = 0x00000001,
} VkOpticalFlowExecuteFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_OPTICAL_FLOW_EXECUTE_DISABLE_TEMPORAL_HINTS_BIT_NV` specifies that
  temporal hints from previously generated flow vectors are not used. If
  temporal hints are enabled, optical flow vectors from previous
  [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdOpticalFlowExecuteNV.html) call are
  automatically used as hints for the current
  [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdOpticalFlowExecuteNV.html) call, to
  take advantage of temporal correlation in a video sequence. Temporal
  hints should be disabled if there is a-priori knowledge of no temporal
  correlation (e.g. a scene change, independent successive frame pairs).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowExecuteFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
