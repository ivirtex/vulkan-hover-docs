# VkOpticalFlowExecuteInfoNV(3) Manual Page

## Name

VkOpticalFlowExecuteInfoNV - Structure specifying parameters of an
optical flow vector calculation



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteInfoNV.html)
structure is defined as:

``` c
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowExecuteInfoNV {
    VkStructureType                sType;
    void*                          pNext;
    VkOpticalFlowExecuteFlagsNV    flags;
    uint32_t                       regionCount;
    const VkRect2D*                pRegions;
} VkOpticalFlowExecuteInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` are the
  [VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteFlagsNV.html) used
  for this command.

- `regionCount` is the number of regions of interest specified in
  `pRegions`.

- `pRegions` is a pointer to `regionCount` `VkRect2D` regions of
  interest.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkOpticalFlowExecuteInfoNV-regionCount-07593"
  id="VUID-VkOpticalFlowExecuteInfoNV-regionCount-07593"></a>
  VUID-VkOpticalFlowExecuteInfoNV-regionCount-07593  
  `regionCount` **must** be 0 if
  `VK_OPTICAL_FLOW_SESSION_CREATE_ALLOW_REGIONS_BIT_NV` was not set for
  `VkOpticalFlowSessionNV` on which this command is operating

Valid Usage (Implicit)

- <a href="#VUID-VkOpticalFlowExecuteInfoNV-sType-sType"
  id="VUID-VkOpticalFlowExecuteInfoNV-sType-sType"></a>
  VUID-VkOpticalFlowExecuteInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OPTICAL_FLOW_EXECUTE_INFO_NV`

- <a href="#VUID-VkOpticalFlowExecuteInfoNV-pNext-pNext"
  id="VUID-VkOpticalFlowExecuteInfoNV-pNext-pNext"></a>
  VUID-VkOpticalFlowExecuteInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkOpticalFlowExecuteInfoNV-flags-parameter"
  id="VUID-VkOpticalFlowExecuteInfoNV-flags-parameter"></a>
  VUID-VkOpticalFlowExecuteInfoNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkOpticalFlowExecuteFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteFlagBitsNV.html)
  values

- <a href="#VUID-VkOpticalFlowExecuteInfoNV-pRegions-parameter"
  id="VUID-VkOpticalFlowExecuteInfoNV-pRegions-parameter"></a>
  VUID-VkOpticalFlowExecuteInfoNV-pRegions-parameter  
  If `regionCount` is not `0`, `pRegions` **must** be a valid pointer to
  an array of `regionCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteFlagsNV.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdOpticalFlowExecuteNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowExecuteInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
