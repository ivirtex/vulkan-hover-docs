# VkOpticalFlowExecuteInfoNV(3) Manual Page

## Name

VkOpticalFlowExecuteInfoNV - Structure specifying parameters of an optical flow vector calculation



## [](#_c_specification)C Specification

The [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowExecuteInfoNV {
    VkStructureType                sType;
    void*                          pNext;
    VkOpticalFlowExecuteFlagsNV    flags;
    uint32_t                       regionCount;
    const VkRect2D*                pRegions;
} VkOpticalFlowExecuteInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` are the [VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteFlagsNV.html) used for this command.
- `regionCount` is the number of regions of interest specified in `pRegions`.
- `pRegions` is a pointer to `regionCount` `VkRect2D` regions of interest.

## [](#_description)Description

Valid Usage

- [](#VUID-VkOpticalFlowExecuteInfoNV-regionCount-07593)VUID-VkOpticalFlowExecuteInfoNV-regionCount-07593  
  `regionCount` **must** be 0 if `VK_OPTICAL_FLOW_SESSION_CREATE_ALLOW_REGIONS_BIT_NV` was not set for `VkOpticalFlowSessionNV` on which this command is operating

Valid Usage (Implicit)

- [](#VUID-VkOpticalFlowExecuteInfoNV-sType-sType)VUID-VkOpticalFlowExecuteInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OPTICAL_FLOW_EXECUTE_INFO_NV`
- [](#VUID-VkOpticalFlowExecuteInfoNV-pNext-pNext)VUID-VkOpticalFlowExecuteInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkOpticalFlowExecuteInfoNV-flags-parameter)VUID-VkOpticalFlowExecuteInfoNV-flags-parameter  
  `flags` **must** be a valid combination of [VkOpticalFlowExecuteFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteFlagBitsNV.html) values
- [](#VUID-VkOpticalFlowExecuteInfoNV-pRegions-parameter)VUID-VkOpticalFlowExecuteInfoNV-pRegions-parameter  
  If `regionCount` is not `0`, `pRegions` **must** be a valid pointer to an array of `regionCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteFlagsNV.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdOpticalFlowExecuteNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowExecuteInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0