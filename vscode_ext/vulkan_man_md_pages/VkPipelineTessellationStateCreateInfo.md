# VkPipelineTessellationStateCreateInfo(3) Manual Page

## Name

VkPipelineTessellationStateCreateInfo - Structure specifying parameters
of a newly created pipeline tessellation state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineTessellationStateCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineTessellationStateCreateInfo {
    VkStructureType                           sType;
    const void*                               pNext;
    VkPipelineTessellationStateCreateFlags    flags;
    uint32_t                                  patchControlPoints;
} VkPipelineTessellationStateCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `patchControlPoints` is the number of control points per patch.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPipelineTessellationStateCreateInfo-patchControlPoints-01214"
  id="VUID-VkPipelineTessellationStateCreateInfo-patchControlPoints-01214"></a>
  VUID-VkPipelineTessellationStateCreateInfo-patchControlPoints-01214  
  `patchControlPoints` **must** be greater than zero and less than or
  equal to `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineTessellationStateCreateInfo-sType-sType"
  id="VUID-VkPipelineTessellationStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineTessellationStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineTessellationStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineTessellationStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineTessellationStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html)

- <a href="#VUID-VkPipelineTessellationStateCreateInfo-sType-unique"
  id="VUID-VkPipelineTessellationStateCreateInfo-sType-unique"></a>
  VUID-VkPipelineTessellationStateCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPipelineTessellationStateCreateInfo-flags-zerobitmask"
  id="VUID-VkPipelineTessellationStateCreateInfo-flags-zerobitmask"></a>
  VUID-VkPipelineTessellationStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsShaderGroupCreateInfoNV.html),
[VkPipelineTessellationStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineTessellationStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
