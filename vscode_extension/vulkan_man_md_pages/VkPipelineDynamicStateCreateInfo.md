# VkPipelineDynamicStateCreateInfo(3) Manual Page

## Name

VkPipelineDynamicStateCreateInfo - Structure specifying parameters of a
newly created pipeline dynamic state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineDynamicStateCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineDynamicStateCreateInfo {
    VkStructureType                      sType;
    const void*                          pNext;
    VkPipelineDynamicStateCreateFlags    flags;
    uint32_t                             dynamicStateCount;
    const VkDynamicState*                pDynamicStates;
} VkPipelineDynamicStateCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `dynamicStateCount` is the number of elements in the `pDynamicStates`
  array.

- `pDynamicStates` is a pointer to an array of
  [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html) values specifying which pieces
  of pipeline state will use the values from dynamic state commands
  rather than from pipeline state creation information.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-01442"
  id="VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-01442"></a>
  VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-01442  
  Each element of `pDynamicStates` **must** be unique

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineDynamicStateCreateInfo-sType-sType"
  id="VUID-VkPipelineDynamicStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineDynamicStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineDynamicStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineDynamicStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineDynamicStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPipelineDynamicStateCreateInfo-flags-zerobitmask"
  id="VUID-VkPipelineDynamicStateCreateInfo-flags-zerobitmask"></a>
  VUID-VkPipelineDynamicStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-parameter"
  id="VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-parameter"></a>
  VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-parameter  
  If `dynamicStateCount` is not `0`, `pDynamicStates` **must** be a
  valid pointer to an array of `dynamicStateCount` valid
  [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineDynamicStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateFlags.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineDynamicStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
