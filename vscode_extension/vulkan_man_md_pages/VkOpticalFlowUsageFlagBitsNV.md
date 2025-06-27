# VkOpticalFlowUsageFlagBitsNV(3) Manual Page

## Name

VkOpticalFlowUsageFlagBitsNV - Bits specifying usage for optical flow
operations



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html)::`usage`,
controlling optical flow usage, are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV` specifies that the image **can**
  be used as input or reference frame for an optical flow operation.

- `VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV` specifies that the image **can**
  be used as output flow vector map for an optical flow operation.

- `VK_OPTICAL_FLOW_USAGE_HINT_BIT_NV` specifies that the image **can**
  be used as hint flow vector map for an optical flow operation.

- `VK_OPTICAL_FLOW_USAGE_COST_BIT_NV` specifies that the image **can**
  be used as output cost map for an optical flow operation.

- `VK_OPTICAL_FLOW_USAGE_GLOBAL_FLOW_BIT_NV` specifies that the image
  **can** be used as global flow vector for an optical flow operation.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowUsageFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowUsageFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowUsageFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
